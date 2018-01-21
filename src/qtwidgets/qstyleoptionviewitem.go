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
type QStyleOptionViewItem struct {
	*QStyleOption
}

func (this *QStyleOptionViewItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func NewQStyleOptionViewItemFromPointer(cthis unsafe.Pointer) *QStyleOptionViewItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionViewItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:442
// index:0
// Public
// void QStyleOptionViewItem()
func NewQStyleOptionViewItem() *QStyleOptionViewItem {
	cthis := qtrt.Calloc(1, 256) // 192
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:446
// index:1
// Protected
// void QStyleOptionViewItem(int)
func NewQStyleOptionViewItem_1(version int) *QStyleOptionViewItem {
	cthis := qtrt.Calloc(1, 256) // 192
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(cthis)
	return gothis
}

//  body block end
