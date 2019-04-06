package qtgui

// /usr/include/qt/QtGui/qfontmetrics.h
// #include <qfontmetrics.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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
type QFontMetricsF struct {
	*qtrt.CObject
}
type QFontMetricsF_ITF interface {
	QFontMetricsF_PTR() *QFontMetricsF
}

func (ptr *QFontMetricsF) QFontMetricsF_PTR() *QFontMetricsF { return ptr }

func (this *QFontMetricsF) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontMetricsF) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontMetricsFFromPointer(cthis unsafe.Pointer) *QFontMetricsF {
	return &QFontMetricsF{&qtrt.CObject{cthis}}
}
func (*QFontMetricsF) NewFromPointer(cthis unsafe.Pointer) *QFontMetricsF {
	return NewQFontMetricsFFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontmetrics.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontMetricsF(const QFont &)

/*

 */
func (*QFontMetricsF) NewForInherit(arg0 QFont_ITF) *QFontMetricsF {
	return NewQFontMetricsF(arg0)
}
func NewQFontMetricsF(arg0 QFont_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFont", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetricsF)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:141
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFontMetricsF(const QFont &, QPaintDevice *)

/*

 */
func (*QFontMetricsF) NewForInherit1(arg0 QFont_ITF, pd QPaintDevice_ITF /*777 QPaintDevice **/) *QFontMetricsF {
	return NewQFontMetricsF1(arg0, pd)
}
func NewQFontMetricsF1(arg0 QFont_ITF, pd QPaintDevice_ITF /*777 QPaintDevice **/) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pd != nil && pd.QPaintDevice_PTR() != nil {
		convArg1 = pd.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK5QFontP12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetricsF)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:142
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QFontMetricsF(const QFontMetrics &)

/*

 */
func (*QFontMetricsF) NewForInherit2(arg0 QFontMetrics_ITF) *QFontMetricsF {
	return NewQFontMetricsF2(arg0)
}
func NewQFontMetricsF2(arg0 QFontMetrics_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontMetrics_PTR() != nil {
		convArg0 = arg0.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFC2ERK12QFontMetrics", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontMetricsF)
	return gothis
}

// /usr/include/qt/QtGui/qfontmetrics.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFontMetricsF()

/*

 */
func DeleteQFontMetricsF(this *QFontMetricsF) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qfontmetrics.h:146
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontMetricsF & operator=(const QFontMetricsF &)

/*

 */
func (this *QFontMetricsF) Operator_equal(arg0 QFontMetricsF_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontMetricsF_PTR() != nil {
		convArg0 = arg0.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetricsF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:147
// index:1
// Public Visibility=Default Availability=Available
// [8] QFontMetricsF & operator=(const QFontMetrics &)

/*

 */
func (this *QFontMetricsF) Operator_equal1(arg0 QFontMetrics_ITF) *QFontMetricsF {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFontMetrics_PTR() != nil {
		convArg0 = arg0.QFontMetrics_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFaSERK12QFontMetrics", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetricsF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:149
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QFontMetricsF & operator=(QFontMetricsF &&)

/*

 */
func (this *QFontMetricsF) Operator_equal2(other unsafe.Pointer /*333*/) *QFontMetricsF {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsFaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontMetricsFFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFontMetricsF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFontMetricsF &)

/*
Swaps this font metrics instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QFontMetricsF) Swap(other QFontMetricsF_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetricsF_PTR() != nil {
		convArg0 = other.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontMetricsF4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontmetrics.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent() const

/*
Returns the ascent of the font.

The ascent of a font is the distance from the baseline to the highest position characters extend to. In practice, some font designers break this rule, e.g. when they put more than one accent on top of a character, or to accommodate an unusual character in an exotic language, so it is possible (though rare) that this value will be too small.

See also descent().
*/
func (this *QFontMetricsF) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal capHeight() const

/*
Returns the cap height of the font.

The cap height of a font is the height of a capital letter above the baseline. It specifically is the height of capital letters that are flat - such as H or I - as opposed to round letters such as O, or pointed letters like A, both of which may display overshoot.

This function was introduced in  Qt 5.8.

See also ascent().
*/
func (this *QFontMetricsF) CapHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF9capHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent() const

/*
Returns the descent of the font.

The descent is the distance from the base line to the lowest point characters extend to. In practice, some font designers break this rule, e.g. to accommodate an unusual character in an exotic language, so it is possible (though rare) that this value will be too small.

See also ascent().
*/
func (this *QFontMetricsF) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal height() const

/*
Returns the height of the font.

This is always equal to ascent()+descent().

See also leading() and lineSpacing().
*/
func (this *QFontMetricsF) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:159
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leading() const

/*
Returns the leading of the font.

This is the natural inter-line spacing.

See also height() and lineSpacing().
*/
func (this *QFontMetricsF) Leading() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF7leadingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lineSpacing() const

/*
Returns the distance from one base line to the next.

This value is always equal to leading()+height().

See also height() and leading().
*/
func (this *QFontMetricsF) LineSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF11lineSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minLeftBearing() const

/*
Returns the minimum left bearing of the font.

This is the smallest leftBearing(char) of all characters in the font.

Note that this function can be very slow if the font is large.

See also minRightBearing() and leftBearing().
*/
func (this *QFontMetricsF) MinLeftBearing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF14minLeftBearingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal minRightBearing() const

/*
Returns the minimum right bearing of the font.

This is the smallest rightBearing(char) of all characters in the font.

Note that this function can be very slow if the font is large.

See also minLeftBearing() and rightBearing().
*/
func (this *QFontMetricsF) MinRightBearing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF15minRightBearingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:163
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maxWidth() const

/*
Returns the width of the widest character in the font.
*/
func (this *QFontMetricsF) MaxWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF8maxWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:165
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal xHeight() const

/*
Returns the 'x' height of the font. This is often but not always the same as the height of the character 'x'.
*/
func (this *QFontMetricsF) XHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF7xHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:166
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal averageCharWidth() const

/*
Returns the average width of glyphs in the font.

This function was introduced in  Qt 4.2.
*/
func (this *QFontMetricsF) AverageCharWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF16averageCharWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:168
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFont(QChar) const

/*
Returns true if character ch is a valid character in the font; otherwise returns false.
*/
func (this *QFontMetricsF) InFont(arg0 qtcore.QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF6inFontE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:169
// index:0
// Public Visibility=Default Availability=Available
// [1] bool inFontUcs4(uint) const

/*
Returns true if the character ucs4 encoded in UCS-4/UTF-32 is a valid character in the font; otherwise returns false.
*/
func (this *QFontMetricsF) InFontUcs4(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF10inFontUcs4Ej", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leftBearing(QChar) const

/*
Returns the left bearing of character ch in the font.

The left bearing is the right-ward distance of the left-most pixel of the character from the logical origin of the character. This value is negative if the pixels of the character extend to the left of the logical origin.

See width() for a graphical description of this metric.

See also rightBearing(), minLeftBearing(), and width().
*/
func (this *QFontMetricsF) LeftBearing(arg0 qtcore.QChar_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF11leftBearingE5QChar", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rightBearing(QChar) const

/*
Returns the right bearing of character ch in the font.

The right bearing is the left-ward distance of the right-most pixel of the character from the logical origin of a subsequent character. This value is negative if the pixels of the character extend to the right of the width() of the character.

See width() for a graphical description of this metric.

See also leftBearing(), minRightBearing(), and width().
*/
func (this *QFontMetricsF) RightBearing(arg0 qtcore.QChar_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12rightBearingE5QChar", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width(const QString &) const

/*

 */
func (this *QFontMetricsF) Width(string string) float64 {
	var tmpArg0 = qtcore.NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthERK7QString", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:176
// index:1
// Public Visibility=Default Availability=Available
// [8] qreal width(QChar) const

/*

 */
func (this *QFontMetricsF) Width1(arg0 qtcore.QChar_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF5widthE5QChar", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal horizontalAdvance(const QString &, int) const

/*
Returns the horizontal advance in pixels of the first len characters of text. If len is negative (the default), the entire string is used.

This is the distance appropriate for drawing a subsequent character after text.

This function was introduced in  Qt 5.11.

See also boundingRect().
*/
func (this *QFontMetricsF) HorizontalAdvance(string string, length int) float64 {
	var tmpArg0 = qtcore.NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF17horizontalAdvanceERK7QStringi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, length)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal horizontalAdvance(const QString &, int) const

/*
Returns the horizontal advance in pixels of the first len characters of text. If len is negative (the default), the entire string is used.

This is the distance appropriate for drawing a subsequent character after text.

This function was introduced in  Qt 5.11.

See also boundingRect().
*/
func (this *QFontMetricsF) HorizontalAdvancep(string string) float64 {
	var tmpArg0 = qtcore.NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	length := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF17horizontalAdvanceERK7QStringi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0, length)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:180
// index:1
// Public Visibility=Default Availability=Available
// [8] qreal horizontalAdvance(QChar) const

/*
Returns the horizontal advance in pixels of the first len characters of text. If len is negative (the default), the entire string is used.

This is the distance appropriate for drawing a subsequent character after text.

This function was introduced in  Qt 5.11.

See also boundingRect().
*/
func (this *QFontMetricsF) HorizontalAdvance1(arg0 qtcore.QChar_ITF /*123*/) float64 {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF17horizontalAdvanceE5QChar", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:182
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QString &) const

/*
Returns the rectangle that is covered by ink if character ch were to be drawn at the origin of the coordinate system.

Note that the bounding rectangle may extend to the left of (0, 0) (e.g., for italicized fonts), and that the text output may cover all pixels in the bounding rectangle. For a space character the rectangle will usually be empty.

Note that the rectangle usually extends both above and below the base line.

Warning: The width of the returned rectangle is not the advance width of the character. Use boundingRect(const QString &) or horizontalAdvance() instead.

See also width().
*/
func (this *QFontMetricsF) BoundingRect(string string) *qtcore.QRectF /*123*/ {
	var tmpArg0 = qtcore.NewQString5(string)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:183
// index:1
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(QChar) const

/*
Returns the rectangle that is covered by ink if character ch were to be drawn at the origin of the coordinate system.

Note that the bounding rectangle may extend to the left of (0, 0) (e.g., for italicized fonts), and that the text output may cover all pixels in the bounding rectangle. For a space character the rectangle will usually be empty.

Note that the rectangle usually extends both above and below the base line.

Warning: The width of the returned rectangle is not the advance width of the character. Use boundingRect(const QString &) or horizontalAdvance() instead.

See also width().
*/
func (this *QFontMetricsF) BoundingRect1(arg0 qtcore.QChar_ITF /*123*/) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:184
// index:2
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &, int, int *) const

/*
Returns the rectangle that is covered by ink if character ch were to be drawn at the origin of the coordinate system.

Note that the bounding rectangle may extend to the left of (0, 0) (e.g., for italicized fonts), and that the text output may cover all pixels in the bounding rectangle. For a space character the rectangle will usually be empty.

Note that the rectangle usually extends both above and below the base line.

Warning: The width of the returned rectangle is not the advance width of the character. Use boundingRect(const QString &) or horizontalAdvance() instead.

See also width().
*/
func (this *QFontMetricsF) BoundingRect2(r qtcore.QRectF_ITF, flags int, string string, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(string)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:184
// index:2
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &, int, int *) const

/*
Returns the rectangle that is covered by ink if character ch were to be drawn at the origin of the coordinate system.

Note that the bounding rectangle may extend to the left of (0, 0) (e.g., for italicized fonts), and that the text output may cover all pixels in the bounding rectangle. For a space character the rectangle will usually be empty.

Note that the rectangle usually extends both above and below the base line.

Warning: The width of the returned rectangle is not the advance width of the character. Use boundingRect(const QString &) or horizontalAdvance() instead.

See also width().
*/
func (this *QFontMetricsF) BoundingRect2p(r qtcore.QRectF_ITF, flags int, string string) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(string)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, int=Int, =Invalid, , Invalid
	tabstops := int(0)
	// arg: 4, int *=Pointer, =Invalid, , Invalid
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:184
// index:2
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(const QRectF &, int, const QString &, int, int *) const

/*
Returns the rectangle that is covered by ink if character ch were to be drawn at the origin of the coordinate system.

Note that the bounding rectangle may extend to the left of (0, 0) (e.g., for italicized fonts), and that the text output may cover all pixels in the bounding rectangle. For a space character the rectangle will usually be empty.

Note that the rectangle usually extends both above and below the base line.

Warning: The width of the returned rectangle is not the advance width of the character. Use boundingRect(const QString &) or horizontalAdvance() instead.

See also width().
*/
func (this *QFontMetricsF) BoundingRect2p1(r qtcore.QRectF_ITF, flags int, string string, tabstops int) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if r != nil && r.QRectF_PTR() != nil {
		convArg0 = r.QRectF_PTR().GetCthis()
	}
	var tmpArg2 = qtcore.NewQString5(string)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 4, int *=Pointer, =Invalid, , Invalid
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, flags, convArg2, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:185
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(int, const QString &, int, int *) const

/*
Returns the size in pixels of text.

The flags argument is the bitwise OR of the following flags:


Qt::TextSingleLine ignores newline characters.
Qt::TextExpandTabs expands tabs (see below)
Qt::TextShowMnemonic interprets "&x" as x; i.e., underlined.
Qt::TextWordWrap breaks the text to fit the rectangle.


If Qt::TextExpandTabs is set in flags, then: if tabArray is non-null, it specifies a 0-terminated sequence of pixel-positions for tabs; otherwise if tabStops is non-zero, it is used as the tab spacing (in pixels).

Newline characters are processed as linebreaks.

Despite the different actual character heights, the heights of the bounding rectangles of "Yes" and "yes" are the same.

See also boundingRect().
*/
func (this *QFontMetricsF) Size(flags int, str string, tabstops int, tabarray unsafe.Pointer /*666*/) *qtcore.QSizeF /*123*/ {
	var tmpArg1 = qtcore.NewQString5(str)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:185
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(int, const QString &, int, int *) const

/*
Returns the size in pixels of text.

The flags argument is the bitwise OR of the following flags:


Qt::TextSingleLine ignores newline characters.
Qt::TextExpandTabs expands tabs (see below)
Qt::TextShowMnemonic interprets "&x" as x; i.e., underlined.
Qt::TextWordWrap breaks the text to fit the rectangle.


If Qt::TextExpandTabs is set in flags, then: if tabArray is non-null, it specifies a 0-terminated sequence of pixel-positions for tabs; otherwise if tabStops is non-zero, it is used as the tab spacing (in pixels).

Newline characters are processed as linebreaks.

Despite the different actual character heights, the heights of the bounding rectangles of "Yes" and "yes" are the same.

See also boundingRect().
*/
func (this *QFontMetricsF) Sizep(flags int, str string) *qtcore.QSizeF /*123*/ {
	var tmpArg1 = qtcore.NewQString5(str)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, int=Int, =Invalid, , Invalid
	tabstops := int(0)
	// arg: 3, int *=Pointer, =Invalid, , Invalid
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:185
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size(int, const QString &, int, int *) const

/*
Returns the size in pixels of text.

The flags argument is the bitwise OR of the following flags:


Qt::TextSingleLine ignores newline characters.
Qt::TextExpandTabs expands tabs (see below)
Qt::TextShowMnemonic interprets "&x" as x; i.e., underlined.
Qt::TextWordWrap breaks the text to fit the rectangle.


If Qt::TextExpandTabs is set in flags, then: if tabArray is non-null, it specifies a 0-terminated sequence of pixel-positions for tabs; otherwise if tabStops is non-zero, it is used as the tab spacing (in pixels).

Newline characters are processed as linebreaks.

Despite the different actual character heights, the heights of the bounding rectangles of "Yes" and "yes" are the same.

See also boundingRect().
*/
func (this *QFontMetricsF) Sizep1(flags int, str string, tabstops int) *qtcore.QSizeF /*123*/ {
	var tmpArg1 = qtcore.NewQString5(str)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 3, int *=Pointer, =Invalid, , Invalid
	var tabarray unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF4sizeEiRK7QStringiPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, convArg1, tabstops, tabarray)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:187
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF tightBoundingRect(const QString &) const

/*
Returns a tight bounding rectangle around the characters in the string specified by text. The bounding rectangle always covers at least the set of pixels the text would cover if drawn at (0, 0).

Note that the bounding rectangle may extend to the left of (0, 0), e.g. for italicized fonts, and that the width of the returned rectangle might be different than what the width() method returns.

If you want to know the advance width of the string (to lay out a set of strings next to each other), use horizontalAdvance() instead.

Newline characters are processed as normal characters, not as linebreaks.

Warning: Calling this method is very slow on Windows.

This function was introduced in  Qt 4.3.

See also width(), height(), and boundingRect().
*/
func (this *QFontMetricsF) TightBoundingRect(text string) *qtcore.QRectF /*123*/ {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF17tightBoundingRectERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qfontmetrics.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elidedText(const QString &, Qt::TextElideMode, qreal, int) const

/*
If the string text is wider than width, returns an elided version of the string (i.e., a string with "..." in it). Otherwise, returns the original string.

The mode parameter specifies whether the text is elided on the left (e.g., "...tech"), in the middle (e.g., "Tr...ch"), or on the right (e.g., "Trol...").

The width is specified in pixels, not characters.

The flags argument is optional and currently only supports Qt::TextShowMnemonic as value.

The elide mark follows the layoutdirection. For example, it will be on the right side of the text for right-to-left layouts if the mode is Qt::ElideLeft, and on the left side of the text if the mode is Qt::ElideRight.

This function was introduced in  Qt 4.2.
*/
func (this *QFontMetricsF) ElidedText(text string, mode int, width float64, flags int) string {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF10elidedTextERK7QStringN2Qt13TextElideModeEdi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, width, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontmetrics.h:189
// index:0
// Public Visibility=Default Availability=Available
// [8] QString elidedText(const QString &, Qt::TextElideMode, qreal, int) const

/*
If the string text is wider than width, returns an elided version of the string (i.e., a string with "..." in it). Otherwise, returns the original string.

The mode parameter specifies whether the text is elided on the left (e.g., "...tech"), in the middle (e.g., "Tr...ch"), or on the right (e.g., "Trol...").

The width is specified in pixels, not characters.

The flags argument is optional and currently only supports Qt::TextShowMnemonic as value.

The elide mark follows the layoutdirection. For example, it will be on the right side of the text for right-to-left layouts if the mode is Qt::ElideLeft, and on the left side of the text if the mode is Qt::ElideRight.

This function was introduced in  Qt 4.2.
*/
func (this *QFontMetricsF) ElidedTextp(text string, mode int, width float64) string {
	var tmpArg0 = qtcore.NewQString5(text)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, int=Int, =Invalid, , Invalid
	flags := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF10elidedTextERK7QStringN2Qt13TextElideModeEdi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, width, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontmetrics.h:191
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal underlinePos() const

/*
Returns the distance from the base line to where an underscore should be drawn.

See also overlinePos(), strikeOutPos(), and lineWidth().
*/
func (this *QFontMetricsF) UnderlinePos() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12underlinePosEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal overlinePos() const

/*
Returns the distance from the base line to where an overline should be drawn.

See also underlinePos(), strikeOutPos(), and lineWidth().
*/
func (this *QFontMetricsF) OverlinePos() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF11overlinePosEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:193
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal strikeOutPos() const

/*
Returns the distance from the base line to where the strikeout line should be drawn.

See also underlinePos(), overlinePos(), and lineWidth().
*/
func (this *QFontMetricsF) StrikeOutPos() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF12strikeOutPosEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:194
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lineWidth() const

/*
Returns the width of the underline and strikeout lines, adjusted for the point size of the font.

See also underlinePos(), overlinePos(), and strikeOutPos().
*/
func (this *QFontMetricsF) LineWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsF9lineWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfontmetrics.h:196
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QFontMetricsF &) const

/*

 */
func (this *QFontMetricsF) Operator_equal_equal(other QFontMetricsF_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetricsF_PTR() != nil {
		convArg0 = other.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsFeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontmetrics.h:197
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QFontMetricsF &) const

/*

 */
func (this *QFontMetricsF) Operator_not_equal(other QFontMetricsF_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFontMetricsF_PTR() != nil {
		convArg0 = other.QFontMetricsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontMetricsFneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

//  body block end

//  keep block begin

func init_unused_10867() {
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
