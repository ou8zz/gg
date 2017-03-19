package main

import (
	"fmt"
	_ "net/http"
	_ "github.com/labstack/echo"
	_ "github.com/labstack/echo/middleware"
	"io/ioutil"
	"encoding/json"
)

var f bool = false

func aa() {
	var i int = 100
	var b string = "text"
	var c = "aaa"

	var g, p = 1, "astr"

	f = true
	fmt.Println("Hello, World:22222!", i, b, c, f, g, p)
}
func bb() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):  // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n" , i1, i2, c1, c2, c3)
	}
}
func cc(n1, n2 *int) (*int, string) {
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	*n1 = *n2
	return n1, "str"
}
func main2() {
	var n1, n2 = 100, 200
	var c, d = cc(&n1, &n2)
	fmt.Println(c, d, n1, n2)
}

var xxx = map[string]string{}

func readFile(filename string) (map[string]string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return nil, err
	}

	if err := json.Unmarshal(bytes, &xxx); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return nil, err
	}

	return xxx, nil
}

func main1() {
	xxxMap, err := readFile("/Users/ole/dev/gopath/src/gg/config/db.json")
	if err != nil {
		fmt.Println("readFile: ", err.Error())
		//return err
	}

	fmt.Println(xxxMap)
}