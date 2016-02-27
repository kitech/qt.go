package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtCore/qstringlist.h
// dst-file: /src/core/qstringlist.go
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
  // proto:  int QStringList::indexOf(const QRegExp & rx, int from);
extern int32_t C_ZNK11QStringList7indexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::indexOf(const QRegularExpression & re, int from);
extern int32_t C_ZNK11QStringList7indexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::indexOf(QRegExp & rx, int from);
extern int32_t C_ZNK11QStringList7indexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  void QStringList::QStringList();
extern void* C_ZN11QStringListC2Ev(); // 1
  // proto:  void QStringList::QStringList(const QString & i);
extern void* C_ZN11QStringListC2ERK7QString(void* arg0); // 1
  // proto:  int QStringList::lastIndexOf(const QRegExp & rx, int from);
extern int32_t C_ZNK11QStringList11lastIndexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::lastIndexOf(const QRegularExpression & re, int from);
extern int32_t C_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QStringList::lastIndexOf(QRegExp & rx, int from);
extern int32_t C_ZNK11QStringList11lastIndexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  QStringList & QListSpecialMethods<QString>::replaceInStrings(const QRegularExpression & re, const QString & after);
extern void C_ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK18QRegularExpressionRKS0_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QStringList & QListSpecialMethods<QString>::replaceInStrings(const QRegExp & rx, const QString & after);
extern void C_ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK7QRegExpRKS0_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QString QListSpecialMethods<QString>::join(const QString & sep);
extern void* C_ZNK19QListSpecialMethodsI7QStringE4joinERKS0_(void* qthis, void* arg0); // 2
  // proto:  QString QListSpecialMethods<QString>::join(QChar sep);
extern void* C_ZNK19QListSpecialMethodsI7QStringE4joinE5QChar(void* qthis, void* arg0); // 2
  // proto:  QStringList QListSpecialMethods<QString>::filter(const QRegularExpression & re);
extern void C_ZNK19QListSpecialMethodsI7QStringE6filterERK18QRegularExpression(void* qthis, void* arg0); // 2
  // proto:  QStringList QListSpecialMethods<QString>::filter(const QRegExp & rx);
extern void C_ZNK19QListSpecialMethodsI7QStringE6filterERK7QRegExp(void* qthis, void* arg0); // 2
  // proto:  int QListSpecialMethods<QString>::removeDuplicates();
extern int32_t C_ZN19QListSpecialMethodsI7QStringE16removeDuplicatesEv(void* qthis); // 2
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

// class sizeof(QStringList)=1
type QStringList struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QListSpecialMethods<QString>)=1
type QListSpecialMethodsLQStringG struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// indexOf(const class QRegExp &, int)
func (this *QStringList) IndexOf(args ...interface{}) (ret interface{}) {
  // indexOf(const class QRegExp &, int)
  // indexOf(const class QRegularExpression &, int)
  // indexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList7indexOfERK7QRegExpi
    // invoke: int indexOf(const class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QStringList7indexOfERK7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK11QStringList7indexOfERK18QRegularExpressioni
    // invoke: int indexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QStringList7indexOfERK18QRegularExpressioni(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK11QStringList7indexOfER7QRegExpi
    // invoke: int indexOf(class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QStringList7indexOfER7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringList", "indexOf", args)
  }

  return
}

// QStringList()
func GcfreeQStringList(this *QStringList) {
  qtrt.UniverseFree(this)
}
func NewQStringList(args ...interface{}) *QStringList {
  // QStringList()
  // QStringList(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QStringListC1Ev
    // invoke: void QStringList()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QStringListC2Ev()
    this := &QStringList{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStringList)
    return this
  case 1:
    // invoke: _ZN11QStringListC1ERK7QString
    // invoke: void QStringList(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QStringListC2ERK7QString(arg0)
    this := &QStringList{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStringList)
    return this
  default:
    qtrt.ErrorResolve("QStringList", "QStringList", args)
  }

  return nil // QStringList{Qclsinst:qthis}
}

// lastIndexOf(const class QRegExp &, int)
func (this *QStringList) LastIndexOf(args ...interface{}) (ret interface{}) {
  // lastIndexOf(const class QRegExp &, int)
  // lastIndexOf(const class QRegularExpression &, int)
  // lastIndexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QStringList11lastIndexOfERK7QRegExpi
    // invoke: int lastIndexOf(const class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QStringList11lastIndexOfERK7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK11QStringList11lastIndexOfERK18QRegularExpressioni
    // invoke: int lastIndexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QStringList11lastIndexOfERK18QRegularExpressioni(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK11QStringList11lastIndexOfER7QRegExpi
    // invoke: int lastIndexOf(class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QStringList11lastIndexOfER7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringList", "lastIndexOf", args)
  }

  return
}

// replaceInStrings(const class QRegularExpression &, const class QString &)
func (this *QListSpecialMethodsLQStringG) ReplaceInStrings(args ...interface{}) () {
  // replaceInStrings(const class QRegularExpression &, const class QString &)
  // replaceInStrings(const class QRegExp &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK18QRegularExpressionRKS0_
    // invoke: QStringList & replaceInStrings(const class QRegularExpression &, const class QString &)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK18QRegularExpressionRKS0_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK7QRegExpRKS0_
    // invoke: QStringList & replaceInStrings(const class QRegExp &, const class QString &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN19QListSpecialMethodsI7QStringE16replaceInStringsERK7QRegExpRKS0_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "replaceInStrings", args)
  }

  return
}

// join(const class QString &)
func (this *QListSpecialMethodsLQStringG) Join(args ...interface{}) (ret interface{}) {
  // join(const class QString &)
  // join(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE4joinERKS0_
    // invoke: QString join(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QListSpecialMethodsI7QStringE4joinERKS0_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE4joinE5QChar
    // invoke: QString join(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK19QListSpecialMethodsI7QStringE4joinE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "join", args)
  }

  return
}

// filter(const class QRegularExpression &)
func (this *QListSpecialMethodsLQStringG) Filter(args ...interface{}) () {
  // filter(const class QRegularExpression &)
  // filter(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE6filterERK18QRegularExpression
    // invoke: QStringList filter(const class QRegularExpression &)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK19QListSpecialMethodsI7QStringE6filterERK18QRegularExpression(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZNK19QListSpecialMethodsI7QStringE6filterERK7QRegExp
    // invoke: QStringList filter(const class QRegExp &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK19QListSpecialMethodsI7QStringE6filterERK7QRegExp(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "filter", args)
  }

  return
}

// removeDuplicates()
func (this *QListSpecialMethodsLQStringG) RemoveDuplicates(args ...interface{}) (ret interface{}) {
  // removeDuplicates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QListSpecialMethodsI7QStringE16removeDuplicatesEv
    // invoke: int removeDuplicates()
    var ret0 = C.C_ZN19QListSpecialMethodsI7QStringE16removeDuplicatesEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListSpecialMethods<QString>", "removeDuplicates", args)
  }

  return
}

// <= body block end

