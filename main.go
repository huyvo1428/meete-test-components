package main

import (
	"bitbucket.org/meete/genesis-framework/component/captcha"
	"bitbucket.org/meete/genesis-framework/component/captcha/mock"
	"bitbucket.org/meete/genesis-framework/component/email"
	mock3 "bitbucket.org/meete/genesis-framework/component/email/mock"
	"bitbucket.org/meete/genesis-framework/component/sms"
	mock4 "bitbucket.org/meete/genesis-framework/component/sms/mock"
	"bitbucket.org/meete/genesis-framework/component/verification"
	mock2 "bitbucket.org/meete/genesis-framework/component/verification/mock"
	"fmt"
)

func main() {
	fmt.Println("--------------------------------")
	fmt.Println("CAPTCHA TEST")
	mockClient := mock.New(mock.Configuration{Debug: true})

	result, err := mockClient.Verify(captcha.Request{
		Token:  "abc",
		UserIP: "",
	})

	fmt.Println(result, err)
	fmt.Println("--------------------------------")
	fmt.Println("OTP TEST")

	otpMockClient := mock2.New(mock2.Configuration{
		Debug:     false,
		SecretKey: "abc",
		OTP:       "0000",
		TokenTTL:  0,
	})

	otp := otpMockClient.Generate(verification.Input{})
	validate, err := otpMockClient.Validate(verification.Input{}, otp)
	if err != nil {
		return
	}

	fmt.Println(otp)
	fmt.Println(validate)

	fmt.Println("--------------------------------")
	fmt.Println("EMAIL TEST")

	emailClient := mock3.New(mock3.Configuration{Debug: true})
	err = emailClient.Send(email.Message{
		Sender: struct {
			Name  string
			Email string
		}{"Chodeli",
			"no-reply@meete.co"},
		Recipient: struct {
			Name  string
			Email string
		}{"Vo Quoc Huy",
			"huyvo.meete@gmail.com"},
		Subject:     "EMAIL VERIFY",
		HtmlContent: "",
		TextContent: "HTTPS://LINK.MEETE.CO",
	})
	if err != nil {
		return
	}

	fmt.Println("--------------------------------")
	fmt.Println("SMS TEST")

	smsClient := mock4.New(mock4.Configuration{Debug: true})
	err = smsClient.Send(sms.Message{
		PhoneNumber: "0987789987",
		TextContent: "MA TRUNG THUONG CUA BAN LA: 0987789098",
	})
	if err != nil {
		return
	}

}
