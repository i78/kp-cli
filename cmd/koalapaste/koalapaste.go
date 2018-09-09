package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/codecyclist/kp-cli/cmd"
	"github.com/codecyclist/kp-cli/domain"
	"github.com/codecyclist/kp-cli/runners"
)

func main() {
	appConfig := cmd.AppConfig{}
	paste := domain.Paste{}

	flag.StringVar(&paste.Title, "title", "unnamed paste", "Specify a name for your paste")
	flag.StringVar(&appConfig.Stage, "stage", "https://paste.m9d.de", "Koalapaste stage to post to (defaults to paste.m9d.de)")
	flag.Parse()

	appConfig.Mode = flag.Arg(0)
	if appConfig.Mode == "" {
		appConfig.Mode = "paste"
	}

	var content bytes.Buffer

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\r\n")
	}
	paste.Content = content.String()

	fmt.Println(appConfig, paste)

	switch appConfig.Mode {
	case "paste":
		runners.PublishPaste(appConfig, paste)
	}

	os.Exit(0)

}
