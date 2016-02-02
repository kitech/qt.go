package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZN5QUuid12createUuidV5ERKS_RK10QByteArray(void* arg0, void* arg1); // 4
  // proto: static QUuid QUuid::createUuidV5(const QUuid & ns, const QString & baseData);
extern void* C_ZN5QUuid12createUuidV5ERKS_RK7QString(void* arg0, void* arg1); // 2
  // proto: static QUuid QUuid::createUuidV3(const QUuid & ns, const QString & baseData);
extern void* C_ZN5QUuid12createUuidV3ERKS_RK7QString(void* arg0, void* arg1); // 2
  // proto: static QUuid QUuid::createUuidV3(const QUuid & ns, const QByteArray & baseData);
extern void* C_ZN5QUuid12createUuidV3ERKS_RK10QByteArray(void* arg0, void* arg1); // 4
  // proto: static QUuid QUuid::fromRfc4122(const QByteArray & );
extern void* C_ZN5QUuid11fromRfc4122ERK10QByteArray(void* arg0); // 4
  // proto:  QUuid::Variant QUuid::variant();
extern void C_ZNK5QUuid7variantEv(void* qthis); // 4
  // proto:  bool QUuid::isNull();
extern bool C_ZNK5QUuid6isNullEv(void* qthis); // 4
  // proto:  QUuid::Version QUuid::version();
extern void C_ZNK5QUuid7versionEv(void* qthis); // 4
  // proto:  QString QUuid::toString();
extern void* C_ZNK5QUuid8toStringEv(void* qthis); // 4
  // proto:  void QUuid::QUuid();
extern void* C_ZN5QUuidC2Ev(); // 1
  // proto:  void QUuid::QUuid(const QString & );
extern void* C_ZN5QUuidC2ERK7QString(void* arg0); // 3
  // proto:  void QUuid::QUuid(const char * );
extern void* C_ZN5QUuidC2EPKc(void* arg0); // 3
  // proto:  void QUuid::QUuid(const QByteArray & );
extern void* C_ZN5QUuidC2ERK10QByteArray(void* arg0); // 3
  // proto:  void QUuid::QUuid(uint l, ushort w1, ushort w2, uchar b1, uchar b2, uchar b3, uchar b4, uchar b5, uchar b6, uchar b7, uchar b8);
extern void* C_ZN5QUuidC2Ejtthhhhhhhh(int32_t arg0, int16_t arg1, int16_t arg2, unsigned char arg3, unsigned char arg4, unsigned char arg5, unsigned char arg6, unsigned char arg7, unsigned char arg8, unsigned char arg9, unsigned char arg10); // 1
  // proto:  QByteArray QUuid::toRfc4122();
extern void* C_ZNK5QUuid9toRfc4122Ev(void* qthis); // 4
  // proto:  QByteArray QUuid::toByteArray();
extern void* C_ZNK5QUuid11toByteArrayEv(void* qthis); // 4
  // proto: static QUuid QUuid::createUuid();
extern void* C_ZN5QUuid10createUuidEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// createUuidV5(const class QUuid &, const class QByteArray &)
func (this *QUuid) Createuuidv5_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QUuid).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QUuid12createUuidV5ERKS_RK10QByteArray(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUuid{}) // "QUuid"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QUuid12createUuidV5ERKS_RK7QString
    // invoke: QUuid createUuidV5(const class QUuid &, const class QString &)
    var arg0 = args[0].(*QUuid).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QUuid12createUuidV5ERKS_RK7QString(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUuid{}) // "QUuid"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "createUuidV5", args)
  }

  return
}

// createUuidV3(const class QUuid &, const class QString &)
func (this *QUuid) Createuuidv3_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QUuid).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QUuid12createUuidV3ERKS_RK7QString(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUuid{}) // "QUuid"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN5QUuid12createUuidV3ERKS_RK10QByteArray
    // invoke: QUuid createUuidV3(const class QUuid &, const class QByteArray &)
    var arg0 = args[0].(*QUuid).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN5QUuid12createUuidV3ERKS_RK10QByteArray(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUuid{}) // "QUuid"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "createUuidV3", args)
  }

  return
}

// fromRfc4122(const class QByteArray &)
func (this *QUuid) Fromrfc4122_S(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN5QUuid11fromRfc4122ERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUuid{}) // "QUuid"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "fromRfc4122", args)
  }

  return
}

// variant()
func (this *QUuid) Variant(args ...interface{}) () {
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
    C.C_ZNK5QUuid7variantEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "variant", args)
  }

  return
}

// isNull()
func (this *QUuid) Isnull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QUuid6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "isNull", args)
  }

  return
}

// version()
func (this *QUuid) Version(args ...interface{}) () {
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
    C.C_ZNK5QUuid7versionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QUuid", "version", args)
  }

  return
}

// toString()
func (this *QUuid) Tostring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QUuid8toStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "toString", args)
  }

  return
}

// QUuid()
func NewQUuid(args ...interface{}) *QUuid {
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
    qthis = C.C_ZN5QUuidC2Ev()
    return &QUuid{Qclsinst:qthis}
  case 1:
    // invoke: _ZN5QUuidC1ERK7QString
    // invoke: void QUuid(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QUuidC2ERK7QString(arg0)
    return &QUuid{Qclsinst:qthis}
  case 2:
    // invoke: _ZN5QUuidC1EPKc
    // invoke: void QUuid(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QUuidC2EPKc(arg0)
    return &QUuid{Qclsinst:qthis}
  case 3:
    // invoke: _ZN5QUuidC1ERK10QByteArray
    // invoke: void QUuid(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN5QUuidC2ERK10QByteArray(arg0)
    return &QUuid{Qclsinst:qthis}
  case 4:
    // invoke: _ZN5QUuidC1Ejtthhhhhhhh
    // invoke: void QUuid(uint, ushort, ushort, uchar, uchar, uchar, uchar, uchar, uchar, uchar, uchar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int16_t(qtrt.PrimConv(args[1], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg1)}
    var arg2 = C.int16_t(qtrt.PrimConv(args[2], qtrt.Int16Ty(false)).(int16))
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
    qthis = C.C_ZN5QUuidC2Ejtthhhhhhhh(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
    return &QUuid{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QUuid", "QUuid", args)
  }

  return nil // QUuid{Qclsinst:qthis}
}

// toRfc4122()
func (this *QUuid) Torfc4122(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QUuid9toRfc4122Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "toRfc4122", args)
  }

  return
}

// toByteArray()
func (this *QUuid) Tobytearray(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK5QUuid11toByteArrayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "toByteArray", args)
  }

  return
}

// createUuid()
func (this *QUuid) Createuuid_S(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN5QUuid10createUuidEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUuid{}) // "QUuid"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QUuid", "createUuid", args)
  }

  return
}

// <= body block end

