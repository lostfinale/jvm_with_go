package classfile


/*
	ConstantValue_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 constantvalue_index;
	}

*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (ca * ConstantValueAttribute) readInfo(reader *ClassReader) {
	ca.constantValueIndex = reader.readUint16()
}

func (ca *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return ca.constantValueIndex
}
