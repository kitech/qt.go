package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qframe.h
// dst-file: /src/widgets/qframe.go
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
  // proto:  void QFrame::setFrameRect(const QRect & );
extern void _ZN6QFrame12setFrameRectERK5QRect(void* qthis, void* arg0);
  // proto:  int QFrame::lineWidth();
extern void _ZNK6QFrame9lineWidthEv(void* qthis);
  // proto:  void QFrame::setFrameStyle(int );
extern void _ZN6QFrame13setFrameStyleEi(void* qthis, int32_t arg0);
  // proto:  QRect QFrame::frameRect();
extern void _ZNK6QFrame9frameRectEv(void* qthis);
  // proto:  QSize QFrame::sizeHint();
extern void _ZNK6QFrame8sizeHintEv(void* qthis);
  // proto:  void QFrame::QFrame(const QFrame & );
extern void* dector_ZN6QFrameC1ERKS_(void* arg0);
extern void _ZN6QFrameC1ERKS_(void* qthis, void* arg0);
  // proto:  int QFrame::frameStyle();
extern void _ZNK6QFrame10frameStyleEv(void* qthis);
  // proto:  int QFrame::midLineWidth();
extern void _ZNK6QFrame12midLineWidthEv(void* qthis);
  // proto:  void QFrame::setLineWidth(int );
extern void _ZN6QFrame12setLineWidthEi(void* qthis, int32_t arg0);
  // proto:  void QFrame::setMidLineWidth(int );
extern void _ZN6QFrame15setMidLineWidthEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QFrame::metaObject();
extern void _ZNK6QFrame10metaObjectEv(void* qthis);
  // proto:  int QFrame::frameWidth();
extern void _ZNK6QFrame10frameWidthEv(void* qthis);
  // proto:  void QFrame::~QFrame();
extern void _ZN6QFrameD0Ev(void* qthis);
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

// class sizeof(QFrame)=1
type QFrame struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QFrame::setFrameRect(const QRect & );
func (this *QFrame) setFrameRect(args ...interface{}) () {
  // setFrameRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame12setFrameRectERK5QRect
    // invoke: void setFrameRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QFrame12setFrameRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setFrameRect", args)
  }

}

  // proto:  int QFrame::lineWidth();
func (this *QFrame) lineWidth(args ...interface{}) () {
  // lineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame9lineWidthEv
    // invoke: int lineWidth()
    C._ZNK6QFrame9lineWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "lineWidth", args)
  }

}

  // proto:  void QFrame::setFrameStyle(int );
func (this *QFrame) setFrameStyle(args ...interface{}) () {
  // setFrameStyle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame13setFrameStyleEi
    // invoke: void setFrameStyle(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QFrame13setFrameStyleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setFrameStyle", args)
  }

}

  // proto:  QRect QFrame::frameRect();
func (this *QFrame) frameRect(args ...interface{}) () {
  // frameRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame9frameRectEv
    // invoke: QRect frameRect()
    C._ZNK6QFrame9frameRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "frameRect", args)
  }

}

  // proto:  QSize QFrame::sizeHint();
func (this *QFrame) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK6QFrame8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "sizeHint", args)
  }

}

  // proto:  void QFrame::QFrame(const QFrame & );
func NewQFrame(args ...interface{}) QFrame {
  return QFrame{}
}

  // proto:  int QFrame::frameStyle();
func (this *QFrame) frameStyle(args ...interface{}) () {
  // frameStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10frameStyleEv
    // invoke: int frameStyle()
    C._ZNK6QFrame10frameStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "frameStyle", args)
  }

}

  // proto:  int QFrame::midLineWidth();
func (this *QFrame) midLineWidth(args ...interface{}) () {
  // midLineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame12midLineWidthEv
    // invoke: int midLineWidth()
    C._ZNK6QFrame12midLineWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "midLineWidth", args)
  }

}

  // proto:  void QFrame::setLineWidth(int );
func (this *QFrame) setLineWidth(args ...interface{}) () {
  // setLineWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame12setLineWidthEi
    // invoke: void setLineWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QFrame12setLineWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setLineWidth", args)
  }

}

  // proto:  void QFrame::setMidLineWidth(int );
func (this *QFrame) setMidLineWidth(args ...interface{}) () {
  // setMidLineWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QFrame15setMidLineWidthEi
    // invoke: void setMidLineWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QFrame15setMidLineWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFrame", "setMidLineWidth", args)
  }

}

  // proto:  const QMetaObject * QFrame::metaObject();
func (this *QFrame) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK6QFrame10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "metaObject", args)
  }

}

  // proto:  int QFrame::frameWidth();
func (this *QFrame) frameWidth(args ...interface{}) () {
  // frameWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QFrame10frameWidthEv
    // invoke: int frameWidth()
    C._ZNK6QFrame10frameWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFrame", "frameWidth", args)
  }

}

  // proto:  void QFrame::~QFrame();
func (this *QFrame) FreeQFrame(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFrame", "~QFrame", args)
  }

}

// <= body block end

