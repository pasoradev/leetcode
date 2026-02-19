package main

import (
	"fmt"
)

type LetterConnection struct {
	letter          byte
	nextConnections []*LetterConnection
	isEnd           bool
}

type WordDictionary struct {
	dict []*LetterConnection
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) add(word string, connection *LetterConnection, i int) {
	if i >= len(word) {
		connection.isEnd = true
		return
	}
	added := false
	for idx := range connection.nextConnections {
		c := connection.nextConnections[idx]
		if c.letter == word[i] {
			this.add(word, c, i+1)
			added = true
		}
	}
	if !added {
		newC := &LetterConnection{letter: word[i]}
		connection.nextConnections = append(connection.nextConnections, newC)
		this.add(word, newC, i+1)
	}
}

func (this *WordDictionary) AddWord(word string) {
	added := false
	for idx := range this.dict {
		c := this.dict[idx]
		if c.letter == word[0] {
			this.add(word, c, 1)
			added = true
		}
	}
	if !added {
		newRootC := &LetterConnection{letter: word[0]}
		this.dict = append(this.dict, newRootC)
		this.add(word, newRootC, 1)
	}
}

func (this *WordDictionary) search(word string, connection *LetterConnection, i int) bool {
	if i >= len(word) {
		return connection.isEnd || len(connection.nextConnections) == 0
	}
	if word[i] == '.' {
		// we have to try every possible way
		for idx := range connection.nextConnections {
			c := connection.nextConnections[idx]
			res := this.search(word, c, i+1)
			if res {
				return true
			}
		}
	} else {
		for idx := range connection.nextConnections {
			c := connection.nextConnections[idx]
			if c.letter != word[i] {
				continue
			}
			return this.search(word, c, i+1)
		}
	}
	return false
}

func (this *WordDictionary) Search(word string) bool {
	if word[0] == '.' {
		// we have to try every possible way
		for idx := range this.dict {
			c := this.dict[idx]
			res := this.search(word, c, 1)
			if res {
				return true
			}
		}
	} else {
		for idx := range this.dict {
			c := this.dict[idx]
			if c.letter != word[0] {
				continue
			}
			return this.search(word, c, 1)
		}
	}

	return false
}

func (this *WordDictionary) PrintTree() {
	for _, root := range this.dict {
		printNode(root, "", true)
	}
}

func printNode(node *LetterConnection, prefix string, isLast bool) {
	if node == nil {
		return
	}

	connector := "├── "
	nextPrefix := prefix + "│   "
	if isLast {
		connector = "└── "
		nextPrefix = prefix + "    "
	}

	fmt.Printf("%s%s%c\n", prefix, connector, node.letter)

	for i, child := range node.nextConnections {
		printNode(child, nextPrefix, i == len(node.nextConnections)-1)
	}
}

func main() {
	obj := Constructor()

	obj.AddWord("bad")
	obj.AddWord("bath")
	obj.AddWord("dad")
	obj.AddWord("mad")
	obj.PrintTree()

	fmt.Println(obj.Search("bad")) // expected true
	fmt.Println(obj.Search("pad")) // expected false
	fmt.Println(obj.Search(".ad")) // expected true
	fmt.Println(obj.Search("b..")) // expected true

	obj2 := Constructor()

	obj2.AddWord("at")
	obj2.AddWord("and")
	obj2.AddWord("an")
	obj2.AddWord("add")
	obj2.PrintTree()

	fmt.Println(obj2.Search("a"))   // expected false
	fmt.Println(obj2.Search(".at")) // expected false

	obj2.AddWord("bat")
	obj2.PrintTree()

	fmt.Println(obj2.Search(".at"))  // expected true
	fmt.Println(obj2.Search("an."))  // expected true
	fmt.Println(obj2.Search("a.d.")) // expected false
	fmt.Println(obj2.Search("b."))   // expected false
	fmt.Println(obj2.Search("a.d"))  // expected true
	fmt.Println(obj2.Search("."))    // expected false

	obj3 := Constructor()

	obj3.AddWord("a")
	obj3.AddWord("ab")
	obj3.PrintTree()

	fmt.Println(obj3.Search("a"))   // expected true
	fmt.Println(obj3.Search("a."))  // expected true
	fmt.Println(obj3.Search("ab"))  // expected true
	fmt.Println(obj3.Search(".a"))  // expected false
	fmt.Println(obj3.Search(".b"))  // expected true
	fmt.Println(obj3.Search("ab.")) // expected false
	fmt.Println(obj3.Search("."))   // expected true
	fmt.Println(obj3.Search(".."))  // expected true
}
