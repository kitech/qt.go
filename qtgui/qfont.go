package qtgui

// /usr/include/qt/QtGui/qfont.h
// #include <qfont.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QFont struct {
	*qtrt.CObject
}
type QFont_ITF interface {
	QFont_PTR() *QFont
}

func (ptr *QFont) QFont_PTR() *QFont { return ptr }

func (this *QFont) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFont) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontFromPointer(cthis unsafe.Pointer) *QFont {
	return &QFont{&qtrt.CObject{cthis}}
}
func (*QFont) NewFromPointer(cthis unsafe.Pointer) *QFont {
	return NewQFontFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfont.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFont()

/*
Constructs a font object that uses the application's default font.

See also QGuiApplication::setFont() and QGuiApplication::font().
*/
func NewQFont() *QFont {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:171
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
Constructs a font object that uses the application's default font.

See also QGuiApplication::setFont() and QGuiApplication::font().
*/
func NewQFont_1(family string, pointSize int, weight int, italic bool) *QFont {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontC2ERK7QStringiib", qtrt.FFI_TYPE_POINTER, convArg0, pointSize, weight, italic)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:171
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
Constructs a font object that uses the application's default font.

See also QGuiApplication::setFont() and QGuiApplication::font().
*/
func NewQFont_1_(family string) *QFont {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	pointSize := int(-1)
	// arg: 2, int=Int, =Invalid, , Invalid
	weight := int(-1)
	// arg: 3, bool=Bool, =Invalid, , Invalid
	italic := false
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontC2ERK7QStringiib", qtrt.FFI_TYPE_POINTER, convArg0, pointSize, weight, italic)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:171
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
Constructs a font object that uses the application's default font.

See also QGuiApplication::setFont() and QGuiApplication::font().
*/
func NewQFont_1_1(family string, pointSize int) *QFont {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, int=Int, =Invalid, , Invalid
	weight := int(-1)
	// arg: 3, bool=Bool, =Invalid, , Invalid
	italic := false
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontC2ERK7QStringiib", qtrt.FFI_TYPE_POINTER, convArg0, pointSize, weight, italic)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:171
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
Constructs a font object that uses the application's default font.

See also QGuiApplication::setFont() and QGuiApplication::font().
*/
func NewQFont_1_2(family string, pointSize int, weight int) *QFont {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, bool=Bool, =Invalid, , Invalid
	italic := false
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontC2ERK7QStringiib", qtrt.FFI_TYPE_POINTER, convArg0, pointSize, weight, italic)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:172
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QFont &, QPaintDevice *)

/*
Constructs a font object that uses the application's default font.

See also QGuiApplication::setFont() and QGuiApplication::font().
*/
func NewQFont_2(arg0 QFont_ITF, pd QPaintDevice_ITF /*777 QPaintDevice **/) *QFont {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if pd != nil && pd.QPaintDevice_PTR() != nil {
		convArg1 = pd.QPaintDevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontC2ERKS_P12QPaintDevice", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QFont()

/*

 */
func DeleteQFont(this *QFont) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qfont.h:176
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QFont &)

/*
Swaps this font instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QFont) Swap(other QFont_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QFont_PTR() != nil {
		convArg0 = other.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QString family() const

/*
Returns the requested font family name, i.e. the name set in the constructor or the last setFont() call.

See also setFamily(), substitutes(), and substitute().
*/
func (this *QFont) Family() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont6familyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:180
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFamily(const QString &)

/*
Sets the family name of the font. The name is case insensitive and may include a foundry name.

The family name may optionally also include a foundry name, e.g. "Helvetica [Cronyx]". If the family is available from more than one foundry and the foundry isn't specified, an arbitrary foundry is chosen. If the family isn't available a family will be set using the font matching algorithm.

See also family(), setStyleHint(), and QFontInfo.
*/
func (this *QFont) SetFamily(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont9setFamilyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleName() const

/*
Returns the requested font style name. This can be used to match the font with irregular styles (that can't be normalized in other style properties).

This function was introduced in  Qt 4.8.

See also setStyleName(), setFamily(), and setStyle().
*/
func (this *QFont) StyleName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9styleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleName(const QString &)

/*
Sets the style name of the font to styleName. When set, other style properties like style() and weight() will be ignored for font matching, though they may be simulated afterwards if supported by the platform's font engine.

Due to the lower quality of artificially simulated styles, and the lack of full cross platform support, it is not recommended to use matching by style name together with matching by style properties

This function was introduced in  Qt 4.8.

See also styleName().
*/
func (this *QFont) SetStyleName(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setStyleNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:185
// index:0
// Public Visibility=Default Availability=Available
// [4] int pointSize() const

/*
Returns the point size of the font. Returns -1 if the font size was specified in pixels.

See also setPointSize() and pointSizeF().
*/
func (this *QFont) PointSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9pointSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPointSize(int)

/*
Sets the point size to pointSize. The point size must be greater than zero.

See also pointSize() and setPointSizeF().
*/
func (this *QFont) SetPointSize(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setPointSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal pointSizeF() const

/*
Returns the point size of the font. Returns -1 if the font size was specified in pixels.

See also pointSize(), setPointSizeF(), pixelSize(), QFontInfo::pointSize(), and QFontInfo::pixelSize().
*/
func (this *QFont) PointSizeF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont10pointSizeFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPointSizeF(qreal)

/*
Sets the point size to pointSize. The point size must be greater than zero. The requested precision may not be achieved on all platforms.

See also pointSizeF(), setPointSize(), and setPixelSize().
*/
func (this *QFont) SetPointSizeF(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont13setPointSizeFEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:190
// index:0
// Public Visibility=Default Availability=Available
// [4] int pixelSize() const

/*
Returns the pixel size of the font if it was set with setPixelSize(). Returns -1 if the size was set with setPointSize() or setPointSizeF().

See also setPixelSize(), pointSize(), QFontInfo::pointSize(), and QFontInfo::pixelSize().
*/
func (this *QFont) PixelSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9pixelSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelSize(int)

/*
Sets the font size to pixelSize pixels.

Using this function makes the font device dependent. Use setPointSize() or setPointSizeF() to set the size of the font in a device independent manner.

See also pixelSize().
*/
func (this *QFont) SetPixelSize(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setPixelSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:193
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight() const

/*
Returns the weight of the font, using the same scale as the QFont::Weight enumeration.

See also setWeight(), Weight, and QFontInfo.
*/
func (this *QFont) Weight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont6weightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWeight(int)

/*
Sets the weight of the font to weight, using the scale defined by QFont::Weight enumeration.

Note: If styleName() is set, this value may be ignored for font selection.

See also weight() and QFontInfo.
*/
func (this *QFont) SetWeight(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont9setWeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool bold() const

/*
Returns true if weight() is a value greater than QFont::Medium; otherwise returns false.

See also weight(), setBold(), and QFontInfo::bold().
*/
func (this *QFont) Bold() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont4boldEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:197
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBold(bool)

/*
If enable is true sets the font's weight to QFont::Bold; otherwise sets the weight to QFont::Normal.

For finer boldness control use setWeight().

Note: If styleName() is set, this value may be ignored, or if supported on the platform, the font artificially embolded.

See also bold() and setWeight().
*/
func (this *QFont) SetBold(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont7setBoldEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(QFont::Style)

/*
Sets the style of the font to style.

See also style(), italic(), and QFontInfo.
*/
func (this *QFont) SetStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont8setStyleENS_5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:200
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Style style() const

/*
Returns the style of the font.

See also setStyle().
*/
func (this *QFont) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool italic() const

/*
Returns true if the style() of the font is not QFont::StyleNormal

See also setItalic() and style().
*/
func (this *QFont) Italic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont6italicEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setItalic(bool)

/*
Sets the style() of the font to QFont::StyleItalic if enable is true; otherwise the style is set to QFont::StyleNormal.

Note: If styleName() is set, this value may be ignored, or if supported on the platform, the font may be rendered tilted instead of picking a designed italic font-variant.

See also italic() and QFontInfo.
*/
func (this *QFont) SetItalic(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont9setItalicEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:205
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underline() const

/*
Returns true if underline has been set; otherwise returns false.

See also setUnderline().
*/
func (this *QFont) Underline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9underlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnderline(bool)

/*
If enable is true, sets underline on; otherwise sets underline off.

See also underline() and QFontInfo.
*/
func (this *QFont) SetUnderline(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setUnderlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:208
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overline() const

/*
Returns true if overline has been set; otherwise returns false.

See also setOverline().
*/
func (this *QFont) Overline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont8overlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverline(bool)

/*
If enable is true, sets overline on; otherwise sets overline off.

See also overline() and QFontInfo.
*/
func (this *QFont) SetOverline(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont11setOverlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strikeOut() const

/*
Returns true if strikeout has been set; otherwise returns false.

See also setStrikeOut().
*/
func (this *QFont) StrikeOut() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9strikeOutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStrikeOut(bool)

/*
If enable is true, sets strikeout on; otherwise sets strikeout off.

See also strikeOut() and QFontInfo.
*/
func (this *QFont) SetStrikeOut(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setStrikeOutEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:214
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fixedPitch() const

/*
Returns true if fixed pitch has been set; otherwise returns false.

See also setFixedPitch() and QFontInfo::fixedPitch().
*/
func (this *QFont) FixedPitch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont10fixedPitchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedPitch(bool)

/*
If enable is true, sets fixed pitch on; otherwise sets fixed pitch off.

See also fixedPitch() and QFontInfo.
*/
func (this *QFont) SetFixedPitch(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont13setFixedPitchEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:217
// index:0
// Public Visibility=Default Availability=Available
// [1] bool kerning() const

/*
Returns true if kerning should be used when drawing text with this font.

See also setKerning().
*/
func (this *QFont) Kerning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7kerningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKerning(bool)

/*
Enables kerning for this font if enable is true; otherwise disables it. By default, kerning is enabled.

When kerning is enabled, glyph metrics do not add up anymore, even for Latin text. In other words, the assumption that width('a') + width('b') is equal to width("ab") is not necessarily true.

See also kerning() and QFontMetrics.
*/
func (this *QFont) SetKerning(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setKerningEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:220
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::StyleHint styleHint() const

/*
Returns the StyleHint.

The style hint affects the font matching algorithm. See QFont::StyleHint for the list of available hints.

See also setStyleHint(), QFont::StyleStrategy, and QFontInfo::styleHint().
*/
func (this *QFont) StyleHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9styleHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:221
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::StyleStrategy styleStrategy() const

/*
Returns the StyleStrategy.

The style strategy affects the font matching algorithm. See QFont::StyleStrategy for the list of available strategies.

See also setStyleStrategy(), setStyleHint(), and QFont::StyleHint.
*/
func (this *QFont) StyleStrategy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont13styleStrategyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleHint(QFont::StyleHint, QFont::StyleStrategy)

/*
Sets the style hint and strategy to hint and strategy, respectively.

If these aren't set explicitly the style hint will default to AnyStyle and the style strategy to PreferDefault.

Qt does not support style hints on X11 since this information is not provided by the window system.

See also StyleHint, styleHint(), StyleStrategy, styleStrategy(), and QFontInfo.
*/
func (this *QFont) SetStyleHint(arg0 int, arg1 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setStyleHintENS_9StyleHintENS_13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleHint(QFont::StyleHint, QFont::StyleStrategy)

/*
Sets the style hint and strategy to hint and strategy, respectively.

If these aren't set explicitly the style hint will default to AnyStyle and the style strategy to PreferDefault.

Qt does not support style hints on X11 since this information is not provided by the window system.

See also StyleHint, styleHint(), StyleStrategy, styleStrategy(), and QFontInfo.
*/
func (this *QFont) SetStyleHint__(arg0 int) {
	// arg: 1, QFont::StyleStrategy=Enum, QFont::StyleStrategy=Enum, , Invalid
	arg1 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setStyleHintENS_9StyleHintENS_13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleStrategy(QFont::StyleStrategy)

/*
Sets the style strategy for the font to s.

See also styleStrategy() and QFont::StyleStrategy.
*/
func (this *QFont) SetStyleStrategy(s int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont16setStyleStrategyENS_13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretch() const

/*
Returns the stretch factor for the font.

See also setStretch().
*/
func (this *QFont) Stretch() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7stretchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretch(int)

/*
Sets the stretch factor for the font.

The stretch factor matches a condensed or expanded version of the font or applies a stretch transform that changes the width of all characters in the font by factor percent. For example, setting factor to 150 results in all characters in the font being 1.5 times (ie. 150%) wider. The minimum stretch factor is 1, and the maximum stretch factor is 4000. The default stretch factor is AnyStretch, which will accept any stretch factor and not apply any transform on the font.

The stretch factor is only applied to outline fonts. The stretch factor is ignored for bitmap fonts.

Note: When matching a font with a native non-default stretch factor, requesting a stretch of 100 will stretch it back to a medium width font.

See also stretch() and QFont::Stretch.
*/
func (this *QFont) SetStretch(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal letterSpacing() const

/*
Returns the letter spacing for the font.

This function was introduced in  Qt 4.4.

See also setLetterSpacing(), letterSpacingType(), and setWordSpacing().
*/
func (this *QFont) LetterSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont13letterSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:229
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::SpacingType letterSpacingType() const

/*
Returns the spacing type used for letter spacing.

This function was introduced in  Qt 4.4.

See also letterSpacing(), setLetterSpacing(), and setWordSpacing().
*/
func (this *QFont) LetterSpacingType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont17letterSpacingTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLetterSpacing(QFont::SpacingType, qreal)

/*
Sets the letter spacing for the font to spacing and the type of spacing to type.

Letter spacing changes the default spacing between individual letters in the font. The spacing between the letters can be made smaller as well as larger either in percentage of the character width or in pixels, depending on the selected spacing type.

This function was introduced in  Qt 4.4.

See also letterSpacing(), letterSpacingType(), and setWordSpacing().
*/
func (this *QFont) SetLetterSpacing(type_ int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont16setLetterSpacingENS_11SpacingTypeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal wordSpacing() const

/*
Returns the word spacing for the font.

This function was introduced in  Qt 4.4.

See also setWordSpacing() and setLetterSpacing().
*/
func (this *QFont) WordSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont11wordSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordSpacing(qreal)

/*
Sets the word spacing for the font to spacing.

Word spacing changes the default spacing between individual words. A positive value increases the word spacing by a corresponding amount of pixels, while a negative value decreases the inter-word spacing accordingly.

Word spacing will not apply to writing systems, where indiviaul words are not separated by white space.

This function was introduced in  Qt 4.4.

See also wordSpacing() and setLetterSpacing().
*/
func (this *QFont) SetWordSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont14setWordSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapitalization(QFont::Capitalization)

/*
Sets the capitalization of the text in this font to caps.

A font's capitalization makes the text appear in the selected capitalization mode.

This function was introduced in  Qt 4.4.

See also capitalization().
*/
func (this *QFont) SetCapitalization(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont17setCapitalizationENS_14CapitalizationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:236
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Capitalization capitalization() const

/*
Returns the current capitalization type of the font.

This function was introduced in  Qt 4.4.

See also setCapitalization().
*/
func (this *QFont) Capitalization() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont14capitalizationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHintingPreference(QFont::HintingPreference)

/*
Set the preference for the hinting level of the glyphs to hintingPreference. This is a hint to the underlying font rendering system to use a certain level of hinting, and has varying support across platforms. See the table in the documentation for QFont::HintingPreference for more details.

The default hinting preference is QFont::PreferDefaultHinting.

This function was introduced in  Qt 4.8.

See also hintingPreference().
*/
func (this *QFont) SetHintingPreference(hintingPreference int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont20setHintingPreferenceENS_17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hintingPreference)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:239
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::HintingPreference hintingPreference() const

/*
Returns the currently preferred hinting level for glyphs rendered with this font.

This function was introduced in  Qt 4.8.

See also setHintingPreference().
*/
func (this *QFont) HintingPreference() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont17hintingPreferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:242
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rawMode() const

/*

 */
func (this *QFont) RawMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7rawModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawMode(bool)

/*

 */
func (this *QFont) SetRawMode(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setRawModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:247
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exactMatch() const

/*
Returns true if a window system font exactly matching the settings of this font is available.

See also QFontInfo.
*/
func (this *QFont) ExactMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont10exactMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:249
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont & operator=(const QFont &)

/*

 */
func (this *QFont) Operator_equal(arg0 QFont_ITF) *QFont {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qfont.h:256
// index:1
// Public inline Visibility=Default Availability=Available
// [16] QFont & operator=(QFont &&)

/*

 */
func (this *QFont) Operator_equal_1(other unsafe.Pointer /*333*/) *QFont {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFontaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qfont.h:250
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QFont &) const

/*

 */
func (this *QFont) Operator_equal_equal(arg0 QFont_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFonteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:251
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QFont &) const

/*

 */
func (this *QFont) Operator_not_equal(arg0 QFont_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFontneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:252
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QFont &) const

/*

 */
func (this *QFont) Operator_less_than(arg0 QFont_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFontltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCopyOf(const QFont &) const

/*
Returns true if this font and f are copies of each other, i.e. one of them was created as a copy of the other and neither has been modified since. This is much stricter than equality.

See also operator=() and operator==().
*/
func (this *QFont) IsCopyOf(arg0 QFont_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont8isCopyOfERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawName(const QString &)

/*

 */
func (this *QFont) SetRawName(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setRawNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QString rawName() const

/*

 */
func (this *QFont) RawName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7rawNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:266
// index:0
// Public Visibility=Default Availability=Available
// [8] QString key() const

/*
Returns the font's key, a textual representation of a font. It is typically used as the key for a cache or dictionary of fonts.

See also QMap.
*/
func (this *QFont) Key() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont3keyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:268
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toString() const

/*
Returns a description of the font. The description is a comma-separated list of the attributes, perfectly suited for use in QSettings.

See also fromString().
*/
func (this *QFont) ToString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont8toStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:269
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fromString(const QString &)

/*
Sets this font to match the description descrip. The description is a comma-separated list of the font attributes, as returned by toString().

See also toString().
*/
func (this *QFont) FromString(arg0 string) bool {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10fromStringERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:271
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString substitute(const QString &)

/*
Returns the first family name to be used whenever familyName is specified. The lookup is case insensitive.

If there is no substitution for familyName, familyName is returned.

To obtain a list of substitutions use substitutes().

See also setFamily(), insertSubstitutions(), insertSubstitution(), and removeSubstitutions().
*/
func (this *QFont) Substitute(arg0 string) string {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10substituteERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QFont_Substitute(arg0 string) string {
	var nilthis *QFont
	rv := nilthis.Substitute(arg0)
	return rv
}

// /usr/include/qt/QtGui/qfont.h:272
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList substitutes(const QString &)

/*
Returns a list of family names to be used whenever familyName is specified. The lookup is case insensitive.

If there is no substitution for familyName, an empty list is returned.

See also substitute(), insertSubstitutions(), insertSubstitution(), and removeSubstitutions().
*/
func (this *QFont) Substitutes(arg0 string) *qtcore.QStringList /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont11substitutesERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QFont_Substitutes(arg0 string) *qtcore.QStringList /*123*/ {
	var nilthis *QFont
	rv := nilthis.Substitutes(arg0)
	return rv
}

// /usr/include/qt/QtGui/qfont.h:273
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList substitutions()

/*
Returns a sorted list of substituted family names.

See also insertSubstitution(), removeSubstitution(), and substitute().
*/
func (this *QFont) Substitutions() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont13substitutionsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QFont_Substitutions() *qtcore.QStringList /*123*/ {
	var nilthis *QFont
	rv := nilthis.Substitutions()
	return rv
}

// /usr/include/qt/QtGui/qfont.h:274
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void insertSubstitution(const QString &, const QString &)

/*
Inserts substituteName into the substitution table for the family familyName.

See also insertSubstitutions(), removeSubstitutions(), substitutions(), substitute(), and substitutes().
*/
func (this *QFont) InsertSubstitution(arg0 string, arg1 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(arg1)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont18insertSubstitutionERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QFont_InsertSubstitution(arg0 string, arg1 string) {
	var nilthis *QFont
	nilthis.InsertSubstitution(arg0, arg1)
}

// /usr/include/qt/QtGui/qfont.h:275
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void insertSubstitutions(const QString &, const QStringList &)

/*
Inserts the list of families substituteNames into the substitution list for familyName.

See also insertSubstitution(), removeSubstitutions(), substitutions(), and substitute().
*/
func (this *QFont) InsertSubstitutions(arg0 string, arg1 qtcore.QStringList_ITF) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QStringList_PTR() != nil {
		convArg1 = arg1.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QFont_InsertSubstitutions(arg0 string, arg1 qtcore.QStringList_ITF) {
	var nilthis *QFont
	nilthis.InsertSubstitutions(arg0, arg1)
}

// /usr/include/qt/QtGui/qfont.h:276
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void removeSubstitutions(const QString &)

/*
Removes all the substitutions for familyName.

This function was introduced in  Qt 5.0.

See also insertSubstitutions(), insertSubstitution(), substitutions(), and substitute().
*/
func (this *QFont) RemoveSubstitutions(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont19removeSubstitutionsERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QFont_RemoveSubstitutions(arg0 string) {
	var nilthis *QFont
	nilthis.RemoveSubstitutions(arg0)
}

// /usr/include/qt/QtGui/qfont.h:280
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void initialize()

/*

 */
func (this *QFont) Initialize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10initializeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QFont_Initialize() {
	var nilthis *QFont
	nilthis.Initialize()
}

// /usr/include/qt/QtGui/qfont.h:281
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void cleanup()

/*

 */
func (this *QFont) Cleanup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont7cleanupEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QFont_Cleanup() {
	var nilthis *QFont
	nilthis.Cleanup()
}

// /usr/include/qt/QtGui/qfont.h:282
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void cacheStatistics()

/*

 */
func (this *QFont) CacheStatistics() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont15cacheStatisticsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QFont_CacheStatistics() {
	var nilthis *QFont
	nilthis.CacheStatistics()
}

// /usr/include/qt/QtGui/qfont.h:284
// index:0
// Public Visibility=Default Availability=Available
// [8] QString defaultFamily() const

/*
Returns the family name that corresponds to the current style hint.

See also StyleHint, styleHint(), and setStyleHint().
*/
func (this *QFont) DefaultFamily() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont13defaultFamilyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:285
// index:0
// Public Visibility=Default Availability=Available
// [8] QString lastResortFamily() const

/*
Returns the "last resort" font family name.

The current implementation tries a wide variety of common fonts, returning the first one it finds. Is is possible that no family is found in which case an empty string is returned.

See also lastResortFont().
*/
func (this *QFont) LastResortFamily() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont16lastResortFamilyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QString lastResortFont() const

/*
Returns a "last resort" font name for the font matching algorithm. This is used if the last resort family is not available. It will always return a name, if necessary returning something like "fixed" or "system".

The current implementation tries a wide variety of common fonts, returning the first one it finds. The implementation may change at any time, but this function will always return a string containing something.

It is theoretically possible that there really isn't a lastResortFont() in which case Qt will abort with an error message. We have not been able to identify a case where this happens. Please report it as a bug if it does, preferably with a list of the fonts you have installed.

See also lastResortFamily().
*/
func (this *QFont) LastResortFont() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont14lastResortFontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:288
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont resolve(const QFont &) const

/*
Returns a new QFont that has attributes copied from other that have not been previously set on this font.
*/
func (this *QFont) Resolve(arg0 QFont_ITF) *QFont /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFont_PTR() != nil {
		convArg0 = arg0.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7resolveERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qfont.h:289
// index:1
// Public inline Visibility=Default Availability=Available
// [4] uint resolve() const

/*
Returns a new QFont that has attributes copied from other that have not been previously set on this font.
*/
func (this *QFont) Resolve_1() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7resolveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qfont.h:290
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void resolve(uint)

/*
Returns a new QFont that has attributes copied from other that have not been previously set on this font.
*/
func (this *QFont) Resolve_2(mask uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont7resolveEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mask)
	qtrt.ErrPrint(err, rv)
}

/*
Style hints are used by the font matching algorithm to find an appropriate default family if a selected font family is not available.

QFont::AnyStyle?leaves the font matching algorithm to choose the family. This is the default.
QFont::SansSerifHelveticathe font matcher prefer sans serif fonts.
QFont::SerifTimesthe font matcher prefers serif fonts.
QFont::Times?is a synonym for Serif.
QFont::TypeWriterCourierthe font matcher prefers fixed pitch fonts.
QFont::Courier?a synonym for TypeWriter.
QFont::OldEnglish?the font matcher prefers decorative fonts.
QFont::DecorativeOldEnglishis a synonym for OldEnglish.
QFont::Monospace?the font matcher prefers fonts that map to the CSS generic font-family 'monospace'.
QFont::Fantasy?the font matcher prefers fonts that map to the CSS generic font-family 'fantasy'.
QFont::Cursive?the font matcher prefers fonts that map to the CSS generic font-family 'cursive'.
QFont::System?the font matcher prefers system fonts.

*/
type QFont__StyleHint = int

// is a synonym for SansSerif.
const QFont__Helvetica QFont__StyleHint = 0

//
const QFont__SansSerif QFont__StyleHint = 0

//
const QFont__Times QFont__StyleHint = 1

//
const QFont__Serif QFont__StyleHint = 1

//
const QFont__Courier QFont__StyleHint = 2

//
const QFont__TypeWriter QFont__StyleHint = 2

//
const QFont__OldEnglish QFont__StyleHint = 3

//
const QFont__Decorative QFont__StyleHint = 3

//
const QFont__System QFont__StyleHint = 4

//
const QFont__AnyStyle QFont__StyleHint = 5

//
const QFont__Cursive QFont__StyleHint = 6

//
const QFont__Monospace QFont__StyleHint = 7

//
const QFont__Fantasy QFont__StyleHint = 8

func (this *QFont) StyleHintItemName(val int) string {
	switch val {
	case QFont__Helvetica: // 0
		return "Helvetica,SansSerif"
		// case QFont__SansSerif: // 0
		// return ""
	case QFont__Times: // 1
		return "Times,Serif"
		// case QFont__Serif: // 1
		// return ""
	case QFont__Courier: // 2
		return "Courier,TypeWriter"
		// case QFont__TypeWriter: // 2
		// return ""
	case QFont__OldEnglish: // 3
		return "OldEnglish,Decorative"
		// case QFont__Decorative: // 3
		// return ""
	case QFont__System: // 4
		return "System"
	case QFont__AnyStyle: // 5
		return "AnyStyle"
	case QFont__Cursive: // 6
		return "Cursive"
	case QFont__Monospace: // 7
		return "Monospace"
	case QFont__Fantasy: // 8
		return "Fantasy"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_StyleHintItemName(val int) string {
	var nilthis *QFont
	return nilthis.StyleHintItemName(val)
}

/*
The style strategy tells the font matching algorithm what type of fonts should be used to find an appropriate default family.

The following strategies are available:



Any of these may be OR-ed with one of these flags:


*/
type QFont__StyleStrategy = int

//
const QFont__PreferDefault QFont__StyleStrategy = 1

//
const QFont__PreferBitmap QFont__StyleStrategy = 2

//
const QFont__PreferDevice QFont__StyleStrategy = 4

//
const QFont__PreferOutline QFont__StyleStrategy = 8

//
const QFont__ForceOutline QFont__StyleStrategy = 16

//
const QFont__PreferMatch QFont__StyleStrategy = 32

//
const QFont__PreferQuality QFont__StyleStrategy = 64

//
const QFont__PreferAntialias QFont__StyleStrategy = 128

//
const QFont__NoAntialias QFont__StyleStrategy = 256

//
const QFont__OpenGLCompatible QFont__StyleStrategy = 512

//
const QFont__ForceIntegerMetrics QFont__StyleStrategy = 1024

//
const QFont__NoSubpixelAntialias QFont__StyleStrategy = 2048

//
const QFont__PreferNoShaping QFont__StyleStrategy = 4096

//
const QFont__NoFontMerging QFont__StyleStrategy = 32768

func (this *QFont) StyleStrategyItemName(val int) string {
	switch val {
	case QFont__PreferDefault: // 1
		return "PreferDefault"
	case QFont__PreferBitmap: // 2
		return "PreferBitmap"
	case QFont__PreferDevice: // 4
		return "PreferDevice"
	case QFont__PreferOutline: // 8
		return "PreferOutline"
	case QFont__ForceOutline: // 16
		return "ForceOutline"
	case QFont__PreferMatch: // 32
		return "PreferMatch"
	case QFont__PreferQuality: // 64
		return "PreferQuality"
	case QFont__PreferAntialias: // 128
		return "PreferAntialias"
	case QFont__NoAntialias: // 256
		return "NoAntialias"
	case QFont__OpenGLCompatible: // 512
		return "OpenGLCompatible"
	case QFont__ForceIntegerMetrics: // 1024
		return "ForceIntegerMetrics"
	case QFont__NoSubpixelAntialias: // 2048
		return "NoSubpixelAntialias"
	case QFont__PreferNoShaping: // 4096
		return "PreferNoShaping"
	case QFont__NoFontMerging: // 32768
		return "NoFontMerging"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_StyleStrategyItemName(val int) string {
	var nilthis *QFont
	return nilthis.StyleStrategyItemName(val)
}

/*
This enum describes the different levels of hinting that can be applied to glyphs to improve legibility on displays where it might be warranted by the density of pixels.



Please note that this enum only describes a preference, as the full range of hinting levels are not supported on all of Qt's supported platforms. The following table details the effect of a given hinting preference on a selected set of target platforms.


 PreferDefaultHintingPreferNoHintingPreferVerticalHintingPreferFullHinting
Windows Vista (w/o Platform Update) and earlierFull hintingFull hintingFull hintingFull hinting
Windows 7 and Windows Vista (w/Platform Update) and DirectWrite enabled in QtFull hintingVertical hintingVertical hintingFull hinting
FreeTypeOperating System settingNo hintingVertical hinting (light)Full hinting
Cocoa on macOSNo hintingNo hintingNo hintingNo hinting


Note: Please be aware that altering the hinting preference on Windows is available through the DirectWrite font engine. This is available on Windows Vista after installing the platform update, and on Windows 7. In order to use this extension, configure Qt using -directwrite. The target application will then depend on the availability of DirectWrite on the target system.

This enum was introduced or modified in  Qt 4.8.

*/
type QFont__HintingPreference = int

// Use the default hinting level for the target platform.
const QFont__PreferDefaultHinting QFont__HintingPreference = 0

// If possible, render text without hinting the outlines of the glyphs. The text layout will be typographically accurate and scalable, using the same metrics as are used e.g. when printing.
const QFont__PreferNoHinting QFont__HintingPreference = 1

// If possible, render text with no horizontal hinting, but align glyphs to the pixel grid in the vertical direction. The text will appear crisper on displays where the density is too low to give an accurate rendering of the glyphs. But since the horizontal metrics of the glyphs are unhinted, the text's layout will be scalable to higher density devices (such as printers) without impacting details such as line breaks.
const QFont__PreferVerticalHinting QFont__HintingPreference = 2

// If possible, render text with hinting in both horizontal and vertical directions. The text will be altered to optimize legibility on the target device, but since the metrics will depend on the target size of the text, the positions of glyphs, line breaks, and other typographical detail will not scale, meaning that a text layout may look different on devices with different pixel densities.
const QFont__PreferFullHinting QFont__HintingPreference = 3

func (this *QFont) HintingPreferenceItemName(val int) string {
	switch val {
	case QFont__PreferDefaultHinting: // 0
		return "PreferDefaultHinting"
	case QFont__PreferNoHinting: // 1
		return "PreferNoHinting"
	case QFont__PreferVerticalHinting: // 2
		return "PreferVerticalHinting"
	case QFont__PreferFullHinting: // 3
		return "PreferFullHinting"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_HintingPreferenceItemName(val int) string {
	var nilthis *QFont
	return nilthis.HintingPreferenceItemName(val)
}

/*
Qt uses a weighting scale from 0 to 99 similar to, but not the same as, the scales used in Windows or CSS. A weight of 0 will be thin, whilst 99 will be extremely black.

This enum contains the predefined font weights:


*/
type QFont__Weight = int

// 0
const QFont__Thin QFont__Weight = 0

//
const QFont__ExtraLight QFont__Weight = 12

//
const QFont__Light QFont__Weight = 25

//
const QFont__Normal QFont__Weight = 50

//
const QFont__Medium QFont__Weight = 57

//
const QFont__DemiBold QFont__Weight = 63

//
const QFont__Bold QFont__Weight = 75

//
const QFont__ExtraBold QFont__Weight = 81

//
const QFont__Black QFont__Weight = 87

func (this *QFont) WeightItemName(val int) string {
	switch val {
	case QFont__Thin: // 0
		return "Thin"
	case QFont__ExtraLight: // 12
		return "ExtraLight"
	case QFont__Light: // 25
		return "Light"
	case QFont__Normal: // 50
		return "Normal"
	case QFont__Medium: // 57
		return "Medium"
	case QFont__DemiBold: // 63
		return "DemiBold"
	case QFont__Bold: // 75
		return "Bold"
	case QFont__ExtraBold: // 81
		return "ExtraBold"
	case QFont__Black: // 87
		return "Black"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_WeightItemName(val int) string {
	var nilthis *QFont
	return nilthis.WeightItemName(val)
}

/*
This enum describes the different styles of glyphs that are used to display text.



See also Weight.

*/
type QFont__Style = int

// Normal glyphs used in unstyled text.
const QFont__StyleNormal QFont__Style = 0

// Italic glyphs that are specifically designed for the purpose of representing italicized text.
const QFont__StyleItalic QFont__Style = 1

// Glyphs with an italic appearance that are typically based on the unstyled glyphs, but are not fine-tuned for the purpose of representing italicized text.
const QFont__StyleOblique QFont__Style = 2

func (this *QFont) StyleItemName(val int) string {
	switch val {
	case QFont__StyleNormal: // 0
		return "StyleNormal"
	case QFont__StyleItalic: // 1
		return "StyleItalic"
	case QFont__StyleOblique: // 2
		return "StyleOblique"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_StyleItemName(val int) string {
	var nilthis *QFont
	return nilthis.StyleItemName(val)
}

/*
Predefined stretch values that follow the CSS naming convention. The higher the value, the more stretched the text is.



See also setStretch() and stretch().

*/
type QFont__Stretch = int

//
const QFont__AnyStretch QFont__Stretch = 0

//
const QFont__UltraCondensed QFont__Stretch = 50

//
const QFont__ExtraCondensed QFont__Stretch = 62

//
const QFont__Condensed QFont__Stretch = 75

//
const QFont__SemiCondensed QFont__Stretch = 87

//
const QFont__Unstretched QFont__Stretch = 100

//
const QFont__SemiExpanded QFont__Stretch = 112

//
const QFont__Expanded QFont__Stretch = 125

//
const QFont__ExtraExpanded QFont__Stretch = 150

//
const QFont__UltraExpanded QFont__Stretch = 200

func (this *QFont) StretchItemName(val int) string {
	switch val {
	case QFont__AnyStretch: // 0
		return "AnyStretch"
	case QFont__UltraCondensed: // 50
		return "UltraCondensed"
	case QFont__ExtraCondensed: // 62
		return "ExtraCondensed"
	case QFont__Condensed: // 75
		return "Condensed"
	case QFont__SemiCondensed: // 87
		return "SemiCondensed"
	case QFont__Unstretched: // 100
		return "Unstretched"
	case QFont__SemiExpanded: // 112
		return "SemiExpanded"
	case QFont__Expanded: // 125
		return "Expanded"
	case QFont__ExtraExpanded: // 150
		return "ExtraExpanded"
	case QFont__UltraExpanded: // 200
		return "UltraExpanded"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_StretchItemName(val int) string {
	var nilthis *QFont
	return nilthis.StretchItemName(val)
}

/*
Rendering option for text this font applies to.



This enum was introduced or modified in  Qt 4.4.

*/
type QFont__Capitalization = int

// This is the normal text rendering option where no capitalization change is applied.
const QFont__MixedCase QFont__Capitalization = 0

// This alters the text to be rendered in all uppercase type.
const QFont__AllUppercase QFont__Capitalization = 1

// This alters the text to be rendered in all lowercase type.
const QFont__AllLowercase QFont__Capitalization = 2

// This alters the text to be rendered in small-caps type.
const QFont__SmallCaps QFont__Capitalization = 3

// This alters the text to be rendered with the first character of each word as an uppercase character.
const QFont__Capitalize QFont__Capitalization = 4

func (this *QFont) CapitalizationItemName(val int) string {
	switch val {
	case QFont__MixedCase: // 0
		return "MixedCase"
	case QFont__AllUppercase: // 1
		return "AllUppercase"
	case QFont__AllLowercase: // 2
		return "AllLowercase"
	case QFont__SmallCaps: // 3
		return "SmallCaps"
	case QFont__Capitalize: // 4
		return "Capitalize"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_CapitalizationItemName(val int) string {
	var nilthis *QFont
	return nilthis.CapitalizationItemName(val)
}

/*


This enum was introduced or modified in  Qt 4.4.

*/
type QFont__SpacingType = int

//
const QFont__PercentageSpacing QFont__SpacingType = 0

// A positive value increases the letter spacing by the corresponding pixels; a negative value decreases the spacing.
const QFont__AbsoluteSpacing QFont__SpacingType = 1

func (this *QFont) SpacingTypeItemName(val int) string {
	switch val {
	case QFont__PercentageSpacing: // 0
		return "PercentageSpacing"
	case QFont__AbsoluteSpacing: // 1
		return "AbsoluteSpacing"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_SpacingTypeItemName(val int) string {
	var nilthis *QFont
	return nilthis.SpacingTypeItemName(val)
}

/*


 */
type QFont__ResolveProperties = int

//
const QFont__FamilyResolved QFont__ResolveProperties = 1

//
const QFont__SizeResolved QFont__ResolveProperties = 2

//
const QFont__StyleHintResolved QFont__ResolveProperties = 4

//
const QFont__StyleStrategyResolved QFont__ResolveProperties = 8

//
const QFont__WeightResolved QFont__ResolveProperties = 16

//
const QFont__StyleResolved QFont__ResolveProperties = 32

//
const QFont__UnderlineResolved QFont__ResolveProperties = 64

//
const QFont__OverlineResolved QFont__ResolveProperties = 128

//
const QFont__StrikeOutResolved QFont__ResolveProperties = 256

//
const QFont__FixedPitchResolved QFont__ResolveProperties = 512

//
const QFont__StretchResolved QFont__ResolveProperties = 1024

//
const QFont__KerningResolved QFont__ResolveProperties = 2048

//
const QFont__CapitalizationResolved QFont__ResolveProperties = 4096

//
const QFont__LetterSpacingResolved QFont__ResolveProperties = 8192

//
const QFont__WordSpacingResolved QFont__ResolveProperties = 16384

//
const QFont__HintingPreferenceResolved QFont__ResolveProperties = 32768

//
const QFont__StyleNameResolved QFont__ResolveProperties = 65536

//
const QFont__AllPropertiesResolved QFont__ResolveProperties = 131071

func (this *QFont) ResolvePropertiesItemName(val int) string {
	switch val {
	case QFont__FamilyResolved: // 1
		return "FamilyResolved"
	case QFont__SizeResolved: // 2
		return "SizeResolved"
	case QFont__StyleHintResolved: // 4
		return "StyleHintResolved"
	case QFont__StyleStrategyResolved: // 8
		return "StyleStrategyResolved"
	case QFont__WeightResolved: // 16
		return "WeightResolved"
	case QFont__StyleResolved: // 32
		return "StyleResolved"
	case QFont__UnderlineResolved: // 64
		return "UnderlineResolved"
	case QFont__OverlineResolved: // 128
		return "OverlineResolved"
	case QFont__StrikeOutResolved: // 256
		return "StrikeOutResolved"
	case QFont__FixedPitchResolved: // 512
		return "FixedPitchResolved"
	case QFont__StretchResolved: // 1024
		return "StretchResolved"
	case QFont__KerningResolved: // 2048
		return "KerningResolved"
	case QFont__CapitalizationResolved: // 4096
		return "CapitalizationResolved"
	case QFont__LetterSpacingResolved: // 8192
		return "LetterSpacingResolved"
	case QFont__WordSpacingResolved: // 16384
		return "WordSpacingResolved"
	case QFont__HintingPreferenceResolved: // 32768
		return "HintingPreferenceResolved"
	case QFont__StyleNameResolved: // 65536
		return "StyleNameResolved"
	case QFont__AllPropertiesResolved: // 131071
		return "AllPropertiesResolved"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QFont_ResolvePropertiesItemName(val int) string {
	var nilthis *QFont
	return nilthis.ResolvePropertiesItemName(val)
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
