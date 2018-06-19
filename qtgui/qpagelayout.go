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
// extern C begin: 33
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
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

/*
Creates an invalid QPageLayout.
*/
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
// [-2] void QPageLayout(const QPageSize &, QPageLayout::Orientation, const QMarginsF &, QPageLayout::Unit, const QMarginsF &)

/*
Creates an invalid QPageLayout.
*/
func NewQPageLayout_1(pageSize QPageSize_ITF, orientation int, margins qtcore.QMarginsF_ITF, units int, minMargins qtcore.QMarginsF_ITF) *QPageLayout {
	var convArg0 unsafe.Pointer
	if pageSize != nil && pageSize.QPageSize_PTR() != nil {
		convArg0 = pageSize.QPageSize_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg2 = margins.QMarginsF_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if minMargins != nil && minMargins.QMarginsF_PTR() != nil {
		convArg4 = minMargins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutC2ERK9QPageSizeNS_11OrientationERK9QMarginsFNS_4UnitES6_", qtrt.FFI_TYPE_POINTER, convArg0, orientation, convArg2, units, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageLayout)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPageLayout(const QPageSize &, QPageLayout::Orientation, const QMarginsF &, QPageLayout::Unit, const QMarginsF &)

/*
Creates an invalid QPageLayout.
*/
func NewQPageLayout_1_(pageSize QPageSize_ITF, orientation int, margins qtcore.QMarginsF_ITF) *QPageLayout {
	var convArg0 unsafe.Pointer
	if pageSize != nil && pageSize.QPageSize_PTR() != nil {
		convArg0 = pageSize.QPageSize_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg2 = margins.QMarginsF_PTR().GetCthis()
	}
	// arg: 3, QPageLayout::Unit=Enum, QPageLayout::Unit=Enum, , Invalid
	units := 0
	// arg: 4, const QMarginsF &=LValueReference, QMarginsF=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutC2ERK9QPageSizeNS_11OrientationERK9QMarginsFNS_4UnitES6_", qtrt.FFI_TYPE_POINTER, convArg0, orientation, convArg2, units, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageLayout)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPageLayout(const QPageSize &, QPageLayout::Orientation, const QMarginsF &, QPageLayout::Unit, const QMarginsF &)

/*
Creates an invalid QPageLayout.
*/
func NewQPageLayout_1_1(pageSize QPageSize_ITF, orientation int, margins qtcore.QMarginsF_ITF, units int) *QPageLayout {
	var convArg0 unsafe.Pointer
	if pageSize != nil && pageSize.QPageSize_PTR() != nil {
		convArg0 = pageSize.QPageSize_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg2 = margins.QMarginsF_PTR().GetCthis()
	}
	// arg: 4, const QMarginsF &=LValueReference, QMarginsF=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutC2ERK9QPageSizeNS_11OrientationERK9QMarginsFNS_4UnitES6_", qtrt.FFI_TYPE_POINTER, convArg0, orientation, convArg2, units, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageLayout)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:86
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPageLayout & operator=(QPageLayout &&)

/*

 */
func (this *QPageLayout) Operator_equal(other unsafe.Pointer /*333*/) *QPageLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:88
// index:1
// Public Visibility=Default Availability=Available
// [8] QPageLayout & operator=(const QPageLayout &)

/*

 */
func (this *QPageLayout) Operator_equal_1(other QPageLayout_ITF) *QPageLayout {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPageLayout_PTR() != nil {
		convArg0 = other.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayoutaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

// /usr/include/qt/QtGui/qpagelayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPageLayout()

/*

 */
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

/*
Swaps this page layout with other. This function is very fast and never fails.
*/
func (this *QPageLayout) Swap(other QPageLayout_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPageLayout_PTR() != nil {
		convArg0 = other.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEquivalentTo(const QPageLayout &) const

/*
Returns true if this page layout is equivalent to the other page layout, i.e. if the page has the same size, margins and orientation.
*/
func (this *QPageLayout) IsEquivalentTo(other QPageLayout_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPageLayout_PTR() != nil {
		convArg0 = other.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout14isEquivalentToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this page layout is valid.
*/
func (this *QPageLayout) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMode(QPageLayout::Mode)

/*
Sets a page layout mode to mode.

See also mode().
*/
func (this *QPageLayout) SetMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout7setModeENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageLayout::Mode mode() const

/*
Returns the page layout mode.

See also setMode().
*/
func (this *QPageLayout) Mode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout4modeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPageSize(const QPageSize &, const QMarginsF &)

/*
Sets the page size of the page layout to pageSize.

Optionally define the minimum allowed margins minMargins, e.g. the minimum margins able to be printed by a physical print device, otherwise the minimum margins will default to 0.

If StandardMode is set then the existing margins will be clamped to the new minimum margins and the maximum margins allowed by the page size. If FullPageMode is set then the existing margins will be unchanged.

See also pageSize().
*/
func (this *QPageLayout) SetPageSize(pageSize QPageSize_ITF, minMargins qtcore.QMarginsF_ITF) {
	var convArg0 unsafe.Pointer
	if pageSize != nil && pageSize.QPageSize_PTR() != nil {
		convArg0 = pageSize.QPageSize_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if minMargins != nil && minMargins.QMarginsF_PTR() != nil {
		convArg1 = minMargins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPageSize(const QPageSize &, const QMarginsF &)

/*
Sets the page size of the page layout to pageSize.

Optionally define the minimum allowed margins minMargins, e.g. the minimum margins able to be printed by a physical print device, otherwise the minimum margins will default to 0.

If StandardMode is set then the existing margins will be clamped to the new minimum margins and the maximum margins allowed by the page size. If FullPageMode is set then the existing margins will be unchanged.

See also pageSize().
*/
func (this *QPageLayout) SetPageSize__(pageSize QPageSize_ITF) {
	var convArg0 unsafe.Pointer
	if pageSize != nil && pageSize.QPageSize_PTR() != nil {
		convArg0 = pageSize.QPageSize_PTR().GetCthis()
	}
	// arg: 1, const QMarginsF &=LValueReference, QMarginsF=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QPageSize pageSize() const

/*
Returns the page size of the page layout.

Note that the QPageSize is always defined in a Portrait orientation. To obtain a size that takes the set orientation into account you must use fullRect().

See also setPageSize().
*/
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
// [-2] void setOrientation(QPageLayout::Orientation)

/*
Sets the page orientation of the page layout to orientation.

Changing the orientation does not affect the current margins or the minimum margins.

See also orientation().
*/
func (this *QPageLayout) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout14setOrientationENS_11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:106
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageLayout::Orientation orientation() const

/*
Returns the page orientation of the page layout.

See also setOrientation().
*/
func (this *QPageLayout) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnits(QPageLayout::Unit)

/*
Sets the units used to define the page layout.

See also units().
*/
func (this *QPageLayout) SetUnits(units int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout8setUnitsENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageLayout::Unit units() const

/*
Returns the units the page layout is currently defined in.

See also setUnits().
*/
func (this *QPageLayout) Units() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout5unitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setMargins(const QMarginsF &)

/*
Sets the page margins of the page layout to margins Returns true if the margins were successfully set.

The units used are those currently defined for the layout. To use different units then call setUnits() first.

If in the default StandardMode then all the new margins must fall between the minimum margins set and the maximum margins allowed by the page size, otherwise the margins will not be set.

If in FullPageMode then any margin values will be accepted.

See also margins() and units().
*/
func (this *QPageLayout) SetMargins(margins qtcore.QMarginsF_ITF) bool {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg0 = margins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout10setMarginsERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setLeftMargin(qreal)

/*
Sets the left page margin of the page layout to leftMargin. Returns true if the margin was successfully set.

The units used are those currently defined for the layout. To use different units call setUnits() first.

If in the default StandardMode then the new margin must fall between the minimum margin set and the maximum margin allowed by the page size, otherwise the margin will not be set.

If in FullPageMode then any margin values will be accepted.

See also setMargins() and margins().
*/
func (this *QPageLayout) SetLeftMargin(leftMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout13setLeftMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), leftMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setRightMargin(qreal)

/*
Sets the right page margin of the page layout to rightMargin. Returns true if the margin was successfully set.

The units used are those currently defined for the layout. To use different units call setUnits() first.

If in the default StandardMode then the new margin must fall between the minimum margin set and the maximum margin allowed by the page size, otherwise the margin will not be set.

If in FullPageMode then any margin values will be accepted.

See also setMargins() and margins().
*/
func (this *QPageLayout) SetRightMargin(rightMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout14setRightMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rightMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:114
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setTopMargin(qreal)

/*
Sets the top page margin of the page layout to topMargin. Returns true if the margin was successfully set.

The units used are those currently defined for the layout. To use different units call setUnits() first.

If in the default StandardMode then the new margin must fall between the minimum margin set and the maximum margin allowed by the page size, otherwise the margin will not be set.

If in FullPageMode then any margin values will be accepted.

See also setMargins() and margins().
*/
func (this *QPageLayout) SetTopMargin(topMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout12setTopMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), topMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:115
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setBottomMargin(qreal)

/*
Sets the bottom page margin of the page layout to bottomMargin. Returns true if the margin was successfully set.

The units used are those currently defined for the layout. To use different units call setUnits() first.

If in the default StandardMode then the new margin must fall between the minimum margin set and the maximum margin allowed by the page size, otherwise the margin will not be set.

If in FullPageMode then any margin values will be accepted.

See also setMargins() and margins().
*/
func (this *QPageLayout) SetBottomMargin(bottomMargin float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout15setBottomMarginEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), bottomMargin)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:117
// index:0
// Public Visibility=Default Availability=Available
// [32] QMarginsF margins() const

/*
Returns the margins of the page layout using the currently set units.

See also setMargins() and units().
*/
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
// [32] QMarginsF margins(QPageLayout::Unit) const

/*
Returns the margins of the page layout using the currently set units.

See also setMargins() and units().
*/
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
// [16] QMargins marginsPoints() const

/*
Returns the margins of the page layout in Postscript Points (1/72 of an inch).

See also setMargins() and margins().
*/
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
// [16] QMargins marginsPixels(int) const

/*
Returns the margins of the page layout in device pixels for the given resolution.

See also setMargins().
*/
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

/*
Sets the minimum page margins of the page layout to minMargins.

It is not recommended to override the default values set for a page size as this may be the minimum printable area for a physical print device.

If the StandardMode mode is set then the existing margins will be clamped to the new minMargins and the maximum allowed by the page size. If the FullPageMode is set then the existing margins will be unchanged.

See also minimumMargins() and setMargins().
*/
func (this *QPageLayout) SetMinimumMargins(minMargins qtcore.QMarginsF_ITF) {
	var convArg0 unsafe.Pointer
	if minMargins != nil && minMargins.QMarginsF_PTR() != nil {
		convArg0 = minMargins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:123
// index:0
// Public Visibility=Default Availability=Available
// [32] QMarginsF minimumMargins() const

/*
Returns the minimum margins of the page layout.

See also setMinimumMargins() and maximumMargins().
*/
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
// [32] QMarginsF maximumMargins() const

/*
Returns the maximum margins that would be applied if the page layout was in StandardMode.

The maximum margins allowed are calculated as the full size of the page minus the minimum margins set. For example, if the page width is 100 points and the minimum right margin is 10 points, then the maximum left margin will be 90 points.

See also setMinimumMargins() and minimumMargins().
*/
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
// [32] QRectF fullRect() const

/*
Returns the full page rectangle in the current layout units.

The page rectangle takes into account the page size and page orientation, but not the page margins.

See also paintRect() and units().
*/
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
// [32] QRectF fullRect(QPageLayout::Unit) const

/*
Returns the full page rectangle in the current layout units.

The page rectangle takes into account the page size and page orientation, but not the page margins.

See also paintRect() and units().
*/
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
// [16] QRect fullRectPoints() const

/*
Returns the full page rectangle in Postscript Points (1/72 of an inch).

The page rectangle takes into account the page size and page orientation, but not the page margins.

See also paintRect().
*/
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
// [16] QRect fullRectPixels(int) const

/*
Returns the full page rectangle in device pixels for the given resolution.

The page rectangle takes into account the page size and page orientation, but not the page margins.

See also paintRect().
*/
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
// [32] QRectF paintRect() const

/*
Returns the page rectangle in the current layout units.

The paintable rectangle takes into account the page size, orientation and margins.

If the FullPageMode mode is set then the fullRect() is returned and the margins must be manually managed.
*/
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
// [32] QRectF paintRect(QPageLayout::Unit) const

/*
Returns the page rectangle in the current layout units.

The paintable rectangle takes into account the page size, orientation and margins.

If the FullPageMode mode is set then the fullRect() is returned and the margins must be manually managed.
*/
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
// [16] QRect paintRectPoints() const

/*
Returns the paintable rectangle in rounded Postscript Points (1/72 of an inch).

The paintable rectangle takes into account the page size, orientation and margins.

If the FullPageMode mode is set then the fullRect() is returned and the margins must be manually managed.
*/
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
// [16] QRect paintRectPixels(int) const

/*
Returns the paintable rectangle in rounded device pixels for the given resolution.

The paintable rectangle takes into account the page size, orientation and margins.

If the FullPageMode mode is set then the fullRect() is returned and the margins must be manually managed.
*/
func (this *QPageLayout) PaintRectPixels(resolution int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPageLayout15paintRectPixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

/*
This enum type is used to specify the measurement unit for page layout and margins.


*/
type QPageLayout__Unit = int

//
const QPageLayout__Millimeter QPageLayout__Unit = 0

//
const QPageLayout__Point QPageLayout__Unit = 1

//
const QPageLayout__Inch QPageLayout__Unit = 2

//
const QPageLayout__Pica QPageLayout__Unit = 3

//
const QPageLayout__Didot QPageLayout__Unit = 4

//
const QPageLayout__Cicero QPageLayout__Unit = 5

/*
This enum type defines the page orientation



Note that some standard page sizes are defined with a width larger than their height, hence the orientation is defined relative to the standard page size and not using the relative page dimensions.

*/
type QPageLayout__Orientation = int

// The page size is used in its default orientation
const QPageLayout__Portrait QPageLayout__Orientation = 0

//
const QPageLayout__Landscape QPageLayout__Orientation = 1

/*
Defines the page layout mode


*/
type QPageLayout__Mode = int

// Paint Rect includes margins, margins must fall between the minimum and maximum.
const QPageLayout__StandardMode QPageLayout__Mode = 0

// Paint Rect excludes margins, margins can be any value and must be managed manually.
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
