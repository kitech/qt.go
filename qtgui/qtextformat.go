package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

type QTextFormat struct {
	*qtrt.CObject
}
type QTextFormat_ITF interface {
	QTextFormat_PTR() *QTextFormat
}

func (ptr *QTextFormat) QTextFormat_PTR() *QTextFormat { return ptr }

func (this *QTextFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextFormat) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextFormatFromPointer(cthis unsafe.Pointer) *QTextFormat {
	return &QTextFormat{&qtrt.CObject{cthis}}
}
func (*QTextFormat) NewFromPointer(cthis unsafe.Pointer) *QTextFormat {
	return NewQTextFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextFormat()
func NewQTextFormat() *QTextFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:290
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextFormat(int)
func NewQTextFormat_1(type_ int) *QTextFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormatC2Ei", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:293
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFormat & operator=(const QTextFormat &)
func (this *QTextFormat) Operator_equal(rhs QTextFormat_ITF) *QTextFormat {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextFormat_PTR() != nil {
		convArg0 = rhs.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormataSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextFormat()
func DeleteQTextFormat(this *QTextFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextformat.h:296
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QTextFormat &)
func (this *QTextFormat) Swap(other QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextFormat_PTR() != nil {
		convArg0 = other.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void merge(const QTextFormat &)
func (this *QTextFormat) Merge(other QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextFormat_PTR() != nil {
		convArg0 = other.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat5mergeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:301
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:302
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QTextFormat) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:304
// index:0
// Public Visibility=Default Availability=Available
// [4] int type()
func (this *QTextFormat) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:306
// index:0
// Public Visibility=Default Availability=Available
// [4] int objectIndex()
func (this *QTextFormat) ObjectIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11objectIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObjectIndex(int)
func (this *QTextFormat) SetObjectIndex(object int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat14setObjectIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), object)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:309
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant property(int)
func (this *QTextFormat) Property(propertyId int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat8propertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:310
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProperty(int, const QVariant &)
func (this *QTextFormat) SetProperty(propertyId int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat11setPropertyEiRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearProperty(int)
func (this *QTextFormat) ClearProperty(propertyId int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13clearPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:312
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasProperty(int)
func (this *QTextFormat) HasProperty(propertyId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11hasPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:314
// index:0
// Public Visibility=Default Availability=Available
// [1] bool boolProperty(int)
func (this *QTextFormat) BoolProperty(propertyId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12boolPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:315
// index:0
// Public Visibility=Default Availability=Available
// [4] int intProperty(int)
func (this *QTextFormat) IntProperty(propertyId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11intPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:316
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal doubleProperty(int)
func (this *QTextFormat) DoubleProperty(propertyId int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat14doublePropertyEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:317
// index:0
// Public Visibility=Default Availability=Available
// [8] QString stringProperty(int)
func (this *QTextFormat) StringProperty(propertyId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat14stringPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:318
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor colorProperty(int)
func (this *QTextFormat) ColorProperty(propertyId int) *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13colorPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QPen penProperty(int)
func (this *QTextFormat) PenProperty(propertyId int) *QPen /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11penPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:320
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brushProperty(int)
func (this *QTextFormat) BrushProperty(propertyId int) *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13brushPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:321
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLength lengthProperty(int)
func (this *QTextFormat) LengthProperty(propertyId int) *QTextLength /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat14lengthPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLength)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:327
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyCount()
func (this *QTextFormat) PropertyCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13propertyCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:329
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setObjectType(int)
func (this *QTextFormat) SetObjectType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13setObjectTypeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:330
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int objectType()
func (this *QTextFormat) ObjectType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat10objectTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:333
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCharFormat()
func (this *QTextFormat) IsCharFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12isCharFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:334
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBlockFormat()
func (this *QTextFormat) IsBlockFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isBlockFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:335
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isListFormat()
func (this *QTextFormat) IsListFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12isListFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:336
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFrameFormat()
func (this *QTextFormat) IsFrameFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isFrameFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:337
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isImageFormat()
func (this *QTextFormat) IsImageFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isImageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:338
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTableFormat()
func (this *QTextFormat) IsTableFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isTableFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:339
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTableCellFormat()
func (this *QTextFormat) IsTableCellFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat17isTableCellFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:341
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlockFormat toBlockFormat()
func (this *QTextFormat) ToBlockFormat() *QTextBlockFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toBlockFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlockFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:342
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat toCharFormat()
func (this *QTextFormat) ToCharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12toCharFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:343
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextListFormat toListFormat()
func (this *QTextFormat) ToListFormat() *QTextListFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12toListFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextListFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:344
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextTableFormat toTableFormat()
func (this *QTextFormat) ToTableFormat() *QTextTableFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toTableFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:345
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFrameFormat toFrameFormat()
func (this *QTextFormat) ToFrameFormat() *QTextFrameFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toFrameFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFrameFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFrameFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:346
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextImageFormat toImageFormat()
func (this *QTextFormat) ToImageFormat() *QTextImageFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toImageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextImageFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:347
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextTableCellFormat toTableCellFormat()
func (this *QTextFormat) ToTableCellFormat() *QTextTableCellFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat17toTableCellFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableCellFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:349
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QTextFormat &)
func (this *QTextFormat) Operator_equal_equal(rhs QTextFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextFormat_PTR() != nil {
		convArg0 = rhs.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormateqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:350
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QTextFormat &)
func (this *QTextFormat) Operator_not_equal(rhs QTextFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextFormat_PTR() != nil {
		convArg0 = rhs.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormatneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:353
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)
func (this *QTextFormat) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:355
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection()
func (this *QTextFormat) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:358
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)
func (this *QTextFormat) SetBackground(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:360
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background()
func (this *QTextFormat) Background() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:362
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearBackground()
func (this *QTextFormat) ClearBackground() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat15clearBackgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:365
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)
func (this *QTextFormat) SetForeground(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13setForegroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:367
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground()
func (this *QTextFormat) Foreground() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat10foregroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:369
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearForeground()
func (this *QTextFormat) ClearForeground() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat15clearForegroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

type QTextFormat__FormatType = int

const QTextFormat__InvalidFormat QTextFormat__FormatType = -1
const QTextFormat__BlockFormat QTextFormat__FormatType = 1
const QTextFormat__CharFormat QTextFormat__FormatType = 2
const QTextFormat__ListFormat QTextFormat__FormatType = 3
const QTextFormat__TableFormat QTextFormat__FormatType = 4
const QTextFormat__FrameFormat QTextFormat__FormatType = 5
const QTextFormat__UserFormat QTextFormat__FormatType = 100

type QTextFormat__Property = int

const QTextFormat__ObjectIndex QTextFormat__Property = 0
const QTextFormat__CssFloat QTextFormat__Property = 2048
const QTextFormat__LayoutDirection QTextFormat__Property = 2049
const QTextFormat__OutlinePen QTextFormat__Property = 2064
const QTextFormat__BackgroundBrush QTextFormat__Property = 2080
const QTextFormat__ForegroundBrush QTextFormat__Property = 2081
const QTextFormat__BackgroundImageUrl QTextFormat__Property = 2083
const QTextFormat__BlockAlignment QTextFormat__Property = 4112
const QTextFormat__BlockTopMargin QTextFormat__Property = 4144
const QTextFormat__BlockBottomMargin QTextFormat__Property = 4145
const QTextFormat__BlockLeftMargin QTextFormat__Property = 4146
const QTextFormat__BlockRightMargin QTextFormat__Property = 4147
const QTextFormat__TextIndent QTextFormat__Property = 4148
const QTextFormat__TabPositions QTextFormat__Property = 4149
const QTextFormat__BlockIndent QTextFormat__Property = 4160
const QTextFormat__LineHeight QTextFormat__Property = 4168
const QTextFormat__LineHeightType QTextFormat__Property = 4169
const QTextFormat__BlockNonBreakableLines QTextFormat__Property = 4176
const QTextFormat__BlockTrailingHorizontalRulerWidth QTextFormat__Property = 4192
const QTextFormat__FirstFontProperty QTextFormat__Property = 8160
const QTextFormat__FontCapitalization QTextFormat__Property = 8160
const QTextFormat__FontLetterSpacingType QTextFormat__Property = 8243
const QTextFormat__FontLetterSpacing QTextFormat__Property = 8161
const QTextFormat__FontWordSpacing QTextFormat__Property = 8162
const QTextFormat__FontStretch QTextFormat__Property = 8244
const QTextFormat__FontStyleHint QTextFormat__Property = 8163
const QTextFormat__FontStyleStrategy QTextFormat__Property = 8164
const QTextFormat__FontKerning QTextFormat__Property = 8165
const QTextFormat__FontHintingPreference QTextFormat__Property = 8166
const QTextFormat__FontFamily QTextFormat__Property = 8192
const QTextFormat__FontPointSize QTextFormat__Property = 8193
const QTextFormat__FontSizeAdjustment QTextFormat__Property = 8194
const QTextFormat__FontSizeIncrement QTextFormat__Property = 8194
const QTextFormat__FontWeight QTextFormat__Property = 8195
const QTextFormat__FontItalic QTextFormat__Property = 8196
const QTextFormat__FontUnderline QTextFormat__Property = 8197
const QTextFormat__FontOverline QTextFormat__Property = 8198
const QTextFormat__FontStrikeOut QTextFormat__Property = 8199
const QTextFormat__FontFixedPitch QTextFormat__Property = 8200
const QTextFormat__FontPixelSize QTextFormat__Property = 8201
const QTextFormat__LastFontProperty QTextFormat__Property = 8201
const QTextFormat__TextUnderlineColor QTextFormat__Property = 8208
const QTextFormat__TextVerticalAlignment QTextFormat__Property = 8225
const QTextFormat__TextOutline QTextFormat__Property = 8226
const QTextFormat__TextUnderlineStyle QTextFormat__Property = 8227
const QTextFormat__TextToolTip QTextFormat__Property = 8228
const QTextFormat__IsAnchor QTextFormat__Property = 8240
const QTextFormat__AnchorHref QTextFormat__Property = 8241
const QTextFormat__AnchorName QTextFormat__Property = 8242
const QTextFormat__ObjectType QTextFormat__Property = 12032
const QTextFormat__ListStyle QTextFormat__Property = 12288
const QTextFormat__ListIndent QTextFormat__Property = 12289
const QTextFormat__ListNumberPrefix QTextFormat__Property = 12290
const QTextFormat__ListNumberSuffix QTextFormat__Property = 12291
const QTextFormat__FrameBorder QTextFormat__Property = 16384
const QTextFormat__FrameMargin QTextFormat__Property = 16385
const QTextFormat__FramePadding QTextFormat__Property = 16386
const QTextFormat__FrameWidth QTextFormat__Property = 16387
const QTextFormat__FrameHeight QTextFormat__Property = 16388
const QTextFormat__FrameTopMargin QTextFormat__Property = 16389
const QTextFormat__FrameBottomMargin QTextFormat__Property = 16390
const QTextFormat__FrameLeftMargin QTextFormat__Property = 16391
const QTextFormat__FrameRightMargin QTextFormat__Property = 16392
const QTextFormat__FrameBorderBrush QTextFormat__Property = 16393
const QTextFormat__FrameBorderStyle QTextFormat__Property = 16400
const QTextFormat__TableColumns QTextFormat__Property = 16640
const QTextFormat__TableColumnWidthConstraints QTextFormat__Property = 16641
const QTextFormat__TableCellSpacing QTextFormat__Property = 16642
const QTextFormat__TableCellPadding QTextFormat__Property = 16643
const QTextFormat__TableHeaderRowCount QTextFormat__Property = 16644
const QTextFormat__TableCellRowSpan QTextFormat__Property = 18448
const QTextFormat__TableCellColumnSpan QTextFormat__Property = 18449
const QTextFormat__TableCellTopPadding QTextFormat__Property = 18450
const QTextFormat__TableCellBottomPadding QTextFormat__Property = 18451
const QTextFormat__TableCellLeftPadding QTextFormat__Property = 18452
const QTextFormat__TableCellRightPadding QTextFormat__Property = 18453
const QTextFormat__ImageName QTextFormat__Property = 20480
const QTextFormat__ImageWidth QTextFormat__Property = 20496
const QTextFormat__ImageHeight QTextFormat__Property = 20497
const QTextFormat__FullWidthSelection QTextFormat__Property = 24576
const QTextFormat__PageBreakPolicy QTextFormat__Property = 28672
const QTextFormat__UserProperty QTextFormat__Property = 1048576

type QTextFormat__ObjectTypes = int

const QTextFormat__NoObject QTextFormat__ObjectTypes = 0
const QTextFormat__ImageObject QTextFormat__ObjectTypes = 1
const QTextFormat__TableObject QTextFormat__ObjectTypes = 2
const QTextFormat__TableCellObject QTextFormat__ObjectTypes = 3
const QTextFormat__UserObject QTextFormat__ObjectTypes = 4096

type QTextFormat__PageBreakFlag = int

const QTextFormat__PageBreak_Auto QTextFormat__PageBreakFlag = 0
const QTextFormat__PageBreak_AlwaysBefore QTextFormat__PageBreakFlag = 1
const QTextFormat__PageBreak_AlwaysAfter QTextFormat__PageBreakFlag = 16

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
