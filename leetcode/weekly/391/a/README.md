请看 [视频讲解](https://www.bilibili.com/video/BV1fq421A7CY/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def sumOfTheDigitsOfHarshadNumber(self, x: int) -> int:
        s = 0
        v = x
        while v:
            v, d = divmod(v, 10)
            s += d
        return -1 if x % s else s
```

```java [sol-Java]
class Solution {
    public int sumOfTheDigitsOfHarshadNumber(int x) {
        int s = 0;
        for (int v = x; v > 0; v /= 10) {
            s += v % 10;
        }
        return x % s > 0 ? -1 : s;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int sumOfTheDigitsOfHarshadNumber(int x) {
        int s = 0;
        for (int v = x; v; v /= 10) {
            s += v % 10;
        }
        return x % s ? -1 : s;
    }
};
```

```go [sol-Go]
func sumOfTheDigitsOfHarshadNumber(x int) int {
	s := 0
	for v := x; v > 0; v /= 10 {
		s += v % 10
	}
	if x%s == 0 {
		return s
	}
	return -1
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(\log x)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
