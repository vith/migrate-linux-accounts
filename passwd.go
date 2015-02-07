package main

type passwd struct {
	loginName         string
	encryptedPassword string
	uid               int
	gid               int
	userName          string // or comment
	homeDir           string
	shell             string
}

func NewPasswd(line string) Passwd
