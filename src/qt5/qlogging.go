package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtCore/qlogging.h
// dst-file: /src/core/qlogging.go
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
  // proto:  void QMessageLogContext::QMessageLogContext(const char * fileName, int lineNumber, const char * functionName, const char * categoryName);
extern void* dector_ZN18QMessageLogContextC1EPKciS1_S1_(unsigned char* arg0, int32_t arg1, unsigned char* arg2, unsigned char* arg3);
extern void _ZN18QMessageLogContextC1EPKciS1_S1_(void* qthis, unsigned char* arg0, int32_t arg1, unsigned char* arg2, unsigned char* arg3);
  // proto:  void QMessageLogContext::copy(const QMessageLogContext & logContext);
extern void _ZN18QMessageLogContext4copyERKS_(void* qthis, void* arg0);
  // proto:  void QMessageLogContext::QMessageLogContext();
extern void* dector_ZN18QMessageLogContextC1Ev();
extern void _ZN18QMessageLogContextC1Ev(void* qthis);
  // proto:  void QMessageLogContext::QMessageLogContext(const QMessageLogContext & );
extern void* dector_ZN18QMessageLogContextC1ERKS_(void* arg0);
extern void _ZN18QMessageLogContextC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMessageLogger::info(const QLoggingCategory & cat, const char * msg);
extern void _ZNK14QMessageLogger4infoERK16QLoggingCategoryPKcz(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  void QMessageLogger::debug(const char * msg);
extern void _ZNK14QMessageLogger5debugEPKcz(void* qthis, unsigned char* arg0);
  // proto:  void QMessageLogger::warning(const QLoggingCategory & cat, const char * msg);
extern void _ZNK14QMessageLogger7warningERK16QLoggingCategoryPKcz(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  void QMessageLogger::fatal(const char * msg);
extern void _ZNK14QMessageLogger5fatalEPKcz(void* qthis, unsigned char* arg0);
  // proto:  QDebug QMessageLogger::info(const QLoggingCategory & cat);
extern void _ZNK14QMessageLogger4infoERK16QLoggingCategory(void* qthis, void* arg0);
  // proto:  void QMessageLogger::debug(const QLoggingCategory & cat, const char * msg);
extern void _ZNK14QMessageLogger5debugERK16QLoggingCategoryPKcz(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  void QMessageLogger::critical(const char * msg);
extern void _ZNK14QMessageLogger8criticalEPKcz(void* qthis, unsigned char* arg0);
  // proto:  void QMessageLogger::QMessageLogger();
extern void* dector_ZN14QMessageLoggerC1Ev();
extern void _ZN14QMessageLoggerC1Ev(void* qthis);
  // proto:  QDebug QMessageLogger::info();
extern void _ZNK14QMessageLogger4infoEv(void* qthis);
  // proto:  void QMessageLogger::warning(const char * msg);
extern void _ZNK14QMessageLogger7warningEPKcz(void* qthis, unsigned char* arg0);
  // proto:  QDebug QMessageLogger::critical(const QLoggingCategory & cat);
extern void _ZNK14QMessageLogger8criticalERK16QLoggingCategory(void* qthis, void* arg0);
  // proto:  QDebug QMessageLogger::critical();
extern void _ZNK14QMessageLogger8criticalEv(void* qthis);
  // proto:  void QMessageLogger::QMessageLogger(const char * file, int line, const char * function);
extern void* dector_ZN14QMessageLoggerC1EPKciS1_(unsigned char* arg0, int32_t arg1, unsigned char* arg2);
extern void _ZN14QMessageLoggerC1EPKciS1_(void* qthis, unsigned char* arg0, int32_t arg1, unsigned char* arg2);
  // proto:  QDebug QMessageLogger::debug();
extern void _ZNK14QMessageLogger5debugEv(void* qthis);
  // proto:  QDebug QMessageLogger::debug(const QLoggingCategory & cat);
extern void _ZNK14QMessageLogger5debugERK16QLoggingCategory(void* qthis, void* arg0);
  // proto:  void QMessageLogger::QMessageLogger(const QMessageLogger & );
extern void* dector_ZN14QMessageLoggerC1ERKS_(void* arg0);
extern void _ZN14QMessageLoggerC1ERKS_(void* qthis, void* arg0);
  // proto:  QDebug QMessageLogger::warning(const QLoggingCategory & cat);
extern void _ZNK14QMessageLogger7warningERK16QLoggingCategory(void* qthis, void* arg0);
  // proto:  void QMessageLogger::info(const char * msg);
extern void _ZNK14QMessageLogger4infoEPKcz(void* qthis, unsigned char* arg0);
  // proto:  void QMessageLogger::noDebug(const char * );
extern void demth_ZNK14QMessageLogger7noDebugEPKcz(void* qthis, unsigned char* arg0);
  // proto:  QDebug QMessageLogger::warning();
extern void _ZNK14QMessageLogger7warningEv(void* qthis);
  // proto:  QNoDebug QMessageLogger::noDebug();
extern void _ZNK14QMessageLogger7noDebugEv(void* qthis);
  // proto:  void QMessageLogger::critical(const QLoggingCategory & cat, const char * msg);
extern void _ZNK14QMessageLogger8criticalERK16QLoggingCategoryPKcz(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  void QMessageLogger::QMessageLogger(const char * file, int line, const char * function, const char * category);
extern void* dector_ZN14QMessageLoggerC1EPKciS1_S1_(unsigned char* arg0, int32_t arg1, unsigned char* arg2, unsigned char* arg3);
extern void _ZN14QMessageLoggerC1EPKciS1_S1_(void* qthis, unsigned char* arg0, int32_t arg1, unsigned char* arg2, unsigned char* arg3);
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

// class sizeof(QMessageLogContext)=32
type QMessageLogContext struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMessageLogger)=32
type QMessageLogger struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QMessageLogContext::QMessageLogContext(const char * fileName, int lineNumber, const char * functionName, const char * categoryName);
func NewQMessageLogContext(args ...interface{}) QMessageLogContext {
  return QMessageLogContext{}
}

  // proto:  void QMessageLogContext::copy(const QMessageLogContext & logContext);
func (this *QMessageLogContext) copy(args ...interface{}) () {
  // copy(const class QMessageLogContext &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMessageLogContext{}) // "const QMessageLogContext &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QMessageLogContext4copyERKS_
    // invoke: void copy(const class QMessageLogContext &)
    var arg0 = args[0].(QMessageLogContext).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QMessageLogContext4copyERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageLogContext", "copy", args)
  }

}

  // proto:  void QMessageLogger::info(const QLoggingCategory & cat, const char * msg);
func (this *QMessageLogger) info(args ...interface{}) () {
  // info(const class QLoggingCategory &, const char *, ...)
  // info(CategoryFunction)
  // info(const class QLoggingCategory &)
  // info(CategoryFunction, const char *, ...)
  // info()
  // info(const char *, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[3][1] = qtrt.ByteTy(true) // "const char *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger4infoERK16QLoggingCategoryPKcz
    // invoke: void info(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C._ZNK14QMessageLogger4infoERK16QLoggingCategoryPKcz(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK14QMessageLogger4infoERK16QLoggingCategory
    // invoke: QDebug info(const class QLoggingCategory &)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger4infoERK16QLoggingCategory(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QMessageLogger4infoEv
    // invoke: QDebug info()
    C._ZNK14QMessageLogger4infoEv(this.qclsinst)
  case 3:
    // invoke: _ZNK14QMessageLogger4infoEPKcz
    // invoke: void info(const char *, ...)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger4infoEPKcz(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageLogger", "info", args)
  }

}

  // proto:  void QMessageLogger::debug(const char * msg);
func (this *QMessageLogger) debug(args ...interface{}) () {
  // debug(const char *, ...)
  // debug(const class QLoggingCategory &, const char *, ...)
  // debug(CategoryFunction, const char *, ...)
  // debug()
  // debug(const class QLoggingCategory &)
  // debug(CategoryFunction)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.VoidpTy() // "CategoryFunction"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger5debugEPKcz
    // invoke: void debug(const char *, ...)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger5debugEPKcz(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QMessageLogger5debugERK16QLoggingCategoryPKcz
    // invoke: void debug(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C._ZNK14QMessageLogger5debugERK16QLoggingCategoryPKcz(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK14QMessageLogger5debugEv
    // invoke: QDebug debug()
    C._ZNK14QMessageLogger5debugEv(this.qclsinst)
  case 3:
    // invoke: _ZNK14QMessageLogger5debugERK16QLoggingCategory
    // invoke: QDebug debug(const class QLoggingCategory &)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger5debugERK16QLoggingCategory(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageLogger", "debug", args)
  }

}

  // proto:  void QMessageLogger::warning(const QLoggingCategory & cat, const char * msg);
func (this *QMessageLogger) warning(args ...interface{}) () {
  // warning(const class QLoggingCategory &, const char *, ...)
  // warning(const char *, ...)
  // warning(CategoryFunction)
  // warning(const class QLoggingCategory &)
  // warning(CategoryFunction, const char *, ...)
  // warning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"
  vtys[5] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger7warningERK16QLoggingCategoryPKcz
    // invoke: void warning(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C._ZNK14QMessageLogger7warningERK16QLoggingCategoryPKcz(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK14QMessageLogger7warningEPKcz
    // invoke: void warning(const char *, ...)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger7warningEPKcz(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QMessageLogger7warningERK16QLoggingCategory
    // invoke: QDebug warning(const class QLoggingCategory &)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger7warningERK16QLoggingCategory(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK14QMessageLogger7warningEv
    // invoke: QDebug warning()
    C._ZNK14QMessageLogger7warningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageLogger", "warning", args)
  }

}

  // proto:  void QMessageLogger::fatal(const char * msg);
func (this *QMessageLogger) fatal(args ...interface{}) () {
  // fatal(const char *, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger5fatalEPKcz
    // invoke: void fatal(const char *, ...)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger5fatalEPKcz(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageLogger", "fatal", args)
  }

}

  // proto:  void QMessageLogger::critical(const char * msg);
func (this *QMessageLogger) critical(args ...interface{}) () {
  // critical(CategoryFunction)
  // critical(const char *, ...)
  // critical(const class QLoggingCategory &)
  // critical()
  // critical(CategoryFunction, const char *, ...)
  // critical(const class QLoggingCategory &, const char *, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.VoidpTy() // "CategoryFunction"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[5][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger8criticalEPKcz
    // invoke: void critical(const char *, ...)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger8criticalEPKcz(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QMessageLogger8criticalERK16QLoggingCategory
    // invoke: QDebug critical(const class QLoggingCategory &)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QMessageLogger8criticalERK16QLoggingCategory(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QMessageLogger8criticalEv
    // invoke: QDebug critical()
    C._ZNK14QMessageLogger8criticalEv(this.qclsinst)
  case 3:
    // invoke: _ZNK14QMessageLogger8criticalERK16QLoggingCategoryPKcz
    // invoke: void critical(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(QLoggingCategory).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C._ZNK14QMessageLogger8criticalERK16QLoggingCategoryPKcz(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageLogger", "critical", args)
  }

}

  // proto:  void QMessageLogger::QMessageLogger();
func NewQMessageLogger(args ...interface{}) QMessageLogger {
  return QMessageLogger{}
}

  // proto:  void QMessageLogger::noDebug(const char * );
func (this *QMessageLogger) noDebug(args ...interface{}) () {
  // noDebug(const char *, ...)
  // noDebug()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger7noDebugEPKcz
    // invoke: void noDebug(const char *, ...)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C.demth_ZNK14QMessageLogger7noDebugEPKcz(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QMessageLogger7noDebugEv
    // invoke: QNoDebug noDebug()
    C._ZNK14QMessageLogger7noDebugEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMessageLogger", "noDebug", args)
  }

}

// <= body block end

