package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qapplication.h
// dst-file: /src/widgets/qapplication.go
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
  // proto: static QWidget * QApplication::activePopupWidget();
extern void* C_ZN12QApplication17activePopupWidgetEv(); // 4
  // proto: static void QApplication::closeAllWindows();
extern void C_ZN12QApplication15closeAllWindowsEv(); // 4
  // proto: static int QApplication::startDragDistance();
extern int32_t C_ZN12QApplication17startDragDistanceEv(); // 4
  // proto: static int QApplication::wheelScrollLines();
extern int32_t C_ZN12QApplication16wheelScrollLinesEv(); // 4
  // proto: static void QApplication::setColorSpec(int );
extern void C_ZN12QApplication12setColorSpecEi(int32_t arg0); // 4
  // proto:  bool QApplication::autoSipEnabled();
extern bool C_ZNK12QApplication14autoSipEnabledEv(void* qthis); // 4
  // proto: static QStyle * QApplication::setStyle(const QString & );
extern void* C_ZN12QApplication8setStyleERK7QString(void* arg0); // 4
  // proto: static void QApplication::setStyle(QStyle * );
extern void C_ZN12QApplication8setStyleEP6QStyle(void* arg0); // 4
  // proto:  bool QApplication::notify(QObject * , QEvent * );
extern bool C_ZN12QApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QWidget * QApplication::focusWidget();
extern void* C_ZN12QApplication11focusWidgetEv(); // 4
  // proto: static QWidgetList QApplication::allWidgets();
extern void C_ZN12QApplication10allWidgetsEv(); // 4
  // proto: static QWidget * QApplication::widgetAt(int x, int y);
extern void* C_ZN12QApplication8widgetAtEii(int32_t arg0, int32_t arg1); // 2
  // proto: static QWidget * QApplication::widgetAt(const QPoint & p);
extern void* C_ZN12QApplication8widgetAtERK6QPoint(void* arg0); // 4
  // proto: static void QApplication::aboutQt();
extern void C_ZN12QApplication7aboutQtEv(); // 4
  // proto: static void QApplication::setPalette(const QPalette & , const char * className);
extern void C_ZN12QApplication10setPaletteERK8QPalettePKc(void* arg0, void* arg1); // 4
  // proto: static QPalette QApplication::palette(const char * className);
extern void* C_ZN12QApplication7paletteEPKc(void* arg0); // 4
  // proto: static QPalette QApplication::palette(const QWidget * );
extern void* C_ZN12QApplication7paletteEPK7QWidget(void* arg0); // 4
  // proto: static QFont QApplication::font();
extern void* C_ZN12QApplication4fontEv(); // 4
  // proto: static QFont QApplication::font(const char * className);
extern void* C_ZN12QApplication4fontEPKc(void* arg0); // 4
  // proto: static QFont QApplication::font(const QWidget * );
extern void* C_ZN12QApplication4fontEPK7QWidget(void* arg0); // 4
  // proto: static void QApplication::setDoubleClickInterval(int );
extern void C_ZN12QApplication22setDoubleClickIntervalEi(int32_t arg0); // 4
  // proto: static QWidgetList QApplication::topLevelWidgets();
extern void C_ZN12QApplication15topLevelWidgetsEv(); // 4
  // proto: static QDesktopWidget * QApplication::desktop();
extern void* C_ZN12QApplication7desktopEv(); // 4
  // proto: static void QApplication::beep();
extern void C_ZN12QApplication4beepEv(); // 4
  // proto:  const QMetaObject * QApplication::metaObject();
extern void C_ZNK12QApplication10metaObjectEv(void* qthis); // 4
  // proto: static void QApplication::setCursorFlashTime(int );
extern void C_ZN12QApplication18setCursorFlashTimeEi(int32_t arg0); // 4
  // proto: static QWidget * QApplication::activeWindow();
extern void* C_ZN12QApplication12activeWindowEv(); // 4
  // proto: static QWidget * QApplication::topLevelAt(const QPoint & p);
extern void* C_ZN12QApplication10topLevelAtERK6QPoint(void* arg0); // 4
  // proto: static QWidget * QApplication::topLevelAt(int x, int y);
extern void* C_ZN12QApplication10topLevelAtEii(int32_t arg0, int32_t arg1); // 2
  // proto: static QFontMetrics QApplication::fontMetrics();
extern void* C_ZN12QApplication11fontMetricsEv(); // 4
  // proto: static void QApplication::setKeyboardInputInterval(int );
extern void C_ZN12QApplication24setKeyboardInputIntervalEi(int32_t arg0); // 4
  // proto: static int QApplication::colorSpec();
extern int32_t C_ZN12QApplication9colorSpecEv(); // 4
  // proto: static QStyle * QApplication::style();
extern void* C_ZN12QApplication5styleEv(); // 4
  // proto: static void QApplication::alert(QWidget * widget, int duration);
extern void C_ZN12QApplication5alertEP7QWidgeti(void* arg0, int32_t arg1); // 4
  // proto: static QWidget * QApplication::activeModalWidget();
extern void* C_ZN12QApplication17activeModalWidgetEv(); // 4
  // proto:  QString QApplication::styleSheet();
extern void* C_ZNK12QApplication10styleSheetEv(void* qthis); // 4
  // proto: static QIcon QApplication::windowIcon();
extern void* C_ZN12QApplication10windowIconEv(); // 4
  // proto: static int QApplication::cursorFlashTime();
extern int32_t C_ZN12QApplication15cursorFlashTimeEv(); // 4
  // proto: static void QApplication::setGlobalStrut(const QSize & );
extern void C_ZN12QApplication14setGlobalStrutERK5QSize(void* arg0); // 4
  // proto: static void QApplication::setActiveWindow(QWidget * act);
extern void C_ZN12QApplication15setActiveWindowEP7QWidget(void* arg0); // 4
  // proto:  void QApplication::~QApplication();
extern void C_ZN12QApplicationD2Ev(void* qthis); // 4
  // proto: static QSize QApplication::globalStrut();
extern void* C_ZN12QApplication11globalStrutEv(); // 4
  // proto: static void QApplication::setStartDragDistance(int l);
extern void C_ZN12QApplication20setStartDragDistanceEi(int32_t arg0); // 4
  // proto: static void QApplication::setWheelScrollLines(int );
extern void C_ZN12QApplication19setWheelScrollLinesEi(int32_t arg0); // 4
  // proto: static int QApplication::keyboardInputInterval();
extern int32_t C_ZN12QApplication21keyboardInputIntervalEv(); // 4
  // proto: static void QApplication::setStartDragTime(int ms);
extern void C_ZN12QApplication16setStartDragTimeEi(int32_t arg0); // 4
  // proto:  void QApplication::QApplication(int & argc, char ** argv, int );
extern void* C_ZN12QApplicationC2ERiPPci(void* arg0, void* arg1, int32_t arg2); // 3
  // proto: static int QApplication::startDragTime();
extern int32_t C_ZN12QApplication13startDragTimeEv(); // 4
  // proto: static void QApplication::setWindowIcon(const QIcon & icon);
extern void C_ZN12QApplication13setWindowIconERK5QIcon(void* arg0); // 4
  // proto:  void QApplication::setAutoSipEnabled(const bool enabled);
extern void C_ZN12QApplication17setAutoSipEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QApplication::setStyleSheet(const QString & sheet);
extern void C_ZN12QApplication13setStyleSheetERK7QString(void* qthis, void* arg0); // 4
  // proto: static int QApplication::exec();
extern int32_t C_ZN12QApplication4execEv(); // 4
  // proto: static void QApplication::setFont(const QFont & , const char * className);
extern void C_ZN12QApplication7setFontERK5QFontPKc(void* arg0, void* arg1); // 4
  // proto: static int QApplication::doubleClickInterval();
extern int32_t C_ZN12QApplication19doubleClickIntervalEv(); // 4
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

// class sizeof(QApplication)=1
type QApplication struct {
  /*qbase*/ qtgui.QGuiApplication;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _focusChanged QApplication_focusChanged_signal;
}

// activePopupWidget()
func (this *QApplication) Activepopupwidget_S(args ...interface{}) (ret interface{}) {
  // activePopupWidget()
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
    // invoke: _ZN12QApplication17activePopupWidgetEv
    // invoke: QWidget * activePopupWidget()
    var ret0 = C.C_ZN12QApplication17activePopupWidgetEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "activePopupWidget", args)
  }

  return
}

// closeAllWindows()
func (this *QApplication) Closeallwindows_S(args ...interface{}) () {
  // closeAllWindows()
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
    // invoke: _ZN12QApplication15closeAllWindowsEv
    // invoke: void closeAllWindows()
    C.C_ZN12QApplication15closeAllWindowsEv()
  default:
    qtrt.ErrorResolve("QApplication", "closeAllWindows", args)
  }

  return
}

// startDragDistance()
func (this *QApplication) Startdragdistance_S(args ...interface{}) (ret interface{}) {
  // startDragDistance()
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
    // invoke: _ZN12QApplication17startDragDistanceEv
    // invoke: int startDragDistance()
    var ret0 = C.C_ZN12QApplication17startDragDistanceEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "startDragDistance", args)
  }

  return
}

// wheelScrollLines()
func (this *QApplication) Wheelscrolllines_S(args ...interface{}) (ret interface{}) {
  // wheelScrollLines()
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
    // invoke: _ZN12QApplication16wheelScrollLinesEv
    // invoke: int wheelScrollLines()
    var ret0 = C.C_ZN12QApplication16wheelScrollLinesEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "wheelScrollLines", args)
  }

  return
}

// setColorSpec(int)
func (this *QApplication) Setcolorspec_S(args ...interface{}) () {
  // setColorSpec(int)
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
    // invoke: _ZN12QApplication12setColorSpecEi
    // invoke: void setColorSpec(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication12setColorSpecEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setColorSpec", args)
  }

  return
}

// autoSipEnabled()
func (this *QApplication) Autosipenabled(args ...interface{}) (ret interface{}) {
  // autoSipEnabled()
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
    // invoke: _ZNK12QApplication14autoSipEnabledEv
    // invoke: bool autoSipEnabled()
    var ret0 = C.C_ZNK12QApplication14autoSipEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "autoSipEnabled", args)
  }

  return
}

// setStyle(const class QString &)
func (this *QApplication) Setstyle_S(args ...interface{}) (ret interface{}) {
  // setStyle(const class QString &)
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyle{}) // "QStyle *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication8setStyleERK7QString
    // invoke: QStyle * setStyle(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QApplication8setStyleERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QApplication8setStyleEP6QStyle
    // invoke: void setStyle(class QStyle *)
    var arg0 = args[0].(*QStyle).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication8setStyleEP6QStyle(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStyle", args)
  }

  return
}

// notify(class QObject *, class QEvent *)
func (this *QApplication) Notify(args ...interface{}) (ret interface{}) {
  // notify(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication6notifyEP7QObjectP6QEvent
    // invoke: bool notify(class QObject *, class QEvent *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QApplication6notifyEP7QObjectP6QEvent(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "notify", args)
  }

  return
}

// focusWidget()
func (this *QApplication) Focuswidget_S(args ...interface{}) (ret interface{}) {
  // focusWidget()
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
    // invoke: _ZN12QApplication11focusWidgetEv
    // invoke: QWidget * focusWidget()
    var ret0 = C.C_ZN12QApplication11focusWidgetEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "focusWidget", args)
  }

  return
}

// allWidgets()
func (this *QApplication) Allwidgets_S(args ...interface{}) () {
  // allWidgets()
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
    // invoke: _ZN12QApplication10allWidgetsEv
    // invoke: QWidgetList allWidgets()
    C.C_ZN12QApplication10allWidgetsEv()
  default:
    qtrt.ErrorResolve("QApplication", "allWidgets", args)
  }

  return
}

// widgetAt(int, int)
func (this *QApplication) Widgetat_S(args ...interface{}) (ret interface{}) {
  // widgetAt(int, int)
  // widgetAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication8widgetAtEii
    // invoke: QWidget * widgetAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QApplication8widgetAtEii(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QApplication8widgetAtERK6QPoint
    // invoke: QWidget * widgetAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QApplication8widgetAtERK6QPoint(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "widgetAt", args)
  }

  return
}

// aboutQt()
func (this *QApplication) Aboutqt_S(args ...interface{}) () {
  // aboutQt()
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
    // invoke: _ZN12QApplication7aboutQtEv
    // invoke: void aboutQt()
    C.C_ZN12QApplication7aboutQtEv()
  default:
    qtrt.ErrorResolve("QApplication", "aboutQt", args)
  }

  return
}

// setPalette(const class QPalette &, const char *)
func (this *QApplication) Setpalette_S(args ...interface{}) () {
  // setPalette(const class QPalette &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPalette{}) // "const QPalette &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication10setPaletteERK8QPalettePKc
    // invoke: void setPalette(const class QPalette &, const char *)
    var arg0 = args[0].(*qtgui.QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN12QApplication10setPaletteERK8QPalettePKc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "setPalette", args)
  }

  return
}

// palette(const char *)
func (this *QApplication) Palette_S(args ...interface{}) (ret interface{}) {
  // palette(const char *)
  // palette(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication7paletteEPKc
    // invoke: QPalette palette(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN12QApplication7paletteEPKc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPalette{}) // "QPalette"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QApplication7paletteEPK7QWidget
    // invoke: QPalette palette(const class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QApplication7paletteEPK7QWidget(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPalette{}) // "QPalette"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "palette", args)
  }

  return
}

// font()
func (this *QApplication) Font_S(args ...interface{}) (ret interface{}) {
  // font()
  // font(const char *)
  // font(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZN12QApplication4fontEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QApplication4fontEPKc
    // invoke: QFont font(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN12QApplication4fontEPKc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN12QApplication4fontEPK7QWidget
    // invoke: QFont font(const class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QApplication4fontEPK7QWidget(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "font", args)
  }

  return
}

// setDoubleClickInterval(int)
func (this *QApplication) Setdoubleclickinterval_S(args ...interface{}) () {
  // setDoubleClickInterval(int)
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
    // invoke: _ZN12QApplication22setDoubleClickIntervalEi
    // invoke: void setDoubleClickInterval(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication22setDoubleClickIntervalEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setDoubleClickInterval", args)
  }

  return
}

// topLevelWidgets()
func (this *QApplication) Toplevelwidgets_S(args ...interface{}) () {
  // topLevelWidgets()
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
    // invoke: _ZN12QApplication15topLevelWidgetsEv
    // invoke: QWidgetList topLevelWidgets()
    C.C_ZN12QApplication15topLevelWidgetsEv()
  default:
    qtrt.ErrorResolve("QApplication", "topLevelWidgets", args)
  }

  return
}

// desktop()
func (this *QApplication) Desktop_S(args ...interface{}) (ret interface{}) {
  // desktop()
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
    // invoke: _ZN12QApplication7desktopEv
    // invoke: QDesktopWidget * desktop()
    var ret0 = C.C_ZN12QApplication7desktopEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDesktopWidget{}) // "QDesktopWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "desktop", args)
  }

  return
}

// beep()
func (this *QApplication) Beep_S(args ...interface{}) () {
  // beep()
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
    // invoke: _ZN12QApplication4beepEv
    // invoke: void beep()
    C.C_ZN12QApplication4beepEv()
  default:
    qtrt.ErrorResolve("QApplication", "beep", args)
  }

  return
}

// metaObject()
func (this *QApplication) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK12QApplication10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QApplication10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QApplication", "metaObject", args)
  }

  return
}

// setCursorFlashTime(int)
func (this *QApplication) Setcursorflashtime_S(args ...interface{}) () {
  // setCursorFlashTime(int)
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
    // invoke: _ZN12QApplication18setCursorFlashTimeEi
    // invoke: void setCursorFlashTime(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication18setCursorFlashTimeEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setCursorFlashTime", args)
  }

  return
}

// activeWindow()
func (this *QApplication) Activewindow_S(args ...interface{}) (ret interface{}) {
  // activeWindow()
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
    // invoke: _ZN12QApplication12activeWindowEv
    // invoke: QWidget * activeWindow()
    var ret0 = C.C_ZN12QApplication12activeWindowEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "activeWindow", args)
  }

  return
}

// topLevelAt(const class QPoint &)
func (this *QApplication) Toplevelat_S(args ...interface{}) (ret interface{}) {
  // topLevelAt(const class QPoint &)
  // topLevelAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication10topLevelAtERK6QPoint
    // invoke: QWidget * topLevelAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QApplication10topLevelAtERK6QPoint(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QApplication10topLevelAtEii
    // invoke: QWidget * topLevelAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QApplication10topLevelAtEii(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "topLevelAt", args)
  }

  return
}

// fontMetrics()
func (this *QApplication) Fontmetrics_S(args ...interface{}) (ret interface{}) {
  // fontMetrics()
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
    // invoke: _ZN12QApplication11fontMetricsEv
    // invoke: QFontMetrics fontMetrics()
    var ret0 = C.C_ZN12QApplication11fontMetricsEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFontMetrics{}) // "QFontMetrics"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "fontMetrics", args)
  }

  return
}

// setKeyboardInputInterval(int)
func (this *QApplication) Setkeyboardinputinterval_S(args ...interface{}) () {
  // setKeyboardInputInterval(int)
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
    // invoke: _ZN12QApplication24setKeyboardInputIntervalEi
    // invoke: void setKeyboardInputInterval(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication24setKeyboardInputIntervalEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setKeyboardInputInterval", args)
  }

  return
}

// colorSpec()
func (this *QApplication) Colorspec_S(args ...interface{}) (ret interface{}) {
  // colorSpec()
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
    // invoke: _ZN12QApplication9colorSpecEv
    // invoke: int colorSpec()
    var ret0 = C.C_ZN12QApplication9colorSpecEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "colorSpec", args)
  }

  return
}

// style()
func (this *QApplication) Style_S(args ...interface{}) (ret interface{}) {
  // style()
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
    // invoke: _ZN12QApplication5styleEv
    // invoke: QStyle * style()
    var ret0 = C.C_ZN12QApplication5styleEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "style", args)
  }

  return
}

// alert(class QWidget *, int)
func (this *QApplication) Alert_S(args ...interface{}) () {
  // alert(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication5alertEP7QWidgeti
    // invoke: void alert(class QWidget *, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QApplication5alertEP7QWidgeti(arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "alert", args)
  }

  return
}

// activeModalWidget()
func (this *QApplication) Activemodalwidget_S(args ...interface{}) (ret interface{}) {
  // activeModalWidget()
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
    // invoke: _ZN12QApplication17activeModalWidgetEv
    // invoke: QWidget * activeModalWidget()
    var ret0 = C.C_ZN12QApplication17activeModalWidgetEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "activeModalWidget", args)
  }

  return
}

// styleSheet()
func (this *QApplication) Stylesheet(args ...interface{}) (ret interface{}) {
  // styleSheet()
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
    // invoke: _ZNK12QApplication10styleSheetEv
    // invoke: QString styleSheet()
    var ret0 = C.C_ZNK12QApplication10styleSheetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "styleSheet", args)
  }

  return
}

// windowIcon()
func (this *QApplication) Windowicon_S(args ...interface{}) (ret interface{}) {
  // windowIcon()
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
    // invoke: _ZN12QApplication10windowIconEv
    // invoke: QIcon windowIcon()
    var ret0 = C.C_ZN12QApplication10windowIconEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "windowIcon", args)
  }

  return
}

// cursorFlashTime()
func (this *QApplication) Cursorflashtime_S(args ...interface{}) (ret interface{}) {
  // cursorFlashTime()
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
    // invoke: _ZN12QApplication15cursorFlashTimeEv
    // invoke: int cursorFlashTime()
    var ret0 = C.C_ZN12QApplication15cursorFlashTimeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "cursorFlashTime", args)
  }

  return
}

// setGlobalStrut(const class QSize &)
func (this *QApplication) Setglobalstrut_S(args ...interface{}) () {
  // setGlobalStrut(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication14setGlobalStrutERK5QSize
    // invoke: void setGlobalStrut(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication14setGlobalStrutERK5QSize(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setGlobalStrut", args)
  }

  return
}

// setActiveWindow(class QWidget *)
func (this *QApplication) Setactivewindow_S(args ...interface{}) () {
  // setActiveWindow(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication15setActiveWindowEP7QWidget
    // invoke: void setActiveWindow(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication15setActiveWindowEP7QWidget(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setActiveWindow", args)
  }

  return
}

// ~QApplication()
func (this *QApplication) Freeqapplication(args ...interface{}) () {
  // ~QApplication()
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
    // invoke: _ZN12QApplicationD0Ev
    // invoke: void ~QApplication()
    C.C_ZN12QApplicationD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QApplication", "~QApplication", args)
  }

  return
}

// globalStrut()
func (this *QApplication) Globalstrut_S(args ...interface{}) (ret interface{}) {
  // globalStrut()
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
    // invoke: _ZN12QApplication11globalStrutEv
    // invoke: QSize globalStrut()
    var ret0 = C.C_ZN12QApplication11globalStrutEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "globalStrut", args)
  }

  return
}

// setStartDragDistance(int)
func (this *QApplication) Setstartdragdistance_S(args ...interface{}) () {
  // setStartDragDistance(int)
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
    // invoke: _ZN12QApplication20setStartDragDistanceEi
    // invoke: void setStartDragDistance(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication20setStartDragDistanceEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStartDragDistance", args)
  }

  return
}

// setWheelScrollLines(int)
func (this *QApplication) Setwheelscrolllines_S(args ...interface{}) () {
  // setWheelScrollLines(int)
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
    // invoke: _ZN12QApplication19setWheelScrollLinesEi
    // invoke: void setWheelScrollLines(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication19setWheelScrollLinesEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setWheelScrollLines", args)
  }

  return
}

// keyboardInputInterval()
func (this *QApplication) Keyboardinputinterval_S(args ...interface{}) (ret interface{}) {
  // keyboardInputInterval()
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
    // invoke: _ZN12QApplication21keyboardInputIntervalEv
    // invoke: int keyboardInputInterval()
    var ret0 = C.C_ZN12QApplication21keyboardInputIntervalEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "keyboardInputInterval", args)
  }

  return
}

// setStartDragTime(int)
func (this *QApplication) Setstartdragtime_S(args ...interface{}) () {
  // setStartDragTime(int)
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
    // invoke: _ZN12QApplication16setStartDragTimeEi
    // invoke: void setStartDragTime(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication16setStartDragTimeEi(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStartDragTime", args)
  }

  return
}

// QApplication(int &, char **, int)
func NewQApplication(args ...interface{}) *QApplication {
  // QApplication(int &, char **, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int &"
  vtys[0][1] = qtrt.ByteTy(true) // "char **"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplicationC1ERiPPci
    // invoke: void QApplication(int &, char **, int)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QApplicationC2ERiPPci(arg0, arg1, arg2)
    return &QApplication{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QApplication", "QApplication", args)
  }

  return nil // QApplication{Qclsinst:qthis}
}

// startDragTime()
func (this *QApplication) Startdragtime_S(args ...interface{}) (ret interface{}) {
  // startDragTime()
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
    // invoke: _ZN12QApplication13startDragTimeEv
    // invoke: int startDragTime()
    var ret0 = C.C_ZN12QApplication13startDragTimeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "startDragTime", args)
  }

  return
}

// setWindowIcon(const class QIcon &)
func (this *QApplication) Setwindowicon_S(args ...interface{}) () {
  // setWindowIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication13setWindowIconERK5QIcon
    // invoke: void setWindowIcon(const class QIcon &)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication13setWindowIconERK5QIcon(arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setWindowIcon", args)
  }

  return
}

// setAutoSipEnabled(const _Bool)
func (this *QApplication) Setautosipenabled(args ...interface{}) () {
  // setAutoSipEnabled(const _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "const bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication17setAutoSipEnabledEb
    // invoke: void setAutoSipEnabled(const _Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication17setAutoSipEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setAutoSipEnabled", args)
  }

  return
}

// setStyleSheet(const class QString &)
func (this *QApplication) Setstylesheet(args ...interface{}) () {
  // setStyleSheet(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication13setStyleSheetERK7QString
    // invoke: void setStyleSheet(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QApplication13setStyleSheetERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QApplication", "setStyleSheet", args)
  }

  return
}

// exec()
func (this *QApplication) Exec_S(args ...interface{}) (ret interface{}) {
  // exec()
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
    // invoke: _ZN12QApplication4execEv
    // invoke: int exec()
    var ret0 = C.C_ZN12QApplication4execEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "exec", args)
  }

  return
}

// setFont(const class QFont &, const char *)
func (this *QApplication) Setfont_S(args ...interface{}) () {
  // setFont(const class QFont &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QApplication7setFontERK5QFontPKc
    // invoke: void setFont(const class QFont &, const char *)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN12QApplication7setFontERK5QFontPKc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QApplication", "setFont", args)
  }

  return
}

// doubleClickInterval()
func (this *QApplication) Doubleclickinterval_S(args ...interface{}) (ret interface{}) {
  // doubleClickInterval()
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
    // invoke: _ZN12QApplication19doubleClickIntervalEv
    // invoke: int doubleClickInterval()
    var ret0 = C.C_ZN12QApplication19doubleClickIntervalEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QApplication", "doubleClickInterval", args)
  }

  return
}

// <= body block end

