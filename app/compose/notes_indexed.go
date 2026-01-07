package compose

import (
	"github.com/kudrykv/latex-yearly-planner/app/components/cal"
	"github.com/kudrykv/latex-yearly-planner/app/components/note"
	"github.com/kudrykv/latex-yearly-planner/app/components/page"
	"github.com/kudrykv/latex-yearly-planner/app/config"
)

func NotesIndexed(cfg config.Config, tpls []string) (page.Modules, error) {
	index := note.NewIndex(cfg.Year, cfg.Layout.Numbers.NotesOnPage, cfg.Layout.Numbers.NotesIndexPages)
	year := cal.NewYear(cfg.WeekStart, cfg.Year)
	modules := make(page.Modules, 0, 1)

	for idx, indexPage := range index.Pages {
		modules = append(modules, page.Module{
			Cfg: cfg,
			Tpl: tpls[0],
			Body: map[string]interface{}{
				"Notes":           indexPage,
				"Breadcrumb":      indexPage.Breadcrumb(cfg.Year, idx),
				"HeadingMOS":      indexPage.HeadingMOS(idx+1, len(index.Pages)),
				"SideQuarters":    year.SideQuarters(0),
				"SideMonths":      year.SideMonths(0),
				"BreadcrumbExtra": index.PrevNext(idx).WithTopRightCorner(cfg.ClearTopRightCorner),
				"DottedExtra":     dottedExtra(cfg.ClearTopRightCorner, false, true, false, nil, 0),
			},
			SortIndex: page.SortWithFooter(
				2,
				"notes-idx",
				idx,
			),
		})
	}

	for idxPage, notes := range index.Pages {
		for noteIndex, nt := range notes {
			modules = append(modules, page.Module{
				Cfg: cfg,
				Tpl: tpls[1],
				Body: map[string]interface{}{
					"Note":         nt,
					"Breadcrumb":   nt.Breadcrumb(),
					"HeadingMOS":   nt.HeadingMOS(idxPage),
					"SideQuarters": year.SideQuarters(0),
					"SideMonths":   year.SideMonths(0),
					"BreadcrumbExtra": nt.
						PrevNext(cfg.Layout.Numbers.NotesOnPage * cfg.Layout.Numbers.NotesIndexPages).
						WithTopRightCorner(cfg.ClearTopRightCorner),
					"DottedExtra": dottedExtra(cfg.ClearTopRightCorner, false, false, idxPage == 0 && noteIndex == 0, nil, idxPage+1),
				},
				SortIndex: page.SortWithFooter(
					3,
					"notes",
					100*(idxPage)+(noteIndex),
				),
			})
		}
	}

	return modules, nil
}
