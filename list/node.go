package list

type Node struct {
	Letter   string
	Word     string
	Count    int
	index    int
	children []*Node
}

func NewNode(letter string) *Node {
	node := &Node{}
	node.Letter = letter
	node.children = make([]*Node, 0)
	node.index = -1
	return node
}
