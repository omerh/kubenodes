package render

import (
	"kubenodes/pkg/resource"
	"strings"

	"github.com/rivo/tview"
)

// func RenderNodeTableWithHeaders() *tview.Table {
// 	table := tview.NewTable().
// 		SetBorders(true).
// 		SetBordersColor(tview.Styles.BorderColor)

// 	headers := []string{"Name", "LifeCycle", "Type", "Arch", "AZ", "Pods"}

// 	for i, header := range headers {
// 		table.SetCell(0, i, tview.NewTableCell(header).
// 			SetTextColor(tview.Styles.PrimaryTextColor).
// 			SetSelectable(false))
// 	}
// 	return table
// }

// func PopulateNodesInfo(table tview.Table, nodesInfo []resource.NodeInfo) *tview.Table {
// 	row := 1
// 	for i, nodeInfo := range nodesInfo {
// 		table.SetCell(row, 0, tview.NewTableCell(nodeInfo.Name).
// 			SetTextColor(tview.Styles.PrimaryTextColor).
// 			SetSelectable(false))
// 		table.SetCell(row, 1, tview.NewTableCell(nodeInfo.CapacityType).
// 			SetTextColor(tview.Styles.PrimaryTextColor).
// 			SetSelectable(false))
// 		table.SetCell(row, 2, tview.NewTableCell(nodeInfo.InstanceType).
// 			SetTextColor(tview.Styles.PrimaryTextColor).
// 			SetSelectable(false))
// 		table.SetCell(row, 3, tview.NewTableCell(nodeInfo.InstanceArch).
// 			SetTextColor(tview.Styles.PrimaryTextColor).
// 			SetSelectable(false))
// 		table.SetCell(row, 4, tview.NewTableCell(nodeInfo.InstanceAZ).
// 			SetTextColor(tview.Styles.PrimaryTextColor).
// 			SetSelectable(false))

// 		for _, p := range nodesInfo[i].Pod {
// 			table.SetCell(row, 5, tview.NewTableCell(p).
// 				SetTextColor(tview.Styles.PrimaryTextColor).
// 				SetSelectable(false))
// 			row++
// 		}
// 		if len(nodesInfo[i].Pod) <= 1 {
// 			row++
// 		}
// 	}
// 	return &table
// }

func NodesPodsFullRender(nodesInfo []resource.NodeInfo, compact bool) *tview.Table {
	table := tview.NewTable().
		SetBorders(true).
		SetBordersColor(tview.Styles.BorderColor)

	headers := []string{"Name", "Type", "Size", "Arch", "AZ", "Pods"}

	for i, header := range headers {
		table.SetCell(0, i, tview.NewTableCell(header).
			SetTextColor(tview.Styles.PrimaryTextColor).
			SetAlign(1).
			SetSelectable(false))
	}
	row := 1
	for i, nodeInfo := range nodesInfo {
		table.SetCell(row, 0, tview.NewTableCell(nodeInfo.Name).
			SetTextColor(tview.Styles.PrimaryTextColor).
			SetSelectable(false))
		table.SetCell(row, 1, tview.NewTableCell(nodeInfo.CapacityType).
			SetTextColor(tview.Styles.PrimaryTextColor).
			SetAlign(1).
			SetSelectable(false))
		table.SetCell(row, 2, tview.NewTableCell(nodeInfo.InstanceType).
			SetTextColor(tview.Styles.PrimaryTextColor).
			SetAlign(1).
			SetSelectable(false))
		table.SetCell(row, 3, tview.NewTableCell(nodeInfo.InstanceArch).
			SetTextColor(tview.Styles.PrimaryTextColor).
			SetAlign(1).
			SetSelectable(false))
		table.SetCell(row, 4, tview.NewTableCell(nodeInfo.InstanceAZ).
			SetTextColor(tview.Styles.PrimaryTextColor).
			SetAlign(1).
			SetSelectable(false))

		if compact {
			table.SetCell(row, 5, tview.NewTableCell(strings.Join(nodeInfo.Pod, "; ")).
				SetTextColor(tview.Styles.PrimaryTextColor).
				SetSelectable(false))
			row++
		} else {
			for j, p := range nodesInfo[i].Pod {
				table.SetCell(row, 5, tview.NewTableCell(p).
					SetTextColor(tview.Styles.PrimaryTextColor).
					SetSelectable(false))
				row++

				// If this is the last object in the slice, add a new row for next node
				if len(nodesInfo[i].Pod) == j {
					row++
				}
			}
		}
	}
	return table
}
