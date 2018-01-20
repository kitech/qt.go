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
	*qtrt.CObject
}

func (this *QRawFont) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQRawFontFromPointer(cthis unsafe.Pointer) *QRawFont {
	return &QRawFont{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qrawfont.h:74
// index:0
// Public
// void QRawFont()
func NewQRawFont() *QRawFont {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:75
// index:1
// Public
// void QRawFont(const class QString &, qreal, class QFont::HintingPreference)
func NewQRawFont_1(fileName *qtcore.QString, pixelSize float64, hintingPreference int) *QRawFont {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2ERK7QStringdN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:78
// index:2
// Public
// void QRawFont(const class QByteArray &, qreal, class QFont::HintingPreference)
func NewQRawFont_2(fontData *qtcore.QByteArray, pixelSize float64, hintingPreference int) *QRawFont {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fontData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2ERK10QByteArraydN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:86
// index:0
// Public
// void ~QRawFont()
func DeleteQRawFont(*QRawFont) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:88
// index:0
// Public inline
// void swap(class QRawFont &)
func (this *QRawFont) Swap(other *QRawFont) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:90
// index:0
// Public
// bool isValid()
func (this *QRawFont) IsValid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:96
// index:0
// Public
// QString familyName()
func (this *QRawFont) FamilyName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont10familyNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:97
// index:0
// Public
// QString styleName()
func (this *QRawFont) StyleName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9styleNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:99
// index:0
// Public
// QFont::Style style()
func (this *QRawFont) Style() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:100
// index:0
// Public
// int weight()
func (this *QRawFont) Weight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont6weightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:102
// index:0
// Public
// QVector<quint32> glyphIndexesForString(const class QString &)
func (this *QRawFont) GlyphIndexesForString(text *qtcore.QString) interface{} {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont21glyphIndexesForStringERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:105
// index:0
// Public
// bool glyphIndexesForChars(const class QChar *, int, quint32 *, int *)
func (this *QRawFont) GlyphIndexesForChars(chars unsafe.Pointer, numChars int, glyphIndexes unsafe.Pointer, numGlyphs unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), chars, &numChars, glyphIndexes, numGlyphs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:106
// index:0
// Public
// bool advancesForGlyphIndexes(const quint32 *, class QPointF *, int)
func (this *QRawFont) AdvancesForGlyphIndexes(glyphIndexes unsafe.Pointer, advances unsafe.Pointer, numGlyphs int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndexes, advances, &numGlyphs)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:109
// index:0
// Public
// QImage alphaMapForGlyph(quint32, enum QRawFont::AntialiasingType, const class QTransform &)
func (this *QRawFont) AlphaMapForGlyph(glyphIndex uint, antialiasingType int, transform *QTransform) interface{} {
	var convArg2 = transform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont16alphaMapForGlyphEjNS_16AntialiasingTypeERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &glyphIndex, &antialiasingType, convArg2)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:112
// index:0
// Public
// QPainterPath pathForGlyph(quint32)
func (this *QRawFont) PathForGlyph(glyphIndex uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12pathForGlyphEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &glyphIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:113
// index:0
// Public
// QRectF boundingRect(quint32)
func (this *QRawFont) BoundingRect(glyphIndex uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12boundingRectEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &glyphIndex)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:115
// index:0
// Public
// void setPixelSize(qreal)
func (this *QRawFont) SetPixelSize(pixelSize float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12setPixelSizeEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &pixelSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:116
// index:0
// Public
// qreal pixelSize()
func (this *QRawFont) PixelSize() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9pixelSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:118
// index:0
// Public
// QFont::HintingPreference hintingPreference()
func (this *QRawFont) HintingPreference() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17hintingPreferenceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:120
// index:0
// Public
// qreal ascent()
func (this *QRawFont) Ascent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont6ascentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:121
// index:0
// Public
// qreal capHeight()
func (this *QRawFont) CapHeight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9capHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:122
// index:0
// Public
// qreal descent()
func (this *QRawFont) Descent() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7descentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:123
// index:0
// Public
// qreal leading()
func (this *QRawFont) Leading() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7leadingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:124
// index:0
// Public
// qreal xHeight()
func (this *QRawFont) XHeight() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7xHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:125
// index:0
// Public
// qreal averageCharWidth()
func (this *QRawFont) AverageCharWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont16averageCharWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:126
// index:0
// Public
// qreal maxCharWidth()
func (this *QRawFont) MaxCharWidth() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12maxCharWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:127
// index:0
// Public
// qreal lineThickness()
func (this *QRawFont) LineThickness() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont13lineThicknessEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:128
// index:0
// Public
// qreal underlinePosition()
func (this *QRawFont) UnderlinePosition() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17underlinePositionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:130
// index:0
// Public
// qreal unitsPerEm()
func (this *QRawFont) UnitsPerEm() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont10unitsPerEmEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:132
// index:0
// Public
// void loadFromFile(const class QString &, qreal, class QFont::HintingPreference)
func (this *QRawFont) LoadFromFile(fileName *qtcore.QString, pixelSize float64, hintingPreference int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12loadFromFileERK7QStringdN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:136
// index:0
// Public
// void loadFromData(const class QByteArray &, qreal, class QFont::HintingPreference)
func (this *QRawFont) LoadFromData(fontData *qtcore.QByteArray, pixelSize float64, hintingPreference int) {
	var convArg0 = fontData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12loadFromDataERK10QByteArraydN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &pixelSize, &hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:140
// index:0
// Public
// bool supportsCharacter(uint)
func (this *QRawFont) SupportsCharacter(ucs4 uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &ucs4)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:141
// index:1
// Public
// bool supportsCharacter(class QChar)
func (this *QRawFont) SupportsCharacter_1(character *qtcore.QChar) interface{} {
	var convArg0 = character.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:142
// index:0
// Public
// QList<QFontDatabase::WritingSystem> supportedWritingSystems()
func (this *QRawFont) SupportedWritingSystems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont23supportedWritingSystemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:144
// index:0
// Public
// QByteArray fontTable(const char *)
func (this *QRawFont) FontTable(tagName string) interface{} {
	var convArg0 = qtrt.CString(tagName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9fontTableEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:146
// index:0
// Public static
// QRawFont fromFont(const class QFont &, class QFontDatabase::WritingSystem)
func (this *QRawFont) FromFont(font *QFont, writingSystem int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont8fromFontERK5QFontN13QFontDatabase13WritingSystemE", ffiqt.FFI_TYPE_POINTER, font, writingSystem)
	gopp.ErrPrint(err, rv)
	return rv
}
func QRawFont_FromFont(font *QFont, writingSystem int) {
	var nilthis *QRawFont
	nilthis.FromFont(font, writingSystem)
}

//  body block end
