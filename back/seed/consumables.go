package seed

import "github.com/Dan-Yyyy/vendingPanel.git/pkg/repository"

func (s Seed) ConsumableSeed() {
	reportTypes := []string{
		"Шоколад",
		"Молоко",
		"Кофе в зернах",
		"Кофе растворимый",
		"Ваниль",
		"Стаканчики",
		"Палочки",
		"Сахар",
	}

	processDictionarySeed(s, reportTypes, repository.ConsumablesTable)
}
