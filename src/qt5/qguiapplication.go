package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QStyleHints * QGuiApplication::styleHints();
extern void _ZN15QGuiApplication10styleHintsEv(); // 4
  // proto: static void QGuiApplication::setOverrideCursor(const QCursor & );
extern void _ZN15QGuiApplication17setOverrideCursorERK7QCursor(void* arg0); // 4
  // proto: static void QGuiApplication::sync();
extern void _ZN15QGuiApplication4syncEv(); // 4
  // proto: static void QGuiApplication::setPalette(const QPalette & pal);
extern void _ZN15QGuiApplication10setPaletteERK8QPalette(void* arg0); // 4
  // proto: static QList<QScreen *> QGuiApplication::screens();
extern void _ZN15QGuiApplication7screensEv(); // 4
  // proto: static Qt::ApplicationState QGuiApplication::applicationState();
extern void _ZN15QGuiApplication16applicationStateEv(); // 4
  // proto:  bool QGuiApplication::notify(QObject * , QEvent * );
extern void _ZN15QGuiApplication6notifyEP7QObjectP6QEvent(void* qthis, void* arg0, void* arg1); // 4
  // proto: static Qt::KeyboardModifiers QGuiApplication::queryKeyboardModifiers();
extern void _ZN15QGuiApplication22queryKeyboardModifiersEv(); // 4
  // proto:  bool QGuiApplication::isSavingSession();
extern void _ZNK15QGuiApplication15isSavingSessionEv(void* qthis); // 4
  // proto: static void QGuiApplication::changeOverrideCursor(const QCursor & );
extern void _ZN15QGuiApplication20changeOverrideCursorERK7QCursor(void* arg0); // 4
  // proto: static QFont QGuiApplication::font();
extern void _ZN15QGuiApplication4fontEv(); // 4
  // proto: static bool QGuiApplication::quitOnLastWindowClosed();
extern void _ZN15QGuiApplication22quitOnLastWindowClosedEv(); // 4
  // proto: static QPalette QGuiApplication::palette();
extern void _ZN15QGuiApplication7paletteEv(); // 4
  // proto: static void QGuiApplication::setApplicationDisplayName(const QString & name);
extern void _ZN15QGuiApplication25setApplicationDisplayNameERK7QString(void* arg0); // 4
  // proto: static Qt::LayoutDirection QGuiApplication::layoutDirection();
extern void _ZN15QGuiApplication15layoutDirectionEv(); // 4
  // proto: static int QGuiApplication::exec();
extern void _ZN15QGuiApplication4execEv(); // 4
  // proto:  QString QGuiApplication::sessionId();
extern void _ZNK15QGuiApplication9sessionIdEv(void* qthis); // 4
  // proto:  QString QGuiApplication::sessionKey();
extern void _ZNK15QGuiApplication10sessionKeyEv(void* qthis); // 4
  // proto: static QObject * QGuiApplication::focusObject();
extern void _ZN15QGuiApplication11focusObjectEv(); // 4
  // proto: static QFunctionPointer QGuiApplication::platformFunction(const QByteArray & function);
extern void _ZN15QGuiApplication16platformFunctionERK10QByteArray(void* arg0); // 4
  // proto: static Qt::KeyboardModifiers QGuiApplication::keyboardModifiers();
extern void _ZN15QGuiApplication17keyboardModifiersEv(); // 4
  // proto: static QString QGuiApplication::platformName();
extern void _ZN15QGuiApplication12platformNameEv(); // 4
  // proto: static QWindow * QGuiApplication::focusWindow();
extern void _ZN15QGuiApplication11focusWindowEv(); // 4
  // proto: static bool QGuiApplication::isRightToLeft();
extern void _ZN15QGuiApplication13isRightToLeftEv(); // 2
  // proto: static QWindow * QGuiApplication::topLevelAt(const QPoint & pos);
extern void _ZN15QGuiApplication10topLevelAtERK6QPoint(void* arg0); // 4
  // proto: static QString QGuiApplication::applicationDisplayName();
extern void _ZN15QGuiApplication22applicationDisplayNameEv(); // 4
  // proto: static QInputMethod * QGuiApplication::inputMethod();
extern void _ZN15QGuiApplication11inputMethodEv(); // 4
  // proto: static QClipboard * QGuiApplication::clipboard();
extern void _ZN15QGuiApplication9clipboardEv(); // 4
  // proto: static QCursor * QGuiApplication::overrideCursor();
extern void _ZN15QGuiApplication14overrideCursorEv(); // 4
  // proto: static QScreen * QGuiApplication::primaryScreen();
extern void _ZN15QGuiApplication13primaryScreenEv(); // 4
  // proto: static QWindowList QGuiApplication::allWindows();
extern void _ZN15QGuiApplication10allWindowsEv(); // 4
  // proto:  void QGuiApplication::~QGuiApplication();
extern void _ZN15QGuiApplicationD2Ev(void* qthis); // 4
  // proto: static void QGuiApplication::setDesktopSettingsAware(bool on);
extern void _ZN15QGuiApplication23setDesktopSettingsAwareEb(bool arg0); // 4
  // proto: static Qt::MouseButtons QGuiApplication::mouseButtons();
extern void _ZN15QGuiApplication12mouseButtonsEv(); // 4
  // proto: static QPlatformNativeInterface * QGuiApplication::platformNativeInterface();
extern void _ZN15QGuiApplication23platformNativeInterfaceEv(); // 4
  // proto:  const QMetaObject * QGuiApplication::metaObject();
extern void _ZNK15QGuiApplication10metaObjectEv(void* qthis); // 4
  // proto:  qreal QGuiApplication::devicePixelRatio();
extern void _ZNK15QGuiApplication16devicePixelRatioEv(void* qthis); // 4
  // proto: static QWindow * QGuiApplication::modalWindow();
extern void _ZN15QGuiApplication11modalWindowEv(); // 4
  // proto:  void QGuiApplication::QGuiApplication(int & argc, char ** argv, int );
extern void _ZN15QGuiApplicationC2ERiPPci(void* qthis, int32_t* arg0, unsigned char* arg1, int32_t arg2); // 3
  // proto: static bool QGuiApplication::isLeftToRight();
extern void _ZN15QGuiApplication13isLeftToRightEv(); // 2
  // proto: static void QGuiApplication::setWindowIcon(const QIcon & icon);
extern void _ZN15QGuiApplication13setWindowIconERK5QIcon(void* arg0); // 4
  // proto: static void QGuiApplication::setQuitOnLastWindowClosed(bool quit);
extern void _ZN15QGuiApplication25setQuitOnLastWindowClosedEb(bool arg0); // 4
  // proto: static bool QGuiApplication::desktopSettingsAware();
extern void _ZN15QGuiApplication20desktopSettingsAwareEv(); // 4
  // proto: static void QGuiApplication::setFont(const QFont & );
extern void _ZN15QGuiApplication7setFontERK5QFont(void* arg0); // 4
  // proto: static QWindowList QGuiApplication::topLevelWindows();
extern void _ZN15QGuiApplication15topLevelWindowsEv(); // 4
  // proto:  bool QGuiApplication::isSessionRestored();
extern void _ZNK15QGuiApplication17isSessionRestoredEv(void* qthis); // 4
  // proto: static void QGuiApplication::restoreOverrideCursor();
extern void _ZN15QGuiApplication21restoreOverrideCursorEv(); // 4
  // proto: static QIcon QGuiApplication::windowIcon();
extern void _ZN15QGuiApplication10windowIconEv(); // 4
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

// class sizeof(QGuiApplication)=1
type QGuiApplication struct {
  /*qbase*/ QCoreApplication;
  qclsinst unsafe.Pointer /* *C.void */;
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
func (this *QGuiApplication) styleHints_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication10styleHintsEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "styleHints", args)
  }

}

// setOverrideCursor(const class QCursor &)
func (this *QGuiApplication) setOverrideCursor_s(args ...interface{}) () {
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
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication17setOverrideCursorERK7QCursor(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setOverrideCursor", args)
  }

}

// sync()
func (this *QGuiApplication) sync_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication4syncEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "sync", args)
  }

}

// setPalette(const class QPalette &)
func (this *QGuiApplication) setPalette_s(args ...interface{}) () {
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
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication10setPaletteERK8QPalette(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setPalette", args)
  }

}

// screens()
func (this *QGuiApplication) screens_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication7screensEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "screens", args)
  }

}

// applicationState()
func (this *QGuiApplication) applicationState_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication16applicationStateEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "applicationState", args)
  }

}

// notify(class QObject *, class QEvent *)
func (this *QGuiApplication) notify(args ...interface{}) () {
  // notify(class QObject *, class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication6notifyEP7QObjectP6QEvent
    // invoke: bool notify(class QObject *, class QEvent *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QEvent).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN15QGuiApplication6notifyEP7QObjectP6QEvent(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QGuiApplication", "notify", args)
  }

}

// queryKeyboardModifiers()
func (this *QGuiApplication) queryKeyboardModifiers_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication22queryKeyboardModifiersEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "queryKeyboardModifiers", args)
  }

}

// isSavingSession()
func (this *QGuiApplication) isSavingSession(args ...interface{}) () {
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
    C._ZNK15QGuiApplication15isSavingSessionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "isSavingSession", args)
  }

}

// changeOverrideCursor(const class QCursor &)
func (this *QGuiApplication) changeOverrideCursor_s(args ...interface{}) () {
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
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication20changeOverrideCursorERK7QCursor(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "changeOverrideCursor", args)
  }

}

// font()
func (this *QGuiApplication) font_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication4fontEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "font", args)
  }

}

// quitOnLastWindowClosed()
func (this *QGuiApplication) quitOnLastWindowClosed_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication22quitOnLastWindowClosedEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "quitOnLastWindowClosed", args)
  }

}

// palette()
func (this *QGuiApplication) palette_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication7paletteEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "palette", args)
  }

}

// setApplicationDisplayName(const class QString &)
func (this *QGuiApplication) setApplicationDisplayName_s(args ...interface{}) () {
  // setApplicationDisplayName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication25setApplicationDisplayNameERK7QString
    // invoke: void setApplicationDisplayName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication25setApplicationDisplayNameERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setApplicationDisplayName", args)
  }

}

// layoutDirection()
func (this *QGuiApplication) layoutDirection_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication15layoutDirectionEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "layoutDirection", args)
  }

}

// exec()
func (this *QGuiApplication) exec_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication4execEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "exec", args)
  }

}

// sessionId()
func (this *QGuiApplication) sessionId(args ...interface{}) () {
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
    C._ZNK15QGuiApplication9sessionIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "sessionId", args)
  }

}

// sessionKey()
func (this *QGuiApplication) sessionKey(args ...interface{}) () {
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
    C._ZNK15QGuiApplication10sessionKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "sessionKey", args)
  }

}

// focusObject()
func (this *QGuiApplication) focusObject_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication11focusObjectEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "focusObject", args)
  }

}

// platformFunction(const class QByteArray &)
func (this *QGuiApplication) platformFunction_s(args ...interface{}) () {
  // platformFunction(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication16platformFunctionERK10QByteArray
    // invoke: QFunctionPointer platformFunction(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication16platformFunctionERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformFunction", args)
  }

}

// keyboardModifiers()
func (this *QGuiApplication) keyboardModifiers_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication17keyboardModifiersEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "keyboardModifiers", args)
  }

}

// platformName()
func (this *QGuiApplication) platformName_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication12platformNameEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformName", args)
  }

}

// focusWindow()
func (this *QGuiApplication) focusWindow_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication11focusWindowEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "focusWindow", args)
  }

}

// isRightToLeft()
func (this *QGuiApplication) isRightToLeft_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication13isRightToLeftEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "isRightToLeft", args)
  }

}

// topLevelAt(const class QPoint &)
func (this *QGuiApplication) topLevelAt_s(args ...interface{}) () {
  // topLevelAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QGuiApplication10topLevelAtERK6QPoint
    // invoke: QWindow * topLevelAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication10topLevelAtERK6QPoint(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "topLevelAt", args)
  }

}

// applicationDisplayName()
func (this *QGuiApplication) applicationDisplayName_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication22applicationDisplayNameEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "applicationDisplayName", args)
  }

}

// inputMethod()
func (this *QGuiApplication) inputMethod_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication11inputMethodEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "inputMethod", args)
  }

}

// clipboard()
func (this *QGuiApplication) clipboard_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication9clipboardEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "clipboard", args)
  }

}

// overrideCursor()
func (this *QGuiApplication) overrideCursor_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication14overrideCursorEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "overrideCursor", args)
  }

}

// primaryScreen()
func (this *QGuiApplication) primaryScreen_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication13primaryScreenEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "primaryScreen", args)
  }

}

// allWindows()
func (this *QGuiApplication) allWindows_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication10allWindowsEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "allWindows", args)
  }

}

// ~QGuiApplication()
func (this *QGuiApplication) FreeQGuiApplication(args ...interface{}) () {
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
    C._ZN15QGuiApplicationD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "~QGuiApplication", args)
  }

}

// setDesktopSettingsAware(_Bool)
func (this *QGuiApplication) setDesktopSettingsAware_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication23setDesktopSettingsAwareEb(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setDesktopSettingsAware", args)
  }

}

// mouseButtons()
func (this *QGuiApplication) mouseButtons_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication12mouseButtonsEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "mouseButtons", args)
  }

}

// platformNativeInterface()
func (this *QGuiApplication) platformNativeInterface_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication23platformNativeInterfaceEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "platformNativeInterface", args)
  }

}

// metaObject()
func (this *QGuiApplication) metaObject(args ...interface{}) () {
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
    C._ZNK15QGuiApplication10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "metaObject", args)
  }

}

// devicePixelRatio()
func (this *QGuiApplication) devicePixelRatio(args ...interface{}) () {
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
    C._ZNK15QGuiApplication16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "devicePixelRatio", args)
  }

}

// modalWindow()
func (this *QGuiApplication) modalWindow_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication11modalWindowEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "modalWindow", args)
  }

}

// QGuiApplication(int &, char **, int)
func NewQGuiApplication(args ...interface{}) QGuiApplication {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QGuiApplicationC2ERiPPci(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QGuiApplication", "QGuiApplication", args)
  }

  return QGuiApplication{}
}

// isLeftToRight()
func (this *QGuiApplication) isLeftToRight_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication13isLeftToRightEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "isLeftToRight", args)
  }

}

// setWindowIcon(const class QIcon &)
func (this *QGuiApplication) setWindowIcon_s(args ...interface{}) () {
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
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication13setWindowIconERK5QIcon(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setWindowIcon", args)
  }

}

// setQuitOnLastWindowClosed(_Bool)
func (this *QGuiApplication) setQuitOnLastWindowClosed_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication25setQuitOnLastWindowClosedEb(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setQuitOnLastWindowClosed", args)
  }

}

// desktopSettingsAware()
func (this *QGuiApplication) desktopSettingsAware_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication20desktopSettingsAwareEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "desktopSettingsAware", args)
  }

}

// setFont(const class QFont &)
func (this *QGuiApplication) setFont_s(args ...interface{}) () {
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
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QGuiApplication7setFontERK5QFont(arg0)
  default:
    qtrt.ErrorResolve("QGuiApplication", "setFont", args)
  }

}

// topLevelWindows()
func (this *QGuiApplication) topLevelWindows_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication15topLevelWindowsEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "topLevelWindows", args)
  }

}

// isSessionRestored()
func (this *QGuiApplication) isSessionRestored(args ...interface{}) () {
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
    C._ZNK15QGuiApplication17isSessionRestoredEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QGuiApplication", "isSessionRestored", args)
  }

}

// restoreOverrideCursor()
func (this *QGuiApplication) restoreOverrideCursor_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication21restoreOverrideCursorEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "restoreOverrideCursor", args)
  }

}

// windowIcon()
func (this *QGuiApplication) windowIcon_s(args ...interface{}) () {
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
    C._ZN15QGuiApplication10windowIconEv()
  default:
    qtrt.ErrorResolve("QGuiApplication", "windowIcon", args)
  }

}

// <= body block end

