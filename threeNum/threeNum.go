package main

import "fmt"

/*
 *@Author: Ryan_Z
 *@Date: 2022-06-11 22:23:50
 *@Description: 查找一个数组中最接近目标数的三个数，并求和
 *@Think:
 */
/*

func threeSumClosest(nums []int, target int) int {
	AbsStepMap := make(map[int]int)

	for k, v := range nums {
		step := AbsStep(v, target)
		AbsStepMap[k] = step
	}
	ret := sortMap(AbsStepMap)
	sum := 0
	for _, v := range ret {
		sum += nums[v]
	}
	return sum
}*/
/*
func AbsStep(num, t int) int {
	step := t - num
	if step > 0 {
		return step
	}
	return -step
}

func sortMap(asm map[int]int) []int {
	var mapIndexSlice []int
	for k, _ := range asm {
		mapIndexSlice = append(mapIndexSlice, k)
	}
	length := len(mapIndexSlice)
	for i := 0; i < 3; i++ {
		tampMin := asm[mapIndexSlice[i]]
		tampMinIndex := i
		for j := i; j < length; j++ {
			if asm[mapIndexSlice[j]] < tampMin {
				tampMin = asm[mapIndexSlice[j]]
				tampMinIndex = j
			}
		}
		if tampMinIndex != i {
			mapIndexSlice[i], mapIndexSlice[tampMinIndex] = mapIndexSlice[tampMinIndex], mapIndexSlice[i]
		}

	}
	return mapIndexSlice[0:3]
}*/
func AbsStep(num, t int) int {
	step := t - num
	if step > 0 {
		return step
	}
	return -step
}
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	abs := AbsStep((nums[0] + nums[1] + nums[length-1]), target)
	var indexNums = [3]int{0, 1, length - 1}
	// fmt.Printf("indexNums: %v\n", indexNums)
	for i := 0; i < length-2; i++ {
		twoNum := target - nums[i]
		for k := length - 1; k > i+1; k-- {
			for j := i + 1; j < k; j++ {
				leftNum := nums[j] + nums[k]
				tampabs := AbsStep(leftNum, twoNum)
				if tampabs < abs {
					// fmt.Printf("tampabs: %v\n", tampabs)
					abs = tampabs
					indexNums[0], indexNums[1], indexNums[2] = i, j, k
					// fmt.Printf("indexNums: %v\n", indexNums)
				}
			}
		}

	}
	sum := 0
	// fmt.Printf("indexNums: %v\n", indexNums)
	for _, v := range indexNums {
		sum += nums[v]
	}
	return sum

}
func main() {
	nums := []int{-1, 2, 1, -3}
	target := 1
	sum := threeSumClosest(nums[:], target)
	fmt.Printf("sun: %v\n", sum)
	// fmt.Printf("nums[index]: %v\n", nums[index])

}
