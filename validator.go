package youtube

import (
	"regexp"
)

func isCorrectUrl(url string) bool {
	r,_ := regexp.Compile("watch\?v=([a-zA-Z0-9\-_]+)")
}
