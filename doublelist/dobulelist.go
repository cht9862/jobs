package doublelist

import (
	"errors"
	"fmt"
)



// 双向链表
// MyData 泛型存储数据容器
type MyData interface{}

//List 定义节点
type List struct {
	title  bool             // 该节点是否为链表头节点
	strId  string           // 节点唯一表示
	data   MyData           // 数据
	next   *List            // 指向下一个节点
	prev   *List            // 指向上一个节点
	unique map[string]*List // 检测唯一id
}


// NewHeadNode 创建一个链表对象
func NewHeadNode() *List {
	// 创建返回链表对象头部实例，非常重要，相当于链表标识
	res := &List{
		title:  true,
		unique: make(map[string]*List),
	}
	res.next = res
	return res
}

// NewList 创建一个新节点
func (l *List) NewList(key string, d MyData) (*List, error) {
	// 判断该key 是否已经存在l 链表中，若存在返回nil, error
	if _, ok := l.unique[key]; ok {
		return nil, errors.New("【该节点已存在】")
	}
	// 创建返回List 指针实例
	return &List{
		strId: key,
		data:  d,
	}, nil
}

func CheckNodeUnique(l, a *List) bool {
	if _, ok := l.unique[a.strId]; ok {
		return false
	}
	return true
}


//InsertNode 向zero链表中的s节点后插入insert节点
func InsertNode(zero *List, s string, insert *List) bool {
	if !CheckNodeUnique(zero, insert) {
		fmt.Println("[该节点已存在]")
		return false
	}
	if zero.strId == s {
		// 将插入节点的key 注册到zero 的strId中，用于确保节点唯一性
		zero.unique[insert.strId] = insert
		insert.next = zero.next
		insert.prev = zero
		zero.next.prev = insert
		zero.next = insert
		return true
	}
	if zero.next == nil {
		return false // 没有符合s 条件的节点，无法插入
	}
	InsertNode(zero.next, s, insert)
	return true
}

//GetNode 查询当前zero链表中是否有key节点，若有返回出来，没有返回错误
func GetNode(zero *List, key string) (*List, error) {
	if zero.next == nil {
		if zero.strId == key {
			return zero, nil
		}
		return nil, fmt.Errorf("【未查询到指定节点】")
	} else {
		if zero.strId == key {
			return zero, nil
		}
		res, _ := GetNode(zero.next, key)
		if res.strId == key {
			return res, nil
		}
	}
	return nil, fmt.Errorf("~【未查询到指定节点】")
}

//AppendNode 在zero链表结尾加上a节点
func AppendNode(zero *List, a *List) {
	if !CheckNodeUnique(zero, a) {
		fmt.Println("[该节点已存在]")
		return
	}
	if zero.title == true {
		a.next = zero
		zero.prev = a
		zero.unique[a.strId] = a
	}
	if zero.next.title == true {
		zero.next = a
		a.prev = zero
		return
	}
	AppendNode(zero.next, a)
}

// DeleteNode 在zero链表中删除s这个节点
func DeleteNode(zero *List, s string) {
	// 将链表中该节点的前后关系修改，此节点无关联后被遗弃后会被gc 回收
	if zero.strId == s {
		zero.prev.next = zero.next
		zero.next.prev = zero.prev
		return
	}
	DeleteNode(zero.next, s)
}

// 实现一个链式栈(底层使用存储使用链表实现)
type Stack struct {
	store *List
}

func (s *Stack)Push(key string,data MyData) {
	res ,err := s.store.NewList(key,data)
	if err != nil {
		fmt.Println(err)
	}
	AppendNode(s.store,res)
}
func (s *Stack)Pop() MyData {
	defer func() {
		DeleteNode(s.store,s.store.prev.strId)
	}()
	return s.store.prev
}

func NewStack() *Stack{
	res := NewHeadNode()
	return &Stack {
		store: res,
	}
}



