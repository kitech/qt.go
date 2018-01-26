package qtcore

// /usr/include/qt/QtCore/qobjectdefs.h
// #include <qobjectdefs.h>
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
import "mkuse/cffiqt"
import "gopp"
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
type QGenericReturnArgument struct {
	*QGenericArgument
}

func (this *QGenericReturnArgument) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGenericArgument.GetCthis()
	}
}
func (this *QGenericReturnArgument) SetCthis(cthis unsafe.Pointer) {
	this.QGenericArgument = NewQGenericArgumentFromPointer(cthis)
}
func NewQGenericReturnArgumentFromPointer(cthis unsafe.Pointer) *QGenericReturnArgument {
	bcthis0 := NewQGenericArgumentFromPointer(cthis)
	return &QGenericReturnArgument{bcthis0}
}
func (*QGenericReturnArgument) NewFromPointer(cthis unsafe.Pointer) *QGenericReturnArgument {
	return NewQGenericReturnArgumentFromPointer(cthis)
}

// /usr/include/qt/QtCore/qobjectdefs.h:310
// index:0
// Public inline
// void QGenericReturnArgument(const char *, void *)
func NewQGenericReturnArgument(aName string, aData unsafe.Pointer /*666*/) *QGenericReturnArgument {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = qtrt.CString(aName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QGenericReturnArgumentC2EPKcPv", ffiqt.FFI_TYPE_VOID, cthis, convArg0, aData)
	gopp.ErrPrint(err, rv)
	gothis := NewQGenericReturnArgumentFromPointer(cthis)
	return gothis
}

//  body block end
