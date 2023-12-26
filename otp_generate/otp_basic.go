package main

import (
	"fmt"
	"github.com/diebietse/gotp/v2"
	"time"
)

func main() {
	// generate time base TOTP, after 30sec automatically change Code
	// So then it will show invalid OTP, google auth use this
	// OTP mechanism.
	otpUrl := generateOTP()
	fmt.Println(otpUrl)
	time.Sleep(28 * time.Second)
	isValid := verifyOTP(otpUrl)
	if isValid {
		fmt.Println("OTP is valid.")
	} else {
		fmt.Println("OTP is not valid.")
	}
}

func generateOTP() string {
	secret, _ := gotp.DecodeBase32("4S62BZNFXXSZLCRO")
	totp, _ := gotp.NewTOTP(secret)
	otpString, _ := totp.Now()
	return otpString
}

func verifyOTP(otp string) bool {
	secret, _ := gotp.DecodeBase32("4S62BZNFXXSZLCRO")
	totp, _ := gotp.NewTOTP(secret)
	currentTime := time.Now().Unix()
	verify, err := totp.Verify(otp, int(currentTime))
	if err != nil {
		return false
	}
	return verify
}
