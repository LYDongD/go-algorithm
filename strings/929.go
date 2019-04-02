package main

import "fmt"

func numUniqueEmails(emails []string) int {
    if len(emails) == 0 {
	return 0
    }

    mailDic := make(map[string]int)
    count := 0
    for _, email := range emails {

	parsedEmail, ok := parseEmail(email)

	if !ok {
	    return count
	}

	if mailDic[parsedEmail] == 0 {
	    count++
	    mailDic[parsedEmail]++
	}
    }

    return count
}

func parseEmail(email string) (string, bool) {
    local, domain := splitEmail(email)
    if len(local) == 0  {
	return email, false
    }
    local = parseLocal(local)
    return local + "@" + domain, true
}

func splitEmail(email string) (string,string) {
    for index, char := range email {
	if char == '@' && index  < len(email) - 1 {
	    return email[0:index], email[index+1:len(email)]
	}
    }

    return "", ""
}

func parseLocal(local string) string {
    runes := []rune{}
    for _, char := range local {
	if char == '+' {
	    break
	}

	if char == '.' {
	    continue
	}

	runes = append(runes, char)
    }

    return string(runes)
}

func main() {
    fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}))
}
