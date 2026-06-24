package config

const (
	SMTPHost     = "localhost"
	SMTPPort     = "1025"
	SenderEmail  = "siddharthjeph10@gmail.com"
	WorkerCount  = 5
	CSVPath      = "./data/emails.csv"
	Subject      = "Test Email"
	Body         = "Testing"
	DelayBetween = 50 // milliseconds between sends
)
