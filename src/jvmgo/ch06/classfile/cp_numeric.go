package classfile



//-------------int----------------
/**
	CONSTANT_Integer_info {
		u1 tag;
		u4 bytes;
	}
 */
type ConstantIntegerInfo struct {
	val int32
}

func (ci *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	ci.val = int32(bytes)
}


//--------------float--------------
/*
	CONSTANT_Float_info {
		u1 tag;
		u4 bytes;
	}
 */
type ConstantFloatInfo struct {
	val float32
}

func (ci * ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	ci.val = float32(bytes)
}

//--------------long-----------------
/*
	CONSTANT_Long_info {
		u1 tag;
		u4 high_bytes;
		u4 low_bytes;
	}
 */
type ConstantLongInfo struct {
	val int64
}

func (ci * ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	ci.val = int64(bytes)
}

//--------------double-----------------
/*
	CONSTANT_Double_info {
		u1 tag;
		u4 high_bytes;
		u4 low_bytes;
	}
 */
type ConstantDoubleInfo struct {
	val float64
}

func (ci * ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	ci.val = float64(bytes)
}


