package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
)

type EncryptData []byte

func (data *EncryptData) HMacMd5(key string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write(*data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
