package qtgui

// /usr/include/qt/QtGui/qrawfont.h
// #include <qrawfont.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QRawFont struct {
	*qtrt.CObject
}
type QRawFont_ITF interface {
	QRawFont_PTR() *QRawFont
}

func (ptr *QRawFont) QRawFont_PTR() *QRawFont { return ptr }

func (this *QRawFont) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRawFont) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRawFontFromPointer(cthis unsafe.Pointer) *QRawFont {
	return &QRawFont{&qtrt.CObject{cthis}}
}
func (*QRawFont) NewFromPointer(cthis unsafe.Pointer) *QRawFont {
	return NewQRawFontFromPointer(cthis)
}

// /usr/include/qt/QtGui/qrawfont.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRawFont()

/*
Constructs an invalid QRawFont.
*/
func NewQRawFont() *QRawFont {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRawFont(const QString &, qreal, QFont::HintingPreference)

/*
Constructs an invalid QRawFont.
*/
func NewQRawFont_1(fileName string, pixelSize float64, hintingPreference int) *QRawFont {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontC2ERK7QStringdN5QFont17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, convArg0, pixelSize, hintingPreference)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRawFont(const QString &, qreal, QFont::HintingPreference)

/*
Constructs an invalid QRawFont.
*/
func NewQRawFont_1_(fileName string, pixelSize float64) *QRawFont {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QFont::HintingPreference=Elaborated, QFont::HintingPreference=Enum,
	hintingPreference := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontC2ERK7QStringdN5QFont17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, convArg0, pixelSize, hintingPreference)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:78
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRawFont(const QByteArray &, qreal, QFont::HintingPreference)

/*
Constructs an invalid QRawFont.
*/
func NewQRawFont_2(fontData qtcore.QByteArray_ITF, pixelSize float64, hintingPreference int) *QRawFont {
	var convArg0 unsafe.Pointer
	if fontData != nil && fontData.QByteArray_PTR() != nil {
		convArg0 = fontData.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontC2ERK10QByteArraydN5QFont17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, convArg0, pixelSize, hintingPreference)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:78
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRawFont(const QByteArray &, qreal, QFont::HintingPreference)

/*
Constructs an invalid QRawFont.
*/
func NewQRawFont_2_(fontData qtcore.QByteArray_ITF, pixelSize float64) *QRawFont {
	var convArg0 unsafe.Pointer
	if fontData != nil && fontData.QByteArray_PTR() != nil {
		convArg0 = fontData.QByteArray_PTR().GetCthis()
	}
	// arg: 2, QFont::HintingPreference=Elaborated, QFont::HintingPreference=Enum,
	hintingPreference := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontC2ERK10QByteArraydN5QFont17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, convArg0, pixelSize, hintingPreference)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:83
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QRawFont & operator=(QRawFont &&)

/*

 */
func (this *QRawFont) Operator_equal(other unsafe.Pointer /*333*/) *QRawFont {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRawFont)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:85
// index:1
// Public Visibility=Default Availability=Available
// [8] QRawFont & operator=(const QRawFont &)

/*

 */
func (this *QRawFont) Operator_equal_1(other QRawFont_ITF) *QRawFont {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRawFont_PTR() != nil {
		convArg0 = other.QRawFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRawFont)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRawFont()

/*

 */
func DeleteQRawFont(this *QRawFont) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFontD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qrawfont.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRawFont &)

/*
Swaps this raw font with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func (this *QRawFont) Swap(other QRawFont_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRawFont_PTR() != nil {
		convArg0 = other.QRawFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFont4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the QRawFont is valid and false otherwise.
*/
func (this *QRawFont) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QRawFont &) const

/*

 */
func (this *QRawFont) Operator_equal_equal(other QRawFont_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRawFont_PTR() != nil {
		convArg0 = other.QRawFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFonteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QRawFont &) const

/*

 */
func (this *QRawFont) Operator_not_equal(other QRawFont_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QRawFont_PTR() != nil {
		convArg0 = other.QRawFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFontneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QString familyName() const

/*
Returns the family name of this QRawFont.
*/
func (this *QRawFont) FamilyName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont10familyNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qrawfont.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleName() const

/*
Returns the style name of this QRawFont.

See also QFont::styleName().
*/
func (this *QRawFont) StyleName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont9styleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qrawfont.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Style style() const

/*
Returns the style of this QRawFont.

See also QFont::style().
*/
func (this *QRawFont) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qrawfont.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight() const

/*
Returns the weight of this QRawFont.

See also QFont::weight().
*/
func (this *QRawFont) Weight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont6weightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool glyphIndexesForChars(const QChar *, int, quint32 *, int *) const

/*
Converts a string of unicode points to glyph indexes using the CMAP table in the underlying font. The function works like glyphIndexesForString() except it take an array (chars), the results will be returned though glyphIndexes array and number of glyphs will be set in numGlyphs. The size of glyphIndexes array must be at least numChars, if that's still not enough, this function will return false, then you can resize glyphIndexes from the size returned in numGlyphs.

See also glyphIndexesForString(), advancesForGlyphIndexes(), QGlyphRun, QTextLayout::glyphRuns(), and QTextFragment::glyphRuns().
*/
func (this *QRawFont) GlyphIndexesForChars(chars qtcore.QChar_ITF /*777 const QChar **/, numChars int, glyphIndexes unsafe.Pointer /*666*/, numGlyphs unsafe.Pointer /*666*/) bool {
	var convArg0 unsafe.Pointer
	if chars != nil && chars.QChar_PTR() != nil {
		convArg0 = chars.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numChars, glyphIndexes, numGlyphs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool advancesForGlyphIndexes(const quint32 *, QPointF *, int) const

/*
Returns the QRawFont's advances for each of the glyphIndexes in pixel units. The advances give the distance from the position of a given glyph to where the next glyph should be drawn to make it appear as if the two glyphs are unspaced. How the advances are calculated is controlled by layoutFlags.

This function was introduced in  Qt 5.1.

See also QTextLine::horizontalAdvance() and QFontMetricsF::width().
*/
func (this *QRawFont) AdvancesForGlyphIndexes(glyphIndexes unsafe.Pointer /*666*/, advances qtcore.QPointF_ITF /*777 QPointF **/, numGlyphs int) bool {
	var convArg1 unsafe.Pointer
	if advances != nil && advances.QPointF_PTR() != nil {
		convArg1 = advances.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndexes, convArg1, numGlyphs)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:107
// index:1
// Public Visibility=Default Availability=Available
// [1] bool advancesForGlyphIndexes(const quint32 *, QPointF *, int, QRawFont::LayoutFlags) const

/*
Returns the QRawFont's advances for each of the glyphIndexes in pixel units. The advances give the distance from the position of a given glyph to where the next glyph should be drawn to make it appear as if the two glyphs are unspaced. How the advances are calculated is controlled by layoutFlags.

This function was introduced in  Qt 5.1.

See also QTextLine::horizontalAdvance() and QFontMetricsF::width().
*/
func (this *QRawFont) AdvancesForGlyphIndexes_1(glyphIndexes unsafe.Pointer /*666*/, advances qtcore.QPointF_ITF /*777 QPointF **/, numGlyphs int, layoutFlags int) bool {
	var convArg1 unsafe.Pointer
	if advances != nil && advances.QPointF_PTR() != nil {
		convArg1 = advances.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi6QFlagsINS_10LayoutFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndexes, convArg1, numGlyphs, layoutFlags)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:109
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage alphaMapForGlyph(quint32, QRawFont::AntialiasingType, const QTransform &) const

/*
This function returns a rasterized image of the glyph at the given glyphIndex in the underlying font, using the transform specified. If the QRawFont is not valid, this function will return an invalid QImage.

If the font is a color font, then the resulting image will contain the rendered glyph at the current pixel size. In this case, the antialiasingType will be ignored.

Otherwise, if antialiasingType is set to QRawFont::SubPixelAntialiasing, then the resulting image will be in QImage::Format_RGB32 and the RGB values of each pixel will represent the subpixel opacities of the pixel in the rasterization of the glyph. Otherwise, the image will be in the format of QImage::Format_Indexed8 and each pixel will contain the opacity of the pixel in the rasterization.

See also pathForGlyph() and QPainter::drawGlyphRun().
*/
func (this *QRawFont) AlphaMapForGlyph(glyphIndex uint, antialiasingType int, transform QTransform_ITF) *QImage /*123*/ {
	var convArg2 unsafe.Pointer
	if transform != nil && transform.QTransform_PTR() != nil {
		convArg2 = transform.QTransform_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont16alphaMapForGlyphEjNS_16AntialiasingTypeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex, antialiasingType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:109
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage alphaMapForGlyph(quint32, QRawFont::AntialiasingType, const QTransform &) const

/*
This function returns a rasterized image of the glyph at the given glyphIndex in the underlying font, using the transform specified. If the QRawFont is not valid, this function will return an invalid QImage.

If the font is a color font, then the resulting image will contain the rendered glyph at the current pixel size. In this case, the antialiasingType will be ignored.

Otherwise, if antialiasingType is set to QRawFont::SubPixelAntialiasing, then the resulting image will be in QImage::Format_RGB32 and the RGB values of each pixel will represent the subpixel opacities of the pixel in the rasterization of the glyph. Otherwise, the image will be in the format of QImage::Format_Indexed8 and each pixel will contain the opacity of the pixel in the rasterization.

See also pathForGlyph() and QPainter::drawGlyphRun().
*/
func (this *QRawFont) AlphaMapForGlyph__(glyphIndex uint) *QImage /*123*/ {
	// arg: 1, QRawFont::AntialiasingType=Enum, QRawFont::AntialiasingType=Enum,
	antialiasingType := 0
	// arg: 2, const QTransform &=LValueReference, QTransform=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont16alphaMapForGlyphEjNS_16AntialiasingTypeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex, antialiasingType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:109
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage alphaMapForGlyph(quint32, QRawFont::AntialiasingType, const QTransform &) const

/*
This function returns a rasterized image of the glyph at the given glyphIndex in the underlying font, using the transform specified. If the QRawFont is not valid, this function will return an invalid QImage.

If the font is a color font, then the resulting image will contain the rendered glyph at the current pixel size. In this case, the antialiasingType will be ignored.

Otherwise, if antialiasingType is set to QRawFont::SubPixelAntialiasing, then the resulting image will be in QImage::Format_RGB32 and the RGB values of each pixel will represent the subpixel opacities of the pixel in the rasterization of the glyph. Otherwise, the image will be in the format of QImage::Format_Indexed8 and each pixel will contain the opacity of the pixel in the rasterization.

See also pathForGlyph() and QPainter::drawGlyphRun().
*/
func (this *QRawFont) AlphaMapForGlyph__1(glyphIndex uint, antialiasingType int) *QImage /*123*/ {
	// arg: 2, const QTransform &=LValueReference, QTransform=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont16alphaMapForGlyphEjNS_16AntialiasingTypeERK10QTransform", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex, antialiasingType, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath pathForGlyph(quint32) const

/*
This function returns the shape of the glyph at a given glyphIndex in the underlying font if the QRawFont is valid. Otherwise, it returns an empty QPainterPath.

The returned glyph will always be unhinted.

See also alphaMapForGlyph() and QPainterPath::addText().
*/
func (this *QRawFont) PathForGlyph(glyphIndex uint) *QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont12pathForGlyphEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:113
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(quint32) const

/*
Returns the smallest rectangle containing the glyph with the given glyphIndex.

This function was introduced in  Qt 5.0.
*/
func (this *QRawFont) BoundingRect(glyphIndex uint) *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont12boundingRectEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelSize(qreal)

/*
Sets the pixel size with which this font should be rendered to pixelSize.

See also pixelSize().
*/
func (this *QRawFont) SetPixelSize(pixelSize float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFont12setPixelSizeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pixelSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal pixelSize() const

/*
Returns the pixel size set for this QRawFont. The pixel size affects how glyphs are rasterized, the size of glyphs returned by pathForGlyph(), and is used to convert internal metrics from design units to logical pixel units.

See also setPixelSize().
*/
func (this *QRawFont) PixelSize() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont9pixelSizeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::HintingPreference hintingPreference() const

/*
Returns the hinting preference used to construct this QRawFont.

See also QFont::hintingPreference().
*/
func (this *QRawFont) HintingPreference() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont17hintingPreferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qrawfont.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent() const

/*
Returns the ascent of this QRawFont in pixel units.

The ascent of a font is the distance from the baseline to the highest position characters extend to. In practice, some font designers break this rule, e.g. when they put more than one accent on top of a character, or to accommodate an unusual character in an exotic language, so it is possible (though rare) that this value will be too small.

See also QFontMetricsF::ascent().
*/
func (this *QRawFont) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal capHeight() const

/*
Returns the cap height of this QRawFont in pixel units.

The cap height of a font is the height of a capital letter above the baseline. It specifically is the height of capital letters that are flat - such as H or I - as opposed to round letters such as O, or pointed letters like A, both of which may display overshoot.

This function was introduced in  Qt 5.8.

See also QFontMetricsF::capHeight().
*/
func (this *QRawFont) CapHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont9capHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent() const

/*
Returns the descent of this QRawFont in pixel units.

The descent is the distance from the base line to the lowest point characters extend to. In practice, some font designers break this rule, e.g. to accommodate an unusual character in an exotic language, so it is possible (though rare) that this value will be too small.

See also QFontMetricsF::descent().
*/
func (this *QRawFont) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leading() const

/*
Returns the leading of this QRawFont in pixel units.

This is the natural inter-line spacing.

See also QFontMetricsF::leading().
*/
func (this *QRawFont) Leading() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7leadingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal xHeight() const

/*
Returns the xHeight of this QRawFont in pixel units.

This is often but not always the same as the height of the character 'x'.

See also QFontMetricsF::xHeight().
*/
func (this *QRawFont) XHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7xHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal averageCharWidth() const

/*
Returns the average character width of this QRawFont in pixel units.

See also QFontMetricsF::averageCharWidth().
*/
func (this *QRawFont) AverageCharWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont16averageCharWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maxCharWidth() const

/*
Returns the width of the widest character in the font.

See also QFontMetricsF::maxWidth().
*/
func (this *QRawFont) MaxCharWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont12maxCharWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lineThickness() const

/*
Returns the thickness for drawing lines (underline, overline, etc.) along with text drawn in this font.
*/
func (this *QRawFont) LineThickness() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont13lineThicknessEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal underlinePosition() const

/*
Returns the position from baseline for drawing underlines below the text rendered with this font.
*/
func (this *QRawFont) UnderlinePosition() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont17underlinePositionEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal unitsPerEm() const

/*
Returns the number of design units define the width and height of the em square for this QRawFont. This value is used together with the pixel size when converting design metrics to pixel units, as the internal metrics are specified in design units and the pixel size gives the size of 1 em in pixels.

See also pixelSize() and setPixelSize().
*/
func (this *QRawFont) UnitsPerEm() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont10unitsPerEmEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFromFile(const QString &, qreal, QFont::HintingPreference)

/*
Replaces the current QRawFont with the contents of the file referenced by fileName for the size (in pixels) given by pixelSize, and using the hinting preference specified by hintingPreference.

The file must reference a TrueType or OpenType font.

See also loadFromData().
*/
func (this *QRawFont) LoadFromFile(fileName string, pixelSize float64, hintingPreference int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFont12loadFromFileERK7QStringdN5QFont17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pixelSize, hintingPreference)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFromData(const QByteArray &, qreal, QFont::HintingPreference)

/*
Replaces the current QRawFont with the font contained in the supplied fontData for the size (in pixels) given by pixelSize, and using the hinting preference specified by hintingPreference.

The fontData must contain a TrueType or OpenType font.

See also loadFromFile().
*/
func (this *QRawFont) LoadFromData(fontData qtcore.QByteArray_ITF, pixelSize float64, hintingPreference int) {
	var convArg0 unsafe.Pointer
	if fontData != nil && fontData.QByteArray_PTR() != nil {
		convArg0 = fontData.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFont12loadFromDataERK10QByteArraydN5QFont17HintingPreferenceE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pixelSize, hintingPreference)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsCharacter(uint) const

/*
Returns true if the font has a glyph that corresponds to the given character.

See also supportedWritingSystems().
*/
func (this *QRawFont) SupportsCharacter(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:141
// index:1
// Public Visibility=Default Availability=Available
// [1] bool supportsCharacter(QChar) const

/*
Returns true if the font has a glyph that corresponds to the given character.

See also supportedWritingSystems().
*/
func (this *QRawFont) SupportsCharacter_1(character qtcore.QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if character != nil && character.QChar_PTR() != nil {
		convArg0 = character.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray fontTable(const char *) const

/*
Retrieves the sfnt table named tagName from the underlying physical font, or an empty byte array if no such table was found. The returned font table's byte order is Big Endian, like the sfnt format specifies. The tagName must be four characters long and should be formatted in the default endianness of the current platform.
*/
func (this *QRawFont) FontTable(tagName string) *qtcore.QByteArray /*123*/ {
	var convArg0 = qtrt.CString(tagName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont9fontTableEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [8] QRawFont fromFont(const QFont &, QFontDatabase::WritingSystem)

/*
Fetches the physical representation based on a font query. The physical font returned is the font that will be preferred by Qt in order to display text in the selected writingSystem.

Warning: This function is potentially expensive and should not be called in performance sensitive code.
*/
func (this *QRawFont) FromFont(font QFont_ITF, writingSystem int) *QRawFont /*123*/ {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFont8fromFontERK5QFontN13QFontDatabase13WritingSystemE", qtrt.FFI_TYPE_POINTER, convArg0, writingSystem)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRawFont)
	return rv2
}
func QRawFont_FromFont(font QFont_ITF, writingSystem int) *QRawFont /*123*/ {
	var nilthis *QRawFont
	rv := nilthis.FromFont(font, writingSystem)
	return rv
}

// /usr/include/qt/QtGui/qrawfont.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [8] QRawFont fromFont(const QFont &, QFontDatabase::WritingSystem)

/*
Fetches the physical representation based on a font query. The physical font returned is the font that will be preferred by Qt in order to display text in the selected writingSystem.

Warning: This function is potentially expensive and should not be called in performance sensitive code.
*/
func (this *QRawFont) FromFont__(font QFont_ITF) *QRawFont /*123*/ {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	// arg: 1, QFontDatabase::WritingSystem=Elaborated, QFontDatabase::WritingSystem=Enum,
	writingSystem := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFont8fromFontERK5QFontN13QFontDatabase13WritingSystemE", qtrt.FFI_TYPE_POINTER, convArg0, writingSystem)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRawFont)
	return rv2
}

/*
This enum represents the different ways a glyph can be rasterized in the function alphaMapForGlyph().


*/
type QRawFont__AntialiasingType = int

// Will rasterize by measuring the coverage of the shape on whole pixels. The returned image contains the alpha values of each pixel based on the coverage of the glyph shape.
const QRawFont__PixelAntialiasing QRawFont__AntialiasingType = 0

// Will rasterize by measuring the coverage of each subpixel, returning a separate alpha value for each of the red, green and blue components of each pixel.
const QRawFont__SubPixelAntialiasing QRawFont__AntialiasingType = 1

/*


 */
type QRawFont__LayoutFlag = int

//
const QRawFont__SeparateAdvances QRawFont__LayoutFlag = 0

//
const QRawFont__KernedAdvances QRawFont__LayoutFlag = 1

//
const QRawFont__UseDesignMetrics QRawFont__LayoutFlag = 2

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
