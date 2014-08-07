package hashmap

import "testing"

func TestEmpty(t *testing.T) {
    h := New()
    if h == nil {
        t.Error("should not be nil")
    }
}
