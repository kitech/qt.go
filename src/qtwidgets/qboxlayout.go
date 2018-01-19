//  header block begin
// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 39
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qboxlayout.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QBoxLayout) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// void QBoxLayout(enum QBoxLayout::Direction, class QWidget *)
func NewQBoxLayout(arg0 int, parent unsafe.Pointer) *QBoxLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, parent)
	gopp.ErrPrint(err, rv)
	return &QBoxLayout{cthis}
}

// /usr/include/qt/QtWidgets/qboxlayout.h:66
// index:0
// virtual
// void ~QBoxLayout()
func DeleteQBoxLayout(*QBoxLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:68
// index:0
// QBoxLayout::Direction direction()
func (this *QBoxLayout) Direction() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout9directionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:69
// index:0
// void setDirection(enum QBoxLayout::Direction)
func (this *QBoxLayout) SetDirection(arg0 int) {
	// 0: (, QBoxLayout::Direction arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout12setDirectionENS_9DirectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:71
// index:0
// void addSpacing(int)
func (this *QBoxLayout) AddSpacing(size int) {
	// 0: (, int size), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10addSpacingEi", ffiqt.FFI_TYPE_VOID, this.cthis, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:72
// index:0
// void addStretch(int)
func (this *QBoxLayout) AddStretch(stretch int) {
	// 0: (, int stretch), (&stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10addStretchEi", ffiqt.FFI_TYPE_VOID, this.cthis, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:73
// index:0
// void addSpacerItem(class QSpacerItem *)
func (this *QBoxLayout) AddSpacerItem(spacerItem unsafe.Pointer) {
	// 0: (, QSpacerItem * spacerItem), (spacerItem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13addSpacerItemEP11QSpacerItem", ffiqt.FFI_TYPE_VOID, this.cthis, spacerItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// void addLayout(class QLayout *, int)
func (this *QBoxLayout) AddLayout(layout unsafe.Pointer, stretch int) {
	// 0: (, QLayout * layout, int stretch), (layout, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout9addLayoutEP7QLayouti", ffiqt.FFI_TYPE_VOID, this.cthis, layout, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:76
// index:0
// void addStrut(int)
func (this *QBoxLayout) AddStrut(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout8addStrutEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:77
// index:0
// virtual
// void addItem(class QLayoutItem *)
func (this *QBoxLayout) AddItem(arg0 unsafe.Pointer) {
	// 0: (, QLayoutItem * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:79
// index:0
// void insertSpacing(int, int)
func (this *QBoxLayout) InsertSpacing(index int, size int) {
	// 0: (, int index, int size), (&index, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13insertSpacingEii", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:80
// index:0
// void insertStretch(int, int)
func (this *QBoxLayout) InsertStretch(index int, stretch int) {
	// 0: (, int index, int stretch), (&index, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13insertStretchEii", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:81
// index:0
// void insertSpacerItem(int, class QSpacerItem *)
func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem unsafe.Pointer) {
	// 0: (, int index, QSpacerItem * spacerItem), (&index, spacerItem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem", ffiqt.FFI_TYPE_VOID, this.cthis, &index, spacerItem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:83
// index:0
// void insertLayout(int, class QLayout *, int)
func (this *QBoxLayout) InsertLayout(index int, layout unsafe.Pointer, stretch int) {
	// 0: (, int index, QLayout * layout, int stretch), (&index, layout, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout12insertLayoutEiP7QLayouti", ffiqt.FFI_TYPE_VOID, this.cthis, &index, layout, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:84
// index:0
// void insertItem(int, class QLayoutItem *)
func (this *QBoxLayout) InsertItem(index int, arg1 unsafe.Pointer) {
	// 0: (, int index, QLayoutItem * arg1), (&index, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10insertItemEiP11QLayoutItem", ffiqt.FFI_TYPE_VOID, this.cthis, &index, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:86
// index:0
// int spacing()
func (this *QBoxLayout) Spacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7spacingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:87
// index:0
// void setSpacing(int)
func (this *QBoxLayout) SetSpacing(spacing int) {
	// 0: (, int spacing), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10setSpacingEi", ffiqt.FFI_TYPE_VOID, this.cthis, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:89
// index:0
// bool setStretchFactor(class QWidget *, int)
func (this *QBoxLayout) SetStretchFactor(w unsafe.Pointer, stretch int) {
	// 0: (, QWidget * w, int stretch), (w, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QWidgeti", ffiqt.FFI_TYPE_VOID, this.cthis, w, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:90
// index:1
// bool setStretchFactor(class QLayout *, int)
func (this *QBoxLayout) SetStretchFactor_1(l unsafe.Pointer, stretch int) {
	// 1: (, QLayout * l, int stretch), (l, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QLayouti", ffiqt.FFI_TYPE_VOID, this.cthis, l, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:91
// index:0
// void setStretch(int, int)
func (this *QBoxLayout) SetStretch(index int, stretch int) {
	// 0: (, int index, int stretch), (&index, &stretch)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10setStretchEii", ffiqt.FFI_TYPE_VOID, this.cthis, &index, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:92
// index:0
// int stretch(int)
func (this *QBoxLayout) Stretch(index int) {
	// 0: (, int index), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7stretchEi", ffiqt.FFI_TYPE_VOID, this.cthis, &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:94
// index:0
// virtual
// QSize sizeHint()
func (this *QBoxLayout) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:95
// index:0
// virtual
// QSize minimumSize()
func (this *QBoxLayout) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:96
// index:0
// virtual
// QSize maximumSize()
func (this *QBoxLayout) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:98
// index:0
// virtual
// bool hasHeightForWidth()
func (this *QBoxLayout) HasHeightForWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:99
// index:0
// virtual
// int heightForWidth(int)
func (this *QBoxLayout) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:100
// index:0
// virtual
// int minimumHeightForWidth(int)
func (this *QBoxLayout) MinimumHeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout21minimumHeightForWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:102
// index:0
// virtual
// Qt::Orientations expandingDirections()
func (this *QBoxLayout) ExpandingDirections() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:103
// index:0
// virtual
// void invalidate()
func (this *QBoxLayout) Invalidate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10invalidateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:104
// index:0
// virtual
// QLayoutItem * itemAt(int)
func (this *QBoxLayout) ItemAt(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout6itemAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:105
// index:0
// virtual
// QLayoutItem * takeAt(int)
func (this *QBoxLayout) TakeAt(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout6takeAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:106
// index:0
// virtual
// int count()
func (this *QBoxLayout) Count() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout5countEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:107
// index:0
// virtual
// void setGeometry(const class QRect &)
func (this *QBoxLayout) SetGeometry(arg0 unsafe.Pointer) {
	// 0: (, const QRect & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
