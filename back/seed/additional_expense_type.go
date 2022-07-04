package seed

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

func (s Seed) AdditionalExpenseTypeSeed() {
	additionalExpenseTypes := []string{
		"Возврат",
		"Другое",
	}

	processDictionarySeed(s, additionalExpenseTypes, repository.AdditionalExpenseTypesTable)
}
