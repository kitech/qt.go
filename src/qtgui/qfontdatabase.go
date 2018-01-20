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

// /usr/include/qt/QtGui/qfontdatabase.h:116
// index:0
// static
// QList<int> standardSizes()
func (this *QFontDatabase) StandardSizes() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase13standardSizesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_StandardSizes() {
	// 0: (), ()
	var nilthis *QFontDatabase
	nilthis.StandardSizes()
}

// /usr/include/qt/QtGui/qfontdatabase.h:118
// index:0
// void QFontDatabase()
func NewQFontDatabase() *QFontDatabase {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabaseC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQFontDatabaseFromPointer(cthis)
	return gothis
}
func NewQFontDatabaseFromPointer(cthis unsafe.Pointer) *QFontDatabase {
	return &QFontDatabase{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qfontdatabase.h:120
// index:0
// QList<QFontDatabase::WritingSystem> writingSystems()
func (this *QFontDatabase) WritingSystems() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase14writingSystemsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:121
// index:1
// QList<QFontDatabase::WritingSystem> writingSystems(const class QString &)
func (this *QFontDatabase) WritingSystems_1(family unsafe.Pointer) {
	// 1: (, family const QString &), (family)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase14writingSystemsERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:123
// index:0
// QStringList families(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) Families(writingSystem int) {
	// 0: (, writingSystem QFontDatabase::WritingSystem), (&writingSystem)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase8familiesENS_13WritingSystemE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &writingSystem)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:124
// index:0
// QStringList styles(const class QString &)
func (this *QFontDatabase) Styles(family unsafe.Pointer) {
	// 0: (, family const QString &), (family)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6stylesERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:125
// index:0
// QList<int> pointSizes(const class QString &, const class QString &)
func (this *QFontDatabase) PointSizes(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase10pointSizesERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:126
// index:0
// QList<int> smoothSizes(const class QString &, const class QString &)
func (this *QFontDatabase) SmoothSizes(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11smoothSizesERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:127
// index:0
// QString styleString(const class QFont &)
func (this *QFontDatabase) StyleString(font unsafe.Pointer) {
	// 0: (, font const QFont &), (font)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK5QFont", ffiqt.FFI_TYPE_VOID, this.GetCthis(), font)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:128
// index:1
// QString styleString(const class QFontInfo &)
func (this *QFontDatabase) StyleString_1(fontInfo unsafe.Pointer) {
	// 1: (, fontInfo const QFontInfo &), (fontInfo)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase11styleStringERK9QFontInfo", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fontInfo)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:130
// index:0
// QFont font(const class QString &, const class QString &, int)
func (this *QFontDatabase) Font(family unsafe.Pointer, style unsafe.Pointer, pointSize int) {
	// 0: (, family const QString &, style const QString &, pointSize int), (family, style, &pointSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase4fontERK7QStringS2_i", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style, &pointSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:132
// index:0
// bool isBitmapScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsBitmapScalable(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase16isBitmapScalableERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:133
// index:0
// bool isSmoothlyScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsSmoothlyScalable(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase18isSmoothlyScalableERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:134
// index:0
// bool isScalable(const class QString &, const class QString &)
func (this *QFontDatabase) IsScalable(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase10isScalableERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:135
// index:0
// bool isFixedPitch(const class QString &, const class QString &)
func (this *QFontDatabase) IsFixedPitch(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase12isFixedPitchERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:137
// index:0
// bool italic(const class QString &, const class QString &)
func (this *QFontDatabase) Italic(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6italicERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:138
// index:0
// bool bold(const class QString &, const class QString &)
func (this *QFontDatabase) Bold(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase4boldERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:139
// index:0
// int weight(const class QString &, const class QString &)
func (this *QFontDatabase) Weight(family unsafe.Pointer, style unsafe.Pointer) {
	// 0: (, family const QString &, style const QString &), (family, style)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase6weightERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family, style)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:141
// index:0
// bool hasFamily(const class QString &)
func (this *QFontDatabase) HasFamily(family unsafe.Pointer) {
	// 0: (, family const QString &), (family)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase9hasFamilyERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:142
// index:0
// bool isPrivateFamily(const class QString &)
func (this *QFontDatabase) IsPrivateFamily(family unsafe.Pointer) {
	// 0: (, family const QString &), (family)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QFontDatabase15isPrivateFamilyERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), family)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qfontdatabase.h:144
// index:0
// static
// QString writingSystemName(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) WritingSystemName(writingSystem int) {
	// 0: (writingSystem QFontDatabase::WritingSystem), (writingSystem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase17writingSystemNameENS_13WritingSystemE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_WritingSystemName(writingSystem int) {
	// 0: (writingSystem QFontDatabase::WritingSystem), (writingSystem)
	var nilthis *QFontDatabase
	nilthis.WritingSystemName(writingSystem)
}

// /usr/include/qt/QtGui/qfontdatabase.h:145
// index:0
// static
// QString writingSystemSample(enum QFontDatabase::WritingSystem)
func (this *QFontDatabase) WritingSystemSample(writingSystem int) {
	// 0: (writingSystem QFontDatabase::WritingSystem), (writingSystem)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase19writingSystemSampleENS_13WritingSystemE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_WritingSystemSample(writingSystem int) {
	// 0: (writingSystem QFontDatabase::WritingSystem), (writingSystem)
	var nilthis *QFontDatabase
	nilthis.WritingSystemSample(writingSystem)
}

// /usr/include/qt/QtGui/qfontdatabase.h:147
// index:0
// static
// int addApplicationFont(const class QString &)
func (this *QFontDatabase) AddApplicationFont(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase18addApplicationFontERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_AddApplicationFont(fileName unsafe.Pointer) {
	// 0: (fileName const QString &), (fileName)
	var nilthis *QFontDatabase
	nilthis.AddApplicationFont(fileName)
}

// /usr/include/qt/QtGui/qfontdatabase.h:148
// index:0
// static
// int addApplicationFontFromData(const class QByteArray &)
func (this *QFontDatabase) AddApplicationFontFromData(fontData unsafe.Pointer) {
	// 0: (fontData const QByteArray &), (fontData)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase26addApplicationFontFromDataERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_AddApplicationFontFromData(fontData unsafe.Pointer) {
	// 0: (fontData const QByteArray &), (fontData)
	var nilthis *QFontDatabase
	nilthis.AddApplicationFontFromData(fontData)
}

// /usr/include/qt/QtGui/qfontdatabase.h:149
// index:0
// static
// QStringList applicationFontFamilies(int)
func (this *QFontDatabase) ApplicationFontFamilies(id int) {
	// 0: (id int), (id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase23applicationFontFamiliesEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_ApplicationFontFamilies(id int) {
	// 0: (id int), (id)
	var nilthis *QFontDatabase
	nilthis.ApplicationFontFamilies(id)
}

// /usr/include/qt/QtGui/qfontdatabase.h:150
// index:0
// static
// bool removeApplicationFont(int)
func (this *QFontDatabase) RemoveApplicationFont(id int) {
	// 0: (id int), (id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase21removeApplicationFontEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_RemoveApplicationFont(id int) {
	// 0: (id int), (id)
	var nilthis *QFontDatabase
	nilthis.RemoveApplicationFont(id)
}

// /usr/include/qt/QtGui/qfontdatabase.h:151
// index:0
// static
// bool removeAllApplicationFonts()
func (this *QFontDatabase) RemoveAllApplicationFonts() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase25removeAllApplicationFontsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_RemoveAllApplicationFonts() {
	// 0: (), ()
	var nilthis *QFontDatabase
	nilthis.RemoveAllApplicationFonts()
}

// /usr/include/qt/QtGui/qfontdatabase.h:154
// index:0
// static
// bool supportsThreadedFontRendering()
func (this *QFontDatabase) SupportsThreadedFontRendering() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase29supportsThreadedFontRenderingEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_SupportsThreadedFontRendering() {
	// 0: (), ()
	var nilthis *QFontDatabase
	nilthis.SupportsThreadedFontRendering()
}

// /usr/include/qt/QtGui/qfontdatabase.h:157
// index:0
// static
// QFont systemFont(enum QFontDatabase::SystemFont)
func (this *QFontDatabase) SystemFont(type_ int) {
	// 0: (type QFontDatabase::SystemFont), (type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QFontDatabase10systemFontENS_10SystemFontE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QFontDatabase_SystemFont(type_ int) {
	// 0: (type QFontDatabase::SystemFont), (type_)
	var nilthis *QFontDatabase
	nilthis.SystemFont(type_)
}

//  body block end
