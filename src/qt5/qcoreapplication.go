package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qcoreapplication.h
// dst-file: /src/core/qcoreapplication.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static void QCoreApplication::setLibraryPaths(const QStringList & );
extern void C_ZN16QCoreApplication15setLibraryPathsERK11QStringList(void* arg0); // 4
  // proto: static QAbstractEventDispatcher * QCoreApplication::eventDispatcher();
extern void C_ZN16QCoreApplication15eventDispatcherEv(); // 4
  // proto: static void QCoreApplication::setOrganizationName(const QString & orgName);
extern void C_ZN16QCoreApplication19setOrganizationNameERK7QString(void* arg0); // 4
  // proto: static void QCoreApplication::setEventDispatcher(QAbstractEventDispatcher * eventDispatcher);
extern void C_ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher(void* arg0); // 4
  // proto: static QString QCoreApplication::applicationDirPath();
extern void C_ZN16QCoreApplication18applicationDirPathEv(); // 4
  // proto: static qint64 QCoreApplication::applicationPid();
extern void C_ZN16QCoreApplication14applicationPidEv(); // 4
  // proto:  bool QCoreApplication::notify(QObject * , QEvent * );
extern void C_ZN16QCoreApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QCoreApplication::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void C_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto: static void QCoreApplication::flush();
extern void C_ZN16QCoreApplication5flushEv(); // 4
  // proto: static bool QCoreApplication::isQuitLockEnabled();
extern void C_ZN16QCoreApplication17isQuitLockEnabledEv(); // 4
  // proto: static bool QCoreApplication::closingDown();
extern void C_ZN16QCoreApplication11closingDownEv(); // 4
  // proto: static QString QCoreApplication::applicationName();
extern void C_ZN16QCoreApplication15applicationNameEv(); // 4
  // proto: static void QCoreApplication::quit();
extern void C_ZN16QCoreApplication4quitEv(); // 4
  // proto: static QString QCoreApplication::applicationVersion();
extern void C_ZN16QCoreApplication18applicationVersionEv(); // 4
  // proto: static int QCoreApplication::exec();
extern void C_ZN16QCoreApplication4execEv(); // 4
  // proto: static QCoreApplication * QCoreApplication::instance();
extern void C_ZN16QCoreApplication8instanceEv(); // 2
  // proto: static QString QCoreApplication::organizationDomain();
extern void C_ZN16QCoreApplication18organizationDomainEv(); // 4
  // proto: static void QCoreApplication::exit(int retcode);
extern void C_ZN16QCoreApplication4exitEi(int32_t arg0); // 4
  // proto: static QStringList QCoreApplication::arguments();
extern void C_ZN16QCoreApplication9argumentsEv(); // 4
  // proto: static QString QCoreApplication::applicationFilePath();
extern void C_ZN16QCoreApplication19applicationFilePathEv(); // 4
  // proto: static void QCoreApplication::setSetuidAllowed(bool allow);
extern void C_ZN16QCoreApplication16setSetuidAllowedEb(bool arg0); // 4
  // proto: static QString QCoreApplication::translate(const char * context, const char * key, const char * disambiguation, int n);
extern void C_ZN16QCoreApplication9translateEPKcS1_S1_i(unsigned char* arg0, unsigned char* arg1, unsigned char* arg2, int32_t arg3); // 4
  // proto: static void QCoreApplication::setQuitLockEnabled(bool enabled);
extern void C_ZN16QCoreApplication18setQuitLockEnabledEb(bool arg0); // 4
  // proto: static void QCoreApplication::setApplicationVersion(const QString & version);
extern void C_ZN16QCoreApplication21setApplicationVersionERK7QString(void* arg0); // 4
  // proto:  void QCoreApplication::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void C_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto: static bool QCoreApplication::installTranslator(QTranslator * messageFile);
extern void C_ZN16QCoreApplication17installTranslatorEP11QTranslator(void* arg0); // 4
  // proto:  void QCoreApplication::QCoreApplication(int & argc, char ** argv, int );
extern void* C_ZN16QCoreApplicationC2ERiPPci(int32_t* arg0, unsigned char* arg1, int32_t arg2); // 3
  // proto: static void QCoreApplication::setOrganizationDomain(const QString & orgDomain);
extern void C_ZN16QCoreApplication21setOrganizationDomainERK7QString(void* arg0); // 4
  // proto: static void QCoreApplication::removePostedEvents(QObject * receiver, int eventType);
extern void C_ZN16QCoreApplication18removePostedEventsEP7QObjecti(void* arg0, int32_t arg1); // 4
  // proto:  const QMetaObject * QCoreApplication::metaObject();
extern void C_ZNK16QCoreApplication10metaObjectEv(void* qthis); // 4
  // proto: static void QCoreApplication::setApplicationName(const QString & application);
extern void C_ZN16QCoreApplication18setApplicationNameERK7QString(void* arg0); // 4
  // proto: static void QCoreApplication::addLibraryPath(const QString & );
extern void C_ZN16QCoreApplication14addLibraryPathERK7QString(void* arg0); // 4
  // proto: static bool QCoreApplication::removeTranslator(QTranslator * messageFile);
extern void C_ZN16QCoreApplication16removeTranslatorEP11QTranslator(void* arg0); // 4
  // proto: static QStringList QCoreApplication::libraryPaths();
extern void C_ZN16QCoreApplication12libraryPathsEv(); // 4
  // proto: static QString QCoreApplication::organizationName();
extern void C_ZN16QCoreApplication16organizationNameEv(); // 4
  // proto: static void QCoreApplication::sendPostedEvents(QObject * receiver, int event_type);
extern void C_ZN16QCoreApplication16sendPostedEventsEP7QObjecti(void* arg0, int32_t arg1); // 4
  // proto:  void QCoreApplication::~QCoreApplication();
extern void C_ZN16QCoreApplicationD2Ev(void* qthis); // 4
  // proto: static bool QCoreApplication::isSetuidAllowed();
extern void C_ZN16QCoreApplication15isSetuidAllowedEv(); // 4
  // proto: static bool QCoreApplication::sendEvent(QObject * receiver, QEvent * event);
extern void C_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent(void* arg0, void* arg1); // 2
  // proto: static bool QCoreApplication::startingUp();
extern void C_ZN16QCoreApplication10startingUpEv(); // 4
  // proto: static bool QCoreApplication::hasPendingEvents();
extern void C_ZN16QCoreApplication16hasPendingEventsEv(); // 4
  // proto: static void QCoreApplication::postEvent(QObject * receiver, QEvent * event, int priority);
extern void C_ZN16QCoreApplication9postEventEP7QObjectP6QEventi(void* arg0, void* arg1, int32_t arg2); // 4
  // proto: static void QCoreApplication::removeLibraryPath(const QString & );
extern void C_ZN16QCoreApplication17removeLibraryPathERK7QString(void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QCoreApplication)=1
type QCoreApplication struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToQuit QCoreApplication_aboutToQuit_signal;
//  _applicationVersionChanged QCoreApplication_applicationVersionChanged_signal;
//  _organizationDomainChanged QCoreApplication_organizationDomainChanged_signal;
//  _applicationNameChanged QCoreApplication_applicationNameChanged_signal;
//  _organizationNameChanged QCoreApplication_organizationNameChanged_signal;
}

// setLibraryPaths(const class QStringList &)
func (this *QCoreApplication) setLibraryPaths_s(args ...interface{}) () {
  // setLibraryPaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication15setLibraryPathsERK11QStringList
    // invoke: void setLibraryPaths(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication15setLibraryPathsERK11QStringList(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setLibraryPaths", args)
  }

}

// eventDispatcher()
func (this *QCoreApplication) eventDispatcher_s(args ...interface{}) () {
  // eventDispatcher()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication15eventDispatcherEv
    // invoke: QAbstractEventDispatcher * eventDispatcher()
    C.C_ZN16QCoreApplication15eventDispatcherEv()
  default:
    qtrt.ErrorResolve("QCoreApplication", "eventDispatcher", args)
  }

}

// setOrganizationName(const class QString &)
func (this *QCoreApplication) setOrganizationName_s(args ...interface{}) () {
  // setOrganizationName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication19setOrganizationNameERK7QString
    // invoke: void setOrganizationName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication19setOrganizationNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setOrganizationName", args)
  }

}

// setEventDispatcher(class QAbstractEventDispatcher *)
func (this *QCoreApplication) setEventDispatcher_s(args ...interface{}) () {
  // setEventDispatcher(class QAbstractEventDispatcher *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractEventDispatcher{}) // "QAbstractEventDispatcher *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher
    // invoke: void setEventDispatcher(class QAbstractEventDispatcher *)
    var arg0 = args[0].(QAbstractEventDispatcher).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setEventDispatcher", args)
  }

}

// applicationDirPath()
func (this *QCoreApplication) applicationDirPath_s(args ...interface{}) () {
  // applicationDirPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication18applicationDirPathEv
    // invoke: QString applicationDirPath()
    var ret = C.C_ZN16QCoreApplication18applicationDirPathEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationDirPath", args)
  }

}

// applicationPid()
func (this *QCoreApplication) applicationPid_s(args ...interface{}) () {
  // applicationPid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication14applicationPidEv
    // invoke: qint64 applicationPid()
    var ret = C.C_ZN16QCoreApplication14applicationPidEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationPid", args)
  }

}

// notify(class QObject *, class QEvent *)
func (this *QCoreApplication) notify(args ...interface{}) () {
  // notify(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication6notifyEP7QObjectP6QEvent
    // invoke: bool notify(class QObject *, class QEvent *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN16QCoreApplication6notifyEP7QObjectP6QEvent(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "notify", args)
  }

}

// installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) installNativeEventFilter(args ...interface{}) () {
  // installNativeEventFilter(class QAbstractNativeEventFilter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractNativeEventFilter{}) // "QAbstractNativeEventFilter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter
    // invoke: void installNativeEventFilter(class QAbstractNativeEventFilter *)
    var arg0 = args[0].(QAbstractNativeEventFilter).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "installNativeEventFilter", args)
  }

}

// flush()
func (this *QCoreApplication) flush_s(args ...interface{}) () {
  // flush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication5flushEv
    // invoke: void flush()
    C.C_ZN16QCoreApplication5flushEv()
  default:
    qtrt.ErrorResolve("QCoreApplication", "flush", args)
  }

}

// isQuitLockEnabled()
func (this *QCoreApplication) isQuitLockEnabled_s(args ...interface{}) () {
  // isQuitLockEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication17isQuitLockEnabledEv
    // invoke: bool isQuitLockEnabled()
    var ret = C.C_ZN16QCoreApplication17isQuitLockEnabledEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "isQuitLockEnabled", args)
  }

}

// closingDown()
func (this *QCoreApplication) closingDown_s(args ...interface{}) () {
  // closingDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication11closingDownEv
    // invoke: bool closingDown()
    var ret = C.C_ZN16QCoreApplication11closingDownEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "closingDown", args)
  }

}

// applicationName()
func (this *QCoreApplication) applicationName_s(args ...interface{}) () {
  // applicationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication15applicationNameEv
    // invoke: QString applicationName()
    var ret = C.C_ZN16QCoreApplication15applicationNameEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationName", args)
  }

}

// quit()
func (this *QCoreApplication) quit_s(args ...interface{}) () {
  // quit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication4quitEv
    // invoke: void quit()
    C.C_ZN16QCoreApplication4quitEv()
  default:
    qtrt.ErrorResolve("QCoreApplication", "quit", args)
  }

}

// applicationVersion()
func (this *QCoreApplication) applicationVersion_s(args ...interface{}) () {
  // applicationVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication18applicationVersionEv
    // invoke: QString applicationVersion()
    var ret = C.C_ZN16QCoreApplication18applicationVersionEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationVersion", args)
  }

}

// exec()
func (this *QCoreApplication) exec_s(args ...interface{}) () {
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication4execEv
    // invoke: int exec()
    var ret = C.C_ZN16QCoreApplication4execEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "exec", args)
  }

}

// instance()
func (this *QCoreApplication) instance_s(args ...interface{}) () {
  // instance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication8instanceEv
    // invoke: QCoreApplication * instance()
    var ret = C.C_ZN16QCoreApplication8instanceEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "instance", args)
  }

}

// organizationDomain()
func (this *QCoreApplication) organizationDomain_s(args ...interface{}) () {
  // organizationDomain()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication18organizationDomainEv
    // invoke: QString organizationDomain()
    var ret = C.C_ZN16QCoreApplication18organizationDomainEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "organizationDomain", args)
  }

}

// exit(int)
func (this *QCoreApplication) exit_s(args ...interface{}) () {
  // exit(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication4exitEi
    // invoke: void exit(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication4exitEi(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "exit", args)
  }

}

// arguments()
func (this *QCoreApplication) arguments_s(args ...interface{}) () {
  // arguments()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication9argumentsEv
    // invoke: QStringList arguments()
    C.C_ZN16QCoreApplication9argumentsEv()
  default:
    qtrt.ErrorResolve("QCoreApplication", "arguments", args)
  }

}

// applicationFilePath()
func (this *QCoreApplication) applicationFilePath_s(args ...interface{}) () {
  // applicationFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication19applicationFilePathEv
    // invoke: QString applicationFilePath()
    var ret = C.C_ZN16QCoreApplication19applicationFilePathEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationFilePath", args)
  }

}

// setSetuidAllowed(_Bool)
func (this *QCoreApplication) setSetuidAllowed_s(args ...interface{}) () {
  // setSetuidAllowed(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication16setSetuidAllowedEb
    // invoke: void setSetuidAllowed(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication16setSetuidAllowedEb(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setSetuidAllowed", args)
  }

}

// translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) translate_s(args ...interface{}) () {
  // translate(const char *, const char *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication9translateEPKcS1_S1_i
    // invoke: QString translate(const char *, const char *, const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN16QCoreApplication9translateEPKcS1_S1_i(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "translate", args)
  }

}

// setQuitLockEnabled(_Bool)
func (this *QCoreApplication) setQuitLockEnabled_s(args ...interface{}) () {
  // setQuitLockEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication18setQuitLockEnabledEb
    // invoke: void setQuitLockEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication18setQuitLockEnabledEb(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setQuitLockEnabled", args)
  }

}

// setApplicationVersion(const class QString &)
func (this *QCoreApplication) setApplicationVersion_s(args ...interface{}) () {
  // setApplicationVersion(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication21setApplicationVersionERK7QString
    // invoke: void setApplicationVersion(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication21setApplicationVersionERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setApplicationVersion", args)
  }

}

// removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) removeNativeEventFilter(args ...interface{}) () {
  // removeNativeEventFilter(class QAbstractNativeEventFilter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractNativeEventFilter{}) // "QAbstractNativeEventFilter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter
    // invoke: void removeNativeEventFilter(class QAbstractNativeEventFilter *)
    var arg0 = args[0].(QAbstractNativeEventFilter).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeNativeEventFilter", args)
  }

}

// installTranslator(class QTranslator *)
func (this *QCoreApplication) installTranslator_s(args ...interface{}) () {
  // installTranslator(class QTranslator *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTranslator{}) // "QTranslator *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication17installTranslatorEP11QTranslator
    // invoke: bool installTranslator(class QTranslator *)
    var arg0 = args[0].(QTranslator).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN16QCoreApplication17installTranslatorEP11QTranslator(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "installTranslator", args)
  }

}

// QCoreApplication(int &, char **, int)
func NewQCoreApplication(args ...interface{}) *QCoreApplication {
  // QCoreApplication(int &, char **, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int &"
  vtys[0][1] = qtrt.ByteTy(true) // "char **"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplicationC1ERiPPci
    // invoke: void QCoreApplication(int &, char **, int)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QCoreApplicationC2ERiPPci(arg0, arg1, arg2)
    return &QCoreApplication{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCoreApplication", "QCoreApplication", args)
  }

  return nil // QCoreApplication{qclsinst:qthis}
}

// setOrganizationDomain(const class QString &)
func (this *QCoreApplication) setOrganizationDomain_s(args ...interface{}) () {
  // setOrganizationDomain(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication21setOrganizationDomainERK7QString
    // invoke: void setOrganizationDomain(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication21setOrganizationDomainERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setOrganizationDomain", args)
  }

}

// removePostedEvents(class QObject *, int)
func (this *QCoreApplication) removePostedEvents_s(args ...interface{}) () {
  // removePostedEvents(class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication18removePostedEventsEP7QObjecti
    // invoke: void removePostedEvents(class QObject *, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN16QCoreApplication18removePostedEventsEP7QObjecti(arg0, arg1)
  default:
    qtrt.ErrorResolve("QCoreApplication", "removePostedEvents", args)
  }

}

// metaObject()
func (this *QCoreApplication) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QCoreApplication10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK16QCoreApplication10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCoreApplication", "metaObject", args)
  }

}

// setApplicationName(const class QString &)
func (this *QCoreApplication) setApplicationName_s(args ...interface{}) () {
  // setApplicationName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication18setApplicationNameERK7QString
    // invoke: void setApplicationName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication18setApplicationNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setApplicationName", args)
  }

}

// addLibraryPath(const class QString &)
func (this *QCoreApplication) addLibraryPath_s(args ...interface{}) () {
  // addLibraryPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication14addLibraryPathERK7QString
    // invoke: void addLibraryPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication14addLibraryPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "addLibraryPath", args)
  }

}

// removeTranslator(class QTranslator *)
func (this *QCoreApplication) removeTranslator_s(args ...interface{}) () {
  // removeTranslator(class QTranslator *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTranslator{}) // "QTranslator *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication16removeTranslatorEP11QTranslator
    // invoke: bool removeTranslator(class QTranslator *)
    var arg0 = args[0].(QTranslator).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN16QCoreApplication16removeTranslatorEP11QTranslator(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeTranslator", args)
  }

}

// libraryPaths()
func (this *QCoreApplication) libraryPaths_s(args ...interface{}) () {
  // libraryPaths()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication12libraryPathsEv
    // invoke: QStringList libraryPaths()
    C.C_ZN16QCoreApplication12libraryPathsEv()
  default:
    qtrt.ErrorResolve("QCoreApplication", "libraryPaths", args)
  }

}

// organizationName()
func (this *QCoreApplication) organizationName_s(args ...interface{}) () {
  // organizationName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication16organizationNameEv
    // invoke: QString organizationName()
    var ret = C.C_ZN16QCoreApplication16organizationNameEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "organizationName", args)
  }

}

// sendPostedEvents(class QObject *, int)
func (this *QCoreApplication) sendPostedEvents_s(args ...interface{}) () {
  // sendPostedEvents(class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication16sendPostedEventsEP7QObjecti
    // invoke: void sendPostedEvents(class QObject *, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN16QCoreApplication16sendPostedEventsEP7QObjecti(arg0, arg1)
  default:
    qtrt.ErrorResolve("QCoreApplication", "sendPostedEvents", args)
  }

}

// ~QCoreApplication()
func (this *QCoreApplication) FreeQCoreApplication(args ...interface{}) () {
  // ~QCoreApplication()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplicationD0Ev
    // invoke: void ~QCoreApplication()
    C.C_ZN16QCoreApplicationD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCoreApplication", "~QCoreApplication", args)
  }

}

// isSetuidAllowed()
func (this *QCoreApplication) isSetuidAllowed_s(args ...interface{}) () {
  // isSetuidAllowed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication15isSetuidAllowedEv
    // invoke: bool isSetuidAllowed()
    var ret = C.C_ZN16QCoreApplication15isSetuidAllowedEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "isSetuidAllowed", args)
  }

}

// sendEvent(class QObject *, class QEvent *)
func (this *QCoreApplication) sendEvent_s(args ...interface{}) () {
  // sendEvent(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication9sendEventEP7QObjectP6QEvent
    // invoke: bool sendEvent(class QObject *, class QEvent *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent(arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "sendEvent", args)
  }

}

// startingUp()
func (this *QCoreApplication) startingUp_s(args ...interface{}) () {
  // startingUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication10startingUpEv
    // invoke: bool startingUp()
    var ret = C.C_ZN16QCoreApplication10startingUpEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "startingUp", args)
  }

}

// hasPendingEvents()
func (this *QCoreApplication) hasPendingEvents_s(args ...interface{}) () {
  // hasPendingEvents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication16hasPendingEventsEv
    // invoke: bool hasPendingEvents()
    var ret = C.C_ZN16QCoreApplication16hasPendingEventsEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QCoreApplication", "hasPendingEvents", args)
  }

}

// postEvent(class QObject *, class QEvent *, int)
func (this *QCoreApplication) postEvent_s(args ...interface{}) () {
  // postEvent(class QObject *, class QEvent *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication9postEventEP7QObjectP6QEventi
    // invoke: void postEvent(class QObject *, class QEvent *, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN16QCoreApplication9postEventEP7QObjectP6QEventi(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QCoreApplication", "postEvent", args)
  }

}

// removeLibraryPath(const class QString &)
func (this *QCoreApplication) removeLibraryPath_s(args ...interface{}) () {
  // removeLibraryPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QCoreApplication17removeLibraryPathERK7QString
    // invoke: void removeLibraryPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication17removeLibraryPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeLibraryPath", args)
  }

}

// <= body block end

