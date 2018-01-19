//  header block begin
// /usr/include/qt/QtCore/qpauseanimation.h
// #include <qpauseanimation.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QPauseAnimation struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qpauseanimation.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPauseAnimation) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QPauseAnimation10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:57
// index:0
// void QPauseAnimation(class QObject *)
func NewQPauseAnimation(parent unsafe.Pointer) *QPauseAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QPauseAnimation{cthis}
}

// /usr/include/qt/QtCore/qpauseanimation.h:58
// index:1
// void QPauseAnimation(int, class QObject *)
func NewQPauseAnimation_1(msecs int, parent unsafe.Pointer) *QPauseAnimation {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationC2EiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, &msecs, parent)
	gopp.ErrPrint(err, rv)
	return &QPauseAnimation{cthis}
}

// /usr/include/qt/QtCore/qpauseanimation.h:59
// index:0
// virtual
// void ~QPauseAnimation()
func DeleteQPauseAnimation(*QPauseAnimation) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:61
// index:0
// virtual
// int duration()
func (this *QPauseAnimation) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QPauseAnimation8durationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qpauseanimation.h:62
// index:0
// void setDuration(int)
func (this *QPauseAnimation) SetDuration(msecs int) {
	// 0: (, int msecs), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QPauseAnimation11setDurationEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs)
	gopp.ErrPrint(err, rv)
}

//  body block end
