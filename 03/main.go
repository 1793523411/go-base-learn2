package main

import "fmt"

func once(nums []int) int {
	eor := 0
	for i := 0; i < len(nums); i++ {
		eor ^= nums[i]
	}
	return eor
}

func main() {
	// fmt.Println('a' > 256)

	// & 	参与运算的两数各对应的二进位相与。
	// （两位均为1才为1）
	// | 	参与运算的两数各对应的二进位相或。
	// （两位有一个为1就为1）
	// ^ 	参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。
	// （两位不一样则为1）
	// << 	左移n位就是乘以2的n次方。
	// “a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。
	// >> 	右移n位就是除以2的n次方。
	// “a>>b”是把a的各二进位全部右移b位。
	//! the number of only once
	nums := []int{1, 1, 3, 3, 7, 4, 4, 5, 5}
	fmt.Println(once(nums))

}
