package progs

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

// ProgsHandler lists out all the programs by the user in /usr/bin
func ProgsHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/applications.html")
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

// CurrProgHandler passes the current program selected by the user
func CurrProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/choice.html")
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
	exec.Command("sudo", "apt", "purge", "-y", m["application"][0]).Run()
	t.Execute(w, p)
}
