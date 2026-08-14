package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"github.com/renji18/rms/utils"
)

func SendEmail(to []string, subject, body string) {
	from := utils.Config.EMAIL_FROM
	password := utils.Config.EMAIL_PASSWORD
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	header := make(map[string]string)
	header["From"] = from
	header["To"] = strings.Join(to, ",")
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/html; charset=\"utf-8\""

	var message strings.Builder
	for k, v := range header {
		line := fmt.Sprintf("%s: %s\r\n", k, v)
		message.WriteString(line)
	}
	message.WriteString("\r\n")
	message.WriteString(body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message.String()))
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return
	}
}
