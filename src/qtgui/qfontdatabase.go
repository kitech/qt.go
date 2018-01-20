//  header block begin
// /usr/include/qt/QtGui/qfontdatabase.h
// #include <qfontdatabase.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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
type QFontDatabase struct {
	*qtrt.CObject
}

func (this *QFontDatabase) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQFontDatabaseFromPointer(cthis unsafe.Pointer) *QFontDatabase {
	return &QFontDatabase{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qfontdatabase.h:116
// index:0
// Public static
// QList<int> standardSizes()
func (this *QFontDatabase) StandardSizes() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase13standardSizesEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_StandardSizes() {
	var nilthis *QFontDatabase
	nilthis.StandardSizes()
}

// /usr/include/qt/QtGui/qfontdatabase.h:118
// index:0
// Public
// void QFontDatabase()
func NewQFontDatabase() *QFontDatabase {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabaseC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDatabaseFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qfontdatabase.h:120
// index:0
// Public
// QList<QFontDatabase::WritingSystem> writingSystems()
func (this *QFontDatabase) WritingSystems() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase14writingSystemsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:121
// index:1
// Public
// QList<QFontDatabase::WritingSystem> writingSystems(const class QString &)
func (this *QFontDatabase) WritingSystems_1(family *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase14writingSystemsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:123
// index:0
// Public
// QStringList families(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) Families(writingSystem int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase8familiesENS_13WritingSystemE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &writingSystem)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:124
// index:0
// Public
// QStringList styles(const class QString &)
func (this *QFontDatabase) Styles(family *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6stylesERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:125
// index:0
// Public
// QList<int> pointSizes(const class QString &, const class QString &)
func (this *QFontDatabase) PointSizes(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase10pointSizesERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:126
// index:0
// Public
// QList<int> smoothSizes(const class QString &, const class QString &)
func (this *QFontDatabase) SmoothSizes(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11smoothSizesERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:127
// index:0
// Public
// QString styleString(const class QFont &)
func (this *QFontDatabase) StyleString(font *QFont) interface{} {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:128
// index:1
// Public
// QString styleString(const class QFontInfo &)
func (this *QFontDatabase) StyleString_1(fontInfo *QFontInfo) interface{} {
	var convArg0 = fontInfo.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK9QFontInfo", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:130
// index:0
// Public
// QFont font(const class QString &, const class QString &, int)
func (this *QFontDatabase) Font(family *qtcore.QString, style *qtcore.QString, pointSize int) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase4fontERK7QStringS2_i", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, &pointSize)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:132
// index:0
// Public
// bool isBitmapScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsBitmapScalable(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:133
// index:0
// Public
// bool isSmoothlyScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsSmoothlyScalable(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:134
// index:0
// Public
// bool isScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsScalable(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase10isScalableERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:135
// index:0
// Public
// bool isFixedPitch(const class QString &, const class QString &)
func (this *QFontDatabase) IsFixedPitch(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:137
// index:0
// Public
// bool italic(const class QString &, const class QString &)
func (this *QFontDatabase) Italic(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6italicERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:138
// index:0
// Public
// bool bold(const class QString &, const class QString &)
func (this *QFontDatabase) Bold(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase4boldERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:139
// index:0
// Public
// int weight(const class QString &, const class QString &)
func (this *QFontDatabase) Weight(family *qtcore.QString, style *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	var convArg1 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6weightERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:141
// index:0
// Public
// bool hasFamily(const class QString &)
func (this *QFontDatabase) HasFamily(family *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase9hasFamilyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:142
// index:0
// Public
// bool isPrivateFamily(const class QString &)
func (this *QFontDatabase) IsPrivateFamily(family *qtcore.QString) interface{} {
	var convArg0 = family.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase15isPrivateFamilyERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qfontdatabase.h:144
// index:0
// Public static
// QString writingSystemName(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) WritingSystemName(writingSystem int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase17writingSystemNameENS_13WritingSystemE", ffiqt.FFI_TYPE_POINTER, writingSystem)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_WritingSystemName(writingSystem int) {
	var nilthis *QFontDatabase
	nilthis.WritingSystemName(writingSystem)
}

// /usr/include/qt/QtGui/qfontdatabase.h:145
// index:0
// Public static
// QString writingSystemSample(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) WritingSystemSample(writingSystem int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase19writingSystemSampleENS_13WritingSystemE", ffiqt.FFI_TYPE_POINTER, writingSystem)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_WritingSystemSample(writingSystem int) {
	var nilthis *QFontDatabase
	nilthis.WritingSystemSample(writingSystem)
}

// /usr/include/qt/QtGui/qfontdatabase.h:147
// index:0
// Public static
// int addApplicationFont(const class QString &)
func (this *QFontDatabase) AddApplicationFont(fileName *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase18addApplicationFontERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_AddApplicationFont(fileName *qtcore.QString) {
	var nilthis *QFontDatabase
	nilthis.AddApplicationFont(fileName)
}

// /usr/include/qt/QtGui/qfontdatabase.h:148
// index:0
// Public static
// int addApplicationFontFromData(const class QByteArray &)
func (this *QFontDatabase) AddApplicationFontFromData(fontData *qtcore.QByteArray) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray", ffiqt.FFI_TYPE_POINTER, fontData)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_AddApplicationFontFromData(fontData *qtcore.QByteArray) {
	var nilthis *QFontDatabase
	nilthis.AddApplicationFontFromData(fontData)
}

// /usr/include/qt/QtGui/qfontdatabase.h:149
// index:0
// Public static
// QStringList applicationFontFamilies(int)
func (this *QFontDatabase) ApplicationFontFamilies(id int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase23applicationFontFamiliesEi", ffiqt.FFI_TYPE_POINTER, id)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_ApplicationFontFamilies(id int) {
	var nilthis *QFontDatabase
	nilthis.ApplicationFontFamilies(id)
}

// /usr/include/qt/QtGui/qfontdatabase.h:150
// index:0
// Public static
// bool removeApplicationFont(int)
func (this *QFontDatabase) RemoveApplicationFont(id int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase21removeApplicationFontEi", ffiqt.FFI_TYPE_POINTER, id)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_RemoveApplicationFont(id int) {
	var nilthis *QFontDatabase
	nilthis.RemoveApplicationFont(id)
}

// /usr/include/qt/QtGui/qfontdatabase.h:151
// index:0
// Public static
// bool removeAllApplicationFonts()
func (this *QFontDatabase) RemoveAllApplicationFonts() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase25removeAllApplicationFontsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_RemoveAllApplicationFonts() {
	var nilthis *QFontDatabase
	nilthis.RemoveAllApplicationFonts()
}

// /usr/include/qt/QtGui/qfontdatabase.h:154
// index:0
// Public static
// bool supportsThreadedFontRendering()
func (this *QFontDatabase) SupportsThreadedFontRendering() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase29supportsThreadedFontRenderingEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_SupportsThreadedFontRendering() {
	var nilthis *QFontDatabase
	nilthis.SupportsThreadedFontRendering()
}

// /usr/include/qt/QtGui/qfontdatabase.h:157
// index:0
// Public static
// QFont systemFont(enum QFontDatabase::SystemFont)
func (this *QFontDatabase) SystemFont(type_ int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase10systemFontENS_10SystemFontE", ffiqt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	return rv
}
func QFontDatabase_SystemFont(type_ int) {
	var nilthis *QFontDatabase
	nilthis.SystemFont(type_)
}

//  body block end
