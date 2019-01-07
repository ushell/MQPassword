package MQPassword

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	_ "image/png"
	"log"
	"strconv"
)

type modelHandler struct {
	RowNums 		int
	ColumnNums 		int
	Data    		InputContent
}

func newModelHandler() *modelHandler {
	m := new(modelHandler)
	//初始化
	m.RowNums = 0
	//列数
	m.ColumnNums = 6
	return m
}

//标题类型
func (mh *modelHandler) ColumnTypes(m *ui.TableModel) []ui.TableValue {
	return []ui.TableValue{
		ui.TableString(""), // column 0 text
		ui.TableString(""), // column 1 text
		ui.TableString(""), // column 2 text
		ui.TableString(""), // column 3 text
		ui.TableString(""), // column 4 button text
		ui.TableString(""), // column 5 button text
	}
}

//行数
func (mh *modelHandler) NumRows(m *ui.TableModel) int {
	return mh.RowNums
}

//数据设置
func (mh *modelHandler) CellValue(m *ui.TableModel, row, column int) ui.TableValue {
	if column == 0 {
		index := row + 1
		return ui.TableString(strconv.Itoa(index))
	}
	if column == 1 {
		return ui.TableString(mh.Data.Name)
	}
	if column == 2 {
		return ui.TableString(mh.Data.CreateAt)
	}
	if column == 3 {
		return ui.TableString(mh.Data.Domain)
	}
	if column == 4 {
		return ui.TableString("查看")
	}
	if column == 5 {
		return ui.TableString("删除")
	}

	return ui.TableString("")
}

//事件
func (mh *modelHandler) SetCellValue(m *ui.TableModel, row, column int, value ui.TableValue) {
	//按钮事件
	if column == 4 {
		item := DATABASE.Data[row]
		showItemWindow(item.Name)
	}

	if column == 5 {
		info := fmt.Sprintf("行号:%s", strconv.Itoa(row))
		ui.MsgBox(mainwin, "数据位置", info)
	}

	//log.Fatal("当前表格数据:", value)
}

func setupTable(data *Model) *ui.Table {
	//表格数据
	mh := newModelHandler()
	//数据行数
	mh.RowNums = DATABASE.Count

	//log.Fatal(DATABASE.Count)

	//初始化model
	model := ui.NewTableModel(mh)

	//渲染数据
	for index, item := range data.List() {
		mh.Data = item
		log.Println(item)
		for i:=0; i< mh.ColumnNums; i++ {
			mh.CellValue(model, index, i)
		}
	}

	//设置表格
	table := ui.NewTable(&ui.TableParams{Model: model, RowBackgroundColorModelColumn: 3,})

	table.AppendTextColumn("序号", 0, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("名称", 1, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("时间", 2, ui.TableModelColumnNeverEditable, nil)
	table.AppendTextColumn("域名", 3, ui.TableModelColumnNeverEditable, nil)
	//按钮
	table.AppendButtonColumn("操作", 4, ui.TableModelColumnAlwaysEditable)
	table.AppendButtonColumn("操作", 5, ui.TableModelColumnAlwaysEditable)

	return table
}
