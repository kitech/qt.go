//  header block begin
// /usr/include/qt/QtGui/qdrag.h
// #include <qdrag.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QDrag struct {
	*qtcore.QObject
}

func (this *QDrag) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qdrag.h:59
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDrag) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:62
// index:0
// void QDrag(class QObject *)
func NewQDrag(dragSource unsafe.Pointer) *QDrag {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDragC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, dragSource)
	gopp.ErrPrint(err, rv)
	gothis := NewQDragFromPointer(cthis)
	return gothis
}
func NewQDragFromPointer(cthis unsafe.Pointer) *QDrag {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDrag{bcthis0}
}

// /usr/include/qt/QtGui/qdrag.h:63
// index:0
// virtual
// void ~QDrag()
func DeleteQDrag(*QDrag) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDragD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:65
// index:0
// void setMimeData(class QMimeData *)
func (this *QDrag) SetMimeData(data unsafe.Pointer) {
	// 0: (, data QMimeData *), (data)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag11setMimeDataEP9QMimeData", ffiqt.FFI_TYPE_VOID, this.GetCthis(), data)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:66
// index:0
// QMimeData * mimeData()
func (this *QDrag) MimeData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag8mimeDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:68
// index:0
// void setPixmap(const class QPixmap &)
func (this *QDrag) SetPixmap(arg0 unsafe.Pointer) {
	// 0: (, const QPixmap & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag9setPixmapERK7QPixmap", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:69
// index:0
// QPixmap pixmap()
func (this *QDrag) Pixmap() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag6pixmapEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:71
// index:0
// void setHotSpot(const class QPoint &)
func (this *QDrag) SetHotSpot(hotspot unsafe.Pointer) {
	// 0: (, hotspot const QPoint &), (hotspot)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag10setHotSpotERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), hotspot)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:72
// index:0
// QPoint hotSpot()
func (this *QDrag) HotSpot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag7hotSpotEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:74
// index:0
// QObject * source()
func (this *QDrag) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag6sourceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:75
// index:0
// QObject * target()
func (this *QDrag) Target() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag6targetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:77
// index:0
// Qt::DropAction start(Qt::DropActions)
func (this *QDrag) Start(supportedActions int) {
	// 0: (, QFlags<Qt::DropAction> supportedActions), (&supportedActions)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag5startE6QFlagsIN2Qt10DropActionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &supportedActions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:78
// index:0
// Qt::DropAction exec(Qt::DropActions)
func (this *QDrag) Exec(supportedActions int) {
	// 0: (, QFlags<Qt::DropAction> supportedActions), (&supportedActions)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag4execE6QFlagsIN2Qt10DropActionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &supportedActions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:79
// index:1
// Qt::DropAction exec(Qt::DropActions, Qt::DropAction)
func (this *QDrag) Exec_1(supportedActions int, defaultAction int) {
	// 1: (, QFlags<Qt::DropAction> supportedActions, defaultAction Qt::DropAction), (&supportedActions, &defaultAction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag4execE6QFlagsIN2Qt10DropActionEES2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &supportedActions, &defaultAction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:81
// index:0
// void setDragCursor(const class QPixmap &, Qt::DropAction)
func (this *QDrag) SetDragCursor(cursor unsafe.Pointer, action int) {
	// 0: (, cursor const QPixmap &, action Qt::DropAction), (cursor, &action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag13setDragCursorERK7QPixmapN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor, &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:82
// index:0
// QPixmap dragCursor(Qt::DropAction)
func (this *QDrag) DragCursor(action int) {
	// 0: (, action Qt::DropAction), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag10dragCursorEN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:84
// index:0
// Qt::DropActions supportedActions()
func (this *QDrag) SupportedActions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag16supportedActionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:85
// index:0
// Qt::DropAction defaultAction()
func (this *QDrag) DefaultAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QDrag13defaultActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:87
// index:0
// static
// void cancel()
func (this *QDrag) Cancel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag6cancelEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QDrag_Cancel() {
	// 0: (), ()
	var nilthis *QDrag
	nilthis.Cancel()
}

// /usr/include/qt/QtGui/qdrag.h:90
// index:0
// void actionChanged(Qt::DropAction)
func (this *QDrag) ActionChanged(action int) {
	// 0: (, action Qt::DropAction), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag13actionChangedEN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qdrag.h:91
// index:0
// void targetChanged(class QObject *)
func (this *QDrag) TargetChanged(newTarget unsafe.Pointer) {
	// 0: (, newTarget QObject *), (newTarget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QDrag13targetChangedEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), newTarget)
	gopp.ErrPrint(err, rv)
}

//  body block end
