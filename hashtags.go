package webtools

import (
	"log"
	"regexp"
	"strings"
)

const (
	HASHTAG_POUND				= "#"
	HASHTAG_MAX_LENGTH	= 33
)

const (
	HASHTAG_REGEXP		= "#[a-z0-9]{1,32}"
	HASHTAG_INVALID  	= "+,?-_.@$%^&*#!"
)


func ExtractHashtags(s string) []string {

	hashtags := []string{}

	if !strings.Contains(s, HASHTAG_POUND) {
		return hashtags
	} else {

		tokens := strings.Fields(s)

		for _, token := range tokens {

			t := strings.TrimSpace(token)

			if t[0] == HASHTAG_POUND[0] {
	
				if len(t) < HASHTAG_MAX_LENGTH {
				
					r, err := regexp.Compile(HASHTAG_REGEXP)

					if err != nil {
						log.Println(err)
					} else {

						if r.MatchString(t) && !strings.ContainsAny(t[1:],
							HASHTAG_INVALID) {
							hashtags = append(hashtags, t[1:])
						}
						
					}

				}

			}

		}

		return hashtags

	}

} // ExtractHashtags
