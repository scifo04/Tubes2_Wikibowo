package back

import (
	"fmt"
	"strings"
)

func findElement(data [][]string,end string) int{
	for i := len(data) - 1; i >= 0; i-- {
		// fmt.Println(data[i][len(data[i])-1]," ",end)
		if (data[i][len(data[i]) - 1] == end){
			return i
		}
	}
	return -1
}

// func goPrint(data [][]string,end string) {
// 	for i := len(data) - 1; i >= 0; i-- {
// 		fmt.Println(data[i][len(data[i])-1]," ",end)
// 	}
// }

// func isInList(data []string, title string) bool {
// 	for i := 0; i < len(data); i++ {
// 		if data[i] == title {
// 			return true;
// 		}
// 	}
// 	return false
// }

func TurnToWikipedia(title string) string {
	tokens := strings.Split(title, " ")
	var temp string = ""
	for i := 0; i < len(tokens); i++ {
		temp = temp + tokens[i]
		if (i < len(tokens)-1) {
			temp = temp + "_"
		}
	}
	temp = "https://en.wikipedia.org/wiki/" + temp
	if (strings.Contains(temp,"–")) {
		temp = strings.Replace(temp,"–","%E2%80%93",-1)
	}
	return temp
}

func TurnToTitle(url string) string {
	url = url[30 : len(url)-0]
	temp := ""
	tokens := strings.Split(url, "_")
	for i := 0; i < len(tokens); i++ {
		temp = temp + tokens[i]
		if (i < len(tokens)-1) {
			temp = temp + " "
		}
	}
	if (strings.Contains(temp,"%E2%80%93")) {
		temp = strings.Replace(temp,"%E2%80%93","–",-1)
	}
	fmt.Println(temp)
	return temp
}