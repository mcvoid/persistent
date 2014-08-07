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
type Map inode

// A Key is any value which can supply a hash value.
type Key interface {
	Hash() uint32
}

// A value is a placeholder for any kind of data.
type Value interface{}

// inode describes a trie node.
type inode interface {
	assoc(shift uint, hash uint32, key Key, val Value) inode
	dissoc(shift uint, hash uint32, key Key) inode
	find(shift uint, hash uint32, key Key) (val Value, ok bool)
	keys() []Key
	vals() []Value
	count() int
}

// New creates a new persistent associative array.
func New() Map {
	return empty
}

// Adds a new key/value pair to the associative array
func Put(m Map, key Key, val Value) Map {
	hash := key.Hash()
	return m.assoc(0, hash, key, val)
}

// Delete removes the key and its related value from the array.
func Delete(m Map, key Key) Map {
	hash := key.Hash()
	return m.dissoc(0, hash, key)
}

// Get returns the value stored in the map at a given key.
func Get(m Map, key Key) (val Value, ok bool) {
	hash := key.Hash()
	val, ok = m.find(0, hash, key)
	return val, ok
}

// Keys returns a list of all the keys stored in the array.
func Keys(m Map) []Key {
	return m.keys()
}

// Values returns a list of all the values stored in the array.
func Values(m Map) []Value {
	return m.vals()
}

// Len determines the number of items in the map.
func Len(m Map) int {
	return m.count()
}
