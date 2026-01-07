package page

import (
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app/config"
)

type Page struct {
	Cfg     config.Config
	Modules Modules
}

type ModuleType int

const (
	TitleModule ModuleType = iota
	AnnualModule
	QuarterModule
	MonthModule
	WeekModule
	DailyNotesModule
	NotesModule
)

type Modules []Module
type Module struct {
	Cfg       config.Config
	Tpl       string
	Body      interface{}
	SortIndex string
}

func SortWith(quarter int, monthInQuarter int, weekInYear int, dayInYear int, offset int) string {
	sorter := fmt.Sprintf("Q%03d", quarter+1)

	if monthInQuarter >= 0 {
		sorter += fmt.Sprintf("M%03d", monthInQuarter)

		if weekInYear >= 0 {
			sorter += fmt.Sprintf("W%03d", weekInYear)
			if dayInYear >= 0 {
				sorter += fmt.Sprintf("D%03d", dayInYear)

			}
		}
	}

	if offset >= 0 {
		sorter += fmt.Sprintf("-%05d", offset)
	}

	return sorter
}
