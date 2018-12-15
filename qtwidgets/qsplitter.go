// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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

// QSplitterHandle * createHandle()
func (this *QSplitter) InheritCreateHandle(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "createHandle", f)
}

// void childEvent(QChildEvent *)
func (this *QSplitter) InheritChildEvent(f func(arg0 *qtcore.QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

// bool event(QEvent *)
func (this *QSplitter) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(QResizeEvent *)
func (this *QSplitter) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void changeEvent(QEvent *)
func (this *QSplitter) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// void moveSplitter(int, int)
func (this *QSplitter) InheritMoveSplitter(f func(pos int, index int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveSplitter", f)
}

// void setRubberBand(int)
func (this *QSplitter) InheritSetRubberBand(f func(position int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setRubberBand", f)
}

// int closestLegalPosition(int, int)
func (this *QSplitter) InheritClosestLegalPosition(f func(arg0 int, arg1 int) int) {
	qtrt.SetAllInheritCallback(this, "closestLegalPosition", f)
}

/*

 */
type QSplitter struct {
	*QFrame
}
type QSplitter_ITF interface {
	QFrame_ITF
	QSplitter_PTR() *QSplitter
}

func (ptr *QSplitter) QSplitter_PTR() *QSplitter { return ptr }

func (this *QSplitter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QSplitter) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQSplitterFromPointer(cthis unsafe.Pointer) *QSplitter {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QSplitter{bcthis0}
}
func (*QSplitter) NewFromPointer(cthis unsafe.Pointer) *QSplitter {
	return NewQSplitterFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsplitter.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSplitter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplitter(QWidget *)

/*
Constructs a horizontal splitter with the parent argument passed on to the QFrame constructor.

See also setOrientation().
*/
func (*QSplitter) NewForInherit(parent QWidget_ITF /*777 QWidget **/) *QSplitter {
	return NewQSplitter(parent)
}
func NewQSplitter(parent QWidget_ITF /*777 QWidget **/) *QSplitter {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitterC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplitter")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplitter(QWidget *)

/*
Constructs a horizontal splitter with the parent argument passed on to the QFrame constructor.

See also setOrientation().
*/
func (*QSplitter) NewForInheritp() *QSplitter {
	return NewQSplitterp()
}
func NewQSplitterp() *QSplitter {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitterC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplitter")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplitter(Qt::Orientation, QWidget *)

/*
Constructs a horizontal splitter with the parent argument passed on to the QFrame constructor.

See also setOrientation().
*/
func (*QSplitter) NewForInherit1(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QSplitter {
	return NewQSplitter1(arg0, parent)
}
func NewQSplitter1(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QSplitter {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitterC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplitter")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplitter(Qt::Orientation, QWidget *)

/*
Constructs a horizontal splitter with the parent argument passed on to the QFrame constructor.

See also setOrientation().
*/
func (*QSplitter) NewForInherit1p(arg0 int) *QSplitter {
	return NewQSplitter1p(arg0)
}
func NewQSplitter1p(arg0 int) *QSplitter {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitterC2EN2Qt11OrientationEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSplitter")
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSplitter()

/*

 */
func DeleteQSplitter(this *QSplitter) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitterD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsplitter.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *)

/*
Adds the given widget to the splitter's layout after all the other items.

If widget is already in the splitter, it will be moved to the new position.

Note: The splitter takes ownership of the widget.

See also insertWidget(), widget(), and indexOf().
*/
func (this *QSplitter) AddWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter9addWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertWidget(int, QWidget *)

/*
Inserts the widget specified into the splitter's layout at the given index.

If widget is already in the splitter, it will be moved to the new position.

If index is an invalid index, then the widget will be inserted at the end.

Note: The splitter takes ownership of the widget.

See also addWidget(), indexOf(), and widget().
*/
func (this *QSplitter) InsertWidget(index int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter12insertWidgetEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * replaceWidget(int, QWidget *)

/*
Replaces the widget in the splitter's layout at the given index by widget.

Returns the widget that has just been replaced if index is valid and widget is not already a child of the splitter. Otherwise, it returns null and no replacement or addition is made.

The geometry of the newly inserted widget will be the same as the widget it replaces. Its visible and collapsed states are also inherited.

Note: The splitter takes ownership of widget and sets the parent of the replaced widget to null.

Note: Because widget gets reparented into the splitter, its geometry may not be set right away, but only after widget will receive the appropriate events.

This function was introduced in  Qt 5.9.

See also insertWidget() and indexOf().
*/
func (this *QSplitter) ReplaceWidget(index int, widget QWidget_ITF /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter13replaceWidgetEiP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*

 */
func (this *QSplitter) SetOrientation(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*

 */
func (this *QSplitter) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChildrenCollapsible(bool)

/*

 */
func (this *QSplitter) SetChildrenCollapsible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter22setChildrenCollapsibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool childrenCollapsible() const

/*

 */
func (this *QSplitter) ChildrenCollapsible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter19childrenCollapsibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCollapsible(int, bool)

/*
Sets whether the child widget at index is collapsible to collapse.

By default, children are collapsible, meaning that the user can resize them down to size 0, even if they have a non-zero minimumSize() or minimumSizeHint(). This behavior can be changed on a per-widget basis by calling this function, or globally for all the widgets in the splitter by setting the childrenCollapsible property.

See also isCollapsible() and childrenCollapsible.
*/
func (this *QSplitter) SetCollapsible(index int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter14setCollapsibleEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCollapsible(int) const

/*
Returns true if the widget at index is collapsible, otherwise returns false.
*/
func (this *QSplitter) IsCollapsible(index int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter13isCollapsibleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaqueResize(bool)

/*

 */
func (this *QSplitter) SetOpaqueResize(opaque bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter15setOpaqueResizeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opaque)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaqueResize(bool)

/*

 */
func (this *QSplitter) SetOpaqueResizep() {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	opaque := true
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter15setOpaqueResizeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opaque)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool opaqueResize() const

/*

 */
func (this *QSplitter) OpaqueResize() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter12opaqueResizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()

/*
Updates the splitter's state. You should not need to call this function.
*/
func (this *QSplitter) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QFrame::sizeHint().
*/
func (this *QSplitter) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint() const

/*
Reimplemented from QWidget::minimumSizeHint().
*/
func (this *QSplitter) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState() const

/*
Saves the state of the splitter's layout.

Typically this is used in conjunction with QSettings to remember the size for a future session. A version number is stored as part of the data. Here is an example:


      QSettings settings;
      settings.setValue("splitterSizes", splitter->saveState());



See also restoreState().
*/
func (this *QSplitter) SaveState() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter9saveStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &)

/*
Restores the splitter's layout to the state specified. Returns true if the state is restored; otherwise returns false.

Typically this is used in conjunction with QSettings to restore the size from a past session. Here is an example:

Restore the splitter's state:


      QSettings settings;
      splitter->restoreState(settings.value("splitterSizes").toByteArray());



A failure to restore the splitter's layout may result from either invalid or out-of-date data in the supplied byte array.

See also saveState().
*/
func (this *QSplitter) RestoreState(state qtcore.QByteArray_ITF) bool {
	var convArg0 unsafe.Pointer
	if state != nil && state.QByteArray_PTR() != nil {
		convArg0 = state.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter12restoreStateERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int handleWidth() const

/*

 */
func (this *QSplitter) HandleWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter11handleWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHandleWidth(int)

/*

 */
func (this *QSplitter) SetHandleWidth(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter14setHandleWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const

/*
Returns the index in the splitter's layout of the specified widget, or -1 if widget is not found. This also works for handles.

Handles are numbered from 0. There are as many handles as there are child widgets, but the handle at position 0 is always hidden.

See also count() and widget().
*/
func (this *QSplitter) IndexOf(w QWidget_ITF /*777 QWidget **/) int {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter7indexOfEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int) const

/*
Returns the widget at the given index in the splitter's layout, or nullptr if there is no such widget.

See also count(), handle(), indexOf(), and insertWidget().
*/
func (this *QSplitter) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter6widgetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*
Returns the number of widgets contained in the splitter's layout.

See also widget() and handle().
*/
func (this *QSplitter) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getRange(int, int *, int *) const

/*
Returns the valid range of the splitter at index in *min and *max if min and max are not 0.
*/
func (this *QSplitter) GetRange(index int, arg1 unsafe.Pointer /*666*/, arg2 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter8getRangeEiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1, arg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QSplitterHandle * handle(int) const

/*
Returns the handle to the left of (or above) the item in the splitter's layout at the given index, or nullptr if there is no such item. The handle at index 0 is always hidden.

For right-to-left languages such as Arabic and Hebrew, the layout of horizontal splitters is reversed. The handle will be to the right of the widget at index.

See also count(), widget(), indexOf(), createHandle(), and setHandleWidth().
*/
func (this *QSplitter) Handle(index int) *QSplitterHandle /*777 QSplitterHandle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter6handleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSplitterHandleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretchFactor(int, int)

/*
Updates the size policy of the widget at position index to have a stretch factor of stretch.

stretch is not the effective stretch factor; the effective stretch factor is calculated by taking the initial size of the widget and multiplying it with stretch.

This function is provided for convenience. It is equivalent to


  QWidget *widget = splitter->widget(index);
  QSizePolicy policy = widget->sizePolicy();
  policy.setHorizontalStretch(stretch);
  policy.setVerticalStretch(stretch);
  widget->setSizePolicy(policy);



See also setSizes() and widget().
*/
func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter16setStretchFactorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void splitterMoved(int, int)

/*
This signal is emitted when the splitter handle at a particular index has been moved to position pos.

For right-to-left languages such as Arabic and Hebrew, the layout of horizontal splitters is reversed. pos is then the distance from the right edge of the widget.

See also moveSplitter().
*/
func (this *QSplitter) SplitterMoved(pos int, index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter13splitterMovedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSplitterHandle * createHandle()

/*
Returns a new splitter handle as a child widget of this splitter. This function can be reimplemented in subclasses to provide support for custom handles.

See also handle() and indexOf().
*/
func (this *QSplitter) CreateHandle() *QSplitterHandle /*777 QSplitterHandle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter12createHandleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSplitterHandleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)

/*
Reimplemented from QObject::childEvent().

Tells the splitter that the child widget described by c has been inserted or removed.

This method is also used to handle the situation where a widget is created with the splitter as a parent but not explicitly added with insertWidget() or addWidget(). This is for compatibility and not the recommended way of putting widgets into a splitter in new code. Please use insertWidget() or addWidget() in new code.

See also addWidget() and insertWidget().
*/
func (this *QSplitter) ChildEvent(arg0 qtcore.QChildEvent_ITF /*777 QChildEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChildEvent_PTR() != nil {
		convArg0 = arg0.QChildEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QFrame::event().
*/
func (this *QSplitter) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QSplitter) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)

/*
Reimplemented from QFrame::changeEvent().
*/
func (this *QSplitter) ChangeEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:120
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void moveSplitter(int, int)

/*
Moves the left or top edge of the splitter handle at index as close as possible to position pos, which is the distance from the left or top edge of the widget.

For right-to-left languages such as Arabic and Hebrew, the layout of horizontal splitters is reversed. pos is then the distance from the right edge of the widget.

See also splitterMoved(), closestLegalPosition(), and getRange().
*/
func (this *QSplitter) MoveSplitter(pos int, index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter12moveSplitterEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:121
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRubberBand(int)

/*
Displays a rubber band at position pos. If pos is negative, the rubber band is removed.
*/
func (this *QSplitter) SetRubberBand(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter13setRubberBandEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:122
// index:0
// Protected Visibility=Default Availability=Available
// [4] int closestLegalPosition(int, int)

/*
Returns the closest legal position to pos of the widget at index.

For right-to-left languages such as Arabic and Hebrew, the layout of horizontal splitters is reversed. Positions are then measured from the right edge of the widget.

See also getRange().
*/
func (this *QSplitter) ClosestLegalPosition(arg0 int, arg1 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter20closestLegalPositionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
