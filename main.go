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
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/progs/", ProgsHandler)
	http.HandleFunc("/currProg", CurrProgHandler)
	http.ListenAndServe(":80", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/progs/\"> Downloaded applications</a>")
}

// ProgsHandler lists out all the programs by the user in /usr/bin
func ProgsHandler(w http.ResponseWriter, r *http.Request) {
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
	t, _ := template.ParseFiles("applications.html")
	t.Execute(w, p)
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
