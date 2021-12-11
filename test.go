package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

func main() {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("token--->", token)
}
