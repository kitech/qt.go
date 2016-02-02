package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QMessageLogContext::copy(const QMessageLogContext & logContext);
extern void C_ZN18QMessageLogContext4copyERKS_(void* qthis, void* arg0); // 4
  // proto:  void QMessageLogContext::QMessageLogContext(const char * fileName, int lineNumber, const char * functionName, const char * categoryName);
extern void* C_ZN18QMessageLogContextC2EPKciS1_S1_(void* arg0, int32_t arg1, void* arg2, void* arg3); // 1
  // proto:  void QMessageLogContext::QMessageLogContext();
extern void* C_ZN18QMessageLogContextC2Ev(); // 1
  // proto:  void QMessageLogger::info(const QLoggingCategory & cat, const char * msg);
extern void C_ZNK14QMessageLogger4infoERK16QLoggingCategoryPKcz(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QMessageLogger::info(const char * msg);
extern void C_ZNK14QMessageLogger4infoEPKcz(void* qthis, void* arg0); // 4
  // proto:  QDebug QMessageLogger::info();
extern void* C_ZNK14QMessageLogger4infoEv(void* qthis); // 4
  // proto:  QDebug QMessageLogger::info(const QLoggingCategory & cat);
extern void* C_ZNK14QMessageLogger4infoERK16QLoggingCategory(void* qthis, void* arg0); // 4
  // proto:  QDebug QMessageLogger::warning(const QLoggingCategory & cat);
extern void* C_ZNK14QMessageLogger7warningERK16QLoggingCategory(void* qthis, void* arg0); // 4
  // proto:  void QMessageLogger::warning(const char * msg);
extern void C_ZNK14QMessageLogger7warningEPKcz(void* qthis, void* arg0); // 4
  // proto:  void QMessageLogger::warning(const QLoggingCategory & cat, const char * msg);
extern void C_ZNK14QMessageLogger7warningERK16QLoggingCategoryPKcz(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QDebug QMessageLogger::warning();
extern void* C_ZNK14QMessageLogger7warningEv(void* qthis); // 4
  // proto:  QDebug QMessageLogger::critical();
extern void* C_ZNK14QMessageLogger8criticalEv(void* qthis); // 4
  // proto:  void QMessageLogger::critical(const char * msg);
extern void C_ZNK14QMessageLogger8criticalEPKcz(void* qthis, void* arg0); // 4
  // proto:  QDebug QMessageLogger::critical(const QLoggingCategory & cat);
extern void* C_ZNK14QMessageLogger8criticalERK16QLoggingCategory(void* qthis, void* arg0); // 4
  // proto:  void QMessageLogger::critical(const QLoggingCategory & cat, const char * msg);
extern void C_ZNK14QMessageLogger8criticalERK16QLoggingCategoryPKcz(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QMessageLogger::debug(const QLoggingCategory & cat, const char * msg);
extern void C_ZNK14QMessageLogger5debugERK16QLoggingCategoryPKcz(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QDebug QMessageLogger::debug();
extern void* C_ZNK14QMessageLogger5debugEv(void* qthis); // 4
  // proto:  void QMessageLogger::debug(const char * msg);
extern void C_ZNK14QMessageLogger5debugEPKcz(void* qthis, void* arg0); // 4
  // proto:  QDebug QMessageLogger::debug(const QLoggingCategory & cat);
extern void* C_ZNK14QMessageLogger5debugERK16QLoggingCategory(void* qthis, void* arg0); // 4
  // proto:  void QMessageLogger::fatal(const char * msg);
extern void C_ZNK14QMessageLogger5fatalEPKcz(void* qthis, void* arg0); // 4
  // proto:  void QMessageLogger::QMessageLogger();
extern void* C_ZN14QMessageLoggerC2Ev(); // 1
  // proto:  void QMessageLogger::QMessageLogger(const char * file, int line, const char * function);
extern void* C_ZN14QMessageLoggerC2EPKciS1_(void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  void QMessageLogger::QMessageLogger(const char * file, int line, const char * function, const char * category);
extern void* C_ZN14QMessageLoggerC2EPKciS1_S1_(void* arg0, int32_t arg1, void* arg2, void* arg3); // 1
  // proto:  void QMessageLogger::noDebug(const char * );
extern void C_ZNK14QMessageLogger7noDebugEPKcz(void* qthis, void* arg0); // 2
  // proto:  QNoDebug QMessageLogger::noDebug();
extern void* C_ZNK14QMessageLogger7noDebugEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QMessageLogger)=32
type QMessageLogger struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// copy(const class QMessageLogContext &)
func (this *QMessageLogContext) Copy(args ...interface{}) () {
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
    var arg0 = args[0].(*QMessageLogContext).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QMessageLogContext4copyERKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageLogContext", "copy", args)
  }

  return
}

// QMessageLogContext(const char *, int, const char *, const char *)
func NewQMessageLogContext(args ...interface{}) *QMessageLogContext {
  // QMessageLogContext(const char *, int, const char *, const char *)
  // QMessageLogContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[0][3] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QMessageLogContextC1EPKciS1_S1_
    // invoke: void QMessageLogContext(const char *, int, const char *, const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    argif3, free3 := qtrt.HandyConvert2c(args[3], vtys[0][3])
    var arg3 = argif3.(unsafe.Pointer)
    if false {fmt.Println(argif3, arg3)}
    if free3 {defer C.free(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QMessageLogContextC2EPKciS1_S1_(arg0, arg1, arg2, arg3)
    return &QMessageLogContext{Qclsinst:qthis}
  case 1:
    // invoke: _ZN18QMessageLogContextC1Ev
    // invoke: void QMessageLogContext()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QMessageLogContextC2Ev()
    return &QMessageLogContext{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMessageLogContext", "QMessageLogContext", args)
  }

  return nil // QMessageLogContext{Qclsinst:qthis}
}

// info(const class QLoggingCategory &, const char *, ...)
func (this *QMessageLogger) Info(args ...interface{}) () {
  // info(const class QLoggingCategory &, const char *, ...)
  // info(const char *, ...)
  // info()
  // info(const class QLoggingCategory &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger4infoERK16QLoggingCategoryPKcz
    // invoke: void info(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZNK14QMessageLogger4infoERK16QLoggingCategoryPKcz(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK14QMessageLogger4infoEPKcz
    // invoke: void info(const char *, ...)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZNK14QMessageLogger4infoEPKcz(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QMessageLogger4infoEv
    // invoke: QDebug info()
    var ret0 = C.C_ZNK14QMessageLogger4infoEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
  case 3:
    // invoke: _ZNK14QMessageLogger4infoERK16QLoggingCategory
    // invoke: QDebug info(const class QLoggingCategory &)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QMessageLogger4infoERK16QLoggingCategory(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QMessageLogger", "info", args)
  }

  return
}

// warning(const class QLoggingCategory &)
func (this *QMessageLogger) Warning(args ...interface{}) (ret interface{}) {
  // warning(const class QLoggingCategory &)
  // warning(const char *, ...)
  // warning(const class QLoggingCategory &, const char *, ...)
  // warning()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger7warningERK16QLoggingCategory
    // invoke: QDebug warning(const class QLoggingCategory &)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QMessageLogger7warningERK16QLoggingCategory(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDebug{}) // "QDebug"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK14QMessageLogger7warningEPKcz
    // invoke: void warning(const char *, ...)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZNK14QMessageLogger7warningEPKcz(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QMessageLogger7warningERK16QLoggingCategoryPKcz
    // invoke: void warning(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[2][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZNK14QMessageLogger7warningERK16QLoggingCategoryPKcz(this.Qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK14QMessageLogger7warningEv
    // invoke: QDebug warning()
    var ret0 = C.C_ZNK14QMessageLogger7warningEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDebug{}) // "QDebug"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMessageLogger", "warning", args)
  }

  return
}

// critical()
func (this *QMessageLogger) Critical(args ...interface{}) (ret interface{}) {
  // critical()
  // critical(const char *, ...)
  // critical(const class QLoggingCategory &)
  // critical(const class QLoggingCategory &, const char *, ...)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[3][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger8criticalEv
    // invoke: QDebug critical()
    var ret0 = C.C_ZNK14QMessageLogger8criticalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDebug{}) // "QDebug"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK14QMessageLogger8criticalEPKcz
    // invoke: void critical(const char *, ...)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZNK14QMessageLogger8criticalEPKcz(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QMessageLogger8criticalERK16QLoggingCategory
    // invoke: QDebug critical(const class QLoggingCategory &)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QMessageLogger8criticalERK16QLoggingCategory(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDebug{}) // "QDebug"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK14QMessageLogger8criticalERK16QLoggingCategoryPKcz
    // invoke: void critical(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[3][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZNK14QMessageLogger8criticalERK16QLoggingCategoryPKcz(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMessageLogger", "critical", args)
  }

  return
}

// debug(const class QLoggingCategory &, const char *, ...)
func (this *QMessageLogger) Debug(args ...interface{}) () {
  // debug(const class QLoggingCategory &, const char *, ...)
  // debug()
  // debug(const char *, ...)
  // debug(const class QLoggingCategory &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QLoggingCategory{}) // "const QLoggingCategory &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QMessageLogger5debugERK16QLoggingCategoryPKcz
    // invoke: void debug(const class QLoggingCategory &, const char *, ...)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZNK14QMessageLogger5debugERK16QLoggingCategoryPKcz(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK14QMessageLogger5debugEv
    // invoke: QDebug debug()
    var ret0 = C.C_ZNK14QMessageLogger5debugEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
  case 2:
    // invoke: _ZNK14QMessageLogger5debugEPKcz
    // invoke: void debug(const char *, ...)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZNK14QMessageLogger5debugEPKcz(this.Qclsinst, arg0)
  case 3:
    // invoke: _ZNK14QMessageLogger5debugERK16QLoggingCategory
    // invoke: QDebug debug(const class QLoggingCategory &)
    var arg0 = args[0].(*QLoggingCategory).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QMessageLogger5debugERK16QLoggingCategory(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QMessageLogger", "debug", args)
  }

  return
}

// fatal(const char *, ...)
func (this *QMessageLogger) Fatal(args ...interface{}) () {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZNK14QMessageLogger5fatalEPKcz(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMessageLogger", "fatal", args)
  }

  return
}

// QMessageLogger()
func NewQMessageLogger(args ...interface{}) *QMessageLogger {
  // QMessageLogger()
  // QMessageLogger(const char *, int, const char *)
  // QMessageLogger(const char *, int, const char *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.ByteTy(true) // "const char *"
  vtys[2][3] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QMessageLoggerC1Ev
    // invoke: void QMessageLogger()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QMessageLoggerC2Ev()
    return &QMessageLogger{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QMessageLoggerC1EPKciS1_
    // invoke: void QMessageLogger(const char *, int, const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[1][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QMessageLoggerC2EPKciS1_(arg0, arg1, arg2)
    return &QMessageLogger{Qclsinst:qthis}
  case 2:
    // invoke: _ZN14QMessageLoggerC1EPKciS1_S1_
    // invoke: void QMessageLogger(const char *, int, const char *, const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[2][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    argif3, free3 := qtrt.HandyConvert2c(args[3], vtys[2][3])
    var arg3 = argif3.(unsafe.Pointer)
    if false {fmt.Println(argif3, arg3)}
    if free3 {defer C.free(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QMessageLoggerC2EPKciS1_S1_(arg0, arg1, arg2, arg3)
    return &QMessageLogger{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMessageLogger", "QMessageLogger", args)
  }

  return nil // QMessageLogger{Qclsinst:qthis}
}

// noDebug(const char *, ...)
func (this *QMessageLogger) Nodebug(args ...interface{}) () {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZNK14QMessageLogger7noDebugEPKcz(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QMessageLogger7noDebugEv
    // invoke: QNoDebug noDebug()
    var ret0 = C.C_ZNK14QMessageLogger7noDebugEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
  default:
    qtrt.ErrorResolve("QMessageLogger", "noDebug", args)
  }

  return
}

// <= body block end

