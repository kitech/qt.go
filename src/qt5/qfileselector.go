package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QStringList QFileSelector::allSelectors();
extern void _ZNK13QFileSelector12allSelectorsEv(void* qthis);
  // proto:  const QMetaObject * QFileSelector::metaObject();
extern void _ZNK13QFileSelector10metaObjectEv(void* qthis);
  // proto:  QUrl QFileSelector::select(const QUrl & filePath);
extern void _ZNK13QFileSelector6selectERK4QUrl(void* qthis, void* arg0);
  // proto:  void QFileSelector::QFileSelector(QObject * parent);
extern void* dector_ZN13QFileSelectorC1EP7QObject(void* arg0);
extern void _ZN13QFileSelectorC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QFileSelector::setExtraSelectors(const QStringList & list);
extern void _ZN13QFileSelector17setExtraSelectorsERK11QStringList(void* qthis, void* arg0);
  // proto:  QString QFileSelector::select(const QString & filePath);
extern void _ZNK13QFileSelector6selectERK7QString(void* qthis, void* arg0);
  // proto:  void QFileSelector::~QFileSelector();
extern void _ZN13QFileSelectorD0Ev(void* qthis);
  // proto:  QStringList QFileSelector::extraSelectors();
extern void _ZNK13QFileSelector14extraSelectorsEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QStringList QFileSelector::allSelectors();
func (this *QFileSelector) allSelectors(args ...interface{}) () {
  // allSelectors()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFileSelector12allSelectorsEv
  default:
    qtrt.ErrorResolve("QFileSelector", "allSelectors", args)
  }

}

  // proto:  const QMetaObject * QFileSelector::metaObject();
func (this *QFileSelector) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFileSelector10metaObjectEv
  default:
    qtrt.ErrorResolve("QFileSelector", "metaObject", args)
  }

}

  // proto:  QUrl QFileSelector::select(const QUrl & filePath);
func (this *QFileSelector) select_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileSelector", "select", args)
  }

}

  // proto:  void QFileSelector::QFileSelector(QObject * parent);
func NewQFileSelector(args ...interface{}) QFileSelector {
  return QFileSelector{}
}

  // proto:  void QFileSelector::setExtraSelectors(const QStringList & list);
func (this *QFileSelector) setExtraSelectors(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QFileSelector", "setExtraSelectors", args)
  }

}

  // proto:  void QFileSelector::~QFileSelector();
func (this *QFileSelector) FreeQFileSelector(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileSelector", "~QFileSelector", args)
  }

}

  // proto:  QStringList QFileSelector::extraSelectors();
func (this *QFileSelector) extraSelectors(args ...interface{}) () {
  // extraSelectors()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFileSelector14extraSelectorsEv
  default:
    qtrt.ErrorResolve("QFileSelector", "extraSelectors", args)
  }

}

// <= body block end

