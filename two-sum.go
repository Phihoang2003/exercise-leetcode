package main

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				break
			}
		}
	}
	return result

}
func main() {
	result := twoSum([]int{3, 2, 3}, 6)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}
