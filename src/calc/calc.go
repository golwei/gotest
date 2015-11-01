package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

func Usage() {
	fmt.Println("ex:add 2 3")
}
func main() {
	args := os.Args
	//	fmt.Println(len(args))
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("Usage:add 3 8")
			return
		}

		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE:add 3 4")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result:", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("sqrt 3-->")
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("这样输入:sqrt 3.")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println(ret)
	default:
		fmt.Println("default")
	}

}
