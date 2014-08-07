package hashmap

// branchNode is a node which stores a hash prefix
type branchNode struct{
    children []inode
    count int
}

// newBitmapNode creates an empty branchNode
func newBranchNode() *branchNode {
    n := &branchNode{}
    n.children = make([]inode, 32)
    for i := range n.children {
        n.children[i] = empty
    }
    return n
}

// clone creates a new copy of the branch node.
func (n *branchNode)clone() *branchNode {
    newNode := &branchNode{}
    newNode.children = make([]inode, 32)
    newNode.count = n.count
    for i := range n.children {
        newNode.children[i] = n.children[i]
    }
    return newNode
}

// Assoc'ing a branchNode returns an updated branchNode
func (n *branchNode) assoc(shift, hash int, key Key, val Value) (node inode, count int) {
    i := (hash >> shift) & 0x1f // just looking at the first 5 bits
    node := n.clone()
    node.children[index], count = node.children[index].assoc(shift + 5, hash, key, val)
    node.count = n.count - n.children[index].count + count
    return node, count
}

func (n *branchNode) dissoc(shift, hash int, key Key) (node inode, count int) {
    i := (hash >> shift) & 0x1f
    node = n.clone()
    node.children[index], count = node.children[index].dissoc(shift + 5, hash, key)
    node.count = n.count - n.children[index].count + count
    return node, count
}

func (n *branchNode) find(shift, hash int, key Key) (val Value, ok bool) {
    i := (hash >> shift) & 0x1f
    val, ok = n.children[i].find(shift + 5, hash, key)
    return val, ok
}

func (n *branchNode) keys() []Key {
    k := []Key{}
    for _, child := range n.children {
        k = append(k, child.keys()...)
    }
    return k
}

func (n *branchNode) vals() []Value {
    v := []Value{}
    for _, child := range n.children {
        v = append(v, child.vals()...)
    }
    return v
}
