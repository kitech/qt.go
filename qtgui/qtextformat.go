package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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

type QTextFormat struct {
	*qtrt.CObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:290
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextFormat(int)
func NewQTextFormat_1(type_ int) *QTextFormat {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatC2Ei", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:294
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextFormat()
func DeleteQTextFormat(this *QTextFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextformat.h:296
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QTextFormat &)
func (this *QTextFormat) Swap(other *QTextFormat) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void merge(const QTextFormat &)
func (this *QTextFormat) Merge(other *QTextFormat) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat5mergeERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:301
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:302
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QTextFormat) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:304
// index:0
// Public Visibility=Default Availability=Available
// [4] int type()
func (this *QTextFormat) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:306
// index:0
// Public Visibility=Default Availability=Available
// [4] int objectIndex()
func (this *QTextFormat) ObjectIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11objectIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObjectIndex(int)
func (this *QTextFormat) SetObjectIndex(object int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat14setObjectIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:309
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant property(int)
func (this *QTextFormat) Property(propertyId int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat8propertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:310
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProperty(int, const QVariant &)
func (this *QTextFormat) SetProperty(propertyId int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat11setPropertyEiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearProperty(int)
func (this *QTextFormat) ClearProperty(propertyId int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13clearPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:312
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasProperty(int)
func (this *QTextFormat) HasProperty(propertyId int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11hasPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:314
// index:0
// Public Visibility=Default Availability=Available
// [1] bool boolProperty(int)
func (this *QTextFormat) BoolProperty(propertyId int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12boolPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:315
// index:0
// Public Visibility=Default Availability=Available
// [4] int intProperty(int)
func (this *QTextFormat) IntProperty(propertyId int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11intPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:316
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal doubleProperty(int)
func (this *QTextFormat) DoubleProperty(propertyId int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14doublePropertyEi", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:317
// index:0
// Public Visibility=Default Availability=Available
// [8] QString stringProperty(int)
func (this *QTextFormat) StringProperty(propertyId int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14stringPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:318
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor colorProperty(int)
func (this *QTextFormat) ColorProperty(propertyId int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13colorPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QPen penProperty(int)
func (this *QTextFormat) PenProperty(propertyId int) *QPen /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11penPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:320
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brushProperty(int)
func (this *QTextFormat) BrushProperty(propertyId int) *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13brushPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:321
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLength lengthProperty(int)
func (this *QTextFormat) LengthProperty(propertyId int) *QTextLength /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14lengthPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLength)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:327
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyCount()
func (this *QTextFormat) PropertyCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13propertyCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:329
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setObjectType(int)
func (this *QTextFormat) SetObjectType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setObjectTypeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:330
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int objectType()
func (this *QTextFormat) ObjectType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10objectTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:333
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCharFormat()
func (this *QTextFormat) IsCharFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12isCharFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:334
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBlockFormat()
func (this *QTextFormat) IsBlockFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isBlockFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:335
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isListFormat()
func (this *QTextFormat) IsListFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12isListFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:336
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFrameFormat()
func (this *QTextFormat) IsFrameFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isFrameFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:337
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isImageFormat()
func (this *QTextFormat) IsImageFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isImageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:338
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTableFormat()
func (this *QTextFormat) IsTableFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isTableFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:339
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTableCellFormat()
func (this *QTextFormat) IsTableCellFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat17isTableCellFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:341
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlockFormat toBlockFormat()
func (this *QTextFormat) ToBlockFormat() *QTextBlockFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toBlockFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlockFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:342
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat toCharFormat()
func (this *QTextFormat) ToCharFormat() *QTextCharFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12toCharFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:343
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextListFormat toListFormat()
func (this *QTextFormat) ToListFormat() *QTextListFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12toListFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextListFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:344
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextTableFormat toTableFormat()
func (this *QTextFormat) ToTableFormat() *QTextTableFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toTableFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextTableFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:345
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFrameFormat toFrameFormat()
func (this *QTextFormat) ToFrameFormat() *QTextFrameFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toFrameFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextFrameFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFrameFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:346
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextImageFormat toImageFormat()
func (this *QTextFormat) ToImageFormat() *QTextImageFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toImageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextImageFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:347
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextTableCellFormat toTableCellFormat()
func (this *QTextFormat) ToTableCellFormat() *QTextTableCellFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat17toTableCellFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableCellFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:353
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)
func (this *QTextFormat) SetLayoutDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:355
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection()
func (this *QTextFormat) LayoutDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat15layoutDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:358
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)
func (this *QTextFormat) SetBackground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:360
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background()
func (this *QTextFormat) Background() *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:362
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearBackground()
func (this *QTextFormat) ClearBackground() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat15clearBackgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:365
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)
func (this *QTextFormat) SetForeground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setForegroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:367
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground()
func (this *QTextFormat) Foreground() *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:369
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearForeground()
func (this *QTextFormat) ClearForeground() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat15clearForegroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QTextFormat__FormatType = int

const QTextFormat__InvalidFormat QTextFormat__FormatType = 4294967295
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
