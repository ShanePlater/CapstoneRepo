package controllers

import (
	"time"
)

// parseTime parse and complete start time and end time.
// It will check whether start and end string are in RFC3339 format.
// If start is empty string, it will parse start as "0001-01-01T00:00:00Z".
// If end is empty string, it will parse start as current local time in RFC3339 format.
func parseTime(start, end *string) error {
	// Check start time.
	t, err := time.Parse(time.RFC3339, *start)
	if *start != "" && err != nil {
		return err
	}

	// Format start time.
	*start = t.Format(time.RFC3339)

	// Check end time.
	t, err = time.Parse(time.RFC3339, *end)
	if *end != "" && err != nil {
		return err
	}

	// Parse end as current local time in RFC3339 format if it is empty string.
	if *end == "" {
		*end = time.Now().Format(time.RFC3339)
		return nil
	}

	// Format end time.
	*end = t.Format(time.RFC3339)

	return nil
}

func authCheck(username, password string) error {
	/* 	bindusername := "readonly"
	   	bindpassword := "password"

	   	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "ldap.example.com", 389))
	   	if err != nil {
	   		return err
	   	}
	   	defer l.Close()

	   	// Reconnect with TLS
	   	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	   	if err != nil {
	   		return err
	   	}

	   	// First bind with a read only user
	   	err = l.Bind(bindusername, bindpassword)
	   	if err != nil {
	   		return err
	   	}

	   	// Search for the given username
	   	searchRequest := ldap.NewSearchRequest(
	   		"dc=example,dc=com",
	   		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
	   		fmt.Sprintf("(&(objectClass=organizationalPerson)&(uid=%s))", username),
	   		[]string{"dn"},
	   		nil,
	   	)

	   	sr, err := l.Search(searchRequest)
	   	if err != nil {
	   		return err
	   	}

	   	if len(sr.Entries) != 1 {
	   		return errors.New("User does not exist or too many entries returned")
	   	}

	   	userdn := sr.Entries[0].DN

	   	// Bind as the user to verify their password
	   	err = l.Bind(userdn, password)
	   	if err != nil {
	   		return err
	   	}

	   	// Rebind as the read only user for any futher queries
	   	err = l.Bind(bindusername, bindpassword)
	   	if err != nil {
	   		return err
	   	} */

	return nil
}
