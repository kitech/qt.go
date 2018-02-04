package qtwidgets

// /usr/include/qt/QtWidgets/qstackedlayout.h
// #include <qstackedlayout.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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

type QStackedLayout struct {
	*QLayout
}

func (this *QStackedLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayout.GetCthis()
	}
}
func (this *QStackedLayout) SetCthis(cthis unsafe.Pointer) {
	this.QLayout = NewQLayoutFromPointer(cthis)
}
func NewQStackedLayoutFromPointer(cthis unsafe.Pointer) *QStackedLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QStackedLayout{bcthis0}
}
func (*QStackedLayout) NewFromPointer(cthis unsafe.Pointer) *QStackedLayout {
	return NewQStackedLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStackedLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStackedLayout()
func NewQStackedLayout() *QStackedLayout {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQStackedLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:67
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStackedLayout(QWidget *)
func NewQStackedLayout_1(parent *QWidget /*777 QWidget **/) *QStackedLayout {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2EP7QWidget", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStackedLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QStackedLayout(QLayout *)
func NewQStackedLayout_2(parentLayout *QLayout /*777 QLayout **/) *QStackedLayout {
	var convArg0 = parentLayout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2EP7QLayout", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStackedLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStackedLayout()
func DeleteQStackedLayout(this *QStackedLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:71
// index:0
// Public Visibility=Default Availability=Available
// [4] int addWidget(QWidget *)
func (this *QStackedLayout) AddWidget(w *QWidget /*777 QWidget **/) int {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout9addWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [4] int insertWidget(int, QWidget *)
func (this *QStackedLayout) InsertWidget(index int, w *QWidget /*777 QWidget **/) int {
	var convArg1 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * currentWidget()
func (this *QStackedLayout) CurrentWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout13currentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex()
func (this *QStackedLayout) CurrentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget(int)
func (this *QStackedLayout) Widget(arg0 int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QStackedLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] QStackedLayout::StackingMode stackingMode()
func (this *QStackedLayout) StackingMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout12stackingModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStackingMode(enum QStackedLayout::StackingMode)
func (this *QStackedLayout) SetStackingMode(stackingMode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout15setStackingModeENS_12StackingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stackingMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)
func (this *QStackedLayout) AddItem(item *QLayoutItem /*777 QLayoutItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QStackedLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QStackedLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int)
func (this *QStackedLayout) ItemAt(arg0 int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)
func (this *QStackedLayout) TakeAt(arg0 int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QStackedLayout) SetGeometry(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QStackedLayout) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QStackedLayout) HeightForWidth(width int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), width)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void widgetRemoved(int)
func (this *QStackedLayout) WidgetRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout13widgetRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentChanged(int)
func (this *QStackedLayout) CurrentChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)
func (this *QStackedLayout) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentWidget(QWidget *)
func (this *QStackedLayout) SetCurrentWidget(w *QWidget /*777 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QStackedLayout__StackingMode = int

const QStackedLayout__StackOne QStackedLayout__StackingMode = 0
const QStackedLayout__StackAll QStackedLayout__StackingMode = 1

//  body block end
