package main

type shadow struct {
	loginName                string
	encryptedPassword        string
	dateOfLastPasswordChange epochDays
	minPasswordAge           days
	maxPasswordAge           days
	passwordWarningPeriod    days
	passwordInactivityPeriod days
	accountExpirationDate    epochDays
	reserved                 string
}
