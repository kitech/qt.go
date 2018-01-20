//  header block begin
// /usr/include/qt/QtCore/qsettings.h
// #include <qsettings.h>
// #include <QtCore>
package qtcore

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
	*QObject
}

func (this *QSettings) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQSettingsFromPointer(cthis unsafe.Pointer) *QSettings {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSettings{bcthis0}
}

// /usr/include/qt/QtCore/qsettings.h:71
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QSettings) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:127
// index:0
// Public
// void QSettings(const class QString &, const class QString &, class QObject *)
func NewQSettings(organization *QString, application *QString, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = organization.GetCthis()
	var convArg1 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringS2_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:129
// index:1
// Public
// void QSettings(enum QSettings::Scope, const class QString &, const class QString &, class QObject *)
func NewQSettings_1(scope int, organization *QString, application *QString, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = organization.GetCthis()
	var convArg2 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ENS_5ScopeERK7QStringS3_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, &scope, convArg1, convArg2, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:131
// index:2
// Public
// void QSettings(enum QSettings::Format, enum QSettings::Scope, const class QString &, const class QString &, class QObject *)
func NewQSettings_2(format int, scope int, organization *QString, application *QString, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg2 = organization.GetCthis()
	var convArg3 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ENS_6FormatENS_5ScopeERK7QStringS4_P7QObject", ffiqt.FFI_TYPE_VOID, cthis, &format, &scope, convArg2, convArg3, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:133
// index:3
// Public
// void QSettings(const class QString &, enum QSettings::Format, class QObject *)
func NewQSettings_3(fileName *QString, format int, parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringNS_6FormatEP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &format, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:134
// index:4
// Public
// void QSettings(class QObject *)
func NewQSettings_4(parent unsafe.Pointer) *QSettings {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:144
// index:0
// Public virtual
// void ~QSettings()
func DeleteQSettings(*QSettings) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettingsD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:146
// index:0
// Public
// void clear()
func (this *QSettings) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:147
// index:0
// Public
// void sync()
func (this *QSettings) Sync() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings4syncEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:148
// index:0
// Public
// QSettings::Status status()
func (this *QSettings) Status() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:149
// index:0
// Public
// bool isAtomicSyncRequired()
func (this *QSettings) IsAtomicSyncRequired() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings20isAtomicSyncRequiredEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:150
// index:0
// Public
// void setAtomicSyncRequired(_Bool)
func (this *QSettings) SetAtomicSyncRequired(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings21setAtomicSyncRequiredEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:152
// index:0
// Public
// void beginGroup(const class QString &)
func (this *QSettings) BeginGroup(prefix *QString) {
	var convArg0 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings10beginGroupERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:153
// index:0
// Public
// void endGroup()
func (this *QSettings) EndGroup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings8endGroupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:154
// index:0
// Public
// QString group()
func (this *QSettings) Group() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings5groupEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:156
// index:0
// Public
// int beginReadArray(const class QString &)
func (this *QSettings) BeginReadArray(prefix *QString) interface{} {
	var convArg0 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings14beginReadArrayERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:157
// index:0
// Public
// void beginWriteArray(const class QString &, int)
func (this *QSettings) BeginWriteArray(prefix *QString, size int) {
	var convArg0 = prefix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings15beginWriteArrayERK7QStringi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:158
// index:0
// Public
// void endArray()
func (this *QSettings) EndArray() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings8endArrayEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:159
// index:0
// Public
// void setArrayIndex(int)
func (this *QSettings) SetArrayIndex(i int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings13setArrayIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:161
// index:0
// Public
// QStringList allKeys()
func (this *QSettings) AllKeys() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings7allKeysEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:162
// index:0
// Public
// QStringList childKeys()
func (this *QSettings) ChildKeys() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings9childKeysEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:163
// index:0
// Public
// QStringList childGroups()
func (this *QSettings) ChildGroups() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings11childGroupsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:164
// index:0
// Public
// bool isWritable()
func (this *QSettings) IsWritable() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings10isWritableEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:166
// index:0
// Public
// void setValue(const class QString &, const class QVariant &)
func (this *QSettings) SetValue(key *QString, value *QVariant) {
	var convArg0 = key.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings8setValueERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:167
// index:0
// Public
// QVariant value(const class QString &, const class QVariant &)
func (this *QSettings) Value(key *QString, defaultValue *QVariant) interface{} {
	var convArg0 = key.GetCthis()
	var convArg1 = defaultValue.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings5valueERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:169
// index:0
// Public
// void remove(const class QString &)
func (this *QSettings) Remove(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings6removeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:170
// index:0
// Public
// bool contains(const class QString &)
func (this *QSettings) Contains(key *QString) interface{} {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings8containsERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:172
// index:0
// Public
// void setFallbacksEnabled(_Bool)
func (this *QSettings) SetFallbacksEnabled(b bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings19setFallbacksEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:173
// index:0
// Public
// bool fallbacksEnabled()
func (this *QSettings) FallbacksEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings16fallbacksEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:175
// index:0
// Public
// QString fileName()
func (this *QSettings) FileName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings8fileNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:176
// index:0
// Public
// QSettings::Format format()
func (this *QSettings) Format() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:177
// index:0
// Public
// QSettings::Scope scope()
func (this *QSettings) Scope() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings5scopeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:178
// index:0
// Public
// QString organizationName()
func (this *QSettings) OrganizationName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings16organizationNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:179
// index:0
// Public
// QString applicationName()
func (this *QSettings) ApplicationName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings15applicationNameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:182
// index:0
// Public
// void setIniCodec(class QTextCodec *)
func (this *QSettings) SetIniCodec(codec unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEP10QTextCodec", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), codec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:183
// index:1
// Public
// void setIniCodec(const char *)
func (this *QSettings) SetIniCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:184
// index:0
// Public
// QTextCodec * iniCodec()
func (this *QSettings) IniCodec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSettings8iniCodecEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:187
// index:0
// Public static
// void setDefaultFormat(enum QSettings::Format)
func (this *QSettings) SetDefaultFormat(format int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings16setDefaultFormatENS_6FormatE", ffiqt.FFI_TYPE_POINTER, format)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetDefaultFormat(format int) {
	var nilthis *QSettings
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtCore/qsettings.h:188
// index:0
// Public static
// QSettings::Format defaultFormat()
func (this *QSettings) DefaultFormat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings13defaultFormatEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QSettings_DefaultFormat() {
	var nilthis *QSettings
	nilthis.DefaultFormat()
}

// /usr/include/qt/QtCore/qsettings.h:189
// index:0
// Public static
// void setSystemIniPath(const class QString &)
func (this *QSettings) SetSystemIniPath(dir *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings16setSystemIniPathERK7QString", ffiqt.FFI_TYPE_POINTER, dir)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetSystemIniPath(dir *QString) {
	var nilthis *QSettings
	nilthis.SetSystemIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:190
// index:0
// Public static
// void setUserIniPath(const class QString &)
func (this *QSettings) SetUserIniPath(dir *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings14setUserIniPathERK7QString", ffiqt.FFI_TYPE_POINTER, dir)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetUserIniPath(dir *QString) {
	var nilthis *QSettings
	nilthis.SetUserIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:191
// index:0
// Public static
// void setPath(enum QSettings::Format, enum QSettings::Scope, const class QString &)
func (this *QSettings) SetPath(format int, scope int, path *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings7setPathENS_6FormatENS_5ScopeERK7QString", ffiqt.FFI_TYPE_POINTER, format, scope, path)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetPath(format int, scope int, path *QString) {
	var nilthis *QSettings
	nilthis.SetPath(format, scope, path)
}

// /usr/include/qt/QtCore/qsettings.h:202
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QSettings) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSettings5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
