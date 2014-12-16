package class

import "gopkg.in/metakeule/goh4.v5"

var (
	ALL = []goh4.Class{}
)

func class(name string) (c goh4.Class) {
	c = goh4.Class(name)
	ALL = append(ALL, c)
	return
}
