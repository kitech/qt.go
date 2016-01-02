package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static void QCoreApplication::sendPostedEvents(QObject * receiver, int event_type);
extern void _ZN16QCoreApplication16sendPostedEventsEP7QObjecti(void* arg0, int arg1);
  // proto: static void QCoreApplication::addLibraryPath(const QString & );
extern void _ZN16QCoreApplication14addLibraryPathERK7QString(void* arg0);
  // proto: static qint64 QCoreApplication::applicationPid();
extern void _ZN16QCoreApplication14applicationPidEv();
  // proto: static void QCoreApplication::setApplicationName(const QString & application);
extern void _ZN16QCoreApplication18setApplicationNameERK7QString(void* arg0);
  // proto: static QString QCoreApplication::organizationName();
extern void _ZN16QCoreApplication16organizationNameEv();
  // proto:  void QCoreApplication::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void _ZN16QCoreApplication24installNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0);
  // proto: static QCoreApplication * QCoreApplication::instance();
extern void _ZN16QCoreApplication8instanceEv();
  // proto: static bool QCoreApplication::isSetuidAllowed();
extern void _ZN16QCoreApplication15isSetuidAllowedEv();
  // proto: static QString QCoreApplication::applicationName();
extern void _ZN16QCoreApplication15applicationNameEv();
  // proto:  void QCoreApplication::QCoreApplication(const QCoreApplication & );
extern void* dector_ZN16QCoreApplicationC1ERKS_(void* arg0);
extern void _ZN16QCoreApplicationC1ERKS_(void* qthis, void* arg0);
  // proto: static void QCoreApplication::setSetuidAllowed(bool allow);
extern void _ZN16QCoreApplication16setSetuidAllowedEb(bool arg0);
  // proto: static void QCoreApplication::postEvent(QObject * receiver, QEvent * event, int priority);
extern void _ZN16QCoreApplication9postEventEP7QObjectP6QEventi(void* arg0, void* arg1, int arg2);
  // proto: static QStringList QCoreApplication::libraryPaths();
extern void _ZN16QCoreApplication12libraryPathsEv();
  // proto: static void QCoreApplication::removeLibraryPath(const QString & );
extern void _ZN16QCoreApplication17removeLibraryPathERK7QString(void* arg0);
  // proto: static QString QCoreApplication::translate(const char * context, const char * key, const char * disambiguation, int n);
extern void _ZN16QCoreApplication9translateEPKcS1_S1_i(char* arg0, char* arg1, char* arg2, int arg3);
  // proto: static QString QCoreApplication::applicationFilePath();
extern void _ZN16QCoreApplication19applicationFilePathEv();
  // proto: static bool QCoreApplication::removeTranslator(QTranslator * messageFile);
extern void _ZN16QCoreApplication16removeTranslatorEP11QTranslator(void* arg0);
  // proto: static void QCoreApplication::setOrganizationName(const QString & orgName);
extern void _ZN16QCoreApplication19setOrganizationNameERK7QString(void* arg0);
  // proto: static void QCoreApplication::exit(int retcode);
extern void _ZN16QCoreApplication4exitEi(int arg0);
  // proto: static QString QCoreApplication::applicationVersion();
extern void _ZN16QCoreApplication18applicationVersionEv();
  // proto: static void QCoreApplication::quit();
extern void _ZN16QCoreApplication4quitEv();
  // proto: static bool QCoreApplication::closingDown();
extern void _ZN16QCoreApplication11closingDownEv();
  // proto: static void QCoreApplication::setQuitLockEnabled(bool enabled);
extern void _ZN16QCoreApplication18setQuitLockEnabledEb(bool arg0);
  // proto: static bool QCoreApplication::hasPendingEvents();
extern void _ZN16QCoreApplication16hasPendingEventsEv();
  // proto: static void QCoreApplication::setOrganizationDomain(const QString & orgDomain);
extern void _ZN16QCoreApplication21setOrganizationDomainERK7QString(void* arg0);
  // proto:  void QCoreApplication::~QCoreApplication();
extern void _ZN16QCoreApplicationD0Ev(void* qthis);
  // proto:  void QCoreApplication::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
extern void _ZN16QCoreApplication23removeNativeEventFilterEP26QAbstractNativeEventFilter(void* qthis, void* arg0);
  // proto: static QString QCoreApplication::organizationDomain();
extern void _ZN16QCoreApplication18organizationDomainEv();
  // proto: static bool QCoreApplication::installTranslator(QTranslator * messageFile);
extern void _ZN16QCoreApplication17installTranslatorEP11QTranslator(void* arg0);
  // proto: static QString QCoreApplication::applicationDirPath();
extern void _ZN16QCoreApplication18applicationDirPathEv();
  // proto: static void QCoreApplication::flush();
extern void _ZN16QCoreApplication5flushEv();
  // proto: static int QCoreApplication::exec();
extern void _ZN16QCoreApplication4execEv();
  // proto: static QStringList QCoreApplication::arguments();
extern void _ZN16QCoreApplication9argumentsEv();
  // proto: static void QCoreApplication::setLibraryPaths(const QStringList & );
extern void _ZN16QCoreApplication15setLibraryPathsERK11QStringList(void* arg0);
  // proto: static QAbstractEventDispatcher * QCoreApplication::eventDispatcher();
extern void _ZN16QCoreApplication15eventDispatcherEv();
  // proto: static bool QCoreApplication::startingUp();
extern void _ZN16QCoreApplication10startingUpEv();
  // proto: static bool QCoreApplication::sendEvent(QObject * receiver, QEvent * event);
extern void _ZN16QCoreApplication9sendEventEP7QObjectP6QEvent(void* arg0, void* arg1);
  // proto:  bool QCoreApplication::notify(QObject * , QEvent * );
extern void _ZN16QCoreApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1);
  // proto: static bool QCoreApplication::isQuitLockEnabled();
extern void _ZN16QCoreApplication17isQuitLockEnabledEv();
  // proto: static void QCoreApplication::removePostedEvents(QObject * receiver, int eventType);
extern void _ZN16QCoreApplication18removePostedEventsEP7QObjecti(void* arg0, int arg1);
  // proto:  const QMetaObject * QCoreApplication::metaObject();
extern void _ZNK16QCoreApplication10metaObjectEv(void* qthis);
  // proto:  void QCoreApplication::QCoreApplication(int & argc, char ** argv, int );
extern void* dector_ZN16QCoreApplicationC1ERiPPci(int* arg0, char* arg1, int arg2);
extern void _ZN16QCoreApplicationC1ERiPPci(void* qthis, int* arg0, char* arg1, int arg2);
  // proto: static void QCoreApplication::setApplicationVersion(const QString & version);
extern void _ZN16QCoreApplication21setApplicationVersionERK7QString(void* arg0);
  // proto: static void QCoreApplication::setEventDispatcher(QAbstractEventDispatcher * eventDispatcher);
extern void _ZN16QCoreApplication18setEventDispatcherEP24QAbstractEventDispatcher(void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToQuit QCoreApplication_aboutToQuit_signal;
//  _applicationVersionChanged QCoreApplication_applicationVersionChanged_signal;
//  _organizationDomainChanged QCoreApplication_organizationDomainChanged_signal;
//  _applicationNameChanged QCoreApplication_applicationNameChanged_signal;
//  _organizationNameChanged QCoreApplication_organizationNameChanged_signal;
}

  // proto: static void QCoreApplication::sendPostedEvents(QObject * receiver, int event_type);
func (this *QCoreApplication) sendPostedEvents_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "sendPostedEvents", args)
  }

}

  // proto: static void QCoreApplication::addLibraryPath(const QString & );
func (this *QCoreApplication) addLibraryPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "addLibraryPath", args)
  }

}

  // proto: static qint64 QCoreApplication::applicationPid();
func (this *QCoreApplication) applicationPid_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationPid", args)
  }

}

  // proto: static void QCoreApplication::setApplicationName(const QString & application);
func (this *QCoreApplication) setApplicationName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setApplicationName", args)
  }

}

  // proto: static QString QCoreApplication::organizationName();
func (this *QCoreApplication) organizationName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "organizationName", args)
  }

}

  // proto:  void QCoreApplication::installNativeEventFilter(QAbstractNativeEventFilter * filterObj);
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
  default:
    qtrt.ErrorResolve("QCoreApplication", "installNativeEventFilter", args)
  }

}

  // proto: static QCoreApplication * QCoreApplication::instance();
func (this *QCoreApplication) instance_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "instance", args)
  }

}

  // proto: static bool QCoreApplication::isSetuidAllowed();
func (this *QCoreApplication) isSetuidAllowed_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "isSetuidAllowed", args)
  }

}

  // proto: static QString QCoreApplication::applicationName();
func (this *QCoreApplication) applicationName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationName", args)
  }

}

  // proto:  void QCoreApplication::QCoreApplication(const QCoreApplication & );
func NewQCoreApplication(args ...interface{}) QCoreApplication {
  return QCoreApplication{}
}

  // proto: static void QCoreApplication::setSetuidAllowed(bool allow);
func (this *QCoreApplication) setSetuidAllowed_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setSetuidAllowed", args)
  }

}

  // proto: static void QCoreApplication::postEvent(QObject * receiver, QEvent * event, int priority);
func (this *QCoreApplication) postEvent_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "postEvent", args)
  }

}

  // proto: static QStringList QCoreApplication::libraryPaths();
func (this *QCoreApplication) libraryPaths_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "libraryPaths", args)
  }

}

  // proto: static void QCoreApplication::removeLibraryPath(const QString & );
func (this *QCoreApplication) removeLibraryPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeLibraryPath", args)
  }

}

  // proto: static QString QCoreApplication::translate(const char * context, const char * key, const char * disambiguation, int n);
func (this *QCoreApplication) translate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "translate", args)
  }

}

  // proto: static QString QCoreApplication::applicationFilePath();
func (this *QCoreApplication) applicationFilePath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationFilePath", args)
  }

}

  // proto: static bool QCoreApplication::removeTranslator(QTranslator * messageFile);
func (this *QCoreApplication) removeTranslator_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeTranslator", args)
  }

}

  // proto: static void QCoreApplication::setOrganizationName(const QString & orgName);
func (this *QCoreApplication) setOrganizationName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setOrganizationName", args)
  }

}

  // proto: static void QCoreApplication::exit(int retcode);
func (this *QCoreApplication) exit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "exit", args)
  }

}

  // proto: static QString QCoreApplication::applicationVersion();
func (this *QCoreApplication) applicationVersion_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationVersion", args)
  }

}

  // proto: static void QCoreApplication::quit();
func (this *QCoreApplication) quit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "quit", args)
  }

}

  // proto: static bool QCoreApplication::closingDown();
func (this *QCoreApplication) closingDown_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "closingDown", args)
  }

}

  // proto: static void QCoreApplication::setQuitLockEnabled(bool enabled);
func (this *QCoreApplication) setQuitLockEnabled_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setQuitLockEnabled", args)
  }

}

  // proto: static bool QCoreApplication::hasPendingEvents();
func (this *QCoreApplication) hasPendingEvents_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "hasPendingEvents", args)
  }

}

  // proto: static void QCoreApplication::setOrganizationDomain(const QString & orgDomain);
func (this *QCoreApplication) setOrganizationDomain_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setOrganizationDomain", args)
  }

}

  // proto:  void QCoreApplication::~QCoreApplication();
func (this *QCoreApplication) FreeQCoreApplication(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "~QCoreApplication", args)
  }

}

  // proto:  void QCoreApplication::removeNativeEventFilter(QAbstractNativeEventFilter * filterObj);
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
  default:
    qtrt.ErrorResolve("QCoreApplication", "removeNativeEventFilter", args)
  }

}

  // proto: static QString QCoreApplication::organizationDomain();
func (this *QCoreApplication) organizationDomain_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "organizationDomain", args)
  }

}

  // proto: static bool QCoreApplication::installTranslator(QTranslator * messageFile);
func (this *QCoreApplication) installTranslator_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "installTranslator", args)
  }

}

  // proto: static QString QCoreApplication::applicationDirPath();
func (this *QCoreApplication) applicationDirPath_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "applicationDirPath", args)
  }

}

  // proto: static void QCoreApplication::flush();
func (this *QCoreApplication) flush_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "flush", args)
  }

}

  // proto: static int QCoreApplication::exec();
func (this *QCoreApplication) exec_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "exec", args)
  }

}

  // proto: static QStringList QCoreApplication::arguments();
func (this *QCoreApplication) arguments_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "arguments", args)
  }

}

  // proto: static void QCoreApplication::setLibraryPaths(const QStringList & );
func (this *QCoreApplication) setLibraryPaths_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setLibraryPaths", args)
  }

}

  // proto: static QAbstractEventDispatcher * QCoreApplication::eventDispatcher();
func (this *QCoreApplication) eventDispatcher_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "eventDispatcher", args)
  }

}

  // proto: static bool QCoreApplication::startingUp();
func (this *QCoreApplication) startingUp_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "startingUp", args)
  }

}

  // proto: static bool QCoreApplication::sendEvent(QObject * receiver, QEvent * event);
func (this *QCoreApplication) sendEvent_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "sendEvent", args)
  }

}

  // proto:  bool QCoreApplication::notify(QObject * , QEvent * );
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
  default:
    qtrt.ErrorResolve("QCoreApplication", "notify", args)
  }

}

  // proto: static bool QCoreApplication::isQuitLockEnabled();
func (this *QCoreApplication) isQuitLockEnabled_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "isQuitLockEnabled", args)
  }

}

  // proto: static void QCoreApplication::removePostedEvents(QObject * receiver, int eventType);
func (this *QCoreApplication) removePostedEvents_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "removePostedEvents", args)
  }

}

  // proto:  const QMetaObject * QCoreApplication::metaObject();
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
  default:
    qtrt.ErrorResolve("QCoreApplication", "metaObject", args)
  }

}

  // proto: static void QCoreApplication::setApplicationVersion(const QString & version);
func (this *QCoreApplication) setApplicationVersion_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setApplicationVersion", args)
  }

}

  // proto: static void QCoreApplication::setEventDispatcher(QAbstractEventDispatcher * eventDispatcher);
func (this *QCoreApplication) setEventDispatcher_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCoreApplication", "setEventDispatcher", args)
  }

}

// <= body block end

