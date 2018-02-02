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
func (this *QBoxLayout) SetCthis(cthis unsafe.Pointer) {
	this.QLayout = NewQLayoutFromPointer(cthis)
}
func NewQBoxLayoutFromPointer(cthis unsafe.Pointer) *QBoxLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QBoxLayout{bcthis0}
}
func (*QBoxLayout) NewFromPointer(cthis unsafe.Pointer) *QBoxLayout {
	return NewQBoxLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QBoxLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBoxLayout(enum QBoxLayout::Direction, QWidget *)
func NewQBoxLayout(arg0 int, parent *QWidget /*777 QWidget **/) *QBoxLayout {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", ffiqt.FFI_TYPE_POINTER, arg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QBoxLayout()
func DeleteQBoxLayout(this *QBoxLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayoutD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] QBoxLayout::Direction direction()
func (this *QBoxLayout) Direction() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout9directionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirection(enum QBoxLayout::Direction)
func (this *QBoxLayout) SetDirection(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout12setDirectionENS_9DirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addSpacing(int)
func (this *QBoxLayout) AddSpacing(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10addSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addStretch(int)
func (this *QBoxLayout) AddStretch(stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10addStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addSpacerItem(QSpacerItem *)
func (this *QBoxLayout) AddSpacerItem(spacerItem *QSpacerItem /*777 QSpacerItem **/) {
	var convArg0 = spacerItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13addSpacerItemEP11QSpacerItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addLayout(QLayout *, int)
func (this *QBoxLayout) AddLayout(layout *QLayout /*777 QLayout **/, stretch int) {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout9addLayoutEP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addStrut(int)
func (this *QBoxLayout) AddStrut(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout8addStrutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)
func (this *QBoxLayout) AddItem(arg0 *QLayoutItem /*777 QLayoutItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout7addItemEP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertSpacing(int, int)
func (this *QBoxLayout) InsertSpacing(index int, size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13insertSpacingEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertStretch(int, int)
func (this *QBoxLayout) InsertStretch(index int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout13insertStretchEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertSpacerItem(int, QSpacerItem *)
func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem *QSpacerItem /*777 QSpacerItem **/) {
	var convArg1 = spacerItem.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertLayout(int, QLayout *, int)
func (this *QBoxLayout) InsertLayout(index int, layout *QLayout /*777 QLayout **/, stretch int) {
	var convArg1 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout12insertLayoutEiP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, QLayoutItem *)
func (this *QBoxLayout) InsertItem(index int, arg1 *QLayoutItem /*777 QLayoutItem **/) {
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10insertItemEiP11QLayoutItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing()
func (this *QBoxLayout) Spacing() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7spacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)
func (this *QBoxLayout) SetSpacing(spacing int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10setSpacingEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setStretchFactor(QWidget *, int)
func (this *QBoxLayout) SetStretchFactor(w *QWidget /*777 QWidget **/, stretch int) bool {
	var convArg0 = w.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:90
// index:1
// Public Visibility=Default Availability=Available
// [1] bool setStretchFactor(QLayout *, int)
func (this *QBoxLayout) SetStretchFactor_1(l *QLayout /*777 QLayout **/, stretch int) bool {
	var convArg0 = l.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QLayouti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretch(int, int)
func (this *QBoxLayout) SetStretch(index int, stretch int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10setStretchEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretch(int)
func (this *QBoxLayout) Stretch(index int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout7stretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QBoxLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QBoxLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize()
func (this *QBoxLayout) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QBoxLayout) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QBoxLayout) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int minimumHeightForWidth(int)
func (this *QBoxLayout) MinimumHeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout21minimumHeightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:102
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections()
func (this *QBoxLayout) ExpandingDirections() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout19expandingDirectionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QBoxLayout) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int)
func (this *QBoxLayout) ItemAt(arg0 int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout6itemAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)
func (this *QBoxLayout) TakeAt(arg0 int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout6takeAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count()
func (this *QBoxLayout) Count() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QBoxLayout5countEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QBoxLayout) SetGeometry(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QBoxLayout11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QBoxLayout__Direction = int

const QBoxLayout__LeftToRight QBoxLayout__Direction = 0
const QBoxLayout__RightToLeft QBoxLayout__Direction = 1
const QBoxLayout__TopToBottom QBoxLayout__Direction = 2
const QBoxLayout__BottomToTop QBoxLayout__Direction = 3
const QBoxLayout__Down QBoxLayout__Direction = 2
const QBoxLayout__Up QBoxLayout__Direction = 3

//  body block end
