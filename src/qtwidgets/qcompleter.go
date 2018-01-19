//  header block begin
// /usr/include/qt/QtWidgets/qcompleter.h
// #include <qcompleter.h>
// #include <QtWidgets>
package qtwidgets

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qcompleter.h:61
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCompleter) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:85
// index:0
// void QCompleter(class QObject *)
func NewQCompleter(parent unsafe.Pointer) *QCompleter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QCompleter{cthis}
}

// /usr/include/qt/QtWidgets/qcompleter.h:86
// index:1
// void QCompleter(class QAbstractItemModel *, class QObject *)
func NewQCompleter_1(model unsafe.Pointer, parent unsafe.Pointer) *QCompleter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterC2EP18QAbstractItemModelP7QObject", ffiqt.FFI_TYPE_VOID, cthis, model, parent)
	gopp.ErrPrint(err, rv)
	return &QCompleter{cthis}
}

// /usr/include/qt/QtWidgets/qcompleter.h:88
// index:2
// void QCompleter(const class QStringList &, class QObject *)
func NewQCompleter_2(completions unsafe.Pointer, parent unsafe.Pointer) *QCompleter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterC2ERK11QStringListP7QObject", ffiqt.FFI_TYPE_VOID, cthis, completions, parent)
	gopp.ErrPrint(err, rv)
	return &QCompleter{cthis}
}

// /usr/include/qt/QtWidgets/qcompleter.h:90
// index:0
// virtual
// void ~QCompleter()
func DeleteQCompleter(*QCompleter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:92
// index:0
// void setWidget(class QWidget *)
func (this *QCompleter) SetWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:93
// index:0
// QWidget * widget()
func (this *QCompleter) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter6widgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:95
// index:0
// void setModel(class QAbstractItemModel *)
func (this *QCompleter) SetModel(c unsafe.Pointer) {
	// 0: (, QAbstractItemModel * c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.cthis, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:96
// index:0
// QAbstractItemModel * model()
func (this *QCompleter) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter5modelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:98
// index:0
// void setCompletionMode(enum QCompleter::CompletionMode)
func (this *QCompleter) SetCompletionMode(mode int) {
	// 0: (, QCompleter::CompletionMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter17setCompletionModeENS_14CompletionModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:99
// index:0
// QCompleter::CompletionMode completionMode()
func (this *QCompleter) CompletionMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter14completionModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:102
// index:0
// Qt::MatchFlags filterMode()
func (this *QCompleter) FilterMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10filterModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:104
// index:0
// QAbstractItemView * popup()
func (this *QCompleter) Popup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter5popupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:105
// index:0
// void setPopup(class QAbstractItemView *)
func (this *QCompleter) SetPopup(popup unsafe.Pointer) {
	// 0: (, QAbstractItemView * popup), (popup)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter8setPopupEP17QAbstractItemView", ffiqt.FFI_TYPE_VOID, this.cthis, popup)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:107
// index:0
// void setCaseSensitivity(Qt::CaseSensitivity)
func (this *QCompleter) SetCaseSensitivity(caseSensitivity int) {
	// 0: (, Qt::CaseSensitivity caseSensitivity), (&caseSensitivity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter18setCaseSensitivityEN2Qt15CaseSensitivityE", ffiqt.FFI_TYPE_VOID, this.cthis, &caseSensitivity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:108
// index:0
// Qt::CaseSensitivity caseSensitivity()
func (this *QCompleter) CaseSensitivity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15caseSensitivityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:110
// index:0
// void setModelSorting(enum QCompleter::ModelSorting)
func (this *QCompleter) SetModelSorting(sorting int) {
	// 0: (, QCompleter::ModelSorting sorting), (&sorting)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter15setModelSortingENS_12ModelSortingE", ffiqt.FFI_TYPE_VOID, this.cthis, &sorting)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:111
// index:0
// QCompleter::ModelSorting modelSorting()
func (this *QCompleter) ModelSorting() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter12modelSortingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:113
// index:0
// void setCompletionColumn(int)
func (this *QCompleter) SetCompletionColumn(column int) {
	// 0: (, int column), (&column)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter19setCompletionColumnEi", ffiqt.FFI_TYPE_VOID, this.cthis, &column)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:114
// index:0
// int completionColumn()
func (this *QCompleter) CompletionColumn() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter16completionColumnEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:116
// index:0
// void setCompletionRole(int)
func (this *QCompleter) SetCompletionRole(role int) {
	// 0: (, int role), (&role)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter17setCompletionRoleEi", ffiqt.FFI_TYPE_VOID, this.cthis, &role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:117
// index:0
// int completionRole()
func (this *QCompleter) CompletionRole() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter14completionRoleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:119
// index:0
// bool wrapAround()
func (this *QCompleter) WrapAround() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10wrapAroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:121
// index:0
// int maxVisibleItems()
func (this *QCompleter) MaxVisibleItems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15maxVisibleItemsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:122
// index:0
// void setMaxVisibleItems(int)
func (this *QCompleter) SetMaxVisibleItems(maxItems int) {
	// 0: (, int maxItems), (&maxItems)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter18setMaxVisibleItemsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &maxItems)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:124
// index:0
// int completionCount()
func (this *QCompleter) CompletionCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15completionCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:125
// index:0
// bool setCurrentRow(int)
func (this *QCompleter) SetCurrentRow(row int) {
	// 0: (, int row), (&row)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter13setCurrentRowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &row)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:126
// index:0
// int currentRow()
func (this *QCompleter) CurrentRow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter10currentRowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:128
// index:0
// QModelIndex currentIndex()
func (this *QCompleter) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:129
// index:0
// QString currentCompletion()
func (this *QCompleter) CurrentCompletion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter17currentCompletionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:131
// index:0
// QAbstractItemModel * completionModel()
func (this *QCompleter) CompletionModel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter15completionModelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:133
// index:0
// QString completionPrefix()
func (this *QCompleter) CompletionPrefix() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter16completionPrefixEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:136
// index:0
// void setCompletionPrefix(const class QString &)
func (this *QCompleter) SetCompletionPrefix(prefix unsafe.Pointer) {
	// 0: (, const QString & prefix), (prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter19setCompletionPrefixERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:137
// index:0
// void complete(const class QRect &)
func (this *QCompleter) Complete(rect unsafe.Pointer) {
	// 0: (, const QRect & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter8completeERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:138
// index:0
// void setWrapAround(_Bool)
func (this *QCompleter) SetWrapAround(wrap bool) {
	// 0: (, bool wrap), (&wrap)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter13setWrapAroundEb", ffiqt.FFI_TYPE_VOID, this.cthis, &wrap)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:141
// index:0
// virtual
// QString pathFromIndex(const class QModelIndex &)
func (this *QCompleter) PathFromIndex(index unsafe.Pointer) {
	// 0: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter13pathFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:142
// index:0
// virtual
// QStringList splitPath(const class QString &)
func (this *QCompleter) SplitPath(path unsafe.Pointer) {
	// 0: (, const QString & path), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QCompleter9splitPathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, path)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:149
// index:0
// void activated(const class QString &)
func (this *QCompleter) Activated(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter9activatedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:150
// index:1
// void activated(const class QModelIndex &)
func (this *QCompleter) Activated_1(index unsafe.Pointer) {
	// 1: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter9activatedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:151
// index:0
// void highlighted(const class QString &)
func (this *QCompleter) Highlighted(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter11highlightedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:152
// index:1
// void highlighted(const class QModelIndex &)
func (this *QCompleter) Highlighted_1(index unsafe.Pointer) {
	// 1: (, const QModelIndex & index), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QCompleter11highlightedERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.cthis, index)
	gopp.ErrPrint(err, rv)
}

//  body block end
