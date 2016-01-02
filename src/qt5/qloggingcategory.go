package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QLoggingCategory::QLoggingCategory(const char * category, QtMsgType severityLevel);
extern void* dector_ZN16QLoggingCategoryC1EPKc9QtMsgType(char* arg0, int arg1);
extern void _ZN16QLoggingCategoryC1EPKc9QtMsgType(void* qthis, char* arg0, int arg1);
  // proto:  void QLoggingCategory::QLoggingCategory(const QLoggingCategory & );
extern void* dector_ZN16QLoggingCategoryC1ERKS_(void* arg0);
extern void _ZN16QLoggingCategoryC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QLoggingCategory::isDebugEnabled();
extern void _ZNK16QLoggingCategory14isDebugEnabledEv(void* qthis);
  // proto:  void QLoggingCategory::~QLoggingCategory();
extern void _ZN16QLoggingCategoryD0Ev(void* qthis);
  // proto:  void QLoggingCategory::QLoggingCategory(const char * category);
extern void* dector_ZN16QLoggingCategoryC1EPKc(char* arg0);
extern void _ZN16QLoggingCategoryC1EPKc(void* qthis, char* arg0);
  // proto:  void QLoggingCategory::setEnabled(QtMsgType type, bool enable);
extern void _ZN16QLoggingCategory10setEnabledE9QtMsgTypeb(void* qthis, int arg0, bool arg1);
  // proto:  bool QLoggingCategory::isEnabled(QtMsgType type);
extern void _ZNK16QLoggingCategory9isEnabledE9QtMsgType(void* qthis, int arg0);
  // proto:  bool QLoggingCategory::isWarningEnabled();
extern void _ZNK16QLoggingCategory16isWarningEnabledEv(void* qthis);
  // proto:  bool QLoggingCategory::isInfoEnabled();
extern void _ZNK16QLoggingCategory13isInfoEnabledEv(void* qthis);
  // proto:  const char * QLoggingCategory::categoryName();
extern void _ZNK16QLoggingCategory12categoryNameEv(void* qthis);
  // proto:  bool QLoggingCategory::isCriticalEnabled();
extern void _ZNK16QLoggingCategory17isCriticalEnabledEv(void* qthis);
  // proto: static QLoggingCategory * QLoggingCategory::defaultCategory();
extern void _ZN16QLoggingCategory15defaultCategoryEv();
  // proto: static void QLoggingCategory::setFilterRules(const QString & rules);
extern void _ZN16QLoggingCategory14setFilterRulesERK7QString(void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QLoggingCategory::QLoggingCategory(const char * category, QtMsgType severityLevel);
func NewQLoggingCategory(args ...interface{}) QLoggingCategory {
  return QLoggingCategory{}
}

  // proto:  bool QLoggingCategory::isDebugEnabled();
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
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isDebugEnabled", args)
  }

}

  // proto:  void QLoggingCategory::~QLoggingCategory();
func (this *QLoggingCategory) FreeQLoggingCategory(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLoggingCategory", "~QLoggingCategory", args)
  }

}

  // proto:  void QLoggingCategory::setEnabled(QtMsgType type, bool enable);
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
  default:
    qtrt.ErrorResolve("QLoggingCategory", "setEnabled", args)
  }

}

  // proto:  bool QLoggingCategory::isEnabled(QtMsgType type);
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
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isEnabled", args)
  }

}

  // proto:  bool QLoggingCategory::isWarningEnabled();
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
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isWarningEnabled", args)
  }

}

  // proto:  bool QLoggingCategory::isInfoEnabled();
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
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isInfoEnabled", args)
  }

}

  // proto:  const char * QLoggingCategory::categoryName();
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
  default:
    qtrt.ErrorResolve("QLoggingCategory", "categoryName", args)
  }

}

  // proto:  bool QLoggingCategory::isCriticalEnabled();
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
  default:
    qtrt.ErrorResolve("QLoggingCategory", "isCriticalEnabled", args)
  }

}

  // proto: static QLoggingCategory * QLoggingCategory::defaultCategory();
func (this *QLoggingCategory) defaultCategory_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLoggingCategory", "defaultCategory", args)
  }

}

  // proto: static void QLoggingCategory::setFilterRules(const QString & rules);
func (this *QLoggingCategory) setFilterRules_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLoggingCategory", "setFilterRules", args)
  }

}

// <= body block end

