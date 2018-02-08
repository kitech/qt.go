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

type QStyleOptionGraphicsItem struct {
	*QStyleOption
}

func (this *QStyleOptionGraphicsItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QStyleOption.GetCthis()
	}
}
func (this *QStyleOptionGraphicsItem) SetCthis(cthis unsafe.Pointer) {
	this.QStyleOption = NewQStyleOptionFromPointer(cthis)
}
func NewQStyleOptionGraphicsItemFromPointer(cthis unsafe.Pointer) *QStyleOptionGraphicsItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionGraphicsItem{bcthis0}
}
func (*QStyleOptionGraphicsItem) NewFromPointer(cthis unsafe.Pointer) *QStyleOptionGraphicsItem {
	return NewQStyleOptionGraphicsItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstyleoption.h:669
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionGraphicsItem()
func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGraphicsItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:673
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionGraphicsItem(int)
func NewQStyleOptionGraphicsItem_1(version int) *QStyleOptionGraphicsItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemC2Ei", qtrt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionGraphicsItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:671
// index:0
// Public static Visibility=Default Availability=Available
// [8] qreal levelOfDetailFromTransform(const QTransform &)
func (this *QStyleOptionGraphicsItem) LevelOfDetailFromTransform(worldTransform *qtgui.QTransform) float64 {
	var convArg0 = worldTransform.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}
func QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform *qtgui.QTransform) float64 {
	var nilthis *QStyleOptionGraphicsItem
	rv := nilthis.LevelOfDetailFromTransform(worldTransform)
	return rv
}

func DeleteQStyleOptionGraphicsItem(this *QStyleOptionGraphicsItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QStyleOptionGraphicsItem__StyleOptionType = int

const QStyleOptionGraphicsItem__Type QStyleOptionGraphicsItem__StyleOptionType = 15

type QStyleOptionGraphicsItem__StyleOptionVersion = int

const QStyleOptionGraphicsItem__Version QStyleOptionGraphicsItem__StyleOptionVersion = 1

//  body block end
