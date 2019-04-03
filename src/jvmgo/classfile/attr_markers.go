package classfile



/*

	Deprecated 属性用于指出类、接口、字段或方法已经不建议使用，
	编译器等工具可以根据 Deprecated 属性输出警告信息
	Deprecated_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
	}

	Synthetic 属性用来标记源文件中不存在、由编译器生成的类成员，
	引入 Synthetic 属性主要是为了支持嵌套类和嵌套接口

	Synthetic_attribute {
		u2 attribute_name_index;
		u4 attribute_length;
	}
*/

type MarkerAttribute struct {}


type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (ma *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}