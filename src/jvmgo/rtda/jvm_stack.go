package rtda


//虚拟机栈
type Stack struct {
	maxSize uint
	size uint
	_top *Frame
}


func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
/*


00 ldc #4
02 fstore_1
03 fconst_2
04 fload_1
05 fmul
06 fload_0
07 fmul
08 fstore_2
09 fload_2
10 return

public static float circumference(float r) {
	float pi = 3.14f;
	float area = 2 * pi * r;
	return area;
}

 */

//入栈
func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if stack._top != nil {
		frame.lower = stack._top
	}
	stack._top = frame
	stack.size++
}


//访问栈顶元素
func (stack *Stack) top() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty")
	}
	return stack._top
}

//出栈
func (stack *Stack) pop() *Frame {
	if stack._top == nil {
		panic("jvm stack is empty")
	}

	top := stack._top
	stack._top = top.lower
	top.lower = nil
	stack.size--
	return top
}

func (stack *Stack) isEmpty() bool {
	return stack._top == nil
}

