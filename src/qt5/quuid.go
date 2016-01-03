package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QUuid::QUuid(const QString & );
extern void* dector_ZN5QUuidC1ERK7QString(void* arg0);
extern void _ZN5QUuidC1ERK7QString(void* qthis, void* arg0);
  // proto:  QByteArray QUuid::toRfc4122();
extern void _ZNK5QUuid9toRfc4122Ev(void* qthis);
  // proto:  QString QUuid::toString();
extern void _ZNK5QUuid8toStringEv(void* qthis);
  // proto:  bool QUuid::isNull();
extern void _ZNK5QUuid6isNullEv(void* qthis);
  // proto: static QUuid QUuid::createUuidV5(const QUuid & ns, const QString & baseData);
extern void demth_ZN5QUuid12createUuidV5ERKS_RK7QString(void* arg0, void* arg1);
  // proto: static QUuid QUuid::createUuid();
extern void _ZN5QUuid10createUuidEv();
  // proto:  void QUuid::QUuid(uint l, ushort w1, ushort w2, uchar b1, uchar b2, uchar b3, uchar b4, uchar b5, uchar b6, uchar b7, uchar b8);
extern void* dector_ZN5QUuidC1Ejtthhhhhhhh(int32_t arg0, int16_t arg1, int16_t arg2, unsigned char arg3, unsigned char arg4, unsigned char arg5, unsigned char arg6, unsigned char arg7, unsigned char arg8, unsigned char arg9, unsigned char arg10);
extern void _ZN5QUuidC1Ejtthhhhhhhh(void* qthis, int32_t arg0, int16_t arg1, int16_t arg2, unsigned char arg3, unsigned char arg4, unsigned char arg5, unsigned char arg6, unsigned char arg7, unsigned char arg8, unsigned char arg9, unsigned char arg10);
  // proto:  void QUuid::QUuid(const QByteArray & );
extern void* dector_ZN5QUuidC1ERK10QByteArray(void* arg0);
extern void _ZN5QUuidC1ERK10QByteArray(void* qthis, void* arg0);
  // proto: static QUuid QUuid::createUuidV3(const QUuid & ns, const QString & baseData);
extern void demth_ZN5QUuid12createUuidV3ERKS_RK7QString(void* arg0, void* arg1);
  // proto:  void QUuid::QUuid();
extern void* dector_ZN5QUuidC1Ev();
extern void _ZN5QUuidC1Ev(void* qthis);
  // proto:  QByteArray QUuid::toByteArray();
extern void _ZNK5QUuid11toByteArrayEv(void* qthis);
  // proto:  void QUuid::QUuid(const char * );
extern void* dector_ZN5QUuidC1EPKc(unsigned char* arg0);
extern void _ZN5QUuidC1EPKc(void* qthis, unsigned char* arg0);
  // proto: static QUuid QUuid::createUuidV5(const QUuid & ns, const QByteArray & baseData);
extern void _ZN5QUuid12createUuidV5ERKS_RK10QByteArray(void* arg0, void* arg1);
  // proto: static QUuid QUuid::fromRfc4122(const QByteArray & );
extern void _ZN5QUuid11fromRfc4122ERK10QByteArray(void* arg0);
  // proto: static QUuid QUuid::createUuidV3(const QUuid & ns, const QByteArray & baseData);
extern void _ZN5QUuid12createUuidV3ERKS_RK10QByteArray(void* arg0, void* arg1);
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

  // proto:  void QUuid::QUuid(const QString & );
func NewQUuid(args ...interface{}) QUuid {
  return QUuid{}
}

  // proto:  QByteArray QUuid::toRfc4122();
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

  // proto:  QString QUuid::toString();
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

  // proto:  bool QUuid::isNull();
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

  // proto: static QUuid QUuid::createUuidV5(const QUuid & ns, const QString & baseData);
func (this *QUuid) createUuidV5_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUuid", "createUuidV5", args)
  }

}

  // proto: static QUuid QUuid::createUuid();
func (this *QUuid) createUuid_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUuid", "createUuid", args)
  }

}

  // proto: static QUuid QUuid::createUuidV3(const QUuid & ns, const QString & baseData);
func (this *QUuid) createUuidV3_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUuid", "createUuidV3", args)
  }

}

  // proto:  QByteArray QUuid::toByteArray();
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

  // proto: static QUuid QUuid::fromRfc4122(const QByteArray & );
func (this *QUuid) fromRfc4122_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QUuid", "fromRfc4122", args)
  }

}

// <= body block end

