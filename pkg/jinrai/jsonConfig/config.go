package jsonConfig

import (
	"path/filepath"

	"github.com/jinrai-js/go/internal/tools"
	"github.com/jinrai-js/go/pkg/jinrai/context"
)

type Route struct {
	Mask             string `json:"mask"`
	context.Content  `json:"content"`
	context.Requests `json:"requests"`
}

type Config struct {
	OutDir string
	Routes []Route
}

func New(outDir string) (Config, error) {
	config := Config{
		outDir,
		[]Route{},
	}
	err := tools.ReadConfig(filepath.Join(outDir, "config.json"), &config.Routes)

	return config, err
}
