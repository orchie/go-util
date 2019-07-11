package goutil

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

//EncodeMD5 EncodeMD5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return fmt.Sprintf("%x", m.Sum(nil))
}

//EncodeSha1 EncodeSha1
func EncodeSha1(value string) string {
	m := sha1.New()
	m.Write([]byte(value))
	return fmt.Sprintf("%x", m.Sum(nil))
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxMask = 1<<6 - 1 // All 1-bits, as many as 6
)

var src = rand.NewSource(time.Now().UnixNano())

//RandStringBytesMaskImprSrc 返回指定长度的随机字符串
//出处 https://colobu.com/2018/09/02/generate-random-string-in-Go/
//评论区 作者 heyuanchao
//实测该页面最快
func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for 10 characters!
	for i, cache, remain := n-1, src.Int63(), 10; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), 10
		}
		b[i] = letterBytes[int(cache&letterIdxMask)%len(letterBytes)]
		i--
		cache >>= 6
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
