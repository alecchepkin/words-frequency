package list

// Trie is an ordered tree data structure
// list - frequency list
// lnum - number of words searching
type Trie struct {
	root *Node
	list []*Node
	lnum int
}

func NewTrie(num int) *Trie {
	trie := &Trie{}
	trie.root = NewNode("")
	trie.list = make([]*Node, 0, num)
	trie.lnum = num
	return trie
}

func (trie *Trie) GetMostFrequent() []*Node {
	return trie.list
}

// Insert insert words in the trie and increase count.
// If a word exists it increase the count only.
// After that do adding to frequency list
func (trie *Trie) Insert(word string) {
	node := trie.root
	for i := 0; i < len(word); i++ {
		letter := string(word[i])
		var child *Node
		for _, childNode := range node.children {
			if letter == childNode.Letter {
				child = childNode
				node = child
				break
			}
		}
		if child == nil {
			child := NewNode(letter)
			node.children = append(node.children, child)
			node = child
		}
	}
	node.Count++
	node.Word = word

	// add to frequency list
	trie.addToList(node)
}

// Adding and sorting frequency list
func (trie *Trie) addToList(node *Node) int {
	index := node.index
	if index < 0 {
		if len(trie.list) < trie.lnum {
			trie.list = append(trie.list, node)
			index = len(trie.list) - 1
		} else if node.Count > trie.list[trie.lnum-1].Count {
			trie.list[trie.lnum-1].index = -1
			trie.list[trie.lnum-1] = node
			index = trie.lnum - 1
		} else {
			return index
		}
	}
	if index > 0 {
		for i := index; i > 0; i-- {
			if trie.list[i].Count >= trie.list[i-1].Count {
				trie.list[i-1], trie.list[i] = trie.list[i], trie.list[i-1]
				trie.list[i-1].index, trie.list[i].index = i-1, i
				index = i - 1
			}
		}
	}
	node.index = index
	return index
}
