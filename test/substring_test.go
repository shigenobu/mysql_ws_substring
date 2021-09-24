package test

import (
	"crypto/md5"
	_ "crypto/md5"
	"fmt"
	"github.com/shigenobu/mysql_ws_substring/func"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSimple1(t *testing.T) {
	input := "aaabbbccc"

	output1 := _func.Substring(input, 3, 3)
	fmt.Println(output1)

	assert.Equal(t, "bbb", output1)
}

func TestSimple2(t *testing.T) {
	input := "𠮷野家で𠮷野がご飯をたべる"

	output1 := _func.Substring(input, 3, 3)
	fmt.Println(output1)

	assert.Equal(t, "で𠮷野", output1)
}

func TestHugeString(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./test1.txt")
	input := string(bytes)
	fmt.Println(md5.Sum([]byte(input)))

	output1 := _func.Substring(input, 0, 100000)
	fmt.Println(md5.Sum([]byte(output1)))

	assert.Equal(t, output1, input)
}


