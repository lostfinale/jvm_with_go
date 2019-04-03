package classfile



/*
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length];
}
*/


type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

//读取所有属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {

	attributesCount := reader.readUint16()

	attributes := make([]AttributeInfo, attributesCount)

	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}

	return attributes

}


//读取单个属性
func readAttribute(reader *ClassReader,  cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {

	switch attrName {
	case "Code":
		return &CodeAttribute{cp:cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}

}



