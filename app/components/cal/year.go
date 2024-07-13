package cal

import (
	"strconv"
	"time"

	"github.com/kudrykv/latex-yearly-planner/app/components/header"
	"github.com/kudrykv/latex-yearly-planner/app/tex"
)

type Years []*Year
type Year struct {
	Number   int
	Quarters Quarters
	Weeks    Weeks
}

func NewYear(wd time.Weekday, year int) *Year {
	out := &Year{Number: year}
	out.Weeks = NewWeeksForYear(wd, out)

	for q := 1; q <= 4; q++ {
		out.Quarters = append(out.Quarters, NewQuarter(wd, out, q))
	}

	return out
}

func (y Year) Breadcrumb() string {
	return header.Items{
		header.NewIntItem(y.Number).Ref(),
		// header.NewItemsGroup(
		// 	header.NewTextItem("Q1"),
		// 	header.NewTextItem("Q2"),
		// 	header.NewTextItem("Q3"),
		// 	header.NewTextItem("Q4"),
		// ),
	}.Table(true)
}

func (y Year) SideQuarters(sel ...int) []header.CellItem {
	out := make([]header.CellItem, 0, len(y.Quarters))

	for i := len(y.Quarters) - 1; i >= 0; i-- {
		mark := false
		for _, oneof := range sel {
			if oneof == y.Quarters[i].Number {
				mark = true

				break
			}
		}

		out = append(out, header.NewCellItem(y.Quarters[i].Name()).Selected(mark))
	}

	return out
}

func (y Year) SideMonths(sel ...time.Month) []header.CellItem {
	out := make([]header.CellItem, 0, 12)

	for i := len(y.Quarters) - 1; i >= 0; i-- {
		for j := len(y.Quarters[i].Months) - 1; j >= 0; j-- {
			mon := y.Quarters[i].Months[j]
			mark := false

			for _, month := range sel {
				if month == mon.Month {
					mark = true

					break
				}
			}

			cell := header.NewCellItem(mon.ShortName()).Refer(mon.Month.String()).Selected(mark)
			out = append(out, cell)
		}
	}

	return out
}

func (y Year) HeadingMOS() string {
	return tex.ResizeBoxW(`\myLenHeaderResizeBox`, tex.Hypertarget("Calendar", strconv.Itoa(y.Number)))
}
