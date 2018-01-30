package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

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
import "gopp"
import "qt.go/cffiqt"
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
type QSGNodeVisitor struct {
	*qtrt.CObject
}

func (this *QSGNodeVisitor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGNodeVisitor) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSGNodeVisitorFromPointer(cthis unsafe.Pointer) *QSGNodeVisitor {
	return &QSGNodeVisitor{&qtrt.CObject{cthis}}
}
func (*QSGNodeVisitor) NewFromPointer(cthis unsafe.Pointer) *QSGNodeVisitor {
	return NewQSGNodeVisitorFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:337
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGNodeVisitor()
func DeleteQSGNodeVisitor(*QSGNodeVisitor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitorD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:340
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterTransformNode(QSGTransformNode *)
func (this *QSGNodeVisitor) EnterTransformNode(arg0 *QSGTransformNode /*777 QSGTransformNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor18enterTransformNodeEP16QSGTransformNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:341
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveTransformNode(QSGTransformNode *)
func (this *QSGNodeVisitor) LeaveTransformNode(arg0 *QSGTransformNode /*777 QSGTransformNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor18leaveTransformNodeEP16QSGTransformNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:342
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterClipNode(QSGClipNode *)
func (this *QSGNodeVisitor) EnterClipNode(arg0 *QSGClipNode /*777 QSGClipNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor13enterClipNodeEP11QSGClipNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:343
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveClipNode(QSGClipNode *)
func (this *QSGNodeVisitor) LeaveClipNode(arg0 *QSGClipNode /*777 QSGClipNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor13leaveClipNodeEP11QSGClipNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:344
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterGeometryNode(QSGGeometryNode *)
func (this *QSGNodeVisitor) EnterGeometryNode(arg0 *QSGGeometryNode /*777 QSGGeometryNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor17enterGeometryNodeEP15QSGGeometryNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:345
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveGeometryNode(QSGGeometryNode *)
func (this *QSGNodeVisitor) LeaveGeometryNode(arg0 *QSGGeometryNode /*777 QSGGeometryNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor17leaveGeometryNodeEP15QSGGeometryNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:346
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterOpacityNode(QSGOpacityNode *)
func (this *QSGNodeVisitor) EnterOpacityNode(arg0 *QSGOpacityNode /*777 QSGOpacityNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor16enterOpacityNodeEP14QSGOpacityNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:347
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveOpacityNode(QSGOpacityNode *)
func (this *QSGNodeVisitor) LeaveOpacityNode(arg0 *QSGOpacityNode /*777 QSGOpacityNode **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor16leaveOpacityNodeEP14QSGOpacityNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:348
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void visitNode(QSGNode *)
func (this *QSGNodeVisitor) VisitNode(n *QSGNode /*777 QSGNode **/) {
	var convArg0 = n.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor9visitNodeEP7QSGNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:349
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void visitChildren(QSGNode *)
func (this *QSGNodeVisitor) VisitChildren(n *QSGNode /*777 QSGNode **/) {
	var convArg0 = n.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QSGNodeVisitor13visitChildrenEP7QSGNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
