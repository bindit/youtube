package youtube

import (
	"regexp"
)

func isValidUrl(url string) bool {
	r,_ := regexp.Compile(`watch\?v=([a-zA-Z0-9\-_]+)`)
	
	return r.MatchString(url)
}
