package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qtoolbox.h
// dst-file: /src/widgets/qtoolbox.go
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
  // proto:  void QToolBox::setItemText(int index, const QString & text);
extern void C_ZN8QToolBox11setItemTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QIcon QToolBox::itemIcon(int index);
extern void* C_ZNK8QToolBox8itemIconEi(void* qthis, int32_t arg0); // 4
  // proto:  void QToolBox::setItemEnabled(int index, bool enabled);
extern void C_ZN8QToolBox14setItemEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QToolBox::setCurrentIndex(int index);
extern void C_ZN8QToolBox15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  QWidget * QToolBox::currentWidget();
extern void* C_ZNK8QToolBox13currentWidgetEv(void* qthis); // 4
  // proto:  void QToolBox::~QToolBox();
extern void C_ZN8QToolBoxD2Ev(void* qthis); // 4
  // proto:  void QToolBox::setCurrentWidget(QWidget * widget);
extern void C_ZN8QToolBox16setCurrentWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  int QToolBox::indexOf(QWidget * widget);
extern int32_t C_ZNK8QToolBox7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  int QToolBox::count();
extern int32_t C_ZNK8QToolBox5countEv(void* qthis); // 4
  // proto:  QWidget * QToolBox::widget(int index);
extern void* C_ZNK8QToolBox6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  void QToolBox::setItemToolTip(int index, const QString & toolTip);
extern void C_ZN8QToolBox14setItemToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QToolBox::setItemIcon(int index, const QIcon & icon);
extern void C_ZN8QToolBox11setItemIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  int QToolBox::insertItem(int index, QWidget * widget, const QIcon & icon, const QString & text);
extern int32_t C_ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  int QToolBox::insertItem(int index, QWidget * widget, const QString & text);
extern int32_t C_ZN8QToolBox10insertItemEiP7QWidgetRK7QString(void* qthis, int32_t arg0, void* arg1, void* arg2); // 2
  // proto:  int QToolBox::addItem(QWidget * widget, const QIcon & icon, const QString & text);
extern int32_t C_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  int QToolBox::addItem(QWidget * widget, const QString & text);
extern int32_t C_ZN8QToolBox7addItemEP7QWidgetRK7QString(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QToolBox::removeItem(int index);
extern void C_ZN8QToolBox10removeItemEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QToolBox::itemText(int index);
extern void* C_ZNK8QToolBox8itemTextEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QToolBox::itemToolTip(int index);
extern void* C_ZNK8QToolBox11itemToolTipEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QToolBox::metaObject();
extern void C_ZNK8QToolBox10metaObjectEv(void* qthis); // 4
  // proto:  int QToolBox::currentIndex();
extern int32_t C_ZNK8QToolBox12currentIndexEv(void* qthis); // 4
  // proto:  bool QToolBox::isItemEnabled(int index);
extern bool C_ZNK8QToolBox13isItemEnabledEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QToolBox)=1
type QToolBox struct {
  /*qbase*/ QFrame;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _currentChanged QToolBox_currentChanged_signal;
}

// setItemText(int, const class QString &)
func (this *QToolBox) SetItemText(args ...interface{}) () {
  // setItemText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox11setItemTextEiRK7QString
    // invoke: void setItemText(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QToolBox11setItemTextEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemText", args)
  }

  return
}

// itemIcon(int)
func (this *QToolBox) ItemIcon(args ...interface{}) (ret interface{}) {
  // itemIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox8itemIconEi
    // invoke: QIcon itemIcon(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBox8itemIconEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "itemIcon", args)
  }

  return
}

// setItemEnabled(int, _Bool)
func (this *QToolBox) SetItemEnabled(args ...interface{}) () {
  // setItemEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox14setItemEnabledEib
    // invoke: void setItemEnabled(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN8QToolBox14setItemEnabledEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemEnabled", args)
  }

  return
}

// setCurrentIndex(int)
func (this *QToolBox) SetCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolBox15setCurrentIndexEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "setCurrentIndex", args)
  }

  return
}

// currentWidget()
func (this *QToolBox) CurrentWidget(args ...interface{}) (ret interface{}) {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox13currentWidgetEv
    // invoke: QWidget * currentWidget()
    var ret0 = C.C_ZNK8QToolBox13currentWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "currentWidget", args)
  }

  return
}

// ~QToolBox()
func (this *QToolBox) Free(args ...interface{}) () {
  // ~QToolBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBoxD0Ev
    // invoke: void ~QToolBox()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN8QToolBoxD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QToolBox", "~QToolBox", args)
  }

  return
}

// setCurrentWidget(class QWidget *)
func (this *QToolBox) SetCurrentWidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolBox16setCurrentWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "setCurrentWidget", args)
  }

  return
}

// indexOf(class QWidget *)
func (this *QToolBox) IndexOf(args ...interface{}) (ret interface{}) {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBox7indexOfEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "indexOf", args)
  }

  return
}

// count()
func (this *QToolBox) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK8QToolBox5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "count", args)
  }

  return
}

// widget(int)
func (this *QToolBox) Widget(args ...interface{}) (ret interface{}) {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBox6widgetEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "widget", args)
  }

  return
}

// setItemToolTip(int, const class QString &)
func (this *QToolBox) SetItemToolTip(args ...interface{}) () {
  // setItemToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox14setItemToolTipEiRK7QString
    // invoke: void setItemToolTip(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QToolBox14setItemToolTipEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemToolTip", args)
  }

  return
}

// setItemIcon(int, const class QIcon &)
func (this *QToolBox) SetItemIcon(args ...interface{}) () {
  // setItemIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox11setItemIconEiRK5QIcon
    // invoke: void setItemIcon(int, const class QIcon &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN8QToolBox11setItemIconEiRK5QIcon(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QToolBox", "setItemIcon", args)
  }

  return
}

// insertItem(int, class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) InsertItem(args ...interface{}) (ret interface{}) {
  // insertItem(int, class QWidget *, const class QIcon &, const class QString &)
  // insertItem(int, class QWidget *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[0][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString
    // invoke: int insertItem(int, class QWidget *, const class QIcon &, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN8QToolBox10insertItemEiP7QWidgetRK5QIconRK7QString(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN8QToolBox10insertItemEiP7QWidgetRK7QString
    // invoke: int insertItem(int, class QWidget *, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN8QToolBox10insertItemEiP7QWidgetRK7QString(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "insertItem", args)
  }

  return
}

// addItem(class QWidget *, const class QIcon &, const class QString &)
func (this *QToolBox) AddItem(args ...interface{}) (ret interface{}) {
  // addItem(class QWidget *, const class QIcon &, const class QString &)
  // addItem(class QWidget *, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString
    // invoke: int addItem(class QWidget *, const class QIcon &, const class QString &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN8QToolBox7addItemEP7QWidgetRK5QIconRK7QString(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN8QToolBox7addItemEP7QWidgetRK7QString
    // invoke: int addItem(class QWidget *, const class QString &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN8QToolBox7addItemEP7QWidgetRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "addItem", args)
  }

  return
}

// removeItem(int)
func (this *QToolBox) RemoveItem(args ...interface{}) () {
  // removeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QToolBox10removeItemEi
    // invoke: void removeItem(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN8QToolBox10removeItemEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QToolBox", "removeItem", args)
  }

  return
}

// itemText(int)
func (this *QToolBox) ItemText(args ...interface{}) (ret interface{}) {
  // itemText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox8itemTextEi
    // invoke: QString itemText(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBox8itemTextEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "itemText", args)
  }

  return
}

// itemToolTip(int)
func (this *QToolBox) ItemToolTip(args ...interface{}) (ret interface{}) {
  // itemToolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox11itemToolTipEi
    // invoke: QString itemToolTip(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBox11itemToolTipEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "itemToolTip", args)
  }

  return
}

// metaObject()
func (this *QToolBox) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK8QToolBox10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QToolBox", "metaObject", args)
  }

  return
}

// currentIndex()
func (this *QToolBox) CurrentIndex(args ...interface{}) (ret interface{}) {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox12currentIndexEv
    // invoke: int currentIndex()
    var ret0 = C.C_ZNK8QToolBox12currentIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "currentIndex", args)
  }

  return
}

// isItemEnabled(int)
func (this *QToolBox) IsItemEnabled(args ...interface{}) (ret interface{}) {
  // isItemEnabled(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QToolBox13isItemEnabledEi
    // invoke: bool isItemEnabled(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK8QToolBox13isItemEnabledEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QToolBox", "isItemEnabled", args)
  }

  return
}

// <= body block end

