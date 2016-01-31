package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void C_ZN5QDrag11setMimeDataEP9QMimeData(void* qthis, void* arg0); // 4
  // proto:  void QDrag::QDrag(QObject * dragSource);
extern void* C_ZN5QDragC2EP7QObject(void* arg0); // 3
  // proto:  QPixmap QDrag::pixmap();
extern void* C_ZNK5QDrag6pixmapEv(void* qthis); // 4
  // proto:  void QDrag::~QDrag();
extern void C_ZN5QDragD2Ev(void* qthis); // 4
  // proto:  Qt::DropAction QDrag::defaultAction();
extern void C_ZNK5QDrag13defaultActionEv(void* qthis); // 4
  // proto:  QObject * QDrag::source();
extern void* C_ZNK5QDrag6sourceEv(void* qthis); // 4
  // proto:  Qt::DropActions QDrag::supportedActions();
extern void C_ZNK5QDrag16supportedActionsEv(void* qthis); // 4
  // proto:  void QDrag::setHotSpot(const QPoint & hotspot);
extern void C_ZN5QDrag10setHotSpotERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QDrag::setPixmap(const QPixmap & );
extern void C_ZN5QDrag9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  QPoint QDrag::hotSpot();
extern void* C_ZNK5QDrag7hotSpotEv(void* qthis); // 4
  // proto:  QMimeData * QDrag::mimeData();
extern void* C_ZNK5QDrag8mimeDataEv(void* qthis); // 4
  // proto:  const QMetaObject * QDrag::metaObject();
extern void C_ZNK5QDrag10metaObjectEv(void* qthis); // 4
  // proto:  QObject * QDrag::target();
extern void* C_ZNK5QDrag6targetEv(void* qthis); // 4
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
func (this *QDrag) Setmimedata(args ...interface{}) () {
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
    C.C_ZN5QDrag11setMimeDataEP9QMimeData(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setMimeData", args)
  }

  return
}

// QDrag(class QObject *)
func NewQDrag(args ...interface{}) *QDrag {
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
    qthis = C.C_ZN5QDragC2EP7QObject(arg0)
    return &QDrag{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDrag", "QDrag", args)
  }

  return nil // QDrag{qclsinst:qthis}
}

// pixmap()
func (this *QDrag) Pixmap(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QDrag6pixmapEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDrag", "pixmap", args)
  }

  return
}

// ~QDrag()
func (this *QDrag) Freeqdrag(args ...interface{}) () {
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
    C.C_ZN5QDragD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "~QDrag", args)
  }

  return
}

// defaultAction()
func (this *QDrag) Defaultaction(args ...interface{}) () {
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
    C.C_ZNK5QDrag13defaultActionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "defaultAction", args)
  }

  return
}

// source()
func (this *QDrag) Source(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QDrag6sourceEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDrag", "source", args)
  }

  return
}

// supportedActions()
func (this *QDrag) Supportedactions(args ...interface{}) () {
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
    C.C_ZNK5QDrag16supportedActionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "supportedActions", args)
  }

  return
}

// setHotSpot(const class QPoint &)
func (this *QDrag) Sethotspot(args ...interface{}) () {
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
    C.C_ZN5QDrag10setHotSpotERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setHotSpot", args)
  }

  return
}

// setPixmap(const class QPixmap &)
func (this *QDrag) Setpixmap(args ...interface{}) () {
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
    C.C_ZN5QDrag9setPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDrag", "setPixmap", args)
  }

  return
}

// hotSpot()
func (this *QDrag) Hotspot(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QDrag7hotSpotEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDrag", "hotSpot", args)
  }

  return
}

// mimeData()
func (this *QDrag) Mimedata(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QDrag8mimeDataEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMimeData{}) // "QMimeData *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDrag", "mimeData", args)
  }

  return
}

// metaObject()
func (this *QDrag) Metaobject(args ...interface{}) () {
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
    C.C_ZNK5QDrag10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDrag", "metaObject", args)
  }

  return
}

// target()
func (this *QDrag) Target(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QDrag6targetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QObject{}) // "QObject *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDrag", "target", args)
  }

  return
}

// <= body block end

