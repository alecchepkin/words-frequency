package list

// Node is a element of trie
// If node is a leaf it increasing counter, has word
// Index is a index of frequency list (Trie.list), default value is -1.
// If a word isn't in the frequent list, then it the index is equal to -1
//
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
