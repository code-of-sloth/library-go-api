package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jackc/pgtype"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomChar(prefix string, length int) (result string) {
	rand.NewSource(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	result = string(b)
	if val := strings.Trim(prefix, " "); val != "" {
		result = fmt.Sprintf("%s_%s", prefix, result)
	}
	return
}

func GetPgTime(t time.Time) (pgt pgtype.Timestamp) {
	pgt.Time = t
	pgt.Status = pgtype.Present
	return
}
