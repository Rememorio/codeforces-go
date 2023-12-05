package copypasta

import "sort"

type Interface[T any] interface {
	sort.Interface
	Push(x T)
	Pop() T
}

// Init 从第一个非叶子节点开始
func Init[T any](h Interface[T]) {
	// heapify
	n := h.Len()                    // 数组长度n
	for i := n/2 - 1; i >= 0; i-- { // n/2 - 1 是第一个非叶子节点
		down(h, i, n)
	}
}

// Push 插入新元素到堆中,自顶向下
// O(log n)
func Push[T any](h Interface[T], x T) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop 删除并返回堆顶元素，小根堆
// Pop 相当于 Remove(h, 0).
func Pop[T any](h Interface[T]) T {
	n := h.Len() - 1 // 最后一个元素下标
	h.Swap(0, n)     // 交换根元素和最后一个元素
	down(h, 0, n)
	return h.Pop()
}

func Remove[T any](h Interface[T], i int) any {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)        // 第i个放到末尾
		if !down(h, i, n) { // 如果向下失败，就向上
			up(h, i)
		}
	}
	return h.Pop()
}

// Fix 从执行位置向下，如果向下不成功就再向上
func Fix[T any](h Interface[T], i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}

// 上滤, 最后元素向上调整
// 与父节点比较，直到不能上移为止
func up[T any](h Interface[T], j int) {
	for {
		i := (j - 1) / 2             // parent
		if i == j || !h.Less(j, i) { // 父节点下标和子节点下标相同 | 父子不满足调整
			break // 这里是 j, i 是为了最小堆， 子<父
		}
		h.Swap(i, j)
		j = i // 交换后，从父节点继续向上
	}
}

// 下滤，根元素向下调整
// 父节点与左右子节点比较，与最大|最小的子节点交换，直到不能交换为止
// h 元素data
// i0 父节点
// n 存储数组的长度
func down[T any](h Interface[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1          // 左节点
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break // 判断左节点是否已超数组长度，或越界
		}
		j := j1                                     // left child 左节点赋值给j
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) { // 这里j2 < j1 为了最小堆
			//左节点右移一位取到右节点下标，小于长度，并且右节点和左节点满足交换条件，不管大顶堆，还是小顶堆
			j = j2 // = 2*i + 2  // right child  // 赋值右节点给j
		}
		if !h.Less(j, i) { // 不满足交换条件退出
			break
		}
		h.Swap(i, j) // 交换子节点和父节点
		i = j        // 子节点赋值为目标，继续往下走，
	}
	return i > i0
}
