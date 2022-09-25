package main

import (
	"fmt"
	"strings"
)

func upperToLower(str string) string {
	var arr []string
	for i := 0; i < len(str); i++ {
		if string(str[i]) >= "A" && string(str[i]) <= "Z" {
			arr = append(arr, string(str[i]+32))
		} else {
			arr = append(arr, string(str[i]))
		}
	}
	return strings.Join(arr, "")
}

func lowerToUpper(str string) string {
	var arr []string
	for i := 0; i < len(str); i++ {
		if string(str[i]) >= "a" && string(str[i]) <= "z" {
			arr = append(arr, string(str[i]-32))
		} else {
			arr = append(arr, string(str[i]))
		}
	}
	return strings.Join(arr, "")
}

func longestPalindrome(s string) int {
	alpha := make([]int, 128)
	n := len(s)
	for _, val := range s {
		alpha[val]++
	}
	odd := 0
	for _, val := range alpha {
		if val%2 != 0 {
			odd++
		}
	}
	if odd > 1 {
		return n - (odd - 1)
	}
	return n
}

func main() {
	str := "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"
	//fmt.Println(upperToLower(str))
	//fmt.Println(lowerToUpper(str))
	fmt.Println(longestPalindrome(str))
}
