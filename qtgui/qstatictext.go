package qtgui

// /usr/include/qt/QtGui/qstatictext.h
// #include <qstatictext.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 50
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
type QStaticText struct {
	*qtrt.CObject
}
type QStaticText_ITF interface {
	QStaticText_PTR() *QStaticText
}

func (ptr *QStaticText) QStaticText_PTR() *QStaticText { return ptr }

func (this *QStaticText) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStaticText) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStaticTextFromPointer(cthis unsafe.Pointer) *QStaticText {
	return &QStaticText{&qtrt.CObject{cthis}}
}
func (*QStaticText) NewFromPointer(cthis unsafe.Pointer) *QStaticText {
	return NewQStaticTextFromPointer(cthis)
}

// /usr/include/qt/QtGui/qstatictext.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStaticText()

/*
Constructs an empty QStaticText
*/
func (*QStaticText) NewForInherit() *QStaticText {
	return NewQStaticText()
}
func NewQStaticText() *QStaticText {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStaticText)
	return gothis
}

// /usr/include/qt/QtGui/qstatictext.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStaticText(const QString &)

/*
Constructs an empty QStaticText
*/
func (*QStaticText) NewForInherit_1(text string) *QStaticText {
	return NewQStaticText_1(text)
}
func NewQStaticText_1(text string) *QStaticText {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStaticText)
	return gothis
}

// /usr/include/qt/QtGui/qstatictext.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QStaticText & operator=(QStaticText &&)

/*

 */
func (this *QStaticText) Operator_equal(other unsafe.Pointer /*333*/) *QStaticText {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStaticText)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:70
// index:1
// Public Visibility=Default Availability=Available
// [8] QStaticText & operator=(const QStaticText &)

/*

 */
func (this *QStaticText) Operator_equal_1(arg0 QStaticText_ITF) *QStaticText {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStaticText_PTR() != nil {
		convArg0 = arg0.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStaticTextFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStaticText)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QStaticText()

/*

 */
func DeleteQStaticText(this *QStaticText) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticTextD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qstatictext.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QStaticText &)

/*
Swaps this static text instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QStaticText) Swap(other QStaticText_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QStaticText_PTR() != nil {
		convArg0 = other.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)

/*
Sets the text of the QStaticText to text.

Note: This function will cause the layout of the text to require recalculation.

See also text().
*/
func (this *QStaticText) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text() const

/*
Returns the text of the QStaticText.

See also setText().
*/
func (this *QStaticText) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qstatictext.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextFormat(Qt::TextFormat)

/*
Sets the text format of the QStaticText to textFormat. If textFormat is set to Qt::AutoText (the default), the format of the text will try to be determined using the function Qt::mightBeRichText(). If the text format is Qt::PlainText, then the text will be displayed as is, whereas it will be interpreted as HTML if the format is Qt::RichText. HTML tags that alter the font of the text, its color, or its layout are supported by QStaticText.

Note: This function will cause the layout of the text to require recalculation.

See also textFormat(), setText(), and text().
*/
func (this *QStaticText) SetTextFormat(textFormat int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText13setTextFormatEN2Qt10TextFormatE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), textFormat)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TextFormat textFormat() const

/*
Returns the text format of the QStaticText.

See also setTextFormat(), setText(), and text().
*/
func (this *QStaticText) TextFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText10textFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstatictext.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextWidth(qreal)

/*
Sets the preferred width for this QStaticText. If the text is wider than the specified width, it will be broken into multiple lines and grow vertically. If the text cannot be split into multiple lines, it will be larger than the specified textWidth.

Setting the preferred text width to a negative number will cause the text to be unbounded.

Use size() to get the actual size of the text.

Note: This function will cause the layout of the text to require recalculation.

See also textWidth() and size().
*/
func (this *QStaticText) SetTextWidth(textWidth float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText12setTextWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), textWidth)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal textWidth() const

/*
Returns the preferred width for this QStaticText.

See also setTextWidth().
*/
func (this *QStaticText) TextWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText9textWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qstatictext.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextOption(const QTextOption &)

/*
Sets the text option structure that controls the layout process to the given textOption.

See also textOption().
*/
func (this *QStaticText) SetTextOption(textOption QTextOption_ITF) {
	var convArg0 unsafe.Pointer
	if textOption != nil && textOption.QTextOption_PTR() != nil {
		convArg0 = textOption.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText13setTextOptionERK11QTextOption", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:85
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextOption textOption() const

/*
Returns the current text option used to control the layout process.

See also setTextOption().
*/
func (this *QStaticText) TextOption() *QTextOption /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText10textOptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:87
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size() const

/*
Returns the size of the bounding rect for this QStaticText.

See also textWidth().
*/
func (this *QStaticText) Size() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepare(const QTransform &, const QFont &)

/*
Prepares the QStaticText object for being painted with the given matrix and the given font to avoid overhead when the actual drawStaticText() call is made.

When drawStaticText() is called, the layout of the QStaticText will be recalculated if any part of the QStaticText object has changed since the last time it was drawn. It will also be recalculated if the painter's font is not the same as when the QStaticText was last drawn, or, on any other paint engine than the OpenGL2 engine, if the painter's matrix has been altered since the static text was last drawn.

To avoid the overhead of creating the layout the first time you draw the QStaticText after making changes, you can use the prepare() function and pass in the matrix and font you expect to use when drawing the text.

See also QPainter::setFont() and QPainter::setMatrix().
*/
func (this *QStaticText) Prepare(matrix QTransform_ITF, font QFont_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg1 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepare(const QTransform &, const QFont &)

/*
Prepares the QStaticText object for being painted with the given matrix and the given font to avoid overhead when the actual drawStaticText() call is made.

When drawStaticText() is called, the layout of the QStaticText will be recalculated if any part of the QStaticText object has changed since the last time it was drawn. It will also be recalculated if the painter's font is not the same as when the QStaticText was last drawn, or, on any other paint engine than the OpenGL2 engine, if the painter's matrix has been altered since the static text was last drawn.

To avoid the overhead of creating the layout the first time you draw the QStaticText after making changes, you can use the prepare() function and pass in the matrix and font you expect to use when drawing the text.

See also QPainter::setFont() and QPainter::setMatrix().
*/
func (this *QStaticText) Prepare__() {
	// arg: 0, const QTransform &=LValueReference, QTransform=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, const QFont &=LValueReference, QFont=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepare(const QTransform &, const QFont &)

/*
Prepares the QStaticText object for being painted with the given matrix and the given font to avoid overhead when the actual drawStaticText() call is made.

When drawStaticText() is called, the layout of the QStaticText will be recalculated if any part of the QStaticText object has changed since the last time it was drawn. It will also be recalculated if the painter's font is not the same as when the QStaticText was last drawn, or, on any other paint engine than the OpenGL2 engine, if the painter's matrix has been altered since the static text was last drawn.

To avoid the overhead of creating the layout the first time you draw the QStaticText after making changes, you can use the prepare() function and pass in the matrix and font you expect to use when drawing the text.

See also QPainter::setFont() and QPainter::setMatrix().
*/
func (this *QStaticText) Prepare__1(matrix QTransform_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QTransform_PTR() != nil {
		convArg0 = matrix.QTransform_PTR().GetCthis()
	}
	// arg: 1, const QFont &=LValueReference, QFont=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText7prepareERK10QTransformRK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPerformanceHint(QStaticText::PerformanceHint)

/*
Sets the performance hint of the QStaticText according to the performanceHint provided. The performanceHint is used to customize how much caching is done internally to improve performance.

The default is QStaticText::ModerateCaching.

Note: This function will cause the layout of the text to require recalculation.

See also performanceHint().
*/
func (this *QStaticText) SetPerformanceHint(performanceHint int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStaticText18setPerformanceHintENS_15PerformanceHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), performanceHint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] QStaticText::PerformanceHint performanceHint() const

/*
Returns which performance hint is set for the QStaticText.

See also setPerformanceHint().
*/
func (this *QStaticText) PerformanceHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticText15performanceHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstatictext.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QStaticText &) const

/*

 */
func (this *QStaticText) Operator_equal_equal(arg0 QStaticText_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStaticText_PTR() != nil {
		convArg0 = arg0.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticTexteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstatictext.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QStaticText &) const

/*

 */
func (this *QStaticText) Operator_not_equal(arg0 QStaticText_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStaticText_PTR() != nil {
		convArg0 = arg0.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStaticTextneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum the different performance hints that can be set on the QStaticText. These hints can be used to indicate that the QStaticText should use additional caches, if possible, to improve performance at the expense of memory. In particular, setting the performance hint AggressiveCaching on the QStaticText will improve performance when using the OpenGL graphics system or when drawing to a QOpenGLWidget.


*/
type QStaticText__PerformanceHint = int

// Do basic caching for high performance at a low memory cost.
const QStaticText__ModerateCaching QStaticText__PerformanceHint = 0

// Use additional caching when available. This may improve performance at a higher memory cost.
const QStaticText__AggressiveCaching QStaticText__PerformanceHint = 1

func (this *QStaticText) PerformanceHintItemName(val int) string {
	switch val {
	case QStaticText__ModerateCaching: // 0
		return "ModerateCaching"
	case QStaticText__AggressiveCaching: // 1
		return "AggressiveCaching"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QStaticText_PerformanceHintItemName(val int) string {
	var nilthis *QStaticText
	return nilthis.PerformanceHintItemName(val)
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
