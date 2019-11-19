package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	git "gopkg.in/src-d/go-git.v4"
)

const repoURL = "https://github.com/jreisinger/homepage"
const repoPath = "/tmp/homepage"
const dataPath = "/tmp/homepage/data"

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("page").Parse(tplPage)
		check(err)
		t.New("head").Parse(tplHead)
		t.New("navbar").Parse(tplNavbar)

		p := page{Title: "home", RepoURL: repoURL}
		p.Generate()

		err = t.Execute(w, p)
		check(err)
	})

	r.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("index").Parse(tplIndex)
		check(err)
		t.New("head").Parse(tplHead)
		t.New("navbar").Parse(tplNavbar)

		p := index{Title: "Index", RepoURL: repoURL, Dir: "notes"}
		p.Generate()

		err = t.Execute(w, p)
		check(err)
	})

	r.HandleFunc("/notes/{what}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		what := vars["what"]

		t, err := template.New("page").Parse(tplPage)
		check(err)
		t.New("head").Parse(tplHead)
		t.New("navbar").Parse(tplNavbar)

		p := page{Title: what, RepoURL: repoURL, Dir: "notes"}
		p.Generate()

		err = t.Execute(w, p)
		check(err)
	})

	go gitPuller()

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":5001", loggedRouter))
}

//
// Git
//

func gitClone() {
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	check(err)
}

func gitPull() {
	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(repoPath)
	check(err)

	// Get the working directory for the repository
	w, err := r.Worktree()
	check(err)

	// Pull the latest changes from the origin remote and merge into the current branch
	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	//check(err)
	log.Print(err)
}

func gitPuller() {
	for {
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			gitClone()
		}

		gitPull()
		time.Sleep(time.Second * 2)
	}
}

// check errors
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//
// Index
//

type index struct {
	Title   string
	Items   []string
	Err     error
	RepoURL string
	Dir     string
}

func (i *index) Generate() {
	dir, err := os.Open(dataPath + "/" + i.Dir)
	if err != nil {
		i.Err = err
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // -1 means return all entries
	if err != nil {
		i.Err = err
		return
	}
	var dirs []string
	for _, fi := range fileInfos {
		basename := fi.Name()
		name := strings.TrimSuffix(basename, filepath.Ext(basename))
		dirs = append(dirs, name)
	}

	sort.Strings(dirs)
	i.Items = dirs
}

//
// Page
//

type page struct {
	Title   string
	Body    template.HTML
	Err     error
	RepoURL string
	Dir     string
}

func (p *page) Generate() {
	md, err := ioutil.ReadFile(dataPath + "/" + p.Dir + "/" + p.Title + ".md")
	if err != nil {
		p.Err = err
	}

	html := markdown.ToHTML(md, nil, nil)
	p.Body = template.HTML(html)
}

//
// Templates for web pages.
//

const tplHead = `
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>{{.Title}}</title>
	<link rel="stylesheet" type="text/css" href="/css/style.css">
`

const tplNavbar = `
	<a href="/">home</a> |
	<a href="/notes">notes</a>
`

const tplIndex = `
<!DOCTYPE html>
<html>
	<head>
		{{ template "head" }}
	</head>
	<body>
		{{ template "navbar" }}
		<ul>
		{{$Dir:=.Dir}}
		{{range .Items}}
			<li><a href="{{ $Dir }}/{{ . }}">{{ . }}</a></li>
		{{end}}
		</ul>
		<a href="{{.RepoURL}}/tree/master/data/{{$Dir}}">source</a>
	</body>
</html>`

const tplPage = `
<!DOCTYPE html>
<html>
	<head>
		{{ template "head" }}
	</head>
	<body>
		{{ template "navbar" }}
        {{.Body}}
		<a href="{{.RepoURL}}/tree/master/data/{{.Dir}}/{{.Title}}.md">source</a>
	</body>
</html>`
