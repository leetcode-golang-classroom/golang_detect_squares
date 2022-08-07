# golang_detect_squares

You are given a stream of points on the X-Y plane. Design an algorithm that:

- **Adds** new points from the stream into a data structure. **Duplicate** points are allowed and should be treated as different points.
- Given a query point, **counts** the number of ways to choose three points from the data structure such that the three points and the query point form an **axis-aligned square** with **positive area**.

An **axis-aligned square** is a square whose edges are all the same length and are either parallel or perpendicular to the x-axis and y-axis.

Implement the `DetectSquares` class:

- `DetectSquares()` Initializes the object with an empty data structure.
- `void add(int[] point)` Adds a new point `point = [x, y]` to the data structure.
- `int count(int[] point)` Counts the number of ways to form **axis-aligned squares** with point `point = [x, y]` as described above.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/09/01/image.png](https://assets.leetcode.com/uploads/2021/09/01/image.png)

```
Input
["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
[[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
Output
[null, null, null, null, 1, 0, null, 2]

Explanation
DetectSquares detectSquares = new DetectSquares();
detectSquares.add([3, 10]);
detectSquares.add([11, 2]);
detectSquares.add([3, 2]);
detectSquares.count([11, 10]); // return 1. You can choose:
                               //   - The first, second, and third points
detectSquares.count([14, 8]);  // return 0. The query point cannot form a square with any points in the data structure.
detectSquares.add([11, 2]);    // Adding duplicate points is allowed.
detectSquares.count([11, 10]); // return 2. You can choose:
                               //   - The first, second, and third points
                               //   - The first, third, and fourth points

```

**Constraints:**

- `point.length == 2`
- `0 <= x, y <= 1000`
- At most `3000` calls **in total** will be made to `add` and `count`.

## 解析

給定一序列 x-y 座標上的點 point 

實做出一個資料結構 DetectSquare 具有以下 method

1. 建構子： 用來初始化整個結構
2. void add(int[] point): 用來新增 座標到此資料結構並且儲存起來
3. int count(int[] point): 用來計算以一個 point 為 square 其中一點 在資料結構中所目前所儲存的座標點選出3點可與 point 形成 square 的方法數

首先這個資料結構要能儲存 add 的座標點

所以可以用類似 list 的方式來做處理

然而 我們思考一下

假設給定一個點 p , 可以形成 square 的有 p1, p2, p3, 出現的次數分別是 2, 3, 1

如下圖

![](https://i.imgur.com/ixCwAet.png)

則其可能方法數 = p1 出現個數 * p2 出現個數 * p3  出現個數

所以透過這個方式，所以知道透過 HashMap 來儲存每個座標點以其出現的個數是最有效率的作法

而剩下的問題是 當固定一個點時 p = {px, py}， 如何去找出其他3個點

但由於是 square 所以已知長與寬是固定的

所以我們只要找到對角線的那個點 另外兩個點就固定住了

![](https://i.imgur.com/FW7BQ1X.png)


因此問題可以簡化為如何檢驗目前的點 c = {cx, cy} 是對角線的點

如果是 square 代表 abs(px - cx) == abs(py -cy) 

且因為 square 面積需要大於 0 所以 px ≠ cx && py ≠ cy

當確認 c 可以是對角線

接著就可以確認兩個 點 cp1(px, cy), cp2(cx, py)有在儲存的 hash 內

![](https://i.imgur.com/Jgc2tYm.png)

然後把3個 hash值相乘所累加就是所求

這樣每次只需要 O(n) 就可以做 count 的運算

每次需要 O(n) 存儲 n 個 input

## 程式碼
```go
package sol

type Point struct {
	x, y int
}
type DetectSquares struct {
	hash map[Point]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		hash: make(map[Point]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	p := Point{x: point[0], y: point[1]}
	(*this).hash[p] += 1
}

func (this *DetectSquares) Count(point []int) int {
	var abs = func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}
	p := Point{x: point[0], y: point[1]}
	// detect all possible diagonal
	res := 0
	for coord, count := range (*this).hash {
		if abs(coord.x-p.x) != abs(coord.y-p.y) || (coord.x == p.x && coord.y == p.y) {
			continue
		}
		res += count * (*this).hash[Point{x: coord.x, y: p.y}] * (*this).hash[Point{x: p.x, y: coord.y}]
	}
	return res
}

```
## 困難點

1. 要知曉如何計算可形成 square 的方法數
2. 需要知道透過對角線來判斷 square 的方法

## Solve Point

- [x]  建立一個 Point 的資料結構 裏面有個兩個 field x, y 用來儲存座標
- [x]  建立一個叫作 hash 的 HashMap 把 Point 當作 key , int 當作 value 用來儲存每個 point 出現的次數
- [x]  DetectSquare 建構子先初始化一個 DetectSquare 物件裏面初始化一個 叫作 hash 的 HashMap 把 Point 當作 key , int 當作 value
- [x]  add method 把 point []int 當作輸入參數 每次把 point 加入這個 hashMap 把 hash[Point{x;point[0], y: point[1]}] += 1
- [x]  count method 把 point []int 當作輸入參數, 每次對所有 hashMap 的所有 key 檢查 是否 key 能成為 point 的對角線 point 如果否則往下一個 key 找 如果是 則找出對應的另外兩點做相乘累加 , 最後回傳累加值