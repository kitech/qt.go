package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QFontDialog::open(QObject * receiver, const char * member);
extern void C_ZN11QFontDialog4openEP7QObjectPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QFont QFontDialog::selectedFont();
extern void* C_ZNK11QFontDialog12selectedFontEv(void* qthis); // 4
  // proto:  void QFontDialog::QFontDialog(QWidget * parent);
extern void* C_ZN11QFontDialogC2EP7QWidget(void* arg0); // 3
  // proto:  void QFontDialog::QFontDialog(const QFont & initial, QWidget * parent);
extern void* C_ZN11QFontDialogC2ERK5QFontP7QWidget(void* arg0, void* arg1); // 3
  // proto: static QFont QFontDialog::getFont(bool * ok, QWidget * parent);
extern void* C_ZN11QFontDialog7getFontEPbP7QWidget(void* arg0, void* arg1); // 4
  // proto:  QFont QFontDialog::currentFont();
extern void* C_ZNK11QFontDialog11currentFontEv(void* qthis); // 4
  // proto:  void QFontDialog::~QFontDialog();
extern void C_ZN11QFontDialogD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QFontDialog::metaObject();
extern void C_ZNK11QFontDialog10metaObjectEv(void* qthis); // 4
  // proto:  void QFontDialog::setCurrentFont(const QFont & font);
extern void C_ZN11QFontDialog14setCurrentFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  FontDialogOptions QFontDialog::options();
extern void C_ZNK11QFontDialog7optionsEv(void* qthis); // 4
  // proto:  void QFontDialog::setVisible(bool visible);
extern void C_ZN11QFontDialog10setVisibleEb(void* qthis, bool arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QFontDialog)=1
type QFontDialog struct {
  /*qbase*/ QDialog;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _fontSelected QFontDialog_fontSelected_signal;
//  _currentFontChanged QFontDialog_currentFontChanged_signal;
}

// open(class QObject *, const char *)
func (this *QFontDialog) Open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialog4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN11QFontDialog4openEP7QObjectPKc(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontDialog", "open", args)
  }

  return
}

// selectedFont()
func (this *QFontDialog) SelectedFont(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QFontDialog12selectedFontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDialog", "selectedFont", args)
  }

  return
}

// QFontDialog(class QWidget *)
func GcfreeQFontDialog(this *QFontDialog) {
  qtrt.UniverseFree(this)
}
func NewQFontDialog(args ...interface{}) *QFontDialog {
  // QFontDialog(class QWidget *)
  // QFontDialog(const class QFont &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialogC1EP7QWidget
    // invoke: void QFontDialog(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QFontDialogC2EP7QWidget(arg0)
    this := &QFontDialog{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFontDialog)
    return this
  case 1:
    // invoke: _ZN11QFontDialogC1ERK5QFontP7QWidget
    // invoke: void QFontDialog(const class QFont &, class QWidget *)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QFontDialogC2ERK5QFontP7QWidget(arg0, arg1)
    this := &QFontDialog{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFontDialog)
    return this
  default:
    qtrt.ErrorResolve("QFontDialog", "QFontDialog", args)
  }

  return nil // QFontDialog{Qclsinst:qthis}
}

// getFont(_Bool *, class QWidget *)
func (this *QFontDialog) GetFont_s(args ...interface{}) (ret interface{}) {
  // getFont(_Bool *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialog7getFontEPbP7QWidget
    // invoke: QFont getFont(_Bool *, class QWidget *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN11QFontDialog7getFontEPbP7QWidget(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDialog", "getFont", args)
  }

  return
}

// currentFont()
func (this *QFontDialog) CurrentFont(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QFontDialog11currentFontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontDialog", "currentFont", args)
  }

  return
}

// ~QFontDialog()
func (this *QFontDialog) Free(args ...interface{}) () {
  // ~QFontDialog()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialogD0Ev
    // invoke: void ~QFontDialog()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QFontDialogD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QFontDialog", "~QFontDialog", args)
  }

  return
}

// metaObject()
func (this *QFontDialog) MetaObject(args ...interface{}) () {
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
    C.C_ZNK11QFontDialog10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontDialog", "metaObject", args)
  }

  return
}

// setCurrentFont(const class QFont &)
func (this *QFontDialog) SetCurrentFont(args ...interface{}) () {
  // setCurrentFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFontDialog14setCurrentFontERK5QFont
    // invoke: void setCurrentFont(const class QFont &)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFontDialog14setCurrentFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDialog", "setCurrentFont", args)
  }

  return
}

// options()
func (this *QFontDialog) Options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFontDialog7optionsEv
    // invoke: FontDialogOptions options()
    C.C_ZNK11QFontDialog7optionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontDialog", "options", args)
  }

  return
}

// setVisible(_Bool)
func (this *QFontDialog) SetVisible(args ...interface{}) () {
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
    C.C_ZN11QFontDialog10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontDialog", "setVisible", args)
  }

  return
}

// <= body block end

