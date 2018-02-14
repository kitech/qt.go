package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h
// #include <qgraphicsanchorlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QGraphicsAnchor struct {
	*qtcore.QObject
}
type QGraphicsAnchor_ITF interface {
	qtcore.QObject_ITF
	QGraphicsAnchor_PTR() *QGraphicsAnchor
}

func (ptr *QGraphicsAnchor) QGraphicsAnchor_PTR() *QGraphicsAnchor { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QGraphicsAnchor) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsAnchor10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(qreal)
func (this *QGraphicsAnchor) SetSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsAnchor10setSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetSpacing()
func (this *QGraphicsAnchor) UnsetSpacing() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsAnchor12unsetSpacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal spacing()
func (this *QGraphicsAnchor) Spacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsAnchor7spacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy::Policy)
func (this *QGraphicsAnchor) SetSizePolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsAnchor13setSizePolicyEN11QSizePolicy6PolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy::Policy sizePolicy()
func (this *QGraphicsAnchor) SizePolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsAnchor10sizePolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsanchorlayout.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsAnchor()
func DeleteQGraphicsAnchor(this *QGraphicsAnchor) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsAnchorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
