package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jammpu/simplebank/db/sqlc"
	"github.com/jammpu/simplebank/token"
	"github.com/lib/pq"
	"net/http"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (s *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Balance:  0,
		Currency: req.Currency,
	}

	account, err := s.store.CreateAccount(ctx, arg)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := s.store.GetAccount(ctx, req.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if account.Owner != authPayload.Username {
		err = errors.New("not authorized to access this account")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
}

func (s *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := s.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

//type updateAccountURIRequest struct {
//	ID int64 `uri:"id" binding:"required,min=1"`
//}
//
//type updateAccountJSONRequest struct {
//	Amount int64 `json:"amount" binding:"required,min=1"`
//}
//
//func (s *Server) updateAccount(ctx *gin.Context) {
//	var reqUri updateAccountURIRequest
//	var reqJSON updateAccountJSONRequest
//
//	if err := ctx.ShouldBindUri(&reqUri); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	if err := ctx.ShouldBindJSON(&reqJSON); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	arg := db.UpdateAccountBalanceParams{
//		ID:     reqUri.ID,
//		Amount: reqJSON.Amount,
//	}
//
//	account, err := s.store.UpdateAccountBalance(ctx, arg)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, err)
//		return
//	}
//	ctx.JSON(http.StatusOK, account)
//}
//
//type deleteAccountRequest struct {
//	ID int64 `uri:"id" binding:"required,min=1"`
//}
//
//func (s *Server) deleteAccount(ctx *gin.Context) {
//	var req deleteAccountRequest
//	if err := ctx.ShouldBindUri(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, errorResponse(err))
//		return
//	}
//
//	err := s.store.DeleteAccountCheckRows(ctx, req.ID)
//	if err != nil {
//		if errors.Is(db.ErrNotFound, err) {
//			ctx.JSON(http.StatusNotFound, errorResponse(err))
//			return
//		}
//		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"account deleted with ID": req.ID})
//}
