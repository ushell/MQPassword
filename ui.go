package MQPassword

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

//主进程
var mainwin *ui.Window

//Event
var uiEvent *Event

//主页
func makeHomeWindow() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	newInputBotton := ui.NewButton("新增")
	newInputBotton.OnClicked(func(button *ui.Button) {
		//新窗口
		inputWin := ui.NewWindow(APP_NAME, 480, 320, false)

		inputWin.OnClosing(func(*ui.Window) bool {
			return true
		})

		tab := ui.NewTab()
		inputWin.SetChild(tab)
		inputWin.SetMargined(true)

		tab.Append("", makeInputWindow(&InputContent{}, inputWin))

		inputWin.Show()
	})
	hbox.Append(newInputBotton, false)

	vbox.Append(hbox, false)

	//分割线
	vbox.Append(ui.NewHorizontalSeparator(), false)

	group := ui.NewGroup("数据栏")
	group.SetMargined(true)
	group.SetChild(ui.NewNonWrappingMultilineEntry())

	//数据区域
	table := setupTable(new(Model))
	group.SetChild(table)
	vbox.Append(group, true)

	//刷新数据
	//newRefreshButton := ui.NewButton("刷新")
	//newRefreshButton.OnClicked(func(button *ui.Button) {
	//	//重新渲染数据
	//	bbox := ui.NewHorizontalBox()
	//	model := new(Model)
	//	for _,item := range model.List() {
	//		bbox.Append(ui.NewButton(item.Name), false)
	//	}
	//	group.SetChild(bbox)
	//})
	//hbox.Append(newRefreshButton, false)

	return vbox
}

func showItemWindow(name string) {
	//获取数据
	data := uiEvent.PasswordItem(name)

	//窗口生成,数据渲染
	inputWin := ui.NewWindow(APP_NAME, 480, 320, false)

	inputWin.OnClosing(func(*ui.Window) bool {
		return true
	})

	tab := ui.NewTab()
	inputWin.SetChild(tab)
	inputWin.SetMargined(true)

	tab.Append("", makeInputWindow(&data, inputWin))

	inputWin.Show()
}

func makeDemo() ui.Control {
	Separator := ui.NewHorizontalSeparator()
	Separator_label_l := ui.NewLabel("left")
	Separator_label_l.SetText("节点")


	Separator_label_r := ui.NewLabel("right")

	Separator_div := ui.NewHorizontalBox()
	Separator_div.Append(Separator_label_l, true)

	Separator_div.Append(Separator, false)

	Separator_div.Append(Separator_label_r, true)

	Separator_div.SetPadded(true)

	div := ui.NewVerticalBox()
	box := ui.NewHorizontalBox()

	group := ui.NewGroup("数据栏")
	group.SetMargined(true)
	group.SetChild(Separator_div)

	box.Append(group, true)
	div.Append(box, true)

	return div
}

//数据存储
func makeInputWindow(content *InputContent, win *ui.Window) ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	group := ui.NewGroup("输入区域")
	group.SetMargined(true)
	vbox.Append(group, false)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	input := ui.NewForm()
	input.SetPadded(true)
	group.SetChild(input)

	name := ui.NewEntry()
	if content.Name != "" {
		name.SetText(content.Name)
	}

	account := ui.NewEntry()
	if content.Account != "" {
		account.SetText(content.Account)
	}

	password := ui.NewEntry()
	if content.Password != "" {
		password.SetText(content.Password)
	}

	domain := ui.NewEntry()
	if content.Domain != "" {
		domain.SetText(content.Domain)
	}

	comment := ui.NewNonWrappingMultilineEntry()
	if content.Comment != "" {
		comment.SetText(content.Comment)
	}

	input.Append("名称", name, false)
	input.Append("账户", account, false)
	input.Append("密码", password, false)
	input.Append("域名", domain, false)
	//富文本
	input.Append("备注", comment, true)

	button := ui.NewButton("确认")
	button.OnClicked(func(button *ui.Button) {
		//存储数据
		data := InputContent{Name:name.Text(),Account:account.Text(), Password:password.Text(), Domain:domain.Text(), Comment:comment.Text()}

		if uiEvent.PasswordSave(data) == false {
			ui.MsgBox(win, "操作提示", "操作失败")
		} else {
			ui.MsgBox(win, "操作提示", "操作成功")
		}
		//关闭输入框
		win.Destroy()
	})

	vbox.Append(button, false)

	return vbox
}

//设置
func makeSettingWindow() ui.Control {
	//水平
	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	//垂直
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	group := ui.NewGroup("输入区域")
	group.SetMargined(true)
	vbox.Append(group, false)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	input := ui.NewForm()
	input.SetPadded(true)
	input.Append("名称", ui.NewEntry(), false)
	input.Append("名称", ui.NewEntry(), false)
	input.Append("名称", ui.NewEntry(), false)
	input.Append("名称", ui.NewEntry(), false)
	group.SetChild(input)

	hbox.Append(vbox, true)

	//分界线
	vbox.Append(ui.NewVerticalSeparator(), false)

	vbox = ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, true)

	grid := ui.NewGrid()
	grid.SetPadded(true)
	vbox.Append(grid, true)
	vbox.Append(ui.NewButton("确认"), false)

	dataButton := ui.NewButton("数据源")
	entry := ui.NewEntry()
	entry.SetReadOnly(true)
	dataButton.OnClicked(func(*ui.Button) {
		filename := ui.OpenFile(mainwin)
		if filename == "" {
			filename = ""
		}
		entry.SetText(filename)
	})

	grid.Append(dataButton, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(entry, 1, 0, 1, 1, true, ui.AlignFill, false, ui.AlignFill)

	return hbox
}

//密码修改
func makePasswordWindow() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	group := ui.NewGroup("")
	group.SetMargined(true)
	vbox.Append(group, false)

	group.SetChild(ui.NewNonWrappingMultilineEntry())

	input := ui.NewForm()
	input.SetPadded(true)
	group.SetChild(input)

	originPassword := ui.NewPasswordEntry()
	password := ui.NewPasswordEntry()
	confirmPassword := ui.NewPasswordEntry()

	input.Append("原密码", originPassword, false)
	input.Append("新密码", password, false)
	input.Append("确认密码", confirmPassword, false)

	originPassword.OnChanged(func(entry *ui.Entry) {
		if entry.Text() == "" {
			ui.MsgBox(mainwin, "操作提示", "原密码为空")
		}
	})
	password.OnChanged(func(entry *ui.Entry) {
		if entry.Text() == "" {
			ui.MsgBox(mainwin, "操作提示", "密码为空")
		}
	})
	confirmPassword.OnChanged(func(entry *ui.Entry) {
		if password.Text() != entry.Text() {
			ui.MsgBox(mainwin, "操作提示", "两次密码不一致")
		}
	})

	//确认
	submit := ui.NewButton("确认")
	submit.OnClicked(func(button *ui.Button) {
		if password.Text() == "" {
			ui.MsgBox(mainwin, "操作提示", "密码为空")
		}
		result, msg := uiEvent.PasswordChange(password.Text())
		if result == true {
			ui.MsgBox(mainwin, "操作提示", "操作成功")
		} else {
			ui.MsgBox(mainwin, "操作提示", msg)
		}
	})

	input.Append("", submit, false)

	return vbox
}

//关于
func makeAboutWindow() ui.Control {
	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)

	hbox.Append(vbox, false)

	data := fmt.Sprintf("%s\n", APP_NAME)
	data += fmt.Sprintf("版本:%s\n", APP_VERSION)
	data += fmt.Sprintf("作者:%s\n", APP_AUTHOR)

	return hbox
}

func InitUI() {
	mainwin = ui.NewWindow(APP_NAME, 720, 480, true)
	//关闭
	mainwin.OnClosing(func(*ui.Window) bool {
		//注册关闭事件
		uiEvent.AppClose()
		ui.Quit()
		return true
	})
	//退出
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return false
	})

	form := ui.NewForm()
	form.SetPadded(true)

	mainwin.SetChild(form)
	mainwin.SetMargined(true)

	passwordInput := ui.NewPasswordEntry()

	loginButton := ui.NewButton("登录")
	loginButton.OnClicked(func(button *ui.Button) {
		//登录事件
		if uiEvent.LoginCheck(passwordInput.Text()) == false {
			ui.MsgBox(mainwin, "消息", "密码错误")
			return
		}

		//主菜单
		tab := ui.NewTab()
		mainwin.SetChild(tab)
		mainwin.SetMargined(true)

		tab.Append("主页", makeHomeWindow())
		//tab.Append("设置", makeSettingWindow())
		tab.Append("密码", makePasswordWindow())
		tab.Append("关于", makeAboutWindow())
	})

	form.Append("密码", passwordInput, false)
	form.Append("", loginButton, false)

	//事件初始化
	uiEvent.AppInit()

	mainwin.Show()
}