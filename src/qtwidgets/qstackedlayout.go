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
func NewQStackedLayoutFromPointer(cthis unsafe.Pointer) *QStackedLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QStackedLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QStackedLayout) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:66
// index:0
// Public
// void QStackedLayout()
func NewQStackedLayout() *QStackedLayout {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQStackedLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:67
// index:1
// Public
// void QStackedLayout(class QWidget *)
func NewQStackedLayout_1(parent *QWidget /*444 QWidget **/) *QStackedLayout {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStackedLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:68
// index:2
// Public
// void QStackedLayout(class QLayout *)
func NewQStackedLayout_2(parentLayout *QLayout /*444 QLayout **/) *QStackedLayout {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = parentLayout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutC2EP7QLayout", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStackedLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:69
// index:0
// Public virtual
// void ~QStackedLayout()
func DeleteQStackedLayout(*QStackedLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:71
// index:0
// Public
// int addWidget(class QWidget *)
func (this *QStackedLayout) AddWidget(w *QWidget /*444 QWidget **/) int {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout9addWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:72
// index:0
// Public
// int insertWidget(int, class QWidget *)
func (this *QStackedLayout) InsertWidget(index int, w *QWidget /*444 QWidget **/) int {
	var convArg1 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout12insertWidgetEiP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:74
// index:0
// Public
// QWidget * currentWidget()
func (this *QStackedLayout) CurrentWidget() *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout13currentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:75
// index:0
// Public
// int currentIndex()
func (this *QStackedLayout) CurrentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:77
// index:0
// Public
// QWidget * widget(int)
func (this *QStackedLayout) Widget(arg0 int) *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout6widgetEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:78
// index:0
// Public virtual
// int count()
func (this *QStackedLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:80
// index:0
// Public
// QStackedLayout::StackingMode stackingMode()
func (this *QStackedLayout) StackingMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout12stackingModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:81
// index:0
// Public
// void setStackingMode(enum QStackedLayout::StackingMode)
func (this *QStackedLayout) SetStackingMode(stackingMode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout15setStackingModeENS_12StackingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stackingMode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:84
// index:0
// Public virtual
// void addItem(class QLayoutItem *)
func (this *QStackedLayout) AddItem(item *QLayoutItem /*444 QLayoutItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:85
// index:0
// Public virtual
// QSize sizeHint()
func (this *QStackedLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:86
// index:0
// Public virtual
// QSize minimumSize()
func (this *QStackedLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:87
// index:0
// Public virtual
// QLayoutItem * itemAt(int)
func (this *QStackedLayout) ItemAt(arg0 int) *QLayoutItem /*444 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:88
// index:0
// Public virtual
// QLayoutItem * takeAt(int)
func (this *QStackedLayout) TakeAt(arg0 int) *QLayoutItem /*444 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:89
// index:0
// Public virtual
// void setGeometry(const class QRect &)
func (this *QStackedLayout) SetGeometry(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:90
// index:0
// Public virtual
// bool hasHeightForWidth()
func (this *QStackedLayout) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:91
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QStackedLayout) HeightForWidth(width int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QStackedLayout14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &width)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:94
// index:0
// Public
// void widgetRemoved(int)
func (this *QStackedLayout) WidgetRemoved(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout13widgetRemovedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:95
// index:0
// Public
// void currentChanged(int)
func (this *QStackedLayout) CurrentChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout14currentChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:98
// index:0
// Public
// void setCurrentIndex(int)
func (this *QStackedLayout) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qstackedlayout.h:99
// index:0
// Public
// void setCurrentWidget(class QWidget *)
func (this *QStackedLayout) SetCurrentWidget(w *QWidget /*444 QWidget **/) {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStackedLayout16setCurrentWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
