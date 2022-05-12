package codec

import (
	"crypto/md5"
	"fmt"
)

// MD5 执行MD5计算
func MD5(src string) string {
	data := []byte(src)
	bytes := md5.Sum(data)
	return fmt.Sprintf("%x", bytes)
}
