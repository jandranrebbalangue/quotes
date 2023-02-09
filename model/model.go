package model

// quotes represents data about a  quotes.
type Quotes struct {
	ID           string   `json:"_id"`
	Author       string   `json:"author"`
	Content      string   `json:"content"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
}
