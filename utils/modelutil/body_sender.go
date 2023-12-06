package modelutil

type BodySender struct {
	To          []string `json:"to"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body"`
	CSVFilePath string
}
