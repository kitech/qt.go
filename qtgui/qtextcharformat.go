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
type QTextCharFormat struct {
	*QTextFormat
}
type QTextCharFormat_ITF interface {
	QTextFormat_ITF
	QTextCharFormat_PTR() *QTextCharFormat
}

func (ptr *QTextCharFormat) QTextCharFormat_PTR() *QTextCharFormat { return ptr }

func (this *QTextCharFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFormat.GetCthis()
	}
}
func (this *QTextCharFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextFormat = NewQTextFormatFromPointer(cthis)
}
func NewQTextCharFormatFromPointer(cthis unsafe.Pointer) *QTextCharFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextCharFormat{bcthis0}
}
func (*QTextCharFormat) NewFromPointer(cthis unsafe.Pointer) *QTextCharFormat {
	return NewQTextCharFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:412
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextCharFormat()

/*

 */
func (*QTextCharFormat) NewForInherit() *QTextCharFormat {
	return NewQTextCharFormat()
}
func NewQTextCharFormat() *QTextCharFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextCharFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:557
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextCharFormat(const QTextFormat &)

/*

 */
func (*QTextCharFormat) NewForInherit1(fmt_ QTextFormat_ITF) *QTextCharFormat {
	return NewQTextCharFormat1(fmt_)
}
func NewQTextCharFormat1(fmt_ QTextFormat_ITF) *QTextCharFormat {
	var convArg0 unsafe.Pointer
	if fmt_ != nil && fmt_.QTextFormat_PTR() != nil {
		convArg0 = fmt_.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextCharFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:414
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the format is valid (i.e. is not InvalidFormat); otherwise returns false.
*/
func (this *QTextCharFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:420
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &, QTextCharFormat::FontPropertiesInheritanceBehavior)

/*

 */
func (this *QTextCharFormat) SetFont(font QFont_ITF, behavior int) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat7setFontERK5QFontNS_33FontPropertiesInheritanceBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, behavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:421
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)

/*

 */
func (this *QTextCharFormat) SetFont1(font QFont_ITF) {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:422
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font() const

/*

 */
func (this *QTextCharFormat) Font() *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:424
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontFamily(const QString &)

/*

 */
func (this *QTextCharFormat) SetFontFamily(family string) {
	var tmpArg0 = qtcore.NewQString5(family)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontFamilyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:426
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString fontFamily() const

/*

 */
func (this *QTextCharFormat) FontFamily() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontFamilyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:429
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontPointSize(qreal)

/*

 */
func (this *QTextCharFormat) SetFontPointSize(size float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontPointSizeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:431
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal fontPointSize() const

/*

 */
func (this *QTextCharFormat) FontPointSize() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontPointSizeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:434
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontWeight(int)

/*

 */
func (this *QTextCharFormat) SetFontWeight(weight int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontWeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), weight)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:436
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int fontWeight() const

/*

 */
func (this *QTextCharFormat) FontWeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontWeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:438
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontItalic(bool)

/*

 */
func (this *QTextCharFormat) SetFontItalic(italic bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontItalicEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), italic)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:440
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool fontItalic() const

/*

 */
func (this *QTextCharFormat) FontItalic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontItalicEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:442
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontCapitalization(QFont::Capitalization)

/*

 */
func (this *QTextCharFormat) SetFontCapitalization(capitalization int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat21setFontCapitalizationEN5QFont14CapitalizationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), capitalization)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:444
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QFont::Capitalization fontCapitalization() const

/*

 */
func (this *QTextCharFormat) FontCapitalization() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat18fontCapitalizationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:446
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontLetterSpacingType(QFont::SpacingType)

/*

 */
func (this *QTextCharFormat) SetFontLetterSpacingType(letterSpacingType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat24setFontLetterSpacingTypeEN5QFont11SpacingTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), letterSpacingType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:448
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QFont::SpacingType fontLetterSpacingType() const

/*

 */
func (this *QTextCharFormat) FontLetterSpacingType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat21fontLetterSpacingTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:450
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontLetterSpacing(qreal)

/*

 */
func (this *QTextCharFormat) SetFontLetterSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat20setFontLetterSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:452
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal fontLetterSpacing() const

/*

 */
func (this *QTextCharFormat) FontLetterSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat17fontLetterSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:454
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontWordSpacing(qreal)

/*

 */
func (this *QTextCharFormat) SetFontWordSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat18setFontWordSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:456
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal fontWordSpacing() const

/*

 */
func (this *QTextCharFormat) FontWordSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat15fontWordSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:459
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontUnderline(bool)

/*

 */
func (this *QTextCharFormat) SetFontUnderline(underline bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontUnderlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), underline)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:461
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fontUnderline() const

/*

 */
func (this *QTextCharFormat) FontUnderline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontUnderlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:463
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontOverline(bool)

/*

 */
func (this *QTextCharFormat) SetFontOverline(overline bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat15setFontOverlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), overline)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:465
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool fontOverline() const

/*

 */
func (this *QTextCharFormat) FontOverline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat12fontOverlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:468
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontStrikeOut(bool)

/*

 */
func (this *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontStrikeOutEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), strikeOut)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:470
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool fontStrikeOut() const

/*

 */
func (this *QTextCharFormat) FontStrikeOut() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontStrikeOutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:473
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setUnderlineColor(const QColor &)

/*

 */
func (this *QTextCharFormat) SetUnderlineColor(color QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat17setUnderlineColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:475
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QColor underlineColor() const

/*

 */
func (this *QTextCharFormat) UnderlineColor() *QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat14underlineColorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:478
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontFixedPitch(bool)

/*

 */
func (this *QTextCharFormat) SetFontFixedPitch(fixedPitch bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat17setFontFixedPitchEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fixedPitch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:480
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool fontFixedPitch() const

/*

 */
func (this *QTextCharFormat) FontFixedPitch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat14fontFixedPitchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:483
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontStretch(int)

/*

 */
func (this *QTextCharFormat) SetFontStretch(factor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat14setFontStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), factor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:485
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int fontStretch() const

/*

 */
func (this *QTextCharFormat) FontStretch() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat11fontStretchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:488
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontStyleHint(QFont::StyleHint, QFont::StyleStrategy)

/*

 */
func (this *QTextCharFormat) SetFontStyleHint(hint int, strategy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontStyleHintEN5QFont9StyleHintENS0_13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, strategy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:488
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontStyleHint(QFont::StyleHint, QFont::StyleStrategy)

/*

 */
func (this *QTextCharFormat) SetFontStyleHintp(hint int) {
	// arg: 1, QFont::StyleStrategy=Elaborated, QFont::StyleStrategy=Enum, , Invalid
	strategy := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontStyleHintEN5QFont9StyleHintENS0_13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hint, strategy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:490
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontStyleStrategy(QFont::StyleStrategy)

/*

 */
func (this *QTextCharFormat) SetFontStyleStrategy(strategy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat20setFontStyleStrategyEN5QFont13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), strategy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:492
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QFont::StyleHint fontStyleHint() const

/*

 */
func (this *QTextCharFormat) FontStyleHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontStyleHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:494
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QFont::StyleStrategy fontStyleStrategy() const

/*

 */
func (this *QTextCharFormat) FontStyleStrategy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat17fontStyleStrategyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:497
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontHintingPreference(QFont::HintingPreference)

/*

 */
func (this *QTextCharFormat) SetFontHintingPreference(hintingPreference int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat24setFontHintingPreferenceEN5QFont17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hintingPreference)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:502
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QFont::HintingPreference fontHintingPreference() const

/*

 */
func (this *QTextCharFormat) FontHintingPreference() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat21fontHintingPreferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:507
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFontKerning(bool)

/*

 */
func (this *QTextCharFormat) SetFontKerning(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat14setFontKerningEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:509
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool fontKerning() const

/*

 */
func (this *QTextCharFormat) FontKerning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat11fontKerningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:512
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnderlineStyle(QTextCharFormat::UnderlineStyle)

/*

 */
func (this *QTextCharFormat) SetUnderlineStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat17setUnderlineStyleENS_14UnderlineStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:513
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextCharFormat::UnderlineStyle underlineStyle() const

/*

 */
func (this *QTextCharFormat) UnderlineStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat14underlineStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:516
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setVerticalAlignment(QTextCharFormat::VerticalAlignment)

/*

 */
func (this *QTextCharFormat) SetVerticalAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat20setVerticalAlignmentENS_17VerticalAlignmentE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:518
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextCharFormat::VerticalAlignment verticalAlignment() const

/*

 */
func (this *QTextCharFormat) VerticalAlignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat17verticalAlignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:521
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextOutline(const QPen &)

/*

 */
func (this *QTextCharFormat) SetTextOutline(pen QPen_ITF) {
	var convArg0 unsafe.Pointer
	if pen != nil && pen.QPen_PTR() != nil {
		convArg0 = pen.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat14setTextOutlineERK4QPen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:523
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPen textOutline() const

/*

 */
func (this *QTextCharFormat) TextOutline() *QPen /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat11textOutlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPen)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:526
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)

/*

 */
func (this *QTextCharFormat) SetToolTip(tip string) {
	var tmpArg0 = qtcore.NewQString5(tip)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:528
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString toolTip() const

/*

 */
func (this *QTextCharFormat) ToolTip() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:531
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAnchor(bool)

/*

 */
func (this *QTextCharFormat) SetAnchor(anchor bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat9setAnchorEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), anchor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:533
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAnchor() const

/*

 */
func (this *QTextCharFormat) IsAnchor() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat8isAnchorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:536
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAnchorHref(const QString &)

/*

 */
func (this *QTextCharFormat) SetAnchorHref(value string) {
	var tmpArg0 = qtcore.NewQString5(value)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat13setAnchorHrefERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:538
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QString anchorHref() const

/*

 */
func (this *QTextCharFormat) AnchorHref() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat10anchorHrefEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:541
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAnchorName(const QString &)

/*

 */
func (this *QTextCharFormat) SetAnchorName(name string) {
	var tmpArg0 = qtcore.NewQString5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat13setAnchorNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:543
// index:0
// Public Visibility=Default Availability=Available
// [8] QString anchorName() const

/*

 */
func (this *QTextCharFormat) AnchorName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat10anchorNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextformat.h:545
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAnchorNames(const QStringList &)

/*

 */
func (this *QTextCharFormat) SetAnchorNames(names qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if names != nil && names.QStringList_PTR() != nil {
		convArg0 = names.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat14setAnchorNamesERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:547
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList anchorNames() const

/*

 */
func (this *QTextCharFormat) AnchorNames() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat11anchorNamesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:549
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTableCellRowSpan(int)

/*

 */
func (this *QTextCharFormat) SetTableCellRowSpan(tableCellRowSpan int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat19setTableCellRowSpanEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tableCellRowSpan)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:550
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int tableCellRowSpan() const

/*

 */
func (this *QTextCharFormat) TableCellRowSpan() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat16tableCellRowSpanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:552
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTableCellColumnSpan(int)

/*

 */
func (this *QTextCharFormat) SetTableCellColumnSpan(tableCellColumnSpan int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormat22setTableCellColumnSpanEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tableCellColumnSpan)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:553
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int tableCellColumnSpan() const

/*

 */
func (this *QTextCharFormat) TableCellColumnSpan() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QTextCharFormat19tableCellColumnSpanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTextCharFormat(this *QTextCharFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QTextCharFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTextCharFormat__VerticalAlignment = int

//
const QTextCharFormat__AlignNormal QTextCharFormat__VerticalAlignment = 0

//
const QTextCharFormat__AlignSuperScript QTextCharFormat__VerticalAlignment = 1

//
const QTextCharFormat__AlignSubScript QTextCharFormat__VerticalAlignment = 2

//
const QTextCharFormat__AlignMiddle QTextCharFormat__VerticalAlignment = 3

//
const QTextCharFormat__AlignTop QTextCharFormat__VerticalAlignment = 4

//
const QTextCharFormat__AlignBottom QTextCharFormat__VerticalAlignment = 5

//
const QTextCharFormat__AlignBaseline QTextCharFormat__VerticalAlignment = 6

func (this *QTextCharFormat) VerticalAlignmentItemName(val int) string {
	switch val {
	case QTextCharFormat__AlignNormal: // 0
		return "AlignNormal"
	case QTextCharFormat__AlignSuperScript: // 1
		return "AlignSuperScript"
	case QTextCharFormat__AlignSubScript: // 2
		return "AlignSubScript"
	case QTextCharFormat__AlignMiddle: // 3
		return "AlignMiddle"
	case QTextCharFormat__AlignTop: // 4
		return "AlignTop"
	case QTextCharFormat__AlignBottom: // 5
		return "AlignBottom"
	case QTextCharFormat__AlignBaseline: // 6
		return "AlignBaseline"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextCharFormat_VerticalAlignmentItemName(val int) string {
	var nilthis *QTextCharFormat
	return nilthis.VerticalAlignmentItemName(val)
}

/*


 */
type QTextCharFormat__UnderlineStyle = int

//
const QTextCharFormat__NoUnderline QTextCharFormat__UnderlineStyle = 0

//
const QTextCharFormat__SingleUnderline QTextCharFormat__UnderlineStyle = 1

//
const QTextCharFormat__DashUnderline QTextCharFormat__UnderlineStyle = 2

//
const QTextCharFormat__DotLine QTextCharFormat__UnderlineStyle = 3

//
const QTextCharFormat__DashDotLine QTextCharFormat__UnderlineStyle = 4

//
const QTextCharFormat__DashDotDotLine QTextCharFormat__UnderlineStyle = 5

//
const QTextCharFormat__WaveUnderline QTextCharFormat__UnderlineStyle = 6

//
const QTextCharFormat__SpellCheckUnderline QTextCharFormat__UnderlineStyle = 7

func (this *QTextCharFormat) UnderlineStyleItemName(val int) string {
	switch val {
	case QTextCharFormat__NoUnderline: // 0
		return "NoUnderline"
	case QTextCharFormat__SingleUnderline: // 1
		return "SingleUnderline"
	case QTextCharFormat__DashUnderline: // 2
		return "DashUnderline"
	case QTextCharFormat__DotLine: // 3
		return "DotLine"
	case QTextCharFormat__DashDotLine: // 4
		return "DashDotLine"
	case QTextCharFormat__DashDotDotLine: // 5
		return "DashDotDotLine"
	case QTextCharFormat__WaveUnderline: // 6
		return "WaveUnderline"
	case QTextCharFormat__SpellCheckUnderline: // 7
		return "SpellCheckUnderline"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextCharFormat_UnderlineStyleItemName(val int) string {
	var nilthis *QTextCharFormat
	return nilthis.UnderlineStyleItemName(val)
}

/*


 */
type QTextCharFormat__FontPropertiesInheritanceBehavior = int

//
const QTextCharFormat__FontPropertiesSpecifiedOnly QTextCharFormat__FontPropertiesInheritanceBehavior = 0

//
const QTextCharFormat__FontPropertiesAll QTextCharFormat__FontPropertiesInheritanceBehavior = 1

func (this *QTextCharFormat) FontPropertiesInheritanceBehaviorItemName(val int) string {
	switch val {
	case QTextCharFormat__FontPropertiesSpecifiedOnly: // 0
		return "FontPropertiesSpecifiedOnly"
	case QTextCharFormat__FontPropertiesAll: // 1
		return "FontPropertiesAll"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextCharFormat_FontPropertiesInheritanceBehaviorItemName(val int) string {
	var nilthis *QTextCharFormat
	return nilthis.FontPropertiesInheritanceBehaviorItemName(val)
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
