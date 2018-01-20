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

// /usr/include/qt/QtCore/qcoreapplication.h:78
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCoreApplication) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QCoreApplication10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:91
// index:0
// void QCoreApplication(int &, char **, int)
func NewQCoreApplication(argc int, argv []string, arg2 int) *QCoreApplication {
	cthis := qtrt.Calloc(1, 256)
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplicationC2ERiPPci", ffiqt.FFI_TYPE_VOID, cthis, &argc, convArg1, &arg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQCoreApplicationFromPointer(cthis)
	return gothis
}
func NewQCoreApplicationFromPointer(cthis unsafe.Pointer) *QCoreApplication {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QCoreApplication{bcthis0}
}

// /usr/include/qt/QtCore/qcoreapplication.h:196
// index:1
// void QCoreApplication(class QCoreApplicationPrivate &)
func NewQCoreApplication_1(p unsafe.Pointer) *QCoreApplication {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplicationC2ER23QCoreApplicationPrivate", ffiqt.FFI_TYPE_VOID, cthis, p)
	gopp.ErrPrint(err, rv)
	gothis := NewQCoreApplicationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qcoreapplication.h:97
// index:0
// virtual
// void ~QCoreApplication()
func DeleteQCoreApplication(*QCoreApplication) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplicationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:99
// index:0
// static
// QStringList arguments()
func (this *QCoreApplication) Arguments() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9argumentsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Arguments() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.Arguments()
}

// /usr/include/qt/QtCore/qcoreapplication.h:101
// index:0
// static
// void setAttribute(Qt::ApplicationAttribute, _Bool)
func (this *QCoreApplication) SetAttribute(attribute int, on bool) {
	// 0: (attribute Qt::ApplicationAttribute, on bool), (attribute, on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication12setAttributeEN2Qt20ApplicationAttributeEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetAttribute(attribute int, on bool) {
	// 0: (attribute Qt::ApplicationAttribute, on bool), (attribute, on)
	var nilthis *QCoreApplication
	nilthis.SetAttribute(attribute, on)
}

// /usr/include/qt/QtCore/qcoreapplication.h:102
// index:0
// static
// bool testAttribute(Qt::ApplicationAttribute)
func (this *QCoreApplication) TestAttribute(attribute int) {
	// 0: (attribute Qt::ApplicationAttribute), (attribute)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13testAttributeEN2Qt20ApplicationAttributeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_TestAttribute(attribute int) {
	// 0: (attribute Qt::ApplicationAttribute), (attribute)
	var nilthis *QCoreApplication
	nilthis.TestAttribute(attribute)
}

// /usr/include/qt/QtCore/qcoreapplication.h:104
// index:0
// static
// void setOrganizationDomain(const class QString &)
func (this *QCoreApplication) SetOrganizationDomain(orgDomain unsafe.Pointer) {
	// 0: (orgDomain const QString &), (orgDomain)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication21setOrganizationDomainERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetOrganizationDomain(orgDomain unsafe.Pointer) {
	// 0: (orgDomain const QString &), (orgDomain)
	var nilthis *QCoreApplication
	nilthis.SetOrganizationDomain(orgDomain)
}

// /usr/include/qt/QtCore/qcoreapplication.h:105
// index:0
// static
// QString organizationDomain()
func (this *QCoreApplication) OrganizationDomain() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18organizationDomainEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_OrganizationDomain() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.OrganizationDomain()
}

// /usr/include/qt/QtCore/qcoreapplication.h:106
// index:0
// static
// void setOrganizationName(const class QString &)
func (this *QCoreApplication) SetOrganizationName(orgName unsafe.Pointer) {
	// 0: (orgName const QString &), (orgName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication19setOrganizationNameERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetOrganizationName(orgName unsafe.Pointer) {
	// 0: (orgName const QString &), (orgName)
	var nilthis *QCoreApplication
	nilthis.SetOrganizationName(orgName)
}

// /usr/include/qt/QtCore/qcoreapplication.h:107
// index:0
// static
// QString organizationName()
func (this *QCoreApplication) OrganizationName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16organizationNameEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_OrganizationName() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.OrganizationName()
}

// /usr/include/qt/QtCore/qcoreapplication.h:108
// index:0
// static
// void setApplicationName(const class QString &)
func (this *QCoreApplication) SetApplicationName(application unsafe.Pointer) {
	// 0: (application const QString &), (application)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setApplicationNameERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetApplicationName(application unsafe.Pointer) {
	// 0: (application const QString &), (application)
	var nilthis *QCoreApplication
	nilthis.SetApplicationName(application)
}

// /usr/include/qt/QtCore/qcoreapplication.h:109
// index:0
// static
// QString applicationName()
func (this *QCoreApplication) ApplicationName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15applicationNameEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ApplicationName() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.ApplicationName()
}

// /usr/include/qt/QtCore/qcoreapplication.h:110
// index:0
// static
// void setApplicationVersion(const class QString &)
func (this *QCoreApplication) SetApplicationVersion(version unsafe.Pointer) {
	// 0: (version const QString &), (version)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication21setApplicationVersionERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetApplicationVersion(version unsafe.Pointer) {
	// 0: (version const QString &), (version)
	var nilthis *QCoreApplication
	nilthis.SetApplicationVersion(version)
}

// /usr/include/qt/QtCore/qcoreapplication.h:111
// index:0
// static
// QString applicationVersion()
func (this *QCoreApplication) ApplicationVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18applicationVersionEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ApplicationVersion() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.ApplicationVersion()
}

// /usr/include/qt/QtCore/qcoreapplication.h:113
// index:0
// static
// void setSetuidAllowed(_Bool)
func (this *QCoreApplication) SetSetuidAllowed(allow bool) {
	// 0: (allow bool), (allow)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16setSetuidAllowedEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetSetuidAllowed(allow bool) {
	// 0: (allow bool), (allow)
	var nilthis *QCoreApplication
	nilthis.SetSetuidAllowed(allow)
}

// /usr/include/qt/QtCore/qcoreapplication.h:114
// index:0
// static
// bool isSetuidAllowed()
func (this *QCoreApplication) IsSetuidAllowed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15isSetuidAllowedEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_IsSetuidAllowed() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.IsSetuidAllowed()
}

// /usr/include/qt/QtCore/qcoreapplication.h:116
// index:0
// static inline
// QCoreApplication * instance()
func (this *QCoreApplication) Instance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication8instanceEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Instance() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.Instance()
}

// /usr/include/qt/QtCore/qcoreapplication.h:119
// index:0
// static
// int exec()
func (this *QCoreApplication) Exec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication4execEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Exec() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.Exec()
}

// /usr/include/qt/QtCore/qcoreapplication.h:120
// index:0
// static
// void processEvents(class QEventLoop::ProcessEventsFlags)
func (this *QCoreApplication) ProcessEvents(flags int) {
	// 0: (QFlags<QEventLoop::ProcessEventsFlag> flags), (flags)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ProcessEvents(flags int) {
	// 0: (QFlags<QEventLoop::ProcessEventsFlag> flags), (flags)
	var nilthis *QCoreApplication
	nilthis.ProcessEvents(flags)
}

// /usr/include/qt/QtCore/qcoreapplication.h:121
// index:1
// static
// void processEvents(class QEventLoop::ProcessEventsFlags, int)
func (this *QCoreApplication) ProcessEvents_1(flags int, maxtime int) {
	// 1: (QFlags<QEventLoop::ProcessEventsFlag> flags, maxtime int), (flags, maxtime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ProcessEvents_1(flags int, maxtime int) {
	// 1: (QFlags<QEventLoop::ProcessEventsFlag> flags, maxtime int), (flags, maxtime)
	var nilthis *QCoreApplication
	nilthis.ProcessEvents_1(flags, maxtime)
}

// /usr/include/qt/QtCore/qcoreapplication.h:122
// index:0
// static
// void exit(int)
func (this *QCoreApplication) Exit(retcode int) {
	// 0: (retcode int), (retcode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication4exitEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Exit(retcode int) {
	// 0: (retcode int), (retcode)
	var nilthis *QCoreApplication
	nilthis.Exit(retcode)
}

// /usr/include/qt/QtCore/qcoreapplication.h:124
// index:0
// static
// bool sendEvent(class QObject *, class QEvent *)
func (this *QCoreApplication) SendEvent(receiver unsafe.Pointer, event unsafe.Pointer) {
	// 0: (receiver QObject *, event QEvent *), (receiver, event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SendEvent(receiver unsafe.Pointer, event unsafe.Pointer) {
	// 0: (receiver QObject *, event QEvent *), (receiver, event)
	var nilthis *QCoreApplication
	nilthis.SendEvent(receiver, event)
}

// /usr/include/qt/QtCore/qcoreapplication.h:125
// index:0
// static
// void postEvent(class QObject *, class QEvent *, int)
func (this *QCoreApplication) PostEvent(receiver unsafe.Pointer, event unsafe.Pointer, priority int) {
	// 0: (receiver QObject *, event QEvent *, priority int), (receiver, event, priority)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9postEventEP7QObjectP6QEventi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_PostEvent(receiver unsafe.Pointer, event unsafe.Pointer, priority int) {
	// 0: (receiver QObject *, event QEvent *, priority int), (receiver, event, priority)
	var nilthis *QCoreApplication
	nilthis.PostEvent(receiver, event, priority)
}

// /usr/include/qt/QtCore/qcoreapplication.h:126
// index:0
// static
// void sendPostedEvents(class QObject *, int)
func (this *QCoreApplication) SendPostedEvents(receiver unsafe.Pointer, event_type int) {
	// 0: (receiver QObject *, event_type int), (receiver, event_type)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SendPostedEvents(receiver unsafe.Pointer, event_type int) {
	// 0: (receiver QObject *, event_type int), (receiver, event_type)
	var nilthis *QCoreApplication
	nilthis.SendPostedEvents(receiver, event_type)
}

// /usr/include/qt/QtCore/qcoreapplication.h:127
// index:0
// static
// void removePostedEvents(class QObject *, int)
func (this *QCoreApplication) RemovePostedEvents(receiver unsafe.Pointer, eventType int) {
	// 0: (receiver QObject *, eventType int), (receiver, eventType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18removePostedEventsEP7QObjecti", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_RemovePostedEvents(receiver unsafe.Pointer, eventType int) {
	// 0: (receiver QObject *, eventType int), (receiver, eventType)
	var nilthis *QCoreApplication
	nilthis.RemovePostedEvents(receiver, eventType)
}

// /usr/include/qt/QtCore/qcoreapplication.h:129
// index:0
// static
// bool hasPendingEvents()
func (this *QCoreApplication) HasPendingEvents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16hasPendingEventsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_HasPendingEvents() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.HasPendingEvents()
}

// /usr/include/qt/QtCore/qcoreapplication.h:131
// index:0
// static
// QAbstractEventDispatcher * eventDispatcher()
func (this *QCoreApplication) EventDispatcher() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15eventDispatcherEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_EventDispatcher() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.EventDispatcher()
}

// /usr/include/qt/QtCore/qcoreapplication.h:132
// index:0
// static
// void setEventDispatcher(class QAbstractEventDispatcher *)
func (this *QCoreApplication) SetEventDispatcher(eventDispatcher unsafe.Pointer) {
	// 0: (eventDispatcher QAbstractEventDispatcher *), (eventDispatcher)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetEventDispatcher(eventDispatcher unsafe.Pointer) {
	// 0: (eventDispatcher QAbstractEventDispatcher *), (eventDispatcher)
	var nilthis *QCoreApplication
	nilthis.SetEventDispatcher(eventDispatcher)
}

// /usr/include/qt/QtCore/qcoreapplication.h:134
// index:0
// virtual
// bool notify(class QObject *, class QEvent *)
func (this *QCoreApplication) Notify(arg0 unsafe.Pointer, arg1 unsafe.Pointer) {
	// 0: (, QObject * arg0, QEvent * arg1), (arg0, arg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:136
// index:0
// static
// bool startingUp()
func (this *QCoreApplication) StartingUp() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication10startingUpEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_StartingUp() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.StartingUp()
}

// /usr/include/qt/QtCore/qcoreapplication.h:137
// index:0
// static
// bool closingDown()
func (this *QCoreApplication) ClosingDown() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication11closingDownEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ClosingDown() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.ClosingDown()
}

// /usr/include/qt/QtCore/qcoreapplication.h:140
// index:0
// static
// QString applicationDirPath()
func (this *QCoreApplication) ApplicationDirPath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18applicationDirPathEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ApplicationDirPath() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.ApplicationDirPath()
}

// /usr/include/qt/QtCore/qcoreapplication.h:141
// index:0
// static
// QString applicationFilePath()
func (this *QCoreApplication) ApplicationFilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication19applicationFilePathEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ApplicationFilePath() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.ApplicationFilePath()
}

// /usr/include/qt/QtCore/qcoreapplication.h:142
// index:0
// static
// qint64 applicationPid()
func (this *QCoreApplication) ApplicationPid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication14applicationPidEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ApplicationPid() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.ApplicationPid()
}

// /usr/include/qt/QtCore/qcoreapplication.h:145
// index:0
// static
// void setLibraryPaths(const class QStringList &)
func (this *QCoreApplication) SetLibraryPaths(arg0 unsafe.Pointer) {
	// 0: (const QStringList & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15setLibraryPathsERK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetLibraryPaths(arg0 unsafe.Pointer) {
	// 0: (const QStringList & arg0), (arg0)
	var nilthis *QCoreApplication
	nilthis.SetLibraryPaths(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:146
// index:0
// static
// QStringList libraryPaths()
func (this *QCoreApplication) LibraryPaths() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication12libraryPathsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_LibraryPaths() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.LibraryPaths()
}

// /usr/include/qt/QtCore/qcoreapplication.h:147
// index:0
// static
// void addLibraryPath(const class QString &)
func (this *QCoreApplication) AddLibraryPath(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication14addLibraryPathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_AddLibraryPath(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	var nilthis *QCoreApplication
	nilthis.AddLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:148
// index:0
// static
// void removeLibraryPath(const class QString &)
func (this *QCoreApplication) RemoveLibraryPath(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17removeLibraryPathERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_RemoveLibraryPath(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	var nilthis *QCoreApplication
	nilthis.RemoveLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:152
// index:0
// static
// bool installTranslator(class QTranslator *)
func (this *QCoreApplication) InstallTranslator(messageFile unsafe.Pointer) {
	// 0: (messageFile QTranslator *), (messageFile)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17installTranslatorEP11QTranslator", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_InstallTranslator(messageFile unsafe.Pointer) {
	// 0: (messageFile QTranslator *), (messageFile)
	var nilthis *QCoreApplication
	nilthis.InstallTranslator(messageFile)
}

// /usr/include/qt/QtCore/qcoreapplication.h:153
// index:0
// static
// bool removeTranslator(class QTranslator *)
func (this *QCoreApplication) RemoveTranslator(messageFile unsafe.Pointer) {
	// 0: (messageFile QTranslator *), (messageFile)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16removeTranslatorEP11QTranslator", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_RemoveTranslator(messageFile unsafe.Pointer) {
	// 0: (messageFile QTranslator *), (messageFile)
	var nilthis *QCoreApplication
	nilthis.RemoveTranslator(messageFile)
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// static
// QString translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) Translate(context unsafe.Pointer, key unsafe.Pointer, disambiguation unsafe.Pointer, n int) {
	// 0: (context const char *, key const char *, disambiguation const char *, n int), (context, key, disambiguation, n)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9translateEPKcS1_S1_i", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Translate(context unsafe.Pointer, key unsafe.Pointer, disambiguation unsafe.Pointer, n int) {
	// 0: (context const char *, key const char *, disambiguation const char *, n int), (context, key, disambiguation, n)
	var nilthis *QCoreApplication
	nilthis.Translate(context, key, disambiguation, n)
}

// /usr/include/qt/QtCore/qcoreapplication.h:169
// index:0
// static
// void flush()
func (this *QCoreApplication) Flush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication5flushEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Flush() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.Flush()
}

// /usr/include/qt/QtCore/qcoreapplication.h:172
// index:0
// void installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) InstallNativeEventFilter(filterObj unsafe.Pointer) {
	// 0: (, filterObj QAbstractNativeEventFilter *), (filterObj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filterObj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:173
// index:0
// void removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) RemoveNativeEventFilter(filterObj unsafe.Pointer) {
	// 0: (, filterObj QAbstractNativeEventFilter *), (filterObj)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_VOID, this.GetCthis(), filterObj)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:175
// index:0
// static
// bool isQuitLockEnabled()
func (this *QCoreApplication) IsQuitLockEnabled() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17isQuitLockEnabledEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_IsQuitLockEnabled() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.IsQuitLockEnabled()
}

// /usr/include/qt/QtCore/qcoreapplication.h:176
// index:0
// static
// void setQuitLockEnabled(_Bool)
func (this *QCoreApplication) SetQuitLockEnabled(enabled bool) {
	// 0: (enabled bool), (enabled)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setQuitLockEnabledEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	// 0: (enabled bool), (enabled)
	var nilthis *QCoreApplication
	nilthis.SetQuitLockEnabled(enabled)
}

// /usr/include/qt/QtCore/qcoreapplication.h:179
// index:0
// static
// void quit()
func (this *QCoreApplication) Quit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication4quitEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_Quit() {
	// 0: (), ()
	var nilthis *QCoreApplication
	nilthis.Quit()
}

// /usr/include/qt/QtCore/qcoreapplication.h:184
// index:0
// void organizationNameChanged()
func (this *QCoreApplication) OrganizationNameChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication23organizationNameChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:185
// index:0
// void organizationDomainChanged()
func (this *QCoreApplication) OrganizationDomainChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication25organizationDomainChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:186
// index:0
// void applicationNameChanged()
func (this *QCoreApplication) ApplicationNameChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication22applicationNameChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:187
// index:0
// void applicationVersionChanged()
func (this *QCoreApplication) ApplicationVersionChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication25applicationVersionChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:190
// index:0
// virtual
// bool event(class QEvent *)
func (this *QCoreApplication) Event(arg0 unsafe.Pointer) {
	// 0: (, QEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:192
// index:0
// virtual
// bool compressEvent(class QEvent *, class QObject *, class QPostEventList *)
func (this *QCoreApplication) CompressEvent(arg0 unsafe.Pointer, receiver unsafe.Pointer, arg2 unsafe.Pointer) {
	// 0: (, QEvent * arg0, receiver QObject *, QPostEventList * arg2), (arg0, receiver, arg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13compressEventEP6QEventP7QObjectP14QPostEventList", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0, receiver, arg2)
	gopp.ErrPrint(err, rv)
}

//  body block end
