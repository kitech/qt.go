//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstyleoption.h:102
// index:0
// void QStyleOption(int, int)
func NewQStyleOption(version int, type_ int) *QStyleOption {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOptionC2Eii", ffiqt.FFI_TYPE_VOID, cthis, &version, &type_)
	gopp.ErrPrint(err, rv)
	return &QStyleOption{cthis}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:104
// index:0
// void ~QStyleOption()
func DeleteQStyleOption(*QStyleOption) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOptionD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:106
// index:0
// void init(const class QWidget *)
func (this *QStyleOption) Init(w unsafe.Pointer) {
	// 0: (, const QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOption4initEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:107
// index:0
// inline
// void initFrom(const class QWidget *)
func (this *QStyleOption) InitFrom(w unsafe.Pointer) {
	// 0: (, const QWidget * w), (w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QStyleOption8initFromEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, w)
	gopp.ErrPrint(err, rv)
}

//  body block end
