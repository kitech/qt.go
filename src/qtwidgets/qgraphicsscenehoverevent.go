//  header block begin
// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QGraphicsSceneHoverEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:201
// index:0
// void QGraphicsSceneHoverEvent(enum QEvent::Type)
func NewQGraphicsSceneHoverEvent(type_ int) *QGraphicsSceneHoverEvent {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEventC2EN6QEvent4TypeE", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	return &QGraphicsSceneHoverEvent{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:202
// index:0
// virtual
// void ~QGraphicsSceneHoverEvent()
func DeleteQGraphicsSceneHoverEvent(*QGraphicsSceneHoverEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:204
// index:0
// QPointF pos()
func (this *QGraphicsSceneHoverEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:205
// index:0
// void setPos(const class QPointF &)
func (this *QGraphicsSceneHoverEvent) SetPos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent6setPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:207
// index:0
// QPointF scenePos()
func (this *QGraphicsSceneHoverEvent) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent8scenePosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:208
// index:0
// void setScenePos(const class QPointF &)
func (this *QGraphicsSceneHoverEvent) SetScenePos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent11setScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:210
// index:0
// QPoint screenPos()
func (this *QGraphicsSceneHoverEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:211
// index:0
// void setScreenPos(const class QPoint &)
func (this *QGraphicsSceneHoverEvent) SetScreenPos(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent12setScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:213
// index:0
// QPointF lastPos()
func (this *QGraphicsSceneHoverEvent) LastPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent7lastPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:214
// index:0
// void setLastPos(const class QPointF &)
func (this *QGraphicsSceneHoverEvent) SetLastPos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent10setLastPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:216
// index:0
// QPointF lastScenePos()
func (this *QGraphicsSceneHoverEvent) LastScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent12lastScenePosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:217
// index:0
// void setLastScenePos(const class QPointF &)
func (this *QGraphicsSceneHoverEvent) SetLastScenePos(pos unsafe.Pointer) {
	// 0: (, const QPointF & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent15setLastScenePosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:219
// index:0
// QPoint lastScreenPos()
func (this *QGraphicsSceneHoverEvent) LastScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent13lastScreenPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:220
// index:0
// void setLastScreenPos(const class QPoint &)
func (this *QGraphicsSceneHoverEvent) SetLastScreenPos(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN24QGraphicsSceneHoverEvent16setLastScreenPosERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:222
// index:0
// Qt::KeyboardModifiers modifiers()
func (this *QGraphicsSceneHoverEvent) Modifiers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK24QGraphicsSceneHoverEvent9modifiersEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
