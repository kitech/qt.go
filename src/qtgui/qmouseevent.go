//  header block begin
// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>
package qtgui

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
type QMouseEvent struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qevent.h:118
// index:0
// virtual
// void ~QMouseEvent()
func DeleteQMouseEvent(*QMouseEvent) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEventD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:121
// index:0
// inline
// QPoint pos()
func (this *QMouseEvent) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent3posEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:122
// index:0
// inline
// QPoint globalPos()
func (this *QMouseEvent) GlobalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9globalPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:123
// index:0
// inline
// int x()
func (this *QMouseEvent) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:124
// index:0
// inline
// int y()
func (this *QMouseEvent) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:125
// index:0
// inline
// int globalX()
func (this *QMouseEvent) GlobalX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7globalXEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:126
// index:0
// inline
// int globalY()
func (this *QMouseEvent) GlobalY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7globalYEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:128
// index:0
// inline
// const QPointF & localPos()
func (this *QMouseEvent) LocalPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent8localPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:129
// index:0
// inline
// const QPointF & windowPos()
func (this *QMouseEvent) WindowPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9windowPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:130
// index:0
// inline
// const QPointF & screenPos()
func (this *QMouseEvent) ScreenPos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent9screenPosEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:132
// index:0
// inline
// Qt::MouseButton button()
func (this *QMouseEvent) Button() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent6buttonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:133
// index:0
// inline
// Qt::MouseButtons buttons()
func (this *QMouseEvent) Buttons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent7buttonsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:135
// index:0
// inline
// void setLocalPos(const class QPointF &)
func (this *QMouseEvent) SetLocalPos(localPosition unsafe.Pointer) {
	// 0: (, const QPointF & localPosition), (localPosition)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QMouseEvent11setLocalPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.cthis, localPosition)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:141
// index:0
// Qt::MouseEventSource source()
func (this *QMouseEvent) Source() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent6sourceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qevent.h:142
// index:0
// Qt::MouseEventFlags flags()
func (this *QMouseEvent) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QMouseEvent5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
