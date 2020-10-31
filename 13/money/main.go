package main

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

//全局变量
var (
	coins = 5000
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	} //声明切片并初始化
	//users 用户

	distribution = make(map[string]int, len(users)) //声明map并初始化
	//distribution 分配
)

func dispatchCoin() (left int) {

	//1.依次拿到每个人的名字  遍历
	for _, name := range users {
		//2.遍历每个人名字中的字母
		for _, zimu := range name {
			//2.1拿到一个人的名字内的每个字母根据分金币的规则分金币
			switch zimu {
			//满足这个条件分1枚金币
			case 'e', 'E':
				distribution[name] += 1 //每个人分的金币保存到distribution中
				coins -= 1              //记录剩余的金币数
			//满足这个条件分2枚金币
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			//满足这个条件分3枚金币
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			//满足这个条件分4枚金币
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}

	}
	//3.整个第二部执行完成就能得到最终每个人分的金币数和剩余金币数
	left = coins
	return
}

func main() {
	left := dispatchCoin()
	for key, value := range distribution {
		//fmt.Println(key,value)
		fmt.Printf("%s:%d\n", key, value)
	}
	fmt.Println("剩余的金币数：", left)
}
