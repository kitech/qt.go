package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern void* C_ZNK9QMimeType7commentEv(void* qthis); // 4
  // proto:  bool QMimeType::inherits(const QString & mimeTypeName);
extern bool C_ZNK9QMimeType8inheritsERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QMimeType::parentMimeTypes();
extern void C_ZNK9QMimeType15parentMimeTypesEv(void* qthis); // 4
  // proto:  QString QMimeType::filterString();
extern void* C_ZNK9QMimeType12filterStringEv(void* qthis); // 4
  // proto:  bool QMimeType::isValid();
extern bool C_ZNK9QMimeType7isValidEv(void* qthis); // 4
  // proto:  QString QMimeType::preferredSuffix();
extern void* C_ZNK9QMimeType15preferredSuffixEv(void* qthis); // 4
  // proto:  QStringList QMimeType::globPatterns();
extern void C_ZNK9QMimeType12globPatternsEv(void* qthis); // 4
  // proto:  QString QMimeType::genericIconName();
extern void* C_ZNK9QMimeType15genericIconNameEv(void* qthis); // 4
  // proto:  QString QMimeType::iconName();
extern void* C_ZNK9QMimeType8iconNameEv(void* qthis); // 4
  // proto:  void QMimeType::swap(QMimeType & other);
extern void C_ZN9QMimeType4swapERS_(void* qthis, void* arg0); // 2
  // proto:  void QMimeType::QMimeType(const QMimeType & other);
extern void* C_ZN9QMimeTypeC2ERKS_(void* arg0); // 3
  // proto:  void QMimeType::QMimeType();
extern void* C_ZN9QMimeTypeC2Ev(); // 3
  // proto:  void QMimeType::~QMimeType();
extern void C_ZN9QMimeTypeD2Ev(void* qthis); // 4
  // proto:  QStringList QMimeType::allAncestors();
extern void C_ZNK9QMimeType12allAncestorsEv(void* qthis); // 4
  // proto:  QStringList QMimeType::suffixes();
extern void C_ZNK9QMimeType8suffixesEv(void* qthis); // 4
  // proto:  QString QMimeType::name();
extern void* C_ZNK9QMimeType4nameEv(void* qthis); // 4
  // proto:  bool QMimeType::isDefault();
extern bool C_ZNK9QMimeType9isDefaultEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// comment()
func (this *QMimeType) Comment(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType7commentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "comment", args)
  }

  return
}

// inherits(const class QString &)
func (this *QMimeType) Inherits(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMimeType8inheritsERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "inherits", args)
  }

  return
}

// parentMimeTypes()
func (this *QMimeType) Parentmimetypes(args ...interface{}) () {
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
    C.C_ZNK9QMimeType15parentMimeTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "parentMimeTypes", args)
  }

  return
}

// filterString()
func (this *QMimeType) Filterstring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType12filterStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "filterString", args)
  }

  return
}

// isValid()
func (this *QMimeType) Isvalid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "isValid", args)
  }

  return
}

// preferredSuffix()
func (this *QMimeType) Preferredsuffix(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType15preferredSuffixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "preferredSuffix", args)
  }

  return
}

// globPatterns()
func (this *QMimeType) Globpatterns(args ...interface{}) () {
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
    C.C_ZNK9QMimeType12globPatternsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "globPatterns", args)
  }

  return
}

// genericIconName()
func (this *QMimeType) Genericiconname(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType15genericIconNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "genericIconName", args)
  }

  return
}

// iconName()
func (this *QMimeType) Iconname(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType8iconNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "iconName", args)
  }

  return
}

// swap(class QMimeType &)
func (this *QMimeType) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QMimeType).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeType4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeType", "swap", args)
  }

  return
}

// QMimeType(const class QMimeType &)
func NewQMimeType(args ...interface{}) *QMimeType {
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
    var arg0 = args[0].(*QMimeType).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMimeTypeC2ERKS_(arg0)
    return &QMimeType{Qclsinst:qthis}
  case 1:
    // invoke: _ZN9QMimeTypeC1Ev
    // invoke: void QMimeType()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMimeTypeC2Ev()
    return &QMimeType{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMimeType", "QMimeType", args)
  }

  return nil // QMimeType{Qclsinst:qthis}
}

// ~QMimeType()
func (this *QMimeType) Freeqmimetype(args ...interface{}) () {
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
    C.C_ZN9QMimeTypeD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "~QMimeType", args)
  }

  return
}

// allAncestors()
func (this *QMimeType) Allancestors(args ...interface{}) () {
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
    C.C_ZNK9QMimeType12allAncestorsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "allAncestors", args)
  }

  return
}

// suffixes()
func (this *QMimeType) Suffixes(args ...interface{}) () {
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
    C.C_ZNK9QMimeType8suffixesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "suffixes", args)
  }

  return
}

// name()
func (this *QMimeType) Name(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType4nameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "name", args)
  }

  return
}

// isDefault()
func (this *QMimeType) Isdefault(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeType9isDefaultEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeType", "isDefault", args)
  }

  return
}

// aliases()
func (this *QMimeType) Aliases(args ...interface{}) () {
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
    C.C_ZNK9QMimeType7aliasesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeType", "aliases", args)
  }

  return
}

// <= body block end

