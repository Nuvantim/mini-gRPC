package config

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func getOSVersion() string {
	switch runtime.GOOS {
	case "linux":
		if out, err := os.ReadFile("/proc/version"); err == nil {
			return strings.Fields(string(out))[0] + " " + strings.Fields(string(out))[2]
		}
	case "windows":
		if out, err := exec.Command("cmd", "/C", "ver").Output(); err == nil {
			return strings.TrimSpace(string(out))
		}
	case "darwin":
		if out, err := exec.Command("sw_vers", "-productVersion").Output(); err == nil {
			return "macOS " + strings.TrimSpace(string(out))
		}
	}
	return "Unknown OS"
}

func Banner() {
	// load app environment
	var appconfig, errs = GetAppConfig()
	if errs != nil {
		log.Fatal(errs)
	}

	fmt.Print(`
    __    __  __       ___ 
   / /_  / /_/ /_____ |__ \
  / __ \/ __/ __/ __ \__/ /
 / / / / /_/ /_/ /_/ / __/ 
/_/ /_/\__/\__/ .___/____/ 
             /_/            
`)

	fmt.Printf(
		"App Name     : %s\nHost Server  : %s\nPID          : %d\nRuntime      : %s\nStartup Time : %s\n",
		appconfig.AppName,
		getOSVersion(),
		os.Getpid(),
		runtime.Version(),
		time.Now().Format("2006-01-02 15:04:05"),
	)
}
