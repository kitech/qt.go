//  header block begin
// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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
type QBoxLayout struct {
	*QLayout
}

func (this *QBoxLayout) GetCthis() unsafe.Pointer {
	return this.QLayout.GetCthis()
}
func NewQBoxLayoutFromPointer(cthis unsafe.Pointer) *QBoxLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QBoxLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qboxlayout.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QBoxLayout) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// Public
// void QBoxLayout(enum QBoxLayout::Direction, class QWidget *)
func NewQBoxLayout(arg0 int, parent unsafe.Pointer) *QBoxLayout {
	cthis := qtrt.Calloc(1, 256) // 32
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQBoxLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:66
// index:0
// Public virtual
// void ~QBoxLayout()
func DeleteQBoxLayout(*QBoxLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:68
// index:0
// Public
// QBoxLayout::Direction direction()
func (this *QBoxLayout) Direction() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:69
// index:0
// Public
// void setDirection(enum QBoxLayout::Direction)
func (this *QBoxLayout) SetDirection(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout12setDirectionENS_9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:71
// index:0
// Public
// void addSpacing(int)
func (this *QBoxLayout) AddSpacing(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10addSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:72
// index:0
// Public
// void addStretch(int)
func (this *QBoxLayout) AddStretch(stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10addStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:73
// index:0
// Public
// void addSpacerItem(class QSpacerItem *)
func (this *QBoxLayout) AddSpacerItem(spacerItem unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13addSpacerItemEP11QSpacerItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacerItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// Public
// void addLayout(class QLayout *, int)
func (this *QBoxLayout) AddLayout(layout unsafe.Pointer, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout9addLayoutEP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), layout, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:76
// index:0
// Public
// void addStrut(int)
func (this *QBoxLayout) AddStrut(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout8addStrutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:77
// index:0
// Public virtual
// void addItem(class QLayoutItem *)
func (this *QBoxLayout) AddItem(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:79
// index:0
// Public
// void insertSpacing(int, int)
func (this *QBoxLayout) InsertSpacing(index int, size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13insertSpacingEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:80
// index:0
// Public
// void insertStretch(int, int)
func (this *QBoxLayout) InsertStretch(index int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13insertStretchEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:81
// index:0
// Public
// void insertSpacerItem(int, class QSpacerItem *)
func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, spacerItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:83
// index:0
// Public
// void insertLayout(int, class QLayout *, int)
func (this *QBoxLayout) InsertLayout(index int, layout unsafe.Pointer, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout12insertLayoutEiP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, layout, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:84
// index:0
// Public
// void insertItem(int, class QLayoutItem *)
func (this *QBoxLayout) InsertItem(index int, arg1 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10insertItemEiP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:86
// index:0
// Public
// int spacing()
func (this *QBoxLayout) Spacing() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:87
// index:0
// Public
// void setSpacing(int)
func (this *QBoxLayout) SetSpacing(spacing int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10setSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:89
// index:0
// Public
// bool setStretchFactor(class QWidget *, int)
func (this *QBoxLayout) SetStretchFactor(w unsafe.Pointer, stretch int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, &stretch)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:90
// index:1
// Public
// bool setStretchFactor(class QLayout *, int)
func (this *QBoxLayout) SetStretchFactor_1(l unsafe.Pointer, stretch int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), l, &stretch)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:91
// index:0
// Public
// void setStretch(int, int)
func (this *QBoxLayout) SetStretch(index int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10setStretchEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:92
// index:0
// Public
// int stretch(int)
func (this *QBoxLayout) Stretch(index int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7stretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:94
// index:0
// Public virtual
// QSize sizeHint()
func (this *QBoxLayout) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:95
// index:0
// Public virtual
// QSize minimumSize()
func (this *QBoxLayout) MinimumSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:96
// index:0
// Public virtual
// QSize maximumSize()
func (this *QBoxLayout) MaximumSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:98
// index:0
// Public virtual
// bool hasHeightForWidth()
func (this *QBoxLayout) HasHeightForWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:99
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QBoxLayout) HeightForWidth(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:100
// index:0
// Public virtual
// int minimumHeightForWidth(int)
func (this *QBoxLayout) MinimumHeightForWidth(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout21minimumHeightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:102
// index:0
// Public virtual
// Qt::Orientations expandingDirections()
func (this *QBoxLayout) ExpandingDirections() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:103
// index:0
// Public virtual
// void invalidate()
func (this *QBoxLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:104
// index:0
// Public virtual
// QLayoutItem * itemAt(int)
func (this *QBoxLayout) ItemAt(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:105
// index:0
// Public virtual
// QLayoutItem * takeAt(int)
func (this *QBoxLayout) TakeAt(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:106
// index:0
// Public virtual
// int count()
func (this *QBoxLayout) Count() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qboxlayout.h:107
// index:0
// Public virtual
// void setGeometry(const class QRect &)
func (this *QBoxLayout) SetGeometry(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
