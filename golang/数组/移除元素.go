/*
从列表 nums 中移除值为 val 的元素，并返回新的列表长度。

参数：

	nums：一个整数列表。
	val：要移除的整数。

返回值：新的列表长度。

测试用例：
>>> nums = [3,2,2,3]
>>> removeElement(nums, 3)
2
>>> nums
[2, 2]
*/
package main

func removeElment(muns []int, val int) int {
	// 快慢指针，快指针遍历，慢指针记录
	n := len(muns)
	slow, fast := 0, 0
	for fast < n {
		if muns[fast] != val {
			muns[slow] = muns[fast]
			slow++
		}

		fast++
	}

	return slow
}

func removeElment2(muns []int, val int) int {
	// 双指针，左右指针
	n := len(muns)
	left, right := 0, n-1
	for left <= right {
		if muns[left] == val {
			muns[left] = muns[right]
			right--
		} else {
			left++
		}
	}

	return left
}

// func main() {
// 	nums := []int{3, 2, 2, 3}
// 	val := 2
// 	fmt.Println(removeElment(nums, val))
// 	fmt.Println(removeElment2(nums, val))
// }
