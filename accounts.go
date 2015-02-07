package main

type accounts struct {
	users   []passwd
	groups  []group
	shadow  []shadow
	gshadow []gshadow
}

func AccountsFromEtc() *accounts {
	accs := new(accounts)
	accs.users = filterByUid(getSystemPasswd())
	accs.groups = filterByMembers(readGroups(), accs.users)
	accs.shadow = filterByUsers(readShadow(), accs.users)
	accs.gshadow = filterByGids(readGshadow(), accs.groups)
}

func getSystemPasswd() []passwd {
	var users []passwd
	file := ioutil.ReadFile(passwdPath)
	lines := strings.Split(file, "\n")
	for _, line := lines {
		
	}
}
