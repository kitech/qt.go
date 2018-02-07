package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin

type QSGRootNode struct {
	*QSGNode
}

func (this *QSGRootNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGNode.GetCthis()
	}
}
func (this *QSGRootNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGNode = NewQSGNodeFromPointer(cthis)
}
func NewQSGRootNodeFromPointer(cthis unsafe.Pointer) *QSGRootNode {
	bcthis0 := NewQSGNodeFromPointer(cthis)
	return &QSGRootNode{bcthis0}
}
func (*QSGRootNode) NewFromPointer(cthis unsafe.Pointer) *QSGRootNode {
	return NewQSGRootNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGRootNode()
func NewQSGRootNode() *QSGRootNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGRootNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGRootNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGRootNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:303
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGRootNode()
func DeleteQSGRootNode(this *QSGRootNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGRootNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 88)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
