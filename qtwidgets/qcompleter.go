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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool eventFilter(QObject *, QEvent *)
func (this *QCompleter) InheritEventFilter(f func(o *qtcore.QObject /*777 QObject **/, e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// bool event(QEvent *)
func (this *QCompleter) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QCompleter struct {
	*qtcore.QObject
}
type QCompleter_ITF interface {
	qtcore.QObject_ITF
	QCompleter_PTR() *QCompleter
}

func (ptr *QCompleter) QCompleter_PTR() *QCompleter { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCompleter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(QObject *)

/*
Constructs a completer object with the given parent.
*/
func NewQCompleter(parent qtcore.QObject_ITF /*777 QObject **/) *QCompleter {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(QObject *)

/*
Constructs a completer object with the given parent.
*/
func NewQCompleter__() *QCompleter {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(QAbstractItemModel *, QObject *)

/*
Constructs a completer object with the given parent.
*/
func NewQCompleter_1(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/, parent qtcore.QObject_ITF /*777 QObject **/) *QCompleter {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2EP18QAbstractItemModelP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(QAbstractItemModel *, QObject *)

/*
Constructs a completer object with the given parent.
*/
func NewQCompleter_1_(model qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) *QCompleter {
	var convArg0 unsafe.Pointer
	if model != nil && model.QAbstractItemModel_PTR() != nil {
		convArg0 = model.QAbstractItemModel_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2EP18QAbstractItemModelP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:88
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(const QStringList &, QObject *)

/*
Constructs a completer object with the given parent.
*/
func NewQCompleter_2(completions qtcore.QStringList_ITF, parent qtcore.QObject_ITF /*777 QObject **/) *QCompleter {
	var convArg0 unsafe.Pointer
	if completions != nil && completions.QStringList_PTR() != nil {
		convArg0 = completions.QStringList_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:88
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCompleter(const QStringList &, QObject *)

/*
Constructs a completer object with the given parent.
*/
func NewQCompleter_2_(completions qtcore.QStringList_ITF) *QCompleter {
	var convArg0 unsafe.Pointer
	if completions != nil && completions.QStringList_PTR() != nil {
		convArg0 = completions.QStringList_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleterC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCompleterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCompleter")
	return gothis
}

// /usr/include/qt/QtWidgets/qcompleter.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCompleter()

/*

 */
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

/*
Sets the widget for which completion are provided for to widget. This function is automatically called when a QCompleter is set on a QLineEdit using QLineEdit::setCompleter() or on a QComboBox using QComboBox::setCompleter(). The widget needs to be set explicitly when providing completions for custom widgets.

See also widget(), setModel(), and setPopup().
*/
func (this *QCompleter) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns the widget for which the completer object is providing completions.

See also setWidget().
*/
func (this *QCompleter) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)

/*
Sets the model which provides completions to model. The model can be list model or a tree model. If a model has been already previously set and it has the QCompleter as its parent, it is deleted.

For convenience, if model is a QFileSystemModel, QCompleter switches its caseSensitivity to Qt::CaseInsensitive on Windows and Qt::CaseSensitive on other platforms.

See also completionModel(), modelSorting, and Handling Tree Models.
*/
func (this *QCompleter) SetModel(c qtcore.QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QAbstractItemModel_PTR() != nil {
		convArg0 = c.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model() const

/*
Returns the model that provides completion strings.

See also setModel() and completionModel().
*/
func (this *QCompleter) Model() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompletionMode(QCompleter::CompletionMode)

/*

 */
func (this *QCompleter) SetCompletionMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter17setCompletionModeENS_14CompletionModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QCompleter::CompletionMode completionMode() const

/*

 */
func (this *QCompleter) CompletionMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter14completionModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilterMode(Qt::MatchFlags)

/*

 */
func (this *QCompleter) SetFilterMode(filterMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter13setFilterModeE6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filterMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MatchFlags filterMode() const

/*

 */
func (this *QCompleter) FilterMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10filterModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemView * popup() const

/*
Returns the popup used to display completions.

See also setPopup().
*/
func (this *QCompleter) Popup() *QAbstractItemView /*777 QAbstractItemView **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter5popupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemViewFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPopup(QAbstractItemView *)

/*
Sets the popup used to display completions to popup. QCompleter takes ownership of the view.

A QListView is automatically created when the completionMode() is set to QCompleter::PopupCompletion or QCompleter::UnfilteredPopupCompletion. The default popup displays the completionColumn().

Ensure that this function is called before the view settings are modified. This is required since view's properties may require that a model has been set on the view (for example, hiding columns in the view requires a model to be set on the view).

See also popup().
*/
func (this *QCompleter) SetPopup(popup QAbstractItemView_ITF /*777 QAbstractItemView **/) {
	var convArg0 unsafe.Pointer
	if popup != nil && popup.QAbstractItemView_PTR() != nil {
		convArg0 = popup.QAbstractItemView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter8setPopupEP17QAbstractItemView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCaseSensitivity(Qt::CaseSensitivity)

/*

 */
func (this *QCompleter) SetCaseSensitivity(caseSensitivity int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter18setCaseSensitivityEN2Qt15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), caseSensitivity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:108
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::CaseSensitivity caseSensitivity() const

/*

 */
func (this *QCompleter) CaseSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15caseSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModelSorting(QCompleter::ModelSorting)

/*

 */
func (this *QCompleter) SetModelSorting(sorting int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter15setModelSortingENS_12ModelSortingE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sorting)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:111
// index:0
// Public Visibility=Default Availability=Available
// [4] QCompleter::ModelSorting modelSorting() const

/*

 */
func (this *QCompleter) ModelSorting() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter12modelSortingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompletionColumn(int)

/*

 */
func (this *QCompleter) SetCompletionColumn(column int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter19setCompletionColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int completionColumn() const

/*

 */
func (this *QCompleter) CompletionColumn() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter16completionColumnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCompletionRole(int)

/*

 */
func (this *QCompleter) SetCompletionRole(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter17setCompletionRoleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int completionRole() const

/*

 */
func (this *QCompleter) CompletionRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter14completionRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:119
// index:0
// Public Visibility=Default Availability=Available
// [1] bool wrapAround() const

/*

 */
func (this *QCompleter) WrapAround() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10wrapAroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:121
// index:0
// Public Visibility=Default Availability=Available
// [4] int maxVisibleItems() const

/*

 */
func (this *QCompleter) MaxVisibleItems() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15maxVisibleItemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaxVisibleItems(int)

/*

 */
func (this *QCompleter) SetMaxVisibleItems(maxItems int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter18setMaxVisibleItemsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxItems)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] int completionCount() const

/*
Returns the number of completions for the current prefix. For an unsorted model with a large number of items this can be expensive. Use setCurrentRow() and currentCompletion() to iterate through all the completions.
*/
func (this *QCompleter) CompletionCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15completionCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setCurrentRow(int)

/*
Sets the current row to the row specified. Returns true if successful; otherwise returns false.

This function may be used along with currentCompletion() to iterate through all the possible completions.

See also currentRow(), currentCompletion(), and completionCount().
*/
func (this *QCompleter) SetCurrentRow(row int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter13setCurrentRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:126
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentRow() const

/*
Returns the current row.

See also setCurrentRow().
*/
func (this *QCompleter) CurrentRow() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter10currentRowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcompleter.h:128
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex currentIndex() const

/*
Returns the model index of the current completion in the completionModel().

See also setCurrentRow(), currentCompletion(), and model().
*/
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
// [8] QString currentCompletion() const

/*
Returns the current completion string. This includes the completionPrefix. When used alongside setCurrentRow(), it can be used to iterate through all the matches.

See also setCurrentRow() and currentIndex().
*/
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
// [8] QAbstractItemModel * completionModel() const

/*
Returns the completion model. The completion model is a read-only list model that contains all the possible matches for the current completion prefix. The completion model is auto-updated to reflect the current completions.

Note: The return value of this function is defined to be an QAbstractItemModel purely for generality. This actual kind of model returned is an instance of an QAbstractProxyModel subclass.

See also completionPrefix and model().
*/
func (this *QCompleter) CompletionModel() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QCompleter15completionModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcompleter.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QString completionPrefix() const

/*

 */
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

/*

 */
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

/*
For QCompleter::PopupCompletion and QCompletion::UnfilteredPopupCompletion modes, calling this function displays the popup displaying the current completions. By default, if rect is not specified, the popup is displayed on the bottom of the widget(). If rect is specified the popup is displayed on the left edge of the rectangle.

For QCompleter::InlineCompletion mode, the highlighted() signal is fired with the current completion.
*/
func (this *QCompleter) Complete(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter8completeERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void complete(const QRect &)

/*
For QCompleter::PopupCompletion and QCompletion::UnfilteredPopupCompletion modes, calling this function displays the popup displaying the current completions. By default, if rect is not specified, the popup is displayed on the bottom of the widget(). If rect is specified the popup is displayed on the left edge of the rectangle.

For QCompleter::InlineCompletion mode, the highlighted() signal is fired with the current completion.
*/
func (this *QCompleter) Complete__() {
	// arg: 0, const QRect &=LValueReference, QRect=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter8completeERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWrapAround(bool)

/*

 */
func (this *QCompleter) SetWrapAround(wrap bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter13setWrapAroundEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), wrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:141
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString pathFromIndex(const QModelIndex &) const

/*
Returns the path for the given index. The completer object uses this to obtain the completion text from the underlying model.

The default implementation returns the edit role of the item for list models. It returns the absolute file path if the model is a QFileSystemModel.

See also splitPath().
*/
func (this *QCompleter) PathFromIndex(index qtcore.QModelIndex_ITF) string {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
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
// [8] QStringList splitPath(const QString &) const

/*
Splits the given path into strings that are used to match at each level in the model().

The default implementation of splitPath() splits a file system path based on QDir::separator() when the sourceModel() is a QFileSystemModel.

When used with list models, the first item in the returned list is used for matching.

See also pathFromIndex() and Handling Tree Models.
*/
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

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QCompleter) EventFilter(o qtcore.QObject_ITF /*777 QObject **/, e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if o != nil && o.QObject_PTR() != nil {
		convArg0 = o.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg1 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:146
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QCompleter) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcompleter.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activated(const QString &)

/*
This signal is sent when an item in the popup() is activated by the user (by clicking or pressing return). The item's text is given.

Note: Signal activated is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(completer, QOverload<const QString &>::of(&QCompleter::activated),
      [=](const QString &text){ \/* ... *\/ });
*/
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

/*
This signal is sent when an item in the popup() is activated by the user (by clicking or pressing return). The item's text is given.

Note: Signal activated is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(completer, QOverload<const QString &>::of(&QCompleter::activated),
      [=](const QString &text){ \/* ... *\/ });
*/
func (this *QCompleter) Activated_1(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter9activatedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcompleter.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void highlighted(const QString &)

/*
This signal is sent when an item in the popup() is highlighted by the user. It is also sent if complete() is called with the completionMode() set to QCompleter::InlineCompletion. The item's text is given.

Note: Signal highlighted is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(completer, QOverload<const QString &>::of(&QCompleter::highlighted),
      [=](const QString &text){ \/* ... *\/ });
*/
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

/*
This signal is sent when an item in the popup() is highlighted by the user. It is also sent if complete() is called with the completionMode() set to QCompleter::InlineCompletion. The item's text is given.

Note: Signal highlighted is overloaded in this class. To connect to this signal by using the function pointer syntax, Qt provides a convenient helper for obtaining the function pointer as shown in this example:


  connect(completer, QOverload<const QString &>::of(&QCompleter::highlighted),
      [=](const QString &text){ \/* ... *\/ });
*/
func (this *QCompleter) Highlighted_1(index qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QCompleter11highlightedERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This enum specifies how completions are provided to the user.



See also setCompletionMode().

*/
type QCompleter__CompletionMode = int

// Current completions are displayed in a popup window.
const QCompleter__PopupCompletion QCompleter__CompletionMode = 0

// All possible completions are displayed in a popup window with the most likely suggestion indicated as current.
const QCompleter__UnfilteredPopupCompletion QCompleter__CompletionMode = 1

// Completions appear inline (as selected text).
const QCompleter__InlineCompletion QCompleter__CompletionMode = 2

/*
This enum specifies how the items in the model are sorted.



See also setModelSorting().

*/
type QCompleter__ModelSorting = int

// The model is unsorted.
const QCompleter__UnsortedModel QCompleter__ModelSorting = 0

// The model is sorted case sensitively.
const QCompleter__CaseSensitivelySortedModel QCompleter__ModelSorting = 1

// The model is sorted case insensitively.
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
