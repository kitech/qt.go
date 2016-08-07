package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QAbstractTextDocumentLayout::unregisterHandler(int objectType, QObject * component);
extern void C_ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QAbstractTextDocumentLayout::setPaintDevice(QPaintDevice * device);
extern void C_ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice(void* qthis, void* arg0); // 4
  // proto:  QString QAbstractTextDocumentLayout::anchorAt(const QPointF & pos);
extern void* C_ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF(void* qthis, void* arg0); // 4
  // proto:  void QAbstractTextDocumentLayout::QAbstractTextDocumentLayout(QTextDocument * doc);
extern void* C_ZN27QAbstractTextDocumentLayoutC2EP13QTextDocument(void* arg0); // 3
  // proto:  QPaintDevice * QAbstractTextDocumentLayout::paintDevice();
extern void* C_ZNK27QAbstractTextDocumentLayout11paintDeviceEv(void* qthis); // 4
  // proto:  QTextDocument * QAbstractTextDocumentLayout::document();
extern void* C_ZNK27QAbstractTextDocumentLayout8documentEv(void* qthis); // 4
  // proto:  void QAbstractTextDocumentLayout::registerHandler(int objectType, QObject * component);
extern void C_ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  const QMetaObject * QAbstractTextDocumentLayout::metaObject();
extern void C_ZNK27QAbstractTextDocumentLayout10metaObjectEv(void* qthis); // 4
  // proto:  void QAbstractTextDocumentLayout::~QAbstractTextDocumentLayout();
extern void C_ZN27QAbstractTextDocumentLayoutD2Ev(void* qthis); // 4
  // proto:  QTextObjectInterface * QAbstractTextDocumentLayout::handlerForObject(int objectType);
extern void* C_ZNK27QAbstractTextDocumentLayout16handlerForObjectEi(void* qthis, int32_t arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextObjectInterface)=8
type QTextObjectInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractTextDocumentLayout)=1
type QAbstractTextDocumentLayout struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _updateBlock QAbstractTextDocumentLayout_updateBlock_signal;
//  _pageCountChanged QAbstractTextDocumentLayout_pageCountChanged_signal;
//  _update QAbstractTextDocumentLayout_update_signal;
//  _documentSizeChanged QAbstractTextDocumentLayout_documentSizeChanged_signal;
}

// unregisterHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) Unregisterhandler(args ...interface{}) () {
  // unregisterHandler(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject
    // invoke: void unregisterHandler(int, class QObject *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN27QAbstractTextDocumentLayout17unregisterHandlerEiP7QObject(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "unregisterHandler", args)
  }

  return
}

// setPaintDevice(class QPaintDevice *)
func (this *QAbstractTextDocumentLayout) Setpaintdevice(args ...interface{}) () {
  // setPaintDevice(class QPaintDevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice
    // invoke: void setPaintDevice(class QPaintDevice *)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QAbstractTextDocumentLayout14setPaintDeviceEP12QPaintDevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "setPaintDevice", args)
  }

  return
}

// anchorAt(const class QPointF &)
func (this *QAbstractTextDocumentLayout) Anchorat(args ...interface{}) (ret interface{}) {
  // anchorAt(const class QPointF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPointF{}) // "const QPointF &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF
    // invoke: QString anchorAt(const class QPointF &)
    var arg0 = args[0].(*qtcore.QPointF).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK27QAbstractTextDocumentLayout8anchorAtERK7QPointF(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "anchorAt", args)
  }

  return
}

// QAbstractTextDocumentLayout(class QTextDocument *)
func NewQAbstractTextDocumentLayout(args ...interface{}) *QAbstractTextDocumentLayout {
  // QAbstractTextDocumentLayout(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayoutC1EP13QTextDocument
    // invoke: void QAbstractTextDocumentLayout(class QTextDocument *)
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QAbstractTextDocumentLayoutC2EP13QTextDocument(arg0)
    return &QAbstractTextDocumentLayout{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "QAbstractTextDocumentLayout", args)
  }

  return nil // QAbstractTextDocumentLayout{Qclsinst:qthis}
}

// paintDevice()
func (this *QAbstractTextDocumentLayout) Paintdevice(args ...interface{}) (ret interface{}) {
  // paintDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout11paintDeviceEv
    // invoke: QPaintDevice * paintDevice()
    var ret0 = C.C_ZNK27QAbstractTextDocumentLayout11paintDeviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "paintDevice", args)
  }

  return
}

// document()
func (this *QAbstractTextDocumentLayout) Document(args ...interface{}) (ret interface{}) {
  // document()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout8documentEv
    // invoke: QTextDocument * document()
    var ret0 = C.C_ZNK27QAbstractTextDocumentLayout8documentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "document", args)
  }

  return
}

// registerHandler(int, class QObject *)
func (this *QAbstractTextDocumentLayout) Registerhandler(args ...interface{}) () {
  // registerHandler(int, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject
    // invoke: void registerHandler(int, class QObject *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN27QAbstractTextDocumentLayout15registerHandlerEiP7QObject(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "registerHandler", args)
  }

  return
}

// metaObject()
func (this *QAbstractTextDocumentLayout) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK27QAbstractTextDocumentLayout10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "metaObject", args)
  }

  return
}

// ~QAbstractTextDocumentLayout()
func (this *QAbstractTextDocumentLayout) Freeqabstracttextdocumentlayout(args ...interface{}) () {
  // ~QAbstractTextDocumentLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAbstractTextDocumentLayoutD0Ev
    // invoke: void ~QAbstractTextDocumentLayout()
    C.C_ZN27QAbstractTextDocumentLayoutD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "~QAbstractTextDocumentLayout", args)
  }

  return
}

// handlerForObject(int)
func (this *QAbstractTextDocumentLayout) Handlerforobject(args ...interface{}) (ret interface{}) {
  // handlerForObject(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout16handlerForObjectEi
    // invoke: QTextObjectInterface * handlerForObject(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK27QAbstractTextDocumentLayout16handlerForObjectEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextObjectInterface{}) // "QTextObjectInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "handlerForObject", args)
  }

  return
}

// <= body block end

