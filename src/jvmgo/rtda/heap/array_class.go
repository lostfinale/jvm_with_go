package heap


/*

public class ArrayDemo {
	public static void main(String[] args) {
		int[] a1 = new int[10]; // newarray
		String[] a2 = new String[10]; // anewarray
		int[][] a3 = new int[10][10]; // multianewarray
		int x = a1.length; // arraylength
		a1[0] = 100; // iastore
		int y = a1[0]; // iaload
		a2[0] = "abc"; // aastore
		String s = a2[0]; // aaload
	}
}



 */
func (self *Class) NewArray(count uint) *Object {
	if !self.IsArray() {
		panic("Not array class: " + self.name)
	}

	switch self.Name() {
	case "[Z":
		return &Object{self, make([]int8, count)}
	case "[B":
		return &Object{self, make([]int8, count)}
	case "[C":
		return &Object{self, make([]uint16, count)}
	case "[S":
		return &Object{self, make([]int16, count)}
	case "[I":
		return &Object{self, make([]int32, count)}
	case "[J":
		return &Object{self, make([]int64, count)}
	case "[F":
		return &Object{self, make([]float32, count)}
	case "[D":
		return &Object{self, make([]float64, count)}
	default:
		return &Object{self, make([]*Object, count)}
	}
}

func (self *Class) IsArray() bool {
	return self.name[0] == '['
}
