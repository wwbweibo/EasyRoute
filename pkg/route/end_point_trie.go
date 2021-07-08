package route

import (
	"errors"
	"github.com/wwbweibo/EasyRoute/pkg/delegates"
	"github.com/wwbweibo/EasyRoute/pkg/http"
	http2 "net/http"
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

// EndPointTrieNode is a TrieTreeNode to save EndPoint information for single partten
type EndPointTrieNode struct {
	isEnd          bool
	endPoint       *EndPoint
	next           []*EndPointTrieNode
	section        string
	defaultHandler delegates.RequestDelegate
}

// create a new instance of EndPointTrie
func NewEndPointTrie() *EndPointTrie {
	root := &EndPointTrieNode{
		isEnd:    true,
		endPoint: nil,
		next:     make([]*EndPointTrieNode, 0),
		section:  "/",
		defaultHandler: func(ctx *http.HttpContext) {
			ctx.Response.WriteHeader(http2.StatusNotFound)
			ctx.Response.Write([]byte("404 Not Found"))
		},
	}

	return &EndPointTrie{root: root}
}

// add a EndPoint to map, if existing, will cover it
func (t *EndPointTrie) AddEndPoint(endPoint *EndPoint) {
	sections := strings.Split(endPoint.Template, "/")[1:]
	t.root.Insert(sections, endPoint)
}

// get the first matched request handler for the given path
func (t *EndPointTrie) GetMatchedRoute(path string) (*EndPointTrieNode, bool, error) {
	if len(path) == 0 {
		return nil, false, errors.New("an empty path is not valid")
	}
	if path[0] != '/' {
		return nil, false, errors.New("a path must begin with '/'")
	}
	sections := strings.Split(path, "/")[1:]
	targetNode, isMatched := t.root.Search(sections)
	return targetNode, isMatched, nil
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

func (n *EndPointTrieNode) Search(sections []string) (*EndPointTrieNode, bool) {
	if len(sections) == 0 {
		// 表明当前节点就是要匹配的节点，直接返回当前节点
		return n, true
	} else {
		// 搜索匹配的下一个节点
		next := n.searchBySection(sections[0])
		var targetNode *EndPointTrieNode
		isMatched := false
		if next != nil {
			// 搜索下一级
			targetNode, isMatched = next.Search(sections[1:])
		}
		if targetNode == nil {
			if n.defaultHandler != nil {
				targetNode = n
			} else {
				targetNode = nil
			}
		}

		return targetNode, isMatched
	}
}

// 在当前节点的下级节点中搜索是否有匹配的节点
// 如果有，返回对应的节点，
// 否则返回空
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
