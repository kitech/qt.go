package qtwidgets

// /usr/include/qt/QtWidgets/qcompleter.h
// #include <qcompleter.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
func NewQCompleterFromPointer(cthis unsafe.Pointer) *QCompleter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QCompleter{bcthis0}
}

// /usr/include/qt/QtWidgets/qcompleter.h:61
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QCompleter) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:85
// index:0
// Public
// void QCompleter(class QObject *)
func NewQCompleter(parent *qtcore.QObject /*444 QObject **/) *QCompleter {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:86
// index:1
// Public
// void QCompleter(class QAbstractItemModel *, class QObject *)
func NewQCompleter_1(model *qtcore.QAbstractItemModel /*444 QAbstractItemModel **/, parent *qtcore.QObject /*444 QObject **/) *QCompleter {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = model.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterC2EP18QAbstractItemModelP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:88
// index:2
// Public
// void QCompleter(const class QStringList &, class QObject *)
func NewQCompleter_2(completions *qtcore.QStringList, parent *qtcore.QObject /*444 QObject **/) *QCompleter {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = completions.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterC2ERK11QStringListP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:90
// index:0
// Public virtual
// void ~QCompleter()
func DeleteQCompleter(*QCompleter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:92
// index:0
// Public
// void setWidget(class QWidget *)
func (this *QCompleter) SetWidget(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:93
// index:0
// Public
// QWidget * widget()
func (this *QCompleter) Widget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:95
// index:0
// Public
// void setModel(class QAbstractItemModel *)
func (this *QCompleter) SetModel(c *qtcore.QAbstractItemModel /*444 QAbstractItemModel **/) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:96
// index:0
// Public
// QAbstractItemModel * model()
func (this *QCompleter) Model() *qtcore.QAbstractItemModel /*444 QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:98
// index:0
// Public
// void setCompletionMode(enum QCompleter::CompletionMode)
func (this *QCompleter) SetCompletionMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter17setCompletionModeENS_14CompletionModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:99
// index:0
// Public
// QCompleter::CompletionMode completionMode()
func (this *QCompleter) CompletionMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter14completionModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:102
// index:0
// Public
// Qt::MatchFlags filterMode()
func (this *QCompleter) FilterMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10filterModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:104
// index:0
// Public
// QAbstractItemView * popup()
func (this *QCompleter) Popup() *QAbstractItemView /*444 QAbstractItemView **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter5popupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:105
// index:0
// Public
// void setPopup(class QAbstractItemView *)
func (this *QCompleter) SetPopup(popup *QAbstractItemView /*444 QAbstractItemView **/) {
	var convArg0 = popup.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter8setPopupEP17QAbstractItemView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:107
// index:0
// Public
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QCompleter) SetCaseSensitivity(caseSensitivity int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &caseSensitivity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:108
// index:0
// Public
// Qt::CaseSensitivity caseSensitivity()
func (this *QCompleter) CaseSensitivity() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15caseSensitivityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:110
// index:0
// Public
// void setModelSorting(enum QCompleter::ModelSorting)
func (this *QCompleter) SetModelSorting(sorting int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter15setModelSortingENS_12ModelSortingE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &sorting)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:111
// index:0
// Public
// QCompleter::ModelSorting modelSorting()
func (this *QCompleter) ModelSorting() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter12modelSortingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:113
// index:0
// Public
// void setCompletionColumn(int)
func (this *QCompleter) SetCompletionColumn(column int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter19setCompletionColumnEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:114
// index:0
// Public
// int completionColumn()
func (this *QCompleter) CompletionColumn() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter16completionColumnEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcompleter.h:116
// index:0
// Public
// void setCompletionRole(int)
func (this *QCompleter) SetCompletionRole(role int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter17setCompletionRoleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:117
// index:0
// Public
// int completionRole()
func (this *QCompleter) CompletionRole() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter14completionRoleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcompleter.h:119
// index:0
// Public
// bool wrapAround()
func (this *QCompleter) WrapAround() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10wrapAroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:121
// index:0
// Public
// int maxVisibleItems()
func (this *QCompleter) MaxVisibleItems() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15maxVisibleItemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcompleter.h:122
// index:0
// Public
// void setMaxVisibleItems(int)
func (this *QCompleter) SetMaxVisibleItems(maxItems int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter18setMaxVisibleItemsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &maxItems)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:124
// index:0
// Public
// int completionCount()
func (this *QCompleter) CompletionCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15completionCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcompleter.h:125
// index:0
// Public
// bool setCurrentRow(int)
func (this *QCompleter) SetCurrentRow(row int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter13setCurrentRowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &row)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:126
// index:0
// Public
// int currentRow()
func (this *QCompleter) CurrentRow() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10currentRowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcompleter.h:128
// index:0
// Public
// QModelIndex currentIndex()
func (this *QCompleter) CurrentIndex() *qtcore.QModelIndex /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:129
// index:0
// Public
// QString currentCompletion()
func (this *QCompleter) CurrentCompletion() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter17currentCompletionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:131
// index:0
// Public
// QAbstractItemModel * completionModel()
func (this *QCompleter) CompletionModel() *qtcore.QAbstractItemModel /*444 QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15completionModelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:133
// index:0
// Public
// QString completionPrefix()
func (this *QCompleter) CompletionPrefix() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter16completionPrefixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:136
// index:0
// Public
// void setCompletionPrefix(const class QString &)
func (this *QCompleter) SetCompletionPrefix(prefix *qtcore.QString) {
	var convArg0 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter19setCompletionPrefixERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:137
// index:0
// Public
// void complete(const class QRect &)
func (this *QCompleter) Complete(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter8completeERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:138
// index:0
// Public
// void setWrapAround(_Bool)
func (this *QCompleter) SetWrapAround(wrap bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter13setWrapAroundEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &wrap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:141
// index:0
// Public virtual
// QString pathFromIndex(const class QModelIndex &)
func (this *QCompleter) PathFromIndex(index *qtcore.QModelIndex) *qtcore.QString /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter13pathFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcompleter.h:145
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QCompleter) EventFilter(o *qtcore.QObject /*444 QObject **/, e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = o.GetCthis()
	var convArg1 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:146
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QCompleter) Event(arg0 *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:149
// index:0
// Public
// void activated(const class QString &)
func (this *QCompleter) Activated(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter9activatedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:150
// index:1
// Public
// void activated(const class QModelIndex &)
func (this *QCompleter) Activated_1(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter9activatedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:151
// index:0
// Public
// void highlighted(const class QString &)
func (this *QCompleter) Highlighted(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter11highlightedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:152
// index:1
// Public
// void highlighted(const class QModelIndex &)
func (this *QCompleter) Highlighted_1(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter11highlightedERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
