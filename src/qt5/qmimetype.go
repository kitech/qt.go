package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qmimetype.h
// dst-file: /src/core/qmimetype.go
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
  // proto:  QString QMimeType::comment();
extern void C_ZNK9QMimeType7commentEv(void* qthis); // 4
  // proto:  bool QMimeType::inherits(const QString & mimeTypeName);
extern void C_ZNK9QMimeType8inheritsERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QMimeType::parentMimeTypes();
extern void C_ZNK9QMimeType15parentMimeTypesEv(void* qthis); // 4
  // proto:  QString QMimeType::filterString();
extern void C_ZNK9QMimeType12filterStringEv(void* qthis); // 4
  // proto:  bool QMimeType::isValid();
extern void C_ZNK9QMimeType7isValidEv(void* qthis); // 4
  // proto:  QString QMimeType::preferredSuffix();
extern void C_ZNK9QMimeType15preferredSuffixEv(void* qthis); // 4
  // proto:  QStringList QMimeType::globPatterns();
extern void C_ZNK9QMimeType12globPatternsEv(void* qthis); // 4
  // proto:  QString QMimeType::genericIconName();
extern void C_ZNK9QMimeType15genericIconNameEv(void* qthis); // 4
  // proto:  QString QMimeType::iconName();
extern void C_ZNK9QMimeType8iconNameEv(void* qthis); // 4
  // proto:  void QMimeType::swap(QMimeType & other);
extern void C_ZN9QMimeType4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QMimeType::QMimeType(const QMimeType & other);
extern void C_ZN9QMimeTypeC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QMimeType::QMimeType();
extern void C_ZN9QMimeTypeC2Ev(void* qthis); // 3
  // proto:  void QMimeType::~QMimeType();
extern void C_ZN9QMimeTypeD2Ev(void* qthis); // 4
  // proto:  QStringList QMimeType::allAncestors();
extern void C_ZNK9QMimeType12allAncestorsEv(void* qthis); // 4
  // proto:  QStringList QMimeType::suffixes();
extern void C_ZNK9QMimeType8suffixesEv(void* qthis); // 4
  // proto:  QString QMimeType::name();
extern void C_ZNK9QMimeType4nameEv(void* qthis); // 4
  // proto:  bool QMimeType::isDefault();
extern void C_ZNK9QMimeType9isDefaultEv(void* qthis); // 4
  // proto:  QStringList QMimeType::aliases();
extern void C_ZNK9QMimeType7aliasesEv(void* qthis); // 4
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

// class sizeof(QMimeType)=1
type QMimeType struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// comment()
func (this *QMimeType) comment(args ...interface{}) () {
  // comment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType7commentEv
    // invoke: QString comment()
    C.C_ZNK9QMimeType7commentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "comment", args)
  }

}

// inherits(const class QString &)
func (this *QMimeType) inherits(args ...interface{}) () {
  // inherits(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType8inheritsERK7QString
    // invoke: bool inherits(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QMimeType8inheritsERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeType", "inherits", args)
  }

}

// parentMimeTypes()
func (this *QMimeType) parentMimeTypes(args ...interface{}) () {
  // parentMimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType15parentMimeTypesEv
    // invoke: QStringList parentMimeTypes()
    C.C_ZNK9QMimeType15parentMimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "parentMimeTypes", args)
  }

}

// filterString()
func (this *QMimeType) filterString(args ...interface{}) () {
  // filterString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType12filterStringEv
    // invoke: QString filterString()
    C.C_ZNK9QMimeType12filterStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "filterString", args)
  }

}

// isValid()
func (this *QMimeType) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType7isValidEv
    // invoke: bool isValid()
    C.C_ZNK9QMimeType7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "isValid", args)
  }

}

// preferredSuffix()
func (this *QMimeType) preferredSuffix(args ...interface{}) () {
  // preferredSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType15preferredSuffixEv
    // invoke: QString preferredSuffix()
    C.C_ZNK9QMimeType15preferredSuffixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "preferredSuffix", args)
  }

}

// globPatterns()
func (this *QMimeType) globPatterns(args ...interface{}) () {
  // globPatterns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType12globPatternsEv
    // invoke: QStringList globPatterns()
    C.C_ZNK9QMimeType12globPatternsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "globPatterns", args)
  }

}

// genericIconName()
func (this *QMimeType) genericIconName(args ...interface{}) () {
  // genericIconName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType15genericIconNameEv
    // invoke: QString genericIconName()
    C.C_ZNK9QMimeType15genericIconNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "genericIconName", args)
  }

}

// iconName()
func (this *QMimeType) iconName(args ...interface{}) () {
  // iconName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType8iconNameEv
    // invoke: QString iconName()
    C.C_ZNK9QMimeType8iconNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "iconName", args)
  }

}

// swap(class QMimeType &)
func (this *QMimeType) swap(args ...interface{}) () {
  // swap(class QMimeType &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMimeType{}) // "QMimeType &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeType4swapERS_
    // invoke: void swap(class QMimeType &)
    var arg0 = args[0].(QMimeType).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeType4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeType", "swap", args)
  }

}

// QMimeType(const class QMimeType &)
func NewQMimeType(args ...interface{}) QMimeType {
  // QMimeType(const class QMimeType &)
  // QMimeType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMimeType{}) // "const QMimeType &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeTypeC1ERKS_
    // invoke: void QMimeType(const class QMimeType &)
    var arg0 = args[0].(QMimeType).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QMimeTypeC2ERKS_(qthis, arg0)
  case 1:
    // invoke: _ZN9QMimeTypeC1Ev
    // invoke: void QMimeType()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN9QMimeTypeC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QMimeType", "QMimeType", args)
  }

  return QMimeType{}
}

// ~QMimeType()
func (this *QMimeType) FreeQMimeType(args ...interface{}) () {
  // ~QMimeType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeTypeD0Ev
    // invoke: void ~QMimeType()
    C.C_ZN9QMimeTypeD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "~QMimeType", args)
  }

}

// allAncestors()
func (this *QMimeType) allAncestors(args ...interface{}) () {
  // allAncestors()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType12allAncestorsEv
    // invoke: QStringList allAncestors()
    C.C_ZNK9QMimeType12allAncestorsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "allAncestors", args)
  }

}

// suffixes()
func (this *QMimeType) suffixes(args ...interface{}) () {
  // suffixes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType8suffixesEv
    // invoke: QStringList suffixes()
    C.C_ZNK9QMimeType8suffixesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "suffixes", args)
  }

}

// name()
func (this *QMimeType) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType4nameEv
    // invoke: QString name()
    C.C_ZNK9QMimeType4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "name", args)
  }

}

// isDefault()
func (this *QMimeType) isDefault(args ...interface{}) () {
  // isDefault()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType9isDefaultEv
    // invoke: bool isDefault()
    C.C_ZNK9QMimeType9isDefaultEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "isDefault", args)
  }

}

// aliases()
func (this *QMimeType) aliases(args ...interface{}) () {
  // aliases()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeType7aliasesEv
    // invoke: QStringList aliases()
    C.C_ZNK9QMimeType7aliasesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "aliases", args)
  }

}

// <= body block end

