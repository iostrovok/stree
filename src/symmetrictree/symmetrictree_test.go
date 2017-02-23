package symmetrictree

import (
	// "fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func makeTestData(add int) (*Tree, []int, []int) {
	/*
				   				  	  - 1 -
				   		  	- 52 -			  - 2 -
				      - 53 -		54		4	  - 3 -
			  - 55 -	      - 56 - 		  - 6 -       - 5 -
			57		58		59		60		10 		9	8		7
		61															11

		LEFT: 52, 53, 55, 57, 61, 59, 4, 6, 10, 8
		RIGHT: 2, 3, 5, 7, 11, 9, 54, 56, 60, 58
	*/
	left := []int{add + 2, add + 3, add + 5, add + 7, add + 11, add + 9, 4, 6, 10, 8}
	right := []int{2, 3, 5, 7, 11, 9, add + 4, add + 6, add + 10, add + 8}

	n10 := &Tree{10, nil, nil}
	n11 := &Tree{11, nil, nil}
	n4 := &Tree{4, nil, nil}
	n54 := &Tree{add + 4, nil, nil}
	n58 := &Tree{add + 8, nil, nil}
	n59 := &Tree{add + 9, nil, nil}
	n60 := &Tree{add + 10, nil, nil}
	n61 := &Tree{add + 11, nil, nil}
	n8 := &Tree{8, nil, nil}
	n9 := &Tree{9, nil, nil}

	n57 := &Tree{add + 7, n61, nil}
	n7 := &Tree{7, nil, n11}

	n56 := &Tree{add + 6, n59, n60}
	n55 := &Tree{add + 5, n57, n58}
	n53 := &Tree{add + 3, n55, n56}
	n52 := &Tree{add + 2, n53, n54}

	n6 := &Tree{6, n10, n9}
	n5 := &Tree{5, n8, n7}
	n3 := &Tree{3, n6, n5}
	n2 := &Tree{2, n4, n3}

	top := &Tree{1, n52, n2}

	return top, left, right
}

func (s *MySuite) Test_CheckLR_Empty(c *C) {

	// t.Skip("no now")

	top := &Tree{10, nil, nil}
	listL := top.CheckL([]int{})
	c.Assert(len(listL), Equals, 0)

	listR := top.CheckL([]int{})
	c.Assert(len(listR), Equals, 0)
}

func (s *MySuite) TTest_CheckLR_Simple(c *C) {

	// t.Skip("no now")

	n2l := &Tree{2, nil, nil}
	n2r := &Tree{2, nil, nil}
	top := &Tree{1, n2l, n2r}

	listL := top.CheckL([]int{})
	c.Assert(len(listL), Equals, 1)
	c.Assert(listL[0], Equals, 2)

	listR := top.CheckL([]int{})
	c.Assert(len(listR), Equals, 1)
	c.Assert(listR[0], Equals, 2)
}

func (s *MySuite) Test_CheckLR(c *C) {

	// t.Skip("no now")

	/*

					1
			2       		2
		3		4		4		3

	*/

	n4l := &Tree{4, nil, nil}
	n4r := &Tree{4, nil, nil}

	n3l := &Tree{3, nil, nil}
	n3r := &Tree{3, nil, nil}

	n2l := &Tree{2, n3l, n4l}
	n2r := &Tree{2, n4r, n3r}

	top := &Tree{1, n2l, n2r}
	//

	res := []int{2, 3, 4}

	listL := top.CheckL([]int{})
	c.Assert(len(listL), Equals, len(res))
	for i := range res {
		c.Assert(listL[i], Equals, res[i])
	}

	listR := top.CheckR([]int{})
	c.Assert(len(listR), Equals, len(res))
	for i := range res {
		c.Assert(listR[i], Equals, res[i])
	}

}

func (s *MySuite) Test_Check_LongTest(c *C) {

	// t.Skip("no now")

	top, left, right := makeTestData(50)

	stackLeft := top.CheckL([]int{})
	stackRight := top.CheckR([]int{})

	c.Assert(len(stackLeft), Equals, len(stackRight))
	c.Assert(len(stackLeft), Equals, len(left))
	c.Assert(len(right), Equals, len(stackRight))

	for i := range left {
		c.Assert(stackLeft[i], Equals, left[i])
	}

	for i := range right {
		c.Assert(right[i], Equals, stackRight[i])
	}
}

func (s *MySuite) Test_CheckSimmtry_Empty(c *C) {

	// t.Skip("no now")
	top := &Tree{10, nil, nil}
	c.Assert(top.CheckSimmtry(), Equals, true)

	top = &Tree{20, &Tree{}, nil}
	c.Assert(top.CheckSimmtry(), Equals, false)

	top = &Tree{20, nil, &Tree{}}
	c.Assert(top.CheckSimmtry(), Equals, false)
}

func (s *MySuite) Test_CheckSimmtry_Simple(c *C) {

	// t.Skip("no now")
	n2l := &Tree{2, nil, nil}
	n2r := &Tree{2, nil, nil}
	top := &Tree{1, n2l, n2r}

	c.Assert(top.CheckSimmtry(), Equals, true)

	top = &Tree{20, n2l, nil}
	c.Assert(top.CheckSimmtry(), Equals, false)

	top = &Tree{20, nil, n2r}
	c.Assert(top.CheckSimmtry(), Equals, false)
}

func (s *MySuite) Test_CheckSimmtry(c *C) {

	// t.Skip("no now")
	top, _, _ := makeTestData(0)
	c.Assert(top.CheckSimmtry(), Equals, true)
}
