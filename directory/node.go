package directory

type Node struct {
	name  string
	icon  string
	ref   string
	isDir bool
}

func (node Node) Name() string {
	return node.name
}

func (node Node) Icon() string {
	return node.icon
}

func (node Node) Reference() string {
	return node.ref
}

func (node Node) IsDir() bool {
	return node.isDir
}
