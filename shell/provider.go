package shell

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() func() *schema.Provider {
	return func() *schema.Provider {

		dir, err := os.Getwd()
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		log.Printf("[INFO] Current directory is: %s", dir)

		paths := os.Getenv("TF_INIT_SCRIPTS")
		if paths == "" {
			log.Printf("[ERROR] Need to provide TF_INIT_SCRIPTS.")
		}

		for _, token := range strings.Split(paths, ",") {
			path := strings.Trim(token, " ")

			err = os.Chmod(path, 0755)
			if err != nil {
				log.Printf("[ERROR] %s", err)
			}

			cmd := exec.Command("./" + path)

			stdout, err := cmd.StdoutPipe()
			if err != nil {
				log.Printf("[ERROR] %s", err)
			}

			stderr, err := cmd.StderrPipe()
			if err != nil {
				log.Printf("[ERROR] %s", err)
			}

			if err := cmd.Start(); err != nil {
				log.Printf("[ERROR] %s", err)
			}

			bb := bytes.NewBuffer([]byte{})
			bb.ReadFrom(stdout)
			for {
				s, err := bb.ReadString('\n')
				if err != nil {
					break
				}
				log.Printf("[INFO] %v\n", strings.Fields(s))
			}

			be := bytes.NewBuffer([]byte{})
			be.ReadFrom(stderr)
			for {
				s, err := be.ReadString('\n')
				if err != nil {
					break
				}
				log.Printf("[INFO] %v\n", strings.Fields(s))
			}
		}

		return &schema.Provider{
			Schema:         map[string]*schema.Schema{},
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap:   map[string]*schema.Resource{},
		}
	}
}
