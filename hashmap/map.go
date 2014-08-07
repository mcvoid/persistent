// Package hashmap implements an persistent hash array-mapped trie
// for use as an immutable associative array.
// Any insertions or deletions on this data structure will not
// change the original map, but rather return a new map with those
// changes in place. This will help eliminate race conditions and make
// maps thread-safe as outside goroutines cannot modify your current data.
// The mutated versions of the maps can then be shared over channels
// or such.
//
// The trie uses structural sharing rather than full copies for
// decent performance: the number of copies is limited to log32(m) for 
// m entries in most cases.
//
// Keys are required to supply their own hashing routine.
//
// Return values are sent back as interface{} and so they have to be cast
// back to their original values for use afterwards.
package hashmap

// A Map is a reference to a persistent hash array-mapped trie
type Map struct{
    root inode
    count int
}

// A Key is any value which can supply a hash value.
type Key interface {
    Hash() int
}
// A value is a placeholder for any kind of data.
type Value interface{}

// inode describes a trie node.
type inode interface {
    assoc(shift, hash int, key Key, val Value) (node inode, count int)
    dissoc(shift, hash int, key Key) (node inode, count int)
    find(shift, hash int, key Key) (val Value, ok bool)
    keys() []Key
    vals() []Value
}

// New creates a new persistent associative array.
func New() *Map {
    return &Map{empty}
}

// Adds a new key/value pair to the associative array
func Put(m *Map, key Key, val Value) *Map {
    hash := key.Hash()
    return &Map{m.root.assoc(0, hash, key, val)}
}

// Delete removes the key and its related value from the array.
func Delete(m *Map, key Key) *Map {
    hash := key.Hash()
    return &Map{m.root.dissoc(0, hash, key, val)}
}

// Get returns the value stored in the map at a given key.
func Get(m *Map, key Key) Value {
    hash := key.Hash()
    return m.root.find(0, hash, key)
}

// Keys returns a list of all the keys stored in the array.
func Keys(m *Map) []Key {
    return m.root.keys()
}

// Values returns a list of all the values stored in the array.
func Values(m *Map) []Value {
    return m.vals()
}

// Len determines the number of items in the map.
func Len(m *Map) int {
    return m.root.count
}
