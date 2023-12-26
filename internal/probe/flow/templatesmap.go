package flow

import "sync"

type TemplateMap struct {
	tMap map[uint16]*Template
	mut  sync.RWMutex
}

func NewTemplateMap() *TemplateMap {
	return &TemplateMap{
		tMap: make(map[uint16]*Template),
	}
}

func (t *TemplateMap) Add(template *Template) {
	t.mut.Lock()
	(t.tMap)[template.TemplateID] = template
	t.mut.Unlock()
}

func (t *TemplateMap) Get(templateID uint16) (*Template, bool) {
	t.mut.RLock()
	template, ok := (t.tMap)[templateID]
	t.mut.RUnlock()
	return template, ok
}
