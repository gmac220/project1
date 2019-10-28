package search

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os/exec"
)

type SearchProgram struct {
	Results  []string
	NoResult bool
}

type InstallProgram struct {
	ProgName string
}

// ProgHandler searches apt for specified program user searches for
func ProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/search.html")
	s := SearchProgram{Results: make([]string, 1), NoResult: false}
	searchVal := r.FormValue("pname")
	var count int = 0
	var skipIntro int = 31
	var lineFeed byte = 10
	var spaceChar byte = 32

	if searchVal != "" {
		usrSearch := exec.Command("apt", "search", searchVal)
		searchOutput, stderr := usrSearch.Output()
		if stderr != nil {
			fmt.Println(stderr)
		}
		if len(searchOutput) > skipIntro {
			s.NoResult = true
			for i := skipIntro; i < len(searchOutput); i++ {
				if searchOutput[i] != lineFeed {
					s.Results[count] += string(searchOutput[i])
				} else if searchOutput[i] != lineFeed && searchOutput[i+1] != lineFeed {
					s.Results[count] += string(searchOutput[i])
				} else if searchOutput[i] == lineFeed && searchOutput[i+1] == lineFeed {
					i++
					count++
					s.Results = append(s.Results, "")
				} else if searchOutput[i] == lineFeed && searchOutput[i+1] == spaceChar {
					s.Results[count] += string(searchOutput[i])
				}
			}
		} else {
			s.NoResult = false
		}

		// fmt.Println(searchOutput)
		// fmt.Println(string(searchOutput))
		// fmt.Println(len(searchOutput))
	}
	t.Execute(w, s)
}

// InstallProgHandler installs the program that is passed in
func InstallProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	installProg := InstallProgram{ProgName: m["pname"][0]}
	exec.Command("sudo", "apt", "install", "-y", m["pname"][0]).Run()
	t.Execute(w, installProg)
}
