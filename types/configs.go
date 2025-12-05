package types

type Config struct {
	Name    string
	IsDir   bool
	PreCmd  []string
	PostCmd []string
	// createdAt and modifiedAt
}

var DefaultConfig = map[string]any{
	"name":    "default-template-name",
	"isDir":   false,
	"preCmd":  []string{},
	"postCmd": []string{},
}
