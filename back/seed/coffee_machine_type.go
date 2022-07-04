package seed

import (
	"github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"
)

func (s Seed) CoffeeMachineSeed() {
	coffeeMachineTypes := []string{
		"Saeco 500",
		"Saeco 700",
		"Necta RY",
		"Necta Colibri",
	}

	processDictionarySeed(s, coffeeMachineTypes, repository.CoffeeMachineTypesTable)
}
