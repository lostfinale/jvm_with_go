package classfile


/*

	CONSTANT_Class_info {
		u1 tag;
		u2 name_index;
	}
 */

 type ConstantClassInfo struct {
 	cp ConstantPool
 	nameIndex uint16
 }

 func (ci *ConstantClassInfo) readInfo(reader *ClassReader) {
 	ci.nameIndex = reader.readUint16()
 }

 func (ci *ConstantClassInfo) Name() string {
 	return ci.cp.getUtf8(ci.nameIndex)
 }