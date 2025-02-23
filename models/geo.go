package models

type GeoNode struct {
	Name     string
	Children map[string]*GeoNode
}

// retrieves an existing child node or creates a new one.
func GetOrCreateChild(parent *GeoNode, name string) *GeoNode {
	if child, exists := parent.Children[name]; exists {
		return child
	}
	child := &GeoNode{Name: name, Children: make(map[string]*GeoNode)}
	parent.Children[name] = child
	return child
}
