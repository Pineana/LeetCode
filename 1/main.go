package main

func main() {
	
}

func twoSum(nums []int, target int) []int {
   mp:= make(map[int]int)
   for i := 0; i < len(nums); i++ {
		lk:= target-nums[i]
		if v,ok:=mp[lk];ok{
			return []int{i,v}
		}else{
			mp[nums[i]] = i
		}
   }
   return nil
}