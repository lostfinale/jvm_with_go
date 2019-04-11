package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"fmt"
)

//StackTraceElement 结构体用来记录 Java 虚拟机栈帧信息： lineNumber 字段给出帧正在执行哪行代码；
//methodName 字段给出方法名； className 字段给出声明方法的类名； fileName 字段给出类所在的文件名。
type StackTraceElement struct {
	fileName string
	className string
	methodName string
	lineNumber int
}


func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}


func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

// private native Throwable fillInStackTrace(int dummy);
func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement{
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func distanceToObject(class *heap.Class) int{
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement {
		fileName : class.SourceFile(),
		className: class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}