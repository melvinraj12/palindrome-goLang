package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := "./textFiles/input/words.txt"
	arrWord := readFile(fileName)
	var postiveCount int
	var negativeCount int
	for _, iniStr := range arrWord {
		lowStr := strings.ToLower(iniStr)
		revStr := reverse(lowStr)
		isPalindrome := lowStr == revStr
		if isPalindrome {
			postiveCount = postiveCount + 1
		} else {
			negativeCount = negativeCount + 1
		}
		writeToFile(iniStr, isPalindrome)
	}
	fmt.Printf("Total number of given words that is a palindrome are %d \n", postiveCount)
	fmt.Printf("Total number of given words that is a palindrome are %d \n", negativeCount)
	fmt.Println("Task executed successfully")
}

func readFile(fileName string) (scannedTextArray []string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func reverse(initialWord string) (reversedWord string) {
	for _, temp := range initialWord {
		reversedWord = string(temp) + reversedWord
	}
	return
}

func writeToFile(finalWord string, flag bool) {
	var fileName string = "./textFiles/results/positive.txt"
	if !flag {
		fileName = "./textFiles/results/negative.txt"
	}
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if nil != err {
		fmt.Println(err)
	}
	fmt.Fprintln(file, finalWord)
	file.Close()
}
