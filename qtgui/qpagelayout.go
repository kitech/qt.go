package qtgui

// /usr/include/qt/QtGui/qpagelayout.h
// #include <qpagelayout.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QPageLayout struct {
	*qtrt.CObject
}
type QPageLayout_ITF interface {
	QPageLayout_PTR() *QPageLayout
}

func (ptr *QPageLayout) QPageLayout_PTR() *QPageLayout { return ptr }

func (this *QPageLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPageLayout) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPageLayoutFromPointer(cthis unsafe.Pointer) *QPageLayout {
	return &QPageLayout{&qtrt.CObject{cthis}}
}
func (*QPageLayout) NewFromPointer(cthis unsafe.Pointer) *QPageLayout {
	return NewQPageLayoutFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpagelayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPageLayout()
func NewQPageLayout() *QPageLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageLayout)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPageLayout(const QPageSize &, enum QPageLayout::Orientation, const QMarginsF &, enum QPageLayout::Unit, const QMarginsF &)
func NewQPageLayout_1(pageSize QPageSize_ITF, orientation int, margins qtcore.QMarginsF_ITF, units int, minMargins qtcore.QMarginsF_ITF) *QPageLayout {
	var convArg0 = pageSize.QPageSize_PTR().GetCthis()
	var convArg2 = margins.QMarginsF_PTR().GetCthis()
	var convArg4 = minMargins.QMarginsF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutC2ERK9QPageSizeNS_11OrientationERK9QMarginsFNS_4UnitES6_", qtrt.FFI_TYPE_POINTER, convArg0, orientation, convArg2, units, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageLayout)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPageLayout()
func DeleteQPageLayout(this *QPageLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpagelayout.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPageLayout &)
func (this *QPageLayout) Swap(other QPageLayout_ITF) {
	var convArg0 = other.QPageLayout_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEquivalentTo(const QPageLayout &)
func (this *QPageLayout) IsEquivalentTo(other QPageLayout_ITF) bool {
	var convArg0 = other.QPageLayout_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout14isEquivalentToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QPageLayout) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMode(enum QPageLayout::Mode)
func (this *QPageLayout) SetMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout7setModeENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageLayout::Mode mode()
func (this *QPageLayout) Mode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout4modeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPageSize(const QPageSize &, const QMarginsF &)
func (this *QPageLayout) SetPageSize(pageSize QPageSize_ITF, minMargins qtcore.QMarginsF_ITF) {
	var convArg0 = pageSize.QPageSize_PTR().GetCthis()
	var convArg1 = minMargins.QMarginsF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QPageSize pageSize()
func (this *QPageLayout) PageSize() *QPageSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout8pageSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageSize)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(enum QPageLayout::Orientation)
func (this *QPageLayout) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout14setOrientationENS_11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageLayout::Orientation orientation()
func (this *QPageLayout) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnits(enum QPageLayout::Unit)
func (this *QPageLayout) SetUnits(units int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout8setUnitsENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageLayout::Unit units()
func (this *QPageLayout) Units() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout5unitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setMargins(const QMarginsF &)
func (this *QPageLayout) SetMargins(margins qtcore.QMarginsF_ITF) bool {
	var convArg0 = margins.QMarginsF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout10setMarginsERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setLeftMargin(qreal)
func (this *QPageLayout) SetLeftMargin(leftMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout13setLeftMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), leftMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setRightMargin(qreal)
func (this *QPageLayout) SetRightMargin(rightMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout14setRightMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rightMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setTopMargin(qreal)
func (this *QPageLayout) SetTopMargin(topMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout12setTopMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), topMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setBottomMargin(qreal)
func (this *QPageLayout) SetBottomMargin(bottomMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout15setBottomMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottomMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:117
// index:0
// Public Visibility=Default Availability=Available
// [32] QMarginsF margins()
func (this *QPageLayout) Margins() *qtcore.QMarginsF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout7marginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMarginsF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:118
// index:1
// Public Visibility=Default Availability=Available
// [32] QMarginsF margins(enum QPageLayout::Unit)
func (this *QPageLayout) Margins_1(units int) *qtcore.QMarginsF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout7marginsENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMarginsF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:119
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins marginsPoints()
func (this *QPageLayout) MarginsPoints() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout13marginsPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:120
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins marginsPixels(int)
func (this *QPageLayout) MarginsPixels(resolution int) *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout13marginsPixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumMargins(const QMarginsF &)
func (this *QPageLayout) SetMinimumMargins(minMargins qtcore.QMarginsF_ITF) {
	var convArg0 = minMargins.QMarginsF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:123
// index:0
// Public Visibility=Default Availability=Available
// [32] QMarginsF minimumMargins()
func (this *QPageLayout) MinimumMargins() *qtcore.QMarginsF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout14minimumMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMarginsF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:124
// index:0
// Public Visibility=Default Availability=Available
// [32] QMarginsF maximumMargins()
func (this *QPageLayout) MaximumMargins() *qtcore.QMarginsF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout14maximumMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMarginsF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:126
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF fullRect()
func (this *QPageLayout) FullRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout8fullRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:127
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF fullRect(enum QPageLayout::Unit)
func (this *QPageLayout) FullRect_1(units int) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout8fullRectENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:128
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect fullRectPoints()
func (this *QPageLayout) FullRectPoints() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout14fullRectPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:129
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect fullRectPixels(int)
func (this *QPageLayout) FullRectPixels(resolution int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout14fullRectPixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:131
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF paintRect()
func (this *QPageLayout) PaintRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout9paintRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:132
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF paintRect(enum QPageLayout::Unit)
func (this *QPageLayout) PaintRect_1(units int) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout9paintRectENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:133
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect paintRectPoints()
func (this *QPageLayout) PaintRectPoints() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout15paintRectPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:134
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect paintRectPixels(int)
func (this *QPageLayout) PaintRectPixels(resolution int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout15paintRectPixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
