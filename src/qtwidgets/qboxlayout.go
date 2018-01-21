package qtwidgets

// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>

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
	if this == nil {
		return nil
	} else {
		return this.QLayout.GetCthis()
	}
}
func NewQBoxLayoutFromPointer(cthis unsafe.Pointer) *QBoxLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QBoxLayout{bcthis0}
}

// /usr/include/qt/QtWidgets/qboxlayout.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QBoxLayout) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// Public
// void QBoxLayout(enum QBoxLayout::Direction, class QWidget *)
func NewQBoxLayout(arg0 int, parent *QWidget /*444 QWidget **/) *QBoxLayout {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &arg0, convArg1)
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
func (this *QBoxLayout) Direction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
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
func (this *QBoxLayout) AddSpacerItem(spacerItem *QSpacerItem /*444 QSpacerItem **/) {
	var convArg0 = spacerItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13addSpacerItemEP11QSpacerItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// Public
// void addLayout(class QLayout *, int)
func (this *QBoxLayout) AddLayout(layout *QLayout /*444 QLayout **/, stretch int) {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout9addLayoutEP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &stretch)
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
func (this *QBoxLayout) AddItem(arg0 *QLayoutItem /*444 QLayoutItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem *QSpacerItem /*444 QSpacerItem **/) {
	var convArg1 = spacerItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:83
// index:0
// Public
// void insertLayout(int, class QLayout *, int)
func (this *QBoxLayout) InsertLayout(index int, layout *QLayout /*444 QLayout **/, stretch int) {
	var convArg1 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout12insertLayoutEiP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1, &stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:84
// index:0
// Public
// void insertItem(int, class QLayoutItem *)
func (this *QBoxLayout) InsertItem(index int, arg1 *QLayoutItem /*444 QLayoutItem **/) {
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10insertItemEiP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:86
// index:0
// Public
// int spacing()
func (this *QBoxLayout) Spacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
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
func (this *QBoxLayout) SetStretchFactor(w *QWidget /*444 QWidget **/, stretch int) bool {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &stretch)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:90
// index:1
// Public
// bool setStretchFactor(class QLayout *, int)
func (this *QBoxLayout) SetStretchFactor_1(l *QLayout /*444 QLayout **/, stretch int) bool {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &stretch)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
func (this *QBoxLayout) Stretch(index int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7stretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:94
// index:0
// Public virtual
// QSize sizeHint()
func (this *QBoxLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:95
// index:0
// Public virtual
// QSize minimumSize()
func (this *QBoxLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:96
// index:0
// Public virtual
// QSize maximumSize()
func (this *QBoxLayout) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:98
// index:0
// Public virtual
// bool hasHeightForWidth()
func (this *QBoxLayout) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:99
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QBoxLayout) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:100
// index:0
// Public virtual
// int minimumHeightForWidth(int)
func (this *QBoxLayout) MinimumHeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout21minimumHeightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:102
// index:0
// Public virtual
// Qt::Orientations expandingDirections()
func (this *QBoxLayout) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
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
func (this *QBoxLayout) ItemAt(arg0 int) *QLayoutItem /*444 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:105
// index:0
// Public virtual
// QLayoutItem * takeAt(int)
func (this *QBoxLayout) TakeAt(arg0 int) *QLayoutItem /*444 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:106
// index:0
// Public virtual
// int count()
func (this *QBoxLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
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
