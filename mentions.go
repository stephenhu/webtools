package webtools

import (
	"log"
	"regexp"
	"strings"
)

const (
	MENTION_AT							= "@"
	USERNAME_MAX_LENGTH			= 25
)

const (
	USERNAME_REGEXP		= "@[a-z][a-z0-9_]{1,23}"
	USERNAME_INVALID  = "+,?-.@$%^&*#!"
)


func ExtractMentions(s string) []string {

	mentions := []string{}

	if !strings.Contains(s, MENTION_AT) {
		return mentions
	} else {

		tokens := strings.Fields(s)

		for _, token := range tokens {

			t := strings.TrimSpace(token)


			if t[0] == MENTION_AT[0] {
	
				if len(t) < USERNAME_MAX_LENGTH {
				
					r, err := regexp.Compile(USERNAME_REGEXP)

					if err != nil {
						log.Println(err)
					} else {

						if r.MatchString(t) && !strings.ContainsAny(t[1:],
							USERNAME_INVALID) {
							mentions = append(mentions, t[1:])
						}
						
					}

				}

			}

		}

		return mentions

	}

} // ExtractMentions
