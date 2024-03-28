package etc

import (
	"fmt"
	"net/smtp"
)

func SendCode(email, code string) {
	from := "nuriddinovdavron2003@gmail.com"
	password := "nkfw ugwe grje nytg"

	to := []string{
		email,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte(code)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
}
