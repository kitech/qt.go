package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qabstracttextdocumentlayout.h
// dst-file: /src/gui/qabstracttextdocumentlayout.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QTextObjectInterface)=8
type QTextObjectInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAbstractTextDocumentLayout)=1
type QAbstractTextDocumentLayout struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _updateBlock QAbstractTextDocumentLayout_updateBlock_signal;
//  _pageCountChanged QAbstractTextDocumentLayout_pageCountChanged_signal;
//  _update QAbstractTextDocumentLayout_update_signal;
//  _documentSizeChanged QAbstractTextDocumentLayout_documentSizeChanged_signal;
}


func (this *QTextObjectInterface) FreeQTextObjectInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextObjectInterface", "~QTextObjectInterface", args)
 }

}


func (this *QTextObjectInterface) intrinsicSize(args ...interface{}) () {
  // intrinsicSize(class QTextDocument *, int, const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QTextObjectInterface13intrinsicSizeEP13QTextDocumentiRK11QTextFormat
  default:
    qtrt.ErrorResolve("QTextObjectInterface", "intrinsicSize", args)
 }

}


func (this *QTextObjectInterface) drawObject(args ...interface{}) () {
  // drawObject(class QPainter *, const class QRectF &, class QTextDocument *, int, const class QTextFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][2] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = reflect.TypeOf(QTextFormat{}) // "const QTextFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN20QTextObjectInterface10drawObjectEP8QPainterRK6QRectFP13QTextDocumentiRK11QTextFormat
  default:
    qtrt.ErrorResolve("QTextObjectInterface", "drawObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "registerHandler", args)
 }

}


func (this *QAbstractTextDocumentLayout) pageCount(args ...interface{}) () {
  // pageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout9pageCountEv
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "pageCount", args)
 }

}


func (this *QAbstractTextDocumentLayout) FreeQAbstractTextDocumentLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "~QAbstractTextDocumentLayout", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "setPaintDevice", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "document", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "unregisterHandler", args)
 }

}


func NewQAbstractTextDocumentLayout(args ...interface{})() {
}


func (this *QAbstractTextDocumentLayout) documentSize(args ...interface{}) () {
  // documentSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout12documentSizeEv
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "documentSize", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "paintDevice", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "anchorAt", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "handlerForObject", args)
 }

}


func (this *QAbstractTextDocumentLayout) frameBoundingRect(args ...interface{}) () {
  // frameBoundingRect(class QTextFrame *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextFrame{}) // "QTextFrame *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout17frameBoundingRectEP10QTextFrame
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "frameBoundingRect", args)
 }

}


func (this *QAbstractTextDocumentLayout) blockBoundingRect(args ...interface{}) () {
  // blockBoundingRect(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK27QAbstractTextDocumentLayout17blockBoundingRectERK10QTextBlock
  default:
    qtrt.ErrorResolve("QAbstractTextDocumentLayout", "blockBoundingRect", args)
 }

}

// <= body block end

