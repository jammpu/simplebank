package api

import (
	"github.com/gin-gonic/gin"
	mockdb "github.com/jammpu/simplebank/db/mock"
	db "github.com/jammpu/simplebank/db/sqlc"
	"github.com/jammpu/simplebank/token"
	"github.com/jammpu/simplebank/util"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

type testCases struct {
	name          string
	accountID     int64
	setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
	body          gin.H
	buildStubs    func(store *mockdb.MockStore)
	checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
}

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymetricKey:    util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}
	server, err := NewServer(config, store)
	require.NoError(t, err)
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
