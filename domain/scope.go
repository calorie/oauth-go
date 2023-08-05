package domain

import (
	"strings"
)

type Scope struct {
	Name string
}

func FilterScope(scope string) *[]Scope {
	m := map[string]bool{}
	for _, s := range *scopes() {
		m[s.Name] = true
	}

	scopes := []Scope{}
	for _, s := range strings.Fields(scope) {
		if m[s] {
			scopes = append(scopes, Scope{Name: s})
		}
	}

	return &scopes
}

func scopes() *[]Scope {
	return &[]Scope{
		{
			Name: "read",
		},
		{
			Name: "write",
		},
	}
}
