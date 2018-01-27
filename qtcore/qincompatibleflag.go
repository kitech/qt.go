package qtcore

// /usr/include/qt/QtCore/qflags.h
// #include <qflags.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
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
type QIncompatibleFlag struct {
	*qtrt.CObject
}

func (this *QIncompatibleFlag) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIncompatibleFlag) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQIncompatibleFlagFromPointer(cthis unsafe.Pointer) *QIncompatibleFlag {
	return &QIncompatibleFlag{&qtrt.CObject{cthis}}
}
func (*QIncompatibleFlag) NewFromPointer(cthis unsafe.Pointer) *QIncompatibleFlag {
	return NewQIncompatibleFlagFromPointer(cthis)
}

// /usr/include/qt/QtCore/qflags.h:80
// index:0
// Public inline
// void QIncompatibleFlag(int)
func NewQIncompatibleFlag(i int) *QIncompatibleFlag {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QIncompatibleFlagC2Ei", ffiqt.FFI_TYPE_VOID, cthis, i)
	gopp.ErrPrint(err, rv)
	gothis := NewQIncompatibleFlagFromPointer(cthis)
	return gothis
}

//  body block end
