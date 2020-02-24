package main

import (
"fmt"
)

func main(){
	months :=[...]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}
	fmt.Printf("%v\n", months)
	fmt.Printf("%v\n", reverse(months[:]))

	nums := []int{0, 1,2,3,4,5,6}
	reverseInt(nums)
	reverseInt(nums[2:])
	reverseInt(nums[:2])
	fmt.Printf("%v\n", nums)

	runes := []rune("HELLO,世界日報")
	fmt.Printf("%q\n", reverseRunes(runes))

	s := []string{"A", "B", "C", "", "D", "E"}
	fmt.Printf("%v\n", removeEmpty(s))
	fmt.Printf("Top Element %v\n", s[len(s)-1])
	fmt.Printf("Top Element %v\n", s[:len(s)-1])

}

func removeEmpty(strings []string)[]string{
	out := strings[:0]
	for _, s := range strings{
		if s != ""{
			out = append(out,s)
		}
	}
	return out
}


func reverseRunes(runes []rune) []rune{
	for i,j:=0,len(runes)-1; i <j ; i,j=i+1,j-1{
		runes[j], runes[i] = runes[i], runes[j]
	}
	return runes
}

func reverseInt(nums []int)[]int{
	for i,j:=0,len(nums)-1; i<j;i,j=i+1, j-1{
		nums[j],nums[i] = nums[i],nums[j]
	}
	return nums
}

func reverse(s []string) []string{
	for i,j:=0,len(s)-1; i < j; i,j=i+1,j-1{
		s[j], s[i] = s[i],s[j]
	}
	return s
}
