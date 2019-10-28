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
	// memStats := new(runtime.MemStats)
	// runtime.ReadMemStats(memStats)
	// fmt.Println("arc", runtime.GOARCH)
	// fmt.Println("os", runtime.GOOS)
	// fmt.Println("go root", runtime.GOROOT())
	// fmt.Println("cpus", strconv.Itoa(runtime.NumCPU()))
	// cdUsr := exec.Command("cd", "..")
	// cdUsr.Stdin = os.Stdin
	// _, err := cdUsr.Output()
	// cdUsr.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cdProj1 := exec.Command("cd", "~/go/src/github.com/gmac220/project1")
	// cdProj1.Run()
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/progs/", ProgsHandler)
	http.HandleFunc("/currProg", CurrProgHandler)
	http.HandleFunc("/install", InstallProgHandler)
	http.HandleFunc("/search/", SearchProgHandler)
	http.HandleFunc("/update", UpdateProgHandler)
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

// SearchProgHandler searches apt
func SearchProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("search.html")
	s := SearchProgram{Results: make([]string, 1), NoResult: false}
	searchVal := r.FormValue("pname")
	var count int = 0
	var skipIntro int = 30
	// var carriageReturn byte = 10
	if searchVal != "" {
		usrSearch := exec.Command("apt", "search", searchVal)
		searchOutput, stderr := usrSearch.Output()
		if stderr != nil {
			fmt.Println(stderr)
		}
		if len(searchOutput) > skipIntro {
			s.NoResult = true
		}
		for i := skipIntro; i < len(searchOutput); i++ {
			if searchOutput[i] != 10 && (searchOutput[i] != 10 && searchOutput[i+1] != 10) {
				s.Results[count] += string(searchOutput[i])
			} else if searchOutput[i] == 10 && searchOutput[i+1] == 10 {
				i++
			} else if searchOutput[i] == 10 && searchOutput[i+1] == 32 {
				s.Results[count] += string(searchOutput[i])
			} else {
				count++
				//fmt.Println("This is count" + strconv.Itoa(count))
				s.Results = append(s.Results, "")
			}
		}
		// fmt.Println(searchOutput)
	}

	// fmt.Println(string(searchOutput))
	// fmt.Println(len(searchOutput))

	//fmt.Println(s.Results)
	t.Execute(w, s)
}

func InstallProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("install.html")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	installProg := InstallProgram{ProgName: m["pname"][0]}
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

func UpdateProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("update.html")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	p := Programs{CurrProg: m["application"][0]}
	//exec.Command("apt", "update", m["application"][0])
	t.Execute(w, p)
}

func UninstallProgHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("uninstall.html")
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	p := Programs{CurrProg: m["application"][0]}
	//exec.Command("apt", "uninstall", m["application"][0])
	t.Execute(w, p)
}
