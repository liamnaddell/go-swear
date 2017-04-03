package main

import "regexp"
import "fmt"
import "strings"

const swear = "I absolutely fucking hate you, kill yourself you little shit"
const fword_regex = `(?i)(f+u+(k+|c+k+))|f *u |f *u[^a-z\.]`
const sword_regex = `(?i)(s+h+(t+|i+t+))|s *h *i *t`

func main() {
	new1 := deleteFword(swear, "f***")
	new2 := deleteSword(new1, "s***")
	fmt.Println(new2)
}

func deleteFword(sentance string, toReplace string) string {
	r, err := regexp.Compile(fword_regex)
	if err != nil {
		fmt.Println(err)
	}
	stringnew := strings.Replace(sentance, r.FindString(sentance), toReplace, -1)
	return stringnew
}

func deleteSword(sentance string, toReplace string) string {
	r, err := regexp.Compile(sword_regex)
	if err != nil {
		fmt.Println(err)
	}
	stringnew := strings.Replace(sentance, r.FindString(sentance), toReplace, -1)
	return stringnew
}
