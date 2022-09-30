package common

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

const (
	secret            = "test"
	OtpEmailPeriodSec = 1800
)

func GenerateOtp() (string, error) {
	secretEncoded := base32.StdEncoding.EncodeToString([]byte(secret))
	code, err := totp.GenerateCodeCustom(secretEncoded, time.Now().UTC(), totp.ValidateOpts{
		Period:    OtpEmailPeriodSec,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		return "", err
	}
	return code, nil
}

func IsValidOtp(code string) bool {

	secretEncoded := base32.StdEncoding.EncodeToString([]byte(secret))

	ok, err := totp.ValidateCustom(code, secretEncoded, time.Now().UTC(), totp.ValidateOpts{
		Period:    OtpEmailPeriodSec,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil || !ok {
		return false
	}
	return true
}
