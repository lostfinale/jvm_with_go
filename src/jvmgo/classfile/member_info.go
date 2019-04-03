package classfile

/**
	field_info {
		u2 access_flags;
		u2 name_index;
		u2 descriptor_index;
		u2 attributes_count;
		attribute_info attributes[attributes_count];
	}
 */



type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16 //访问权限标志
	nameIndex uint16 //名字索引
	descriptorIndex uint16 //描述符索引
	attributes []AttributeInfo //属性表
}

//读取成员里列表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)

	for i := range members {
		members[i] = readMember(reader, cp)
	}


	return members
}

//读取单个成员
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	accessFlags := reader.readUint16()
	nameIndex := reader.readUint16()
	descriptorIndex := reader.readUint16()
	attributes := readAttributes(reader, cp)
	return &MemberInfo{
		cp:              cp,
		accessFlags:     accessFlags,
		nameIndex:       nameIndex,
		descriptorIndex: descriptorIndex,
		attributes:      attributes,
	}
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.accessFlags
}

//从常量池查找字段
func (mi *MemberInfo) Name() string{
	return mi.cp.getUtf8(mi.nameIndex)
}

//从常量池查找方法描述符
func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (mi *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

