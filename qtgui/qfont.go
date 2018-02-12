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
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QFont struct {
	*qtrt.CObject
}

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
// [-2] void QFont(const QString &, int, int, _Bool)
func NewQFont_1(family string, pointSize int, weight int, italic bool) *QFont {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
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
func NewQFont_2(arg0 *QFont, pd *QPaintDevice /*777 QPaintDevice **/) *QFont {
	var convArg0 = arg0.GetCthis()
	var convArg1 = pd.GetCthis()
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
func (this *QFont) Swap(other *QFont) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QString family()
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
func (this *QFont) SetFamily(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont9setFamilyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleName()
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
func (this *QFont) SetStyleName(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setStyleNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:185
// index:0
// Public Visibility=Default Availability=Available
// [4] int pointSize()
func (this *QFont) PointSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9pointSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPointSize(int)
func (this *QFont) SetPointSize(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setPointSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal pointSizeF()
func (this *QFont) PointSizeF() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont10pointSizeFEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPointSizeF(qreal)
func (this *QFont) SetPointSizeF(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont13setPointSizeFEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:190
// index:0
// Public Visibility=Default Availability=Available
// [4] int pixelSize()
func (this *QFont) PixelSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9pixelSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelSize(int)
func (this *QFont) SetPixelSize(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setPixelSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:193
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight()
func (this *QFont) Weight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont6weightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:194
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWeight(int)
func (this *QFont) SetWeight(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont9setWeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:196
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool bold()
func (this *QFont) Bold() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont4boldEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:197
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBold(_Bool)
func (this *QFont) SetBold(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont7setBoldEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(enum QFont::Style)
func (this *QFont) SetStyle(style int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont8setStyleENS_5StyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), style)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:200
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Style style()
func (this *QFont) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:202
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool italic()
func (this *QFont) Italic() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont6italicEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:203
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setItalic(_Bool)
func (this *QFont) SetItalic(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont9setItalicEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:205
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underline()
func (this *QFont) Underline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9underlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUnderline(_Bool)
func (this *QFont) SetUnderline(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setUnderlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:208
// index:0
// Public Visibility=Default Availability=Available
// [1] bool overline()
func (this *QFont) Overline() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont8overlineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:209
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOverline(_Bool)
func (this *QFont) SetOverline(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont11setOverlineEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool strikeOut()
func (this *QFont) StrikeOut() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9strikeOutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStrikeOut(_Bool)
func (this *QFont) SetStrikeOut(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setStrikeOutEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:214
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fixedPitch()
func (this *QFont) FixedPitch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont10fixedPitchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedPitch(_Bool)
func (this *QFont) SetFixedPitch(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont13setFixedPitchEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:217
// index:0
// Public Visibility=Default Availability=Available
// [1] bool kerning()
func (this *QFont) Kerning() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7kerningEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKerning(_Bool)
func (this *QFont) SetKerning(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setKerningEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:220
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::StyleHint styleHint()
func (this *QFont) StyleHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont9styleHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:221
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::StyleStrategy styleStrategy()
func (this *QFont) StyleStrategy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont13styleStrategyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:222
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleHint(enum QFont::StyleHint, enum QFont::StyleStrategy)
func (this *QFont) SetStyleHint(arg0 int, arg1 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont12setStyleHintENS_9StyleHintENS_13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:223
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleStrategy(enum QFont::StyleStrategy)
func (this *QFont) SetStyleStrategy(s int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont16setStyleStrategyENS_13StyleStrategyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), s)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:225
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretch()
func (this *QFont) Stretch() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7stretchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretch(int)
func (this *QFont) SetStretch(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal letterSpacing()
func (this *QFont) LetterSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont13letterSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:229
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::SpacingType letterSpacingType()
func (this *QFont) LetterSpacingType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont17letterSpacingTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLetterSpacing(enum QFont::SpacingType, qreal)
func (this *QFont) SetLetterSpacing(type_ int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont16setLetterSpacingENS_11SpacingTypeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal wordSpacing()
func (this *QFont) WordSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont11wordSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWordSpacing(qreal)
func (this *QFont) SetWordSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont14setWordSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCapitalization(enum QFont::Capitalization)
func (this *QFont) SetCapitalization(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont17setCapitalizationENS_14CapitalizationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:236
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Capitalization capitalization()
func (this *QFont) Capitalization() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont14capitalizationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHintingPreference(enum QFont::HintingPreference)
func (this *QFont) SetHintingPreference(hintingPreference int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont20setHintingPreferenceENS_17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hintingPreference)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:239
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::HintingPreference hintingPreference()
func (this *QFont) HintingPreference() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont17hintingPreferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:242
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rawMode()
func (this *QFont) RawMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7rawModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawMode(_Bool)
func (this *QFont) SetRawMode(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setRawModeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:247
// index:0
// Public Visibility=Default Availability=Available
// [1] bool exactMatch()
func (this *QFont) ExactMatch() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont10exactMatchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isCopyOf(const QFont &)
func (this *QFont) IsCopyOf(arg0 *QFont) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont8isCopyOfERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRawName(const QString &)
func (this *QFont) SetRawName(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont10setRawNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QString rawName()
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
// [8] QString key()
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
// [8] QString toString()
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
func (this *QFont) InsertSubstitutions(arg0 string, arg1 *qtcore.QStringList) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}
func QFont_InsertSubstitutions(arg0 string, arg1 *qtcore.QStringList) {
	var nilthis *QFont
	nilthis.InsertSubstitutions(arg0, arg1)
}

// /usr/include/qt/QtGui/qfont.h:276
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void removeSubstitutions(const QString &)
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
// [8] QString defaultFamily()
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
// [8] QString lastResortFamily()
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
// [8] QString lastResortFont()
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
// [16] QFont resolve(const QFont &)
func (this *QFont) Resolve(arg0 *QFont) *QFont /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7resolveERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qfont.h:289
// index:1
// Public inline Visibility=Default Availability=Available
// [4] uint resolve()
func (this *QFont) Resolve_1() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QFont7resolveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qfont.h:290
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void resolve(uint)
func (this *QFont) Resolve_2(mask uint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QFont7resolveEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mask)
	qtrt.ErrPrint(err, rv)
}

type QFont__StyleHint = int

const QFont__Helvetica QFont__StyleHint = 0
const QFont__SansSerif QFont__StyleHint = 0
const QFont__Times QFont__StyleHint = 1
const QFont__Serif QFont__StyleHint = 1
const QFont__Courier QFont__StyleHint = 2
const QFont__TypeWriter QFont__StyleHint = 2
const QFont__OldEnglish QFont__StyleHint = 3
const QFont__Decorative QFont__StyleHint = 3
const QFont__System QFont__StyleHint = 4
const QFont__AnyStyle QFont__StyleHint = 5
const QFont__Cursive QFont__StyleHint = 6
const QFont__Monospace QFont__StyleHint = 7
const QFont__Fantasy QFont__StyleHint = 8

type QFont__StyleStrategy = int

const QFont__PreferDefault QFont__StyleStrategy = 1
const QFont__PreferBitmap QFont__StyleStrategy = 2
const QFont__PreferDevice QFont__StyleStrategy = 4
const QFont__PreferOutline QFont__StyleStrategy = 8
const QFont__ForceOutline QFont__StyleStrategy = 16
const QFont__PreferMatch QFont__StyleStrategy = 32
const QFont__PreferQuality QFont__StyleStrategy = 64
const QFont__PreferAntialias QFont__StyleStrategy = 128
const QFont__NoAntialias QFont__StyleStrategy = 256
const QFont__OpenGLCompatible QFont__StyleStrategy = 512
const QFont__ForceIntegerMetrics QFont__StyleStrategy = 1024
const QFont__NoSubpixelAntialias QFont__StyleStrategy = 2048
const QFont__PreferNoShaping QFont__StyleStrategy = 4096
const QFont__NoFontMerging QFont__StyleStrategy = 32768

type QFont__HintingPreference = int

const QFont__PreferDefaultHinting QFont__HintingPreference = 0
const QFont__PreferNoHinting QFont__HintingPreference = 1
const QFont__PreferVerticalHinting QFont__HintingPreference = 2
const QFont__PreferFullHinting QFont__HintingPreference = 3

type QFont__Weight = int

const QFont__Thin QFont__Weight = 0
const QFont__ExtraLight QFont__Weight = 12
const QFont__Light QFont__Weight = 25
const QFont__Normal QFont__Weight = 50
const QFont__Medium QFont__Weight = 57
const QFont__DemiBold QFont__Weight = 63
const QFont__Bold QFont__Weight = 75
const QFont__ExtraBold QFont__Weight = 81
const QFont__Black QFont__Weight = 87

type QFont__Style = int

const QFont__StyleNormal QFont__Style = 0
const QFont__StyleItalic QFont__Style = 1
const QFont__StyleOblique QFont__Style = 2

type QFont__Stretch = int

const QFont__AnyStretch QFont__Stretch = 0
const QFont__UltraCondensed QFont__Stretch = 50
const QFont__ExtraCondensed QFont__Stretch = 62
const QFont__Condensed QFont__Stretch = 75
const QFont__SemiCondensed QFont__Stretch = 87
const QFont__Unstretched QFont__Stretch = 100
const QFont__SemiExpanded QFont__Stretch = 112
const QFont__Expanded QFont__Stretch = 125
const QFont__ExtraExpanded QFont__Stretch = 150
const QFont__UltraExpanded QFont__Stretch = 200

type QFont__Capitalization = int

const QFont__MixedCase QFont__Capitalization = 0
const QFont__AllUppercase QFont__Capitalization = 1
const QFont__AllLowercase QFont__Capitalization = 2
const QFont__SmallCaps QFont__Capitalization = 3
const QFont__Capitalize QFont__Capitalization = 4

type QFont__SpacingType = int

const QFont__PercentageSpacing QFont__SpacingType = 0
const QFont__AbsoluteSpacing QFont__SpacingType = 1

type QFont__ResolveProperties = int

const QFont__FamilyResolved QFont__ResolveProperties = 1
const QFont__SizeResolved QFont__ResolveProperties = 2
const QFont__StyleHintResolved QFont__ResolveProperties = 4
const QFont__StyleStrategyResolved QFont__ResolveProperties = 8
const QFont__WeightResolved QFont__ResolveProperties = 16
const QFont__StyleResolved QFont__ResolveProperties = 32
const QFont__UnderlineResolved QFont__ResolveProperties = 64
const QFont__OverlineResolved QFont__ResolveProperties = 128
const QFont__StrikeOutResolved QFont__ResolveProperties = 256
const QFont__FixedPitchResolved QFont__ResolveProperties = 512
const QFont__StretchResolved QFont__ResolveProperties = 1024
const QFont__KerningResolved QFont__ResolveProperties = 2048
const QFont__CapitalizationResolved QFont__ResolveProperties = 4096
const QFont__LetterSpacingResolved QFont__ResolveProperties = 8192
const QFont__WordSpacingResolved QFont__ResolveProperties = 16384
const QFont__HintingPreferenceResolved QFont__ResolveProperties = 32768
const QFont__StyleNameResolved QFont__ResolveProperties = 65536
const QFont__AllPropertiesResolved QFont__ResolveProperties = 131071

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
