package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qsize.h
// dst-file: /src/core/qsize.go
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
  // proto:  QSize QSize::boundedTo(const QSize & );
extern void* C_ZNK5QSize9boundedToERKS_(void* qthis, void* arg0); // 2
  // proto:  void QSize::QSize(int w, int h);
extern void* C_ZN5QSizeC2Eii(int32_t arg0, int32_t arg1); // 1
  // proto:  void QSize::QSize();
extern void* C_ZN5QSizeC2Ev(); // 1
  // proto:  QSize QSize::transposed();
extern void* C_ZNK5QSize10transposedEv(void* qthis); // 2
  // proto:  int QSize::width();
extern int32_t C_ZNK5QSize5widthEv(void* qthis); // 2
  // proto:  bool QSize::isValid();
extern bool C_ZNK5QSize7isValidEv(void* qthis); // 2
  // proto:  void QSize::transpose();
extern void C_ZN5QSize9transposeEv(void* qthis); // 4
  // proto:  int QSize::height();
extern int32_t C_ZNK5QSize6heightEv(void* qthis); // 2
  // proto:  bool QSize::isNull();
extern bool C_ZNK5QSize6isNullEv(void* qthis); // 2
  // proto:  QSize QSize::expandedTo(const QSize & );
extern void* C_ZNK5QSize10expandedToERKS_(void* qthis, void* arg0); // 2
  // proto:  bool QSize::isEmpty();
extern bool C_ZNK5QSize7isEmptyEv(void* qthis); // 2
  // proto:  int & QSize::rwidth();
extern void C_ZN5QSize6rwidthEv(void* qthis); // 2
  // proto:  int & QSize::rheight();
extern void C_ZN5QSize7rheightEv(void* qthis); // 2
  // proto:  void QSize::setWidth(int w);
extern void C_ZN5QSize8setWidthEi(void* qthis, int32_t arg0); // 2
  // proto:  void QSize::setHeight(int h);
extern void C_ZN5QSize9setHeightEi(void* qthis, int32_t arg0); // 2
  // proto:  qreal & QSizeF::rwidth();
extern void C_ZN6QSizeF6rwidthEv(void* qthis); // 2
  // proto:  QSizeF QSizeF::transposed();
extern void* C_ZNK6QSizeF10transposedEv(void* qthis); // 2
  // proto:  bool QSizeF::isNull();
extern bool C_ZNK6QSizeF6isNullEv(void* qthis); // 2
  // proto:  bool QSizeF::isValid();
extern bool C_ZNK6QSizeF7isValidEv(void* qthis); // 2
  // proto:  QSizeF QSizeF::boundedTo(const QSizeF & );
extern void* C_ZNK6QSizeF9boundedToERKS_(void* qthis, void* arg0); // 2
  // proto:  void QSizeF::transpose();
extern void C_ZN6QSizeF9transposeEv(void* qthis); // 4
  // proto:  QSizeF QSizeF::expandedTo(const QSizeF & );
extern void* C_ZNK6QSizeF10expandedToERKS_(void* qthis, void* arg0); // 2
  // proto:  qreal QSizeF::height();
extern double C_ZNK6QSizeF6heightEv(void* qthis); // 2
  // proto:  qreal QSizeF::width();
extern double C_ZNK6QSizeF5widthEv(void* qthis); // 2
  // proto:  void QSizeF::QSizeF();
extern void* C_ZN6QSizeFC2Ev(); // 1
  // proto:  void QSizeF::QSizeF(const QSize & sz);
extern void* C_ZN6QSizeFC2ERK5QSize(void* arg0); // 1
  // proto:  void QSizeF::QSizeF(qreal w, qreal h);
extern void* C_ZN6QSizeFC2Edd(double arg0, double arg1); // 1
  // proto:  bool QSizeF::isEmpty();
extern bool C_ZNK6QSizeF7isEmptyEv(void* qthis); // 2
  // proto:  void QSizeF::setHeight(qreal h);
extern void C_ZN6QSizeF9setHeightEd(void* qthis, double arg0); // 2
  // proto:  qreal & QSizeF::rheight();
extern void C_ZN6QSizeF7rheightEv(void* qthis); // 2
  // proto:  void QSizeF::setWidth(qreal w);
extern void C_ZN6QSizeF8setWidthEd(void* qthis, double arg0); // 2
  // proto:  QSize QSizeF::toSize();
extern void* C_ZNK6QSizeF6toSizeEv(void* qthis); // 2
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

// class sizeof(QSize)=8
type QSize struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QSizeF)=16
type QSizeF struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// boundedTo(const class QSize &)
func (this *QSize) Boundedto(args ...interface{}) (ret interface{}) {
  // boundedTo(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize9boundedToERKS_
    // invoke: QSize boundedTo(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QSize9boundedToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "boundedTo", args)
  }

  return
}

// QSize(int, int)
func NewQSize(args ...interface{}) *QSize {
  // QSize(int, int)
  // QSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSizeC1Eii
    // invoke: void QSize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QSizeC2Eii(arg0, arg1)
    return &QSize{qclsinst:qthis}
  case 1:
    // invoke: _ZN5QSizeC1Ev
    // invoke: void QSize()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QSizeC2Ev()
    return &QSize{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSize", "QSize", args)
  }

  return nil // QSize{qclsinst:qthis}
}

// transposed()
func (this *QSize) Transposed(args ...interface{}) (ret interface{}) {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize10transposedEv
    // invoke: QSize transposed()
    var ret0 = C.C_ZNK5QSize10transposedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "transposed", args)
  }

  return
}

// width()
func (this *QSize) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize5widthEv
    // invoke: int width()
    var ret0 = C.C_ZNK5QSize5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "width", args)
  }

  return
}

// isValid()
func (this *QSize) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK5QSize7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "isValid", args)
  }

  return
}

// transpose()
func (this *QSize) Transpose(args ...interface{}) () {
  // transpose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize9transposeEv
    // invoke: void transpose()
    C.C_ZN5QSize9transposeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSize", "transpose", args)
  }

  return
}

// height()
func (this *QSize) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize6heightEv
    // invoke: int height()
    var ret0 = C.C_ZNK5QSize6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "height", args)
  }

  return
}

// isNull()
func (this *QSize) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK5QSize6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "isNull", args)
  }

  return
}

// expandedTo(const class QSize &)
func (this *QSize) Expandedto(args ...interface{}) (ret interface{}) {
  // expandedTo(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize10expandedToERKS_
    // invoke: QSize expandedTo(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK5QSize10expandedToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "expandedTo", args)
  }

  return
}

// isEmpty()
func (this *QSize) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK5QSize7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSize", "isEmpty", args)
  }

  return
}

// rwidth()
func (this *QSize) Rwidth(args ...interface{}) () {
  // rwidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize6rwidthEv
    // invoke: int & rwidth()
    C.C_ZN5QSize6rwidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSize", "rwidth", args)
  }

  return
}

// rheight()
func (this *QSize) Rheight(args ...interface{}) () {
  // rheight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize7rheightEv
    // invoke: int & rheight()
    C.C_ZN5QSize7rheightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSize", "rheight", args)
  }

  return
}

// setWidth(int)
func (this *QSize) Setwidth(args ...interface{}) () {
  // setWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize8setWidthEi
    // invoke: void setWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QSize8setWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSize", "setWidth", args)
  }

  return
}

// setHeight(int)
func (this *QSize) Setheight(args ...interface{}) () {
  // setHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize9setHeightEi
    // invoke: void setHeight(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN5QSize9setHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSize", "setHeight", args)
  }

  return
}

// rwidth()
func (this *QSizeF) Rwidth(args ...interface{}) () {
  // rwidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF6rwidthEv
    // invoke: qreal & rwidth()
    C.C_ZN6QSizeF6rwidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeF", "rwidth", args)
  }

  return
}

// transposed()
func (this *QSizeF) Transposed(args ...interface{}) (ret interface{}) {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF10transposedEv
    // invoke: QSizeF transposed()
    var ret0 = C.C_ZNK6QSizeF10transposedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "transposed", args)
  }

  return
}

// isNull()
func (this *QSizeF) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK6QSizeF6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "isNull", args)
  }

  return
}

// isValid()
func (this *QSizeF) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK6QSizeF7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "isValid", args)
  }

  return
}

// boundedTo(const class QSizeF &)
func (this *QSizeF) Boundedto(args ...interface{}) (ret interface{}) {
  // boundedTo(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF9boundedToERKS_
    // invoke: QSizeF boundedTo(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QSizeF9boundedToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "boundedTo", args)
  }

  return
}

// transpose()
func (this *QSizeF) Transpose(args ...interface{}) () {
  // transpose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF9transposeEv
    // invoke: void transpose()
    C.C_ZN6QSizeF9transposeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeF", "transpose", args)
  }

  return
}

// expandedTo(const class QSizeF &)
func (this *QSizeF) Expandedto(args ...interface{}) (ret interface{}) {
  // expandedTo(const class QSizeF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizeF{}) // "const QSizeF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF10expandedToERKS_
    // invoke: QSizeF expandedTo(const class QSizeF &)
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QSizeF10expandedToERKS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "expandedTo", args)
  }

  return
}

// height()
func (this *QSizeF) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF6heightEv
    // invoke: qreal height()
    var ret0 = C.C_ZNK6QSizeF6heightEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "height", args)
  }

  return
}

// width()
func (this *QSizeF) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF5widthEv
    // invoke: qreal width()
    var ret0 = C.C_ZNK6QSizeF5widthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "width", args)
  }

  return
}

// QSizeF()
func NewQSizeF(args ...interface{}) *QSizeF {
  // QSizeF()
  // QSizeF(const class QSize &)
  // QSizeF(qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[2][1] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeFC1Ev
    // invoke: void QSizeF()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QSizeFC2Ev()
    return &QSizeF{qclsinst:qthis}
  case 1:
    // invoke: _ZN6QSizeFC1ERK5QSize
    // invoke: void QSizeF(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QSizeFC2ERK5QSize(arg0)
    return &QSizeF{qclsinst:qthis}
  case 2:
    // invoke: _ZN6QSizeFC1Edd
    // invoke: void QSizeF(qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QSizeFC2Edd(arg0, arg1)
    return &QSizeF{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QSizeF", "QSizeF", args)
  }

  return nil // QSizeF{qclsinst:qthis}
}

// isEmpty()
func (this *QSizeF) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK6QSizeF7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "isEmpty", args)
  }

  return
}

// setHeight(qreal)
func (this *QSizeF) Setheight(args ...interface{}) () {
  // setHeight(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF9setHeightEd
    // invoke: void setHeight(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QSizeF9setHeightEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizeF", "setHeight", args)
  }

  return
}

// rheight()
func (this *QSizeF) Rheight(args ...interface{}) () {
  // rheight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF7rheightEv
    // invoke: qreal & rheight()
    C.C_ZN6QSizeF7rheightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSizeF", "rheight", args)
  }

  return
}

// setWidth(qreal)
func (this *QSizeF) Setwidth(args ...interface{}) () {
  // setWidth(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF8setWidthEd
    // invoke: void setWidth(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QSizeF8setWidthEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSizeF", "setWidth", args)
  }

  return
}

// toSize()
func (this *QSizeF) Tosize(args ...interface{}) (ret interface{}) {
  // toSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF6toSizeEv
    // invoke: QSize toSize()
    var ret0 = C.C_ZNK6QSizeF6toSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QSizeF", "toSize", args)
  }

  return
}

// <= body block end

