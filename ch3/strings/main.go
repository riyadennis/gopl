package main

import "fmt"

func main(){
	fmt.Printf("%s\n", reverseString("riya"))
	fmt.Printf("%v\n", hasPrefix("Mr", "Mr John"))
	fmt.Printf("%v\n", hasPrefix("Mr", "M John"))
}

func reverseString(s string) string{
	runes := []rune(s)
	for i,j:=0,len(s)-1; i<j;i,j=i+1,j-1{
		runes[j], runes[i]= runes[i], runes[j]
	}
	return string(runes)
}

func hasPrefix(prefix, s string) bool{
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}