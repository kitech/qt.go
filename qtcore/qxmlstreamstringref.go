package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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

type QXmlStreamStringRef struct {
	*qtrt.CObject
}

func (this *QXmlStreamStringRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamStringRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamStringRefFromPointer(cthis unsafe.Pointer) *QXmlStreamStringRef {
	return &QXmlStreamStringRef{&qtrt.CObject{cthis}}
}
func (*QXmlStreamStringRef) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamStringRef {
	return NewQXmlStreamStringRefFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef()
func NewQXmlStreamStringRef() *QXmlStreamStringRef {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:59
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef(const QStringRef &)
func NewQXmlStreamStringRef_1(aString *QStringRef) *QXmlStreamStringRef {
	var convArg0 = aString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2ERK10QStringRef", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:61
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef(const QString &)
func NewQXmlStreamStringRef_2(aString *QString) *QXmlStreamStringRef {
	var convArg0 = aString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:63
// index:3
// Public inline Visibility=Default Availability=Available
// [-2] void QXmlStreamStringRef(QString &&)
func NewQXmlStreamStringRef_3(aString unsafe.Pointer /*333*/) *QXmlStreamStringRef {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2EO7QString", ffiqt.FFI_TYPE_POINTER, aString)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamStringRef)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QXmlStreamStringRef()
func DeleteQXmlStreamStringRef(this *QXmlStreamStringRef) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:82
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QXmlStreamStringRef &)
func (this *QXmlStreamStringRef) Swap(other *QXmlStreamStringRef) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRef4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:89
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clear()
func (this *QXmlStreamStringRef) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRef5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QString * string()
func (this *QXmlStreamStringRef) String() *QString /*777 const QString **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef6stringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int position()
func (this *QXmlStreamStringRef) Position() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qxmlstream.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int size()
func (this *QXmlStreamStringRef) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
