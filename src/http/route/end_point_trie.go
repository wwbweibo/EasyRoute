package route

import (
	"errors"
	"strings"
)

// we use a trie tree to save all the EndPoint information
// 									/
//        |                 |                  |                          |
//  Controller1       Controller2         Controller3 				Controller4
//  |    |    |       |    |     |       |     |     |              |      |     |
// m1   m2   m3     m1    m2     m3     m1    m2     m3            m1     m2     m3
type EndPointTrie struct {
	root *EndPointTrieNode
}

// a EndPointTrieNode is a TrieTreeNode to save EndPoint information for single partten
type EndPointTrieNode struct {
	isEnd    bool
	endPoint *EndPoint
	next     []*EndPointTrieNode
	section  string
}

// create a new instance of EndPointTrie
func NewEndPointTrie() *EndPointTrie {
	root := &EndPointTrieNode{
		isEnd:    true,
		endPoint: nil,
		next:     make([]*EndPointTrieNode, 0),
		section:  "/",
	}

	return &EndPointTrie{root: root}
}

// add a EndPoint to map, if existing, will cover it
func (t *EndPointTrie) AddEndPoint(endPoint *EndPoint) {
	sections := strings.Split(endPoint.Template, "/")[1:]
	t.root.Insert(sections, endPoint)
}

// get the first matched request handler for the given path
func (t *EndPointTrie) GetMatchedRoute(path string) (*EndPoint, error) {
	if len(path) == 0 {
		return nil, errors.New("an empty path is not valid")
	}
	if path[0] != '/' {
		return nil, errors.New("a path must begin with '/'")
	}
	sections := strings.Split(path, "/")[1:]
	return t.root.Search(sections), nil
}

func (t *EndPointTrie) GetRoot() *EndPointTrieNode {
	return t.root
}

// add a EndPoint to EndPointTrie,
func (n *EndPointTrieNode) Insert(routeSections []string, endPoint *EndPoint) {
	if len(routeSections) == 1 {
		// will add to the root
		node := &EndPointTrieNode{
			isEnd:    true,
			endPoint: endPoint,
			next:     nil,
			section:  routeSections[0],
		}
		n.isEnd = false
		if n.next == nil {
			n.next = make([]*EndPointTrieNode, 0)
		}
		n.next = append(n.next, node)
	} else {
		section := routeSections[0]
		nextNode := n.searchBySection(section)

		if nextNode == nil {
			nextNode = &EndPointTrieNode{
				isEnd:    true,
				endPoint: nil,
				next:     nil,
				section:  section,
			}
			n.next = append(n.next, nextNode)
		}
		nextNode.Insert(routeSections[1:], endPoint)
	}
}

func (n *EndPointTrieNode) Search(sections []string) *EndPoint {
	if len(sections) == 0 {
		return n.endPoint
	} else {
		next := n.searchBySection(sections[0])
		if next != nil {
			return next.Search(sections[1:])
		} else {
			return nil
		}
	}
}

func (n *EndPointTrieNode) searchBySection(section string) *EndPointTrieNode {
	if n.next == nil {
		n.next = make([]*EndPointTrieNode, 0)
	}
	for _, item := range n.next {
		if item.section == section {
			return item
		}
	}
	return nil
}

func (n *EndPointTrieNode) GetEndPoint() *EndPoint {
	return n.endPoint
}

func (n *EndPointTrieNode) GetNext() []*EndPointTrieNode {
	return n.next
}
