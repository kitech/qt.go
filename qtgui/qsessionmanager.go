package qtgui

// /usr/include/qt/QtGui/qsessionmanager.h
// #include <qsessionmanager.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QSessionManager struct {
	*qtcore.QObject
}
type QSessionManager_ITF interface {
	qtcore.QObject_ITF
	QSessionManager_PTR() *QSessionManager
}

func (ptr *QSessionManager) QSessionManager_PTR() *QSessionManager { return ptr }

func (this *QSessionManager) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSessionManager) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSessionManagerFromPointer(cthis unsafe.Pointer) *QSessionManager {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSessionManager{bcthis0}
}
func (*QSessionManager) NewFromPointer(cthis unsafe.Pointer) *QSessionManager {
	return NewQSessionManagerFromPointer(cthis)
}

// /usr/include/qt/QtGui/qsessionmanager.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QSessionManager) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qsessionmanager.h:65
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sessionId() const

/*
Returns the identifier of the current session.

If the application has been restored from an earlier session, this identifier is the same as it was in the earlier session.

See also sessionKey() and QGuiApplication::sessionId().
*/
func (this *QSessionManager) SessionId() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager9sessionIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qsessionmanager.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sessionKey() const

/*
Returns the session key in the current session.

If the application has been restored from an earlier session, this key is the same as it was when the previous session ended.

The session key changes with every call of commitData() or saveState().

See also sessionId() and QGuiApplication::sessionKey().
*/
func (this *QSessionManager) SessionKey() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager10sessionKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qsessionmanager.h:68
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allowsInteraction()

/*
Asks the session manager for permission to interact with the user. Returns true if interaction is permitted; otherwise returns false.

The rationale behind this mechanism is to make it possible to synchronize user interaction during a shutdown. Advanced session managers may ask all applications simultaneously to commit their data, resulting in a much faster shutdown.

When the interaction is completed we strongly recommend releasing the user interaction semaphore with a call to release(). This way, other applications may get the chance to interact with the user while your application is still busy saving data. (The semaphore is implicitly released when the application exits.)

If the user decides to cancel the shutdown process during the interaction phase, you must tell the session manager that this has happened by calling cancel().

Here's an example of how an application's QGuiApplication::commitDataRequest() might be implemented:


  MyMainWidget::MyMainWidget(QWidget *parent)
      :QWidget(parent)
  {
      QGuiApplication::setFallbackSessionManagementEnabled(false);
      connect(qApp, SIGNAL(commitDataRequest(QSessionManager)), SLOT(commitData(QSessionManager)));
  }

  void MyMainWidget::commitData(QSessionManager& manager)
  {
      if (manager.allowsInteraction()) {
          int ret = QMessageBox::warning(
                      mainWindow,
                      tr("My Application"),
                      tr("Save changes to document?"),
                      QMessageBox::Save | QMessageBox::Discard | QMessageBox::Cancel);

          switch (ret) {
          case QMessageBox::Save:
              manager.release();
              if (!saveDocument())
                  manager.cancel();
              break;
          case QMessageBox::Discard:
              break;
          case QMessageBox::Cancel:
          default:
              manager.cancel();
          }
      } else {
          // we did not get permission to interact, then
          // do something reasonable instead
      }
  }



If an error occurred within the application while saving its data, you may want to try allowsErrorInteraction() instead.

See also QGuiApplication::commitDataRequest(), release(), and cancel().
*/
func (this *QSessionManager) AllowsInteraction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager17allowsInteractionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsessionmanager.h:69
// index:0
// Public Visibility=Default Availability=Available
// [1] bool allowsErrorInteraction()

/*
Returns true if error interaction is permitted; otherwise returns false.

This is similar to allowsInteraction(), but also enables the application to tell the user about any errors that occur. Session managers may give error interaction requests higher priority, which means that it is more likely that an error interaction is permitted. However, you are still not guaranteed that the session manager will allow interaction.

See also allowsInteraction(), release(), and cancel().
*/
func (this *QSessionManager) AllowsErrorInteraction() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager22allowsErrorInteractionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsessionmanager.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void release()

/*
Releases the session manager's interaction semaphore after an interaction phase.

See also allowsInteraction() and allowsErrorInteraction().
*/
func (this *QSessionManager) Release() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager7releaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancel()

/*
Tells the session manager to cancel the shutdown process. Applications should not call this function without asking the user first.

See also allowsInteraction() and allowsErrorInteraction().
*/
func (this *QSessionManager) Cancel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager6cancelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRestartHint(QSessionManager::RestartHint)

/*
Sets the application's restart hint to hint. On application startup, the hint is set to RestartIfRunning.

Note: These flags are only hints, a session manager may or may not respect them.

We recommend setting the restart hint in QGuiApplication::saveStateRequest() because most session managers perform a checkpoint shortly after an application's startup.

See also restartHint().
*/
func (this *QSessionManager) SetRestartHint(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager14setRestartHintENS_11RestartHintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] QSessionManager::RestartHint restartHint() const

/*
Returns the application's current restart hint. The default is RestartIfRunning.

See also setRestartHint().
*/
func (this *QSessionManager) RestartHint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager11restartHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRestartCommand(const QStringList &)

/*
If the session manager is capable of restoring sessions it will execute command in order to restore the application. The command defaults to


  appname -session id



The -session option is mandatory; otherwise QGuiApplication cannot tell whether it has been restored or what the current session identifier is. See QGuiApplication::isSessionRestored() and QGuiApplication::sessionId() for details.

If your application is very simple, it may be possible to store the entire application state in additional command line options. This is usually a very bad idea because command lines are often limited to a few hundred bytes. Instead, use QSettings, temporary files, or a database for this purpose. By marking the data with the unique sessionId(), you will be able to restore the application in a future session.

See also restartCommand(), setDiscardCommand(), and setRestartHint().
*/
func (this *QSessionManager) SetRestartCommand(arg0 qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStringList_PTR() != nil {
		convArg0 = arg0.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager17setRestartCommandERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:84
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList restartCommand() const

/*
Returns the currently set restart command.

To iterate over the list, you can use the foreach pseudo-keyword:


  foreach (const QString &command, mySession.restartCommand())
      do_something(command);



See also setRestartCommand() and restartHint().
*/
func (this *QSessionManager) RestartCommand() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager14restartCommandEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qsessionmanager.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDiscardCommand(const QStringList &)

/*
Sets the discard command to the given command.

See also discardCommand() and setRestartCommand().
*/
func (this *QSessionManager) SetDiscardCommand(arg0 qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStringList_PTR() != nil {
		convArg0 = arg0.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager17setDiscardCommandERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList discardCommand() const

/*
Returns the currently set discard command.

To iterate over the list, you can use the foreach pseudo-keyword:


  foreach (const QString &command, mySession.discardCommand())
      do_something(command);



See also setDiscardCommand(), restartCommand(), and setRestartCommand().
*/
func (this *QSessionManager) DiscardCommand() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager14discardCommandEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qsessionmanager.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setManagerProperty(const QString &, const QString &)

/*
Low-level write access to the application's identification and state record are kept in the session manager.

The property called name has its value set to the string list value.
*/
func (this *QSessionManager) SetManagerProperty(name string, value string) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(value)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setManagerProperty(const QString &, const QStringList &)

/*
Low-level write access to the application's identification and state record are kept in the session manager.

The property called name has its value set to the string list value.
*/
func (this *QSessionManager) SetManagerProperty_1(name string, value qtcore.QStringList_ITF) {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if value != nil && value.QStringList_PTR() != nil {
		convArg1 = value.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager18setManagerPropertyERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qsessionmanager.h:91
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPhase2() const

/*
Returns true if the session manager is currently performing a second session management phase; otherwise returns false.

See also requestPhase2().
*/
func (this *QSessionManager) IsPhase2() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSessionManager8isPhase2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsessionmanager.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestPhase2()

/*
Requests a second session management phase for the application. The application may then return immediately from the QGuiApplication::commitDataRequest() or QApplication::saveStateRequest() function, and they will be called again once most or all other applications have finished their session management.

The two phases are useful for applications such as the X11 window manager that need to store information about another application's windows and therefore have to wait until these applications have completed their respective session management tasks.

Note: If another application has requested a second phase it may get called before, simultaneously with, or after your application's second phase.

See also isPhase2().
*/
func (this *QSessionManager) RequestPhase2() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManager13requestPhase2Ev", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

func DeleteQSessionManager(this *QSessionManager) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSessionManagerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
This enum type defines the circumstances under which this application wants to be restarted by the session manager. The current values are:



The default hint is RestartIfRunning.

*/
type QSessionManager__RestartHint = int

// If the application is still running when the session is shut down, it wants to be restarted at the start of the next session.
const QSessionManager__RestartIfRunning QSessionManager__RestartHint = 0

// The application wants to be started at the start of the next session, no matter what. (This is useful for utilities that run just after startup and then quit.)
const QSessionManager__RestartAnyway QSessionManager__RestartHint = 1

// The application wants to be started immediately whenever it is not running.
const QSessionManager__RestartImmediately QSessionManager__RestartHint = 2

// The application does not want to be restarted automatically.
const QSessionManager__RestartNever QSessionManager__RestartHint = 3

func (this *QSessionManager) RestartHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QSessionManager_RestartHintItemName(val int) string {
	var nilthis *QSessionManager
	return nilthis.RestartHintItemName(val)
}

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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
