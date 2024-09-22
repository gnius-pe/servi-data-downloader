package utils

import (
	"strings"
	"time"
)

func Concat(args ...string) string {
	return strings.Join(args, "")
}

func GetCurrentDateTime() string {
	return time.Now().Format("02-01-2006 15:04")
}
