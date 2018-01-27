package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
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
type QSGNode struct {
	*qtrt.CObject
}

func (this *QSGNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGNode) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSGNodeFromPointer(cthis unsafe.Pointer) *QSGNode {
	return &QSGNode{&qtrt.CObject{cthis}}
}
func (*QSGNode) NewFromPointer(cthis unsafe.Pointer) *QSGNode {
	return NewQSGNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:130
// index:0
// Public
// void QSGNode()
func NewQSGNode() *QSGNode {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNodeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:165
// index:1
// Protected
// void QSGNode(QSGNode::NodeType)
func NewQSGNode_1(type_ int) *QSGNode {
	cthis := qtrt.Calloc(1, 256) // 80
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNodeC2ENS_8NodeTypeE", ffiqt.FFI_TYPE_VOID, cthis, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:131
// index:0
// Public virtual
// void ~QSGNode()
func DeleteQSGNode(*QSGNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:133
// index:0
// Public inline
// QSGNode * parent()
func (this *QSGNode) Parent() *QSGNode /*777 QSGNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode6parentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:135
// index:0
// Public
// void removeChildNode(QSGNode *)
func (this *QSGNode) RemoveChildNode(node *QSGNode /*777 QSGNode **/) {
	var convArg0 = node.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode15removeChildNodeEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:136
// index:0
// Public
// void removeAllChildNodes()
func (this *QSGNode) RemoveAllChildNodes() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode19removeAllChildNodesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:137
// index:0
// Public
// void prependChildNode(QSGNode *)
func (this *QSGNode) PrependChildNode(node *QSGNode /*777 QSGNode **/) {
	var convArg0 = node.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode16prependChildNodeEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:138
// index:0
// Public
// void appendChildNode(QSGNode *)
func (this *QSGNode) AppendChildNode(node *QSGNode /*777 QSGNode **/) {
	var convArg0 = node.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode15appendChildNodeEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:139
// index:0
// Public
// void insertChildNodeBefore(QSGNode *, QSGNode *)
func (this *QSGNode) InsertChildNodeBefore(node *QSGNode /*777 QSGNode **/, before *QSGNode /*777 QSGNode **/) {
	var convArg0 = node.GetCthis()
	var convArg1 = before.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode21insertChildNodeBeforeEPS_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:140
// index:0
// Public
// void insertChildNodeAfter(QSGNode *, QSGNode *)
func (this *QSGNode) InsertChildNodeAfter(node *QSGNode /*777 QSGNode **/, after *QSGNode /*777 QSGNode **/) {
	var convArg0 = node.GetCthis()
	var convArg1 = after.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode20insertChildNodeAfterEPS_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:141
// index:0
// Public
// void reparentChildNodesTo(QSGNode *)
func (this *QSGNode) ReparentChildNodesTo(newParent *QSGNode /*777 QSGNode **/) {
	var convArg0 = newParent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode20reparentChildNodesToEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:143
// index:0
// Public
// int childCount()
func (this *QSGNode) ChildCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode10childCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQuick/qsgnode.h:144
// index:0
// Public
// QSGNode * childAtIndex(int)
func (this *QSGNode) ChildAtIndex(i int) *QSGNode /*777 QSGNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode12childAtIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:145
// index:0
// Public inline
// QSGNode * firstChild()
func (this *QSGNode) FirstChild() *QSGNode /*777 QSGNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode10firstChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:146
// index:0
// Public inline
// QSGNode * lastChild()
func (this *QSGNode) LastChild() *QSGNode /*777 QSGNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode9lastChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:147
// index:0
// Public inline
// QSGNode * nextSibling()
func (this *QSGNode) NextSibling() *QSGNode /*777 QSGNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode11nextSiblingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:148
// index:0
// Public inline
// QSGNode * previousSibling()
func (this *QSGNode) PreviousSibling() *QSGNode /*777 QSGNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode15previousSiblingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:150
// index:0
// Public inline
// QSGNode::NodeType type()
func (this *QSGNode) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:152
// index:0
// Public inline
// void clearDirty()
func (this *QSGNode) ClearDirty() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode10clearDirtyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:154
// index:0
// Public inline
// QSGNode::DirtyState dirtyState()
func (this *QSGNode) DirtyState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode10dirtyStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:156
// index:0
// Public virtual
// bool isSubtreeBlocked()
func (this *QSGNode) IsSubtreeBlocked() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode16isSubtreeBlockedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgnode.h:158
// index:0
// Public inline
// QSGNode::Flags flags()
func (this *QSGNode) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QSGNode5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:159
// index:0
// Public
// void setFlag(QSGNode::Flag, bool)
func (this *QSGNode) SetFlag(arg0 int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode7setFlagENS_4FlagEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:160
// index:0
// Public
// void setFlags(QSGNode::Flags, bool)
func (this *QSGNode) SetFlags(arg0 int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode8setFlagsE6QFlagsINS_4FlagEEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:162
// index:0
// Public inline virtual
// void preprocess()
func (this *QSGNode) Preprocess() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QSGNode10preprocessEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
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
