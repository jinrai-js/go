package jinrai

import (
	"encoding/json"
	"net/url"

	"github.com/jinrai-js/go/pkg/jinrai/context"
	"github.com/jinrai-js/go/pkg/jinrai/jsonConfig"
)

func (c Static) Generate(url *url.URL, route *jsonConfig.Route) (string, string) {
	var context = context.New(url, c.OutDir)

	context.ExecuteRequests(c.Proxy, route.Requests, c.Rewrite)
	html := context.GetHTML(route.Content, []string{})

	export, _ := json.Marshal(context.Output.Export)

	return html, string(export)
}
