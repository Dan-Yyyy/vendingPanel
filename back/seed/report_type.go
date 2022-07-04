package seed

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

func (s Seed) ReportTypeSeed() {
	reportTypes := []string{
		"Поломка",
		"Обслуживание",
	}

	processDictionarySeed(s, reportTypes, repository.ReportTypesTable)
}
