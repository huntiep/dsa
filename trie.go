package main

type Trie struct {
    wordp bool
    children [26]*Trie
}

func (t *Trie) insert(s string) {
    n := t
    for _, char := range(s) {
        idx := char - 'a'
        if n.children[idx] == nil {
            n.children[idx] = &Trie { wordp: false, children: [26]*Trie{} }
        }
        n = n.children[idx]
    }
    n.wordp = true
}

func (t *Trie) remove(s string) {
    // TODO
    /*
    var path []*Trie
    n := t
    for _, char := range(s) {
        idx := char = 'a'
        if n.children[idx] == nil {
            return
        }
        path = append(path, n)
        n = n.children[idx]
    }
    path[len(path)-1].wordp = false
    for i := len(path)-1; i > 0; i-- {
        if path[i].wordp {
            return
        }
        path[i-1]
    }
    */
}

func (t *Trie) contains(s string) bool {
    n := t
    for _, char := range(s) {
        idx := char - 'a'
        if n.children[idx] == nil {
            return false
        }
        n = n.children[idx]
    }

    return n.wordp
}

func testTrie() {
    t := Trie { wordp: false, children: [26]*Trie{} }
    t.insert("food")
    assert(!t.contains("foo"), "foo isn't a word")
    assert(t.contains("food"), "abc")
    t.insert("foodie")
    t.insert("foot")
    assert(!t.contains("foo"), "foo isn't a word")
    assert(t.contains("food"), "abc")
    assert(t.contains("foot"), "abc")
    assert(t.contains("foodie"), "abc")
}
