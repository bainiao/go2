package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())
	// fmt.Println(example.PublicConst)
	// fmt.Println(example.GetPrivateConst())
	// pkg1.Printpkg1()
	// u := pkg1.User{Id: 1, Name: "Alice", Age: 20}
	// student := pkg1.Student{User: u, School: "XYZ University"}
	// pkg1.SayHello(student)
	// season := pkg1.Summer
	// fmt.Println("Current season is:", season.String())
	// example.Print("This is a message from main.")

	// dest := make([]int, 5)
	// src := []int{1, 2, 3, 4, 5}
	// copy(dest, src)
	// fmt.Println("dest:", dest)

	// var nums [5][5]int
	// for index, n := range nums {
	// 	fmt.Printf("index: %d, value: %v\n", index, n)
	// }

	// str := "Hello, 世界"
	// bytes := []byte(str)
	// // fmt.Println(bytes)
	// fmt.Println(string(bytes[12:13]))

	// str := "This is a string"
	// fmt.Println(&str)
	// bytes := []byte(str)
	// bytes = append(bytes, 96, 97, 98)
	// bytes[0] = []byte("t")[0]
	// str2 := str[:] // string(bytes)
	// fmt.Println(&str2)
	// fmt.Println(str)
	// fmt.Println(str2)
	// fmt.Println(reflect.TypeOf(str))
	// fmt.Println(reflect.TypeOf(str2))
	// fmt.Println(string(bytes))
	// copy value of a string to another string
	// s1 := "hello world"
	// b1 := unsafe.Slice(unsafe.StringData(s1), len(s1))
	// fmt.Printf("%p, %p\n", unsafe.StringData(s1), unsafe.SliceData(b1))
	// s1 := "hello world"
	// s2 := s1
	// fmt.Printf("s1: %p, s2: %p\n", unsafe.StringData(s1), unsafe.StringData(s2))
	// s2 = "hello go"
	// fmt.Printf("s1: %s, s2: %s\n", s1, s2)
	// str := "hello, 世界"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("byte %d, %x, %s\n", str[i], str[i], string(str[i]))
	// }
	// for _, value := range str {
	// 	fmt.Printf("%d, %x, %s\n", value, value, string(value))
	// }
	// runes := []rune(str)
	// for i := 0; i < len(runes); i++ {
	// 	fmt.Printf("rune %d, %x, %s\n", runes[i], runes[i], string(runes[i]))
	// }
	// mp := make(map[string]int, 5)
	// mp["one"] = 1
	// mp["two"] = 2
	// mp["three"] = 3
	// // for k, v := range mp {
	// // 	fmt.Printf("key: %s, value: %d\n", k, v)
	// // }
	// fmt.Println(len(mp))
	// fmt.Println(mp["four"])
	// f, exists := mp["four"]
	// if exists {
	// 	fmt.Println("four exists:", f)
	// } else {
	// 	fmt.Println("four does not exist")
	// }
	// set := make(map[int]struct{}, 10)
	// for i := 0; i < 10; i++ {
	// 	set[rand.IntN(100)] = struct{}{}
	// }
	// fmt.Println(set)
	// arr := make([]int, 5)
	// fmt.Println("main:", &arr[0])
	// printPtr(arr)
	// fmt.Println(arr[0])
	// people := []Person{
	// 	{Name: "Bob", Age: 25, Salary: 5000.0},
	// 	{Name: "Alice", Age: 26, Salary: 7000.0},
	// 	{Name: "Mike", Age: 23, Salary: 3000.0},
	// }
	// slices.SortFunc(people, func(p1 Person, p2 Person) int {
	// 	if p1.Name > p2.Name {
	// 		return 1
	// 	} else if p1.Name < p2.Name {
	// 		return -1
	// 	}
	// 	return 0
	// })
	// for i := 0; i < len(people); i++ {
	// 	fmt.Println(people[i].Name)
	// }

	// grow := Exp(2)
	// for i := range 10 {
	// 	fmt.Printf("2^%d = %d\n", i, grow())
	// }
	// fib := Fib(10)
	// for i, next := fib(); next; i, next = fib() {
	// 	fmt.Println(i)
	// }
	// deferDo()
	// company := Company[int, int]{Name: "Raysov", Id: 1, Stuff: []int{1, 2}}
	// fmt.Printf("Company Name: %s, Id: %d, Stuff Id: %v.\n", company.Name, company.Id, company.Stuff)
	// a := 1
	// ok, res := Assert[int](a)
	// if ok {
	// 	fmt.Println(res)
	// }
	// Do[int32](2)
	// Do[TinyInt](2)

	// bufferPool := NewPool(func() *bytes.Buffer {
	// 	return bytes.NewBuffer(nil)
	// })
	// for i := range 5 {
	// 	buffer := bufferPool.Get()
	// 	buffer.WriteString(strconv.Itoa(i) + "Hello, World !\n")
	// 	fmt.Println(buffer.String())
	// 	buffer.Reset()
	// 	bufferPool.Put(buffer)
	// }
	// next, stop := iter.Pull(Fibonacci(10))
	// defer stop()
	// for {
	// 	fibn, ok := next()
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(fibn)
	// }
	// s := []int{1, 2, 3, 4, 5}
	// for i, n := range slices.All(s) {
	// 	fmt.Printf("index: %d, value: %d,", i, n)
	// }
	// for n := range slices.Values(s) {
	// 	fmt.Printf("value: %d,", n)
	// }
	// slices.Reverse(s)
	// // slices.Chunk(s,3)
	// // s = slices.Clip(s)
	// m := map[string]int{"one":1, "two":2, "three":3}
	// keys := slices.Collect(maps.Keys(m))
	// // maps.Values(m)
	// // maps.Collect(m)
	// for _, k := range keys {
	// 	fmt.Printf("key: %s, value: %d\n", k, m[k])
	// }
	// ch := make(chan int, 5)
	// chW := make(chan struct{})
	// chR := make(chan struct{})
	// defer func() {
	// 	close(ch)
	// 	close(chW)
	// 	close(chR)
	// }()
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 		fmt.Println("write:", i)
	// 	}
	// 	chW <- struct{}{}
	// }()
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(time.Millisecond)
	// 		fmt.Println("read:", <-ch)
	// 	}
	// 	chR <- struct{}{}
	// }()
	// fmt.Println("write done:", <-chW)
	// fmt.Println("read done:", <-chR)
	// ch := make(chan struct{})
	// defer close(ch)
	// go func() {
	// 	fmt.Println("2")
	// 	ch <- struct{}{}
	// }()
	// fmt.Println("waiting for goroutine...")
	// <-ch
	// fmt.Println("goroutine finished.")
	// ch := make(chan int, 1)
	// // defer close(ch)
	// go func(chan int) {
	// 	ch <- 1
	// 	close(ch)
	// 	fmt.Println("chan closed")
	// }(ch)
	// time.Sleep(time.Millisecond * 10)
	// fmt.Println("begin to read chan")
	// fmt.Println(<-ch)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func(ctx context.Context) {
	// 	ticker := time.NewTimer(time.Second)
	// 	defer func() {
	// 		wg.Done()
	// 		ticker.Stop()
	// 	}()
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			return
	// 		case <-ticker.C:
	// 			fmt.Println("time out")
	// 			return
	// 		default:
	// 			fmt.Println(ctx.Value(1))
	// 		}
	// 		time.Sleep(time.Millisecond * 200)
	// 	}
	// }(context.WithValue(context.Background(), 1, 2))
	// wg.Wait()
	// nilchan := make(chan int, 1)
	// select {
	// case <-nilchan:
	// 	fmt.Println("read from nil chan")
	// case nilchan <- 1:
	// 	fmt.Println("write to nil chan")
	// case <-time.After(time.Second):
	// 	fmt.Println("timeout after 1 second")
	// }
	// d := time.Now().Add(2 * time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), d)
	// defer cancel()
	// select {
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("operation completed")
	// case <-ctx.Done():
	// 	fmt.Println("deadline exceeded:", ctx.Err())
	// }
	// l := findLongestTime([]rune{'a', 'a', 'a', 'b', 'b', 'b', 'c', 'd'}, 3)
	// fmt.Println(l)

	// writeOddEven(10)

	// arrays := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5},
	// 	{1, 2, 3, 4},
	// 	{10, 11},
	// 	{-1, 0},
	// }
	// fmt.Println(maxDistance(arrays))

	// ch := make(chan int, 2)
	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	// for n := range ch {
	// 	fmt.Println(n)
	// }
	// fmt.Println(isValid("()[]{}"))

	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

func findKthLargest1(nums []int, k int) int {
	res := make([]int, k)
	for i := 0; i < k; i++ {
		j := 0
		for ; j < i && res[j] < nums[i]; j++ {
		}
		res = append(res[:j], append([]int{nums[i]}, res[j:k-1]...)...)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > res[0] {
			j := 0
			for ; j < k && res[j] < nums[i]; j++ {
			}
			res = append(res[1:j], append([]int{nums[i]}, res[j:k]...)...)
		}
	}
	return res[0]
}

// type MinStack struct {
//     s []int;
// }

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// func Constructor() MinStack {
//     ms := MinStack{
//         s: make([]int, 0, 10),
//     };
// 	return ms
// }
// func (this *MinStack) Push(val int)  {
//     this.s = append(this.s, val);
// }

// func (this *MinStack) Pop()  {
//     this.s = this.s[:len(this.s)-1];
// }

// func (this *MinStack) Top() int {
//     return this.s[len(this.s)-1];
// }

//	func (this *MinStack) GetMin() int{
//	    if len(this.s) == 0 {
//	        return 0
//	    }
//	    min := this.s[0]
//	    for i:=0;i<len(this.s);i++{
//	        if this.s[i] < min {
//	            min = this.s[i]
//	        }
//	    }
//	    return min
//	}
func isValid(s string) bool {
	left := []rune{'(', '{', '['}
	rightMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0, len(s))
	for _, c := range s {
		isLeft := false
		for _, l := range left {
			if c == l {
				stack = append(stack, c)
				isLeft = true
				break
			}
		}
		if !isLeft {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != rightMap[c] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return true && len(stack) == 0
}
func orangesRotting(grid [][]int) int {
	res := -1
	rotten := make([][2]int, 0, 10)
	freshApple := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch grid[i][j] {
			case 2:
				rotten = append(rotten, [2]int{i, j})
			case 1:
				freshApple++
			}
		}
	}
	for len(rotten) > 0 {
		r := len(rotten)
		for i := 0; i < r; i++ {
			x, y := rotten[i][0], rotten[i][1]
			directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
			for _, d := range directions {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					freshApple--
					rotten = append(rotten, [2]int{nx, ny})
				}
			}
		}
		res++
		rotten = rotten[r:]
	}
	return res
}
func maxDistance(array [][]int) int {
	if len(array) == 0 {
		return 0
	}
	min := array[0][0]
	max := array[0][len(array[0])-1]
	maxD := 0
	for i := 1; i < len(array); i++ {
		if array[i][len(array[i])-1]-min >= max-array[i][0] {
			if array[i][len(array[i])-1]-min > maxD {
				max = array[i][len(array[i])-1]
			}
		} else {
			if max-array[i][0] > maxD {
				min = array[i][0]
			}
		}
	}
	return max - min
}
func writeOddEven(n int) {
	evench := make(chan int, n)
	defer close(evench)
	go func() {
		for i := 0; i < n; i += 2 {
			evench <- i
			time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		for i := 1; i < n; i += 2 {
			evench <- i
			time.Sleep(time.Millisecond)
		}
	}()
	for i := 0; i < n; i++ {
		fmt.Println(<-evench)
	}
}
func findLongestTime(letters []rune, n int) int {
	res := 0
	l_n := make(map[rune]int)
	longest := 0
	kind := 0
	for _, l := range letters {
		n, ok := l_n[l-'a']
		if !ok {
			kind++
			l_n[l-'a'] = 1
		} else {
			l_n[l-'a'] = n + 1
		}
		if l_n[l-'a'] > longest {
			longest = l_n[l-'a']
		}
	}
	if kind <= n {
		return (longest-1)*n + kind
	} else {
		for longest > 0 {
			for k, v := range l_n {
				res++
				v--
				if v == 0 {
					delete(l_n, k)
					kind--
				}
			}
			longest--
			if kind <= n {
				return res + (longest-1)*n + kind
			}
		}
	}
	return res
}
func Fibonacci(n int) func(yield func(int) bool) {
	a, b, c := 0, 1, 1
	return func(yield func(int) bool) {
		for range n {
			if !yield(a) {
				return
			}
			a, b = b, c
			c = a + b
		}
	}
}
func NewPool[T any](newFn func() T) *Pool[T] {
	return &Pool[T]{
		pool: &sync.Pool{
			New: func() interface{} {
				return newFn()
			},
		},
	}
}

type Pool[T any] struct {
	pool *sync.Pool
}

func (p *Pool[T]) Put(v T) {
	p.pool.Put(v)
}
func (p *Pool[T]) Get() T {
	var v T
	get := p.pool.Get()
	if get != nil {
		v, _ = get.(T)
	}
	return v
}

type Comparator[T any] func(a, b T) int
type BinaryHeap[T any] struct {
	s []T
	c Comparator[T]
}
type Queue[T any] []T

func (q *Queue[T]) Push(e T) {
	*q = append(*q, e)
}
func (q *Queue[T]) Size() int {
	return len(*q)
}
func (q *Queue[T]) Pop() (T, bool) {
	var res T
	ok := false
	if q.Size() > 0 {
		res = (*q)[0]
		*q = (*q)[1:]
		ok = true
	}
	return res, ok
}
func (q *Queue[T]) Peek() (_ T) {
	if q.Size() > 0 {
		return (*q)[0]
	}
	return
}

type TinyInt int8
type SignedInt interface {
	~int8 | int16 | int | int32 | int64
}
type UnSignedInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}
type Integer interface {
	SignedInt | UnSignedInt
}
type NilInteget interface {
	SignedInt
	UnSignedInt
}

func Do2[T NilInteget](n T) T {
	return n
}
func Do[T Integer](n T) T {
	return n
}
func Assert[T any](v any) (bool, T) {
	var res T
	if v == nil {
		return false, res
	}
	res, ok := v.(T)
	return ok, res
}

type Company[T, S int | string] struct {
	Name  string
	Id    T
	Stuff []S
}

func deferDo() {
	defer func() {
		fmt.Println("1")
	}()
	fmt.Println("2")
}

func Fib(n int) func() (int, bool) {
	a, b, c := 1, 1, 2
	i := 0
	return func() (int, bool) {
		if i >= n {
			return 0, false
		} else if i < 2 {
			temp := i
			i++
			return temp, true
		}
		a, b = b, c
		c = a + b
		i++
		return a, true
	}
}

func Exp(n int) func() int {
	e := 1
	return func() int {
		temp := e
		e *= n
		return temp
	}
}

type Person struct {
	Name   string
	Age    int
	Salary float64
}

func printPtr(a []int) {
	fmt.Println("inner func:", &a[0])
	a[0] = 1
}
