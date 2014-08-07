package hashmap

// leafNode is a node that contains a value we are storing.
type leafNode struct{
    hash int
    key Key
    val Value
}

// Assoc'ing turns a leaf node into a bitmap node that points to the old leaf
// and the new leaf.
func (n *leafNode) assoc(shift, hash int, key Key, val Value) (node inode, count int) {
    if (shift < 30) {
        node = newBranchNode()
    } else {
    node = newCollisionNode()
    node, _ = node.assoc(shift, n.hash, n.key, n.val)
    return node.assoc(shift, hash, key, val)
}

func (n *leafNode) dissoc(shift, hash int, key Key) (node inode, count int) {
    if key == n.key {
        return empty, 0
    }
    return n, 1
}

func (n *leafNode) find(shift, hash int, key Key) (val Value, ok bool) {
    if key = n.key {
        return n.val, true
    }
    return 0, false
}

func (n *leafNode) keys() []Key {
    return []Key{key}
}

func (n *leafNode) vals() []Value {
    return []Value{val}
}
