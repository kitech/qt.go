package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h
// #include <qgraphicsanchorlayout.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QGraphicsAnchor struct {
	*qtcore.QObject
}

func (this *QGraphicsAnchor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGraphicsAnchor) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGraphicsAnchorFromPointer(cthis unsafe.Pointer) *QGraphicsAnchor {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGraphicsAnchor{bcthis0}
}
func (*QGraphicsAnchor) NewFromPointer(cthis unsafe.Pointer) *QGraphicsAnchor {
	return NewQGraphicsAnchorFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsAnchor) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsAnchor10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:61
// index:0
// Public
// void setSpacing(qreal)
func (this *QGraphicsAnchor) SetSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchor10setSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:62
// index:0
// Public
// void unsetSpacing()
func (this *QGraphicsAnchor) UnsetSpacing() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchor12unsetSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:63
// index:0
// Public
// qreal spacing()
func (this *QGraphicsAnchor) Spacing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsAnchor7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:64
// index:0
// Public
// void setSizePolicy(class QSizePolicy::Policy)
func (this *QGraphicsAnchor) SetSizePolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchor13setSizePolicyEN11QSizePolicy6PolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:65
// index:0
// Public
// QSizePolicy::Policy sizePolicy()
func (this *QGraphicsAnchor) SizePolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsAnchor10sizePolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:66
// index:0
// Public virtual
// void ~QGraphicsAnchor()
func DeleteQGraphicsAnchor(*QGraphicsAnchor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsAnchorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end
