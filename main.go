package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os/exec"
)

// Programs contains list of downloaded programs and the current program selected by the user.
type Programs struct {
	Progs    []string
	CurrProg string
}

type SearchProgram struct {
	Results  []string
	NoResult bool
}

type InstallProgram struct {
	ProgName string
}

func main() {
	// fmt.Println("arc", runtime.GOARCH)
	// fmt.Println("os", runtime.GOOS)
	// fmt.Println("go root", runtime.GOROOT())
	// fmt.Println("cpus", strconv.Itoa(runtime.NumCPU()))

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/progs/", ProgsHandler)
	http.HandleFunc("/currProg", CurrProgHandler)
	http.HandleFunc("/install", InstallProgHandler)
	http.HandleFunc("/search/", SearchProgHandler)
	http.HandleFunc("/upgrade", UpgradeProgHandler)
	http.HandleFunc("/uninstall", UninstallProgHandler)
	http.ListenAndServe(":80", nil)
}

// ProgsHandler lists out all the programs by the user in /usr/bin
func ProgsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("applications.html")
	lsUsr := exec.Command("ls", "/usr/bin")
	lsOutput, stderr := lsUsr.Output()
	p := Programs{Progs: make([]string, 1)}
	var count int = 0

	if stderr != nil {
		fmt.Println(stderr)
	}
	for i := 1; i < len(lsOutput); i++ {
		if lsOutput[i] != 10 {
			p.Progs[count] += string(lsOutput[i])
		} else {
			count++
			p.Progs = append(p.Progs, "")
		}
	}

	t.Execute(w, p)
}

// SearchProgHandler searches apt for specified program user searches for
func SearchProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("search.html")
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

// CurrProgHandler passes the current program selected by the user
func CurrProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("choice.html")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	p := Programs{CurrProg: m["application"][0]}
	t.Execute(w, p)
}

// UpgradeProgHandler upgrades the program to the latest version in apt
func UpgradeProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	p := Programs{CurrProg: m["application"][0]}
	exec.Command("sudo", "apt", "upgrade", "-y", m["application"][0]).Run()
	t.Execute(w, p)
}

// UninstallProgHandler removes the program that is passed in
func UninstallProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	p := Programs{CurrProg: m["application"][0]}
	exec.Command("sudo", "apt", "remove", "-y", m["application"][0]).Run()
	t.Execute(w, p)
}
