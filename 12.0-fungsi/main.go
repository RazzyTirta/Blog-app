package main

import (
	"fmt"
	"math"
	// "strings"
	// "math/rand"
	// "time"
)

// func main()  {
// 	names := []string{"razzi", "tirta"}
// 	printMessage("Hai", names)
// }

// func printMessage(message string, arr []string)  {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// fungsi dengan return value
// func main()  {
// 	rand.Seed(time.Now().Unix())

// 	var randomValue int
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number :", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number :", randomValue)
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number :", randomValue)
// }

// func randomWithRange(min, max int) int  {
// 	var value = rand.Int() % (max - min + 1) + min
// 	return value
// }

// func main()  {
// 	divideNumber(10, 2)
// 	divideNumber(2, 0)
// 	divideNumber(23, 2)
// }

// func divideNumber(m, n int)  {
// 	if n == 0 {
// 		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
// 		return
// 	}

// 	var res = m / n
// 	fmt.Printf("%d / %d = %d\n",m, n, res )
// }


// fungsi multiple return
func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran \t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran \t: %.2f \n", circumference)
}

// // predefined return value
// func calculate (d float64) (area float64, circumference float64)
func calculate(d float64) (area float64, circumference float64) {
	// hitung luas
	area = math.Pi * math.Pow(d /  2, 2)
	// var area = math.Pi * math.Pow(d /  2, 2)
	// hitung keliling
	circumference = math.Pi * d

	// kembalikan 2 nilai
	// var circumference = math.Pi * d
	// return area, circumference
	return
}