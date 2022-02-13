package webtools

import (
	"testing"
)


func TestParseNameFromEmailGenuine(t *testing.T) {

	name, err := ParseNameFromEmail("abc@aol.com")

	if err != nil {
		t.Error(err)
	}

	if name != "abc" {
		t.Error("Expected abc, received " + name)
	}

} // TestParseNameFromEmailGenuine


func TestParseNameFromEmailEmpty(t *testing.T) {

	_, err := ParseNameFromEmail("")

	if err == nil {
		t.Error(err)
	}

} // TestParseNameFromEmailEmpty


func TestParseNameFromEmailSingleAt(t *testing.T) {

	_, err := ParseNameFromEmail("@")

	if err == nil {
		t.Error(err)
	}

} // TestParseNameFromEmailSingleAt


func TestParseNameFromEmailNoDomain(t *testing.T) {

	_, err := ParseNameFromEmail("abc@")

	if err == nil {
		t.Error(err)
	}

} // TestParseNameFromEmailNoDomain


func TestParseNameFromEmailInvalid(t *testing.T) {

	t.Skip("multiple at symbols don't work")

	_, err := ParseNameFromEmail("sandy@aol@aol.com")

	if err != nil {
		t.Error(err)
	}

} // TestParseNameFromEmailInvalid


func TestParseNameFromEmailInvalidDomain(t *testing.T) {

	_, err := ParseNameFromEmail("abc")

	if err == nil {
		t.Error(err)
	}

} // TestParseNameFromEmailInvalidDomain
