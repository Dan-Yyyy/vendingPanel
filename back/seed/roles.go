package seed

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

func (s Seed) RoleSeed() {
	roles := []string{
		"Администратор",
		"Оператор",
	}

	processDictionarySeed(s, roles, repository.RolesTable)
}
