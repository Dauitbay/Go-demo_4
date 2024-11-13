package encrypter

import "os"

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Not given parameter KEY to env file")
	}
	return &Encrypter{
		Key: key,
	}
}

func (enc *Encrypter) Encrypt(plainStr string) string {
	return ""
}

func (enc *Encrypter) Decrypt(encryptedStr string) string{
	return ""
}
