package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
// src-file: /QtCore/quuid.h
// dst-file: /src/core/quuid.go
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
  // proto: static QUuid QUuid::createUuidV5(const QUuid & ns, const QByteArray & baseData);
extern void _ZN5QUuid12createUuidV5ERKS_RK10QByteArray(void* arg0, void* arg1); // 4
  // proto: static QUuid QUuid::createUuidV5(const QUuid & ns, const QString & baseData);
extern void _ZN5QUuid12createUuidV5ERKS_RK7QString(void* arg0, void* arg1); // 2
  // proto: static QUuid QUuid::createUuidV3(const QUuid & ns, const QString & baseData);
extern void _ZN5QUuid12createUuidV3ERKS_RK7QString(void* arg0, void* arg1); // 2
  // proto: static QUuid QUuid::createUuidV3(const QUuid & ns, const QByteArray & baseData);
extern void _ZN5QUuid12createUuidV3ERKS_RK10QByteArray(void* arg0, void* arg1); // 4
  // proto: static QUuid QUuid::fromRfc4122(const QByteArray & );
extern void _ZN5QUuid11fromRfc4122ERK10QByteArray(void* arg0); // 4
  // proto:  QUuid::Variant QUuid::variant();
extern void _ZNK5QUuid7variantEv(void* qthis); // 4
  // proto:  bool QUuid::isNull();
extern void _ZNK5QUuid6isNullEv(void* qthis); // 4
  // proto:  QUuid::Version QUuid::version();
extern void _ZNK5QUuid7versionEv(void* qthis); // 4
  // proto:  QString QUuid::toString();
extern void _ZNK5QUuid8toStringEv(void* qthis); // 4
  // proto:  void QUuid::QUuid();
extern void _ZN5QUuidC2Ev(void* qthis); // 1
  // proto:  void QUuid::QUuid(const QString & );
extern void _ZN5QUuidC2ERK7QString(void* qthis, void* arg0); // 3
  // proto:  void QUuid::QUuid(const char * );
extern void _ZN5QUuidC2EPKc(void* qthis, unsigned char* arg0); // 3
  // proto:  void QUuid::QUuid(const QByteArray & );
extern void _ZN5QUuidC2ERK10QByteArray(void* qthis, void* arg0); // 3
  // proto:  void QUuid::QUuid(uint l, ushort w1, ushort w2, uchar b1, uchar b2, uchar b3, uchar b4, uchar b5, uchar b6, uchar b7, uchar b8);
extern void _ZN5QUuidC2Ejtthhhhhhhh(void* qthis, int32_t arg0, int16_t arg1, int16_t arg2, unsigned char arg3, unsigned char arg4, unsigned char arg5, unsigned char arg6, unsigned char arg7, unsigned char arg8, unsigned char arg9, unsigned char arg10); // 1
  // proto:  QByteArray QUuid::toRfc4122();
extern void _ZNK5QUuid9toRfc4122Ev(void* qthis); // 4
  // proto:  QByteArray QUuid::toByteArray();
extern void _ZNK5QUuid11toByteArrayEv(void* qthis); // 4
  // proto: static QUuid QUuid::createUuid();
extern void _ZN5QUuid10createUuidEv(); // 4
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

// class sizeof(QUuid)=16
type QUuid struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// createUuidV5(const class QUuid &, const class QByteArray &)
func (this *QUuid) createUuidV5_s(args ...interface{}) () {
  // createUuidV5(const class QUuid &, const class QByteArray &)
  // createUuidV5(const class QUuid &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUuid{}) // "const QUuid &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUuid{}) // "const QUuid &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QUuid12createUuidV5ERKS_RK10QByteArray
    // invoke: QUuid createUuidV5(const class QUuid &, const class QByteArray &)
    var arg0 = args[0].(QUuid).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QUuid12createUuidV5ERKS_RK10QByteArray(arg0, arg1)
  case 1:
    // invoke: _ZN5QUuid12createUuidV5ERKS_RK7QString
    // invoke: QUuid createUuidV5(const class QUuid &, const class QString &)
    var arg0 = args[0].(QUuid).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QUuid12createUuidV5ERKS_RK7QString(arg0, arg1)
  default:
    qtrt.ErrorResolve("QUuid", "createUuidV5", args)
  }

}

// createUuidV3(const class QUuid &, const class QString &)
func (this *QUuid) createUuidV3_s(args ...interface{}) () {
  // createUuidV3(const class QUuid &, const class QString &)
  // createUuidV3(const class QUuid &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUuid{}) // "const QUuid &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QUuid{}) // "const QUuid &"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QUuid12createUuidV3ERKS_RK7QString
    // invoke: QUuid createUuidV3(const class QUuid &, const class QString &)
    var arg0 = args[0].(QUuid).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QUuid12createUuidV3ERKS_RK7QString(arg0, arg1)
  case 1:
    // invoke: _ZN5QUuid12createUuidV3ERKS_RK10QByteArray
    // invoke: QUuid createUuidV3(const class QUuid &, const class QByteArray &)
    var arg0 = args[0].(QUuid).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN5QUuid12createUuidV3ERKS_RK10QByteArray(arg0, arg1)
  default:
    qtrt.ErrorResolve("QUuid", "createUuidV3", args)
  }

}

// fromRfc4122(const class QByteArray &)
func (this *QUuid) fromRfc4122_s(args ...interface{}) () {
  // fromRfc4122(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QUuid11fromRfc4122ERK10QByteArray
    // invoke: QUuid fromRfc4122(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN5QUuid11fromRfc4122ERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QUuid", "fromRfc4122", args)
  }

}

// variant()
func (this *QUuid) variant(args ...interface{}) () {
  // variant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QUuid7variantEv
    // invoke: QUuid::Variant variant()
    C._ZNK5QUuid7variantEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "variant", args)
  }

}

// isNull()
func (this *QUuid) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QUuid6isNullEv
    // invoke: bool isNull()
    C._ZNK5QUuid6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "isNull", args)
  }

}

// version()
func (this *QUuid) version(args ...interface{}) () {
  // version()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QUuid7versionEv
    // invoke: QUuid::Version version()
    C._ZNK5QUuid7versionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "version", args)
  }

}

// toString()
func (this *QUuid) toString(args ...interface{}) () {
  // toString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QUuid8toStringEv
    // invoke: QString toString()
    C._ZNK5QUuid8toStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "toString", args)
  }

}

// QUuid()
func NewQUuid(args ...interface{}) QUuid {
  // QUuid()
  // QUuid(const class QString &)
  // QUuid(const char *)
  // QUuid(const class QByteArray &)
  // QUuid(uint, ushort, ushort, uchar, uchar, uchar, uchar, uchar, uchar, uchar, uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "uint"
  vtys[4][1] = qtrt.Int16Ty(false) // "ushort"
  vtys[4][2] = qtrt.Int16Ty(false) // "ushort"
  vtys[4][3] = qtrt.ByteTy(false) // "uchar"
  vtys[4][4] = qtrt.ByteTy(false) // "uchar"
  vtys[4][5] = qtrt.ByteTy(false) // "uchar"
  vtys[4][6] = qtrt.ByteTy(false) // "uchar"
  vtys[4][7] = qtrt.ByteTy(false) // "uchar"
  vtys[4][8] = qtrt.ByteTy(false) // "uchar"
  vtys[4][9] = qtrt.ByteTy(false) // "uchar"
  vtys[4][10] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QUuidC1Ev
    // invoke: void QUuid()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QUuidC2Ev(qthis)
  case 1:
    // invoke: _ZN5QUuidC1ERK7QString
    // invoke: void QUuid(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QUuidC2ERK7QString(qthis, arg0)
  case 2:
    // invoke: _ZN5QUuidC1EPKc
    // invoke: void QUuid(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QUuidC2EPKc(qthis, arg0)
  case 3:
    // invoke: _ZN5QUuidC1ERK10QByteArray
    // invoke: void QUuid(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QUuidC2ERK10QByteArray(qthis, arg0)
  case 4:
    // invoke: _ZN5QUuidC1Ejtthhhhhhhh
    // invoke: void QUuid(uint, ushort, ushort, uchar, uchar, uchar, uchar, uchar, uchar, uchar, uchar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int16_t(args[1].(int16))
    if false {fmt.Println(arg1)}
    var arg2 = C.int16_t(args[2].(int16))
    if false {fmt.Println(arg2)}
    var arg3 = C.uchar(args[3].(byte))
    if false {fmt.Println(arg3)}
    var arg4 = C.uchar(args[4].(byte))
    if false {fmt.Println(arg4)}
    var arg5 = C.uchar(args[5].(byte))
    if false {fmt.Println(arg5)}
    var arg6 = C.uchar(args[6].(byte))
    if false {fmt.Println(arg6)}
    var arg7 = C.uchar(args[7].(byte))
    if false {fmt.Println(arg7)}
    var arg8 = C.uchar(args[8].(byte))
    if false {fmt.Println(arg8)}
    var arg9 = C.uchar(args[9].(byte))
    if false {fmt.Println(arg9)}
    var arg10 = C.uchar(args[10].(byte))
    if false {fmt.Println(arg10)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN5QUuidC2Ejtthhhhhhhh(qthis, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
  default:
    qtrt.ErrorResolve("QUuid", "QUuid", args)
  }

  return QUuid{}
}

// toRfc4122()
func (this *QUuid) toRfc4122(args ...interface{}) () {
  // toRfc4122()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QUuid9toRfc4122Ev
    // invoke: QByteArray toRfc4122()
    C._ZNK5QUuid9toRfc4122Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "toRfc4122", args)
  }

}

// toByteArray()
func (this *QUuid) toByteArray(args ...interface{}) () {
  // toByteArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK5QUuid11toByteArrayEv
    // invoke: QByteArray toByteArray()
    C._ZNK5QUuid11toByteArrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "toByteArray", args)
  }

}

// createUuid()
func (this *QUuid) createUuid_s(args ...interface{}) () {
  // createUuid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN5QUuid10createUuidEv
    // invoke: QUuid createUuid()
    C._ZN5QUuid10createUuidEv()
  default:
    qtrt.ErrorResolve("QUuid", "createUuid", args)
  }

}

// <= body block end

