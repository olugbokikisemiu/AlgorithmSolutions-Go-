package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	fmt.Println(NoOfOccurrence([]int{2, 3, 4, 3, 2, 1}, 2))
	//crawlSite()
	//fmt.Println(BirthChocolate([]int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}, 18, 7))
}

func findWords(word []string) {
	wordLength := []string{}
	for i := 0; i < len(word); i++ {
		if len(wordLength) == 0 {
			wordLength = append(wordLength, word[i])
		}

	}
}

func scannerCode() {
	scanner := bufio.NewScanner(os.Stdin)
	var scanWords []string
	for scanner.Scan() {
		input := scanner.Text()
		scanWords = append(scanWords, input)
		if len(scanWords) == 2 {
			break
		}
	}
	overAll, _ := strconv.Atoi(scanWords[0])
	diff, _ := strconv.Atoi(scanWords[1])

	share := (overAll - diff) / 2

	fmt.Println(share + diff)
	fmt.Println(share)
}

func stairs(n int) {
	d := "#"
	for i := 1; i <= n; i++ {
		r := strings.Repeat(d, n)
		s := strings.Replace(r, "#", " ", n-i)
		fmt.Println(strings.Trim(s, "\n"))
	}
}

func diagonalDifference(arr [][]int32) int32 {

	var fromLeft, fromRight int32

	for i := 0; i < len(arr); i++ {

		fromLeft += arr[i][i]
		fromRight += arr[i][(len(arr)-1)-i]
	}
	ans := math.Abs(float64(fromLeft - fromRight))
	diag := int32(ans)
	return diag

}

func plusMinus(arr []int32) {
	var plusValue, minusValue, zeroValue float64
	for i := 0; i < len(arr); i++ {
		strValue := strconv.Itoa(int(arr[i]))
		if strings.EqualFold(strValue, "0") {
			zeroValue++
		} else if strings.Contains(strValue, "-") {
			minusValue++
		} else {
			plusValue++
		}
	}
	arrayLength := float64(len(arr))
	fmt.Println(plusValue / arrayLength)
	fmt.Println(minusValue / arrayLength)
	fmt.Println(zeroValue / arrayLength)
}

func plusMinus2(arr []int32) {
	var plusValue, minusValue, zeroValue float64
	for i := 0; i < len(arr); i++ {
		if math.Signbit(float64(arr[i])) {
			minusValue++
		} else if arr[i] == 0 {
			zeroValue++
		} else {
			plusValue++
		}
	}

	arrayLength := float64(len(arr))
	fmt.Println(plusValue / arrayLength)
	fmt.Println(minusValue / arrayLength)
	fmt.Println(zeroValue / arrayLength)
}

func miniMaxSum(arr []int) {
	var minValue, maxValue int
	sort.Ints(arr)
	for i := 0; i < 4; i++ {
		minValue += arr[i]
		maxValue += arr[len(arr)-(i+1)]
	}

	fmt.Printf("%d %d", minValue, maxValue)
}

func birthdayCakeCandles(ar []int) int32 {
	var counter = int32(0)
	sort.Ints([]int(ar))
	max := ar[len(ar)-1]
	for i := 0; i < len(ar); i++ {
		println("Max ", max)
		if max == ar[i] {
			counter++
		}
	}
	return counter
}

func isValid1(s string) string {
	sArray := strings.Split(s, "")
	mapString := make(map[string]int, len(sArray))
	var counter = 1
	for _, v := range sArray {
		if _, ok := mapString[v]; ok {
			mapString[v] = mapString[v] + 1
		} else {
			mapString[v] = counter
		}
	}
	var min = len(sArray)
	var max int
	var maxCounter, minCounter int
	for _, v := range mapString {
		if v > max {
			max = v
			maxCounter++
		} else if max == v {
			maxCounter++
		}
		if v < min {
			min = v
			minCounter++
		} else if min == v {
			minCounter++
		}
	}
	fmt.Println(maxCounter, " ", minCounter)
	if (max-min) <= 1 && math.Abs((float64(maxCounter)-float64(minCounter))) <= 1 {
		return "YES"
	}
	return "NO"
}

func isValid(s string) string {
	var min = len(s)
	var max int
	var maxCounter, minCounter int
	for _, v := range s {
		occurence := len(strings.Split(s, string(v)))
		if occurence > max {
			max = occurence
			maxCounter++
		} else if max == occurence {
			maxCounter++
		}
		if occurence < min {
			min = occurence
			minCounter++
		} else if min == occurence {
			minCounter++
		}
	}
	fmt.Println(maxCounter, " ", minCounter)
	if (max-min) <= 1 && math.Abs((float64(maxCounter)-float64(minCounter))) <= 1 {
		return "YES"
	}
	return "NO"
}

func gradingStudents(grades []int32) []int32 {
	var result []int32
	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			result = append(result, grades[i])
		} else {
			module := grades[i] % 5
			round := 5 - module
			if round < 3 {
				grades[i] += round
				result = append(result, grades[i])
			} else {
				result = append(result, grades[i])
			}
		}
	}
	return result
}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var appleCount int
	var orangeCount int
	for i := 0; i < len(apples); i++ {
		if a+apples[i] >= s && a+apples[i] <= t {
			appleCount++
		}
	}

	for i := 0; i < len(oranges); i++ {
		if b+oranges[i] >= s && b+oranges[i] <= t {
			orangeCount++
		}
	}

	fmt.Println(appleCount)
	fmt.Println(orangeCount)
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	i := int32(0)
	b2 := x2
	for i < b2 {
		x1 += v1
		x2 += v2
		if x1 == x2 {
			return "YES"
		}
		i++
	}
	return "NO"
}

// func getTotalX(a []int32, b []int32) int32 {
// 	counter := int32(0)
// 	c := append(a, b...)
// 	var value int32
// 	fmt.Println("Join ", c)
// 	for i := a[0]; i <= b[len(b)-1]; i++ {
// 		lengthCounter := 0
// 		for j := 0; j < len(c); j++ {
// 			if c[j]%i == 0 {
// 				fmt.Println("a ", i)
// 				fmt.Println("b ", c[j])
// 				value = c[j]
// 				lengthCounter++
// 			}
// 		}
// 		if lengthCounter == len(b) {
// 			fmt.Println("V: ", value)
// 			counter++
// 		}
// 	}
// 	return counter
// }

func getTotalX(a []int32, b []int32) int32 {
	var validCount int32
	counter := int32(0)
	for i := a[0]; i <= b[len(b)-1]; i++ {
		if PassesFirstRule(a[0]+counter, a) && PassesSecondRule(a[0]+counter, b) {
			validCount++
		}
		counter++
	}

	return validCount
}

func PassesFirstRule(element int32, firstArray []int32) bool {
	var goodStudents []int32
	for _, val := range firstArray {
		if element%val == 0 {
			goodStudents = append(goodStudents, element)
		}
	}

	return len(goodStudents) == len(firstArray)
}

func PassesSecondRule(element int32, secondArray []int32) bool {
	var goodStudents []int32
	for _, val := range secondArray {
		if val%element == 0 {
			goodStudents = append(goodStudents, element)
		}
	}

	return len(goodStudents) == len(secondArray)
}

func breakingRecords(stat []int32) []int32 {
	var minRecordsBroken int32
	var maxRecordsBroken int32

	min := stat[0]
	max := int32(-1)

	for _, val := range stat {
		if val < min {
			min = val
			minRecordsBroken++
		}

		if val > max {
			if max == -1 {
				max = val
				continue
			}
			max = val
			maxRecordsBroken++
		}
	}

	return []int32{maxRecordsBroken, minRecordsBroken}
}

func LongestEvenWords(words string) string {
	maxLength := 0
	longestString := "00"
	wordSplit := strings.Split(words, " ")
	for i := 0; i < len(wordSplit); i++ {
		if len(wordSplit[i])%2 == 0 && len(wordSplit[i]) > maxLength {
			maxLength = len(wordSplit[i])
			longestString = wordSplit[i]
		}
	}
	return longestString
}

func BirthChocolate(bar []int32, d, m int32) int32 {
	counter := int32(0)
	forCounter := int32(0)
	startP := int32(0)
	cLength := m - 1
	for forCounter < int32(len(bar)) {
		sumValue := int32(0)
		for i := startP; i <= cLength; i++ {
			sumValue += bar[i]
			forCounter++

		}
		if sumValue == d {
			counter++
		}
		startP++
		cLength = (cLength + m) - 1
	}
	return counter
}

func BinarySearch(list []int, item int) string {
	low := 0
	high := len(list) - 1
	counter := 0
	for low <= high {
		mid := low + high
		guess := list[mid]
		if guess == item {
			fmt.Printf("No of search %d \n", counter)
			return strconv.Itoa(mid)
		} else if guess > item {
			high = mid - 1
		} else if guess < item {
			low = mid + 1
		}
		counter++
		fmt.Printf("No of search %d \n", counter)
	}
	return "NONE"
}

func smallest(arr []int) int {
	//var sortArray []int
	smallest := 9999
	smallestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func SortAsc(arr []int) []int {
	var sortArray []int
	for _ = range arr {
		small := smallest(arr)
		sortArray = append(sortArray, arr[small])
		arr = append(arr[:small], arr[small+1:]...)
	}
	return sortArray
}

func Letter(l, n string) bool {
	s := strings.Fields(l)
	var count int
	for i := 0; i < len(s); i++ {
		if strings.Contains(n, s[i]) {
			count++
		}
	}
	if count == len(s) {
		return true
	}
	return false
}

func Solution(a []int) int {
	var count int
	for i := 1; i < len(a); i++ {
		if a[i-1] == a[i] {
			count++
		} else if a[i-1] > a[i] && a[i] > a[i+1] {
			count++
		}
	}
	return count / 2
}

func SockMerchant(n int32, ar []int32) int {
	var pair int
	sockArray := make(map[int32]int, len(ar))
	for i := 0; i < len(ar); i++ {
		if _, ok := sockArray[ar[i]]; !ok {
			sockArray[ar[i]] = i
		} else {
			for k := range sockArray {
				if ar[i] == k {
					pair++
					delete(sockArray, k)
				}
			}
		}
	}
	return pair
}

func countingValleys(n int32, s string) int32 {
	path := strings.Split(s, "")
	var level, valley int32
	for i := 0; i < len(path); i++ {
		if path[i] == "U" {
			level++
		} else if path[i] == "D" {
			level--
		}

		if level == 0 && path[i] == "U" {
			valley++
		}
	}
	return valley
}

func crawlSite() {
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("edmundmartin.com", "edmundmartin.com/writing-a-web-crawler-in-golang/"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://edmundmartin.com/writing-a-web-crawler-in-golang/")
}
func averageDistance(pointa, pointb, pointc []float32) float32 {
	xCordinate := pointa[0] + pointb[0] + pointc[0]
	yCordinate := pointa[1] + pointb[1] + pointc[1]

	return xCordinate + yCordinate/2
}

func MostPopularNumber(ar []int, l int) int {
	if len(ar) == 0 {
		return 0
	} else if len(ar) == 1 {
		return ar[0]
	}
	sockArray := make(map[int]int, len(ar))
	counter := 1
	max := 0
	var popular int
	for i := 0; i < len(ar); i++ {
		if _, ok := sockArray[ar[i]]; !ok {
			sockArray[ar[i]] = counter
		} else {
			for _, v := range sockArray {
				sockArray[ar[i]] = v + 1
			}
		}
	}
	for k, v := range sockArray {
		if v > max {
			max = v
			popular = k
		}
	}
	return popular
}

func NoOfOccurrence(arr []int, a int) int {
	var occur int
	for i := 0; i < len(arr); i++ {
		if arr[i] == a {
			occur++
		}
	}
	return occur
}
