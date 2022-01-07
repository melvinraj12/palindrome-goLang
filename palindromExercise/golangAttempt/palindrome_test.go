package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	expected := []string{"Axa", "a", "axc", "Cba"}
	result := readFile("./textFiles/testInput/test.txt")
	if len(result) != 4 {
		t.Errorf("Total number of elements in array, got: %d , want: %d.", len(result), len(expected))
	}

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Array not the same the one present in test file")
	}
}

func TestReverse(t *testing.T) {
	expected := "cba"
	result := reverse("abc")
	if result != expected {
		t.Errorf("Reverse function doesnt work, got:%s, want:%s.", result, expected)
	}
}

func TestWriteToFile(t *testing.T) {
	writeTest := func(flag bool) {
		switch flag {
		case true:
			err := os.Truncate("./textFiles/results/positive.txt", 0)
			if err != nil {
				fmt.Println(err)
			}
			var sample string = "Axa"
			expected := []string{"Axa"}
			writeToFile(sample, true)
			result := readFile("./textFiles/results/positive.txt")
			if !reflect.DeepEqual(expected, result) {
				t.Errorf("File content did not get written to positive text file.")
			}
		case false:
			err := os.Truncate("./textFiles/results/negative.txt", 0)
			if err != nil {
				fmt.Println(err)
			}
			var sample2 string = "Axa"
			expected2 := []string{"Axa"}
			writeToFile(sample2, false)
			result2 := readFile("./textFiles/results/negative.txt")
			if !reflect.DeepEqual(expected2, result2) {
				t.Errorf("File content did not get written to negative text file.")
			}
		}
	}

	//test write to positive text file
	writeTest(true)
	//test write to negative text file
	writeTest(false)
}
