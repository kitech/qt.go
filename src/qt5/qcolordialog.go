package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QColorDialog::QColorDialog(const QColorDialog & );
extern void* dector_ZN12QColorDialogC1ERKS_(void* arg0);
extern void _ZN12QColorDialogC1ERKS_(void* qthis, void* arg0);
  // proto:  QColor QColorDialog::currentColor();
extern void _ZNK12QColorDialog12currentColorEv(void* qthis);
  // proto: static QColor QColorDialog::customColor(int index);
extern void _ZN12QColorDialog11customColorEi(int arg0);
  // proto:  const QMetaObject * QColorDialog::metaObject();
extern void _ZNK12QColorDialog10metaObjectEv(void* qthis);
  // proto:  void QColorDialog::QColorDialog(const QColor & initial, QWidget * parent);
extern void* dector_ZN12QColorDialogC1ERK6QColorP7QWidget(void* arg0, void* arg1);
extern void _ZN12QColorDialogC1ERK6QColorP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto: static void QColorDialog::setStandardColor(int index, QColor color);
extern void _ZN12QColorDialog16setStandardColorEi6QColor(int arg0, void* arg1);
  // proto:  void QColorDialog::open(QObject * receiver, const char * member);
extern void _ZN12QColorDialog4openEP7QObjectPKc(void* qthis, void* arg0, char* arg1);
  // proto:  QColor QColorDialog::selectedColor();
extern void _ZNK12QColorDialog13selectedColorEv(void* qthis);
  // proto:  void QColorDialog::~QColorDialog();
extern void _ZN12QColorDialogD0Ev(void* qthis);
  // proto:  void QColorDialog::setVisible(bool visible);
extern void _ZN12QColorDialog10setVisibleEb(void* qthis, bool arg0);
  // proto:  void QColorDialog::setCurrentColor(const QColor & color);
extern void _ZN12QColorDialog15setCurrentColorERK6QColor(void* qthis, void* arg0);
  // proto: static QColor QColorDialog::standardColor(int index);
extern void _ZN12QColorDialog13standardColorEi(int arg0);
  // proto: static QRgb QColorDialog::getRgba(QRgb rgba, bool * ok, QWidget * parent);
extern void _ZN12QColorDialog7getRgbaEjPbP7QWidget(unsigned int arg0, bool* arg1, void* arg2);
  // proto: static void QColorDialog::setCustomColor(int index, QColor color);
extern void _ZN12QColorDialog14setCustomColorEi6QColor(int arg0, void* arg1);
  // proto:  void QColorDialog::QColorDialog(QWidget * parent);
extern void* dector_ZN12QColorDialogC1EP7QWidget(void* arg0);
extern void _ZN12QColorDialogC1EP7QWidget(void* qthis, void* arg0);
  // proto: static int QColorDialog::customCount();
extern void _ZN12QColorDialog11customCountEv();
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
  qclsinst uint64 /* *mut c_void*/;
//  _colorSelected QColorDialog_colorSelected_signal;
//  _currentColorChanged QColorDialog_currentColorChanged_signal;
}

  // proto:  void QColorDialog::QColorDialog(const QColorDialog & );
func NewQColorDialog(args ...interface{}) QColorDialog {
  return QColorDialog{}
}

  // proto:  QColor QColorDialog::currentColor();
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
    C._ZNK12QColorDialog12currentColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "currentColor", args)
  }

}

  // proto: static QColor QColorDialog::customColor(int index);
func (this *QColorDialog) customColor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColorDialog", "customColor", args)
  }

}

  // proto:  const QMetaObject * QColorDialog::metaObject();
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
    C._ZNK12QColorDialog10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "metaObject", args)
  }

}

  // proto: static void QColorDialog::setStandardColor(int index, QColor color);
func (this *QColorDialog) setStandardColor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColorDialog", "setStandardColor", args)
  }

}

  // proto:  void QColorDialog::open(QObject * receiver, const char * member);
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
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C._ZN12QColorDialog4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "open", args)
  }

}

  // proto:  QColor QColorDialog::selectedColor();
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
    C._ZNK12QColorDialog13selectedColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "selectedColor", args)
  }

}

  // proto:  void QColorDialog::~QColorDialog();
func (this *QColorDialog) FreeQColorDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColorDialog", "~QColorDialog", args)
  }

}

  // proto:  void QColorDialog::setVisible(bool visible);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN12QColorDialog10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "setVisible", args)
  }

}

  // proto:  void QColorDialog::setCurrentColor(const QColor & color);
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
    C._ZN12QColorDialog15setCurrentColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "setCurrentColor", args)
  }

}

  // proto: static QColor QColorDialog::standardColor(int index);
func (this *QColorDialog) standardColor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColorDialog", "standardColor", args)
  }

}

  // proto: static QRgb QColorDialog::getRgba(QRgb rgba, bool * ok, QWidget * parent);
func (this *QColorDialog) getRgba_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColorDialog", "getRgba", args)
  }

}

  // proto: static void QColorDialog::setCustomColor(int index, QColor color);
func (this *QColorDialog) setCustomColor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColorDialog", "setCustomColor", args)
  }

}

  // proto: static int QColorDialog::customCount();
func (this *QColorDialog) customCount_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColorDialog", "customCount", args)
  }

}

// <= body block end

