package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qtextboundaryfinder.h
// dst-file: /src/core/qtextboundaryfinder.go
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
  // proto:  bool QTextBoundaryFinder::isAtBoundary();
extern void _ZNK19QTextBoundaryFinder12isAtBoundaryEv(void* qthis);
  // proto:  int QTextBoundaryFinder::toNextBoundary();
extern void _ZN19QTextBoundaryFinder14toNextBoundaryEv(void* qthis);
  // proto:  void QTextBoundaryFinder::toEnd();
extern void _ZN19QTextBoundaryFinder5toEndEv(void* qthis);
  // proto:  void QTextBoundaryFinder::QTextBoundaryFinder(const QTextBoundaryFinder & other);
extern void* dector_ZN19QTextBoundaryFinderC1ERKS_(void* arg0);
extern void _ZN19QTextBoundaryFinderC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextBoundaryFinder::setPosition(int position);
extern void _ZN19QTextBoundaryFinder11setPositionEi(void* qthis, int arg0);
  // proto:  void QTextBoundaryFinder::QTextBoundaryFinder();
extern void* dector_ZN19QTextBoundaryFinderC1Ev();
extern void _ZN19QTextBoundaryFinderC1Ev(void* qthis);
  // proto:  int QTextBoundaryFinder::toPreviousBoundary();
extern void _ZN19QTextBoundaryFinder18toPreviousBoundaryEv(void* qthis);
  // proto:  bool QTextBoundaryFinder::isValid();
extern void demth_ZNK19QTextBoundaryFinder7isValidEv(void* qthis);
  // proto:  void QTextBoundaryFinder::~QTextBoundaryFinder();
extern void _ZN19QTextBoundaryFinderD0Ev(void* qthis);
  // proto:  QString QTextBoundaryFinder::string();
extern void _ZNK19QTextBoundaryFinder6stringEv(void* qthis);
  // proto:  void QTextBoundaryFinder::toStart();
extern void _ZN19QTextBoundaryFinder7toStartEv(void* qthis);
  // proto:  int QTextBoundaryFinder::position();
extern void _ZNK19QTextBoundaryFinder8positionEv(void* qthis);
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

// class sizeof(QTextBoundaryFinder)=48
type QTextBoundaryFinder struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QTextBoundaryFinder::isAtBoundary();
func (this *QTextBoundaryFinder) isAtBoundary(args ...interface{}) () {
  // isAtBoundary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder12isAtBoundaryEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isAtBoundary", args)
  }

}

  // proto:  int QTextBoundaryFinder::toNextBoundary();
func (this *QTextBoundaryFinder) toNextBoundary(args ...interface{}) () {
  // toNextBoundary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder14toNextBoundaryEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toNextBoundary", args)
  }

}

  // proto:  void QTextBoundaryFinder::toEnd();
func (this *QTextBoundaryFinder) toEnd(args ...interface{}) () {
  // toEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder5toEndEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toEnd", args)
  }

}

  // proto:  void QTextBoundaryFinder::QTextBoundaryFinder(const QTextBoundaryFinder & other);
func NewQTextBoundaryFinder(args ...interface{}) QTextBoundaryFinder {
  return QTextBoundaryFinder{}
}

  // proto:  void QTextBoundaryFinder::setPosition(int position);
func (this *QTextBoundaryFinder) setPosition(args ...interface{}) () {
  // setPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder11setPositionEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "setPosition", args)
  }

}

  // proto:  int QTextBoundaryFinder::toPreviousBoundary();
func (this *QTextBoundaryFinder) toPreviousBoundary(args ...interface{}) () {
  // toPreviousBoundary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder18toPreviousBoundaryEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toPreviousBoundary", args)
  }

}

  // proto:  bool QTextBoundaryFinder::isValid();
func (this *QTextBoundaryFinder) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder7isValidEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isValid", args)
  }

}

  // proto:  void QTextBoundaryFinder::~QTextBoundaryFinder();
func (this *QTextBoundaryFinder) FreeQTextBoundaryFinder(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "~QTextBoundaryFinder", args)
  }

}

  // proto:  QString QTextBoundaryFinder::string();
func (this *QTextBoundaryFinder) string(args ...interface{}) () {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder6stringEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "string", args)
  }

}

  // proto:  void QTextBoundaryFinder::toStart();
func (this *QTextBoundaryFinder) toStart(args ...interface{}) () {
  // toStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinder7toStartEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toStart", args)
  }

}

  // proto:  int QTextBoundaryFinder::position();
func (this *QTextBoundaryFinder) position(args ...interface{}) () {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder8positionEv
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "position", args)
  }

}

// <= body block end

