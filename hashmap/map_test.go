package hashmap

import "testing"

type intKey int

func (i intKey) Hash() uint32 {
	return uint32(i)
}

func getInt(m Map, k int) (int, bool) {
	v, ok := Get(m, intKey(k))
	return v.(int), ok
}

func putInt(m Map, k int, v int) Map {
	return Put(m, intKey(k), v)
}

func deleteInt(m Map, k int) Map {
    return Delete(m, intKey(k))
}

func TestNew(t *testing.T) {
	m := New()
	if m == nil {
		t.Error("should not be nil")
	}
	_, ok := getInt(m, 0)
	if ok {
		t.Error("reported key 0 found in empty map")
	}
}

func TestPut(t *testing.T) {
	m1 := New()
	m2 := putInt(m1, 0, 5)
	v, ok := getInt(m2, 0)
	if !ok {
		t.Error("key 0 not found after insertion")
	}
	if v != 5 {
		t.Error("expected 5, got ", v)
	}
	if _, ok = getInt(m2, 5); ok {
		t.Error("returned non-existent key after inserting something else")
	}
	if _, ok = getInt(m1, 0); ok {
		t.Error("Put() mutated original map")
	}
    m3 := putInt(m2, 0, 4)
    if v, _ = getInt(m3, 0); v != 4 {
        t.Error("returned wrong value on re-assoc, expected 4, got ", v)
    }
}

func TestDelete(t *testing.T) {
    m := New()
    m = putInt(putInt(putInt(m, 0, 0), 1, 1), 2, 2)
    m = deleteInt(m, 1)
    if _, ok := getInt(m, 1); ok {
        t.Error("returned non-existent key that was just deleted")
    }
    if v, k := getInt(m, 2); !k || v!= 2 {
        t.Error("Wrong keys missing after delete")
    }
}

func TestLen(t *testing.T) {
    m := New()
    if Len(m) != 0 {
        t.Error("empty map not length 0")
    }
    for i := 0; i < 100; i++ {
        m = putInt(m, i, 100 - i)
    }
    if l := Len(m); l != 100 {
        t.Error("incorrect length: expected 100, actual ", l)
    }
}

func TestKeys(t *testing.T) {
    m := New()
    m = putInt(putInt(putInt(m, 0, 0), 1, 1), 2, 2)
    for _, v := range []intKey{0, 1, 2} {
        found := false
        for _, k := range Keys(m) {
            if v == k.(intKey) {
                found = true
            }
        }
        if !found {
            t.Error("Key not found: ", v)
        }
    }
}

func TestValues(t *testing.T) {
    m := New()
    m = putInt(putInt(putInt(m, 0, 0), 1, 1), 2, 2)
    for _, v := range []int{0, 1, 2} {
        found := false
        for _, u := range Values(m) {
            if v == u.(int) {
                found = true
            }
        }
        if !found {
            t.Error("Key not found: ", v)
        }
    }
}
