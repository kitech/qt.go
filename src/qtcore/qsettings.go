//  header block begin
// /usr/include/qt/QtCore/qsettings.h
// #include <qsettings.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

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
}

//  ext block end

//  body block begin
type QSettings struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtCore/qsettings.h:71
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QSettings) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:127
// index:0
// void QSettings(const class QString &, const class QString &, class QObject *)
func NewQSettings(organization unsafe.Pointer, application unsafe.Pointer, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringS2_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, organization, application, parent)
	gopp.ErrPrint(err, rv)
	return &QSettings{cthis}
}

// /usr/include/qt/QtCore/qsettings.h:129
// index:1
// void QSettings(enum QSettings::Scope, const class QString &, const class QString &, class QObject *)
func NewQSettings_1(scope int, organization unsafe.Pointer, application unsafe.Pointer, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ENS_5ScopeERK7QStringS3_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, &scope, organization, application, parent)
	gopp.ErrPrint(err, rv)
	return &QSettings{cthis}
}

// /usr/include/qt/QtCore/qsettings.h:131
// index:2
// void QSettings(enum QSettings::Format, enum QSettings::Scope, const class QString &, const class QString &, class QObject *)
func NewQSettings_2(format int, scope int, organization unsafe.Pointer, application unsafe.Pointer, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ENS_6FormatENS_5ScopeERK7QStringS4_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, &format, &scope, organization, application, parent)
	gopp.ErrPrint(err, rv)
	return &QSettings{cthis}
}

// /usr/include/qt/QtCore/qsettings.h:133
// index:3
// void QSettings(const class QString &, enum QSettings::Format, class QObject *)
func NewQSettings_3(fileName unsafe.Pointer, format int, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringNS_6FormatEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, fileName, &format, parent)
	gopp.ErrPrint(err, rv)
	return &QSettings{cthis}
}

// /usr/include/qt/QtCore/qsettings.h:134
// index:4
// void QSettings(class QObject *)
func NewQSettings_4(parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QSettings{cthis}
}

// /usr/include/qt/QtCore/qsettings.h:144
// index:0
// virtual
// void ~QSettings()
func DeleteQSettings(*QSettings) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:146
// index:0
// void clear()
func (this *QSettings) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings5clearEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:147
// index:0
// void sync()
func (this *QSettings) Sync() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings4syncEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:148
// index:0
// QSettings::Status status()
func (this *QSettings) Status() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings6statusEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:149
// index:0
// bool isAtomicSyncRequired()
func (this *QSettings) IsAtomicSyncRequired() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings20isAtomicSyncRequiredEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:150
// index:0
// void setAtomicSyncRequired(_Bool)
func (this *QSettings) SetAtomicSyncRequired(enable bool) {
	// 0: (, bool enable), (&enable)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings21setAtomicSyncRequiredEb", ffiqt.FFI_TYPE_VOID, this.cthis, &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:152
// index:0
// void beginGroup(const class QString &)
func (this *QSettings) BeginGroup(prefix unsafe.Pointer) {
	// 0: (, const QString & prefix), (prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings10beginGroupERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:153
// index:0
// void endGroup()
func (this *QSettings) EndGroup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings8endGroupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:154
// index:0
// QString group()
func (this *QSettings) Group() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings5groupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:156
// index:0
// int beginReadArray(const class QString &)
func (this *QSettings) BeginReadArray(prefix unsafe.Pointer) {
	// 0: (, const QString & prefix), (prefix)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings14beginReadArrayERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, prefix)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:157
// index:0
// void beginWriteArray(const class QString &, int)
func (this *QSettings) BeginWriteArray(prefix unsafe.Pointer, size int) {
	// 0: (, const QString & prefix, int size), (prefix, &size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings15beginWriteArrayERK7QStringi", ffiqt.FFI_TYPE_VOID, this.cthis, prefix, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:158
// index:0
// void endArray()
func (this *QSettings) EndArray() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings8endArrayEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:159
// index:0
// void setArrayIndex(int)
func (this *QSettings) SetArrayIndex(i int) {
	// 0: (, int i), (&i)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings13setArrayIndexEi", ffiqt.FFI_TYPE_VOID, this.cthis, &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:161
// index:0
// QStringList allKeys()
func (this *QSettings) AllKeys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings7allKeysEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:162
// index:0
// QStringList childKeys()
func (this *QSettings) ChildKeys() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings9childKeysEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:163
// index:0
// QStringList childGroups()
func (this *QSettings) ChildGroups() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings11childGroupsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:164
// index:0
// bool isWritable()
func (this *QSettings) IsWritable() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings10isWritableEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:166
// index:0
// void setValue(const class QString &, const class QVariant &)
func (this *QSettings) SetValue(key unsafe.Pointer, value unsafe.Pointer) {
	// 0: (, const QString & key, const QVariant & value), (key, value)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings8setValueERK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, key, value)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:167
// index:0
// QVariant value(const class QString &, const class QVariant &)
func (this *QSettings) Value(key unsafe.Pointer, defaultValue unsafe.Pointer) {
	// 0: (, const QString & key, const QVariant & defaultValue), (key, defaultValue)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings5valueERK7QStringRK8QVariant", ffiqt.FFI_TYPE_VOID, this.cthis, key, defaultValue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:169
// index:0
// void remove(const class QString &)
func (this *QSettings) Remove(key unsafe.Pointer) {
	// 0: (, const QString & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings6removeERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:170
// index:0
// bool contains(const class QString &)
func (this *QSettings) Contains(key unsafe.Pointer) {
	// 0: (, const QString & key), (key)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings8containsERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, key)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:172
// index:0
// void setFallbacksEnabled(_Bool)
func (this *QSettings) SetFallbacksEnabled(b bool) {
	// 0: (, bool b), (&b)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings19setFallbacksEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:173
// index:0
// bool fallbacksEnabled()
func (this *QSettings) FallbacksEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings16fallbacksEnabledEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:175
// index:0
// QString fileName()
func (this *QSettings) FileName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings8fileNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:176
// index:0
// QSettings::Format format()
func (this *QSettings) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:177
// index:0
// QSettings::Scope scope()
func (this *QSettings) Scope() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings5scopeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:178
// index:0
// QString organizationName()
func (this *QSettings) OrganizationName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings16organizationNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:179
// index:0
// QString applicationName()
func (this *QSettings) ApplicationName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings15applicationNameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:182
// index:0
// void setIniCodec(class QTextCodec *)
func (this *QSettings) SetIniCodec(codec unsafe.Pointer) {
	// 0: (, QTextCodec * codec), (codec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEP10QTextCodec", ffiqt.FFI_TYPE_VOID, this.cthis, codec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:183
// index:1
// void setIniCodec(const char *)
func (this *QSettings) SetIniCodec_1(codecName unsafe.Pointer) {
	// 1: (, const char * codecName), (codecName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEPKc", ffiqt.FFI_TYPE_VOID, this.cthis, codecName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:184
// index:0
// QTextCodec * iniCodec()
func (this *QSettings) IniCodec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings8iniCodecEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:187
// index:0
// static
// void setDefaultFormat(enum QSettings::Format)
func (this *QSettings) SetDefaultFormat(format int) {
	// 0: (QSettings::Format format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings16setDefaultFormatENS_6FormatE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetDefaultFormat(format int) {
	// 0: (QSettings::Format format), (format)
	var nilthis *QSettings
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtCore/qsettings.h:188
// index:0
// static
// QSettings::Format defaultFormat()
func (this *QSettings) DefaultFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings13defaultFormatEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSettings_DefaultFormat() {
	// 0: (), ()
	var nilthis *QSettings
	nilthis.DefaultFormat()
}

// /usr/include/qt/QtCore/qsettings.h:189
// index:0
// static
// void setSystemIniPath(const class QString &)
func (this *QSettings) SetSystemIniPath(dir unsafe.Pointer) {
	// 0: (const QString & dir), (dir)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings16setSystemIniPathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetSystemIniPath(dir unsafe.Pointer) {
	// 0: (const QString & dir), (dir)
	var nilthis *QSettings
	nilthis.SetSystemIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:190
// index:0
// static
// void setUserIniPath(const class QString &)
func (this *QSettings) SetUserIniPath(dir unsafe.Pointer) {
	// 0: (const QString & dir), (dir)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings14setUserIniPathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetUserIniPath(dir unsafe.Pointer) {
	// 0: (const QString & dir), (dir)
	var nilthis *QSettings
	nilthis.SetUserIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:191
// index:0
// static
// void setPath(enum QSettings::Format, enum QSettings::Scope, const class QString &)
func (this *QSettings) SetPath(format int, scope int, path unsafe.Pointer) {
	// 0: (QSettings::Format format, QSettings::Scope scope, const QString & path), (format, scope, path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings7setPathENS_6FormatENS_5ScopeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetPath(format int, scope int, path unsafe.Pointer) {
	// 0: (QSettings::Format format, QSettings::Scope scope, const QString & path), (format, scope, path)
	var nilthis *QSettings
	nilthis.SetPath(format, scope, path)
}

//  body block end
