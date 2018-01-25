package qtcore

// /usr/include/qt/QtCore/qflags.h
// #include <qflags.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QFlag struct {
	*qtrt.CObject
}

func (this *QFlag) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFlag) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQFlagFromPointer(cthis unsafe.Pointer) *QFlag {
	return &QFlag{&qtrt.CObject{cthis}}
}
func (*QFlag) NewFromPointer(cthis unsafe.Pointer) *QFlag {
	return NewQFlagFromPointer(cthis)
}

// /usr/include/qt/QtCore/qflags.h:57
// index:0
// Public inline
// void QFlag(int)
func NewQFlag(ai int) *QFlag {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFlagC2Ei", ffiqt.FFI_TYPE_VOID, cthis, ai)
	gopp.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qflags.h:68
// index:1
// Public inline
// void QFlag(uint)
func NewQFlag_1(ai uint) *QFlag {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFlagC2Ej", ffiqt.FFI_TYPE_VOID, cthis, ai)
	gopp.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qflags.h:69
// index:2
// Public inline
// void QFlag(short)
func NewQFlag_2(ai int16) *QFlag {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFlagC2Es", ffiqt.FFI_TYPE_VOID, cthis, ai)
	gopp.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qflags.h:70
// index:3
// Public inline
// void QFlag(ushort)
func NewQFlag_3(ai uint16) *QFlag {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFlagC2Et", ffiqt.FFI_TYPE_VOID, cthis, ai)
	gopp.ErrPrint(err, rv)
	gothis := NewQFlagFromPointer(cthis)
	return gothis
}

//  body block end
