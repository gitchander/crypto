package base26

import (
	"fmt"
)

const mask4bitValueMax = 9

func testBase26Values() {
	// 4 bits value <= 9

	// with 5 bit is 0:
	fmt.Println(0) // 0_0000
	fmt.Println(1) // 0_0001
	fmt.Println(2) // 0_0010
	fmt.Println(3) // 0_0011
	// ...
	fmt.Println(8) // 0_1000
	fmt.Println(9) // 0_1001

	// 4 bits value > 9
	fmt.Println(10) // 1010
	fmt.Println(11) // 1011
	fmt.Println(12) // 1100
	fmt.Println(13) // 1101
	fmt.Println(14) // 1110
	fmt.Println(15) // 1111

	// with 5 bit is 1:
	fmt.Println((1 << 4) + 0) // 1_0000: 16
	fmt.Println((1 << 4) + 1) // 1_0001: 17
	fmt.Println((1 << 4) + 2) // 1_0010: 18
	fmt.Println((1 << 4) + 3) // 1_0011: 19
	// ...
	fmt.Println((1 << 4) + 8) // 1_1000: 24
	fmt.Println((1 << 4) + 9) // 1_1001: 25
}
