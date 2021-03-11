package main

import "fmt"
//https://leetcode.com/problems/adding-two-negabinary-numbers/discuss/303842/C%2B%2B-O(N)-Time-with-Explanation-(carry-(sum-greatergreater-1))
func addNegabinary(arr1 []int, arr2 []int) []int {
	i := len(arr1) - 1
	j := len(arr2) - 1
	carry := 0
	stack := make([]int, 0, len(arr1))
	for i >=0 || j >=0 || carry != 0 {
		if i >= 0 {
			carry += arr1[i]
			i--
		}
		if j >=0 {
			carry += arr2[j]
			j--
		}
		stack = append(stack, carry & 1)
		carry = - (carry >> 1)
	}
	t := len(stack) - 1
	for ;t >= 0 && stack[t] == 0; t-- {
	}
	stack = stack[0:t + 1]
	if len(stack) == 0 {
		return []int{0}
	}
	for start, end := 0, len(stack) - 1; start < end; start,end = start+1, end-1 {
		temp := stack[end]
		stack[end] = stack[start]
		stack[start] = temp
	}
	return stack
}

func main() {
	fmt.Printf("%+v", addNegabinary([]int{0}, []int{1}))
}
/*
   The remainder operator can be used with negative integers. The rule is:

   Perform the operation as if both operands were positive.
   If the left operand is negative, then make the result negative.
   If the left operand is positive, then make the result positive.
   Ignore the sign of the right operand in all cases.
   For example:

   17 %  3 == 2     -17 %  3 == -2
   17 % -3 == 2     -17 % -3 == -2
*/
//private String baseNegative(int n, int base) {
//StringBuilder sb = new StringBuilder();
//while (n != 0) {
//int remainder = n % base;
//n = n / base;
////This is the only difference compared to regular base conversion
//if (remainder < 0) {
//remainder += -base;
//n += 1;
//}
//sb.append(remainder);
//}
//return sb.length() == 0 ? "0" : sb.reverse().toString();
//}
