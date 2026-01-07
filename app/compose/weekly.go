package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

const WEEK_INDEX_SIZE = 10000

func Weekly(cfg config.Config, tpls []string) (page.Modules, error) {
	modules := make(page.Modules, 0, 53)
	year := cal.NewYear(cfg.WeekStart, cfg.Year)

	for _, week := range year.Weeks {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Year":            year,
				"Week":            week,
				"Breadcrumb":      week.Breadcrumb(),
				"HeadingMOS":      week.HeadingMOS(),
				"SideQuarters":    year.SideQuarters(week.Quarters.Numbers()...),
				"SideMonths":      year.SideMonths(week.Months.Months()...),
				"BreadcrumbExtra": week.PrevNext().WithTopRightCorner(cfg.ClearTopRightCorner),
				"DottedExtra":     dottedExtra(cfg.ClearTopRightCorner, false, false, false, nil, 0),
			},

			SortIndex: page.SortWith(
				week.Quarters.Numbers()[0],
				int(week.Months[0].Month)%4,
				week.WeekNumberInt(),
				-1,
				0,
			),
		})
	}

	return modules, nil
}
