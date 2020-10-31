package leetcode

import (
	"sort"
	"strings"
)

// 两数之和
func TwoSum(nums []int, target int) []int {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//两数相加
// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret = new(ListNode)
	current := ret
	carry := 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return ret.Next
}

// 无重复字符的最长子串
func LengthOfLongestSubstring(s string) int {
	i := 0
	max := 0
	a := []rune(s)
	for m, c := range a {
		for n := i; n < m; n++ {
			if a[n] == c {
				i = n + 1
			}
		}
		if m-i+1 > max {
			max = m - i + 1
		}
	}
	return max
}

//寻找两个有序数组的中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1 := len(nums1)
	len2 := len(nums2)
	lenSum := len1 + len2
	if lenSum == 0 {
		return float64(0)
	}
	l, r := 0, 0
	a := make([]int, 0, lenSum)
	for l < len1 && r < len2 {
		if nums1[l] < nums2[r] {
			a = append(a, nums1[l])
			l++
		} else {
			a = append(a, nums2[r])
			r++
		}
	}
	a = append(a, nums1[l:]...)
	a = append(a, nums2[r:]...)
	if lenSum%2 != 0 {
		return float64(a[lenSum/2])
	} else {
		return (float64(a[lenSum/2-1]) + float64(a[lenSum/2])) / 2
	}
}

//整数反转
func Reverse(x int) int {
	var recv int
	for x != 0 {
		pop := x % 10
		x /= 10
		if recv*10 > 1<<31-1 || (recv*10 == 1<<31-1 && pop > 7) {
			return 0
		}
		if recv*10 < -1<<31 || (recv*10 == -1<<31 && pop < 8) {
			return 0
		}
		recv = recv*10 + pop
	}
	return recv
}

//回文数
func IsPalindrome(x int) bool {
	// 如果是负数或类似10、100、1000这种就直接返回false
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	// 反转后半部分的数字，和前半部分做比较
	var right int
	for x > right {
		right = right*10 + x%10
		x /= 10
	}
	// 注意前半部分和后半部分刚好相等（1221）或正好差一位（121）
	return x == right || x == right/10
}

//盛最多水的容器
func MaxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var ret int
	for i < j {
		var area int
		if height[i] > height[j] {
			area = (j - i) * height[j]
		} else {
			area = (j - i) * height[i]
		}
		if area > ret {
			ret = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ret
}

//罗马数字转整数
func RomanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	var ret int
	for i := 0; i < len(s); {
		// 先尝试读两个字符，注意索引不要越界
		if i+2 <= len(s) && m[s[i:i+2]] != 0 {
			ret += m[s[i:i+2]]
			i += 2
			continue
		} else {
			ret += m[s[i:i+1]]
			i++
			continue
		}
	}
	return ret
}

//最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = strs[0][0 : len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

//三数之和
func ThreeSum(nums []int) [][]int {
	lenNums := len(nums)
	ret := make([][]int, 0, 0)
	if lenNums < 3 {
		return ret
	}
	// 排序
	sort.Ints(nums)

	for i := 0; i < lenNums; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} // 去重
		l, r := i+1, lenNums-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				} // 左边去重
				for r < r && nums[r] == nums[r-1] {
					r--
				} // 后边去重
				l++
				r--
			}
			if sum > 0 {
				r--
			}
			if sum < 0 {
				l++
			}
		}
	}
	return ret
}

//two
func threeSum(nums []int) [][]int {
	var ret [][]int
	l := len(nums)
	if l < 3 {
		return ret
	}
	sort.Ints(nums) // 排序

	for k := 0; k < l; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] { // 去重
			continue
		}
		for i, j := k+1, l-1; i < j; {
			if j < l-1 && nums[j] == nums[j+1] { // 去重
				j--
				continue
			}
			sum := nums[k] + nums[i] + nums[j]
			if sum > 0 {
				j--
			} else if sum < 0 {
				i++
			} else {
				ret = append(ret, []int{nums[k], nums[i], nums[j]})
				i++
				j--
			}
		}
	}
	return ret
}

//有效的括号
func IsValid(s string) bool {
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	a := make([]byte, 0, len(s)/2)
	for _, b := range []byte(s) {
		if b == '(' || b == '{' || b == '[' {
			a = append(a, b)
			continue
		}
		if b == ')' || b == '}' || b == ']' {
			if len(a) > 0 && m[b] == a[len(a)-1] {
				a = a[:len(a)-1]
				continue
			} else {
				return false
			}
		}
	}
	if len(a) == 0 {
		return true
	} else {
		return false
	}
}

//合并两个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret = new(ListNode) // 前置虚拟节点法
	cur := ret              // 定义一个保存当前节点的变量
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next // 当前节点向后移
	}
	// 追加没遍历到的链表
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return ret.Next
}

//删除排序数组中的重复项
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret := 1
	for i := 0; i < len(nums)-1; {
		if nums[i+1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
		ret++
	}
	return ret
}

//移除元素
func RemoveElement(nums []int, val int) int {
	ret := 0
	for i := 0; i < len(nums); {
		if nums[i] != val {
			nums[ret] = nums[i]
			ret++
		}
		i++
	}
	return ret
}

//删除排序链表中的重复元素
func DeleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

//删除排序链表中的重复元素 II
func DeleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next != nil && head.Val == head.Next.Val {
		for head != nil && head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return DeleteDuplicates2(head.Next)
	} else {
		head.Next = DeleteDuplicates2(head.Next)
	}
	return head
}

//two
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head

	pre := dummy
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			for cur.Next != nil && cur.Val == cur.Next.Val {
				cur = cur.Next
			}
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}

//环形链表
func HasCycle(head *ListNode) bool {
	n1 := head // 慢指针一步一格
	n2 := head // 快指针一步两格
	if head == nil || head.Next == nil {
		return false
	}
	for n2 != nil && n2.Next != nil {
		n1 = n1.Next
		n2 = n2.Next.Next
		if n1 == n2 {
			return true
		}
	}
	return false
}

// 反转链表
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

//two
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
