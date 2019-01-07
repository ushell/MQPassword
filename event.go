package MQPassword


type Event struct {

}

func (this *Event) LoginCheck(password string) bool {
	model := new(Model)
	return model.PasswordCheck(password)
}

func (this *Event) AppClose() bool {
	model := new(Model)
	return model.Store()
}

func (this *Event) PasswordChange(password string) (result bool, msg string) {
	return true, ""
}

func (this *Event) AppInit() {
	//数据初始化
	model := new(Model)
	model.EnvInit()
}

func (this *Event) PasswordSave(data InputContent) bool {
	model := new(Model)
	return model.Save(data)
}

func (this *Event) PasswordItem(data string) InputContent {
	model := new(Model)
	return model.Item(data)
}
