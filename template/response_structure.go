package template

type TemplateData struct {
	Outputs []Output `json:"outputs"`
}

type Output struct {
	SimpleText *SimpleText `json:"simpleText,omitempty"`
	ListCard   *ListCard   `json:"listCard,omitempty"`
}

type SimpleText struct {
	Text string `json:"text"`
}

type ListCard struct {
	Header  Header   `json:"header"`
	Items   []Item   `json:"items"`
	Buttons []Button `json:"buttons"`
}

type Header struct {
	Title string `json:"title"`
}

type Item struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Link        *Link  `json:"link"`
}

type Link struct {
	Web string `json:"web"`
}

type Button struct {
	Action     string `json:"action"`
	Label      string `json:"label"`
	WebLinkUrl string `json:"webLinkUrl,omitempty"`
}
