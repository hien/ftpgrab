package notifier

import "github.com/ftpgrab/ftpgrab/internal/journal"

// Handler is a notifier interface
type Handler interface {
	Name() string
	Send(jnl journal.Client) error
}

// Notifier represents an active notifier object
type Notifier struct {
	Handler
}
