package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qguiapplication.h
// dst-file: /src/gui/qguiapplication.go
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QStyleHints * QGuiApplication::styleHints();
extern void* C_ZN15QGuiApplication10styleHintsEv(); // 4
  // proto: static void QGuiApplication::setOverrideCursor(const QCursor & );
extern void C_ZN15QGuiApplication17setOverrideCursorERK7QCursor(void* arg0); // 4
  // proto: static void QGuiApplication::sync();
extern void C_ZN15QGuiApplication4syncEv(); // 4
  // proto: static void QGuiApplication::setPalette(const QPalette & pal);
extern void C_ZN15QGuiApplication10setPaletteERK8QPalette(void* arg0); // 4
  // proto: static QList<QScreen *> QGuiApplication::screens();
extern void C_ZN15QGuiApplication7screensEv(); // 4
  // proto: static Qt::ApplicationState QGuiApplication::applicationState();
extern void C_ZN15QGuiApplication16applicationStateEv(); // 4
  // proto:  bool QGuiApplication::notify(QObject * , QEvent * );
extern bool C_ZN15QGuiApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto: static Qt::KeyboardModifiers QGuiApplication::queryKeyboardModifiers();
extern void C_ZN15QGuiApplication22queryKeyboardModifiersEv(); // 4
  // proto:  bool QGuiApplication::isSavingSession();
extern bool C_ZNK15QGuiApplication15isSavingSessionEv(void* qthis); // 4
  // proto: static void QGuiApplication::changeOverrideCursor(const QCursor & );
extern void C_ZN15QGuiApplication20changeOverrideCursorERK7QCursor(void* arg0); // 4
  // proto: static QFont QGuiApplication::font();
extern void* C_ZN15QGuiApplication4fontEv(); // 4
  // proto: static bool QGuiApplication::quitOnLastWindowClosed();
extern bool C_ZN15QGuiApplication22quitOnLastWindowClosedEv(); // 4
  // proto: static QPalette QGuiApplication::palette();
extern void* C_ZN15QGuiApplication7paletteEv(); // 4
  // proto: static void QGuiApplication::setApplicationDisplayName(const QString & name);
extern void C_ZN15QGuiApplication25setApplicationDisplayNameERK7QString(void* arg0); // 4
  // proto: static Qt::LayoutDirection QGuiApplication::layoutDirection();
extern void C_ZN15QGuiApplication15layoutDirectionEv(); // 4
  // proto: static int QGuiApplication::exec();
extern int32_t C_ZN15QGuiApplication4execEv(); // 4
  // proto:  QString QGuiApplication::sessionId();
extern void* C_ZNK15QGuiApplication9sessionIdEv(void* qthis); // 4
  // proto:  QString QGuiApplication::sessionKey();
extern void* C_ZNK15QGuiApplication10sessionKeyEv(void* qthis); // 4
  // proto: static QObject * QGuiApplication::focusObject();
extern void* C_ZN15QGuiApplication11focusObjectEv(); // 4
  // proto: static QFunctionPointer QGuiApplication::platformFunction(const QByteArray & function);
extern void C_ZN15QGuiApplication16platformFunctionERK10QByteArray(void* arg0); // 4
  // proto: static Qt::KeyboardModifiers QGuiApplication::keyboardModifiers();
extern void C_ZN15QGuiApplication17keyboardModifiersEv(); // 4
  // proto: static QString QGuiApplication::platformName();
extern void* C_ZN15QGuiApplication12platformNameEv(); // 4
  // proto: static QWindow * QGuiApplication::focusWindow();
extern void* C_ZN15QGuiApplication11focusWindowEv(); // 4
  // proto: static bool QGuiApplication::isRightToLeft();
extern bool C_ZN15QGuiApplication13isRightToLeftEv(); // 2
  // proto: static QWindow * QGuiApplication::topLevelAt(const QPoint & pos);
extern void* C_ZN15QGuiApplication10topLevelAtERK6QPoint(void* arg0); // 4
  // proto: static QString QGuiApplication::applicationDisplayName();
extern void* C_ZN15QGuiApplication22applicationDisplayNameEv(); // 4
  // proto: static QInputMethod * QGuiApplication::inputMethod();
extern void* C_ZN15QGuiApplication11inputMethodEv(); // 4
  // proto: static QClipboard * QGuiApplication::clipboard();
extern void* C_ZN15QGuiApplication9clipboardEv(); // 4
  // proto: static QCursor * QGuiApplication::overrideCursor();
extern void* C_ZN15QGuiApplication14overrideCursorEv(); // 4
  // proto: static QScreen * QGuiApplication::primaryScreen();
extern void* C_ZN15QGuiApplication13primaryScreenEv(); // 4
  // proto: static QWindowList QGuiApplication::allWindows();
extern void C_ZN15QGuiApplication10allWindowsEv(); // 4
  // proto:  void QGuiApplication::~QGuiApplication();
extern void C_ZN15QGuiApplicationD2Ev(void* qthis); // 4
  // proto: static void QGuiApplication::setDesktopSettingsAware(bool on);
extern void C_ZN15QGuiApplication23setDesktopSettingsAwareEb(bool arg0); // 4
  // proto: static Qt::MouseButtons QGuiApplication::mouseButtons();
extern void C_ZN15QGuiApplication12mouseButtonsEv(); // 4
  // proto: static QPlatformNativeInterface * QGuiApplication::platformNativeInterface();
extern void C_ZN15QGuiApplication23platformNativeInterfaceEv(); // 4
  // proto:  const QMetaObject * QGuiApplication::metaObject();
extern void C_ZNK15QGuiApplication10metaObjectEv(void* qthis); // 4
  // proto:  qreal QGuiApplication::devicePixelRatio();
extern double C_ZNK15QGuiApplication16devicePixelRatioEv(void* qthis); // 4
  // proto: static QWindow * QGuiApplication::modalWindow();
extern void* C_ZN15QGuiApplication11modalWindowEv(); // 4
  // proto:  void QGuiApplication::QGuiApplication(int & argc, char ** argv, int );
extern void* C_ZN15QGuiApplicationC2ERiPPci(void* arg0, void* arg1, int32_t arg2); // 3
  // proto: static bool QGuiApplication::isLeftToRight();
extern bool C_ZN15QGuiApplication13isLeftToRightEv(); // 2
  // proto: static void QGuiApplication::setWindowIcon(const QIcon & icon);
extern void C_ZN15QGuiApplication13setWindowIconERK5QIcon(void* arg0); // 4
  // proto: static void QGuiApplication::setQuitOnLastWindowClosed(bool quit);
extern void C_ZN15QGuiApplication25setQuitOnLastWindowClosedEb(bool arg0); // 4
  // proto: static bool QGuiApplication::desktopSettingsAware();
extern bool C_ZN15QGuiApplication20desktopSettingsAwareEv(); // 4
  // proto: static void QGuiApplication::setFont(const QFont & );
extern void C_ZN15QGuiApplication7setFontERK5QFont(void* arg0); // 4
  // proto: static QWindowList QGuiApplication::topLevelWindows();
extern void C_ZN15QGuiApplication15topLevelWindowsEv(); // 4
  // proto:  bool QGuiApplication::isSessionRestored();
extern bool C_ZNK15QGuiApplication17isSessionRestoredEv(void* qthis); // 4
  // proto: static void QGuiApplication::restoreOverrideCursor();
extern void C_ZN15QGuiApplication21restoreOverrideCursorEv(); // 4
  // proto: static QIcon QGuiApplication::windowIcon();
extern void* C_ZN15QGuiApplication10windowIconEv(); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QGuiApplication)=1
type QGuiApplication struct {
  /*qbase*/ qtcore.QCoreApplication;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _saveStateRequest QGuiApplication_saveStateRequest_signal;
//  _applicationStateChanged QGuiApplication_applicationStateChanged_signal;
//  _screenAdded QGuiApplication_screenAdded_signal;
//  _screenRemoved QGuiApplication_screenRemoved_signal;
//  _focusObjectChanged QGuiApplication_focusObjectChanged_signal;
//  _paletteChanged QGuiApplication_paletteChanged_signal;
//  _focusWindowChanged QGuiApplication_focusWindowChanged_signal;
//  _lastWindowClosed QGuiApplication_lastWindowClosed_signal;
//  _fontDatabaseChanged QGuiApplication_fontDatabaseChanged_signal;
//  _layoutDirectionChanged QGuiApplication_layoutDirectionChanged_signal;
//  _commitDataRequest QGuiApplication_commitDataRequest_signal;
}

// styleHints()
func (this *QGuiApplication) StyleHints_s(args ...interface{}) (ret interface{}) {
  // styleHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication10styleHintsEv
    // invoke: QStyleHints * styleHints()
    var ret0 = C.C_ZN15QGuiApplication10styleHintsEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyleHints{}) // "QStyleHints *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "styleHints", args)
  }

  return
}

// setOverrideCursor(const class QCursor &)
func (this *QGuiApplication) SetOverrideCursor_s(args ...interface{}) () {
  // setOverrideCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication17setOverrideCursorERK7QCursor
    // invoke: void setOverrideCursor(const class QCursor &)
    var arg0 = args[0].(*QCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication17setOverrideCursorERK7QCursor(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setOverrideCursor", args)
  }

  return
}

// sync()
func (this *QGuiApplication) Sync_s(args ...interface{}) () {
  // sync()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication4syncEv
    // invoke: void sync()
    C.C_ZN15QGuiApplication4syncEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "sync", args)
  }

  return
}

// setPalette(const class QPalette &)
func (this *QGuiApplication) SetPalette_s(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication10setPaletteERK8QPalette
    // invoke: void setPalette(const class QPalette &)
    var arg0 = args[0].(*QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication10setPaletteERK8QPalette(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setPalette", args)
  }

  return
}

// screens()
func (this *QGuiApplication) Screens_s(args ...interface{}) () {
  // screens()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication7screensEv
    // invoke: QList<QScreen *> screens()
    C.C_ZN15QGuiApplication7screensEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "screens", args)
  }

  return
}

// applicationState()
func (this *QGuiApplication) ApplicationState_s(args ...interface{}) () {
  // applicationState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication16applicationStateEv
    // invoke: Qt::ApplicationState applicationState()
    C.C_ZN15QGuiApplication16applicationStateEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "applicationState", args)
  }

  return
}

// notify(class QObject *, class QEvent *)
func (this *QGuiApplication) Notify(args ...interface{}) (ret interface{}) {
  // notify(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(qtcore.QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication6notifyEP7QObjectP6QEvent
    // invoke: bool notify(class QObject *, class QEvent *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QEvent).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN15QGuiApplication6notifyEP7QObjectP6QEvent(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "notify", args)
  }

  return
}

// queryKeyboardModifiers()
func (this *QGuiApplication) QueryKeyboardModifiers_s(args ...interface{}) () {
  // queryKeyboardModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication22queryKeyboardModifiersEv
    // invoke: Qt::KeyboardModifiers queryKeyboardModifiers()
    C.C_ZN15QGuiApplication22queryKeyboardModifiersEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "queryKeyboardModifiers", args)
  }

  return
}

// isSavingSession()
func (this *QGuiApplication) IsSavingSession(args ...interface{}) (ret interface{}) {
  // isSavingSession()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication15isSavingSessionEv
    // invoke: bool isSavingSession()
    var ret0 = C.C_ZNK15QGuiApplication15isSavingSessionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "isSavingSession", args)
  }

  return
}

// changeOverrideCursor(const class QCursor &)
func (this *QGuiApplication) ChangeOverrideCursor_s(args ...interface{}) () {
  // changeOverrideCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication20changeOverrideCursorERK7QCursor
    // invoke: void changeOverrideCursor(const class QCursor &)
    var arg0 = args[0].(*QCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication20changeOverrideCursorERK7QCursor(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "changeOverrideCursor", args)
  }

  return
}

// font()
func (this *QGuiApplication) Font_s(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZN15QGuiApplication4fontEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "font", args)
  }

  return
}

// quitOnLastWindowClosed()
func (this *QGuiApplication) QuitOnLastWindowClosed_s(args ...interface{}) (ret interface{}) {
  // quitOnLastWindowClosed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication22quitOnLastWindowClosedEv
    // invoke: bool quitOnLastWindowClosed()
    var ret0 = C.C_ZN15QGuiApplication22quitOnLastWindowClosedEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "quitOnLastWindowClosed", args)
  }

  return
}

// palette()
func (this *QGuiApplication) Palette_s(args ...interface{}) (ret interface{}) {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication7paletteEv
    // invoke: QPalette palette()
    var ret0 = C.C_ZN15QGuiApplication7paletteEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPalette{}) // "QPalette"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "palette", args)
  }

  return
}

// setApplicationDisplayName(const class QString &)
func (this *QGuiApplication) SetApplicationDisplayName_s(args ...interface{}) () {
  // setApplicationDisplayName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication25setApplicationDisplayNameERK7QString
    // invoke: void setApplicationDisplayName(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication25setApplicationDisplayNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setApplicationDisplayName", args)
  }

  return
}

// layoutDirection()
func (this *QGuiApplication) LayoutDirection_s(args ...interface{}) () {
  // layoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication15layoutDirectionEv
    // invoke: Qt::LayoutDirection layoutDirection()
    C.C_ZN15QGuiApplication15layoutDirectionEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "layoutDirection", args)
  }

  return
}

// exec()
func (this *QGuiApplication) Exec_s(args ...interface{}) (ret interface{}) {
  // exec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication4execEv
    // invoke: int exec()
    var ret0 = C.C_ZN15QGuiApplication4execEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "exec", args)
  }

  return
}

// sessionId()
func (this *QGuiApplication) SessionId(args ...interface{}) (ret interface{}) {
  // sessionId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication9sessionIdEv
    // invoke: QString sessionId()
    var ret0 = C.C_ZNK15QGuiApplication9sessionIdEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "sessionId", args)
  }

  return
}

// sessionKey()
func (this *QGuiApplication) SessionKey(args ...interface{}) (ret interface{}) {
  // sessionKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication10sessionKeyEv
    // invoke: QString sessionKey()
    var ret0 = C.C_ZNK15QGuiApplication10sessionKeyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "sessionKey", args)
  }

  return
}

// focusObject()
func (this *QGuiApplication) FocusObject_s(args ...interface{}) (ret interface{}) {
  // focusObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication11focusObjectEv
    // invoke: QObject * focusObject()
    var ret0 = C.C_ZN15QGuiApplication11focusObjectEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "focusObject", args)
  }

  return
}

// platformFunction(const class QByteArray &)
func (this *QGuiApplication) PlatformFunction_s(args ...interface{}) () {
  // platformFunction(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication16platformFunctionERK10QByteArray
    // invoke: QFunctionPointer platformFunction(const class QByteArray &)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication16platformFunctionERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformFunction", args)
  }

  return
}

// keyboardModifiers()
func (this *QGuiApplication) KeyboardModifiers_s(args ...interface{}) () {
  // keyboardModifiers()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication17keyboardModifiersEv
    // invoke: Qt::KeyboardModifiers keyboardModifiers()
    C.C_ZN15QGuiApplication17keyboardModifiersEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "keyboardModifiers", args)
  }

  return
}

// platformName()
func (this *QGuiApplication) PlatformName_s(args ...interface{}) (ret interface{}) {
  // platformName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication12platformNameEv
    // invoke: QString platformName()
    var ret0 = C.C_ZN15QGuiApplication12platformNameEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformName", args)
  }

  return
}

// focusWindow()
func (this *QGuiApplication) FocusWindow_s(args ...interface{}) (ret interface{}) {
  // focusWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication11focusWindowEv
    // invoke: QWindow * focusWindow()
    var ret0 = C.C_ZN15QGuiApplication11focusWindowEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "focusWindow", args)
  }

  return
}

// isRightToLeft()
func (this *QGuiApplication) IsRightToLeft_s(args ...interface{}) (ret interface{}) {
  // isRightToLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication13isRightToLeftEv
    // invoke: bool isRightToLeft()
    var ret0 = C.C_ZN15QGuiApplication13isRightToLeftEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "isRightToLeft", args)
  }

  return
}

// topLevelAt(const class QPoint &)
func (this *QGuiApplication) TopLevelAt_s(args ...interface{}) (ret interface{}) {
  // topLevelAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication10topLevelAtERK6QPoint
    // invoke: QWindow * topLevelAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN15QGuiApplication10topLevelAtERK6QPoint(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "topLevelAt", args)
  }

  return
}

// applicationDisplayName()
func (this *QGuiApplication) ApplicationDisplayName_s(args ...interface{}) (ret interface{}) {
  // applicationDisplayName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication22applicationDisplayNameEv
    // invoke: QString applicationDisplayName()
    var ret0 = C.C_ZN15QGuiApplication22applicationDisplayNameEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "applicationDisplayName", args)
  }

  return
}

// inputMethod()
func (this *QGuiApplication) InputMethod_s(args ...interface{}) (ret interface{}) {
  // inputMethod()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication11inputMethodEv
    // invoke: QInputMethod * inputMethod()
    var ret0 = C.C_ZN15QGuiApplication11inputMethodEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QInputMethod{}) // "QInputMethod *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "inputMethod", args)
  }

  return
}

// clipboard()
func (this *QGuiApplication) Clipboard_s(args ...interface{}) (ret interface{}) {
  // clipboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication9clipboardEv
    // invoke: QClipboard * clipboard()
    var ret0 = C.C_ZN15QGuiApplication9clipboardEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QClipboard{}) // "QClipboard *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "clipboard", args)
  }

  return
}

// overrideCursor()
func (this *QGuiApplication) OverrideCursor_s(args ...interface{}) (ret interface{}) {
  // overrideCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication14overrideCursorEv
    // invoke: QCursor * overrideCursor()
    var ret0 = C.C_ZN15QGuiApplication14overrideCursorEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCursor{}) // "QCursor *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "overrideCursor", args)
  }

  return
}

// primaryScreen()
func (this *QGuiApplication) PrimaryScreen_s(args ...interface{}) (ret interface{}) {
  // primaryScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication13primaryScreenEv
    // invoke: QScreen * primaryScreen()
    var ret0 = C.C_ZN15QGuiApplication13primaryScreenEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScreen{}) // "QScreen *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "primaryScreen", args)
  }

  return
}

// allWindows()
func (this *QGuiApplication) AllWindows_s(args ...interface{}) () {
  // allWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication10allWindowsEv
    // invoke: QWindowList allWindows()
    C.C_ZN15QGuiApplication10allWindowsEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "allWindows", args)
  }

  return
}

// ~QGuiApplication()
func (this *QGuiApplication) Free(args ...interface{}) () {
  // ~QGuiApplication()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplicationD0Ev
    // invoke: void ~QGuiApplication()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN15QGuiApplicationD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "~QGuiApplication", args)
  }

  return
}

// setDesktopSettingsAware(_Bool)
func (this *QGuiApplication) SetDesktopSettingsAware_s(args ...interface{}) () {
  // setDesktopSettingsAware(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication23setDesktopSettingsAwareEb
    // invoke: void setDesktopSettingsAware(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication23setDesktopSettingsAwareEb(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setDesktopSettingsAware", args)
  }

  return
}

// mouseButtons()
func (this *QGuiApplication) MouseButtons_s(args ...interface{}) () {
  // mouseButtons()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication12mouseButtonsEv
    // invoke: Qt::MouseButtons mouseButtons()
    C.C_ZN15QGuiApplication12mouseButtonsEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "mouseButtons", args)
  }

  return
}

// platformNativeInterface()
func (this *QGuiApplication) PlatformNativeInterface_s(args ...interface{}) () {
  // platformNativeInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication23platformNativeInterfaceEv
    // invoke: QPlatformNativeInterface * platformNativeInterface()
    C.C_ZN15QGuiApplication23platformNativeInterfaceEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformNativeInterface", args)
  }

  return
}

// metaObject()
func (this *QGuiApplication) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK15QGuiApplication10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "metaObject", args)
  }

  return
}

// devicePixelRatio()
func (this *QGuiApplication) DevicePixelRatio(args ...interface{}) (ret interface{}) {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication16devicePixelRatioEv
    // invoke: qreal devicePixelRatio()
    var ret0 = C.C_ZNK15QGuiApplication16devicePixelRatioEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "devicePixelRatio", args)
  }

  return
}

// modalWindow()
func (this *QGuiApplication) ModalWindow_s(args ...interface{}) (ret interface{}) {
  // modalWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication11modalWindowEv
    // invoke: QWindow * modalWindow()
    var ret0 = C.C_ZN15QGuiApplication11modalWindowEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "modalWindow", args)
  }

  return
}

// QGuiApplication(int &, char **, int)
func GcfreeQGuiApplication(this *QGuiApplication) {
  qtrt.UniverseFree(this)
}
func NewQGuiApplication(args ...interface{}) *QGuiApplication {
  // QGuiApplication(int &, char **, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int &"
  vtys[0][1] = qtrt.ByteTy(true) // "char **"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplicationC1ERiPPci
    // invoke: void QGuiApplication(int &, char **, int)
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
    qthis = C.C_ZN15QGuiApplicationC2ERiPPci(arg0, arg1, arg2)
    this := &QGuiApplication{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQGuiApplication)
    return this
  default:
    qtrt.ErrorResolve("QGuiApplication", "QGuiApplication", args)
  }

  return nil // QGuiApplication{Qclsinst:qthis}
}

// isLeftToRight()
func (this *QGuiApplication) IsLeftToRight_s(args ...interface{}) (ret interface{}) {
  // isLeftToRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication13isLeftToRightEv
    // invoke: bool isLeftToRight()
    var ret0 = C.C_ZN15QGuiApplication13isLeftToRightEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "isLeftToRight", args)
  }

  return
}

// setWindowIcon(const class QIcon &)
func (this *QGuiApplication) SetWindowIcon_s(args ...interface{}) () {
  // setWindowIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication13setWindowIconERK5QIcon
    // invoke: void setWindowIcon(const class QIcon &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication13setWindowIconERK5QIcon(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setWindowIcon", args)
  }

  return
}

// setQuitOnLastWindowClosed(_Bool)
func (this *QGuiApplication) SetQuitOnLastWindowClosed_s(args ...interface{}) () {
  // setQuitOnLastWindowClosed(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication25setQuitOnLastWindowClosedEb
    // invoke: void setQuitOnLastWindowClosed(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication25setQuitOnLastWindowClosedEb(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setQuitOnLastWindowClosed", args)
  }

  return
}

// desktopSettingsAware()
func (this *QGuiApplication) DesktopSettingsAware_s(args ...interface{}) (ret interface{}) {
  // desktopSettingsAware()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication20desktopSettingsAwareEv
    // invoke: bool desktopSettingsAware()
    var ret0 = C.C_ZN15QGuiApplication20desktopSettingsAwareEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "desktopSettingsAware", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QGuiApplication) SetFont_s(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QGuiApplication7setFontERK5QFont(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setFont", args)
  }

  return
}

// topLevelWindows()
func (this *QGuiApplication) TopLevelWindows_s(args ...interface{}) () {
  // topLevelWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication15topLevelWindowsEv
    // invoke: QWindowList topLevelWindows()
    C.C_ZN15QGuiApplication15topLevelWindowsEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "topLevelWindows", args)
  }

  return
}

// isSessionRestored()
func (this *QGuiApplication) IsSessionRestored(args ...interface{}) (ret interface{}) {
  // isSessionRestored()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QGuiApplication17isSessionRestoredEv
    // invoke: bool isSessionRestored()
    var ret0 = C.C_ZNK15QGuiApplication17isSessionRestoredEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "isSessionRestored", args)
  }

  return
}

// restoreOverrideCursor()
func (this *QGuiApplication) RestoreOverrideCursor_s(args ...interface{}) () {
  // restoreOverrideCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication21restoreOverrideCursorEv
    // invoke: void restoreOverrideCursor()
    C.C_ZN15QGuiApplication21restoreOverrideCursorEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "restoreOverrideCursor", args)
  }

  return
}

// windowIcon()
func (this *QGuiApplication) WindowIcon_s(args ...interface{}) (ret interface{}) {
  // windowIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication10windowIconEv
    // invoke: QIcon windowIcon()
    var ret0 = C.C_ZN15QGuiApplication10windowIconEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGuiApplication", "windowIcon", args)
  }

  return
}

// <= body block end

