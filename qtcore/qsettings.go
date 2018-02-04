package qtcore

// /usr/include/qt/QtCore/qsettings.h
// #include <qsettings.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/qtrt"

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
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QSettings) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QSettings struct {
	*QObject
}

func (this *QSettings) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSettings) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQSettingsFromPointer(cthis unsafe.Pointer) *QSettings {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QSettings{bcthis0}
}
func (*QSettings) NewFromPointer(cthis unsafe.Pointer) *QSettings {
	return NewQSettingsFromPointer(cthis)
}

// /usr/include/qt/QtCore/qsettings.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSettings) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSettings(const QString &, const QString &, QObject *)
func NewQSettings(organization *QString, application *QString, parent *QObject /*777 QObject **/) *QSettings {
	var convArg0 = organization.GetCthis()
	var convArg1 = application.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringS2_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:129
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSettings(enum QSettings::Scope, const QString &, const QString &, QObject *)
func NewQSettings_1(scope int, organization *QString, application *QString, parent *QObject /*777 QObject **/) *QSettings {
	var convArg1 = organization.GetCthis()
	var convArg2 = application.GetCthis()
	var convArg3 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_5ScopeERK7QStringS3_P7QObject", qtrt.FFI_TYPE_POINTER, scope, convArg1, convArg2, convArg3)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:131
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSettings(enum QSettings::Format, enum QSettings::Scope, const QString &, const QString &, QObject *)
func NewQSettings_2(format int, scope int, organization *QString, application *QString, parent *QObject /*777 QObject **/) *QSettings {
	var convArg2 = organization.GetCthis()
	var convArg3 = application.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_6FormatENS_5ScopeERK7QStringS4_P7QObject", qtrt.FFI_TYPE_POINTER, format, scope, convArg2, convArg3, convArg4)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:133
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QSettings(const QString &, enum QSettings::Format, QObject *)
func NewQSettings_3(fileName *QString, format int, parent *QObject /*777 QObject **/) *QSettings {
	var convArg0 = fileName.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringNS_6FormatEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, format, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:134
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QObject *)
func NewQSettings_4(parent *QObject /*777 QObject **/) *QSettings {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSettings()
func DeleteQSettings(this *QSettings) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsettings.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QSettings) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sync()
func (this *QSettings) Sync() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings4syncEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:148
// index:0
// Public Visibility=Default Availability=Available
// [4] QSettings::Status status()
func (this *QSettings) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qsettings.h:149
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAtomicSyncRequired()
func (this *QSettings) IsAtomicSyncRequired() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings20isAtomicSyncRequiredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAtomicSyncRequired(_Bool)
func (this *QSettings) SetAtomicSyncRequired(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings21setAtomicSyncRequiredEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginGroup(const QString &)
func (this *QSettings) BeginGroup(prefix *QString) {
	var convArg0 = prefix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings10beginGroupERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endGroup()
func (this *QSettings) EndGroup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings8endGroupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] QString group()
func (this *QSettings) Group() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] int beginReadArray(const QString &)
func (this *QSettings) BeginReadArray(prefix *QString) int {
	var convArg0 = prefix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings14beginReadArrayERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsettings.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginWriteArray(const QString &, int)
func (this *QSettings) BeginWriteArray(prefix *QString, size int) {
	var convArg0 = prefix.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings15beginWriteArrayERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endArray()
func (this *QSettings) EndArray() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings8endArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArrayIndex(int)
func (this *QSettings) SetArrayIndex(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings13setArrayIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable()
func (this *QSettings) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(const QString &, const QVariant &)
func (this *QSettings) SetValue(key *QString, value *QVariant) {
	var convArg0 = key.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings8setValueERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:167
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QString &, const QVariant &)
func (this *QSettings) Value(key *QString, defaultValue *QVariant) *QVariant /*123*/ {
	var convArg0 = key.GetCthis()
	var convArg1 = defaultValue.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings5valueERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(const QString &)
func (this *QSettings) Remove(key *QString) {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &)
func (this *QSettings) Contains(key *QString) bool {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFallbacksEnabled(_Bool)
func (this *QSettings) SetFallbacksEnabled(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings19setFallbacksEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fallbacksEnabled()
func (this *QSettings) FallbacksEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings16fallbacksEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName()
func (this *QSettings) FileName() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:176
// index:0
// Public Visibility=Default Availability=Available
// [4] QSettings::Format format()
func (this *QSettings) Format() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qsettings.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] QSettings::Scope scope()
func (this *QSettings) Scope() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings5scopeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtCore/qsettings.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QString organizationName()
func (this *QSettings) OrganizationName() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings16organizationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QString applicationName()
func (this *QSettings) ApplicationName() *QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings15applicationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIniCodec(QTextCodec *)
func (this *QSettings) SetIniCodec(codec *QTextCodec /*777 QTextCodec **/) {
	var convArg0 = codec.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:183
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setIniCodec(const char *)
func (this *QSettings) SetIniCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * iniCodec()
func (this *QSettings) IniCodec() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings8iniCodecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:187
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultFormat(enum QSettings::Format)
func (this *QSettings) SetDefaultFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings16setDefaultFormatENS_6FormatE", qtrt.FFI_TYPE_POINTER, format)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetDefaultFormat(format int) {
	var nilthis *QSettings
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtCore/qsettings.h:188
// index:0
// Public static Visibility=Default Availability=Available
// [4] QSettings::Format defaultFormat()
func (this *QSettings) DefaultFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings13defaultFormatEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QSettings_DefaultFormat() int {
	var nilthis *QSettings
	rv := nilthis.DefaultFormat()
	return rv
}

// /usr/include/qt/QtCore/qsettings.h:189
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSystemIniPath(const QString &)
func (this *QSettings) SetSystemIniPath(dir *QString) {
	var convArg0 = dir.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings16setSystemIniPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetSystemIniPath(dir *QString) {
	var nilthis *QSettings
	nilthis.SetSystemIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:190
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setUserIniPath(const QString &)
func (this *QSettings) SetUserIniPath(dir *QString) {
	var convArg0 = dir.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings14setUserIniPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetUserIniPath(dir *QString) {
	var nilthis *QSettings
	nilthis.SetUserIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:191
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPath(enum QSettings::Format, enum QSettings::Scope, const QString &)
func (this *QSettings) SetPath(format int, scope int, path *QString) {
	var convArg2 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings7setPathENS_6FormatENS_5ScopeERK7QString", qtrt.FFI_TYPE_POINTER, format, scope, convArg2)
	gopp.ErrPrint(err, rv)
}
func QSettings_SetPath(format int, scope int, path *QString) {
	var nilthis *QSettings
	nilthis.SetPath(format, scope, path)
}

// /usr/include/qt/QtCore/qsettings.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QSettings) Event(event *QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QSettings__Status = int

const QSettings__NoError QSettings__Status = 0
const QSettings__AccessError QSettings__Status = 1
const QSettings__FormatError QSettings__Status = 2

type QSettings__Format = int

const QSettings__NativeFormat QSettings__Format = 0
const QSettings__IniFormat QSettings__Format = 1
const QSettings__InvalidFormat QSettings__Format = 16
const QSettings__CustomFormat1 QSettings__Format = 17
const QSettings__CustomFormat2 QSettings__Format = 18
const QSettings__CustomFormat3 QSettings__Format = 19
const QSettings__CustomFormat4 QSettings__Format = 20
const QSettings__CustomFormat5 QSettings__Format = 21
const QSettings__CustomFormat6 QSettings__Format = 22
const QSettings__CustomFormat7 QSettings__Format = 23
const QSettings__CustomFormat8 QSettings__Format = 24
const QSettings__CustomFormat9 QSettings__Format = 25
const QSettings__CustomFormat10 QSettings__Format = 26
const QSettings__CustomFormat11 QSettings__Format = 27
const QSettings__CustomFormat12 QSettings__Format = 28
const QSettings__CustomFormat13 QSettings__Format = 29
const QSettings__CustomFormat14 QSettings__Format = 30
const QSettings__CustomFormat15 QSettings__Format = 31
const QSettings__CustomFormat16 QSettings__Format = 32

type QSettings__Scope = int

const QSettings__UserScope QSettings__Scope = 0
const QSettings__SystemScope QSettings__Scope = 1

//  body block end
