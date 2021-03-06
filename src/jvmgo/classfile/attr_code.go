package classfile


/*
	Code 是变长属性，只存在于 method_info 结构中。
	Code 属性中存放字节码等方法相关信息。
	相比前面介绍的几种属性， Code 属性比较复杂，其结构定义如下：

	Code_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 max_stack;
		u2 max_locals;
		u4 code_length;
		u1 code[code_length];
		u2 exception_table_length;
		{ 	u2 start_pc;
			u2 end_pc;
			u2 handler_pc;
			u2 catch_type;
		} exception_table[exception_table_length];
		u2 attributes_count;
		attribute_info attributes[attributes_count];
	}

*/


type CodeAttribute struct {
	cp ConstantPool
	maxStack uint16
	maxLocals uint16
	code []byte
	exceptionTable []*ExceptionTableEntry
	attributes []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}

func (ca * CodeAttribute)readInfo(reader *ClassReader) {
	ca.maxStack = reader.readUint16()
	ca.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	ca.code = reader.readBytes(codeLength)
	ca.exceptionTable = readExceptionTable(reader)
	ca.attributes = readAttributes(reader, ca.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:reader.readUint16(),
			endPc:reader.readUint16(),
			handlerPc:reader.readUint16(),
			catchType:reader.readUint16(),
		}
	}
	return exceptionTable
}


func (self *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

func (ca *CodeAttribute) MaxStack() uint {
	return uint(ca.maxStack)
}
func (ca *CodeAttribute) MaxLocals() uint {
	return uint(ca.maxLocals)
}
func (ca *CodeAttribute) Code() []byte {
	return ca.code
}
func (ca *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return ca.exceptionTable
}



func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
