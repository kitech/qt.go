package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qstringbuilder.h
// dst-file: /src/core/qstringbuilder.go
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
  // proto: static void QConcatenable<QCharRef>::appendTo(QCharRef c, QChar *& out);
extern void C_ZN13QConcatenableI8QCharRefE8appendToES0_RP5QChar(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<QCharRef>::size(QCharRef );
extern int32_t C_ZN13QConcatenableI8QCharRefE4sizeES0_(void* arg0); // 2
  // proto: static void QConcatenable<QString>::appendTo(const QString & a, QChar *& out);
extern void C_ZN13QConcatenableI7QStringE8appendToERKS0_RP5QChar(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<QString>::size(const QString & a);
extern int32_t C_ZN13QConcatenableI7QStringE4sizeERKS0_(void* arg0); // 2
  // proto: static void QConcatenable<char>::appendTo(const char c, QChar *& out);
extern void C_ZN13QConcatenableIcE8appendToEcRP5QChar(unsigned char arg0, void* arg1); // 2
  // proto: static void QConcatenable<char>::appendTo(const char c, char *& out);
extern void C_ZN13QConcatenableIcE8appendToEcRPc(unsigned char arg0, void* arg1); // 2
  // proto: static int QConcatenable<char>::size(const char );
extern int32_t C_ZN13QConcatenableIcE4sizeEc(unsigned char arg0); // 2
  // proto: static void QConcatenable<QByteArray>::appendTo(const QByteArray & ba, char *& out);
extern void C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RPc(void* arg0, void* arg1); // 2
  // proto: static void QConcatenable<QByteArray>::appendTo(const QByteArray & ba, QChar *& out);
extern void C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RP5QChar(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<QByteArray>::size(const QByteArray & ba);
extern int32_t C_ZN13QConcatenableI10QByteArrayE4sizeERKS0_(void* arg0); // 2
  // proto: static void QConcatenable<QChar>::appendTo(const QChar c, QChar *& out);
extern void C_ZN13QConcatenableI5QCharE8appendToES0_RPS0_(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<QChar>::size(const QChar );
extern int32_t C_ZN13QConcatenableI5QCharE4sizeES0_(void* arg0); // 2
  // proto:  void QStringBuilder<QByteArray, QByteArray>::QStringBuilder(const QByteArray & a_, const QByteArray & b_);
extern void* C_ZN14QStringBuilderI10QByteArrayS0_EC2ERKS0_S3_(void* arg0, void* arg1); // 1
  // proto: static void QConcatenable<const char *>::appendTo(const char * a, QChar *& out);
extern void C_ZN13QConcatenableIPKcE8appendToES1_RP5QChar(void* arg0, void* arg1); // 2
  // proto: static void QConcatenable<const char *>::appendTo(const char * a, char *& out);
extern void C_ZN13QConcatenableIPKcE8appendToES1_RPc(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<const char *>::size(const char * a);
extern int32_t C_ZN13QConcatenableIPKcE4sizeES1_(void* arg0); // 2
  // proto:  void QStringBuilder<QString, QString>::QStringBuilder(const QString & a_, const QString & b_);
extern void* C_ZN14QStringBuilderI7QStringS0_EC2ERKS0_S3_(void* arg0, void* arg1); // 1
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

// class sizeof(QConcatenable<QLatin1String>)=1
type QConcatenableLQLatin1StringG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QCharRef>)=1
type QConcatenableLQCharRefG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QStringRef>)=1
type QConcatenableLQStringRefG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractConcatenable)=1
type QAbstractConcatenable struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QString>)=1
type QConcatenableLQStringG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QLatin1Char>)=1
type QConcatenableLQLatin1CharG struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<char>)=1
type QConcatenableLcharG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QByteArray>)=1
type QConcatenableLQByteArrayG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QChar::SpecialCharacter>)=1
type QConcatenableLQChar__SpecialCharacterG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QChar>)=1
type QConcatenableLQCharG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStringBuilder<QByteArray, QByteArray>)=1
type QStringBuilderLQByteArray_EQByteArrayG struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<const char *>)=1
type QConcatenableLconstEcharEPG struct {
  /*qbase*/ QAbstractConcatenable;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStringBuilder<QString, QString>)=1
type QStringBuilderLQString_EQStringG struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// appendTo(class QCharRef, class QChar *&)
func (this *QConcatenableLQCharRefG) AppendTo_s(args ...interface{}) () {
  // appendTo(class QCharRef, class QChar *&)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCharRef{}) // "QCharRef"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar *&"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI8QCharRefE8appendToES0_RP5QChar
    // invoke: void appendTo(class QCharRef, class QChar *&)
    var arg0 = args[0].(*QCharRef).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI8QCharRefE8appendToES0_RP5QChar(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QCharRef>", "appendTo", args)
  }

  return
}

// size(class QCharRef)
func (this *QConcatenableLQCharRefG) Size_s(args ...interface{}) (ret interface{}) {
  // size(class QCharRef)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCharRef{}) // "QCharRef"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI8QCharRefE4sizeES0_
    // invoke: int size(class QCharRef)
    var arg0 = args[0].(*QCharRef).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QConcatenableI8QCharRefE4sizeES0_(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConcatenable<QCharRef>", "size", args)
  }

  return
}

// appendTo(const class QString &, class QChar *&)
func (this *QConcatenableLQStringG) AppendTo_s(args ...interface{}) () {
  // appendTo(const class QString &, class QChar *&)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar *&"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI7QStringE8appendToERKS0_RP5QChar
    // invoke: void appendTo(const class QString &, class QChar *&)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI7QStringE8appendToERKS0_RP5QChar(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QString>", "appendTo", args)
  }

  return
}

// size(const class QString &)
func (this *QConcatenableLQStringG) Size_s(args ...interface{}) (ret interface{}) {
  // size(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI7QStringE4sizeERKS0_
    // invoke: int size(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QConcatenableI7QStringE4sizeERKS0_(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConcatenable<QString>", "size", args)
  }

  return
}

// appendTo(const char, class QChar *&)
func (this *QConcatenableLcharG) AppendTo_s(args ...interface{}) () {
  // appendTo(const char, class QChar *&)
  // appendTo(const char, char *&)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "const char"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar *&"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "const char"
  vtys[1][1] = qtrt.StringTy(false) // "char *&"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableIcE8appendToEcRP5QChar
    // invoke: void appendTo(const char, class QChar *&)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableIcE8appendToEcRP5QChar(arg0, arg1)
  case 1:
    // invoke: _ZN13QConcatenableIcE8appendToEcRPc
    // invoke: void appendTo(const char, char *&)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN13QConcatenableIcE8appendToEcRPc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<char>", "appendTo", args)
  }

  return
}

// size(const char)
func (this *QConcatenableLcharG) Size_s(args ...interface{}) (ret interface{}) {
  // size(const char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "const char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableIcE4sizeEc
    // invoke: int size(const char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QConcatenableIcE4sizeEc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConcatenable<char>", "size", args)
  }

  return
}

// appendTo(const class QByteArray &, char *&)
func (this *QConcatenableLQByteArrayG) AppendTo_s(args ...interface{}) () {
  // appendTo(const class QByteArray &, char *&)
  // appendTo(const class QByteArray &, class QChar *&)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.StringTy(false) // "char *&"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = reflect.TypeOf(QChar{}) // "QChar *&"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI10QByteArrayE8appendToERKS0_RPc
    // invoke: void appendTo(const class QByteArray &, char *&)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RPc(arg0, arg1)
  case 1:
    // invoke: _ZN13QConcatenableI10QByteArrayE8appendToERKS0_RP5QChar
    // invoke: void appendTo(const class QByteArray &, class QChar *&)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RP5QChar(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QByteArray>", "appendTo", args)
  }

  return
}

// size(const class QByteArray &)
func (this *QConcatenableLQByteArrayG) Size_s(args ...interface{}) (ret interface{}) {
  // size(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI10QByteArrayE4sizeERKS0_
    // invoke: int size(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QConcatenableI10QByteArrayE4sizeERKS0_(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConcatenable<QByteArray>", "size", args)
  }

  return
}

// appendTo(const class QChar, class QChar *&)
func (this *QConcatenableLQCharG) AppendTo_s(args ...interface{}) () {
  // appendTo(const class QChar, class QChar *&)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar *&"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI5QCharE8appendToES0_RPS0_
    // invoke: void appendTo(const class QChar, class QChar *&)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI5QCharE8appendToES0_RPS0_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QChar>", "appendTo", args)
  }

  return
}

// size(const class QChar)
func (this *QConcatenableLQCharG) Size_s(args ...interface{}) (ret interface{}) {
  // size(const class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableI5QCharE4sizeES0_
    // invoke: int size(const class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QConcatenableI5QCharE4sizeES0_(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConcatenable<QChar>", "size", args)
  }

  return
}

// QStringBuilder(const class QByteArray &, const class QByteArray &)
func GcfreeQStringBuilderLQByteArray_EQByteArrayG(this *QStringBuilderLQByteArray_EQByteArrayG) {
  qtrt.UniverseFree(this)
}
func NewQStringBuilderLQByteArray_EQByteArrayG(args ...interface{}) *QStringBuilderLQByteArray_EQByteArrayG {
  // QStringBuilder(const class QByteArray &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStringBuilderI10QByteArrayS0_EC1ERKS0_S3_
    // invoke: void QStringBuilder(const class QByteArray &, const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStringBuilderI10QByteArrayS0_EC2ERKS0_S3_(arg0, arg1)
    this := &QStringBuilderLQByteArray_EQByteArrayG{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStringBuilderLQByteArray_EQByteArrayG)
    return this
  default:
    qtrt.ErrorResolve("QStringBuilder<QByteArray, QByteArray>", "QStringBuilder", args)
  }

  return nil // QStringBuilderLQByteArray_EQByteArrayG{Qclsinst:qthis}
}

// appendTo(const char *, class QChar *&)
func (this *QConcatenableLconstEcharEPG) AppendTo_s(args ...interface{}) () {
  // appendTo(const char *, class QChar *&)
  // appendTo(const char *, char *&)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar *&"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.StringTy(false) // "char *&"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableIPKcE8appendToES1_RP5QChar
    // invoke: void appendTo(const char *, class QChar *&)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableIPKcE8appendToES1_RP5QChar(arg0, arg1)
  case 1:
    // invoke: _ZN13QConcatenableIPKcE8appendToES1_RPc
    // invoke: void appendTo(const char *, char *&)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN13QConcatenableIPKcE8appendToES1_RPc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<const char *>", "appendTo", args)
  }

  return
}

// size(const char *)
func (this *QConcatenableLconstEcharEPG) Size_s(args ...interface{}) (ret interface{}) {
  // size(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QConcatenableIPKcE4sizeES1_
    // invoke: int size(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN13QConcatenableIPKcE4sizeES1_(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QConcatenable<const char *>", "size", args)
  }

  return
}

// QStringBuilder(const class QString &, const class QString &)
func GcfreeQStringBuilderLQString_EQStringG(this *QStringBuilderLQString_EQStringG) {
  qtrt.UniverseFree(this)
}
func NewQStringBuilderLQString_EQStringG(args ...interface{}) *QStringBuilderLQString_EQStringG {
  // QStringBuilder(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStringBuilderI7QStringS0_EC1ERKS0_S3_
    // invoke: void QStringBuilder(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStringBuilderI7QStringS0_EC2ERKS0_S3_(arg0, arg1)
    this := &QStringBuilderLQString_EQStringG{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStringBuilderLQString_EQStringG)
    return this
  default:
    qtrt.ErrorResolve("QStringBuilder<QString, QString>", "QStringBuilder", args)
  }

  return nil // QStringBuilderLQString_EQStringG{Qclsinst:qthis}
}

// <= body block end

