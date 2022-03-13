package deliver

import (
	"errors"
	"log"
	"net/smtp"
	"net/textproto"
	"os"

	"github.com/jordan-wright/email"
	"github.com/junxxx/read.news/env"
)

var to = []string{"312866238@qq.com", "jinyanhuohuo@163.com"}

const (
	from     = "hprjunxxx@gmail.com"
	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

func addAttachFile(e *email.Email, filenames []string) {
	for _, filename := range filenames {
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			log.Println("file " + filename + "doesn't exist")
			return
		}

		e.AttachFile(filename)
	}
}

func DeliverDoc(filenames []string) {
	subject := "[VOA]Daily Reads\r\n\r\n"

	e := &email.Email{
		To:      to,
		From:    from,
		Subject: subject,
		Text:    []byte("Read news, have a nice day!"),
		// HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	addAttachFile(e, filenames)

	auth := smtp.PlainAuth("", from, *env.SmtpPwd, smtpHost)
	err := e.Send(smtpHost+":"+smtpPort, auth)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("email send successfully!")
}
