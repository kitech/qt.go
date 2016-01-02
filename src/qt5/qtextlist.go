package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QTextBlock QTextList::item(int i);
extern void _ZNK9QTextList4itemEi(void* qthis, int arg0);
  // proto:  void QTextList::remove(const QTextBlock & );
extern void _ZN9QTextList6removeERK10QTextBlock(void* qthis, void* arg0);
  // proto:  void QTextList::setFormat(const QTextListFormat & format);
extern void demth_ZN9QTextList9setFormatERK15QTextListFormat(void* qthis, void* arg0);
  // proto:  void QTextList::QTextList(QTextDocument * doc);
extern void* dector_ZN9QTextListC1EP13QTextDocument(void* arg0);
extern void _ZN9QTextListC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  void QTextList::QTextList(const QTextList & );
extern void* dector_ZN9QTextListC1ERKS_(void* arg0);
extern void _ZN9QTextListC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextList::add(const QTextBlock & block);
extern void _ZN9QTextList3addERK10QTextBlock(void* qthis, void* arg0);
  // proto:  QString QTextList::itemText(const QTextBlock & );
extern void _ZNK9QTextList8itemTextERK10QTextBlock(void* qthis, void* arg0);
  // proto:  void QTextList::removeItem(int i);
extern void _ZN9QTextList10removeItemEi(void* qthis, int arg0);
  // proto:  int QTextList::itemNumber(const QTextBlock & );
extern void _ZNK9QTextList10itemNumberERK10QTextBlock(void* qthis, void* arg0);
  // proto:  int QTextList::count();
extern void _ZNK9QTextList5countEv(void* qthis);
  // proto:  QTextListFormat QTextList::format();
extern void _ZNK9QTextList6formatEv(void* qthis);
  // proto:  void QTextList::~QTextList();
extern void _ZN9QTextListD0Ev(void* qthis);
  // proto:  bool QTextList::isEmpty();
extern void demth_ZNK9QTextList7isEmptyEv(void* qthis);
  // proto:  const QMetaObject * QTextList::metaObject();
extern void _ZNK9QTextList10metaObjectEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QTextBlock QTextList::item(int i);
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
  default:
    qtrt.ErrorResolve("QTextList", "item", args)
  }

}

  // proto:  void QTextList::remove(const QTextBlock & );
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
  default:
    qtrt.ErrorResolve("QTextList", "remove", args)
  }

}

  // proto:  void QTextList::setFormat(const QTextListFormat & format);
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
  default:
    qtrt.ErrorResolve("QTextList", "setFormat", args)
  }

}

  // proto:  void QTextList::QTextList(QTextDocument * doc);
func NewQTextList(args ...interface{}) QTextList {
  return QTextList{}
}

  // proto:  void QTextList::add(const QTextBlock & block);
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
  default:
    qtrt.ErrorResolve("QTextList", "add", args)
  }

}

  // proto:  QString QTextList::itemText(const QTextBlock & );
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
  default:
    qtrt.ErrorResolve("QTextList", "itemText", args)
  }

}

  // proto:  void QTextList::removeItem(int i);
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
  default:
    qtrt.ErrorResolve("QTextList", "removeItem", args)
  }

}

  // proto:  int QTextList::itemNumber(const QTextBlock & );
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
  default:
    qtrt.ErrorResolve("QTextList", "itemNumber", args)
  }

}

  // proto:  int QTextList::count();
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
  default:
    qtrt.ErrorResolve("QTextList", "count", args)
  }

}

  // proto:  QTextListFormat QTextList::format();
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
  default:
    qtrt.ErrorResolve("QTextList", "format", args)
  }

}

  // proto:  void QTextList::~QTextList();
func (this *QTextList) FreeQTextList(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextList", "~QTextList", args)
  }

}

  // proto:  bool QTextList::isEmpty();
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
  default:
    qtrt.ErrorResolve("QTextList", "isEmpty", args)
  }

}

  // proto:  const QMetaObject * QTextList::metaObject();
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
  default:
    qtrt.ErrorResolve("QTextList", "metaObject", args)
  }

}

// <= body block end

