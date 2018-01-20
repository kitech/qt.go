//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QGraphicsSceneHelpEvent struct {
	*QGraphicsSceneEvent
}

func (this *QGraphicsSceneHelpEvent) GetCthis() unsafe.Pointer {
	return this.QGraphicsSceneEvent.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:234
// index:0
// void QGraphicsSceneHelpEvent(enum QEvent::Type)
func NewQGraphicsSceneHelpEvent(type_ int) *QGraphicsSceneHelpEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneHelpEventFromPointer(cthis)
	return gothis
}
func NewQGraphicsSceneHelpEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneHelpEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneHelpEvent{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:235
// index:0
// virtual
// void ~QGraphicsSceneHelpEvent()
func DeleteQGraphicsSceneHelpEvent(*QGraphicsSceneHelpEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:237
// index:0
// QPointF scenePos()
func (this *QGraphicsSceneHelpEvent) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSceneHelpEvent8scenePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:238
// index:0
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneHelpEvent) SetScenePos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:240
// index:0
// QPoint screenPos()
func (this *QGraphicsSceneHelpEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSceneHelpEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:241
// index:0
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneHelpEvent) SetScreenPos(pos unsafe.Pointer) {
	// 0: (, pos const QPoint &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSceneHelpEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

//  body block end
