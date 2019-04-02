package rtda

import (
	"math"
	"jvmgo/ch06/heap"
)

//操作数栈

type OperandStack struct {

	size uint //size 字段用于记录栈顶
	slots []Slot
}



func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots : make([]Slot, maxStack),
		}
	}
	return nil
}


func (os *OperandStack) PushSlot(slot Slot) {
	os.slots[os.size] = slot
	os.size++
}

func (os *OperandStack) PopSlot() Slot{
	os.size--
	return os.slots[os.size]
}

func (os *OperandStack) PushInt(val int32) {
	os.slots[os.size].num = val
	os.size++
}

func (os *OperandStack) PopInt() int32 {
	os.size--
	return os.slots[os.size].num
}

func (os *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	os.slots[os.size].num = int32(bits)
	os.size++
}

func (os *OperandStack) PopFloat() float32 {
	os.size--
	bits := uint32(os.slots[os.size].num)
	return math.Float32frombits(bits)
}

func (os *OperandStack) PushLong(val int64) {
	os.slots[os.size].num = int32(val)
	os.slots[os.size+1].num = int32(val >> 32)
	os.size += 2
}

func (os *OperandStack) PopLong() int64 {
	os.size -= 2
	low := uint32(os.slots[os.size].num)
	high := uint32(os.slots[os.size+1].num)
	return int64(high) << 32 | int64(low)
}


func (os *OperandStack)  PushDouble(val float64) {
	bits := int64(math.Float64bits(val))
	os.PushLong(bits)
}

func (os *OperandStack) PopDouble() float64 {
	bits := uint64(os.PopLong())
	return math.Float64frombits(bits)
}

func (os *OperandStack) PushRef(ref *heap.Object) {
	os.slots[os.size].ref = ref
	os.size++
}

func (os *OperandStack) PopRef() *heap.Object {
	os.size--
	ref := os.slots[os.size].ref
	//弹出引用后，把Slot 结构体的 ref 字段设置成 nil
	//这样做是为了帮助 Go 的垃圾收集器回收 Object 结构体实例
	os.slots[os.size].ref = nil
	return ref
}