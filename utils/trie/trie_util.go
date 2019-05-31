package trie

import (
	"fmt"
)

// 前缀树
type Trie struct {
	next  map[rune]*Trie
	isEnd bool
}

func NewTrie() *Trie {
	root := new(Trie)
	root.next = make(map[rune]*Trie)
	root.isEnd = false
	return root
}

func (this *Trie) Insert(word string) {
	tmp := this
	index := len([]rune(word)) - 1
	for i, v := range []rune(word) {
		if _, exist := tmp.next[v]; !exist {
			node := new(Trie)
			node.next = make(map[rune]*Trie)
			if i == index {
				node.isEnd = true
			}
			tmp.next[v] = node
		}
		tmp = tmp.next[v]
	}
	tmp.isEnd = true
}

func (this *Trie) Search(word string) bool {
	tmp := this
	for _, v := range word {
		if tmp.next[v] == nil {
			return false
		}
		tmp = tmp.next[v]
	}

	return tmp.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	tmp := this
	for _, v := range prefix {
		if tmp.next[v] == nil {
			return false
		}
		tmp = tmp.next[v]
	}
	return true
}

func (this *Trie) GetStartsWith(prefix string) (result []string) {
	tmp := this
	if !tmp.StartsWith(prefix) {
		return
	}
	for _, v := range prefix {
		tmp = tmp.next[v]
	}
	result = tmp.getKey(prefix)
	return
}

func (this *Trie) getKey(prefix string) (result []string) {
	if this.isEnd {
		result = append(result, prefix)
	}
	for key, val := range this.next {
		result = append(result, val.getKey(fmt.Sprintf("%s%s", prefix, string(key)))...)
	}
	return
}