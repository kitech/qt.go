//  header block begin
// /usr/include/qt/QtGui/qrawfont.h
// #include <qrawfont.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 29
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
type QRawFont struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qrawfont.h:74
// index:0
// void QRawFont()
func NewQRawFont() *QRawFont {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QRawFont{cthis}
}

// /usr/include/qt/QtGui/qrawfont.h:75
// index:1
// void QRawFont(const class QString &, qreal, class QFont::HintingPreference)
func NewQRawFont_1(fileName unsafe.Pointer, pixelSize float64, hintingPreference int) *QRawFont {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2ERK7QStringdN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, cthis, fileName, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
	return &QRawFont{cthis}
}

// /usr/include/qt/QtGui/qrawfont.h:78
// index:2
// void QRawFont(const class QByteArray &, qreal, class QFont::HintingPreference)
func NewQRawFont_2(fontData unsafe.Pointer, pixelSize float64, hintingPreference int) *QRawFont {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2ERK10QByteArraydN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, cthis, fontData, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
	return &QRawFont{cthis}
}

// /usr/include/qt/QtGui/qrawfont.h:86
// index:0
// void ~QRawFont()
func DeleteQRawFont(*QRawFont) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:88
// index:0
// inline
// void swap(class QRawFont &)
func (this *QRawFont) Swap(other unsafe.Pointer) {
	// 0: (, QRawFont & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:90
// index:0
// bool isValid()
func (this *QRawFont) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:96
// index:0
// QString familyName()
func (this *QRawFont) FamilyName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont10familyNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:97
// index:0
// QString styleName()
func (this *QRawFont) StyleName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9styleNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:99
// index:0
// QFont::Style style()
func (this *QRawFont) Style() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont5styleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:100
// index:0
// int weight()
func (this *QRawFont) Weight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont6weightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:102
// index:0
// QVector<quint32> glyphIndexesForString(const class QString &)
func (this *QRawFont) GlyphIndexesForString(text unsafe.Pointer) {
	// 0: (, const QString & text), (text)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont21glyphIndexesForStringERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, text)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:105
// index:0
// bool glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
func (this *QRawFont) GlyphIndexesForChars(chars unsafe.Pointer, numChars int, glyphIndexes unsafe.Pointer, numGlyphs unsafe.Pointer) {
	// 0: (, const QChar * chars, int numChars, quint32 * glyphIndexes, int * numGlyphs), (chars, &numChars, glyphIndexes, numGlyphs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi", ffiqt.FFI_TYPE_VOID, this.cthis, chars, &numChars, glyphIndexes, numGlyphs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:106
// index:0
// bool advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
func (this *QRawFont) AdvancesForGlyphIndexes(glyphIndexes unsafe.Pointer, advances unsafe.Pointer, numGlyphs int) {
	// 0: (, const quint32 * glyphIndexes, QPointF * advances, int numGlyphs), (glyphIndexes, advances, &numGlyphs)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi", ffiqt.FFI_TYPE_VOID, this.cthis, glyphIndexes, advances, &numGlyphs)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:109
// index:0
// QImage alphaMapForGlyph(quint32, enum QRawFont::AntialiasingType, const class QTransform &)
func (this *QRawFont) AlphaMapForGlyph(glyphIndex uint, antialiasingType int, transform unsafe.Pointer) {
	// 0: (, quint32 glyphIndex, QRawFont::AntialiasingType antialiasingType, const QTransform & transform), (&glyphIndex, &antialiasingType, transform)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont16alphaMapForGlyphEjNS_16AntialiasingTypeERK10QTransform", ffiqt.FFI_TYPE_VOID, this.cthis, &glyphIndex, &antialiasingType, transform)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:112
// index:0
// QPainterPath pathForGlyph(quint32)
func (this *QRawFont) PathForGlyph(glyphIndex uint) {
	// 0: (, quint32 glyphIndex), (&glyphIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12pathForGlyphEj", ffiqt.FFI_TYPE_VOID, this.cthis, &glyphIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:113
// index:0
// QRectF boundingRect(quint32)
func (this *QRawFont) BoundingRect(glyphIndex uint) {
	// 0: (, quint32 glyphIndex), (&glyphIndex)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12boundingRectEj", ffiqt.FFI_TYPE_VOID, this.cthis, &glyphIndex)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:115
// index:0
// void setPixelSize(qreal)
func (this *QRawFont) SetPixelSize(pixelSize float64) {
	// 0: (, qreal pixelSize), (&pixelSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12setPixelSizeEd", ffiqt.FFI_TYPE_VOID, this.cthis, &pixelSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:116
// index:0
// qreal pixelSize()
func (this *QRawFont) PixelSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9pixelSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:118
// index:0
// QFont::HintingPreference hintingPreference()
func (this *QRawFont) HintingPreference() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17hintingPreferenceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:120
// index:0
// qreal ascent()
func (this *QRawFont) Ascent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont6ascentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:121
// index:0
// qreal capHeight()
func (this *QRawFont) CapHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9capHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:122
// index:0
// qreal descent()
func (this *QRawFont) Descent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7descentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:123
// index:0
// qreal leading()
func (this *QRawFont) Leading() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7leadingEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:124
// index:0
// qreal xHeight()
func (this *QRawFont) XHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7xHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:125
// index:0
// qreal averageCharWidth()
func (this *QRawFont) AverageCharWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont16averageCharWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:126
// index:0
// qreal maxCharWidth()
func (this *QRawFont) MaxCharWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12maxCharWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:127
// index:0
// qreal lineThickness()
func (this *QRawFont) LineThickness() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont13lineThicknessEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:128
// index:0
// qreal underlinePosition()
func (this *QRawFont) UnderlinePosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17underlinePositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:130
// index:0
// qreal unitsPerEm()
func (this *QRawFont) UnitsPerEm() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont10unitsPerEmEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:132
// index:0
// void loadFromFile(const class QString &, qreal, class QFont::HintingPreference)
func (this *QRawFont) LoadFromFile(fileName unsafe.Pointer, pixelSize float64, hintingPreference int) {
	// 0: (, const QString & fileName, qreal pixelSize, QFont::HintingPreference hintingPreference), (fileName, &pixelSize, &hintingPreference)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12loadFromFileERK7QStringdN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, this.cthis, fileName, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:136
// index:0
// void loadFromData(const class QByteArray &, qreal, class QFont::HintingPreference)
func (this *QRawFont) LoadFromData(fontData unsafe.Pointer, pixelSize float64, hintingPreference int) {
	// 0: (, const QByteArray & fontData, qreal pixelSize, QFont::HintingPreference hintingPreference), (fontData, &pixelSize, &hintingPreference)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12loadFromDataERK10QByteArraydN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, this.cthis, fontData, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:140
// index:0
// bool supportsCharacter(uint)
func (this *QRawFont) SupportsCharacter(ucs4 uint) {
	// 0: (, uint ucs4), (&ucs4)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterEj", ffiqt.FFI_TYPE_VOID, this.cthis, &ucs4)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:141
// index:1
// bool supportsCharacter(class QChar)
func (this *QRawFont) SupportsCharacter_1(character unsafe.Pointer) {
	// 1: (, QChar character), (character)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterE5QChar", ffiqt.FFI_TYPE_VOID, this.cthis, character)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:142
// index:0
// QList<QFontDatabase::WritingSystem> supportedWritingSystems()
func (this *QRawFont) SupportedWritingSystems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont23supportedWritingSystemsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:144
// index:0
// QByteArray fontTable(const char *)
func (this *QRawFont) FontTable(tagName unsafe.Pointer) {
	// 0: (, const char * tagName), (tagName)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9fontTableEPKc", ffiqt.FFI_TYPE_VOID, this.cthis, tagName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:146
// index:0
// static
// QRawFont fromFont(const class QFont &, class QFontDatabase::WritingSystem)
func (this *QRawFont) FromFont(font unsafe.Pointer, writingSystem int) {
	// 0: (const QFont & font, QFontDatabase::WritingSystem writingSystem), (font, writingSystem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont8fromFontERK5QFontN13QFontDatabase13WritingSystemE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QRawFont_FromFont(font unsafe.Pointer, writingSystem int) {
	// 0: (const QFont & font, QFontDatabase::WritingSystem writingSystem), (font, writingSystem)
	var nilthis *QRawFont
	nilthis.FromFont(font, writingSystem)
}

//  body block end
