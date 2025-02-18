package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
