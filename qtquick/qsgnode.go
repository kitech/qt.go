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
// extern C begin: 39
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

type QSGNode struct {
	*qtrt.CObject
}
type QSGNode_ITF interface {
	QSGNode_PTR() *QSGNode
}

func (ptr *QSGNode) QSGNode_PTR() *QSGNode { return ptr }

func (this *QSGNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGNode) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSGNodeFromPointer(cthis unsafe.Pointer) *QSGNode {
	return &QSGNode{&qtrt.CObject{cthis}}
}
func (*QSGNode) NewFromPointer(cthis unsafe.Pointer) *QSGNode {
	return NewQSGNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGNode()
func NewQSGNode() *QSGNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:165
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QSGNode(enum QSGNode::NodeType)
func NewQSGNode_1(type_ int) *QSGNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNodeC2ENS_8NodeTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:131
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGNode()
func DeleteQSGNode(this *QSGNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 80)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * parent()
func (this *QSGNode) Parent() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeChildNode(QSGNode *)
func (this *QSGNode) RemoveChildNode(node QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode15removeChildNodeEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAllChildNodes()
func (this *QSGNode) RemoveAllChildNodes() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode19removeAllChildNodesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prependChildNode(QSGNode *)
func (this *QSGNode) PrependChildNode(node QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode16prependChildNodeEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendChildNode(QSGNode *)
func (this *QSGNode) AppendChildNode(node QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode15appendChildNodeEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertChildNodeBefore(QSGNode *, QSGNode *)
func (this *QSGNode) InsertChildNodeBefore(node QSGNode_ITF /*777 QSGNode **/, before QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if before != nil && before.QSGNode_PTR() != nil {
		convArg1 = before.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode21insertChildNodeBeforeEPS_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertChildNodeAfter(QSGNode *, QSGNode *)
func (this *QSGNode) InsertChildNodeAfter(node QSGNode_ITF /*777 QSGNode **/, after QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if after != nil && after.QSGNode_PTR() != nil {
		convArg1 = after.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode20insertChildNodeAfterEPS_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:141
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reparentChildNodesTo(QSGNode *)
func (this *QSGNode) ReparentChildNodesTo(newParent QSGNode_ITF /*777 QSGNode **/) {
	var convArg0 unsafe.Pointer
	if newParent != nil && newParent.QSGNode_PTR() != nil {
		convArg0 = newParent.QSGNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode20reparentChildNodesToEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:143
// index:0
// Public Visibility=Default Availability=Available
// [4] int childCount()
func (this *QSGNode) ChildCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode10childCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgnode.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGNode * childAtIndex(int)
func (this *QSGNode) ChildAtIndex(i int) *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode12childAtIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:145
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * firstChild()
func (this *QSGNode) FirstChild() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode10firstChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * lastChild()
func (this *QSGNode) LastChild() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode9lastChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:147
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * nextSibling()
func (this *QSGNode) NextSibling() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode11nextSiblingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:148
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * previousSibling()
func (this *QSGNode) PreviousSibling() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode15previousSiblingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:150
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGNode::NodeType type()
func (this *QSGNode) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearDirty()
func (this *QSGNode) ClearDirty() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode10clearDirtyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markDirty(QSGNode::DirtyState)
func (this *QSGNode) MarkDirty(bits int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode9markDirtyE6QFlagsINS_13DirtyStateBitEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bits)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGNode::DirtyState dirtyState()
func (this *QSGNode) DirtyState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode10dirtyStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:156
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSubtreeBlocked()
func (this *QSGNode) IsSubtreeBlocked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode16isSubtreeBlockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgnode.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGNode::Flags flags()
func (this *QSGNode) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(enum QSGNode::Flag, _Bool)
func (this *QSGNode) SetFlag(arg0 int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode7setFlagENS_4FlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QSGNode::Flags, _Bool)
func (this *QSGNode) SetFlags(arg0 int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode8setFlagsE6QFlagsINS_4FlagEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:162
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void preprocess()
func (this *QSGNode) Preprocess() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode10preprocessEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QSGNode__NodeType = int

const QSGNode__BasicNodeType QSGNode__NodeType = 0
const QSGNode__GeometryNodeType QSGNode__NodeType = 1
const QSGNode__TransformNodeType QSGNode__NodeType = 2
const QSGNode__ClipNodeType QSGNode__NodeType = 3
const QSGNode__OpacityNodeType QSGNode__NodeType = 4
const QSGNode__RootNodeType QSGNode__NodeType = 5
const QSGNode__RenderNodeType QSGNode__NodeType = 6

type QSGNode__Flag = int

const QSGNode__OwnedByParent QSGNode__Flag = 1
const QSGNode__UsePreprocess QSGNode__Flag = 2
const QSGNode__OwnsGeometry QSGNode__Flag = 65536
const QSGNode__OwnsMaterial QSGNode__Flag = 131072
const QSGNode__OwnsOpaqueMaterial QSGNode__Flag = 262144
const QSGNode__IsVisitableNode QSGNode__Flag = 16777216

type QSGNode__DirtyStateBit = int

const QSGNode__DirtySubtreeBlocked QSGNode__DirtyStateBit = 128
const QSGNode__DirtyMatrix QSGNode__DirtyStateBit = 256
const QSGNode__DirtyNodeAdded QSGNode__DirtyStateBit = 1024
const QSGNode__DirtyNodeRemoved QSGNode__DirtyStateBit = 2048
const QSGNode__DirtyGeometry QSGNode__DirtyStateBit = 4096
const QSGNode__DirtyMaterial QSGNode__DirtyStateBit = 8192
const QSGNode__DirtyOpacity QSGNode__DirtyStateBit = 16384
const QSGNode__DirtyForceUpdate QSGNode__DirtyStateBit = 32768
const QSGNode__DirtyUsePreprocess QSGNode__DirtyStateBit = 2
const QSGNode__DirtyPropagationMask QSGNode__DirtyStateBit = 50432

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
