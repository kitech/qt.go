//  header block begin
// /usr/include/qt/QtWidgets/qabstractitemview.h
// #include <qabstractitemview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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
type QAbstractItemView struct {
	*QAbstractScrollArea
}

func (this *QAbstractItemView) GetCthis() unsafe.Pointer {
	return this.QAbstractScrollArea.GetCthis()
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:63
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QAbstractItemView) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:127
// index:0
// void QAbstractItemView(class QWidget *)
func NewQAbstractItemView(parent unsafe.Pointer) *QAbstractItemView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemViewC1EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemViewFromPointer(cthis)
	return gothis
}
func NewQAbstractItemViewFromPointer(cthis unsafe.Pointer) *QAbstractItemView {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QAbstractItemView{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:273
// index:1
// void QAbstractItemView(class QAbstractItemViewPrivate &, class QWidget *)
func NewQAbstractItemView_1(arg0 unsafe.Pointer, parent unsafe.Pointer) *QAbstractItemView {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemViewC1ER24QAbstractItemViewPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:128
// index:0
// virtual
// void ~QAbstractItemView()
func DeleteQAbstractItemView(*QAbstractItemView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:130
// index:0
// virtual
// void setModel(class QAbstractItemModel *)
func (this *QAbstractItemView) SetModel(model unsafe.Pointer) {
	// 0: (, model QAbstractItemModel *), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:131
// index:0
// QAbstractItemModel * model()
func (this *QAbstractItemView) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView5modelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:133
// index:0
// virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QAbstractItemView) SetSelectionModel(selectionModel unsafe.Pointer) {
	// 0: (, selectionModel QItemSelectionModel *), (selectionModel)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:134
// index:0
// QItemSelectionModel * selectionModel()
func (this *QAbstractItemView) SelectionModel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView14selectionModelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:136
// index:0
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegate(delegate unsafe.Pointer) {
	// 0: (, delegate QAbstractItemDelegate *), (delegate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:137
// index:0
// QAbstractItemDelegate * itemDelegate()
func (this *QAbstractItemView) ItemDelegate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:226
// index:1
// QAbstractItemDelegate * itemDelegate(const class QModelIndex &)
func (this *QAbstractItemView) ItemDelegate_1(index unsafe.Pointer) {
	// 1: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:139
// index:0
// void setSelectionMode(class QAbstractItemView::SelectionMode)
func (this *QAbstractItemView) SetSelectionMode(mode int) {
	// 0: (, mode QAbstractItemView::SelectionMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16setSelectionModeENS_13SelectionModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:140
// index:0
// QAbstractItemView::SelectionMode selectionMode()
func (this *QAbstractItemView) SelectionMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13selectionModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:142
// index:0
// void setSelectionBehavior(class QAbstractItemView::SelectionBehavior)
func (this *QAbstractItemView) SetSelectionBehavior(behavior int) {
	// 0: (, behavior QAbstractItemView::SelectionBehavior), (&behavior)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20setSelectionBehaviorENS_17SelectionBehaviorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:143
// index:0
// QAbstractItemView::SelectionBehavior selectionBehavior()
func (this *QAbstractItemView) SelectionBehavior() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17selectionBehaviorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:145
// index:0
// QModelIndex currentIndex()
func (this *QAbstractItemView) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:146
// index:0
// QModelIndex rootIndex()
func (this *QAbstractItemView) RootIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView9rootIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:148
// index:0
// void setEditTriggers(QAbstractItemView::EditTriggers)
func (this *QAbstractItemView) SetEditTriggers(triggers int) {
	// 0: (, QFlags<QAbstractItemView::EditTrigger> triggers), (triggers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15setEditTriggersE6QFlagsINS_11EditTriggerEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), triggers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:149
// index:0
// QAbstractItemView::EditTriggers editTriggers()
func (this *QAbstractItemView) EditTriggers() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12editTriggersEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:151
// index:0
// void setVerticalScrollMode(enum QAbstractItemView::ScrollMode)
func (this *QAbstractItemView) SetVerticalScrollMode(mode int) {
	// 0: (, mode QAbstractItemView::ScrollMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21setVerticalScrollModeENS_10ScrollModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:152
// index:0
// QAbstractItemView::ScrollMode verticalScrollMode()
func (this *QAbstractItemView) VerticalScrollMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView18verticalScrollModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:153
// index:0
// void resetVerticalScrollMode()
func (this *QAbstractItemView) ResetVerticalScrollMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23resetVerticalScrollModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:155
// index:0
// void setHorizontalScrollMode(enum QAbstractItemView::ScrollMode)
func (this *QAbstractItemView) SetHorizontalScrollMode(mode int) {
	// 0: (, mode QAbstractItemView::ScrollMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23setHorizontalScrollModeENS_10ScrollModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:156
// index:0
// QAbstractItemView::ScrollMode horizontalScrollMode()
func (this *QAbstractItemView) HorizontalScrollMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView20horizontalScrollModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:157
// index:0
// void resetHorizontalScrollMode()
func (this *QAbstractItemView) ResetHorizontalScrollMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25resetHorizontalScrollModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:159
// index:0
// void setAutoScroll(_Bool)
func (this *QAbstractItemView) SetAutoScroll(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13setAutoScrollEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:160
// index:0
// bool hasAutoScroll()
func (this *QAbstractItemView) HasAutoScroll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13hasAutoScrollEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:162
// index:0
// void setAutoScrollMargin(int)
func (this *QAbstractItemView) SetAutoScrollMargin(margin int) {
	// 0: (, margin int), (&margin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView19setAutoScrollMarginEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:163
// index:0
// int autoScrollMargin()
func (this *QAbstractItemView) AutoScrollMargin() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16autoScrollMarginEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:165
// index:0
// void setTabKeyNavigation(_Bool)
func (this *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView19setTabKeyNavigationEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:166
// index:0
// bool tabKeyNavigation()
func (this *QAbstractItemView) TabKeyNavigation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16tabKeyNavigationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:169
// index:0
// void setDropIndicatorShown(_Bool)
func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21setDropIndicatorShownEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:170
// index:0
// bool showDropIndicator()
func (this *QAbstractItemView) ShowDropIndicator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17showDropIndicatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:172
// index:0
// void setDragEnabled(_Bool)
func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14setDragEnabledEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:173
// index:0
// bool dragEnabled()
func (this *QAbstractItemView) DragEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView11dragEnabledEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:175
// index:0
// void setDragDropOverwriteMode(_Bool)
func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	// 0: (, overwrite bool), (&overwrite)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView24setDragDropOverwriteModeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &overwrite)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:176
// index:0
// bool dragDropOverwriteMode()
func (this *QAbstractItemView) DragDropOverwriteMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView21dragDropOverwriteModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:187
// index:0
// void setDragDropMode(enum QAbstractItemView::DragDropMode)
func (this *QAbstractItemView) SetDragDropMode(behavior int) {
	// 0: (, behavior QAbstractItemView::DragDropMode), (&behavior)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15setDragDropModeENS_12DragDropModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:188
// index:0
// QAbstractItemView::DragDropMode dragDropMode()
func (this *QAbstractItemView) DragDropMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12dragDropModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:190
// index:0
// void setDefaultDropAction(Qt::DropAction)
func (this *QAbstractItemView) SetDefaultDropAction(dropAction int) {
	// 0: (, dropAction Qt::DropAction), (&dropAction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20setDefaultDropActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dropAction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:191
// index:0
// Qt::DropAction defaultDropAction()
func (this *QAbstractItemView) DefaultDropAction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17defaultDropActionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:194
// index:0
// void setAlternatingRowColors(_Bool)
func (this *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	// 0: (, enable bool), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23setAlternatingRowColorsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:195
// index:0
// bool alternatingRowColors()
func (this *QAbstractItemView) AlternatingRowColors() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView20alternatingRowColorsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:197
// index:0
// void setIconSize(const class QSize &)
func (this *QAbstractItemView) SetIconSize(size unsafe.Pointer) {
	// 0: (, size const QSize &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11setIconSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:198
// index:0
// QSize iconSize()
func (this *QAbstractItemView) IconSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView8iconSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:200
// index:0
// void setTextElideMode(Qt::TextElideMode)
func (this *QAbstractItemView) SetTextElideMode(mode int) {
	// 0: (, mode Qt::TextElideMode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16setTextElideModeEN2Qt13TextElideModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:201
// index:0
// Qt::TextElideMode textElideMode()
func (this *QAbstractItemView) TextElideMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13textElideModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:203
// index:0
// virtual
// void keyboardSearch(const class QString &)
func (this *QAbstractItemView) KeyboardSearch(search unsafe.Pointer) {
	// 0: (, search const QString &), (search)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14keyboardSearchERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), search)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:205
// index:0
// pure virtual
// QRect visualRect(const class QModelIndex &)
func (this *QAbstractItemView) VisualRect(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:206
// index:0
// pure virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QAbstractItemView) ScrollTo(index unsafe.Pointer, hint int) {
	// 0: (, index const QModelIndex &, hint QAbstractItemView::ScrollHint), (index, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView8scrollToERK11QModelIndexNS_10ScrollHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:207
// index:0
// pure virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QAbstractItemView) IndexAt(point unsafe.Pointer) {
	// 0: (, point const QPoint &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView7indexAtERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:209
// index:0
// QSize sizeHintForIndex(const class QModelIndex &)
func (this *QAbstractItemView) SizeHintForIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:210
// index:0
// virtual
// int sizeHintForRow(int)
func (this *QAbstractItemView) SizeHintForRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView14sizeHintForRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:211
// index:0
// virtual
// int sizeHintForColumn(int)
func (this *QAbstractItemView) SizeHintForColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17sizeHintForColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:213
// index:0
// void openPersistentEditor(const class QModelIndex &)
func (this *QAbstractItemView) OpenPersistentEditor(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:214
// index:0
// void closePersistentEditor(const class QModelIndex &)
func (this *QAbstractItemView) ClosePersistentEditor(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:215
// index:0
// bool isPersistentEditorOpen(const class QModelIndex &)
func (this *QAbstractItemView) IsPersistentEditorOpen(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView22isPersistentEditorOpenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:217
// index:0
// void setIndexWidget(const class QModelIndex &, class QWidget *)
func (this *QAbstractItemView) SetIndexWidget(index unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, index const QModelIndex &, widget QWidget *), (index, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:218
// index:0
// QWidget * indexWidget(const class QModelIndex &)
func (this *QAbstractItemView) IndexWidget(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView11indexWidgetERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:220
// index:0
// void setItemDelegateForRow(int, class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForRow(row int, delegate unsafe.Pointer) {
	// 0: (, row int, delegate QAbstractItemDelegate *), (&row, delegate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row, delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:221
// index:0
// QAbstractItemDelegate * itemDelegateForRow(int)
func (this *QAbstractItemView) ItemDelegateForRow(row int) {
	// 0: (, row int), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView18itemDelegateForRowEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:223
// index:0
// void setItemDelegateForColumn(int, class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForColumn(column int, delegate unsafe.Pointer) {
	// 0: (, column int, delegate QAbstractItemDelegate *), (&column, delegate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column, delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:224
// index:0
// QAbstractItemDelegate * itemDelegateForColumn(int)
func (this *QAbstractItemView) ItemDelegateForColumn(column int) {
	// 0: (, column int), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView21itemDelegateForColumnEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:228
// index:0
// virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QAbstractItemView) InputMethodQuery(query int) {
	// 0: (, query Qt::InputMethodQuery), (&query)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:233
// index:0
// virtual
// void reset()
func (this *QAbstractItemView) Reset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView5resetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:234
// index:0
// virtual
// void setRootIndex(const class QModelIndex &)
func (this *QAbstractItemView) SetRootIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:235
// index:0
// virtual
// void doItemsLayout()
func (this *QAbstractItemView) DoItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13doItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:236
// index:0
// virtual
// void selectAll()
func (this *QAbstractItemView) SelectAll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView9selectAllEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:237
// index:0
// void edit(const class QModelIndex &)
func (this *QAbstractItemView) Edit(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:295
// index:1
// virtual
// bool edit(const class QModelIndex &, enum QAbstractItemView::EditTrigger, class QEvent *)
func (this *QAbstractItemView) Edit_1(index unsafe.Pointer, trigger int, event unsafe.Pointer) {
	// 1: (, index const QModelIndex &, trigger QAbstractItemView::EditTrigger, event QEvent *), (index, &trigger, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, &trigger, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:238
// index:0
// void clearSelection()
func (this *QAbstractItemView) ClearSelection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14clearSelectionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:239
// index:0
// void setCurrentIndex(const class QModelIndex &)
func (this *QAbstractItemView) SetCurrentIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:240
// index:0
// void scrollToTop()
func (this *QAbstractItemView) ScrollToTop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11scrollToTopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:241
// index:0
// void scrollToBottom()
func (this *QAbstractItemView) ScrollToBottom() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14scrollToBottomEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:242
// index:0
// void update(const class QModelIndex &)
func (this *QAbstractItemView) Update(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView6updateERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:246
// index:0
// virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QAbstractItemView) RowsInserted(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:247
// index:0
// virtual
// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QAbstractItemView) RowsAboutToBeRemoved(parent unsafe.Pointer, start int, end int) {
	// 0: (, parent const QModelIndex &, start int, end int), (parent, &start, &end)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), parent, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:248
// index:0
// virtual
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QAbstractItemView) SelectionChanged(selected unsafe.Pointer, deselected unsafe.Pointer) {
	// 0: (, selected const QItemSelection &, deselected const QItemSelection &), (selected, deselected)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selected, deselected)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:249
// index:0
// virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QAbstractItemView) CurrentChanged(current unsafe.Pointer, previous unsafe.Pointer) {
	// 0: (, current const QModelIndex &, previous const QModelIndex &), (current, previous)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), current, previous)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:250
// index:0
// virtual
// void updateEditorData()
func (this *QAbstractItemView) UpdateEditorData() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16updateEditorDataEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:251
// index:0
// virtual
// void updateEditorGeometries()
func (this *QAbstractItemView) UpdateEditorGeometries() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView22updateEditorGeometriesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:252
// index:0
// virtual
// void updateGeometries()
func (this *QAbstractItemView) UpdateGeometries() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16updateGeometriesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:253
// index:0
// virtual
// void verticalScrollbarAction(int)
func (this *QAbstractItemView) VerticalScrollbarAction(action int) {
	// 0: (, action int), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23verticalScrollbarActionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:254
// index:0
// virtual
// void horizontalScrollbarAction(int)
func (this *QAbstractItemView) HorizontalScrollbarAction(action int) {
	// 0: (, action int), (&action)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25horizontalScrollbarActionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:255
// index:0
// virtual
// void verticalScrollbarValueChanged(int)
func (this *QAbstractItemView) VerticalScrollbarValueChanged(value int) {
	// 0: (, value int), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView29verticalScrollbarValueChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:256
// index:0
// virtual
// void horizontalScrollbarValueChanged(int)
func (this *QAbstractItemView) HorizontalScrollbarValueChanged(value int) {
	// 0: (, value int), (&value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:257
// index:0
// virtual
// void closeEditor(class QWidget *, class QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemView) CloseEditor(editor unsafe.Pointer, hint int) {
	// 0: (, editor QWidget *, hint QAbstractItemDelegate::EndEditHint), (editor, &hint)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:258
// index:0
// virtual
// void commitData(class QWidget *)
func (this *QAbstractItemView) CommitData(editor unsafe.Pointer) {
	// 0: (, editor QWidget *), (editor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView10commitDataEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:259
// index:0
// virtual
// void editorDestroyed(class QObject *)
func (this *QAbstractItemView) EditorDestroyed(editor unsafe.Pointer) {
	// 0: (, editor QObject *), (editor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15editorDestroyedEP7QObject", ffiqt.FFI_TYPE_VOID, this.GetCthis(), editor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:262
// index:0
// void pressed(const class QModelIndex &)
func (this *QAbstractItemView) Pressed(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView7pressedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:263
// index:0
// void clicked(const class QModelIndex &)
func (this *QAbstractItemView) Clicked(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView7clickedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:264
// index:0
// void doubleClicked(const class QModelIndex &)
func (this *QAbstractItemView) DoubleClicked(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13doubleClickedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:266
// index:0
// void activated(const class QModelIndex &)
func (this *QAbstractItemView) Activated(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView9activatedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:267
// index:0
// void entered(const class QModelIndex &)
func (this *QAbstractItemView) Entered(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView7enteredERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:268
// index:0
// void viewportEntered()
func (this *QAbstractItemView) ViewportEntered() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15viewportEnteredEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:270
// index:0
// void iconSizeChanged(const class QSize &)
func (this *QAbstractItemView) IconSizeChanged(size unsafe.Pointer) {
	// 0: (, size const QSize &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15iconSizeChangedERK5QSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:275
// index:0
// void setHorizontalStepsPerItem(int)
func (this *QAbstractItemView) SetHorizontalStepsPerItem(steps int) {
	// 0: (, steps int), (&steps)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25setHorizontalStepsPerItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:276
// index:0
// int horizontalStepsPerItem()
func (this *QAbstractItemView) HorizontalStepsPerItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView22horizontalStepsPerItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:277
// index:0
// void setVerticalStepsPerItem(int)
func (this *QAbstractItemView) SetVerticalStepsPerItem(steps int) {
	// 0: (, steps int), (&steps)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23setVerticalStepsPerItemEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:278
// index:0
// int verticalStepsPerItem()
func (this *QAbstractItemView) VerticalStepsPerItem() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView20verticalStepsPerItemEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:283
// index:0
// pure virtual
// QModelIndex moveCursor(enum QAbstractItemView::CursorAction, Qt::KeyboardModifiers)
func (this *QAbstractItemView) MoveCursor(cursorAction int, modifiers int) {
	// 0: (, cursorAction QAbstractItemView::CursorAction, QFlags<Qt::KeyboardModifier> modifiers), (&cursorAction, &modifiers)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView10moveCursorENS_12CursorActionE6QFlagsIN2Qt16KeyboardModifierEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cursorAction, &modifiers)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:286
// index:0
// pure virtual
// int horizontalOffset()
func (this *QAbstractItemView) HorizontalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16horizontalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:287
// index:0
// pure virtual
// int verticalOffset()
func (this *QAbstractItemView) VerticalOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView14verticalOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:289
// index:0
// pure virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QAbstractItemView) IsIndexHidden(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:291
// index:0
// pure virtual
// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
func (this *QAbstractItemView) SetSelection(rect unsafe.Pointer, command int) {
	// 0: (, rect const QRect &, QFlags<QItemSelectionModel::SelectionFlag> command), (rect, &command)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), rect, &command)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:292
// index:0
// pure virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QAbstractItemView) VisualRegionForSelection(selection unsafe.Pointer) {
	// 0: (, selection const QItemSelection &), (selection)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_VOID, this.GetCthis(), selection)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:293
// index:0
// virtual
// QModelIndexList selectedIndexes()
func (this *QAbstractItemView) SelectedIndexes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView15selectedIndexesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:297
// index:0
// virtual
// QItemSelectionModel::SelectionFlags selectionCommand(const class QModelIndex &, const class QEvent *)
func (this *QAbstractItemView) SelectionCommand(index unsafe.Pointer, event unsafe.Pointer) {
	// 0: (, index const QModelIndex &, event const QEvent *), (index, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index, event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:301
// index:0
// virtual
// void startDrag(Qt::DropActions)
func (this *QAbstractItemView) StartDrag(supportedActions int) {
	// 0: (, QFlags<Qt::DropAction> supportedActions), (&supportedActions)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView9startDragE6QFlagsIN2Qt10DropActionEE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &supportedActions)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:304
// index:0
// virtual
// QStyleOptionViewItem viewOptions()
func (this *QAbstractItemView) ViewOptions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView11viewOptionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:316
// index:0
// QAbstractItemView::State state()
func (this *QAbstractItemView) State() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView5stateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:317
// index:0
// void setState(enum QAbstractItemView::State)
func (this *QAbstractItemView) SetState(state int) {
	// 0: (, state QAbstractItemView::State), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView8setStateENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:319
// index:0
// void scheduleDelayedItemsLayout()
func (this *QAbstractItemView) ScheduleDelayedItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:320
// index:0
// void executeDelayedItemsLayout()
func (this *QAbstractItemView) ExecuteDelayedItemsLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25executeDelayedItemsLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:322
// index:0
// void setDirtyRegion(const class QRegion &)
func (this *QAbstractItemView) SetDirtyRegion(region unsafe.Pointer) {
	// 0: (, region const QRegion &), (region)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14setDirtyRegionERK7QRegion", ffiqt.FFI_TYPE_VOID, this.GetCthis(), region)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:323
// index:0
// void scrollDirtyRegion(int, int)
func (this *QAbstractItemView) ScrollDirtyRegion(dx int, dy int) {
	// 0: (, dx int, dy int), (&dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView17scrollDirtyRegionEii", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:324
// index:0
// QPoint dirtyRegionOffset()
func (this *QAbstractItemView) DirtyRegionOffset() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17dirtyRegionOffsetEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:326
// index:0
// void startAutoScroll()
func (this *QAbstractItemView) StartAutoScroll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15startAutoScrollEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:327
// index:0
// void stopAutoScroll()
func (this *QAbstractItemView) StopAutoScroll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14stopAutoScrollEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:328
// index:0
// void doAutoScroll()
func (this *QAbstractItemView) DoAutoScroll() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12doAutoScrollEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:330
// index:0
// virtual
// bool focusNextPrevChild(_Bool)
func (this *QAbstractItemView) FocusNextPrevChild(next bool) {
	// 0: (, next bool), (&next)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView18focusNextPrevChildEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:331
// index:0
// virtual
// bool event(class QEvent *)
func (this *QAbstractItemView) Event(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:332
// index:0
// virtual
// bool viewportEvent(class QEvent *)
func (this *QAbstractItemView) ViewportEvent(event unsafe.Pointer) {
	// 0: (, event QEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:333
// index:0
// virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractItemView) MousePressEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:334
// index:0
// virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractItemView) MouseMoveEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:335
// index:0
// virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractItemView) MouseReleaseEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:336
// index:0
// virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QAbstractItemView) MouseDoubleClickEvent(event unsafe.Pointer) {
	// 0: (, event QMouseEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:338
// index:0
// virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QAbstractItemView) DragEnterEvent(event unsafe.Pointer) {
	// 0: (, event QDragEnterEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:339
// index:0
// virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QAbstractItemView) DragMoveEvent(event unsafe.Pointer) {
	// 0: (, event QDragMoveEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:340
// index:0
// virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QAbstractItemView) DragLeaveEvent(event unsafe.Pointer) {
	// 0: (, event QDragLeaveEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:341
// index:0
// virtual
// void dropEvent(class QDropEvent *)
func (this *QAbstractItemView) DropEvent(event unsafe.Pointer) {
	// 0: (, event QDropEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:343
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QAbstractItemView) FocusInEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:344
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractItemView) FocusOutEvent(event unsafe.Pointer) {
	// 0: (, event QFocusEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:345
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractItemView) KeyPressEvent(event unsafe.Pointer) {
	// 0: (, event QKeyEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:346
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QAbstractItemView) ResizeEvent(event unsafe.Pointer) {
	// 0: (, event QResizeEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:347
// index:0
// virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractItemView) TimerEvent(event unsafe.Pointer) {
	// 0: (, event QTimerEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:348
// index:0
// virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QAbstractItemView) InputMethodEvent(event unsafe.Pointer) {
	// 0: (, event QInputMethodEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:352
// index:0
// QAbstractItemView::DropIndicatorPosition dropIndicatorPosition()
func (this *QAbstractItemView) DropIndicatorPosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView21dropIndicatorPositionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:355
// index:0
// virtual
// QSize viewportSizeHint()
func (this *QAbstractItemView) ViewportSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16viewportSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
