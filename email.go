package webtools

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	AT_SYMBOL				=	"@"
	SPACE_SYMBOL		= " "
	EMAIL_REGEXP    = "([\\w-.+]+)@([\\w-.]+)"
)

func IsValidEmail(email string) bool {

	if email == "" {
		return false
	} else {

		r, err := regexp.Compile(EMAIL_REGEXP)

		if err != nil {
			log.Println(err)
			return false
		} else {
			return r.MatchString(email)
		}

		

	}

} // IsValidEmail


func ParseNameFromEmail(email string) (string, error) {

	if IsValidEmail(email) {

		strs := strings.Split(email, AT_SYMBOL)

		return strs[0], nil
	
	} else {
		return "", errors.New(fmt.Sprintf("%s ParseNameFromEmail(): %s",
		  APP_NAME, "invalid email format"))
	}

} // ParseNameFromEmail

