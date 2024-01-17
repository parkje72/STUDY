package main

import "fmt"

var lineCount int
var spaceCount int
var starCount int

func main() {
	// 1번
	lineCount = 10
	spaceCount = lineCount/2 + 1
	starCount = 1

	for i := 0; i < lineCount; i++ {
		for j := 0; j < starCount; j++ {
			fmt.Print("*")
		}
		for j := 0; j < spaceCount; j++ {
			fmt.Print(" ")
		}
		if i < lineCount/2-1 { //1 ~ 4 줄까지
			spaceCount -= 1
			starCount += 1
			fmt.Println()
		} else if i == lineCount/2-1 { // 5줄
			fmt.Println()
		} else { //6 ~ 10 줄까지
			spaceCount += 1
			starCount -= 1
			fmt.Println()
		}
	}
	// 2번
	lineCount = 8
	spaceCount = lineCount/2 - 1
	starCount = 1

	for i := 0; i < lineCount; i++ {
		//공백 + 별 + 공백
		for j := 0; j < spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < starCount; j++ {
			fmt.Print("*")
		}
		for j := 0; j < spaceCount; j++ {
			fmt.Print(" ")
		}
		if i < lineCount/2-1 {
			spaceCount -= 1
			starCount += 2
			fmt.Println()
		} else if i == lineCount/2-1 {
			fmt.Println()
		} else {
			spaceCount += 1
			starCount -= 2
			fmt.Println()
		}
	}
	//3번
	lineCount = 7
	spaceCount = lineCount / 2
	starCount = 1

	for i := 0; i < lineCount; i++ {
		//공백 + 별 + 공백
		for j := 0; j < spaceCount; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < starCount; j++ {
			fmt.Print("*")
		}
		for j := 0; j < spaceCount; j++ {
			fmt.Print(" ")
		}
		if i < lineCount/2 {
			spaceCount -= 1
			starCount += 2
			fmt.Println()
		} else {
			spaceCount += 1
			starCount -= 2
			fmt.Println()
		}
	}
}
