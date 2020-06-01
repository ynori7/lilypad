package view

import (
	"html/template"
	"sync"
)

type config struct {
	mutex   sync.RWMutex
	funcMap template.FuncMap
}

var defaultConfig = config{
	funcMap: make(template.FuncMap),
}

// RegisterGlobalTemplateFuncs adds functions which can be used globally for all templates
func RegisterGlobalTemplateFuncs(f template.FuncMap) {
	defaultConfig.mutex.Lock()
	for k, v := range f {
		defaultConfig.funcMap[k] = v
	}
	defaultConfig.mutex.Unlock()
}

func getFuncMap() template.FuncMap {
	defaultConfig.mutex.RLock()
	defer defaultConfig.mutex.RUnlock()
	return defaultConfig.funcMap
}
