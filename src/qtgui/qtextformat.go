//  header block begin
// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>
package qtgui

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtextformat.h:288
// index:0
// void QTextFormat()
func NewQTextFormat() *QTextFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QTextFormat{cthis}
}

// /usr/include/qt/QtGui/qtextformat.h:290
// index:1
// void QTextFormat(int)
func NewQTextFormat_1(type_ int) *QTextFormat {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatC2Ei", ffiqt.FFI_TYPE_VOID, cthis, &type_)
	gopp.ErrPrint(err, rv)
	return &QTextFormat{cthis}
}

// /usr/include/qt/QtGui/qtextformat.h:294
// index:0
// void ~QTextFormat()
func DeleteQTextFormat(*QTextFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormatD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:296
// index:0
// inline
// void swap(class QTextFormat &)
func (this *QTextFormat) Swap(other unsafe.Pointer) {
	// 0: (, QTextFormat & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:299
// index:0
// void merge(const class QTextFormat &)
func (this *QTextFormat) Merge(other unsafe.Pointer) {
	// 0: (, const QTextFormat & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat5mergeERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:301
// index:0
// inline
// bool isValid()
func (this *QTextFormat) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:302
// index:0
// inline
// bool isEmpty()
func (this *QTextFormat) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:304
// index:0
// int type()
func (this *QTextFormat) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:306
// index:0
// int objectIndex()
func (this *QTextFormat) ObjectIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11objectIndexEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:307
// index:0
// void setObjectIndex(int)
func (this *QTextFormat) SetObjectIndex(object int) {
	// 0: (, int object), (&object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat14setObjectIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:309
// index:0
// QVariant property(int)
func (this *QTextFormat) Property(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat8propertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:310
// index:0
// void setProperty(int, const class QVariant &)
func (this *QTextFormat) SetProperty(propertyId int, value unsafe.Pointer) {
	// 0: (, int propertyId, const QVariant & value), (&propertyId, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat11setPropertyEiRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:311
// index:0
// void clearProperty(int)
func (this *QTextFormat) ClearProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13clearPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:312
// index:0
// bool hasProperty(int)
func (this *QTextFormat) HasProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11hasPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:314
// index:0
// bool boolProperty(int)
func (this *QTextFormat) BoolProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12boolPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:315
// index:0
// int intProperty(int)
func (this *QTextFormat) IntProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11intPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:316
// index:0
// qreal doubleProperty(int)
func (this *QTextFormat) DoubleProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14doublePropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:317
// index:0
// QString stringProperty(int)
func (this *QTextFormat) StringProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14stringPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:318
// index:0
// QColor colorProperty(int)
func (this *QTextFormat) ColorProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13colorPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:319
// index:0
// QPen penProperty(int)
func (this *QTextFormat) PenProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat11penPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:320
// index:0
// QBrush brushProperty(int)
func (this *QTextFormat) BrushProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13brushPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:321
// index:0
// QTextLength lengthProperty(int)
func (this *QTextFormat) LengthProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat14lengthPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:322
// index:0
// QVector<QTextLength> lengthVectorProperty(int)
func (this *QTextFormat) LengthVectorProperty(propertyId int) {
	// 0: (, int propertyId), (&propertyId)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat20lengthVectorPropertyEi", ffiqt.FFI_TYPE_VOID, this.cthis, &propertyId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:326
// index:0
// QMap<int, QVariant> properties()
func (this *QTextFormat) Properties() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10propertiesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:327
// index:0
// int propertyCount()
func (this *QTextFormat) PropertyCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13propertyCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:329
// index:0
// inline
// void setObjectType(int)
func (this *QTextFormat) SetObjectType(type_ int) {
	// 0: (, int type), (&type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setObjectTypeEi", ffiqt.FFI_TYPE_VOID, this.cthis, &type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:330
// index:0
// inline
// int objectType()
func (this *QTextFormat) ObjectType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10objectTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:333
// index:0
// inline
// bool isCharFormat()
func (this *QTextFormat) IsCharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12isCharFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:334
// index:0
// inline
// bool isBlockFormat()
func (this *QTextFormat) IsBlockFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isBlockFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:335
// index:0
// inline
// bool isListFormat()
func (this *QTextFormat) IsListFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12isListFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:336
// index:0
// inline
// bool isFrameFormat()
func (this *QTextFormat) IsFrameFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isFrameFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:337
// index:0
// inline
// bool isImageFormat()
func (this *QTextFormat) IsImageFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isImageFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:338
// index:0
// inline
// bool isTableFormat()
func (this *QTextFormat) IsTableFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13isTableFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:339
// index:0
// inline
// bool isTableCellFormat()
func (this *QTextFormat) IsTableCellFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat17isTableCellFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:341
// index:0
// QTextBlockFormat toBlockFormat()
func (this *QTextFormat) ToBlockFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toBlockFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:342
// index:0
// QTextCharFormat toCharFormat()
func (this *QTextFormat) ToCharFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12toCharFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:343
// index:0
// QTextListFormat toListFormat()
func (this *QTextFormat) ToListFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat12toListFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:344
// index:0
// QTextTableFormat toTableFormat()
func (this *QTextFormat) ToTableFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toTableFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:345
// index:0
// QTextFrameFormat toFrameFormat()
func (this *QTextFormat) ToFrameFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toFrameFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:346
// index:0
// QTextImageFormat toImageFormat()
func (this *QTextFormat) ToImageFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat13toImageFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:347
// index:0
// QTextTableCellFormat toTableCellFormat()
func (this *QTextFormat) ToTableCellFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat17toTableCellFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:353
// index:0
// inline
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QTextFormat) SetLayoutDirection(direction int) {
	// 0: (, Qt::LayoutDirection direction), (&direction)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_VOID, this.cthis, &direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:355
// index:0
// inline
// Qt::LayoutDirection layoutDirection()
func (this *QTextFormat) LayoutDirection() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat15layoutDirectionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:358
// index:0
// inline
// void setBackground(const class QBrush &)
func (this *QTextFormat) SetBackground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setBackgroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:360
// index:0
// inline
// QBrush background()
func (this *QTextFormat) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10backgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:362
// index:0
// inline
// void clearBackground()
func (this *QTextFormat) ClearBackground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat15clearBackgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:365
// index:0
// inline
// void setForeground(const class QBrush &)
func (this *QTextFormat) SetForeground(brush unsafe.Pointer) {
	// 0: (, const QBrush & brush), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat13setForegroundERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:367
// index:0
// inline
// QBrush foreground()
func (this *QTextFormat) Foreground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QTextFormat10foregroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:369
// index:0
// inline
// void clearForeground()
func (this *QTextFormat) ClearForeground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QTextFormat15clearForegroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
