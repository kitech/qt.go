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

// bool event(QEvent *)
func (this *QCoreApplication) InheritEvent(f func(arg0 *QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

/*

 */
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

/*

 */
func (this *QCoreApplication) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QCoreApplication10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qcoreapplication.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCoreApplication(int &, char **, int)

/*
Constructs a Qt core application. Core applications are applications without a graphical user interface. Such applications are used at the console or as server processes.

The argc and argv arguments are processed by the application, and made available in a more convenient form by the arguments() function.

Warning: The data referred to by argc and argv must stay valid for the entire lifetime of the QCoreApplication object. In addition, argc must be greater than zero and argv must contain at least one valid character string.
*/
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

/*
Constructs a Qt core application. Core applications are applications without a graphical user interface. Such applications are used at the console or as server processes.

The argc and argv arguments are processed by the application, and made available in a more convenient form by the arguments() function.

Warning: The data referred to by argc and argv must stay valid for the entire lifetime of the QCoreApplication object. In addition, argc must be greater than zero and argv must contain at least one valid character string.
*/
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

/*

 */
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

/*
Returns the list of command-line arguments.

Usually arguments().at(0) is the program name, arguments().at(1) is the first argument, and arguments().last() is the last argument. See the note below about Windows.

Calling this function is slow - you should store the result in a variable when parsing the command line.

Warning: On Unix, this list is built from the argc and argv parameters passed to the constructor in the main() function. The string-data in argv is interpreted using QString::fromLocal8Bit(); hence it is not possible to pass, for example, Japanese command line arguments on a system that runs in a Latin1 locale. Most modern Unix systems do not have this limitation, as they are Unicode-based.

On Windows, the list is built from the argc and argv parameters only if modified argv/argc parameters are passed to the constructor. In that case, encoding problems might occur.

Otherwise, the arguments() are constructed from the return value of GetCommandLine(). As a result of this, the string given by arguments().at(0) might not be the program name on Windows, depending on how the application was started.

This function was introduced in  Qt 4.1.

See also applicationFilePath() and QCommandLineParser.
*/
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
// [-2] void setAttribute(Qt::ApplicationAttribute, bool)

/*
Sets the attribute attribute if on is true; otherwise clears the attribute.

See also testAttribute().
*/
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
// [-2] void setAttribute(Qt::ApplicationAttribute, bool)

/*
Sets the attribute attribute if on is true; otherwise clears the attribute.

See also testAttribute().
*/
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

/*
Returns true if attribute attribute is set; otherwise returns false.

See also setAttribute().
*/
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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

/*

 */
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
// [-2] void setSetuidAllowed(bool)

/*
Allows the application to run setuid on UNIX platforms if allow is true.

If allow is false (the default) and Qt detects the application is running with an effective user id different than the real user id, the application will be aborted when a QCoreApplication instance is created.

Qt is not an appropriate solution for setuid programs due to its large attack surface. However some applications may be required to run in this manner for historical reasons. This flag will prevent Qt from aborting the application when this is detected, and must be set before a QCoreApplication instance is created.

Note: It is strongly recommended not to enable this option since it introduces security risks.

This function was introduced in  Qt 5.3.

See also isSetuidAllowed().
*/
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

/*
Returns true if the application is allowed to run setuid on UNIX platforms.

This function was introduced in  Qt 5.3.

See also QCoreApplication::setSetuidAllowed().
*/
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

/*
Returns a pointer to the application's QCoreApplication (or QGuiApplication/QApplication) instance.

If no instance has been allocated, null is returned.
*/
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

/*
Enters the main event loop and waits until exit() is called. Returns the value that was passed to exit() (which is 0 if exit() is called via quit()).

It is necessary to call this function to start event handling. The main event loop receives events from the window system and dispatches these to the application widgets.

To make your application perform idle processing (by executing a special function whenever there are no pending events), use a QTimer with 0 timeout. More advanced idle processing schemes can be achieved using processEvents().

We recommend that you connect clean-up code to the aboutToQuit() signal, instead of putting it in your application's main() function because on some platforms the exec() call may not return. For example, on Windows when the user logs off, the system terminates the process after Qt closes all top-level windows. Hence, there is no guarantee that the application will have time to exit its event loop and execute code at the end of the main() function after the exec() call.

See also quit(), exit(), processEvents(), and QApplication::exec().
*/
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

/*
Processes all pending events for the calling thread according to the specified flags until there are no more events to process.

You can call this function occasionally when your program is busy performing a long operation (e.g. copying a file).

In the event that you are running a local loop which calls this function continuously, without an event loop, the DeferredDelete events will not be processed. This can affect the behaviour of widgets, e.g. QToolTip, that rely on DeferredDelete events to function properly. An alternative would be to call sendPostedEvents() from within that local loop.

Calling this function processes events only for the calling thread.

Note: This function is thread-safe.

See also exec(), QTimer, QEventLoop::processEvents(), flush(), and sendPostedEvents().
*/
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

/*
Processes all pending events for the calling thread according to the specified flags until there are no more events to process.

You can call this function occasionally when your program is busy performing a long operation (e.g. copying a file).

In the event that you are running a local loop which calls this function continuously, without an event loop, the DeferredDelete events will not be processed. This can affect the behaviour of widgets, e.g. QToolTip, that rely on DeferredDelete events to function properly. An alternative would be to call sendPostedEvents() from within that local loop.

Calling this function processes events only for the calling thread.

Note: This function is thread-safe.

See also exec(), QTimer, QEventLoop::processEvents(), flush(), and sendPostedEvents().
*/
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

/*
Processes all pending events for the calling thread according to the specified flags until there are no more events to process.

You can call this function occasionally when your program is busy performing a long operation (e.g. copying a file).

In the event that you are running a local loop which calls this function continuously, without an event loop, the DeferredDelete events will not be processed. This can affect the behaviour of widgets, e.g. QToolTip, that rely on DeferredDelete events to function properly. An alternative would be to call sendPostedEvents() from within that local loop.

Calling this function processes events only for the calling thread.

Note: This function is thread-safe.

See also exec(), QTimer, QEventLoop::processEvents(), flush(), and sendPostedEvents().
*/
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

/*
Tells the application to exit with a return code.

After this function has been called, the application leaves the main event loop and returns from the call to exec(). The exec() function returns returnCode. If the event loop is not running, this function does nothing.

By convention, a returnCode of 0 means success, and any non-zero value indicates an error.

Note that unlike the C library function of the same name, this function does return to the caller -- it is event processing that stops.

See also quit() and exec().
*/
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

/*
Tells the application to exit with a return code.

After this function has been called, the application leaves the main event loop and returns from the call to exec(). The exec() function returns returnCode. If the event loop is not running, this function does nothing.

By convention, a returnCode of 0 means success, and any non-zero value indicates an error.

Note that unlike the C library function of the same name, this function does return to the caller -- it is event processing that stops.

See also quit() and exec().
*/
func (this *QCoreApplication) Exit__() {
	// arg: 0, int=Int, =Invalid,
	retcode := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication4exitEi", qtrt.FFI_TYPE_POINTER, retcode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:124
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool sendEvent(QObject *, QEvent *)

/*
Sends event event directly to receiver receiver, using the notify() function. Returns the value that was returned from the event handler.

The event is not deleted when the event has been sent. The normal approach is to create the event on the stack, for example:


  QMouseEvent event(QEvent::MouseButtonPress, pos, 0, 0, 0);
  QApplication::sendEvent(mainWindow, &event);



See also postEvent() and notify().
*/
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

/*
Adds the event event, with the object receiver as the receiver of the event, to an event queue and returns immediately.

The event must be allocated on the heap since the post event queue will take ownership of the event and delete it once it has been posted. It is not safe to access the event after it has been posted.

When control returns to the main event loop, all events that are stored in the queue will be sent using the notify() function.

Events are sorted in descending priority order, i.e. events with a high priority are queued before events with a lower priority. The priority can be any integer value, i.e. between INT_MAX and INT_MIN, inclusive; see Qt::EventPriority for more details. Events with equal priority will be processed in the order posted.

Note: This function is thread-safe.

This function was introduced in  Qt 4.3.

See also sendEvent(), notify(), sendPostedEvents(), and Qt::EventPriority.
*/
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

/*
Adds the event event, with the object receiver as the receiver of the event, to an event queue and returns immediately.

The event must be allocated on the heap since the post event queue will take ownership of the event and delete it once it has been posted. It is not safe to access the event after it has been posted.

When control returns to the main event loop, all events that are stored in the queue will be sent using the notify() function.

Events are sorted in descending priority order, i.e. events with a high priority are queued before events with a lower priority. The priority can be any integer value, i.e. between INT_MAX and INT_MIN, inclusive; see Qt::EventPriority for more details. Events with equal priority will be processed in the order posted.

Note: This function is thread-safe.

This function was introduced in  Qt 4.3.

See also sendEvent(), notify(), sendPostedEvents(), and Qt::EventPriority.
*/
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

/*
Immediately dispatches all events which have been previously queued with QCoreApplication::postEvent() and which are for the object receiver and have the event type event_type.

Events from the window system are not dispatched by this function, but by processEvents().

If receiver is null, the events of event_type are sent for all objects. If event_type is 0, all the events are sent for receiver.

Note: This method must be called from the thread in which its QObject parameter, receiver, lives.

See also flush() and postEvent().
*/
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

/*
Immediately dispatches all events which have been previously queued with QCoreApplication::postEvent() and which are for the object receiver and have the event type event_type.

Events from the window system are not dispatched by this function, but by processEvents().

If receiver is null, the events of event_type are sent for all objects. If event_type is 0, all the events are sent for receiver.

Note: This method must be called from the thread in which its QObject parameter, receiver, lives.

See also flush() and postEvent().
*/
func (this *QCoreApplication) SendPostedEvents__() {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	// arg: 1, int=Int, =Invalid,
	event_type := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, event_type)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:126
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void sendPostedEvents(QObject *, int)

/*
Immediately dispatches all events which have been previously queued with QCoreApplication::postEvent() and which are for the object receiver and have the event type event_type.

Events from the window system are not dispatched by this function, but by processEvents().

If receiver is null, the events of event_type are sent for all objects. If event_type is 0, all the events are sent for receiver.

Note: This method must be called from the thread in which its QObject parameter, receiver, lives.

See also flush() and postEvent().
*/
func (this *QCoreApplication) SendPostedEvents__1(receiver QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	event_type := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication16sendPostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, event_type)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:127
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void removePostedEvents(QObject *, int)

/*
Removes all events of the given eventType that were posted using postEvent() for receiver.

The events are not dispatched, instead they are removed from the queue. You should never need to call this function. If you do call it, be aware that killing events may cause receiver to break one or more invariants.

If receiver is null, the events of eventType are removed for all objects. If eventType is 0, all the events are removed for receiver. You should never call this function with eventType of 0. If you do call it in this way, be aware that killing events may cause receiver to break one or more invariants.

Note: This function is thread-safe.

This function was introduced in  Qt 4.3.
*/
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

/*
Removes all events of the given eventType that were posted using postEvent() for receiver.

The events are not dispatched, instead they are removed from the queue. You should never need to call this function. If you do call it, be aware that killing events may cause receiver to break one or more invariants.

If receiver is null, the events of eventType are removed for all objects. If eventType is 0, all the events are removed for receiver. You should never call this function with eventType of 0. If you do call it in this way, be aware that killing events may cause receiver to break one or more invariants.

Note: This function is thread-safe.

This function was introduced in  Qt 4.3.
*/
func (this *QCoreApplication) RemovePostedEvents__(receiver QObject_ITF /*777 QObject **/) {
	var convArg0 unsafe.Pointer
	if receiver != nil && receiver.QObject_PTR() != nil {
		convArg0 = receiver.QObject_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	eventType := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication18removePostedEventsEP7QObjecti", qtrt.FFI_TYPE_POINTER, convArg0, eventType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:129
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasPendingEvents()

/*

 */
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

/*
Returns a pointer to the event dispatcher object for the main thread. If no event dispatcher exists for the thread, this function returns 0.

See also setEventDispatcher().
*/
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

/*
Sets the event dispatcher for the main thread to eventDispatcher. This is only possible as long as there is no event dispatcher installed yet. That is, before QCoreApplication has been instantiated. This method takes ownership of the object.

See also eventDispatcher().
*/
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

/*
Sends event to receiver: receiver->event(event). Returns the value that is returned from the receiver's event handler. Note that this function is called for all events sent to any object in any thread.

For certain types of events (e.g. mouse and key events), the event will be propagated to the receiver's parent and so on up to the top-level object if the receiver is not interested in the event (i.e., it returns false).

There are five different ways that events can be processed; reimplementing this virtual function is just one of them. All five approaches are listed below:
*/
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

/*
Returns true if an application object has not been created yet; otherwise returns false.

See also closingDown().
*/
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

/*
Returns true if the application objects are being destroyed; otherwise returns false.

See also startingUp().
*/
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

/*
Returns the directory that contains the application executable.

For example, if you have installed Qt in the C:\Qt directory, and you run the regexp example, this function will return "C:/Qt/examples/tools/regexp".

On macOS and iOS this will point to the directory actually containing the executable, which may be inside an application bundle (if the application is bundled).

Warning: On Linux, this function will try to get the path from the /proc file system. If that fails, it assumes that argv[0] contains the absolute file name of the executable. The function also assumes that the current directory has not been changed by the application.

See also applicationFilePath().
*/
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

/*
Returns the file path of the application executable.

For example, if you have installed Qt in the /usr/local/qt directory, and you run the regexp example, this function will return "/usr/local/qt/examples/tools/regexp/regexp".

Warning: On Linux, this function will try to get the path from the /proc file system. If that fails, it assumes that argv[0] contains the absolute file name of the executable. The function also assumes that the current directory has not been changed by the application.

See also applicationDirPath().
*/
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

/*
Returns the current process ID for the application.

This function was introduced in  Qt 4.4.
*/
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

/*
Sets the list of directories to search when loading libraries to paths. All existing paths will be deleted and the path list will consist of the paths given in paths.

The library paths are reset to the default when an instance of QCoreApplication is destructed.

See also libraryPaths(), addLibraryPath(), removeLibraryPath(), and QLibrary.
*/
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

/*
Returns a list of paths that the application will search when dynamically loading libraries.

The return value of this function may change when a QCoreApplication is created. It is not recommended to call it before creating a QCoreApplication. The directory of the application executable (not the working directory) is part of the list if it is known. In order to make it known a QCoreApplication has to be constructed as it will use argv[0] to find it.

Qt provides default library paths, but they can also be set using a qt.conf file. Paths specified in this file will override default values. Note that if the qt.conf file is in the directory of the application executable, it may not be found until a QCoreApplication is created. If it is not found when calling this function, the default library paths will be used.

The list will include the installation directory for plugins if it exists (the default installation directory for plugins is INSTALL/plugins, where INSTALL is the directory where Qt was installed). The colon separated entries of the QT_PLUGIN_PATH environment variable are always added. The plugin installation directory (and its existence) may change when the directory of the application executable becomes known.

If you want to iterate over the list, you can use the foreach pseudo-keyword:


  foreach (const QString &path, app.libraryPaths())
      do_something(path);



See also setLibraryPaths(), addLibraryPath(), removeLibraryPath(), QLibrary, and How to Create Qt Plugins.
*/
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

/*
Prepends path to the beginning of the library path list, ensuring that it is searched for libraries first. If path is empty or already in the path list, the path list is not changed.

The default path list consists of a single entry, the installation directory for plugins. The default installation directory for plugins is INSTALL/plugins, where INSTALL is the directory where Qt was installed.

The library paths are reset to the default when an instance of QCoreApplication is destructed.

See also removeLibraryPath(), libraryPaths(), and setLibraryPaths().
*/
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

/*
Removes path from the library path list. If path is empty or not in the path list, the list is not changed.

The library paths are reset to the default when an instance of QCoreApplication is destructed.

See also addLibraryPath(), libraryPaths(), and setLibraryPaths().
*/
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

/*
Adds the translation file translationFile to the list of translation files to be used for translations.

Multiple translation files can be installed. Translations are searched for in the reverse order in which they were installed, so the most recently installed translation file is searched first and the first translation file installed is searched last. The search stops as soon as a translation containing a matching string is found.

Installing or removing a QTranslator, or changing an installed QTranslator generates a LanguageChange event for the QCoreApplication instance. A QApplication instance will propagate the event to all toplevel widgets, where a reimplementation of changeEvent can re-translate the user interface by passing user-visible strings via the tr() function to the respective property setters. User-interface classes generated by Qt Designer provide a retranslateUi() function that can be called.

The function returns true on success and false on failure.

See also removeTranslator(), translate(), QTranslator::load(), and Dynamic Translation.
*/
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

/*
Removes the translation file translationFile from the list of translation files used by this application. (It does not delete the translation file from the file system.)

The function returns true on success and false on failure.

See also installTranslator(), translate(), and QObject::tr().
*/
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

/*
Returns the translation text for sourceText, by querying the installed translation files. The translation files are searched from the most recently installed file back to the first installed file.

QObject::tr() provides this functionality more conveniently.

context is typically a class name (e.g., "MyDialog") and sourceText is either English text or a short identifying text.

disambiguation is an identifying string, for when the same sourceText is used in different roles within the same context. By default, it is null.

See the QTranslator and QObject::tr() documentation for more information about contexts, disambiguations and comments.

n is used in conjunction with %n to support plural forms. See QObject::tr() for details.

If none of the translation files contain a translation for sourceText in context, this function returns a QString equivalent of sourceText.

This function is not virtual. You can use alternative translation techniques by subclassing QTranslator.

Note: This function is thread-safe.

See also QObject::tr(), installTranslator(), removeTranslator(), and translate().
*/
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

/*
Returns the translation text for sourceText, by querying the installed translation files. The translation files are searched from the most recently installed file back to the first installed file.

QObject::tr() provides this functionality more conveniently.

context is typically a class name (e.g., "MyDialog") and sourceText is either English text or a short identifying text.

disambiguation is an identifying string, for when the same sourceText is used in different roles within the same context. By default, it is null.

See the QTranslator and QObject::tr() documentation for more information about contexts, disambiguations and comments.

n is used in conjunction with %n to support plural forms. See QObject::tr() for details.

If none of the translation files contain a translation for sourceText in context, this function returns a QString equivalent of sourceText.

This function is not virtual. You can use alternative translation techniques by subclassing QTranslator.

Note: This function is thread-safe.

See also QObject::tr(), installTranslator(), removeTranslator(), and translate().
*/
func (this *QCoreApplication) Translate__(context string, key string) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, const char *=Pointer, =Invalid,
	var convArg2 unsafe.Pointer
	// arg: 3, int=Int, =Invalid,
	n := int(-1)
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

/*
Returns the translation text for sourceText, by querying the installed translation files. The translation files are searched from the most recently installed file back to the first installed file.

QObject::tr() provides this functionality more conveniently.

context is typically a class name (e.g., "MyDialog") and sourceText is either English text or a short identifying text.

disambiguation is an identifying string, for when the same sourceText is used in different roles within the same context. By default, it is null.

See the QTranslator and QObject::tr() documentation for more information about contexts, disambiguations and comments.

n is used in conjunction with %n to support plural forms. See QObject::tr() for details.

If none of the translation files contain a translation for sourceText in context, this function returns a QString equivalent of sourceText.

This function is not virtual. You can use alternative translation techniques by subclassing QTranslator.

Note: This function is thread-safe.

See also QObject::tr(), installTranslator(), removeTranslator(), and translate().
*/
func (this *QCoreApplication) Translate__1(context string, key string, disambiguation string) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(key)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, int=Int, =Invalid,
	n := int(-1)
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

/*

 */
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

/*
Installs an event filter filterObj for all native events received by the application in the main thread.

The event filter filterObj receives events via its nativeEventFilter() function, which is called for all native events received in the main thread.

The QAbstractNativeEventFilter::nativeEventFilter() function should return true if the event should be filtered, i.e. stopped. It should return false to allow normal Qt processing to continue: the native event can then be translated into a QEvent and handled by the standard Qt event filtering, e.g. QObject::installEventFilter().

If multiple event filters are installed, the filter that was installed last is activated first.

Note: The filter function set here receives native messages, i.e. MSG or XCB event structs.

Note: Native event filters will be disabled in the application when the Qt::AA_PluginApplication attribute is set.

For maximum portability, you should always try to use QEvent and QObject::installEventFilter() whenever possible.

This function was introduced in  Qt 5.0.

See also QObject::installEventFilter().
*/
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

/*
Removes an event filterObject from this object. The request is ignored if such an event filter has not been installed.

All event filters for this object are automatically removed when this object is destroyed.

It is always safe to remove an event filter, even during event filter activation (i.e. from the nativeEventFilter() function).

This function was introduced in  Qt 5.0.

See also installNativeEventFilter().
*/
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

/*

 */
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
// [-2] void setQuitLockEnabled(bool)

/*

 */
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

/*
Tells the application to exit with return code 0 (success). Equivalent to calling QCoreApplication::exit(0).

It's common to connect the QGuiApplication::lastWindowClosed() signal to quit(), and you also often connect e.g. QAbstractButton::clicked() or signals in QAction, QMenu, or QMenuBar to it.

Example:


  QPushButton *quitButton = new QPushButton("Quit");
  connect(quitButton, SIGNAL(clicked()), &app, SLOT(quit()));



See also exit(), aboutToQuit(), and QGuiApplication::lastWindowClosed().
*/
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

/*

 */
func (this *QCoreApplication) OrganizationNameChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication23organizationNameChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void organizationDomainChanged()

/*

 */
func (this *QCoreApplication) OrganizationDomainChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication25organizationDomainChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationNameChanged()

/*

 */
func (this *QCoreApplication) ApplicationNameChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication22applicationNameChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void applicationVersionChanged()

/*

 */
func (this *QCoreApplication) ApplicationVersionChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication25applicationVersionChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:190
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QCoreApplication) Event(arg0 QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QCoreApplication5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*


 */
type QCoreApplication__ = int

//
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
