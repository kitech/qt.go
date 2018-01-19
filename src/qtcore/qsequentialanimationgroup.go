//  header block begin
// /usr/include/qt/QtCore/qsequentialanimationgroup.h
// #include <qsequentialanimationgroup.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QSequentialAnimationGroup struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSequentialAnimationGroup) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:59
// index:0
// void QSequentialAnimationGroup(class QObject *)
func NewQSequentialAnimationGroup(parent unsafe.Pointer) *QSequentialAnimationGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSequentialAnimationGroup{cthis}
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:60
// index:0
// virtual
// void ~QSequentialAnimationGroup()
func DeleteQSequentialAnimationGroup(*QSequentialAnimationGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:62
// index:0
// QPauseAnimation * addPause(int)
func (this *QSequentialAnimationGroup) AddPause(msecs int) {
	// 0: (, int msecs), (&msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup8addPauseEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:63
// index:0
// QPauseAnimation * insertPause(int, int)
func (this *QSequentialAnimationGroup) InsertPause(index int, msecs int) {
	// 0: (, int index, int msecs), (&index, &msecs)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup11insertPauseEii", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &msecs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:65
// index:0
// QAbstractAnimation * currentAnimation()
func (this *QSequentialAnimationGroup) CurrentAnimation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup16currentAnimationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:66
// index:0
// virtual
// int duration()
func (this *QSequentialAnimationGroup) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK25QSequentialAnimationGroup8durationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsequentialanimationgroup.h:69
// index:0
// void currentAnimationChanged(class QAbstractAnimation *)
func (this *QSequentialAnimationGroup) CurrentAnimationChanged(current unsafe.Pointer) {
	// 0: (, QAbstractAnimation * current), (current)
	rv, err := ffiqt.InvokeQtFunc6("_ZN25QSequentialAnimationGroup23currentAnimationChangedEP18QAbstractAnimation", ffiqt.FFI_TYPE_VOID, this.cthis, current)
	gopp.ErrPrint(err, rv)
}

//  body block end
