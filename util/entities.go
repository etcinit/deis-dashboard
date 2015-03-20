package util

// Page represents data sent back to the browser
type Page struct {
	Title string
	JSON  string //[]byte
	App   string
}
