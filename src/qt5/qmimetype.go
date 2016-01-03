package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QMimeType::~QMimeType();
extern void _ZN9QMimeTypeD0Ev(void* qthis);
  // proto:  QString QMimeType::comment();
extern void _ZNK9QMimeType7commentEv(void* qthis);
  // proto:  QStringList QMimeType::aliases();
extern void _ZNK9QMimeType7aliasesEv(void* qthis);
  // proto:  QString QMimeType::filterString();
extern void _ZNK9QMimeType12filterStringEv(void* qthis);
  // proto:  QStringList QMimeType::parentMimeTypes();
extern void _ZNK9QMimeType15parentMimeTypesEv(void* qthis);
  // proto:  void QMimeType::QMimeType(const QMimeType & other);
extern void* dector_ZN9QMimeTypeC1ERKS_(void* arg0);
extern void _ZN9QMimeTypeC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QMimeType::inherits(const QString & mimeTypeName);
extern void _ZNK9QMimeType8inheritsERK7QString(void* qthis, void* arg0);
  // proto:  bool QMimeType::isDefault();
extern void _ZNK9QMimeType9isDefaultEv(void* qthis);
  // proto:  bool QMimeType::isValid();
extern void _ZNK9QMimeType7isValidEv(void* qthis);
  // proto:  void QMimeType::QMimeType();
extern void* dector_ZN9QMimeTypeC1Ev();
extern void _ZN9QMimeTypeC1Ev(void* qthis);
  // proto:  void QMimeType::swap(QMimeType & other);
extern void demth_ZN9QMimeType4swapERS_(void* qthis, void* arg0);
  // proto:  QStringList QMimeType::suffixes();
extern void _ZNK9QMimeType8suffixesEv(void* qthis);
  // proto:  QString QMimeType::genericIconName();
extern void _ZNK9QMimeType15genericIconNameEv(void* qthis);
  // proto:  QString QMimeType::iconName();
extern void _ZNK9QMimeType8iconNameEv(void* qthis);
  // proto:  QStringList QMimeType::allAncestors();
extern void _ZNK9QMimeType12allAncestorsEv(void* qthis);
  // proto:  QStringList QMimeType::globPatterns();
extern void _ZNK9QMimeType12globPatternsEv(void* qthis);
  // proto:  QString QMimeType::name();
extern void _ZNK9QMimeType4nameEv(void* qthis);
  // proto:  QString QMimeType::preferredSuffix();
extern void _ZNK9QMimeType15preferredSuffixEv(void* qthis);
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

  // proto:  void QMimeType::~QMimeType();
func (this *QMimeType) FreeQMimeType(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMimeType", "~QMimeType", args)
  }

}

  // proto:  QString QMimeType::comment();
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
    C._ZNK9QMimeType7commentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "comment", args)
  }

}

  // proto:  QStringList QMimeType::aliases();
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
    C._ZNK9QMimeType7aliasesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "aliases", args)
  }

}

  // proto:  QString QMimeType::filterString();
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
    C._ZNK9QMimeType12filterStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "filterString", args)
  }

}

  // proto:  QStringList QMimeType::parentMimeTypes();
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
    C._ZNK9QMimeType15parentMimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "parentMimeTypes", args)
  }

}

  // proto:  void QMimeType::QMimeType(const QMimeType & other);
func NewQMimeType(args ...interface{}) QMimeType {
  return QMimeType{}
}

  // proto:  bool QMimeType::inherits(const QString & mimeTypeName);
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
    C._ZNK9QMimeType8inheritsERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeType", "inherits", args)
  }

}

  // proto:  bool QMimeType::isDefault();
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
    C._ZNK9QMimeType9isDefaultEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "isDefault", args)
  }

}

  // proto:  bool QMimeType::isValid();
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
    C._ZNK9QMimeType7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "isValid", args)
  }

}

  // proto:  void QMimeType::swap(QMimeType & other);
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
    C.demth_ZN9QMimeType4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeType", "swap", args)
  }

}

  // proto:  QStringList QMimeType::suffixes();
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
    C._ZNK9QMimeType8suffixesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "suffixes", args)
  }

}

  // proto:  QString QMimeType::genericIconName();
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
    C._ZNK9QMimeType15genericIconNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "genericIconName", args)
  }

}

  // proto:  QString QMimeType::iconName();
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
    C._ZNK9QMimeType8iconNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "iconName", args)
  }

}

  // proto:  QStringList QMimeType::allAncestors();
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
    C._ZNK9QMimeType12allAncestorsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "allAncestors", args)
  }

}

  // proto:  QStringList QMimeType::globPatterns();
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
    C._ZNK9QMimeType12globPatternsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "globPatterns", args)
  }

}

  // proto:  QString QMimeType::name();
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
    C._ZNK9QMimeType4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "name", args)
  }

}

  // proto:  QString QMimeType::preferredSuffix();
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
    C._ZNK9QMimeType15preferredSuffixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "preferredSuffix", args)
  }

}

// <= body block end

