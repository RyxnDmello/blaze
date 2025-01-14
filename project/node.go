package project

type Node struct {
	icon  string
	name  string
	path  string
	isDir bool
}

func (node *Node) Icon() string {
	return node.icon
}

func (node *Node) Name() string {
	return node.name
}

func (node *Node) Path() string {
	return node.path
}

func (node *Node) IsDir() bool {
	return node.isDir
}
