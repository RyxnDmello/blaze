package directory

type Node struct {
	name     string
	icon     string
	isDir    bool
	isOpen   bool
	children *Directory
	prev     *Node
	next     *Node
}

func (node Node) Name() string {
	return node.name
}

func (node Node) Icon() string {
	return node.icon
}

func (node Node) IsDir() bool {
	return node.isDir
}

func (node Node) IsOpen() bool {
	return node.isOpen
}

func (node Node) Children() *Directory {
	return node.children
}

func (node Node) Prev() *Node {
	return node.prev
}

func (node Node) Next() *Node {
	return node.next
}
