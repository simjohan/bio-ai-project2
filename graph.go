package main

type Element struct {
	Parent *Element
	Rank   int
	Value  interface{}
}

func MakeSet(value interface{}) *Element {
	e := new(Element)
	e.Parent = e
	e.Rank = 0
	e.Value = value
	return e
}

func FindSet(e *Element) *Element {
	if e == e.Parent {
		return e
	}
	e.Parent = FindSet(e.Parent)
	return e.Parent
}

func Union(e1, e2 *Element) {
	r1 := FindSet(e1)
	r2 := FindSet(e2)

	if r1 == r2 {
		return
	}

	if r1.Rank > r2.Rank {
		r2.Parent = r1
	} else if r1.Rank < r2.Rank {
		r1.Parent = r2
	} else {
		r1.Parent = r2
		r2.Rank = r2.Rank + 1
	}

}