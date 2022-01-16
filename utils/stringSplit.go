package utils

import (
	"strings"
)

func SplitString(move string, index int) string {
	return strings.Split(move, "\n")[index]
}
