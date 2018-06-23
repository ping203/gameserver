package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func MD5(str string) string {
	data := []byte(str)
	hash := md5.New()
	h := hash.Sum(data)
	return hex.EncodeToString(h)
}

// GeneratePKID todo 唯一ID生成
func GeneratePKID() uint64 {
	return 1
}
