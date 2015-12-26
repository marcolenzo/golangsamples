package main

import (
	"flag"
	"fmt"
	"net/smtp"
)

var from, to, server_host, server_port, username, passwd string

func main() {
	flag.StringVar(&from, "from", "", "Sender email address, e.g. user@domain.com")
	flag.StringVar(&to, "to", "", "Recipient email address, e.g. recipient@domain.com")
	flag.StringVar(&server_host, "server_host", "smtp.gmail.com", "Mail server hostname. Default is smtp.gmail.com")
	flag.StringVar(&server_port, "server_port", "587", "Mail server port. Default is 587")
	flag.StringVar(&username, "user", "", "Username @ Mail server. If not specified the sender email address is used.")
	flag.StringVar(&passwd, "passwd", "", "Password @ Mail server.")
	flag.Parse()
	if from == "" || to == "" {
		panic("-from and -to parameters are required.")
	}
	sendMail()
}

func sendMail() {
	auth := smtp.PlainAuth("", from, passwd, server_host)
	fmt.Println("Sending mail to: " + to)
	err := smtp.SendMail(server_host+":"+server_port, auth, from, []string{to}, []byte("This is the email body."))
	if err != nil {
		panic(err)
	}
}
