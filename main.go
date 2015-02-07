package main

import (
	"fmt"
	"os"
)

var saveFile string

const passwdPath = "/etc/passwd"
const groupsPath = "/etc/groups"
const shadowPath = "/etc/shadow"
const gshadowPath = "/etc/gshadow"

func init() {
	if len(os.Args) != 3 {
		fmt.Println("usage: migrate-accounts save|import filename")
	}

	command := os.Args[1]
	switch command {
	case "save":
	case "import":
		break
	default:
		fmt.Println("valid commands: save, import")
	}

	saveFile = os.Args[2]
}

func main() {

}
