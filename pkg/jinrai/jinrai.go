package jinrai

import (
	"github.com/jinrai-js/go/internal/server"
	"github.com/jinrai-js/go/pkg/jinrai/jsonConfig"
)

type Static struct {
	Templates  string
	Api        string
	Meta       string
	Components map[string]func(props any) string
	jsonConfig.Config
}

func New(templates, api, meta string) (Static, error) {
	jconfig, err := jsonConfig.New(templates)
	if err != nil {
		return Static{}, err
	}

	config := Static{
		templates,
		api,
		meta,
		make(map[string]func(props any) string),
		jconfig,
	}

	return config, nil
}

func (c Static) Serve(port int) error {
	return server.Run(port, c.Handler)
}

func (c Static) AddComponent(component string, handler func(props any) string) {
	c.Components[component] = handler
}
