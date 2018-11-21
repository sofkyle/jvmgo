package classfile

/*
 * CONSTANT_String_info本身不存放字符串数据，只存放了常量池索引，
 * 这个索引指向一个CONSTANT_UTF8_info常量
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	cp					ConstantPool
	stringIndex			uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}