package hashmap

// collisionNode is a node for holding all values whose keys hash to the same first 30 bits.
type collisionNode map[Key]Value

// newCollisionNode creates a new collision node that has its map already allocated.
func newCollisionNode() collisionNode {
	return make(map[Key]Value)
}

func (n collisionNode) clone() collisionNode {
	node := newCollisionNode()
	for k, v := range n {
		node[k] = v
	}
	return node
}

// Assoc'in a collision node is just managing a map- this map should be very small
// so copying it should be cheap.
func (n collisionNode) assoc(shift uint, hash uint32, key Key, val Value) inode {
	newnode := n.clone()
	newnode[key] = val
	return newnode
}

func (n collisionNode) dissoc(shift uint, hash uint32, key Key) inode {
	if _, ok := n[key]; !ok {
		return n
	}
	newnode := n.clone()
	if delete(newnode, key); len(newnode) == 0 {
		return empty
	} else if len(newnode) == 1 {
		return &leafNode{key.Hash(), key, newnode[key]}
	} else if shift >= 30 {
		var b inode = newBranchNode()
		for k, v := range newnode {
			b = b.assoc(shift, k.Hash(), k, v)
		}
		return b
	}
	return newnode
}

func (n collisionNode) find(shift uint, hash uint32, key Key) (val Value, ok bool) {
	v, ok := n[key]
	return v, ok
}

func (n collisionNode) keys() []Key {
	k := make([]Key, len(n))
	i := 0
	for key := range n {
		k[i] = key
		i++
	}
	return k
}

func (n collisionNode) vals() []Value {
	v := make([]Value, len(n))
	i := 0
	for _, val := range n {
		v[i] = val
		i++
	}
	return v
}

func (n collisionNode) count() int {
	return len(n)
}
