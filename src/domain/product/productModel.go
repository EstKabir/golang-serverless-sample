package product

type Model struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
