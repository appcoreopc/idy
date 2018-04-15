package main

import (
	"github.com/pquerna/otp/totp"
	"github.com/pquerna/otp"
)
// otpauth://totp/Example:alice@google.com?secret=JBSWY3DPEHPK3PXP&issuer=Example


//Issuer string
//// Name of the User's Account (eg, email address)
//AccountName string
//// Number of seconds a TOTP hash is valid for. Defaults to 30 seconds.
//Period uint
//// Size in size of the generated Secret. Defaults to 10 bytes.
//SecretSize uint
//// Digits to request. Defaults to 6.
//Digits otp.Digits

func GenerateQr(issuer string, accountname string, timing uint) *otp.Key {

	opts := totp.GenerateOpts{ Issuer: issuer, AccountName: accountname, Period: timing}
	totp, err  := totp.Generate(opts)

	if err != nil {
		panic("error")
	}
	return totp
}

func Validate(key string, secret string) bool {
	return totp.Validate(key, secret);
}
