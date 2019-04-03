package classfile

/*
	SourceFile 是可选定长属性，只会出现在 ClassFile 结构中，用于指出源文件名
	SourceFile_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 sourcefile_index;
	}
*/

type SourceFileAttribute struct {
	cp ConstantPool
	sourceFileIndex uint16
}

func (sa *SourceFileAttribute) readInfo(reader *ClassReader) {
	sa.sourceFileIndex = reader.readUint16()
}

func (sa *SourceFileAttribute) FileName() string {
	return sa.cp.getUtf8(sa.sourceFileIndex)
}