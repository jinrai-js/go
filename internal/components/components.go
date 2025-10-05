package components

var components = make(map[string]func(props any) string)

func Add(name string, handler func(props any) string) {
	components[name] = handler
}

func Get(name string, props any) string {
	handler, exists := components[name]

	if !exists {
		return ""
	}

	return handler(props)
}
