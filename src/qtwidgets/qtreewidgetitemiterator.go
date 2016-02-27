package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qtreewidgetitemiterator.h
// dst-file: /src/widgets/qtreewidgetitemiterator.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTreeWidgetItemIterator::~QTreeWidgetItemIterator();
extern void C_ZN23QTreeWidgetItemIteratorD2Ev(void* qthis); // 4
  // proto:  void QTreeWidgetItemIterator::QTreeWidgetItemIterator(const QTreeWidgetItemIterator & it);
extern void* C_ZN23QTreeWidgetItemIteratorC2ERKS_(void* arg0); // 3
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTreeWidgetItemIterator)=1
type QTreeWidgetItemIterator struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QTreeWidgetItemIterator()
func (this *QTreeWidgetItemIterator) Free(args ...interface{}) () {
  // ~QTreeWidgetItemIterator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QTreeWidgetItemIteratorD0Ev
    // invoke: void ~QTreeWidgetItemIterator()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN23QTreeWidgetItemIteratorD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTreeWidgetItemIterator", "~QTreeWidgetItemIterator", args)
  }

  return
}

// QTreeWidgetItemIterator(const class QTreeWidgetItemIterator &)
func GcfreeQTreeWidgetItemIterator(this *QTreeWidgetItemIterator) {
  qtrt.UniverseFree(this)
}
func NewQTreeWidgetItemIterator(args ...interface{}) *QTreeWidgetItemIterator {
  // QTreeWidgetItemIterator(const class QTreeWidgetItemIterator &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItemIterator{}) // "const QTreeWidgetItemIterator &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QTreeWidgetItemIteratorC1ERKS_
    // invoke: void QTreeWidgetItemIterator(const class QTreeWidgetItemIterator &)
    var arg0 = args[0].(*QTreeWidgetItemIterator).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QTreeWidgetItemIteratorC2ERKS_(arg0)
    this := &QTreeWidgetItemIterator{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTreeWidgetItemIterator)
    return this
  default:
    qtrt.ErrorResolve("QTreeWidgetItemIterator", "QTreeWidgetItemIterator", args)
  }

  return nil // QTreeWidgetItemIterator{Qclsinst:qthis}
}

// <= body block end

