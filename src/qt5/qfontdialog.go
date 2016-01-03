package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtWidgets/qfontdialog.h
// dst-file: /src/widgets/qfontdialog.go
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
  // proto:  void QFontDialog::QFontDialog(QWidget * parent);
extern void* dector_ZN11QFontDialogC1EP7QWidget(void* arg0);
extern void _ZN11QFontDialogC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QFontDialog::QFontDialog(const QFontDialog & );
extern void* dector_ZN11QFontDialogC1ERKS_(void* arg0);
extern void _ZN11QFontDialogC1ERKS_(void* qthis, void* arg0);
  // proto:  void QFontDialog::open(QObject * receiver, const char * member);
extern void _ZN11QFontDialog4openEP7QObjectPKc(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  void QFontDialog::QFontDialog(const QFont & initial, QWidget * parent);
extern void* dector_ZN11QFontDialogC1ERK5QFontP7QWidget(void* arg0, void* arg1);
extern void _ZN11QFontDialogC1ERK5QFontP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  QFont QFontDialog::currentFont();
extern void _ZNK11QFontDialog11currentFontEv(void* qthis);
  // proto:  void QFontDialog::setVisible(bool visible);
extern void _ZN11QFontDialog10setVisibleEb(void* qthis, bool arg0);
  // proto:  void QFontDialog::~QFontDialog();
extern void _ZN11QFontDialogD0Ev(void* qthis);
  // proto: static QFont QFontDialog::getFont(bool * ok, QWidget * parent);
extern void _ZN11QFontDialog7getFontEPbP7QWidget(bool* arg0, void* arg1);
  // proto:  const QMetaObject * QFontDialog::metaObject();
extern void _ZNK11QFontDialog10metaObjectEv(void* qthis);
  // proto:  QFont QFontDialog::selectedFont();
extern void _ZNK11QFontDialog12selectedFontEv(void* qthis);
  // proto:  void QFontDialog::setCurrentFont(const QFont & font);
extern void _ZN11QFontDialog14setCurrentFontERK5QFont(void* qthis, void* arg0);
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

// class sizeof(QFontDialog)=1
type QFontDialog struct {
  /*qbase*/ QDialog;
  qclsinst unsafe.Pointer /* *C.void */;
//  _fontSelected QFontDialog_fontSelected_signal;
//  _currentFontChanged QFontDialog_currentFontChanged_signal;
}

  // proto:  void QFontDialog::QFontDialog(QWidget * parent);
func NewQFontDialog(args ...interface{}) QFontDialog {
  return QFontDialog{}
}

  // proto:  void QFontDialog::open(QObject * receiver, const char * member);
func (this *QFontDialog) open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialog4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C._ZN11QFontDialog4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDialog", "open", args)
  }

}

  // proto:  QFont QFontDialog::currentFont();
func (this *QFontDialog) currentFont(args ...interface{}) () {
  // currentFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFontDialog11currentFontEv
    // invoke: QFont currentFont()
    C._ZNK11QFontDialog11currentFontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontDialog", "currentFont", args)
  }

}

  // proto:  void QFontDialog::setVisible(bool visible);
func (this *QFontDialog) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialog10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QFontDialog10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDialog", "setVisible", args)
  }

}

  // proto:  void QFontDialog::~QFontDialog();
func (this *QFontDialog) FreeQFontDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDialog", "~QFontDialog", args)
  }

}

  // proto: static QFont QFontDialog::getFont(bool * ok, QWidget * parent);
func (this *QFontDialog) getFont_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontDialog", "getFont", args)
  }

}

  // proto:  const QMetaObject * QFontDialog::metaObject();
func (this *QFontDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFontDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QFontDialog10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontDialog", "metaObject", args)
  }

}

  // proto:  QFont QFontDialog::selectedFont();
func (this *QFontDialog) selectedFont(args ...interface{}) () {
  // selectedFont()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFontDialog12selectedFontEv
    // invoke: QFont selectedFont()
    C._ZNK11QFontDialog12selectedFontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontDialog", "selectedFont", args)
  }

}

  // proto:  void QFontDialog::setCurrentFont(const QFont & font);
func (this *QFontDialog) setCurrentFont(args ...interface{}) () {
  // setCurrentFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialog14setCurrentFontERK5QFont
    // invoke: void setCurrentFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFontDialog14setCurrentFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDialog", "setCurrentFont", args)
  }

}

// <= body block end

