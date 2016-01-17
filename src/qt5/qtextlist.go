package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK9QTextList8itemTextERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  int QTextList::count();
extern void _ZNK9QTextList5countEv(void* qthis); // 4
  // proto:  void QTextList::setFormat(const QTextListFormat & format);
extern void _ZN9QTextList9setFormatERK15QTextListFormat(void* qthis, void* arg0); // 2
  // proto:  QTextListFormat QTextList::format();
extern void _ZNK9QTextList6formatEv(void* qthis); // 2
  // proto:  void QTextList::remove(const QTextBlock & );
extern void _ZN9QTextList6removeERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  int QTextList::itemNumber(const QTextBlock & );
extern void _ZNK9QTextList10itemNumberERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  QTextBlock QTextList::item(int i);
extern void _ZNK9QTextList4itemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextList::add(const QTextBlock & block);
extern void _ZN9QTextList3addERK10QTextBlock(void* qthis, void* arg0); // 4
  // proto:  bool QTextList::isEmpty();
extern void _ZNK9QTextList7isEmptyEv(void* qthis); // 2
  // proto:  void QTextList::removeItem(int i);
extern void _ZN9QTextList10removeItemEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QTextList::metaObject();
extern void _ZNK9QTextList10metaObjectEv(void* qthis); // 4
  // proto:  void QTextList::~QTextList();
extern void _ZN9QTextListD2Ev(void* qthis); // 4
  // proto:  void QTextList::QTextList(QTextDocument * doc);
extern void _ZN9QTextListC2EP13QTextDocument(void* qthis, void* arg0); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// itemText(const class QTextBlock &)
func (this *QTextList) itemText(args ...interface{}) () {
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
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTextList8itemTextERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "itemText", args)
  }

}

// count()
func (this *QTextList) count(args ...interface{}) () {
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
    C._ZNK9QTextList5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextList", "count", args)
  }

}

// setFormat(const class QTextListFormat &)
func (this *QTextList) setFormat(args ...interface{}) () {
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
    var arg0 = args[0].(QTextListFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextList9setFormatERK15QTextListFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "setFormat", args)
  }

}

// format()
func (this *QTextList) format(args ...interface{}) () {
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
    C._ZNK9QTextList6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextList", "format", args)
  }

}

// remove(const class QTextBlock &)
func (this *QTextList) remove(args ...interface{}) () {
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
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextList6removeERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "remove", args)
  }

}

// itemNumber(const class QTextBlock &)
func (this *QTextList) itemNumber(args ...interface{}) () {
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
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK9QTextList10itemNumberERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "itemNumber", args)
  }

}

// item(int)
func (this *QTextList) item(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QTextList4itemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "item", args)
  }

}

// add(const class QTextBlock &)
func (this *QTextList) add(args ...interface{}) () {
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
    var arg0 = args[0].(QTextBlock).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QTextList3addERK10QTextBlock(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "add", args)
  }

}

// isEmpty()
func (this *QTextList) isEmpty(args ...interface{}) () {
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
    C._ZNK9QTextList7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextList", "isEmpty", args)
  }

}

// removeItem(int)
func (this *QTextList) removeItem(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QTextList10removeItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "removeItem", args)
  }

}

// metaObject()
func (this *QTextList) metaObject(args ...interface{}) () {
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
    C._ZNK9QTextList10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextList", "metaObject", args)
  }

}

// ~QTextList()
func (this *QTextList) FreeQTextList(args ...interface{}) () {
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
    C._ZN9QTextListD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextList", "~QTextList", args)
  }

}

// QTextList(class QTextDocument *)
func NewQTextList(args ...interface{}) QTextList {
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
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN9QTextListC2EP13QTextDocument(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextList", "QTextList", args)
  }

  return QTextList{}
}

// <= body block end

