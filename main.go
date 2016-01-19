package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"regexp"
	"time"
)

type Config struct {
	Dirname       string
	ComplaintDept []string
	Complaintent  string
	SmtpAuthPass  string
	SmtpHost      string
	Subject       string
	SmtpPort      int
}

//["tlee@osd.wednet.edu", "deilers@osd.wednet.edu", "tfrazier@osd.wednet.edu"],

func main() {
	errorPreamble := "Transfinder Import reporter failed\r\n\r\nError message follows:\r\n\r\n"

	d, err := os.Open(C.Dirname)
	if err != nil {
		emailComplaint("TransFinder Import !ERROR! ", errorPreamble+err.Error())
		log.Print(err)
		return
	}

	files, err := d.Readdir(0)
	if err != nil {
		emailComplaint("TransFinder Import !ERROR! ", errorPreamble+err.Error())
		log.Print(err)
		return
	}

	FileRegex := `TFX_db22_` + time.Now().Format(`060201`) + `[0-9]{4,4}\.log` // TFX_db22_1522100350.log // October 10, 2015 3:50am

	re := regexp.MustCompile(FileRegex)
	if err != nil {
		emailComplaint("TransFinder Import !ERROR! ", errorPreamble+err.Error())
		log.Print(err)
		return
	}

	for _, val := range files {
		if nil != re.FindStringIndex(val.Name()) {
			// What do I do now?
			contents, err := ioutil.ReadFile(C.Dirname + string(os.PathSeparator) + val.Name())
			if err != nil {
				emailComplaint("TransFinder Import !ERROR! ", errorPreamble+err.Error())
				log.Print(err)
				return
			}

			ImportMsg := "Oh boy! I found an import log. I suppose I'll just include it below for your convenience. Have a wonderful day! \r\n\r\n" + string(contents)

			log.Println(ImportMsg)
			emailComplaint("TransFinder Import results ", ImportMsg)
			return
		}
	}

	emailComplaint("Transfinder Import File Log Not Found", "No import file found. \r\n\r\nIt likely didn't run. \r\n\r\nPerhaps someone forgot to log out. :-(\r\n")
	return

}

func emailComplaint(subject, txt string) {
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	auth := smtp.PlainAuth("", C.Complaintent, C.SmtpAuthPass, C.SmtpHost)
	to := C.ComplaintDept
	msg := []byte("To: " + C.Complaintent + "\r\n" +
		"Subject: " + C.Subject + "\r\n" +
		"\r\n" +
		txt +
		"\r\n\r\n")
	log.Print(auth, to, string(msg))
	err := smtp.SendMail(fmt.Sprintf("%s:%d", C.SmtpHost, C.SmtpPort), auth, C.Complaintent, to, msg)
	if err != nil {
		log.Fatal(err)
	}

}
