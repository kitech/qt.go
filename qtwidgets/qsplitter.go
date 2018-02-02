package qtwidgets

// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
// QSplitterHandle * createHandle()
func (this *QSplitter) InheritCreateHandle(f func() unsafe.Pointer /*666*/) {
	ffiqt.SetAllInheritCallback(this, "createHandle", f)
}

// void childEvent(class QChildEvent *)
func (this *QSplitter) InheritChildEvent(f func(arg0 *qtcore.QChildEvent /*777 QChildEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "childEvent", f)
}

// bool event(class QEvent *)
func (this *QSplitter) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	ffiqt.SetAllInheritCallback(this, "event", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QSplitter) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QSplitter) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	ffiqt.SetAllInheritCallback(this, "changeEvent", f)
}

// void moveSplitter(int, int)
func (this *QSplitter) InheritMoveSplitter(f func(pos int, index int)) {
	ffiqt.SetAllInheritCallback(this, "moveSplitter", f)
}

// void setRubberBand(int)
func (this *QSplitter) InheritSetRubberBand(f func(position int)) {
	ffiqt.SetAllInheritCallback(this, "setRubberBand", f)
}

// int closestLegalPosition(int, int)
func (this *QSplitter) InheritClosestLegalPosition(f func(arg0 int, arg1 int) int) {
	ffiqt.SetAllInheritCallback(this, "closestLegalPosition", f)
}

type QSplitter struct {
	*QFrame
}

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
// [8] const QMetaObject * metaObject()
func (this *QSplitter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSplitter(QWidget *)
func NewQSplitter(parent *QWidget /*777 QWidget **/) *QSplitter {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:68
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSplitter(Qt::Orientation, QWidget *)
func NewQSplitter_1(arg0 int, parent *QWidget /*777 QWidget **/) *QSplitter {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_POINTER, arg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSplitter()
func DeleteQSplitter(this *QSplitter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qsplitter.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *)
func (this *QSplitter) AddWidget(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter9addWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertWidget(int, QWidget *)
func (this *QSplitter) InsertWidget(index int, widget *QWidget /*777 QWidget **/) {
	var convArg1 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * replaceWidget(int, QWidget *)
func (this *QSplitter) ReplaceWidget(index int, widget *QWidget /*777 QWidget **/) *QWidget /*777 QWidget **/ {
	var convArg1 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13replaceWidgetEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QSplitter) SetOrientation(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QSplitter) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setChildrenCollapsible(_Bool)
func (this *QSplitter) SetChildrenCollapsible(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter22setChildrenCollapsibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool childrenCollapsible()
func (this *QSplitter) ChildrenCollapsible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter19childrenCollapsibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCollapsible(int, _Bool)
func (this *QSplitter) SetCollapsible(index int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setCollapsibleEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:82
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCollapsible(int)
func (this *QSplitter) IsCollapsible(index int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter13isCollapsibleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaqueResize(_Bool)
func (this *QSplitter) SetOpaqueResize(opaque bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter15setOpaqueResizeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), opaque)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool opaqueResize()
func (this *QSplitter) OpaqueResize() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter12opaqueResizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh()
func (this *QSplitter) Refresh() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter7refreshEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QSplitter) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QSplitter) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveState()
func (this *QSplitter) SaveState() *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter9saveStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreState(const QByteArray &)
func (this *QSplitter) RestoreState(state *qtcore.QByteArray) bool {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] int handleWidth()
func (this *QSplitter) HandleWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter11handleWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHandleWidth(int)
func (this *QSplitter) SetHandleWidth(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setHandleWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int indexOf(QWidget *)
func (this *QSplitter) IndexOf(w *QWidget /*777 QWidget **/) int {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int)
func (this *QSplitter) Widget(index int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QSplitter) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsplitter.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getRange(int, int *, int *)
func (this *QSplitter) GetRange(index int, arg1 unsafe.Pointer /*666*/, arg2 unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter8getRangeEiPiS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, &arg1, &arg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QSplitterHandle * handle(int)
func (this *QSplitter) Handle(index int) *QSplitterHandle /*777 QSplitterHandle **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter6handleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSplitterHandleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretchFactor(int, int)
func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter16setStretchFactorEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void splitterMoved(int, int)
func (this *QSplitter) SplitterMoved(pos int, index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13splitterMovedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSplitterHandle * createHandle()
func (this *QSplitter) CreateHandle() *QSplitterHandle /*777 QSplitterHandle **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12createHandleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSplitterHandleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qsplitter.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)
func (this *QSplitter) ChildEvent(arg0 *qtcore.QChildEvent /*777 QChildEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSplitter) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsplitter.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QSplitter) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QSplitter) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:120
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void moveSplitter(int, int)
func (this *QSplitter) MoveSplitter(pos int, index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12moveSplitterEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pos, index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:121
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setRubberBand(int)
func (this *QSplitter) SetRubberBand(position int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13setRubberBandEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:122
// index:0
// Protected Visibility=Default Availability=Available
// [4] int closestLegalPosition(int, int)
func (this *QSplitter) ClosestLegalPosition(arg0 int, arg1 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter20closestLegalPositionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
