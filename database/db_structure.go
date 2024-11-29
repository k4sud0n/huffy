package database

type Menu struct {
	ID       int    `json:"id"`
	Date     string `json:"title"`
	Location string `json:"location"`
	Content  string `json:"content"`
}

type Notice struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Link  string `json:"link"`
	Date  string `json:"date"`
}
