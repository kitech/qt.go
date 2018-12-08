package qtgui

// /usr/include/qt/QtGui/qpagesize.h
// #include <qpagesize.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QPageSize struct {
	*qtrt.CObject
}
type QPageSize_ITF interface {
	QPageSize_PTR() *QPageSize
}

func (ptr *QPageSize) QPageSize_PTR() *QPageSize { return ptr }

func (this *QPageSize) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPageSize) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPageSizeFromPointer(cthis unsafe.Pointer) *QPageSize {
	return &QPageSize{&qtrt.CObject{cthis}}
}
func (*QPageSize) NewFromPointer(cthis unsafe.Pointer) *QPageSize {
	return NewQPageSizeFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpagesize.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPageSize()

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit() *QPageSize {
	return NewQPageSize()
}
func NewQPageSize() *QPageSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:231
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(QPageSize::PageSizeId)

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit1(pageSizeId int) *QPageSize {
	return NewQPageSize1(pageSizeId)
}
func NewQPageSize1(pageSizeId int) *QPageSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:232
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSize &, const QString &, QPageSize::SizeMatchPolicy)

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit2(pointSize qtcore.QSize_ITF, name string, matchPolicy int) *QPageSize {
	return NewQPageSize2(pointSize, name, matchPolicy)
}
func NewQPageSize2(pointSize qtcore.QSize_ITF, name string, matchPolicy int) *QPageSize {
	var convArg0 unsafe.Pointer
	if pointSize != nil && pointSize.QSize_PTR() != nil {
		convArg0 = pointSize.QSize_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK5QSizeRK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, matchPolicy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:232
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSize &, const QString &, QPageSize::SizeMatchPolicy)

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit2p(pointSize qtcore.QSize_ITF) *QPageSize {
	return NewQPageSize2p(pointSize)
}
func NewQPageSize2p(pointSize qtcore.QSize_ITF) *QPageSize {
	var convArg0 unsafe.Pointer
	if pointSize != nil && pointSize.QSize_PTR() != nil {
		convArg0 = pointSize.QSize_PTR().GetCthis()
	}
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = qtcore.NewQString()
	// arg: 2, QPageSize::SizeMatchPolicy=Enum, QPageSize::SizeMatchPolicy=Enum, , Invalid
	matchPolicy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK5QSizeRK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, matchPolicy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:232
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSize &, const QString &, QPageSize::SizeMatchPolicy)

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit2p1(pointSize qtcore.QSize_ITF, name string) *QPageSize {
	return NewQPageSize2p1(pointSize, name)
}
func NewQPageSize2p1(pointSize qtcore.QSize_ITF, name string) *QPageSize {
	var convArg0 unsafe.Pointer
	if pointSize != nil && pointSize.QSize_PTR() != nil {
		convArg0 = pointSize.QSize_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QPageSize::SizeMatchPolicy=Enum, QPageSize::SizeMatchPolicy=Enum, , Invalid
	matchPolicy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK5QSizeRK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, matchPolicy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:235
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSizeF &, QPageSize::Unit, const QString &, QPageSize::SizeMatchPolicy)

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit3(size qtcore.QSizeF_ITF, units int, name string, matchPolicy int) *QPageSize {
	return NewQPageSize3(size, units, name, matchPolicy)
}
func NewQPageSize3(size qtcore.QSizeF_ITF, units int, name string, matchPolicy int) *QPageSize {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(name)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK6QSizeFNS_4UnitERK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, units, convArg2, matchPolicy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:235
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSizeF &, QPageSize::Unit, const QString &, QPageSize::SizeMatchPolicy)

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit3p(size qtcore.QSizeF_ITF, units int) *QPageSize {
	return NewQPageSize3p(size, units)
}
func NewQPageSize3p(size qtcore.QSizeF_ITF, units int) *QPageSize {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = qtcore.NewQString()
	// arg: 3, QPageSize::SizeMatchPolicy=Enum, QPageSize::SizeMatchPolicy=Enum, , Invalid
	matchPolicy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK6QSizeFNS_4UnitERK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, units, convArg2, matchPolicy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:235
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSizeF &, QPageSize::Unit, const QString &, QPageSize::SizeMatchPolicy)

/*
Creates a null QPageSize.
*/
func (*QPageSize) NewForInherit3p1(size qtcore.QSizeF_ITF, units int, name string) *QPageSize {
	return NewQPageSize3p1(size, units, name)
}
func NewQPageSize3p1(size qtcore.QSizeF_ITF, units int, name string) *QPageSize {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(name)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QPageSize::SizeMatchPolicy=Enum, QPageSize::SizeMatchPolicy=Enum, , Invalid
	matchPolicy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK6QSizeFNS_4UnitERK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, units, convArg2, matchPolicy)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:240
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPageSize & operator=(QPageSize &&)

/*

 */
func (this *QPageSize) Operator_equal(other unsafe.Pointer /*333*/) *QPageSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageSize)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:242
// index:1
// Public Visibility=Default Availability=Available
// [8] QPageSize & operator=(const QPageSize &)

/*

 */
func (this *QPageSize) Operator_equal1(other QPageSize_ITF) *QPageSize {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPageSize_PTR() != nil {
		convArg0 = other.QPageSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageSize)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPageSize()

/*

 */
func DeleteQPageSize(this *QPageSize) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpagesize.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPageSize &)

/*
Swaps this QPageSize with other. This function is very fast and never fails.
*/
func (this *QPageSize) Swap(other QPageSize_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPageSize_PTR() != nil {
		convArg0 = other.QPageSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:249
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEquivalentTo(const QPageSize &) const

/*
Returns true if this page is equivalent to the other page, i.e. if the page has the same size regardless of other attributes like name.
*/
func (this *QPageSize) IsEquivalentTo(other QPageSize_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPageSize_PTR() != nil {
		convArg0 = other.QPageSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize14isEquivalentToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagesize.h:251
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if this page size is valid.

The page size may be invalid if created with an invalid PageSizeId, or a negative or invalid QSize or QSizeF, or the null constructor.
*/
func (this *QPageSize) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagesize.h:253
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key() const

/*
Returns the unique key of the page size.

By default this is the PPD standard mediaOption keyword for the page size, or the PPD custom format key. If the QPageSize instance was obtained from a print device then this will be the key provided by the print device and may differ from the standard key.

If the QPageSize is invalid then the key will be an empty string.

This key should never be shown to end users, it is an internal key only. For a human-readable name use name().

See also name().
*/
func (this *QPageSize) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpagesize.h:271
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString key(QPageSize::PageSizeId)

/*
Returns the unique key of the page size.

By default this is the PPD standard mediaOption keyword for the page size, or the PPD custom format key. If the QPageSize instance was obtained from a print device then this will be the key provided by the print device and may differ from the standard key.

If the QPageSize is invalid then the key will be an empty string.

This key should never be shown to end users, it is an internal key only. For a human-readable name use name().

See also name().
*/
func (this *QPageSize) Key1(pageSizeId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize3keyENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QPageSize_Key1(pageSizeId int) string {
	var nilthis *QPageSize
	rv := nilthis.Key1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:254
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*
Returns a localized human-readable name for the page size.

If the QPageSize instance was obtained from a print device then the name used is that provided by the print device. Note that a print device may not support the current default locale language.

If the QPageSize is invalid then the name will be an empty string.
*/
func (this *QPageSize) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpagesize.h:272
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString name(QPageSize::PageSizeId)

/*
Returns a localized human-readable name for the page size.

If the QPageSize instance was obtained from a print device then the name used is that provided by the print device. Note that a print device may not support the current default locale language.

If the QPageSize is invalid then the name will be an empty string.
*/
func (this *QPageSize) Name1(pageSizeId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize4nameENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QPageSize_Name1(pageSizeId int) string {
	var nilthis *QPageSize
	rv := nilthis.Name1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:256
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id() const

/*
Returns the standard QPageSize::PageSizeId of the page, or QPageSize::Custom.

If the QPageSize is invalid then the ID will be QPageSize::Custom.
*/
func (this *QPageSize) Id() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize2idEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:274
// index:1
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(const QSize &, QPageSize::SizeMatchPolicy)

/*
Returns the standard QPageSize::PageSizeId of the page, or QPageSize::Custom.

If the QPageSize is invalid then the ID will be QPageSize::Custom.
*/
func (this *QPageSize) Id1(pointSize qtcore.QSize_ITF, matchPolicy int) int {
	var convArg0 unsafe.Pointer
	if pointSize != nil && pointSize.QSize_PTR() != nil {
		convArg0 = pointSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idERK5QSizeNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, matchPolicy)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_Id1(pointSize qtcore.QSize_ITF, matchPolicy int) int {
	var nilthis *QPageSize
	rv := nilthis.Id1(pointSize, matchPolicy)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:274
// index:1
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(const QSize &, QPageSize::SizeMatchPolicy)

/*
Returns the standard QPageSize::PageSizeId of the page, or QPageSize::Custom.

If the QPageSize is invalid then the ID will be QPageSize::Custom.
*/
func (this *QPageSize) Id1p(pointSize qtcore.QSize_ITF) int {
	var convArg0 unsafe.Pointer
	if pointSize != nil && pointSize.QSize_PTR() != nil {
		convArg0 = pointSize.QSize_PTR().GetCthis()
	}
	// arg: 1, QPageSize::SizeMatchPolicy=Enum, QPageSize::SizeMatchPolicy=Enum, , Invalid
	matchPolicy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idERK5QSizeNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, matchPolicy)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:276
// index:2
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(const QSizeF &, QPageSize::Unit, QPageSize::SizeMatchPolicy)

/*
Returns the standard QPageSize::PageSizeId of the page, or QPageSize::Custom.

If the QPageSize is invalid then the ID will be QPageSize::Custom.
*/
func (this *QPageSize) Id2(size qtcore.QSizeF_ITF, units int, matchPolicy int) int {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idERK6QSizeFNS_4UnitENS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, units, matchPolicy)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_Id2(size qtcore.QSizeF_ITF, units int, matchPolicy int) int {
	var nilthis *QPageSize
	rv := nilthis.Id2(size, units, matchPolicy)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:276
// index:2
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(const QSizeF &, QPageSize::Unit, QPageSize::SizeMatchPolicy)

/*
Returns the standard QPageSize::PageSizeId of the page, or QPageSize::Custom.

If the QPageSize is invalid then the ID will be QPageSize::Custom.
*/
func (this *QPageSize) Id2p(size qtcore.QSizeF_ITF, units int) int {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	// arg: 2, QPageSize::SizeMatchPolicy=Enum, QPageSize::SizeMatchPolicy=Enum, , Invalid
	matchPolicy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idERK6QSizeFNS_4UnitENS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, units, matchPolicy)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:279
// index:3
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(int)

/*
Returns the standard QPageSize::PageSizeId of the page, or QPageSize::Custom.

If the QPageSize is invalid then the ID will be QPageSize::Custom.
*/
func (this *QPageSize) Id3(windowsId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idEi", qtrt.FFI_TYPE_POINTER, windowsId)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_Id3(windowsId int) int {
	var nilthis *QPageSize
	rv := nilthis.Id3(windowsId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:258
// index:0
// Public Visibility=Default Availability=Available
// [4] int windowsId() const

/*
Returns the Windows DMPAPER enum value for the page size.

Not all valid PPD page sizes have a Windows equivalent, in which case 0 will be returned.

If the QPageSize is invalid then the Windows ID will be 0.

See also id().
*/
func (this *QPageSize) WindowsId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize9windowsIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpagesize.h:280
// index:1
// Public static Visibility=Default Availability=Available
// [4] int windowsId(QPageSize::PageSizeId)

/*
Returns the Windows DMPAPER enum value for the page size.

Not all valid PPD page sizes have a Windows equivalent, in which case 0 will be returned.

If the QPageSize is invalid then the Windows ID will be 0.

See also id().
*/
func (this *QPageSize) WindowsId1(pageSizeId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize9windowsIdENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QPageSize_WindowsId1(pageSizeId int) int {
	var nilthis *QPageSize
	rv := nilthis.WindowsId1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:260
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF definitionSize() const

/*
Returns the definition size of the page size.

For a standard page size this will be the size as defined in the relevant standard, i.e. ISO A4 will be defined in millimeters while ANSI Letter will be defined in inches.

For a custom page size this will be the original size used to create the page size object.

If the QPageSize is invalid then the QSizeF will be invalid.

See also definitionUnits().
*/
func (this *QPageSize) DefinitionSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize14definitionSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:282
// index:1
// Public static Visibility=Default Availability=Available
// [16] QSizeF definitionSize(QPageSize::PageSizeId)

/*
Returns the definition size of the page size.

For a standard page size this will be the size as defined in the relevant standard, i.e. ISO A4 will be defined in millimeters while ANSI Letter will be defined in inches.

For a custom page size this will be the original size used to create the page size object.

If the QPageSize is invalid then the QSizeF will be invalid.

See also definitionUnits().
*/
func (this *QPageSize) DefinitionSize1(pageSizeId int) *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize14definitionSizeENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}
func QPageSize_DefinitionSize1(pageSizeId int) *qtcore.QSizeF /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.DefinitionSize1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:261
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageSize::Unit definitionUnits() const

/*
Returns the definition units of the page size.

For a standard page size this will be the units as defined in the relevant standard, i.e. ISO A4 will be defined in millimeters while ANSI Letter will be defined in inches.

For a custom page size this will be the original units used to create the page size object.

If the QPageSize is invalid then the QPageSize::Unit will be invalid.

See also definitionSize().
*/
func (this *QPageSize) DefinitionUnits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize15definitionUnitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:283
// index:1
// Public static Visibility=Default Availability=Available
// [4] QPageSize::Unit definitionUnits(QPageSize::PageSizeId)

/*
Returns the definition units of the page size.

For a standard page size this will be the units as defined in the relevant standard, i.e. ISO A4 will be defined in millimeters while ANSI Letter will be defined in inches.

For a custom page size this will be the original units used to create the page size object.

If the QPageSize is invalid then the QPageSize::Unit will be invalid.

See also definitionSize().
*/
func (this *QPageSize) DefinitionUnits1(pageSizeId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize15definitionUnitsENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_DefinitionUnits1(pageSizeId int) int {
	var nilthis *QPageSize
	rv := nilthis.DefinitionUnits1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:263
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(QPageSize::Unit) const

/*
Returns the size of the page in the required units.

If the QPageSize is invalid then the QSizeF will be invalid.
*/
func (this *QPageSize) Size(units int) *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize4sizeENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:285
// index:1
// Public static Visibility=Default Availability=Available
// [16] QSizeF size(QPageSize::PageSizeId, QPageSize::Unit)

/*
Returns the size of the page in the required units.

If the QPageSize is invalid then the QSizeF will be invalid.
*/
func (this *QPageSize) Size1(pageSizeId int, units int) *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize4sizeENS_10PageSizeIdENS_4UnitE", qtrt.FFI_TYPE_POINTER, pageSizeId, units)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}
func QPageSize_Size1(pageSizeId int, units int) *qtcore.QSizeF /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.Size1(pageSizeId, units)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:264
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizePoints() const

/*
Returns the size of the page in Postscript Points (1/72 of an inch).

If the QPageSize is invalid then the QSize will be invalid.
*/
func (this *QPageSize) SizePoints() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10sizePointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:286
// index:1
// Public static Visibility=Default Availability=Available
// [8] QSize sizePoints(QPageSize::PageSizeId)

/*
Returns the size of the page in Postscript Points (1/72 of an inch).

If the QPageSize is invalid then the QSize will be invalid.
*/
func (this *QPageSize) SizePoints1(pageSizeId int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize10sizePointsENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}
func QPageSize_SizePoints1(pageSizeId int) *qtcore.QSize /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.SizePoints1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:265
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizePixels(int) const

/*
Returns the size of the page in Device Pixels at the given resolution.

If the QPageSize is invalid then the QSize will be invalid.
*/
func (this *QPageSize) SizePixels(resolution int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10sizePixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:287
// index:1
// Public static Visibility=Default Availability=Available
// [8] QSize sizePixels(QPageSize::PageSizeId, int)

/*
Returns the size of the page in Device Pixels at the given resolution.

If the QPageSize is invalid then the QSize will be invalid.
*/
func (this *QPageSize) SizePixels1(pageSizeId int, resolution int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize10sizePixelsENS_10PageSizeIdEi", qtrt.FFI_TYPE_POINTER, pageSizeId, resolution)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}
func QPageSize_SizePixels1(pageSizeId int, resolution int) *qtcore.QSize /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.SizePixels1(pageSizeId, resolution)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:267
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect(QPageSize::Unit) const

/*
Returns the page rectangle in the required units.

If the QPageSize is invalid then the QRect will be invalid.
*/
func (this *QPageSize) Rect(units int) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize4rectENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:268
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rectPoints() const

/*
Returns the page rectangle in Postscript Points (1/72 of an inch).

If the QPageSize is invalid then the QRect will be invalid.
*/
func (this *QPageSize) RectPoints() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10rectPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:269
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rectPixels(int) const

/*
Returns the page rectangle in Device Pixels at the given resolution.

If the QPageSize is invalid then the QRect will be invalid.
*/
func (this *QPageSize) RectPixels(resolution int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10rectPixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

/*
This enum type lists the available page sizes as defined in the Postscript PPD standard. These values are duplicated in QPagedPaintDevice and QPrinter.

The defined sizes are:

QPageSize::AnsiALetter= Letter
QPageSize::AnsiBLedger= Ledger
QPageSize::EnvelopeDLDLE= DLE


Due to historic reasons QPageSize::Executive is not the same as the standard Postscript and Windows Executive size, use QPageSize::ExecutiveStandard instead.

The Postscript standard size QPageSize::Folio is different to the Windows DMPAPER_FOLIO size, use the Postscript standard size QPageSize::FanFoldGermanLegal if needed.

*/
type QPageSize__PageSizeId = int

//
const QPageSize__A4 QPageSize__PageSizeId = 0

//
const QPageSize__B5 QPageSize__PageSizeId = 1

//
const QPageSize__Letter QPageSize__PageSizeId = 2

//
const QPageSize__Legal QPageSize__PageSizeId = 3

//
const QPageSize__Executive QPageSize__PageSizeId = 4

//
const QPageSize__A0 QPageSize__PageSizeId = 5

//
const QPageSize__A1 QPageSize__PageSizeId = 6

//
const QPageSize__A2 QPageSize__PageSizeId = 7

//
const QPageSize__A3 QPageSize__PageSizeId = 8

//
const QPageSize__A5 QPageSize__PageSizeId = 9

//
const QPageSize__A6 QPageSize__PageSizeId = 10

//
const QPageSize__A7 QPageSize__PageSizeId = 11

//
const QPageSize__A8 QPageSize__PageSizeId = 12

//
const QPageSize__A9 QPageSize__PageSizeId = 13

//
const QPageSize__B0 QPageSize__PageSizeId = 14

//
const QPageSize__B1 QPageSize__PageSizeId = 15

//
const QPageSize__B10 QPageSize__PageSizeId = 16

//
const QPageSize__B2 QPageSize__PageSizeId = 17

//
const QPageSize__B3 QPageSize__PageSizeId = 18

//
const QPageSize__B4 QPageSize__PageSizeId = 19

//
const QPageSize__B6 QPageSize__PageSizeId = 20

//
const QPageSize__B7 QPageSize__PageSizeId = 21

//
const QPageSize__B8 QPageSize__PageSizeId = 22

//
const QPageSize__B9 QPageSize__PageSizeId = 23

//
const QPageSize__C5E QPageSize__PageSizeId = 24

//
const QPageSize__Comm10E QPageSize__PageSizeId = 25

//
const QPageSize__DLE QPageSize__PageSizeId = 26

//
const QPageSize__Folio QPageSize__PageSizeId = 27

//
const QPageSize__Ledger QPageSize__PageSizeId = 28

//
const QPageSize__Tabloid QPageSize__PageSizeId = 29

//
const QPageSize__Custom QPageSize__PageSizeId = 30

//
const QPageSize__A10 QPageSize__PageSizeId = 31

//
const QPageSize__A3Extra QPageSize__PageSizeId = 32

//
const QPageSize__A4Extra QPageSize__PageSizeId = 33

//
const QPageSize__A4Plus QPageSize__PageSizeId = 34

//
const QPageSize__A4Small QPageSize__PageSizeId = 35

//
const QPageSize__A5Extra QPageSize__PageSizeId = 36

//
const QPageSize__B5Extra QPageSize__PageSizeId = 37

//
const QPageSize__JisB0 QPageSize__PageSizeId = 38

//
const QPageSize__JisB1 QPageSize__PageSizeId = 39

//
const QPageSize__JisB2 QPageSize__PageSizeId = 40

//
const QPageSize__JisB3 QPageSize__PageSizeId = 41

//
const QPageSize__JisB4 QPageSize__PageSizeId = 42

//
const QPageSize__JisB5 QPageSize__PageSizeId = 43

//
const QPageSize__JisB6 QPageSize__PageSizeId = 44

//
const QPageSize__JisB7 QPageSize__PageSizeId = 45

//
const QPageSize__JisB8 QPageSize__PageSizeId = 46

//
const QPageSize__JisB9 QPageSize__PageSizeId = 47

//
const QPageSize__JisB10 QPageSize__PageSizeId = 48

//
const QPageSize__AnsiC QPageSize__PageSizeId = 49

//
const QPageSize__AnsiD QPageSize__PageSizeId = 50

//
const QPageSize__AnsiE QPageSize__PageSizeId = 51

//
const QPageSize__LegalExtra QPageSize__PageSizeId = 52

//
const QPageSize__LetterExtra QPageSize__PageSizeId = 53

//
const QPageSize__LetterPlus QPageSize__PageSizeId = 54

//
const QPageSize__LetterSmall QPageSize__PageSizeId = 55

//
const QPageSize__TabloidExtra QPageSize__PageSizeId = 56

//
const QPageSize__ArchA QPageSize__PageSizeId = 57

//
const QPageSize__ArchB QPageSize__PageSizeId = 58

//
const QPageSize__ArchC QPageSize__PageSizeId = 59

//
const QPageSize__ArchD QPageSize__PageSizeId = 60

//
const QPageSize__ArchE QPageSize__PageSizeId = 61

//
const QPageSize__Imperial7x9 QPageSize__PageSizeId = 62

//
const QPageSize__Imperial8x10 QPageSize__PageSizeId = 63

//
const QPageSize__Imperial9x11 QPageSize__PageSizeId = 64

//
const QPageSize__Imperial9x12 QPageSize__PageSizeId = 65

//
const QPageSize__Imperial10x11 QPageSize__PageSizeId = 66

//
const QPageSize__Imperial10x13 QPageSize__PageSizeId = 67

//
const QPageSize__Imperial10x14 QPageSize__PageSizeId = 68

//
const QPageSize__Imperial12x11 QPageSize__PageSizeId = 69

//
const QPageSize__Imperial15x11 QPageSize__PageSizeId = 70

//
const QPageSize__ExecutiveStandard QPageSize__PageSizeId = 71

//
const QPageSize__Note QPageSize__PageSizeId = 72

//
const QPageSize__Quarto QPageSize__PageSizeId = 73

//
const QPageSize__Statement QPageSize__PageSizeId = 74

//
const QPageSize__SuperA QPageSize__PageSizeId = 75

//
const QPageSize__SuperB QPageSize__PageSizeId = 76

//
const QPageSize__Postcard QPageSize__PageSizeId = 77

//
const QPageSize__DoublePostcard QPageSize__PageSizeId = 78

//
const QPageSize__Prc16K QPageSize__PageSizeId = 79

//
const QPageSize__Prc32K QPageSize__PageSizeId = 80

//
const QPageSize__Prc32KBig QPageSize__PageSizeId = 81

//
const QPageSize__FanFoldUS QPageSize__PageSizeId = 82

//
const QPageSize__FanFoldGerman QPageSize__PageSizeId = 83

//
const QPageSize__FanFoldGermanLegal QPageSize__PageSizeId = 84

//
const QPageSize__EnvelopeB4 QPageSize__PageSizeId = 85

//
const QPageSize__EnvelopeB5 QPageSize__PageSizeId = 86

//
const QPageSize__EnvelopeB6 QPageSize__PageSizeId = 87

//
const QPageSize__EnvelopeC0 QPageSize__PageSizeId = 88

//
const QPageSize__EnvelopeC1 QPageSize__PageSizeId = 89

//
const QPageSize__EnvelopeC2 QPageSize__PageSizeId = 90

//
const QPageSize__EnvelopeC3 QPageSize__PageSizeId = 91

//
const QPageSize__EnvelopeC4 QPageSize__PageSizeId = 92

//
const QPageSize__EnvelopeC6 QPageSize__PageSizeId = 93

//
const QPageSize__EnvelopeC65 QPageSize__PageSizeId = 94

//
const QPageSize__EnvelopeC7 QPageSize__PageSizeId = 95

//
const QPageSize__Envelope9 QPageSize__PageSizeId = 96

//
const QPageSize__Envelope11 QPageSize__PageSizeId = 97

//
const QPageSize__Envelope12 QPageSize__PageSizeId = 98

//
const QPageSize__Envelope14 QPageSize__PageSizeId = 99

//
const QPageSize__EnvelopeMonarch QPageSize__PageSizeId = 100

//
const QPageSize__EnvelopePersonal QPageSize__PageSizeId = 101

//
const QPageSize__EnvelopeChou3 QPageSize__PageSizeId = 102

//
const QPageSize__EnvelopeChou4 QPageSize__PageSizeId = 103

//
const QPageSize__EnvelopeInvite QPageSize__PageSizeId = 104

//
const QPageSize__EnvelopeItalian QPageSize__PageSizeId = 105

//
const QPageSize__EnvelopeKaku2 QPageSize__PageSizeId = 106

//
const QPageSize__EnvelopeKaku3 QPageSize__PageSizeId = 107

//
const QPageSize__EnvelopePrc1 QPageSize__PageSizeId = 108

//
const QPageSize__EnvelopePrc2 QPageSize__PageSizeId = 109

//
const QPageSize__EnvelopePrc3 QPageSize__PageSizeId = 110

//
const QPageSize__EnvelopePrc4 QPageSize__PageSizeId = 111

//
const QPageSize__EnvelopePrc5 QPageSize__PageSizeId = 112

//
const QPageSize__EnvelopePrc6 QPageSize__PageSizeId = 113

//
const QPageSize__EnvelopePrc7 QPageSize__PageSizeId = 114

//
const QPageSize__EnvelopePrc8 QPageSize__PageSizeId = 115

//
const QPageSize__EnvelopePrc9 QPageSize__PageSizeId = 116

//
const QPageSize__EnvelopePrc10 QPageSize__PageSizeId = 117

//
const QPageSize__EnvelopeYou4 QPageSize__PageSizeId = 118

//
const QPageSize__LastPageSize QPageSize__PageSizeId = 118

//
const QPageSize__NPageSize QPageSize__PageSizeId = 118

//
const QPageSize__NPaperSize QPageSize__PageSizeId = 118

//
const QPageSize__AnsiA QPageSize__PageSizeId = 2

//
const QPageSize__AnsiB QPageSize__PageSizeId = 28

//
const QPageSize__EnvelopeC5 QPageSize__PageSizeId = 24

//
const QPageSize__EnvelopeDL QPageSize__PageSizeId = 26

//
const QPageSize__Envelope10 QPageSize__PageSizeId = 25

func (this *QPageSize) PageSizeIdItemName(val int) string {
	switch val {
	case QPageSize__A4: // 0
		return "A4"
	case QPageSize__B5: // 1
		return "B5"
	case QPageSize__Letter: // 2
		return "Letter,AnsiA"
	case QPageSize__Legal: // 3
		return "Legal"
	case QPageSize__Executive: // 4
		return "Executive"
	case QPageSize__A0: // 5
		return "A0"
	case QPageSize__A1: // 6
		return "A1"
	case QPageSize__A2: // 7
		return "A2"
	case QPageSize__A3: // 8
		return "A3"
	case QPageSize__A5: // 9
		return "A5"
	case QPageSize__A6: // 10
		return "A6"
	case QPageSize__A7: // 11
		return "A7"
	case QPageSize__A8: // 12
		return "A8"
	case QPageSize__A9: // 13
		return "A9"
	case QPageSize__B0: // 14
		return "B0"
	case QPageSize__B1: // 15
		return "B1"
	case QPageSize__B10: // 16
		return "B10"
	case QPageSize__B2: // 17
		return "B2"
	case QPageSize__B3: // 18
		return "B3"
	case QPageSize__B4: // 19
		return "B4"
	case QPageSize__B6: // 20
		return "B6"
	case QPageSize__B7: // 21
		return "B7"
	case QPageSize__B8: // 22
		return "B8"
	case QPageSize__B9: // 23
		return "B9"
	case QPageSize__C5E: // 24
		return "C5E,EnvelopeC5"
	case QPageSize__Comm10E: // 25
		return "Comm10E,Envelope10"
	case QPageSize__DLE: // 26
		return "DLE,EnvelopeDL"
	case QPageSize__Folio: // 27
		return "Folio"
	case QPageSize__Ledger: // 28
		return "Ledger,AnsiB"
	case QPageSize__Tabloid: // 29
		return "Tabloid"
	case QPageSize__Custom: // 30
		return "Custom"
	case QPageSize__A10: // 31
		return "A10"
	case QPageSize__A3Extra: // 32
		return "A3Extra"
	case QPageSize__A4Extra: // 33
		return "A4Extra"
	case QPageSize__A4Plus: // 34
		return "A4Plus"
	case QPageSize__A4Small: // 35
		return "A4Small"
	case QPageSize__A5Extra: // 36
		return "A5Extra"
	case QPageSize__B5Extra: // 37
		return "B5Extra"
	case QPageSize__JisB0: // 38
		return "JisB0"
	case QPageSize__JisB1: // 39
		return "JisB1"
	case QPageSize__JisB2: // 40
		return "JisB2"
	case QPageSize__JisB3: // 41
		return "JisB3"
	case QPageSize__JisB4: // 42
		return "JisB4"
	case QPageSize__JisB5: // 43
		return "JisB5"
	case QPageSize__JisB6: // 44
		return "JisB6"
	case QPageSize__JisB7: // 45
		return "JisB7"
	case QPageSize__JisB8: // 46
		return "JisB8"
	case QPageSize__JisB9: // 47
		return "JisB9"
	case QPageSize__JisB10: // 48
		return "JisB10"
	case QPageSize__AnsiC: // 49
		return "AnsiC"
	case QPageSize__AnsiD: // 50
		return "AnsiD"
	case QPageSize__AnsiE: // 51
		return "AnsiE"
	case QPageSize__LegalExtra: // 52
		return "LegalExtra"
	case QPageSize__LetterExtra: // 53
		return "LetterExtra"
	case QPageSize__LetterPlus: // 54
		return "LetterPlus"
	case QPageSize__LetterSmall: // 55
		return "LetterSmall"
	case QPageSize__TabloidExtra: // 56
		return "TabloidExtra"
	case QPageSize__ArchA: // 57
		return "ArchA"
	case QPageSize__ArchB: // 58
		return "ArchB"
	case QPageSize__ArchC: // 59
		return "ArchC"
	case QPageSize__ArchD: // 60
		return "ArchD"
	case QPageSize__ArchE: // 61
		return "ArchE"
	case QPageSize__Imperial7x9: // 62
		return "Imperial7x9"
	case QPageSize__Imperial8x10: // 63
		return "Imperial8x10"
	case QPageSize__Imperial9x11: // 64
		return "Imperial9x11"
	case QPageSize__Imperial9x12: // 65
		return "Imperial9x12"
	case QPageSize__Imperial10x11: // 66
		return "Imperial10x11"
	case QPageSize__Imperial10x13: // 67
		return "Imperial10x13"
	case QPageSize__Imperial10x14: // 68
		return "Imperial10x14"
	case QPageSize__Imperial12x11: // 69
		return "Imperial12x11"
	case QPageSize__Imperial15x11: // 70
		return "Imperial15x11"
	case QPageSize__ExecutiveStandard: // 71
		return "ExecutiveStandard"
	case QPageSize__Note: // 72
		return "Note"
	case QPageSize__Quarto: // 73
		return "Quarto"
	case QPageSize__Statement: // 74
		return "Statement"
	case QPageSize__SuperA: // 75
		return "SuperA"
	case QPageSize__SuperB: // 76
		return "SuperB"
	case QPageSize__Postcard: // 77
		return "Postcard"
	case QPageSize__DoublePostcard: // 78
		return "DoublePostcard"
	case QPageSize__Prc16K: // 79
		return "Prc16K"
	case QPageSize__Prc32K: // 80
		return "Prc32K"
	case QPageSize__Prc32KBig: // 81
		return "Prc32KBig"
	case QPageSize__FanFoldUS: // 82
		return "FanFoldUS"
	case QPageSize__FanFoldGerman: // 83
		return "FanFoldGerman"
	case QPageSize__FanFoldGermanLegal: // 84
		return "FanFoldGermanLegal"
	case QPageSize__EnvelopeB4: // 85
		return "EnvelopeB4"
	case QPageSize__EnvelopeB5: // 86
		return "EnvelopeB5"
	case QPageSize__EnvelopeB6: // 87
		return "EnvelopeB6"
	case QPageSize__EnvelopeC0: // 88
		return "EnvelopeC0"
	case QPageSize__EnvelopeC1: // 89
		return "EnvelopeC1"
	case QPageSize__EnvelopeC2: // 90
		return "EnvelopeC2"
	case QPageSize__EnvelopeC3: // 91
		return "EnvelopeC3"
	case QPageSize__EnvelopeC4: // 92
		return "EnvelopeC4"
	case QPageSize__EnvelopeC6: // 93
		return "EnvelopeC6"
	case QPageSize__EnvelopeC65: // 94
		return "EnvelopeC65"
	case QPageSize__EnvelopeC7: // 95
		return "EnvelopeC7"
	case QPageSize__Envelope9: // 96
		return "Envelope9"
	case QPageSize__Envelope11: // 97
		return "Envelope11"
	case QPageSize__Envelope12: // 98
		return "Envelope12"
	case QPageSize__Envelope14: // 99
		return "Envelope14"
	case QPageSize__EnvelopeMonarch: // 100
		return "EnvelopeMonarch"
	case QPageSize__EnvelopePersonal: // 101
		return "EnvelopePersonal"
	case QPageSize__EnvelopeChou3: // 102
		return "EnvelopeChou3"
	case QPageSize__EnvelopeChou4: // 103
		return "EnvelopeChou4"
	case QPageSize__EnvelopeInvite: // 104
		return "EnvelopeInvite"
	case QPageSize__EnvelopeItalian: // 105
		return "EnvelopeItalian"
	case QPageSize__EnvelopeKaku2: // 106
		return "EnvelopeKaku2"
	case QPageSize__EnvelopeKaku3: // 107
		return "EnvelopeKaku3"
	case QPageSize__EnvelopePrc1: // 108
		return "EnvelopePrc1"
	case QPageSize__EnvelopePrc2: // 109
		return "EnvelopePrc2"
	case QPageSize__EnvelopePrc3: // 110
		return "EnvelopePrc3"
	case QPageSize__EnvelopePrc4: // 111
		return "EnvelopePrc4"
	case QPageSize__EnvelopePrc5: // 112
		return "EnvelopePrc5"
	case QPageSize__EnvelopePrc6: // 113
		return "EnvelopePrc6"
	case QPageSize__EnvelopePrc7: // 114
		return "EnvelopePrc7"
	case QPageSize__EnvelopePrc8: // 115
		return "EnvelopePrc8"
	case QPageSize__EnvelopePrc9: // 116
		return "EnvelopePrc9"
	case QPageSize__EnvelopePrc10: // 117
		return "EnvelopePrc10"
	case QPageSize__EnvelopeYou4: // 118
		return "EnvelopeYou4,LastPageSize,NPageSize,NPaperSize"
		// case QPageSize__LastPageSize: // 118
		// return ""
		// case QPageSize__NPageSize: // 118
		// return ""
		// case QPageSize__NPaperSize: // 118
		// return ""
		// case QPageSize__AnsiA: // 2
		// return ""
		// case QPageSize__AnsiB: // 28
		// return ""
		// case QPageSize__EnvelopeC5: // 24
		// return ""
		// case QPageSize__EnvelopeDL: // 26
		// return ""
		// case QPageSize__Envelope10: // 25
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPageSize_PageSizeIdItemName(val int) string {
	var nilthis *QPageSize
	return nilthis.PageSizeIdItemName(val)
}

/*
This enum type is used to specify the measurement unit for page sizes.


*/
type QPageSize__Unit = int

//
const QPageSize__Millimeter QPageSize__Unit = 0

//
const QPageSize__Point QPageSize__Unit = 1

//
const QPageSize__Inch QPageSize__Unit = 2

//
const QPageSize__Pica QPageSize__Unit = 3

//
const QPageSize__Didot QPageSize__Unit = 4

//
const QPageSize__Cicero QPageSize__Unit = 5

func (this *QPageSize) UnitItemName(val int) string {
	switch val {
	case QPageSize__Millimeter: // 0
		return "Millimeter"
	case QPageSize__Point: // 1
		return "Point"
	case QPageSize__Inch: // 2
		return "Inch"
	case QPageSize__Pica: // 3
		return "Pica"
	case QPageSize__Didot: // 4
		return "Didot"
	case QPageSize__Cicero: // 5
		return "Cicero"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPageSize_UnitItemName(val int) string {
	var nilthis *QPageSize
	return nilthis.UnitItemName(val)
}

/*

 */
type QPageSize__SizeMatchPolicy = int

// Match to a standard page size if within the margin of tolerance.
const QPageSize__FuzzyMatch QPageSize__SizeMatchPolicy = 0

// Match to a standard page size if within the margin of tolerance regardless of orientation.
const QPageSize__FuzzyOrientationMatch QPageSize__SizeMatchPolicy = 1

// Only match to a standard page size if the sizes match exactly.
const QPageSize__ExactMatch QPageSize__SizeMatchPolicy = 2

func (this *QPageSize) SizeMatchPolicyItemName(val int) string {
	switch val {
	case QPageSize__FuzzyMatch: // 0
		return "FuzzyMatch"
	case QPageSize__FuzzyOrientationMatch: // 1
		return "FuzzyOrientationMatch"
	case QPageSize__ExactMatch: // 2
		return "ExactMatch"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPageSize_SizeMatchPolicyItemName(val int) string {
	var nilthis *QPageSize
	return nilthis.SizeMatchPolicyItemName(val)
}

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
