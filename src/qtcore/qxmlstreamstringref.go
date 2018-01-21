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
func NewQXmlStreamStringRefFromPointer(cthis unsafe.Pointer) *QXmlStreamStringRef {
	return &QXmlStreamStringRef{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qxmlstream.h:58
// index:0
// Public inline
// void QXmlStreamStringRef()
func NewQXmlStreamStringRef() *QXmlStreamStringRef {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:59
// index:1
// Public inline
// void QXmlStreamStringRef(const class QStringRef &)
func NewQXmlStreamStringRef_1(aString *QStringRef) *QXmlStreamStringRef {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = aString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2ERK10QStringRef", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:61
// index:2
// Public inline
// void QXmlStreamStringRef(const class QString &)
func NewQXmlStreamStringRef_2(aString *QString) *QXmlStreamStringRef {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = aString.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:63
// index:3
// Public inline
// void QXmlStreamStringRef(class QString &&)
func NewQXmlStreamStringRef_3(aString unsafe.Pointer /*333*/) *QXmlStreamStringRef {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefC2EO7QString", ffiqt.FFI_TYPE_VOID, cthis, aString)
	gopp.ErrPrint(err, rv)
	gothis := NewQXmlStreamStringRefFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:77
// index:0
// Public inline
// void ~QXmlStreamStringRef()
func DeleteQXmlStreamStringRef(*QXmlStreamStringRef) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRefD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:82
// index:0
// Public inline
// void swap(class QXmlStreamStringRef &)
func (this *QXmlStreamStringRef) Swap(other *QXmlStreamStringRef) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRef4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:89
// index:0
// Public inline
// void clear()
func (this *QXmlStreamStringRef) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QXmlStreamStringRef5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qxmlstream.h:91
// index:0
// Public inline
// const QString * string()
func (this *QXmlStreamStringRef) String() *QString /*444 const QString **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef6stringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:92
// index:0
// Public inline
// int position()
func (this *QXmlStreamStringRef) Position() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef8positionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtCore/qxmlstream.h:93
// index:0
// Public inline
// int size()
func (this *QXmlStreamStringRef) Size() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QXmlStreamStringRef4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
