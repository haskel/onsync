package transfer

import (
	"fmt"
	"log"
	"os/exec"
	config "ysync/configuration"
)

type Rsync struct {
}

func (*Rsync) Sync(paths []string, source config.Target, target config.Target, syncConfig config.Sync) {
	log.Printf("Rsync")

	var excludes []string
	for _, dir := range syncConfig.Directories.Excluded {
		excludes = append(excludes, fmt.Sprintf("--exclude=%s", dir))
	}

	args := []string{
		"-azv",
		source.Path + "/",
		fmt.Sprintf("%s@%s:%s", target.Credentials.User, target.Host, target.Path),
	}

	args = append(args, excludes...)

	cmd := exec.Command("rsync", args...)

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Fatal(err)
	}
	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	//output, _ := cmd.CombinedOutput()
	//fmt.Println(string(output))

	//
	//err := cmd.Run()
	//
	//if err != nil {
	//	log.Println(err)
	//}
}
