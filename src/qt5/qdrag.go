package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtGui/qdrag.h
// dst-file: /src/gui/qdrag.go
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
  // proto:  void QDrag::setMimeData(QMimeData * data);
extern void _ZN5QDrag11setMimeDataEP9QMimeData(void* qthis, void* arg0); // 4
  // proto:  void QDrag::QDrag(QObject * dragSource);
extern void _ZN5QDragC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  QPixmap QDrag::pixmap();
extern void _ZNK5QDrag6pixmapEv(void* qthis); // 4
  // proto:  void QDrag::~QDrag();
extern void _ZN5QDragD2Ev(void* qthis); // 4
  // proto:  Qt::DropAction QDrag::defaultAction();
extern void _ZNK5QDrag13defaultActionEv(void* qthis); // 4
  // proto:  QObject * QDrag::source();
extern void _ZNK5QDrag6sourceEv(void* qthis); // 4
  // proto:  Qt::DropActions QDrag::supportedActions();
extern void _ZNK5QDrag16supportedActionsEv(void* qthis); // 4
  // proto:  void QDrag::setHotSpot(const QPoint & hotspot);
extern void _ZN5QDrag10setHotSpotERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QDrag::setPixmap(const QPixmap & );
extern void _ZN5QDrag9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  QPoint QDrag::hotSpot();
extern void _ZNK5QDrag7hotSpotEv(void* qthis); // 4
  // proto:  QMimeData * QDrag::mimeData();
extern void _ZNK5QDrag8mimeDataEv(void* qthis); // 4
  // proto:  const QMetaObject * QDrag::metaObject();
extern void _ZNK5QDrag10metaObjectEv(void* qthis); // 4
  // proto:  QObject * QDrag::target();
extern void _ZNK5QDrag6targetEv(void* qthis); // 4
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

// class sizeof(QDrag)=1
type QDrag struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
//  _targetChanged QDrag_targetChanged_signal;
//  _actionChanged QDrag_actionChanged_signal;
}

// setMimeData(class QMimeData *)
func (this *QDrag) setMimeData(args ...interface{}) () {
  // setMimeData(class QMimeData *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMimeData{}) // "QMimeData *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDrag11setMimeDataEP9QMimeData
    // invoke: void setMimeData(class QMimeData *)
    var arg0 = args[0].(QMimeData).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QDrag11setMimeDataEP9QMimeData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setMimeData", args)
  }

}

// QDrag(class QObject *)
func NewQDrag(args ...interface{}) QDrag {
  // QDrag(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDragC1EP7QObject
    // invoke: void QDrag(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QDragC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "QDrag", args)
  }

  return QDrag{}
}

// pixmap()
func (this *QDrag) pixmap(args ...interface{}) () {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag6pixmapEv
    // invoke: QPixmap pixmap()
    C._ZNK5QDrag6pixmapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "pixmap", args)
  }

}

// ~QDrag()
func (this *QDrag) FreeQDrag(args ...interface{}) () {
  // ~QDrag()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDragD0Ev
    // invoke: void ~QDrag()
    C._ZN5QDragD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "~QDrag", args)
  }

}

// defaultAction()
func (this *QDrag) defaultAction(args ...interface{}) () {
  // defaultAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag13defaultActionEv
    // invoke: Qt::DropAction defaultAction()
    C._ZNK5QDrag13defaultActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "defaultAction", args)
  }

}

// source()
func (this *QDrag) source(args ...interface{}) () {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag6sourceEv
    // invoke: QObject * source()
    C._ZNK5QDrag6sourceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "source", args)
  }

}

// supportedActions()
func (this *QDrag) supportedActions(args ...interface{}) () {
  // supportedActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag16supportedActionsEv
    // invoke: Qt::DropActions supportedActions()
    C._ZNK5QDrag16supportedActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "supportedActions", args)
  }

}

// setHotSpot(const class QPoint &)
func (this *QDrag) setHotSpot(args ...interface{}) () {
  // setHotSpot(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDrag10setHotSpotERK6QPoint
    // invoke: void setHotSpot(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QDrag10setHotSpotERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setHotSpot", args)
  }

}

// setPixmap(const class QPixmap &)
func (this *QDrag) setPixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QDrag9setPixmapERK7QPixmap
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QDrag9setPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setPixmap", args)
  }

}

// hotSpot()
func (this *QDrag) hotSpot(args ...interface{}) () {
  // hotSpot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag7hotSpotEv
    // invoke: QPoint hotSpot()
    C._ZNK5QDrag7hotSpotEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "hotSpot", args)
  }

}

// mimeData()
func (this *QDrag) mimeData(args ...interface{}) () {
  // mimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag8mimeDataEv
    // invoke: QMimeData * mimeData()
    C._ZNK5QDrag8mimeDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "mimeData", args)
  }

}

// metaObject()
func (this *QDrag) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK5QDrag10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "metaObject", args)
  }

}

// target()
func (this *QDrag) target(args ...interface{}) () {
  // target()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QDrag6targetEv
    // invoke: QObject * target()
    C._ZNK5QDrag6targetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "target", args)
  }

}

// <= body block end

