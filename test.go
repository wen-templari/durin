package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	timeUnix := time.Now().UnixMilli() //单位秒
	s := strconv.FormatInt(timeUnix, 10)
	fmt.Println(s)

}
