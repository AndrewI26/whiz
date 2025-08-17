package routing

type RouteNode struct {
	methodToHandler map[string]Handler
	children        map[string]*RouteNode
	params          []string
}

func NewRouteNode() *RouteNode {
	return &RouteNode{map[string]Handler{}, make(map[string]*RouteNode), make([]string, 0)}
}
