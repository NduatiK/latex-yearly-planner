package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func Monthly(cfg config.Config, tpls []string) (page.Modules, error) {
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 12)

	for quarterIndex, quarter := range year.Quarters {
		for monthIndex, month := range quarter.Months {
			modules = append(modules, page.Module{
				Cfg: cfg,
				Tpl: tpls[0],
				Body: map[string]interface{}{
					"Year":            year,
					"Quarter":         quarter,
					"Month":           month,
					"Breadcrumb":      month.Breadcrumb(),
					"HeadingMOS":      month.HeadingMOS(),
					"SideQuarters":    year.SideQuarters(quarter.Number),
					"SideMonths":      year.SideMonths(month.Month),
					"BreadcrumbExtra": month.PrevNext().WithTopRightCorner(cfg.ClearTopRightCorner),
					"DottedExtra":     dottedExtra(cfg.ClearTopRightCorner, false, false, false, nil, 0),
				},
				SortIndex: page.SortWith(
					quarterIndex,
					monthIndex,
					-2,
					-1,
					0,
				),
			})
		}
	}

	return modules, nil
}
