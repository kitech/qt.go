package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
type QStyleOptionDockWidget struct {
	*QStyleOption
}

func (this *QStyleOptionDockWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionDockWidget) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionDockWidgetFromPointer(cthis unsafe.Pointer) *QStyleOptionDockWidget {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionDockWidget{bcthis0}
}
func (*QStyleOptionDockWidget) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionDockWidget {
	return NewQStyleOptionDockWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:391
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionDockWidget()
func NewQStyleOptionDockWidget() *QStyleOptionDockWidget {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QStyleOptionDockWidgetC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:395
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionDockWidget(int)
func NewQStyleOptionDockWidget_1(version int) *QStyleOptionDockWidget {
	rv, err := ffiqt.InvokeQtFunc6("_ZN22QStyleOptionDockWidgetC2Ei", ffiqt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionDockWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

type QStyleOptionDockWidget__StyleOptionType = int

const QStyleOptionDockWidget__Type QStyleOptionDockWidget__StyleOptionType = 9

type QStyleOptionDockWidget__StyleOptionVersion = int

const QStyleOptionDockWidget__Version QStyleOptionDockWidget__StyleOptionVersion = 2

//  body block end
