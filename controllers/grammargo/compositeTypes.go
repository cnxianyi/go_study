package grammargo

/* 复合数据类型

TODO 数组
	数组的长度是 固定 的 , 在编译时就已经被固定
	支持 下标访问 支持 len()
	默认情况下 数组的元素会被初始化为 零值
		var arr [3]int = [3]int{1, 2}
		println(arr[2], len(arr)) // 0 3

	可以使用 ... 来 自动创建值长度的数组
		arr := [...]int{1, 2, 3}
		println(len(arr)) // 3

	定义部分位置的数组
		arr := []int{50: 50, 99: 1}
		println(len(arr), arr[50]) // 100 50

TODO Slice 切片
	Slice是 对一个底层数组的 引用或视图
	Slice 由三个部分构成 指针、长度、容量
		指针: slice第一个元素对应的底层数组的地址
		长度: slice中元素的数目 *长度不能超过容量*
		容量: slice中开始位置 到 结束位置的位置

	Slice 的创建与数组类似,只是没有指定长度
		s := []int{1, 2}
		println(len(s), cap(s)) // 2 2

	Slice是底层数组的引用
		arr := [2]int{1, 2}
		s := arr[0:1]
		println(&s[0], &arr[0]) // 0xc0000bf788 0xc0000bf788 地址相同
		s[0] = 2
		println(s[0], arr[0]) // 2 2

	Slice 不支持 ==  Slice仅支持通过 == nil 来判断slice是否有底层数组
		arr := [2]int{1, 2}
		arr1 := [2]int{1, 2}
		println(arr == arr1) // true

		s1 := []int{1, 2}
		s2 := []int{1, 2}
		println(s1 == s2) // invalid operation: s1 == s2 (slice can only be compared to nil)

		println(s1 == nil) // false
		s2 = nil
		println(s2 == nil) // true

		println(len([]int{})) // 0
	判断slice是否为零值应该使用 len来判断长度

	make 函数初始化slice 语法 make([]T , len)
		创建的默认为零值
		s := make([]int, 3)
		println(len(s), s[0]) // 3 0

	append 方法用于向slice 添加元素
		s := make([]int, 3)
		s = append(s, 1)
		println(len(s), s[3]) // 4 1

TODO Map 哈希表
	无序的key/value 集合
	唯一key
	每个map都是 对一个哈希表的引用
	key 仅支持 支持 == 的类型作为key
	map可以通过测试key是否相等来判断是否已经存在

	make或字面量创建 Map
		m1 := make(map[string]int)
		m2 := map[string]int{}

	delete 删除
		m1 := map[string]int{"a": 1}
		delete(m1, "a")
		println(m1["a"]) // 0 零值

	无法 对map元素进行获取指针
		map 元素的地址不能被取出，因为 map 的底层实现允许其内部数据结构在插入、删除或更新元素时动态扩展和重定位。也就是说，map 中的元素可能在操作时发生移动，因此直接取地址是不安全的。
		m1 := map[string]int{"a": 1}
		b1 := &m1["a"] // delcared and not used

	map 的遍历顺序是不确定的
	map 的零值是 nil

	如何确定一个元素是否真的在map中,如这个值恰好是0
		m1 := map[string]int{"a": 0}
		if m, ok := m1["a"]; ok { // ok 的值是 map中是否真的有这个值
			println(m) // 0 只有真的存在才会执行,而不是返回零值
		}

	map 之前也不能全等比较. 只能通过循环来实现
	map 只能跟 nil ==

TODO 结构体
*/

func compositeType() {
	m1 := map[string]int{"a": 0}
	if m, ok := m1["a"]; ok { // ok 的值是 map中是否真的有这个值
		println(m) // 0 只有真的存在才会执行,而不是返回零值
	}

}
