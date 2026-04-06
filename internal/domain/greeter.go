package domain

// Greeter represents a greeter entity
type Greeter struct {
	Name string
}

// NewGreeter creates a new greeter
func NewGreeter(name string) *Greeter {
	return &Greeter{Name: name}
}

// Greet returns a greeting message
func (g *Greeter) Greet() string {
	if g.Name == "" {
		return "Hello, World!"
	}
	return "Hello, " + g.Name + "!"
}
