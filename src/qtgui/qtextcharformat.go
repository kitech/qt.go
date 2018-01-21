package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
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
type QTextCharFormat struct {
	*QTextFormat
}

func (this *QTextCharFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFormat.GetCthis()
	}
}
func NewQTextCharFormatFromPointer(cthis unsafe.Pointer) *QTextCharFormat {
	bcthis0 := NewQTextFormatFromPointer(cthis)
	return &QTextCharFormat{bcthis0}
}

// /usr/include/qt/QtGui/qtextformat.h:412
// index:0
// Public
// void QTextCharFormat()
func NewQTextCharFormat() *QTextCharFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormatC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCharFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:557
// index:1
// Protected
// void QTextCharFormat(const class QTextFormat &)
func NewQTextCharFormat_1(fmt *QTextFormat) *QTextCharFormat {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = fmt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextCharFormatFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:414
// index:0
// Public inline
// bool isValid()
func (this *QTextCharFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:420
// index:0
// Public
// void setFont(const class QFont &, enum QTextCharFormat::FontPropertiesInheritanceBehavior)
func (this *QTextCharFormat) SetFont(font *QFont, behavior int) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat7setFontERK5QFontNS_33FontPropertiesInheritanceBehaviorE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &behavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:421
// index:1
// Public
// void setFont(const class QFont &)
func (this *QTextCharFormat) SetFont_1(font *QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:422
// index:0
// Public
// QFont font()
func (this *QTextCharFormat) Font() *QFont /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:424
// index:0
// Public inline
// void setFontFamily(const class QString &)
func (this *QTextCharFormat) SetFontFamily(family *qtcore.QString) {
	var convArg0 = family.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontFamilyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:426
// index:0
// Public inline
// QString fontFamily()
func (this *QTextCharFormat) FontFamily() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontFamilyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:429
// index:0
// Public inline
// void setFontPointSize(qreal)
func (this *QTextCharFormat) SetFontPointSize(size float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontPointSizeEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:431
// index:0
// Public inline
// qreal fontPointSize()
func (this *QTextCharFormat) FontPointSize() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontPointSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextformat.h:434
// index:0
// Public inline
// void setFontWeight(int)
func (this *QTextCharFormat) SetFontWeight(weight int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontWeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &weight)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:436
// index:0
// Public inline
// int fontWeight()
func (this *QTextCharFormat) FontWeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontWeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:438
// index:0
// Public inline
// void setFontItalic(_Bool)
func (this *QTextCharFormat) SetFontItalic(italic bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setFontItalicEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &italic)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:440
// index:0
// Public inline
// bool fontItalic()
func (this *QTextCharFormat) FontItalic() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10fontItalicEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:442
// index:0
// Public inline
// void setFontCapitalization(class QFont::Capitalization)
func (this *QTextCharFormat) SetFontCapitalization(capitalization int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat21setFontCapitalizationEN5QFont14CapitalizationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &capitalization)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:444
// index:0
// Public inline
// QFont::Capitalization fontCapitalization()
func (this *QTextCharFormat) FontCapitalization() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat18fontCapitalizationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:446
// index:0
// Public inline
// void setFontLetterSpacingType(class QFont::SpacingType)
func (this *QTextCharFormat) SetFontLetterSpacingType(letterSpacingType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat24setFontLetterSpacingTypeEN5QFont11SpacingTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &letterSpacingType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:448
// index:0
// Public inline
// QFont::SpacingType fontLetterSpacingType()
func (this *QTextCharFormat) FontLetterSpacingType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat21fontLetterSpacingTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:450
// index:0
// Public inline
// void setFontLetterSpacing(qreal)
func (this *QTextCharFormat) SetFontLetterSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat20setFontLetterSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:452
// index:0
// Public inline
// qreal fontLetterSpacing()
func (this *QTextCharFormat) FontLetterSpacing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat17fontLetterSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextformat.h:454
// index:0
// Public inline
// void setFontWordSpacing(qreal)
func (this *QTextCharFormat) SetFontWordSpacing(spacing float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat18setFontWordSpacingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:456
// index:0
// Public inline
// qreal fontWordSpacing()
func (this *QTextCharFormat) FontWordSpacing() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat15fontWordSpacingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qtextformat.h:459
// index:0
// Public inline
// void setFontUnderline(_Bool)
func (this *QTextCharFormat) SetFontUnderline(underline bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontUnderlineEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &underline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:461
// index:0
// Public
// bool fontUnderline()
func (this *QTextCharFormat) FontUnderline() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontUnderlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:463
// index:0
// Public inline
// void setFontOverline(_Bool)
func (this *QTextCharFormat) SetFontOverline(overline bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat15setFontOverlineEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &overline)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:465
// index:0
// Public inline
// bool fontOverline()
func (this *QTextCharFormat) FontOverline() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat12fontOverlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:468
// index:0
// Public inline
// void setFontStrikeOut(_Bool)
func (this *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontStrikeOutEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &strikeOut)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:470
// index:0
// Public inline
// bool fontStrikeOut()
func (this *QTextCharFormat) FontStrikeOut() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontStrikeOutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:473
// index:0
// Public inline
// void setUnderlineColor(const class QColor &)
func (this *QTextCharFormat) SetUnderlineColor(color *QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat17setUnderlineColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:475
// index:0
// Public inline
// QColor underlineColor()
func (this *QTextCharFormat) UnderlineColor() *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat14underlineColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:478
// index:0
// Public inline
// void setFontFixedPitch(_Bool)
func (this *QTextCharFormat) SetFontFixedPitch(fixedPitch bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat17setFontFixedPitchEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &fixedPitch)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:480
// index:0
// Public inline
// bool fontFixedPitch()
func (this *QTextCharFormat) FontFixedPitch() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat14fontFixedPitchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:483
// index:0
// Public inline
// void setFontStretch(int)
func (this *QTextCharFormat) SetFontStretch(factor int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setFontStretchEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &factor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:485
// index:0
// Public inline
// int fontStretch()
func (this *QTextCharFormat) FontStretch() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat11fontStretchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:488
// index:0
// Public inline
// void setFontStyleHint(class QFont::StyleHint, class QFont::StyleStrategy)
func (this *QTextCharFormat) SetFontStyleHint(hint int, strategy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat16setFontStyleHintEN5QFont9StyleHintENS0_13StyleStrategyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &hint, &strategy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:490
// index:0
// Public inline
// void setFontStyleStrategy(class QFont::StyleStrategy)
func (this *QTextCharFormat) SetFontStyleStrategy(strategy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat20setFontStyleStrategyEN5QFont13StyleStrategyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &strategy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:492
// index:0
// Public inline
// QFont::StyleHint fontStyleHint()
func (this *QTextCharFormat) FontStyleHint() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat13fontStyleHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:494
// index:0
// Public inline
// QFont::StyleStrategy fontStyleStrategy()
func (this *QTextCharFormat) FontStyleStrategy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat17fontStyleStrategyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:497
// index:0
// Public inline
// void setFontHintingPreference(class QFont::HintingPreference)
func (this *QTextCharFormat) SetFontHintingPreference(hintingPreference int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat24setFontHintingPreferenceEN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:502
// index:0
// Public inline
// QFont::HintingPreference fontHintingPreference()
func (this *QTextCharFormat) FontHintingPreference() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat21fontHintingPreferenceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:507
// index:0
// Public inline
// void setFontKerning(_Bool)
func (this *QTextCharFormat) SetFontKerning(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setFontKerningEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:509
// index:0
// Public inline
// bool fontKerning()
func (this *QTextCharFormat) FontKerning() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat11fontKerningEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:512
// index:0
// Public
// void setUnderlineStyle(enum QTextCharFormat::UnderlineStyle)
func (this *QTextCharFormat) SetUnderlineStyle(style int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat17setUnderlineStyleENS_14UnderlineStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:513
// index:0
// Public inline
// QTextCharFormat::UnderlineStyle underlineStyle()
func (this *QTextCharFormat) UnderlineStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat14underlineStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:516
// index:0
// Public inline
// void setVerticalAlignment(enum QTextCharFormat::VerticalAlignment)
func (this *QTextCharFormat) SetVerticalAlignment(alignment int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat20setVerticalAlignmentENS_17VerticalAlignmentE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &alignment)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:518
// index:0
// Public inline
// QTextCharFormat::VerticalAlignment verticalAlignment()
func (this *QTextCharFormat) VerticalAlignment() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat17verticalAlignmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:521
// index:0
// Public inline
// void setTextOutline(const class QPen &)
func (this *QTextCharFormat) SetTextOutline(pen *QPen) {
	var convArg0 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setTextOutlineERK4QPen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:523
// index:0
// Public inline
// QPen textOutline()
func (this *QTextCharFormat) TextOutline() *QPen /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat11textOutlineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:526
// index:0
// Public inline
// void setToolTip(const class QString &)
func (this *QTextCharFormat) SetToolTip(tip *qtcore.QString) {
	var convArg0 = tip.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:528
// index:0
// Public inline
// QString toolTip()
func (this *QTextCharFormat) ToolTip() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat7toolTipEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:531
// index:0
// Public inline
// void setAnchor(_Bool)
func (this *QTextCharFormat) SetAnchor(anchor bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat9setAnchorEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &anchor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:533
// index:0
// Public inline
// bool isAnchor()
func (this *QTextCharFormat) IsAnchor() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat8isAnchorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:536
// index:0
// Public inline
// void setAnchorHref(const class QString &)
func (this *QTextCharFormat) SetAnchorHref(value *qtcore.QString) {
	var convArg0 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setAnchorHrefERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:538
// index:0
// Public inline
// QString anchorHref()
func (this *QTextCharFormat) AnchorHref() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10anchorHrefEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:541
// index:0
// Public inline
// void setAnchorName(const class QString &)
func (this *QTextCharFormat) SetAnchorName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat13setAnchorNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:543
// index:0
// Public
// QString anchorName()
func (this *QTextCharFormat) AnchorName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat10anchorNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:545
// index:0
// Public inline
// void setAnchorNames(const class QStringList &)
func (this *QTextCharFormat) SetAnchorNames(names *qtcore.QStringList) {
	var convArg0 = names.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat14setAnchorNamesERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:549
// index:0
// Public inline
// void setTableCellRowSpan(int)
func (this *QTextCharFormat) SetTableCellRowSpan(tableCellRowSpan int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat19setTableCellRowSpanEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &tableCellRowSpan)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:550
// index:0
// Public inline
// int tableCellRowSpan()
func (this *QTextCharFormat) TableCellRowSpan() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat16tableCellRowSpanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qtextformat.h:552
// index:0
// Public inline
// void setTableCellColumnSpan(int)
func (this *QTextCharFormat) SetTableCellColumnSpan(tableCellColumnSpan int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QTextCharFormat22setTableCellColumnSpanEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &tableCellColumnSpan)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:553
// index:0
// Public inline
// int tableCellColumnSpan()
func (this *QTextCharFormat) TableCellColumnSpan() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QTextCharFormat19tableCellColumnSpanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

//  body block end
