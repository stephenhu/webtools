package webtools

import (
	"testing"
)


var (
	blockKey1 		= "0123456789abcdef"
  blockKey2			= "0123456789abcde"
	blockKey3     = ""
  iv1       		= "abcdeaaaaaaaaaaa"
	iv2         	= "abcd"
	iv3           = ""
)


var (
	text1             = []byte("i need some sleep")
	text2							= []byte("defrag the drive")
	text3             = []byte("")
)


func TestEncryptValid(t *testing.T) {

	_, err := encrypt(text1, blockKey1, iv1)
	
	if err != nil {
		t.Error(err)
	}
	
} // TestEncryptValid


func TestEncryptMismatchBlockIVLength(t *testing.T) {

	_, err := encrypt(text1, blockKey1, iv2)
	
	if err == nil {
		t.Error(err)
	}
	
} // TestEncryptMismatchBlockIVLength

 
// This might be a valid case
func TestEncryptEmptyBlockIv(t *testing.T) {

	_, err := encrypt(text1, blockKey3, iv3)
	
	if err == nil {
		t.Error(err)
	}
	
} // TestEncryptEmptyBlockIv


func TestEncryptEmptyText(t *testing.T) {

	_, err := encrypt(text3, blockKey1, iv1)
	
	if err == nil {
		t.Error(err)
	}
	
} // TestEncryptEmptyText
