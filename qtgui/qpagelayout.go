package qtgui

// /usr/include/qt/QtGui/qpagelayout.h
// #include <qpagelayout.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QPageLayout struct {
	*qtrt.CObject
}

func (this *QPageLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPageLayout) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPageLayoutFromPointer(cthis unsafe.Pointer) *QPageLayout {
	return &QPageLayout{&qtrt.CObject{cthis}}
}
func (*QPageLayout) NewFromPointer(cthis unsafe.Pointer) *QPageLayout {
	return NewQPageLayoutFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpagelayout.h:80
// index:0
// Public
// void QPageLayout()
func NewQPageLayout() *QPageLayout {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:81
// index:1
// Public
// void QPageLayout(const class QPageSize &, enum QPageLayout::Orientation, const class QMarginsF &, enum QPageLayout::Unit, const class QMarginsF &)
func NewQPageLayout_1(pageSize *QPageSize, orientation int, margins *qtcore.QMarginsF, units int, minMargins *qtcore.QMarginsF) *QPageLayout {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pageSize.GetCthis()
	var convArg2 = margins.GetCthis()
	var convArg4 = minMargins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayoutC2ERK9QPageSizeNS_11OrientationERK9QMarginsFNS_4UnitES6_", ffiqt.FFI_TYPE_VOID, cthis, convArg0, orientation, convArg2, units, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:89
// index:0
// Public
// void ~QPageLayout()
func DeleteQPageLayout(*QPageLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:91
// index:0
// Public inline
// void swap(class QPageLayout &)
func (this *QPageLayout) Swap(other *QPageLayout) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:94
// index:0
// Public
// bool isEquivalentTo(const class QPageLayout &)
func (this *QPageLayout) IsEquivalentTo(other *QPageLayout) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14isEquivalentToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:96
// index:0
// Public
// bool isValid()
func (this *QPageLayout) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:98
// index:0
// Public
// void setMode(enum QPageLayout::Mode)
func (this *QPageLayout) SetMode(mode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout7setModeENS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:99
// index:0
// Public
// QPageLayout::Mode mode()
func (this *QPageLayout) Mode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout4modeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:101
// index:0
// Public
// void setPageSize(const class QPageSize &, const class QMarginsF &)
func (this *QPageLayout) SetPageSize(pageSize *QPageSize, minMargins *qtcore.QMarginsF) {
	var convArg0 = pageSize.GetCthis()
	var convArg1 = minMargins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:103
// index:0
// Public
// QPageSize pageSize()
func (this *QPageLayout) PageSize() *QPageSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout8pageSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:105
// index:0
// Public
// void setOrientation(enum QPageLayout::Orientation)
func (this *QPageLayout) SetOrientation(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout14setOrientationENS_11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:106
// index:0
// Public
// QPageLayout::Orientation orientation()
func (this *QPageLayout) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:108
// index:0
// Public
// void setUnits(enum QPageLayout::Unit)
func (this *QPageLayout) SetUnits(units int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout8setUnitsENS_4UnitE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:109
// index:0
// Public
// QPageLayout::Unit units()
func (this *QPageLayout) Units() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout5unitsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:111
// index:0
// Public
// bool setMargins(const class QMarginsF &)
func (this *QPageLayout) SetMargins(margins *qtcore.QMarginsF) bool {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout10setMarginsERK9QMarginsF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:112
// index:0
// Public
// bool setLeftMargin(qreal)
func (this *QPageLayout) SetLeftMargin(leftMargin float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout13setLeftMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), leftMargin)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:113
// index:0
// Public
// bool setRightMargin(qreal)
func (this *QPageLayout) SetRightMargin(rightMargin float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout14setRightMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rightMargin)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:114
// index:0
// Public
// bool setTopMargin(qreal)
func (this *QPageLayout) SetTopMargin(topMargin float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout12setTopMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), topMargin)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:115
// index:0
// Public
// bool setBottomMargin(qreal)
func (this *QPageLayout) SetBottomMargin(bottomMargin float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout15setBottomMarginEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), bottomMargin)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:117
// index:0
// Public
// QMarginsF margins()
func (this *QPageLayout) Margins() *qtcore.QMarginsF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout7marginsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:118
// index:1
// Public
// QMarginsF margins(enum QPageLayout::Unit)
func (this *QPageLayout) Margins_1(units int) *qtcore.QMarginsF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout7marginsENS_4UnitE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), units)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:119
// index:0
// Public
// QMargins marginsPoints()
func (this *QPageLayout) MarginsPoints() *qtcore.QMargins /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout13marginsPointsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:120
// index:0
// Public
// QMargins marginsPixels(int)
func (this *QPageLayout) MarginsPixels(resolution int) *qtcore.QMargins /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout13marginsPixelsEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), resolution)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:122
// index:0
// Public
// void setMinimumMargins(const class QMarginsF &)
func (this *QPageLayout) SetMinimumMargins(minMargins *qtcore.QMarginsF) {
	var convArg0 = minMargins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:123
// index:0
// Public
// QMarginsF minimumMargins()
func (this *QPageLayout) MinimumMargins() *qtcore.QMarginsF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14minimumMarginsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:124
// index:0
// Public
// QMarginsF maximumMargins()
func (this *QPageLayout) MaximumMargins() *qtcore.QMarginsF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14maximumMarginsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:126
// index:0
// Public
// QRectF fullRect()
func (this *QPageLayout) FullRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout8fullRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:127
// index:1
// Public
// QRectF fullRect(enum QPageLayout::Unit)
func (this *QPageLayout) FullRect_1(units int) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout8fullRectENS_4UnitE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), units)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:128
// index:0
// Public
// QRect fullRectPoints()
func (this *QPageLayout) FullRectPoints() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14fullRectPointsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:129
// index:0
// Public
// QRect fullRectPixels(int)
func (this *QPageLayout) FullRectPixels(resolution int) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14fullRectPixelsEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), resolution)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:131
// index:0
// Public
// QRectF paintRect()
func (this *QPageLayout) PaintRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout9paintRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:132
// index:1
// Public
// QRectF paintRect(enum QPageLayout::Unit)
func (this *QPageLayout) PaintRect_1(units int) *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout9paintRectENS_4UnitE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), units)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:133
// index:0
// Public
// QRect paintRectPoints()
func (this *QPageLayout) PaintRectPoints() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout15paintRectPointsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:134
// index:0
// Public
// QRect paintRectPixels(int)
func (this *QPageLayout) PaintRectPixels(resolution int) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout15paintRectPixelsEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), resolution)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

type QPageLayout__Unit = int

const QPageLayout__Millimeter QPageLayout__Unit = 0
const QPageLayout__Point QPageLayout__Unit = 1
const QPageLayout__Inch QPageLayout__Unit = 2
const QPageLayout__Pica QPageLayout__Unit = 3
const QPageLayout__Didot QPageLayout__Unit = 4
const QPageLayout__Cicero QPageLayout__Unit = 5

type QPageLayout__Orientation = int

const QPageLayout__Portrait QPageLayout__Orientation = 0
const QPageLayout__Landscape QPageLayout__Orientation = 1

type QPageLayout__Mode = int

const QPageLayout__StandardMode QPageLayout__Mode = 0
const QPageLayout__FullPageMode QPageLayout__Mode = 1

//  body block end
