package repository

var url string
var title string
var header string

func init() {
	url = "https://www.youtube.com/watch?v=7BJ3ZXpserc&list=RDMM&index=16"
	title = "hallo"
	header = "x-form-code-xxrl"
}

func GetURL() string {
	return url
}

func GetTitle() string {
	return title
}

func GetHeader() string {
	return header
}
