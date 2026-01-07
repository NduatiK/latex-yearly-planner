package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

var Daily = DailyStuff("", "", page.WeekModule, 0)
var DailyReflect = DailyStuff("Reflect", "Reflect", page.WeekModule, 1)
var DailyNotes = DailyStuff("More", "Notes", page.DailyNotesModule, 0)

func DailyStuff(prefix, leaf string, group page.ModuleType, offset int) func(cfg config.Config, tpls []string) (page.Modules, error) {
	return func(cfg config.Config, tpls []string) (page.Modules, error) {
		year := cal.NewYear(cfg.WeekStart, cfg.Year)
		modules := make(page.Modules, 0, 366)

		for quarterIndex, quarter := range year.Quarters {
			for monthIndex, month := range quarter.Months {
				for _, week := range month.Weeks {
					for _, day := range week.Days {
						if day.Time.IsZero() {
							continue
						}

						modules = append(modules, page.Module{
							Cfg: cfg,
							Tpl: tpls[0],
							Body: map[string]interface{}{
								"Year":            year,
								"Quarter":         quarter,
								"Month":           month,
								"Week":            week,
								"Day":             day,
								"Breadcrumb":      day.Breadcrumb(prefix, leaf, cfg.ClearTopRightCorner && len(leaf) > 0),
								"HeadingMOS":      day.HeadingMOS(prefix, leaf),
								"SideQuarters":    year.SideQuarters(day.Quarter()),
								"SideMonths":      year.SideMonths(day.Month()),
								"BreadcrumbExtra": day.PrevNext(prefix).WithTopRightCorner(cfg.ClearTopRightCorner),
								"DottedExtra":     dottedExtra(cfg.ClearTopRightCorner, false, false, false, week, 0),
							},
							SortIndex: page.SortWith(
								quarterIndex,
								monthIndex,
								week.WeekNumberInt(),
								day.Time.Day(),
								offset,
							),
						})
					}
				}
			}
		}

		return modules, nil
	}
}
