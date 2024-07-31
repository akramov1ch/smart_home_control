package utils

import (
    "crypto/tls"
    "fmt"
    "gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
    mailer := gomail.NewMessage()
    mailer.SetHeader("From", "shaxbozakramovic@gmail.com")
    mailer.SetHeader("To", to)
    mailer.SetHeader("Subject", subject)
    mailer.SetBody("text/html", body)

    dialer := gomail.NewDialer("smtp.gmail.com", 587, "shaxbozakramovic@example.com", "vakhaboff2502")
    dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

    if err := dialer.DialAndSend(mailer); err != nil {
        return fmt.Errorf("could not send email: %v", err)
    }
    return nil
}
