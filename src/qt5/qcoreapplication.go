package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZN16QCoreApplication18applicationDirPathEv(); // 4
  // proto: static qint64 QCoreApplication::applicationPid();
extern int64_t C_ZN16QCoreApplication14applicationPidEv(); // 4
  // proto:  bool QCoreApplication::notify(QObject * , QEvent * );
extern bool C_ZN16QCoreApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QCoreApplication::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void C_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto: static void QCoreApplication::flush();
extern void C_ZN16QCoreApplication5flushEv(); // 4
  // proto: static bool QCoreApplication::isQuitLockEnabled();
extern bool C_ZN16QCoreApplication17isQuitLockEnabledEv(); // 4
  // proto: static bool QCoreApplication::closingDown();
extern bool C_ZN16QCoreApplication11closingDownEv(); // 4
  // proto: static QString QCoreApplication::applicationName();
extern void* C_ZN16QCoreApplication15applicationNameEv(); // 4
  // proto: static void QCoreApplication::quit();
extern void C_ZN16QCoreApplication4quitEv(); // 4
  // proto: static QString QCoreApplication::applicationVersion();
extern void* C_ZN16QCoreApplication18applicationVersionEv(); // 4
  // proto: static int QCoreApplication::exec();
extern int32_t C_ZN16QCoreApplication4execEv(); // 4
  // proto: static QCoreApplication * QCoreApplication::instance();
extern void* C_ZN16QCoreApplication8instanceEv(); // 2
  // proto: static QString QCoreApplication::organizationDomain();
extern void* C_ZN16QCoreApplication18organizationDomainEv(); // 4
  // proto: static void QCoreApplication::exit(int retcode);
extern void C_ZN16QCoreApplication4exitEi(int32_t arg0); // 4
  // proto: static QStringList QCoreApplication::arguments();
extern void C_ZN16QCoreApplication9argumentsEv(); // 4
  // proto: static QString QCoreApplication::applicationFilePath();
extern void* C_ZN16QCoreApplication19applicationFilePathEv(); // 4
  // proto: static void QCoreApplication::setSetuidAllowed(bool allow);
extern void C_ZN16QCoreApplication16setSetuidAllowedEb(bool arg0); // 4
  // proto: static QString QCoreApplication::translate(const char * context, const char * key, const char * disambiguation, int n);
extern void* C_ZN16QCoreApplication9translateEPKcS1_S1_i(void* arg0, void* arg1, void* arg2, int32_t arg3); // 4
  // proto: static void QCoreApplication::setQuitLockEnabled(bool enabled);
extern void C_ZN16QCoreApplication18setQuitLockEnabledEb(bool arg0); // 4
  // proto: static void QCoreApplication::setApplicationVersion(const QString & version);
extern void C_ZN16QCoreApplication21setApplicationVersionERK7QString(void* arg0); // 4
  // proto:  void QCoreApplication::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void C_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0); // 4
  // proto: static bool QCoreApplication::installTranslator(QTranslator * messageFile);
extern bool C_ZN16QCoreApplication17installTranslatorEP11QTranslator(void* arg0); // 4
  // proto:  void QCoreApplication::QCoreApplication(int & argc, char ** argv, int );
extern void* C_ZN16QCoreApplicationC2ERiPPci(void* arg0, void* arg1, int32_t arg2); // 3
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
extern bool C_ZN16QCoreApplication16removeTranslatorEP11QTranslator(void* arg0); // 4
  // proto: static QStringList QCoreApplication::libraryPaths();
extern void C_ZN16QCoreApplication12libraryPathsEv(); // 4
  // proto: static QString QCoreApplication::organizationName();
extern void* C_ZN16QCoreApplication16organizationNameEv(); // 4
  // proto: static void QCoreApplication::sendPostedEvents(QObject * receiver, int event_type);
extern void C_ZN16QCoreApplication16sendPostedEventsEP7QObjecti(void* arg0, int32_t arg1); // 4
  // proto:  void QCoreApplication::~QCoreApplication();
extern void C_ZN16QCoreApplicationD2Ev(void* qthis); // 4
  // proto: static bool QCoreApplication::isSetuidAllowed();
extern bool C_ZN16QCoreApplication15isSetuidAllowedEv(); // 4
  // proto: static bool QCoreApplication::sendEvent(QObject * receiver, QEvent * event);
extern bool C_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent(void* arg0, void* arg1); // 2
  // proto: static bool QCoreApplication::startingUp();
extern bool C_ZN16QCoreApplication10startingUpEv(); // 4
  // proto: static bool QCoreApplication::hasPendingEvents();
extern bool C_ZN16QCoreApplication16hasPendingEventsEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToQuit QCoreApplication_aboutToQuit_signal;
//  _applicationVersionChanged QCoreApplication_applicationVersionChanged_signal;
//  _organizationDomainChanged QCoreApplication_organizationDomainChanged_signal;
//  _applicationNameChanged QCoreApplication_applicationNameChanged_signal;
//  _organizationNameChanged QCoreApplication_organizationNameChanged_signal;
}

// setLibraryPaths(const class QStringList &)
func (this *QCoreApplication) Setlibrarypaths_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication15setLibraryPathsERK11QStringList(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setLibraryPaths", args)
  }

  return
}

// eventDispatcher()
func (this *QCoreApplication) Eventdispatcher_S(args ...interface{}) () {
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

  return
}

// setOrganizationName(const class QString &)
func (this *QCoreApplication) Setorganizationname_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication19setOrganizationNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setOrganizationName", args)
  }

  return
}

// setEventDispatcher(class QAbstractEventDispatcher *)
func (this *QCoreApplication) Seteventdispatcher_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractEventDispatcher).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setEventDispatcher", args)
  }

  return
}

// applicationDirPath()
func (this *QCoreApplication) Applicationdirpath_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication18applicationDirPathEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationDirPath", args)
  }

  return
}

// applicationPid()
func (this *QCoreApplication) Applicationpid_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication14applicationPidEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationPid", args)
  }

  return
}

// notify(class QObject *, class QEvent *)
func (this *QCoreApplication) Notify(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QEvent).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN16QCoreApplication6notifyEP7QObjectP6QEvent(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "notify", args)
  }

  return
}

// installNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) Installnativeeventfilter(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractNativeEventFilter).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "installNativeEventFilter", args)
  }

  return
}

// flush()
func (this *QCoreApplication) Flush_S(args ...interface{}) () {
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

  return
}

// isQuitLockEnabled()
func (this *QCoreApplication) Isquitlockenabled_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication17isQuitLockEnabledEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "isQuitLockEnabled", args)
  }

  return
}

// closingDown()
func (this *QCoreApplication) Closingdown_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication11closingDownEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "closingDown", args)
  }

  return
}

// applicationName()
func (this *QCoreApplication) Applicationname_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication15applicationNameEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationName", args)
  }

  return
}

// quit()
func (this *QCoreApplication) Quit_S(args ...interface{}) () {
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

  return
}

// applicationVersion()
func (this *QCoreApplication) Applicationversion_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication18applicationVersionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationVersion", args)
  }

  return
}

// exec()
func (this *QCoreApplication) Exec_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication4execEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "exec", args)
  }

  return
}

// instance()
func (this *QCoreApplication) Instance_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication8instanceEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCoreApplication{}) // "QCoreApplication *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "instance", args)
  }

  return
}

// organizationDomain()
func (this *QCoreApplication) Organizationdomain_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication18organizationDomainEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "organizationDomain", args)
  }

  return
}

// exit(int)
func (this *QCoreApplication) Exit_S(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication4exitEi(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "exit", args)
  }

  return
}

// arguments()
func (this *QCoreApplication) Arguments_S(args ...interface{}) () {
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

  return
}

// applicationFilePath()
func (this *QCoreApplication) Applicationfilepath_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication19applicationFilePathEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationFilePath", args)
  }

  return
}

// setSetuidAllowed(_Bool)
func (this *QCoreApplication) Setsetuidallowed_S(args ...interface{}) () {
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

  return
}

// translate(const char *, const char *, const char *, int)
func (this *QCoreApplication) Translate_S(args ...interface{}) (ret interface{}) {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN16QCoreApplication9translateEPKcS1_S1_i(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "translate", args)
  }

  return
}

// setQuitLockEnabled(_Bool)
func (this *QCoreApplication) Setquitlockenabled_S(args ...interface{}) () {
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

  return
}

// setApplicationVersion(const class QString &)
func (this *QCoreApplication) Setapplicationversion_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication21setApplicationVersionERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setApplicationVersion", args)
  }

  return
}

// removeNativeEventFilter(class QAbstractNativeEventFilter *)
func (this *QCoreApplication) Removenativeeventfilter(args ...interface{}) () {
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
    var arg0 = args[0].(*QAbstractNativeEventFilter).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeNativeEventFilter", args)
  }

  return
}

// installTranslator(class QTranslator *)
func (this *QCoreApplication) Installtranslator_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTranslator).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN16QCoreApplication17installTranslatorEP11QTranslator(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "installTranslator", args)
  }

  return
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
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QCoreApplicationC2ERiPPci(arg0, arg1, arg2)
    return &QCoreApplication{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QCoreApplication", "QCoreApplication", args)
  }

  return nil // QCoreApplication{Qclsinst:qthis}
}

// setOrganizationDomain(const class QString &)
func (this *QCoreApplication) Setorganizationdomain_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication21setOrganizationDomainERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setOrganizationDomain", args)
  }

  return
}

// removePostedEvents(class QObject *, int)
func (this *QCoreApplication) Removepostedevents_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN16QCoreApplication18removePostedEventsEP7QObjecti(arg0, arg1)
  default:
    qtrt.ErrorResolve("QCoreApplication", "removePostedEvents", args)
  }

  return
}

// metaObject()
func (this *QCoreApplication) Metaobject(args ...interface{}) () {
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
    C.C_ZNK16QCoreApplication10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCoreApplication", "metaObject", args)
  }

  return
}

// setApplicationName(const class QString &)
func (this *QCoreApplication) Setapplicationname_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication18setApplicationNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "setApplicationName", args)
  }

  return
}

// addLibraryPath(const class QString &)
func (this *QCoreApplication) Addlibrarypath_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication14addLibraryPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "addLibraryPath", args)
  }

  return
}

// removeTranslator(class QTranslator *)
func (this *QCoreApplication) Removetranslator_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTranslator).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN16QCoreApplication16removeTranslatorEP11QTranslator(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeTranslator", args)
  }

  return
}

// libraryPaths()
func (this *QCoreApplication) Librarypaths_S(args ...interface{}) () {
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

  return
}

// organizationName()
func (this *QCoreApplication) Organizationname_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication16organizationNameEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "organizationName", args)
  }

  return
}

// sendPostedEvents(class QObject *, int)
func (this *QCoreApplication) Sendpostedevents_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN16QCoreApplication16sendPostedEventsEP7QObjecti(arg0, arg1)
  default:
    qtrt.ErrorResolve("QCoreApplication", "sendPostedEvents", args)
  }

  return
}

// ~QCoreApplication()
func (this *QCoreApplication) Freeqcoreapplication(args ...interface{}) () {
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
    C.C_ZN16QCoreApplicationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCoreApplication", "~QCoreApplication", args)
  }

  return
}

// isSetuidAllowed()
func (this *QCoreApplication) Issetuidallowed_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication15isSetuidAllowedEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "isSetuidAllowed", args)
  }

  return
}

// sendEvent(class QObject *, class QEvent *)
func (this *QCoreApplication) Sendevent_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QEvent).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN16QCoreApplication9sendEventEP7QObjectP6QEvent(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "sendEvent", args)
  }

  return
}

// startingUp()
func (this *QCoreApplication) Startingup_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication10startingUpEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "startingUp", args)
  }

  return
}

// hasPendingEvents()
func (this *QCoreApplication) Haspendingevents_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN16QCoreApplication16hasPendingEventsEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCoreApplication", "hasPendingEvents", args)
  }

  return
}

// postEvent(class QObject *, class QEvent *, int)
func (this *QCoreApplication) Postevent_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QEvent).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN16QCoreApplication9postEventEP7QObjectP6QEventi(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QCoreApplication", "postEvent", args)
  }

  return
}

// removeLibraryPath(const class QString &)
func (this *QCoreApplication) Removelibrarypath_S(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QCoreApplication17removeLibraryPathERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeLibraryPath", args)
  }

  return
}

// <= body block end

