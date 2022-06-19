package tool

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	val := md5.New()
	val.Write([]byte(str))
	//fmt.Println("sum", val.Sum(nil))
	return hex.EncodeToString(val.Sum(nil))
}

func MD5secert(str string, cret string) string {
	if !IsStrNull(str) && !IsStrNull(cret) {
		return MD5(MD5(str) + MD5(cret))
	}
	if !IsStrNull(str) {
		return MD5(str)
	}
	return ""
}
