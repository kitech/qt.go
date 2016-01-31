package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QJsonObject::const_iterator QJsonObject::constBegin();
extern void C_ZNK11QJsonObject10constBeginEv(void* qthis); // 2
  // proto:  QJsonValue QJsonObject::value(const QString & key);
extern void C_ZNK11QJsonObject5valueERK7QString(void* qthis, void* arg0); // 4
  // proto:  QJsonObject::const_iterator QJsonObject::constFind(const QString & key);
extern void C_ZNK11QJsonObject9constFindERK7QString(void* qthis, void* arg0); // 4
  // proto:  QJsonObject::iterator QJsonObject::find(const QString & key);
extern void C_ZN11QJsonObject4findERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QJsonObject::size();
extern void C_ZNK11QJsonObject4sizeEv(void* qthis); // 4
  // proto:  QJsonObject::iterator QJsonObject::end();
extern void C_ZN11QJsonObject3endEv(void* qthis); // 2
  // proto:  bool QJsonObject::contains(const QString & key);
extern void C_ZNK11QJsonObject8containsERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QJsonObject::isEmpty();
extern void C_ZNK11QJsonObject7isEmptyEv(void* qthis); // 4
  // proto:  QJsonValue QJsonObject::take(const QString & key);
extern void C_ZN11QJsonObject4takeERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QJsonObject::empty();
extern void C_ZNK11QJsonObject5emptyEv(void* qthis); // 2
  // proto:  QJsonObject::iterator QJsonObject::begin();
extern void C_ZN11QJsonObject5beginEv(void* qthis); // 2
  // proto:  QVariantHash QJsonObject::toVariantHash();
extern void C_ZNK11QJsonObject13toVariantHashEv(void* qthis); // 4
  // proto:  QStringList QJsonObject::keys();
extern void C_ZNK11QJsonObject4keysEv(void* qthis); // 4
  // proto:  void QJsonObject::~QJsonObject();
extern void C_ZN11QJsonObjectD2Ev(void* qthis); // 4
  // proto:  QVariantMap QJsonObject::toVariantMap();
extern void C_ZNK11QJsonObject12toVariantMapEv(void* qthis); // 4
  // proto:  int QJsonObject::count();
extern void C_ZNK11QJsonObject5countEv(void* qthis); // 2
  // proto:  void QJsonObject::QJsonObject();
extern void* C_ZN11QJsonObjectC2Ev(); // 3
  // proto:  void QJsonObject::remove(const QString & key);
extern void C_ZN11QJsonObject6removeERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QJsonObject::length();
extern void C_ZNK11QJsonObject6lengthEv(void* qthis); // 2
  // proto:  QJsonObject::const_iterator QJsonObject::constEnd();
extern void C_ZNK11QJsonObject8constEndEv(void* qthis); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// constBegin()
func (this *QJsonObject) constBegin(args ...interface{}) () {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject10constBeginEv
    // invoke: QJsonObject::const_iterator constBegin()
    C.C_ZNK11QJsonObject10constBeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "constBegin", args)
  }

}

// value(const class QString &)
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
    // invoke: QJsonValue value(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QJsonObject5valueERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "value", args)
  }

}

// constFind(const class QString &)
func (this *QJsonObject) constFind(args ...interface{}) () {
  // constFind(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject9constFindERK7QString
    // invoke: QJsonObject::const_iterator constFind(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QJsonObject9constFindERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "constFind", args)
  }

}

// find(const class QString &)
func (this *QJsonObject) find(args ...interface{}) () {
  // find(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QJsonObject4findERK7QString
    // invoke: QJsonObject::iterator find(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QJsonObject4findERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "find", args)
  }

}

// size()
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
    // invoke: int size()
    var ret = C.C_ZNK11QJsonObject4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonObject", "size", args)
  }

}

// end()
func (this *QJsonObject) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QJsonObject3endEv
    // invoke: QJsonObject::iterator end()
    C.C_ZN11QJsonObject3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "end", args)
  }

}

// contains(const class QString &)
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
    // invoke: bool contains(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QJsonObject8containsERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonObject", "contains", args)
  }

}

// isEmpty()
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
    // invoke: bool isEmpty()
    var ret = C.C_ZNK11QJsonObject7isEmptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonObject", "isEmpty", args)
  }

}

// take(const class QString &)
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
    // invoke: QJsonValue take(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QJsonObject4takeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "take", args)
  }

}

// empty()
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
    // invoke: bool empty()
    var ret = C.C_ZNK11QJsonObject5emptyEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonObject", "empty", args)
  }

}

// begin()
func (this *QJsonObject) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QJsonObject5beginEv
    // invoke: QJsonObject::iterator begin()
    C.C_ZN11QJsonObject5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "begin", args)
  }

}

// toVariantHash()
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
    // invoke: QVariantHash toVariantHash()
    C.C_ZNK11QJsonObject13toVariantHashEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "toVariantHash", args)
  }

}

// keys()
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
    // invoke: QStringList keys()
    C.C_ZNK11QJsonObject4keysEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "keys", args)
  }

}

// ~QJsonObject()
func (this *QJsonObject) FreeQJsonObject(args ...interface{}) () {
  // ~QJsonObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QJsonObjectD0Ev
    // invoke: void ~QJsonObject()
    C.C_ZN11QJsonObjectD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "~QJsonObject", args)
  }

}

// toVariantMap()
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
    // invoke: QVariantMap toVariantMap()
    C.C_ZNK11QJsonObject12toVariantMapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "toVariantMap", args)
  }

}

// count()
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
    // invoke: int count()
    var ret = C.C_ZNK11QJsonObject5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonObject", "count", args)
  }

}

// QJsonObject()
func NewQJsonObject(args ...interface{}) *QJsonObject {
  // QJsonObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QJsonObjectC1Ev
    // invoke: void QJsonObject()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QJsonObjectC2Ev()
    return &QJsonObject{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QJsonObject", "QJsonObject", args)
  }

  return nil // QJsonObject{qclsinst:qthis}
}

// remove(const class QString &)
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
    // invoke: void remove(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QJsonObject6removeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "remove", args)
  }

}

// length()
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
    // invoke: int length()
    var ret = C.C_ZNK11QJsonObject6lengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QJsonObject", "length", args)
  }

}

// constEnd()
func (this *QJsonObject) constEnd(args ...interface{}) () {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QJsonObject8constEndEv
    // invoke: QJsonObject::const_iterator constEnd()
    C.C_ZNK11QJsonObject8constEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "constEnd", args)
  }

}

// <= body block end

