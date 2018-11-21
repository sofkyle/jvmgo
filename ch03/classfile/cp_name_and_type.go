package classfile

/*
 * 字段、方法的名称和描述符
 * name_index、descriptor_index都指向CONSTANT_Utf8_info常量
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex					uint16
	descriptorIndex				uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}