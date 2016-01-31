package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZNK12QDirIterator8filePathEv(void* qthis); // 4
  // proto:  QString QDirIterator::next();
extern void* C_ZN12QDirIterator4nextEv(void* qthis); // 4
  // proto:  bool QDirIterator::hasNext();
extern bool C_ZNK12QDirIterator7hasNextEv(void* qthis); // 4
  // proto:  QString QDirIterator::path();
extern void* C_ZNK12QDirIterator4pathEv(void* qthis); // 4
  // proto:  QFileInfo QDirIterator::fileInfo();
extern void* C_ZNK12QDirIterator8fileInfoEv(void* qthis); // 4
  // proto:  QString QDirIterator::fileName();
extern void* C_ZNK12QDirIterator8fileNameEv(void* qthis); // 4
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
func (this *QDirIterator) Freeqdiriterator(args ...interface{}) () {
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

  return
}

// filePath()
func (this *QDirIterator) Filepath(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QDirIterator8filePathEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDirIterator", "filePath", args)
  }

  return
}

// next()
func (this *QDirIterator) Next(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN12QDirIterator4nextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDirIterator", "next", args)
  }

  return
}

// hasNext()
func (this *QDirIterator) Hasnext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QDirIterator7hasNextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDirIterator", "hasNext", args)
  }

  return
}

// path()
func (this *QDirIterator) Path(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QDirIterator4pathEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDirIterator", "path", args)
  }

  return
}

// fileInfo()
func (this *QDirIterator) Fileinfo(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QDirIterator8fileInfoEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFileInfo{}) // "QFileInfo"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDirIterator", "fileInfo", args)
  }

  return
}

// fileName()
func (this *QDirIterator) Filename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QDirIterator8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDirIterator", "fileName", args)
  }

  return
}

// <= body block end

