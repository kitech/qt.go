package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qtextoption.h
// dst-file: /src/gui/qtextoption.go
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
  // proto:  void QTextOption::QTextOption(const QTextOption & o);
extern void* dector_ZN11QTextOptionC1ERKS_(void* arg0);
extern void _ZN11QTextOptionC1ERKS_(void* qthis, void* arg0);
  // proto:  qreal QTextOption::tabStop();
extern void demth_ZNK11QTextOption7tabStopEv(void* qthis);
  // proto:  void QTextOption::setUseDesignMetrics(bool b);
extern void _ZN11QTextOption19setUseDesignMetricsEb(void* qthis, bool arg0);
  // proto:  void QTextOption::setTabStop(qreal tabStop);
extern void demth_ZN11QTextOption10setTabStopEd(void* qthis, double arg0);
  // proto:  bool QTextOption::useDesignMetrics();
extern void _ZNK11QTextOption16useDesignMetricsEv(void* qthis);
  // proto:  void QTextOption::QTextOption();
extern void* dector_ZN11QTextOptionC1Ev();
extern void _ZN11QTextOptionC1Ev(void* qthis);
  // proto:  QList<qreal> QTextOption::tabArray();
extern void _ZNK11QTextOption8tabArrayEv(void* qthis);
  // proto:  void QTextOption::~QTextOption();
extern void _ZN11QTextOptionD0Ev(void* qthis);
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

// class sizeof(QTextOption)=32
type QTextOption struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QTextOption::QTextOption(const QTextOption & o);
func NewQTextOption(args ...interface{}) QTextOption {
  return QTextOption{}
}

  // proto:  qreal QTextOption::tabStop();
func (this *QTextOption) tabStop(args ...interface{}) () {
  // tabStop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption7tabStopEv
  default:
    qtrt.ErrorResolve("QTextOption", "tabStop", args)
  }

}

  // proto:  void QTextOption::setUseDesignMetrics(bool b);
func (this *QTextOption) setUseDesignMetrics(args ...interface{}) () {
  // setUseDesignMetrics(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextOption19setUseDesignMetricsEb
  default:
    qtrt.ErrorResolve("QTextOption", "setUseDesignMetrics", args)
  }

}

  // proto:  void QTextOption::setTabStop(qreal tabStop);
func (this *QTextOption) setTabStop(args ...interface{}) () {
  // setTabStop(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextOption10setTabStopEd
  default:
    qtrt.ErrorResolve("QTextOption", "setTabStop", args)
  }

}

  // proto:  bool QTextOption::useDesignMetrics();
func (this *QTextOption) useDesignMetrics(args ...interface{}) () {
  // useDesignMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption16useDesignMetricsEv
  default:
    qtrt.ErrorResolve("QTextOption", "useDesignMetrics", args)
  }

}

  // proto:  QList<qreal> QTextOption::tabArray();
func (this *QTextOption) tabArray(args ...interface{}) () {
  // tabArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption8tabArrayEv
  default:
    qtrt.ErrorResolve("QTextOption", "tabArray", args)
  }

}

  // proto:  void QTextOption::~QTextOption();
func (this *QTextOption) FreeQTextOption(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextOption", "~QTextOption", args)
  }

}

// <= body block end

