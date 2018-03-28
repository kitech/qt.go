package qtgui

// /usr/include/qt/QtGui/qfontdatabase.h
// #include <qfontdatabase.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QFontDatabase struct {
	*qtrt.CObject
}
type QFontDatabase_ITF interface {
	QFontDatabase_PTR() *QFontDatabase
}

func (ptr *QFontDatabase) QFontDatabase_PTR() *QFontDatabase { return ptr }

func (this *QFontDatabase) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QFontDatabase) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQFontDatabaseFromPointer(cthis unsafe.Pointer) *QFontDatabase {
	return &QFontDatabase{&qtrt.CObject{cthis}}
}
func (*QFontDatabase) NewFromPointer(cthis unsafe.Pointer) *QFontDatabase {
	return NewQFontDatabaseFromPointer(cthis)
}

// /usr/include/qt/QtGui/qfontdatabase.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QFontDatabase()

/*
Creates a font database object.
*/
func NewQFontDatabase() *QFontDatabase {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabaseC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQFontDatabaseFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQFontDatabase)
	return gothis
}

// /usr/include/qt/QtGui/qfontdatabase.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList families(QFontDatabase::WritingSystem) const

/*
Returns a sorted list of the available font families which support the writingSystem.

If a family exists in several foundries, the returned name for that font is in the form "family [foundry]". Examples: "Times [Adobe]", "Times [Cronyx]", "Palatino".

See also writingSystems().
*/
func (this *QFontDatabase) Families(writingSystem int) *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase8familiesENS_13WritingSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), writingSystem)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qfontdatabase.h:123
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList families(QFontDatabase::WritingSystem) const

/*
Returns a sorted list of the available font families which support the writingSystem.

If a family exists in several foundries, the returned name for that font is in the form "family [foundry]". Examples: "Times [Adobe]", "Times [Cronyx]", "Palatino".

See also writingSystems().
*/
func (this *QFontDatabase) Families__() *qtcore.QStringList /*123*/ {
	// arg: 0, QFontDatabase::WritingSystem=Enum, QFontDatabase::WritingSystem=Enum,
	writingSystem := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase8familiesENS_13WritingSystemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), writingSystem)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qfontdatabase.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList styles(const QString &) const

/*
Returns a list of the styles available for the font family family. Some example styles: "Light", "Light Italic", "Bold", "Oblique", "Demi". The list may be empty.

See also families().
*/
func (this *QFontDatabase) Styles(family string) *qtcore.QStringList /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase6stylesERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qfontdatabase.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleString(const QFont &)

/*
Returns a string that describes the style of the font. For example, "Bold Italic", "Bold", "Italic" or "Normal". An empty string may be returned.
*/
func (this *QFontDatabase) StyleString(font QFont_ITF) string {
	var convArg0 unsafe.Pointer
	if font != nil && font.QFont_PTR() != nil {
		convArg0 = font.QFont_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontdatabase.h:128
// index:1
// Public Visibility=Default Availability=Available
// [8] QString styleString(const QFontInfo &)

/*
Returns a string that describes the style of the font. For example, "Bold Italic", "Bold", "Italic" or "Normal". An empty string may be returned.
*/
func (this *QFontDatabase) StyleString_1(fontInfo QFontInfo_ITF) string {
	var convArg0 unsafe.Pointer
	if fontInfo != nil && fontInfo.QFontInfo_PTR() != nil {
		convArg0 = fontInfo.QFontInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK9QFontInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qfontdatabase.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font(const QString &, const QString &, int) const

/*
Returns a QFont object that has family family, style style and point size pointSize. If no matching font could be created, a QFont object that uses the application's default font is returned.
*/
func (this *QFontDatabase) Font(family string, style string, pointSize int) *QFont /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase4fontERK7QStringS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, pointSize)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}

// /usr/include/qt/QtGui/qfontdatabase.h:132
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBitmapScalable(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is a scalable bitmap font; otherwise returns false. Scaling a bitmap font usually produces an unattractive hardly readable result, because the pixels of the font are scaled. If you need to scale a bitmap font it is better to scale it to one of the fixed sizes returned by smoothSizes().

See also isScalable() and isSmoothlyScalable().
*/
func (this *QFontDatabase) IsBitmapScalable(family string, style string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:132
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBitmapScalable(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is a scalable bitmap font; otherwise returns false. Scaling a bitmap font usually produces an unattractive hardly readable result, because the pixels of the font are scaled. If you need to scale a bitmap font it is better to scale it to one of the fixed sizes returned by smoothSizes().

See also isScalable() and isSmoothlyScalable().
*/
func (this *QFontDatabase) IsBitmapScalable__(family string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSmoothlyScalable(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is smoothly scalable; otherwise returns false. If this function returns true, it's safe to scale this font to any size, and the result will always look attractive.

See also isScalable() and isBitmapScalable().
*/
func (this *QFontDatabase) IsSmoothlyScalable(family string, style string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:133
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSmoothlyScalable(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is smoothly scalable; otherwise returns false. If this function returns true, it's safe to scale this font to any size, and the result will always look attractive.

See also isScalable() and isBitmapScalable().
*/
func (this *QFontDatabase) IsSmoothlyScalable__(family string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:134
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScalable(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is scalable; otherwise returns false.

See also isBitmapScalable() and isSmoothlyScalable().
*/
func (this *QFontDatabase) IsScalable(family string, style string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase10isScalableERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:134
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isScalable(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is scalable; otherwise returns false.

See also isBitmapScalable() and isSmoothlyScalable().
*/
func (this *QFontDatabase) IsScalable__(family string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase10isScalableERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFixedPitch(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is fixed pitch; otherwise returns false.
*/
func (this *QFontDatabase) IsFixedPitch(family string, style string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:135
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFixedPitch(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is fixed pitch; otherwise returns false.
*/
func (this *QFontDatabase) IsFixedPitch__(family string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record,
	var convArg1 = qtcore.NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:137
// index:0
// Public Visibility=Default Availability=Available
// [1] bool italic(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is italic; otherwise returns false.

See also weight() and bold().
*/
func (this *QFontDatabase) Italic(family string, style string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase6italicERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:138
// index:0
// Public Visibility=Default Availability=Available
// [1] bool bold(const QString &, const QString &) const

/*
Returns true if the font that has family family and style style is bold; otherwise returns false.

See also italic() and weight().
*/
func (this *QFontDatabase) Bold(family string, style string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase4boldERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] int weight(const QString &, const QString &) const

/*
Returns the weight of the font that has family family and style style. If there is no such family and style combination, returns -1.

See also italic() and bold().
*/
func (this *QFontDatabase) Weight(family string, style string) int {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(style)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase6weightERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qfontdatabase.h:141
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFamily(const QString &) const

/*

 */
func (this *QFontDatabase) HasFamily(family string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase9hasFamilyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:142
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPrivateFamily(const QString &) const

/*
Returns true if and only if the family font family is private.

This happens, for instance, on macOS and iOS, where the system UI fonts are not accessible to the user. For completeness, QFontDatabase::families() returns all font families, including the private ones. You should use this function if you are developing a font selection control in order to keep private fonts hidden.

This function was introduced in  Qt 5.5.

See also families().
*/
func (this *QFontDatabase) IsPrivateFamily(family string) bool {
	var tmpArg0 = qtcore.NewQString_5(family)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QFontDatabase15isPrivateFamilyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qfontdatabase.h:144
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString writingSystemName(QFontDatabase::WritingSystem)

/*
Returns the names the writingSystem (e.g. for displaying to the user in a dialog).
*/
func (this *QFontDatabase) WritingSystemName(writingSystem int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase17writingSystemNameENS_13WritingSystemE", qtrt.FFI_TYPE_POINTER, writingSystem)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QFontDatabase_WritingSystemName(writingSystem int) string {
	var nilthis *QFontDatabase
	rv := nilthis.WritingSystemName(writingSystem)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:145
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString writingSystemSample(QFontDatabase::WritingSystem)

/*
Returns a string with sample characters from writingSystem.
*/
func (this *QFontDatabase) WritingSystemSample(writingSystem int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase19writingSystemSampleENS_13WritingSystemE", qtrt.FFI_TYPE_POINTER, writingSystem)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QFontDatabase_WritingSystemSample(writingSystem int) string {
	var nilthis *QFontDatabase
	rv := nilthis.WritingSystemSample(writingSystem)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:147
// index:0
// Public static Visibility=Default Availability=Available
// [4] int addApplicationFont(const QString &)

/*
Loads the font from the file specified by fileName and makes it available to the application. An ID is returned that can be used to remove the font again with removeApplicationFont() or to retrieve the list of family names contained in the font.

The function returns -1 if the font could not be loaded.

Currently only TrueType fonts, TrueType font collections, and OpenType fonts are supported.

Note: Adding application fonts on Unix/X11 platforms without fontconfig is currently not supported.

This function was introduced in  Qt 4.2.

See also addApplicationFontFromData(), applicationFontFamilies(), and removeApplicationFont().
*/
func (this *QFontDatabase) AddApplicationFont(fileName string) int {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase18addApplicationFontERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QFontDatabase_AddApplicationFont(fileName string) int {
	var nilthis *QFontDatabase
	rv := nilthis.AddApplicationFont(fileName)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:148
// index:0
// Public static Visibility=Default Availability=Available
// [4] int addApplicationFontFromData(const QByteArray &)

/*
Loads the font from binary data specified by fontData and makes it available to the application. An ID is returned that can be used to remove the font again with removeApplicationFont() or to retrieve the list of family names contained in the font.

The function returns -1 if the font could not be loaded.

Currently only TrueType fonts and TrueType font collections are supported.

Note: Adding application fonts on Unix/X11 platforms without fontconfig is currently not supported.

This function was introduced in  Qt 4.2.

See also addApplicationFont(), applicationFontFamilies(), and removeApplicationFont().
*/
func (this *QFontDatabase) AddApplicationFontFromData(fontData qtcore.QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if fontData != nil && fontData.QByteArray_PTR() != nil {
		convArg0 = fontData.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QFontDatabase_AddApplicationFontFromData(fontData qtcore.QByteArray_ITF) int {
	var nilthis *QFontDatabase
	rv := nilthis.AddApplicationFontFromData(fontData)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:149
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList applicationFontFamilies(int)

/*
Returns a list of font families for the given application font identified by id.

This function was introduced in  Qt 4.2.

See also addApplicationFont() and addApplicationFontFromData().
*/
func (this *QFontDatabase) ApplicationFontFamilies(id int) *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase23applicationFontFamiliesEi", qtrt.FFI_TYPE_POINTER, id)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QFontDatabase_ApplicationFontFamilies(id int) *qtcore.QStringList /*123*/ {
	var nilthis *QFontDatabase
	rv := nilthis.ApplicationFontFamilies(id)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:150
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool removeApplicationFont(int)

/*
Removes the previously loaded application font identified by id. Returns true if unloading of the font succeeded; otherwise returns false.

This function was introduced in  Qt 4.2.

See also removeAllApplicationFonts(), addApplicationFont(), and addApplicationFontFromData().
*/
func (this *QFontDatabase) RemoveApplicationFont(id int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase21removeApplicationFontEi", qtrt.FFI_TYPE_POINTER, id)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFontDatabase_RemoveApplicationFont(id int) bool {
	var nilthis *QFontDatabase
	rv := nilthis.RemoveApplicationFont(id)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:151
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool removeAllApplicationFonts()

/*
Removes all application-local fonts previously added using addApplicationFont() and addApplicationFontFromData().

Returns true if unloading of the fonts succeeded; otherwise returns false.

This function was introduced in  Qt 4.2.

See also removeApplicationFont(), addApplicationFont(), and addApplicationFontFromData().
*/
func (this *QFontDatabase) RemoveAllApplicationFonts() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase25removeAllApplicationFontsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFontDatabase_RemoveAllApplicationFonts() bool {
	var nilthis *QFontDatabase
	rv := nilthis.RemoveAllApplicationFonts()
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:154
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool supportsThreadedFontRendering()

/*

 */
func (this *QFontDatabase) SupportsThreadedFontRendering() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase29supportsThreadedFontRenderingEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QFontDatabase_SupportsThreadedFontRendering() bool {
	var nilthis *QFontDatabase
	rv := nilthis.SupportsThreadedFontRendering()
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:157
// index:0
// Public static Visibility=Default Availability=Available
// [16] QFont systemFont(QFontDatabase::SystemFont)

/*
Returns the most adequate font for a given type case for proper integration with the system's look and feel.

This function was introduced in  Qt 5.2.

See also QGuiApplication::font().
*/
func (this *QFontDatabase) SystemFont(type_ int) *QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabase10systemFontENS_10SystemFontE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQFont)
	return rv2
}
func QFontDatabase_SystemFont(type_ int) *QFont /*123*/ {
	var nilthis *QFontDatabase
	rv := nilthis.SystemFont(type_)
	return rv
}

func DeleteQFontDatabase(this *QFontDatabase) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QFontDatabaseD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
QFontDatabase::OtherSymbol(the same as Symbol)
QFontDatabase::Ogham?
QFontDatabase::Runic?
QFontDatabase::Nko?

*/
type QFontDatabase__WritingSystem = int

//
const QFontDatabase__Any QFontDatabase__WritingSystem = 0

//
const QFontDatabase__Latin QFontDatabase__WritingSystem = 1

//
const QFontDatabase__Greek QFontDatabase__WritingSystem = 2

//
const QFontDatabase__Cyrillic QFontDatabase__WritingSystem = 3

//
const QFontDatabase__Armenian QFontDatabase__WritingSystem = 4

//
const QFontDatabase__Hebrew QFontDatabase__WritingSystem = 5

//
const QFontDatabase__Arabic QFontDatabase__WritingSystem = 6

//
const QFontDatabase__Syriac QFontDatabase__WritingSystem = 7

//
const QFontDatabase__Thaana QFontDatabase__WritingSystem = 8

//
const QFontDatabase__Devanagari QFontDatabase__WritingSystem = 9

//
const QFontDatabase__Bengali QFontDatabase__WritingSystem = 10

//
const QFontDatabase__Gurmukhi QFontDatabase__WritingSystem = 11

//
const QFontDatabase__Gujarati QFontDatabase__WritingSystem = 12

//
const QFontDatabase__Oriya QFontDatabase__WritingSystem = 13

//
const QFontDatabase__Tamil QFontDatabase__WritingSystem = 14

//
const QFontDatabase__Telugu QFontDatabase__WritingSystem = 15

//
const QFontDatabase__Kannada QFontDatabase__WritingSystem = 16

//
const QFontDatabase__Malayalam QFontDatabase__WritingSystem = 17

//
const QFontDatabase__Sinhala QFontDatabase__WritingSystem = 18

//
const QFontDatabase__Thai QFontDatabase__WritingSystem = 19

//
const QFontDatabase__Lao QFontDatabase__WritingSystem = 20

//
const QFontDatabase__Tibetan QFontDatabase__WritingSystem = 21

//
const QFontDatabase__Myanmar QFontDatabase__WritingSystem = 22

//
const QFontDatabase__Georgian QFontDatabase__WritingSystem = 23

//
const QFontDatabase__Khmer QFontDatabase__WritingSystem = 24

//
const QFontDatabase__SimplifiedChinese QFontDatabase__WritingSystem = 25

//
const QFontDatabase__TraditionalChinese QFontDatabase__WritingSystem = 26

//
const QFontDatabase__Japanese QFontDatabase__WritingSystem = 27

//
const QFontDatabase__Korean QFontDatabase__WritingSystem = 28

//
const QFontDatabase__Vietnamese QFontDatabase__WritingSystem = 29

//
const QFontDatabase__Symbol QFontDatabase__WritingSystem = 30

//
const QFontDatabase__Other QFontDatabase__WritingSystem = 30

//
const QFontDatabase__Ogham QFontDatabase__WritingSystem = 31

//
const QFontDatabase__Runic QFontDatabase__WritingSystem = 32

//
const QFontDatabase__Nko QFontDatabase__WritingSystem = 33

//
const QFontDatabase__WritingSystemsCount QFontDatabase__WritingSystem = 34

/*


This enum was introduced or modified in  Qt 5.2.

*/
type QFontDatabase__SystemFont = int

// The default system font.
const QFontDatabase__GeneralFont QFontDatabase__SystemFont = 0

// The fixed font that the system recommends.
const QFontDatabase__FixedFont QFontDatabase__SystemFont = 1

// The system standard font for titles.
const QFontDatabase__TitleFont QFontDatabase__SystemFont = 2

// The smallest readable system font.
const QFontDatabase__SmallestReadableFont QFontDatabase__SystemFont = 3

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
