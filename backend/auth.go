package backend

import (
	"crypto/sha512"
	"encoding/base64"
	"os"
)

/*
* Register user is data file not present
*
* first `64` bits of file are sha512 hash
* of password.
*
* permission of data file is set to `700`
*
 */
func Register(pass []byte) {
	encoded := base64.StdEncoding.EncodeToString(pass)
	hash := sha512.New()
	hash.Write([]byte(encoded))

	file, err := os.OpenFile(DF_NAME, os.O_CREATE | os.O_RDWR, 0700)
	if err != nil {
		panic(err)
	}
	file.Seek(0,0)
	file.Write(hash.Sum(nil))
}

/*
 * authenticate user and logges him in
 */

func Login(pass []byte) bool {
	encoded := base64.StdEncoding.EncodeToString(pass)
	hash := sha512.New()
	hash.Write([]byte(encoded))
	gpass := hash.Sum(nil)
	file, _ := os.Open(DF_NAME)
	rpass := make([]byte, 64)
	file.Read(rpass)

	for i := 0; i < 64; i++ {
		if gpass[i] != rpass[i] {
			return false
		}
	}
	Logged_in = true
	_passcode = string(pass)
	return true
}