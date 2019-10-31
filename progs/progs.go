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
	Version  string
}

// ProgsHandler lists out all the programs by the user in /usr/bin
func ProgramHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/applications.html")
	lsUsr := exec.Command("ls", "/usr/bin")
	lsOutput, stderr := lsUsr.Output()
	p := Programs{Progs: make([]string, 1)}
	var count int = 0

	if stderr != nil {
		fmt.Println(stderr)
	}
	for i := 0; i < len(lsOutput); i++ {
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
	version := exec.Command(m["application"][0], "--version")
	verOutput, stderr := version.Output()
	p := Programs{CurrProg: m["application"][0], Version: ""}
	if stderr != nil {
		fmt.Println(stderr)
	}
	if string(verOutput) != "" {
		for i := 0; verOutput[i] != 10; i++ {
			p.Version += string(verOutput[i])
		}
	}
	t.Execute(w, p)
}

// UpgradeProgHandler upgrades the program to the latest version in apt
func UpgradeProgHandler(w http.ResponseWriter, r *http.Request) {
	t, fErr := template.ParseFiles("index.html")
	if fErr != nil {
		panic(fErr)
	}
	u, pErr := url.Parse(r.URL.String())
	if pErr != nil {
		panic(pErr)
	}
	m, qErr := url.ParseQuery(u.RawQuery)
	if qErr != nil {
		panic(qErr)
	}
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
