package util

import (
	"math/rand"
	"strings"
	"time"
)

var rng *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))

}

func RandomInt(min, max int64) int64 {
	if min >= max {
		panic("invalid range (min >= max)")
	}
	return min + rng.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for range n {
		c := alphabet[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 20000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "MAD"}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}
