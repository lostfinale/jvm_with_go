package heap


//符号常量
type SymRef struct {
	cp *ConstantPool//运行时常量池指针
	className string //类的完全限定名
	class *Class //类结构体指针
}


//获取解析过后的class
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}


//解析class
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}