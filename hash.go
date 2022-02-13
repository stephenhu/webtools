package webtools

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)


func GenerateRandomHex(length int) (string, error) {

	if length < 1 {
		return "", errors.New(fmt.Sprintf(
			"%s GenerateRandomHex(): length must be greater than 0", APP_NAME))
	}
	
	buf := make([]byte, length)

	_, err := rand.Read(buf)

	if err != nil {
		return "", err
	} else {
		return hex.EncodeToString(buf), nil
	}

} // GenerateRandomHex


func GenerateHash(s string, salt string, hk string, p string,
	length int) (string, error) {

		if s == "" || hk == "" || p == "" || salt == "" {
			return "", errors.New(fmt.Sprintf(
				"%s GenerateHash(): empty string not allowed", APP_NAME))
		}
	
		if length < 1 {
			return "", errors.New(fmt.Sprintf(
				"%s GenerateHash(): length must be greater than 0", APP_NAME))
		}

		digest 	:= hmac.New(sha256.New, []byte(hk))
		
		digest.Write([]byte(s + salt + p))
	
		hash := hex.EncodeToString(digest.Sum(nil))
		
		return hash[:length], nil
	
} // GenerateHash


func GenerateHashAndSalt(s string, hk string, p string, length int) (
	string, string, error) {

	if s == "" || hk == "" || p == "" {
		return "", "", errors.New(fmt.Sprintf(
			"%s GenerateHashAndSalt(): empty string not allowed", APP_NAME))
	}

	if length < 1 {
		return "", "", errors.New(fmt.Sprintf(
			"%s GenerateHashAndSalt(): length must be greater than 0", APP_NAME))
	}

	salt, err := GenerateRandomHex(length)

	if err != nil {
		return "", "", err
	} else {

		hash, err := GenerateHash(s, salt, hk, p, length)

		if err != nil {
			return "", "", err
		} else {
			return hash, salt, nil
		}

	}

} // GenerateHashAndSalt


func GenerateToken(key string, length int) (string, error) {

	if key == "" {
		return "", errors.New(fmt.Sprintf(
			"%s GenerateHashAndSalt(): empty string not allowed", APP_NAME))
	}

	if length < 1 {
		return "", errors.New(fmt.Sprintf(
			"%s GenerateToken(): length must be greater than 0", APP_NAME))
	}

	salt, err := GenerateRandomHex(32)

	if err != nil {
		return "", err
	} else {

		text := []byte(time.Now().String() + salt)

		digest := hmac.New(sha256.New, []byte(key))
	
		digest.Write([]byte(text))
	
		hash := hex.EncodeToString(digest.Sum(nil))
	
		return hash[:length], nil
	
	}

} // GenerateToken


func GenerateSimpleHash(s string, length int) (string, error) {

		if s == "" {
			return "", errors.New(fmt.Sprintf(
				"%s GenerateSimpleHash(): empty string not allowed", APP_NAME))
		}

		if length < 1 {
			return "", errors.New(fmt.Sprintf(
				"%s GenerateSimpleHash(): length must be greater than 0", APP_NAME))
		}
	
		digest 	:= hmac.New(sha256.New, []byte(s))
		
		digest.Write([]byte(s))
	
		hash := hex.EncodeToString(digest.Sum(nil))
		
		return hash[:length], nil
	
} // GenerateSimpleHash
