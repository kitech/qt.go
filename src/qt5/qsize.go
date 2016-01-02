package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  QSize QSize::boundedTo(const QSize & );
extern void _ZNK5QSize9boundedToERKS_(void* qthis, void* arg0);
  // proto:  bool QSize::isValid();
extern void _ZNK5QSize7isValidEv(void* qthis);
  // proto:  bool QSize::isNull();
extern void _ZNK5QSize6isNullEv(void* qthis);
  // proto:  void QSize::QSize();
extern void* dector_ZN5QSizeC1Ev();
extern void _ZN5QSizeC1Ev(void* qthis);
  // proto:  QSize QSize::expandedTo(const QSize & );
extern void _ZNK5QSize10expandedToERKS_(void* qthis, void* arg0);
  // proto:  int QSize::height();
extern void _ZNK5QSize6heightEv(void* qthis);
  // proto:  int & QSize::rheight();
extern void demth_ZN5QSize7rheightEv(void* qthis);
  // proto:  void QSize::QSize(int w, int h);
extern void* dector_ZN5QSizeC1Eii(int arg0, int arg1);
extern void _ZN5QSizeC1Eii(void* qthis, int arg0, int arg1);
  // proto:  int QSize::width();
extern void _ZNK5QSize5widthEv(void* qthis);
  // proto:  QSize QSize::transposed();
extern void _ZNK5QSize10transposedEv(void* qthis);
  // proto:  int & QSize::rwidth();
extern void demth_ZN5QSize6rwidthEv(void* qthis);
  // proto:  void QSize::setHeight(int h);
extern void demth_ZN5QSize9setHeightEi(void* qthis, int arg0);
  // proto:  bool QSize::isEmpty();
extern void _ZNK5QSize7isEmptyEv(void* qthis);
  // proto:  void QSize::setWidth(int w);
extern void demth_ZN5QSize8setWidthEi(void* qthis, int arg0);
  // proto:  void QSize::transpose();
extern void _ZN5QSize9transposeEv(void* qthis);
  // proto:  qreal & QSizeF::rheight();
extern void demth_ZN6QSizeF7rheightEv(void* qthis);
  // proto:  qreal & QSizeF::rwidth();
extern void demth_ZN6QSizeF6rwidthEv(void* qthis);
  // proto:  QSizeF QSizeF::transposed();
extern void _ZNK6QSizeF10transposedEv(void* qthis);
  // proto:  bool QSizeF::isValid();
extern void _ZNK6QSizeF7isValidEv(void* qthis);
  // proto:  void QSizeF::setHeight(qreal h);
extern void demth_ZN6QSizeF9setHeightEd(void* qthis, double arg0);
  // proto:  void QSizeF::QSizeF();
extern void* dector_ZN6QSizeFC1Ev();
extern void _ZN6QSizeFC1Ev(void* qthis);
  // proto:  qreal QSizeF::width();
extern void _ZNK6QSizeF5widthEv(void* qthis);
  // proto:  bool QSizeF::isNull();
extern void demth_ZNK6QSizeF6isNullEv(void* qthis);
  // proto:  QSizeF QSizeF::boundedTo(const QSizeF & );
extern void _ZNK6QSizeF9boundedToERKS_(void* qthis, void* arg0);
  // proto:  qreal QSizeF::height();
extern void _ZNK6QSizeF6heightEv(void* qthis);
  // proto:  void QSizeF::transpose();
extern void _ZN6QSizeF9transposeEv(void* qthis);
  // proto:  void QSizeF::QSizeF(const QSize & sz);
extern void* dector_ZN6QSizeFC1ERK5QSize(void* arg0);
extern void _ZN6QSizeFC1ERK5QSize(void* qthis, void* arg0);
  // proto:  QSizeF QSizeF::expandedTo(const QSizeF & );
extern void _ZNK6QSizeF10expandedToERKS_(void* qthis, void* arg0);
  // proto:  bool QSizeF::isEmpty();
extern void _ZNK6QSizeF7isEmptyEv(void* qthis);
  // proto:  void QSizeF::setWidth(qreal w);
extern void demth_ZN6QSizeF8setWidthEd(void* qthis, double arg0);
  // proto:  QSize QSizeF::toSize();
extern void _ZNK6QSizeF6toSizeEv(void* qthis);
  // proto:  void QSizeF::QSizeF(qreal w, qreal h);
extern void* dector_ZN6QSizeFC1Edd(double arg0, double arg1);
extern void _ZN6QSizeFC1Edd(void* qthis, double arg0, double arg1);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QSizeF)=16
type QSizeF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QSize QSize::boundedTo(const QSize & );
func (this *QSize) boundedTo(args ...interface{}) () {
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
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSize", "boundedTo", args)
  }

}

  // proto:  bool QSize::isValid();
func (this *QSize) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize7isValidEv
  default:
    qtrt.ErrorResolve("QSize", "isValid", args)
  }

}

  // proto:  bool QSize::isNull();
func (this *QSize) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize6isNullEv
  default:
    qtrt.ErrorResolve("QSize", "isNull", args)
  }

}

  // proto:  void QSize::QSize();
func NewQSize(args ...interface{}) QSize {
  return QSize{}
}

  // proto:  QSize QSize::expandedTo(const QSize & );
func (this *QSize) expandedTo(args ...interface{}) () {
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
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSize", "expandedTo", args)
  }

}

  // proto:  int QSize::height();
func (this *QSize) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize6heightEv
  default:
    qtrt.ErrorResolve("QSize", "height", args)
  }

}

  // proto:  int & QSize::rheight();
func (this *QSize) rheight(args ...interface{}) () {
  // rheight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize7rheightEv
  default:
    qtrt.ErrorResolve("QSize", "rheight", args)
  }

}

  // proto:  int QSize::width();
func (this *QSize) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize5widthEv
  default:
    qtrt.ErrorResolve("QSize", "width", args)
  }

}

  // proto:  QSize QSize::transposed();
func (this *QSize) transposed(args ...interface{}) () {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize10transposedEv
  default:
    qtrt.ErrorResolve("QSize", "transposed", args)
  }

}

  // proto:  int & QSize::rwidth();
func (this *QSize) rwidth(args ...interface{}) () {
  // rwidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize6rwidthEv
  default:
    qtrt.ErrorResolve("QSize", "rwidth", args)
  }

}

  // proto:  void QSize::setHeight(int h);
func (this *QSize) setHeight(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSize", "setHeight", args)
  }

}

  // proto:  bool QSize::isEmpty();
func (this *QSize) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QSize7isEmptyEv
  default:
    qtrt.ErrorResolve("QSize", "isEmpty", args)
  }

}

  // proto:  void QSize::setWidth(int w);
func (this *QSize) setWidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSize", "setWidth", args)
  }

}

  // proto:  void QSize::transpose();
func (this *QSize) transpose(args ...interface{}) () {
  // transpose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QSize9transposeEv
  default:
    qtrt.ErrorResolve("QSize", "transpose", args)
  }

}

  // proto:  qreal & QSizeF::rheight();
func (this *QSizeF) rheight(args ...interface{}) () {
  // rheight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF7rheightEv
  default:
    qtrt.ErrorResolve("QSizeF", "rheight", args)
  }

}

  // proto:  qreal & QSizeF::rwidth();
func (this *QSizeF) rwidth(args ...interface{}) () {
  // rwidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF6rwidthEv
  default:
    qtrt.ErrorResolve("QSizeF", "rwidth", args)
  }

}

  // proto:  QSizeF QSizeF::transposed();
func (this *QSizeF) transposed(args ...interface{}) () {
  // transposed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF10transposedEv
  default:
    qtrt.ErrorResolve("QSizeF", "transposed", args)
  }

}

  // proto:  bool QSizeF::isValid();
func (this *QSizeF) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF7isValidEv
  default:
    qtrt.ErrorResolve("QSizeF", "isValid", args)
  }

}

  // proto:  void QSizeF::setHeight(qreal h);
func (this *QSizeF) setHeight(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizeF", "setHeight", args)
  }

}

  // proto:  void QSizeF::QSizeF();
func NewQSizeF(args ...interface{}) QSizeF {
  return QSizeF{}
}

  // proto:  qreal QSizeF::width();
func (this *QSizeF) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF5widthEv
  default:
    qtrt.ErrorResolve("QSizeF", "width", args)
  }

}

  // proto:  bool QSizeF::isNull();
func (this *QSizeF) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF6isNullEv
  default:
    qtrt.ErrorResolve("QSizeF", "isNull", args)
  }

}

  // proto:  QSizeF QSizeF::boundedTo(const QSizeF & );
func (this *QSizeF) boundedTo(args ...interface{}) () {
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
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizeF", "boundedTo", args)
  }

}

  // proto:  qreal QSizeF::height();
func (this *QSizeF) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF6heightEv
  default:
    qtrt.ErrorResolve("QSizeF", "height", args)
  }

}

  // proto:  void QSizeF::transpose();
func (this *QSizeF) transpose(args ...interface{}) () {
  // transpose()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QSizeF9transposeEv
  default:
    qtrt.ErrorResolve("QSizeF", "transpose", args)
  }

}

  // proto:  QSizeF QSizeF::expandedTo(const QSizeF & );
func (this *QSizeF) expandedTo(args ...interface{}) () {
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
    var arg0 = args[0].(QSizeF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizeF", "expandedTo", args)
  }

}

  // proto:  bool QSizeF::isEmpty();
func (this *QSizeF) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF7isEmptyEv
  default:
    qtrt.ErrorResolve("QSizeF", "isEmpty", args)
  }

}

  // proto:  void QSizeF::setWidth(qreal w);
func (this *QSizeF) setWidth(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QSizeF", "setWidth", args)
  }

}

  // proto:  QSize QSizeF::toSize();
func (this *QSizeF) toSize(args ...interface{}) () {
  // toSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QSizeF6toSizeEv
  default:
    qtrt.ErrorResolve("QSizeF", "toSize", args)
  }

}

// <= body block end

