package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	str := "23-ab-48-caba-56-haha"
	result, err := wholeStory(str)
	if err != nil {
		log.Println("Error : " + err.Error())
		os.Exit(-1)
	}
	fmt.Println(result)
	strstats, errstat := storyStats(result)
	if errstat != nil {
		os.Exit(-1)

	}
	fmt.Println("shortest : " + strstats["shortest"])
	fmt.Println("longest  : " + strstats["longest"])
	fmt.Println("average  : " + strstats["average"])
	x := round(strstats["average"])
	fmt.Println(fmt.Sprintf("average round : %v", x))

	strrandom := RandStringBytesMask(true, result)
	fmt.Println(strrandom)
	strrandom = RandStringBytesMask(false, result)
	fmt.Println(strrandom)

}
func round(x string) float64 {
	i, _ := strconv.ParseFloat(x, 10)
	up := math.RoundToEven(float64(i)) // 2
	return up

}

/*
 text that is composed from all the text words separated by spaces
*/
func wholeStory(str string) (string, error) {
	var err error
	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("can't work empty string")
		return "", err
	}
	// patern for extract -1 -2- -3

	patern := `[-]?\d[\d]*[\]?[\d{2}]*?[-]`
	re := regexp.MustCompile(patern)

	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	fmt.Printf("String contains any match: %v\n", re.MatchString(str)) // True
	// Finding the number from
	// the given string
	// Using FindAllStrings() method
	submatchall := re.FindAllString(str, -1)
	for _, element := range submatchall {
		str = strings.Replace(str, element, " ", -1)
	}
	str = strings.TrimSpace(str)

	return str, err

}

/*
  * Write a function storyStats that returns four things:

* the shortest word

* the longest word

* the average word length

*/
func storyStats(str string) (map[string]string, error) {
	var err error
	m := make(map[string]string)
	m["shortest"] = ""
	m["longest"] = ""
	m["average"] = ""

	str = strings.TrimSpace(str)
	if str == "" {
		err = errors.New("can't work empty string")
		return nil, err
	}
	var shortest int
	var longest int
	words := strings.Fields(str)

	for index, element := range words {
		if index == 0 {
			// set intail values
			shortest = len(element)
			longest = len(element)
			m["shortest"] = element
			m["longest"] = element
			continue
		}
		// check longest and shotest wors
		if shortest > len(element) {
			m["shortest"] = element
			shortest = len(element)
		}
		if longest < len(element) {
			m["longest"] = element
			longest = len(element)
		}

	}
	m["average"] = fmt.Sprint(averageNumber(str))

	return m, err

}

// averageNumber that takes the string, and returns the average number from all the numbers

func averageNumber(str string) float32 {
	str = strings.TrimSpace(str)
	words := strings.Fields(str)
	var countwords float32
	var totallen float32
	for _, element := range words {
		countwords = countwords + 1
		totallen = totallen + float32(len(element))

	}
	avr := totallen / countwords
	return avr
}

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(flag bool, letterBytes string) string {
	n := len(letterBytes)
	if !flag {
		letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	}
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}
