package webhook

// Webhook implements the Notifier interface
type Webhook struct {
	URL   string
	Token string
}
