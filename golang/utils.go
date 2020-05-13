package golang

func min(first int, second int) int {
	if first < second {
		return first
	} else {
		return second
	}
}

func max(first int, second int) int {
	if first < second {
		return second
	} else {
		return first
	}
}

func toString() {

}

type QueryRes struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	LocalURL string `json:"local_url"`
	Content  string `json:"content"`
}
