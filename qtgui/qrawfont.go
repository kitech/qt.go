package qtgui

// /usr/include/qt/QtGui/qrawfont.h
// #include <qrawfont.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QRawFont(const QString &, qreal, QFont::HintingPreference)
func NewQRawFont_1(fileName *qtcore.QString, pixelSize float64, hintingPreference int) *QRawFont {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2ERK7QStringdN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_POINTER, convArg0, pixelSize, hintingPreference)
	gopp.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:78
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QRawFont(const QByteArray &, qreal, QFont::HintingPreference)
func NewQRawFont_2(fontData *qtcore.QByteArray, pixelSize float64, hintingPreference int) *QRawFont {
	var convArg0 = fontData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontC2ERK10QByteArraydN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_POINTER, convArg0, pixelSize, hintingPreference)
	gopp.ErrPrint(err, rv)
	gothis := NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRawFont)
	return gothis
}

// /usr/include/qt/QtGui/qrawfont.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QRawFont()
func DeleteQRawFont(this *QRawFont) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFontD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qrawfont.h:88
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QRawFont &)
func (this *QRawFont) Swap(other *QRawFont) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QRawFont) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QString familyName()
func (this *QRawFont) FamilyName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont10familyNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleName()
func (this *QRawFont) StyleName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9styleNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::Style style()
func (this *QRawFont) Style() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qrawfont.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight()
func (this *QRawFont) Weight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont6weightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool glyphIndexesForChars(const QChar *, int, quint32 *, int *)
func (this *QRawFont) GlyphIndexesForChars(chars *qtcore.QChar /*777 const QChar **/, numChars int, glyphIndexes unsafe.Pointer /*666*/, numGlyphs unsafe.Pointer /*666*/) bool {
	var convArg0 = chars.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont20glyphIndexesForCharsEPK5QChariPjPi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, numChars, &glyphIndexes, &numGlyphs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:106
// index:0
// Public Visibility=Default Availability=Available
// [1] bool advancesForGlyphIndexes(const quint32 *, QPointF *, int)
func (this *QRawFont) AdvancesForGlyphIndexes(glyphIndexes unsafe.Pointer /*666*/, advances *qtcore.QPointF /*777 QPointF **/, numGlyphs int) bool {
	var convArg1 = advances.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &glyphIndexes, convArg1, numGlyphs)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:107
// index:1
// Public Visibility=Default Availability=Available
// [1] bool advancesForGlyphIndexes(const quint32 *, QPointF *, int, QRawFont::LayoutFlags)
func (this *QRawFont) AdvancesForGlyphIndexes_1(glyphIndexes unsafe.Pointer /*666*/, advances *qtcore.QPointF /*777 QPointF **/, numGlyphs int, layoutFlags int) bool {
	var convArg1 = advances.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont23advancesForGlyphIndexesEPKjP7QPointFi6QFlagsINS_10LayoutFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &glyphIndexes, convArg1, numGlyphs, layoutFlags)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:109
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage alphaMapForGlyph(quint32, enum QRawFont::AntialiasingType, const QTransform &)
func (this *QRawFont) AlphaMapForGlyph(glyphIndex uint, antialiasingType int, transform *QTransform) *QImage /*123*/ {
	var convArg2 = transform.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont16alphaMapForGlyphEjNS_16AntialiasingTypeERK10QTransform", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex, antialiasingType, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQImage)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath pathForGlyph(quint32)
func (this *QRawFont) PathForGlyph(glyphIndex uint) *QPainterPath /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12pathForGlyphEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:113
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF boundingRect(quint32)
func (this *QRawFont) BoundingRect(glyphIndex uint) *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12boundingRectEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), glyphIndex)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPixelSize(qreal)
func (this *QRawFont) SetPixelSize(pixelSize float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12setPixelSizeEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), pixelSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal pixelSize()
func (this *QRawFont) PixelSize() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9pixelSizeEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] QFont::HintingPreference hintingPreference()
func (this *QRawFont) HintingPreference() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17hintingPreferenceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qrawfont.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal ascent()
func (this *QRawFont) Ascent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont6ascentEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:121
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal capHeight()
func (this *QRawFont) CapHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9capHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal descent()
func (this *QRawFont) Descent() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7descentEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal leading()
func (this *QRawFont) Leading() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7leadingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal xHeight()
func (this *QRawFont) XHeight() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont7xHeightEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal averageCharWidth()
func (this *QRawFont) AverageCharWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont16averageCharWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal maxCharWidth()
func (this *QRawFont) MaxCharWidth() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont12maxCharWidthEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lineThickness()
func (this *QRawFont) LineThickness() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont13lineThicknessEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal underlinePosition()
func (this *QRawFont) UnderlinePosition() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17underlinePositionEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal unitsPerEm()
func (this *QRawFont) UnitsPerEm() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont10unitsPerEmEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qrawfont.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFromFile(const QString &, qreal, QFont::HintingPreference)
func (this *QRawFont) LoadFromFile(fileName *qtcore.QString, pixelSize float64, hintingPreference int) {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12loadFromFileERK7QStringdN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pixelSize, hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void loadFromData(const QByteArray &, qreal, QFont::HintingPreference)
func (this *QRawFont) LoadFromData(fontData *qtcore.QByteArray, pixelSize float64, hintingPreference int) {
	var convArg0 = fontData.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont12loadFromDataERK10QByteArraydN5QFont17HintingPreferenceE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, pixelSize, hintingPreference)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrawfont.h:140
// index:0
// Public Visibility=Default Availability=Available
// [1] bool supportsCharacter(uint)
func (this *QRawFont) SupportsCharacter(ucs4 uint) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), ucs4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:141
// index:1
// Public Visibility=Default Availability=Available
// [1] bool supportsCharacter(QChar)
func (this *QRawFont) SupportsCharacter_1(character *qtcore.QChar /*123*/) bool {
	var convArg0 = character.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont17supportsCharacterE5QChar", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qrawfont.h:144
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray fontTable(const char *)
func (this *QRawFont) FontTable(tagName string) *qtcore.QByteArray /*123*/ {
	var convArg0 = qtrt.CString(tagName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QRawFont9fontTableEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qrawfont.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [8] QRawFont fromFont(const QFont &, QFontDatabase::WritingSystem)
func (this *QRawFont) FromFont(font *QFont, writingSystem int) *QRawFont /*123*/ {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QRawFont8fromFontERK5QFontN13QFontDatabase13WritingSystemE", ffiqt.FFI_TYPE_POINTER, convArg0, writingSystem)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQRawFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRawFont)
	return rv2
}
func QRawFont_FromFont(font *QFont, writingSystem int) *QRawFont /*123*/ {
	var nilthis *QRawFont
	rv := nilthis.FromFont(font, writingSystem)
	return rv
}

type QRawFont__AntialiasingType = int

const QRawFont__PixelAntialiasing QRawFont__AntialiasingType = 0
const QRawFont__SubPixelAntialiasing QRawFont__AntialiasingType = 1

type QRawFont__LayoutFlag = int

const QRawFont__SeparateAdvances QRawFont__LayoutFlag = 0
const QRawFont__KernedAdvances QRawFont__LayoutFlag = 1
const QRawFont__UseDesignMetrics QRawFont__LayoutFlag = 2

//  body block end
