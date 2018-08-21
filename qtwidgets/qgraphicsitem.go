package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
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

// void updateMicroFocus()
func (this *QGraphicsItem) InheritUpdateMicroFocus(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateMicroFocus", f)
}

// bool sceneEventFilter(QGraphicsItem *, QEvent *)
func (this *QGraphicsItem) InheritSceneEventFilter(f func(watched *QGraphicsItem /*777 QGraphicsItem **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEventFilter", f)
}

// bool sceneEvent(QEvent *)
func (this *QGraphicsItem) InheritSceneEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEvent", f)
}

// void contextMenuEvent(QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsItem) InheritContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDropEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QGraphicsItem) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QGraphicsItem) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void hoverEnterEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) InheritHoverEnterEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverEnterEvent", f)
}

// void hoverMoveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) InheritHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) InheritHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QGraphicsItem) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsItem) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMousePressEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void wheelEvent(QGraphicsSceneWheelEvent *)
func (this *QGraphicsItem) InheritWheelEvent(f func(event *QGraphicsSceneWheelEvent /*777 QGraphicsSceneWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsItem) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsItem) InheritInputMethodQuery(f func(query int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "inputMethodQuery", f)
}

// QVariant itemChange(QGraphicsItem::GraphicsItemChange, const QVariant &)
func (this *QGraphicsItem) InheritItemChange(f func(change int, value *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "itemChange", f)
}

// bool supportsExtension(QGraphicsItem::Extension)
func (this *QGraphicsItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const QVariant &)
func (this *QGraphicsItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

// void addToIndex()
func (this *QGraphicsItem) InheritAddToIndex(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "addToIndex", f)
}

// void removeFromIndex()
func (this *QGraphicsItem) InheritRemoveFromIndex(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "removeFromIndex", f)
}

// void prepareGeometryChange()
func (this *QGraphicsItem) InheritPrepareGeometryChange(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "prepareGeometryChange", f)
}

/*

 */
type QGraphicsItem struct {
	*qtrt.CObject
}
type QGraphicsItem_ITF interface {
	QGraphicsItem_PTR() *QGraphicsItem
}

func (ptr *QGraphicsItem) QGraphicsItem_PTR() *QGraphicsItem { return ptr }

func (this *QGraphicsItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGraphicsItemFromPointer(cthis unsafe.Pointer) *QGraphicsItem {
	return &QGraphicsItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsItem {
	return NewQGraphicsItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsItem(QGraphicsItem *)

/*
Constructs a QGraphicsItem with the given parent item. It does not modify the parent object returned by QObject::parent().

If parent is 0, you can add the item to a scene by calling QGraphicsScene::addItem(). The item will then become a top-level item.

See also QGraphicsScene::addItem() and setParentItem().
*/
func NewQGraphicsItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItemC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsItem(QGraphicsItem *)

/*
Constructs a QGraphicsItem with the given parent item. It does not modify the parent object returned by QObject::parent().

If parent is 0, you can add the item to a scene by calling QGraphicsScene::addItem(). The item will then become a top-level item.

See also QGraphicsScene::addItem() and setParentItem().
*/
func NewQGraphicsItem__() *QGraphicsItem {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItemC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:162
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsItem()

/*

 */
func DeleteQGraphicsItem(this *QGraphicsItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsScene * scene() const

/*
Returns the current scene for the item, or 0 if the item is not stored in a scene.

To add or move an item to a scene, call QGraphicsScene::addItem().
*/
func (this *QGraphicsItem) Scene() *QGraphicsScene /*777 QGraphicsScene **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5sceneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * parentItem() const

/*
Returns a pointer to this item's parent item. If this item does not have a parent, 0 is returned.

See also setParentItem() and childItems().
*/
func (this *QGraphicsItem) ParentItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10parentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * topLevelItem() const

/*
Returns this item's top-level item. The top-level item is the item's topmost ancestor item whose parent is 0. If an item has no parent, its own pointer is returned (i.e., a top-level item is its own top-level item).

See also parentItem().
*/
func (this *QGraphicsItem) TopLevelItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12topLevelItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * parentWidget() const

/*
Returns a pointer to the item's parent widget. The item's parent widget is the closest parent item that is a widget.

This function was introduced in  Qt 4.4.

See also parentItem() and childItems().
*/
func (this *QGraphicsItem) ParentWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12parentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * topLevelWidget() const

/*
Returns a pointer to the item's top level widget (i.e., the item's ancestor whose parent is 0, or whose parent is not a widget), or 0 if this item does not have a top level widget. If the item is its own top level widget, this function returns a pointer to the item itself.

This function was introduced in  Qt 4.4.
*/
func (this *QGraphicsItem) TopLevelWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14topLevelWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * window() const

/*
Returns the item's window, or 0 if this item does not have a window. If the item is a window, it will return itself. Otherwise it will return the closest ancestor that is a window.

This function was introduced in  Qt 4.4.

See also QGraphicsWidget::isWindow().
*/
func (this *QGraphicsItem) Window() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * panel() const

/*
Returns the item's panel, or 0 if this item does not have a panel. If the item is a panel, it will return itself. Otherwise it will return the closest ancestor that is a panel.

This function was introduced in  Qt 4.6.

See also isPanel() and ItemIsPanel.
*/
func (this *QGraphicsItem) Panel() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5panelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParentItem(QGraphicsItem *)

/*
Sets this item's parent item to newParent. If this item already has a parent, it is first removed from the previous parent. If newParent is 0, this item will become a top-level item.

Note that this implicitly adds this graphics item to the scene of the parent. You should not add the item to the scene yourself.

The behavior when calling this function on an item that is an ancestor of newParent is undefined.

See also parentItem() and childItems().
*/
func (this *QGraphicsItem) SetParentItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13setParentItemEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:177
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> childItems() const

/*
Returns a list of this item's children.

The items are sorted by stacking order. This takes into account both the items' insertion order and their Z-values.

This function was introduced in  Qt 4.4.

See also setParentItem(), zValue(), and Sorting.
*/
func (this *QGraphicsItem) ChildItems() *QGraphicsItemList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10childItemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:178
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWidget() const

/*
Returns true if this item is a widget (i.e., QGraphicsWidget); otherwise, returns false.

This function was introduced in  Qt 4.4.
*/
func (this *QGraphicsItem) IsWidget() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:179
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWindow() const

/*
Returns true if the item is a QGraphicsWidget window, otherwise returns false.

This function was introduced in  Qt 4.4.

See also QGraphicsWidget::windowFlags().
*/
func (this *QGraphicsItem) IsWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:180
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPanel() const

/*
Returns true if the item is a panel; otherwise returns false.

This function was introduced in  Qt 4.6.

See also QGraphicsItem::panel() and ItemIsPanel.
*/
func (this *QGraphicsItem) IsPanel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem7isPanelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItemGroup * group() const

/*
Returns a pointer to this item's item group, or 0 if this item is not member of a group.

See also setGroup(), QGraphicsItemGroup, and QGraphicsScene::createItemGroup().
*/
func (this *QGraphicsItem) Group() *QGraphicsItemGroup /*777 QGraphicsItemGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroup(QGraphicsItemGroup *)

/*
Adds this item to the item group group. If group is 0, this item is removed from any current group and added as a child of the previous group's parent.

See also group() and QGraphicsScene::createItemGroup().
*/
func (this *QGraphicsItem) SetGroup(group QGraphicsItemGroup_ITF /*777 QGraphicsItemGroup **/) {
	var convArg0 unsafe.Pointer
	if group != nil && group.QGraphicsItemGroup_PTR() != nil {
		convArg0 = group.QGraphicsItemGroup_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:188
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsItem::GraphicsItemFlags flags() const

/*
Returns this item's flags. The flags describe what configurable features of the item are enabled and not. For example, if the flags include ItemIsFocusable, the item can accept input focus.

By default, no flags are enabled.

See also setFlags() and setFlag().
*/
func (this *QGraphicsItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QGraphicsItem::GraphicsItemFlag, bool)

/*
If enabled is true, the item flag flag is enabled; otherwise, it is disabled.

See also flags() and setFlags().
*/
func (this *QGraphicsItem) SetFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem7setFlagENS_16GraphicsItemFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QGraphicsItem::GraphicsItemFlag, bool)

/*
If enabled is true, the item flag flag is enabled; otherwise, it is disabled.

See also flags() and setFlags().
*/
func (this *QGraphicsItem) SetFlag__(flag int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem7setFlagENS_16GraphicsItemFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QGraphicsItem::GraphicsItemFlags)

/*
Sets the item flags to flags. All flags in flags are enabled; all flags not in flags are disabled.

If the item had focus and flags does not enable ItemIsFocusable, the item loses focus as a result of calling this function. Similarly, if the item was selected, and flags does not enabled ItemIsSelectable, the item is automatically unselected.

By default, no flags are enabled. (QGraphicsWidget enables the ItemSendsGeometryChanges flag by default in order to track position changes.)

See also flags() and setFlag().
*/
func (this *QGraphicsItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setFlagsE6QFlagsINS_16GraphicsItemFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:192
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsItem::CacheMode cacheMode() const

/*
Returns the cache mode for this item. The default mode is NoCache (i.e., cache is disabled and all painting is immediate).

This function was introduced in  Qt 4.4.

See also setCacheMode().
*/
func (this *QGraphicsItem) CacheMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9cacheModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheMode(QGraphicsItem::CacheMode, const QSize &)

/*
Sets the item's cache mode to mode.

The optional logicalCacheSize argument is used only by ItemCoordinateCache mode, and describes the resolution of the cache buffer; if logicalCacheSize is (100, 100), QGraphicsItem will fit the item into 100x100 pixels in graphics memory, regardless of the logical size of the item itself. By default QGraphicsItem uses the size of boundingRect(). For all other cache modes than ItemCoordinateCache, logicalCacheSize is ignored.

Caching can speed up rendering if your item spends a significant time redrawing itself. In some cases the cache can also slow down rendering, in particular when the item spends less time redrawing than QGraphicsItem spends redrawing from the cache.

When caching is enabled, an item's paint() function will generally draw into an offscreen pixmap cache; for any subsequent repaint requests, the Graphics View framework will redraw from the cache. This approach works particularly well with QGLWidget, which stores all the cache as OpenGL textures.

Be aware that QPixmapCache's cache limit may need to be changed to obtain optimal performance.

You can read more about the different cache modes in the CacheMode documentation.

Note: Enabling caching does not imply that the item's paint() function will be called only in response to an explicit update() call. For instance, under memory pressure, Qt may decide to drop some of the cache information; in such cases an item's paint() function will be called even if there was no update() call (that is, exactly as if there were no caching enabled).

This function was introduced in  Qt 4.4.

See also cacheMode(), CacheMode, and QPixmapCache::setCacheLimit().
*/
func (this *QGraphicsItem) SetCacheMode(mode int, cacheSize qtcore.QSize_ITF) {
	var convArg1 unsafe.Pointer
	if cacheSize != nil && cacheSize.QSize_PTR() != nil {
		convArg1 = cacheSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12setCacheModeENS_9CacheModeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheMode(QGraphicsItem::CacheMode, const QSize &)

/*
Sets the item's cache mode to mode.

The optional logicalCacheSize argument is used only by ItemCoordinateCache mode, and describes the resolution of the cache buffer; if logicalCacheSize is (100, 100), QGraphicsItem will fit the item into 100x100 pixels in graphics memory, regardless of the logical size of the item itself. By default QGraphicsItem uses the size of boundingRect(). For all other cache modes than ItemCoordinateCache, logicalCacheSize is ignored.

Caching can speed up rendering if your item spends a significant time redrawing itself. In some cases the cache can also slow down rendering, in particular when the item spends less time redrawing than QGraphicsItem spends redrawing from the cache.

When caching is enabled, an item's paint() function will generally draw into an offscreen pixmap cache; for any subsequent repaint requests, the Graphics View framework will redraw from the cache. This approach works particularly well with QGLWidget, which stores all the cache as OpenGL textures.

Be aware that QPixmapCache's cache limit may need to be changed to obtain optimal performance.

You can read more about the different cache modes in the CacheMode documentation.

Note: Enabling caching does not imply that the item's paint() function will be called only in response to an explicit update() call. For instance, under memory pressure, Qt may decide to drop some of the cache information; in such cases an item's paint() function will be called even if there was no update() call (that is, exactly as if there were no caching enabled).

This function was introduced in  Qt 4.4.

See also cacheMode(), CacheMode, and QPixmapCache::setCacheLimit().
*/
func (this *QGraphicsItem) SetCacheMode__(mode int) {
	// arg: 1, const QSize &=LValueReference, QSize=Record, , Invalid
	var convArg1 = qtcore.NewQSize()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12setCacheModeENS_9CacheModeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:195
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsItem::PanelModality panelModality() const

/*
Returns the modality for this item.

This function was introduced in  Qt 4.6.

See also setPanelModality().
*/
func (this *QGraphicsItem) PanelModality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13panelModalityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPanelModality(QGraphicsItem::PanelModality)

/*
Sets the modality for this item to panelModality.

Changing the modality of a visible item takes effect immediately.

This function was introduced in  Qt 4.6.

See also panelModality().
*/
func (this *QGraphicsItem) SetPanelModality(panelModality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem16setPanelModalityENS_13PanelModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), panelModality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:197
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBlockedByModalPanel(QGraphicsItem **) const

/*
Returns true if this item is blocked by a modal panel, false otherwise. If blockingPanel is non-zero, blockingPanel will be set to the modal panel that is blocking this item. If this item is not blocked, blockingPanel will not be set by this function.

This function always returns false for items not in a scene.

This function was introduced in  Qt 4.6.

See also panelModality(), setPanelModality(), and PanelModality.
*/
func (this *QGraphicsItem) IsBlockedByModalPanel(blockingPanel QGraphicsItem_ITF /*777 QGraphicsItem ***/) bool {
	var convArg0 unsafe.Pointer
	if blockingPanel != nil && blockingPanel.QGraphicsItem_PTR() != nil {
		convArg0 = blockingPanel.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:197
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBlockedByModalPanel(QGraphicsItem **) const

/*
Returns true if this item is blocked by a modal panel, false otherwise. If blockingPanel is non-zero, blockingPanel will be set to the modal panel that is blocking this item. If this item is not blocked, blockingPanel will not be set by this function.

This function always returns false for items not in a scene.

This function was introduced in  Qt 4.6.

See also panelModality(), setPanelModality(), and PanelModality.
*/
func (this *QGraphicsItem) IsBlockedByModalPanel__() bool {
	// arg: 0, QGraphicsItem **=Pointer, QGraphicsItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip() const

/*
Returns the item's tool tip, or an empty QString if no tool tip has been set.

See also setToolTip() and QToolTip.
*/
func (this *QGraphicsItem) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*
Sets the item's tool tip to toolTip. If toolTip is empty, the item's tool tip is cleared.

See also toolTip() and QToolTip.
*/
func (this *QGraphicsItem) SetToolTip(toolTip string) {
	var tmpArg0 = qtcore.NewQString_5(toolTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:205
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor() const

/*
Returns the current cursor shape for the item. The mouse cursor will assume this shape when it's over this item. See the list of predefined cursor objects for a range of useful shapes.

An editor item might want to use an I-beam cursor:


  item->setCursor(Qt::IBeamCursor);



If no cursor has been set, the cursor of the item beneath is used.

See also setCursor(), hasCursor(), unsetCursor(), QWidget::cursor, and QApplication::overrideCursor().
*/
func (this *QGraphicsItem) Cursor() *qtgui.QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)

/*
Sets the current cursor shape for the item to cursor. The mouse cursor will assume this shape when it's over this item. See the list of predefined cursor objects for a range of useful shapes.

An editor item might want to use an I-beam cursor:


  item->setCursor(Qt::IBeamCursor);



If no cursor has been set, the cursor of the item beneath is used.

See also cursor(), hasCursor(), unsetCursor(), QWidget::cursor, and QApplication::overrideCursor().
*/
func (this *QGraphicsItem) SetCursor(cursor qtgui.QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QCursor_PTR() != nil {
		convArg0 = cursor.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:207
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasCursor() const

/*
Returns true if this item has a cursor set; otherwise, false is returned.

By default, items don't have any cursor set. cursor() will return a standard pointing arrow cursor.

See also unsetCursor().
*/
func (this *QGraphicsItem) HasCursor() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9hasCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()

/*
Clears the cursor from this item.

See also hasCursor() and setCursor().
*/
func (this *QGraphicsItem) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*
Returns true if the item is visible; otherwise, false is returned.

Note that the item's general visibility is unrelated to whether or not it is actually being visualized by a QGraphicsView.

See also setVisible().
*/
func (this *QGraphicsItem) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:212
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisibleTo(const QGraphicsItem *) const

/*
Returns true if the item is visible to parent; otherwise, false is returned. parent can be 0, in which case this function will return whether the item is visible to the scene or not.

An item may not be visible to its ancestors even if isVisible() is true. It may also be visible to its ancestors even if isVisible() is false. If any ancestor is hidden, the item itself will be implicitly hidden, in which case this function will return false.

This function was introduced in  Qt 4.4.

See also isVisible() and setVisible().
*/
func (this *QGraphicsItem) IsVisibleTo(parent QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11isVisibleToEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:213
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*
If visible is true, the item is made visible. Otherwise, the item is made invisible. Invisible items are not painted, nor do they receive any events. In particular, mouse events pass right through invisible items, and are delivered to any item that may be behind. Invisible items are also unselectable, they cannot take input focus, and are not detected by QGraphicsScene's item location functions.

If an item becomes invisible while grabbing the mouse, (i.e., while it is receiving mouse events,) it will automatically lose the mouse grab, and the grab is not regained by making the item visible again; it must receive a new mouse press to regain the mouse grab.

Similarly, an invisible item cannot have focus, so if the item has focus when it becomes invisible, it will lose focus, and the focus is not regained by simply making the item visible again.

If you hide a parent item, all its children will also be hidden. If you show a parent item, all children will be shown, unless they have been explicitly hidden (i.e., if you call setVisible(false) on a child, it will not be reshown even if its parent is hidden, and then shown again).

Items are visible by default; it is unnecessary to call setVisible() on a new item.

Note: An item with opacity set to 0 will still be considered visible, although it will be treated like an invisible item: mouse events will pass through it, it will not be included in the items returned by QGraphicsView::items(), and so on. However, the item will retain the focus.

See also isVisible(), show(), hide(), and setOpacity().
*/
func (this *QGraphicsItem) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hide()

/*
Hides the item (items are visible by default).

This convenience function is equivalent to calling setVisible(false).

See also show() and setVisible().
*/
func (this *QGraphicsItem) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void show()

/*
Shows the item (items are visible by default).

This convenience function is equivalent to calling setVisible(true).

See also hide() and setVisible().
*/
func (this *QGraphicsItem) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:217
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*
Returns true if the item is enabled; otherwise, false is returned.

See also setEnabled().
*/
func (this *QGraphicsItem) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*
If enabled is true, the item is enabled; otherwise, it is disabled.

Disabled items are visible, but they do not receive any events, and cannot take focus nor be selected. Mouse events are discarded; they are not propagated unless the item is also invisible, or if it does not accept mouse events (see acceptedMouseButtons()). A disabled item cannot become the mouse grabber, and as a result of this, an item loses the grab if it becomes disabled when grabbing the mouse, just like it loses focus if it had focus when it was disabled.

Disabled items are traditionally drawn using grayed-out colors (see QPalette::Disabled).

If you disable a parent item, all its children will also be disabled. If you enable a parent item, all children will be enabled, unless they have been explicitly disabled (i.e., if you call setEnabled(false) on a child, it will not be reenabled if its parent is disabled, and then enabled again).

Items are enabled by default.

Note: If you install an event filter, you can still intercept events before they are delivered to items; this mechanism disregards the item's enabled state.

See also isEnabled().
*/
func (this *QGraphicsItem) SetEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:220
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSelected() const

/*
Returns true if this item is selected; otherwise, false is returned.

Items that are in a group inherit the group's selected state.

Items are not selected by default.

See also setSelected() and QGraphicsScene::setSelectionArea().
*/
func (this *QGraphicsItem) IsSelected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10isSelectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelected(bool)

/*
If selected is true and this item is selectable, this item is selected; otherwise, it is unselected.

If the item is in a group, the whole group's selected state is toggled by this function. If the group is selected, all items in the group are also selected, and if the group is not selected, no item in the group is selected.

Only visible, enabled, selectable items can be selected. If selected is true and this item is either invisible or disabled or unselectable, this function does nothing.

By default, items cannot be selected. To enable selection, set the ItemIsSelectable flag.

This function is provided for convenience, allowing individual toggling of the selected state of an item. However, a more common way of selecting items is to call QGraphicsScene::setSelectionArea(), which will call this function for all visible, enabled, and selectable items within a specified area on the scene.

See also isSelected() and QGraphicsScene::selectedItems().
*/
func (this *QGraphicsItem) SetSelected(selected bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11setSelectedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selected)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:223
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptDrops() const

/*
Returns true if this item can accept drag and drop events; otherwise, returns false. By default, items do not accept drag and drop events; items are transparent to drag and drop.

See also setAcceptDrops().
*/
func (this *QGraphicsItem) AcceptDrops() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11acceptDropsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptDrops(bool)

/*
If on is true, this item will accept drag and drop events; otherwise, it is transparent for drag and drop events. By default, items do not accept drag and drop events.

See also acceptDrops().
*/
func (this *QGraphicsItem) SetAcceptDrops(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14setAcceptDropsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:226
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity() const

/*
Returns this item's local opacity, which is between 0.0 (transparent) and 1.0 (opaque). This value is combined with parent and ancestor values into the effectiveOpacity(). The effective opacity decides how the item is rendered and also affects its visibility when queried by functions such as QGraphicsView::items().

The opacity property decides the state of the painter passed to the paint() function. If the item is cached, i.e., ItemCoordinateCache or DeviceCoordinateCache, the effective property will be applied to the item's cache as it is rendered.

The default opacity is 1.0; fully opaque.

This function was introduced in  Qt 4.5.

See also setOpacity(), paint(), ItemIgnoresParentOpacity, and ItemDoesntPropagateOpacityToChildren.
*/
func (this *QGraphicsItem) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:227
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal effectiveOpacity() const

/*
Returns this item's effective opacity, which is between 0.0 (transparent) and 1.0 (opaque). This value is a combination of this item's local opacity, and its parent and ancestors' opacities. The effective opacity decides how the item is rendered.

This function was introduced in  Qt 4.5.

See also opacity(), setOpacity(), paint(), ItemIgnoresParentOpacity, and ItemDoesntPropagateOpacityToChildren.
*/
func (this *QGraphicsItem) EffectiveOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16effectiveOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)

/*
Sets this item's local opacity, between 0.0 (transparent) and 1.0 (opaque). The item's local opacity is combined with parent and ancestor opacities into the effectiveOpacity().

By default, opacity propagates from parent to child, so if a parent's opacity is 0.5 and the child is also 0.5, the child's effective opacity will be 0.25.

The opacity property decides the state of the painter passed to the paint() function. If the item is cached, i.e., ItemCoordinateCache or DeviceCoordinateCache, the effective property will be applied to the item's cache as it is rendered.

There are two item flags that affect how the item's opacity is combined with the parent: ItemIgnoresParentOpacity and ItemDoesntPropagateOpacityToChildren.

Note: Setting the opacity of an item to 0 will not make the item invisible (according to isVisible()), but the item will be treated like an invisible one. See the documentation of setVisible() for more information.

This function was introduced in  Qt 4.5.

See also opacity(), effectiveOpacity(), and setVisible().
*/
func (this *QGraphicsItem) SetOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEffect * graphicsEffect() const

/*
Returns a pointer to this item's effect if it has one; otherwise 0.

This function was introduced in  Qt 4.6.

See also setGraphicsEffect().
*/
func (this *QGraphicsItem) GraphicsEffect() *QGraphicsEffect /*777 QGraphicsEffect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14graphicsEffectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGraphicsEffect(QGraphicsEffect *)

/*
Sets effect as the item's effect. If there already is an effect installed on this item, QGraphicsItem will delete the existing effect before installing the new effect. You can delete an existing effect by calling setGraphicsEffect(0).

If effect is the installed effect on a different item, setGraphicsEffect() will remove the effect from the item and install it on this item.

QGraphicsItem takes ownership of effect.

Note: This function will apply the effect on itself and all its children.

This function was introduced in  Qt 4.6.

See also graphicsEffect().
*/
func (this *QGraphicsItem) SetGraphicsEffect(effect QGraphicsEffect_ITF /*777 QGraphicsEffect **/) {
	var convArg0 unsafe.Pointer
	if effect != nil && effect.QGraphicsEffect_PTR() != nil {
		convArg0 = effect.QGraphicsEffect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:236
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons acceptedMouseButtons() const

/*
Returns the mouse buttons that this item accepts mouse events for. By default, all mouse buttons are accepted.

If an item accepts a mouse button, it will become the mouse grabber item when a mouse press event is delivered for that mouse button. However, if the item does not accept the button, QGraphicsScene will forward the mouse events to the first item beneath it that does.

See also setAcceptedMouseButtons() and mousePressEvent().
*/
func (this *QGraphicsItem) AcceptedMouseButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem20acceptedMouseButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptedMouseButtons(Qt::MouseButtons)

/*
Sets the mouse buttons that this item accepts mouse events for.

By default, all mouse buttons are accepted. If an item accepts a mouse button, it will become the mouse grabber item when a mouse press event is delivered for that button. However, if the item does not accept the mouse button, QGraphicsScene will forward the mouse events to the first item beneath it that does.

To disable mouse events for an item (i.e., make it transparent for mouse events), call setAcceptedMouseButtons(0).

See also acceptedMouseButtons() and mousePressEvent().
*/
func (this *QGraphicsItem) SetAcceptedMouseButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem23setAcceptedMouseButtonsE6QFlagsIN2Qt11MouseButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:242
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptHoverEvents() const

/*
Returns true if an item accepts hover events (QGraphicsSceneHoverEvent); otherwise, returns false. By default, items do not accept hover events.

This function was introduced in  Qt 4.4.

See also setAcceptHoverEvents() and setAcceptedMouseButtons().
*/
func (this *QGraphicsItem) AcceptHoverEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptHoverEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptHoverEvents(bool)

/*
If enabled is true, this item will accept hover events; otherwise, it will ignore them. By default, items do not accept hover events.

Hover events are delivered when there is no current mouse grabber item. They are sent when the mouse cursor enters an item, when it moves around inside the item, and when the cursor leaves an item. Hover events are commonly used to highlight an item when it's entered, and for tracking the mouse cursor as it hovers over the item (equivalent to QWidget::mouseTracking).

Parent items receive hover enter events before their children, and leave events after their children. The parent does not receive a hover leave event if the cursor enters a child, though; the parent stays "hovered" until the cursor leaves its area, including its children's areas.

If a parent item handles child events, it will receive hover move, drag move, and drop events as the cursor passes through its children, but it does not receive hover enter and hover leave, nor drag enter and drag leave events on behalf of its children.

A QGraphicsWidget with window decorations will accept hover events regardless of the value of acceptHoverEvents().

This function was introduced in  Qt 4.4.

See also acceptHoverEvents(), hoverEnterEvent(), hoverMoveEvent(), and hoverLeaveEvent().
*/
func (this *QGraphicsItem) SetAcceptHoverEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptHoverEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:244
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptTouchEvents() const

/*
Returns true if an item accepts touch events; otherwise, returns false. By default, items do not accept touch events.

This function was introduced in  Qt 4.6.

See also setAcceptTouchEvents().
*/
func (this *QGraphicsItem) AcceptTouchEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptTouchEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptTouchEvents(bool)

/*
If enabled is true, this item will accept touch events; otherwise, it will ignore them. By default, items do not accept touch events.

This function was introduced in  Qt 4.6.

See also acceptTouchEvents().
*/
func (this *QGraphicsItem) SetAcceptTouchEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptTouchEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:247
// index:0
// Public Visibility=Default Availability=Available
// [1] bool filtersChildEvents() const

/*
Returns true if this item filters child events (i.e., all events intended for any of its children are instead sent to this item); otherwise, false is returned.

The default value is false; child events are not filtered.

This function was introduced in  Qt 4.6.

See also setFiltersChildEvents().
*/
func (this *QGraphicsItem) FiltersChildEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem18filtersChildEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltersChildEvents(bool)

/*
If enabled is true, this item is set to filter all events for all its children (i.e., all events intented for any of its children are instead sent to this item); otherwise, if enabled is false, this item will only handle its own events. The default value is false.

This function was introduced in  Qt 4.6.

See also filtersChildEvents().
*/
func (this *QGraphicsItem) SetFiltersChildEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem21setFiltersChildEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:250
// index:0
// Public Visibility=Default Availability=Available
// [1] bool handlesChildEvents() const

/*

 */
func (this *QGraphicsItem) HandlesChildEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem18handlesChildEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHandlesChildEvents(bool)

/*

 */
func (this *QGraphicsItem) SetHandlesChildEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem21setHandlesChildEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:253
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive() const

/*
Returns true if this item is active; otherwise returns false.

An item can only be active if the scene is active. An item is active if it is, or is a descendent of, an active panel. Items in non-active panels are not active.

Items that are not part of a panel follow scene activation when the scene has no active panel.

Only active items can gain input focus.

This function was introduced in  Qt 4.6.

See also QGraphicsScene::isActive(), QGraphicsScene::activePanel(), panel(), and isPanel().
*/
func (this *QGraphicsItem) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:254
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActive(bool)

/*
If active is true, and the scene is active, this item's panel will be activated. Otherwise, the panel is deactivated.

If the item is not part of an active scene, active will decide what happens to the panel when the scene becomes active or the item is added to the scene. If true, the item's panel will be activated when the item is either added to the scene or the scene is activated. Otherwise, the item will stay inactive independent of the scene's activated state.

This function was introduced in  Qt 4.6.

See also isPanel(), QGraphicsScene::setActivePanel(), and QGraphicsScene::isActive().
*/
func (this *QGraphicsItem) SetActive(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setActiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), active)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:256
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus() const

/*
Returns true if this item is active, and it or its focus proxy has keyboard input focus; otherwise, returns false.

See also focusItem(), setFocus(), QGraphicsScene::setFocusItem(), and isActive().
*/
func (this *QGraphicsItem) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)

/*
Gives keyboard input focus to this item. The focusReason argument will be passed into any focus event generated by this function; it is used to give an explanation of what caused the item to get focus.

Only enabled items that set the ItemIsFocusable flag can accept keyboard focus.

If this item is not visible, not active, or not associated with a scene, it will not gain immediate input focus. However, it will be registered as the preferred focus item for its subtree of items, should it later become visible.

As a result of calling this function, this item will receive a focus in event with focusReason. If another item already has focus, that item will first receive a focus out event indicating that it has lost input focus.

See also clearFocus(), hasFocus(), focusItem(), and focusProxy().
*/
func (this *QGraphicsItem) SetFocus(focusReason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focusReason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)

/*
Gives keyboard input focus to this item. The focusReason argument will be passed into any focus event generated by this function; it is used to give an explanation of what caused the item to get focus.

Only enabled items that set the ItemIsFocusable flag can accept keyboard focus.

If this item is not visible, not active, or not associated with a scene, it will not gain immediate input focus. However, it will be registered as the preferred focus item for its subtree of items, should it later become visible.

As a result of calling this function, this item will receive a focus in event with focusReason. If another item already has focus, that item will first receive a focus out event indicating that it has lost input focus.

See also clearFocus(), hasFocus(), focusItem(), and focusProxy().
*/
func (this *QGraphicsItem) SetFocus__() {
	// arg: 0, Qt::FocusReason=Elaborated, Qt::FocusReason=Enum, , Invalid
	focusReason := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focusReason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFocus()

/*
Takes keyboard input focus from the item.

If it has focus, a focus out event is sent to this item to tell it that it is about to lose the focus.

Only items that set the ItemIsFocusable flag, or widgets that set an appropriate focus policy, can accept keyboard focus.

See also setFocus(), hasFocus(), and QGraphicsWidget::focusPolicy.
*/
func (this *QGraphicsItem) ClearFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10clearFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:260
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * focusProxy() const

/*
Returns this item's focus proxy, or 0 if this item has no focus proxy.

This function was introduced in  Qt 4.6.

See also setFocusProxy(), setFocus(), and hasFocus().
*/
func (this *QGraphicsItem) FocusProxy() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10focusProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusProxy(QGraphicsItem *)

/*
Sets the item's focus proxy to item.

If an item has a focus proxy, the focus proxy will receive input focus when the item gains input focus. The item itself will still have focus (i.e., hasFocus() will return true), but only the focus proxy will receive the keyboard input.

A focus proxy can itself have a focus proxy, and so on. In such case, keyboard input will be handled by the outermost focus proxy.

The focus proxy item must belong to the same scene as this item.

This function was introduced in  Qt 4.6.

See also focusProxy(), setFocus(), and hasFocus().
*/
func (this *QGraphicsItem) SetFocusProxy(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13setFocusProxyEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * focusItem() const

/*
If this item, a child or descendant of this item currently has input focus, this function will return a pointer to that item. If no descendant has input focus, 0 is returned.

This function was introduced in  Qt 4.6.

See also hasFocus(), setFocus(), and QWidget::focusWidget().
*/
func (this *QGraphicsItem) FocusItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9focusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:264
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * focusScopeItem() const

/*

 */
func (this *QGraphicsItem) FocusScopeItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14focusScopeItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabMouse()

/*
Grabs the mouse input.

This item will receive all mouse events for the scene until any of the following events occurs:


The item becomes invisible
The item is removed from the scene
The item is deleted
The item call ungrabMouse()
Another item calls grabMouse(); the item will regain the mouse grab when the other item calls ungrabMouse().


When an item gains the mouse grab, it receives a QEvent::GrabMouse event. When it loses the mouse grab, it receives a QEvent::UngrabMouse event. These events can be used to detect when your item gains or loses the mouse grab through other means than receiving mouse button events.

It is almost never necessary to explicitly grab the mouse in Qt, as Qt grabs and releases it sensibly. In particular, Qt grabs the mouse when you press a mouse button, and keeps the mouse grabbed until you release the last mouse button. Also, Qt::Popup widgets implicitly call grabMouse() when shown, and ungrabMouse() when hidden.

Note that only visible items can grab mouse input. Calling grabMouse() on an invisible item has no effect.

Keyboard events are not affected.

This function was introduced in  Qt 4.4.

See also QGraphicsScene::mouseGrabberItem(), ungrabMouse(), and grabKeyboard().
*/
func (this *QGraphicsItem) GrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9grabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabMouse()

/*
Releases the mouse grab.

This function was introduced in  Qt 4.4.

See also grabMouse() and ungrabKeyboard().
*/
func (this *QGraphicsItem) UngrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11ungrabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabKeyboard()

/*
Grabs the keyboard input.

The item will receive all keyboard input to the scene until one of the following events occur:


The item becomes invisible
The item is removed from the scene
The item is deleted
The item calls ungrabKeyboard()
Another item calls grabKeyboard(); the item will regain the keyboard grab when the other item calls ungrabKeyboard().


When an item gains the keyboard grab, it receives a QEvent::GrabKeyboard event. When it loses the keyboard grab, it receives a QEvent::UngrabKeyboard event. These events can be used to detect when your item gains or loses the keyboard grab through other means than gaining input focus.

It is almost never necessary to explicitly grab the keyboard in Qt, as Qt grabs and releases it sensibly. In particular, Qt grabs the keyboard when your item gains input focus, and releases it when your item loses input focus, or when the item is hidden.

Note that only visible items can grab keyboard input. Calling grabKeyboard() on an invisible item has no effect.

Keyboard events are not affected.

This function was introduced in  Qt 4.4.

See also ungrabKeyboard(), grabMouse(), and setFocus().
*/
func (this *QGraphicsItem) GrabKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12grabKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabKeyboard()

/*
Releases the keyboard grab.

This function was introduced in  Qt 4.4.

See also grabKeyboard() and ungrabMouse().
*/
func (this *QGraphicsItem) UngrabKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14ungrabKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:272
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos() const

/*
Returns the position of the item in parent coordinates. If the item has no parent, its position is given in scene coordinates.

The position of the item describes its origin (local coordinate (0, 0)) in parent coordinates; this function returns the same as mapToParent(0, 0).

For convenience, you can also call scenePos() to determine the item's position in scene coordinates, regardless of its parent.

See also x(), y(), setPos(), transform(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) Pos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:273
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal x() const

/*
This convenience function is equivalent to calling pos().x().

See also setX() and y().
*/
func (this *QGraphicsItem) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(qreal)

/*
Set's the x coordinate of the item's position. Equivalent to calling setPos(x, y()).

This function was introduced in  Qt 4.6.

See also x() and setPos().
*/
func (this *QGraphicsItem) SetX(x float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:275
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y() const

/*
This convenience function is equivalent to calling pos().y().

See also setY() and x().
*/
func (this *QGraphicsItem) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(qreal)

/*
Set's the y coordinate of the item's position. Equivalent to calling setPos(x(), y).

This function was introduced in  Qt 4.6.

See also y(), x(), and setPos().
*/
func (this *QGraphicsItem) SetY(y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:277
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos() const

/*
Returns the item's position in scene coordinates. This is equivalent to calling mapToScene(0, 0).

See also pos(), sceneTransform(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) ScenePos() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8scenePosEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPos(const QPointF &)

/*
Sets the position of the item to pos, which is in parent coordinates. For items with no parent, pos is in scene coordinates.

The position of the item describes its origin (local coordinate (0, 0)) in parent coordinates.

See also pos(), scenePos(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) SetPos(pos qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPointF_PTR() != nil {
		convArg0 = pos.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6setPosERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:279
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setPos(qreal, qreal)

/*
Sets the position of the item to pos, which is in parent coordinates. For items with no parent, pos is in scene coordinates.

The position of the item describes its origin (local coordinate (0, 0)) in parent coordinates.

See also pos(), scenePos(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) SetPos_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6setPosEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:280
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBy(qreal, qreal)

/*
Moves the item by dx points horizontally, and dy point vertically. This function is equivalent to calling setPos(pos() + QPointF(dx, dy)).
*/
func (this *QGraphicsItem) MoveBy(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6moveByEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)

/*
If this item is part of a scene that is viewed by a QGraphicsView, this convenience function will attempt to scroll the view to ensure that rect is visible inside the view's viewport. If rect is a null rect (the default), QGraphicsItem will default to the item's bounding rect. xmargin and ymargin are the number of pixels the view should use for margins.

If the specified rect cannot be reached, the contents are scrolled to the nearest valid position.

If this item is not viewed by a QGraphicsView, this function does nothing.

See also QGraphicsView::ensureVisible().
*/
func (this *QGraphicsItem) EnsureVisible(rect qtcore.QRectF_ITF, xmargin int, ymargin int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)

/*
If this item is part of a scene that is viewed by a QGraphicsView, this convenience function will attempt to scroll the view to ensure that rect is visible inside the view's viewport. If rect is a null rect (the default), QGraphicsItem will default to the item's bounding rect. xmargin and ymargin are the number of pixels the view should use for margins.

If the specified rect cannot be reached, the contents are scrolled to the nearest valid position.

If this item is not viewed by a QGraphicsView, this function does nothing.

See also QGraphicsView::ensureVisible().
*/
func (this *QGraphicsItem) EnsureVisible__() {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)

/*
If this item is part of a scene that is viewed by a QGraphicsView, this convenience function will attempt to scroll the view to ensure that rect is visible inside the view's viewport. If rect is a null rect (the default), QGraphicsItem will default to the item's bounding rect. xmargin and ymargin are the number of pixels the view should use for margins.

If the specified rect cannot be reached, the contents are scrolled to the nearest valid position.

If this item is not viewed by a QGraphicsView, this function does nothing.

See also QGraphicsView::ensureVisible().
*/
func (this *QGraphicsItem) EnsureVisible__1(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)

/*
If this item is part of a scene that is viewed by a QGraphicsView, this convenience function will attempt to scroll the view to ensure that rect is visible inside the view's viewport. If rect is a null rect (the default), QGraphicsItem will default to the item's bounding rect. xmargin and ymargin are the number of pixels the view should use for margins.

If the specified rect cannot be reached, the contents are scrolled to the nearest valid position.

If this item is not viewed by a QGraphicsView, this function does nothing.

See also QGraphicsView::ensureVisible().
*/
func (this *QGraphicsItem) EnsureVisible__2(rect qtcore.QRectF_ITF, xmargin int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)

/*
If this item is part of a scene that is viewed by a QGraphicsView, this convenience function will attempt to scroll the view to ensure that rect is visible inside the view's viewport. If rect is a null rect (the default), QGraphicsItem will default to the item's bounding rect. xmargin and ymargin are the number of pixels the view should use for margins.

If the specified rect cannot be reached, the contents are scrolled to the nearest valid position.

If this item is not viewed by a QGraphicsView, this function does nothing.

See also QGraphicsView::ensureVisible().
*/
func (this *QGraphicsItem) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)

/*
If this item is part of a scene that is viewed by a QGraphicsView, this convenience function will attempt to scroll the view to ensure that rect is visible inside the view's viewport. If rect is a null rect (the default), QGraphicsItem will default to the item's bounding rect. xmargin and ymargin are the number of pixels the view should use for margins.

If the specified rect cannot be reached, the contents are scrolled to the nearest valid position.

If this item is not viewed by a QGraphicsView, this function does nothing.

See also QGraphicsView::ensureVisible().
*/
func (this *QGraphicsItem) EnsureVisible_1_(x float64, y float64, w float64, h float64) {
	// arg: 4, int=Int, =Invalid, , Invalid
	xmargin := int(50)
	// arg: 5, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)

/*
If this item is part of a scene that is viewed by a QGraphicsView, this convenience function will attempt to scroll the view to ensure that rect is visible inside the view's viewport. If rect is a null rect (the default), QGraphicsItem will default to the item's bounding rect. xmargin and ymargin are the number of pixels the view should use for margins.

If the specified rect cannot be reached, the contents are scrolled to the nearest valid position.

If this item is not viewed by a QGraphicsView, this function does nothing.

See also QGraphicsView::ensureVisible().
*/
func (this *QGraphicsItem) EnsureVisible_1_1(x float64, y float64, w float64, h float64, xmargin int) {
	// arg: 5, int=Int, =Invalid, , Invalid
	ymargin := int(50)
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:286
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix matrix() const

/*

 */
func (this *QGraphicsItem) Matrix() *qtgui.QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:287
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix sceneMatrix() const

/*

 */
func (this *QGraphicsItem) SceneMatrix() *qtgui.QMatrix /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11sceneMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &, bool)

/*

 */
func (this *QGraphicsItem) SetMatrix(matrix qtgui.QMatrix_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix &, bool)

/*

 */
func (this *QGraphicsItem) SetMatrix__(matrix qtgui.QMatrix_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid, , Invalid
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetMatrix()

/*

 */
func (this *QGraphicsItem) ResetMatrix() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11resetMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:290
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transform() const

/*
Returns this item's transformation matrix.

The transformation matrix is combined with the item's rotation(), scale() and transformations() into a combined transformations for the item.

The default transformation matrix is an identity matrix.

This function was introduced in  Qt 4.3.

See also setTransform() and sceneTransform().
*/
func (this *QGraphicsItem) Transform() *qtgui.QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9transformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:291
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform sceneTransform() const

/*
Returns this item's scene transformation matrix. This matrix can be used to map coordinates and geometrical shapes from this item's local coordinate system to the scene's coordinate system. To map coordinates from the scene, you must first invert the returned matrix.

Example:


  QGraphicsRectItem rect;
  rect.setPos(100, 100);

  rect.sceneTransform().map(QPointF(0, 0));
  // returns QPointF(100, 100);

  rect.sceneTransform().inverted().map(QPointF(100, 100));
  // returns QPointF(0, 0);



Unlike transform(), which returns only an item's local transformation, this function includes the item's (and any parents') position, and all the transfomation properties.

This function was introduced in  Qt 4.3.

See also transform(), setTransform(), scenePos(), The Graphics View Coordinate System, and Transformations.
*/
func (this *QGraphicsItem) SceneTransform() *qtgui.QTransform /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14sceneTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:292
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform deviceTransform(const QTransform &) const

/*
Returns this item's device transformation matrix, using viewportTransform to map from scene to device coordinates. This matrix can be used to map coordinates and geometrical shapes from this item's local coordinate system to the viewport's (or any device's) coordinate system. To map coordinates from the viewport, you must first invert the returned matrix.

Example:


  QGraphicsRectItem rect;
  rect.setPos(100, 100);

  rect.deviceTransform(view->viewportTransform()).map(QPointF(0, 0));
  // returns the item's (0, 0) point in view's viewport coordinates

  rect.deviceTransform(view->viewportTransform()).inverted().map(QPointF(100, 100));
  // returns view's viewport's (100, 100) coordinate in item coordinates



This function is the same as combining this item's scene transform with the view's viewport transform, but it also understands the ItemIgnoresTransformations flag. The device transform can be used to do accurate coordinate mapping (and collision detection) for untransformable items.

This function was introduced in  Qt 4.3.

See also transform(), setTransform(), scenePos(), The Graphics View Coordinate System, and itemTransform().
*/
func (this *QGraphicsItem) DeviceTransform(viewportTransform qtgui.QTransform_ITF) *qtgui.QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if viewportTransform != nil && viewportTransform.QTransform_PTR() != nil {
		convArg0 = viewportTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem15deviceTransformERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:293
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform itemTransform(const QGraphicsItem *, bool *) const

/*
Returns a QTransform that maps coordinates from this item to other. If ok is not null, and if there is no such transform, the boolean pointed to by ok will be set to false; otherwise it will be set to true.

This transform provides an alternative to the mapToItem() or mapFromItem() functions, by returning the appropriate transform so that you can map shapes and coordinates yourself. It also helps you write more efficient code when repeatedly mapping between the same two items.

Note: In rare circumstances, there is no transform that maps between two items.

This function was introduced in  Qt 4.5.

See also mapToItem(), mapFromItem(), and deviceTransform().
*/
func (this *QGraphicsItem) ItemTransform(other QGraphicsItem_ITF /*777 const QGraphicsItem **/, ok *bool) *qtgui.QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGraphicsItem_PTR() != nil {
		convArg0 = other.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13itemTransformEPKS_Pb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:293
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform itemTransform(const QGraphicsItem *, bool *) const

/*
Returns a QTransform that maps coordinates from this item to other. If ok is not null, and if there is no such transform, the boolean pointed to by ok will be set to false; otherwise it will be set to true.

This transform provides an alternative to the mapToItem() or mapFromItem() functions, by returning the appropriate transform so that you can map shapes and coordinates yourself. It also helps you write more efficient code when repeatedly mapping between the same two items.

Note: In rare circumstances, there is no transform that maps between two items.

This function was introduced in  Qt 4.5.

See also mapToItem(), mapFromItem(), and deviceTransform().
*/
func (this *QGraphicsItem) ItemTransform__(other QGraphicsItem_ITF /*777 const QGraphicsItem **/) *qtgui.QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGraphicsItem_PTR() != nil {
		convArg0 = other.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, bool *=Pointer, =Invalid, , Invalid
	var ok unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13itemTransformEPKS_Pb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, ok)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, bool)

/*
Sets the item's current transformation matrix to matrix.

If combine is true, then matrix is combined with the current matrix; otherwise, matrix replaces the current matrix. combine is false by default.

To simplify interaction with items using a transformed view, QGraphicsItem provides mapTo... and mapFrom... functions that can translate between items' and the scene's coordinates. For example, you can call mapToScene() to map an item coordiate to a scene coordinate, or mapFromScene() to map from scene coordinates to item coordinates.

The transformation matrix is combined with the item's rotation(), scale() and transformations() into a combined transformation that maps the item's coordinate system to its parent.

This function was introduced in  Qt 4.3.

See also transform(), setRotation(), setScale(), setTransformOriginPoint(), The Graphics View Coordinate System, and Transformations.
*/
func (this *QGraphicsItem) SetTransform(matrix qtgui.QTransform_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, bool)

/*
Sets the item's current transformation matrix to matrix.

If combine is true, then matrix is combined with the current matrix; otherwise, matrix replaces the current matrix. combine is false by default.

To simplify interaction with items using a transformed view, QGraphicsItem provides mapTo... and mapFrom... functions that can translate between items' and the scene's coordinates. For example, you can call mapToScene() to map an item coordiate to a scene coordinate, or mapFromScene() to map from scene coordinates to item coordinates.

The transformation matrix is combined with the item's rotation(), scale() and transformations() into a combined transformation that maps the item's coordinate system to its parent.

This function was introduced in  Qt 4.3.

See also transform(), setRotation(), setScale(), setTransformOriginPoint(), The Graphics View Coordinate System, and Transformations.
*/
func (this *QGraphicsItem) SetTransform__(matrix qtgui.QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid, , Invalid
	combine := false
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:295
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetTransform()

/*
Resets this item's transformation matrix to the identity matrix or all the transformation properties to their default values. This is equivalent to calling setTransform(QTransform()).

This function was introduced in  Qt 4.3.

See also setTransform() and transform().
*/
func (this *QGraphicsItem) ResetTransform() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14resetTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRotation(qreal)

/*
Sets the clockwise rotation angle, in degrees, around the Z axis. The default value is 0 (i.e., the item is not rotated). Assigning a negative value will rotate the item counter-clockwise. Normally the rotation angle is in the range (-360, 360), but it's also possible to assign values outside of this range (e.g., a rotation of 370 degrees is the same as a rotation of 10 degrees).

The item is rotated around its transform origin point, which by default is (0, 0). You can select a different transformation origin by calling setTransformOriginPoint().

The rotation is combined with the item's scale(), transform() and transformations() to map the item's coordinate system to the parent item.

This function was introduced in  Qt 4.6.

See also rotation(), setTransformOriginPoint(), and Transformations.
*/
func (this *QGraphicsItem) SetRotation(angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11setRotationEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:303
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rotation() const

/*
Returns the clockwise rotation, in degrees, around the Z axis. The default value is 0 (i.e., the item is not rotated).

The rotation is combined with the item's scale(), transform() and transformations() to map the item's coordinate system to the parent item.

This function was introduced in  Qt 4.6.

See also setRotation(), transformOriginPoint(), and Transformations.
*/
func (this *QGraphicsItem) Rotation() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8rotationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScale(qreal)

/*
Sets the scale factor of the item. The default scale factor is 1.0 (i.e., the item is not scaled). A scale factor of 0.0 will collapse the item to a single point. If you provide a negative scale factor, the item will be flipped and mirrored (i.e., rotated 180 degrees).

The item is scaled around its transform origin point, which by default is (0, 0). You can select a different transformation origin by calling setTransformOriginPoint().

The scale is combined with the item's rotation(), transform() and transformations() to map the item's coordinate system to the parent item.

This function was introduced in  Qt 4.6.

See also scale(), setTransformOriginPoint(), and Transformations Example.
*/
func (this *QGraphicsItem) SetScale(scale float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scale)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:306
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal scale() const

/*
Returns the scale factor of the item. The default scale factor is 1.0 (i.e., the item is not scaled).

The scale is combined with the item's rotation(), transform() and transformations() to map the item's coordinate system to the parent item.

This function was introduced in  Qt 4.6.

See also setScale(), rotation(), and Transformations.
*/
func (this *QGraphicsItem) Scale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5scaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:311
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF transformOriginPoint() const

/*
Returns the origin point for the transformation in item coordinates.

The default is QPointF(0,0).

This function was introduced in  Qt 4.6.

See also setTransformOriginPoint() and Transformations.
*/
func (this *QGraphicsItem) TransformOriginPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem20transformOriginPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:312
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformOriginPoint(const QPointF &)

/*
Sets the origin point for the transformation in item coordinates.

This function was introduced in  Qt 4.6.

See also transformOriginPoint() and Transformations.
*/
func (this *QGraphicsItem) SetTransformOriginPoint(origin qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if origin != nil && origin.QPointF_PTR() != nil {
		convArg0 = origin.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem23setTransformOriginPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:313
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setTransformOriginPoint(qreal, qreal)

/*
Sets the origin point for the transformation in item coordinates.

This function was introduced in  Qt 4.6.

See also transformOriginPoint() and Transformations.
*/
func (this *QGraphicsItem) SetTransformOriginPoint_1(ax float64, ay float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem23setTransformOriginPointEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ax, ay)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:316
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void advance(int)

/*
This virtual function is called twice for all items by the QGraphicsScene::advance() slot. In the first phase, all items are called with phase == 0, indicating that items on the scene are about to advance, and then all items are called with phase == 1. Reimplement this function to update your item if you need simple scene-controlled animation.

The default implementation does nothing.

This function is intended for animations. An alternative is to multiple-inherit from QObject and QGraphicsItem and use the Animation Framework.

See also QGraphicsScene::advance() and QTimeLine.
*/
func (this *QGraphicsItem) Advance(phase int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem7advanceEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), phase)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal zValue() const

/*
Returns the Z-value of the item. The Z-value affects the stacking order of sibling (neighboring) items.

The default Z-value is 0.

See also setZValue(), Sorting, stackBefore(), and ItemStacksBehindParent.
*/
func (this *QGraphicsItem) ZValue() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem6zValueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZValue(qreal)

/*
Sets the Z-value of the item to z. The Z value decides the stacking order of sibling (neighboring) items. A sibling item of high Z value will always be drawn on top of another sibling item with a lower Z value.

If you restore the Z value, the item's insertion order will decide its stacking order.

The Z-value does not affect the item's size in any way.

The default Z-value is 0.

See also zValue(), Sorting, stackBefore(), and ItemStacksBehindParent.
*/
func (this *QGraphicsItem) SetZValue(z float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setZValueEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:321
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackBefore(const QGraphicsItem *)

/*
Stacks this item before sibling, which must be a sibling item (i.e., the two items must share the same parent item, or must both be toplevel items). The sibling must have the same Z value as this item, otherwise calling this function will have no effect.

By default, all sibling items are stacked by insertion order (i.e., the first item you add is drawn before the next item you add). If two items' Z values are different, then the item with the highest Z value is drawn on top. When the Z values are the same, the insertion order will decide the stacking order.

This function was introduced in  Qt 4.6.

See also setZValue(), ItemStacksBehindParent, and Sorting.
*/
func (this *QGraphicsItem) StackBefore(sibling QGraphicsItem_ITF /*777 const QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if sibling != nil && sibling.QGraphicsItem_PTR() != nil {
		convArg0 = sibling.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11stackBeforeEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:324
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect() const

/*
This pure virtual function defines the outer bounds of the item as a rectangle; all painting must be restricted to inside an item's bounding rect. QGraphicsView uses this to determine whether the item requires redrawing.

Although the item's shape can be arbitrary, the bounding rect is always rectangular, and it is unaffected by the items' transformation.

If you want to change the item's bounding rectangle, you must first call prepareGeometryChange(). This notifies the scene of the imminent change, so that it can update its item geometry index; otherwise, the scene will be unaware of the item's new geometry, and the results are undefined (typically, rendering artifacts are left within the view).

Reimplement this function to let QGraphicsView determine what parts of the widget, if any, need to be redrawn.

Note: For shapes that paint an outline / stroke, it is important to include half the pen width in the bounding rect. It is not necessary to compensate for antialiasing, though.

Example:


  QRectF CircleItem::boundingRect() const
  {
      qreal penWidth = 1;
      return QRectF(-radius - penWidth / 2, -radius - penWidth / 2,
                    diameter + penWidth, diameter + penWidth);
  }



See also boundingRegion(), shape(), contains(), The Graphics View Coordinate System, and prepareGeometryChange().
*/
func (this *QGraphicsItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:325
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF childrenBoundingRect() const

/*
Returns the bounding rect of this item's descendants (i.e., its children, their children, etc.) in local coordinates. The rectangle will contain all descendants after they have been mapped to local coordinates. If the item has no children, this function returns an empty QRectF.

This does not include this item's own bounding rect; it only returns its descendants' accumulated bounding rect. If you need to include this item's bounding rect, you can add boundingRect() to childrenBoundingRect() using QRectF::operator|().

This function is linear in complexity; it determines the size of the returned bounding rect by iterating through all descendants.

See also boundingRect() and sceneBoundingRect().
*/
func (this *QGraphicsItem) ChildrenBoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem20childrenBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:326
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF sceneBoundingRect() const

/*
Returns the bounding rect of this item in scene coordinates, by combining sceneTransform() with boundingRect().

See also boundingRect() and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) SceneBoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17sceneBoundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:327
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape() const

/*
Returns the shape of this item as a QPainterPath in local coordinates. The shape is used for many things, including collision detection, hit tests, and for the QGraphicsScene::items() functions.

The default implementation calls boundingRect() to return a simple rectangular shape, but subclasses can reimplement this function to return a more accurate shape for non-rectangular items. For example, a round item may choose to return an elliptic shape for better collision detection. For example:


  QPainterPath RoundItem::shape() const
  {
      QPainterPath path;
      path.addEllipse(boundingRect());
      return path;
  }



The outline of a shape can vary depending on the width and style of the pen used when drawing. If you want to include this outline in the item's shape, you can create a shape from the stroke using QPainterPathStroker.

This function is called by the default implementations of contains() and collidesWithPath().

See also boundingRect(), contains(), prepareGeometryChange(), and QPainterPathStroker.
*/
func (this *QGraphicsItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:328
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isClipped() const

/*
Returns true if this item is clipped. An item is clipped if it has either set the ItemClipsToShape flag, or if it or any of its ancestors has set the ItemClipsChildrenToShape flag.

Clipping affects the item's appearance (i.e., painting), as well as mouse and hover event delivery.

See also clipPath(), shape(), and setFlags().
*/
func (this *QGraphicsItem) IsClipped() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9isClippedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:329
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath clipPath() const

/*
Returns this item's clip path, or an empty QPainterPath if this item is not clipped. The clip path constrains the item's appearance and interaction (i.e., restricts the area the item can draw within and receive events for).

You can enable clipping by setting the ItemClipsToShape or ItemClipsChildrenToShape flags. The item's clip path is calculated by intersecting all clipping ancestors' shapes. If the item sets ItemClipsToShape, the final clip is intersected with the item's own shape.

Note: Clipping introduces a performance penalty for all items involved; you should generally avoid using clipping if you can (e.g., if your items always draw inside boundingRect() or shape() boundaries, clipping is not necessary).

This function was introduced in  Qt 4.5.

See also isClipped(), shape(), and setFlags().
*/
func (this *QGraphicsItem) ClipPath() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8clipPathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:330
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &) const

/*
Returns true if this item contains point, which is in local coordinates; otherwise, false is returned. It is most often called from QGraphicsView to determine what item is under the cursor, and for that reason, the implementation of this function should be as light-weight as possible.

By default, this function calls shape(), but you can reimplement it in a subclass to provide a (perhaps more efficient) implementation.

See also shape(), boundingRect(), and collidesWithPath().
*/
func (this *QGraphicsItem) Contains(point qtcore.QPointF_ITF) bool {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:331
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool collidesWithItem(const QGraphicsItem *, Qt::ItemSelectionMode) const

/*
Returns true if this item collides with other; otherwise returns false.

The mode is applied to other, and the resulting shape or bounding rectangle is then compared to this item's shape. The default value for mode is Qt::IntersectsItemShape; other collides with this item if it either intersects, contains, or is contained by this item's shape (see Qt::ItemSelectionMode for details).

The default implementation is based on shape intersection, and it calls shape() on both items. Because the complexity of arbitrary shape-shape intersection grows with an order of magnitude when the shapes are complex, this operation can be noticably time consuming. You have the option of reimplementing this function in a subclass of QGraphicsItem to provide a custom algorithm. This allows you to make use of natural constraints in the shapes of your own items, in order to improve the performance of the collision detection. For instance, two untransformed perfectly circular items' collision can be determined very efficiently by comparing their positions and radii.

Keep in mind that when reimplementing this function and calling shape() or boundingRect() on other, the returned coordinates must be mapped to this item's coordinate system before any intersection can take place.

See also contains() and shape().
*/
func (this *QGraphicsItem) CollidesWithItem(other QGraphicsItem_ITF /*777 const QGraphicsItem **/, mode int) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGraphicsItem_PTR() != nil {
		convArg0 = other.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithItemEPKS_N2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:331
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool collidesWithItem(const QGraphicsItem *, Qt::ItemSelectionMode) const

/*
Returns true if this item collides with other; otherwise returns false.

The mode is applied to other, and the resulting shape or bounding rectangle is then compared to this item's shape. The default value for mode is Qt::IntersectsItemShape; other collides with this item if it either intersects, contains, or is contained by this item's shape (see Qt::ItemSelectionMode for details).

The default implementation is based on shape intersection, and it calls shape() on both items. Because the complexity of arbitrary shape-shape intersection grows with an order of magnitude when the shapes are complex, this operation can be noticably time consuming. You have the option of reimplementing this function in a subclass of QGraphicsItem to provide a custom algorithm. This allows you to make use of natural constraints in the shapes of your own items, in order to improve the performance of the collision detection. For instance, two untransformed perfectly circular items' collision can be determined very efficiently by comparing their positions and radii.

Keep in mind that when reimplementing this function and calling shape() or boundingRect() on other, the returned coordinates must be mapped to this item's coordinate system before any intersection can take place.

See also contains() and shape().
*/
func (this *QGraphicsItem) CollidesWithItem__(other QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGraphicsItem_PTR() != nil {
		convArg0 = other.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithItemEPKS_N2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:332
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool collidesWithPath(const QPainterPath &, Qt::ItemSelectionMode) const

/*
Returns true if this item collides with path.

The collision is determined by mode. The default value for mode is Qt::IntersectsItemShape; path collides with this item if it either intersects, contains, or is contained by this item's shape.

Note that this function checks whether the item's shape or bounding rectangle (depending on mode) is contained within path, and not whether path is contained within the items shape or bounding rectangle.

See also collidesWithItem(), contains(), and shape().
*/
func (this *QGraphicsItem) CollidesWithPath(path qtgui.QPainterPath_ITF, mode int) bool {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithPathERK12QPainterPathN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:332
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool collidesWithPath(const QPainterPath &, Qt::ItemSelectionMode) const

/*
Returns true if this item collides with path.

The collision is determined by mode. The default value for mode is Qt::IntersectsItemShape; path collides with this item if it either intersects, contains, or is contained by this item's shape.

Note that this function checks whether the item's shape or bounding rectangle (depending on mode) is contained within path, and not whether path is contained within the items shape or bounding rectangle.

See also collidesWithItem(), contains(), and shape().
*/
func (this *QGraphicsItem) CollidesWithPath__(path qtgui.QPainterPath_ITF) bool {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	// arg: 1, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithPathERK12QPainterPathN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:333
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> collidingItems(Qt::ItemSelectionMode) const

/*
Returns a list of all items that collide with this item.

The way collisions are detected is determined by applying mode to items that are compared to this item, i.e., each item's shape or bounding rectangle is checked against this item's shape. The default value for mode is Qt::IntersectsItemShape.

See also collidesWithItem().
*/
func (this *QGraphicsItem) CollidingItems(mode int) *QGraphicsItemList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14collidingItemsEN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:333
// index:0
// Public Visibility=Default Availability=Available
// [8] QList<QGraphicsItem *> collidingItems(Qt::ItemSelectionMode) const

/*
Returns a list of all items that collide with this item.

The way collisions are detected is determined by applying mode to items that are compared to this item, i.e., each item's shape or bounding rectangle is checked against this item's shape. The default value for mode is Qt::IntersectsItemShape.

See also collidesWithItem().
*/
func (this *QGraphicsItem) CollidingItems__() *QGraphicsItemList /*lll*/ {
	// arg: 0, Qt::ItemSelectionMode=Elaborated, Qt::ItemSelectionMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14collidingItemsEN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGraphicsItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:334
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObscured(const QRectF &) const

/*
This is an overloaded function.

Returns true if rect is completely obscured by the opaque shape of any of colliding items above it (i.e., with a higher Z value than this item).

This function was introduced in  Qt 4.3.

See also opaqueArea().
*/
func (this *QGraphicsItem) IsObscured(rect qtcore.QRectF_ITF) bool {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:334
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObscured(const QRectF &) const

/*
This is an overloaded function.

Returns true if rect is completely obscured by the opaque shape of any of colliding items above it (i.e., with a higher Z value than this item).

This function was introduced in  Qt 4.3.

See also opaqueArea().
*/
func (this *QGraphicsItem) IsObscured__() bool {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:335
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool isObscured(qreal, qreal, qreal, qreal) const

/*
This is an overloaded function.

Returns true if rect is completely obscured by the opaque shape of any of colliding items above it (i.e., with a higher Z value than this item).

This function was introduced in  Qt 4.3.

See also opaqueArea().
*/
func (this *QGraphicsItem) IsObscured_1(x float64, y float64, w float64, h float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:336
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *) const

/*
Returns true if this item's bounding rect is completely obscured by the opaque shape of item.

The base implementation maps item's opaqueArea() to this item's coordinate system, and then checks if this item's boundingRect() is fully contained within the mapped shape.

You can reimplement this function to provide a custom algorithm for determining whether this item is obscured by item.

See also opaqueArea() and isObscured().
*/
func (this *QGraphicsItem) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12isObscuredByEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:337
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea() const

/*
This virtual function returns a shape representing the area where this item is opaque. An area is opaque if it is filled using an opaque brush or color (i.e., not transparent).

This function is used by isObscuredBy(), which is called by underlying items to determine if they are obscured by this item.

The default implementation returns an empty QPainterPath, indicating that this item is completely transparent and does not obscure any other items.

See also isObscuredBy(), isObscured(), and shape().
*/
func (this *QGraphicsItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:339
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion boundingRegion(const QTransform &) const

/*
Returns the bounding region for this item. The coordinate space of the returned region depends on itemToDeviceTransform. If you pass an identity QTransform as a parameter, this function will return a local coordinate region.

The bounding region describes a coarse outline of the item's visual contents. Although it's expensive to calculate, it's also more precise than boundingRect(), and it can help to avoid unnecessary repainting when an item is updated. This is particularly efficient for thin items (e.g., lines or simple polygons). You can tune the granularity for the bounding region by calling setBoundingRegionGranularity(). The default granularity is 0; in which the item's bounding region is the same as its bounding rect.

itemToDeviceTransform is the transformation from item coordinates to device coordinates. If you want this function to return a QRegion in scene coordinates, you can pass sceneTransform() as an argument.

This function was introduced in  Qt 4.4.

See also boundingRegionGranularity().
*/
func (this *QGraphicsItem) BoundingRegion(itemToDeviceTransform qtgui.QTransform_ITF) *qtgui.QRegion /*123*/ {
	var convArg0 unsafe.Pointer
	if itemToDeviceTransform != nil && itemToDeviceTransform.QTransform_PTR() != nil {
		convArg0 = itemToDeviceTransform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14boundingRegionERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:340
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal boundingRegionGranularity() const

/*
Returns the item's bounding region granularity; a value between and including 0 and 1. The default value is 0 (i.e., the lowest granularity, where the bounding region corresponds to the item's bounding rectangle).

This function was introduced in  Qt 4.4.

See also setBoundingRegionGranularity().
*/
func (this *QGraphicsItem) BoundingRegionGranularity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem25boundingRegionGranularityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:341
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundingRegionGranularity(qreal)

/*
Sets the bounding region granularity to granularity; a value between and including 0 and 1. The default value is 0 (i.e., the lowest granularity, where the bounding region corresponds to the item's bounding rectangle).

The granularity is used by boundingRegion() to calculate how fine the bounding region of the item should be. The highest achievable granularity is 1, where boundingRegion() will return the finest outline possible for the respective device (e.g., for a QGraphicsView viewport, this gives you a pixel-perfect bounding region). The lowest possible granularity is 0. The value of granularity describes the ratio between device resolution and the resolution of the bounding region (e.g., a value of 0.25 will provide a region where each chunk corresponds to 4x4 device units / pixels).

This function was introduced in  Qt 4.4.

See also boundingRegionGranularity().
*/
func (this *QGraphicsItem) SetBoundingRegionGranularity(granularity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem28setBoundingRegionGranularityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), granularity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:344
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
This function, which is usually called by QGraphicsView, paints the contents of an item in local coordinates.

Reimplement this function in a QGraphicsItem subclass to provide the item's painting implementation, using painter. The option parameter provides style options for the item, such as its state, exposed area and its level-of-detail hints. The widget argument is optional. If provided, it points to the widget that is being painted on; otherwise, it is 0. For cached painting, widget is always 0.


  void RoundRectItem::paint(QPainter *painter,
                            const QStyleOptionGraphicsItem *option,
                            QWidget *widget)
  {
      painter->drawRoundedRect(-10, -10, 20, 20, 5, 5);
  }



The painter's pen is 0-width by default, and its pen is initialized to the QPalette::Text brush from the paint device's palette. The brush is initialized to QPalette::Window.

Make sure to constrain all painting inside the boundaries of boundingRect() to avoid rendering artifacts (as QGraphicsView does not clip the painter for you). In particular, when QPainter renders the outline of a shape using an assigned QPen, half of the outline will be drawn outside, and half inside, the shape you're rendering (e.g., with a pen width of 2 units, you must draw outlines 1 unit inside boundingRect()). QGraphicsItem does not support use of cosmetic pens with a non-zero width.

All painting is done in local coordinates.

Note: It is mandatory that an item will always redraw itself in the exact same way, unless update() was called; otherwise visual artifacts may occur. In other words, two subsequent calls to paint() must always produce the same output, unless update() was called between them.

Note: Enabling caching for an item does not guarantee that paint() will be invoked only once by the Graphics View framework, even without any explicit call to update(). See the documentation of setCacheMode() for more details.

See also setCacheMode(), QPen::width(), Item Coordinates, and ItemUsesExtendedStyleOption.
*/
func (this *QGraphicsItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:344
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
This function, which is usually called by QGraphicsView, paints the contents of an item in local coordinates.

Reimplement this function in a QGraphicsItem subclass to provide the item's painting implementation, using painter. The option parameter provides style options for the item, such as its state, exposed area and its level-of-detail hints. The widget argument is optional. If provided, it points to the widget that is being painted on; otherwise, it is 0. For cached painting, widget is always 0.


  void RoundRectItem::paint(QPainter *painter,
                            const QStyleOptionGraphicsItem *option,
                            QWidget *widget)
  {
      painter->drawRoundedRect(-10, -10, 20, 20, 5, 5);
  }



The painter's pen is 0-width by default, and its pen is initialized to the QPalette::Text brush from the paint device's palette. The brush is initialized to QPalette::Window.

Make sure to constrain all painting inside the boundaries of boundingRect() to avoid rendering artifacts (as QGraphicsView does not clip the painter for you). In particular, when QPainter renders the outline of a shape using an assigned QPen, half of the outline will be drawn outside, and half inside, the shape you're rendering (e.g., with a pen width of 2 units, you must draw outlines 1 unit inside boundingRect()). QGraphicsItem does not support use of cosmetic pens with a non-zero width.

All painting is done in local coordinates.

Note: It is mandatory that an item will always redraw itself in the exact same way, unless update() was called; otherwise visual artifacts may occur. In other words, two subsequent calls to paint() must always produce the same output, unless update() was called between them.

Note: Enabling caching for an item does not guarantee that paint() will be invoked only once by the Graphics View framework, even without any explicit call to update(). See the documentation of setCacheMode() for more details.

See also setCacheMode(), QPen::width(), Item Coordinates, and ItemUsesExtendedStyleOption.
*/
func (this *QGraphicsItem) Paint__(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:345
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)

/*
Schedules a redraw of the area covered by rect in this item. You can call this function whenever your item needs to be redrawn, such as if it changes appearance or size.

This function does not cause an immediate paint; instead it schedules a paint request that is processed by QGraphicsView after control reaches the event loop. The item will only be redrawn if it is visible in any associated view.

As a side effect of the item being repainted, other items that overlap the area rect may also be repainted.

If the item is invisible (i.e., isVisible() returns false), this function does nothing.

See also paint() and boundingRect().
*/
func (this *QGraphicsItem) Update(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:345
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)

/*
Schedules a redraw of the area covered by rect in this item. You can call this function whenever your item needs to be redrawn, such as if it changes appearance or size.

This function does not cause an immediate paint; instead it schedules a paint request that is processed by QGraphicsView after control reaches the event loop. The item will only be redrawn if it is visible in any associated view.

As a side effect of the item being repainted, other items that overlap the area rect may also be repainted.

If the item is invisible (i.e., isVisible() returns false), this function does nothing.

See also paint() and boundingRect().
*/
func (this *QGraphicsItem) Update__() {
	// arg: 0, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:346
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void update(qreal, qreal, qreal, qreal)

/*
Schedules a redraw of the area covered by rect in this item. You can call this function whenever your item needs to be redrawn, such as if it changes appearance or size.

This function does not cause an immediate paint; instead it schedules a paint request that is processed by QGraphicsView after control reaches the event loop. The item will only be redrawn if it is visible in any associated view.

As a side effect of the item being repainted, other items that overlap the area rect may also be repainted.

If the item is invisible (i.e., isVisible() returns false), this function does nothing.

See also paint() and boundingRect().
*/
func (this *QGraphicsItem) Update_1(x float64, y float64, width float64, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6updateEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:347
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scroll(qreal, qreal, const QRectF &)

/*
Scrolls the contents of rect by dx, dy. If rect is a null rect (the default), the item's bounding rect is scrolled.

Scrolling provides a fast alternative to simply redrawing when the contents of the item (or parts of the item) are shifted vertically or horizontally. Depending on the current transformation and the capabilities of the paint device (i.e., the viewport), this operation may consist of simply moving pixels from one location to another using memmove(). In most cases this is faster than rerendering the entire area.

After scrolling, the item will issue an update for the newly exposed areas. If scrolling is not supported (e.g., you are rendering to an OpenGL viewport, which does not benefit from scroll optimizations), this function is equivalent to calling update(rect).

Note: Scrolling is only supported when QGraphicsItem::ItemCoordinateCache is enabled; in all other cases calling this function is equivalent to calling update(rect). If you for sure know that the item is opaque and not overlapped by other items, you can map the rect to viewport coordinates and scroll the viewport.


  QTransform xform = item->deviceTransform(view->viewportTransform());
  QRect deviceRect = xform.mapRect(rect).toAlignedRect();
  view->viewport()->scroll(dx, dy, deviceRect);



This function was introduced in  Qt 4.4.

See also boundingRect().
*/
func (this *QGraphicsItem) Scroll(dx float64, dy float64, rect qtcore.QRectF_ITF) {
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg2 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6scrollEddRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:347
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scroll(qreal, qreal, const QRectF &)

/*
Scrolls the contents of rect by dx, dy. If rect is a null rect (the default), the item's bounding rect is scrolled.

Scrolling provides a fast alternative to simply redrawing when the contents of the item (or parts of the item) are shifted vertically or horizontally. Depending on the current transformation and the capabilities of the paint device (i.e., the viewport), this operation may consist of simply moving pixels from one location to another using memmove(). In most cases this is faster than rerendering the entire area.

After scrolling, the item will issue an update for the newly exposed areas. If scrolling is not supported (e.g., you are rendering to an OpenGL viewport, which does not benefit from scroll optimizations), this function is equivalent to calling update(rect).

Note: Scrolling is only supported when QGraphicsItem::ItemCoordinateCache is enabled; in all other cases calling this function is equivalent to calling update(rect). If you for sure know that the item is opaque and not overlapped by other items, you can map the rect to viewport coordinates and scroll the viewport.


  QTransform xform = item->deviceTransform(view->viewportTransform());
  QRect deviceRect = xform.mapRect(rect).toAlignedRect();
  view->viewport()->scroll(dx, dy, deviceRect);



This function was introduced in  Qt 4.4.

See also boundingRect().
*/
func (this *QGraphicsItem) Scroll__(dx float64, dy float64) {
	// arg: 2, const QRectF &=LValueReference, QRectF=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6scrollEddRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:350
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToItem(const QGraphicsItem *, const QPointF &) const

/*
Maps the point point, which is in this item's coordinate system, to item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapToScene().

See also itemTransform(), mapToParent(), mapToScene(), transform(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToItem(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg1 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:353
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToItem(const QGraphicsItem *, const QRectF &) const

/*
Maps the point point, which is in this item's coordinate system, to item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapToScene().

See also itemTransform(), mapToParent(), mapToScene(), transform(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToItem_1(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, rect qtcore.QRectF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:359
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToItem(const QGraphicsItem *, const QPolygonF &) const

/*
Maps the point point, which is in this item's coordinate system, to item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapToScene().

See also itemTransform(), mapToParent(), mapToScene(), transform(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToItem_2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, polygon qtgui.QPolygonF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg1 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:362
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapToItem(const QGraphicsItem *, const QPainterPath &) const

/*
Maps the point point, which is in this item's coordinate system, to item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapToScene().

See also itemTransform(), mapToParent(), mapToScene(), transform(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToItem_3(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg1 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:381
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QPointF mapToItem(const QGraphicsItem *, qreal, qreal) const

/*
Maps the point point, which is in this item's coordinate system, to item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapToScene().

See also itemTransform(), mapToParent(), mapToScene(), transform(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToItem_4(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, x float64, y float64) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_dd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:384
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF mapToItem(const QGraphicsItem *, qreal, qreal, qreal, qreal) const

/*
Maps the point point, which is in this item's coordinate system, to item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapToScene().

See also itemTransform(), mapToParent(), mapToScene(), transform(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToItem_5(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_dddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:351
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToParent(const QPointF &) const

/*
Maps the point point, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped coordinate. If the item has no parent, point will be mapped to the scene's coordinate system.

See also mapToItem(), mapToScene(), transform(), mapFromParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToParent(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:354
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToParent(const QRectF &) const

/*
Maps the point point, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped coordinate. If the item has no parent, point will be mapped to the scene's coordinate system.

See also mapToItem(), mapToScene(), transform(), mapFromParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToParent_1(rect qtcore.QRectF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:360
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToParent(const QPolygonF &) const

/*
Maps the point point, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped coordinate. If the item has no parent, point will be mapped to the scene's coordinate system.

See also mapToItem(), mapToScene(), transform(), mapFromParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToParent_2(polygon qtgui.QPolygonF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:363
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapToParent(const QPainterPath &) const

/*
Maps the point point, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped coordinate. If the item has no parent, point will be mapped to the scene's coordinate system.

See also mapToItem(), mapToScene(), transform(), mapFromParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToParent_3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:382
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QPointF mapToParent(qreal, qreal) const

/*
Maps the point point, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped coordinate. If the item has no parent, point will be mapped to the scene's coordinate system.

See also mapToItem(), mapToScene(), transform(), mapFromParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToParent_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:385
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF mapToParent(qreal, qreal, qreal, qreal) const

/*
Maps the point point, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped coordinate. If the item has no parent, point will be mapped to the scene's coordinate system.

See also mapToItem(), mapToScene(), transform(), mapFromParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToParent_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:352
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToScene(const QPointF &) const

/*
Maps the point point, which is in this item's coordinate system, to the scene's coordinate system, and returns the mapped coordinate.

See also mapToItem(), mapToParent(), transform(), mapFromScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToScene(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:355
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToScene(const QRectF &) const

/*
Maps the point point, which is in this item's coordinate system, to the scene's coordinate system, and returns the mapped coordinate.

See also mapToItem(), mapToParent(), transform(), mapFromScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToScene_1(rect qtcore.QRectF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:361
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapToScene(const QPolygonF &) const

/*
Maps the point point, which is in this item's coordinate system, to the scene's coordinate system, and returns the mapped coordinate.

See also mapToItem(), mapToParent(), transform(), mapFromScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToScene_2(polygon qtgui.QPolygonF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:364
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapToScene(const QPainterPath &) const

/*
Maps the point point, which is in this item's coordinate system, to the scene's coordinate system, and returns the mapped coordinate.

See also mapToItem(), mapToParent(), transform(), mapFromScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToScene_3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:383
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QPointF mapToScene(qreal, qreal) const

/*
Maps the point point, which is in this item's coordinate system, to the scene's coordinate system, and returns the mapped coordinate.

See also mapToItem(), mapToParent(), transform(), mapFromScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToScene_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:386
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF mapToScene(qreal, qreal, qreal, qreal) const

/*
Maps the point point, which is in this item's coordinate system, to the scene's coordinate system, and returns the mapped coordinate.

See also mapToItem(), mapToParent(), transform(), mapFromScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapToScene_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:356
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectToItem(const QGraphicsItem *, const QRectF &) const

/*
Maps the rectangle rect, which is in this item's coordinate system, to item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

If item is 0, this function returns the same as mapRectToScene().

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectToItem(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:387
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QRectF mapRectToItem(const QGraphicsItem *, qreal, qreal, qreal, qreal) const

/*
Maps the rectangle rect, which is in this item's coordinate system, to item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

If item is 0, this function returns the same as mapRectToScene().

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectToItem_1(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:357
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectToParent(const QRectF &) const

/*
Maps the rectangle rect, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectToParent(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectToParentERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:388
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QRectF mapRectToParent(qreal, qreal, qreal, qreal) const

/*
Maps the rectangle rect, which is in this item's coordinate system, to its parent's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectToParent_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectToParentEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:358
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectToScene(const QRectF &) const

/*
Maps the rectangle rect, which is in this item's coordinate system, to the scene coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectToScene(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14mapRectToSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:389
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QRectF mapRectToScene(qreal, qreal, qreal, qreal) const

/*
Maps the rectangle rect, which is in this item's coordinate system, to the scene coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectToScene_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14mapRectToSceneEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:365
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromItem(const QGraphicsItem *, const QPointF &) const

/*
Maps the point point, which is in item's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapFromScene().

See also itemTransform(), mapFromParent(), mapFromScene(), transform(), mapToItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromItem(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg1 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:368
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapFromItem(const QGraphicsItem *, const QRectF &) const

/*
Maps the point point, which is in item's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapFromScene().

See also itemTransform(), mapFromParent(), mapFromScene(), transform(), mapToItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromItem_1(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, rect qtcore.QRectF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:374
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapFromItem(const QGraphicsItem *, const QPolygonF &) const

/*
Maps the point point, which is in item's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapFromScene().

See also itemTransform(), mapFromParent(), mapFromScene(), transform(), mapToItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromItem_2(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, polygon qtgui.QPolygonF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg1 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:377
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapFromItem(const QGraphicsItem *, const QPainterPath &) const

/*
Maps the point point, which is in item's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapFromScene().

See also itemTransform(), mapFromParent(), mapFromScene(), transform(), mapToItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromItem_3(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg1 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:390
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QPointF mapFromItem(const QGraphicsItem *, qreal, qreal) const

/*
Maps the point point, which is in item's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapFromScene().

See also itemTransform(), mapFromParent(), mapFromScene(), transform(), mapToItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromItem_4(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, x float64, y float64) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_dd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:393
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF mapFromItem(const QGraphicsItem *, qreal, qreal, qreal, qreal) const

/*
Maps the point point, which is in item's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

If item is 0, this function returns the same as mapFromScene().

See also itemTransform(), mapFromParent(), mapFromScene(), transform(), mapToItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromItem_5(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_dddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:366
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromParent(const QPointF &) const

/*
Maps the point point, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromScene(), transform(), mapToParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromParent(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:369
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapFromParent(const QRectF &) const

/*
Maps the point point, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromScene(), transform(), mapToParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromParent_1(rect qtcore.QRectF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:375
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapFromParent(const QPolygonF &) const

/*
Maps the point point, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromScene(), transform(), mapToParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromParent_2(polygon qtgui.QPolygonF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:378
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapFromParent(const QPainterPath &) const

/*
Maps the point point, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromScene(), transform(), mapToParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromParent_3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:391
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QPointF mapFromParent(qreal, qreal) const

/*
Maps the point point, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromScene(), transform(), mapToParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromParent_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:394
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF mapFromParent(qreal, qreal, qreal, qreal) const

/*
Maps the point point, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromScene(), transform(), mapToParent(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromParent_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:367
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromScene(const QPointF &) const

/*
Maps the point point, which is in this item's scene's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromParent(), transform(), mapToScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromScene(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:370
// index:1
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapFromScene(const QRectF &) const

/*
Maps the point point, which is in this item's scene's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromParent(), transform(), mapToScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromScene_1(rect qtcore.QRectF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:376
// index:2
// Public Visibility=Default Availability=Available
// [8] QPolygonF mapFromScene(const QPolygonF &) const

/*
Maps the point point, which is in this item's scene's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromParent(), transform(), mapToScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromScene_2(polygon qtgui.QPolygonF_ITF) *qtgui.QPolygonF /*123*/ {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:379
// index:3
// Public Visibility=Default Availability=Available
// [8] QPainterPath mapFromScene(const QPainterPath &) const

/*
Maps the point point, which is in this item's scene's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromParent(), transform(), mapToScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromScene_3(path qtgui.QPainterPath_ITF) *qtgui.QPainterPath /*123*/ {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:392
// index:4
// Public inline Visibility=Default Availability=Available
// [16] QPointF mapFromScene(qreal, qreal) const

/*
Maps the point point, which is in this item's scene's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromParent(), transform(), mapToScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromScene_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:395
// index:5
// Public inline Visibility=Default Availability=Available
// [8] QPolygonF mapFromScene(qreal, qreal, qreal, qreal) const

/*
Maps the point point, which is in this item's scene's coordinate system, to this item's coordinate system, and returns the mapped coordinate.

See also mapFromItem(), mapFromParent(), transform(), mapToScene(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapFromScene_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:371
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectFromItem(const QGraphicsItem *, const QRectF &) const

/*
Maps the rectangle rect, which is in item's coordinate system, to this item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

If item is 0, this function returns the same as mapRectFromScene().

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectFromItem(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:396
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QRectF mapRectFromItem(const QGraphicsItem *, qreal, qreal, qreal, qreal) const

/*
Maps the rectangle rect, which is in item's coordinate system, to this item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

If item is 0, this function returns the same as mapRectFromScene().

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectFromItem_1(item QGraphicsItem_ITF /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:372
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectFromParent(const QRectF &) const

/*
Maps the rectangle rect, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectFromParent(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17mapRectFromParentERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:397
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QRectF mapRectFromParent(qreal, qreal, qreal, qreal) const

/*
Maps the rectangle rect, which is in this item's parent's coordinate system, to this item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectFromParent_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17mapRectFromParentEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:373
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectFromScene(const QRectF &) const

/*
Maps the rectangle rect, which is in scene coordinates, to this item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectFromScene(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:398
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QRectF mapRectFromScene(qreal, qreal, qreal, qreal) const

/*
Maps the rectangle rect, which is in scene coordinates, to this item's coordinate system, and returns the mapped rectangle as a new rectangle (i.e., the bounding rectangle of the resulting polygon).

This function was introduced in  Qt 4.5.

See also itemTransform(), mapToParent(), mapToScene(), mapFromItem(), and The Graphics View Coordinate System.
*/
func (this *QGraphicsItem) MapRectFromScene_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16mapRectFromSceneEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:400
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QGraphicsItem *) const

/*
Returns true if this item is an ancestor of child (i.e., if this item is child's parent, or one of child's parent's ancestors).

See also parentItem().
*/
func (this *QGraphicsItem) IsAncestorOf(child QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if child != nil && child.QGraphicsItem_PTR() != nil {
		convArg0 = child.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12isAncestorOfEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:401
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * commonAncestorItem(const QGraphicsItem *) const

/*
Returns the closest common ancestor item of this item and other, or 0 if either other is 0, or there is no common ancestor.

This function was introduced in  Qt 4.4.

See also isAncestorOf().
*/
func (this *QGraphicsItem) CommonAncestorItem(other QGraphicsItem_ITF /*777 const QGraphicsItem **/) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGraphicsItem_PTR() != nil {
		convArg0 = other.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem18commonAncestorItemEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:402
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUnderMouse() const

/*
Returns true if this item is currently under the mouse cursor in one of the views; otherwise, false is returned.

This function was introduced in  Qt 4.4.

See also QGraphicsScene::views() and QCursor::pos().
*/
func (this *QGraphicsItem) IsUnderMouse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12isUnderMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:405
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*
Returns this item's custom data for the key key as a QVariant.

Custom item data is useful for storing arbitrary properties in any item. Example:


  static const int ObjectName = 0;

  QGraphicsItem *item = scene.itemAt(100, 50);
  if (item->data(ObjectName).toString().isEmpty()) {
      if (qgraphicsitem_cast<ButtonItem *>(item))
          item->setData(ObjectName, "Button");
  }



Qt does not use this feature for storing data; it is provided solely for the convenience of the user.

See also setData().
*/
func (this *QGraphicsItem) Data(key int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:406
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setData(int, const QVariant &)

/*
Sets this item's custom data for the key key to value.

Custom item data is useful for storing arbitrary properties for any item. Qt does not use this feature for storing data; it is provided solely for the convenience of the user.

See also data().
*/
func (this *QGraphicsItem) SetData(key int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem7setDataEiRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), key, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:408
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::InputMethodHints inputMethodHints() const

/*
Returns the current input method hints of this item.

Input method hints are only relevant for input items. The hints are used by the input method to indicate how it should operate. For example, if the Qt::ImhNumbersOnly flag is set, the input method may change its visual components to reflect that only numbers can be entered.

The effect may vary between input method implementations.

This function was introduced in  Qt 4.6.

See also setInputMethodHints() and inputMethodQuery().
*/
func (this *QGraphicsItem) InputMethodHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16inputMethodHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:409
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMethodHints(Qt::InputMethodHints)

/*
Sets the current input method hints of this item to hints.

This function was introduced in  Qt 4.6.

See also inputMethodHints() and inputMethodQuery().
*/
func (this *QGraphicsItem) SetInputMethodHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem19setInputMethodHintsE6QFlagsIN2Qt15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:415
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const

/*
Returns the type of an item as an int. All standard graphicsitem classes are associated with a unique value; see QGraphicsItem::Type. This type information is used by qgraphicsitem_cast() to distinguish between types.

The default implementation (in QGraphicsItem) returns UserType.

To enable use of qgraphicsitem_cast() with a custom item, reimplement this function and declare a Type enum value equal to your custom item's type. Custom items must return a value larger than or equal to UserType (65536).

For example:


  class CustomItem : public QGraphicsItem
  {
  public:
     enum { Type = UserType + 1 };

     int type() const
     {
         // Enable the use of qgraphicsitem_cast with this item.
         return Type;
     }
     ...
  };



See also UserType.
*/
func (this *QGraphicsItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:417
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installSceneEventFilter(QGraphicsItem *)

/*
Installs an event filter for this item on filterItem, causing all events for this item to first pass through filterItem's sceneEventFilter() function.

To filter another item's events, install this item as an event filter for the other item. Example:


  QGraphicsScene scene;
  QGraphicsEllipseItem *ellipse = scene.addEllipse(QRectF(-10, -10, 20, 20));
  QGraphicsLineItem *line = scene.addLine(QLineF(-10, -10, 20, 20));

  line->installSceneEventFilter(ellipse);
  // line's events are filtered by ellipse's sceneEventFilter() function.

  ellipse->installSceneEventFilter(line);
  // ellipse's events are filtered by line's sceneEventFilter() function.



An item can only filter events for other items in the same scene. Also, an item cannot filter its own events; instead, you can reimplement sceneEvent() directly.

Items must belong to a scene for scene event filters to be installed and used.

See also removeSceneEventFilter(), sceneEventFilter(), and sceneEvent().
*/
func (this *QGraphicsItem) InstallSceneEventFilter(filterItem QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if filterItem != nil && filterItem.QGraphicsItem_PTR() != nil {
		convArg0 = filterItem.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem23installSceneEventFilterEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:418
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeSceneEventFilter(QGraphicsItem *)

/*
Removes an event filter on this item from filterItem.

See also installSceneEventFilter().
*/
func (this *QGraphicsItem) RemoveSceneEventFilter(filterItem QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if filterItem != nil && filterItem.QGraphicsItem_PTR() != nil {
		convArg0 = filterItem.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem22removeSceneEventFilterEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:421
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateMicroFocus()

/*
Updates the item's micro focus.

This function was introduced in  Qt 4.7.

See also QInputMethod.
*/
func (this *QGraphicsItem) UpdateMicroFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem16updateMicroFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:422
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool sceneEventFilter(QGraphicsItem *, QEvent *)

/*
Filters events for the item watched. event is the filtered event.

Reimplementing this function in a subclass makes it possible for the item to be used as an event filter for other items, intercepting all the events sent to those items before they are able to respond.

Reimplementations must return true to prevent further processing of a given event, ensuring that it will not be delivered to the watched item, or return false to indicate that the event should be propagated further by the event system.

See also installSceneEventFilter().
*/
func (this *QGraphicsItem) SceneEventFilter(watched QGraphicsItem_ITF /*777 QGraphicsItem **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if watched != nil && watched.QGraphicsItem_PTR() != nil {
		convArg0 = watched.QGraphicsItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:423
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool sceneEvent(QEvent *)

/*
This virtual function receives events to this item. Reimplement this function to intercept events before they are dispatched to the specialized event handlers contextMenuEvent(), focusInEvent(), focusOutEvent(), hoverEnterEvent(), hoverMoveEvent(), hoverLeaveEvent(), keyPressEvent(), keyReleaseEvent(), mousePressEvent(), mouseReleaseEvent(), mouseMoveEvent(), and mouseDoubleClickEvent().

Returns true if the event was recognized and handled; otherwise, (e.g., if the event type was not recognized,) false is returned.

event is the intercepted event.
*/
func (this *QGraphicsItem) SceneEvent(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10sceneEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:424
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QGraphicsSceneContextMenuEvent *)

/*
This event handler can be reimplemented in a subclass to process context menu events. The event parameter contains details about the event to be handled.

If you ignore the event (i.e., by calling QEvent::ignore()), event will propagate to any item beneath this item. If no items accept the event, it will be ignored by the scene and propagate to the view.

It's common to open a QMenu in response to receiving a context menu event. Example:


  void CustomItem::contextMenuEvent(QGraphicsSceneContextMenuEvent *event)
  {
      QMenu menu;
      QAction *removeAction = menu.addAction("Remove");
      QAction *markAction = menu.addAction("Mark");
      QAction *selectedAction = menu.exec(event->screenPos());
      // ...
  }



The default implementation ignores the event.

See also sceneEvent().
*/
func (this *QGraphicsItem) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF /*777 QGraphicsSceneContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneContextMenuEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:425
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drag enter events for this item. Drag enter events are generated as the cursor enters the item's area.

By accepting the event (i.e., by calling QEvent::accept()), the item will accept drop events, in addition to receiving drag move and drag leave. Otherwise, the event will be ignored and propagate to the item beneath. If the event is accepted, the item will receive a drag move event before control goes back to the event loop.

A common implementation of dragEnterEvent accepts or ignores event depending on the associated mime data in event. Example:


  CustomItem::CustomItem()
  {
      setAcceptDrops(true);
      ...
  }

  void CustomItem::dragEnterEvent(QGraphicsSceneDragDropEvent *event)
  {
      event->setAccepted(event->mimeData()->hasFormat("text/plain"));
  }



Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dropEvent(), dragMoveEvent(), and dragLeaveEvent().
*/
func (this *QGraphicsItem) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:426
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drag leave events for this item. Drag leave events are generated as the cursor leaves the item's area. Most often you will not need to reimplement this function, but it can be useful for resetting state in your item (e.g., highlighting).

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dragEnterEvent(), dropEvent(), and dragMoveEvent().
*/
func (this *QGraphicsItem) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:427
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drag move events for this item. Drag move events are generated as the cursor moves around inside the item's area. Most often you will not need to reimplement this function; it is used to indicate that only parts of the item can accept drops.

Calling QEvent::ignore() or QEvent::accept() on event toggles whether or not the item will accept drops at the position from the event. By default, event is accepted, indicating that the item allows drops at the specified position.

Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dropEvent(), dragEnterEvent(), and dragLeaveEvent().
*/
func (this *QGraphicsItem) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:428
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QGraphicsSceneDragDropEvent *)

/*
This event handler, for event event, can be reimplemented to receive drop events for this item. Items can only receive drop events if the last drag move event was accepted.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

Items do not receive drag and drop events by default; to enable this feature, call setAcceptDrops(true).

The default implementation does nothing.

See also dragEnterEvent(), dragMoveEvent(), and dragLeaveEvent().
*/
func (this *QGraphicsItem) DropEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:429
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
This event handler, for event event, can be reimplemented to receive focus in events for this item. The default implementation calls ensureVisible().

See also focusOutEvent(), sceneEvent(), and setFocus().
*/
func (this *QGraphicsItem) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:430
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
This event handler, for event event, can be reimplemented to receive focus out events for this item. The default implementation does nothing.

See also focusInEvent(), sceneEvent(), and setFocus().
*/
func (this *QGraphicsItem) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:431
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverEnterEvent(QGraphicsSceneHoverEvent *)

/*
This event handler, for event event, can be reimplemented to receive hover enter events for this item. The default implementation calls update(); otherwise it does nothing.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

See also hoverMoveEvent(), hoverLeaveEvent(), sceneEvent(), and setAcceptHoverEvents().
*/
func (this *QGraphicsItem) HoverEnterEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:432
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QGraphicsSceneHoverEvent *)

/*
This event handler, for event event, can be reimplemented to receive hover move events for this item. The default implementation does nothing.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

See also hoverEnterEvent(), hoverLeaveEvent(), sceneEvent(), and setAcceptHoverEvents().
*/
func (this *QGraphicsItem) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:433
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QGraphicsSceneHoverEvent *)

/*
This event handler, for event event, can be reimplemented to receive hover leave events for this item. The default implementation calls update(); otherwise it does nothing.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

See also hoverEnterEvent(), hoverMoveEvent(), sceneEvent(), and setAcceptHoverEvents().
*/
func (this *QGraphicsItem) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:434
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
This event handler, for event event, can be reimplemented to receive key press events for this item. The default implementation ignores the event. If you reimplement this handler, the event will by default be accepted.

Note that key events are only received for items that set the ItemIsFocusable flag, and that have keyboard input focus.

See also keyReleaseEvent(), setFocus(), QGraphicsScene::setFocusItem(), and sceneEvent().
*/
func (this *QGraphicsItem) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:435
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
This event handler, for event event, can be reimplemented to receive key release events for this item. The default implementation ignores the event. If you reimplement this handler, the event will by default be accepted.

Note that key events are only received for items that set the ItemIsFocusable flag, and that have keyboard input focus.

See also keyPressEvent(), setFocus(), QGraphicsScene::setFocusItem(), and sceneEvent().
*/
func (this *QGraphicsItem) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:436
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse press events for this item. Mouse press events are only delivered to items that accept the mouse button that is pressed. By default, an item accepts all mouse buttons, but you can change this by calling setAcceptedMouseButtons().

The mouse press event decides which item should become the mouse grabber (see QGraphicsScene::mouseGrabberItem()). If you do not reimplement this function, the press event will propagate to any topmost item beneath this item, and no other mouse events will be delivered to this item.

If you do reimplement this function, event will by default be accepted (see QEvent::accept()), and this item is then the mouse grabber. This allows the item to receive future move, release and doubleclick events. If you call QEvent::ignore() on event, this item will lose the mouse grab, and event will propagate to any topmost item beneath. No further mouse events will be delivered to this item unless a new mouse press event is received.

The default implementation handles basic item interaction, such as selection and moving. If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mousePressEvent() in your reimplementation.

The event is QEvent::ignore()d for items that are neither movable nor selectable.

See also mouseMoveEvent(), mouseReleaseEvent(), mouseDoubleClickEvent(), and sceneEvent().
*/
func (this *QGraphicsItem) MousePressEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:437
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse move events for this item. If you do receive this event, you can be certain that this item also received a mouse press event, and that this item is the current mouse grabber.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

The default implementation handles basic item interaction, such as selection and moving. If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mouseMoveEvent() in your reimplementation.

Please note that mousePressEvent() decides which graphics item it is that receives mouse events. See the mousePressEvent() description for details.

See also mousePressEvent(), mouseReleaseEvent(), mouseDoubleClickEvent(), and sceneEvent().
*/
func (this *QGraphicsItem) MouseMoveEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:438
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse release events for this item.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

The default implementation handles basic item interaction, such as selection and moving. If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mouseReleaseEvent() in your reimplementation.

Please note that mousePressEvent() decides which graphics item it is that receives mouse events. See the mousePressEvent() description for details.

See also mousePressEvent(), mouseMoveEvent(), mouseDoubleClickEvent(), and sceneEvent().
*/
func (this *QGraphicsItem) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:439
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)

/*
This event handler, for event event, can be reimplemented to receive mouse doubleclick events for this item.

When doubleclicking an item, the item will first receive a mouse press event, followed by a release event (i.e., a click), then a doubleclick event, and finally a release event.

Calling QEvent::ignore() or QEvent::accept() on event has no effect.

The default implementation calls mousePressEvent(). If you want to keep the base implementation when reimplementing this function, call QGraphicsItem::mouseDoubleClickEvent() in your reimplementation.

Note that an item will not receive double click events if it is neither selectable nor movable (single mouse clicks are ignored in this case, and that stops the generation of double clicks).

See also mousePressEvent(), mouseMoveEvent(), mouseReleaseEvent(), and sceneEvent().
*/
func (this *QGraphicsItem) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:440
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QGraphicsSceneWheelEvent *)

/*
This event handler, for event event, can be reimplemented to receive wheel events for this item. If you reimplement this function, event will be accepted by default.

If you ignore the event, (i.e., by calling QEvent::ignore(),) it will propagate to any item beneath this item. If no items accept the event, it will be ignored by the scene, and propagate to the view (e.g., the view's vertical scroll bar).

The default implementation ignores the event.

See also sceneEvent().
*/
func (this *QGraphicsItem) WheelEvent(event QGraphicsSceneWheelEvent_ITF /*777 QGraphicsSceneWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneWheelEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:441
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
This event handler, for event event, can be reimplemented to receive input method events for this item. The default implementation ignores the event.

See also inputMethodQuery() and sceneEvent().
*/
func (this *QGraphicsItem) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QInputMethodEvent_PTR() != nil {
		convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:442
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
This method is only relevant for input items. It is used by the input method to query a set of properties of the item to be able to support complex input method operations, such as support for surrounding text and reconversions. query specifies which property is queried.

See also inputMethodEvent() and QInputMethodEvent.
*/
func (this *QGraphicsItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:444
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant itemChange(QGraphicsItem::GraphicsItemChange, const QVariant &)

/*
This virtual function is called by QGraphicsItem to notify custom items that some part of the item's state changes. By reimplementing this function, you can react to a change, and in some cases (depending on change), adjustments can be made.

change is the parameter of the item that is changing. value is the new value; the type of the value depends on change.

Example:


  QVariant Component::itemChange(GraphicsItemChange change, const QVariant &value)
  {
      if (change == ItemPositionChange && scene()) {
          // value is the new position.
          QPointF newPos = value.toPointF();
          QRectF rect = scene()->sceneRect();
          if (!rect.contains(newPos)) {
              // Keep the item inside the scene rect.
              newPos.setX(qMin(rect.right(), qMax(newPos.x(), rect.left())));
              newPos.setY(qMin(rect.bottom(), qMax(newPos.y(), rect.top())));
              return newPos;
          }
      }
      return QGraphicsItem::itemChange(change, value);
  }



The default implementation does nothing, and returns value.

Note: Certain QGraphicsItem functions cannot be called in a reimplementation of this function; see the GraphicsItemChange documentation for details.

See also GraphicsItemChange.
*/
func (this *QGraphicsItem) ItemChange(change int, value qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:449
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(QGraphicsItem::Extension) const

/*

 */
func (this *QGraphicsItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:450
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(QGraphicsItem::Extension, const QVariant &)

/*

 */
func (this *QGraphicsItem) SetExtension(extension int, variant qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg1 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:451
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &) const

/*

 */
func (this *QGraphicsItem) Extension(variant qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:457
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addToIndex()

/*

 */
func (this *QGraphicsItem) AddToIndex() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10addToIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:458
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void removeFromIndex()

/*

 */
func (this *QGraphicsItem) RemoveFromIndex() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem15removeFromIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:459
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void prepareGeometryChange()

/*
Prepares the item for a geometry change. Call this function before changing the bounding rect of an item to keep QGraphicsScene's index up to date.

prepareGeometryChange() will call update() if this is necessary.

Example:


  void CircleItem::setRadius(qreal newRadius)
  {
      if (radius != newRadius) {
          prepareGeometryChange();
          radius = newRadius;
      }
  }



See also boundingRect().
*/
func (this *QGraphicsItem) PrepareGeometryChange() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem21prepareGeometryChangeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QGraphicsItem__GraphicsItemFlag = int

//
const QGraphicsItem__ItemIsMovable QGraphicsItem__GraphicsItemFlag = 1

//
const QGraphicsItem__ItemIsSelectable QGraphicsItem__GraphicsItemFlag = 2

//
const QGraphicsItem__ItemIsFocusable QGraphicsItem__GraphicsItemFlag = 4

//
const QGraphicsItem__ItemClipsToShape QGraphicsItem__GraphicsItemFlag = 8

//
const QGraphicsItem__ItemClipsChildrenToShape QGraphicsItem__GraphicsItemFlag = 16

//
const QGraphicsItem__ItemIgnoresTransformations QGraphicsItem__GraphicsItemFlag = 32

//
const QGraphicsItem__ItemIgnoresParentOpacity QGraphicsItem__GraphicsItemFlag = 64

//
const QGraphicsItem__ItemDoesntPropagateOpacityToChildren QGraphicsItem__GraphicsItemFlag = 128

//
const QGraphicsItem__ItemStacksBehindParent QGraphicsItem__GraphicsItemFlag = 256

//
const QGraphicsItem__ItemUsesExtendedStyleOption QGraphicsItem__GraphicsItemFlag = 512

//
const QGraphicsItem__ItemHasNoContents QGraphicsItem__GraphicsItemFlag = 1024

//
const QGraphicsItem__ItemSendsGeometryChanges QGraphicsItem__GraphicsItemFlag = 2048

//
const QGraphicsItem__ItemAcceptsInputMethod QGraphicsItem__GraphicsItemFlag = 4096

//
const QGraphicsItem__ItemNegativeZStacksBehindParent QGraphicsItem__GraphicsItemFlag = 8192

//
const QGraphicsItem__ItemIsPanel QGraphicsItem__GraphicsItemFlag = 16384

//
const QGraphicsItem__ItemIsFocusScope QGraphicsItem__GraphicsItemFlag = 32768

//
const QGraphicsItem__ItemSendsScenePositionChanges QGraphicsItem__GraphicsItemFlag = 65536

//
const QGraphicsItem__ItemStopsClickFocusPropagation QGraphicsItem__GraphicsItemFlag = 131072

//
const QGraphicsItem__ItemStopsFocusHandling QGraphicsItem__GraphicsItemFlag = 262144

//
const QGraphicsItem__ItemContainsChildrenInShape QGraphicsItem__GraphicsItemFlag = 524288

func (this *QGraphicsItem) GraphicsItemFlagItemName(val int) string {
	switch val {
	case QGraphicsItem__ItemIsMovable: // 1
		return "ItemIsMovable"
	case QGraphicsItem__ItemIsSelectable: // 2
		return "ItemIsSelectable"
	case QGraphicsItem__ItemIsFocusable: // 4
		return "ItemIsFocusable"
	case QGraphicsItem__ItemClipsToShape: // 8
		return "ItemClipsToShape"
	case QGraphicsItem__ItemClipsChildrenToShape: // 16
		return "ItemClipsChildrenToShape"
	case QGraphicsItem__ItemIgnoresTransformations: // 32
		return "ItemIgnoresTransformations"
	case QGraphicsItem__ItemIgnoresParentOpacity: // 64
		return "ItemIgnoresParentOpacity"
	case QGraphicsItem__ItemDoesntPropagateOpacityToChildren: // 128
		return "ItemDoesntPropagateOpacityToChildren"
	case QGraphicsItem__ItemStacksBehindParent: // 256
		return "ItemStacksBehindParent"
	case QGraphicsItem__ItemUsesExtendedStyleOption: // 512
		return "ItemUsesExtendedStyleOption"
	case QGraphicsItem__ItemHasNoContents: // 1024
		return "ItemHasNoContents"
	case QGraphicsItem__ItemSendsGeometryChanges: // 2048
		return "ItemSendsGeometryChanges"
	case QGraphicsItem__ItemAcceptsInputMethod: // 4096
		return "ItemAcceptsInputMethod"
	case QGraphicsItem__ItemNegativeZStacksBehindParent: // 8192
		return "ItemNegativeZStacksBehindParent"
	case QGraphicsItem__ItemIsPanel: // 16384
		return "ItemIsPanel"
	case QGraphicsItem__ItemIsFocusScope: // 32768
		return "ItemIsFocusScope"
	case QGraphicsItem__ItemSendsScenePositionChanges: // 65536
		return "ItemSendsScenePositionChanges"
	case QGraphicsItem__ItemStopsClickFocusPropagation: // 131072
		return "ItemStopsClickFocusPropagation"
	case QGraphicsItem__ItemStopsFocusHandling: // 262144
		return "ItemStopsFocusHandling"
	case QGraphicsItem__ItemContainsChildrenInShape: // 524288
		return "ItemContainsChildrenInShape"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_GraphicsItemFlagItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.GraphicsItemFlagItemName(val)
}

/*
This enum describes the state changes that are notified by QGraphicsItem::itemChange(). The notifications are sent as the state changes, and in some cases, adjustments can be made (see the documentation for each change for details).

Note: Be careful with calling functions on the QGraphicsItem itself inside itemChange(), as certain function calls can lead to unwanted recursion. For example, you cannot call setPos() in itemChange() on an ItemPositionChange notification, as the setPos() function will again call itemChange(ItemPositionChange). Instead, you can return the new, adjusted position from itemChange().


*/
type QGraphicsItem__GraphicsItemChange = int

// The item's position changes. This notification is sent if the ItemSendsGeometryChanges flag is enabled, and when the item's local position changes, relative to its parent (i.e., as a result of calling setPos() or moveBy()). The value argument is the new position (i.e., a QPointF). You can call pos() to get the original position. Do not call setPos() or moveBy() in itemChange() as this notification is delivered; instead, you can return the new, adjusted position from itemChange(). After this notification, QGraphicsItem immediately sends the ItemPositionHasChanged notification if the position changed.
const QGraphicsItem__ItemPositionChange QGraphicsItem__GraphicsItemChange = 0

// The item's affine transformation matrix is changing. This value is obsolete; you can use ItemTransformChange instead.
const QGraphicsItem__ItemMatrixChange QGraphicsItem__GraphicsItemChange = 1

// The item's visible state changes. If the item is presently visible, it will become invisible, and vice verca. The value argument is the new visible state (i.e., true or false). Do not call setVisible() in itemChange() as this notification is delivered; instead, you can return the new visible state from itemChange().
const QGraphicsItem__ItemVisibleChange QGraphicsItem__GraphicsItemChange = 2

// The item's enabled state changes. If the item is presently enabled, it will become disabled, and vice verca. The value argument is the new enabled state (i.e., true or false). Do not call setEnabled() in itemChange() as this notification is delivered. Instead, you can return the new state from itemChange().
const QGraphicsItem__ItemEnabledChange QGraphicsItem__GraphicsItemChange = 3

// The item's selected state changes. If the item is presently selected, it will become unselected, and vice verca. The value argument is the new selected state (i.e., true or false). Do not call setSelected() in itemChange() as this notification is delivered; instead, you can return the new selected state from itemChange().
const QGraphicsItem__ItemSelectedChange QGraphicsItem__GraphicsItemChange = 4

// The item's parent changes. The value argument is the new parent item (i.e., a QGraphicsItem pointer). Do not call setParentItem() in itemChange() as this notification is delivered; instead, you can return the new parent from itemChange().
const QGraphicsItem__ItemParentChange QGraphicsItem__GraphicsItemChange = 5

// A child is added to this item. The value argument is the new child item (i.e., a QGraphicsItem pointer). Do not pass this item to any item's setParentItem() function as this notification is delivered. The return value is unused; you cannot adjust anything in this notification. Note that the new child might not be fully constructed when this notification is sent; calling pure virtual functions on the child can lead to a crash.
const QGraphicsItem__ItemChildAddedChange QGraphicsItem__GraphicsItemChange = 6

// A child is removed from this item. The value argument is the child item that is about to be removed (i.e., a QGraphicsItem pointer). The return value is unused; you cannot adjust anything in this notification.
const QGraphicsItem__ItemChildRemovedChange QGraphicsItem__GraphicsItemChange = 7

// The item's transformation matrix changes. This notification is sent if the ItemSendsGeometryChanges flag is enabled, and when the item's local transformation matrix changes (i.e., as a result of calling setTransform(). The value argument is the new matrix (i.e., a QTransform); to get the old matrix, call transform(). Do not call setTransform() or set any of the transformation properties in itemChange() as this notification is delivered; instead, you can return the new matrix from itemChange(). This notification is not sent if you change the transformation properties.
const QGraphicsItem__ItemTransformChange QGraphicsItem__GraphicsItemChange = 8

// The item's position has changed. This notification is sent if the ItemSendsGeometryChanges flag is enabled, and after the item's local position, relative to its parent, has changed. The value argument is the new position (the same as pos()), and QGraphicsItem ignores the return value for this notification (i.e., a read-only notification).
const QGraphicsItem__ItemPositionHasChanged QGraphicsItem__GraphicsItemChange = 9

//
const QGraphicsItem__ItemTransformHasChanged QGraphicsItem__GraphicsItemChange = 10

//
const QGraphicsItem__ItemSceneChange QGraphicsItem__GraphicsItemChange = 11

//
const QGraphicsItem__ItemVisibleHasChanged QGraphicsItem__GraphicsItemChange = 12

//
const QGraphicsItem__ItemEnabledHasChanged QGraphicsItem__GraphicsItemChange = 13

//
const QGraphicsItem__ItemSelectedHasChanged QGraphicsItem__GraphicsItemChange = 14

//
const QGraphicsItem__ItemParentHasChanged QGraphicsItem__GraphicsItemChange = 15

//
const QGraphicsItem__ItemSceneHasChanged QGraphicsItem__GraphicsItemChange = 16

//
const QGraphicsItem__ItemCursorChange QGraphicsItem__GraphicsItemChange = 17

//
const QGraphicsItem__ItemCursorHasChanged QGraphicsItem__GraphicsItemChange = 18

//
const QGraphicsItem__ItemToolTipChange QGraphicsItem__GraphicsItemChange = 19

//
const QGraphicsItem__ItemToolTipHasChanged QGraphicsItem__GraphicsItemChange = 20

//
const QGraphicsItem__ItemFlagsChange QGraphicsItem__GraphicsItemChange = 21

//
const QGraphicsItem__ItemFlagsHaveChanged QGraphicsItem__GraphicsItemChange = 22

//
const QGraphicsItem__ItemZValueChange QGraphicsItem__GraphicsItemChange = 23

//
const QGraphicsItem__ItemZValueHasChanged QGraphicsItem__GraphicsItemChange = 24

//
const QGraphicsItem__ItemOpacityChange QGraphicsItem__GraphicsItemChange = 25

//
const QGraphicsItem__ItemOpacityHasChanged QGraphicsItem__GraphicsItemChange = 26

//
const QGraphicsItem__ItemScenePositionHasChanged QGraphicsItem__GraphicsItemChange = 27

//
const QGraphicsItem__ItemRotationChange QGraphicsItem__GraphicsItemChange = 28

//
const QGraphicsItem__ItemRotationHasChanged QGraphicsItem__GraphicsItemChange = 29

//
const QGraphicsItem__ItemScaleChange QGraphicsItem__GraphicsItemChange = 30

//
const QGraphicsItem__ItemScaleHasChanged QGraphicsItem__GraphicsItemChange = 31

//
const QGraphicsItem__ItemTransformOriginPointChange QGraphicsItem__GraphicsItemChange = 32

//
const QGraphicsItem__ItemTransformOriginPointHasChanged QGraphicsItem__GraphicsItemChange = 33

func (this *QGraphicsItem) GraphicsItemChangeItemName(val int) string {
	switch val {
	case QGraphicsItem__ItemPositionChange: // 0
		return "ItemPositionChange"
	case QGraphicsItem__ItemMatrixChange: // 1
		return "ItemMatrixChange"
	case QGraphicsItem__ItemVisibleChange: // 2
		return "ItemVisibleChange"
	case QGraphicsItem__ItemEnabledChange: // 3
		return "ItemEnabledChange"
	case QGraphicsItem__ItemSelectedChange: // 4
		return "ItemSelectedChange"
	case QGraphicsItem__ItemParentChange: // 5
		return "ItemParentChange"
	case QGraphicsItem__ItemChildAddedChange: // 6
		return "ItemChildAddedChange"
	case QGraphicsItem__ItemChildRemovedChange: // 7
		return "ItemChildRemovedChange"
	case QGraphicsItem__ItemTransformChange: // 8
		return "ItemTransformChange"
	case QGraphicsItem__ItemPositionHasChanged: // 9
		return "ItemPositionHasChanged"
	case QGraphicsItem__ItemTransformHasChanged: // 10
		return "ItemTransformHasChanged"
	case QGraphicsItem__ItemSceneChange: // 11
		return "ItemSceneChange"
	case QGraphicsItem__ItemVisibleHasChanged: // 12
		return "ItemVisibleHasChanged"
	case QGraphicsItem__ItemEnabledHasChanged: // 13
		return "ItemEnabledHasChanged"
	case QGraphicsItem__ItemSelectedHasChanged: // 14
		return "ItemSelectedHasChanged"
	case QGraphicsItem__ItemParentHasChanged: // 15
		return "ItemParentHasChanged"
	case QGraphicsItem__ItemSceneHasChanged: // 16
		return "ItemSceneHasChanged"
	case QGraphicsItem__ItemCursorChange: // 17
		return "ItemCursorChange"
	case QGraphicsItem__ItemCursorHasChanged: // 18
		return "ItemCursorHasChanged"
	case QGraphicsItem__ItemToolTipChange: // 19
		return "ItemToolTipChange"
	case QGraphicsItem__ItemToolTipHasChanged: // 20
		return "ItemToolTipHasChanged"
	case QGraphicsItem__ItemFlagsChange: // 21
		return "ItemFlagsChange"
	case QGraphicsItem__ItemFlagsHaveChanged: // 22
		return "ItemFlagsHaveChanged"
	case QGraphicsItem__ItemZValueChange: // 23
		return "ItemZValueChange"
	case QGraphicsItem__ItemZValueHasChanged: // 24
		return "ItemZValueHasChanged"
	case QGraphicsItem__ItemOpacityChange: // 25
		return "ItemOpacityChange"
	case QGraphicsItem__ItemOpacityHasChanged: // 26
		return "ItemOpacityHasChanged"
	case QGraphicsItem__ItemScenePositionHasChanged: // 27
		return "ItemScenePositionHasChanged"
	case QGraphicsItem__ItemRotationChange: // 28
		return "ItemRotationChange"
	case QGraphicsItem__ItemRotationHasChanged: // 29
		return "ItemRotationHasChanged"
	case QGraphicsItem__ItemScaleChange: // 30
		return "ItemScaleChange"
	case QGraphicsItem__ItemScaleHasChanged: // 31
		return "ItemScaleHasChanged"
	case QGraphicsItem__ItemTransformOriginPointChange: // 32
		return "ItemTransformOriginPointChange"
	case QGraphicsItem__ItemTransformOriginPointHasChanged: // 33
		return "ItemTransformOriginPointHasChanged"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_GraphicsItemChangeItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.GraphicsItemChangeItemName(val)
}

/*
This enum describes QGraphicsItem's cache modes. Caching is used to speed up rendering by allocating and rendering to an off-screen pixel buffer, which can be reused when the item requires redrawing. For some paint devices, the cache is stored directly in graphics memory, which makes rendering very quick.



This enum was introduced or modified in  Qt 4.4.

See also QGraphicsItem::setCacheMode().

*/
type QGraphicsItem__CacheMode = int

// The default; all item caching is disabled. QGraphicsItem::paint() is called every time the item needs redrawing.
const QGraphicsItem__NoCache QGraphicsItem__CacheMode = 0

// Caching is enabled for the item's logical (local) coordinate system. QGraphicsItem creates an off-screen pixel buffer with a configurable size / resolution that you can pass to QGraphicsItem::setCacheMode(). Rendering quality will typically degrade, depending on the resolution of the cache and the item transformation. The first time the item is redrawn, it will render itself into the cache, and the cache is then reused for every subsequent expose. The cache is also reused as the item is transformed. To adjust the resolution of the cache, you can call setCacheMode() again.
const QGraphicsItem__ItemCoordinateCache QGraphicsItem__CacheMode = 1

// Caching is enabled at the paint device level, in device coordinates. This mode is for items that can move, but are not rotated, scaled or sheared. If the item is transformed directly or indirectly, the cache will be regenerated automatically. Unlike ItemCoordinateCacheMode, DeviceCoordinateCache always renders at maximum quality.
const QGraphicsItem__DeviceCoordinateCache QGraphicsItem__CacheMode = 2

func (this *QGraphicsItem) CacheModeItemName(val int) string {
	switch val {
	case QGraphicsItem__NoCache: // 0
		return "NoCache"
	case QGraphicsItem__ItemCoordinateCache: // 1
		return "ItemCoordinateCache"
	case QGraphicsItem__DeviceCoordinateCache: // 2
		return "DeviceCoordinateCache"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_CacheModeItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.CacheModeItemName(val)
}

/*
This enum specifies the behavior of a modal panel. A modal panel is one that blocks input to other panels. Note that items that are children of a modal panel are not blocked.

The values are:



This enum was introduced or modified in  Qt 4.6.

See also QGraphicsItem::setPanelModality(), QGraphicsItem::panelModality(), and QGraphicsItem::ItemIsPanel.

*/
type QGraphicsItem__PanelModality = int

// The panel is not modal and does not block input to other panels. This is the default value for panels.
const QGraphicsItem__NonModal QGraphicsItem__PanelModality = 0

// The panel is modal to a single item hierarchy and blocks input to its parent pane, all grandparent panels, and all siblings of its parent and grandparent panels.
const QGraphicsItem__PanelModal QGraphicsItem__PanelModality = 1

// The window is modal to the entire scene and blocks input to all panels.
const QGraphicsItem__SceneModal QGraphicsItem__PanelModality = 2

func (this *QGraphicsItem) PanelModalityItemName(val int) string {
	switch val {
	case QGraphicsItem__NonModal: // 0
		return "NonModal"
	case QGraphicsItem__PanelModal: // 1
		return "PanelModal"
	case QGraphicsItem__SceneModal: // 2
		return "SceneModal"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_PanelModalityItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.PanelModalityItemName(val)
}

/*


 */
type QGraphicsItem__ = int

//
const QGraphicsItem__Type QGraphicsItem__ = 1

//
const QGraphicsItem__UserType QGraphicsItem__ = 65536

func (this *QGraphicsItem) ItemName(val int) string {
	switch val {
	case QGraphicsItem__Type: // 1
		return "Type"
	case QGraphicsItem__UserType: // 65536
		return "UserType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_ItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.ItemName(val)
}

/*


 */
type QGraphicsItem__Extension = int

//
const QGraphicsItem__UserExtension QGraphicsItem__Extension = -2147483648

func (this *QGraphicsItem) ExtensionItemName(val int) string {
	switch val {
	case QGraphicsItem__UserExtension: // -2147483648
		return "UserExtension"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QGraphicsItem_ExtensionItemName(val int) string {
	var nilthis *QGraphicsItem
	return nilthis.ExtensionItemName(val)
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
