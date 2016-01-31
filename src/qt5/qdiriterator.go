package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qdiriterator.h
// dst-file: /src/core/qdiriterator.go
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
  // proto:  void QDirIterator::~QDirIterator();
extern void C_ZN12QDirIteratorD2Ev(void* qthis); // 4
  // proto:  QString QDirIterator::filePath();
extern void C_ZNK12QDirIterator8filePathEv(void* qthis); // 4
  // proto:  QString QDirIterator::next();
extern void C_ZN12QDirIterator4nextEv(void* qthis); // 4
  // proto:  bool QDirIterator::hasNext();
extern void C_ZNK12QDirIterator7hasNextEv(void* qthis); // 4
  // proto:  QString QDirIterator::path();
extern void C_ZNK12QDirIterator4pathEv(void* qthis); // 4
  // proto:  QFileInfo QDirIterator::fileInfo();
extern void C_ZNK12QDirIterator8fileInfoEv(void* qthis); // 4
  // proto:  QString QDirIterator::fileName();
extern void C_ZNK12QDirIterator8fileNameEv(void* qthis); // 4
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

// class sizeof(QDirIterator)=1
type QDirIterator struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QDirIterator()
func (this *QDirIterator) FreeQDirIterator(args ...interface{}) () {
  // ~QDirIterator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QDirIteratorD0Ev
    // invoke: void ~QDirIterator()
    C.C_ZN12QDirIteratorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDirIterator", "~QDirIterator", args)
  }

}

// filePath()
func (this *QDirIterator) filePath(args ...interface{}) () {
  // filePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QDirIterator8filePathEv
    // invoke: QString filePath()
    var ret = C.C_ZNK12QDirIterator8filePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDirIterator", "filePath", args)
  }

}

// next()
func (this *QDirIterator) next(args ...interface{}) () {
  // next()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QDirIterator4nextEv
    // invoke: QString next()
    var ret = C.C_ZN12QDirIterator4nextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDirIterator", "next", args)
  }

}

// hasNext()
func (this *QDirIterator) hasNext(args ...interface{}) () {
  // hasNext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QDirIterator7hasNextEv
    // invoke: bool hasNext()
    var ret = C.C_ZNK12QDirIterator7hasNextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDirIterator", "hasNext", args)
  }

}

// path()
func (this *QDirIterator) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QDirIterator4pathEv
    // invoke: QString path()
    var ret = C.C_ZNK12QDirIterator4pathEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDirIterator", "path", args)
  }

}

// fileInfo()
func (this *QDirIterator) fileInfo(args ...interface{}) () {
  // fileInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QDirIterator8fileInfoEv
    // invoke: QFileInfo fileInfo()
    var ret = C.C_ZNK12QDirIterator8fileInfoEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDirIterator", "fileInfo", args)
  }

}

// fileName()
func (this *QDirIterator) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QDirIterator8fileNameEv
    // invoke: QString fileName()
    var ret = C.C_ZNK12QDirIterator8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QDirIterator", "fileName", args)
  }

}

// <= body block end

