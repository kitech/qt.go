package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTreeWidgetItemIterator::~QTreeWidgetItemIterator();
extern void _ZN23QTreeWidgetItemIteratorD2Ev(void* qthis); // 4
  // proto:  void QTreeWidgetItemIterator::QTreeWidgetItemIterator(const QTreeWidgetItemIterator & it);
extern void _ZN23QTreeWidgetItemIteratorC2ERKS_(void* qthis, void* arg0); // 3
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

// class sizeof(QTreeWidgetItemIterator)=1
type QTreeWidgetItemIterator struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QTreeWidgetItemIterator()
func (this *QTreeWidgetItemIterator) FreeQTreeWidgetItemIterator(args ...interface{}) () {
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
    C._ZN23QTreeWidgetItemIteratorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItemIterator", "~QTreeWidgetItemIterator", args)
  }

}

// QTreeWidgetItemIterator(const class QTreeWidgetItemIterator &)
func NewQTreeWidgetItemIterator(args ...interface{}) QTreeWidgetItemIterator {
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
    var arg0 = args[0].(QTreeWidgetItemIterator).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN23QTreeWidgetItemIteratorC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItemIterator", "QTreeWidgetItemIterator", args)
  }

  return QTreeWidgetItemIterator{}
}

// <= body block end

