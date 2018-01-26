package qtcore

// /usr/include/qt/QtCore/qmap.h
// #include <qmap.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
type QMapNodeBase struct {
	*qtrt.CObject
}

func (this *QMapNodeBase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMapNodeBase) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQMapNodeBaseFromPointer(cthis unsafe.Pointer) *QMapNodeBase {
	return &QMapNodeBase{&qtrt.CObject{cthis}}
}
func (*QMapNodeBase) NewFromPointer(cthis unsafe.Pointer) *QMapNodeBase {
	return NewQMapNodeBaseFromPointer(cthis)
}

// /usr/include/qt/QtCore/qmap.h:91
// index:0
// Public
// const QMapNodeBase * nextNode()
func (this *QMapNodeBase) NextNode() *QMapNodeBase /*777 const QMapNodeBase **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QMapNodeBase8nextNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmap.h:92
// index:1
// Public inline
// QMapNodeBase * nextNode()
func (this *QMapNodeBase) NextNode_1() *QMapNodeBase /*777 QMapNodeBase **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMapNodeBase8nextNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmap.h:93
// index:0
// Public
// const QMapNodeBase * previousNode()
func (this *QMapNodeBase) PreviousNode() *QMapNodeBase /*777 const QMapNodeBase **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QMapNodeBase12previousNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmap.h:94
// index:1
// Public inline
// QMapNodeBase * previousNode()
func (this *QMapNodeBase) PreviousNode_1() *QMapNodeBase /*777 QMapNodeBase **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMapNodeBase12previousNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qmap.h:96
// index:0
// Public inline
// QMapNodeBase::Color color()
func (this *QMapNodeBase) Color() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QMapNodeBase5colorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qmap.h:97
// index:0
// Public inline
// void setColor(enum QMapNodeBase::Color)
func (this *QMapNodeBase) SetColor(c int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QMapNodeBase8setColorENS_5ColorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmap.h:98
// index:0
// Public inline
// QMapNodeBase * parent()
func (this *QMapNodeBase) Parent() *QMapNodeBase /*777 QMapNodeBase **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QMapNodeBase6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMapNodeBaseFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QMapNodeBase__Color = int

const QMapNodeBase__Red QMapNodeBase__Color = 0
const QMapNodeBase__Black QMapNodeBase__Color = 1

type QMapNodeBase__ = int

const QMapNodeBase__Mask QMapNodeBase__ = 3

//  body block end
