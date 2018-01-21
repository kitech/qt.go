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
func NewQTextFormatFromPointer(cthis unsafe.Pointer) *QTextFormat {
	return &QTextFormat{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextformat.h:288
// index:0
// Public
// void QTextFormat()
func NewQTextFormat() *QTextFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:290
// index:1
// Public
// void QTextFormat(int)
func NewQTextFormat_1(type_ int) *QTextFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:294
// index:0
// Public
// void ~QTextFormat()
func DeleteQTextFormat(*QTextFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:296
// index:0
// Public inline
// void swap(class QTextFormat &)
func (this *QTextFormat) Swap(other *QTextFormat) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:299
// index:0
// Public
// void merge(const class QTextFormat &)
func (this *QTextFormat) Merge(other *QTextFormat) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat5mergeERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:301
// index:0
// Public inline
// bool isValid()
func (this *QTextFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:302
// index:0
// Public inline
// bool isEmpty()
func (this *QTextFormat) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:304
// index:0
// Public
// int type()
func (this *QTextFormat) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:306
// index:0
// Public
// int objectIndex()
func (this *QTextFormat) ObjectIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11objectIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:307
// index:0
// Public
// void setObjectIndex(int)
func (this *QTextFormat) SetObjectIndex(object int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat14setObjectIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:309
// index:0
// Public
// QVariant property(int)
func (this *QTextFormat) Property(propertyId int) *qtcore.QVariant /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat8propertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:310
// index:0
// Public
// void setProperty(int, const class QVariant &)
func (this *QTextFormat) SetProperty(propertyId int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat11setPropertyEiRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:311
// index:0
// Public
// void clearProperty(int)
func (this *QTextFormat) ClearProperty(propertyId int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13clearPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:312
// index:0
// Public
// bool hasProperty(int)
func (this *QTextFormat) HasProperty(propertyId int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11hasPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:314
// index:0
// Public
// bool boolProperty(int)
func (this *QTextFormat) BoolProperty(propertyId int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12boolPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:315
// index:0
// Public
// int intProperty(int)
func (this *QTextFormat) IntProperty(propertyId int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11intPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:316
// index:0
// Public
// qreal doubleProperty(int)
func (this *QTextFormat) DoubleProperty(propertyId int) float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14doublePropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextformat.h:317
// index:0
// Public
// QString stringProperty(int)
func (this *QTextFormat) StringProperty(propertyId int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14stringPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:318
// index:0
// Public
// QColor colorProperty(int)
func (this *QTextFormat) ColorProperty(propertyId int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13colorPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:319
// index:0
// Public
// QPen penProperty(int)
func (this *QTextFormat) PenProperty(propertyId int) *QPen /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11penPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:320
// index:0
// Public
// QBrush brushProperty(int)
func (this *QTextFormat) BrushProperty(propertyId int) *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13brushPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:321
// index:0
// Public
// QTextLength lengthProperty(int)
func (this *QTextFormat) LengthProperty(propertyId int) *QTextLength /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14lengthPropertyEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &propertyId)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:327
// index:0
// Public
// int propertyCount()
func (this *QTextFormat) PropertyCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13propertyCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:329
// index:0
// Public inline
// void setObjectType(int)
func (this *QTextFormat) SetObjectType(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setObjectTypeEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:330
// index:0
// Public inline
// int objectType()
func (this *QTextFormat) ObjectType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10objectTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:333
// index:0
// Public inline
// bool isCharFormat()
func (this *QTextFormat) IsCharFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12isCharFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:334
// index:0
// Public inline
// bool isBlockFormat()
func (this *QTextFormat) IsBlockFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isBlockFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:335
// index:0
// Public inline
// bool isListFormat()
func (this *QTextFormat) IsListFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12isListFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:336
// index:0
// Public inline
// bool isFrameFormat()
func (this *QTextFormat) IsFrameFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isFrameFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:337
// index:0
// Public inline
// bool isImageFormat()
func (this *QTextFormat) IsImageFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isImageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:338
// index:0
// Public inline
// bool isTableFormat()
func (this *QTextFormat) IsTableFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isTableFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:339
// index:0
// Public inline
// bool isTableCellFormat()
func (this *QTextFormat) IsTableCellFormat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat17isTableCellFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:341
// index:0
// Public
// QTextBlockFormat toBlockFormat()
func (this *QTextFormat) ToBlockFormat() *QTextBlockFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toBlockFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:342
// index:0
// Public
// QTextCharFormat toCharFormat()
func (this *QTextFormat) ToCharFormat() *QTextCharFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12toCharFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:343
// index:0
// Public
// QTextListFormat toListFormat()
func (this *QTextFormat) ToListFormat() *QTextListFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12toListFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:344
// index:0
// Public
// QTextTableFormat toTableFormat()
func (this *QTextFormat) ToTableFormat() *QTextTableFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toTableFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextTableFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:345
// index:0
// Public
// QTextFrameFormat toFrameFormat()
func (this *QTextFormat) ToFrameFormat() *QTextFrameFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toFrameFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextFrameFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:346
// index:0
// Public
// QTextImageFormat toImageFormat()
func (this *QTextFormat) ToImageFormat() *QTextImageFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toImageFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:347
// index:0
// Public
// QTextTableCellFormat toTableCellFormat()
func (this *QTextFormat) ToTableCellFormat() *QTextTableCellFormat /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat17toTableCellFormatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:353
// index:0
// Public inline
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QTextFormat) SetLayoutDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:355
// index:0
// Public inline
// Qt::LayoutDirection layoutDirection()
func (this *QTextFormat) LayoutDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat15layoutDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:358
// index:0
// Public inline
// void setBackground(const class QBrush &)
func (this *QTextFormat) SetBackground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:360
// index:0
// Public inline
// QBrush background()
func (this *QTextFormat) Background() *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10backgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:362
// index:0
// Public inline
// void clearBackground()
func (this *QTextFormat) ClearBackground() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat15clearBackgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:365
// index:0
// Public inline
// void setForeground(const class QBrush &)
func (this *QTextFormat) SetForeground(brush *QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setForegroundERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:367
// index:0
// Public inline
// QBrush foreground()
func (this *QTextFormat) Foreground() *QBrush /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10foregroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:369
// index:0
// Public inline
// void clearForeground()
func (this *QTextFormat) ClearForeground() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat15clearForegroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
