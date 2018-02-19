package qtcore

// /usr/include/qt/QtCore/qcoreapplication.h
// #include <qcoreapplication.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// bool event(class QEvent *)
func (this *QCoreApplication) InheritEvent(f func(arg0 *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

type QCoreApplication struct {
	*QObject
}
type QCoreApplication_ITF interface {
	QObject_ITF
	QCoreApplication_PTR() *QCoreApplication
}

func (ptr *QCoreApplication) QCoreApplication_PTR() *QCoreApplication { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const
func (this *QCoreApplication) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCoreApplication10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qcoreapplication.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCoreApplication(int &, char **, int)
func NewQCoreApplication(argc int, argv []string, arg2 int) *QCoreApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplicationC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCoreApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCoreApplication")
	return gothis
}

// /usr/include/qt/QtCore/qcoreapplication.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCoreApplication(int &, char **, int)
func NewQCoreApplication__(argc int, argv []string) *QCoreApplication {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	// arg: 2, int=Int, =Invalid,
	arg2 := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplicationC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, arg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCoreApplicationFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCoreApplication")
	return gothis
}

// /usr/include/qt/QtCore/qcoreapplication.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCoreApplication()
func DeleteQCoreApplication(this *QCoreApplication) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplicationD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcoreapplication.h:99
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList arguments()
func (this *QCoreApplication) Arguments() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication9argumentsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QCoreApplication_Arguments() *QStringList /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.Arguments()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::ApplicationAttribute, _Bool)
func (this *QCoreApplication) SetAttribute(attribute int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication12setAttributeEN2Qt20ApplicationAttributeEb", qtrt.FFI_TYPE_POINTER, attribute, on)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetAttribute(attribute int, on bool) {
	var nilthis *QCoreApplication
	nilthis.SetAttribute(attribute, on)
}

// /usr/include/qt/QtCore/qcoreapplication.h:101
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::ApplicationAttribute, _Bool)
func (this *QCoreApplication) SetAttribute__(attribute int) {
	// arg: 1, bool=Bool, =Invalid,
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication12setAttributeEN2Qt20ApplicationAttributeEb", qtrt.FFI_TYPE_POINTER, attribute, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:102
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool testAttribute(Qt::ApplicationAttribute)
func (this *QCoreApplication) TestAttribute(attribute int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication13testAttributeEN2Qt20ApplicationAttributeE", qtrt.FFI_TYPE_POINTER, attribute)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_TestAttribute(attribute int) bool {
	var nilthis *QCoreApplication
	rv := nilthis.TestAttribute(attribute)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:104
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setOrganizationDomain(const QString &)
func (this *QCoreApplication) SetOrganizationDomain(orgDomain string) {
	var tmpArg0 = NewQString_5(orgDomain)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication21setOrganizationDomainERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetOrganizationDomain(orgDomain string) {
	var nilthis *QCoreApplication
	nilthis.SetOrganizationDomain(orgDomain)
}

// /usr/include/qt/QtCore/qcoreapplication.h:105
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString organizationDomain()
func (this *QCoreApplication) OrganizationDomain() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18organizationDomainEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_OrganizationDomain() string {
	var nilthis *QCoreApplication
	rv := nilthis.OrganizationDomain()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setOrganizationName(const QString &)
func (this *QCoreApplication) SetOrganizationName(orgName string) {
	var tmpArg0 = NewQString_5(orgName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication19setOrganizationNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetOrganizationName(orgName string) {
	var nilthis *QCoreApplication
	nilthis.SetOrganizationName(orgName)
}

// /usr/include/qt/QtCore/qcoreapplication.h:107
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString organizationName()
func (this *QCoreApplication) OrganizationName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16organizationNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_OrganizationName() string {
	var nilthis *QCoreApplication
	rv := nilthis.OrganizationName()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:108
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setApplicationName(const QString &)
func (this *QCoreApplication) SetApplicationName(application string) {
	var tmpArg0 = NewQString_5(application)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18setApplicationNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetApplicationName(application string) {
	var nilthis *QCoreApplication
	nilthis.SetApplicationName(application)
}

// /usr/include/qt/QtCore/qcoreapplication.h:109
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString applicationName()
func (this *QCoreApplication) ApplicationName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication15applicationNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_ApplicationName() string {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationName()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:110
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setApplicationVersion(const QString &)
func (this *QCoreApplication) SetApplicationVersion(version string) {
	var tmpArg0 = NewQString_5(version)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication21setApplicationVersionERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetApplicationVersion(version string) {
	var nilthis *QCoreApplication
	nilthis.SetApplicationVersion(version)
}

// /usr/include/qt/QtCore/qcoreapplication.h:111
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString applicationVersion()
func (this *QCoreApplication) ApplicationVersion() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18applicationVersionEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_ApplicationVersion() string {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationVersion()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:113
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSetuidAllowed(_Bool)
func (this *QCoreApplication) SetSetuidAllowed(allow bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16setSetuidAllowedEb", qtrt.FFI_TYPE_POINTER, allow)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetSetuidAllowed(allow bool) {
	var nilthis *QCoreApplication
	nilthis.SetSetuidAllowed(allow)
}

// /usr/include/qt/QtCore/qcoreapplication.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isSetuidAllowed()
func (this *QCoreApplication) IsSetuidAllowed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication15isSetuidAllowedEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_IsSetuidAllowed() bool {
	var nilthis *QCoreApplication
	rv := nilthis.IsSetuidAllowed()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:116
// index:0
// Public static inline Visibility=Default Availability=Available
// [8] QCoreApplication * instance()
func (this *QCoreApplication) Instance() *QCoreApplication /*777 QCoreApplication **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication8instanceEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQCoreApplicationFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QCoreApplication_Instance() *QCoreApplication /*777 QCoreApplication **/ {
	var nilthis *QCoreApplication
	rv := nilthis.Instance()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:119
// index:0
// Public static Visibility=Default Availability=Available
// [4] int exec()
func (this *QCoreApplication) Exec() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication4execEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QCoreApplication_Exec() int {
	var nilthis *QCoreApplication
	rv := nilthis.Exec()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void processEvents(QEventLoop::ProcessEventsFlags)
func (this *QCoreApplication) ProcessEvents(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, flags)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_ProcessEvents(flags int) {
	var nilthis *QCoreApplication
	nilthis.ProcessEvents(flags)
}

// /usr/include/qt/QtCore/qcoreapplication.h:120
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void processEvents(QEventLoop::ProcessEventsFlags)
func (this *QCoreApplication) ProcessEvents__() {
	// arg: 0, QEventLoop::ProcessEventsFlags=Elaborated, QEventLoop::ProcessEventsFlags=Typedef, QFlags<QEventLoop::ProcessEventsFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEE", qtrt.FFI_TYPE_POINTER, flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:121
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void processEvents(QEventLoop::ProcessEventsFlags, int)
func (this *QCoreApplication) ProcessEvents_1(flags int, maxtime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication13processEventsE6QFlagsIN10QEventLoop17ProcessEventsFlagEEi", qtrt.FFI_TYPE_POINTER, flags, maxtime)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_ProcessEvents_1(flags int, maxtime int) {
	var nilthis *QCoreApplication
	nilthis.ProcessEvents_1(flags, maxtime)
}

// /usr/include/qt/QtCore/qcoreapplication.h:122
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void exit(int)
func (this *QCoreApplication) Exit(retcode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication4exitEi", qtrt.FFI_TYPE_POINTER, retcode)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_Exit(retcode int) {
	var nilthis *QCoreApplication
	nilthis.Exit(retcode)
}

// /usr/include/qt/QtCore/qcoreapplication.h:122
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void exit(int)
func (this *QCoreApplication) Exit__() {
	// arg: 0, int=Int, =Invalid,
	retcode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication4exitEi", qtrt.FFI_TYPE_POINTER, retcode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool sendEvent(QObject *, QEvent *)
func (this *QCoreApplication) SendEvent(receiver QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_SendEvent(receiver QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/) bool {
	var nilthis *QCoreApplication
	rv := nilthis.SendEvent(receiver, event)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void postEvent(QObject *, QEvent *, int)
func (this *QCoreApplication) PostEvent(receiver QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/, priority int) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication9postEventEP7QObjectP6QEventi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, priority)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_PostEvent(receiver QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/, priority int) {
	var nilthis *QCoreApplication
	nilthis.PostEvent(receiver, event, priority)
}

// /usr/include/qt/QtCore/qcoreapplication.h:125
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void postEvent(QObject *, QEvent *, int)
func (this *QCoreApplication) PostEvent__(receiver QObject_ITF /*777 QObject **/, event QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	priority := 0 /*Qt::NormalEventPriority*/
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication9postEventEP7QObjectP6QEventi", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, priority)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:126
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void sendPostedEvents(QObject *, int)
func (this *QCoreApplication) SendPostedEvents(receiver QObject_ITF /*777 QObject **/, event_type int) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, event_type)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SendPostedEvents(receiver QObject_ITF /*777 QObject **/, event_type int) {
	var nilthis *QCoreApplication
	nilthis.SendPostedEvents(receiver, event_type)
}

// /usr/include/qt/QtCore/qcoreapplication.h:126
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void sendPostedEvents(QObject *, int)
func (this *QCoreApplication) SendPostedEvents__() {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	event_type := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, event_type)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:126
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void sendPostedEvents(QObject *, int)
func (this *QCoreApplication) SendPostedEvents__1(receiver QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	event_type := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, event_type)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void removePostedEvents(QObject *, int)
func (this *QCoreApplication) RemovePostedEvents(receiver QObject_ITF /*777 QObject **/, eventType int) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18removePostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, eventType)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_RemovePostedEvents(receiver QObject_ITF /*777 QObject **/, eventType int) {
	var nilthis *QCoreApplication
	nilthis.RemovePostedEvents(receiver, eventType)
}

// /usr/include/qt/QtCore/qcoreapplication.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void removePostedEvents(QObject *, int)
func (this *QCoreApplication) RemovePostedEvents__(receiver QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	eventType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18removePostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, eventType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:129
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasPendingEvents()
func (this *QCoreApplication) HasPendingEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16hasPendingEventsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_HasPendingEvents() bool {
	var nilthis *QCoreApplication
	rv := nilthis.HasPendingEvents()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:131
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAbstractEventDispatcher * eventDispatcher()
func (this *QCoreApplication) EventDispatcher() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication15eventDispatcherEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractEventDispatcherFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QCoreApplication_EventDispatcher() *QAbstractEventDispatcher /*777 QAbstractEventDispatcher **/ {
	var nilthis *QCoreApplication
	rv := nilthis.EventDispatcher()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:132
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setEventDispatcher(QAbstractEventDispatcher *)
func (this *QCoreApplication) SetEventDispatcher(eventDispatcher QAbstractEventDispatcher_ITF /*777 QAbstractEventDispatcher **/) {
	var convArg0 unsafe.Pointer
	if eventDispatcher != nil && eventDispatcher.QAbstractEventDispatcher_PTR() != nil {
		convArg0 = eventDispatcher.QAbstractEventDispatcher_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetEventDispatcher(eventDispatcher QAbstractEventDispatcher_ITF /*777 QAbstractEventDispatcher **/) {
	var nilthis *QCoreApplication
	nilthis.SetEventDispatcher(eventDispatcher)
}

// /usr/include/qt/QtCore/qcoreapplication.h:134
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool notify(QObject *, QEvent *)
func (this *QCoreApplication) Notify(arg0 QObject_ITF /*777 QObject **/, arg1 QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QObject_PTR() != nil {
		convArg0 = arg0.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication6notifyEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreapplication.h:136
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool startingUp()
func (this *QCoreApplication) StartingUp() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication10startingUpEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_StartingUp() bool {
	var nilthis *QCoreApplication
	rv := nilthis.StartingUp()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:137
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool closingDown()
func (this *QCoreApplication) ClosingDown() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication11closingDownEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_ClosingDown() bool {
	var nilthis *QCoreApplication
	rv := nilthis.ClosingDown()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:140
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString applicationDirPath()
func (this *QCoreApplication) ApplicationDirPath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18applicationDirPathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_ApplicationDirPath() string {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationDirPath()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:141
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString applicationFilePath()
func (this *QCoreApplication) ApplicationFilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication19applicationFilePathEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_ApplicationFilePath() string {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationFilePath()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:142
// index:0
// Public static Visibility=Default Availability=Available
// [8] qint64 applicationPid()
func (this *QCoreApplication) ApplicationPid() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication14applicationPidEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}
func QCoreApplication_ApplicationPid() int64 {
	var nilthis *QCoreApplication
	rv := nilthis.ApplicationPid()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:145
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setLibraryPaths(const QStringList &)
func (this *QCoreApplication) SetLibraryPaths(arg0 QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStringList_PTR() != nil {
		convArg0 = arg0.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication15setLibraryPathsERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetLibraryPaths(arg0 QStringList_ITF) {
	var nilthis *QCoreApplication
	nilthis.SetLibraryPaths(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:146
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList libraryPaths()
func (this *QCoreApplication) LibraryPaths() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication12libraryPathsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QCoreApplication_LibraryPaths() *QStringList /*123*/ {
	var nilthis *QCoreApplication
	rv := nilthis.LibraryPaths()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:147
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void addLibraryPath(const QString &)
func (this *QCoreApplication) AddLibraryPath(arg0 string) {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication14addLibraryPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_AddLibraryPath(arg0 string) {
	var nilthis *QCoreApplication
	nilthis.AddLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:148
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void removeLibraryPath(const QString &)
func (this *QCoreApplication) RemoveLibraryPath(arg0 string) {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication17removeLibraryPathERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_RemoveLibraryPath(arg0 string) {
	var nilthis *QCoreApplication
	nilthis.RemoveLibraryPath(arg0)
}

// /usr/include/qt/QtCore/qcoreapplication.h:152
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool installTranslator(QTranslator *)
func (this *QCoreApplication) InstallTranslator(messageFile QTranslator_ITF /*777 QTranslator **/) bool {
	var convArg0 unsafe.Pointer
	if messageFile != nil && messageFile.QTranslator_PTR() != nil {
		convArg0 = messageFile.QTranslator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication17installTranslatorEP11QTranslator", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_InstallTranslator(messageFile QTranslator_ITF /*777 QTranslator **/) bool {
	var nilthis *QCoreApplication
	rv := nilthis.InstallTranslator(messageFile)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool removeTranslator(QTranslator *)
func (this *QCoreApplication) RemoveTranslator(messageFile QTranslator_ITF /*777 QTranslator **/) bool {
	var convArg0 unsafe.Pointer
	if messageFile != nil && messageFile.QTranslator_PTR() != nil {
		convArg0 = messageFile.QTranslator_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16removeTranslatorEP11QTranslator", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_RemoveTranslator(messageFile QTranslator_ITF /*777 QTranslator **/) bool {
	var nilthis *QCoreApplication
	rv := nilthis.RemoveTranslator(messageFile)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) Translate(context string, key string, disambiguation string, n int) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication9translateEPKcS1_S1_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_Translate(context string, key string, disambiguation string, n int) string {
	var nilthis *QCoreApplication
	rv := nilthis.Translate(context, key, disambiguation, n)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) Translate__(context string, key string) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	// arg: 3, int=Int, =Invalid,
	n := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication9translateEPKcS1_S1_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) Translate__1(context string, key string, disambiguation string) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, int=Int, =Invalid,
	n := -1
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication9translateEPKcS1_S1_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcoreapplication.h:169
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void flush()
func (this *QCoreApplication) Flush() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication5flushEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_Flush() {
	var nilthis *QCoreApplication
	nilthis.Flush()
}

// /usr/include/qt/QtCore/qcoreapplication.h:172
// index:0
// Public Visibility=Default Availability=Available
// [-2] void installNativeEventFilter(QAbstractNativeEventFilter *)
func (this *QCoreApplication) InstallNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF /*777 QAbstractNativeEventFilter **/) {
	var convArg0 unsafe.Pointer
	if filterObj != nil && filterObj.QAbstractNativeEventFilter_PTR() != nil {
		convArg0 = filterObj.QAbstractNativeEventFilter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:173
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeNativeEventFilter(QAbstractNativeEventFilter *)
func (this *QCoreApplication) RemoveNativeEventFilter(filterObj QAbstractNativeEventFilter_ITF /*777 QAbstractNativeEventFilter **/) {
	var convArg0 unsafe.Pointer
	if filterObj != nil && filterObj.QAbstractNativeEventFilter_PTR() != nil {
		convArg0 = filterObj.QAbstractNativeEventFilter_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:175
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isQuitLockEnabled()
func (this *QCoreApplication) IsQuitLockEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication17isQuitLockEnabledEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QCoreApplication_IsQuitLockEnabled() bool {
	var nilthis *QCoreApplication
	rv := nilthis.IsQuitLockEnabled()
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:176
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setQuitLockEnabled(_Bool)
func (this *QCoreApplication) SetQuitLockEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18setQuitLockEnabledEb", qtrt.FFI_TYPE_POINTER, enabled)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	var nilthis *QCoreApplication
	nilthis.SetQuitLockEnabled(enabled)
}

// /usr/include/qt/QtCore/qcoreapplication.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void quit()
func (this *QCoreApplication) Quit() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication4quitEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QCoreApplication_Quit() {
	var nilthis *QCoreApplication
	nilthis.Quit()
}

// /usr/include/qt/QtCore/qcoreapplication.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void organizationNameChanged()
func (this *QCoreApplication) OrganizationNameChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication23organizationNameChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void organizationDomainChanged()
func (this *QCoreApplication) OrganizationDomainChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication25organizationDomainChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationNameChanged()
func (this *QCoreApplication) ApplicationNameChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication22applicationNameChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationVersionChanged()
func (this *QCoreApplication) ApplicationVersionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication25applicationVersionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:190
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QCoreApplication) Event(arg0 QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QCoreApplication__ = int

const QCoreApplication__ApplicationFlags QCoreApplication__ = 330241

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
