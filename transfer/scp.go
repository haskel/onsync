package transfer

import (
	"fmt"
	"log"
	"os/exec"
	config "ysync/configuration"
)

type Scp struct {
}

func (*Scp) Sync(path string, source config.Target, target config.Target) {
	log.Printf("Scp path: %s", path)

	cmd := exec.Command(
		"scp",
		fmt.Sprintf("%s/%s", source.Path, path),
		fmt.Sprintf("%s@%s:%s/%s", target.Credentials.User, target.Host, target.Path, path),
	)

	//fmt.Println(cmd.String())

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

	fmt.Println("")
}
