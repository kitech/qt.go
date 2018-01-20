//  header block begin
// /usr/include/qt/QtCore/qcoreapplication.h
// #include <qcoreapplication.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
type QCoreApplication struct {
	*QObject
}

func (this *QCoreApplication) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQCoreApplicationFromPointer(cthis unsafe.Pointer) *QCoreApplication {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QCoreApplication{bcthis0}
}

// /usr/include/qt/QtCore/qcoreapplication.h:78
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QCoreApplication) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QCoreApplication10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:91
// index:0
// Public
// void QCoreApplication(int &, char **, int)
func NewQCoreApplication(argc int, argv []string, arg2 int) *QCoreApplication {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplicationC2ERiPPci", ffiqt.FFI_TYPE_VOID, cthis, &argc, convArg1, &arg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQCoreApplicationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcoreapplication.h:97
// index:0
// Public virtual
// void ~QCoreApplication()
func DeleteQCoreApplication(*QCoreApplication) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplicationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:99
// index:0
// Public static
// QStringList arguments()
func (this *QCoreApplication) Arguments() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9argumentsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_Arguments() {
	var nilthis *QCoreApplication
	nilthis.Arguments()
}

// /usr/include/qt/QtCore/qcoreapplication.h:101
// index:0
// Public static
// void setAttribute(Qt::ApplicationAttribute, _Bool)
func (this *QCoreApplication) SetAttribute(attribute int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication12setAttributeEN2Qt20ApplicationAttributeEb", ffiqt.FFI_TYPE_POINTER, attribute, on)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetAttribute(attribute int, on bool) {
	var nilthis *QCoreApplication
	nilthis.SetAttribute(attribute, on)
}

// /usr/include/qt/QtCore/qcoreapplication.h:102
// index:0
// Public static
// bool testAttribute(Qt::ApplicationAttribute)
func (this *QCoreApplication) TestAttribute(attribute int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13testAttributeEN2Qt20ApplicationAttributeE", ffiqt.FFI_TYPE_POINTER, attribute)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_TestAttribute(attribute int) {
	var nilthis *QCoreApplication
	nilthis.TestAttribute(attribute)
}

// /usr/include/qt/QtCore/qcoreapplication.h:104
// index:0
// Public static
// void setOrganizationDomain(const class QString &)
func (this *QCoreApplication) SetOrganizationDomain(orgDomain *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication21setOrganizationDomainERK7QString", ffiqt.FFI_TYPE_POINTER, orgDomain)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetOrganizationDomain(orgDomain *QString) {
	var nilthis *QCoreApplication
	nilthis.SetOrganizationDomain(orgDomain)
}

// /usr/include/qt/QtCore/qcoreapplication.h:105
// index:0
// Public static
// QString organizationDomain()
func (this *QCoreApplication) OrganizationDomain() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18organizationDomainEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_OrganizationDomain() {
	var nilthis *QCoreApplication
	nilthis.OrganizationDomain()
}

// /usr/include/qt/QtCore/qcoreapplication.h:106
// index:0
// Public static
// void setOrganizationName(const class QString &)
func (this *QCoreApplication) SetOrganizationName(orgName *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication19setOrganizationNameERK7QString", ffiqt.FFI_TYPE_POINTER, orgName)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetOrganizationName(orgName *QString) {
	var nilthis *QCoreApplication
	nilthis.SetOrganizationName(orgName)
}

// /usr/include/qt/QtCore/qcoreapplication.h:107
// index:0
// Public static
// QString organizationName()
func (this *QCoreApplication) OrganizationName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16organizationNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_OrganizationName() {
	var nilthis *QCoreApplication
	nilthis.OrganizationName()
}

// /usr/include/qt/QtCore/qcoreapplication.h:108
// index:0
// Public static
// void setApplicationName(const class QString &)
func (this *QCoreApplication) SetApplicationName(application *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setApplicationNameERK7QString", ffiqt.FFI_TYPE_POINTER, application)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetApplicationName(application *QString) {
	var nilthis *QCoreApplication
	nilthis.SetApplicationName(application)
}

// /usr/include/qt/QtCore/qcoreapplication.h:109
// index:0
// Public static
// QString applicationName()
func (this *QCoreApplication) ApplicationName() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15applicationNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_ApplicationName() {
	var nilthis *QCoreApplication
	nilthis.ApplicationName()
}

// /usr/include/qt/QtCore/qcoreapplication.h:110
// index:0
// Public static
// void setApplicationVersion(const class QString &)
func (this *QCoreApplication) SetApplicationVersion(version *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication21setApplicationVersionERK7QString", ffiqt.FFI_TYPE_POINTER, version)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetApplicationVersion(version *QString) {
	var nilthis *QCoreApplication
	nilthis.SetApplicationVersion(version)
}

// /usr/include/qt/QtCore/qcoreapplication.h:111
// index:0
// Public static
// QString applicationVersion()
func (this *QCoreApplication) ApplicationVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18applicationVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_ApplicationVersion() {
	var nilthis *QCoreApplication
	nilthis.ApplicationVersion()
}

// /usr/include/qt/QtCore/qcoreapplication.h:113
// index:0
// Public static
// void setSetuidAllowed(_Bool)
func (this *QCoreApplication) SetSetuidAllowed(allow bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16setSetuidAllowedEb", ffiqt.FFI_TYPE_POINTER, allow)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetSetuidAllowed(allow bool) {
	var nilthis *QCoreApplication
	nilthis.SetSetuidAllowed(allow)
}

// /usr/include/qt/QtCore/qcoreapplication.h:114
// index:0
// Public static
// bool isSetuidAllowed()
func (this *QCoreApplication) IsSetuidAllowed() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15isSetuidAllowedEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_IsSetuidAllowed() {
	var nilthis *QCoreApplication
	nilthis.IsSetuidAllowed()
}

// /usr/include/qt/QtCore/qcoreapplication.h:116
// index:0
// Public static inline
// QCoreApplication * instance()
func (this *QCoreApplication) Instance() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication8instanceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_Instance() {
	var nilthis *QCoreApplication
	nilthis.Instance()
}

// /usr/include/qt/QtCore/qcoreapplication.h:119
// index:0
// Public static
// int exec()
func (this *QCoreApplication) Exec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication4execEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_Exec() {
	var nilthis *QCoreApplication
	nilthis.Exec()
}

// /usr/include/qt/QtCore/qcoreapplication.h:122
// index:0
// Public static
// void exit(int)
func (this *QCoreApplication) Exit(retcode int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication4exitEi", ffiqt.FFI_TYPE_POINTER, retcode)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Exit(retcode int) {
	var nilthis *QCoreApplication
	nilthis.Exit(retcode)
}

// /usr/include/qt/QtCore/qcoreapplication.h:124
// index:0
// Public static
// bool sendEvent(class QObject *, class QEvent *)
func (this *QCoreApplication) SendEvent(receiver unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, receiver, event)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_SendEvent(receiver unsafe.Pointer, event unsafe.Pointer) {
	var nilthis *QCoreApplication
	nilthis.SendEvent(receiver, event)
}

// /usr/include/qt/QtCore/qcoreapplication.h:125
// index:0
// Public static
// void postEvent(class QObject *, class QEvent *, int)
func (this *QCoreApplication) PostEvent(receiver unsafe.Pointer, event unsafe.Pointer, priority int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9postEventEP7QObjectP6QEventi", ffiqt.FFI_TYPE_POINTER, receiver, event, priority)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_PostEvent(receiver unsafe.Pointer, event unsafe.Pointer, priority int) {
	var nilthis *QCoreApplication
	nilthis.PostEvent(receiver, event, priority)
}

// /usr/include/qt/QtCore/qcoreapplication.h:126
// index:0
// Public static
// void sendPostedEvents(class QObject *, int)
func (this *QCoreApplication) SendPostedEvents(receiver unsafe.Pointer, event_type int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", ffiqt.FFI_TYPE_POINTER, receiver, event_type)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SendPostedEvents(receiver unsafe.Pointer, event_type int) {
	var nilthis *QCoreApplication
	nilthis.SendPostedEvents(receiver, event_type)
}

// /usr/include/qt/QtCore/qcoreapplication.h:127
// index:0
// Public static
// void removePostedEvents(class QObject *, int)
func (this *QCoreApplication) RemovePostedEvents(receiver unsafe.Pointer, eventType int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18removePostedEventsEP7QObjecti", ffiqt.FFI_TYPE_POINTER, receiver, eventType)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_RemovePostedEvents(receiver unsafe.Pointer, eventType int) {
	var nilthis *QCoreApplication
	nilthis.RemovePostedEvents(receiver, eventType)
}

// /usr/include/qt/QtCore/qcoreapplication.h:129
// index:0
// Public static
// bool hasPendingEvents()
func (this *QCoreApplication) HasPendingEvents() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16hasPendingEventsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_HasPendingEvents() {
	var nilthis *QCoreApplication
	nilthis.HasPendingEvents()
}

// /usr/include/qt/QtCore/qcoreapplication.h:131
// index:0
// Public static
// QAbstractEventDispatcher * eventDispatcher()
func (this *QCoreApplication) EventDispatcher() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15eventDispatcherEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_EventDispatcher() {
	var nilthis *QCoreApplication
	nilthis.EventDispatcher()
}

// /usr/include/qt/QtCore/qcoreapplication.h:132
// index:0
// Public static
// void setEventDispatcher(class QAbstractEventDispatcher *)
func (this *QCoreApplication) SetEventDispatcher(eventDispatcher unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher", ffiqt.FFI_TYPE_POINTER, eventDispatcher)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetEventDispatcher(eventDispatcher unsafe.Pointer) {
	var nilthis *QCoreApplication
	nilthis.SetEventDispatcher(eventDispatcher)
}

// /usr/include/qt/QtCore/qcoreapplication.h:134
// index:0
// Public virtual
// bool notify(class QObject *, class QEvent *)
func (this *QCoreApplication) Notify(arg0 unsafe.Pointer, arg1 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:136
// index:0
// Public static
// bool startingUp()
func (this *QCoreApplication) StartingUp() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication10startingUpEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_StartingUp() {
	var nilthis *QCoreApplication
	nilthis.StartingUp()
}

// /usr/include/qt/QtCore/qcoreapplication.h:137
// index:0
// Public static
// bool closingDown()
func (this *QCoreApplication) ClosingDown() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication11closingDownEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_ClosingDown() {
	var nilthis *QCoreApplication
	nilthis.ClosingDown()
}

// /usr/include/qt/QtCore/qcoreapplication.h:140
// index:0
// Public static
// QString applicationDirPath()
func (this *QCoreApplication) ApplicationDirPath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18applicationDirPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_ApplicationDirPath() {
	var nilthis *QCoreApplication
	nilthis.ApplicationDirPath()
}

// /usr/include/qt/QtCore/qcoreapplication.h:141
// index:0
// Public static
// QString applicationFilePath()
func (this *QCoreApplication) ApplicationFilePath() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication19applicationFilePathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_ApplicationFilePath() {
	var nilthis *QCoreApplication
	nilthis.ApplicationFilePath()
}

// /usr/include/qt/QtCore/qcoreapplication.h:142
// index:0
// Public static
// qint64 applicationPid()
func (this *QCoreApplication) ApplicationPid() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication14applicationPidEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_ApplicationPid() {
	var nilthis *QCoreApplication
	nilthis.ApplicationPid()
}

// /usr/include/qt/QtCore/qcoreapplication.h:145
// index:0
// Public static
// void setLibraryPaths(const class QStringList &)
func (this *QCoreApplication) SetLibraryPaths(arg0 *QStringList) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15setLibraryPathsERK11QStringList", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetLibraryPaths(arg0 *QStringList) {
	var nilthis *QCoreApplication
	nilthis.SetLibraryPaths(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:146
// index:0
// Public static
// QStringList libraryPaths()
func (this *QCoreApplication) LibraryPaths() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication12libraryPathsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_LibraryPaths() {
	var nilthis *QCoreApplication
	nilthis.LibraryPaths()
}

// /usr/include/qt/QtCore/qcoreapplication.h:147
// index:0
// Public static
// void addLibraryPath(const class QString &)
func (this *QCoreApplication) AddLibraryPath(arg0 *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication14addLibraryPathERK7QString", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_AddLibraryPath(arg0 *QString) {
	var nilthis *QCoreApplication
	nilthis.AddLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:148
// index:0
// Public static
// void removeLibraryPath(const class QString &)
func (this *QCoreApplication) RemoveLibraryPath(arg0 *QString) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17removeLibraryPathERK7QString", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_RemoveLibraryPath(arg0 *QString) {
	var nilthis *QCoreApplication
	nilthis.RemoveLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:152
// index:0
// Public static
// bool installTranslator(class QTranslator *)
func (this *QCoreApplication) InstallTranslator(messageFile unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17installTranslatorEP11QTranslator", ffiqt.FFI_TYPE_POINTER, messageFile)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_InstallTranslator(messageFile unsafe.Pointer) {
	var nilthis *QCoreApplication
	nilthis.InstallTranslator(messageFile)
}

// /usr/include/qt/QtCore/qcoreapplication.h:153
// index:0
// Public static
// bool removeTranslator(class QTranslator *)
func (this *QCoreApplication) RemoveTranslator(messageFile unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16removeTranslatorEP11QTranslator", ffiqt.FFI_TYPE_POINTER, messageFile)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_RemoveTranslator(messageFile unsafe.Pointer) {
	var nilthis *QCoreApplication
	nilthis.RemoveTranslator(messageFile)
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static
// QString translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) Translate(context string, key string, disambiguation string, n int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9translateEPKcS1_S1_i", ffiqt.FFI_TYPE_POINTER, context, key, disambiguation, n)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_Translate(context string, key string, disambiguation string, n int) {
	var nilthis *QCoreApplication
	nilthis.Translate(context, key, disambiguation, n)
}

// /usr/include/qt/QtCore/qcoreapplication.h:169
// index:0
// Public static
// void flush()
func (this *QCoreApplication) Flush() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication5flushEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Flush() {
	var nilthis *QCoreApplication
	nilthis.Flush()
}

// /usr/include/qt/QtCore/qcoreapplication.h:172
// index:0
// Public
// void installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) InstallNativeEventFilter(filterObj unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filterObj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:173
// index:0
// Public
// void removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) RemoveNativeEventFilter(filterObj unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), filterObj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:175
// index:0
// Public static
// bool isQuitLockEnabled()
func (this *QCoreApplication) IsQuitLockEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17isQuitLockEnabledEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QCoreApplication_IsQuitLockEnabled() {
	var nilthis *QCoreApplication
	nilthis.IsQuitLockEnabled()
}

// /usr/include/qt/QtCore/qcoreapplication.h:176
// index:0
// Public static
// void setQuitLockEnabled(_Bool)
func (this *QCoreApplication) SetQuitLockEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setQuitLockEnabledEb", ffiqt.FFI_TYPE_POINTER, enabled)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	var nilthis *QCoreApplication
	nilthis.SetQuitLockEnabled(enabled)
}

// /usr/include/qt/QtCore/qcoreapplication.h:179
// index:0
// Public static
// void quit()
func (this *QCoreApplication) Quit() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication4quitEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Quit() {
	var nilthis *QCoreApplication
	nilthis.Quit()
}

// /usr/include/qt/QtCore/qcoreapplication.h:184
// index:0
// Public
// void organizationNameChanged()
func (this *QCoreApplication) OrganizationNameChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication23organizationNameChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:185
// index:0
// Public
// void organizationDomainChanged()
func (this *QCoreApplication) OrganizationDomainChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication25organizationDomainChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:186
// index:0
// Public
// void applicationNameChanged()
func (this *QCoreApplication) ApplicationNameChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication22applicationNameChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:187
// index:0
// Public
// void applicationVersionChanged()
func (this *QCoreApplication) ApplicationVersionChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication25applicationVersionChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:190
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QCoreApplication) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:192
// index:0
// Protected virtual
// bool compressEvent(class QEvent *, class QObject *, class QPostEventList *)
func (this *QCoreApplication) CompressEvent(arg0 unsafe.Pointer, receiver unsafe.Pointer, arg2 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13compressEventEP6QEventP7QObjectP14QPostEventList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, receiver, arg2)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
