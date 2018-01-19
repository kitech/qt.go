//  header block begin
// /usr/include/qt/QtCore/qparallelanimationgroup.h
// #include <qparallelanimationgroup.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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
type QParallelAnimationGroup struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QParallelAnimationGroup) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:56
// index:0
// void QParallelAnimationGroup(class QObject *)
func NewQParallelAnimationGroup(parent unsafe.Pointer) *QParallelAnimationGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroupC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QParallelAnimationGroup{cthis}
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:57
// index:0
// virtual
// void ~QParallelAnimationGroup()
func DeleteQParallelAnimationGroup(*QParallelAnimationGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QParallelAnimationGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qparallelanimationgroup.h:59
// index:0
// virtual
// int duration()
func (this *QParallelAnimationGroup) Duration() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QParallelAnimationGroup8durationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
