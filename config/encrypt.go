package config

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	"strings"
)

func generateKey() string {
	c := 32
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("error:", err)
		return ""
	}
	// The slice should now contain random bytes instead of only zeroes.
	return base64.StdEncoding.EncodeToString(b)
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

/**
 * @private
 * Encrypt a string using standardized encyryption of AES256
 * @param plainText value to encrypt
 * @param secret secret to use
 */
func encryptString(plainText string, secret string) string {
	if plainText == "" {
		return plainText
	}

	if secret == "" {
		log.Fatalln("you must pass a secret")
	}

	// Generates 16 byte cryptographically strong pseudo-random data as IV
	// https://nodejs.org/api/crypto.html#crypto_crypto_randombytes_size_callback
	ivBytes := make([]byte, 16)
	ivText := base64.StdEncoding.EncodeToString(ivBytes)

	// encrypt using aes256 iv + key + plainText = encryptedText

	block, _ := aes.NewCipher([]byte(createHash(secret)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	encryptedValue := base64.StdEncoding.EncodeToString(ciphertext)

	// store base64(ivBytes)!base64(encryptedValue)
	return "${" + ivText + "}!${" + encryptedValue + "}";
}

/**
 * @private
 * Decrypt a string using standardized encyryption of AES256
 * @param enryptedValue value to decrypt
 * @param secret secret to use
 */
func decryptString(encryptedValue string, secret string) string {
	if encryptedValue == "" {
		return encryptedValue
	}

	if secret == "" {
		log.Println("you must pass a secret")
	}

	// enrypted value = base64(ivBytes)!base64(encryptedValue)
	var parts [] string = strings.Split(encryptedValue, "!")
	if (len(parts) != 2) {
		log.Println("The encrypted value is not a valid format")
	}

	ivText := parts[0]

	ivBytes := bytes.NewBufferString(ivText).Bytes()
	keyBytes := bytes.NewBufferString(secret).Bytes()

	if (len(ivBytes) != 16) {
		log.Println("The encrypted value is not a valid format")
	}

	if (len(keyBytes) != 32) {
		log.Println("The secret is not valid format")
	}

	// decrypt using aes256 iv + key + encryptedText = decryptedText

	key := []byte(createHash(secret))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := encryptedValue[:nonceSize], encryptedValue[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		panic(err.Error())
	}

	return string(plaintext)
}

/**
 * @private
 * @param encryptedValue
 * @param secret
 */
func legacyDecrypt(encryptedValue string, secret string) string {
	// LEGACY for pre standardized SHA256 encryption, this uses some undocumented nodejs MD5 hash internally and is deprecated
	block, _ := aes.NewCipher([]byte(createHash(secret)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(encryptedValue), nil)
	value := string(ciphertext)
	return value
}
