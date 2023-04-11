package todo

type Model struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
