package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

/*

 */
type QTapGesture struct {
	*QGesture
}
type QTapGesture_ITF interface {
	QGesture_ITF
	QTapGesture_PTR() *QTapGesture
}

func (ptr *QTapGesture) QTapGesture_PTR() *QTapGesture { return ptr }

func (this *QTapGesture) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGesture.GetCthis()
	}
}
func (this *QTapGesture) SetCthis(cthis unsafe.Pointer) {
	this.QGesture = NewQGestureFromPointer(cthis)
}
func NewQTapGestureFromPointer(cthis unsafe.Pointer) *QTapGesture {
	bcthis0 := NewQGestureFromPointer(cthis)
	return &QTapGesture{bcthis0}
}
func (*QTapGesture) NewFromPointer(cthis unsafe.Pointer) *QTapGesture {
	return NewQTapGestureFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgesture.h:236
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTapGesture) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTapGesture10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgesture.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTapGesture(QObject *)

/*

 */
func NewQTapGesture(parent qtcore.QObject_ITF /*777 QObject **/) *QTapGesture {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTapGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTapGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTapGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTapGesture(QObject *)

/*

 */
func NewQTapGesture__() *QTapGesture {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTapGestureC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTapGestureFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTapGesture")
	return gothis
}

// /usr/include/qt/QtWidgets/qgesture.h:243
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTapGesture()

/*

 */
func DeleteQTapGesture(this *QTapGesture) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTapGestureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgesture.h:245
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position() const

/*

 */
func (this *QTapGesture) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTapGesture8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgesture.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)

/*

 */
func (this *QTapGesture) SetPosition(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTapGesture11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
