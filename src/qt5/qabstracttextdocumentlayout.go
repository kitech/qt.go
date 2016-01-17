package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtGui/qabstracttextdocumentlayout.h
// dst-file: /src/gui/qabstracttextdocumentlayout.go
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
  // proto:  void QAbstractTextDocumentLayout::unregisterHandler(int objectType, QObject * component);
extern void _ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QAbstractTextDocumentLayout::setPaintDevice(QPaintDevice * device);
extern void _ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QString QAbstractTextDocumentLayout::anchorAt(const QPointF & pos);
extern void _ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QAbstractTextDocumentLayout::QAbstractTextDocumentLayout(QTextDocument * doc);
extern void _ZN27QAbstractTextDocumentLayoutC2EP13QTextDocument(void* qthis, void* arg0); // 3
  // proto:  QPaintDevice * QAbstractTextDocumentLayout::paintDevice();
extern void _ZNK27QAbstractTextDocumentLayout11paintDeviceEv(void* qthis); // 4
  // proto:  QTextDocument * QAbstractTextDocumentLayout::document();
extern void _ZNK27QAbstractTextDocumentLayout8documentEv(void* qthis); // 4
  // proto:  void QAbstractTextDocumentLayout::registerHandler(int objectType, QObject * component);
extern void _ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  const QMetaObject * QAbstractTextDocumentLayout::metaObject();
extern void _ZNK27QAbstractTextDocumentLayout10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractTextDocumentLayout::~QAbstractTextDocumentLayout();
extern void _ZN27QAbstractTextDocumentLayoutD2Ev(void* qthis); // 4
  // proto:  QTextObjectInterface * QAbstractTextDocumentLayout::handlerForObject(int objectType);
extern void _ZNK27QAbstractTextDocumentLayout16handlerForObjectEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QTextObjectInterface)=8
type QTextObjectInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractTextDocumentLayout)=1
type QAbstractTextDocumentLayout struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _updateBlock QAbstractTextDocumentLayout_updateBlock_signal;
//  _pageCountChanged QAbstractTextDocumentLayout_pageCountChanged_signal;
//  _update QAbstractTextDocumentLayout_update_signal;
//  _documentSizeChanged QAbstractTextDocumentLayout_documentSizeChanged_signal;
}

// unregisterHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) unregisterHandler(args ...interface{}) () {
  // unregisterHandler(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject
    // invoke: void unregisterHandler(int, class QObject *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "unregisterHandler", args)
  }

}

// setPaintDevice(class QPaintDevice *)
func (this *QAbstractTextDocumentLayout) setPaintDevice(args ...interface{}) () {
  // setPaintDevice(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice
    // invoke: void setPaintDevice(class QPaintDevice *)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "setPaintDevice", args)
  }

}

// anchorAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) anchorAt(args ...interface{}) () {
  // anchorAt(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPointF{}) // "const QPointF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF
    // invoke: QString anchorAt(const class QPointF &)
    var arg0 = args[0].(QPointF).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "anchorAt", args)
  }

}

// QAbstractTextDocumentLayout(class QTextDocument *)
func NewQAbstractTextDocumentLayout(args ...interface{}) QAbstractTextDocumentLayout {
  // QAbstractTextDocumentLayout(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayoutC1EP13QTextDocument
    // invoke: void QAbstractTextDocumentLayout(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN27QAbstractTextDocumentLayoutC2EP13QTextDocument(qthis, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "QAbstractTextDocumentLayout", args)
  }

  return QAbstractTextDocumentLayout{}
}

// paintDevice()
func (this *QAbstractTextDocumentLayout) paintDevice(args ...interface{}) () {
  // paintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout11paintDeviceEv
    // invoke: QPaintDevice * paintDevice()
    C._ZNK27QAbstractTextDocumentLayout11paintDeviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "paintDevice", args)
  }

}

// document()
func (this *QAbstractTextDocumentLayout) document(args ...interface{}) () {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout8documentEv
    // invoke: QTextDocument * document()
    C._ZNK27QAbstractTextDocumentLayout8documentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "document", args)
  }

}

// registerHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) registerHandler(args ...interface{}) () {
  // registerHandler(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject
    // invoke: void registerHandler(int, class QObject *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "registerHandler", args)
  }

}

// metaObject()
func (this *QAbstractTextDocumentLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK27QAbstractTextDocumentLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "metaObject", args)
  }

}

// ~QAbstractTextDocumentLayout()
func (this *QAbstractTextDocumentLayout) FreeQAbstractTextDocumentLayout(args ...interface{}) () {
  // ~QAbstractTextDocumentLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayoutD0Ev
    // invoke: void ~QAbstractTextDocumentLayout()
    C._ZN27QAbstractTextDocumentLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "~QAbstractTextDocumentLayout", args)
  }

}

// handlerForObject(int)
func (this *QAbstractTextDocumentLayout) handlerForObject(args ...interface{}) () {
  // handlerForObject(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout16handlerForObjectEi
    // invoke: QTextObjectInterface * handlerForObject(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK27QAbstractTextDocumentLayout16handlerForObjectEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "handlerForObject", args)
  }

}

// <= body block end

