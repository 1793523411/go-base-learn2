package main

import (
	"fmt"
)

// func count(s string) map {
// 	ss := strings.Split(s, " ")
// 	m := make(map[string]int, len(ss))
// 	for _, v := range ss {
// 		_, ok := m[v]
// 		if !ok {
// 			m[v] = 1
// 		} else {
// 			m[v]++
// 		}
// 	}
// 	return m
// }

func main() {
	// map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用

	//!map基本使用
	// scoreMap := make(map[string]int, 8)
	// scoreMap["张三"] = 90
	// scoreMap["小明"] = 100
	// fmt.Println(scoreMap)
	// fmt.Println(scoreMap["小明"])
	// fmt.Printf("type of a:%T\n", scoreMap)

	// userInfo := map[string]string{
	// 	"username": "沙河小王子",
	// 	"password": "123456",
	// }
	// fmt.Println(userInfo) //

	//!判断某个键是否存在

	// scoreMap := make(map[string]int)
	// scoreMap["张三"] = 90
	// scoreMap["小明"] = 100
	// // 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	// v, ok := scoreMap["张三"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("查无此人")
	// }

	//!map的遍历

	// scoreMap := make(map[string]int)
	// scoreMap["张三"] = 90
	// scoreMap["小明"] = 100
	// scoreMap["娜扎"] = 60
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }
	// for k := range scoreMap {
	// 	fmt.Println(k)
	// }

	// // 遍历map时的元素顺序与添加键值对的顺序无关

	// //!使用delete()函数删除键值对

	// delete(scoreMap, "小明") //将小明:100从map中删除
	// for k, v := range scoreMap {
	// 	fmt.Println(k, v)
	// }

	//!按照指定顺序遍历map
	// rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	// var scoreMap = make(map[string]int, 200)

	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
	// 	value := rand.Intn(100)          //生成0~99的随机整数
	// 	scoreMap[key] = value
	// }
	// //取出map中的所有key存入切片keys
	// var keys = make([]string, 0, 200)
	// for key := range scoreMap {
	// 	keys = append(keys, key)
	// }
	// //对切片进行排序
	// sort.Strings(keys)
	// //按照排序后的key遍历map
	// for _, key := range keys {
	// 	fmt.Println(key, scoreMap[key])
	// }

	//!元素为map类型的切片
	// var mapSlice = make([]map[string]string, 3)
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }
	// fmt.Println("after init")
	// // 对切片中的map元素进行初始化
	// mapSlice[0] = make(map[string]string, 10)
	// mapSlice[0]["name"] = "小王子"
	// mapSlice[0]["password"] = "123456"
	// mapSlice[0]["address"] = "沙河"
	// for index, value := range mapSlice {
	// 	fmt.Printf("index:%d value:%v\n", index, value)
	// }

	//!值为切片类型的map
	// var sliceMap = make(map[string][]string, 3)
	// fmt.Println(sliceMap)
	// fmt.Println("after init")
	// key := "中国"
	// value, ok := sliceMap[key]
	// if !ok {
	// 	value = make([]string, 0, 2)
	// }
	// value = append(value, "北京", "上海")
	// sliceMap[key] = value
	// fmt.Println(sliceMap)

	//other
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])

	// count()
}
