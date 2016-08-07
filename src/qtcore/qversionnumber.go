package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtCore/qversionnumber.h
// dst-file: /src/core/qversionnumber.go
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
  // proto:  int QVersionNumber::majorVersion();
extern int32_t C_ZNK14QVersionNumber12majorVersionEv(void* qthis); // 2
  // proto: static int QVersionNumber::compare(const QVersionNumber & v1, const QVersionNumber & v2);
extern int32_t C_ZN14QVersionNumber7compareERKS_S1_(void* arg0, void* arg1); // 4
  // proto: static QVersionNumber QVersionNumber::commonPrefix(const QVersionNumber & v1, const QVersionNumber & v2);
extern void* C_ZN14QVersionNumber12commonPrefixERKS_S1_(void* arg0, void* arg1); // 4
  // proto:  int QVersionNumber::segmentCount();
extern int32_t C_ZNK14QVersionNumber12segmentCountEv(void* qthis); // 2
  // proto:  int QVersionNumber::minorVersion();
extern int32_t C_ZNK14QVersionNumber12minorVersionEv(void* qthis); // 2
  // proto:  void QVersionNumber::QVersionNumber();
extern void* C_ZN14QVersionNumberC2Ev(); // 1
  // proto:  void QVersionNumber::QVersionNumber(int maj, int min, int mic);
extern void* C_ZN14QVersionNumberC2Eiii(int32_t arg0, int32_t arg1, int32_t arg2); // 1
  // proto:  void QVersionNumber::QVersionNumber(int maj);
extern void* C_ZN14QVersionNumberC2Ei(int32_t arg0); // 1
  // proto:  void QVersionNumber::QVersionNumber(int maj, int min);
extern void* C_ZN14QVersionNumberC2Eii(int32_t arg0, int32_t arg1); // 1
  // proto: static QVersionNumber QVersionNumber::fromString(const QString & string, int * suffixIndex);
extern void* C_ZN14QVersionNumber10fromStringERK7QStringPi(void* arg0, void* arg1); // 4
  // proto:  QVector<int> QVersionNumber::segments();
extern void C_ZNK14QVersionNumber8segmentsEv(void* qthis); // 4
  // proto:  bool QVersionNumber::isNormalized();
extern bool C_ZNK14QVersionNumber12isNormalizedEv(void* qthis); // 2
  // proto:  int QVersionNumber::microVersion();
extern int32_t C_ZNK14QVersionNumber12microVersionEv(void* qthis); // 2
  // proto:  int QVersionNumber::segmentAt(int index);
extern int32_t C_ZNK14QVersionNumber9segmentAtEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QVersionNumber::isNull();
extern bool C_ZNK14QVersionNumber6isNullEv(void* qthis); // 2
  // proto:  QString QVersionNumber::toString();
extern void* C_ZNK14QVersionNumber8toStringEv(void* qthis); // 4
  // proto:  QVersionNumber QVersionNumber::normalized();
extern void* C_ZNK14QVersionNumber10normalizedEv(void* qthis); // 4
  // proto:  bool QVersionNumber::isPrefixOf(const QVersionNumber & other);
extern bool C_ZNK14QVersionNumber10isPrefixOfERKS_(void* qthis, void* arg0); // 4
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

// class sizeof(QVersionNumber)=8
type QVersionNumber struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// majorVersion()
func (this *QVersionNumber) Majorversion(args ...interface{}) (ret interface{}) {
  // majorVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber12majorVersionEv
    // invoke: int majorVersion()
    var ret0 = C.C_ZNK14QVersionNumber12majorVersionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "majorVersion", args)
  }

  return
}

// compare(const class QVersionNumber &, const class QVersionNumber &)
func (this *QVersionNumber) Compare_S(args ...interface{}) (ret interface{}) {
  // compare(const class QVersionNumber &, const class QVersionNumber &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVersionNumber{}) // "const QVersionNumber &"
  vtys[0][1] = reflect.TypeOf(QVersionNumber{}) // "const QVersionNumber &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QVersionNumber7compareERKS_S1_
    // invoke: int compare(const class QVersionNumber &, const class QVersionNumber &)
    var arg0 = args[0].(*QVersionNumber).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVersionNumber).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QVersionNumber7compareERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "compare", args)
  }

  return
}

// commonPrefix(const class QVersionNumber &, const class QVersionNumber &)
func (this *QVersionNumber) Commonprefix_S(args ...interface{}) (ret interface{}) {
  // commonPrefix(const class QVersionNumber &, const class QVersionNumber &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVersionNumber{}) // "const QVersionNumber &"
  vtys[0][1] = reflect.TypeOf(QVersionNumber{}) // "const QVersionNumber &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QVersionNumber12commonPrefixERKS_S1_
    // invoke: QVersionNumber commonPrefix(const class QVersionNumber &, const class QVersionNumber &)
    var arg0 = args[0].(*QVersionNumber).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QVersionNumber).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QVersionNumber12commonPrefixERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVersionNumber{}) // "QVersionNumber"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "commonPrefix", args)
  }

  return
}

// segmentCount()
func (this *QVersionNumber) Segmentcount(args ...interface{}) (ret interface{}) {
  // segmentCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber12segmentCountEv
    // invoke: int segmentCount()
    var ret0 = C.C_ZNK14QVersionNumber12segmentCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "segmentCount", args)
  }

  return
}

// minorVersion()
func (this *QVersionNumber) Minorversion(args ...interface{}) (ret interface{}) {
  // minorVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber12minorVersionEv
    // invoke: int minorVersion()
    var ret0 = C.C_ZNK14QVersionNumber12minorVersionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "minorVersion", args)
  }

  return
}

// QVersionNumber()
func NewQVersionNumber(args ...interface{}) *QVersionNumber {
  // QVersionNumber()
  // QVersionNumber(int, int, int)
  // QVersionNumber(int)
  // QVersionNumber(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QVersionNumberC1Ev
    // invoke: void QVersionNumber()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QVersionNumberC2Ev()
    return &QVersionNumber{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QVersionNumberC1Eiii
    // invoke: void QVersionNumber(int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QVersionNumberC2Eiii(arg0, arg1, arg2)
    return &QVersionNumber{Qclsinst:qthis}
  case 2:
    // invoke: _ZN14QVersionNumberC1Ei
    // invoke: void QVersionNumber(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QVersionNumberC2Ei(arg0)
    return &QVersionNumber{Qclsinst:qthis}
  case 3:
    // invoke: _ZN14QVersionNumberC1Eii
    // invoke: void QVersionNumber(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QVersionNumberC2Eii(arg0, arg1)
    return &QVersionNumber{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QVersionNumber", "QVersionNumber", args)
  }

  return nil // QVersionNumber{Qclsinst:qthis}
}

// fromString(const class QString &, int *)
func (this *QVersionNumber) Fromstring_S(args ...interface{}) (ret interface{}) {
  // fromString(const class QString &, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QVersionNumber10fromStringERK7QStringPi
    // invoke: QVersionNumber fromString(const class QString &, int *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QVersionNumber10fromStringERK7QStringPi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVersionNumber{}) // "QVersionNumber"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "fromString", args)
  }

  return
}

// segments()
func (this *QVersionNumber) Segments(args ...interface{}) () {
  // segments()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber8segmentsEv
    // invoke: QVector<int> segments()
    C.C_ZNK14QVersionNumber8segmentsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QVersionNumber", "segments", args)
  }

  return
}

// isNormalized()
func (this *QVersionNumber) Isnormalized(args ...interface{}) (ret interface{}) {
  // isNormalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber12isNormalizedEv
    // invoke: bool isNormalized()
    var ret0 = C.C_ZNK14QVersionNumber12isNormalizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "isNormalized", args)
  }

  return
}

// microVersion()
func (this *QVersionNumber) Microversion(args ...interface{}) (ret interface{}) {
  // microVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber12microVersionEv
    // invoke: int microVersion()
    var ret0 = C.C_ZNK14QVersionNumber12microVersionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "microVersion", args)
  }

  return
}

// segmentAt(int)
func (this *QVersionNumber) Segmentat(args ...interface{}) (ret interface{}) {
  // segmentAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber9segmentAtEi
    // invoke: int segmentAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QVersionNumber9segmentAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "segmentAt", args)
  }

  return
}

// isNull()
func (this *QVersionNumber) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK14QVersionNumber6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "isNull", args)
  }

  return
}

// toString()
func (this *QVersionNumber) Tostring(args ...interface{}) (ret interface{}) {
  // toString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber8toStringEv
    // invoke: QString toString()
    var ret0 = C.C_ZNK14QVersionNumber8toStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "toString", args)
  }

  return
}

// normalized()
func (this *QVersionNumber) Normalized(args ...interface{}) (ret interface{}) {
  // normalized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber10normalizedEv
    // invoke: QVersionNumber normalized()
    var ret0 = C.C_ZNK14QVersionNumber10normalizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVersionNumber{}) // "QVersionNumber"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "normalized", args)
  }

  return
}

// isPrefixOf(const class QVersionNumber &)
func (this *QVersionNumber) Isprefixof(args ...interface{}) (ret interface{}) {
  // isPrefixOf(const class QVersionNumber &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVersionNumber{}) // "const QVersionNumber &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QVersionNumber10isPrefixOfERKS_
    // invoke: bool isPrefixOf(const class QVersionNumber &)
    var arg0 = args[0].(*QVersionNumber).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QVersionNumber10isPrefixOfERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QVersionNumber", "isPrefixOf", args)
  }

  return
}

// <= body block end

