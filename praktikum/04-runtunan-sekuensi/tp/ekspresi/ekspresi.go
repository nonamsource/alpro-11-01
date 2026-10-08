package main

import "fmt"

func main () {
	intNum := 5 
	intOther := 10 
	var sngNum float64 = -3.0
	fmt.Println(intOther + 2 * intNum != 30 || !(sngNum > 0))
	fmt.Println(4 / 2 == intOther / intNum )
	fmt.Println(intOther + 2 * intNum != 30 || !(sngNum > 0))
	fmt.Println(intNum > 0 || (sngNum <= 0 && intOther == 13))
}