package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func md5s(s string) string {
	r := md5.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func md5f(fName string) string {
	f, e := os.Open(fName)
	if e != nil {
		log.Fatal(e)
	}
	h := md5.New()
	_, e = io.Copy(h, f)
	if e != nil {
		log.Fatal(e)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func sha1f(fName string) string {
	f, e := os.Open(fName)
	if e != nil {
		log.Fatal(e)
	}
	h := sha1.New()
	_, e = io.Copy(h, f)
	if e != nil {
		log.Fatal(e)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func main() {

	fmt.Println(md5s("Hello, Gopher!"))
	//fmt.Println(md5f("E:\\testfile"))

	fmt.Println(sha1s("Hello, Gopher!"))
	//fmt.Println(sha1f("E:\\testfile"))
}
