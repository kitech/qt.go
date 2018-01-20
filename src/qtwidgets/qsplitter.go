//  header block begin
// /usr/include/qt/QtWidgets/qsplitter.h
// #include <qsplitter.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QSplitter struct {
	*QFrame
}

func (this *QSplitter) GetCthis() unsafe.Pointer {
	return this.QFrame.GetCthis()
}
func NewQSplitterFromPointer(cthis unsafe.Pointer) *QSplitter {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QSplitter{bcthis0}
}

// /usr/include/qt/QtWidgets/qsplitter.h:59
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSplitter) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:67
// index:0
// Public
// void QSplitter(class QWidget *)
func NewQSplitter(parent unsafe.Pointer) *QSplitter {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:68
// index:1
// Public
// void QSplitter(Qt::Orientation, class QWidget *)
func NewQSplitter_1(arg0 int, parent unsafe.Pointer) *QSplitter {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterC2EN2Qt11OrientationEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSplitterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qsplitter.h:69
// index:0
// Public virtual
// void ~QSplitter()
func DeleteQSplitter(*QSplitter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:71
// index:0
// Public
// void addWidget(class QWidget *)
func (this *QSplitter) AddWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter9addWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:72
// index:0
// Public
// void insertWidget(int, class QWidget *)
func (this *QSplitter) InsertWidget(index int, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:73
// index:0
// Public
// QWidget * replaceWidget(int, class QWidget *)
func (this *QSplitter) ReplaceWidget(index int, widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13replaceWidgetEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:75
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QSplitter) SetOrientation(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:76
// index:0
// Public
// Qt::Orientation orientation()
func (this *QSplitter) Orientation() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:78
// index:0
// Public
// void setChildrenCollapsible(_Bool)
func (this *QSplitter) SetChildrenCollapsible(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter22setChildrenCollapsibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:79
// index:0
// Public
// bool childrenCollapsible()
func (this *QSplitter) ChildrenCollapsible() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter19childrenCollapsibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:81
// index:0
// Public
// void setCollapsible(int, _Bool)
func (this *QSplitter) SetCollapsible(index int, arg1 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setCollapsibleEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:82
// index:0
// Public
// bool isCollapsible(int)
func (this *QSplitter) IsCollapsible(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter13isCollapsibleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:83
// index:0
// Public
// void setOpaqueResize(_Bool)
func (this *QSplitter) SetOpaqueResize(opaque bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter15setOpaqueResizeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &opaque)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:84
// index:0
// Public
// bool opaqueResize()
func (this *QSplitter) OpaqueResize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter12opaqueResizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:85
// index:0
// Public
// void refresh()
func (this *QSplitter) Refresh() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter7refreshEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:87
// index:0
// Public virtual
// QSize sizeHint()
func (this *QSplitter) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:88
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QSplitter) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:90
// index:0
// Public
// QList<int> sizes()
func (this *QSplitter) Sizes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter5sizesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:93
// index:0
// Public
// QByteArray saveState()
func (this *QSplitter) SaveState() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter9saveStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:94
// index:0
// Public
// bool restoreState(const class QByteArray &)
func (this *QSplitter) RestoreState(state *qtcore.QByteArray) interface{} {
	var convArg0 = state.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12restoreStateERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:96
// index:0
// Public
// int handleWidth()
func (this *QSplitter) HandleWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter11handleWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:97
// index:0
// Public
// void setHandleWidth(int)
func (this *QSplitter) SetHandleWidth(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter14setHandleWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:99
// index:0
// Public
// int indexOf(class QWidget *)
func (this *QSplitter) IndexOf(w unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter7indexOfEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:100
// index:0
// Public
// QWidget * widget(int)
func (this *QSplitter) Widget(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:101
// index:0
// Public
// int count()
func (this *QSplitter) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:103
// index:0
// Public
// void getRange(int, int *, int *)
func (this *QSplitter) GetRange(index int, arg1 unsafe.Pointer, arg2 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter8getRangeEiPiS0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, arg1, arg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:104
// index:0
// Public
// QSplitterHandle * handle(int)
func (this *QSplitter) Handle(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSplitter6handleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:106
// index:0
// Public
// void setStretchFactor(int, int)
func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter16setStretchFactorEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:109
// index:0
// Public
// void splitterMoved(int, int)
func (this *QSplitter) SplitterMoved(pos int, index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13splitterMovedEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:112
// index:0
// Protected virtual
// QSplitterHandle * createHandle()
func (this *QSplitter) CreateHandle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12createHandleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:114
// index:0
// Protected virtual
// void childEvent(class QChildEvent *)
func (this *QSplitter) ChildEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter10childEventEP11QChildEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:116
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSplitter) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qsplitter.h:117
// index:0
// Protected virtual
// void resizeEvent(class QResizeEvent *)
func (this *QSplitter) ResizeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:119
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QSplitter) ChangeEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:120
// index:0
// Protected
// void moveSplitter(int, int)
func (this *QSplitter) MoveSplitter(pos int, index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter12moveSplitterEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pos, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:121
// index:0
// Protected
// void setRubberBand(int)
func (this *QSplitter) SetRubberBand(position int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter13setRubberBandEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsplitter.h:122
// index:0
// Protected
// int closestLegalPosition(int, int)
func (this *QSplitter) ClosestLegalPosition(arg0 int, arg1 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSplitter20closestLegalPositionEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0, &arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
