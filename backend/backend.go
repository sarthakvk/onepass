package backend

import (
	"crypto/aes"
	// "crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const DF_NAME string = "data.op"

var Logged_in bool = false

var _passcode string

type Password struct {
	Name     string
	Url      string
	Password string
}

var _data map[string]Password

func loadData() {
	hash, err := ioutil.ReadFile(DF_NAME)
	if err != nil {
		panic(err)
	}

	hash = hash[64:]
	if len(hash) > 0 {
		// iv := hash[:aes.BlockSize]
		// hash = hash[aes.BlockSize:]
		// block, _ := aes.NewCipher([]byte(_passcode))
		// cfb := cipher.NewCFBDecrypter(block, iv)
		// cfb.XORKeyStream(hash, hash)
		jsondata, err := base64.StdEncoding.DecodeString(string(hash))
		if err != nil {
			panic(err)
		}
		var tmpdata map[string]Password
		json.Unmarshal(jsondata, &tmpdata)
		fmt.Println(tmpdata)
	} else {
		_data = make(map[string]Password)
	}
}

func saveData() {
	jsonData, err := json.Marshal(_data)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(DF_NAME, os.O_WRONLY, 0700)
	if err != nil {
		panic(err)
	}
	file.Seek(64, 0)
	_, er := aes.NewCipher([]byte(_passcode))
	if er != nil {
		panic(er)
	}
	b := base64.StdEncoding.EncodeToString(jsonData)
	cipherText := make([]byte, aes.BlockSize+len(b))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	// cfb := cipher.NewCFBEncrypter(block, iv)
	// cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(b))
	file.Write(cipherText)
	
}
