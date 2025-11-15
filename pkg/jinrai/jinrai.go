package jinrai

import (
	"log"
	"path"
	"strings"

	"github.com/jinrai-js/go/pkg/jinrai/jsonConfig"
)

type Static struct {
	Dist       string
	Meta       *string
	components map[string]func(props any) string
	jsonConfig.Config
	Rewrite  *func(string) string
	Assets   *bool
	Verbose  bool
	Proxy    *map[string]string
	Chashing *[]string
}

const (
	Cached = ".cached"
)

func New(dist string) (Static, error) {
	jconfig, err := jsonConfig.New(path.Join(dist, Cached))
	if err != nil {
		return Static{}, err
	}

	config := Static{
		dist,
		nil,
		make(map[string]func(props any) string),
		jconfig,
		nil,
		nil,
		false,
		nil,
		nil,
	}

	return config, nil
}

func NewX(templates string) Static {
	static, err := New(templates)
	if err != nil {
		log.Fatal(err)
	}

	return static
}

func (c *Static) ServeX(port int) {
	if err := c.Serve(port); err != nil {
		log.Fatal(err)
	}
}

func (c *Static) AddComponent(component string, handler func(props any) string) {
	c.components[component] = handler
}

func (c *Static) SetRewrite(rewrite func(path string) string) {
	c.Rewrite = &rewrite
}

func (c *Static) Debug() {
	c.Verbose = true
}

func (c *Static) ServeAssets(assets bool) {
	c.Assets = &assets
}

func (c *Static) SetProxy(proxy map[string]string) {
	c.Proxy = &proxy
}

func (c *Static) SetStringProxy(str string) {
	c.Log("+ proxy:", str)
	var proxy = make(map[string]string)

	for _, service := range strings.Split(str, ",") {
		values := strings.Split(service, "=")
		proxy[values[0]] = values[1]
	}

	c.Proxy = &proxy
}

func (c *Static) SetMeta(meta string) {
	c.Meta = &meta
}

func (c *Static) SetChashing(chashing []string) {
	c.Chashing = &chashing
}
