package tool

var GV *GlobalVariable

type GlobalVariable struct {
	GetMd5State  bool
	DiffMd5State bool
}

func InitGlobalVariable() {
	gv := new(GlobalVariable)
	gv.GetMd5State = false
	gv.DiffMd5State = false
	GV = gv
}
