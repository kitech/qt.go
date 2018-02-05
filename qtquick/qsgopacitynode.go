package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSGOpacityNode struct {
	*QSGNode
}

func (this *QSGOpacityNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGNode.GetCthis()
	}
}
func (this *QSGOpacityNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGNode = NewQSGNodeFromPointer(cthis)
}
func NewQSGOpacityNodeFromPointer(cthis unsafe.Pointer) *QSGOpacityNode {
	bcthis0 := NewQSGNodeFromPointer(cthis)
	return &QSGOpacityNode{bcthis0}
}
func (*QSGOpacityNode) NewFromPointer(cthis unsafe.Pointer) *QSGOpacityNode {
	return NewQSGOpacityNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGOpacityNode()
func NewQSGOpacityNode() *QSGOpacityNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGOpacityNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGOpacityNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGOpacityNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:320
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGOpacityNode()
func DeleteQSGOpacityNode(this *QSGOpacityNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGOpacityNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 96)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)
func (this *QSGOpacityNode) SetOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGOpacityNode10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:323
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal opacity()
func (this *QSGOpacityNode) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSGOpacityNode7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qsgnode.h:325
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCombinedOpacity(qreal)
func (this *QSGOpacityNode) SetCombinedOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGOpacityNode18setCombinedOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:326
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal combinedOpacity()
func (this *QSGOpacityNode) CombinedOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSGOpacityNode15combinedOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qsgnode.h:328
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSubtreeBlocked()
func (this *QSGOpacityNode) IsSubtreeBlocked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QSGOpacityNode16isSubtreeBlockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

//  body block end
