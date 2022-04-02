package env

import (
	"flag"
	"fmt"
	"os"
)

//SmtpPwd is the google smtp secret key
var SmtpPwd = flag.String("pwd", "", "the smtp password")

//Once is the flag that run the task manual, in case normal service run failed
var Once = flag.String("once", "", "manunal run task once")

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
