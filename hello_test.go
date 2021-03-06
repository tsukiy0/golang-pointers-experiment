package goptr

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HelloSuite struct {
	suite.Suite
}

func (s *HelloSuite) TestSum() {
	s.Equal(3, Sum(1, 2))
}

func (s *HelloSuite) TestGoToPointerValue() {
	a := "something"
	aPtr := &a

	s.Assert().Equal("something", *aPtr)
	s.Assert().Equal(a, *aPtr)
	s.Assert().Equal(&a, aPtr)
}

func (s *HelloSuite) TestModifyReference() {
	modify := func(s *string) {
		*s = "something else"
	}

	a := "something"
	modify(&a)

	s.Assert().Equal("something else", a)
}

func (s *HelloSuite) TestModifyCopy() {
	modify := func(s string) {
		s = "something else"
	}

	a := "something"
	modify(a)

	s.Assert().Equal("something", a)
}

func (s *HelloSuite) TestModifySliceReference() {
	modify := func(s []string) {
		s[0] = "something else"
	}

	a := []string{"something"}
	modify(a)

	s.Assert().Equal("something else", a[0])
}

func (s *HelloSuite) TestModifyStruct() {
	type Person struct {
		Name string
	}

	p := Person{Name: "John"}
	pPtr := &p
	p.Name = "Bob"

	s.Assert().Equal("Bob", p.Name)
	s.Assert().Equal("Bob", pPtr.Name)

	pPtr.Name = "James"

	s.Assert().Equal("James", p.Name)

	modifyCopy := func(pp Person) {
		pp.Name = "Chris"
	}
	modifyCopy(p)

	s.Assert().NotEqual("Chris", p.Name)

	modifyPtr := func(pp *Person) {
		pp.Name = "Chris"
	}
	modifyPtr(pPtr)

	s.Assert().Equal("Chris", p.Name)

	p2 := *pPtr

	s.Assert().Equal("Chris", p2.Name)

	p2.Name = "Luke"

	s.Assert().Equal("Chris", p.Name)
	s.Assert().Equal("Luke", p2.Name)
}

func (s *HelloSuite) TestUninitialisedSlicesAreNil() {
	var a []string

	s.Assert().Nil(a)
}

func (s *HelloSuite) TestModifySliceByIndex() {
	a := []int{1, 1}

	modify := func(i []int) {
		i[0] = 2
		i[1] = 2
		s.Assert().Equal(a, i)
	}

	modify(a)

	s.Assert().Equal([]int{2, 2}, a)
}

func (s *HelloSuite) TestAppendToSliceAfter() {
	a := []int{1, 1}

	modify := func(i []int) {
		i[0] = 2
		i[1] = 2
		s.Assert().Equal(a, i)
		i = append(i, 3)
		s.Assert().NotEqual(a, i)
	}

	modify(a)

	s.Assert().Equal([]int{2, 2}, a)
}

func (s *HelloSuite) TestAppendToSliceBefore() {
	a := []int{1, 1}

	modify := func(i []int) {
		i = append(i, 3)
		s.Assert().NotEqual(a, i)
		i[0] = 2
		i[1] = 2
	}

	modify(a)

	s.Assert().Equal([]int{1, 1}, a)
}

func (s *HelloSuite) TestAppendToSlice() {
	a := []int{1, 1}

	modify := func(i *[]int) {
		*i = append(*i, 2)
	}

	modify(&a)

	s.Assert().Equal([]int{1, 1, 2}, a)
}

func (s *HelloSuite) TestReturnAppendedSlice() {
	a := []int{1, 1}

	modify := func(i []int) []int {
		return append(i, 2)
	}

	b := modify(a)

	s.Assert().Equal([]int{1, 1}, a)
	s.Assert().Equal([]int{1, 1, 2}, b)
}

func TestHelloSuite(t *testing.T) {
	suite.Run(t, new(HelloSuite))
}
