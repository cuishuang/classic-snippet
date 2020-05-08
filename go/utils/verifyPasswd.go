package main

import (
	"fmt"
)

func main() {

	password := "cuAishuang"

	rs := verifyPasswdIsLegal(password)

	fmt.Println(rs)
}

// 验证密码是否合法. 不使用正则,使用ascii码硬匹配方式,性能其实会比正则还要好一些
func verifyPasswdIsLegal(passwd string) bool {
	if len(passwd) < 8 || len(passwd) > 48 {
		return false
	}
	uppercaseLetter := 0
	lowerLetter := 0
	symbol := 0
	for _, v := range passwd {
		switch {
		case v >= 33 && v <= 64:
			{
				symbol = 1
			}
		case v >= 65 && v <= 90:
			{
				uppercaseLetter = 1
			}
		case v >= 91 && v <= 96:
			{
				symbol = 1
			}
		case v >= 97 && v <= 122:
			{
				lowerLetter = 1
			}
		case v >= 123 && v <= 126:
			{
				symbol = 1
			}
		default:
			return false
		}
	}
	if uppercaseLetter+lowerLetter+symbol < 2 {
		return false
	}
	return true
}
