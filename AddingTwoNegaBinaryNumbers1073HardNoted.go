package main

import "fmt"

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
//public int[] addNegabinary(int[] arr1, int[] arr2) {
//int i = arr1.length - 1;
//int  j = arr2.length - 1;
//int carry = 0;
//Stack<Integer> stack = new Stack<>();
//
//while (i >= 0 || j >= 0 || carry != 0) {
//if (i >= 0) carry += arr1[i--];
//if (j >= 0) carry += arr2[j--];
//stack.push(carry & 1);
//carry = - (carry >> 1);
//
//}
//while (!stack.isEmpty() && stack.peek() == 0) {
//stack.pop();
//
//}
//int[] res = new int[stack.size()];
//int index = 0;
//if (stack.isEmpty()) return new int[]{0};
//while (!stack.isEmpty()) {
//res[index++] = stack.pop();
//}
//return res;
//}
