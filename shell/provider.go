package shell

import (
	"bytes"
	"fmt"
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
			log.Fatal(err)
		}
		fmt.Println(dir)
	
		path := os.Getenv("TF_INIT_PATH")
		if path != "" {
			log.Fatal("Need to provide TF_INIT_PATH")
		}

		cmd := exec.Command(path)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			log.Fatal(err)
		}

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		bb := bytes.NewBuffer([]byte{})
		bb.ReadFrom(stdout)
		for {
			s, err := bb.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Printf("%v\n", strings.Fields(s))
		}

		be := bytes.NewBuffer([]byte{})
		be.ReadFrom(stderr)
		for {
			s, err := be.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Printf("%v\n", strings.Fields(s))
		}

		return &schema.Provider{
			Schema:         map[string]*schema.Schema{},
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap:   map[string]*schema.Resource{},
		}
	}
}
