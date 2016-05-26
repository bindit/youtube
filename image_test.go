package youtube

import (
	"testing"
)

type Param struct {
	Id string
	Version int
}

func TestGet(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param Param
		expected string
	} {
		{Param{"id", 1}, "http://img.youtube.com/vi/id/1.jpg"},
		{Param{"test", 2}, "http://img.youtube.com/vi/test/2.jpg"},
		{Param{"true", 3}, "http://img.youtube.com/vi/true/3.jpg"},
	}

	image := NewImage()

	for _, test := range tests {
		result := image.Get(test.param.Id, test.param.Version)
		if result != test.expected {
			t.Error(
				"For", test.param,
				"expected", test.expected,
				"got", result,
			)
		}
	}
}

func TestGetV1(t *testing.T) {
	t.Parallel()

	image := NewImage()
	result := image.GetV1("test")
	if result != "http://img.youtube.com/vi/test/1.jpg" {
		t.Error(
			"For", "test",
			"expected", "http://img.youtube.com/vi/test/1.jpg",
			"got", result,
		)
	}
}

func TestGetV2(t *testing.T) {
	t.Parallel()

	image := NewImage()
	result := image.GetV2("test")
	if result != "http://img.youtube.com/vi/test/2.jpg" {
		t.Error(
			"For", "test",
			"expected", "http://img.youtube.com/vi/test/2.jpg",
			"got", result,
		)
	}
}

func TestGetV3(t *testing.T) {
	t.Parallel()

	image := NewImage()
	result := image.GetV3("test")
	if result != "http://img.youtube.com/vi/test/3.jpg" {
		t.Error(
			"For", "test",
			"expected", "http://img.youtube.com/vi/test/3.jpg",
			"got", result,
		)
	}
}
