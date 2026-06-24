package producer

import (
	"encoding/csv"
	"os"

	"github.com/SiddJeph/email-dispatcher/internal/email"
)

func LoadRecipients(filePath string, ch chan email.Recipient) error {
	defer close(ch)
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		ch <- email.Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}
	return nil
}
