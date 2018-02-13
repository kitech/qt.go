package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QStyleOptionViewItem struct {
	*QStyleOption
}
type QStyleOptionViewItem_ITF interface {
	QStyleOption_ITF
	QStyleOptionViewItem_PTR() *QStyleOptionViewItem
}

func (ptr *QStyleOptionViewItem) QStyleOptionViewItem_PTR() *QStyleOptionViewItem { return ptr }

func (this *QStyleOptionViewItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionViewItem) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionViewItemFromPointer(cthis unsafe.Pointer) *QStyleOptionViewItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionViewItem{bcthis0}
}
func (*QStyleOptionViewItem) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionViewItem {
	return NewQStyleOptionViewItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:442
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionViewItem()
func NewQStyleOptionViewItem() *QStyleOptionViewItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionViewItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:446
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionViewItem(int)
func NewQStyleOptionViewItem_1(version int) *QStyleOptionViewItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionViewItem)
	return gothis
}

func DeleteQStyleOptionViewItem(this *QStyleOptionViewItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionViewItem__StyleOptionType = int

const QStyleOptionViewItem__Type QStyleOptionViewItem__StyleOptionType = 10

type QStyleOptionViewItem__StyleOptionVersion = int

const QStyleOptionViewItem__Version QStyleOptionViewItem__StyleOptionVersion = 4

type QStyleOptionViewItem__Position = int

const QStyleOptionViewItem__Left QStyleOptionViewItem__Position = 0
const QStyleOptionViewItem__Right QStyleOptionViewItem__Position = 1
const QStyleOptionViewItem__Top QStyleOptionViewItem__Position = 2
const QStyleOptionViewItem__Bottom QStyleOptionViewItem__Position = 3

type QStyleOptionViewItem__ViewItemFeature = int

const QStyleOptionViewItem__None QStyleOptionViewItem__ViewItemFeature = 0
const QStyleOptionViewItem__WrapText QStyleOptionViewItem__ViewItemFeature = 1
const QStyleOptionViewItem__Alternate QStyleOptionViewItem__ViewItemFeature = 2
const QStyleOptionViewItem__HasCheckIndicator QStyleOptionViewItem__ViewItemFeature = 4
const QStyleOptionViewItem__HasDisplay QStyleOptionViewItem__ViewItemFeature = 8
const QStyleOptionViewItem__HasDecoration QStyleOptionViewItem__ViewItemFeature = 16

type QStyleOptionViewItem__ViewItemPosition = int

const QStyleOptionViewItem__Invalid QStyleOptionViewItem__ViewItemPosition = 0
const QStyleOptionViewItem__Beginning QStyleOptionViewItem__ViewItemPosition = 1
const QStyleOptionViewItem__Middle QStyleOptionViewItem__ViewItemPosition = 2
const QStyleOptionViewItem__End QStyleOptionViewItem__ViewItemPosition = 3
const QStyleOptionViewItem__OnlyOne QStyleOptionViewItem__ViewItemPosition = 4

//  body block end

//  keep block begin

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
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
