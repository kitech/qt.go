//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

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
	return this.Cthis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:161
// index:0
// void QGraphicsItem(class QGraphicsItem *)
func NewQGraphicsItem(parent unsafe.Pointer) *QGraphicsItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItemC1EPS_", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsItemFromPointer(cthis)
	return gothis
}
func NewQGraphicsItemFromPointer(cthis unsafe.Pointer) *QGraphicsItem {
	return &QGraphicsItem{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:454
// index:1
// void QGraphicsItem(class QGraphicsItemPrivate &, class QGraphicsItem *)
func NewQGraphicsItem_1(dd unsafe.Pointer, parent unsafe.Pointer) *QGraphicsItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItemC1ER20QGraphicsItemPrivatePS_", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:162
// index:0
// virtual
// void ~QGraphicsItem()
func DeleteQGraphicsItem(*QGraphicsItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:164
// index:0
// QGraphicsScene * scene()
func (this *QGraphicsItem) Scene() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5sceneEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:166
// index:0
// QGraphicsItem * parentItem()
func (this *QGraphicsItem) ParentItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10parentItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:167
// index:0
// QGraphicsItem * topLevelItem()
func (this *QGraphicsItem) TopLevelItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12topLevelItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:168
// index:0
// QGraphicsObject * parentObject()
func (this *QGraphicsItem) ParentObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12parentObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:169
// index:0
// QGraphicsWidget * parentWidget()
func (this *QGraphicsItem) ParentWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12parentWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:170
// index:0
// QGraphicsWidget * topLevelWidget()
func (this *QGraphicsItem) TopLevelWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14topLevelWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:171
// index:0
// QGraphicsWidget * window()
func (this *QGraphicsItem) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6windowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:172
// index:0
// QGraphicsItem * panel()
func (this *QGraphicsItem) Panel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5panelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:173
// index:0
// void setParentItem(class QGraphicsItem *)
func (this *QGraphicsItem) SetParentItem(parent unsafe.Pointer) {
	// 0: (, parent QGraphicsItem *), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13setParentItemEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:177
// index:0
// QList<QGraphicsItem *> childItems()
func (this *QGraphicsItem) ChildItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10childItemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:178
// index:0
// bool isWidget()
func (this *QGraphicsItem) IsWidget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWidgetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:179
// index:0
// bool isWindow()
func (this *QGraphicsItem) IsWindow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8isWindowEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:180
// index:0
// bool isPanel()
func (this *QGraphicsItem) IsPanel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem7isPanelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:182
// index:0
// QGraphicsObject * toGraphicsObject()
func (this *QGraphicsItem) ToGraphicsObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16toGraphicsObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:183
// index:1
// const QGraphicsObject * toGraphicsObject()
func (this *QGraphicsItem) ToGraphicsObject_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16toGraphicsObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:185
// index:0
// QGraphicsItemGroup * group()
func (this *QGraphicsItem) Group() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5groupEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:186
// index:0
// void setGroup(class QGraphicsItemGroup *)
func (this *QGraphicsItem) SetGroup(group unsafe.Pointer) {
	// 0: (, group QGraphicsItemGroup *), (group)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setGroupEP18QGraphicsItemGroup", ffiqt.FFI_TYPE_VOID, this.GetCthis(), group)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:188
// index:0
// QGraphicsItem::GraphicsItemFlags flags()
func (this *QGraphicsItem) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5flagsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:189
// index:0
// void setFlag(enum QGraphicsItem::GraphicsItemFlag, _Bool)
func (this *QGraphicsItem) SetFlag(flag int, enabled bool) {
	// 0: (, flag QGraphicsItem::GraphicsItemFlag, enabled bool), (&flag, &enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem7setFlagENS_16GraphicsItemFlagEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &flag, &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:190
// index:0
// void setFlags(QGraphicsItem::GraphicsItemFlags)
func (this *QGraphicsItem) SetFlags(flags int) {
	// 0: (, QFlags<QGraphicsItem::GraphicsItemFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setFlagsE6QFlagsINS_16GraphicsItemFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:192
// index:0
// QGraphicsItem::CacheMode cacheMode()
func (this *QGraphicsItem) CacheMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9cacheModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:193
// index:0
// void setCacheMode(enum QGraphicsItem::CacheMode, const class QSize &)
func (this *QGraphicsItem) SetCacheMode(mode int, cacheSize unsafe.Pointer) {
	// 0: (, mode QGraphicsItem::CacheMode, cacheSize const QSize &), (&mode, cacheSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12setCacheModeENS_9CacheModeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode, cacheSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:195
// index:0
// QGraphicsItem::PanelModality panelModality()
func (this *QGraphicsItem) PanelModality() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13panelModalityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:196
// index:0
// void setPanelModality(enum QGraphicsItem::PanelModality)
func (this *QGraphicsItem) SetPanelModality(panelModality int) {
	// 0: (, panelModality QGraphicsItem::PanelModality), (&panelModality)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16setPanelModalityENS_13PanelModalityE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &panelModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:197
// index:0
// bool isBlockedByModalPanel(class QGraphicsItem **)
func (this *QGraphicsItem) IsBlockedByModalPanel(blockingPanel unsafe.Pointer) {
	// 0: (, blockingPanel QGraphicsItem **), (blockingPanel)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem21isBlockedByModalPanelEPPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), blockingPanel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:200
// index:0
// QString toolTip()
func (this *QGraphicsItem) ToolTip() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem7toolTipEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:201
// index:0
// void setToolTip(const class QString &)
func (this *QGraphicsItem) SetToolTip(toolTip unsafe.Pointer) {
	// 0: (, toolTip const QString &), (toolTip)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setToolTipERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), toolTip)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:205
// index:0
// QCursor cursor()
func (this *QGraphicsItem) Cursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6cursorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:206
// index:0
// void setCursor(const class QCursor &)
func (this *QGraphicsItem) SetCursor(cursor unsafe.Pointer) {
	// 0: (, cursor const QCursor &), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setCursorERK7QCursor", ffiqt.FFI_TYPE_VOID, this.GetCthis(), cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:207
// index:0
// bool hasCursor()
func (this *QGraphicsItem) HasCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9hasCursorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:208
// index:0
// void unsetCursor()
func (this *QGraphicsItem) UnsetCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11unsetCursorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:211
// index:0
// bool isVisible()
func (this *QGraphicsItem) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:212
// index:0
// bool isVisibleTo(const class QGraphicsItem *)
func (this *QGraphicsItem) IsVisibleTo(parent unsafe.Pointer) {
	// 0: (, parent const QGraphicsItem *), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11isVisibleToEPKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:213
// index:0
// void setVisible(_Bool)
func (this *QGraphicsItem) SetVisible(visible bool) {
	// 0: (, visible bool), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:214
// index:0
// inline
// void hide()
func (this *QGraphicsItem) Hide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4hideEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:215
// index:0
// inline
// void show()
func (this *QGraphicsItem) Show() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4showEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:217
// index:0
// bool isEnabled()
func (this *QGraphicsItem) IsEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9isEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:218
// index:0
// void setEnabled(_Bool)
func (this *QGraphicsItem) SetEnabled(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:220
// index:0
// bool isSelected()
func (this *QGraphicsItem) IsSelected() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10isSelectedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:221
// index:0
// void setSelected(_Bool)
func (this *QGraphicsItem) SetSelected(selected bool) {
	// 0: (, selected bool), (&selected)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11setSelectedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &selected)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:223
// index:0
// bool acceptDrops()
func (this *QGraphicsItem) AcceptDrops() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11acceptDropsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:224
// index:0
// void setAcceptDrops(_Bool)
func (this *QGraphicsItem) SetAcceptDrops(on bool) {
	// 0: (, on bool), (&on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14setAcceptDropsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:226
// index:0
// qreal opacity()
func (this *QGraphicsItem) Opacity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem7opacityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:227
// index:0
// qreal effectiveOpacity()
func (this *QGraphicsItem) EffectiveOpacity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16effectiveOpacityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:228
// index:0
// void setOpacity(qreal)
func (this *QGraphicsItem) SetOpacity(opacity float64) {
	// 0: (, opacity qreal), (&opacity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10setOpacityEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:232
// index:0
// QGraphicsEffect * graphicsEffect()
func (this *QGraphicsItem) GraphicsEffect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14graphicsEffectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:233
// index:0
// void setGraphicsEffect(class QGraphicsEffect *)
func (this *QGraphicsItem) SetGraphicsEffect(effect unsafe.Pointer) {
	// 0: (, effect QGraphicsEffect *), (effect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem17setGraphicsEffectEP15QGraphicsEffect", ffiqt.FFI_TYPE_VOID, this.GetCthis(), effect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:236
// index:0
// Qt::MouseButtons acceptedMouseButtons()
func (this *QGraphicsItem) AcceptedMouseButtons() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem20acceptedMouseButtonsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:237
// index:0
// void setAcceptedMouseButtons(Qt::MouseButtons)
func (this *QGraphicsItem) SetAcceptedMouseButtons(buttons int) {
	// 0: (, QFlags<Qt::MouseButton> buttons), (&buttons)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem23setAcceptedMouseButtonsE6QFlagsIN2Qt11MouseButtonEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:242
// index:0
// bool acceptHoverEvents()
func (this *QGraphicsItem) AcceptHoverEvents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptHoverEventsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:243
// index:0
// void setAcceptHoverEvents(_Bool)
func (this *QGraphicsItem) SetAcceptHoverEvents(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptHoverEventsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:244
// index:0
// bool acceptTouchEvents()
func (this *QGraphicsItem) AcceptTouchEvents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17acceptTouchEventsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:245
// index:0
// void setAcceptTouchEvents(_Bool)
func (this *QGraphicsItem) SetAcceptTouchEvents(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem20setAcceptTouchEventsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:247
// index:0
// bool filtersChildEvents()
func (this *QGraphicsItem) FiltersChildEvents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem18filtersChildEventsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:248
// index:0
// void setFiltersChildEvents(_Bool)
func (this *QGraphicsItem) SetFiltersChildEvents(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21setFiltersChildEventsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:250
// index:0
// bool handlesChildEvents()
func (this *QGraphicsItem) HandlesChildEvents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem18handlesChildEventsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:251
// index:0
// void setHandlesChildEvents(_Bool)
func (this *QGraphicsItem) SetHandlesChildEvents(enabled bool) {
	// 0: (, enabled bool), (&enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21setHandlesChildEventsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:253
// index:0
// bool isActive()
func (this *QGraphicsItem) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8isActiveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:254
// index:0
// void setActive(_Bool)
func (this *QGraphicsItem) SetActive(active bool) {
	// 0: (, active bool), (&active)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setActiveEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &active)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:256
// index:0
// bool hasFocus()
func (this *QGraphicsItem) HasFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8hasFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:257
// index:0
// void setFocus(Qt::FocusReason)
func (this *QGraphicsItem) SetFocus(focusReason int) {
	// 0: (, focusReason Qt::FocusReason), (&focusReason)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &focusReason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:258
// index:0
// void clearFocus()
func (this *QGraphicsItem) ClearFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10clearFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:260
// index:0
// QGraphicsItem * focusProxy()
func (this *QGraphicsItem) FocusProxy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10focusProxyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:261
// index:0
// void setFocusProxy(class QGraphicsItem *)
func (this *QGraphicsItem) SetFocusProxy(item unsafe.Pointer) {
	// 0: (, item QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13setFocusProxyEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:263
// index:0
// QGraphicsItem * focusItem()
func (this *QGraphicsItem) FocusItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9focusItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:264
// index:0
// QGraphicsItem * focusScopeItem()
func (this *QGraphicsItem) FocusScopeItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14focusScopeItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:266
// index:0
// void grabMouse()
func (this *QGraphicsItem) GrabMouse() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9grabMouseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:267
// index:0
// void ungrabMouse()
func (this *QGraphicsItem) UngrabMouse() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11ungrabMouseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:268
// index:0
// void grabKeyboard()
func (this *QGraphicsItem) GrabKeyboard() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12grabKeyboardEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:269
// index:0
// void ungrabKeyboard()
func (this *QGraphicsItem) UngrabKeyboard() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14ungrabKeyboardEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:272
// index:0
// QPointF pos()
func (this *QGraphicsItem) Pos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem3posEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:273
// index:0
// inline
// qreal x()
func (this *QGraphicsItem) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem1xEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:274
// index:0
// void setX(qreal)
func (this *QGraphicsItem) SetX(x float64) {
	// 0: (, x qreal), (&x)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4setXEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:275
// index:0
// inline
// qreal y()
func (this *QGraphicsItem) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem1yEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:276
// index:0
// void setY(qreal)
func (this *QGraphicsItem) SetY(y float64) {
	// 0: (, y qreal), (&y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem4setYEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:277
// index:0
// QPointF scenePos()
func (this *QGraphicsItem) ScenePos() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8scenePosEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:278
// index:0
// void setPos(const class QPointF &)
func (this *QGraphicsItem) SetPos(pos unsafe.Pointer) {
	// 0: (, pos const QPointF &), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6setPosERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:279
// index:1
// inline
// void setPos(qreal, qreal)
func (this *QGraphicsItem) SetPos_1(x float64, y float64) {
	// 1: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6setPosEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:280
// index:0
// inline
// void moveBy(qreal, qreal)
func (this *QGraphicsItem) MoveBy(dx float64, dy float64) {
	// 0: (, dx qreal, dy qreal), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6moveByEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:282
// index:0
// void ensureVisible(const class QRectF &, int, int)
func (this *QGraphicsItem) EnsureVisible(rect unsafe.Pointer, xmargin int, ymargin int) {
	// 0: (, rect const QRectF &, xmargin int, ymargin int), (rect, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleERK6QRectFii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:283
// index:1
// inline
// void ensureVisible(qreal, qreal, qreal, qreal, int, int)
func (this *QGraphicsItem) EnsureVisible_1(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	// 1: (, x qreal, y qreal, w qreal, h qreal, xmargin int, ymargin int), (&x, &y, &w, &h, &xmargin, &ymargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13ensureVisibleEddddii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h, &xmargin, &ymargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:286
// index:0
// QMatrix matrix()
func (this *QGraphicsItem) Matrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6matrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:287
// index:0
// QMatrix sceneMatrix()
func (this *QGraphicsItem) SceneMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11sceneMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:288
// index:0
// void setMatrix(const class QMatrix &, _Bool)
func (this *QGraphicsItem) SetMatrix(matrix unsafe.Pointer, combine bool) {
	// 0: (, matrix const QMatrix &, combine bool), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setMatrixERK7QMatrixb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:289
// index:0
// void resetMatrix()
func (this *QGraphicsItem) ResetMatrix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11resetMatrixEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:290
// index:0
// QTransform transform()
func (this *QGraphicsItem) Transform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9transformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:291
// index:0
// QTransform sceneTransform()
func (this *QGraphicsItem) SceneTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14sceneTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:292
// index:0
// QTransform deviceTransform(const class QTransform &)
func (this *QGraphicsItem) DeviceTransform(viewportTransform unsafe.Pointer) {
	// 0: (, viewportTransform const QTransform &), (viewportTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15deviceTransformERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), viewportTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:293
// index:0
// QTransform itemTransform(const class QGraphicsItem *, _Bool *)
func (this *QGraphicsItem) ItemTransform(other unsafe.Pointer, ok unsafe.Pointer) {
	// 0: (, other const QGraphicsItem *, ok bool *), (other, ok)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13itemTransformEPKS_Pb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other, ok)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:294
// index:0
// void setTransform(const class QTransform &, _Bool)
func (this *QGraphicsItem) SetTransform(matrix unsafe.Pointer, combine bool) {
	// 0: (, matrix const QTransform &, combine bool), (matrix, &combine)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12setTransformERK10QTransformb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), matrix, &combine)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:295
// index:0
// void resetTransform()
func (this *QGraphicsItem) ResetTransform() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14resetTransformEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:302
// index:0
// void setRotation(qreal)
func (this *QGraphicsItem) SetRotation(angle float64) {
	// 0: (, angle qreal), (&angle)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11setRotationEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &angle)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:303
// index:0
// qreal rotation()
func (this *QGraphicsItem) Rotation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8rotationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:305
// index:0
// void setScale(qreal)
func (this *QGraphicsItem) SetScale(scale float64) {
	// 0: (, scale qreal), (&scale)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem8setScaleEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &scale)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:306
// index:0
// qreal scale()
func (this *QGraphicsItem) Scale() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5scaleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:308
// index:0
// QList<QGraphicsTransform *> transformations()
func (this *QGraphicsItem) Transformations() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15transformationsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:311
// index:0
// QPointF transformOriginPoint()
func (this *QGraphicsItem) TransformOriginPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem20transformOriginPointEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:312
// index:0
// void setTransformOriginPoint(const class QPointF &)
func (this *QGraphicsItem) SetTransformOriginPoint(origin unsafe.Pointer) {
	// 0: (, origin const QPointF &), (origin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem23setTransformOriginPointERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), origin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:313
// index:1
// inline
// void setTransformOriginPoint(qreal, qreal)
func (this *QGraphicsItem) SetTransformOriginPoint_1(ax float64, ay float64) {
	// 1: (, ax qreal, ay qreal), (&ax, &ay)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem23setTransformOriginPointEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &ax, &ay)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:316
// index:0
// virtual
// void advance(int)
func (this *QGraphicsItem) Advance(phase int) {
	// 0: (, phase int), (&phase)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem7advanceEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &phase)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:319
// index:0
// qreal zValue()
func (this *QGraphicsItem) ZValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem6zValueEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:320
// index:0
// void setZValue(qreal)
func (this *QGraphicsItem) SetZValue(z float64) {
	// 0: (, z qreal), (&z)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9setZValueEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &z)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:321
// index:0
// void stackBefore(const class QGraphicsItem *)
func (this *QGraphicsItem) StackBefore(sibling unsafe.Pointer) {
	// 0: (, sibling const QGraphicsItem *), (sibling)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem11stackBeforeEPKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), sibling)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:324
// index:0
// pure virtual
// QRectF boundingRect()
func (this *QGraphicsItem) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:325
// index:0
// QRectF childrenBoundingRect()
func (this *QGraphicsItem) ChildrenBoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem20childrenBoundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:326
// index:0
// QRectF sceneBoundingRect()
func (this *QGraphicsItem) SceneBoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17sceneBoundingRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:327
// index:0
// virtual
// QPainterPath shape()
func (this *QGraphicsItem) Shape() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem5shapeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:328
// index:0
// bool isClipped()
func (this *QGraphicsItem) IsClipped() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9isClippedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:329
// index:0
// QPainterPath clipPath()
func (this *QGraphicsItem) ClipPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8clipPathEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:330
// index:0
// virtual
// bool contains(const class QPointF &)
func (this *QGraphicsItem) Contains(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem8containsERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:331
// index:0
// virtual
// bool collidesWithItem(const class QGraphicsItem *, Qt::ItemSelectionMode)
func (this *QGraphicsItem) CollidesWithItem(other unsafe.Pointer, mode int) {
	// 0: (, other const QGraphicsItem *, mode Qt::ItemSelectionMode), (other, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithItemEPKS_N2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:332
// index:0
// virtual
// bool collidesWithPath(const class QPainterPath &, Qt::ItemSelectionMode)
func (this *QGraphicsItem) CollidesWithPath(path unsafe.Pointer, mode int) {
	// 0: (, path const QPainterPath &, mode Qt::ItemSelectionMode), (path, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16collidesWithPathERK12QPainterPathN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:333
// index:0
// QList<QGraphicsItem *> collidingItems(Qt::ItemSelectionMode)
func (this *QGraphicsItem) CollidingItems(mode int) {
	// 0: (, mode Qt::ItemSelectionMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14collidingItemsEN2Qt17ItemSelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:334
// index:0
// bool isObscured(const class QRectF &)
func (this *QGraphicsItem) IsObscured(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:335
// index:1
// inline
// bool isObscured(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) IsObscured_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10isObscuredEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:336
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12isObscuredByEPKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:337
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:339
// index:0
// QRegion boundingRegion(const class QTransform &)
func (this *QGraphicsItem) BoundingRegion(itemToDeviceTransform unsafe.Pointer) {
	// 0: (, itemToDeviceTransform const QTransform &), (itemToDeviceTransform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14boundingRegionERK10QTransform", ffiqt.FFI_TYPE_VOID, this.GetCthis(), itemToDeviceTransform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:340
// index:0
// qreal boundingRegionGranularity()
func (this *QGraphicsItem) BoundingRegionGranularity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem25boundingRegionGranularityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:341
// index:0
// void setBoundingRegionGranularity(qreal)
func (this *QGraphicsItem) SetBoundingRegionGranularity(granularity float64) {
	// 0: (, granularity qreal), (&granularity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem28setBoundingRegionGranularityEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &granularity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:344
// index:0
// pure virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, painter QPainter *, option const QStyleOptionGraphicsItem *, widget QWidget *), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:345
// index:0
// void update(const class QRectF &)
func (this *QGraphicsItem) Update(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6updateERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:346
// index:1
// inline
// void update(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) Update_1(x float64, y float64, width float64, height float64) {
	// 1: (, x qreal, y qreal, width qreal, height qreal), (&x, &y, &width, &height)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6updateEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &width, &height)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:347
// index:0
// void scroll(qreal, qreal, const class QRectF &)
func (this *QGraphicsItem) Scroll(dx float64, dy float64, rect unsafe.Pointer) {
	// 0: (, dx qreal, dy qreal, rect const QRectF &), (&dx, &dy, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem6scrollEddRK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:350
// index:0
// QPointF mapToItem(const class QGraphicsItem *, const class QPointF &)
func (this *QGraphicsItem) MapToItem(item unsafe.Pointer, point unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *, point const QPointF &), (item, point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:353
// index:1
// QPolygonF mapToItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapToItem_1(item unsafe.Pointer, rect unsafe.Pointer) {
	// 1: (, item const QGraphicsItem *, rect const QRectF &), (item, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:359
// index:2
// QPolygonF mapToItem(const class QGraphicsItem *, const class QPolygonF &)
func (this *QGraphicsItem) MapToItem_2(item unsafe.Pointer, polygon unsafe.Pointer) {
	// 2: (, item const QGraphicsItem *, polygon const QPolygonF &), (item, polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:362
// index:3
// QPainterPath mapToItem(const class QGraphicsItem *, const class QPainterPath &)
func (this *QGraphicsItem) MapToItem_3(item unsafe.Pointer, path unsafe.Pointer) {
	// 3: (, item const QGraphicsItem *, path const QPainterPath &), (item, path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_RK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:381
// index:4
// inline
// QPointF mapToItem(const class QGraphicsItem *, qreal, qreal)
func (this *QGraphicsItem) MapToItem_4(item unsafe.Pointer, x float64, y float64) {
	// 4: (, item const QGraphicsItem *, x qreal, y qreal), (item, &x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_dd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:384
// index:5
// inline
// QPolygonF mapToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapToItem_5(item unsafe.Pointer, x float64, y float64, w float64, h float64) {
	// 5: (, item const QGraphicsItem *, x qreal, y qreal, w qreal, h qreal), (item, &x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9mapToItemEPKS_dddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:351
// index:0
// QPointF mapToParent(const class QPointF &)
func (this *QGraphicsItem) MapToParent(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:354
// index:1
// QPolygonF mapToParent(const class QRectF &)
func (this *QGraphicsItem) MapToParent_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:360
// index:2
// QPolygonF mapToParent(const class QPolygonF &)
func (this *QGraphicsItem) MapToParent_2(polygon unsafe.Pointer) {
	// 2: (, polygon const QPolygonF &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:363
// index:3
// QPainterPath mapToParent(const class QPainterPath &)
func (this *QGraphicsItem) MapToParent_3(path unsafe.Pointer) {
	// 3: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:382
// index:4
// inline
// QPointF mapToParent(qreal, qreal)
func (this *QGraphicsItem) MapToParent_4(x float64, y float64) {
	// 4: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:385
// index:5
// inline
// QPolygonF mapToParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapToParent_5(x float64, y float64, w float64, h float64) {
	// 5: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapToParentEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:352
// index:0
// QPointF mapToScene(const class QPointF &)
func (this *QGraphicsItem) MapToScene(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:355
// index:1
// QPolygonF mapToScene(const class QRectF &)
func (this *QGraphicsItem) MapToScene_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:361
// index:2
// QPolygonF mapToScene(const class QPolygonF &)
func (this *QGraphicsItem) MapToScene_2(polygon unsafe.Pointer) {
	// 2: (, polygon const QPolygonF &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:364
// index:3
// QPainterPath mapToScene(const class QPainterPath &)
func (this *QGraphicsItem) MapToScene_3(path unsafe.Pointer) {
	// 3: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:383
// index:4
// inline
// QPointF mapToScene(qreal, qreal)
func (this *QGraphicsItem) MapToScene_4(x float64, y float64) {
	// 4: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:386
// index:5
// inline
// QPolygonF mapToScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapToScene_5(x float64, y float64, w float64, h float64) {
	// 5: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem10mapToSceneEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:356
// index:0
// QRectF mapRectToItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapRectToItem(item unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *, rect const QRectF &), (item, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapRectToItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:387
// index:1
// inline
// QRectF mapRectToItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectToItem_1(item unsafe.Pointer, x float64, y float64, w float64, h float64) {
	// 1: (, item const QGraphicsItem *, x qreal, y qreal, w qreal, h qreal), (item, &x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapRectToItemEPKS_dddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:357
// index:0
// QRectF mapRectToParent(const class QRectF &)
func (this *QGraphicsItem) MapRectToParent(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectToParentERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:388
// index:1
// inline
// QRectF mapRectToParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectToParent_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectToParentEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:358
// index:0
// QRectF mapRectToScene(const class QRectF &)
func (this *QGraphicsItem) MapRectToScene(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14mapRectToSceneERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:389
// index:1
// inline
// QRectF mapRectToScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectToScene_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem14mapRectToSceneEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:365
// index:0
// QPointF mapFromItem(const class QGraphicsItem *, const class QPointF &)
func (this *QGraphicsItem) MapFromItem(item unsafe.Pointer, point unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *, point const QPointF &), (item, point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:368
// index:1
// QPolygonF mapFromItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapFromItem_1(item unsafe.Pointer, rect unsafe.Pointer) {
	// 1: (, item const QGraphicsItem *, rect const QRectF &), (item, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:374
// index:2
// QPolygonF mapFromItem(const class QGraphicsItem *, const class QPolygonF &)
func (this *QGraphicsItem) MapFromItem_2(item unsafe.Pointer, polygon unsafe.Pointer) {
	// 2: (, item const QGraphicsItem *, polygon const QPolygonF &), (item, polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:377
// index:3
// QPainterPath mapFromItem(const class QGraphicsItem *, const class QPainterPath &)
func (this *QGraphicsItem) MapFromItem_3(item unsafe.Pointer, path unsafe.Pointer) {
	// 3: (, item const QGraphicsItem *, path const QPainterPath &), (item, path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_RK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:390
// index:4
// inline
// QPointF mapFromItem(const class QGraphicsItem *, qreal, qreal)
func (this *QGraphicsItem) MapFromItem_4(item unsafe.Pointer, x float64, y float64) {
	// 4: (, item const QGraphicsItem *, x qreal, y qreal), (item, &x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_dd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:393
// index:5
// inline
// QPolygonF mapFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapFromItem_5(item unsafe.Pointer, x float64, y float64, w float64, h float64) {
	// 5: (, item const QGraphicsItem *, x qreal, y qreal, w qreal, h qreal), (item, &x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem11mapFromItemEPKS_dddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:366
// index:0
// QPointF mapFromParent(const class QPointF &)
func (this *QGraphicsItem) MapFromParent(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:369
// index:1
// QPolygonF mapFromParent(const class QRectF &)
func (this *QGraphicsItem) MapFromParent_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:375
// index:2
// QPolygonF mapFromParent(const class QPolygonF &)
func (this *QGraphicsItem) MapFromParent_2(polygon unsafe.Pointer) {
	// 2: (, polygon const QPolygonF &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:378
// index:3
// QPainterPath mapFromParent(const class QPainterPath &)
func (this *QGraphicsItem) MapFromParent_3(path unsafe.Pointer) {
	// 3: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:391
// index:4
// inline
// QPointF mapFromParent(qreal, qreal)
func (this *QGraphicsItem) MapFromParent_4(x float64, y float64) {
	// 4: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:394
// index:5
// inline
// QPolygonF mapFromParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapFromParent_5(x float64, y float64, w float64, h float64) {
	// 5: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem13mapFromParentEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:367
// index:0
// QPointF mapFromScene(const class QPointF &)
func (this *QGraphicsItem) MapFromScene(point unsafe.Pointer) {
	// 0: (, point const QPointF &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK7QPointF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:370
// index:1
// QPolygonF mapFromScene(const class QRectF &)
func (this *QGraphicsItem) MapFromScene_1(rect unsafe.Pointer) {
	// 1: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:376
// index:2
// QPolygonF mapFromScene(const class QPolygonF &)
func (this *QGraphicsItem) MapFromScene_2(polygon unsafe.Pointer) {
	// 2: (, polygon const QPolygonF &), (polygon)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK9QPolygonF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), polygon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:379
// index:3
// QPainterPath mapFromScene(const class QPainterPath &)
func (this *QGraphicsItem) MapFromScene_3(path unsafe.Pointer) {
	// 3: (, path const QPainterPath &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneERK12QPainterPath", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:392
// index:4
// inline
// QPointF mapFromScene(qreal, qreal)
func (this *QGraphicsItem) MapFromScene_4(x float64, y float64) {
	// 4: (, x qreal, y qreal), (&x, &y)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneEdd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:395
// index:5
// inline
// QPolygonF mapFromScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapFromScene_5(x float64, y float64, w float64, h float64) {
	// 5: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12mapFromSceneEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:371
// index:0
// QRectF mapRectFromItem(const class QGraphicsItem *, const class QRectF &)
func (this *QGraphicsItem) MapRectFromItem(item unsafe.Pointer, rect unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *, rect const QRectF &), (item, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectFromItemEPKS_RK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:396
// index:1
// inline
// QRectF mapRectFromItem(const class QGraphicsItem *, qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectFromItem_1(item unsafe.Pointer, x float64, y float64, w float64, h float64) {
	// 1: (, item const QGraphicsItem *, x qreal, y qreal, w qreal, h qreal), (item, &x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem15mapRectFromItemEPKS_dddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:372
// index:0
// QRectF mapRectFromParent(const class QRectF &)
func (this *QGraphicsItem) MapRectFromParent(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17mapRectFromParentERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:397
// index:1
// inline
// QRectF mapRectFromParent(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectFromParent_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17mapRectFromParentEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:373
// index:0
// QRectF mapRectFromScene(const class QRectF &)
func (this *QGraphicsItem) MapRectFromScene(rect unsafe.Pointer) {
	// 0: (, rect const QRectF &), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16mapRectFromSceneERK6QRectF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:398
// index:1
// inline
// QRectF mapRectFromScene(qreal, qreal, qreal, qreal)
func (this *QGraphicsItem) MapRectFromScene_1(x float64, y float64, w float64, h float64) {
	// 1: (, x qreal, y qreal, w qreal, h qreal), (&x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16mapRectFromSceneEdddd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:400
// index:0
// bool isAncestorOf(const class QGraphicsItem *)
func (this *QGraphicsItem) IsAncestorOf(child unsafe.Pointer) {
	// 0: (, child const QGraphicsItem *), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12isAncestorOfEPKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:401
// index:0
// QGraphicsItem * commonAncestorItem(const class QGraphicsItem *)
func (this *QGraphicsItem) CommonAncestorItem(other unsafe.Pointer) {
	// 0: (, other const QGraphicsItem *), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem18commonAncestorItemEPKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:402
// index:0
// bool isUnderMouse()
func (this *QGraphicsItem) IsUnderMouse() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem12isUnderMouseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:405
// index:0
// QVariant data(int)
func (this *QGraphicsItem) Data(key int) {
	// 0: (, key int), (&key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem4dataEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:406
// index:0
// void setData(int, const class QVariant &)
func (this *QGraphicsItem) SetData(key int, value unsafe.Pointer) {
	// 0: (, key int, value const QVariant &), (&key, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem7setDataEiRK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &key, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:408
// index:0
// Qt::InputMethodHints inputMethodHints()
func (this *QGraphicsItem) InputMethodHints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16inputMethodHintsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:409
// index:0
// void setInputMethodHints(Qt::InputMethodHints)
func (this *QGraphicsItem) SetInputMethodHints(hints int) {
	// 0: (, QFlags<Qt::InputMethodHint> hints), (&hints)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem19setInputMethodHintsE6QFlagsIN2Qt15InputMethodHintEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:415
// index:0
// virtual
// int type()
func (this *QGraphicsItem) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem4typeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:417
// index:0
// void installSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) InstallSceneEventFilter(filterItem unsafe.Pointer) {
	// 0: (, filterItem QGraphicsItem *), (filterItem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem23installSceneEventFilterEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filterItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:418
// index:0
// void removeSceneEventFilter(class QGraphicsItem *)
func (this *QGraphicsItem) RemoveSceneEventFilter(filterItem unsafe.Pointer) {
	// 0: (, filterItem QGraphicsItem *), (filterItem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem22removeSceneEventFilterEPS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filterItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:421
// index:0
// void updateMicroFocus()
func (this *QGraphicsItem) UpdateMicroFocus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16updateMicroFocusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:422
// index:0
// virtual
// bool sceneEventFilter(class QGraphicsItem *, class QEvent *)
func (this *QGraphicsItem) SceneEventFilter(watched unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, watched QGraphicsItem *, event QEvent *), (watched, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), watched, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:423
// index:0
// virtual
// bool sceneEvent(class QEvent *)
func (this *QGraphicsItem) SceneEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10sceneEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:424
// index:0
// virtual
// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsItem) ContextMenuEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneContextMenuEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:425
// index:0
// virtual
// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DragEnterEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:426
// index:0
// virtual
// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DragLeaveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:427
// index:0
// virtual
// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DragMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:428
// index:0
// virtual
// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsItem) DropEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneDragDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:429
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsItem) FocusInEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:430
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsItem) FocusOutEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:431
// index:0
// virtual
// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) HoverEnterEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneHoverEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:432
// index:0
// virtual
// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) HoverMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneHoverEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:433
// index:0
// virtual
// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsItem) HoverLeaveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneHoverEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:434
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsItem) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:435
// index:0
// virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsItem) KeyReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:436
// index:0
// virtual
// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:437
// index:0
// virtual
// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MouseMoveEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:438
// index:0
// virtual
// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MouseReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:439
// index:0
// virtual
// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsItem) MouseDoubleClickEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:440
// index:0
// virtual
// void wheelEvent(class QGraphicsSceneWheelEvent *)
func (this *QGraphicsItem) WheelEvent(event unsafe.Pointer) {
	// 0: (, event QGraphicsSceneWheelEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:441
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsItem) InputMethodEvent(event unsafe.Pointer) {
	// 0: (, event QInputMethodEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:442
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsItem) InputMethodQuery(query int) {
	// 0: (, query Qt::InputMethodQuery), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:444
// index:0
// virtual
// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
func (this *QGraphicsItem) ItemChange(change int, value unsafe.Pointer) {
	// 0: (, change QGraphicsItem::GraphicsItemChange, value const QVariant &), (&change, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &change, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:449
// index:0
// virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsItem) SupportsExtension(extension int) {
	// 0: (, extension QGraphicsItem::Extension), (&extension)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:450
// index:0
// virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsItem) SetExtension(extension int, variant unsafe.Pointer) {
	// 0: (, extension QGraphicsItem::Extension, variant const QVariant &), (&extension, variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &extension, variant)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:451
// index:0
// virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsItem) Extension(variant unsafe.Pointer) {
	// 0: (, variant const QVariant &), (variant)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QGraphicsItem9extensionERK8QVariant", ffiqt.FFI_TYPE_VOID, this.GetCthis(), variant)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:457
// index:0
// void addToIndex()
func (this *QGraphicsItem) AddToIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem10addToIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:458
// index:0
// void removeFromIndex()
func (this *QGraphicsItem) RemoveFromIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem15removeFromIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:459
// index:0
// void prepareGeometryChange()
func (this *QGraphicsItem) PrepareGeometryChange() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QGraphicsItem21prepareGeometryChangeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
