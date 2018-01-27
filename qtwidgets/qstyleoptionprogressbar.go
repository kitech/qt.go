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
type QStyleOptionProgressBar struct {
	*QStyleOption
}

func (this *QStyleOptionProgressBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionProgressBar) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionProgressBarFromPointer(cthis unsafe.Pointer) *QStyleOptionProgressBar {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionProgressBar{bcthis0}
}
func (*QStyleOptionProgressBar) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionProgressBar {
	return NewQStyleOptionProgressBarFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:342
// index:0
// Public
// void QStyleOptionProgressBar()
func NewQStyleOptionProgressBar() *QStyleOptionProgressBar {
	cthis := qtrt.Calloc(1, 256) // 104
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QStyleOptionProgressBarC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionProgressBarFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:346
// index:1
// Protected
// void QStyleOptionProgressBar(int)
func NewQStyleOptionProgressBar_1(version int) *QStyleOptionProgressBar {
	cthis := qtrt.Calloc(1, 256) // 104
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QStyleOptionProgressBarC2Ei", ffiqt.FFI_TYPE_VOID, cthis, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionProgressBarFromPointer(cthis)
	return gothis
}

type QStyleOptionProgressBar__StyleOptionType = int

const QStyleOptionProgressBar__Type QStyleOptionProgressBar__StyleOptionType = 6

type QStyleOptionProgressBar__StyleOptionVersion = int

const QStyleOptionProgressBar__Version QStyleOptionProgressBar__StyleOptionVersion = 2

//  body block end
