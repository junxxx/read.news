package env

import (
	"flag"
	"fmt"
	"os"
)

var SmtpPwd = flag.String("pwd", "", "the smtp password")

func init() {
	flag.Parse()
	if len(*SmtpPwd) <= 0 {
		fmt.Println("usage: -pwd smtppassword")
		os.Exit(0)
	}
}
