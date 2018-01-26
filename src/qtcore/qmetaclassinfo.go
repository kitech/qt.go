package qtcore

// /usr/include/qt/QtCore/qmetaobject.h
// #include <qmetaobject.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QMetaClassInfo struct {
	*qtrt.CObject
}

func (this *QMetaClassInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMetaClassInfo) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMetaClassInfoFromPointer(cthis unsafe.Pointer) *QMetaClassInfo {
	return &QMetaClassInfo{&qtrt.CObject{cthis}}
}
func (*QMetaClassInfo) NewFromPointer(cthis unsafe.Pointer) *QMetaClassInfo {
	return NewQMetaClassInfoFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmetaobject.h:303
// index:0
// Public inline
// void QMetaClassInfo()
func NewQMetaClassInfo() *QMetaClassInfo {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QMetaClassInfoC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQMetaClassInfoFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qmetaobject.h:304
// index:0
// Public
// const char * name()
func (this *QMetaClassInfo) Name() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMetaClassInfo4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:305
// index:0
// Public
// const char * value()
func (this *QMetaClassInfo) Value() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMetaClassInfo5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qmetaobject.h:306
// index:0
// Public inline
// const QMetaObject * enclosingMetaObject()
func (this *QMetaClassInfo) EnclosingMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QMetaClassInfo19enclosingMetaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

//  body block end
