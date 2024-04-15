[视频讲解](https://www.bilibili.com/video/BV1bV4y1e72v/) 第三题。

## 提示 1

最终所有元素一定变成了一个在 $\textit{nums}$ 中的数。

枚举这个数。

## 提示 2

考虑把数字 $x$「扩散」到其它位置，那么每一秒 $x$ 都可以向左右扩散一位。

多个相同数字 $x$ 同时扩散，那么扩散完整个数组的耗时，就取决于相距**最远**的两个相邻的 $x$。

假设这两个 $x$ 的下标分别为 $i$ 和 $j$，且 $i<j$，那么耗时为：

$$
\left\lfloor\dfrac{j-i}{2}\right\rfloor
$$

枚举不同的 $x$，计算相应的耗时，更新答案的最小值。

## 提示 3

统计所有相同数字的下标，记到一个哈希表 $\textit{pos}$ 中。

设 $\textit{pos}[x]$ 列表第一个下标是 $p$，最后一个下标是 $q$。本题数组可以视作是**环形**的，所以 $p$ 和 $q$ 也是相邻的，耗时为 $\left\lfloor\dfrac{n-(q-p)}{2}\right\rfloor$。

也可以在 $\textit{pos}[x]$ 列表末尾添加一个 $p+n$，就可以转换成非环形数组处理了。

```py [sol-Python3]
class Solution:
    def minimumSeconds(self, nums: List[int]) -> int:
        pos = defaultdict(list)
        for i, x in enumerate(nums):
            pos[x].append(i)

        ans = n = len(nums)
        for a in pos.values():
            a.append(a[0] + n)
            mx = max(j - i for i, j in pairwise(a))
            ans = min(ans, mx)
        return ans // 2  # 最后再除 2
```

```java [sol-Java]
public class Solution {
    public int minimumSeconds(List<Integer> nums) {
        int n = nums.size();
        Map<Integer, List<Integer>> pos = new HashMap<>();
        for (int i = 0; i < n; i++) {
            pos.computeIfAbsent(nums.get(i), k -> new ArrayList<>()).add(i);
        }

        int ans = n;
        for (List<Integer> a : pos.values()) {
            int mx = n - a.get(a.size() - 1) + a.get(0);
            for (int i = 1; i < a.size(); i++) {
                mx = Math.max(mx, a.get(i) - a.get(i - 1));
            }
            ans = Math.min(ans, mx);
        }
        return ans / 2; // 最后再除 2
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimumSeconds(vector<int> &nums) {
        int n = nums.size();
        unordered_map<int, vector<int>> pos;
        for (int i = 0; i < n; i++) {
            pos[nums[i]].push_back(i);
        }

        int ans = n;
        for (auto &[_, a] : pos) {
            int mx = n - a.back() + a[0];
            for (int i = 1; i < a.size(); ++i) {
                mx = max(mx, a[i] - a[i - 1]);
            }
            ans = min(ans, mx);
        }
        return ans / 2; // 最后再除 2
    }
};
```

```go [sol-Go]
func minimumSeconds(nums []int) int {
	pos := map[int][]int{}
	for i, x := range nums {
		pos[x] = append(pos[x], i)
	}

	n := len(nums)
	ans := n
	for _, a := range pos {
		mx := n - a[len(a)-1] + a[0]
		for i := 1; i < len(a); i++ {
			mx = max(mx, a[i]-a[i-1])
		}
		ans = min(ans, mx)
	}
	return ans / 2 // 最后再除 2
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

[2023 下半年周赛题目总结](https://leetcode.cn/circle/discuss/lUu0KB/)
