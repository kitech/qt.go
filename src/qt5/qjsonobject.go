package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtCore/qjsonobject.h
// dst-file: /src/core/qjsonobject.go
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
  // proto:  bool QJsonObject::isEmpty();
extern void _ZNK11QJsonObject7isEmptyEv(void* qthis);
  // proto:  int QJsonObject::length();
extern void demth_ZNK11QJsonObject6lengthEv(void* qthis);
  // proto:  void QJsonObject::remove(const QString & key);
extern void _ZN11QJsonObject6removeERK7QString(void* qthis, void* arg0);
  // proto:  void QJsonObject::~QJsonObject();
extern void _ZN11QJsonObjectD0Ev(void* qthis);
  // proto:  QJsonValue QJsonObject::value(const QString & key);
extern void _ZNK11QJsonObject5valueERK7QString(void* qthis, void* arg0);
  // proto:  int QJsonObject::size();
extern void _ZNK11QJsonObject4sizeEv(void* qthis);
  // proto:  int QJsonObject::count();
extern void demth_ZNK11QJsonObject5countEv(void* qthis);
  // proto:  bool QJsonObject::empty();
extern void demth_ZNK11QJsonObject5emptyEv(void* qthis);
  // proto:  QJsonValue QJsonObject::take(const QString & key);
extern void _ZN11QJsonObject4takeERK7QString(void* qthis, void* arg0);
  // proto:  QVariantHash QJsonObject::toVariantHash();
extern void _ZNK11QJsonObject13toVariantHashEv(void* qthis);
  // proto:  QStringList QJsonObject::keys();
extern void _ZNK11QJsonObject4keysEv(void* qthis);
  // proto:  bool QJsonObject::contains(const QString & key);
extern void _ZNK11QJsonObject8containsERK7QString(void* qthis, void* arg0);
  // proto:  void QJsonObject::QJsonObject();
extern void* dector_ZN11QJsonObjectC1Ev();
extern void _ZN11QJsonObjectC1Ev(void* qthis);
  // proto:  QVariantMap QJsonObject::toVariantMap();
extern void _ZNK11QJsonObject12toVariantMapEv(void* qthis);
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

// class sizeof(QJsonObject)=16
type QJsonObject struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QJsonObject::isEmpty();
func (this *QJsonObject) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject7isEmptyEv
  default:
    qtrt.ErrorResolve("QJsonObject", "isEmpty", args)
  }

}

  // proto:  int QJsonObject::length();
func (this *QJsonObject) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject6lengthEv
  default:
    qtrt.ErrorResolve("QJsonObject", "length", args)
  }

}

  // proto:  void QJsonObject::remove(const QString & key);
func (this *QJsonObject) remove(args ...interface{}) () {
  // remove(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QJsonObject6removeERK7QString
  default:
    qtrt.ErrorResolve("QJsonObject", "remove", args)
  }

}

  // proto:  void QJsonObject::~QJsonObject();
func (this *QJsonObject) FreeQJsonObject(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QJsonObject", "~QJsonObject", args)
  }

}

  // proto:  QJsonValue QJsonObject::value(const QString & key);
func (this *QJsonObject) value(args ...interface{}) () {
  // value(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject5valueERK7QString
  default:
    qtrt.ErrorResolve("QJsonObject", "value", args)
  }

}

  // proto:  int QJsonObject::size();
func (this *QJsonObject) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject4sizeEv
  default:
    qtrt.ErrorResolve("QJsonObject", "size", args)
  }

}

  // proto:  int QJsonObject::count();
func (this *QJsonObject) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject5countEv
  default:
    qtrt.ErrorResolve("QJsonObject", "count", args)
  }

}

  // proto:  bool QJsonObject::empty();
func (this *QJsonObject) empty(args ...interface{}) () {
  // empty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject5emptyEv
  default:
    qtrt.ErrorResolve("QJsonObject", "empty", args)
  }

}

  // proto:  QJsonValue QJsonObject::take(const QString & key);
func (this *QJsonObject) take(args ...interface{}) () {
  // take(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QJsonObject4takeERK7QString
  default:
    qtrt.ErrorResolve("QJsonObject", "take", args)
  }

}

  // proto:  QVariantHash QJsonObject::toVariantHash();
func (this *QJsonObject) toVariantHash(args ...interface{}) () {
  // toVariantHash()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject13toVariantHashEv
  default:
    qtrt.ErrorResolve("QJsonObject", "toVariantHash", args)
  }

}

  // proto:  QStringList QJsonObject::keys();
func (this *QJsonObject) keys(args ...interface{}) () {
  // keys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject4keysEv
  default:
    qtrt.ErrorResolve("QJsonObject", "keys", args)
  }

}

  // proto:  bool QJsonObject::contains(const QString & key);
func (this *QJsonObject) contains(args ...interface{}) () {
  // contains(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject8containsERK7QString
  default:
    qtrt.ErrorResolve("QJsonObject", "contains", args)
  }

}

  // proto:  void QJsonObject::QJsonObject();
func NewQJsonObject(args ...interface{}) QJsonObject {
  return QJsonObject{}
}

  // proto:  QVariantMap QJsonObject::toVariantMap();
func (this *QJsonObject) toVariantMap(args ...interface{}) () {
  // toVariantMap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject12toVariantMapEv
  default:
    qtrt.ErrorResolve("QJsonObject", "toVariantMap", args)
  }

}

// <= body block end

