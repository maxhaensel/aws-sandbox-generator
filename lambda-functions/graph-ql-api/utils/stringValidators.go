package utils

import (
	"regexp"
)

func ProofPexonMail(s string) bool {
	var validPexonMail = regexp.MustCompile(`\w+\.\w+\@pexon-consulting\.de`)
	return validPexonMail.MatchString(s)
}

func Lease_time_Input(s string) bool {
	var validPexonMail = regexp.MustCompile(`(201[4-9]|202[0-9])-(0[1-9]|1[0-2])-(0[1-9]|1[0-9]|2[0-9]|3[0-1])`)
	return validPexonMail.MatchString(s)
}
