package classfile

import "fmt"


/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type ClassFile struct {
	//magic uint32 //魔数用于检验该文件是否是java class 文件
	minorVersion uint16 //次要版本
	majorVersion uint16 //主要版本
	constantPool ConstantPool //常量池
	accessFlags  uint16
	thisClass uint16 //当前类的描述
	superClass uint16 //父类的描述
	interfaces []uint16 //接口列表
	fields []*MemberInfo //成员属性
	methods []*MemberInfo //方法列表
	attributes []AttributeInfo //


}


func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func(){
		if r := recover();  r != nil {
			//var ok bool
			_, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr :=&ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)
	cf.readAndCheckVersion(reader)
	cf.constantPool = readConstantPool(reader)
	cf.accessFlags = reader.readUint16()
	cf.thisClass = reader.readUint16()
	cf.superClass = reader.readUint16()
	cf.interfaces = reader.readUint16s()
	cf.fields = readMembers(reader, cf.constantPool)
	cf.methods = readMembers(reader, cf.constantPool)
	cf.attributes = readAttributes(reader, cf.constantPool)
}


//获取类名
func (cf *ClassFile) ClassName() string {
	//
	return cf.constantPool.getClassName(cf.thisClass)
}

//查找父类
func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""//只有java.lang.Object 没有超类
}

//查找接口
func (cf *ClassFile) InterfaceName() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaceNames

}


//读取并检查魔数（验证文件是否正确）
func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	//魔数是一个u4
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

//读取并检查版本
func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	//版本是两个 u2 分别对应 次要版本和主要版本
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49,50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")

}

func (self *ClassFile) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}


//----------------------getter and setter-------------------------
func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlags
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}
func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}
