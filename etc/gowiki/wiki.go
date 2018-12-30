// ref <https://golang.org/doc/articles/wiki/>
package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

type Page struct {
	Title string
	Body  []byte
}

const (
	dataDir = "data"
	tmplDir = "tmpl"
)

func (p *Page) save() error {
	filename := filepath.Join(dataDir, p.Title+".txt")
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := filepath.Join(dataDir, title+".txt")
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// for updateHandler
type Templates struct {
	sync.Mutex
	T       *template.Template
	tmplDir string
}

func (t *Templates) ExecuteTemplate(w io.Writer, tmp string, v interface{}) error {
	t.Lock()
	defer t.Unlock()
	return t.T.ExecuteTemplate(w, tmp, v)
}

func (t *Templates) Update() error {
	t.Lock()
	defer t.Unlock()
	var err error
	t.T, err = template.ParseGlob(tmplDir + "/*")
	return err
}

func (t *Templates) updateHandler(w http.ResponseWriter, r *http.Request) {
	if err := t.Update(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

var templates = &Templates{
	T:       template.Must(template.ParseGlob(tmplDir + "/*")),
	tmplDir: tmplDir,
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// TODO: fix?
func frontPageHandler(w http.ResponseWriter, r *http.Request) {
	infos, err := ioutil.ReadDir(dataDir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var pages []Page
	for _, info := range infos {
		if info.Mode().IsRegular() && strings.HasSuffix(info.Name(), ".txt") {
			page, err := loadPage(strings.TrimSuffix(info.Name(), ".txt"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			pages = append(pages, *page)
		}
	}
	if err := templates.ExecuteTemplate(w, "front.html", pages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/view/FrontPage/", frontPageHandler)
	http.HandleFunc("/update/", templates.updateHandler)

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	addr := ":8080"
	log.Printf("Listen: \"%s\"\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
