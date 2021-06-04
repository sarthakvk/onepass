package main

import (
	"fmt"
	"encoding/base64"
	"crypto/sha512"
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
func register() {
	fmt.Println("Welcome to OnePass! please set master password for using")
	
	var pass []byte
	fmt.Scan(&pass)
	encoded := base64.StdEncoding.EncodeToString(pass)
	hash := sha512.New()
	hash.Write([]byte(encoded))
	
	os.WriteFile(DF_NAME, hash.Sum(nil), os.FileMode(0700))
	
	rkey, _ := GenerateRandomString(64);

	file, err := os.OpenFile(DF_NAME, os.O_APPEND, 0700);
	if err != nil {
		panic(err);
	}

	defer file.Close();

	hash.Reset()
	hash.Write([]byte(rkey))
	file.Write(hash.Sum(nil))


}

/*
 * authenticate user and logges him in
 */
func login() bool {
	fmt.Print("Enter Passcode: ")
	var pass []byte
	fmt.Scan(&pass)
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
	logged_in = true
	return true
}