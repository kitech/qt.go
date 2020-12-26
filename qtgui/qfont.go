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
// extern C begin: 0
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
// size 16
type QFont struct {
	*qtrt.CObject
}
type QFont_ITF interface {
	QFont_PTR() *QFont
}

func (ptr *QFont) QFont_PTR() *QFont { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
// ignore GetCthis for 0 base
func QFontFromptr(cthis Voidptr) *QFont {
	return &QFont{&qtrt.CObject{cthis}}
}
func (*QFont) Fromptr(cthis Voidptr) *QFont {
	return QFontFromptr(cthis)
}

// /usr/include/qt/QtGui/qfont.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
 */
func (*QFont) NewForInherit(family string, pointSize int, weight int, italic bool) *QFont {
	return NewQFont(family, pointSize, weight, italic)
}
func NewQFont(family string, pointSize int, weight int, italic bool) *QFont {
	var tmpArg0 = qtcore.NewQString5(family)
	var convArg0 = tmpArg0.GetCthis()
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1453900756, "_ZN5QFontC2ERK7QStringiib", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&pointSize), Voidptr(&weight), Voidptr(&italic))
	qtrt.ErrPrint2(err, rv)
	gothis := QFontFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
 */
func (*QFont) NewForInheritp(family string) *QFont {
	return NewQFontp(family)
}
func NewQFontp(family string) *QFont {
	var tmpArg0 = qtcore.NewQString5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	pointSize := int(-1)
	// arg: 2, int=Int, =Invalid, , Invalid
	weight := int(-1)
	// arg: 3, bool=Bool, =Invalid, , Invalid
	italic := false
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1453900756, "_ZN5QFontC2ERK7QStringiib", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&pointSize), Voidptr(&weight), Voidptr(&italic))
	qtrt.ErrPrint2(err, rv)
	gothis := QFontFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
 */
func (*QFont) NewForInheritp1(family string, pointSize int) *QFont {
	return NewQFontp1(family, pointSize)
}
func NewQFontp1(family string, pointSize int) *QFont {
	var tmpArg0 = qtcore.NewQString5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, int=Int, =Invalid, , Invalid
	weight := int(-1)
	// arg: 3, bool=Bool, =Invalid, , Invalid
	italic := false
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1453900756, "_ZN5QFontC2ERK7QStringiib", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&pointSize), Voidptr(&weight), Voidptr(&italic))
	qtrt.ErrPrint2(err, rv)
	gothis := QFontFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFont(const QString &, int, int, bool)

/*
 */
func (*QFont) NewForInheritp2(family string, pointSize int, weight int) *QFont {
	return NewQFontp2(family, pointSize, weight)
}
func NewQFontp2(family string, pointSize int, weight int) *QFont {
	var tmpArg0 = qtcore.NewQString5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 3, bool=Bool, =Invalid, , Invalid
	italic := false
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1453900756, "_ZN5QFontC2ERK7QStringiib", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0), Voidptr(&pointSize), Voidptr(&weight), Voidptr(&italic))
	qtrt.ErrPrint2(err, rv)
	gothis := QFontFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQFont)
	return gothis
}

// /usr/include/qt/QtGui/qfont.h:187
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString family() const

/*
 */
func (this *QFont) Family() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2393185171, "_ZNK5QFont6familyEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:188
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFamily(const QString &)

/*
 */
func (this *QFont) SetFamily(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(3823958997, "_ZN5QFont9setFamilyERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:193
// index:0
// Public Indirect Visibility=Default Availability=Available
// [8] QString styleName() const

/*
 */
func (this *QFont) StyleName() string {
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2670374555, "_ZNK5QFont9styleNameEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr())
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := qtcore.QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfont.h:194
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStyleName(const QString &)

/*
 */
func (this *QFont) SetStyleName(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc3(814338029, "_ZN5QFont12setStyleNameERK7QString", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:196
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int pointSize() const

/*
 */
func (this *QFont) PointSize() int {
	rv, err := qtrt.Qtcc3(3023852881, "_ZNK5QFont9pointSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:197
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPointSize(int)

/*
 */
func (this *QFont) SetPointSize(arg0 int) {
	rv, err := qtrt.Qtcc3(1776764880, "_ZN5QFont12setPointSizeEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:198
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qreal pointSizeF() const

/*
 */
func (this *QFont) PointSizeF() float64 {
	rv, err := qtrt.Qtcc3(314855354, "_ZNK5QFont10pointSizeFEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:199
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPointSizeF(qreal)

/*
 */
func (this *QFont) SetPointSizeF(arg0 float64) {
	rv, err := qtrt.Qtcc3(2339912505, "_ZN5QFont13setPointSizeFEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:201
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int pixelSize() const

/*
 */
func (this *QFont) PixelSize() int {
	rv, err := qtrt.Qtcc3(197311963, "_ZNK5QFont9pixelSizeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:202
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setPixelSize(int)

/*
 */
func (this *QFont) SetPixelSize(arg0 int) {
	rv, err := qtrt.Qtcc3(3592007514, "_ZN5QFont12setPixelSizeEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:204
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int weight() const

/*
 */
func (this *QFont) Weight() int {
	rv, err := qtrt.Qtcc3(1776350790, "_ZNK5QFont6weightEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:205
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWeight(int)

/*
 */
func (this *QFont) SetWeight(arg0 int) {
	rv, err := qtrt.Qtcc3(133781457, "_ZN5QFont9setWeightEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:207
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool bold() const

/*
 */
func (this *QFont) Bold() bool {
	rv, err := qtrt.Qtcc3(1200580973, "_ZNK5QFont4boldEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:208
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setBold(bool)

/*
 */
func (this *QFont) SetBold(arg0 bool) {
	rv, err := qtrt.Qtcc3(1947504972, "_ZN5QFont7setBoldEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:210
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStyle(QFont::Style)

/*
 */
func (this *QFont) SetStyle(style int) {
	rv, err := qtrt.Qtcc3(2307513941, "_ZN5QFont8setStyleENS_5StyleE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&style))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:211
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QFont::Style style() const

/*
 */
func (this *QFont) Style() int {
	rv, err := qtrt.Qtcc3(2644565881, "_ZNK5QFont5styleEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:213
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool italic() const

/*
 */
func (this *QFont) Italic() bool {
	rv, err := qtrt.Qtcc3(4214105033, "_ZNK5QFont6italicEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:214
// index:0
// Public inline Ignore Visibility=Default Availability=Available
// [-2] void setItalic(bool)

/*
 */
func (this *QFont) SetItalic(b bool) {
	rv, err := qtrt.Qtcc3(48585686, "_ZN5QFont9setItalicEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&b))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:216
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool underline() const

/*
 */
func (this *QFont) Underline() bool {
	rv, err := qtrt.Qtcc3(3473615270, "_ZNK5QFont9underlineEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:217
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setUnderline(bool)

/*
 */
func (this *QFont) SetUnderline(arg0 bool) {
	rv, err := qtrt.Qtcc3(2231558831, "_ZN5QFont12setUnderlineEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:219
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool overline() const

/*
 */
func (this *QFont) Overline() bool {
	rv, err := qtrt.Qtcc3(2096479559, "_ZNK5QFont8overlineEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:220
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setOverline(bool)

/*
 */
func (this *QFont) SetOverline(arg0 bool) {
	rv, err := qtrt.Qtcc3(229563217, "_ZN5QFont11setOverlineEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:222
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool strikeOut() const

/*
 */
func (this *QFont) StrikeOut() bool {
	rv, err := qtrt.Qtcc3(2658375640, "_ZNK5QFont9strikeOutEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:223
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStrikeOut(bool)

/*
 */
func (this *QFont) SetStrikeOut(arg0 bool) {
	rv, err := qtrt.Qtcc3(3564788945, "_ZN5QFont12setStrikeOutEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:225
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool fixedPitch() const

/*
 */
func (this *QFont) FixedPitch() bool {
	rv, err := qtrt.Qtcc3(3065387023, "_ZNK5QFont10fixedPitchEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:226
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setFixedPitch(bool)

/*
 */
func (this *QFont) SetFixedPitch(arg0 bool) {
	rv, err := qtrt.Qtcc3(3328821689, "_ZN5QFont13setFixedPitchEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:228
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool kerning() const

/*
 */
func (this *QFont) Kerning() bool {
	rv, err := qtrt.Qtcc3(3727955333, "_ZNK5QFont7kerningEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfont.h:229
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setKerning(bool)

/*
 */
func (this *QFont) SetKerning(arg0 bool) {
	rv, err := qtrt.Qtcc3(43048096, "_ZN5QFont10setKerningEb", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:231
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QFont::StyleHint styleHint() const

/*
 */
func (this *QFont) StyleHint() int {
	rv, err := qtrt.Qtcc3(1985110254, "_ZNK5QFont9styleHintEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:232
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QFont::StyleStrategy styleStrategy() const

/*
 */
func (this *QFont) StyleStrategy() int {
	rv, err := qtrt.Qtcc3(2409632431, "_ZNK5QFont13styleStrategyEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:233
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStyleHint(QFont::StyleHint, QFont::StyleStrategy)

/*
 */
func (this *QFont) SetStyleHint(arg0 int, arg1 int) {
	rv, err := qtrt.Qtcc3(4164244041, "_ZN5QFont12setStyleHintENS_9StyleHintENS_13StyleStrategyE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0), Voidptr(&arg1))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:233
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStyleHint(QFont::StyleHint, QFont::StyleStrategy)

/*
 */
func (this *QFont) SetStyleHintp(arg0 int) {
	// arg: 1, QFont::StyleStrategy=Enum, QFont::StyleStrategy=Enum, , Invalid
	arg1 := 0
	rv, err := qtrt.Qtcc3(4164244041, "_ZN5QFont12setStyleHintENS_9StyleHintENS_13StyleStrategyE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0), Voidptr(&arg1))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:234
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStyleStrategy(QFont::StyleStrategy)

/*
 */
func (this *QFont) SetStyleStrategy(s int) {
	rv, err := qtrt.Qtcc3(3706440996, "_ZN5QFont16setStyleStrategyENS_13StyleStrategyE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&s))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:236
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] int stretch() const

/*
 */
func (this *QFont) Stretch() int {
	rv, err := qtrt.Qtcc3(1924986953, "_ZNK5QFont7stretchEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfont.h:237
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setStretch(int)

/*
 */
func (this *QFont) SetStretch(arg0 int) {
	rv, err := qtrt.Qtcc3(969599204, "_ZN5QFont10setStretchEi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:239
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qreal letterSpacing() const

/*
 */
func (this *QFont) LetterSpacing() float64 {
	rv, err := qtrt.Qtcc3(2641819184, "_ZNK5QFont13letterSpacingEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:240
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QFont::SpacingType letterSpacingType() const

/*
 */
func (this *QFont) LetterSpacingType() int {
	rv, err := qtrt.Qtcc3(2189521662, "_ZNK5QFont17letterSpacingTypeEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:241
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setLetterSpacing(QFont::SpacingType, qreal)

/*
 */
func (this *QFont) SetLetterSpacing(type_ int, spacing float64) {
	rv, err := qtrt.Qtcc3(4000308159, "_ZN5QFont16setLetterSpacingENS_11SpacingTypeEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&type_), Voidptr(&spacing))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:243
// index:0
// Public Direct Visibility=Default Availability=Available
// [8] qreal wordSpacing() const

/*
 */
func (this *QFont) WordSpacing() float64 {
	rv, err := qtrt.Qtcc3(3552615348, "_ZNK5QFont11wordSpacingEv", qtrt.FFITO_DOUBLE,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qfont.h:244
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setWordSpacing(qreal)

/*
 */
func (this *QFont) SetWordSpacing(spacing float64) {
	rv, err := qtrt.Qtcc3(2436393468, "_ZN5QFont14setWordSpacingEd", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&spacing))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:246
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setCapitalization(QFont::Capitalization)

/*
 */
func (this *QFont) SetCapitalization(arg0 int) {
	rv, err := qtrt.Qtcc3(2579133516, "_ZN5QFont17setCapitalizationENS_14CapitalizationE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&arg0))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:247
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QFont::Capitalization capitalization() const

/*
 */
func (this *QFont) Capitalization() int {
	rv, err := qtrt.Qtcc3(74477930, "_ZNK5QFont14capitalizationEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:249
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setHintingPreference(QFont::HintingPreference)

/*
 */
func (this *QFont) SetHintingPreference(hintingPreference int) {
	rv, err := qtrt.Qtcc3(2526211871, "_ZN5QFont20setHintingPreferenceENS_17HintingPreferenceE", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&hintingPreference))
	qtrt.ErrPrint2(err, rv)
}

// /usr/include/qt/QtGui/qfont.h:250
// index:0
// Public Direct Visibility=Default Availability=Available
// [4] QFont::HintingPreference hintingPreference() const

/*
 */
func (this *QFont) HintingPreference() int {
	rv, err := qtrt.Qtcc3(1448856926, "_ZNK5QFont17hintingPreferenceEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qfont.h:258
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool exactMatch() const

/*
 */
func (this *QFont) ExactMatch() bool {
	rv, err := qtrt.Qtcc3(2416987156, "_ZNK5QFont10exactMatchEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint2(err, rv)
	return rv != 0
}

func DeleteQFont(this *QFont) {
	rv, err := qtrt.Qtcc1(0, "_ZN5QFontD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QFont__StyleHint = int

//
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


 */
type QFont__HintingPreference = int

//
const QFont__PreferDefaultHinting QFont__HintingPreference = 0

//
const QFont__PreferNoHinting QFont__HintingPreference = 1

//
const QFont__PreferVerticalHinting QFont__HintingPreference = 2

//
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


 */
type QFont__Weight = int

//
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


 */
type QFont__Style = int

//
const QFont__StyleNormal QFont__Style = 0

//
const QFont__StyleItalic QFont__Style = 1

//
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


 */
type QFont__Capitalization = int

//
const QFont__MixedCase QFont__Capitalization = 0

//
const QFont__AllUppercase QFont__Capitalization = 1

//
const QFont__AllLowercase QFont__Capitalization = 2

//
const QFont__SmallCaps QFont__Capitalization = 3

//
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


 */
type QFont__SpacingType = int

//
const QFont__PercentageSpacing QFont__SpacingType = 0

//
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
const QFont__NoPropertiesResolved QFont__ResolveProperties = 0

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
const QFont__FamiliesResolved QFont__ResolveProperties = 131072

//
const QFont__AllPropertiesResolved QFont__ResolveProperties = 262143

func (this *QFont) ResolvePropertiesItemName(val int) string {
	switch val {
	case QFont__NoPropertiesResolved: // 0
		return "NoPropertiesResolved"
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
	case QFont__FamiliesResolved: // 131072
		return "FamiliesResolved"
	case QFont__AllPropertiesResolved: // 262143
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

func init_unused_10135() {
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
