package main

import (
	"fmt"

	"github.com/kamran-hassan/ciphersuitcase/utils/readconfig"
)

func main() {
	fmt.Println("Hi")
	configuration := readconfig.GetConf()

	fmt.Println(configuration.Backup.Directories[1])

}
