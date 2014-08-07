package hashmap

// collisionNode is a node for holding all values whose keys hash to the same first 30 bits.
type collisionNode struct {
    children map[Key]Value
}

// newCollisionNode creates a new collision node that has its map already allocated.
func newCollisionNode() *collisionNode {
    m := make(map[Key]Value)
    return &collisionNode{m}
}


// Assoc'in a collision node is just managing a map- this map should be very small
// so copying it should be cheap.
func (n *collisionNode) assoc(shift, hash int, key Key, val Value) (node inode, count int) {
    node = newCollisionNode()
    for k, v := range n.children {
        node.children[k] = v
    }
    node.children[key] = val
    return node, len(node.children)
}

func (n *collisionNode) dissoc(shift, hash int, key Key) (node inode, count int) {
    if _, ok = n.children[key]; !ok {
        return n, len(n.children)
    }
    node = newCollisionNode()
    for k, v := range n.children {
        node.children[k] = v
    }
    if delete(node.children, key); len(node.children) == 0 {
        return empty, 0
    } else if len(node.children) == 1 {
        return &leafNode{key, val}, 1
    }
    return node, len(node.children)
}

func (n *collisionNode) find(shift, hash int, key Key) (val Value, ok bool) {
    val, ok = n.children[key]
    return val, ok
}

func (n *collisionNode) keys() []Key {
    k := make([]Key, len(n.children))
    i := 0
    for key := range n.children {
        k[i] = key
        i++
    }
    return k
}

func (n *collisionNode) vals() []Value {
    v := make([]Value, len(n.children))
    i := 0
    for _, val := range n.children {
        v[i] = val
        i++
    }
    return v
}
