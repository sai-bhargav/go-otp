package authenticator

import (
	"crypto/hmac"
	"crypto/sha1"
	"time"
	"math"
	"strings"
	"encoding/base32"
	"encoding/binary"
)

func AuthenticatorTotp(accessToken string) int32 {

	d, err := base32.StdEncoding.DecodeString(strings.ToUpper(accessToken))

	if err != nil{
		panic("Cannot decode the access token")

	}

	h := hmac.New(sha1.New, d)

	counter := int(time.Now().Unix() / int64(30))

	byteArr := make([]byte, 8)
	binary.BigEndian.PutUint64(byteArr, uint64(counter))

	h.Write(byteArr)

	sum := h.Sum(nil)

	offset := sum[len(sum)-1] & 0xf
	value := int64(((int(sum[offset]) & 0x7f) << 24) |
		((int(sum[offset+1] & 0xff)) << 16) |
		((int(sum[offset+2] & 0xff)) << 8) |
		(int(sum[offset+3]) & 0xff))

	l := 6
	otp := int32(value % int64(math.Pow10(l)))

	return otp
}
