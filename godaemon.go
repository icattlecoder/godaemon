package godaemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var godaemon = flag.Bool("d", false, "run app as a daemon with -d=true")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *godaemon {
		args := os.Args[1:]
		i := 0
		for ; i < len(args); i++ {
			if strings.HasPrefix(args[i], "-d") {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}
