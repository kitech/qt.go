package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  void QColorDialog::setCurrentColor(const QColor & color);
extern void C_ZN12QColorDialog15setCurrentColorERK6QColor(void* qthis, void* arg0); // 4
  // proto: static void QColorDialog::setCustomColor(int index, QColor color);
extern void C_ZN12QColorDialog14setCustomColorEi6QColor(int32_t arg0, void* arg1); // 4
  // proto: static QRgb QColorDialog::getRgba(QRgb rgba, bool * ok, QWidget * parent);
extern int32_t C_ZN12QColorDialog7getRgbaEjPbP7QWidget(int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QColorDialog::open(QObject * receiver, const char * member);
extern void C_ZN12QColorDialog4openEP7QObjectPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QColor QColorDialog::standardColor(int index);
extern void* C_ZN12QColorDialog13standardColorEi(int32_t arg0); // 4
  // proto: static void QColorDialog::setStandardColor(int index, QColor color);
extern void C_ZN12QColorDialog16setStandardColorEi6QColor(int32_t arg0, void* arg1); // 4
  // proto:  void QColorDialog::~QColorDialog();
extern void C_ZN12QColorDialogD2Ev(void* qthis); // 4
  // proto: static int QColorDialog::customCount();
extern int32_t C_ZN12QColorDialog11customCountEv(); // 4
  // proto:  const QMetaObject * QColorDialog::metaObject();
extern void C_ZNK12QColorDialog10metaObjectEv(void* qthis); // 4
  // proto:  QColor QColorDialog::currentColor();
extern void* C_ZNK12QColorDialog12currentColorEv(void* qthis); // 4
  // proto:  QColor QColorDialog::selectedColor();
extern void* C_ZNK12QColorDialog13selectedColorEv(void* qthis); // 4
  // proto:  void QColorDialog::QColorDialog(QWidget * parent);
extern void* C_ZN12QColorDialogC2EP7QWidget(void* arg0); // 3
  // proto:  void QColorDialog::QColorDialog(const QColor & initial, QWidget * parent);
extern void* C_ZN12QColorDialogC2ERK6QColorP7QWidget(void* arg0, void* arg1); // 3
  // proto: static QColor QColorDialog::customColor(int index);
extern void* C_ZN12QColorDialog11customColorEi(int32_t arg0); // 4
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
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QColorDialog)=1
type QColorDialog struct {
  /*qbase*/ QDialog;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _colorSelected QColorDialog_colorSelected_signal;
//  _currentColorChanged QColorDialog_currentColorChanged_signal;
}

// setCurrentColor(const class QColor &)
func (this *QColorDialog) Setcurrentcolor(args ...interface{}) () {
  // setCurrentColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog15setCurrentColorERK6QColor
    // invoke: void setCurrentColor(const class QColor &)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QColorDialog15setCurrentColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "setCurrentColor", args)
  }

  return
}

// setCustomColor(int, class QColor)
func (this *QColorDialog) Setcustomcolor_S(args ...interface{}) () {
  // setCustomColor(int, class QColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtgui.QColor{}) // "QColor"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog14setCustomColorEi6QColor
    // invoke: void setCustomColor(int, class QColor)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QColorDialog14setCustomColorEi6QColor(arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "setCustomColor", args)
  }

  return
}

// getRgba(QRgb, _Bool *, class QWidget *)
func (this *QColorDialog) Getrgba_S(args ...interface{}) (ret interface{}) {
  // getRgba(QRgb, _Bool *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QRgb"
  vtys[0][1] = qtrt.BoolTy(true) // "bool *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog7getRgbaEjPbP7QWidget
    // invoke: QRgb getRgba(QRgb, _Bool *, class QWidget *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*bool))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN12QColorDialog7getRgbaEjPbP7QWidget(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "QRgb"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColorDialog", "getRgba", args)
  }

  return
}

// open(class QObject *, const char *)
func (this *QColorDialog) Open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN12QColorDialog4openEP7QObjectPKc(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "open", args)
  }

  return
}

// standardColor(int)
func (this *QColorDialog) Standardcolor_S(args ...interface{}) (ret interface{}) {
  // standardColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog13standardColorEi
    // invoke: QColor standardColor(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QColorDialog13standardColorEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColorDialog", "standardColor", args)
  }

  return
}

// setStandardColor(int, class QColor)
func (this *QColorDialog) Setstandardcolor_S(args ...interface{}) () {
  // setStandardColor(int, class QColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtgui.QColor{}) // "QColor"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog16setStandardColorEi6QColor
    // invoke: void setStandardColor(int, class QColor)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QColorDialog16setStandardColorEi6QColor(arg0, arg1)
  default:
    qtrt.ErrorResolve("QColorDialog", "setStandardColor", args)
  }

  return
}

// ~QColorDialog()
func (this *QColorDialog) Freeqcolordialog(args ...interface{}) () {
  // ~QColorDialog()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialogD0Ev
    // invoke: void ~QColorDialog()
    C.C_ZN12QColorDialogD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "~QColorDialog", args)
  }

  return
}

// customCount()
func (this *QColorDialog) Customcount_S(args ...interface{}) (ret interface{}) {
  // customCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog11customCountEv
    // invoke: int customCount()
    var ret0 = C.C_ZN12QColorDialog11customCountEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColorDialog", "customCount", args)
  }

  return
}

// metaObject()
func (this *QColorDialog) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QColorDialog10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "metaObject", args)
  }

  return
}

// currentColor()
func (this *QColorDialog) Currentcolor(args ...interface{}) (ret interface{}) {
  // currentColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog12currentColorEv
    // invoke: QColor currentColor()
    var ret0 = C.C_ZNK12QColorDialog12currentColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColorDialog", "currentColor", args)
  }

  return
}

// selectedColor()
func (this *QColorDialog) Selectedcolor(args ...interface{}) (ret interface{}) {
  // selectedColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog13selectedColorEv
    // invoke: QColor selectedColor()
    var ret0 = C.C_ZNK12QColorDialog13selectedColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColorDialog", "selectedColor", args)
  }

  return
}

// QColorDialog(class QWidget *)
func NewQColorDialog(args ...interface{}) *QColorDialog {
  // QColorDialog(class QWidget *)
  // QColorDialog(const class QColor &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialogC1EP7QWidget
    // invoke: void QColorDialog(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QColorDialogC2EP7QWidget(arg0)
    return &QColorDialog{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QColorDialogC1ERK6QColorP7QWidget
    // invoke: void QColorDialog(const class QColor &, class QWidget *)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QColorDialogC2ERK6QColorP7QWidget(arg0, arg1)
    return &QColorDialog{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QColorDialog", "QColorDialog", args)
  }

  return nil // QColorDialog{Qclsinst:qthis}
}

// customColor(int)
func (this *QColorDialog) Customcolor_S(args ...interface{}) (ret interface{}) {
  // customColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog11customColorEi
    // invoke: QColor customColor(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QColorDialog11customColorEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColorDialog", "customColor", args)
  }

  return
}

// setVisible(_Bool)
func (this *QColorDialog) Setvisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QColorDialog10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QColorDialog10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColorDialog", "setVisible", args)
  }

  return
}

// options()
func (this *QColorDialog) Options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QColorDialog7optionsEv
    // invoke: ColorDialogOptions options()
    C.C_ZNK12QColorDialog7optionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColorDialog", "options", args)
  }

  return
}

// <= body block end

