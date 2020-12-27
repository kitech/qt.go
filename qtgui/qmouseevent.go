package qtgui

// /usr/include/qt/QtGui/qevent.h
// #include <qevent.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*
 */
// size 104
type QMouseEvent struct {
	*QInputEvent
}
type QMouseEvent_ITF interface {
	QInputEvent_ITF
	QMouseEvent_PTR() *QMouseEvent
}

func (ptr *QMouseEvent) QMouseEvent_PTR() *QMouseEvent { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QMouseEventFromptr(cthis Voidptr) *QMouseEvent {
	bcthis0 := QInputEventFromptr(cthis)
	return &QMouseEvent{bcthis0}
}
func (*QMouseEvent) Fromptr(cthis Voidptr) *QMouseEvent {
	return QMouseEventFromptr(cthis)
}

// /usr/include/qt/QtGui/qevent.h:121
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] QPoint pos() const

/*
 */
func (this *QMouseEvent) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.Qtcc3(1659492746, "_ZNK11QMouseEvent3posEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := qtcore.QPointFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:122
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [8] QPoint globalPos() const

/*
 */
func (this *QMouseEvent) GlobalPos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.Qtcc3(2976136822, "_ZNK11QMouseEvent9globalPosEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	cthis := qtrt.Malloc(8)
	qtrt.Cmemcpy(cthis, rv.Addr(), 8)
	rv2 := qtcore.QPointFromptr(cthis) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:123
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int x() const

/*
 */
func (this *QMouseEvent) X() int {
	rv, err := qtrt.Qtcc3(1953764823, "_ZNK11QMouseEvent1xEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qevent.h:124
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int y() const

/*
 */
func (this *QMouseEvent) Y() int {
	rv, err := qtrt.Qtcc3(1974894560, "_ZNK11QMouseEvent1yEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qevent.h:125
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int globalX() const

/*
 */
func (this *QMouseEvent) GlobalX() int {
	rv, err := qtrt.Qtcc3(220043088, "_ZNK11QMouseEvent7globalXEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qevent.h:126
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [4] int globalY() const

/*
 */
func (this *QMouseEvent) GlobalY() int {
	rv, err := qtrt.Qtcc3(216005991, "_ZNK11QMouseEvent7globalYEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtGui/qevent.h:128
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [16] const QPointF & localPos() const

/*
 */
func (this *QMouseEvent) LocalPos() *qtcore.QPointF {
	rv, err := qtrt.Qtcc3(1441673618, "_ZNK11QMouseEvent8localPosEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv2 := qtcore.QPointFFromptr(rv.Ptr()) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:129
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [16] const QPointF & windowPos() const

/*
 */
func (this *QMouseEvent) WindowPos() *qtcore.QPointF {
	rv, err := qtrt.Qtcc3(293709112, "_ZNK11QMouseEvent9windowPosEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv2 := qtcore.QPointFFromptr(rv.Ptr()) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:130
// index:0
// Public inline Direct Visibility=Default Availability=Available
// [16] const QPointF & screenPos() const

/*
 */
func (this *QMouseEvent) ScreenPos() *qtcore.QPointF {
	rv, err := qtrt.Qtcc3(2962094997, "_ZNK11QMouseEvent9screenPosEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	rv2 := qtcore.QPointFFromptr(rv.Ptr()) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

func DeleteQMouseEvent(this *QMouseEvent) {
	rv, err := qtrt.Qtcc3(3707978826, "_ZN11QMouseEventD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10071() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
