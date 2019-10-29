package main

import (
	"net/http"

	"github.com/gmac220/project1/progs"
	"github.com/gmac220/project1/search"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/progs/", progs.ProgramHandler)
	http.HandleFunc("/currProg", progs.CurrProgHandler)
	http.HandleFunc("/install", search.InstallProgHandler)
	http.HandleFunc("/search/", search.ProgHandler)
	http.HandleFunc("/upgrade", progs.UpgradeProgHandler)
	http.HandleFunc("/uninstall", progs.UninstallProgHandler)
	http.ListenAndServe(":80", nil)
}
