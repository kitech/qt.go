package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qloggingcategory.h
// dst-file: /src/core/qloggingcategory.go
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
  // proto:  void QLoggingCategory::setEnabled(QtMsgType type, bool enable);
extern void C_ZN16QLoggingCategory10setEnabledE9QtMsgTypeb(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  const char * QLoggingCategory::categoryName();
extern void C_ZNK16QLoggingCategory12categoryNameEv(void* qthis); // 2
  // proto:  bool QLoggingCategory::isCriticalEnabled();
extern void C_ZNK16QLoggingCategory17isCriticalEnabledEv(void* qthis); // 2
  // proto:  bool QLoggingCategory::isEnabled(QtMsgType type);
extern void C_ZNK16QLoggingCategory9isEnabledE9QtMsgType(void* qthis, int32_t arg0); // 4
  // proto:  void QLoggingCategory::QLoggingCategory(const char * category);
extern void C_ZN16QLoggingCategoryC2EPKc(void* qthis, unsigned char* arg0); // 3
  // proto:  void QLoggingCategory::QLoggingCategory(const char * category, QtMsgType severityLevel);
extern void C_ZN16QLoggingCategoryC2EPKc9QtMsgType(void* qthis, unsigned char* arg0, int32_t arg1); // 3
  // proto:  void QLoggingCategory::~QLoggingCategory();
extern void C_ZN16QLoggingCategoryD2Ev(void* qthis); // 4
  // proto:  bool QLoggingCategory::isDebugEnabled();
extern void C_ZNK16QLoggingCategory14isDebugEnabledEv(void* qthis); // 2
  // proto:  bool QLoggingCategory::isWarningEnabled();
extern void C_ZNK16QLoggingCategory16isWarningEnabledEv(void* qthis); // 2
  // proto: static void QLoggingCategory::setFilterRules(const QString & rules);
extern void C_ZN16QLoggingCategory14setFilterRulesERK7QString(void* arg0); // 4
  // proto: static QLoggingCategory * QLoggingCategory::defaultCategory();
extern void C_ZN16QLoggingCategory15defaultCategoryEv(); // 4
  // proto:  bool QLoggingCategory::isInfoEnabled();
extern void C_ZNK16QLoggingCategory13isInfoEnabledEv(void* qthis); // 2
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

// class sizeof(QLoggingCategory)=24
type QLoggingCategory struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setEnabled(enum QtMsgType, _Bool)
func (this *QLoggingCategory) setEnabled(args ...interface{}) () {
  // setEnabled(enum QtMsgType, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QtMsgType"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QLoggingCategory10setEnabledE9QtMsgTypeb
    // invoke: void setEnabled(enum QtMsgType, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN16QLoggingCategory10setEnabledE9QtMsgTypeb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "setEnabled", args)
  }

}

// categoryName()
func (this *QLoggingCategory) categoryName(args ...interface{}) () {
  // categoryName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory12categoryNameEv
    // invoke: const char * categoryName()
    C.C_ZNK16QLoggingCategory12categoryNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "categoryName", args)
  }

}

// isCriticalEnabled()
func (this *QLoggingCategory) isCriticalEnabled(args ...interface{}) () {
  // isCriticalEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory17isCriticalEnabledEv
    // invoke: bool isCriticalEnabled()
    C.C_ZNK16QLoggingCategory17isCriticalEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isCriticalEnabled", args)
  }

}

// isEnabled(enum QtMsgType)
func (this *QLoggingCategory) isEnabled(args ...interface{}) () {
  // isEnabled(enum QtMsgType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QtMsgType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory9isEnabledE9QtMsgType
    // invoke: bool isEnabled(enum QtMsgType)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK16QLoggingCategory9isEnabledE9QtMsgType(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isEnabled", args)
  }

}

// QLoggingCategory(const char *)
func NewQLoggingCategory(args ...interface{}) QLoggingCategory {
  // QLoggingCategory(const char *)
  // QLoggingCategory(const char *, enum QtMsgType)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "QtMsgType"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QLoggingCategoryC1EPKc
    // invoke: void QLoggingCategory(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QLoggingCategoryC2EPKc(qthis, arg0)
  case 1:
    // invoke: _ZN16QLoggingCategoryC1EPKc9QtMsgType
    // invoke: void QLoggingCategory(const char *, enum QtMsgType)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN16QLoggingCategoryC2EPKc9QtMsgType(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "QLoggingCategory", args)
  }

  return QLoggingCategory{}
}

// ~QLoggingCategory()
func (this *QLoggingCategory) FreeQLoggingCategory(args ...interface{}) () {
  // ~QLoggingCategory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QLoggingCategoryD0Ev
    // invoke: void ~QLoggingCategory()
    C.C_ZN16QLoggingCategoryD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "~QLoggingCategory", args)
  }

}

// isDebugEnabled()
func (this *QLoggingCategory) isDebugEnabled(args ...interface{}) () {
  // isDebugEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory14isDebugEnabledEv
    // invoke: bool isDebugEnabled()
    C.C_ZNK16QLoggingCategory14isDebugEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isDebugEnabled", args)
  }

}

// isWarningEnabled()
func (this *QLoggingCategory) isWarningEnabled(args ...interface{}) () {
  // isWarningEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory16isWarningEnabledEv
    // invoke: bool isWarningEnabled()
    C.C_ZNK16QLoggingCategory16isWarningEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isWarningEnabled", args)
  }

}

// setFilterRules(const class QString &)
func (this *QLoggingCategory) setFilterRules_s(args ...interface{}) () {
  // setFilterRules(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QLoggingCategory14setFilterRulesERK7QString
    // invoke: void setFilterRules(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QLoggingCategory14setFilterRulesERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "setFilterRules", args)
  }

}

// defaultCategory()
func (this *QLoggingCategory) defaultCategory_s(args ...interface{}) () {
  // defaultCategory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QLoggingCategory15defaultCategoryEv
    // invoke: QLoggingCategory * defaultCategory()
    C.C_ZN16QLoggingCategory15defaultCategoryEv()
  default:
    qtrt.ErrorResolve("QLoggingCategory", "defaultCategory", args)
  }

}

// isInfoEnabled()
func (this *QLoggingCategory) isInfoEnabled(args ...interface{}) () {
  // isInfoEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QLoggingCategory13isInfoEnabledEv
    // invoke: bool isInfoEnabled()
    C.C_ZNK16QLoggingCategory13isInfoEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isInfoEnabled", args)
  }

}

// <= body block end

