# Persistent Data Structures For Go

This package contains several immutable containers to help with concurrency issues like race conditions. Any operation which changes the data in the collection will return a new collection with the changes made, rather thanmutating the original collection.

Currently supported data structures are a stack and an associative array.

## PACKAGE DOCUMENTATION

### package hashmap
    `import "github.com/mcvoid/persistent/hashmap"`

    Package hashmap implements an persistent hash array-mapped trie for use
    as an immutable associative array. Any insertions or deletions on this
    data structure will not change the original map, but rather return a new
    map with those changes in place. This will help eliminate race
    conditions and make maps thread-safe as outside goroutines cannot modify
    your current data. The mutated versions of the maps can then be shared
    over channels or such.

    The trie uses structural sharing rather than full copies for decent
    performance: the number of copies is limited to log32(m) for m entries
    in most cases.

    Keys are required to supply their own hashing routine.

    Return values are sent back as interface{} and so they have to be cast
    back to their original values for use afterwards.

#### FUNCTIONS

`func Keys(m Map) []Key`
    Keys returns a list of all the keys stored in the array.

`func Len(m Map) int`
    Len determines the number of items in the map.

`func Values(m Map) []Value`
    Values returns a list of all the values stored in the array.

`func Delete(m Map, key Key) Map`
    Delete removes the key and its related value from the array.

`func New() Map`
    New creates a new persistent associative array.

`func Put(m Map, key Key, val Value) Map`
    Adds a new key/value pair to the associative array

`func Get(m Map, key Key) (val Value, ok bool)`
    Get returns the value stored in the map at a given key.

#### TYPES

`type Map inode`
    A Map is a reference to a persistent hash array-mapped trie

```
type Key interface {
    Hash() uint32
}
```
    A Key is any value which can supply a hash value.

`type Value interface{}`
    A value is a placeholder for any kind of data.



## License

The MIT License (MIT)

Copyright (c) 2014 Sean Wolcott

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
