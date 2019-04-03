package base

import "jvmgo/rtda"

//跳到指定offset的语句位置
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
