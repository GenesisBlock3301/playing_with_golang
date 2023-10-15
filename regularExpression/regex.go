package main

import (
	"fmt"
	"regexp"
)

func main() {
	// if any uppercase letter exist, it will give `true` else `false`
	re := regexp.MustCompile(`[A-Z]`)
	matches := re.Match([]byte("bangladesh"))
	fmt.Println(matches)

	/////////////////////////////////
	// `^` matches the start of string
	text := "The quick brown fox jump over the lazy dog."
	re = regexp.MustCompile(`^The`)
	matches = re.Match([]byte(text))
	fmt.Println(matches)

	/////////////////////////////////
	// `\.` ignore . using escape
	text = "The quick brown fox jump over the lazy dog."
	re = regexp.MustCompile(`dog\.`)
	matches = re.Match([]byte(text))
	fmt.Println("escape", matches)

	/////////////////////////////////
	// `^` matches the start of string
	text = "The quick brown fox jump over the lazy dog."
	re = regexp.MustCompile(`dog\.`)
	matches = re.Match([]byte(text))
	fmt.Println("^", matches)

	/////////////////////////////////
	// `*` matches zero or more occurrences. `a` can exist or not but
	// `+` means one or more occurrences, must `a` exist or not work.
	text = "ab"
	re = regexp.MustCompile(`a*b`)
	matches = re.Match([]byte(text))
	fmt.Println("+,*", matches)

	/////////////////////////////////
	// `?` matches zero or one. `u` exists or not it will work.
	text = "color or colour?"
	re = regexp.MustCompile(`colou?r`)
	mes := re.FindAll([]byte(text), -1)

	if mes != nil {
		for i, match := range mes {
			fmt.Printf("?-%s : %s\n", i, match)
		}
	} else {
		fmt.Println("No matches found")
	}

	/////////////////////////////////
	// `?` matches zero or one. `u` exists or not it will work.
	text = "colorrrrr or colour?"
	re = regexp.MustCompile(`colou?r`)
	mes = re.FindAll([]byte(text), -1)

	if mes != nil {
		for _, match := range mes {
			fmt.Printf("Match: %s\n", match)
		}
	} else {
		fmt.Println("No matches found")
	}

	/////////////////////////////////
	fmt.Println("{1,2} matches zero or one. `u` exists or not it will work.")
	text = "color colourr?"
	re = regexp.MustCompile(`colou?r{2,2}`)
	mes = re.FindAll([]byte(text), -1)

	if mes != nil {
		for _, match := range mes {
			fmt.Printf("Match: %s\n", match)
		}
	} else {
		fmt.Println("No matches found")
	}

	/////////////////////////////////
	fmt.Println("[aeiou] matches vowels")
	text = "The quick brown fox jump over the lazy dog."
	re = regexp.MustCompile(`[aeiouAEIOU]`)
	mes = re.FindAll([]byte(text), -1)

	if mes != nil {
		for _, match := range mes {
			fmt.Printf("Match: %s\n", match)
		}
	} else {
		fmt.Println("No matches found")
	}

	/////////////////////////////////
	fmt.Println("[0-9] matches vowels")
	text = "The price is is $1000"
	re = regexp.MustCompile(`[0-9]`)
	mes = re.FindAll([]byte(text), -1)

	if mes != nil {
		for _, match := range mes {
			fmt.Printf("Match: %s\n", match)
		}
	} else {
		fmt.Println("No matches found")
	}
	/////////////////////////////////
	fmt.Println("[^0-9] matches which is not belongs 0-9 with extra space")
	text = "The price is is $1000"
	re = regexp.MustCompile(`[^0-9 ]`)
	mes = re.FindAll([]byte(text), -1)

	if mes != nil {
		for _, match := range mes {
			fmt.Printf("Match: %s\n", match)
		}
	} else {
		fmt.Println("No matches found")
	}

	/////////////////////////////////
	fmt.Println("Email validation")
	text = "nas@gmail.com"
	re = regexp.MustCompile(`[a-zA-Z0-9@a-z.com]`)
	mes = re.FindAll([]byte(text), -1)

	if mes != nil {
		for _, match := range mes {
			fmt.Printf("Match: %s\n", match)
		}
	} else {
		fmt.Println("No matches found")
	}
}
