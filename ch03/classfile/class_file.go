package classfile

import "fmt"

type ClassFile struct {
	minorVersion		uint16
	majorVersion 		uint16
	constantPool		ConstantPool
	accessFlage			uint16
	thisClass			uint16
	superClass			uint16
	interfaces			uint16
	fields				[] uint16
	methods				[] *MemberInfo
	attributes			[] AttributeInfo
}

func Parse(classData [] byte) (cf *ClassFile, err error) {

}

func (self *ClassFile) read(reader *ClassReader) {

}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {

}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {

}

func (self *ClassFile) MinorVersion() uint16 {

}

func (self *ClassFile) MajorVersion() uint16 {

}

func (self *ClassFile) ConstantPool() ConstantPool {

}

func (self *ClassFile) AccessFlags() uint16 {

}

func (self *ClassFile) Fields() [] *MemberInfo {

}

func (self *ClassFile) Methods() [] *Memberinfo {

}

func (self *ClassFile) ClassName() string {

}

func (self *ClassFile) SuperClassName() string {

}

func (self *ClassFile) InterfaceNames() [] string {

}