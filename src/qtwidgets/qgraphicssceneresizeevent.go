//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
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
type QGraphicsSceneResizeEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneResizeEvent) GetCthis() unsafe.Pointer {
	return this.QGraphicsSceneEvent.GetCthis()
}
func NewQGraphicsSceneResizeEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneResizeEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneResizeEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:297
// index:0
// Public
// void QGraphicsSceneResizeEvent()
func NewQGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneResizeEventFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:298
// index:0
// Public virtual
// void ~QGraphicsSceneResizeEvent()
func DeleteQGraphicsSceneResizeEvent(*QGraphicsSceneResizeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:300
// index:0
// Public
// QSizeF oldSize()
func (this *QGraphicsSceneResizeEvent) OldSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsSceneResizeEvent7oldSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:301
// index:0
// Public
// void setOldSize(const class QSizeF &)
func (this *QGraphicsSceneResizeEvent) SetOldSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEvent10setOldSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:303
// index:0
// Public
// QSizeF newSize()
func (this *QGraphicsSceneResizeEvent) NewSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsSceneResizeEvent7newSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:304
// index:0
// Public
// void setNewSize(const class QSizeF &)
func (this *QGraphicsSceneResizeEvent) SetNewSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEvent10setNewSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
