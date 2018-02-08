package qtgui

// /usr/include/qt/QtGui/qpagesize.h
// #include <qpagesize.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QPageSize struct {
	*qtrt.CObject
}

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
func NewQPageSize() *QPageSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:231
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(enum QPageSize::PageSizeId)
func NewQPageSize_1(pageSizeId int) *QPageSize {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:232
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSize &, const QString &, enum QPageSize::SizeMatchPolicy)
func NewQPageSize_2(pointSize *qtcore.QSize, name string, matchPolicy int) *QPageSize {
	var convArg0 = pointSize.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK5QSizeRK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, matchPolicy)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:235
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QPageSize(const QSizeF &, enum QPageSize::Unit, const QString &, enum QPageSize::SizeMatchPolicy)
func NewQPageSize_3(size *qtcore.QSizeF, units int, name string, matchPolicy int) *QPageSize {
	var convArg0 = size.GetCthis()
	var tmpArg2 = qtcore.NewQString_5(name)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeC2ERK6QSizeFNS_4UnitERK7QStringNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, units, convArg2, matchPolicy)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPageSize)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPageSize()
func DeleteQPageSize(this *QPageSize) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSizeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpagesize.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPageSize &)
func (this *QPageSize) Swap(other *QPageSize) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:249
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEquivalentTo(const QPageSize &)
func (this *QPageSize) IsEquivalentTo(other *QPageSize) bool {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize14isEquivalentToERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagesize.h:251
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QPageSize) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagesize.h:253
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key()
func (this *QPageSize) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpagesize.h:271
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString key(enum QPageSize::PageSizeId)
func (this *QPageSize) Key_1(pageSizeId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize3keyENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QPageSize_Key_1(pageSizeId int) string {
	var nilthis *QPageSize
	rv := nilthis.Key_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:254
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QPageSize) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpagesize.h:272
// index:1
// Public static Visibility=Default Availability=Available
// [8] QString name(enum QPageSize::PageSizeId)
func (this *QPageSize) Name_1(pageSizeId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize4nameENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QPageSize_Name_1(pageSizeId int) string {
	var nilthis *QPageSize
	rv := nilthis.Name_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:256
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id()
func (this *QPageSize) Id() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize2idEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:274
// index:1
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(const QSize &, enum QPageSize::SizeMatchPolicy)
func (this *QPageSize) Id_1(pointSize *qtcore.QSize, matchPolicy int) int {
	var convArg0 = pointSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idERK5QSizeNS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, matchPolicy)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_Id_1(pointSize *qtcore.QSize, matchPolicy int) int {
	var nilthis *QPageSize
	rv := nilthis.Id_1(pointSize, matchPolicy)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:276
// index:2
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(const QSizeF &, enum QPageSize::Unit, enum QPageSize::SizeMatchPolicy)
func (this *QPageSize) Id_2(size *qtcore.QSizeF, units int, matchPolicy int) int {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idERK6QSizeFNS_4UnitENS_15SizeMatchPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, units, matchPolicy)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_Id_2(size *qtcore.QSizeF, units int, matchPolicy int) int {
	var nilthis *QPageSize
	rv := nilthis.Id_2(size, units, matchPolicy)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:279
// index:3
// Public static Visibility=Default Availability=Available
// [4] QPageSize::PageSizeId id(int)
func (this *QPageSize) Id_3(windowsId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize2idEi", qtrt.FFI_TYPE_POINTER, windowsId)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_Id_3(windowsId int) int {
	var nilthis *QPageSize
	rv := nilthis.Id_3(windowsId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:258
// index:0
// Public Visibility=Default Availability=Available
// [4] int windowsId()
func (this *QPageSize) WindowsId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize9windowsIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpagesize.h:280
// index:1
// Public static Visibility=Default Availability=Available
// [4] int windowsId(enum QPageSize::PageSizeId)
func (this *QPageSize) WindowsId_1(pageSizeId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize9windowsIdENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QPageSize_WindowsId_1(pageSizeId int) int {
	var nilthis *QPageSize
	rv := nilthis.WindowsId_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:260
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF definitionSize()
func (this *QPageSize) DefinitionSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize14definitionSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:282
// index:1
// Public static Visibility=Default Availability=Available
// [16] QSizeF definitionSize(enum QPageSize::PageSizeId)
func (this *QPageSize) DefinitionSize_1(pageSizeId int) *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize14definitionSizeENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}
func QPageSize_DefinitionSize_1(pageSizeId int) *qtcore.QSizeF /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.DefinitionSize_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:261
// index:0
// Public Visibility=Default Availability=Available
// [4] QPageSize::Unit definitionUnits()
func (this *QPageSize) DefinitionUnits() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize15definitionUnitsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:283
// index:1
// Public static Visibility=Default Availability=Available
// [4] QPageSize::Unit definitionUnits(enum QPageSize::PageSizeId)
func (this *QPageSize) DefinitionUnits_1(pageSizeId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize15definitionUnitsENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	return int(rv)
}
func QPageSize_DefinitionUnits_1(pageSizeId int) int {
	var nilthis *QPageSize
	rv := nilthis.DefinitionUnits_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:263
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(enum QPageSize::Unit)
func (this *QPageSize) Size(units int) *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize4sizeENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:285
// index:1
// Public static Visibility=Default Availability=Available
// [16] QSizeF size(enum QPageSize::PageSizeId, enum QPageSize::Unit)
func (this *QPageSize) Size_1(pageSizeId int, units int) *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize4sizeENS_10PageSizeIdENS_4UnitE", qtrt.FFI_TYPE_POINTER, pageSizeId, units)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}
func QPageSize_Size_1(pageSizeId int, units int) *qtcore.QSizeF /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.Size_1(pageSizeId, units)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:264
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizePoints()
func (this *QPageSize) SizePoints() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10sizePointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:286
// index:1
// Public static Visibility=Default Availability=Available
// [8] QSize sizePoints(enum QPageSize::PageSizeId)
func (this *QPageSize) SizePoints_1(pageSizeId int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize10sizePointsENS_10PageSizeIdE", qtrt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}
func QPageSize_SizePoints_1(pageSizeId int) *qtcore.QSize /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.SizePoints_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:265
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizePixels(int)
func (this *QPageSize) SizePixels(resolution int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10sizePixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:287
// index:1
// Public static Visibility=Default Availability=Available
// [8] QSize sizePixels(enum QPageSize::PageSizeId, int)
func (this *QPageSize) SizePixels_1(pageSizeId int, resolution int) *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QPageSize10sizePixelsENS_10PageSizeIdEi", qtrt.FFI_TYPE_POINTER, pageSizeId, resolution)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}
func QPageSize_SizePixels_1(pageSizeId int, resolution int) *qtcore.QSize /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.SizePixels_1(pageSizeId, resolution)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:267
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect(enum QPageSize::Unit)
func (this *QPageSize) Rect(units int) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize4rectENS_4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), units)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:268
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rectPoints()
func (this *QPageSize) RectPoints() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10rectPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:269
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect rectPixels(int)
func (this *QPageSize) RectPixels(resolution int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QPageSize10rectPixelsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

type QPageSize__PageSizeId = int

const QPageSize__A4 QPageSize__PageSizeId = 0
const QPageSize__B5 QPageSize__PageSizeId = 1
const QPageSize__Letter QPageSize__PageSizeId = 2
const QPageSize__Legal QPageSize__PageSizeId = 3
const QPageSize__Executive QPageSize__PageSizeId = 4
const QPageSize__A0 QPageSize__PageSizeId = 5
const QPageSize__A1 QPageSize__PageSizeId = 6
const QPageSize__A2 QPageSize__PageSizeId = 7
const QPageSize__A3 QPageSize__PageSizeId = 8
const QPageSize__A5 QPageSize__PageSizeId = 9
const QPageSize__A6 QPageSize__PageSizeId = 10
const QPageSize__A7 QPageSize__PageSizeId = 11
const QPageSize__A8 QPageSize__PageSizeId = 12
const QPageSize__A9 QPageSize__PageSizeId = 13
const QPageSize__B0 QPageSize__PageSizeId = 14
const QPageSize__B1 QPageSize__PageSizeId = 15
const QPageSize__B10 QPageSize__PageSizeId = 16
const QPageSize__B2 QPageSize__PageSizeId = 17
const QPageSize__B3 QPageSize__PageSizeId = 18
const QPageSize__B4 QPageSize__PageSizeId = 19
const QPageSize__B6 QPageSize__PageSizeId = 20
const QPageSize__B7 QPageSize__PageSizeId = 21
const QPageSize__B8 QPageSize__PageSizeId = 22
const QPageSize__B9 QPageSize__PageSizeId = 23
const QPageSize__C5E QPageSize__PageSizeId = 24
const QPageSize__Comm10E QPageSize__PageSizeId = 25
const QPageSize__DLE QPageSize__PageSizeId = 26
const QPageSize__Folio QPageSize__PageSizeId = 27
const QPageSize__Ledger QPageSize__PageSizeId = 28
const QPageSize__Tabloid QPageSize__PageSizeId = 29
const QPageSize__Custom QPageSize__PageSizeId = 30
const QPageSize__A10 QPageSize__PageSizeId = 31
const QPageSize__A3Extra QPageSize__PageSizeId = 32
const QPageSize__A4Extra QPageSize__PageSizeId = 33
const QPageSize__A4Plus QPageSize__PageSizeId = 34
const QPageSize__A4Small QPageSize__PageSizeId = 35
const QPageSize__A5Extra QPageSize__PageSizeId = 36
const QPageSize__B5Extra QPageSize__PageSizeId = 37
const QPageSize__JisB0 QPageSize__PageSizeId = 38
const QPageSize__JisB1 QPageSize__PageSizeId = 39
const QPageSize__JisB2 QPageSize__PageSizeId = 40
const QPageSize__JisB3 QPageSize__PageSizeId = 41
const QPageSize__JisB4 QPageSize__PageSizeId = 42
const QPageSize__JisB5 QPageSize__PageSizeId = 43
const QPageSize__JisB6 QPageSize__PageSizeId = 44
const QPageSize__JisB7 QPageSize__PageSizeId = 45
const QPageSize__JisB8 QPageSize__PageSizeId = 46
const QPageSize__JisB9 QPageSize__PageSizeId = 47
const QPageSize__JisB10 QPageSize__PageSizeId = 48
const QPageSize__AnsiC QPageSize__PageSizeId = 49
const QPageSize__AnsiD QPageSize__PageSizeId = 50
const QPageSize__AnsiE QPageSize__PageSizeId = 51
const QPageSize__LegalExtra QPageSize__PageSizeId = 52
const QPageSize__LetterExtra QPageSize__PageSizeId = 53
const QPageSize__LetterPlus QPageSize__PageSizeId = 54
const QPageSize__LetterSmall QPageSize__PageSizeId = 55
const QPageSize__TabloidExtra QPageSize__PageSizeId = 56
const QPageSize__ArchA QPageSize__PageSizeId = 57
const QPageSize__ArchB QPageSize__PageSizeId = 58
const QPageSize__ArchC QPageSize__PageSizeId = 59
const QPageSize__ArchD QPageSize__PageSizeId = 60
const QPageSize__ArchE QPageSize__PageSizeId = 61
const QPageSize__Imperial7x9 QPageSize__PageSizeId = 62
const QPageSize__Imperial8x10 QPageSize__PageSizeId = 63
const QPageSize__Imperial9x11 QPageSize__PageSizeId = 64
const QPageSize__Imperial9x12 QPageSize__PageSizeId = 65
const QPageSize__Imperial10x11 QPageSize__PageSizeId = 66
const QPageSize__Imperial10x13 QPageSize__PageSizeId = 67
const QPageSize__Imperial10x14 QPageSize__PageSizeId = 68
const QPageSize__Imperial12x11 QPageSize__PageSizeId = 69
const QPageSize__Imperial15x11 QPageSize__PageSizeId = 70
const QPageSize__ExecutiveStandard QPageSize__PageSizeId = 71
const QPageSize__Note QPageSize__PageSizeId = 72
const QPageSize__Quarto QPageSize__PageSizeId = 73
const QPageSize__Statement QPageSize__PageSizeId = 74
const QPageSize__SuperA QPageSize__PageSizeId = 75
const QPageSize__SuperB QPageSize__PageSizeId = 76
const QPageSize__Postcard QPageSize__PageSizeId = 77
const QPageSize__DoublePostcard QPageSize__PageSizeId = 78
const QPageSize__Prc16K QPageSize__PageSizeId = 79
const QPageSize__Prc32K QPageSize__PageSizeId = 80
const QPageSize__Prc32KBig QPageSize__PageSizeId = 81
const QPageSize__FanFoldUS QPageSize__PageSizeId = 82
const QPageSize__FanFoldGerman QPageSize__PageSizeId = 83
const QPageSize__FanFoldGermanLegal QPageSize__PageSizeId = 84
const QPageSize__EnvelopeB4 QPageSize__PageSizeId = 85
const QPageSize__EnvelopeB5 QPageSize__PageSizeId = 86
const QPageSize__EnvelopeB6 QPageSize__PageSizeId = 87
const QPageSize__EnvelopeC0 QPageSize__PageSizeId = 88
const QPageSize__EnvelopeC1 QPageSize__PageSizeId = 89
const QPageSize__EnvelopeC2 QPageSize__PageSizeId = 90
const QPageSize__EnvelopeC3 QPageSize__PageSizeId = 91
const QPageSize__EnvelopeC4 QPageSize__PageSizeId = 92
const QPageSize__EnvelopeC6 QPageSize__PageSizeId = 93
const QPageSize__EnvelopeC65 QPageSize__PageSizeId = 94
const QPageSize__EnvelopeC7 QPageSize__PageSizeId = 95
const QPageSize__Envelope9 QPageSize__PageSizeId = 96
const QPageSize__Envelope11 QPageSize__PageSizeId = 97
const QPageSize__Envelope12 QPageSize__PageSizeId = 98
const QPageSize__Envelope14 QPageSize__PageSizeId = 99
const QPageSize__EnvelopeMonarch QPageSize__PageSizeId = 100
const QPageSize__EnvelopePersonal QPageSize__PageSizeId = 101
const QPageSize__EnvelopeChou3 QPageSize__PageSizeId = 102
const QPageSize__EnvelopeChou4 QPageSize__PageSizeId = 103
const QPageSize__EnvelopeInvite QPageSize__PageSizeId = 104
const QPageSize__EnvelopeItalian QPageSize__PageSizeId = 105
const QPageSize__EnvelopeKaku2 QPageSize__PageSizeId = 106
const QPageSize__EnvelopeKaku3 QPageSize__PageSizeId = 107
const QPageSize__EnvelopePrc1 QPageSize__PageSizeId = 108
const QPageSize__EnvelopePrc2 QPageSize__PageSizeId = 109
const QPageSize__EnvelopePrc3 QPageSize__PageSizeId = 110
const QPageSize__EnvelopePrc4 QPageSize__PageSizeId = 111
const QPageSize__EnvelopePrc5 QPageSize__PageSizeId = 112
const QPageSize__EnvelopePrc6 QPageSize__PageSizeId = 113
const QPageSize__EnvelopePrc7 QPageSize__PageSizeId = 114
const QPageSize__EnvelopePrc8 QPageSize__PageSizeId = 115
const QPageSize__EnvelopePrc9 QPageSize__PageSizeId = 116
const QPageSize__EnvelopePrc10 QPageSize__PageSizeId = 117
const QPageSize__EnvelopeYou4 QPageSize__PageSizeId = 118
const QPageSize__LastPageSize QPageSize__PageSizeId = 118
const QPageSize__NPageSize QPageSize__PageSizeId = 118
const QPageSize__NPaperSize QPageSize__PageSizeId = 118
const QPageSize__AnsiA QPageSize__PageSizeId = 2
const QPageSize__AnsiB QPageSize__PageSizeId = 28
const QPageSize__EnvelopeC5 QPageSize__PageSizeId = 24
const QPageSize__EnvelopeDL QPageSize__PageSizeId = 26
const QPageSize__Envelope10 QPageSize__PageSizeId = 25

type QPageSize__Unit = int

const QPageSize__Millimeter QPageSize__Unit = 0
const QPageSize__Point QPageSize__Unit = 1
const QPageSize__Inch QPageSize__Unit = 2
const QPageSize__Pica QPageSize__Unit = 3
const QPageSize__Didot QPageSize__Unit = 4
const QPageSize__Cicero QPageSize__Unit = 5

type QPageSize__SizeMatchPolicy = int

const QPageSize__FuzzyMatch QPageSize__SizeMatchPolicy = 0
const QPageSize__FuzzyOrientationMatch QPageSize__SizeMatchPolicy = 1
const QPageSize__ExactMatch QPageSize__SizeMatchPolicy = 2

//  body block end
