package heap


//代表一个java类的实例
type Object struct {
	class *Class //存放对象的 Class 指针
	fields Slots //存放实例变量
}

func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}
