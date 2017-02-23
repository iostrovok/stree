package main

import (
	"fmt"

	t "symmetrictree"
)

func main() {

	/*
				   				  	  - 1 -
				   		  	-  2 -			  - 2 -
				      -  3 -		  4    4	  		- 3 -
			  -  5 -	      -  6 - 		  		- 6 -       - 5 -
			- 7		 8		 9		10			  10 	9	  8		    7 -
		11															        11
	*/

	n10 := &t.Tree{10, nil, nil}
	n11 := &t.Tree{11, nil, nil}
	n4 := &t.Tree{4, nil, nil}
	n54 := &t.Tree{4, nil, nil}
	n58 := &t.Tree{8, nil, nil}
	n59 := &t.Tree{9, nil, nil}
	n60 := &t.Tree{10, nil, nil}
	n61 := &t.Tree{11, nil, nil}
	n8 := &t.Tree{8, nil, nil}
	n9 := &t.Tree{9, nil, nil}

	n57 := &t.Tree{7, n61, nil}
	n7 := &t.Tree{7, nil, n11}

	n56 := &t.Tree{6, n59, n60}
	n55 := &t.Tree{5, n57, n58}
	n53 := &t.Tree{3, n55, n56}
	n52 := &t.Tree{2, n53, n54}

	n6 := &t.Tree{6, n10, n9}
	n5 := &t.Tree{5, n8, n7}
	n3 := &t.Tree{3, n6, n5}
	n2 := &t.Tree{2, n4, n3}

	top := &t.Tree{1, n52, n2}

	left := top.CheckL([]int{})
	right := top.CheckR([]int{})

	for i := range left {
		fmt.Printf("%d. Left / Right ===> %d / %d\n", i, left[i], right[i])
	}

	fmt.Printf("CheckSimmtry: %t\n", top.CheckSimmtry())
}
