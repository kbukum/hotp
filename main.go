package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"math"
	"strconv"
	"time"
)

type OTP struct {
	SharedSecret string
	Digits       int
	Crypto       Crypto
}

func (hotp *OTP) Password(counter int) string {
	encodedSecret := []byte(hotp.SharedSecret)
	T := fmt.Sprintf("%X", counter)
	T = lpad(T, "0", 16)
	msg, _ := hex.DecodeString(T)
	hash := encrypt(msg, encodedSecret, hotp.Crypto)
	return truncate(hash, hotp.Digits)
}

func lpad(value string, pad string, length int) string {
	for len(value) < length {
		value = pad + value
	}
	return value
}

type TOTP struct {
	OTP
	StartTime int64
	TimeStep  int64
}

func (totp *TOTP) Password() string {
	currentTime := time.Now().UTC()
	counter := int(math.Floor(float64((currentTime.Unix() - totp.StartTime) / totp.TimeStep)))
	return totp.OTP.Password(counter)
}

type Crypto int

const (
	SHA1 Crypto = iota + 1
	SHA256
	SHA512
)

func encrypt(msg, key []byte, crypto Crypto) []byte {
	var mac hash.Hash
	switch crypto {
	case SHA1:
		mac = hmac.New(sha1.New, key)
	case SHA256:
		mac = hmac.New(sha256.New, key)
	case SHA512:
		mac = hmac.New(sha512.New, key)
	}
	mac.Write(msg)
	return mac.Sum(nil)
}

var (
	DigitPowers = []int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000}
)

func truncate(hash []byte, digits int) string {
	offset := hash[len(hash)-1] & 0xf
	binCode := (int64(hash[offset])&0x7f)<<24 |
		(int64(hash[offset+1])&0xff)<<16 |
		(int64(hash[offset+2])&0xff)<<8 |
		(int64(hash[offset+3]) & 0xff)

	otp := binCode % DigitPowers[digits]
	return strconv.FormatInt(otp, 10)
}
