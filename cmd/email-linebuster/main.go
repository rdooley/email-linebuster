package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/jordan-wright/email"
)

func main() {
	timeStr := flag.String("time", "2020-07-07T14:48:22-07:00", "time to send the email, expressed in RFC3339 format")
	from := flag.String("from", "", "from email")
	to := flag.String("to", "", "to email")
	subj := flag.String("subj", "", "email subj")
	pass := flag.String("pass", "", "gmail app password, see https://support.google.com/accounts/answer/185833?hl=en")
	emailPath := flag.String("email", "", "path to email file (txt only)")
	attachmentPath := flag.String("attachement", "", "path to attachment")

	flag.Parse()

	emailFile, err := os.Open(*emailPath)
	if err != nil {
		log.Fatal(err)
	}
	defer emailFile.Close()
	emailText, err := ioutil.ReadAll(emailFile)
	if err != nil {
		log.Fatal("error reading email file")
	}
	t, err := time.Parse(time.RFC3339, *timeStr)
	if err != nil {
		log.Fatal("Couldn't parse timestr\t" + *timeStr)
		return
	}
	duration := time.Until(t)
	// Check its in future
	if duration.Seconds() < 0 {
		log.Print(t)
		log.Fatal("this is in the past, cant send email then")
		return
	}
	// wait correct amount and do stuff"
	auth := smtp.PlainAuth("", *from, *pass, "smtp.gmail.com")
	e := email.NewEmail()
	e.AttachFile(*attachmentPath)
	e.From = *from
	e.To = []string{*to}
	e.Subject = *subj
	e.Text = append(emailText, []byte("\r\n Email sent  at "+t.String())...)
	log.Print("Wanted date is\t" + t.String() + "\t so waiting \t" + duration.String())
	time.Sleep(duration)
	log.Print("waited the duration")
	e.Send("smtp.gmail.com:587", auth)
	log.Print("Sent email from " + *from + " to " + *to)
}
