package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QFontComboBox::QFontComboBox(const QFontComboBox & );
extern void* dector_ZN13QFontComboBoxC1ERKS_(void* arg0);
extern void _ZN13QFontComboBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QFontComboBox::~QFontComboBox();
extern void _ZN13QFontComboBoxD0Ev(void* qthis);
  // proto:  const QMetaObject * QFontComboBox::metaObject();
extern void _ZNK13QFontComboBox10metaObjectEv(void* qthis);
  // proto:  void QFontComboBox::QFontComboBox(QWidget * parent);
extern void* dector_ZN13QFontComboBoxC1EP7QWidget(void* arg0);
extern void _ZN13QFontComboBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QSize QFontComboBox::sizeHint();
extern void _ZNK13QFontComboBox8sizeHintEv(void* qthis);
  // proto:  QFont QFontComboBox::currentFont();
extern void _ZNK13QFontComboBox11currentFontEv(void* qthis);
  // proto:  void QFontComboBox::setCurrentFont(const QFont & f);
extern void _ZN13QFontComboBox14setCurrentFontERK5QFont(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _currentFontChanged QFontComboBox_currentFontChanged_signal;
}

  // proto:  void QFontComboBox::QFontComboBox(const QFontComboBox & );
func NewQFontComboBox(args ...interface{}) QFontComboBox {
  return QFontComboBox{}
}

  // proto:  void QFontComboBox::~QFontComboBox();
func (this *QFontComboBox) FreeQFontComboBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontComboBox", "~QFontComboBox", args)
  }

}

  // proto:  const QMetaObject * QFontComboBox::metaObject();
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

  // proto:  QSize QFontComboBox::sizeHint();
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

  // proto:  QFont QFontComboBox::currentFont();
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

  // proto:  void QFontComboBox::setCurrentFont(const QFont & f);
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

// <= body block end

