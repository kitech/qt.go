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

// bool sceneEventFilter(class QGraphicsItem *, class QEvent *)
func (this *QGraphicsItem) InheritSceneEventFilter(f func(watched *QGraphicsItem /*777 QGraphicsItem **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEventFilter", f)
}

// bool sceneEvent(class QEvent *)
func (this *QGraphicsItem) InheritSceneEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "sceneEvent", f)
}

// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsItem) InheritContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) InheritDropEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsItem) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsItem) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) InheritHoverEnterEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverEnterEvent", f)
}

// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) InheritHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) InheritHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsItem) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsItem) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMousePressEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) InheritMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void wheelEvent(class QGraphicsSceneWheelEvent *)
func (this *QGraphicsItem) InheritWheelEvent(f func(event *QGraphicsSceneWheelEvent /*777 QGraphicsSceneWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsItem) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsItem) InheritInputMethodQuery(f func(query int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "inputMethodQuery", f)
}

// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
func (this *QGraphicsItem) InheritItemChange(f func(change int, value *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "itemChange", f)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
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

// /usr/include/qt/QtWidgets/qgraphicsitem.h:162
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsItem()
func DeleteQGraphicsItem(this *QGraphicsItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:164
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsScene * scene()
func (this *QGraphicsItem) Scene() *QGraphicsScene /*777 QGraphicsScene **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5sceneEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * parentItem()
func (this *QGraphicsItem) ParentItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10parentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * topLevelItem()
func (this *QGraphicsItem) TopLevelItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12topLevelItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * parentWidget()
func (this *QGraphicsItem) ParentWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12parentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * topLevelWidget()
func (this *QGraphicsItem) TopLevelWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14topLevelWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsWidget * window()
func (this *QGraphicsItem) Window() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * panel()
func (this *QGraphicsItem) Panel() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5panelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParentItem(QGraphicsItem *)
func (this *QGraphicsItem) SetParentItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13setParentItemEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:178
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWidget()
func (this *QGraphicsItem) IsWidget() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:179
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWindow()
func (this *QGraphicsItem) IsWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:180
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPanel()
func (this *QGraphicsItem) IsPanel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem7isPanelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItemGroup * group()
func (this *QGraphicsItem) Group() *QGraphicsItemGroup /*777 QGraphicsItemGroup **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGroup(QGraphicsItemGroup *)
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
// [4] QGraphicsItem::GraphicsItemFlags flags()
func (this *QGraphicsItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(enum QGraphicsItem::GraphicsItemFlag, _Bool)
func (this *QGraphicsItem) SetFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem7setFlagENS_16GraphicsItemFlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QGraphicsItem::GraphicsItemFlags)
func (this *QGraphicsItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setFlagsE6QFlagsINS_16GraphicsItemFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:192
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsItem::CacheMode cacheMode()
func (this *QGraphicsItem) CacheMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9cacheModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCacheMode(enum QGraphicsItem::CacheMode, const QSize &)
func (this *QGraphicsItem) SetCacheMode(mode int, cacheSize qtcore.QSize_ITF) {
	var convArg1 unsafe.Pointer
	if cacheSize != nil && cacheSize.QSize_PTR() != nil {
		convArg1 = cacheSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12setCacheModeENS_9CacheModeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:195
// index:0
// Public Visibility=Default Availability=Available
// [4] QGraphicsItem::PanelModality panelModality()
func (this *QGraphicsItem) PanelModality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem13panelModalityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPanelModality(enum QGraphicsItem::PanelModality)
func (this *QGraphicsItem) SetPanelModality(panelModality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem16setPanelModalityENS_13PanelModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), panelModality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:197
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBlockedByModalPanel(QGraphicsItem **)
func (this *QGraphicsItem) IsBlockedByModalPanel(blockingPanel QGraphicsItem_ITF /*777 QGraphicsItem ***/) bool {
	var convArg0 unsafe.Pointer
	if blockingPanel != nil && blockingPanel.QGraphicsItem_PTR() != nil {
		convArg0 = blockingPanel.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip()
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
func (this *QGraphicsItem) SetToolTip(toolTip string) {
	var tmpArg0 = qtcore.NewQString_5(toolTip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:205
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor()
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
// [1] bool hasCursor()
func (this *QGraphicsItem) HasCursor() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9hasCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()
func (this *QGraphicsItem) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QGraphicsItem) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:212
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisibleTo(const QGraphicsItem *)
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
// [-2] void setVisible(_Bool)
func (this *QGraphicsItem) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:214
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void hide()
func (this *QGraphicsItem) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:215
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void show()
func (this *QGraphicsItem) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:217
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QGraphicsItem) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QGraphicsItem) SetEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:220
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSelected()
func (this *QGraphicsItem) IsSelected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10isSelectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelected(_Bool)
func (this *QGraphicsItem) SetSelected(selected bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11setSelectedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), selected)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:223
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptDrops()
func (this *QGraphicsItem) AcceptDrops() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem11acceptDropsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptDrops(_Bool)
func (this *QGraphicsItem) SetAcceptDrops(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14setAcceptDropsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:226
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity()
func (this *QGraphicsItem) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:227
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal effectiveOpacity()
func (this *QGraphicsItem) EffectiveOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16effectiveOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:228
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)
func (this *QGraphicsItem) SetOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEffect * graphicsEffect()
func (this *QGraphicsItem) GraphicsEffect() *QGraphicsEffect /*777 QGraphicsEffect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14graphicsEffectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGraphicsEffect(QGraphicsEffect *)
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
// [4] Qt::MouseButtons acceptedMouseButtons()
func (this *QGraphicsItem) AcceptedMouseButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem20acceptedMouseButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptedMouseButtons(Qt::MouseButtons)
func (this *QGraphicsItem) SetAcceptedMouseButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem23setAcceptedMouseButtonsE6QFlagsIN2Qt11MouseButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:242
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptHoverEvents()
func (this *QGraphicsItem) AcceptHoverEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptHoverEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptHoverEvents(_Bool)
func (this *QGraphicsItem) SetAcceptHoverEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptHoverEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:244
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptTouchEvents()
func (this *QGraphicsItem) AcceptTouchEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptTouchEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptTouchEvents(_Bool)
func (this *QGraphicsItem) SetAcceptTouchEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptTouchEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:247
// index:0
// Public Visibility=Default Availability=Available
// [1] bool filtersChildEvents()
func (this *QGraphicsItem) FiltersChildEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem18filtersChildEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltersChildEvents(_Bool)
func (this *QGraphicsItem) SetFiltersChildEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem21setFiltersChildEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:250
// index:0
// Public Visibility=Default Availability=Available
// [1] bool handlesChildEvents()
func (this *QGraphicsItem) HandlesChildEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem18handlesChildEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHandlesChildEvents(_Bool)
func (this *QGraphicsItem) SetHandlesChildEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem21setHandlesChildEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:253
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QGraphicsItem) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:254
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActive(_Bool)
func (this *QGraphicsItem) SetActive(active bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setActiveEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), active)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:256
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus()
func (this *QGraphicsItem) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)
func (this *QGraphicsItem) SetFocus(focusReason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focusReason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFocus()
func (this *QGraphicsItem) ClearFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10clearFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:260
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * focusProxy()
func (this *QGraphicsItem) FocusProxy() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10focusProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusProxy(QGraphicsItem *)
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
// [8] QGraphicsItem * focusItem()
func (this *QGraphicsItem) FocusItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9focusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:264
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * focusScopeItem()
func (this *QGraphicsItem) FocusScopeItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem14focusScopeItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:266
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabMouse()
func (this *QGraphicsItem) GrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9grabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabMouse()
func (this *QGraphicsItem) UngrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11ungrabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabKeyboard()
func (this *QGraphicsItem) GrabKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12grabKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabKeyboard()
func (this *QGraphicsItem) UngrabKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14ungrabKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:272
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF pos()
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
// [8] qreal x()
func (this *QGraphicsItem) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(qreal)
func (this *QGraphicsItem) SetX(x float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:275
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal y()
func (this *QGraphicsItem) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(qreal)
func (this *QGraphicsItem) SetY(y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:277
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF scenePos()
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
func (this *QGraphicsItem) SetPos_1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6setPosEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:280
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void moveBy(qreal, qreal)
func (this *QGraphicsItem) MoveBy(dx float64, dy float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6moveByEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:282
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensureVisible(const QRectF &, int, int)
func (this *QGraphicsItem) EnsureVisible(rect qtcore.QRectF_ITF, xmargin int, ymargin int) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleERK6QRectFii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:283
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsItem) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleEddddii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:286
// index:0
// Public Visibility=Default Availability=Available
// [48] QMatrix matrix()
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
// [48] QMatrix sceneMatrix()
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
// [-2] void setMatrix(const QMatrix &, _Bool)
func (this *QGraphicsItem) SetMatrix(matrix qtgui.QMatrix_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix_PTR() != nil {
		convArg0 = matrix.QMatrix_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setMatrixERK7QMatrixb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetMatrix()
func (this *QGraphicsItem) ResetMatrix() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11resetMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:290
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transform()
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
// [88] QTransform sceneTransform()
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
// [88] QTransform deviceTransform(const QTransform &)
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
// [88] QTransform itemTransform(const QGraphicsItem *, _Bool *)
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

// /usr/include/qt/QtWidgets/qgraphicsitem.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransform(const QTransform &, _Bool)
func (this *QGraphicsItem) SetTransform(matrix qtgui.QTransform_ITF, combine bool) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem12setTransformERK10QTransformb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:295
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetTransform()
func (this *QGraphicsItem) ResetTransform() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem14resetTransformEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRotation(qreal)
func (this *QGraphicsItem) SetRotation(angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem11setRotationEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:303
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rotation()
func (this *QGraphicsItem) Rotation() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem8rotationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScale(qreal)
func (this *QGraphicsItem) SetScale(scale float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem8setScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scale)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:306
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal scale()
func (this *QGraphicsItem) Scale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem5scaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:311
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF transformOriginPoint()
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
func (this *QGraphicsItem) SetTransformOriginPoint_1(ax float64, ay float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem23setTransformOriginPointEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ax, ay)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:316
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void advance(int)
func (this *QGraphicsItem) Advance(phase int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem7advanceEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), phase)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal zValue()
func (this *QGraphicsItem) ZValue() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem6zValueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZValue(qreal)
func (this *QGraphicsItem) SetZValue(z float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem9setZValueEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), z)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:321
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackBefore(const QGraphicsItem *)
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
// [32] QRectF boundingRect()
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
// [32] QRectF childrenBoundingRect()
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
// [32] QRectF sceneBoundingRect()
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
// [8] QPainterPath shape()
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
// [1] bool isClipped()
func (this *QGraphicsItem) IsClipped() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem9isClippedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:329
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath clipPath()
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
// [1] bool contains(const QPointF &)
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
// [1] bool collidesWithItem(const QGraphicsItem *, Qt::ItemSelectionMode)
func (this *QGraphicsItem) CollidesWithItem(other QGraphicsItem_ITF /*777 const QGraphicsItem **/, mode int) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QGraphicsItem_PTR() != nil {
		convArg0 = other.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithItemEPKS_N2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:332
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool collidesWithPath(const QPainterPath &, Qt::ItemSelectionMode)
func (this *QGraphicsItem) CollidesWithPath(path qtgui.QPainterPath_ITF, mode int) bool {
	var convArg0 unsafe.Pointer
	if path != nil && path.QPainterPath_PTR() != nil {
		convArg0 = path.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithPathERK12QPainterPathN2Qt17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:334
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isObscured(const QRectF &)
func (this *QGraphicsItem) IsObscured(rect qtcore.QRectF_ITF) bool {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:335
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool isObscured(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) IsObscured_1(x float64, y float64, w float64, h float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:336
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
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
// [8] QPainterPath opaqueArea()
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
// [8] QRegion boundingRegion(const QTransform &)
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
// [8] qreal boundingRegionGranularity()
func (this *QGraphicsItem) BoundingRegionGranularity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem25boundingRegionGranularityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:341
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBoundingRegionGranularity(qreal)
func (this *QGraphicsItem) SetBoundingRegionGranularity(granularity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem28setBoundingRegionGranularityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), granularity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:344
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
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

// /usr/include/qt/QtWidgets/qgraphicsitem.h:345
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update(const QRectF &)
func (this *QGraphicsItem) Update(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6updateERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:346
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void update(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) Update_1(x float64, y float64, width float64, height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6updateEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, width, height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:347
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scroll(qreal, qreal, const QRectF &)
func (this *QGraphicsItem) Scroll(dx float64, dy float64, rect qtcore.QRectF_ITF) {
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg2 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem6scrollEddRK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:350
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToItem(const QGraphicsItem *, const QPointF &)
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
// [8] QPolygonF mapToItem(const QGraphicsItem *, const QRectF &)
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
// [8] QPolygonF mapToItem(const QGraphicsItem *, const QPolygonF &)
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
// [8] QPainterPath mapToItem(const QGraphicsItem *, const QPainterPath &)
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
// [16] QPointF mapToItem(const QGraphicsItem *, qreal, qreal)
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
// [8] QPolygonF mapToItem(const QGraphicsItem *, qreal, qreal, qreal, qreal)
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
// [16] QPointF mapToParent(const QPointF &)
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
// [8] QPolygonF mapToParent(const QRectF &)
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
// [8] QPolygonF mapToParent(const QPolygonF &)
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
// [8] QPainterPath mapToParent(const QPainterPath &)
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
// [16] QPointF mapToParent(qreal, qreal)
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
// [8] QPolygonF mapToParent(qreal, qreal, qreal, qreal)
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
// [16] QPointF mapToScene(const QPointF &)
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
// [8] QPolygonF mapToScene(const QRectF &)
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
// [8] QPolygonF mapToScene(const QPolygonF &)
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
// [8] QPainterPath mapToScene(const QPainterPath &)
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
// [16] QPointF mapToScene(qreal, qreal)
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
// [8] QPolygonF mapToScene(qreal, qreal, qreal, qreal)
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
// [32] QRectF mapRectToItem(const QGraphicsItem *, const QRectF &)
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
// [32] QRectF mapRectToItem(const QGraphicsItem *, qreal, qreal, qreal, qreal)
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
// [32] QRectF mapRectToParent(const QRectF &)
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
// [32] QRectF mapRectToParent(qreal, qreal, qreal, qreal)
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
// [32] QRectF mapRectToScene(const QRectF &)
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
// [32] QRectF mapRectToScene(qreal, qreal, qreal, qreal)
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
// [16] QPointF mapFromItem(const QGraphicsItem *, const QPointF &)
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
// [8] QPolygonF mapFromItem(const QGraphicsItem *, const QRectF &)
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
// [8] QPolygonF mapFromItem(const QGraphicsItem *, const QPolygonF &)
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
// [8] QPainterPath mapFromItem(const QGraphicsItem *, const QPainterPath &)
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
// [16] QPointF mapFromItem(const QGraphicsItem *, qreal, qreal)
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
// [8] QPolygonF mapFromItem(const QGraphicsItem *, qreal, qreal, qreal, qreal)
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
// [16] QPointF mapFromParent(const QPointF &)
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
// [8] QPolygonF mapFromParent(const QRectF &)
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
// [8] QPolygonF mapFromParent(const QPolygonF &)
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
// [8] QPainterPath mapFromParent(const QPainterPath &)
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
// [16] QPointF mapFromParent(qreal, qreal)
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
// [8] QPolygonF mapFromParent(qreal, qreal, qreal, qreal)
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
// [16] QPointF mapFromScene(const QPointF &)
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
// [8] QPolygonF mapFromScene(const QRectF &)
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
// [8] QPolygonF mapFromScene(const QPolygonF &)
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
// [8] QPainterPath mapFromScene(const QPainterPath &)
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
// [16] QPointF mapFromScene(qreal, qreal)
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
// [8] QPolygonF mapFromScene(qreal, qreal, qreal, qreal)
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
// [32] QRectF mapRectFromItem(const QGraphicsItem *, const QRectF &)
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
// [32] QRectF mapRectFromItem(const QGraphicsItem *, qreal, qreal, qreal, qreal)
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
// [32] QRectF mapRectFromParent(const QRectF &)
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
// [32] QRectF mapRectFromParent(qreal, qreal, qreal, qreal)
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
// [32] QRectF mapRectFromScene(const QRectF &)
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
// [32] QRectF mapRectFromScene(qreal, qreal, qreal, qreal)
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
// [1] bool isAncestorOf(const QGraphicsItem *)
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
// [8] QGraphicsItem * commonAncestorItem(const QGraphicsItem *)
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
// [1] bool isUnderMouse()
func (this *QGraphicsItem) IsUnderMouse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem12isUnderMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:405
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant data(int)
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
// [4] Qt::InputMethodHints inputMethodHints()
func (this *QGraphicsItem) InputMethodHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem16inputMethodHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:409
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMethodHints(Qt::InputMethodHints)
func (this *QGraphicsItem) SetInputMethodHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem19setInputMethodHintsE6QFlagsIN2Qt15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:415
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:417
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installSceneEventFilter(QGraphicsItem *)
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
func (this *QGraphicsItem) UpdateMicroFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem16updateMicroFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:422
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool sceneEventFilter(QGraphicsItem *, QEvent *)
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
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
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
// [16] QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const QVariant &)
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
// [1] bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:450
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
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
// [16] QVariant extension(const QVariant &)
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
func (this *QGraphicsItem) AddToIndex() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem10addToIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:458
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void removeFromIndex()
func (this *QGraphicsItem) RemoveFromIndex() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem15removeFromIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:459
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void prepareGeometryChange()
func (this *QGraphicsItem) PrepareGeometryChange() {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QGraphicsItem21prepareGeometryChangeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QGraphicsItem__GraphicsItemFlag = int

const QGraphicsItem__ItemIsMovable QGraphicsItem__GraphicsItemFlag = 1
const QGraphicsItem__ItemIsSelectable QGraphicsItem__GraphicsItemFlag = 2
const QGraphicsItem__ItemIsFocusable QGraphicsItem__GraphicsItemFlag = 4
const QGraphicsItem__ItemClipsToShape QGraphicsItem__GraphicsItemFlag = 8
const QGraphicsItem__ItemClipsChildrenToShape QGraphicsItem__GraphicsItemFlag = 16
const QGraphicsItem__ItemIgnoresTransformations QGraphicsItem__GraphicsItemFlag = 32
const QGraphicsItem__ItemIgnoresParentOpacity QGraphicsItem__GraphicsItemFlag = 64
const QGraphicsItem__ItemDoesntPropagateOpacityToChildren QGraphicsItem__GraphicsItemFlag = 128
const QGraphicsItem__ItemStacksBehindParent QGraphicsItem__GraphicsItemFlag = 256
const QGraphicsItem__ItemUsesExtendedStyleOption QGraphicsItem__GraphicsItemFlag = 512
const QGraphicsItem__ItemHasNoContents QGraphicsItem__GraphicsItemFlag = 1024
const QGraphicsItem__ItemSendsGeometryChanges QGraphicsItem__GraphicsItemFlag = 2048
const QGraphicsItem__ItemAcceptsInputMethod QGraphicsItem__GraphicsItemFlag = 4096
const QGraphicsItem__ItemNegativeZStacksBehindParent QGraphicsItem__GraphicsItemFlag = 8192
const QGraphicsItem__ItemIsPanel QGraphicsItem__GraphicsItemFlag = 16384
const QGraphicsItem__ItemIsFocusScope QGraphicsItem__GraphicsItemFlag = 32768
const QGraphicsItem__ItemSendsScenePositionChanges QGraphicsItem__GraphicsItemFlag = 65536
const QGraphicsItem__ItemStopsClickFocusPropagation QGraphicsItem__GraphicsItemFlag = 131072
const QGraphicsItem__ItemStopsFocusHandling QGraphicsItem__GraphicsItemFlag = 262144
const QGraphicsItem__ItemContainsChildrenInShape QGraphicsItem__GraphicsItemFlag = 524288

type QGraphicsItem__GraphicsItemChange = int

const QGraphicsItem__ItemPositionChange QGraphicsItem__GraphicsItemChange = 0
const QGraphicsItem__ItemMatrixChange QGraphicsItem__GraphicsItemChange = 1
const QGraphicsItem__ItemVisibleChange QGraphicsItem__GraphicsItemChange = 2
const QGraphicsItem__ItemEnabledChange QGraphicsItem__GraphicsItemChange = 3
const QGraphicsItem__ItemSelectedChange QGraphicsItem__GraphicsItemChange = 4
const QGraphicsItem__ItemParentChange QGraphicsItem__GraphicsItemChange = 5
const QGraphicsItem__ItemChildAddedChange QGraphicsItem__GraphicsItemChange = 6
const QGraphicsItem__ItemChildRemovedChange QGraphicsItem__GraphicsItemChange = 7
const QGraphicsItem__ItemTransformChange QGraphicsItem__GraphicsItemChange = 8
const QGraphicsItem__ItemPositionHasChanged QGraphicsItem__GraphicsItemChange = 9
const QGraphicsItem__ItemTransformHasChanged QGraphicsItem__GraphicsItemChange = 10
const QGraphicsItem__ItemSceneChange QGraphicsItem__GraphicsItemChange = 11
const QGraphicsItem__ItemVisibleHasChanged QGraphicsItem__GraphicsItemChange = 12
const QGraphicsItem__ItemEnabledHasChanged QGraphicsItem__GraphicsItemChange = 13
const QGraphicsItem__ItemSelectedHasChanged QGraphicsItem__GraphicsItemChange = 14
const QGraphicsItem__ItemParentHasChanged QGraphicsItem__GraphicsItemChange = 15
const QGraphicsItem__ItemSceneHasChanged QGraphicsItem__GraphicsItemChange = 16
const QGraphicsItem__ItemCursorChange QGraphicsItem__GraphicsItemChange = 17
const QGraphicsItem__ItemCursorHasChanged QGraphicsItem__GraphicsItemChange = 18
const QGraphicsItem__ItemToolTipChange QGraphicsItem__GraphicsItemChange = 19
const QGraphicsItem__ItemToolTipHasChanged QGraphicsItem__GraphicsItemChange = 20
const QGraphicsItem__ItemFlagsChange QGraphicsItem__GraphicsItemChange = 21
const QGraphicsItem__ItemFlagsHaveChanged QGraphicsItem__GraphicsItemChange = 22
const QGraphicsItem__ItemZValueChange QGraphicsItem__GraphicsItemChange = 23
const QGraphicsItem__ItemZValueHasChanged QGraphicsItem__GraphicsItemChange = 24
const QGraphicsItem__ItemOpacityChange QGraphicsItem__GraphicsItemChange = 25
const QGraphicsItem__ItemOpacityHasChanged QGraphicsItem__GraphicsItemChange = 26
const QGraphicsItem__ItemScenePositionHasChanged QGraphicsItem__GraphicsItemChange = 27
const QGraphicsItem__ItemRotationChange QGraphicsItem__GraphicsItemChange = 28
const QGraphicsItem__ItemRotationHasChanged QGraphicsItem__GraphicsItemChange = 29
const QGraphicsItem__ItemScaleChange QGraphicsItem__GraphicsItemChange = 30
const QGraphicsItem__ItemScaleHasChanged QGraphicsItem__GraphicsItemChange = 31
const QGraphicsItem__ItemTransformOriginPointChange QGraphicsItem__GraphicsItemChange = 32
const QGraphicsItem__ItemTransformOriginPointHasChanged QGraphicsItem__GraphicsItemChange = 33

type QGraphicsItem__CacheMode = int

const QGraphicsItem__NoCache QGraphicsItem__CacheMode = 0
const QGraphicsItem__ItemCoordinateCache QGraphicsItem__CacheMode = 1
const QGraphicsItem__DeviceCoordinateCache QGraphicsItem__CacheMode = 2

type QGraphicsItem__PanelModality = int

const QGraphicsItem__NonModal QGraphicsItem__PanelModality = 0
const QGraphicsItem__PanelModal QGraphicsItem__PanelModality = 1
const QGraphicsItem__SceneModal QGraphicsItem__PanelModality = 2

type QGraphicsItem__ = int

const QGraphicsItem__Type QGraphicsItem__ = 1
const QGraphicsItem__UserType QGraphicsItem__ = 65536

type QGraphicsItem__Extension = int

const QGraphicsItem__UserExtension QGraphicsItem__Extension = -2147483648

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
