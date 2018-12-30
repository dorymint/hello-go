// +build ignore

package main

import (
	"io/ioutil"
	"strings"
	"sync"
)

// PageCache cache on memory
// can convert loadPage to *PageCache.LoadPage
// TODO: consider to use?
type PageCache struct {
	sync.Mutex
	cache   map[string]*Page // map[Page.Title]*Page
	dataDir string
}

func (p *PageCache) Update() error {
	infos, err := ioutil.ReadDir(p.dataDir)
	if err != nil {
		return err
	}
	p.Lock()
	defer p.Unlock()
	// reset is need?
	//p.cache = make(map[string]*Page)
	for _, info := range infos {
		if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), ".txt") {
			page, err := loadPage(strings.TrimSuffix(info.Name(), ".txt"))
			if err != nil {
				return err
			}
			p.cache[info.Name()] = page
		}
	}
	return nil
}

// need call at saveHandler
func (p *PageCache) Save(page *Page) error {
	p.Lock()
	defer p.Unlock()
	err := page.save()
	if err != nil {
		return err
	}
	p.cache[page.Title] = page
	return nil
}

func (p *PageCache) LoadPage(title string) (*Page, error) {
	page, ok := p.cache[title]
	if ok {
		return page, nil
	}
	var err error
	page, err = loadPage(title)
	if err != nil {
		return nil, err
	}
	p.Lock()
	p.cache[title] = page
	p.Unlock()
	return page, nil
}

var pages = func() *PageCache {
	pages := &PageCache{
		cache:   make(map[string]*Page),
		dataDir: dataDir,
	}
	if err := pages.Update(); err != nil {
		panic(err)
	}
	return pages
}()
