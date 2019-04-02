package classfile

import "fmt"

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	//常量池大小
	u := reader.readUint16()
	cpCount := int(u)
	cp := make([]ConstantInfo, cpCount)

	//索引从1开始，0号位置是废弃不用的
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		//fmt.Printf("读取常量index:[%d]->[%T]:[%v]\n", i, cp[i], cp[i] )
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			//Long和Double占两个常量索引
			i++
		}
	}


	return cp
}

//按索引查找常量
func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	info := fmt.Sprintf("Invalid constant pool index:%d", index)
	panic(info)
}

//从常量池查找字段或方法的名字和描述符
func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	nameType := cp.getUtf8(ntInfo.descriptorIndex)
	return name, nameType
}

//从常量池查找类名
func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

//常量池查找 UTF-8 字符串
func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
