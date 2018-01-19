//  header block begin
// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qstyleoption.h:669
// index:0
// void QStyleOptionGraphicsItem()
func NewQStyleOptionGraphicsItem() *QStyleOptionGraphicsItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItemC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QStyleOptionGraphicsItem{cthis}
}

// /usr/include/qt/QtWidgets/qstyleoption.h:671
// index:0
// static
// qreal levelOfDetailFromTransform(const class QTransform &)
func (this *QStyleOptionGraphicsItem) LevelOfDetailFromTransform(worldTransform unsafe.Pointer) {
	// 0: (const QTransform & worldTransform), (worldTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QStyleOptionGraphicsItem_LevelOfDetailFromTransform(worldTransform unsafe.Pointer) {
	// 0: (const QTransform & worldTransform), (worldTransform)
	var nilthis *QStyleOptionGraphicsItem
	nilthis.LevelOfDetailFromTransform(worldTransform)
}

//  body block end
