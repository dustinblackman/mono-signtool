package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	var keyPath, password, filePath string
	for idx, param := range os.Args {
		if param == "/f" {
			keyPath = os.Args[idx+1]
		}
		if param == "/p" {
			password = os.Args[idx+1]
		}
		if idx == (len(os.Args) - 1) {
			filePath = param
		}
	}

	if keyPath == "" {
		log.Fatal("keyPath not set")
	}
	if password == "" {
		log.Fatal("password not set")
	}
	if filePath == "" {
		log.Fatal("exe file path not set")
	}

	for idx, hash := range []string{"sha1", "sha256"} {
		args := []string{
			"/c",
			"start",
			"/wait",
			"/unix",
			"/usr/local/bin/osslsigncode", // Path for osslsigncode on OSX/Linux
			"-in",
			filePath,
			"-out",
			filePath + ".signed",
			"-t",
			"http://timestamp.verisign.com/scripts/timstamp.dll",
			"-pkcs12",
			keyPath,
			"-pass",
			password}

		if idx == 1 {
			args = append(args, "-nest")
		}
		args = append(args, "-h", hash)

		cmd := exec.Command("cmd", args...)
		cmd.Start()
		for {
			if _, err := os.Stat(filePath + ".signed"); err == nil {
				break
			}
		}

		err := cmd.Process.Kill()
		if err != nil {
			log.Fatal(err)
		}
		err = os.Remove(filePath)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Rename(filePath+".signed", filePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
