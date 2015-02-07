package main

type gshadow struct {
	groupName         string
	encryptedPassword string
	administrators    []string // csv
	members           []string // csv
}
