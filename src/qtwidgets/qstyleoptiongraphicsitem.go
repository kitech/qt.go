//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

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
type QStyleOptionGraphicsItem struct {
	*QStyleOption
}

func (this *QStyleOptionGraphicsItem) GetCthis() unsafe.Pointer {
	return this.QStyleOption.GetCthis()
}

// /usr/include/qt/QtWidgets/qstyleoption.h:669
// index:0
// void QStyleOptionGraphicsItem()
func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGraphicsItemFromPointer(cthis)
	return gothis
}
func NewQStyleOptionGraphicsItemFromPointer(cthis unsafe.Pointer) *QStyleOptionGraphicsItem {
	bcthis0 := NewQStyleOptionFromPointer(cthis)
	return &QStyleOptionGraphicsItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:673
// index:1
// void QStyleOptionGraphicsItem(int)
func NewQStyleOptionGraphicsItem_1(version int) *QStyleOptionGraphicsItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &version)
	gopp.ErrPrint(err, rv)
	gothis := NewQStyleOptionGraphicsItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:671
// index:0
// static
// qreal levelOfDetailFromTransform(const class QTransform &)
func (this *QStyleOptionGraphicsItem) LevelOfDetailFromTransform(worldTransform unsafe.Pointer) {
	// 0: (worldTransform const QTransform &), (worldTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform unsafe.Pointer) {
	// 0: (worldTransform const QTransform &), (worldTransform)
	var nilthis *QStyleOptionGraphicsItem
	nilthis.LevelOfDetailFromTransform(worldTransform)
}

//  body block end
