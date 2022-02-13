package webtools

import (
	"testing"
)

var validHashtags = []string{
	"hello world #abc",
  "#abcxyz",
	"#go123",
	"#abcdefghijklmn",
	"#23623",
	"#me  too",
}

var invalidHashtags = []string{
	"#abc-xyz",
	"##asdb",
	"#_2as",
	"#abcdefghijklmntoolongboywhatevatoolongm",
	"#golf#mole",
}

func TestExtractHashtagsValid(t *testing.T) {

	for _, m := range validHashtags {

		hashtags := ExtractHashtags(m)

		if len(hashtags) == 0 {
			t.Error(hashtags)
		}
	
	}

} // TestExtractHashtagsValid


func TestExtractHashtagsInvalid(t *testing.T) {

	for _, m := range invalidHashtags {

		hashtags := ExtractHashtags(m)

		if len(hashtags) > 0 {
			t.Error(hashtags)
		}
	
	}

} // TestExtractHashtagsInvalid
