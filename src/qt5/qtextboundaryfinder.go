package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QTextBoundaryFinder::toNextBoundary();
extern void _ZN19QTextBoundaryFinder14toNextBoundaryEv(void* qthis); // 4
  // proto:  int QTextBoundaryFinder::position();
extern void _ZNK19QTextBoundaryFinder8positionEv(void* qthis); // 4
  // proto:  QString QTextBoundaryFinder::string();
extern void _ZNK19QTextBoundaryFinder6stringEv(void* qthis); // 4
  // proto:  bool QTextBoundaryFinder::isValid();
extern void _ZNK19QTextBoundaryFinder7isValidEv(void* qthis); // 2
  // proto:  BoundaryReasons QTextBoundaryFinder::boundaryReasons();
extern void _ZNK19QTextBoundaryFinder15boundaryReasonsEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::toEnd();
extern void _ZN19QTextBoundaryFinder5toEndEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::~QTextBoundaryFinder();
extern void _ZN19QTextBoundaryFinderD2Ev(void* qthis); // 4
  // proto:  int QTextBoundaryFinder::toPreviousBoundary();
extern void _ZN19QTextBoundaryFinder18toPreviousBoundaryEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::QTextBoundaryFinder();
extern void _ZN19QTextBoundaryFinderC2Ev(void* qthis); // 3
  // proto:  void QTextBoundaryFinder::QTextBoundaryFinder(const QTextBoundaryFinder & other);
extern void _ZN19QTextBoundaryFinderC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  bool QTextBoundaryFinder::isAtBoundary();
extern void _ZNK19QTextBoundaryFinder12isAtBoundaryEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::setPosition(int position);
extern void _ZN19QTextBoundaryFinder11setPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextBoundaryFinder::BoundaryType QTextBoundaryFinder::type();
extern void _ZNK19QTextBoundaryFinder4typeEv(void* qthis); // 2
  // proto:  void QTextBoundaryFinder::toStart();
extern void _ZN19QTextBoundaryFinder7toStartEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// toNextBoundary()
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
    // invoke: int toNextBoundary()
    C._ZN19QTextBoundaryFinder14toNextBoundaryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toNextBoundary", args)
  }

}

// position()
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
    // invoke: int position()
    C._ZNK19QTextBoundaryFinder8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "position", args)
  }

}

// string()
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
    // invoke: QString string()
    C._ZNK19QTextBoundaryFinder6stringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "string", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    C._ZNK19QTextBoundaryFinder7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isValid", args)
  }

}

// boundaryReasons()
func (this *QTextBoundaryFinder) boundaryReasons(args ...interface{}) () {
  // boundaryReasons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder15boundaryReasonsEv
    // invoke: BoundaryReasons boundaryReasons()
    C._ZNK19QTextBoundaryFinder15boundaryReasonsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "boundaryReasons", args)
  }

}

// toEnd()
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
    // invoke: void toEnd()
    C._ZN19QTextBoundaryFinder5toEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toEnd", args)
  }

}

// ~QTextBoundaryFinder()
func (this *QTextBoundaryFinder) FreeQTextBoundaryFinder(args ...interface{}) () {
  // ~QTextBoundaryFinder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinderD0Ev
    // invoke: void ~QTextBoundaryFinder()
    C._ZN19QTextBoundaryFinderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "~QTextBoundaryFinder", args)
  }

}

// toPreviousBoundary()
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
    // invoke: int toPreviousBoundary()
    C._ZN19QTextBoundaryFinder18toPreviousBoundaryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toPreviousBoundary", args)
  }

}

// QTextBoundaryFinder()
func NewQTextBoundaryFinder(args ...interface{}) QTextBoundaryFinder {
  // QTextBoundaryFinder()
  // QTextBoundaryFinder(const class QTextBoundaryFinder &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextBoundaryFinder{}) // "const QTextBoundaryFinder &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextBoundaryFinderC1Ev
    // invoke: void QTextBoundaryFinder()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QTextBoundaryFinderC2Ev(qthis)
  case 1:
    // invoke: _ZN19QTextBoundaryFinderC1ERKS_
    // invoke: void QTextBoundaryFinder(const class QTextBoundaryFinder &)
    var arg0 = args[0].(QTextBoundaryFinder).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN19QTextBoundaryFinderC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "QTextBoundaryFinder", args)
  }

  return QTextBoundaryFinder{}
}

// isAtBoundary()
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
    // invoke: bool isAtBoundary()
    C._ZNK19QTextBoundaryFinder12isAtBoundaryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isAtBoundary", args)
  }

}

// setPosition(int)
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
    // invoke: void setPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN19QTextBoundaryFinder11setPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "setPosition", args)
  }

}

// type()
func (this *QTextBoundaryFinder) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QTextBoundaryFinder4typeEv
    // invoke: QTextBoundaryFinder::BoundaryType type()
    C._ZNK19QTextBoundaryFinder4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "type", args)
  }

}

// toStart()
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
    // invoke: void toStart()
    C._ZN19QTextBoundaryFinder7toStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toStart", args)
  }

}

// <= body block end

