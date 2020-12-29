package go_rope

type Rope interface {
	String() string
	Len() int
	Index(idx int) byte
	Depth() byte
}
