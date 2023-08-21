package libDingtalkBot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"unsafe"
)

func string2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

func bytes2String(b []byte) string {
	return unsafe.String(&b[0], len(b))
}

func timestampSign(timestamp, secret []byte) string {
	msg := bytes.Join([][]byte{
		timestamp,
		secret,
	}, []byte{'\n'})

	h := hmac.New(sha256.New, secret)
	h.Write(msg)
	sign := h.Sum(nil)

	b := make([]byte, base64.StdEncoding.EncodedLen(len(sign)))
	base64.StdEncoding.Encode(b, sign)

	return bytes2String(b)
}
