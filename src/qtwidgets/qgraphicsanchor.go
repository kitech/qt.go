//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h
// #include <qgraphicsanchorlayout.h>
// #include <QtWidgets>
package qtwidgets

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
type QGraphicsAnchor struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsAnchor) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsAnchor10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:61
// index:0
// void setSpacing(qreal)
func (this *QGraphicsAnchor) SetSpacing(spacing float64) {
	// 0: (, qreal spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchor10setSpacingEd", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:62
// index:0
// void unsetSpacing()
func (this *QGraphicsAnchor) UnsetSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchor12unsetSpacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:63
// index:0
// qreal spacing()
func (this *QGraphicsAnchor) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsAnchor7spacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:64
// index:0
// void setSizePolicy(class QSizePolicy::Policy)
func (this *QGraphicsAnchor) SetSizePolicy(policy int) {
	// 0: (, QSizePolicy::Policy policy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchor13setSizePolicyEN11QSizePolicy6PolicyE", ffiqt.FFI_TYPE_VOID, this.cthis, &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:65
// index:0
// QSizePolicy::Policy sizePolicy()
func (this *QGraphicsAnchor) SizePolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsAnchor10sizePolicyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:66
// index:0
// virtual
// void ~QGraphicsAnchor()
func DeleteQGraphicsAnchor(*QGraphicsAnchor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
