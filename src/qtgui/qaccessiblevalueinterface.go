//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QAccessibleValueInterface struct {
	*qtrt.CObject
}

func (this *QAccessibleValueInterface) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qaccessible.h:566
// index:0
// virtual
// void ~QAccessibleValueInterface()
func DeleteQAccessibleValueInterface(*QAccessibleValueInterface) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleValueInterfaceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:568
// index:0
// pure virtual
// QVariant currentValue()
func (this *QAccessibleValueInterface) CurrentValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12currentValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:569
// index:0
// pure virtual
// void setCurrentValue(const class QVariant &)
func (this *QAccessibleValueInterface) SetCurrentValue(value unsafe.Pointer) {
	// 0: (, value const QVariant &), (value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QAccessibleValueInterface15setCurrentValueERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:570
// index:0
// pure virtual
// QVariant maximumValue()
func (this *QAccessibleValueInterface) MaximumValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12maximumValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:571
// index:0
// pure virtual
// QVariant minimumValue()
func (this *QAccessibleValueInterface) MinimumValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface12minimumValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:572
// index:0
// pure virtual
// QVariant minimumStepSize()
func (this *QAccessibleValueInterface) MinimumStepSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QAccessibleValueInterface15minimumStepSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
