package MQPassword

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
	"math/rand"
)

type CryptService struct {

}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func (this *CryptService) Encrypt(data []byte) []byte {
	if MAIN_KEY == "" {
		log.Fatal("主密钥未设置!")
	}

	contentEncrypt, err := _AesEncrypt(data, []byte(MAIN_KEY))
	if err != nil {
		log.Fatal("加密失败:", err)
	}

	return contentEncrypt
}

func (this *CryptService) Decrypt(data []byte) []byte {
	if MAIN_KEY == "" {
		log.Fatal("主密钥未设置!")
	}

	contentDecrypt, err := _AesDecrypt(data, []byte(MAIN_KEY))

	if err != nil {
		log.Println("解密失败:", err)
		return []byte("")
	}

	return contentDecrypt
}

func (this *CryptService) RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}

	return string(b)
}

/**********************AES加密********************************/
func _PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func _AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = _PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

func _AesDecrypt(cipherContent, key []byte) ([]byte, error) {
	//crypted, _ := base64.StdEncoding.DecodeString(string(cipherContent))
	crypted := cipherContent

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = _PKCS7UnPadding(origData)

	return origData, nil
}