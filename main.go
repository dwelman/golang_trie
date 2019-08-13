package main

import "fmt"

type Node struct {
	Children map[rune]*Node
	Value    rune
	Priority int
}

func MakeNode(value rune, priority int) *Node {
	return &Node{Children: make(map[rune]*Node, 1), Value: value, Priority: priority}
}

func AddWord(head *Node, word string) *Node {
	letters := []rune(word)
	length := len(letters)
	head.Priority++

	crawl := head
	for i := 0; i <= length; i++ {
		if i == length {
			crawl.Children[-1] = MakeNode(-1, 0)
			break
		}
		if val, ok := crawl.Children[letters[i]]; ok {
			crawl.Children[letters[i]] = AddWord(val, string(letters[i:]))
			break
		} else {
			n := MakeNode(letters[i], 1)
			n = AddWord(n, string(letters[i+1:]))
			crawl.Children[letters[i]] = n
			break
		}
	}

	return crawl
}

func PrintTrie(head *Node) {
	crawl := head

	for _, v := range crawl.Children {
		if v.Value == -1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%c", v.Value)
			PrintTrie(v)
		}
	}
}

func main() {
	head := MakeNode(-1, 0)
	head = AddWord(head, "apple")
	head = AddWord(head, "and")
	head = AddWord(head, "anvil")
	head = AddWord(head, "dog")
	head = AddWord(head, "doodle")
	head = AddWord(head, "addictive")
	PrintTrie(head)
}
