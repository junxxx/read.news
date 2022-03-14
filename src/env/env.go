package env

import (
	"flag"
	"fmt"
	"os"
)

var SmtpPwd = flag.String("pwd", "", "the smtp password")

// store the send log  date -> true
var SendLog = make(map[string]bool)

func init() {
	flag.Parse()
	if len(*SmtpPwd) <= 0 {
		*SmtpPwd = os.Getenv("SMTPPWD")
		if len(*SmtpPwd) <= 0 {
			fmt.Println("usage: -pwd smtppassword")
			fmt.Println("or set env SMTPPWD")
			os.Exit(0)
		}
	}
}
