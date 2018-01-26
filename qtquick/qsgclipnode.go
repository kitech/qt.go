package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QSGClipNode struct {
	*QSGBasicGeometryNode
}

func (this *QSGClipNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGBasicGeometryNode.GetCthis()
	}
}
func (this *QSGClipNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGBasicGeometryNode = NewQSGBasicGeometryNodeFromPointer(cthis)
}
func NewQSGClipNodeFromPointer(cthis unsafe.Pointer) *QSGClipNode {
	bcthis0 := NewQSGBasicGeometryNodeFromPointer(cthis)
	return &QSGClipNode{bcthis0}
}
func (*QSGClipNode) NewFromPointer(cthis unsafe.Pointer) *QSGClipNode {
	return NewQSGClipNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:264
// index:0
// Public
// void QSGClipNode()
func NewQSGClipNode() *QSGClipNode {
	cthis := qtrt.Calloc(1, 256) // 152
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGClipNodeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGClipNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:265
// index:0
// Public virtual
// void ~QSGClipNode()
func DeleteQSGClipNode(*QSGClipNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGClipNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:267
// index:0
// Public
// void setIsRectangular(_Bool)
func (this *QSGClipNode) SetIsRectangular(rectHint bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGClipNode16setIsRectangularEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rectHint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:268
// index:0
// Public inline
// bool isRectangular()
func (this *QSGClipNode) IsRectangular() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGClipNode13isRectangularEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgnode.h:270
// index:0
// Public
// void setClipRect(const class QRectF &)
func (this *QSGClipNode) SetClipRect(arg0 *qtcore.QRectF) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSGClipNode11setClipRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:271
// index:0
// Public inline
// QRectF clipRect()
func (this *QSGClipNode) ClipRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QSGClipNode8clipRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
