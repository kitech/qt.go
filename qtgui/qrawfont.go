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
func (this *QRawFont) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QRawFont &) const
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
func (this *QRawFont) Style() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qrawfont.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight() const
func (this *QRawFont) Weight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont6weightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool glyphIndexesForChars(const QChar *, int, quint32 *, int *) const
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
// [32] QImage alphaMapForGlyph(quint32, enum QRawFont::AntialiasingType, const QTransform &) const
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
// [32] QImage alphaMapForGlyph(quint32, enum QRawFont::AntialiasingType, const QTransform &) const
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
// [32] QImage alphaMapForGlyph(quint32, enum QRawFont::AntialiasingType, const QTransform &) const
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
func (this *QRawFont) SetPixelSize(pixelSize float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QRawFont12setPixelSizeEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pixelSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal pixelSize() const
func (this *QRawFont) PixelSize() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont9pixelSizeEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::HintingPreference hintingPreference() const
func (this *QRawFont) HintingPreference() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont17hintingPreferenceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qrawfont.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent() const
func (this *QRawFont) Ascent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont6ascentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal capHeight() const
func (this *QRawFont) CapHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont9capHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent() const
func (this *QRawFont) Descent() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7descentEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leading() const
func (this *QRawFont) Leading() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7leadingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal xHeight() const
func (this *QRawFont) XHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont7xHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal averageCharWidth() const
func (this *QRawFont) AverageCharWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont16averageCharWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maxCharWidth() const
func (this *QRawFont) MaxCharWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont12maxCharWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lineThickness() const
func (this *QRawFont) LineThickness() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont13lineThicknessEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal underlinePosition() const
func (this *QRawFont) UnderlinePosition() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont17underlinePositionEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal unitsPerEm() const
func (this *QRawFont) UnitsPerEm() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont10unitsPerEmEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFromFile(const QString &, qreal, QFont::HintingPreference)
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
func (this *QRawFont) SupportsCharacter(ucs4 uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterEj", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:141
// index:1
// Public Visibility=Default Availability=Available
// [1] bool supportsCharacter(QChar) const
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

type QRawFont__AntialiasingType = int

const QRawFont__PixelAntialiasing QRawFont__AntialiasingType = 0
const QRawFont__SubPixelAntialiasing QRawFont__AntialiasingType = 1

type QRawFont__LayoutFlag = int

const QRawFont__SeparateAdvances QRawFont__LayoutFlag = 0
const QRawFont__KernedAdvances QRawFont__LayoutFlag = 1
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
