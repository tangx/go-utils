package utils

import (
	"fmt"
	"testing"
)

var (
	i      int = 3
	iSlice     = []int{1, 2, 3, 4, 5}
	sSlice     = []string{"a", "b", "c", "d", "e"}

	s    string = "a"
	sMap        = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	iMap        = map[int]string{1: "b", 2: "c", 3: "d"}
)

func Test_Contain(t *testing.T) {

	ok, err := Contains(i, iSlice)
	fmt.Printf("i in iSlice = %t , err(%v)\n", ok, err)

	ok, err = Contains(i, sSlice)
	fmt.Printf("i in sSlice = %t , err(%v)\n", ok, err)

	ok, err = Contains(s, sMap)
	fmt.Printf("s in sMap = %t , err(%v)\n", ok, err)

	ok, err = Contains(s, iMap)
	fmt.Printf("i in iMap = %t , err(%v)\n", ok, err)

	ok, err = Contains(i, iMap)
	fmt.Printf("i in iMap = %t , err(%v)\n", ok, err)

	ok, err = Contains(i, sMap)
	fmt.Printf("i in iMap = %t , err(%v)\n", ok, err)

}
