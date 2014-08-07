package hashmap

// emptynode is equivalent to a nil pointer.
type emptyNode struct{}

// Assoc'ing turns an empty node into a leaf node.
func (n *emptyNode) assoc(shift uint, hash uint32, key Key, val Value) inode {
    return &leafNode{hash, key, val}
}

func (n *emptyNode) dissoc(shift uint, hash uint32, key Key) inode {
    return n
}

func (n *emptyNode) find(shift uint, hash uint32, key Key) (val Value, ok bool) {
    return 0, false
}

func (n *emptyNode) keys() []Key {
    return []Key{}
}

func (n *emptyNode) vals() []Value {
    return []Value{}
}

func (n *emptyNode) count() int {
    return 0
}

// a holder for a non-existent value
var empty = &emptyNode{}
