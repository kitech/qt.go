//  header block begin
// /usr/include/qt/QtGui/qfont.h
// #include <qfont.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QFont struct {
	*qtrt.CObject
}

func (this *QFont) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qfont.h:170
// index:0
// void QFont()
func NewQFont() *QFont {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFontC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(cthis)
	return gothis
}
func NewQFontFromPointer(cthis unsafe.Pointer) *QFont {
	return &QFont{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qfont.h:171
// index:1
// void QFont(const class QString &, int, int, _Bool)
func NewQFont_1(family unsafe.Pointer, pointSize int, weight int, italic bool) *QFont {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFontC2ERK7QStringiib", ffiqt.FFI_TYPE_VOID, cthis, family, &pointSize, &weight, &italic)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:172
// index:2
// void QFont(const class QFont &, class QPaintDevice *)
func NewQFont_2(arg0 unsafe.Pointer, pd unsafe.Pointer) *QFont {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFontC2ERKS_P12QPaintDevice", ffiqt.FFI_TYPE_VOID, cthis, arg0, pd)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:174
// index:0
// void ~QFont()
func DeleteQFont(*QFont) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFontD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:176
// index:0
// inline
// void swap(class QFont &)
func (this *QFont) Swap(other unsafe.Pointer) {
	// 0: (, other QFont &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:179
// index:0
// QString family()
func (this *QFont) Family() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont6familyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:180
// index:0
// void setFamily(const class QString &)
func (this *QFont) SetFamily(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont9setFamilyERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:182
// index:0
// QString styleName()
func (this *QFont) StyleName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont9styleNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:183
// index:0
// void setStyleName(const class QString &)
func (this *QFont) SetStyleName(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont12setStyleNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:185
// index:0
// int pointSize()
func (this *QFont) PointSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont9pointSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:186
// index:0
// void setPointSize(int)
func (this *QFont) SetPointSize(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont12setPointSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:187
// index:0
// qreal pointSizeF()
func (this *QFont) PointSizeF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont10pointSizeFEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:188
// index:0
// void setPointSizeF(qreal)
func (this *QFont) SetPointSizeF(arg0 float64) {
	// 0: (, qreal arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont13setPointSizeFEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:190
// index:0
// int pixelSize()
func (this *QFont) PixelSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont9pixelSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:191
// index:0
// void setPixelSize(int)
func (this *QFont) SetPixelSize(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont12setPixelSizeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:193
// index:0
// int weight()
func (this *QFont) Weight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont6weightEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:194
// index:0
// void setWeight(int)
func (this *QFont) SetWeight(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont9setWeightEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:196
// index:0
// inline
// bool bold()
func (this *QFont) Bold() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont4boldEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:197
// index:0
// inline
// void setBold(_Bool)
func (this *QFont) SetBold(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont7setBoldEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:199
// index:0
// void setStyle(enum QFont::Style)
func (this *QFont) SetStyle(style int) {
	// 0: (, style QFont::Style), (&style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont8setStyleENS_5StyleE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:200
// index:0
// QFont::Style style()
func (this *QFont) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont5styleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:202
// index:0
// inline
// bool italic()
func (this *QFont) Italic() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont6italicEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:203
// index:0
// inline
// void setItalic(_Bool)
func (this *QFont) SetItalic(b bool) {
	// 0: (, b bool), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont9setItalicEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:205
// index:0
// bool underline()
func (this *QFont) Underline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont9underlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:206
// index:0
// void setUnderline(_Bool)
func (this *QFont) SetUnderline(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont12setUnderlineEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:208
// index:0
// bool overline()
func (this *QFont) Overline() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont8overlineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:209
// index:0
// void setOverline(_Bool)
func (this *QFont) SetOverline(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont11setOverlineEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:211
// index:0
// bool strikeOut()
func (this *QFont) StrikeOut() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont9strikeOutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:212
// index:0
// void setStrikeOut(_Bool)
func (this *QFont) SetStrikeOut(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont12setStrikeOutEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:214
// index:0
// bool fixedPitch()
func (this *QFont) FixedPitch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont10fixedPitchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:215
// index:0
// void setFixedPitch(_Bool)
func (this *QFont) SetFixedPitch(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont13setFixedPitchEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:217
// index:0
// bool kerning()
func (this *QFont) Kerning() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont7kerningEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:218
// index:0
// void setKerning(_Bool)
func (this *QFont) SetKerning(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont10setKerningEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:220
// index:0
// QFont::StyleHint styleHint()
func (this *QFont) StyleHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont9styleHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:221
// index:0
// QFont::StyleStrategy styleStrategy()
func (this *QFont) StyleStrategy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont13styleStrategyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:222
// index:0
// void setStyleHint(enum QFont::StyleHint, enum QFont::StyleStrategy)
func (this *QFont) SetStyleHint(arg0 int, arg1 int) {
	// 0: (, QFont::StyleHint arg0, QFont::StyleStrategy arg1), (&arg0, &arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont12setStyleHintENS_9StyleHintENS_13StyleStrategyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0, &arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:223
// index:0
// void setStyleStrategy(enum QFont::StyleStrategy)
func (this *QFont) SetStyleStrategy(s int) {
	// 0: (, s QFont::StyleStrategy), (&s)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont16setStyleStrategyENS_13StyleStrategyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &s)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:225
// index:0
// int stretch()
func (this *QFont) Stretch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont7stretchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:226
// index:0
// void setStretch(int)
func (this *QFont) SetStretch(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont10setStretchEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:228
// index:0
// qreal letterSpacing()
func (this *QFont) LetterSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont13letterSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:229
// index:0
// QFont::SpacingType letterSpacingType()
func (this *QFont) LetterSpacingType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont17letterSpacingTypeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:230
// index:0
// void setLetterSpacing(enum QFont::SpacingType, qreal)
func (this *QFont) SetLetterSpacing(type_ int, spacing float64) {
	// 0: (, type QFont::SpacingType, spacing qreal), (&type_, &spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont16setLetterSpacingENS_11SpacingTypeEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &type_, &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:232
// index:0
// qreal wordSpacing()
func (this *QFont) WordSpacing() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont11wordSpacingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:233
// index:0
// void setWordSpacing(qreal)
func (this *QFont) SetWordSpacing(spacing float64) {
	// 0: (, spacing qreal), (&spacing)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont14setWordSpacingEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &spacing)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:235
// index:0
// void setCapitalization(enum QFont::Capitalization)
func (this *QFont) SetCapitalization(arg0 int) {
	// 0: (, QFont::Capitalization arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont17setCapitalizationENS_14CapitalizationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:236
// index:0
// QFont::Capitalization capitalization()
func (this *QFont) Capitalization() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont14capitalizationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:238
// index:0
// void setHintingPreference(enum QFont::HintingPreference)
func (this *QFont) SetHintingPreference(hintingPreference int) {
	// 0: (, hintingPreference QFont::HintingPreference), (&hintingPreference)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont20setHintingPreferenceENS_17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:239
// index:0
// QFont::HintingPreference hintingPreference()
func (this *QFont) HintingPreference() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont17hintingPreferenceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:242
// index:0
// bool rawMode()
func (this *QFont) RawMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont7rawModeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:243
// index:0
// void setRawMode(_Bool)
func (this *QFont) SetRawMode(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont10setRawModeEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:247
// index:0
// bool exactMatch()
func (this *QFont) ExactMatch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont10exactMatchEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:254
// index:0
// bool isCopyOf(const class QFont &)
func (this *QFont) IsCopyOf(arg0 unsafe.Pointer) {
	// 0: (, const QFont & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont8isCopyOfERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:262
// index:0
// void setRawName(const class QString &)
func (this *QFont) SetRawName(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont10setRawNameERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:263
// index:0
// QString rawName()
func (this *QFont) RawName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont7rawNameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:266
// index:0
// QString key()
func (this *QFont) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont3keyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:268
// index:0
// QString toString()
func (this *QFont) ToString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont8toStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:269
// index:0
// bool fromString(const class QString &)
func (this *QFont) FromString(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont10fromStringERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:271
// index:0
// static
// QString substitute(const class QString &)
func (this *QFont) Substitute(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont10substituteERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_Substitute(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	var nilthis *QFont
	nilthis.Substitute(arg0)
}

// /usr/include/qt/QtGui/qfont.h:272
// index:0
// static
// QStringList substitutes(const class QString &)
func (this *QFont) Substitutes(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont11substitutesERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_Substitutes(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	var nilthis *QFont
	nilthis.Substitutes(arg0)
}

// /usr/include/qt/QtGui/qfont.h:273
// index:0
// static
// QStringList substitutions()
func (this *QFont) Substitutions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont13substitutionsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_Substitutions() {
	// 0: (), ()
	var nilthis *QFont
	nilthis.Substitutions()
}

// /usr/include/qt/QtGui/qfont.h:274
// index:0
// static
// void insertSubstitution(const class QString &, const class QString &)
func (this *QFont) InsertSubstitution(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (const QString & arg0, const QString & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont18insertSubstitutionERK7QStringS2_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_InsertSubstitution(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (const QString & arg0, const QString & arg1), (arg0, arg1)
	var nilthis *QFont
	nilthis.InsertSubstitution(arg0, arg1)
}

// /usr/include/qt/QtGui/qfont.h:275
// index:0
// static
// void insertSubstitutions(const class QString &, const class QStringList &)
func (this *QFont) InsertSubstitutions(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (const QString & arg0, const QStringList & arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont19insertSubstitutionsERK7QStringRK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_InsertSubstitutions(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (const QString & arg0, const QStringList & arg1), (arg0, arg1)
	var nilthis *QFont
	nilthis.InsertSubstitutions(arg0, arg1)
}

// /usr/include/qt/QtGui/qfont.h:276
// index:0
// static
// void removeSubstitutions(const class QString &)
func (this *QFont) RemoveSubstitutions(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont19removeSubstitutionsERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_RemoveSubstitutions(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	var nilthis *QFont
	nilthis.RemoveSubstitutions(arg0)
}

// /usr/include/qt/QtGui/qfont.h:280
// index:0
// static
// void initialize()
func (this *QFont) Initialize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont10initializeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_Initialize() {
	// 0: (), ()
	var nilthis *QFont
	nilthis.Initialize()
}

// /usr/include/qt/QtGui/qfont.h:281
// index:0
// static
// void cleanup()
func (this *QFont) Cleanup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont7cleanupEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_Cleanup() {
	// 0: (), ()
	var nilthis *QFont
	nilthis.Cleanup()
}

// /usr/include/qt/QtGui/qfont.h:282
// index:0
// static
// void cacheStatistics()
func (this *QFont) CacheStatistics() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont15cacheStatisticsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFont_CacheStatistics() {
	// 0: (), ()
	var nilthis *QFont
	nilthis.CacheStatistics()
}

// /usr/include/qt/QtGui/qfont.h:284
// index:0
// QString defaultFamily()
func (this *QFont) DefaultFamily() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont13defaultFamilyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:285
// index:0
// QString lastResortFamily()
func (this *QFont) LastResortFamily() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont16lastResortFamilyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:286
// index:0
// QString lastResortFont()
func (this *QFont) LastResortFont() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont14lastResortFontEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:288
// index:0
// QFont resolve(const class QFont &)
func (this *QFont) Resolve(arg0 unsafe.Pointer) {
	// 0: (, const QFont & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont7resolveERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:289
// index:1
// inline
// uint resolve()
func (this *QFont) Resolve_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QFont7resolveEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:290
// index:2
// inline
// void resolve(uint)
func (this *QFont) Resolve_2(mask uint) {
	// 2: (, mask uint), (&mask)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QFont7resolveEj", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mask)
	gopp.ErrPrint(err, rv)
}

//  body block end
