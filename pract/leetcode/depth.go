package leetcode

import (
	"strings"
)

func maxDepthAfterSplit(seq string) []int {
	// 输入：seq = "(()())"
	//输出：[0,1,1,1,1,0]有效括号字符串 定义：对于每个左括号，都能找到与之对应的右括号，反之亦然。详情参见题末「有效括号字符串」部分。
	//嵌套深度 depth 定义：即有效括号字符串嵌套的层数，depth(A) 表示有效括号字符串 A 的嵌套深度。详情参见题末「嵌套深度」部分。
	//给你一个「有效括号字符串」 seq，请你将其分成两个不相交的有效括号字符串，A 和 B，并使这两个字符串的深度最小。
	//不相交：每个 seq[i] 只能分给 A 和 B 二者中的一个，不能既属于 A 也属于 B 。
	//A 或 B 中的元素在原字符串中可以不连续。
	//A.length + B.length = seq.length
	//max(depth(A), depth(B)) 的可能取值最小。
	//划分方案用一个长度为 seq.length 的答案数组 answer 表示，编码规则如下：
	//
	//answer[i] = 0，seq[i] 分给 A 。
	//answer[i] = 1，seq[i] 分给 B 。
	//如果存在多个满足要求的答案，只需返回其中任意 一个 即可。

	/*我们保证栈内一半的括号属于序列 A，一半的括号属于序列 B，那么就能保证拆分后最大的嵌套深度最小，
	是当前最大嵌套深度的一半。要实现这样的对半分配，
	我们只需要把奇数层的 ( 分配给 A，偶数层的 ( 分配给 B 即可。对于上面的例子，我们将嵌套深度为 1 和 3 的所有括号 (())
	分配给 A，嵌套深度为 2 的所有括号 ()()() 分配给 B。
	*/
	sli := make([]int, len(seq))
	strList := strings.Split(seq, "")
	i := 0
	for j, v:= range strList{
		if v == "("{
			i ++
			sli[j] = i%2
		}
		if v == ")"{
			sli[j] = i%2
			i--
		}
	}
	return sli
}

// 两数之和
func twoSum(num []int, target int) []int{
	numMap := make(map[int] int, len(num))
	for i, v := range num{
		if k, ok := numMap[target-v]; ok{
			return []int{k, i}
		}
		numMap[v] = i
	}
	return nil
}

// 生命游戏
func gameOfLife(board [][]int) {
	// leetcode 编译器有问题 该方法得不到正确答案
	for i, _ := range board{
		board[i] = append([]int{0}, board[i]...) // 头插
		board[i] = append(board[i], 0)  // 尾插
	}
	a := make([][]int, 1)
	for i, _ := range a{
		a[i] = make([]int, len(board[0]))
	}
	board = append(a, board...)
	board = append(board, a...)
	reallist := make([][]int, len(board)-2)
	onecount := func(a []int) int {
		count := 0
		for _, v := range a{
			if v == 1{
				count++
			}
		}
		return count
	}
	for i:= 1; i<=len(board)-2; i++ {
		reallist[i-1] = make([]int, len(board[0])-2)
		for j:= 1; j<=len(board[0])-2;j++{
			count := 0
			up := onecount(board[i-1][j-1:j+2])
			down := onecount(board[i+1][j-1:j+2])
			if board[i][j-1] == 1{
				count += 1
			}
			if board[i][j+1] == 1{
				count += 1
			}
			sum := up+down+count
			if sum<2 && board[i][j] == 1{
				reallist[i-1][j-1] = 0
			}
			if sum== 2 || sum==3{
				if board[i][j] == 1{
					reallist[i-1][j-1] = 1
				}
			}
			if sum>3 && board[i][j] == 1{
				reallist[i-1][j-1] = 0
			}
			if sum==3 && board[i][j] == 0{
				reallist[i-1][j-1] = 1
			}
		}
	}
    board = reallist
}
func GameOfLife(board [][]int)  {
	realList := make([][]int, len(board))
	for i := 0;i < len(board);i++ {
		realList[i] = make([]int, len(board[i]))
	}
	for i := 0;i < len(board);i++ {
		for j := 0;j < len(board[i]);j++ {
			nums := 0
			// 上面
			if i - 1 >= 0 {
				nums += board[i-1][j]
			}
			// 左面
			if j - 1 >= 0 {
				nums += board[i][j-1]
			}
			// 下面
			if i + 1 < len(board) {
				nums += board[i+1][j]
			}
			// 右面
			if j + 1 < len(board[i]) {
				nums += board[i][j+1]
			}
			// 左上
			if i - 1 >= 0 && j - 1 >= 0 {
				nums += board[i-1][j-1]
			}
			// 右上
			if i - 1 >= 0 && j + 1 < len(board[i]) {
				nums += board[i-1][j+1]
			}
			// 左下
			if i + 1 < len(board) && j - 1 >= 0  {
				nums += board[i+1][j-1]
			}
			// 右下
			if j + 1 < len(board[i]) && i + 1 < len(board) {
				nums += board[i+1][j+1]
			}
			realList[i][j] = board[i][j]
			switch {
			case nums < 2: realList[i][j] = 0
			case nums == 3 && realList[i][j] == 0: realList[i][j] = 1
			case nums > 3: realList[i][j] = 0
			}
		}
	}
	copy(board, realList)
}

