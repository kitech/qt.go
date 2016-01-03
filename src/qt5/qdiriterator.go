package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  QString QDirIterator::fileName();
extern void _ZNK12QDirIterator8fileNameEv(void* qthis);
  // proto:  QString QDirIterator::path();
extern void _ZNK12QDirIterator4pathEv(void* qthis);
  // proto:  void QDirIterator::QDirIterator(const QDirIterator & );
extern void* dector_ZN12QDirIteratorC1ERKS_(void* arg0);
extern void _ZN12QDirIteratorC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QDirIterator::next();
extern void _ZN12QDirIterator4nextEv(void* qthis);
  // proto:  QString QDirIterator::filePath();
extern void _ZNK12QDirIterator8filePathEv(void* qthis);
  // proto:  void QDirIterator::~QDirIterator();
extern void _ZN12QDirIteratorD0Ev(void* qthis);
  // proto:  QFileInfo QDirIterator::fileInfo();
extern void _ZNK12QDirIterator8fileInfoEv(void* qthis);
  // proto:  bool QDirIterator::hasNext();
extern void _ZNK12QDirIterator7hasNextEv(void* qthis);
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

  // proto:  QString QDirIterator::fileName();
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
    C._ZNK12QDirIterator8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDirIterator", "fileName", args)
  }

}

  // proto:  QString QDirIterator::path();
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
    C._ZNK12QDirIterator4pathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDirIterator", "path", args)
  }

}

  // proto:  void QDirIterator::QDirIterator(const QDirIterator & );
func NewQDirIterator(args ...interface{}) QDirIterator {
  return QDirIterator{}
}

  // proto:  QString QDirIterator::next();
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
    C._ZN12QDirIterator4nextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDirIterator", "next", args)
  }

}

  // proto:  QString QDirIterator::filePath();
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
    C._ZNK12QDirIterator8filePathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDirIterator", "filePath", args)
  }

}

  // proto:  void QDirIterator::~QDirIterator();
func (this *QDirIterator) FreeQDirIterator(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDirIterator", "~QDirIterator", args)
  }

}

  // proto:  QFileInfo QDirIterator::fileInfo();
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
    C._ZNK12QDirIterator8fileInfoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDirIterator", "fileInfo", args)
  }

}

  // proto:  bool QDirIterator::hasNext();
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
    C._ZNK12QDirIterator7hasNextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDirIterator", "hasNext", args)
  }

}

// <= body block end

