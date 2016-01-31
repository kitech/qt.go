package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTextOption::~QTextOption();
extern void C_ZN11QTextOptionD2Ev(void* qthis); // 4
  // proto:  void QTextOption::setUseDesignMetrics(bool b);
extern void C_ZN11QTextOption19setUseDesignMetricsEb(void* qthis, bool arg0); // 2
  // proto:  bool QTextOption::useDesignMetrics();
extern void C_ZNK11QTextOption16useDesignMetricsEv(void* qthis); // 2
  // proto:  void QTextOption::setTabStop(qreal tabStop);
extern void C_ZN11QTextOption10setTabStopEd(void* qthis, double arg0); // 2
  // proto:  QList<QTextOption::Tab> QTextOption::tabs();
extern void C_ZNK11QTextOption4tabsEv(void* qthis); // 4
  // proto:  qreal QTextOption::tabStop();
extern void C_ZNK11QTextOption7tabStopEv(void* qthis); // 2
  // proto:  QList<qreal> QTextOption::tabArray();
extern void C_ZNK11QTextOption8tabArrayEv(void* qthis); // 4
  // proto:  Flags QTextOption::flags();
extern void C_ZNK11QTextOption5flagsEv(void* qthis); // 2
  // proto:  void QTextOption::QTextOption(const QTextOption & o);
extern void C_ZN11QTextOptionC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QTextOption::QTextOption();
extern void C_ZN11QTextOptionC2Ev(void* qthis); // 3
  // proto:  QTextOption::WrapMode QTextOption::wrapMode();
extern void C_ZNK11QTextOption8wrapModeEv(void* qthis); // 2
  // proto:  Qt::Alignment QTextOption::alignment();
extern void C_ZNK11QTextOption9alignmentEv(void* qthis); // 2
  // proto:  Qt::LayoutDirection QTextOption::textDirection();
extern void C_ZNK11QTextOption13textDirectionEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QTextOption()
func (this *QTextOption) FreeQTextOption(args ...interface{}) () {
  // ~QTextOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextOptionD0Ev
    // invoke: void ~QTextOption()
    C.C_ZN11QTextOptionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "~QTextOption", args)
  }

}

// setUseDesignMetrics(_Bool)
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
    // invoke: void setUseDesignMetrics(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextOption19setUseDesignMetricsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextOption", "setUseDesignMetrics", args)
  }

}

// useDesignMetrics()
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
    // invoke: bool useDesignMetrics()
    C.C_ZNK11QTextOption16useDesignMetricsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "useDesignMetrics", args)
  }

}

// setTabStop(qreal)
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
    // invoke: void setTabStop(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextOption10setTabStopEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextOption", "setTabStop", args)
  }

}

// tabs()
func (this *QTextOption) tabs(args ...interface{}) () {
  // tabs()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption4tabsEv
    // invoke: QList<QTextOption::Tab> tabs()
    C.C_ZNK11QTextOption4tabsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "tabs", args)
  }

}

// tabStop()
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
    // invoke: qreal tabStop()
    C.C_ZNK11QTextOption7tabStopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "tabStop", args)
  }

}

// tabArray()
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
    // invoke: QList<qreal> tabArray()
    C.C_ZNK11QTextOption8tabArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "tabArray", args)
  }

}

// flags()
func (this *QTextOption) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption5flagsEv
    // invoke: Flags flags()
    C.C_ZNK11QTextOption5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "flags", args)
  }

}

// QTextOption(const class QTextOption &)
func NewQTextOption(args ...interface{}) QTextOption {
  // QTextOption(const class QTextOption &)
  // QTextOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextOption{}) // "const QTextOption &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextOptionC1ERKS_
    // invoke: void QTextOption(const class QTextOption &)
    var arg0 = args[0].(QTextOption).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextOptionC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN11QTextOptionC1Ev
    // invoke: void QTextOption()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextOptionC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextOption", "QTextOption", args)
  }

  return QTextOption{}
}

// wrapMode()
func (this *QTextOption) wrapMode(args ...interface{}) () {
  // wrapMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption8wrapModeEv
    // invoke: QTextOption::WrapMode wrapMode()
    C.C_ZNK11QTextOption8wrapModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "wrapMode", args)
  }

}

// alignment()
func (this *QTextOption) alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK11QTextOption9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "alignment", args)
  }

}

// textDirection()
func (this *QTextOption) textDirection(args ...interface{}) () {
  // textDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextOption13textDirectionEv
    // invoke: Qt::LayoutDirection textDirection()
    C.C_ZNK11QTextOption13textDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextOption", "textDirection", args)
  }

}

// <= body block end

