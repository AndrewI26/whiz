package routing

import (
	"fmt"
	"strings"

	logging "github.com/AndrewI26/whiz/logger"
)

type Handler func(params map[string]string) *Response

type Router struct {
	Root   *RouteNode
	logger *logging.Logger
}

func NewRouter(logger *logging.Logger) *Router {
	return &Router{Root: NewRouteNode(), logger: logger}
}

// AddRoute inserts a route node into the routes tree.
func (r *Router) AddRoute(method string, path string, handler Handler) error {
	if path[0] != '/' {
		return fmt.Errorf("path must start with '/' (%s)", path)
	}

	currentNode := r.Root
	routeParts := strings.Split(path, "/")
	dynamicParams := make([]string, 0)

	for _, rawSegment := range routeParts {
		segment := strings.ToLower(rawSegment)
		if strings.Contains(segment, " ") {
			return fmt.Errorf("path segment must not contain ' ' (%s)", segment)
		}

		if segment == "" {
			continue // skip empty parts, e.g. leading "/"
		}

		isDynamic := segment[0] == ':'
		var key string
		if isDynamic {
			key = ":"
			dynamicParams = append(dynamicParams, segment[1:])
		} else {
			key = segment
		}

		childNode, ok := currentNode.children[key]
		if !ok {
			childNode = NewRouteNode()
			currentNode.children[key] = childNode
		}

		currentNode = childNode

	}

	currentNode.methodToHandler[method] = handler
	currentNode.params = dynamicParams

	return nil
}

// FindRoute traverses a the routes tree to find the dynamic params and handler associated with a given route.
func (r *Router) FindRoute(method string, path string) (map[string]string, Handler) {
	segments := strings.Split(path, "/")
	currentNode := r.Root
	extractedParams := make([]string, 0)

	for _, segment := range segments {
		if segment == "" {
			continue
		}

		childNode, ok := currentNode.children[segment]
		if ok {
			currentNode = childNode
		} else if _, ok := currentNode.children[":"]; ok {
			childNode = currentNode.children[":"]
			extractedParams = append(extractedParams, segment)
			currentNode = childNode
		} else {
			return nil, nil
		}
	}

	params := map[string]string{}
	for i, extractedParam := range extractedParams {
		key := currentNode.params[i]
		value := extractedParam
		params[key] = value
	}

	r.logger.Info(fmt.Sprintf("%s %s", strings.ToUpper(method), path))
	return params, currentNode.methodToHandler[method]
}

func (r *Router) Get(path string, handler Handler) {
	r.AddRoute("GET", path, handler)
}

func (r *Router) Post(path string, handler Handler) {
	r.AddRoute("POST", path, handler)
}

func (r *Router) Put(path string, handler Handler) {
	r.AddRoute("PUT", path, handler)
}

func (r *Router) Delete(path string, handler Handler) {
	r.AddRoute("DELETE", path, handler)
}

func (r *Router) Patch(path string, handler Handler) {
	r.AddRoute("PATCH", path, handler)
}

func (r *Router) Options(path string, handler Handler) {
	r.AddRoute("OPTIONS", path, handler)
}

func (r *Router) Connect(path string, handler Handler) {
	r.AddRoute("CONNECT", path, handler)
}

// PrintRoutes displays all available routes
func (r *Router) PrintRoutes(node *RouteNode, indent int) {
	indentDisplay := strings.Repeat("->", indent)

	for segment, child := range node.children {
		fmt.Printf("%s %s Dynamic: %d\n", indentDisplay, segment, len(child.params))
		r.PrintRoutes(child, indent+1)
	}
}
