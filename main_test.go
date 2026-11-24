package main

import (
	"reflect"
	"testing"
)

func TestInorder(t *testing.T) {
	// Full tree:      4
	//               2   6
	//              1 3 5 7
	root := tn(4, tn(2, tn(1, nil, nil), tn(3, nil, nil)), tn(6, tn(5, nil, nil), tn(7, nil, nil)))
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if got := Inorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("Inorder = %v, want %v", got, want)
	}

	if got := Inorder(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("Inorder(nil) = %v, want []", got)
	}
	// Right-leaning chain.
	chain := tn(1, nil, tn(2, tn(3, nil, nil), nil))
	if got := Inorder(chain); !reflect.DeepEqual(got, []int{1, 3, 2}) {
		t.Errorf("Inorder(chain) = %v, want [1 3 2]", got)
	}
}
