package main

import "fmt"

//public class TrieNode {
//public TrieNode[] children = new TrieNode[26];
//public boolean isWord;
//}
//
//
///** Initialize your data structure here. */
//private TrieNode root = new TrieNode();
//
//public AddSearchWord211() {
//
//}
//
///** Adds a word into the data structure. */
//public void addWord(String word) {
//TrieNode node = root;
//for (char c : word.toCharArray()) {
//if (node.children[c - 'a'] == null) {
//node.children[c - 'a'] = new TrieNode();
//}
//node = node.children[c - 'a'];
//}
//node.isWord = true;
//}
//
///** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
//public boolean search(String word) {
//
//return match(word.toCharArray(), 0, root);
//
//}
//
//private boolean match(char[] chs, int k, TrieNode node) {
//if (k == chs.length) return node.isWord;
//
//if (chs[k] != '.') {
//return node.children[chs[k] - 'a'] != null && match(chs, k + 1, node.children[chs[k] - 'a']);
//} else {
//for (int i = 0; i < node.children.length; i++) {
//if (node.children[i] != null) {
//if (match(chs, k + 1, node.children[i])) {
//return true;
//}
//}
//}
//}
//return false;
//}

type TrieNode struct {
	Children  []*TrieNode
	IsWord    bool
}



type WordDictionary struct {
	root *TrieNode
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	w := WordDictionary{}
	w.root = &TrieNode{Children: make([]*TrieNode, 26)}
	return w
}


func (this *WordDictionary) AddWord(word string)  {
	node := this.root
	for _, char := range word {
		c := char - 'a'
		if node.Children[c] == nil {
			node.Children[c] = &TrieNode{Children: make([]*TrieNode, 26)}
		}
		node = node.Children[c]
	}
	node.IsWord = true
}


func (this *WordDictionary) Search(word string) bool {
	return this.Match([]byte(word), 0, this.root)
}

func (this *WordDictionary) Match(word []byte, index int, root *TrieNode) bool {
	fmt.Println(string(word))
	if index == len(word) {
		return root.IsWord
	}
	if word[index] != '.' {
		c := word[index] - 'a'
		return root.Children[c] != nil && this.Match(word, index + 1, root.Children[c])
	} else {
		for _, child := range root.Children {
			if child != nil {
				if this.Match(word, index + 1, child) {
					return true
				}
			}
		}
	}
	return false

}

func main() {
	dict := Constructor()
	dict.AddWord("at")
	dict.AddWord("and")
	dict.AddWord("an")
	dict.AddWord("add")
	fmt.Println(dict.Search("a"))
	fmt.Println(dict.Search(".at"))
	dict.AddWord("bat")

	fmt.Println(dict.Search(".at"))
	fmt.Println(dict.Search("an."))
	fmt.Println(dict.Search("a.d."))
	fmt.Println(dict.Search("b."))
	fmt.Println(dict.Search("a.d"))
	fmt.Println(dict.Search("."))

}
