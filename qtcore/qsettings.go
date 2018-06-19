package qtcore

// /usr/include/qt/QtCore/qsettings.h
// #include <qsettings.h>
// #include <QtCore>

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

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QSettings) InheritEvent(f func(event *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
type QSettings struct {
	*QObject
}
type QSettings_ITF interface {
	QObject_ITF
	QSettings_PTR() *QSettings
}

func (ptr *QSettings) QSettings_PTR() *QSettings { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSettings) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsettings.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSettings(const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings(organization string, application string, parent QObject_ITF /*777 QObject **/) *QSettings {
	var tmpArg0 = NewQString_5(organization)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(application)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringS2_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSettings(const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings__(organization string) *QSettings {
	var tmpArg0 = NewQString_5(organization)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringS2_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSettings(const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings__1(organization string, application string) *QSettings {
	var tmpArg0 = NewQString_5(organization)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(application)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringS2_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:129
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QSettings::Scope, const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_1(scope int, organization string, application string, parent QObject_ITF /*777 QObject **/) *QSettings {
	var tmpArg1 = NewQString_5(organization)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(application)
	var convArg2 = tmpArg2.GetCthis()
	var convArg3 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg3 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_5ScopeERK7QStringS3_P7QObject", qtrt.FFI_TYPE_POINTER, scope, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:129
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QSettings::Scope, const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_1_(scope int, organization string) *QSettings {
	var tmpArg1 = NewQString_5(organization)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = NewQString()
	// arg: 3, QObject *=Pointer, QObject=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_5ScopeERK7QStringS3_P7QObject", qtrt.FFI_TYPE_POINTER, scope, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:129
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QSettings::Scope, const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_1_1(scope int, organization string, application string) *QSettings {
	var tmpArg1 = NewQString_5(organization)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(application)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, QObject *=Pointer, QObject=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_5ScopeERK7QStringS3_P7QObject", qtrt.FFI_TYPE_POINTER, scope, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:131
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QSettings::Format, QSettings::Scope, const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_2(format int, scope int, organization string, application string, parent QObject_ITF /*777 QObject **/) *QSettings {
	var tmpArg2 = NewQString_5(organization)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(application)
	var convArg3 = tmpArg3.GetCthis()
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg4 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_6FormatENS_5ScopeERK7QStringS4_P7QObject", qtrt.FFI_TYPE_POINTER, format, scope, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:131
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QSettings::Format, QSettings::Scope, const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_2_(format int, scope int, organization string) *QSettings {
	var tmpArg2 = NewQString_5(organization)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = NewQString()
	// arg: 4, QObject *=Pointer, QObject=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_6FormatENS_5ScopeERK7QStringS4_P7QObject", qtrt.FFI_TYPE_POINTER, format, scope, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:131
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QSettings::Format, QSettings::Scope, const QString &, const QString &, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_2_1(format int, scope int, organization string, application string) *QSettings {
	var tmpArg2 = NewQString_5(organization)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString_5(application)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, QObject *=Pointer, QObject=Record, , Invalid
	var convArg4 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ENS_6FormatENS_5ScopeERK7QStringS4_P7QObject", qtrt.FFI_TYPE_POINTER, format, scope, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:133
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QSettings(const QString &, QSettings::Format, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_3(fileName string, format int, parent QObject_ITF /*777 QObject **/) *QSettings {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringNS_6FormatEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, format, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:133
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QSettings(const QString &, QSettings::Format, QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_3_(fileName string, format int) *QSettings {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2ERK7QStringNS_6FormatEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, format, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:134
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_4(parent QObject_ITF /*777 QObject **/) *QSettings {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:134
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QSettings(QObject *)

/*
Constructs a QSettings object for accessing settings of the application called application from the organization called organization, and with parent parent.

Example:


  QSettings settings("Moose Tech", "Facturo-Pro");



The scope is set to QSettings::UserScope, and the format is set to QSettings::NativeFormat (i.e. calling setDefaultFormat() before calling this constructor has no effect).

See also setDefaultFormat() and Fallback Mechanism.
*/
func NewQSettings_4_() *QSettings {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSettingsFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QSettings")
	return gothis
}

// /usr/include/qt/QtCore/qsettings.h:144
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSettings()

/*

 */
func DeleteQSettings(this *QSettings) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettingsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qsettings.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all entries in the primary location associated to this QSettings object.

Entries in fallback locations are not removed.

If you only want to remove the entries in the current group(), use remove("") instead.

See also remove() and setFallbacksEnabled().
*/
func (this *QSettings) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sync()

/*
Writes any unsaved changes to permanent storage, and reloads any settings that have been changed in the meantime by another application.

This function is called automatically from QSettings's destructor and by the event loop at regular intervals, so you normally don't need to call it yourself.

See also status().
*/
func (this *QSettings) Sync() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings4syncEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:148
// index:0
// Public Visibility=Default Availability=Available
// [4] QSettings::Status status() const

/*
Returns a status code indicating the first error that was met by QSettings, or QSettings::NoError if no error occurred.

Be aware that QSettings delays performing some operations. For this reason, you might want to call sync() to ensure that the data stored in QSettings is written to disk before calling status().

See also sync().
*/
func (this *QSettings) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsettings.h:149
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAtomicSyncRequired() const

/*
Returns true if QSettings is only allowed to perform atomic saving and reloading (synchronization) of the settings. Returns false if it is allowed to save the settings contents directly to the configuration file.

The default is true.

This function was introduced in  Qt 5.10.

See also setAtomicSyncRequired() and QSaveFile.
*/
func (this *QSettings) IsAtomicSyncRequired() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings20isAtomicSyncRequiredEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAtomicSyncRequired(bool)

/*
Configures whether QSettings is required to perform atomic saving and reloading (synchronization) of the settings. If the enable argument is true (the default), sync() will only perform synchronization operations that are atomic. If this is not possible, sync() will fail and status() will be an error condition.

Setting this property to false will allow QSettings to write directly to the configuration file and ignore any errors trying to lock it against other processes trying to write at the same time. Because of the potential for corruption, this option should be used with care, but is required in certain conditions, like a QSettings::IniFormat configuration file that exists in an otherwise non-writeable directory or NTFS Alternate Data Streams.

See QSaveFile for more information on the feature.

This function was introduced in  Qt 5.10.

See also isAtomicSyncRequired() and QSaveFile.
*/
func (this *QSettings) SetAtomicSyncRequired(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings21setAtomicSyncRequiredEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginGroup(const QString &)

/*
Appends prefix to the current group.

The current group is automatically prepended to all keys specified to QSettings. In addition, query functions such as childGroups(), childKeys(), and allKeys() are based on the group. By default, no group is set.

Groups are useful to avoid typing in the same setting paths over and over. For example:


  settings.beginGroup("mainwindow");
  settings.setValue("size", win->size());
  settings.setValue("fullScreen", win->isFullScreen());
  settings.endGroup();

  settings.beginGroup("outputpanel");
  settings.setValue("visible", panel->isVisible());
  settings.endGroup();



This will set the value of three settings:


mainwindow/size
mainwindow/fullScreen
outputpanel/visible


Call endGroup() to reset the current group to what it was before the corresponding beginGroup() call. Groups can be nested.

See also endGroup() and group().
*/
func (this *QSettings) BeginGroup(prefix string) {
	var tmpArg0 = NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings10beginGroupERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endGroup()

/*
Resets the group to what it was before the corresponding beginGroup() call.

Example:


  settings.beginGroup("alpha");
  // settings.group() == "alpha"

  settings.beginGroup("beta");
  // settings.group() == "alpha/beta"

  settings.endGroup();
  // settings.group() == "alpha"

  settings.endGroup();
  // settings.group() == ""



See also beginGroup() and group().
*/
func (this *QSettings) EndGroup() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings8endGroupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] QString group() const

/*
Returns the current group.

See also beginGroup() and endGroup().
*/
func (this *QSettings) Group() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings5groupEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsettings.h:156
// index:0
// Public Visibility=Default Availability=Available
// [4] int beginReadArray(const QString &)

/*
Adds prefix to the current group and starts reading from an array. Returns the size of the array.

Example:


  struct Login {
      QString userName;
      QString password;
  };
  QList<Login> logins;
  ...

  QSettings settings;
  int size = settings.beginReadArray("logins");
  for (int i = 0; i < size; ++i) {
      settings.setArrayIndex(i);
      Login login;
      login.userName = settings.value("userName").toString();
      login.password = settings.value("password").toString();
      logins.append(login);
  }
  settings.endArray();



Use beginWriteArray() to write the array in the first place.

See also beginWriteArray(), endArray(), and setArrayIndex().
*/
func (this *QSettings) BeginReadArray(prefix string) int {
	var tmpArg0 = NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings14beginReadArrayERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qsettings.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginWriteArray(const QString &, int)

/*
Adds prefix to the current group and starts writing an array of size size. If size is -1 (the default), it is automatically determined based on the indexes of the entries written.

If you have many occurrences of a certain set of keys, you can use arrays to make your life easier. For example, let's suppose that you want to save a variable-length list of user names and passwords. You could then write:


  struct Login {
      QString userName;
      QString password;
  };
  QList<Login> logins;
  ...

  QSettings settings;
  settings.beginWriteArray("logins");
  for (int i = 0; i < logins.size(); ++i) {
      settings.setArrayIndex(i);
      settings.setValue("userName", list.at(i).userName);
      settings.setValue("password", list.at(i).password);
  }
  settings.endArray();



The generated keys will have the form


logins/size
logins/1/userName
logins/1/password
logins/2/userName
logins/2/password
logins/3/userName
logins/3/password
...


To read back an array, use beginReadArray().

See also beginReadArray(), endArray(), and setArrayIndex().
*/
func (this *QSettings) BeginWriteArray(prefix string, size int) {
	var tmpArg0 = NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings15beginWriteArrayERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beginWriteArray(const QString &, int)

/*
Adds prefix to the current group and starts writing an array of size size. If size is -1 (the default), it is automatically determined based on the indexes of the entries written.

If you have many occurrences of a certain set of keys, you can use arrays to make your life easier. For example, let's suppose that you want to save a variable-length list of user names and passwords. You could then write:


  struct Login {
      QString userName;
      QString password;
  };
  QList<Login> logins;
  ...

  QSettings settings;
  settings.beginWriteArray("logins");
  for (int i = 0; i < logins.size(); ++i) {
      settings.setArrayIndex(i);
      settings.setValue("userName", list.at(i).userName);
      settings.setValue("password", list.at(i).password);
  }
  settings.endArray();



The generated keys will have the form


logins/size
logins/1/userName
logins/1/password
logins/2/userName
logins/2/password
logins/3/userName
logins/3/password
...


To read back an array, use beginReadArray().

See also beginReadArray(), endArray(), and setArrayIndex().
*/
func (this *QSettings) BeginWriteArray__(prefix string) {
	var tmpArg0 = NewQString_5(prefix)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	size := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings15beginWriteArrayERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void endArray()

/*
Closes the array that was started using beginReadArray() or beginWriteArray().

See also beginReadArray() and beginWriteArray().
*/
func (this *QSettings) EndArray() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings8endArrayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setArrayIndex(int)

/*
Sets the current array index to i. Calls to functions such as setValue(), value(), remove(), and contains() will operate on the array entry at that index.

You must call beginReadArray() or beginWriteArray() before you can call this function.
*/
func (this *QSettings) SetArrayIndex(i int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings13setArrayIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:161
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList allKeys() const

/*
Returns a list of all keys, including subkeys, that can be read using the QSettings object.

Example:


  QSettings settings;
  settings.setValue("fridge/color", QColor(Qt::white));
  settings.setValue("fridge/size", QSize(32, 96));
  settings.setValue("sofa", true);
  settings.setValue("tv", false);

  QStringList keys = settings.allKeys();
  // keys: ["fridge/color", "fridge/size", "sofa", "tv"]



If a group is set using beginGroup(), only the keys in the group are returned, without the group prefix:


  settings.beginGroup("fridge");
  keys = settings.allKeys();
  // keys: ["color", "size"]



See also childGroups() and childKeys().
*/
func (this *QSettings) AllKeys() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings7allKeysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:162
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList childKeys() const

/*
Returns a list of all top-level keys that can be read using the QSettings object.

Example:


  QSettings settings;
  settings.setValue("fridge/color", QColor(Qt::white));
  settings.setValue("fridge/size", QSize(32, 96));
  settings.setValue("sofa", true);
  settings.setValue("tv", false);

  QStringList keys = settings.childKeys();
  // keys: ["sofa", "tv"]



If a group is set using beginGroup(), the top-level keys in that group are returned, without the group prefix:


  settings.beginGroup("fridge");
  keys = settings.childKeys();
  // keys: ["color", "size"]



You can navigate through the entire setting hierarchy using childKeys() and childGroups() recursively.

See also childGroups() and allKeys().
*/
func (this *QSettings) ChildKeys() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings9childKeysEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:163
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList childGroups() const

/*
Returns a list of all key top-level groups that contain keys that can be read using the QSettings object.

Example:


  QSettings settings;
  settings.setValue("fridge/color", QColor(Qt::white));
  settings.setValue("fridge/size", QSize(32, 96));
  settings.setValue("sofa", true);
  settings.setValue("tv", false);

  QStringList groups = settings.childGroups();
  // groups: ["fridge"]



If a group is set using beginGroup(), the first-level keys in that group are returned, without the group prefix.


  settings.beginGroup("fridge");
  groups = settings.childGroups();
  // groups: []



You can navigate through the entire setting hierarchy using childKeys() and childGroups() recursively.

See also childKeys() and allKeys().
*/
func (this *QSettings) ChildGroups() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings11childGroupsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWritable() const

/*
Returns true if settings can be written using this QSettings object; returns false otherwise.

One reason why isWritable() might return false is if QSettings operates on a read-only file.

Warning: This function is not perfectly reliable, because the file permissions can change at any time.

See also fileName(), status(), and sync().
*/
func (this *QSettings) IsWritable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings10isWritableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setValue(const QString &, const QVariant &)

/*
Sets the value of setting key to value. If the key already exists, the previous value is overwritten.

Note that the Windows registry and INI files use case-insensitive keys, whereas the CFPreferences API on macOS and iOS uses case-sensitive keys. To avoid portability problems, see the Section and Key Syntax rules.

Example:


  QSettings settings;
  settings.setValue("interval", 30);
  settings.value("interval").toInt();     // returns 30

  settings.setValue("interval", 6.55);
  settings.value("interval").toDouble();  // returns 6.55



See also value(), remove(), and contains().
*/
func (this *QSettings) SetValue(key string, value QVariant_ITF) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings8setValueERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:167
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QString &, const QVariant &) const

/*
Returns the value for setting key. If the setting doesn't exist, returns defaultValue.

If no default value is specified, a default QVariant is returned.

Note that the Windows registry and INI files use case-insensitive keys, whereas the CFPreferences API on macOS and iOS uses case-sensitive keys. To avoid portability problems, see the Section and Key Syntax rules.

Example:


  QSettings settings;
  settings.setValue("animal/snake", 58);
  settings.value("animal/snake", 1024).toInt();   // returns 58
  settings.value("animal/zebra", 1024).toInt();   // returns 1024
  settings.value("animal/zebra").toInt();         // returns 0



See also setValue(), contains(), and remove().
*/
func (this *QSettings) Value(key string, defaultValue QVariant_ITF) *QVariant /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if defaultValue != nil && defaultValue.QVariant_PTR() != nil {
		convArg1 = defaultValue.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings5valueERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:167
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant value(const QString &, const QVariant &) const

/*
Returns the value for setting key. If the setting doesn't exist, returns defaultValue.

If no default value is specified, a default QVariant is returned.

Note that the Windows registry and INI files use case-insensitive keys, whereas the CFPreferences API on macOS and iOS uses case-sensitive keys. To avoid portability problems, see the Section and Key Syntax rules.

Example:


  QSettings settings;
  settings.setValue("animal/snake", 58);
  settings.value("animal/snake", 1024).toInt();   // returns 58
  settings.value("animal/zebra", 1024).toInt();   // returns 1024
  settings.value("animal/zebra").toInt();         // returns 0



See also setValue(), contains(), and remove().
*/
func (this *QSettings) Value__(key string) *QVariant /*123*/ {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QVariant &=LValueReference, QVariant=Record, , Invalid
	var convArg1 = NewQVariant()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings5valueERK7QStringRK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qsettings.h:169
// index:0
// Public Visibility=Default Availability=Available
// [-2] void remove(const QString &)

/*
Removes the setting key and any sub-settings of key.

Example:


  QSettings settings;
  settings.setValue("ape");
  settings.setValue("monkey", 1);
  settings.setValue("monkey/sea", 2);
  settings.setValue("monkey/doe", 4);

  settings.remove("monkey");
  QStringList keys = settings.allKeys();
  // keys: ["ape"]



Be aware that if one of the fallback locations contains a setting with the same key, that setting will be visible after calling remove().

If key is an empty string, all keys in the current group() are removed. For example:


  QSettings settings;
  settings.setValue("ape");
  settings.setValue("monkey", 1);
  settings.setValue("monkey/sea", 2);
  settings.setValue("monkey/doe", 4);

  settings.beginGroup("monkey");
  settings.remove("");
  settings.endGroup();

  QStringList keys = settings.allKeys();
  // keys: ["ape"]



Note that the Windows registry and INI files use case-insensitive keys, whereas the CFPreferences API on macOS and iOS uses case-sensitive keys. To avoid portability problems, see the Section and Key Syntax rules.

See also setValue(), value(), and contains().
*/
func (this *QSettings) Remove(key string) {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings6removeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:170
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(const QString &) const

/*
Returns true if there exists a setting called key; returns false otherwise.

If a group is set using beginGroup(), key is taken to be relative to that group.

Note that the Windows registry and INI files use case-insensitive keys, whereas the CFPreferences API on macOS and iOS uses case-sensitive keys. To avoid portability problems, see the Section and Key Syntax rules.

See also value() and setValue().
*/
func (this *QSettings) Contains(key string) bool {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings8containsERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFallbacksEnabled(bool)

/*
Sets whether fallbacks are enabled to b.

By default, fallbacks are enabled.

See also fallbacksEnabled().
*/
func (this *QSettings) SetFallbacksEnabled(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings19setFallbacksEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:173
// index:0
// Public Visibility=Default Availability=Available
// [1] bool fallbacksEnabled() const

/*
Returns true if fallbacks are enabled; returns false otherwise.

By default, fallbacks are enabled.

See also setFallbacksEnabled().
*/
func (this *QSettings) FallbacksEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings16fallbacksEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qsettings.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*
Returns the path where settings written using this QSettings object are stored.

On Windows, if the format is QSettings::NativeFormat, the return value is a system registry path, not a file path.

See also isWritable() and format().
*/
func (this *QSettings) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsettings.h:176
// index:0
// Public Visibility=Default Availability=Available
// [4] QSettings::Format format() const

/*
Returns the format used for storing the settings.

This function was introduced in  Qt 4.4.

See also defaultFormat(), fileName(), scope(), organizationName(), and applicationName().
*/
func (this *QSettings) Format() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsettings.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] QSettings::Scope scope() const

/*
Returns the scope used for storing the settings.

This function was introduced in  Qt 4.4.

See also format(), organizationName(), and applicationName().
*/
func (this *QSettings) Scope() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings5scopeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qsettings.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QString organizationName() const

/*
Returns the organization name used for storing the settings.

This function was introduced in  Qt 4.4.

See also QCoreApplication::organizationName(), format(), scope(), and applicationName().
*/
func (this *QSettings) OrganizationName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings16organizationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsettings.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QString applicationName() const

/*
Returns the application name used for storing the settings.

This function was introduced in  Qt 4.4.

See also QCoreApplication::applicationName(), format(), scope(), and organizationName().
*/
func (this *QSettings) ApplicationName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings15applicationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qsettings.h:182
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIniCodec(QTextCodec *)

/*
Sets the codec for accessing INI files (including .conf files on Unix) to codec. The codec is used for decoding any data that is read from the INI file, and for encoding any data that is written to the file. By default, no codec is used, and non-ASCII characters are encoded using standard INI escape sequences.

Warning: The codec must be set immediately after creating the QSettings object, before accessing any data.

This function was introduced in  Qt 4.5.

See also iniCodec().
*/
func (this *QSettings) SetIniCodec(codec QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 unsafe.Pointer
	if codec != nil && codec.QTextCodec_PTR() != nil {
		convArg0 = codec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEP10QTextCodec", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:183
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setIniCodec(const char *)

/*
Sets the codec for accessing INI files (including .conf files on Unix) to codec. The codec is used for decoding any data that is read from the INI file, and for encoding any data that is written to the file. By default, no codec is used, and non-ASCII characters are encoded using standard INI escape sequences.

Warning: The codec must be set immediately after creating the QSettings object, before accessing any data.

This function was introduced in  Qt 4.5.

See also iniCodec().
*/
func (this *QSettings) SetIniCodec_1(codecName string) {
	var convArg0 = qtrt.CString(codecName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings11setIniCodecEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qsettings.h:184
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCodec * iniCodec() const

/*
Returns the codec that is used for accessing INI files. By default, no codec is used, so a null pointer is returned.

This function was introduced in  Qt 4.5.

See also setIniCodec().
*/
func (this *QSettings) IniCodec() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSettings8iniCodecEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qsettings.h:187
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultFormat(QSettings::Format)

/*
Sets the default file format to the given format, which is used for storing settings for the QSettings(QObject *) constructor.

If no default format is set, QSettings::NativeFormat is used. See the documentation for the QSettings constructor you are using to see if that constructor will ignore this function.

This function was introduced in  Qt 4.4.

See also defaultFormat() and format().
*/
func (this *QSettings) SetDefaultFormat(format int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings16setDefaultFormatENS_6FormatE", qtrt.FFI_TYPE_POINTER, format)
	qtrt.ErrPrint(err, rv)
}
func QSettings_SetDefaultFormat(format int) {
	var nilthis *QSettings
	nilthis.SetDefaultFormat(format)
}

// /usr/include/qt/QtCore/qsettings.h:188
// index:0
// Public static Visibility=Default Availability=Available
// [4] QSettings::Format defaultFormat()

/*
Returns default file format used for storing settings for the QSettings(QObject *) constructor. If no default format is set, QSettings::NativeFormat is used.

This function was introduced in  Qt 4.4.

See also setDefaultFormat() and format().
*/
func (this *QSettings) DefaultFormat() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings13defaultFormatEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
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

/*

 */
func (this *QSettings) SetSystemIniPath(dir string) {
	var tmpArg0 = NewQString_5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings16setSystemIniPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSettings_SetSystemIniPath(dir string) {
	var nilthis *QSettings
	nilthis.SetSystemIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:190
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setUserIniPath(const QString &)

/*

 */
func (this *QSettings) SetUserIniPath(dir string) {
	var tmpArg0 = NewQString_5(dir)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings14setUserIniPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QSettings_SetUserIniPath(dir string) {
	var nilthis *QSettings
	nilthis.SetUserIniPath(dir)
}

// /usr/include/qt/QtCore/qsettings.h:191
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setPath(QSettings::Format, QSettings::Scope, const QString &)

/*
Sets the path used for storing settings for the given format and scope, to path. The format can be a custom format.

The table below summarizes the default values:


 PlatformFormatScopePath
WindowsIniFormatUserScopeFOLDERID_RoamingAppData
SystemScopeFOLDERID_ProgramData
UnixNativeFormat, IniFormatUserScope$HOME/.config
SystemScope/etc/xdg
Qt for Embedded LinuxNativeFormat, IniFormatUserScope$HOME/Settings
SystemScope/etc/xdg
macOS and iOSIniFormatUserScope$HOME/.config
SystemScope/etc/xdg


The default UserScope paths on Unix, macOS, and iOS ($HOME/.config or $HOME/Settings) can be overridden by the user by setting the XDG_CONFIG_HOME environment variable. The default SystemScope paths on Unix, macOS, and iOS (/etc/xdg) can be overridden when building the Qt library using the configure script's -sysconfdir flag (see QLibraryInfo for details).

Setting the NativeFormat paths on Windows, macOS, and iOS has no effect.

Warning: This function doesn't affect existing QSettings objects.

This function was introduced in  Qt 4.1.

See also registerFormat().
*/
func (this *QSettings) SetPath(format int, scope int, path string) {
	var tmpArg2 = NewQString_5(path)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings7setPathENS_6FormatENS_5ScopeERK7QString", qtrt.FFI_TYPE_POINTER, format, scope, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QSettings_SetPath(format int, scope int, path string) {
	var nilthis *QSettings
	nilthis.SetPath(format, scope, path)
}

// /usr/include/qt/QtCore/qsettings.h:202
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QSettings) Event(event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSettings5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
The following status values are possible:



See also status().

*/
type QSettings__Status = int

// No error occurred.
const QSettings__NoError QSettings__Status = 0

// An access error occurred (e.g. trying to write to a read-only file).
const QSettings__AccessError QSettings__Status = 1

// A format error occurred (e.g. loading a malformed INI file).
const QSettings__FormatError QSettings__Status = 2

/*
This enum type specifies the storage format used by QSettings.



On Unix, NativeFormat and IniFormat mean the same thing, except that the file extension is different (.conf for NativeFormat, .ini for IniFormat).

The INI file format is a Windows file format that Qt supports on all platforms. In the absence of an INI standard, we try to follow what Microsoft does, with the following exceptions:


If you store types that QVariant can't convert to QString (e.g., QPoint, QRect, and QSize), Qt uses an @-based syntax to encode the type. For example:
  pos = @Point(100 100)


To minimize compatibility issues, any @ that doesn't appear at the first position in the value or that isn't followed by a Qt type (Point, Rect, Size, etc.) is treated as a normal character.

Although backslash is a special character in INI files, most Windows applications don't escape backslashes (\) in file paths:
  windir = C:\Windows


QSettings always treats backslash as a special character and provides no API for reading or writing such entries.

The INI file format has severe restrictions on the syntax of a key. Qt works around this by using % as an escape character in keys. In addition, if you save a top-level setting (a key with no slashes in it, e.g., "someKey"), it will appear in the INI file's "General" section. To avoid overwriting other keys, if you save something using a key such as "General/someKey", the key will be located in the "%General" section, not in the "General" section.
Following the philosophy that we should be liberal in what we accept and conservative in what we generate, QSettings will accept Latin-1 encoded INI files, but generate pure ASCII files, where non-ASCII values are encoded using standard INI escape sequences. To make the INI files more readable (but potentially less compatible), call setIniCodec().


See also registerFormat() and setPath().

*/
type QSettings__Format = int

// Store the settings using the most appropriate storage format for the platform. On Windows, this means the system registry; on macOS and iOS, this means the CFPreferences API; on Unix, this means textual configuration files in INI format.
const QSettings__NativeFormat QSettings__Format = 0

// Store the settings in INI files.
const QSettings__IniFormat QSettings__Format = 1

//
const QSettings__InvalidFormat QSettings__Format = 16

//
const QSettings__CustomFormat1 QSettings__Format = 17

//
const QSettings__CustomFormat2 QSettings__Format = 18

//
const QSettings__CustomFormat3 QSettings__Format = 19

//
const QSettings__CustomFormat4 QSettings__Format = 20

//
const QSettings__CustomFormat5 QSettings__Format = 21

//
const QSettings__CustomFormat6 QSettings__Format = 22

//
const QSettings__CustomFormat7 QSettings__Format = 23

//
const QSettings__CustomFormat8 QSettings__Format = 24

//
const QSettings__CustomFormat9 QSettings__Format = 25

//
const QSettings__CustomFormat10 QSettings__Format = 26

//
const QSettings__CustomFormat11 QSettings__Format = 27

//
const QSettings__CustomFormat12 QSettings__Format = 28

//
const QSettings__CustomFormat13 QSettings__Format = 29

//
const QSettings__CustomFormat14 QSettings__Format = 30

//
const QSettings__CustomFormat15 QSettings__Format = 31

//
const QSettings__CustomFormat16 QSettings__Format = 32

/*
This enum specifies whether settings are user-specific or shared by all users of the same system.



See also setPath().

*/
type QSettings__Scope = int

// Store settings in a location specific to the current user (e.g., in the user's home directory).
const QSettings__UserScope QSettings__Scope = 0

// Store settings in a global location, so that all users on the same machine access the same set of settings.
const QSettings__SystemScope QSettings__Scope = 1

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
}

//  keep block end
