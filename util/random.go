package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt genera un entero aleatorio entre min y max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString genera un string aleatorio de longitud n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner genera un dueño aleatorio
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney genera dinero aleatorio
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency genera tipo de modena aleatorio
func RandomCurrency() string {
	currencies := []string{EUR, USD}
	return currencies[rand.Intn(len(currencies))]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))

}
