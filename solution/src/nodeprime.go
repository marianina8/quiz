package main

import (
    "bufio"
    "fmt"
    "os"
)

var longest string

func main() {

	fmt.Print("Enter input file: ")
	    var input string
	    fmt.Scanln(&input)

	file, err := os.Open(input)
	    if err != nil {
	        panic(err)
	    }

	defer file.Close()

	scanner := bufio.NewScanner(file)
	      for scanner.Scan() {
			  compare(scanner.Text())
	      }

	fmt.Println(longest)

}

func compare(word string) {
	if(len(word) > len(longest)){
		longest=word
	}	
}	

