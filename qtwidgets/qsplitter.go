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

// void childEvent(class QChildEvent *)
func (this *QSplitter) InheritChildEvent(f func(arg0 *qtcore.QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

// bool event(class QEvent *)
func (this *QSplitter) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QSplitter) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void changeEvent(class QEvent *)
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
func (this *QSplitter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplitter(QWidget *)
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
func NewQSplitter__() *QSplitter {
	// arg: 0, QWidget *=Pointer, QWidget=Record,
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
func NewQSplitter_1(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QSplitter {
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
func NewQSplitter_1_(arg0 int) *QSplitter {
	// arg: 1, QWidget *=Pointer, QWidget=Record,
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
func (this *QSplitter) SetOrientation(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const
func (this *QSplitter) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChildrenCollapsible(_Bool)
func (this *QSplitter) SetChildrenCollapsible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter22setChildrenCollapsibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool childrenCollapsible() const
func (this *QSplitter) ChildrenCollapsible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter19childrenCollapsibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCollapsible(int, _Bool)
func (this *QSplitter) SetCollapsible(index int, arg1 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter14setCollapsibleEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCollapsible(int) const
func (this *QSplitter) IsCollapsible(index int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter13isCollapsibleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaqueResize(_Bool)
func (this *QSplitter) SetOpaqueResize(opaque bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter15setOpaqueResizeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opaque)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaqueResize(_Bool)
func (this *QSplitter) SetOpaqueResize__() {
	// arg: 0, bool=Bool, =Invalid,
	opaque := true
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter15setOpaqueResizeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opaque)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool opaqueResize() const
func (this *QSplitter) OpaqueResize() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter12opaqueResizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()
func (this *QSplitter) Refresh() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter7refreshEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const
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
func (this *QSplitter) HandleWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter11handleWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHandleWidth(int)
func (this *QSplitter) SetHandleWidth(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter14setHandleWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const
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
func (this *QSplitter) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter6widgetEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const
func (this *QSplitter) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getRange(int, int *, int *) const
func (this *QSplitter) GetRange(index int, arg1 unsafe.Pointer /*666*/, arg2 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter8getRangeEiPiS0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1, arg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QSplitterHandle * handle(int) const
func (this *QSplitter) Handle(index int) *QSplitterHandle /*777 QSplitterHandle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSplitter6handleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSplitterHandleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretchFactor(int, int)
func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter16setStretchFactorEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void splitterMoved(int, int)
func (this *QSplitter) SplitterMoved(pos int, index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter13splitterMovedEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSplitterHandle * createHandle()
func (this *QSplitter) CreateHandle() *QSplitterHandle /*777 QSplitterHandle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter12createHandleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSplitterHandleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qsplitter.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)
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
func (this *QSplitter) MoveSplitter(pos int, index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter12moveSplitterEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:121
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRubberBand(int)
func (this *QSplitter) SetRubberBand(position int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSplitter13setRubberBandEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:122
// index:0
// Protected Visibility=Default Availability=Available
// [4] int closestLegalPosition(int, int)
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
