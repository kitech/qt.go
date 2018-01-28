package qtcore

// /usr/include/qt/QtCore/qcoreapplication.h
// #include <qcoreapplication.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/cffiqt"
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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QCoreApplication) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQCoreApplicationFromPointer(cthis unsafe.Pointer) *QCoreApplication {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QCoreApplication{bcthis0}
}
func (*QCoreApplication) NewFromPointer(cthis unsafe.Pointer) *QCoreApplication {
	return NewQCoreApplicationFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcoreapplication.h:78
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QCoreApplication) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QCoreApplication10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qcoreapplication.h:91
// index:0
// Public
// void QCoreApplication(int &, char **, int)
func NewQCoreApplication(argc int, argv []string, arg2 int) *QCoreApplication {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplicationC2ERiPPci", ffiqt.FFI_TYPE_VOID, cthis, &argc, convArg1, arg2)
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
func (this *QCoreApplication) TestAttribute(attribute int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13testAttributeEN2Qt20ApplicationAttributeE", ffiqt.FFI_TYPE_POINTER, attribute)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_TestAttribute(attribute int) bool {
	var nilthis *QCoreApplication
	rv := nilthis.TestAttribute(attribute)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:104
// index:0
// Public static
// void setOrganizationDomain(const QString &)
func (this *QCoreApplication) SetOrganizationDomain(orgDomain *QString) {
	var convArg0 = orgDomain.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication21setOrganizationDomainERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QCoreApplication) OrganizationDomain() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18organizationDomainEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCoreApplication_OrganizationDomain() *QString /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.OrganizationDomain()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:106
// index:0
// Public static
// void setOrganizationName(const QString &)
func (this *QCoreApplication) SetOrganizationName(orgName *QString) {
	var convArg0 = orgName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication19setOrganizationNameERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QCoreApplication) OrganizationName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16organizationNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCoreApplication_OrganizationName() *QString /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.OrganizationName()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:108
// index:0
// Public static
// void setApplicationName(const QString &)
func (this *QCoreApplication) SetApplicationName(application *QString) {
	var convArg0 = application.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setApplicationNameERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QCoreApplication) ApplicationName() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15applicationNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCoreApplication_ApplicationName() *QString /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationName()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:110
// index:0
// Public static
// void setApplicationVersion(const QString &)
func (this *QCoreApplication) SetApplicationVersion(version *QString) {
	var convArg0 = version.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication21setApplicationVersionERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
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
func (this *QCoreApplication) ApplicationVersion() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18applicationVersionEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCoreApplication_ApplicationVersion() *QString /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationVersion()
	return rv
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
func (this *QCoreApplication) IsSetuidAllowed() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15isSetuidAllowedEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_IsSetuidAllowed() bool {
	var nilthis *QCoreApplication
	rv := nilthis.IsSetuidAllowed()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:116
// index:0
// Public static inline
// QCoreApplication * instance()
func (this *QCoreApplication) Instance() *QCoreApplication /*777 QCoreApplication **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication8instanceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQCoreApplicationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QCoreApplication_Instance() *QCoreApplication /*777 QCoreApplication **/ {
	var nilthis *QCoreApplication
	rv := nilthis.Instance()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:119
// index:0
// Public static
// int exec()
func (this *QCoreApplication) Exec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication4execEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QCoreApplication_Exec() int {
	var nilthis *QCoreApplication
	rv := nilthis.Exec()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:120
// index:0
// Public static
// void processEvents(QEventLoop::ProcessEventsFlags)
func (this *QCoreApplication) ProcessEvents(flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEE", ffiqt.FFI_TYPE_POINTER, flags)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ProcessEvents(flags int) {
	var nilthis *QCoreApplication
	nilthis.ProcessEvents(flags)
}

// /usr/include/qt/QtCore/qcoreapplication.h:121
// index:1
// Public static
// void processEvents(QEventLoop::ProcessEventsFlags, int)
func (this *QCoreApplication) ProcessEvents_1(flags int, maxtime int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEEi", ffiqt.FFI_TYPE_POINTER, flags, maxtime)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_ProcessEvents_1(flags int, maxtime int) {
	var nilthis *QCoreApplication
	nilthis.ProcessEvents_1(flags, maxtime)
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
// bool sendEvent(QObject *, QEvent *)
func (this *QCoreApplication) SendEvent(receiver *QObject /*777 QObject **/, event *QEvent /*777 QEvent **/) bool {
	var convArg0 = receiver.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_SendEvent(receiver *QObject /*777 QObject **/, event *QEvent /*777 QEvent **/) bool {
	var nilthis *QCoreApplication
	rv := nilthis.SendEvent(receiver, event)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:125
// index:0
// Public static
// void postEvent(QObject *, QEvent *, int)
func (this *QCoreApplication) PostEvent(receiver *QObject /*777 QObject **/, event *QEvent /*777 QEvent **/, priority int) {
	var convArg0 = receiver.GetCthis()
	var convArg1 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9postEventEP7QObjectP6QEventi", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, priority)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_PostEvent(receiver *QObject /*777 QObject **/, event *QEvent /*777 QEvent **/, priority int) {
	var nilthis *QCoreApplication
	nilthis.PostEvent(receiver, event, priority)
}

// /usr/include/qt/QtCore/qcoreapplication.h:126
// index:0
// Public static
// void sendPostedEvents(QObject *, int)
func (this *QCoreApplication) SendPostedEvents(receiver *QObject /*777 QObject **/, event_type int) {
	var convArg0 = receiver.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", ffiqt.FFI_TYPE_POINTER, convArg0, event_type)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SendPostedEvents(receiver *QObject /*777 QObject **/, event_type int) {
	var nilthis *QCoreApplication
	nilthis.SendPostedEvents(receiver, event_type)
}

// /usr/include/qt/QtCore/qcoreapplication.h:127
// index:0
// Public static
// void removePostedEvents(QObject *, int)
func (this *QCoreApplication) RemovePostedEvents(receiver *QObject /*777 QObject **/, eventType int) {
	var convArg0 = receiver.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18removePostedEventsEP7QObjecti", ffiqt.FFI_TYPE_POINTER, convArg0, eventType)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_RemovePostedEvents(receiver *QObject /*777 QObject **/, eventType int) {
	var nilthis *QCoreApplication
	nilthis.RemovePostedEvents(receiver, eventType)
}

// /usr/include/qt/QtCore/qcoreapplication.h:129
// index:0
// Public static
// bool hasPendingEvents()
func (this *QCoreApplication) HasPendingEvents() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16hasPendingEventsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_HasPendingEvents() bool {
	var nilthis *QCoreApplication
	rv := nilthis.HasPendingEvents()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:131
// index:0
// Public static
// QAbstractEventDispatcher * eventDispatcher()
func (this *QCoreApplication) EventDispatcher() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15eventDispatcherEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QCoreApplication_EventDispatcher() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	var nilthis *QCoreApplication
	rv := nilthis.EventDispatcher()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:132
// index:0
// Public static
// void setEventDispatcher(QAbstractEventDispatcher *)
func (this *QCoreApplication) SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/) {
	var convArg0 = eventDispatcher.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/) {
	var nilthis *QCoreApplication
	nilthis.SetEventDispatcher(eventDispatcher)
}

// /usr/include/qt/QtCore/qcoreapplication.h:134
// index:0
// Public virtual
// bool notify(QObject *, QEvent *)
func (this *QCoreApplication) Notify(arg0 *QObject /*777 QObject **/, arg1 *QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreapplication.h:136
// index:0
// Public static
// bool startingUp()
func (this *QCoreApplication) StartingUp() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication10startingUpEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_StartingUp() bool {
	var nilthis *QCoreApplication
	rv := nilthis.StartingUp()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:137
// index:0
// Public static
// bool closingDown()
func (this *QCoreApplication) ClosingDown() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication11closingDownEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_ClosingDown() bool {
	var nilthis *QCoreApplication
	rv := nilthis.ClosingDown()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:140
// index:0
// Public static
// QString applicationDirPath()
func (this *QCoreApplication) ApplicationDirPath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication18applicationDirPathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCoreApplication_ApplicationDirPath() *QString /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationDirPath()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:141
// index:0
// Public static
// QString applicationFilePath()
func (this *QCoreApplication) ApplicationFilePath() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication19applicationFilePathEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCoreApplication_ApplicationFilePath() *QString /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationFilePath()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:142
// index:0
// Public static
// qint64 applicationPid()
func (this *QCoreApplication) ApplicationPid() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication14applicationPidEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return int64(rv) // 222
}
func QCoreApplication_ApplicationPid() int64 {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationPid()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:145
// index:0
// Public static
// void setLibraryPaths(const QStringList &)
func (this *QCoreApplication) SetLibraryPaths(arg0 *QStringList) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication15setLibraryPathsERK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_SetLibraryPaths(arg0 *QStringList) {
	var nilthis *QCoreApplication
	nilthis.SetLibraryPaths(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:147
// index:0
// Public static
// void addLibraryPath(const QString &)
func (this *QCoreApplication) AddLibraryPath(arg0 *QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication14addLibraryPathERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_AddLibraryPath(arg0 *QString) {
	var nilthis *QCoreApplication
	nilthis.AddLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:148
// index:0
// Public static
// void removeLibraryPath(const QString &)
func (this *QCoreApplication) RemoveLibraryPath(arg0 *QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17removeLibraryPathERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QCoreApplication_RemoveLibraryPath(arg0 *QString) {
	var nilthis *QCoreApplication
	nilthis.RemoveLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:152
// index:0
// Public static
// bool installTranslator(QTranslator *)
func (this *QCoreApplication) InstallTranslator(messageFile *QTranslator /*777 QTranslator **/) bool {
	var convArg0 = messageFile.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17installTranslatorEP11QTranslator", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_InstallTranslator(messageFile *QTranslator /*777 QTranslator **/) bool {
	var nilthis *QCoreApplication
	rv := nilthis.InstallTranslator(messageFile)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:153
// index:0
// Public static
// bool removeTranslator(QTranslator *)
func (this *QCoreApplication) RemoveTranslator(messageFile *QTranslator /*777 QTranslator **/) bool {
	var convArg0 = messageFile.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication16removeTranslatorEP11QTranslator", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_RemoveTranslator(messageFile *QTranslator /*777 QTranslator **/) bool {
	var nilthis *QCoreApplication
	rv := nilthis.RemoveTranslator(messageFile)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static
// QString translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) Translate(context string, key string, disambiguation string, n int) *QString /*123*/ {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication9translateEPKcS1_S1_i", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, n)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QCoreApplication_Translate(context string, key string, disambiguation string, n int) *QString /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.Translate(context, key, disambiguation, n)
	return rv
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
// void installNativeEventFilter(QAbstractNativeEventFilter *)
func (this *QCoreApplication) InstallNativeEventFilter(filterObj *QAbstractNativeEventFilter /*777 QAbstractNativeEventFilter **/) {
	var convArg0 = filterObj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:173
// index:0
// Public
// void removeNativeEventFilter(QAbstractNativeEventFilter *)
func (this *QCoreApplication) RemoveNativeEventFilter(filterObj *QAbstractNativeEventFilter /*777 QAbstractNativeEventFilter **/) {
	var convArg0 = filterObj.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:175
// index:0
// Public static
// bool isQuitLockEnabled()
func (this *QCoreApplication) IsQuitLockEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication17isQuitLockEnabledEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QCoreApplication_IsQuitLockEnabled() bool {
	var nilthis *QCoreApplication
	rv := nilthis.IsQuitLockEnabled()
	return rv
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
// bool event(QEvent *)
func (this *QCoreApplication) Event(arg0 *QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QCoreApplication5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QCoreApplication__ = int

const QCoreApplication__ApplicationFlags QCoreApplication__ = 330240

//  body block end
