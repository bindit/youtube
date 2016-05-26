package youtube

import (
	"strconv"
)

type Image struct {
	url string
}

func NewImage() *Image {
	return &Image{
		"http://img.youtube.com/vi/",
	}
}

func (i Image) Get(id string, v int) string {
	return i.url + id + "/" + strconv.Itoa(v) + ".jpg"
}

func (i Image) GetV1(id string) string {
	return i.Get(id, 1)
}

func (i Image) GetV2(id string) string {
	return i.Get(id, 2)
}

func (i Image) GetV3(id string) string {
	return i.Get(id, 3)
}
