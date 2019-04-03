package classfile

/*
	CONSTANT_String_info {
		u1 tag;
		u2 string_index;
	}

 */
type ConstantStringInfo struct {
	cp ConstantPool
	stringIndex uint16
}

//读取常量池索引
func (ci *ConstantStringInfo) readInfo(reader *ClassReader) {
	ci.stringIndex = reader.readUint16()
}

//按索引从常量池中查找字符串
func (ci *ConstantStringInfo) String() string {
	return ci.cp.getUtf8(ci.stringIndex)
}