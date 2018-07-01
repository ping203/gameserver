package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

func GetCurrentMicroTimestamp() int64 {
	return int64(time.Now().UnixNano()) / 1e6
}

func MD5(str string) string {
	data := []byte(str)
	hash := md5.New()
	h := hash.Sum(data)
	return hex.EncodeToString(h)
}
