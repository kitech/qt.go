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

/*

 */
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

/*
Constructs a new node
*/
func (*QSGNode) NewForInherit() *QSGNode {
	return NewQSGNode()
}
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
// [-2] void QSGNode(QSGNode::NodeType)

/*
Constructs a new node
*/
func (*QSGNode) NewForInherit1(type_ int) *QSGNode {
	return NewQSGNode1(type_)
}
func NewQSGNode1(type_ int) *QSGNode {
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

/*

 */
func DeleteQSGNode(this *QSGNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 80)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * parent() const

/*
Returns the parent node of this node.
*/
func (this *QSGNode) Parent() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeChildNode(QSGNode *)

/*
Removes node from this node's list of children.
*/
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

/*
Removes all child nodes from this node's list of children.
*/
func (this *QSGNode) RemoveAllChildNodes() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode19removeAllChildNodesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prependChildNode(QSGNode *)

/*
Prepends node to this node's the list of children.

Ordering of nodes is important as geometry nodes will be rendered in the order they are added to the scene graph.
*/
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

/*
Appends node to this node's list of children.

Ordering of nodes is important as geometry nodes will be rendered in the order they are added to the scene graph.
*/
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

/*
Inserts node to this node's list of children before the node specified with before.

Ordering of nodes is important as geometry nodes will be rendered in the order they are added to the scene graph.
*/
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

/*
Inserts node to this node's list of children after the node specified with after.

Ordering of nodes is important as geometry nodes will be rendered in the order they are added to the scene graph.
*/
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

/*

 */
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
// [4] int childCount() const

/*
Returns the number of child nodes.
*/
func (this *QSGNode) ChildCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode10childCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgnode.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGNode * childAtIndex(int) const

/*
Returns the child at index i.

Children are stored internally as a linked list, so iterating over the children via the index is suboptimal.
*/
func (this *QSGNode) ChildAtIndex(i int) *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode12childAtIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:145
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * firstChild() const

/*
Returns the first child of this node.

The children are stored in a linked list.
*/
func (this *QSGNode) FirstChild() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode10firstChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * lastChild() const

/*
Returns the last child of this node.

The children are stored as a linked list.
*/
func (this *QSGNode) LastChild() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode9lastChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:147
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * nextSibling() const

/*
Returns the node after this in the parent's list of children.

The children are stored as a linked list.
*/
func (this *QSGNode) NextSibling() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode11nextSiblingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:148
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGNode * previousSibling() const

/*
Returns the node before this in the parent's list of children.

The children are stored as a linked list.
*/
func (this *QSGNode) PreviousSibling() *QSGNode /*777 QSGNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode15previousSiblingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:150
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGNode::NodeType type() const

/*
Returns the type of this node. The node type must be one of the predefined types defined in QSGNode::NodeType and can safely be used to cast to the corresponding class.
*/
func (this *QSGNode) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:152
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearDirty()

/*

 */
func (this *QSGNode) ClearDirty() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode10clearDirtyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void markDirty(QSGNode::DirtyState)

/*
Notifies all connected renderers that the node has dirty bits.
*/
func (this *QSGNode) MarkDirty(bits int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode9markDirtyE6QFlagsINS_13DirtyStateBitEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bits)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGNode::DirtyState dirtyState() const

/*

 */
func (this *QSGNode) DirtyState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode10dirtyStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:156
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isSubtreeBlocked() const

/*
Returns whether this node and its subtree is available for use.

Blocked subtrees will not get their dirty states updated and they will not be rendered.

The QSGOpacityNode will return a blocked subtree when accumulated opacity is 0, for instance.
*/
func (this *QSGNode) IsSubtreeBlocked() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode16isSubtreeBlockedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgnode.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGNode::Flags flags() const

/*
Returns the set of flags for this node.

See also setFlags().
*/
func (this *QSGNode) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QSGNode5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QSGNode::Flag, bool)

/*
Sets the flag f on this node if enabled is true; otherwise clears the flag.

See also flags().
*/
func (this *QSGNode) SetFlag(arg0 int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode7setFlagENS_4FlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QSGNode::Flag, bool)

/*
Sets the flag f on this node if enabled is true; otherwise clears the flag.

See also flags().
*/
func (this *QSGNode) SetFlagp(arg0 int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	arg1 := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode7setFlagENS_4FlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QSGNode::Flags, bool)

/*
Sets the flags f on this node if enabled is true; otherwise clears the flags.

See also flags().
*/
func (this *QSGNode) SetFlags(arg0 int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode8setFlagsE6QFlagsINS_4FlagEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QSGNode::Flags, bool)

/*
Sets the flags f on this node if enabled is true; otherwise clears the flags.

See also flags().
*/
func (this *QSGNode) SetFlagsp(arg0 int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	arg1 := true
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode8setFlagsE6QFlagsINS_4FlagEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:162
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void preprocess()

/*
Override this function to do processing on the node before it is rendered.

Preprocessing needs to be explicitly enabled by setting the flag QSGNode::UsePreprocess. The flag needs to be set before the node is added to the scene graph and will cause the preprocess() function to be called for every frame the node is rendered.

Warning: Beware of deleting nodes while they are being preprocessed. It is possible, with a small performance hit, to delete a single node during its own preprocess call. Deleting a subtree which has nodes that also use preprocessing may result in a segmentation fault. This is done for performance reasons.
*/
func (this *QSGNode) Preprocess() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QSGNode10preprocessEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
Can be used to figure out the type of node.



See also type().

*/
type QSGNode__NodeType = int

// The type of QSGNode
const QSGNode__BasicNodeType QSGNode__NodeType = 0

// The type of QSGGeometryNode
const QSGNode__GeometryNodeType QSGNode__NodeType = 1

// The type of QSGTransformNode
const QSGNode__TransformNodeType QSGNode__NodeType = 2

// The type of QSGClipNode
const QSGNode__ClipNodeType QSGNode__NodeType = 3

// The type of QSGOpacityNode
const QSGNode__OpacityNodeType QSGNode__NodeType = 4

//
const QSGNode__RootNodeType QSGNode__NodeType = 5

// The type of QSGRenderNode
const QSGNode__RenderNodeType QSGNode__NodeType = 6

func (this *QSGNode) NodeTypeItemName(val int) string {
	switch val {
	case QSGNode__BasicNodeType: // 0
		return "BasicNodeType"
	case QSGNode__GeometryNodeType: // 1
		return "GeometryNodeType"
	case QSGNode__TransformNodeType: // 2
		return "TransformNodeType"
	case QSGNode__ClipNodeType: // 3
		return "ClipNodeType"
	case QSGNode__OpacityNodeType: // 4
		return "OpacityNodeType"
	case QSGNode__RootNodeType: // 5
		return "RootNodeType"
	case QSGNode__RenderNodeType: // 6
		return "RenderNodeType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGNode_NodeTypeItemName(val int) string {
	var nilthis *QSGNode
	return nilthis.NodeTypeItemName(val)
}

/*


 */
type QSGNode__Flag = int

//
const QSGNode__OwnedByParent QSGNode__Flag = 1

//
const QSGNode__UsePreprocess QSGNode__Flag = 2

//
const QSGNode__OwnsGeometry QSGNode__Flag = 65536

//
const QSGNode__OwnsMaterial QSGNode__Flag = 131072

//
const QSGNode__OwnsOpaqueMaterial QSGNode__Flag = 262144

//
const QSGNode__IsVisitableNode QSGNode__Flag = 16777216

func (this *QSGNode) FlagItemName(val int) string {
	switch val {
	case QSGNode__OwnedByParent: // 1
		return "OwnedByParent"
	case QSGNode__UsePreprocess: // 2
		return "UsePreprocess"
	case QSGNode__OwnsGeometry: // 65536
		return "OwnsGeometry"
	case QSGNode__OwnsMaterial: // 131072
		return "OwnsMaterial"
	case QSGNode__OwnsOpaqueMaterial: // 262144
		return "OwnsOpaqueMaterial"
	case QSGNode__IsVisitableNode: // 16777216
		return "IsVisitableNode"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGNode_FlagItemName(val int) string {
	var nilthis *QSGNode
	return nilthis.FlagItemName(val)
}

/*


 */
type QSGNode__DirtyStateBit = int

//
const QSGNode__DirtySubtreeBlocked QSGNode__DirtyStateBit = 128

//
const QSGNode__DirtyMatrix QSGNode__DirtyStateBit = 256

//
const QSGNode__DirtyNodeAdded QSGNode__DirtyStateBit = 1024

//
const QSGNode__DirtyNodeRemoved QSGNode__DirtyStateBit = 2048

//
const QSGNode__DirtyGeometry QSGNode__DirtyStateBit = 4096

//
const QSGNode__DirtyMaterial QSGNode__DirtyStateBit = 8192

//
const QSGNode__DirtyOpacity QSGNode__DirtyStateBit = 16384

//
const QSGNode__DirtyForceUpdate QSGNode__DirtyStateBit = 32768

//
const QSGNode__DirtyUsePreprocess QSGNode__DirtyStateBit = 2

//
const QSGNode__DirtyPropagationMask QSGNode__DirtyStateBit = 50432

func (this *QSGNode) DirtyStateBitItemName(val int) string {
	switch val {
	case QSGNode__DirtySubtreeBlocked: // 128
		return "DirtySubtreeBlocked"
	case QSGNode__DirtyMatrix: // 256
		return "DirtyMatrix"
	case QSGNode__DirtyNodeAdded: // 1024
		return "DirtyNodeAdded"
	case QSGNode__DirtyNodeRemoved: // 2048
		return "DirtyNodeRemoved"
	case QSGNode__DirtyGeometry: // 4096
		return "DirtyGeometry"
	case QSGNode__DirtyMaterial: // 8192
		return "DirtyMaterial"
	case QSGNode__DirtyOpacity: // 16384
		return "DirtyOpacity"
	case QSGNode__DirtyForceUpdate: // 32768
		return "DirtyForceUpdate"
	case QSGNode__DirtyUsePreprocess: // 2
		return "DirtyUsePreprocess"
	case QSGNode__DirtyPropagationMask: // 50432
		return "DirtyPropagationMask"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGNode_DirtyStateBitItemName(val int) string {
	var nilthis *QSGNode
	return nilthis.DirtyStateBitItemName(val)
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
