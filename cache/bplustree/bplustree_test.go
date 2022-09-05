package bplustree

import (
	"pkgx/utils"
	"testing"
)

func TestNewWith(t *testing.T) {
	tree := NewWith[int](3, utils.IntComparator[int])
	tree.Put(1, 1)
	tree.Put(2, 2)
	tree.Put(3, 3)
	tree.Put(4, 4)
	tree.Put(6, 6)
	tree.Put(7, 7)
	tree.Put(8, 8)
	tree.Put(9, 9)
	tree.Put(10, 10)
}
