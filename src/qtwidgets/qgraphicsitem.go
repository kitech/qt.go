package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
}

//  ext block end

//  body block begin
type QGraphicsItem struct {
	*qtrt.CObject
}

func (this *QGraphicsItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsItem) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQGraphicsItemFromPointer(cthis unsafe.Pointer) *QGraphicsItem {
	return &QGraphicsItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsItem {
	return NewQGraphicsItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:161
// index:0
// Public
// void QGraphicsItem(class QGraphicsItem *)
func NewQGraphicsItem(parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItemC1EPS_", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:162
// index:0
// Public virtual
// void ~QGraphicsItem()
func DeleteQGraphicsItem(*QGraphicsItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:164
// index:0
// Public
// QGraphicsScene * scene()
func (this *QGraphicsItem) Scene() *QGraphicsScene /*777 QGraphicsScene **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5sceneEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsSceneFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:166
// index:0
// Public
// QGraphicsItem * parentItem()
func (this *QGraphicsItem) ParentItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10parentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:167
// index:0
// Public
// QGraphicsItem * topLevelItem()
func (this *QGraphicsItem) TopLevelItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12topLevelItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:169
// index:0
// Public
// QGraphicsWidget * parentWidget()
func (this *QGraphicsItem) ParentWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12parentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:170
// index:0
// Public
// QGraphicsWidget * topLevelWidget()
func (this *QGraphicsItem) TopLevelWidget() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14topLevelWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:171
// index:0
// Public
// QGraphicsWidget * window()
func (this *QGraphicsItem) Window() *QGraphicsWidget /*777 QGraphicsWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:172
// index:0
// Public
// QGraphicsItem * panel()
func (this *QGraphicsItem) Panel() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5panelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:173
// index:0
// Public
// void setParentItem(class QGraphicsItem *)
func (this *QGraphicsItem) SetParentItem(parent *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13setParentItemEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:178
// index:0
// Public
// bool isWidget()
func (this *QGraphicsItem) IsWidget() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:179
// index:0
// Public
// bool isWindow()
func (this *QGraphicsItem) IsWindow() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:180
// index:0
// Public
// bool isPanel()
func (this *QGraphicsItem) IsPanel() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem7isPanelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:185
// index:0
// Public
// QGraphicsItemGroup * group()
func (this *QGraphicsItem) Group() *QGraphicsItemGroup /*777 QGraphicsItemGroup **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5groupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemGroupFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:186
// index:0
// Public
// void setGroup(class QGraphicsItemGroup *)
func (this *QGraphicsItem) SetGroup(group *QGraphicsItemGroup /*777 QGraphicsItemGroup **/) {
	var convArg0 = group.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:188
// index:0
// Public
// QGraphicsItem::GraphicsItemFlags flags()
func (this *QGraphicsItem) Flags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5flagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:189
// index:0
// Public
// void setFlag(enum QGraphicsItem::GraphicsItemFlag, _Bool)
func (this *QGraphicsItem) SetFlag(flag int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem7setFlagENS_16GraphicsItemFlagEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:190
// index:0
// Public
// void setFlags(QGraphicsItem::GraphicsItemFlags)
func (this *QGraphicsItem) SetFlags(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setFlagsE6QFlagsINS_16GraphicsItemFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:192
// index:0
// Public
// QGraphicsItem::CacheMode cacheMode()
func (this *QGraphicsItem) CacheMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9cacheModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:193
// index:0
// Public
// void setCacheMode(enum QGraphicsItem::CacheMode, const class QSize &)
func (this *QGraphicsItem) SetCacheMode(mode int, cacheSize *qtcore.QSize) {
	var convArg1 = cacheSize.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12setCacheModeENS_9CacheModeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:195
// index:0
// Public
// QGraphicsItem::PanelModality panelModality()
func (this *QGraphicsItem) PanelModality() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13panelModalityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:196
// index:0
// Public
// void setPanelModality(enum QGraphicsItem::PanelModality)
func (this *QGraphicsItem) SetPanelModality(panelModality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16setPanelModalityENS_13PanelModalityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), panelModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:197
// index:0
// Public
// bool isBlockedByModalPanel(class QGraphicsItem **)
func (this *QGraphicsItem) IsBlockedByModalPanel(blockingPanel *QGraphicsItem /*777 QGraphicsItem ***/) bool {
	var convArg0 = blockingPanel.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:200
// index:0
// Public
// QString toolTip()
func (this *QGraphicsItem) ToolTip() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem7toolTipEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:201
// index:0
// Public
// void setToolTip(const class QString &)
func (this *QGraphicsItem) SetToolTip(toolTip *qtcore.QString) {
	var convArg0 = toolTip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:205
// index:0
// Public
// QCursor cursor()
func (this *QGraphicsItem) Cursor() *qtgui.QCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6cursorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:206
// index:0
// Public
// void setCursor(const class QCursor &)
func (this *QGraphicsItem) SetCursor(cursor *qtgui.QCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:207
// index:0
// Public
// bool hasCursor()
func (this *QGraphicsItem) HasCursor() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9hasCursorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:208
// index:0
// Public
// void unsetCursor()
func (this *QGraphicsItem) UnsetCursor() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11unsetCursorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:211
// index:0
// Public
// bool isVisible()
func (this *QGraphicsItem) IsVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:212
// index:0
// Public
// bool isVisibleTo(const class QGraphicsItem *)
func (this *QGraphicsItem) IsVisibleTo(parent *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11isVisibleToEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:213
// index:0
// Public
// void setVisible(_Bool)
func (this *QGraphicsItem) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:214
// index:0
// Public inline
// void hide()
func (this *QGraphicsItem) Hide() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4hideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:215
// index:0
// Public inline
// void show()
func (this *QGraphicsItem) Show() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4showEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:217
// index:0
// Public
// bool isEnabled()
func (this *QGraphicsItem) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:218
// index:0
// Public
// void setEnabled(_Bool)
func (this *QGraphicsItem) SetEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:220
// index:0
// Public
// bool isSelected()
func (this *QGraphicsItem) IsSelected() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10isSelectedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:221
// index:0
// Public
// void setSelected(_Bool)
func (this *QGraphicsItem) SetSelected(selected bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11setSelectedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selected)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:223
// index:0
// Public
// bool acceptDrops()
func (this *QGraphicsItem) AcceptDrops() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11acceptDropsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:224
// index:0
// Public
// void setAcceptDrops(_Bool)
func (this *QGraphicsItem) SetAcceptDrops(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14setAcceptDropsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:226
// index:0
// Public
// qreal opacity()
func (this *QGraphicsItem) Opacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem7opacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:227
// index:0
// Public
// qreal effectiveOpacity()
func (this *QGraphicsItem) EffectiveOpacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16effectiveOpacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:228
// index:0
// Public
// void setOpacity(qreal)
func (this *QGraphicsItem) SetOpacity(opacity float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setOpacityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:232
// index:0
// Public
// QGraphicsEffect * graphicsEffect()
func (this *QGraphicsItem) GraphicsEffect() *QGraphicsEffect /*777 QGraphicsEffect **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14graphicsEffectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:233
// index:0
// Public
// void setGraphicsEffect(class QGraphicsEffect *)
func (this *QGraphicsItem) SetGraphicsEffect(effect *QGraphicsEffect /*777 QGraphicsEffect **/) {
	var convArg0 = effect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:236
// index:0
// Public
// Qt::MouseButtons acceptedMouseButtons()
func (this *QGraphicsItem) AcceptedMouseButtons() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem20acceptedMouseButtonsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:242
// index:0
// Public
// bool acceptHoverEvents()
func (this *QGraphicsItem) AcceptHoverEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptHoverEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:243
// index:0
// Public
// void setAcceptHoverEvents(_Bool)
func (this *QGraphicsItem) SetAcceptHoverEvents(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptHoverEventsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:244
// index:0
// Public
// bool acceptTouchEvents()
func (this *QGraphicsItem) AcceptTouchEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptTouchEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:245
// index:0
// Public
// void setAcceptTouchEvents(_Bool)
func (this *QGraphicsItem) SetAcceptTouchEvents(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptTouchEventsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:247
// index:0
// Public
// bool filtersChildEvents()
func (this *QGraphicsItem) FiltersChildEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem18filtersChildEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:248
// index:0
// Public
// void setFiltersChildEvents(_Bool)
func (this *QGraphicsItem) SetFiltersChildEvents(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21setFiltersChildEventsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:250
// index:0
// Public
// bool handlesChildEvents()
func (this *QGraphicsItem) HandlesChildEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem18handlesChildEventsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:251
// index:0
// Public
// void setHandlesChildEvents(_Bool)
func (this *QGraphicsItem) SetHandlesChildEvents(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21setHandlesChildEventsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:253
// index:0
// Public
// bool isActive()
func (this *QGraphicsItem) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8isActiveEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:254
// index:0
// Public
// void setActive(_Bool)
func (this *QGraphicsItem) SetActive(active bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setActiveEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), active)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:256
// index:0
// Public
// bool hasFocus()
func (this *QGraphicsItem) HasFocus() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8hasFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:257
// index:0
// Public
// void setFocus(Qt::FocusReason)
func (this *QGraphicsItem) SetFocus(focusReason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:258
// index:0
// Public
// void clearFocus()
func (this *QGraphicsItem) ClearFocus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10clearFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:260
// index:0
// Public
// QGraphicsItem * focusProxy()
func (this *QGraphicsItem) FocusProxy() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10focusProxyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:261
// index:0
// Public
// void setFocusProxy(class QGraphicsItem *)
func (this *QGraphicsItem) SetFocusProxy(item *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13setFocusProxyEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:263
// index:0
// Public
// QGraphicsItem * focusItem()
func (this *QGraphicsItem) FocusItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9focusItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:264
// index:0
// Public
// QGraphicsItem * focusScopeItem()
func (this *QGraphicsItem) FocusScopeItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14focusScopeItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:266
// index:0
// Public
// void grabMouse()
func (this *QGraphicsItem) GrabMouse() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9grabMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:267
// index:0
// Public
// void ungrabMouse()
func (this *QGraphicsItem) UngrabMouse() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11ungrabMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:268
// index:0
// Public
// void grabKeyboard()
func (this *QGraphicsItem) GrabKeyboard() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12grabKeyboardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:269
// index:0
// Public
// void ungrabKeyboard()
func (this *QGraphicsItem) UngrabKeyboard() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14ungrabKeyboardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:272
// index:0
// Public
// QPointF pos()
func (this *QGraphicsItem) Pos() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:273
// index:0
// Public inline
// qreal x()
func (this *QGraphicsItem) X() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:274
// index:0
// Public
// void setX(qreal)
func (this *QGraphicsItem) SetX(x float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4setXEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:275
// index:0
// Public inline
// qreal y()
func (this *QGraphicsItem) Y() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:276
// index:0
// Public
// void setY(qreal)
func (this *QGraphicsItem) SetY(y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4setYEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:277
// index:0
// Public
// QPointF scenePos()
func (this *QGraphicsItem) ScenePos() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8scenePosEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:278
// index:0
// Public
// void setPos(const class QPointF &)
func (this *QGraphicsItem) SetPos(pos *qtcore.QPointF) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6setPosERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:279
// index:1
// Public inline
// void setPos(qreal, qreal)
func (this *QGraphicsItem) SetPos_1(x float64, y float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6setPosEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:280
// index:0
// Public inline
// void moveBy(qreal, qreal)
func (this *QGraphicsItem) MoveBy(dx float64, dy float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6moveByEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:282
// index:0
// Public
// void ensureVisible(const class QRectF &, int, int)
func (this *QGraphicsItem) EnsureVisible(rect *qtcore.QRectF, xmargin int, ymargin int) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleERK6QRectFii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:283
// index:1
// Public inline
// void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsItem) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleEddddii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h, xmargin, ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:286
// index:0
// Public
// QMatrix matrix()
func (this *QGraphicsItem) Matrix() *qtgui.QMatrix /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6matrixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:287
// index:0
// Public
// QMatrix sceneMatrix()
func (this *QGraphicsItem) SceneMatrix() *qtgui.QMatrix /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11sceneMatrixEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQMatrixFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:288
// index:0
// Public
// void setMatrix(const class QMatrix &, _Bool)
func (this *QGraphicsItem) SetMatrix(matrix *qtgui.QMatrix, combine bool) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setMatrixERK7QMatrixb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:289
// index:0
// Public
// void resetMatrix()
func (this *QGraphicsItem) ResetMatrix() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11resetMatrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:290
// index:0
// Public
// QTransform transform()
func (this *QGraphicsItem) Transform() *qtgui.QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9transformEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:291
// index:0
// Public
// QTransform sceneTransform()
func (this *QGraphicsItem) SceneTransform() *qtgui.QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14sceneTransformEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:292
// index:0
// Public
// QTransform deviceTransform(const class QTransform &)
func (this *QGraphicsItem) DeviceTransform(viewportTransform *qtgui.QTransform) *qtgui.QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = viewportTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15deviceTransformERK10QTransform", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:293
// index:0
// Public
// QTransform itemTransform(const class QGraphicsItem *, _Bool *)
func (this *QGraphicsItem) ItemTransform(other *QGraphicsItem /*777 const QGraphicsItem **/, ok unsafe.Pointer /*666*/) *qtgui.QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13itemTransformEPKS_Pb", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, &ok)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:294
// index:0
// Public
// void setTransform(const class QTransform &, _Bool)
func (this *QGraphicsItem) SetTransform(matrix *qtgui.QTransform, combine bool) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12setTransformERK10QTransformb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:295
// index:0
// Public
// void resetTransform()
func (this *QGraphicsItem) ResetTransform() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14resetTransformEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:302
// index:0
// Public
// void setRotation(qreal)
func (this *QGraphicsItem) SetRotation(angle float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11setRotationEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:303
// index:0
// Public
// qreal rotation()
func (this *QGraphicsItem) Rotation() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8rotationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:305
// index:0
// Public
// void setScale(qreal)
func (this *QGraphicsItem) SetScale(scale float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setScaleEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), scale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:306
// index:0
// Public
// qreal scale()
func (this *QGraphicsItem) Scale() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5scaleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:311
// index:0
// Public
// QPointF transformOriginPoint()
func (this *QGraphicsItem) TransformOriginPoint() *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem20transformOriginPointEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:312
// index:0
// Public
// void setTransformOriginPoint(const class QPointF &)
func (this *QGraphicsItem) SetTransformOriginPoint(origin *qtcore.QPointF) {
	var convArg0 = origin.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem23setTransformOriginPointERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:313
// index:1
// Public inline
// void setTransformOriginPoint(qreal, qreal)
func (this *QGraphicsItem) SetTransformOriginPoint_1(ax float64, ay float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem23setTransformOriginPointEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ax, ay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:316
// index:0
// Public virtual
// void advance(int)
func (this *QGraphicsItem) Advance(phase int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem7advanceEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), phase)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:319
// index:0
// Public
// qreal zValue()
func (this *QGraphicsItem) ZValue() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6zValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:320
// index:0
// Public
// void setZValue(qreal)
func (this *QGraphicsItem) SetZValue(z float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setZValueEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:321
// index:0
// Public
// void stackBefore(const class QGraphicsItem *)
func (this *QGraphicsItem) StackBefore(sibling *QGraphicsItem /*777 const QGraphicsItem **/) {
	var convArg0 = sibling.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11stackBeforeEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:324
// index:0
// Public pure virtual
// QRectF boundingRect()
func (this *QGraphicsItem) BoundingRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:325
// index:0
// Public
// QRectF childrenBoundingRect()
func (this *QGraphicsItem) ChildrenBoundingRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem20childrenBoundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:326
// index:0
// Public
// QRectF sceneBoundingRect()
func (this *QGraphicsItem) SceneBoundingRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17sceneBoundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:327
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsItem) Shape() *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5shapeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:328
// index:0
// Public
// bool isClipped()
func (this *QGraphicsItem) IsClipped() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9isClippedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:329
// index:0
// Public
// QPainterPath clipPath()
func (this *QGraphicsItem) ClipPath() *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8clipPathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:330
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:331
// index:0
// Public virtual
// bool collidesWithItem(const class QGraphicsItem *, Qt::ItemSelectionMode)
func (this *QGraphicsItem) CollidesWithItem(other *QGraphicsItem /*777 const QGraphicsItem **/, mode int) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithItemEPKS_N2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:332
// index:0
// Public virtual
// bool collidesWithPath(const class QPainterPath &, Qt::ItemSelectionMode)
func (this *QGraphicsItem) CollidesWithPath(path *qtgui.QPainterPath, mode int) bool {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithPathERK12QPainterPathN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:334
// index:0
// Public
// bool isObscured(const class QRectF &)
func (this *QGraphicsItem) IsObscured(rect *qtcore.QRectF) bool {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:335
// index:1
// Public inline
// bool isObscured(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) IsObscured_1(x float64, y float64, w float64, h float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:336
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItem) IsObscuredBy(item *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12isObscuredByEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:337
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:339
// index:0
// Public
// QRegion boundingRegion(const class QTransform &)
func (this *QGraphicsItem) BoundingRegion(itemToDeviceTransform *qtgui.QTransform) *qtgui.QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = itemToDeviceTransform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14boundingRegionERK10QTransform", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:340
// index:0
// Public
// qreal boundingRegionGranularity()
func (this *QGraphicsItem) BoundingRegionGranularity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem25boundingRegionGranularityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:341
// index:0
// Public
// void setBoundingRegionGranularity(qreal)
func (this *QGraphicsItem) SetBoundingRegionGranularity(granularity float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem28setBoundingRegionGranularityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), granularity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:344
// index:0
// Public pure virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsItem) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:345
// index:0
// Public
// void update(const class QRectF &)
func (this *QGraphicsItem) Update(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6updateERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:346
// index:1
// Public inline
// void update(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) Update_1(x float64, y float64, width float64, height float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6updateEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, width, height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:347
// index:0
// Public
// void scroll(qreal, qreal, const class QRectF &)
func (this *QGraphicsItem) Scroll(dx float64, dy float64, rect *qtcore.QRectF) {
	var convArg2 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6scrollEddRK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:350
// index:0
// Public
// QPointF mapToItem(const class QGraphicsItem *, const class QPointF &)
func (this *QGraphicsItem) MapToItem(item *QGraphicsItem /*777 const QGraphicsItem **/, point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:353
// index:1
// Public
// QPolygonF mapToItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapToItem_1(item *QGraphicsItem /*777 const QGraphicsItem **/, rect *qtcore.QRectF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:359
// index:2
// Public
// QPolygonF mapToItem(const class QGraphicsItem *, const class QPolygonF &)
func (this *QGraphicsItem) MapToItem_2(item *QGraphicsItem /*777 const QGraphicsItem **/, polygon *qtgui.QPolygonF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:362
// index:3
// Public
// QPainterPath mapToItem(const class QGraphicsItem *, const class QPainterPath &)
func (this *QGraphicsItem) MapToItem_3(item *QGraphicsItem /*777 const QGraphicsItem **/, path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:381
// index:4
// Public inline
// QPointF mapToItem(const class QGraphicsItem *, qreal, qreal)
func (this *QGraphicsItem) MapToItem_4(item *QGraphicsItem /*777 const QGraphicsItem **/, x float64, y float64) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_dd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:384
// index:5
// Public inline
// QPolygonF mapToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapToItem_5(item *QGraphicsItem /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_dddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:351
// index:0
// Public
// QPointF mapToParent(const class QPointF &)
func (this *QGraphicsItem) MapToParent(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:354
// index:1
// Public
// QPolygonF mapToParent(const class QRectF &)
func (this *QGraphicsItem) MapToParent_1(rect *qtcore.QRectF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:360
// index:2
// Public
// QPolygonF mapToParent(const class QPolygonF &)
func (this *QGraphicsItem) MapToParent_2(polygon *qtgui.QPolygonF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:363
// index:3
// Public
// QPainterPath mapToParent(const class QPainterPath &)
func (this *QGraphicsItem) MapToParent_3(path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:382
// index:4
// Public inline
// QPointF mapToParent(qreal, qreal)
func (this *QGraphicsItem) MapToParent_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentEdd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:385
// index:5
// Public inline
// QPolygonF mapToParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapToParent_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:352
// index:0
// Public
// QPointF mapToScene(const class QPointF &)
func (this *QGraphicsItem) MapToScene(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:355
// index:1
// Public
// QPolygonF mapToScene(const class QRectF &)
func (this *QGraphicsItem) MapToScene_1(rect *qtcore.QRectF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:361
// index:2
// Public
// QPolygonF mapToScene(const class QPolygonF &)
func (this *QGraphicsItem) MapToScene_2(polygon *qtgui.QPolygonF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:364
// index:3
// Public
// QPainterPath mapToScene(const class QPainterPath &)
func (this *QGraphicsItem) MapToScene_3(path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:383
// index:4
// Public inline
// QPointF mapToScene(qreal, qreal)
func (this *QGraphicsItem) MapToScene_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneEdd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:386
// index:5
// Public inline
// QPolygonF mapToScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapToScene_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:356
// index:0
// Public
// QRectF mapRectToItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapRectToItem(item *QGraphicsItem /*777 const QGraphicsItem **/, rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:387
// index:1
// Public inline
// QRectF mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectToItem_1(item *QGraphicsItem /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:357
// index:0
// Public
// QRectF mapRectToParent(const class QRectF &)
func (this *QGraphicsItem) MapRectToParent(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectToParentERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:388
// index:1
// Public inline
// QRectF mapRectToParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectToParent_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectToParentEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:358
// index:0
// Public
// QRectF mapRectToScene(const class QRectF &)
func (this *QGraphicsItem) MapRectToScene(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14mapRectToSceneERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:389
// index:1
// Public inline
// QRectF mapRectToScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectToScene_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14mapRectToSceneEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:365
// index:0
// Public
// QPointF mapFromItem(const class QGraphicsItem *, const class QPointF &)
func (this *QGraphicsItem) MapFromItem(item *QGraphicsItem /*777 const QGraphicsItem **/, point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:368
// index:1
// Public
// QPolygonF mapFromItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapFromItem_1(item *QGraphicsItem /*777 const QGraphicsItem **/, rect *qtcore.QRectF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:374
// index:2
// Public
// QPolygonF mapFromItem(const class QGraphicsItem *, const class QPolygonF &)
func (this *QGraphicsItem) MapFromItem_2(item *QGraphicsItem /*777 const QGraphicsItem **/, polygon *qtgui.QPolygonF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:377
// index:3
// Public
// QPainterPath mapFromItem(const class QGraphicsItem *, const class QPainterPath &)
func (this *QGraphicsItem) MapFromItem_3(item *QGraphicsItem /*777 const QGraphicsItem **/, path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:390
// index:4
// Public inline
// QPointF mapFromItem(const class QGraphicsItem *, qreal, qreal)
func (this *QGraphicsItem) MapFromItem_4(item *QGraphicsItem /*777 const QGraphicsItem **/, x float64, y float64) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_dd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:393
// index:5
// Public inline
// QPolygonF mapFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapFromItem_5(item *QGraphicsItem /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_dddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:366
// index:0
// Public
// QPointF mapFromParent(const class QPointF &)
func (this *QGraphicsItem) MapFromParent(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:369
// index:1
// Public
// QPolygonF mapFromParent(const class QRectF &)
func (this *QGraphicsItem) MapFromParent_1(rect *qtcore.QRectF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:375
// index:2
// Public
// QPolygonF mapFromParent(const class QPolygonF &)
func (this *QGraphicsItem) MapFromParent_2(polygon *qtgui.QPolygonF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:378
// index:3
// Public
// QPainterPath mapFromParent(const class QPainterPath &)
func (this *QGraphicsItem) MapFromParent_3(path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:391
// index:4
// Public inline
// QPointF mapFromParent(qreal, qreal)
func (this *QGraphicsItem) MapFromParent_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentEdd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:394
// index:5
// Public inline
// QPolygonF mapFromParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapFromParent_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:367
// index:0
// Public
// QPointF mapFromScene(const class QPointF &)
func (this *QGraphicsItem) MapFromScene(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK7QPointF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:370
// index:1
// Public
// QPolygonF mapFromScene(const class QRectF &)
func (this *QGraphicsItem) MapFromScene_1(rect *qtcore.QRectF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:376
// index:2
// Public
// QPolygonF mapFromScene(const class QPolygonF &)
func (this *QGraphicsItem) MapFromScene_2(polygon *qtgui.QPolygonF) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = polygon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:379
// index:3
// Public
// QPainterPath mapFromScene(const class QPainterPath &)
func (this *QGraphicsItem) MapFromScene_3(path *qtgui.QPainterPath) *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:392
// index:4
// Public inline
// QPointF mapFromScene(qreal, qreal)
func (this *QGraphicsItem) MapFromScene_4(x float64, y float64) *qtcore.QPointF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneEdd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:395
// index:5
// Public inline
// QPolygonF mapFromScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapFromScene_5(x float64, y float64, w float64, h float64) *qtgui.QPolygonF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:371
// index:0
// Public
// QRectF mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapRectFromItem(item *QGraphicsItem /*777 const QGraphicsItem **/, rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:396
// index:1
// Public inline
// QRectF mapRectFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectFromItem_1(item *QGraphicsItem /*777 const QGraphicsItem **/, x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:372
// index:0
// Public
// QRectF mapRectFromParent(const class QRectF &)
func (this *QGraphicsItem) MapRectFromParent(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17mapRectFromParentERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:397
// index:1
// Public inline
// QRectF mapRectFromParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectFromParent_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17mapRectFromParentEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:373
// index:0
// Public
// QRectF mapRectFromScene(const class QRectF &)
func (this *QGraphicsItem) MapRectFromScene(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:398
// index:1
// Public inline
// QRectF mapRectFromScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectFromScene_1(x float64, y float64, w float64, h float64) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16mapRectFromSceneEdddd", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:400
// index:0
// Public
// bool isAncestorOf(const class QGraphicsItem *)
func (this *QGraphicsItem) IsAncestorOf(child *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12isAncestorOfEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:401
// index:0
// Public
// QGraphicsItem * commonAncestorItem(const class QGraphicsItem *)
func (this *QGraphicsItem) CommonAncestorItem(other *QGraphicsItem /*777 const QGraphicsItem **/) *QGraphicsItem /*777 QGraphicsItem **/ {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem18commonAncestorItemEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:402
// index:0
// Public
// bool isUnderMouse()
func (this *QGraphicsItem) IsUnderMouse() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12isUnderMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:405
// index:0
// Public
// QVariant data(int)
func (this *QGraphicsItem) Data(key int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem4dataEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), key)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:406
// index:0
// Public
// void setData(int, const class QVariant &)
func (this *QGraphicsItem) SetData(key int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem7setDataEiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), key, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:408
// index:0
// Public
// Qt::InputMethodHints inputMethodHints()
func (this *QGraphicsItem) InputMethodHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16inputMethodHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:415
// index:0
// Public virtual
// int type()
func (this *QGraphicsItem) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:417
// index:0
// Public
// void installSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) InstallSceneEventFilter(filterItem *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = filterItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem23installSceneEventFilterEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:418
// index:0
// Public
// void removeSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) RemoveSceneEventFilter(filterItem *QGraphicsItem /*777 QGraphicsItem **/) {
	var convArg0 = filterItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem22removeSceneEventFilterEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:421
// index:0
// Protected
// void updateMicroFocus()
func (this *QGraphicsItem) UpdateMicroFocus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16updateMicroFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:422
// index:0
// Protected virtual
// bool sceneEventFilter(class QGraphicsItem *, class QEvent *)
func (this *QGraphicsItem) SceneEventFilter(watched *QGraphicsItem /*777 QGraphicsItem **/, event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = watched.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:423
// index:0
// Protected virtual
// bool sceneEvent(class QEvent *)
func (this *QGraphicsItem) SceneEvent(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10sceneEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:424
// index:0
// Protected virtual
// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsItem) ContextMenuEvent(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:425
// index:0
// Protected virtual
// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DragEnterEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:426
// index:0
// Protected virtual
// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DragLeaveEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:427
// index:0
// Protected virtual
// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DragMoveEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:428
// index:0
// Protected virtual
// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DropEvent(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:429
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsItem) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:430
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsItem) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:431
// index:0
// Protected virtual
// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) HoverEnterEvent(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:432
// index:0
// Protected virtual
// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) HoverMoveEvent(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:433
// index:0
// Protected virtual
// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) HoverLeaveEvent(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:434
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsItem) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:435
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsItem) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:436
// index:0
// Protected virtual
// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MousePressEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:437
// index:0
// Protected virtual
// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MouseMoveEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:438
// index:0
// Protected virtual
// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MouseReleaseEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:439
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MouseDoubleClickEvent(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:440
// index:0
// Protected virtual
// void wheelEvent(class QGraphicsSceneWheelEvent *)
func (this *QGraphicsItem) WheelEvent(event *QGraphicsSceneWheelEvent /*777 QGraphicsSceneWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:441
// index:0
// Protected virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsItem) InputMethodEvent(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:442
// index:0
// Protected virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), query)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:444
// index:0
// Protected virtual
// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
func (this *QGraphicsItem) ItemChange(change int, value *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), change, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:449
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsItem) SupportsExtension(extension int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:450
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:451
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsItem) Extension(variant *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:457
// index:0
// Protected
// void addToIndex()
func (this *QGraphicsItem) AddToIndex() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10addToIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:458
// index:0
// Protected
// void removeFromIndex()
func (this *QGraphicsItem) RemoveFromIndex() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15removeFromIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:459
// index:0
// Protected
// void prepareGeometryChange()
func (this *QGraphicsItem) PrepareGeometryChange() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21prepareGeometryChangeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
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

const QGraphicsItem__UserExtension QGraphicsItem__Extension = 2147483648

//  body block end
