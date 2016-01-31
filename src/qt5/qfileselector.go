package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtCore/qfileselector.h
// dst-file: /src/core/qfileselector.go
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
  // proto:  void QFileSelector::QFileSelector(QObject * parent);
extern void* C_ZN13QFileSelectorC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QFileSelector::metaObject();
extern void C_ZNK13QFileSelector10metaObjectEv(void* qthis); // 4
  // proto:  QStringList QFileSelector::allSelectors();
extern void C_ZNK13QFileSelector12allSelectorsEv(void* qthis); // 4
  // proto:  QStringList QFileSelector::extraSelectors();
extern void C_ZNK13QFileSelector14extraSelectorsEv(void* qthis); // 4
  // proto:  void QFileSelector::setExtraSelectors(const QStringList & list);
extern void C_ZN13QFileSelector17setExtraSelectorsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QFileSelector::~QFileSelector();
extern void C_ZN13QFileSelectorD2Ev(void* qthis); // 4
  // proto:  QString QFileSelector::select(const QString & filePath);
extern void* C_ZNK13QFileSelector6selectERK7QString(void* qthis, void* arg0); // 4
  // proto:  QUrl QFileSelector::select(const QUrl & filePath);
extern void* C_ZNK13QFileSelector6selectERK4QUrl(void* qthis, void* arg0); // 4
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

// class sizeof(QFileSelector)=1
type QFileSelector struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QFileSelector(class QObject *)
func NewQFileSelector(args ...interface{}) *QFileSelector {
  // QFileSelector(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFileSelectorC1EP7QObject
    // invoke: void QFileSelector(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QFileSelectorC2EP7QObject(arg0)
    return &QFileSelector{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFileSelector", "QFileSelector", args)
  }

  return nil // QFileSelector{qclsinst:qthis}
}

// metaObject()
func (this *QFileSelector) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFileSelector10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QFileSelector10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSelector", "metaObject", args)
  }

  return
}

// allSelectors()
func (this *QFileSelector) Allselectors(args ...interface{}) () {
  // allSelectors()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFileSelector12allSelectorsEv
    // invoke: QStringList allSelectors()
    C.C_ZNK13QFileSelector12allSelectorsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSelector", "allSelectors", args)
  }

  return
}

// extraSelectors()
func (this *QFileSelector) Extraselectors(args ...interface{}) () {
  // extraSelectors()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFileSelector14extraSelectorsEv
    // invoke: QStringList extraSelectors()
    C.C_ZNK13QFileSelector14extraSelectorsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSelector", "extraSelectors", args)
  }

  return
}

// setExtraSelectors(const class QStringList &)
func (this *QFileSelector) Setextraselectors(args ...interface{}) () {
  // setExtraSelectors(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFileSelector17setExtraSelectorsERK11QStringList
    // invoke: void setExtraSelectors(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QFileSelector17setExtraSelectorsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileSelector", "setExtraSelectors", args)
  }

  return
}

// ~QFileSelector()
func (this *QFileSelector) Freeqfileselector(args ...interface{}) () {
  // ~QFileSelector()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFileSelectorD0Ev
    // invoke: void ~QFileSelector()
    C.C_ZN13QFileSelectorD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileSelector", "~QFileSelector", args)
  }

  return
}

// select(const class QString &)
func (this *QFileSelector) Select_(args ...interface{}) (ret interface{}) {
  // select(const class QString &)
  // select(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFileSelector6selectERK7QString
    // invoke: QString select(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFileSelector6selectERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK13QFileSelector6selectERK4QUrl
    // invoke: QUrl select(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFileSelector6selectERK4QUrl(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUrl{}) // "QUrl"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QFileSelector", "select", args)
  }

  return
}

// <= body block end

