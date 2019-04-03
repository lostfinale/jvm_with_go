package classfile

/*
	Exceptions_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
		u2 number_of_exceptions;
		u2 exception_index_table[number_of_exceptions];
	}

*/

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (ea * ExceptionsAttribute) readInfo(reader *ClassReader) {
	//readUint16s已经读取了长度，所以这里不用读取number_of_exceptions
	ea.exceptionIndexTable = reader.readUint16s()
}

func (ea *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return ea.exceptionIndexTable
}