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

/*

 */
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

// /usr/include/qt/QtGui/qtextformat.h:290
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextFormat()

/*
Creates a new text format with an InvalidFormat.

See also FormatType.
*/
func (*QTextFormat) NewForInherit() *QTextFormat {
	return NewQTextFormat()
}
func NewQTextFormat() *QTextFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:292
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextFormat(int)

/*
Creates a new text format with an InvalidFormat.

See also FormatType.
*/
func (*QTextFormat) NewForInherit1(type_ int) *QTextFormat {
	return NewQTextFormat1(type_)
}
func NewQTextFormat1(type_ int) *QTextFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormatC2Ei", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:295
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFormat & operator=(const QTextFormat &)

/*

 */
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

// /usr/include/qt/QtGui/qtextformat.h:296
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextFormat()

/*

 */
func DeleteQTextFormat(this *QTextFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextformat.h:298
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QTextFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QTextFormat) Swap(other QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextFormat_PTR() != nil {
		convArg0 = other.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void merge(const QTextFormat &)

/*
Merges the other format with this format; where there are conflicts the other format takes precedence.
*/
func (this *QTextFormat) Merge(other QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextFormat_PTR() != nil {
		convArg0 = other.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat5mergeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:303
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the format is valid (i.e. is not InvalidFormat); otherwise returns false.
*/
func (this *QTextFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:304
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the format does not store any properties; false otherwise.

This function was introduced in  Qt 5.3.

See also propertyCount() and properties().
*/
func (this *QTextFormat) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:306
// index:0
// Public Visibility=Default Availability=Available
// [4] int type() const

/*
Returns the type of this format.

See also FormatType.
*/
func (this *QTextFormat) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:308
// index:0
// Public Visibility=Default Availability=Available
// [4] int objectIndex() const

/*
Returns the index of the format object, or -1 if the format object is invalid.

See also setObjectIndex().
*/
func (this *QTextFormat) ObjectIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11objectIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:309
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setObjectIndex(int)

/*
Sets the format object's object index.

See also objectIndex().
*/
func (this *QTextFormat) SetObjectIndex(object int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat14setObjectIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), object)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:311
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant property(int) const

/*
Returns the property specified by the given propertyId.

See also setProperty() and Property.
*/
func (this *QTextFormat) Property(propertyId int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat8propertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:312
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProperty(int, const QVariant &)

/*
Sets the property specified by the propertyId to the given value.

See also property() and Property.
*/
func (this *QTextFormat) SetProperty(propertyId int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat11setPropertyEiRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:313
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearProperty(int)

/*
Clears the value of the property given by propertyId

See also Property.
*/
func (this *QTextFormat) ClearProperty(propertyId int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13clearPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:314
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasProperty(int) const

/*
Returns true if the text format has a property with the given propertyId; otherwise returns false.

See also properties() and Property.
*/
func (this *QTextFormat) HasProperty(propertyId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11hasPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:316
// index:0
// Public Visibility=Default Availability=Available
// [1] bool boolProperty(int) const

/*
Returns the value of the property specified by propertyId. If the property isn't of QTextFormat::Bool type, false is returned instead.

See also setProperty(), intProperty(), doubleProperty(), stringProperty(), colorProperty(), lengthProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) BoolProperty(propertyId int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12boolPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:317
// index:0
// Public Visibility=Default Availability=Available
// [4] int intProperty(int) const

/*
Returns the value of the property specified by propertyId. If the property is not of QTextFormat::Integer type, 0 is returned instead.

See also setProperty(), boolProperty(), doubleProperty(), stringProperty(), colorProperty(), lengthProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) IntProperty(propertyId int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11intPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:318
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal doubleProperty(int) const

/*
Returns the value of the property specified by propertyId. If the property isn't of QVariant::Double or QMetaType::Float type, 0 is returned instead.

See also setProperty(), boolProperty(), intProperty(), stringProperty(), colorProperty(), lengthProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) DoubleProperty(propertyId int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat14doublePropertyEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QString stringProperty(int) const

/*
Returns the value of the property given by propertyId; if the property isn't of QVariant::String type, an empty string is returned instead.

See also setProperty(), boolProperty(), intProperty(), doubleProperty(), colorProperty(), lengthProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) StringProperty(propertyId int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat14stringPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:320
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor colorProperty(int) const

/*
Returns the value of the property given by propertyId; if the property isn't of QVariant::Color type, an invalid color is returned instead.

See also setProperty(), boolProperty(), intProperty(), doubleProperty(), stringProperty(), lengthProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) ColorProperty(propertyId int) *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13colorPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:321
// index:0
// Public Visibility=Default Availability=Available
// [8] QPen penProperty(int) const

/*
Returns the value of the property given by propertyId; if the property isn't of QVariant::Pen type, Qt::NoPen is returned instead.

See also setProperty(), boolProperty(), intProperty(), doubleProperty(), stringProperty(), lengthProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) PenProperty(propertyId int) *QPen /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat11penPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:322
// index:0
// Public Visibility=Default Availability=Available
// [8] QBrush brushProperty(int) const

/*
Returns the value of the property given by propertyId; if the property isn't of QVariant::Brush type, Qt::NoBrush is returned instead.

See also setProperty(), boolProperty(), intProperty(), doubleProperty(), stringProperty(), lengthProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) BrushProperty(propertyId int) *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13brushPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:323
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextLength lengthProperty(int) const

/*
Returns the value of the property given by propertyId.

See also setProperty(), boolProperty(), intProperty(), doubleProperty(), stringProperty(), colorProperty(), lengthVectorProperty(), and Property.
*/
func (this *QTextFormat) LengthProperty(propertyId int) *QTextLength /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat14lengthPropertyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), propertyId)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextLengthFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextLength)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:329
// index:0
// Public Visibility=Default Availability=Available
// [4] int propertyCount() const

/*
Returns the number of properties stored in the format.

This function was introduced in  Qt 4.3.
*/
func (this *QTextFormat) PropertyCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13propertyCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:331
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setObjectType(int)

/*
Sets the text format's object type to type.

See also ObjectTypes and objectType().
*/
func (this *QTextFormat) SetObjectType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13setObjectTypeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:332
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int objectType() const

/*
Returns the text format's object type.

See also ObjectTypes and setObjectType().
*/
func (this *QTextFormat) ObjectType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat10objectTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:335
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isCharFormat() const

/*
Returns true if this text format is a CharFormat; otherwise returns false.
*/
func (this *QTextFormat) IsCharFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12isCharFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:336
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isBlockFormat() const

/*
Returns true if this text format is a BlockFormat; otherwise returns false.
*/
func (this *QTextFormat) IsBlockFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isBlockFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:337
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isListFormat() const

/*
Returns true if this text format is a ListFormat; otherwise returns false.
*/
func (this *QTextFormat) IsListFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12isListFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:338
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isFrameFormat() const

/*
Returns true if this text format is a FrameFormat; otherwise returns false.
*/
func (this *QTextFormat) IsFrameFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isFrameFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:339
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isImageFormat() const

/*
Returns true if this text format is an image format; otherwise returns false.
*/
func (this *QTextFormat) IsImageFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isImageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:340
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTableFormat() const

/*
Returns true if this text format is a TableFormat; otherwise returns false.
*/
func (this *QTextFormat) IsTableFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13isTableFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:341
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isTableCellFormat() const

/*
Returns true if this text format is a TableCellFormat; otherwise returns false.

This function was introduced in  Qt 4.4.
*/
func (this *QTextFormat) IsTableCellFormat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat17isTableCellFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:343
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextBlockFormat toBlockFormat() const

/*
Returns this format as a block format.
*/
func (this *QTextFormat) ToBlockFormat() *QTextBlockFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toBlockFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextBlockFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextBlockFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:344
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat toCharFormat() const

/*
Returns this format as a character format.
*/
func (this *QTextFormat) ToCharFormat() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12toCharFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:345
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextListFormat toListFormat() const

/*
Returns this format as a list format.
*/
func (this *QTextFormat) ToListFormat() *QTextListFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat12toListFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextListFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextListFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:346
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextTableFormat toTableFormat() const

/*
Returns this format as a table format.
*/
func (this *QTextFormat) ToTableFormat() *QTextTableFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toTableFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:347
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextFrameFormat toFrameFormat() const

/*
Returns this format as a frame format.
*/
func (this *QTextFormat) ToFrameFormat() *QTextFrameFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toFrameFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFrameFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFrameFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:348
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextImageFormat toImageFormat() const

/*
Returns this format as an image format.
*/
func (this *QTextFormat) ToImageFormat() *QTextImageFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat13toImageFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextImageFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextImageFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:349
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextTableCellFormat toTableCellFormat() const

/*
Returns this format as a table cell format.

This function was introduced in  Qt 4.4.
*/
func (this *QTextFormat) ToTableCellFormat() *QTextTableCellFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat17toTableCellFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableCellFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:351
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QTextFormat &) const

/*

 */
func (this *QTextFormat) Operator_equal_equal(rhs QTextFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextFormat_PTR() != nil {
		convArg0 = rhs.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormateqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:352
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QTextFormat &) const

/*

 */
func (this *QTextFormat) Operator_not_equal(rhs QTextFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if rhs != nil && rhs.QTextFormat_PTR() != nil {
		convArg0 = rhs.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormatneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:355
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)

/*
Sets the document's layout direction to the specified direction.

See also layoutDirection().
*/
func (this *QTextFormat) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:357
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection() const

/*
Returns the document's layout direction.

See also setLayoutDirection().
*/
func (this *QTextFormat) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:360
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBackground(const QBrush &)

/*
Sets the brush use to paint the document's background to the brush specified.

See also background(), clearBackground(), and setForeground().
*/
func (this *QTextFormat) SetBackground(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13setBackgroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:362
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush background() const

/*
Returns the brush used to paint the document's background.

See also setBackground(), clearBackground(), and foreground().
*/
func (this *QTextFormat) Background() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat10backgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:364
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearBackground()

/*
Clears the brush used to paint the document's background. The default brush will be used.

See also background(), setBackground(), and clearForeground().
*/
func (this *QTextFormat) ClearBackground() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat15clearBackgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:367
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setForeground(const QBrush &)

/*
Sets the foreground brush to the specified brush. The foreground brush is mostly used to render text.

See also foreground(), clearForeground(), and setBackground().
*/
func (this *QTextFormat) SetForeground(brush QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if brush != nil && brush.QBrush_PTR() != nil {
		convArg0 = brush.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat13setForegroundERK6QBrush", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:369
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QBrush foreground() const

/*
Returns the brush used to render foreground details, such as text, frame outlines, and table borders.

See also setForeground(), clearForeground(), and background().
*/
func (this *QTextFormat) Foreground() *QBrush /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextFormat10foregroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQBrush)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:371
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearForeground()

/*
Clears the brush used to paint the document's foreground. The default brush will be used.

See also foreground(), setForeground(), and clearBackground().
*/
func (this *QTextFormat) ClearForeground() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextFormat15clearForegroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the text item a QTextFormat object is formatting.



See also QTextCharFormat, QTextBlockFormat, QTextListFormat, QTextTableFormat, and type().

*/
type QTextFormat__FormatType = int

//
const QTextFormat__InvalidFormat QTextFormat__FormatType = -1

// The object formats a text block
const QTextFormat__BlockFormat QTextFormat__FormatType = 1

// The object formats a single character
const QTextFormat__CharFormat QTextFormat__FormatType = 2

// The object formats a list
const QTextFormat__ListFormat QTextFormat__FormatType = 3

//
const QTextFormat__TableFormat QTextFormat__FormatType = 4

// The object formats a frame
const QTextFormat__FrameFormat QTextFormat__FormatType = 5

//
const QTextFormat__UserFormat QTextFormat__FormatType = 100

func (this *QTextFormat) FormatTypeItemName(val int) string {
	switch val {
	case QTextFormat__InvalidFormat: // -1
		return "InvalidFormat"
	case QTextFormat__BlockFormat: // 1
		return "BlockFormat"
	case QTextFormat__CharFormat: // 2
		return "CharFormat"
	case QTextFormat__ListFormat: // 3
		return "ListFormat"
	case QTextFormat__TableFormat: // 4
		return "TableFormat"
	case QTextFormat__FrameFormat: // 5
		return "FrameFormat"
	case QTextFormat__UserFormat: // 100
		return "UserFormat"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextFormat_FormatTypeItemName(val int) string {
	var nilthis *QTextFormat
	return nilthis.FormatTypeItemName(val)
}

/*
This enum describes the different properties a format can have.



Paragraph and character properties



Paragraph properties



Character properties

QTextFormat::FontCapitalizationFirstFontPropertySpecifies the capitalization type that is to be applied to the text.


List properties



Table and frame properties



Table cell properties

ConstantValue


Image properties

ConstantValue


Selection properties



Page break properties



See also property() and setProperty().

*/
type QTextFormat__Property = int

//
const QTextFormat__ObjectIndex QTextFormat__Property = 0

//
const QTextFormat__CssFloat QTextFormat__Property = 2048

//
const QTextFormat__LayoutDirection QTextFormat__Property = 2049

//
const QTextFormat__OutlinePen QTextFormat__Property = 2064

//
const QTextFormat__BackgroundBrush QTextFormat__Property = 2080

//
const QTextFormat__ForegroundBrush QTextFormat__Property = 2081

//
const QTextFormat__BackgroundImageUrl QTextFormat__Property = 2083

//
const QTextFormat__BlockAlignment QTextFormat__Property = 4112

//
const QTextFormat__BlockTopMargin QTextFormat__Property = 4144

//
const QTextFormat__BlockBottomMargin QTextFormat__Property = 4145

//
const QTextFormat__BlockLeftMargin QTextFormat__Property = 4146

//
const QTextFormat__BlockRightMargin QTextFormat__Property = 4147

//
const QTextFormat__TextIndent QTextFormat__Property = 4148

//
const QTextFormat__TabPositions QTextFormat__Property = 4149

//
const QTextFormat__BlockIndent QTextFormat__Property = 4160

//
const QTextFormat__LineHeight QTextFormat__Property = 4168

//
const QTextFormat__LineHeightType QTextFormat__Property = 4169

//
const QTextFormat__BlockNonBreakableLines QTextFormat__Property = 4176

//
const QTextFormat__BlockTrailingHorizontalRulerWidth QTextFormat__Property = 4192

//
const QTextFormat__HeadingLevel QTextFormat__Property = 4208

//
const QTextFormat__FirstFontProperty QTextFormat__Property = 8160

//
const QTextFormat__FontCapitalization QTextFormat__Property = 8160

//
const QTextFormat__FontLetterSpacingType QTextFormat__Property = 8243

//
const QTextFormat__FontLetterSpacing QTextFormat__Property = 8161

//
const QTextFormat__FontWordSpacing QTextFormat__Property = 8162

//
const QTextFormat__FontStretch QTextFormat__Property = 8244

//
const QTextFormat__FontStyleHint QTextFormat__Property = 8163

//
const QTextFormat__FontStyleStrategy QTextFormat__Property = 8164

//
const QTextFormat__FontKerning QTextFormat__Property = 8165

//
const QTextFormat__FontHintingPreference QTextFormat__Property = 8166

//
const QTextFormat__FontFamily QTextFormat__Property = 8192

//
const QTextFormat__FontPointSize QTextFormat__Property = 8193

//
const QTextFormat__FontSizeAdjustment QTextFormat__Property = 8194

//
const QTextFormat__FontSizeIncrement QTextFormat__Property = 8194

//
const QTextFormat__FontWeight QTextFormat__Property = 8195

//
const QTextFormat__FontItalic QTextFormat__Property = 8196

//
const QTextFormat__FontUnderline QTextFormat__Property = 8197

//
const QTextFormat__FontOverline QTextFormat__Property = 8198

//
const QTextFormat__FontStrikeOut QTextFormat__Property = 8199

//
const QTextFormat__FontFixedPitch QTextFormat__Property = 8200

//
const QTextFormat__FontPixelSize QTextFormat__Property = 8201

//
const QTextFormat__LastFontProperty QTextFormat__Property = 8201

//
const QTextFormat__TextUnderlineColor QTextFormat__Property = 8208

//
const QTextFormat__TextVerticalAlignment QTextFormat__Property = 8225

//
const QTextFormat__TextOutline QTextFormat__Property = 8226

//
const QTextFormat__TextUnderlineStyle QTextFormat__Property = 8227

//
const QTextFormat__TextToolTip QTextFormat__Property = 8228

//
const QTextFormat__IsAnchor QTextFormat__Property = 8240

//
const QTextFormat__AnchorHref QTextFormat__Property = 8241

//
const QTextFormat__AnchorName QTextFormat__Property = 8242

//
const QTextFormat__ObjectType QTextFormat__Property = 12032

//
const QTextFormat__ListStyle QTextFormat__Property = 12288

//
const QTextFormat__ListIndent QTextFormat__Property = 12289

//
const QTextFormat__ListNumberPrefix QTextFormat__Property = 12290

//
const QTextFormat__ListNumberSuffix QTextFormat__Property = 12291

//
const QTextFormat__FrameBorder QTextFormat__Property = 16384

//
const QTextFormat__FrameMargin QTextFormat__Property = 16385

//
const QTextFormat__FramePadding QTextFormat__Property = 16386

//
const QTextFormat__FrameWidth QTextFormat__Property = 16387

//
const QTextFormat__FrameHeight QTextFormat__Property = 16388

//
const QTextFormat__FrameTopMargin QTextFormat__Property = 16389

//
const QTextFormat__FrameBottomMargin QTextFormat__Property = 16390

//
const QTextFormat__FrameLeftMargin QTextFormat__Property = 16391

//
const QTextFormat__FrameRightMargin QTextFormat__Property = 16392

//
const QTextFormat__FrameBorderBrush QTextFormat__Property = 16393

//
const QTextFormat__FrameBorderStyle QTextFormat__Property = 16400

//
const QTextFormat__TableColumns QTextFormat__Property = 16640

//
const QTextFormat__TableColumnWidthConstraints QTextFormat__Property = 16641

//
const QTextFormat__TableCellSpacing QTextFormat__Property = 16642

//
const QTextFormat__TableCellPadding QTextFormat__Property = 16643

//
const QTextFormat__TableHeaderRowCount QTextFormat__Property = 16644

//
const QTextFormat__TableCellRowSpan QTextFormat__Property = 18448

//
const QTextFormat__TableCellColumnSpan QTextFormat__Property = 18449

//
const QTextFormat__TableCellTopPadding QTextFormat__Property = 18450

//
const QTextFormat__TableCellBottomPadding QTextFormat__Property = 18451

//
const QTextFormat__TableCellLeftPadding QTextFormat__Property = 18452

//
const QTextFormat__TableCellRightPadding QTextFormat__Property = 18453

//
const QTextFormat__ImageName QTextFormat__Property = 20480

//
const QTextFormat__ImageWidth QTextFormat__Property = 20496

//
const QTextFormat__ImageHeight QTextFormat__Property = 20497

//
const QTextFormat__ImageQuality QTextFormat__Property = 20500

//
const QTextFormat__FullWidthSelection QTextFormat__Property = 24576

//
const QTextFormat__PageBreakPolicy QTextFormat__Property = 28672

//
const QTextFormat__UserProperty QTextFormat__Property = 1048576

func (this *QTextFormat) PropertyItemName(val int) string {
	switch val {
	case QTextFormat__ObjectIndex: // 0
		return "ObjectIndex"
	case QTextFormat__CssFloat: // 2048
		return "CssFloat"
	case QTextFormat__LayoutDirection: // 2049
		return "LayoutDirection"
	case QTextFormat__OutlinePen: // 2064
		return "OutlinePen"
	case QTextFormat__BackgroundBrush: // 2080
		return "BackgroundBrush"
	case QTextFormat__ForegroundBrush: // 2081
		return "ForegroundBrush"
	case QTextFormat__BackgroundImageUrl: // 2083
		return "BackgroundImageUrl"
	case QTextFormat__BlockAlignment: // 4112
		return "BlockAlignment"
	case QTextFormat__BlockTopMargin: // 4144
		return "BlockTopMargin"
	case QTextFormat__BlockBottomMargin: // 4145
		return "BlockBottomMargin"
	case QTextFormat__BlockLeftMargin: // 4146
		return "BlockLeftMargin"
	case QTextFormat__BlockRightMargin: // 4147
		return "BlockRightMargin"
	case QTextFormat__TextIndent: // 4148
		return "TextIndent"
	case QTextFormat__TabPositions: // 4149
		return "TabPositions"
	case QTextFormat__BlockIndent: // 4160
		return "BlockIndent"
	case QTextFormat__LineHeight: // 4168
		return "LineHeight"
	case QTextFormat__LineHeightType: // 4169
		return "LineHeightType"
	case QTextFormat__BlockNonBreakableLines: // 4176
		return "BlockNonBreakableLines"
	case QTextFormat__BlockTrailingHorizontalRulerWidth: // 4192
		return "BlockTrailingHorizontalRulerWidth"
	case QTextFormat__HeadingLevel: // 4208
		return "HeadingLevel"
	case QTextFormat__FirstFontProperty: // 8160
		return "FirstFontProperty,FontCapitalization"
		// case QTextFormat__FontCapitalization: // 8160
		// return ""
	case QTextFormat__FontLetterSpacingType: // 8243
		return "FontLetterSpacingType"
	case QTextFormat__FontLetterSpacing: // 8161
		return "FontLetterSpacing"
	case QTextFormat__FontWordSpacing: // 8162
		return "FontWordSpacing"
	case QTextFormat__FontStretch: // 8244
		return "FontStretch"
	case QTextFormat__FontStyleHint: // 8163
		return "FontStyleHint"
	case QTextFormat__FontStyleStrategy: // 8164
		return "FontStyleStrategy"
	case QTextFormat__FontKerning: // 8165
		return "FontKerning"
	case QTextFormat__FontHintingPreference: // 8166
		return "FontHintingPreference"
	case QTextFormat__FontFamily: // 8192
		return "FontFamily"
	case QTextFormat__FontPointSize: // 8193
		return "FontPointSize"
	case QTextFormat__FontSizeAdjustment: // 8194
		return "FontSizeAdjustment,FontSizeIncrement"
		// case QTextFormat__FontSizeIncrement: // 8194
		// return ""
	case QTextFormat__FontWeight: // 8195
		return "FontWeight"
	case QTextFormat__FontItalic: // 8196
		return "FontItalic"
	case QTextFormat__FontUnderline: // 8197
		return "FontUnderline"
	case QTextFormat__FontOverline: // 8198
		return "FontOverline"
	case QTextFormat__FontStrikeOut: // 8199
		return "FontStrikeOut"
	case QTextFormat__FontFixedPitch: // 8200
		return "FontFixedPitch"
	case QTextFormat__FontPixelSize: // 8201
		return "FontPixelSize,LastFontProperty"
		// case QTextFormat__LastFontProperty: // 8201
		// return ""
	case QTextFormat__TextUnderlineColor: // 8208
		return "TextUnderlineColor"
	case QTextFormat__TextVerticalAlignment: // 8225
		return "TextVerticalAlignment"
	case QTextFormat__TextOutline: // 8226
		return "TextOutline"
	case QTextFormat__TextUnderlineStyle: // 8227
		return "TextUnderlineStyle"
	case QTextFormat__TextToolTip: // 8228
		return "TextToolTip"
	case QTextFormat__IsAnchor: // 8240
		return "IsAnchor"
	case QTextFormat__AnchorHref: // 8241
		return "AnchorHref"
	case QTextFormat__AnchorName: // 8242
		return "AnchorName"
	case QTextFormat__ObjectType: // 12032
		return "ObjectType"
	case QTextFormat__ListStyle: // 12288
		return "ListStyle"
	case QTextFormat__ListIndent: // 12289
		return "ListIndent"
	case QTextFormat__ListNumberPrefix: // 12290
		return "ListNumberPrefix"
	case QTextFormat__ListNumberSuffix: // 12291
		return "ListNumberSuffix"
	case QTextFormat__FrameBorder: // 16384
		return "FrameBorder"
	case QTextFormat__FrameMargin: // 16385
		return "FrameMargin"
	case QTextFormat__FramePadding: // 16386
		return "FramePadding"
	case QTextFormat__FrameWidth: // 16387
		return "FrameWidth"
	case QTextFormat__FrameHeight: // 16388
		return "FrameHeight"
	case QTextFormat__FrameTopMargin: // 16389
		return "FrameTopMargin"
	case QTextFormat__FrameBottomMargin: // 16390
		return "FrameBottomMargin"
	case QTextFormat__FrameLeftMargin: // 16391
		return "FrameLeftMargin"
	case QTextFormat__FrameRightMargin: // 16392
		return "FrameRightMargin"
	case QTextFormat__FrameBorderBrush: // 16393
		return "FrameBorderBrush"
	case QTextFormat__FrameBorderStyle: // 16400
		return "FrameBorderStyle"
	case QTextFormat__TableColumns: // 16640
		return "TableColumns"
	case QTextFormat__TableColumnWidthConstraints: // 16641
		return "TableColumnWidthConstraints"
	case QTextFormat__TableCellSpacing: // 16642
		return "TableCellSpacing"
	case QTextFormat__TableCellPadding: // 16643
		return "TableCellPadding"
	case QTextFormat__TableHeaderRowCount: // 16644
		return "TableHeaderRowCount"
	case QTextFormat__TableCellRowSpan: // 18448
		return "TableCellRowSpan"
	case QTextFormat__TableCellColumnSpan: // 18449
		return "TableCellColumnSpan"
	case QTextFormat__TableCellTopPadding: // 18450
		return "TableCellTopPadding"
	case QTextFormat__TableCellBottomPadding: // 18451
		return "TableCellBottomPadding"
	case QTextFormat__TableCellLeftPadding: // 18452
		return "TableCellLeftPadding"
	case QTextFormat__TableCellRightPadding: // 18453
		return "TableCellRightPadding"
	case QTextFormat__ImageName: // 20480
		return "ImageName"
	case QTextFormat__ImageWidth: // 20496
		return "ImageWidth"
	case QTextFormat__ImageHeight: // 20497
		return "ImageHeight"
	case QTextFormat__ImageQuality: // 20500
		return "ImageQuality"
	case QTextFormat__FullWidthSelection: // 24576
		return "FullWidthSelection"
	case QTextFormat__PageBreakPolicy: // 28672
		return "PageBreakPolicy"
	case QTextFormat__UserProperty: // 1048576
		return "UserProperty"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextFormat_PropertyItemName(val int) string {
	var nilthis *QTextFormat
	return nilthis.PropertyItemName(val)
}

/*
This enum describes what kind of QTextObject this format is associated with.



See also QTextObject, QTextTable, and QTextObject::format().

*/
type QTextFormat__ObjectTypes = int

//
const QTextFormat__NoObject QTextFormat__ObjectTypes = 0

//
const QTextFormat__ImageObject QTextFormat__ObjectTypes = 1

//
const QTextFormat__TableObject QTextFormat__ObjectTypes = 2

//
const QTextFormat__TableCellObject QTextFormat__ObjectTypes = 3

//
const QTextFormat__UserObject QTextFormat__ObjectTypes = 4096

func (this *QTextFormat) ObjectTypesItemName(val int) string {
	switch val {
	case QTextFormat__NoObject: // 0
		return "NoObject"
	case QTextFormat__ImageObject: // 1
		return "ImageObject"
	case QTextFormat__TableObject: // 2
		return "TableObject"
	case QTextFormat__TableCellObject: // 3
		return "TableCellObject"
	case QTextFormat__UserObject: // 4096
		return "UserObject"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextFormat_ObjectTypesItemName(val int) string {
	var nilthis *QTextFormat
	return nilthis.ObjectTypesItemName(val)
}

/*


 */
type QTextFormat__PageBreakFlag = int

//
const QTextFormat__PageBreak_Auto QTextFormat__PageBreakFlag = 0

//
const QTextFormat__PageBreak_AlwaysBefore QTextFormat__PageBreakFlag = 1

//
const QTextFormat__PageBreak_AlwaysAfter QTextFormat__PageBreakFlag = 16

func (this *QTextFormat) PageBreakFlagItemName(val int) string {
	switch val {
	case QTextFormat__PageBreak_Auto: // 0
		return "PageBreak_Auto"
	case QTextFormat__PageBreak_AlwaysBefore: // 1
		return "PageBreak_AlwaysBefore"
	case QTextFormat__PageBreak_AlwaysAfter: // 16
		return "PageBreak_AlwaysAfter"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextFormat_PageBreakFlagItemName(val int) string {
	var nilthis *QTextFormat
	return nilthis.PageBreakFlagItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10761() {
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
