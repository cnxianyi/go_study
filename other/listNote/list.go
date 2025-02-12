package listnote_test

import (
	"fmt"
)

/*
链表
*/

type ListNode struct {
	val  int       // 值
	next *ListNode // 下一个节点的地址
}

var son = ListNode{
	val:  2,
	next: nil,
}

// 初始化
var list = ListNode{
	val:  1,
	next: &son,
}

// 打印所有链表
func GetNextA(l *ListNode) *ListNode {
	fmt.Println(l.val)
	if l.next != nil {
		return GetNextA(l.next)
	}

	return &ListNode{}
}

// 尾部添加
func PushA(l *ListNode, v int) error {
	if l.next != nil {
		return PushA(l.next, v)
	} else {
		l.next = &ListNode{
			val: v,
		}
		return nil
	}
}

// 尾部删除
func PopA(l *ListNode) error {
	if l.next.next == nil {
		l.next = nil
		return nil
	} else {
		return PopA(l.next)
	}
}

// 头部添加
func UnshiftA(l *ListNode, v interface{}) error {

	switch v.(type) {
	case int:

		h := *l

		l.val = v.(int)
		l.next = &h

		return nil

	case []int:
		vd := v.([]int)
		// 创建第一个
		head := &ListNode{
			val: vd[0],
		}

		// 保存原链表头
		oldHead := *l

		// 初始化 cur
		cur := head

		for i := 1; i < len(vd); i++ {
			cur.next = &ListNode{
				val: vd[i],
			}
			cur = cur.next
		}

		// 追加原链表
		cur.next = &oldHead

		// 更新传入的链表
		*l = *head

		return nil

	default:
		return fmt.Errorf("未知类型")
	}
}

// 头部删除
func ShiftA(l *ListNode) {
	*l = *l.next
}

func InitA() {
	GetNextA(&list)
	PushA(&list, 3)
	GetNextA(&list)
	PopA(&list)
	PopA(&list)

	GetNextA(&list)
	UnshiftA(&list, []int{2, 3})
	UnshiftA(&list, 4)
	GetNextA(&list)
	ShiftA(&list)
	GetNextA(&list)
}
