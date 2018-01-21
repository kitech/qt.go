package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QStyleOption struct {
	*qtrt.CObject
}

func (this *QStyleOption) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQStyleOptionFromPointer(cthis unsafe.Pointer) *QStyleOption {
	return &QStyleOption{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:102
// index:0
// Public
// void QStyleOption(int, int)
func NewQStyleOption(version int, type_ int) *QStyleOption {
	cthis := qtrt.Calloc(1, 256) // 64
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOptionC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &version, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:104
// index:0
// Public
// void ~QStyleOption()
func DeleteQStyleOption(*QStyleOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOptionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:106
// index:0
// Public
// void init(const class QWidget *)
func (this *QStyleOption) Init(w *QWidget /*444 const QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOption4initEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:107
// index:0
// Public inline
// void initFrom(const class QWidget *)
func (this *QStyleOption) InitFrom(w *QWidget /*444 const QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOption8initFromEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
