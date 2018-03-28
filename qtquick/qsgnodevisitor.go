package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

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
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// void enterTransformNode(QSGTransformNode *)
func (this *QSGNodeVisitor) InheritEnterTransformNode(f func(arg0 *QSGTransformNode /*777 QSGTransformNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "enterTransformNode", f)
}

// void leaveTransformNode(QSGTransformNode *)
func (this *QSGNodeVisitor) InheritLeaveTransformNode(f func(arg0 *QSGTransformNode /*777 QSGTransformNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveTransformNode", f)
}

// void enterClipNode(QSGClipNode *)
func (this *QSGNodeVisitor) InheritEnterClipNode(f func(arg0 *QSGClipNode /*777 QSGClipNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "enterClipNode", f)
}

// void leaveClipNode(QSGClipNode *)
func (this *QSGNodeVisitor) InheritLeaveClipNode(f func(arg0 *QSGClipNode /*777 QSGClipNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveClipNode", f)
}

// void enterGeometryNode(QSGGeometryNode *)
func (this *QSGNodeVisitor) InheritEnterGeometryNode(f func(arg0 *QSGGeometryNode /*777 QSGGeometryNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "enterGeometryNode", f)
}

// void leaveGeometryNode(QSGGeometryNode *)
func (this *QSGNodeVisitor) InheritLeaveGeometryNode(f func(arg0 *QSGGeometryNode /*777 QSGGeometryNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveGeometryNode", f)
}

// void enterOpacityNode(QSGOpacityNode *)
func (this *QSGNodeVisitor) InheritEnterOpacityNode(f func(arg0 *QSGOpacityNode /*777 QSGOpacityNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "enterOpacityNode", f)
}

// void leaveOpacityNode(QSGOpacityNode *)
func (this *QSGNodeVisitor) InheritLeaveOpacityNode(f func(arg0 *QSGOpacityNode /*777 QSGOpacityNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "leaveOpacityNode", f)
}

// void visitNode(QSGNode *)
func (this *QSGNodeVisitor) InheritVisitNode(f func(n *QSGNode /*777 QSGNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "visitNode", f)
}

// void visitChildren(QSGNode *)
func (this *QSGNodeVisitor) InheritVisitChildren(f func(n *QSGNode /*777 QSGNode **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "visitChildren", f)
}

/*

 */
type QSGNodeVisitor struct {
	*qtrt.CObject
}
type QSGNodeVisitor_ITF interface {
	QSGNodeVisitor_PTR() *QSGNodeVisitor
}

func (ptr *QSGNodeVisitor) QSGNodeVisitor_PTR() *QSGNodeVisitor { return ptr }

func (this *QSGNodeVisitor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGNodeVisitor) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
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

/*

 */
func DeleteQSGNodeVisitor(this *QSGNodeVisitor) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:340
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterTransformNode(QSGTransformNode *)

/*

 */
func (this *QSGNodeVisitor) EnterTransformNode(arg0 QSGTransformNode_ITF /*777 QSGTransformNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGTransformNode_PTR() != nil {
		convArg0 = arg0.QSGTransformNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor18enterTransformNodeEP16QSGTransformNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:341
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveTransformNode(QSGTransformNode *)

/*

 */
func (this *QSGNodeVisitor) LeaveTransformNode(arg0 QSGTransformNode_ITF /*777 QSGTransformNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGTransformNode_PTR() != nil {
		convArg0 = arg0.QSGTransformNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor18leaveTransformNodeEP16QSGTransformNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:342
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterClipNode(QSGClipNode *)

/*

 */
func (this *QSGNodeVisitor) EnterClipNode(arg0 QSGClipNode_ITF /*777 QSGClipNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGClipNode_PTR() != nil {
		convArg0 = arg0.QSGClipNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor13enterClipNodeEP11QSGClipNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:343
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveClipNode(QSGClipNode *)

/*

 */
func (this *QSGNodeVisitor) LeaveClipNode(arg0 QSGClipNode_ITF /*777 QSGClipNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGClipNode_PTR() != nil {
		convArg0 = arg0.QSGClipNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor13leaveClipNodeEP11QSGClipNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:344
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterGeometryNode(QSGGeometryNode *)

/*

 */
func (this *QSGNodeVisitor) EnterGeometryNode(arg0 QSGGeometryNode_ITF /*777 QSGGeometryNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGGeometryNode_PTR() != nil {
		convArg0 = arg0.QSGGeometryNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor17enterGeometryNodeEP15QSGGeometryNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:345
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveGeometryNode(QSGGeometryNode *)

/*

 */
func (this *QSGNodeVisitor) LeaveGeometryNode(arg0 QSGGeometryNode_ITF /*777 QSGGeometryNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGGeometryNode_PTR() != nil {
		convArg0 = arg0.QSGGeometryNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor17leaveGeometryNodeEP15QSGGeometryNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:346
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void enterOpacityNode(QSGOpacityNode *)

/*

 */
func (this *QSGNodeVisitor) EnterOpacityNode(arg0 QSGOpacityNode_ITF /*777 QSGOpacityNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGOpacityNode_PTR() != nil {
		convArg0 = arg0.QSGOpacityNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor16enterOpacityNodeEP14QSGOpacityNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:347
// index:0
// Protected inline virtual Visibility=Default Availability=Available
// [-2] void leaveOpacityNode(QSGOpacityNode *)

/*

 */
func (this *QSGNodeVisitor) LeaveOpacityNode(arg0 QSGOpacityNode_ITF /*777 QSGOpacityNode **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSGOpacityNode_PTR() != nil {
		convArg0 = arg0.QSGOpacityNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor16leaveOpacityNodeEP14QSGOpacityNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:348
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void visitNode(QSGNode *)

/*

 */
func (this *QSGNodeVisitor) VisitNode(n QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if n != nil && n.QSGNode_PTR() != nil {
		convArg0 = n.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor9visitNodeEP7QSGNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:349
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void visitChildren(QSGNode *)

/*

 */
func (this *QSGNodeVisitor) VisitChildren(n QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if n != nil && n.QSGNode_PTR() != nil {
		convArg0 = n.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QSGNodeVisitor13visitChildrenEP7QSGNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
