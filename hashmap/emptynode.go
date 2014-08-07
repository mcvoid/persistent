package hashmap

// emptynode is equivalent to a nil pointer.
type emptyNode struct{}

// Assoc'ing turns an empty node into a leaf node.
func (n *emptyNode) assoc(shift, hash int, key Key, val Value) (node inode, count int) {
    return &leafNode{hash, key, val}, 1
}

func (n *branchNode) dissoc(shift, hash int, key Key) (node inode, count int) {
    return n, 0
}

func (n *branchNode) find(shift, hash int, key Key) (val Value, ok bool) {
    return 0, false
}

func (n *branchNode) keys() []Key {
    return []Key{}
}

func (n *branchNode) vals() []Value {
    return []Value{}
}

// a holder for a non-existent value
var empty = &emptyNode{}
