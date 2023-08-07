package repository

import (
	"strings"

	"github.com/calorie/oauth-go/domain"
)

type ScopeRepository struct {
}

func NewScopeRepositoty() *ScopeRepository {
	return &ScopeRepository{}
}

func (r *ScopeRepository) FilterScope(scope string) *[]domain.Scope {
	m := map[string]bool{}
	for _, s := range *r.scopes() {
		m[s.Name] = true
	}

	scopes := []domain.Scope{}
	for _, s := range strings.Fields(scope) {
		if m[s] {
			scopes = append(scopes, domain.Scope{Name: s})
		}
	}

	return &scopes
}

func (r *ScopeRepository) JoinScopes(scopes *[]domain.Scope) string {
	s := []string{}
	for _, scope := range *scopes {
		s = append(s, scope.Name)
	}
	return strings.Join(s, " ")
}

func (r *ScopeRepository) scopes() *[]domain.Scope {
	return &[]domain.Scope{
		{
			Name: "read",
		},
		{
			Name: "write",
		},
	}
}
