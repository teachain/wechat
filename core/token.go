package core

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

//token赋值为基本配置中的信息
//token,timestamp,nonce字典排序得字符串list
//哈希算法加密list得到hashcode
func VerifyToken(token string, timestamp string, nonce string, signature string) bool {
	list := make([]string, 0, 3)
	list = append(list, token)
	list = append(list, timestamp)
	list = append(list, nonce)
	sort.Strings(list)
	preSign := strings.Join(list, "")
	hashcode := Sha1(preSign)
	if strings.Compare(hashcode, signature) == 0 {
		return true
	} else {
		return false
	}
}

func Sha1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
