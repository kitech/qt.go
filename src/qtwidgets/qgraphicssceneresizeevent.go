//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:297
// index:0
// void QGraphicsSceneResizeEvent()
func NewQGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneResizeEventFromPointer(cthis)
	return gothis
}
func NewQGraphicsSceneResizeEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneResizeEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneResizeEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:298
// index:0
// virtual
// void ~QGraphicsSceneResizeEvent()
func DeleteQGraphicsSceneResizeEvent(*QGraphicsSceneResizeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:300
// index:0
// QSizeF oldSize()
func (this *QGraphicsSceneResizeEvent) OldSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsSceneResizeEvent7oldSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:301
// index:0
// void setOldSize(const class QSizeF &)
func (this *QGraphicsSceneResizeEvent) SetOldSize(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEvent10setOldSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:303
// index:0
// QSizeF newSize()
func (this *QGraphicsSceneResizeEvent) NewSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsSceneResizeEvent7newSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:304
// index:0
// void setNewSize(const class QSizeF &)
func (this *QGraphicsSceneResizeEvent) SetNewSize(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEvent10setNewSizeERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

//  body block end
