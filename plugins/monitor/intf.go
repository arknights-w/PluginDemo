package monitor

type MT interface{
	Command(cmd string,path string) CallBack
	Commands(cmd string,paths []string) CallBack
}

type CallBack struct{
	Reslut bool
	Msg string 
}