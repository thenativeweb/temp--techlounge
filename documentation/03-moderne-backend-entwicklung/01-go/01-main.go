package main

import "fmt"

func main1() {
	const hallo = "Hallo "

	fmt.Println(hallo + "Welt!")
	fmt.Println(23.0 / 42)
	fmt.Println(true && false)

	// Zählschleife
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// While-Schleife
	j := 10
	for j > -5 {
		j -= 2
		fmt.Println(j)
	}

	// Do-While-Schleife
	k := 10
	for {
		k -= 2
		if k < -5 {
			break
		} else {
			fmt.Println(k)
		}
	}

	// Endlosschleife
	for {
		fmt.Println("Endlosschleife")
	}

	// useHttp := true
	// port := 80
	// if !useHttp {
	// 	port = 443
	// }

	// l := 23
	// switch l {
	// case 1:
	// 	// ...
	// 	fallthrough
	// case 2:
	// 	// ...
	// default:
	// 	// ...
	// }
}
