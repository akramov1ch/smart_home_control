package utils

import (
    "gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", "shaxbozakramovic@gmail.com")
    m.SetHeader("To", to)
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)

    d := gomail.NewDialer("smtp.gmail.com", 587, "shaxbozakramovic@gmail.com", "vakhaboff2502")

    return d.DialAndSend(m)
}
