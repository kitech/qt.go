package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qcolordialog.h
// dst-file: /src/widgets/qcolordialog.go
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
  // proto:  void QColorDialog::setCurrentColor(const QColor & color);
extern void C_ZN12QColorDialog15setCurrentColorERK6QColor(void* qthis, void* arg0); // 4
  // proto: static void QColorDialog::setCustomColor(int index, QColor color);
extern void C_ZN12QColorDialog14setCustomColorEi6QColor(int32_t arg0, void* arg1); // 4
  // proto: static QRgb QColorDialog::getRgba(QRgb rgba, bool * ok, QWidget * parent);
extern void C_ZN12QColorDialog7getRgbaEjPbP7QWidget(int32_t arg0, bool* arg1, void* arg2); // 4
  // proto:  void QColorDialog::open(QObject * receiver, const char * member);
extern void C_ZN12QColorDialog4openEP7QObjectPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto: static QColor QColorDialog::standardColor(int index);
extern void C_ZN12QColorDialog13standardColorEi(int32_t arg0); // 4
  // proto: static void QColorDialog::setStandardColor(int index, QColor color);
extern void C_ZN12QColorDialog16setStandardColorEi6QColor(int32_t arg0, void* arg1); // 4
  // proto:  void QColorDialog::~QColorDialog();
extern void C_ZN12QColorDialogD2Ev(void* qthis); // 4
  // proto: static int QColorDialog::customCount();
extern void C_ZN12QColorDialog11customCountEv(); // 4
  // proto:  const QMetaObject * QColorDialog::metaObject();
extern void C_ZNK12QColorDialog10metaObjectEv(void* qthis); // 4
  // proto:  QColor QColorDialog::currentColor();
extern void C_ZNK12QColorDialog12currentColorEv(void* qthis); // 4
  // proto:  QColor QColorDialog::selectedColor();
extern void C_ZNK12QColorDialog13selectedColorEv(void* qthis); // 4
  // proto:  void QColorDialog::QColorDialog(QWidget * parent);
extern void C_ZN12QColorDialogC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QColorDialog::QColorDialog(const QColor & initial, QWidget * parent);
extern void C_ZN12QColorDialogC2ERK6QColorP7QWidget(void* qthis, void* arg0, void* arg1); // 3
  // proto: static QColor QColorDialog::customColor(int index);
extern void C_ZN12QColorDialog11customColorEi(int32_t arg0); // 4
  // proto:  void QColorDialog::setVisible(bool visible);
extern void C_ZN12QColorDialog10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  ColorDialogOptions QColorDialog::options();
extern void C_ZNK12QColorDialog7optionsEv(void* qthis); // 4
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

// class sizeof(QColorDialog)=1
type QColorDialog struct {
  /*qbase*/ QDialog;
  qclsinst unsafe.Pointer /* *C.void */;
//  _colorSelected QColorDialog_colorSelected_signal;
//  _currentColorChanged QColorDialog_currentColorChanged_signal;
}

// setCurrentColor(const class QColor &)
func (this *QColorDialog) setCurrentColor(args ...interface{}) () {
  // setCurrentColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog15setCurrentColorERK6QColor
    // invoke: void setCurrentColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QColorDialog15setCurrentColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "setCurrentColor", args)
  }

}

// setCustomColor(int, class QColor)
func (this *QColorDialog) setCustomColor_s(args ...interface{}) () {
  // setCustomColor(int, class QColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "QColor"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog14setCustomColorEi6QColor
    // invoke: void setCustomColor(int, class QColor)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QColorDialog14setCustomColorEi6QColor(arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "setCustomColor", args)
  }

}

// getRgba(QRgb, _Bool *, class QWidget *)
func (this *QColorDialog) getRgba_s(args ...interface{}) () {
  // getRgba(QRgb, _Bool *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QRgb"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog7getRgbaEjPbP7QWidget
    // invoke: QRgb getRgba(QRgb, _Bool *, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.bool)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QColorDialog7getRgbaEjPbP7QWidget(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QColorDialog", "getRgba", args)
  }

}

// open(class QObject *, const char *)
func (this *QColorDialog) open(args ...interface{}) () {
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
    // invoke: _ZN12QColorDialog4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN12QColorDialog4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "open", args)
  }

}

// standardColor(int)
func (this *QColorDialog) standardColor_s(args ...interface{}) () {
  // standardColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog13standardColorEi
    // invoke: QColor standardColor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QColorDialog13standardColorEi(arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "standardColor", args)
  }

}

// setStandardColor(int, class QColor)
func (this *QColorDialog) setStandardColor_s(args ...interface{}) () {
  // setStandardColor(int, class QColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "QColor"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog16setStandardColorEi6QColor
    // invoke: void setStandardColor(int, class QColor)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QColorDialog16setStandardColorEi6QColor(arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "setStandardColor", args)
  }

}

// ~QColorDialog()
func (this *QColorDialog) FreeQColorDialog(args ...interface{}) () {
  // ~QColorDialog()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialogD0Ev
    // invoke: void ~QColorDialog()
    C.C_ZN12QColorDialogD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "~QColorDialog", args)
  }

}

// customCount()
func (this *QColorDialog) customCount_s(args ...interface{}) () {
  // customCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog11customCountEv
    // invoke: int customCount()
    C.C_ZN12QColorDialog11customCountEv()
  default:
    qtrt.ErrorResolve("QColorDialog", "customCount", args)
  }

}

// metaObject()
func (this *QColorDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QColorDialog10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "metaObject", args)
  }

}

// currentColor()
func (this *QColorDialog) currentColor(args ...interface{}) () {
  // currentColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog12currentColorEv
    // invoke: QColor currentColor()
    C.C_ZNK12QColorDialog12currentColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "currentColor", args)
  }

}

// selectedColor()
func (this *QColorDialog) selectedColor(args ...interface{}) () {
  // selectedColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog13selectedColorEv
    // invoke: QColor selectedColor()
    C.C_ZNK12QColorDialog13selectedColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "selectedColor", args)
  }

}

// QColorDialog(class QWidget *)
func NewQColorDialog(args ...interface{}) QColorDialog {
  // QColorDialog(class QWidget *)
  // QColorDialog(const class QColor &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialogC1EP7QWidget
    // invoke: void QColorDialog(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QColorDialogC2EP7QWidget(qthis, arg0)
  case 1:
    // invoke: _ZN12QColorDialogC1ERK6QColorP7QWidget
    // invoke: void QColorDialog(const class QColor &, class QWidget *)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN12QColorDialogC2ERK6QColorP7QWidget(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "QColorDialog", args)
  }

  return QColorDialog{}
}

// customColor(int)
func (this *QColorDialog) customColor_s(args ...interface{}) () {
  // customColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog11customColorEi
    // invoke: QColor customColor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QColorDialog11customColorEi(arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "customColor", args)
  }

}

// setVisible(_Bool)
func (this *QColorDialog) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QColorDialog10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "setVisible", args)
  }

}

// options()
func (this *QColorDialog) options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog7optionsEv
    // invoke: ColorDialogOptions options()
    C.C_ZNK12QColorDialog7optionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "options", args)
  }

}

// <= body block end

