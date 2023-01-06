package main

import (
	"fmt"
)

type Parent struct {
	ID     int
	Name   string
	Childs Childs
}

type Child struct {
	ID       int
	ParentID int
	Name     string
}
type Childs []*Child

func (cs Childs) len() int {
	return len(cs)
}

func (cs Childs) filterBy(f func(c *Child) bool) Childs {
	childs := make([]*Child, 0)
	for i := range cs {
		if f(cs[i]) {
			childs = append(childs, cs[i])
		}
	}
	return childs
}

func main() {
	fmt.Println("array")

	childs := make([]*Child, 0)
	childs = append(childs, &Child{ID: 1, Name: "name_1", ParentID: 1})
	childs = append(childs, &Child{ID: 2, Name: "name_1", ParentID: 1})
	childs = append(childs, &Child{ID: 3, Name: "name_2", ParentID: 1})
	childs = append(childs, &Child{ID: 4, Name: "name_2", ParentID: 1})
	childs = append(childs, &Child{ID: 5, Name: "name_3", ParentID: 1})
	childs = append(childs, &Child{ID: 6, Name: "name_4", ParentID: 1})

	parent := &Parent{ID: 1, Name: "pname_1", Childs: childs}

	fmt.Println(parent.Childs)
	fmt.Println(parent.Childs.len())
	name1childs := parent.Childs.filterBy(func(m *Child) bool {
		return m.Name == "name_1"
	})
	fmt.Println(name1childs)
	fmt.Println(name1childs.len())
	fmt.Println(name1childs[0].Name)
	fmt.Println(name1childs[1].Name)
}
