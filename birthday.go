package birthday

import (
	"encoding/csv"
	"io"
	"net"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

func SendGreetings(filename string, now time.Time, smtpHost string, smtpPort int) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.TrimLeadingSpace = true
	header := false
	for {
		rec, err := r.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if !header {
			header = true
			continue
		}
		// Lastname, Firstname, dateOfBirth, email
		e, err := NewEmployee(rec[0], rec[1], rec[2], rec[3])
		if err != nil {
			return err
		}
		if e.IsBirthday(now) {
			body := strings.Replace("Happy Birthday, dear %NAME%", "%NAME%", e.Firstname, -1)
			err = sendMessage(smtpHost, smtpPort, "sender@here.com", "Happy Birthday!", body, e.Email)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func sendMessage(smtpHost string, smtpPort int, sender, subject, body, recipient string) error {
	smtpAddr := net.JoinHostPort(smtpHost, strconv.Itoa(smtpPort))
	return smtp.SendMail(smtpAddr, nil, sender, []string{recipient}, []byte(body))
}
