package qtmeta

import (
	"log"
	"reflect"
)

// eg. obj.SomeFunc1
func (this *QtMetaData) AddMethodAsSlot(objmth interface{}) {

}

// eg. obj.SomeField1. must a func type
func (this *QtMetaData) AddVarAsSignal(name string, objfld interface{}) {
	fldval := reflect.ValueOf(objfld)
	log.Println(fldval, fldval.Kind())
	this.AddFuncTypeAsSignal(name, fldval.Type())
}
func (this *QtMetaData) AddFieldAsSignal(fld reflect.StructField) {
	this.AddVarAsSignal(fld.Name, fld.Type)
}
func (this *QtMetaData) AddFuncTypeAsSignal(name string, fldty reflect.Type) {
	rettys, argtys := functyproto(fldty)
	this.AddSignal(name, rettys, argtys)
}

// eg. obj.SomeField1. simple type: primity and string
func (this *QtMetaData) AddFieldAsProperty(objfld interface{}) {

}

func functyproto(fty reflect.Type) (rettys, argtys []reflect.Type) {
	for i := 0; i < fty.NumOut(); i++ {
		rettys = append(rettys, fty.Out(i))
	}
	for i := 0; i < fty.NumIn(); i++ {
		argtys = append(argtys, fty.In(i))
	}
	return
}
