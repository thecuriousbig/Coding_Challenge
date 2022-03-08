package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"time"
)

const URL string = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
const NUMBER_OF_REPEAT int = 31

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	log.Println(fmt.Sprintf("%s took %s", name, elapsed))
}

func FindNTimeRepeatWord(words string, repeat int) []string {
	TimeTrack(time.Now())
	dict := make(map[string]int)
	var word []string

	// slice := strings.FieldsFunc(words, func(r rune) bool { return strings.ContainsRune(" .,", r) })
	slice := regexp.MustCompile("[\\,\\.\\s]+").Split(words, -1)
	for _, element := range slice {
		dict[element]++
	}

	for key, value := range dict {
		if value == repeat {
			word = append(word, key)
		}
	}

	return word
}

func main() {
	response, err := http.Get(URL)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	word := FindNTimeRepeatWord(string(responseData), NUMBER_OF_REPEAT)
	fmt.Println(word)
}
