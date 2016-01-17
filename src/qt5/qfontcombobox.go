package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qfontcombobox.h
// dst-file: /src/widgets/qfontcombobox.go
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
  // proto:  FontFilters QFontComboBox::fontFilters();
extern void _ZNK13QFontComboBox11fontFiltersEv(void* qthis); // 4
  // proto:  void QFontComboBox::setCurrentFont(const QFont & f);
extern void _ZN13QFontComboBox14setCurrentFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  void QFontComboBox::QFontComboBox(QWidget * parent);
extern void _ZN13QFontComboBoxC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  QSize QFontComboBox::sizeHint();
extern void _ZNK13QFontComboBox8sizeHintEv(void* qthis); // 4
  // proto:  const QMetaObject * QFontComboBox::metaObject();
extern void _ZNK13QFontComboBox10metaObjectEv(void* qthis); // 4
  // proto:  QFont QFontComboBox::currentFont();
extern void _ZNK13QFontComboBox11currentFontEv(void* qthis); // 4
  // proto:  void QFontComboBox::~QFontComboBox();
extern void _ZN13QFontComboBoxD2Ev(void* qthis); // 4
  // proto:  QFontDatabase::WritingSystem QFontComboBox::writingSystem();
extern void _ZNK13QFontComboBox13writingSystemEv(void* qthis); // 4
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

// class sizeof(QFontComboBox)=1
type QFontComboBox struct {
  /*qbase*/ QComboBox;
  qclsinst unsafe.Pointer /* *C.void */;
//  _currentFontChanged QFontComboBox_currentFontChanged_signal;
}

// fontFilters()
func (this *QFontComboBox) fontFilters(args ...interface{}) () {
  // fontFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox11fontFiltersEv
    // invoke: FontFilters fontFilters()
    C._ZNK13QFontComboBox11fontFiltersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "fontFilters", args)
  }

}

// setCurrentFont(const class QFont &)
func (this *QFontComboBox) setCurrentFont(args ...interface{}) () {
  // setCurrentFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontComboBox14setCurrentFontERK5QFont
    // invoke: void setCurrentFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QFontComboBox14setCurrentFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontComboBox", "setCurrentFont", args)
  }

}

// QFontComboBox(class QWidget *)
func NewQFontComboBox(args ...interface{}) QFontComboBox {
  // QFontComboBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontComboBoxC1EP7QWidget
    // invoke: void QFontComboBox(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QFontComboBoxC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QFontComboBox", "QFontComboBox", args)
  }

  return QFontComboBox{}
}

// sizeHint()
func (this *QFontComboBox) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK13QFontComboBox8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "sizeHint", args)
  }

}

// metaObject()
func (this *QFontComboBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QFontComboBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "metaObject", args)
  }

}

// currentFont()
func (this *QFontComboBox) currentFont(args ...interface{}) () {
  // currentFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox11currentFontEv
    // invoke: QFont currentFont()
    C._ZNK13QFontComboBox11currentFontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "currentFont", args)
  }

}

// ~QFontComboBox()
func (this *QFontComboBox) FreeQFontComboBox(args ...interface{}) () {
  // ~QFontComboBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontComboBoxD0Ev
    // invoke: void ~QFontComboBox()
    C._ZN13QFontComboBoxD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "~QFontComboBox", args)
  }

}

// writingSystem()
func (this *QFontComboBox) writingSystem(args ...interface{}) () {
  // writingSystem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontComboBox13writingSystemEv
    // invoke: QFontDatabase::WritingSystem writingSystem()
    C._ZNK13QFontComboBox13writingSystemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontComboBox", "writingSystem", args)
  }

}

// <= body block end

