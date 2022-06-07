package repository

var url string
var title string

func init() {
	url = "https://www.youtube.com/watch?v=7BJ3ZXpserc&list=RDMM&index=16"
	title = "hallo"
}

func GetURL() string {
	return url
}

func GetTitle() string {
	return title
}
