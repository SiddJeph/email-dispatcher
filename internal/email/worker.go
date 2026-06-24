package email

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"

	"github.com/SiddJeph/email-dispatcher/internal/config"
)

type Recipient struct {
	Name  string
	Email string
}

func Worker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := range ch {
		formattedMsg := fmt.Sprintf(
			"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\n",
			config.SenderEmail, r.Email, config.Subject, config.Body,
		)
		msg := []byte(formattedMsg)
		err := smtp.SendMail(
			config.SMTPHost+":"+config.SMTPPort, nil,
			config.SenderEmail, []string{r.Email}, msg,
		)
		if err != nil {
			log.Printf("Worker %d: failed to send to %s: %v", id, r.Email, err)
			continue
		}
		time.Sleep(config.DelayBetween * time.Millisecond)
		fmt.Printf("Worker %d: sent email to %s\n", id, r.Email)
	}
}
