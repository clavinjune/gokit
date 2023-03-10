package envutil

var (
	// DefaultOption is used when Option is not defined
	DefaultOption = Option{
		IsRequired:   false,
		DefaultValue: "",
	}
)

// Option sets configuration behavior
type Option struct {
	_            struct{}
	IsRequired   bool
	DefaultValue string
}
