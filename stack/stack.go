package stack

type element struct {
    value interface{}
    next *element
}

type Stack struct {
    top *element
    size int
}

func New() *Stack {
    return &Stack{nil, 0}
}

func Peek(s *Stack) interface{} {
    if top == nil { return nil }
    return top.value
}

func Push(s *Stack, val interface{}) *Stack {
    return &Stack{&element{val, s.top}, s.size + 1}
}

func Pop(s *Stack) (*Stack, interface{}) {
    if top == nil { return s, nil }
    return &Stack{s.top.next, s.size - 1}, s.top.value
}
