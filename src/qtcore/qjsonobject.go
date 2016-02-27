package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
extern int32_t C_ZNK11QJsonObject4sizeEv(void* qthis); // 4
  // proto:  QJsonObject::iterator QJsonObject::end();
extern void C_ZN11QJsonObject3endEv(void* qthis); // 2
  // proto:  bool QJsonObject::contains(const QString & key);
extern bool C_ZNK11QJsonObject8containsERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QJsonObject::isEmpty();
extern bool C_ZNK11QJsonObject7isEmptyEv(void* qthis); // 4
  // proto:  QJsonValue QJsonObject::take(const QString & key);
extern void C_ZN11QJsonObject4takeERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QJsonObject::empty();
extern bool C_ZNK11QJsonObject5emptyEv(void* qthis); // 2
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
extern int32_t C_ZNK11QJsonObject5countEv(void* qthis); // 2
  // proto:  void QJsonObject::QJsonObject();
extern void* C_ZN11QJsonObjectC2Ev(); // 3
  // proto:  void QJsonObject::remove(const QString & key);
extern void C_ZN11QJsonObject6removeERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QJsonObject::length();
extern int32_t C_ZNK11QJsonObject6lengthEv(void* qthis); // 2
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QJsonObject)=16
type QJsonObject struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// constBegin()
func (this *QJsonObject) ConstBegin(args ...interface{}) () {
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
    C.C_ZNK11QJsonObject10constBeginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "constBegin", args)
  }

  return
}

// value(const class QString &)
func (this *QJsonObject) Value(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QJsonObject5valueERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "value", args)
  }

  return
}

// constFind(const class QString &)
func (this *QJsonObject) ConstFind(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QJsonObject9constFindERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "constFind", args)
  }

  return
}

// find(const class QString &)
func (this *QJsonObject) Find(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QJsonObject4findERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "find", args)
  }

  return
}

// size()
func (this *QJsonObject) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QJsonObject4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonObject", "size", args)
  }

  return
}

// end()
func (this *QJsonObject) End(args ...interface{}) () {
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
    C.C_ZN11QJsonObject3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "end", args)
  }

  return
}

// contains(const class QString &)
func (this *QJsonObject) Contains(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QJsonObject8containsERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonObject", "contains", args)
  }

  return
}

// isEmpty()
func (this *QJsonObject) IsEmpty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QJsonObject7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonObject", "isEmpty", args)
  }

  return
}

// take(const class QString &)
func (this *QJsonObject) Take(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QJsonObject4takeERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "take", args)
  }

  return
}

// empty()
func (this *QJsonObject) Empty(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QJsonObject5emptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonObject", "empty", args)
  }

  return
}

// begin()
func (this *QJsonObject) Begin(args ...interface{}) () {
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
    C.C_ZN11QJsonObject5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "begin", args)
  }

  return
}

// toVariantHash()
func (this *QJsonObject) ToVariantHash(args ...interface{}) () {
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
    C.C_ZNK11QJsonObject13toVariantHashEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "toVariantHash", args)
  }

  return
}

// keys()
func (this *QJsonObject) Keys(args ...interface{}) () {
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
    C.C_ZNK11QJsonObject4keysEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "keys", args)
  }

  return
}

// ~QJsonObject()
func (this *QJsonObject) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QJsonObjectD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QJsonObject", "~QJsonObject", args)
  }

  return
}

// toVariantMap()
func (this *QJsonObject) ToVariantMap(args ...interface{}) () {
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
    C.C_ZNK11QJsonObject12toVariantMapEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "toVariantMap", args)
  }

  return
}

// count()
func (this *QJsonObject) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QJsonObject5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonObject", "count", args)
  }

  return
}

// QJsonObject()
func GcfreeQJsonObject(this *QJsonObject) {
  qtrt.UniverseFree(this)
}
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
    this := &QJsonObject{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQJsonObject)
    return this
  default:
    qtrt.ErrorResolve("QJsonObject", "QJsonObject", args)
  }

  return nil // QJsonObject{Qclsinst:qthis}
}

// remove(const class QString &)
func (this *QJsonObject) Remove(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QJsonObject6removeERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QJsonObject", "remove", args)
  }

  return
}

// length()
func (this *QJsonObject) Length(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QJsonObject6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QJsonObject", "length", args)
  }

  return
}

// constEnd()
func (this *QJsonObject) ConstEnd(args ...interface{}) () {
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
    C.C_ZNK11QJsonObject8constEndEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QJsonObject", "constEnd", args)
  }

  return
}

// <= body block end

