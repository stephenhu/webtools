package webtools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"log"
)


func Decrypt(encText string, blockKey string,
	iv string) ([]byte, error) {

	if len(blockKey) != len(iv) {
		return nil, errors.New(ERR_BLOCK_IV_LENGTH_MISMATCH)
	} else if len(blockKey) == 0 || len(iv) == 0 {
		return nil, errors.New(ERR_EMPTY_BLOCK_IV)
	} else if len(encText) == 0 {
		return nil, errors.New(ERR_EMPTY_TEXT)
	}

	block, err := aes.NewCipher([]byte(blockKey))

	if err != nil {
	
		log.Println(err)
		return nil, err
	
	} else {

		cfb := cipher.NewCFBDecrypter(block, []byte(iv))

		decoded, err := hex.DecodeString(encText)

		if err != nil {

			log.Println(err)
			return nil, err

		} else {

			clearText := make([]byte, len(decoded))

			cfb.XORKeyStream(clearText, decoded)

			return clearText, nil

		}

	}

} // decrypt


func Encrypt(clearBuf []byte, blockKey string,
	iv string) (string, error) {

	if len(blockKey) != len(iv) {
		return STR_EMPTY, errors.New(ERR_BLOCK_IV_LENGTH_MISMATCH)
	} else if len(blockKey) == 0 || len(iv) == 0 {
		return STR_EMPTY, errors.New(ERR_EMPTY_BLOCK_IV)
	} else if len(clearBuf) == 0 {
		return STR_EMPTY, errors.New(ERR_EMPTY_TEXT)
	}

	block, err := aes.NewCipher([]byte(blockKey))

	if err != nil {
		return STR_EMPTY, err
	} else {

		cfb := cipher.NewCFBEncrypter(block, []byte(iv))

		encText := make([]byte, len(clearBuf))

		cfb.XORKeyStream(encText, clearBuf)

		return hex.EncodeToString(encText), nil

	}

} // encrypt
