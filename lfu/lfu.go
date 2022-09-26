/* ----------------------------------
*  @author suyame 2022-09-26 14:50:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package lfu

import (
	"container/list"
)

// implement of lfu (least frequency use)
// No concurrent safe!
// O(1)
// reference: https://blog.csdn.net/c183662101/article/details/106364749/

type KeyType int
type ValueType int

type LFU struct {
	capacity int
	items    map[KeyType]*list.Element
	freq     map[int]*list.Element
	head     list.List // save freq, elem seem as *HNode

}

type HNode struct {
	freq int
	ll   list.List // elem seem as *VNode
}

type VNode struct {
	father *HNode
	Key    KeyType
	Value  ValueType
}

func NewLFU(capacity int) *LFU {
	return &LFU{
		capacity: capacity,
		items:    make(map[KeyType]*list.Element, 0),
		freq:     make(map[int]*list.Element, 0),
	}
}

// Get a value from LFU cache
// ok = true if exists else false
func (l *LFU) Get(key KeyType) (v interface{}) {
	if v, ok := l.items[key]; !ok {
		return nil
	} else {
		// extract this v to freq+1 list
		l.visit(v)
		return v.Value.(*VNode).Value
	}
}

func (l *LFU) visit(v *list.Element) {
	hnode := v.Value.(*VNode).father
	hnode.ll.Remove(v)
	if _, ok := l.freq[hnode.freq+1]; !ok {
		// if freq + 1 not exist, create it!
		newnode := &HNode{freq: hnode.freq + 1}

		elem := l.head.InsertAfter(newnode, l.freq[hnode.freq])
		l.freq[hnode.freq+1] = elem
		// 是否删除频率低的hnode？

	}
	if hnode.ll.Len() == 0 {
		// if this hnode is null. must delete it
		l.head.Remove(l.freq[hnode.freq])
		delete(l.freq, hnode.freq)
	}
	// move v into new hnode
	hnode = l.freq[hnode.freq+1].Value.(*HNode)
	elem := hnode.ll.PushFront(v.Value)
	elem.Value.(*VNode).father = hnode
	// attention! need update items!
	l.items[v.Value.(*VNode).Key] = elem
}

// Put a kv into LFU cache
// it may occur replacement when capacity is full.
func (l *LFU) Put(key KeyType, value ValueType) {
	// 1. judge if is the key exists
	if v, ok := l.items[key]; ok {
		// visit it
		l.visit(v)
		// change its value
		v.Value.(*VNode).Value = value
		return
	}
	// 2. judge if its capacity is full
	if len(l.items) == l.capacity {
		// need to replace
		// delete min-freq-used first.
		// if min-freq-uesd has many, delete least recently used

		minfreqHNode := l.head.Front()
		hnode := minfreqHNode.Value.(*HNode)
		bck := hnode.ll.Back()

		deleteKey := bck.Value.(*VNode).Key
		hnode.ll.Remove(bck)
		// 是否删除频率低的hnode？
		if hnode.ll.Len() == 0 {
			l.head.Remove(l.freq[hnode.freq])
			delete(l.freq, hnode.freq)
		}

		delete(l.items, deleteKey)

	}
	// add new item
	// 1. find hnode
	if _, ok := l.freq[1]; !ok {
		newnode := &HNode{freq: 1}
		elem := l.head.PushFront(newnode)
		l.freq[1] = elem
	}
	// 2. insert vnode into hnode
	hnode := l.freq[1].Value.(*HNode)
	vnode := &VNode{father: hnode, Key: key, Value: value}
	velem := hnode.ll.PushFront(vnode)
	// 3. update items
	l.items[key] = velem

}
