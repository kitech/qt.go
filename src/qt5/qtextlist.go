package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtGui/qtextlist.h
// dst-file: /src/gui/qtextlist.go
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
  // proto:  QString QTextList::itemText(const QTextBlock & );
extern void* C_ZNK9QTextList8itemTextERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  int QTextList::count();
extern int32_t C_ZNK9QTextList5countEv(void* qthis); // 4
  // proto:  void QTextList::setFormat(const QTextListFormat & format);
extern void C_ZN9QTextList9setFormatERK15QTextListFormat(void* qthis, void* arg0); // 2
  // proto:  QTextListFormat QTextList::format();
extern void* C_ZNK9QTextList6formatEv(void* qthis); // 2
  // proto:  void QTextList::remove(const QTextBlock & );
extern void C_ZN9QTextList6removeERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  int QTextList::itemNumber(const QTextBlock & );
extern int32_t C_ZNK9QTextList10itemNumberERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  QTextBlock QTextList::item(int i);
extern void* C_ZNK9QTextList4itemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextList::add(const QTextBlock & block);
extern void C_ZN9QTextList3addERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  bool QTextList::isEmpty();
extern bool C_ZNK9QTextList7isEmptyEv(void* qthis); // 2
  // proto:  void QTextList::removeItem(int i);
extern void C_ZN9QTextList10removeItemEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QTextList::metaObject();
extern void C_ZNK9QTextList10metaObjectEv(void* qthis); // 4
  // proto:  void QTextList::~QTextList();
extern void C_ZN9QTextListD2Ev(void* qthis); // 4
  // proto:  void QTextList::QTextList(QTextDocument * doc);
extern void* C_ZN9QTextListC2EP13QTextDocument(void* arg0); // 3
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

// class sizeof(QTextList)=1
type QTextList struct {
  /*qbase*/ QTextBlockGroup;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// itemText(const class QTextBlock &)
func (this *QTextList) Itemtext(args ...interface{}) (ret interface{}) {
  // itemText(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList8itemTextERK10QTextBlock
    // invoke: QString itemText(const class QTextBlock &)
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTextList8itemTextERK10QTextBlock(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextList", "itemText", args)
  }

  return
}

// count()
func (this *QTextList) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK9QTextList5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextList", "count", args)
  }

  return
}

// setFormat(const class QTextListFormat &)
func (this *QTextList) Setformat(args ...interface{}) () {
  // setFormat(const class QTextListFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextListFormat{}) // "const QTextListFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList9setFormatERK15QTextListFormat
    // invoke: void setFormat(const class QTextListFormat &)
    var arg0 = args[0].(*QTextListFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextList9setFormatERK15QTextListFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "setFormat", args)
  }

  return
}

// format()
func (this *QTextList) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList6formatEv
    // invoke: QTextListFormat format()
    var ret0 = C.C_ZNK9QTextList6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextListFormat{}) // "QTextListFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextList", "format", args)
  }

  return
}

// remove(const class QTextBlock &)
func (this *QTextList) Remove(args ...interface{}) () {
  // remove(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList6removeERK10QTextBlock
    // invoke: void remove(const class QTextBlock &)
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextList6removeERK10QTextBlock(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "remove", args)
  }

  return
}

// itemNumber(const class QTextBlock &)
func (this *QTextList) Itemnumber(args ...interface{}) (ret interface{}) {
  // itemNumber(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList10itemNumberERK10QTextBlock
    // invoke: int itemNumber(const class QTextBlock &)
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTextList10itemNumberERK10QTextBlock(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextList", "itemNumber", args)
  }

  return
}

// item(int)
func (this *QTextList) Item(args ...interface{}) (ret interface{}) {
  // item(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList4itemEi
    // invoke: QTextBlock item(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QTextList4itemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextBlock{}) // "QTextBlock"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextList", "item", args)
  }

  return
}

// add(const class QTextBlock &)
func (this *QTextList) Add(args ...interface{}) () {
  // add(const class QTextBlock &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextBlock{}) // "const QTextBlock &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList3addERK10QTextBlock
    // invoke: void add(const class QTextBlock &)
    var arg0 = args[0].(*QTextBlock).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextList3addERK10QTextBlock(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "add", args)
  }

  return
}

// isEmpty()
func (this *QTextList) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK9QTextList7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextList", "isEmpty", args)
  }

  return
}

// removeItem(int)
func (this *QTextList) Removeitem(args ...interface{}) () {
  // removeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextList10removeItemEi
    // invoke: void removeItem(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN9QTextList10removeItemEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "removeItem", args)
  }

  return
}

// metaObject()
func (this *QTextList) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QTextList10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QTextList10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextList", "metaObject", args)
  }

  return
}

// ~QTextList()
func (this *QTextList) Freeqtextlist(args ...interface{}) () {
  // ~QTextList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextListD0Ev
    // invoke: void ~QTextList()
    C.C_ZN9QTextListD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextList", "~QTextList", args)
  }

  return
}

// QTextList(class QTextDocument *)
func NewQTextList(args ...interface{}) *QTextList {
  // QTextList(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QTextListC1EP13QTextDocument
    // invoke: void QTextList(class QTextDocument *)
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QTextListC2EP13QTextDocument(arg0)
    return &QTextList{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextList", "QTextList", args)
  }

  return nil // QTextList{Qclsinst:qthis}
}

// <= body block end

