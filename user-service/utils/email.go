package utils

import (
    "gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", "your-email@example.com")
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-password")

    return d.DialAndSend(m)
}
