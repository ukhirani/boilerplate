package types

type Config struct {
	name    string
	isDir   bool
	preCmd  []string
	postCmd []string
	// createdAt and modifiedAt
}

var DefaultConfig = map[string]any{
	"name":    "default-template-name",
	"isDir":   false,
	"preCmd":  []string{},
	"postCmd": []string{},
}
