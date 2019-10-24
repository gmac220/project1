package main

import (
	"fmt"
	"os/exec"
)

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
	lsUsr := exec.Command("ls", "/usr/bin")
	lsOutput, stderr := lsUsr.Output()
	programs := make([]string, 1)
	var count int = 0
	if stderr != nil {
		fmt.Println(stderr)
	}
	for i := 0; i < len(lsOutput); i++ {
		if lsOutput[i] != 10 {
			programs[count] += string(lsOutput[i])
		} else {
			count++
			programs = append(programs, "")
		}
	}
	fmt.Println(programs)
	// cdProj1 := exec.Command("cd", "~/go/src/github.com/gmac220/project1")
	// cdProj1.Run()
}
