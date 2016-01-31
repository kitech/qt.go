package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZN13QConcatenableI8QCharRefE4sizeES0_(void* arg0); // 2
  // proto: static void QConcatenable<QString>::appendTo(const QString & a, QChar *& out);
extern void C_ZN13QConcatenableI7QStringE8appendToERKS0_RP5QChar(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<QString>::size(const QString & a);
extern void C_ZN13QConcatenableI7QStringE4sizeERKS0_(void* arg0); // 2
  // proto: static void QConcatenable<char>::appendTo(const char c, QChar *& out);
extern void C_ZN13QConcatenableIcE8appendToEcRP5QChar(unsigned char arg0, void* arg1); // 2
  // proto: static void QConcatenable<char>::appendTo(const char c, char *& out);
extern void C_ZN13QConcatenableIcE8appendToEcRPc(unsigned char arg0, unsigned char* arg1); // 2
  // proto: static int QConcatenable<char>::size(const char );
extern void C_ZN13QConcatenableIcE4sizeEc(unsigned char arg0); // 2
  // proto: static void QConcatenable<QByteArray>::appendTo(const QByteArray & ba, char *& out);
extern void C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RPc(void* arg0, unsigned char* arg1); // 2
  // proto: static void QConcatenable<QByteArray>::appendTo(const QByteArray & ba, QChar *& out);
extern void C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RP5QChar(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<QByteArray>::size(const QByteArray & ba);
extern void C_ZN13QConcatenableI10QByteArrayE4sizeERKS0_(void* arg0); // 2
  // proto: static void QConcatenable<QChar>::appendTo(const QChar c, QChar *& out);
extern void C_ZN13QConcatenableI5QCharE8appendToES0_RPS0_(void* arg0, void* arg1); // 2
  // proto: static int QConcatenable<QChar>::size(const QChar );
extern void C_ZN13QConcatenableI5QCharE4sizeES0_(void* arg0); // 2
  // proto:  void QStringBuilder<QByteArray, QByteArray>::QStringBuilder(const QByteArray & a_, const QByteArray & b_);
extern void C_ZN14QStringBuilderI10QByteArrayS0_EC2ERKS0_S3_(void* qthis, void* arg0, void* arg1); // 1
  // proto: static void QConcatenable<const char *>::appendTo(const char * a, QChar *& out);
extern void C_ZN13QConcatenableIPKcE8appendToES1_RP5QChar(unsigned char* arg0, void* arg1); // 2
  // proto: static void QConcatenable<const char *>::appendTo(const char * a, char *& out);
extern void C_ZN13QConcatenableIPKcE8appendToES1_RPc(unsigned char* arg0, unsigned char* arg1); // 2
  // proto: static int QConcatenable<const char *>::size(const char * a);
extern void C_ZN13QConcatenableIPKcE4sizeES1_(unsigned char* arg0); // 2
  // proto:  void QStringBuilder<QString, QString>::QStringBuilder(const QString & a_, const QString & b_);
extern void C_ZN14QStringBuilderI7QStringS0_EC2ERKS0_S3_(void* qthis, void* arg0, void* arg1); // 1
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

// class sizeof(QConcatenable<QLatin1String>)=1
type QConcatenableLQLatin1StringG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QCharRef>)=1
type QConcatenableLQCharRefG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QStringRef>)=1
type QConcatenableLQStringRefG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAbstractConcatenable)=1
type QAbstractConcatenable struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QString>)=1
type QConcatenableLQStringG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QLatin1Char>)=1
type QConcatenableLQLatin1CharG struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<char>)=1
type QConcatenableLcharG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QByteArray>)=1
type QConcatenableLQByteArrayG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QChar::SpecialCharacter>)=1
type QConcatenableLQChar__SpecialCharacterG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<QChar>)=1
type QConcatenableLQCharG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStringBuilder<QByteArray, QByteArray>)=1
type QStringBuilderLQByteArray_EQByteArrayG struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QConcatenable<const char *>)=1
type QConcatenableLconstEcharEPG struct {
  /*qbase*/ QAbstractConcatenable;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStringBuilder<QString, QString>)=1
type QStringBuilderLQString_EQStringG struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// appendTo(class QCharRef, class QChar *&)
func (this *QConcatenableLQCharRefG) appendTo_s(args ...interface{}) () {
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
    var arg0 = args[0].(QCharRef).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI8QCharRefE8appendToES0_RP5QChar(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QCharRef>", "appendTo", args)
  }

}

// size(class QCharRef)
func (this *QConcatenableLQCharRefG) size_s(args ...interface{}) () {
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
    var arg0 = args[0].(QCharRef).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QConcatenableI8QCharRefE4sizeES0_(arg0)
  default:
    qtrt.ErrorResolve("QConcatenable<QCharRef>", "size", args)
  }

}

// appendTo(const class QString &, class QChar *&)
func (this *QConcatenableLQStringG) appendTo_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI7QStringE8appendToERKS0_RP5QChar(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QString>", "appendTo", args)
  }

}

// size(const class QString &)
func (this *QConcatenableLQStringG) size_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QConcatenableI7QStringE4sizeERKS0_(arg0)
  default:
    qtrt.ErrorResolve("QConcatenable<QString>", "size", args)
  }

}

// appendTo(const char, class QChar *&)
func (this *QConcatenableLcharG) appendTo_s(args ...interface{}) () {
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
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableIcE8appendToEcRP5QChar(arg0, arg1)
  case 1:
    // invoke: _ZN13QConcatenableIcE8appendToEcRPc
    // invoke: void appendTo(const char, char *&)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableIcE8appendToEcRPc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<char>", "appendTo", args)
  }

}

// size(const char)
func (this *QConcatenableLcharG) size_s(args ...interface{}) () {
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
    C.C_ZN13QConcatenableIcE4sizeEc(arg0)
  default:
    qtrt.ErrorResolve("QConcatenable<char>", "size", args)
  }

}

// appendTo(const class QByteArray &, char *&)
func (this *QConcatenableLQByteArrayG) appendTo_s(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RPc(arg0, arg1)
  case 1:
    // invoke: _ZN13QConcatenableI10QByteArrayE8appendToERKS0_RP5QChar
    // invoke: void appendTo(const class QByteArray &, class QChar *&)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI10QByteArrayE8appendToERKS0_RP5QChar(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QByteArray>", "appendTo", args)
  }

}

// size(const class QByteArray &)
func (this *QConcatenableLQByteArrayG) size_s(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QConcatenableI10QByteArrayE4sizeERKS0_(arg0)
  default:
    qtrt.ErrorResolve("QConcatenable<QByteArray>", "size", args)
  }

}

// appendTo(const class QChar, class QChar *&)
func (this *QConcatenableLQCharG) appendTo_s(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableI5QCharE8appendToES0_RPS0_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<QChar>", "appendTo", args)
  }

}

// size(const class QChar)
func (this *QConcatenableLQCharG) size_s(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QConcatenableI5QCharE4sizeES0_(arg0)
  default:
    qtrt.ErrorResolve("QConcatenable<QChar>", "size", args)
  }

}

// QStringBuilder(const class QByteArray &, const class QByteArray &)
func NewQStringBuilderLQByteArray_EQByteArrayG(args ...interface{}) QStringBuilderLQByteArray_EQByteArrayG {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN14QStringBuilderI10QByteArrayS0_EC2ERKS0_S3_(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringBuilder<QByteArray, QByteArray>", "QStringBuilder", args)
  }

  return QStringBuilderLQByteArray_EQByteArrayG{}
}

// appendTo(const char *, class QChar *&)
func (this *QConcatenableLconstEcharEPG) appendTo_s(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableIPKcE8appendToES1_RP5QChar(arg0, arg1)
  case 1:
    // invoke: _ZN13QConcatenableIPKcE8appendToES1_RPc
    // invoke: void appendTo(const char *, char *&)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN13QConcatenableIPKcE8appendToES1_RPc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QConcatenable<const char *>", "appendTo", args)
  }

}

// size(const char *)
func (this *QConcatenableLconstEcharEPG) size_s(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZN13QConcatenableIPKcE4sizeES1_(arg0)
  default:
    qtrt.ErrorResolve("QConcatenable<const char *>", "size", args)
  }

}

// QStringBuilder(const class QString &, const class QString &)
func NewQStringBuilderLQString_EQStringG(args ...interface{}) QStringBuilderLQString_EQStringG {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN14QStringBuilderI7QStringS0_EC2ERKS0_S3_(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringBuilder<QString, QString>", "QStringBuilder", args)
  }

  return QStringBuilderLQString_EQStringG{}
}

// <= body block end

