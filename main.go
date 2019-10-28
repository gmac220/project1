package main

import (
	"net/http"

	"github.com/gmac220/project1/progs"
	"github.com/gmac220/project1/search"
)

func main() {
	// fmt.Println("arc", runtime.GOARCH)
	// fmt.Println("os", runtime.GOOS)
	// fmt.Println("go root", runtime.GOROOT())
	// fmt.Println("cpus", strconv.Itoa(runtime.NumCPU()))

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/progs/", progs.ProgsHandler)
	http.HandleFunc("/currProg", progs.CurrProgHandler)
	http.HandleFunc("/install", search.InstallProgHandler)
	http.HandleFunc("/search/", search.ProgHandler)
	http.HandleFunc("/upgrade", progs.UpgradeProgHandler)
	http.HandleFunc("/uninstall", progs.UninstallProgHandler)
	http.ListenAndServe(":80", nil)
}
