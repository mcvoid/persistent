package hashmap

// leafNode is a node that contains a value we are storing.
type leafNode struct{
    hash uint32
    key Key
    val Value
}

// Assoc'ing turns a leaf node into a bitmap node that points to the old leaf
// and the new leaf.
func (n *leafNode) assoc(shift uint, hash uint32, key Key, val Value) (node inode) {
    if (shift < 30) {
        node = newBranchNode()
    } else {
        node = newCollisionNode()
    }
    node = node.assoc(shift, n.hash, n.key, n.val)
    return node.assoc(shift, hash, key, val)
}

func (n *leafNode) dissoc(shift uint, hash uint32, key Key) inode {
    if key == n.key {
        return empty
    }
    return n
}

func (n *leafNode) find(shift uint, hash uint32, key Key) (val Value, ok bool) {
    if key == n.key {
        return n.val, true
    }
    return 0, false
}

func (n *leafNode) keys() []Key {
    return []Key{n.key}
}

func (n *leafNode) vals() []Value {
    return []Value{n.val}
}

func (n *leafNode) count() int {
    return 1
}
