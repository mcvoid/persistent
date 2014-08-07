package hashmap

// branchNode is a node which stores a hash prefix
type branchNode []inode

// newBitmapNode creates an empty branchNode
func newBranchNode() branchNode {
	n := make([]inode, 32)
	for i := range n {
		n[i] = empty
	}
	return n
}

// clone creates a new copy of the branch node.
func (n branchNode) clone() branchNode {
	return append([]inode{}, n...)
}

// Assoc'ing a branchNode returns an updated branchNode
func (n branchNode) assoc(shift uint, hash uint32, key Key, val Value) inode {
	i := (hash >> shift) & 0x1f // just looking at the first 5 bits
	newnode := n.clone()
	newnode[i] = newnode[i].assoc(shift+5, hash, key, val)
	return newnode
}

func (n branchNode) dissoc(shift uint, hash uint32, key Key) inode {
	i := (hash >> shift) & 0x1f
	newnode := n.clone()
	newnode[i] = newnode[i].dissoc(shift+5, hash, key)
	return newnode
}

func (n branchNode) find(shift uint, hash uint32, key Key) (val Value, ok bool) {
	i := (hash >> shift) & 0x1f
	val, ok = n[i].find(shift+5, hash, key)
	return val, ok
}

func (n branchNode) keys() []Key {
	k := []Key{}
	for _, child := range n {
		k = append(k, child.keys()...)
	}
	return k
}

func (n branchNode) vals() []Value {
	v := []Value{}
	for _, child := range n {
		v = append(v, child.vals()...)
	}
	return v
}

func (n branchNode) count() int {
	c := 0
	for _, child := range n {
		c += child.count()
	}
	return c
}
