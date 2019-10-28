package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(kangaroo(2, 1, 1, 2))
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
