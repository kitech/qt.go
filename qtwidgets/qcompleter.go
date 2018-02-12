package qtwidgets

// /usr/include/qt/QtWidgets/qcompleter.h
// #include <qcompleter.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool eventFilter(class QObject *, class QEvent *)
func (this *QCompleter) InheritEventFilter(f func(o *qtcore.QObject /*777 QObject **/, e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(class QEvent *)
func (this *QCompleter) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QCompleter struct {
	*qtcore.QObject
}

func (this *QCompleter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QCompleter) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQCompleterFromPointer(cthis unsafe.Pointer) *QCompleter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QCompleter{bcthis0}
}
func (*QCompleter) NewFromPointer(cthis unsafe.Pointer) *QCompleter {
	return NewQCompleterFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcompleter.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QCompleter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(QObject *)
func NewQCompleter(parent *qtcore.QObject /*777 QObject **/) *QCompleter {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(QAbstractItemModel *, QObject *)
func NewQCompleter_1(model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/, parent *qtcore.QObject /*777 QObject **/) *QCompleter {
	var convArg0 = model.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2EP18QAbstractItemModelP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:88
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(const QStringList &, QObject *)
func NewQCompleter_2(completions *qtcore.QStringList, parent *qtcore.QObject /*777 QObject **/) *QCompleter {
	var convArg0 = completions.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCompleter()
func DeleteQCompleter(this *QCompleter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcompleter.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)
func (this *QCompleter) SetWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget()
func (this *QCompleter) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)
func (this *QCompleter) SetModel(c *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = c.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model()
func (this *QCompleter) Model() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompletionMode(enum QCompleter::CompletionMode)
func (this *QCompleter) SetCompletionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter17setCompletionModeENS_14CompletionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QCompleter::CompletionMode completionMode()
func (this *QCompleter) CompletionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter14completionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterMode(Qt::MatchFlags)
func (this *QCompleter) SetFilterMode(filterMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter13setFilterModeE6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filterMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MatchFlags filterMode()
func (this *QCompleter) FilterMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10filterModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemView * popup()
func (this *QCompleter) Popup() *QAbstractItemView /*777 QAbstractItemView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter5popupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPopup(QAbstractItemView *)
func (this *QCompleter) SetPopup(popup *QAbstractItemView /*777 QAbstractItemView **/) {
	var convArg0 = popup.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter8setPopupEP17QAbstractItemView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QCompleter) SetCaseSensitivity(caseSensitivity int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter18setCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), caseSensitivity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity caseSensitivity()
func (this *QCompleter) CaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15caseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModelSorting(enum QCompleter::ModelSorting)
func (this *QCompleter) SetModelSorting(sorting int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter15setModelSortingENS_12ModelSortingE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sorting)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QCompleter::ModelSorting modelSorting()
func (this *QCompleter) ModelSorting() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter12modelSortingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompletionColumn(int)
func (this *QCompleter) SetCompletionColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter19setCompletionColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int completionColumn()
func (this *QCompleter) CompletionColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter16completionColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompletionRole(int)
func (this *QCompleter) SetCompletionRole(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter17setCompletionRoleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int completionRole()
func (this *QCompleter) CompletionRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter14completionRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wrapAround()
func (this *QCompleter) WrapAround() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10wrapAroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxVisibleItems()
func (this *QCompleter) MaxVisibleItems() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15maxVisibleItemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxVisibleItems(int)
func (this *QCompleter) SetMaxVisibleItems(maxItems int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter18setMaxVisibleItemsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxItems)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] int completionCount()
func (this *QCompleter) CompletionCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15completionCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setCurrentRow(int)
func (this *QCompleter) SetCurrentRow(row int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter13setCurrentRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentRow()
func (this *QCompleter) CurrentRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10currentRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:128
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex currentIndex()
func (this *QCompleter) CurrentIndex() *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QString currentCompletion()
func (this *QCompleter) CurrentCompletion() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter17currentCompletionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcompleter.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * completionModel()
func (this *QCompleter) CompletionModel() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15completionModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QString completionPrefix()
func (this *QCompleter) CompletionPrefix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter16completionPrefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcompleter.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompletionPrefix(const QString &)
func (this *QCompleter) SetCompletionPrefix(prefix string) {
	var tmpArg0 = qtcore.NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter19setCompletionPrefixERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void complete(const QRect &)
func (this *QCompleter) Complete(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter8completeERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapAround(_Bool)
func (this *QCompleter) SetWrapAround(wrap bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter13setWrapAroundEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), wrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString pathFromIndex(const QModelIndex &)
func (this *QCompleter) PathFromIndex(index *qtcore.QModelIndex) string {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter13pathFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcompleter.h:142
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList splitPath(const QString &)
func (this *QCompleter) SplitPath(path string) *qtcore.QStringList /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter9splitPathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:145
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)
func (this *QCompleter) EventFilter(o *qtcore.QObject /*777 QObject **/, e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = o.GetCthis()
	var convArg1 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:146
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QCompleter) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(const QString &)
func (this *QCompleter) Activated(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter9activatedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:150
// index:1
// Public Visibility=Default Availability=Available
// [-2] void activated(const QModelIndex &)
func (this *QCompleter) Activated_1(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter9activatedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QString &)
func (this *QCompleter) Highlighted(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter11highlightedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:152
// index:1
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QModelIndex &)
func (this *QCompleter) Highlighted_1(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter11highlightedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QCompleter__CompletionMode = int

const QCompleter__PopupCompletion QCompleter__CompletionMode = 0
const QCompleter__UnfilteredPopupCompletion QCompleter__CompletionMode = 1
const QCompleter__InlineCompletion QCompleter__CompletionMode = 2

type QCompleter__ModelSorting = int

const QCompleter__UnsortedModel QCompleter__ModelSorting = 0
const QCompleter__CaseSensitivelySortedModel QCompleter__ModelSorting = 1
const QCompleter__CaseInsensitivelySortedModel QCompleter__ModelSorting = 2

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
