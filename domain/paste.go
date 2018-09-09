package domain

// Paste represents a Paste Entity
type Paste struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
