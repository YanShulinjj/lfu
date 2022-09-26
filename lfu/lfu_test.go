/* ----------------------------------
*  @author suyame 2022-09-26 16:19:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/

package lfu

import (
	"fmt"
	"testing"
)

// func TestLFU(t *testing.T) {
// 	cache := NewLFU(2)
// 	_, ok := cache.Get("key1")
// 	if ok {
// 		t.Error("shoud not found!")
// 	}
// 	// put
// 	cache.Put("key1", "value1")
// 	// Again get
// 	v, ok := cache.Get("key1")
// 	if !ok {
// 		t.Error("shoud be found!")
// 	}
// 	if v != ValueType("value1") {
// 		t.Error(v)
// 	}
//
// 	// Put key2~6
// 	cache.Get("key1")
// 	cache.Get("key1")
// 	cache.Get("key1")
//
// 	cache.Put("key2", "value2")
// 	cache.Put("key3", "value3")
// 	cache.Put("key4", "value4")
// 	cache.Put("key5", "value5")
// 	cache.Put("key6", "value6")
// 	cache.Put("key2", "value2")
// 	cache.Put("key3", "value3")
// 	cache.Put("key4", "value4")
// 	cache.Put("key5", "value5")
// 	cache.Put("key6", "value6")
// 	// get key1
// 	v, ok = cache.Get("key1")
// 	if !ok {
// 		t.Error("shoud be found")
// 	}
// 	if v != ValueType("value1") {
// 		t.Error(v)
// 	}
// 	// get key2 should not found
// 	v, ok = cache.Get("key2")
// 	if ok {
// 		t.Error("shoud be not found")
// 	}
// }

var ops = []string{
	"LFUCache", "put", "put", "put", "put",
	"put", "get", "put", "get", "get", "put",
	"get", "put", "put", "put", "get", "put",
	"get", "get", "get", "get", "put", "put",
	"get", "get", "get", "put", "put", "get",
	"put", "get", "put", "get", "get", "get",
	"put", "put", "put", "get", "put", "get",
	"get", "put", "put", "get", "put", "put",
	"put", "put", "get", "put", "put", "get",
	"put", "put", "get", "put", "put", "put",
	"put", "put", "get", "put", "put", "get",
	"put", "get", "get", "get", "put", "get",
	"get", "put", "put", "put", "put", "get",
	"put", "put", "put", "put", "get", "get",
	"get", "put", "put", "put", "get", "put",
	"put", "put", "get", "put", "put", "put",
	"get", "get", "get", "put", "put", "put",
	"put", "get", "put", "put", "put", "put",
	"put", "put", "put",
}
var args = [][]int{
	{10}, {10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2},
	{3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8},
	{9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8},
	{2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12},
	{3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21},
	{5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5},
	{3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28},
	{1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17},
	{2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19},
	{4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24},
	{9, 26}, {13, 28}, {11, 26},
}

func TestS1(t *testing.T) {

	if len(ops) != len(args) {
		t.Fatal("ops don't match!")
	}
	var c *LFU
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		arg := args[i]
		switch op {
		case "LFUCache":
			c = NewLFU(arg[0])
		case "put":
			c.Put(KeyType(arg[0]), ValueType(arg[1]))
		case "get":
			v := c.Get(KeyType(arg[0]))
			fmt.Printf("%d: %d\n", arg[0], v)
		}
	}

	// c := NewLFU(2)
	// c.Put(1, 1)
	// c.Put(2, 2)
	// c.Get(2)
	// c.Get(1)
	// // if pre, tgt := c.Get(1), 1; pre != tgt {
	// // 	t.Error("not equal", pre, tgt)
	// // }
	// c.Put(3, 3)
	// if pre, tgt := c.Get(1), 1; pre != tgt {
	// 	t.Error("not equal", pre, tgt)
	// }
	// // if pre, tgt := c.Get(3), 3; pre != tgt {
	// // 	t.Error("not equal", pre, tgt)
	// // }
	// // c.Put(4, 4)
	// // if pre, tgt := c.Get(1), -1; pre != tgt {
	// // 	t.Error("not equal", pre, tgt)
	// // }
	// // if pre, tgt := c.Get(3), 3; pre != tgt {
	// // 	t.Error("not equal", pre, tgt)
	// // }
	// // if pre, tgt := c.Get(4), 4; pre != tgt {
	// // 	t.Error("not equal", pre, tgt)
	// // }
}
