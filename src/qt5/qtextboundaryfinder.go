package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern int32_t C_ZN19QTextBoundaryFinder14toNextBoundaryEv(void* qthis); // 4
  // proto:  int QTextBoundaryFinder::position();
extern int32_t C_ZNK19QTextBoundaryFinder8positionEv(void* qthis); // 4
  // proto:  QString QTextBoundaryFinder::string();
extern void* C_ZNK19QTextBoundaryFinder6stringEv(void* qthis); // 4
  // proto:  bool QTextBoundaryFinder::isValid();
extern bool C_ZNK19QTextBoundaryFinder7isValidEv(void* qthis); // 2
  // proto:  BoundaryReasons QTextBoundaryFinder::boundaryReasons();
extern void C_ZNK19QTextBoundaryFinder15boundaryReasonsEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::toEnd();
extern void C_ZN19QTextBoundaryFinder5toEndEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::~QTextBoundaryFinder();
extern void C_ZN19QTextBoundaryFinderD2Ev(void* qthis); // 4
  // proto:  int QTextBoundaryFinder::toPreviousBoundary();
extern int32_t C_ZN19QTextBoundaryFinder18toPreviousBoundaryEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::QTextBoundaryFinder();
extern void* C_ZN19QTextBoundaryFinderC2Ev(); // 3
  // proto:  void QTextBoundaryFinder::QTextBoundaryFinder(const QTextBoundaryFinder & other);
extern void* C_ZN19QTextBoundaryFinderC2ERKS_(void* arg0); // 3
  // proto:  bool QTextBoundaryFinder::isAtBoundary();
extern bool C_ZNK19QTextBoundaryFinder12isAtBoundaryEv(void* qthis); // 4
  // proto:  void QTextBoundaryFinder::setPosition(int position);
extern void C_ZN19QTextBoundaryFinder11setPositionEi(void* qthis, int32_t arg0); // 4
  // proto:  QTextBoundaryFinder::BoundaryType QTextBoundaryFinder::type();
extern void C_ZNK19QTextBoundaryFinder4typeEv(void* qthis); // 2
  // proto:  void QTextBoundaryFinder::toStart();
extern void C_ZN19QTextBoundaryFinder7toStartEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// toNextBoundary()
func (this *QTextBoundaryFinder) Tonextboundary(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN19QTextBoundaryFinder14toNextBoundaryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toNextBoundary", args)
  }

  return
}

// position()
func (this *QTextBoundaryFinder) Position(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextBoundaryFinder8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "position", args)
  }

  return
}

// string()
func (this *QTextBoundaryFinder) String(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextBoundaryFinder6stringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "string", args)
  }

  return
}

// isValid()
func (this *QTextBoundaryFinder) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextBoundaryFinder7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isValid", args)
  }

  return
}

// boundaryReasons()
func (this *QTextBoundaryFinder) Boundaryreasons(args ...interface{}) () {
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
    C.C_ZNK19QTextBoundaryFinder15boundaryReasonsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "boundaryReasons", args)
  }

  return
}

// toEnd()
func (this *QTextBoundaryFinder) Toend(args ...interface{}) () {
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
    C.C_ZN19QTextBoundaryFinder5toEndEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toEnd", args)
  }

  return
}

// ~QTextBoundaryFinder()
func (this *QTextBoundaryFinder) Freeqtextboundaryfinder(args ...interface{}) () {
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
    C.C_ZN19QTextBoundaryFinderD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "~QTextBoundaryFinder", args)
  }

  return
}

// toPreviousBoundary()
func (this *QTextBoundaryFinder) Topreviousboundary(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN19QTextBoundaryFinder18toPreviousBoundaryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toPreviousBoundary", args)
  }

  return
}

// QTextBoundaryFinder()
func NewQTextBoundaryFinder(args ...interface{}) *QTextBoundaryFinder {
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
    qthis = C.C_ZN19QTextBoundaryFinderC2Ev()
    return &QTextBoundaryFinder{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QTextBoundaryFinderC1ERKS_
    // invoke: void QTextBoundaryFinder(const class QTextBoundaryFinder &)
    var arg0 = args[0].(*QTextBoundaryFinder).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QTextBoundaryFinderC2ERKS_(arg0)
    return &QTextBoundaryFinder{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "QTextBoundaryFinder", args)
  }

  return nil // QTextBoundaryFinder{Qclsinst:qthis}
}

// isAtBoundary()
func (this *QTextBoundaryFinder) Isatboundary(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextBoundaryFinder12isAtBoundaryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "isAtBoundary", args)
  }

  return
}

// setPosition(int)
func (this *QTextBoundaryFinder) Setposition(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN19QTextBoundaryFinder11setPositionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "setPosition", args)
  }

  return
}

// type()
func (this *QTextBoundaryFinder) Type_(args ...interface{}) () {
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
    C.C_ZNK19QTextBoundaryFinder4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "type", args)
  }

  return
}

// toStart()
func (this *QTextBoundaryFinder) Tostart(args ...interface{}) () {
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
    C.C_ZN19QTextBoundaryFinder7toStartEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextBoundaryFinder", "toStart", args)
  }

  return
}

// <= body block end

