//  header block begin
// /usr/include/qt/QtWidgets/qabstractitemview.h
// #include <qabstractitemview.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 42
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
func NewQAbstractItemViewFromPointer(cthis unsafe.Pointer) *QAbstractItemView {
	bcthis0 := NewQAbstractScrollAreaFromPointer(cthis)
	return &QAbstractItemView{bcthis0}
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:63
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QAbstractItemView) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:127
// index:0
// Public
// void QAbstractItemView(class QWidget *)
func NewQAbstractItemView(parent unsafe.Pointer) *QAbstractItemView {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemViewC1EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractItemViewFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:128
// index:0
// Public virtual
// void ~QAbstractItemView()
func DeleteQAbstractItemView(*QAbstractItemView) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemViewD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:130
// index:0
// Public virtual
// void setModel(class QAbstractItemModel *)
func (this *QAbstractItemView) SetModel(model unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:131
// index:0
// Public
// QAbstractItemModel * model()
func (this *QAbstractItemView) Model() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:133
// index:0
// Public virtual
// void setSelectionModel(class QItemSelectionModel *)
func (this *QAbstractItemView) SetSelectionModel(selectionModel unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView17setSelectionModelEP19QItemSelectionModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), selectionModel)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:134
// index:0
// Public
// QItemSelectionModel * selectionModel()
func (this *QAbstractItemView) SelectionModel() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView14selectionModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:136
// index:0
// Public
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegate(delegate unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:137
// index:0
// Public
// QAbstractItemDelegate * itemDelegate()
func (this *QAbstractItemView) ItemDelegate() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:226
// index:1
// Public
// QAbstractItemDelegate * itemDelegate(const class QModelIndex &)
func (this *QAbstractItemView) ItemDelegate_1(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12itemDelegateERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:139
// index:0
// Public
// void setSelectionMode(class QAbstractItemView::SelectionMode)
func (this *QAbstractItemView) SetSelectionMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16setSelectionModeENS_13SelectionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:140
// index:0
// Public
// QAbstractItemView::SelectionMode selectionMode()
func (this *QAbstractItemView) SelectionMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13selectionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:142
// index:0
// Public
// void setSelectionBehavior(class QAbstractItemView::SelectionBehavior)
func (this *QAbstractItemView) SetSelectionBehavior(behavior int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20setSelectionBehaviorENS_17SelectionBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:143
// index:0
// Public
// QAbstractItemView::SelectionBehavior selectionBehavior()
func (this *QAbstractItemView) SelectionBehavior() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17selectionBehaviorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:145
// index:0
// Public
// QModelIndex currentIndex()
func (this *QAbstractItemView) CurrentIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:146
// index:0
// Public
// QModelIndex rootIndex()
func (this *QAbstractItemView) RootIndex() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView9rootIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:149
// index:0
// Public
// QAbstractItemView::EditTriggers editTriggers()
func (this *QAbstractItemView) EditTriggers() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12editTriggersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:151
// index:0
// Public
// void setVerticalScrollMode(enum QAbstractItemView::ScrollMode)
func (this *QAbstractItemView) SetVerticalScrollMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21setVerticalScrollModeENS_10ScrollModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:152
// index:0
// Public
// QAbstractItemView::ScrollMode verticalScrollMode()
func (this *QAbstractItemView) VerticalScrollMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView18verticalScrollModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:153
// index:0
// Public
// void resetVerticalScrollMode()
func (this *QAbstractItemView) ResetVerticalScrollMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23resetVerticalScrollModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:155
// index:0
// Public
// void setHorizontalScrollMode(enum QAbstractItemView::ScrollMode)
func (this *QAbstractItemView) SetHorizontalScrollMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23setHorizontalScrollModeENS_10ScrollModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:156
// index:0
// Public
// QAbstractItemView::ScrollMode horizontalScrollMode()
func (this *QAbstractItemView) HorizontalScrollMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView20horizontalScrollModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:157
// index:0
// Public
// void resetHorizontalScrollMode()
func (this *QAbstractItemView) ResetHorizontalScrollMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25resetHorizontalScrollModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:159
// index:0
// Public
// void setAutoScroll(_Bool)
func (this *QAbstractItemView) SetAutoScroll(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13setAutoScrollEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:160
// index:0
// Public
// bool hasAutoScroll()
func (this *QAbstractItemView) HasAutoScroll() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13hasAutoScrollEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:162
// index:0
// Public
// void setAutoScrollMargin(int)
func (this *QAbstractItemView) SetAutoScrollMargin(margin int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView19setAutoScrollMarginEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &margin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:163
// index:0
// Public
// int autoScrollMargin()
func (this *QAbstractItemView) AutoScrollMargin() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16autoScrollMarginEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:165
// index:0
// Public
// void setTabKeyNavigation(_Bool)
func (this *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView19setTabKeyNavigationEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:166
// index:0
// Public
// bool tabKeyNavigation()
func (this *QAbstractItemView) TabKeyNavigation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16tabKeyNavigationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:169
// index:0
// Public
// void setDropIndicatorShown(_Bool)
func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21setDropIndicatorShownEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:170
// index:0
// Public
// bool showDropIndicator()
func (this *QAbstractItemView) ShowDropIndicator() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17showDropIndicatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:172
// index:0
// Public
// void setDragEnabled(_Bool)
func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14setDragEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:173
// index:0
// Public
// bool dragEnabled()
func (this *QAbstractItemView) DragEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView11dragEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:175
// index:0
// Public
// void setDragDropOverwriteMode(_Bool)
func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView24setDragDropOverwriteModeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &overwrite)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:176
// index:0
// Public
// bool dragDropOverwriteMode()
func (this *QAbstractItemView) DragDropOverwriteMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView21dragDropOverwriteModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:187
// index:0
// Public
// void setDragDropMode(enum QAbstractItemView::DragDropMode)
func (this *QAbstractItemView) SetDragDropMode(behavior int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15setDragDropModeENS_12DragDropModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:188
// index:0
// Public
// QAbstractItemView::DragDropMode dragDropMode()
func (this *QAbstractItemView) DragDropMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView12dragDropModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:190
// index:0
// Public
// void setDefaultDropAction(Qt::DropAction)
func (this *QAbstractItemView) SetDefaultDropAction(dropAction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20setDefaultDropActionEN2Qt10DropActionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dropAction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:191
// index:0
// Public
// Qt::DropAction defaultDropAction()
func (this *QAbstractItemView) DefaultDropAction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17defaultDropActionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:194
// index:0
// Public
// void setAlternatingRowColors(_Bool)
func (this *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23setAlternatingRowColorsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:195
// index:0
// Public
// bool alternatingRowColors()
func (this *QAbstractItemView) AlternatingRowColors() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView20alternatingRowColorsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:197
// index:0
// Public
// void setIconSize(const class QSize &)
func (this *QAbstractItemView) SetIconSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11setIconSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:198
// index:0
// Public
// QSize iconSize()
func (this *QAbstractItemView) IconSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView8iconSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:200
// index:0
// Public
// void setTextElideMode(Qt::TextElideMode)
func (this *QAbstractItemView) SetTextElideMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16setTextElideModeEN2Qt13TextElideModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:201
// index:0
// Public
// Qt::TextElideMode textElideMode()
func (this *QAbstractItemView) TextElideMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13textElideModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:203
// index:0
// Public virtual
// void keyboardSearch(const class QString &)
func (this *QAbstractItemView) KeyboardSearch(search *qtcore.QString) {
	var convArg0 = search.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14keyboardSearchERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:205
// index:0
// Public pure virtual
// QRect visualRect(const class QModelIndex &)
func (this *QAbstractItemView) VisualRect(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView10visualRectERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:206
// index:0
// Public pure virtual
// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
func (this *QAbstractItemView) ScrollTo(index *qtcore.QModelIndex, hint int) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView8scrollToERK11QModelIndexNS_10ScrollHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:207
// index:0
// Public pure virtual
// QModelIndex indexAt(const class QPoint &)
func (this *QAbstractItemView) IndexAt(point *qtcore.QPoint) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView7indexAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:209
// index:0
// Public
// QSize sizeHintForIndex(const class QModelIndex &)
func (this *QAbstractItemView) SizeHintForIndex(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16sizeHintForIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:210
// index:0
// Public virtual
// int sizeHintForRow(int)
func (this *QAbstractItemView) SizeHintForRow(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView14sizeHintForRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:211
// index:0
// Public virtual
// int sizeHintForColumn(int)
func (this *QAbstractItemView) SizeHintForColumn(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17sizeHintForColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:213
// index:0
// Public
// void openPersistentEditor(const class QModelIndex &)
func (this *QAbstractItemView) OpenPersistentEditor(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20openPersistentEditorERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:214
// index:0
// Public
// void closePersistentEditor(const class QModelIndex &)
func (this *QAbstractItemView) ClosePersistentEditor(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21closePersistentEditorERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:215
// index:0
// Public
// bool isPersistentEditorOpen(const class QModelIndex &)
func (this *QAbstractItemView) IsPersistentEditorOpen(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView22isPersistentEditorOpenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:217
// index:0
// Public
// void setIndexWidget(const class QModelIndex &, class QWidget *)
func (this *QAbstractItemView) SetIndexWidget(index *qtcore.QModelIndex, widget unsafe.Pointer) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14setIndexWidgetERK11QModelIndexP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:218
// index:0
// Public
// QWidget * indexWidget(const class QModelIndex &)
func (this *QAbstractItemView) IndexWidget(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView11indexWidgetERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:220
// index:0
// Public
// void setItemDelegateForRow(int, class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForRow(row int, delegate unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21setItemDelegateForRowEiP21QAbstractItemDelegate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row, delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:221
// index:0
// Public
// QAbstractItemDelegate * itemDelegateForRow(int)
func (this *QAbstractItemView) ItemDelegateForRow(row int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView18itemDelegateForRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:223
// index:0
// Public
// void setItemDelegateForColumn(int, class QAbstractItemDelegate *)
func (this *QAbstractItemView) SetItemDelegateForColumn(column int, delegate unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView24setItemDelegateForColumnEiP21QAbstractItemDelegate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column, delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:224
// index:0
// Public
// QAbstractItemDelegate * itemDelegateForColumn(int)
func (this *QAbstractItemView) ItemDelegateForColumn(column int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView21itemDelegateForColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:228
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QAbstractItemView) InputMethodQuery(query int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:233
// index:0
// Public virtual
// void reset()
func (this *QAbstractItemView) Reset() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView5resetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:234
// index:0
// Public virtual
// void setRootIndex(const class QModelIndex &)
func (this *QAbstractItemView) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:235
// index:0
// Public virtual
// void doItemsLayout()
func (this *QAbstractItemView) DoItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13doItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:236
// index:0
// Public virtual
// void selectAll()
func (this *QAbstractItemView) SelectAll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView9selectAllEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:237
// index:0
// Public
// void edit(const class QModelIndex &)
func (this *QAbstractItemView) Edit(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:295
// index:1
// Protected virtual
// bool edit(const class QModelIndex &, enum QAbstractItemView::EditTrigger, class QEvent *)
func (this *QAbstractItemView) Edit_1(index *qtcore.QModelIndex, trigger int, event unsafe.Pointer) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &trigger, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:238
// index:0
// Public
// void clearSelection()
func (this *QAbstractItemView) ClearSelection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14clearSelectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:239
// index:0
// Public
// void setCurrentIndex(const class QModelIndex &)
func (this *QAbstractItemView) SetCurrentIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15setCurrentIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:240
// index:0
// Public
// void scrollToTop()
func (this *QAbstractItemView) ScrollToTop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11scrollToTopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:241
// index:0
// Public
// void scrollToBottom()
func (this *QAbstractItemView) ScrollToBottom() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14scrollToBottomEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:242
// index:0
// Public
// void update(const class QModelIndex &)
func (this *QAbstractItemView) Update(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView6updateERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:246
// index:0
// Protected virtual
// void rowsInserted(const class QModelIndex &, int, int)
func (this *QAbstractItemView) RowsInserted(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:247
// index:0
// Protected virtual
// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
func (this *QAbstractItemView) RowsAboutToBeRemoved(parent *qtcore.QModelIndex, start int, end int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &start, &end)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:248
// index:0
// Protected virtual
// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
func (this *QAbstractItemView) SelectionChanged(selected *qtcore.QItemSelection, deselected *qtcore.QItemSelection) {
	var convArg0 = selected.GetCthis()
	var convArg1 = deselected.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:249
// index:0
// Protected virtual
// void currentChanged(const class QModelIndex &, const class QModelIndex &)
func (this *QAbstractItemView) CurrentChanged(current *qtcore.QModelIndex, previous *qtcore.QModelIndex) {
	var convArg0 = current.GetCthis()
	var convArg1 = previous.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:250
// index:0
// Protected virtual
// void updateEditorData()
func (this *QAbstractItemView) UpdateEditorData() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16updateEditorDataEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:251
// index:0
// Protected virtual
// void updateEditorGeometries()
func (this *QAbstractItemView) UpdateEditorGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView22updateEditorGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:252
// index:0
// Protected virtual
// void updateGeometries()
func (this *QAbstractItemView) UpdateGeometries() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16updateGeometriesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:253
// index:0
// Protected virtual
// void verticalScrollbarAction(int)
func (this *QAbstractItemView) VerticalScrollbarAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23verticalScrollbarActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:254
// index:0
// Protected virtual
// void horizontalScrollbarAction(int)
func (this *QAbstractItemView) HorizontalScrollbarAction(action int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25horizontalScrollbarActionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &action)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:255
// index:0
// Protected virtual
// void verticalScrollbarValueChanged(int)
func (this *QAbstractItemView) VerticalScrollbarValueChanged(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView29verticalScrollbarValueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:256
// index:0
// Protected virtual
// void horizontalScrollbarValueChanged(int)
func (this *QAbstractItemView) HorizontalScrollbarValueChanged(value int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:257
// index:0
// Protected virtual
// void closeEditor(class QWidget *, class QAbstractItemDelegate::EndEditHint)
func (this *QAbstractItemView) CloseEditor(editor unsafe.Pointer, hint int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor, &hint)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:258
// index:0
// Protected virtual
// void commitData(class QWidget *)
func (this *QAbstractItemView) CommitData(editor unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView10commitDataEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:259
// index:0
// Protected virtual
// void editorDestroyed(class QObject *)
func (this *QAbstractItemView) EditorDestroyed(editor unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15editorDestroyedEP7QObject", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), editor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:262
// index:0
// Public
// void pressed(const class QModelIndex &)
func (this *QAbstractItemView) Pressed(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView7pressedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:263
// index:0
// Public
// void clicked(const class QModelIndex &)
func (this *QAbstractItemView) Clicked(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView7clickedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:264
// index:0
// Public
// void doubleClicked(const class QModelIndex &)
func (this *QAbstractItemView) DoubleClicked(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13doubleClickedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:266
// index:0
// Public
// void activated(const class QModelIndex &)
func (this *QAbstractItemView) Activated(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView9activatedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:267
// index:0
// Public
// void entered(const class QModelIndex &)
func (this *QAbstractItemView) Entered(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView7enteredERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:268
// index:0
// Public
// void viewportEntered()
func (this *QAbstractItemView) ViewportEntered() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15viewportEnteredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:270
// index:0
// Public
// void iconSizeChanged(const class QSize &)
func (this *QAbstractItemView) IconSizeChanged(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15iconSizeChangedERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:275
// index:0
// Protected
// void setHorizontalStepsPerItem(int)
func (this *QAbstractItemView) SetHorizontalStepsPerItem(steps int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25setHorizontalStepsPerItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:276
// index:0
// Protected
// int horizontalStepsPerItem()
func (this *QAbstractItemView) HorizontalStepsPerItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView22horizontalStepsPerItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:277
// index:0
// Protected
// void setVerticalStepsPerItem(int)
func (this *QAbstractItemView) SetVerticalStepsPerItem(steps int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView23setVerticalStepsPerItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &steps)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:278
// index:0
// Protected
// int verticalStepsPerItem()
func (this *QAbstractItemView) VerticalStepsPerItem() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView20verticalStepsPerItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:286
// index:0
// Protected pure virtual
// int horizontalOffset()
func (this *QAbstractItemView) HorizontalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16horizontalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:287
// index:0
// Protected pure virtual
// int verticalOffset()
func (this *QAbstractItemView) VerticalOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView14verticalOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:289
// index:0
// Protected pure virtual
// bool isIndexHidden(const class QModelIndex &)
func (this *QAbstractItemView) IsIndexHidden(index *qtcore.QModelIndex) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:292
// index:0
// Protected pure virtual
// QRegion visualRegionForSelection(const class QItemSelection &)
func (this *QAbstractItemView) VisualRegionForSelection(selection *qtcore.QItemSelection) interface{} {
	var convArg0 = selection.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:293
// index:0
// Protected virtual
// QModelIndexList selectedIndexes()
func (this *QAbstractItemView) SelectedIndexes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView15selectedIndexesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:297
// index:0
// Protected virtual
// QItemSelectionModel::SelectionFlags selectionCommand(const class QModelIndex &, const class QEvent *)
func (this *QAbstractItemView) SelectionCommand(index *qtcore.QModelIndex, event unsafe.Pointer) interface{} {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:304
// index:0
// Protected virtual
// QStyleOptionViewItem viewOptions()
func (this *QAbstractItemView) ViewOptions() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView11viewOptionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:316
// index:0
// Protected
// QAbstractItemView::State state()
func (this *QAbstractItemView) State() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView5stateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:317
// index:0
// Protected
// void setState(enum QAbstractItemView::State)
func (this *QAbstractItemView) SetState(state int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView8setStateENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:319
// index:0
// Protected
// void scheduleDelayedItemsLayout()
func (this *QAbstractItemView) ScheduleDelayedItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:320
// index:0
// Protected
// void executeDelayedItemsLayout()
func (this *QAbstractItemView) ExecuteDelayedItemsLayout() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView25executeDelayedItemsLayoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:322
// index:0
// Protected
// void setDirtyRegion(const class QRegion &)
func (this *QAbstractItemView) SetDirtyRegion(region *qtgui.QRegion) {
	var convArg0 = region.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14setDirtyRegionERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:323
// index:0
// Protected
// void scrollDirtyRegion(int, int)
func (this *QAbstractItemView) ScrollDirtyRegion(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView17scrollDirtyRegionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:324
// index:0
// Protected
// QPoint dirtyRegionOffset()
func (this *QAbstractItemView) DirtyRegionOffset() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView17dirtyRegionOffsetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:326
// index:0
// Protected
// void startAutoScroll()
func (this *QAbstractItemView) StartAutoScroll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15startAutoScrollEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:327
// index:0
// Protected
// void stopAutoScroll()
func (this *QAbstractItemView) StopAutoScroll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14stopAutoScrollEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:328
// index:0
// Protected
// void doAutoScroll()
func (this *QAbstractItemView) DoAutoScroll() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12doAutoScrollEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:330
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QAbstractItemView) FocusNextPrevChild(next bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:331
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QAbstractItemView) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:332
// index:0
// Protected virtual
// bool viewportEvent(class QEvent *)
func (this *QAbstractItemView) ViewportEvent(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13viewportEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:333
// index:0
// Protected virtual
// void mousePressEvent(class QMouseEvent *)
func (this *QAbstractItemView) MousePressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:334
// index:0
// Protected virtual
// void mouseMoveEvent(class QMouseEvent *)
func (this *QAbstractItemView) MouseMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:335
// index:0
// Protected virtual
// void mouseReleaseEvent(class QMouseEvent *)
func (this *QAbstractItemView) MouseReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:336
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QAbstractItemView) MouseDoubleClickEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:338
// index:0
// Protected virtual
// void dragEnterEvent(class QDragEnterEvent *)
func (this *QAbstractItemView) DragEnterEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:339
// index:0
// Protected virtual
// void dragMoveEvent(class QDragMoveEvent *)
func (this *QAbstractItemView) DragMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:340
// index:0
// Protected virtual
// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QAbstractItemView) DragLeaveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:341
// index:0
// Protected virtual
// void dropEvent(class QDropEvent *)
func (this *QAbstractItemView) DropEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:343
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QAbstractItemView) FocusInEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:344
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QAbstractItemView) FocusOutEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:345
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QAbstractItemView) KeyPressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:346
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QAbstractItemView) ResizeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:347
// index:0
// Protected virtual
// void timerEvent(class QTimerEvent *)
func (this *QAbstractItemView) TimerEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView10timerEventEP11QTimerEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:348
// index:0
// Protected virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QAbstractItemView) InputMethodEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:352
// index:0
// Protected
// QAbstractItemView::DropIndicatorPosition dropIndicatorPosition()
func (this *QAbstractItemView) DropIndicatorPosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView21dropIndicatorPositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qabstractitemview.h:355
// index:0
// Protected virtual
// QSize viewportSizeHint()
func (this *QAbstractItemView) ViewportSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QAbstractItemView16viewportSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
