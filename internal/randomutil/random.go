package randomutil

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomFullName() string {
	return RandomString(5) + " " + RandomString(5)
}

func RandomBirthDate() time.Time {
	return time.Date(int(RandomInt(1970, 2010)), time.Month(RandomInt(1, 12)), int(RandomInt(1, 28)), 0, 0, 0, 0, time.Now().UTC().Location())
}

func RandomPassword() string {
	return RandomString(7)
}

func RandomEmail() string {
	return RandomString(5) + "@gmail.com"
}
