package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h
// #include <qgraphicssceneevent.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
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
	if this == nil {
		return nil
	} else {
		return this.QGraphicsSceneEvent.GetCthis()
	}
}
func (this *QGraphicsSceneResizeEvent) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsSceneEvent = NewQGraphicsSceneEventFromPointer(cthis)
}
func NewQGraphicsSceneResizeEventFromPointer(cthis unsafe.Pointer) *QGraphicsSceneResizeEvent {
	bcthis0 := NewQGraphicsSceneEventFromPointer(cthis)
	return &QGraphicsSceneResizeEvent{bcthis0}
}
func (*QGraphicsSceneResizeEvent) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSceneResizeEvent {
	return NewQGraphicsSceneResizeEventFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSceneResizeEvent()
func NewQGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEventC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSceneResizeEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSceneResizeEvent)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:298
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSceneResizeEvent()
func DeleteQGraphicsSceneResizeEvent(this *QGraphicsSceneResizeEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEventD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:300
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF oldSize()
func (this *QGraphicsSceneResizeEvent) OldSize() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsSceneResizeEvent7oldSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOldSize(const QSizeF &)
func (this *QGraphicsSceneResizeEvent) SetOldSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEvent10setOldSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:303
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF newSize()
func (this *QGraphicsSceneResizeEvent) NewSize() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QGraphicsSceneResizeEvent7newSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicssceneevent.h:304
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNewSize(const QSizeF &)
func (this *QGraphicsSceneResizeEvent) SetNewSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QGraphicsSceneResizeEvent10setNewSizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
