package route_test

import (
	"github.com/wwbweibo/EasyRoute/src/http/route"
	"strings"
	"testing"
)

func TestNewEndpointTrie(t *testing.T) {
	instance := route.NewEndPointTrie()
	if instance == nil {
		t.Failed()
	}
	root := instance.GetRoot()

	if root == nil {
		t.Failed()
	}
	endpoint := root.GetEndPoint()
	if endpoint != nil {
		t.Failed()
	}
}

func TestAddNewEndpoint(t *testing.T) {
	root := route.NewEndPointTrie().GetRoot()
	endpoint := route.EndPoint{
		Template: "/Home",
	}
	root.Insert(strings.Split(endpoint.Template, "/")[1:], &endpoint)
	next := root.GetNext()
	if next == nil || len(next) == 0 {
		t.Fail()
	}
	if next[0].GetEndPoint().Template != "/Home" {
		t.Fail()
	}

	endpoint2 := route.EndPoint{
		Template: "/User",
	}
	root.Insert(strings.Split(endpoint2.Template, "/")[1:], &endpoint2)
	next = root.GetNext()
	if next == nil || len(next) != 2 {
		t.Fail()
	}
	if next[1].GetEndPoint().Template != "/User" {
		t.Fail()
	}

	endpoint3 := route.EndPoint{
		Template: "/Home/Index",
	}
	root.Insert(strings.Split(endpoint3.Template, "/")[1:], &endpoint3)
	next = root.GetNext()
	if next == nil || len(next) != 2 {
		t.Fail()
	}
	next = next[0].GetNext()
	if next[0].GetEndPoint().Template != "/Home/Index" {
		t.Fail()
	}
}

func TestAddEndPointTest(t *testing.T) {
	tree := route.NewEndPointTrie()
	endPoint := route.EndPoint{Template: "/Home/Index"}
	tree.AddEndPoint(&endPoint)
	e, _, err := tree.GetMatchedRoute("/Home/Index")
	if err != nil {
		t.Fail()
	}
	if e == nil {
		t.Error("endpoint is  nil")
	}
	if e.GetEndPoint().Template != "/Home/Index" {
		t.Fail()
	}

}
