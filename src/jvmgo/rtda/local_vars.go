package rtda

import (
	"math"
	"jvmgo/rtda/heap"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (lv LocalVars) SetInt(index uint, val int32) {
	lv[index].num = val
}

func (lv LocalVars) GetInt(index uint) int32 {
	return lv[index].num
}


//float 变量可以先转成 int 类型，然后按 int 变量来处理。
func (lv LocalVars) SetFloat(index uint, val float32) {
	//将float转换成32位bit对应的uint数值
	bits := math.Float32bits(val)
	//将uint32转换为int32进行存储
	lv[index].num = int32(bits)
}

func (lv LocalVars) GetFloat(index uint) float32 {
	bits := uint32(lv[index].num)
	return math.Float32frombits(bits)
}



//long 变量则需要拆成两个 int 变量。
func (lv LocalVars) SetLong(index uint, val int64) {
	lv[index].num = int32(val)
	lv[index+1].num = int32(val >> 32)
}

func (lv LocalVars) GetLong(index uint) int64 {
	low := uint32(lv[index].num)
	high := uint32(lv[index+1].num)
	return int64(high) << 32 | int64(low)
}


//double 变量可以先转成 long 类型，然后按照 long 变量来处理。
func (lv LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	lv.SetLong(index, int64(bits))
}

func (lv LocalVars) GetDouble(index uint) float64 {
	bits := uint64(lv.GetLong(index))
	return math.Float64frombits(bits)
}


func (lv LocalVars) SetRef(index uint, ref *heap.Object) {
	lv[index].ref = ref
}

func (lv LocalVars) GetRef(index uint) *heap.Object {
	return lv[index].ref
}

func (lv LocalVars) SetSlot(index uint, slot Slot) {
	lv[index] = slot
}


func (self LocalVars) GetThis() *heap.Object {
	return self.GetRef(0)
}
