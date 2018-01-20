//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QGraphicsSceneMoveEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneMoveEvent) GetCthis() unsafe.Pointer {
	return this.QGraphicsSceneEvent.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:313
// index:0
// void QGraphicsSceneMoveEvent()
func NewQGraphicsSceneMoveEvent() *QGraphicsSceneMoveEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEventC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneMoveEventFromPointer(cthis)
	return gothis
}
func NewQGraphicsSceneMoveEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneMoveEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneMoveEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:314
// index:0
// virtual
// void ~QGraphicsSceneMoveEvent()
func DeleteQGraphicsSceneMoveEvent(*QGraphicsSceneMoveEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:316
// index:0
// QPointF oldPos()
func (this *QGraphicsSceneMoveEvent) OldPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSceneMoveEvent6oldPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:317
// index:0
// void setOldPos(const class QPointF &)
func (this *QGraphicsSceneMoveEvent) SetOldPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEvent9setOldPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:319
// index:0
// QPointF newPos()
func (this *QGraphicsSceneMoveEvent) NewPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSceneMoveEvent6newPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:320
// index:0
// void setNewPos(const class QPointF &)
func (this *QGraphicsSceneMoveEvent) SetNewPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneMoveEvent9setNewPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

//  body block end
