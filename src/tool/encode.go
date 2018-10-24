package tool

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"math/rand"
	"time"
)

const (
	CHARS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Base64Encode(data string) (result string) {
	result = base64.StdEncoding.EncodeToString([]byte(data))
	return result
}

func Base64Decode(data string) string {
	result, _ := base64.StdEncoding.DecodeString(data)
	return string(result[:])
}

func Md5(data string, encode bool) string {
	result := md5.Sum([]byte(data))
	if encode {
		return Base64Encode(string(result[:]))
	}
	return string(result[:])
}

func Sha1(data string, encode bool) string {
	result := sha1.Sum([]byte(data))
	if encode {
		return Base64Encode(string(result[:]))
	}
	return string(result[:])
}

func RandomStr(count int) string {
	rand.Seed(time.Now().UnixNano())
	size := len(CHARS)
	var result []byte = make([]byte, count)
	for i := 0; i < count; i++ {
		idx := rand.Intn(size)
		result[i] = CHARS[idx]
	}
	return string(result)
}
