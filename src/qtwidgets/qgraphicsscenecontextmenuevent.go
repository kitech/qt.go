//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QGraphicsSceneContextMenuEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:174
// index:0
// void QGraphicsSceneContextMenuEvent(enum QEvent::Type)
func NewQGraphicsSceneContextMenuEvent(type_ int) *QGraphicsSceneContextMenuEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	return &QGraphicsSceneContextMenuEvent{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:175
// index:0
// virtual
// void ~QGraphicsSceneContextMenuEvent()
func DeleteQGraphicsSceneContextMenuEvent(*QGraphicsSceneContextMenuEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:177
// index:0
// QPointF pos()
func (this *QGraphicsSceneContextMenuEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:178
// index:0
// void setPos(const class QPointF &)
func (this *QGraphicsSceneContextMenuEvent) SetPos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:180
// index:0
// QPointF scenePos()
func (this *QGraphicsSceneContextMenuEvent) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent8scenePosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:181
// index:0
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneContextMenuEvent) SetScenePos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:183
// index:0
// QPoint screenPos()
func (this *QGraphicsSceneContextMenuEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:184
// index:0
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneContextMenuEvent) SetScreenPos(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:186
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneContextMenuEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:189
// index:0
// QGraphicsSceneContextMenuEvent::Reason reason()
func (this *QGraphicsSceneContextMenuEvent) Reason() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK30QGraphicsSceneContextMenuEvent6reasonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:190
// index:0
// void setReason(enum QGraphicsSceneContextMenuEvent::Reason)
func (this *QGraphicsSceneContextMenuEvent) SetReason(reason int) {
	// 0: (, QGraphicsSceneContextMenuEvent::Reason reason), (&reason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN30QGraphicsSceneContextMenuEvent9setReasonENS_6ReasonE", ffiqt.FFI_TYPE_VOID, this.cthis, &reason)
	gopp.ErrPrint(err, rv)
}

//  body block end
