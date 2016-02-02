package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qwidget.h
// dst-file: /src/widgets/qwidget.go
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
  // proto:  void QWidget::setWindowIconText(const QString & );
extern void C_ZN7QWidget17setWindowIconTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::show();
extern void C_ZN7QWidget4showEv(void* qthis); // 4
  // proto:  bool QWidget::hasFocus();
extern bool C_ZNK7QWidget8hasFocusEv(void* qthis); // 4
  // proto:  void QWidget::insertAction(QAction * before, QAction * action);
extern void C_ZN7QWidget12insertActionEP7QActionS1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QWidget::move(int x, int y);
extern void C_ZN7QWidget4moveEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QWidget::move(const QPoint & );
extern void C_ZN7QWidget4moveERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QList<QAction *> QWidget::actions();
extern void C_ZNK7QWidget7actionsEv(void* qthis); // 4
  // proto:  void QWidget::setStyle(QStyle * );
extern void C_ZN7QWidget8setStyleEP6QStyle(void* qthis, void* arg0); // 4
  // proto: static QWidget * QWidget::keyboardGrabber();
extern void* C_ZN7QWidget15keyboardGrabberEv(); // 4
  // proto:  void QWidget::setHidden(bool hidden);
extern void C_ZN7QWidget9setHiddenEb(void* qthis, bool arg0); // 4
  // proto:  QWidget * QWidget::childAt(int x, int y);
extern void* C_ZNK7QWidget7childAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QWidget * QWidget::childAt(const QPoint & p);
extern void* C_ZNK7QWidget7childAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QWidget * QWidget::focusWidget();
extern void* C_ZNK7QWidget11focusWidgetEv(void* qthis); // 4
  // proto:  int QWidget::minimumHeight();
extern int32_t C_ZNK7QWidget13minimumHeightEv(void* qthis); // 2
  // proto:  const QFont & QWidget::font();
extern void* C_ZNK7QWidget4fontEv(void* qthis); // 2
  // proto:  void QWidget::setLayout(QLayout * );
extern void C_ZN7QWidget9setLayoutEP7QLayout(void* qthis, void* arg0); // 4
  // proto:  QLayout * QWidget::layout();
extern void* C_ZNK7QWidget6layoutEv(void* qthis); // 4
  // proto:  void QWidget::activateWindow();
extern void C_ZN7QWidget14activateWindowEv(void* qthis); // 4
  // proto:  void QWidget::unsetLayoutDirection();
extern void C_ZN7QWidget20unsetLayoutDirectionEv(void* qthis); // 4
  // proto:  bool QWidget::isEnabled();
extern bool C_ZNK7QWidget9isEnabledEv(void* qthis); // 2
  // proto:  Qt::ContextMenuPolicy QWidget::contextMenuPolicy();
extern void C_ZNK7QWidget17contextMenuPolicyEv(void* qthis); // 4
  // proto:  void QWidget::setWindowFilePath(const QString & filePath);
extern void C_ZN7QWidget17setWindowFilePathERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QWidget::accessibleName();
extern void* C_ZNK7QWidget14accessibleNameEv(void* qthis); // 4
  // proto:  QWidget * QWidget::window();
extern void* C_ZNK7QWidget6windowEv(void* qthis); // 4
  // proto:  WId QWidget::internalWinId();
extern void C_ZNK7QWidget13internalWinIdEv(void* qthis); // 2
  // proto:  void QWidget::setWindowTitle(const QString & );
extern void C_ZN7QWidget14setWindowTitleERK7QString(void* qthis, void* arg0); // 4
  // proto: static QWidget * QWidget::mouseGrabber();
extern void* C_ZN7QWidget12mouseGrabberEv(); // 4
  // proto:  bool QWidget::isModal();
extern bool C_ZNK7QWidget7isModalEv(void* qthis); // 2
  // proto:  QPalette::ColorRole QWidget::foregroundRole();
extern void C_ZNK7QWidget14foregroundRoleEv(void* qthis); // 4
  // proto:  void QWidget::setShortcutAutoRepeat(int id, bool enable);
extern void C_ZN7QWidget21setShortcutAutoRepeatEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QWidget::setGraphicsEffect(QGraphicsEffect * effect);
extern void C_ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setAccessibleName(const QString & name);
extern void C_ZN7QWidget17setAccessibleNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMinimumHeight(int minh);
extern void C_ZN7QWidget16setMinimumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  QPoint QWidget::mapToParent(const QPoint & );
extern void* C_ZNK7QWidget11mapToParentERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setGeometry(int x, int y, int w, int h);
extern void C_ZN7QWidget11setGeometryEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QWidget::setGeometry(const QRect & );
extern void C_ZN7QWidget11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QPoint QWidget::pos();
extern void* C_ZNK7QWidget3posEv(void* qthis); // 4
  // proto:  QSize QWidget::sizeHint();
extern void* C_ZNK7QWidget8sizeHintEv(void* qthis); // 4
  // proto:  Qt::WindowType QWidget::windowType();
extern void C_ZNK7QWidget10windowTypeEv(void* qthis); // 2
  // proto:  void QWidget::removeAction(QAction * action);
extern void C_ZN7QWidget12removeActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  void QWidget::grabMouse();
extern void C_ZN7QWidget9grabMouseEv(void* qthis); // 4
  // proto:  void QWidget::grabMouse(const QCursor & );
extern void C_ZN7QWidget9grabMouseERK7QCursor(void* qthis, void* arg0); // 4
  // proto:  QRect QWidget::frameGeometry();
extern void* C_ZNK7QWidget13frameGeometryEv(void* qthis); // 4
  // proto:  void QWidget::setWindowRole(const QString & );
extern void C_ZN7QWidget13setWindowRoleERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::repaint(const QRect & );
extern void C_ZN7QWidget7repaintERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QWidget::repaint();
extern void C_ZN7QWidget7repaintEv(void* qthis); // 4
  // proto:  void QWidget::repaint(int x, int y, int w, int h);
extern void C_ZN7QWidget7repaintEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QWidget::repaint(const QRegion & );
extern void C_ZN7QWidget7repaintERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  Qt::InputMethodHints QWidget::inputMethodHints();
extern void C_ZNK7QWidget16inputMethodHintsEv(void* qthis); // 4
  // proto:  const QMetaObject * QWidget::metaObject();
extern void C_ZNK7QWidget10metaObjectEv(void* qthis); // 4
  // proto:  QWidget * QWidget::topLevelWidget();
extern void* C_ZNK7QWidget14topLevelWidgetEv(void* qthis); // 2
  // proto:  void QWidget::releaseMouse();
extern void C_ZN7QWidget12releaseMouseEv(void* qthis); // 4
  // proto:  void QWidget::setLocale(const QLocale & locale);
extern void C_ZN7QWidget9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  QRect QWidget::contentsRect();
extern void* C_ZNK7QWidget12contentsRectEv(void* qthis); // 4
  // proto:  Qt::FocusPolicy QWidget::focusPolicy();
extern void C_ZNK7QWidget11focusPolicyEv(void* qthis); // 4
  // proto:  bool QWidget::isVisibleTo(const QWidget * );
extern bool C_ZNK7QWidget11isVisibleToEPKS_(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::updatesEnabled();
extern bool C_ZNK7QWidget14updatesEnabledEv(void* qthis); // 2
  // proto:  Qt::WindowStates QWidget::windowState();
extern void C_ZNK7QWidget11windowStateEv(void* qthis); // 4
  // proto:  bool QWidget::isWindowModified();
extern bool C_ZNK7QWidget16isWindowModifiedEv(void* qthis); // 4
  // proto:  int QWidget::devType();
extern int32_t C_ZNK7QWidget7devTypeEv(void* qthis); // 4
  // proto:  void QWidget::clearFocus();
extern void C_ZN7QWidget10clearFocusEv(void* qthis); // 4
  // proto:  void QWidget::setStyleSheet(const QString & styleSheet);
extern void C_ZN7QWidget13setStyleSheetERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMaximumHeight(int maxh);
extern void C_ZN7QWidget16setMaximumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWidget::x();
extern void C_ZNK7QWidget1xEv(void* qthis); // 4
  // proto:  QPixmap QWidget::grab(const QRect & rectangle);
extern void* C_ZN7QWidget4grabERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QSize QWidget::maximumSize();
extern void* C_ZNK7QWidget11maximumSizeEv(void* qthis); // 4
  // proto:  QLocale QWidget::locale();
extern void* C_ZNK7QWidget6localeEv(void* qthis); // 4
  // proto:  QSize QWidget::minimumSize();
extern void* C_ZNK7QWidget11minimumSizeEv(void* qthis); // 4
  // proto:  QString QWidget::windowFilePath();
extern void* C_ZNK7QWidget14windowFilePathEv(void* qthis); // 4
  // proto:  void QWidget::setStatusTip(const QString & );
extern void C_ZN7QWidget12setStatusTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  QWidget * QWidget::focusProxy();
extern void* C_ZNK7QWidget10focusProxyEv(void* qthis); // 4
  // proto:  void QWidget::createWinId();
extern void C_ZN7QWidget11createWinIdEv(void* qthis); // 4
  // proto:  QPaintEngine * QWidget::paintEngine();
extern void* C_ZNK7QWidget11paintEngineEv(void* qthis); // 4
  // proto:  QByteArray QWidget::saveGeometry();
extern void* C_ZNK7QWidget12saveGeometryEv(void* qthis); // 4
  // proto:  QBackingStore * QWidget::backingStore();
extern void* C_ZNK7QWidget12backingStoreEv(void* qthis); // 4
  // proto:  void QWidget::setToolTipDuration(int msec);
extern void C_ZN7QWidget18setToolTipDurationEi(void* qthis, int32_t arg0); // 4
  // proto:  QCursor QWidget::cursor();
extern void* C_ZNK7QWidget6cursorEv(void* qthis); // 4
  // proto:  const QPalette & QWidget::palette();
extern void* C_ZNK7QWidget7paletteEv(void* qthis); // 4
  // proto:  void QWidget::hide();
extern void C_ZN7QWidget4hideEv(void* qthis); // 4
  // proto:  QString QWidget::windowTitle();
extern void* C_ZNK7QWidget11windowTitleEv(void* qthis); // 4
  // proto:  void QWidget::setFocusProxy(QWidget * );
extern void C_ZN7QWidget13setFocusProxyEPS_(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setParent(QWidget * parent);
extern void C_ZN7QWidget9setParentEPS_(void* qthis, void* arg0); // 4
  // proto:  QSizePolicy QWidget::sizePolicy();
extern void* C_ZNK7QWidget10sizePolicyEv(void* qthis); // 4
  // proto:  QRegion QWidget::visibleRegion();
extern void* C_ZNK7QWidget13visibleRegionEv(void* qthis); // 4
  // proto:  Qt::WindowFlags QWidget::windowFlags();
extern void C_ZNK7QWidget11windowFlagsEv(void* qthis); // 2
  // proto:  void QWidget::setContentsMargins(int left, int top, int right, int bottom);
extern void C_ZN7QWidget18setContentsMarginsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QWidget::setContentsMargins(const QMargins & margins);
extern void C_ZN7QWidget18setContentsMarginsERK8QMargins(void* qthis, void* arg0); // 4
  // proto:  QGraphicsEffect * QWidget::graphicsEffect();
extern void C_ZNK7QWidget14graphicsEffectEv(void* qthis); // 4
  // proto:  QPoint QWidget::mapFromParent(const QPoint & );
extern void* C_ZNK7QWidget13mapFromParentERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QString QWidget::windowRole();
extern void* C_ZNK7QWidget10windowRoleEv(void* qthis); // 4
  // proto:  QStyle * QWidget::style();
extern void* C_ZNK7QWidget5styleEv(void* qthis); // 4
  // proto:  QString QWidget::toolTip();
extern void* C_ZNK7QWidget7toolTipEv(void* qthis); // 4
  // proto:  int QWidget::maximumWidth();
extern int32_t C_ZNK7QWidget12maximumWidthEv(void* qthis); // 2
  // proto:  bool QWidget::isEnabledToTLW();
extern bool C_ZNK7QWidget14isEnabledToTLWEv(void* qthis); // 2
  // proto:  QIcon QWidget::windowIcon();
extern void* C_ZNK7QWidget10windowIconEv(void* qthis); // 4
  // proto:  bool QWidget::isMinimized();
extern bool C_ZNK7QWidget11isMinimizedEv(void* qthis); // 4
  // proto:  QRect QWidget::rect();
extern void* C_ZNK7QWidget4rectEv(void* qthis); // 2
  // proto:  void QWidget::raise();
extern void C_ZN7QWidget5raiseEv(void* qthis); // 4
  // proto:  void QWidget::stackUnder(QWidget * );
extern void C_ZN7QWidget10stackUnderEPS_(void* qthis, void* arg0); // 4
  // proto:  QWidget * QWidget::parentWidget();
extern void* C_ZNK7QWidget12parentWidgetEv(void* qthis); // 2
  // proto:  WId QWidget::effectiveWinId();
extern void C_ZNK7QWidget14effectiveWinIdEv(void* qthis); // 4
  // proto:  void QWidget::setToolTip(const QString & );
extern void C_ZN7QWidget10setToolTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setSizePolicy(QSizePolicy );
extern void C_ZN7QWidget13setSizePolicyE11QSizePolicy(void* qthis, void* arg0); // 4
  // proto:  const QRect & QWidget::geometry();
extern void* C_ZNK7QWidget8geometryEv(void* qthis); // 2
  // proto:  QWindow * QWidget::windowHandle();
extern void* C_ZNK7QWidget12windowHandleEv(void* qthis); // 4
  // proto:  void QWidget::setAcceptDrops(bool on);
extern void C_ZN7QWidget14setAcceptDropsEb(void* qthis, bool arg0); // 4
  // proto:  bool QWidget::isEnabledTo(const QWidget * );
extern bool C_ZNK7QWidget11isEnabledToEPKS_(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isVisible();
extern bool C_ZNK7QWidget9isVisibleEv(void* qthis); // 2
  // proto:  void QWidget::setWindowModified(bool );
extern void C_ZN7QWidget17setWindowModifiedEb(void* qthis, bool arg0); // 4
  // proto:  QSize QWidget::size();
extern void* C_ZNK7QWidget4sizeEv(void* qthis); // 2
  // proto:  void QWidget::setMaximumWidth(int maxw);
extern void C_ZN7QWidget15setMaximumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::addAction(QAction * action);
extern void C_ZN7QWidget9addActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QRect QWidget::normalGeometry();
extern void* C_ZNK7QWidget14normalGeometryEv(void* qthis); // 4
  // proto:  void QWidget::setMinimumSize(int minw, int minh);
extern void C_ZN7QWidget14setMinimumSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::setMinimumSize(const QSize & );
extern void C_ZN7QWidget14setMinimumSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  int QWidget::heightForWidth(int );
extern int32_t C_ZNK7QWidget14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::setPalette(const QPalette & );
extern void C_ZN7QWidget10setPaletteERK8QPalette(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isAncestorOf(const QWidget * child);
extern bool C_ZNK7QWidget12isAncestorOfEPKS_(void* qthis, void* arg0); // 4
  // proto:  void QWidget::clearMask();
extern void C_ZN7QWidget9clearMaskEv(void* qthis); // 4
  // proto:  void QWidget::setWindowOpacity(qreal level);
extern void C_ZN7QWidget16setWindowOpacityEd(void* qthis, double arg0); // 4
  // proto:  bool QWidget::isActiveWindow();
extern bool C_ZNK7QWidget14isActiveWindowEv(void* qthis); // 4
  // proto:  void QWidget::setFont(const QFont & );
extern void C_ZN7QWidget7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setAccessibleDescription(const QString & description);
extern void C_ZN7QWidget24setAccessibleDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMaximumSize(const QSize & );
extern void C_ZN7QWidget14setMaximumSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QWidget::setMaximumSize(int maxw, int maxh);
extern void C_ZN7QWidget14setMaximumSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::unsetLocale();
extern void C_ZN7QWidget11unsetLocaleEv(void* qthis); // 4
  // proto:  bool QWidget::autoFillBackground();
extern bool C_ZNK7QWidget18autoFillBackgroundEv(void* qthis); // 4
  // proto:  void QWidget::unsetCursor();
extern void C_ZN7QWidget11unsetCursorEv(void* qthis); // 4
  // proto:  WId QWidget::winId();
extern void C_ZNK7QWidget5winIdEv(void* qthis); // 4
  // proto:  Qt::LayoutDirection QWidget::layoutDirection();
extern void C_ZNK7QWidget15layoutDirectionEv(void* qthis); // 4
  // proto:  void QWidget::setShortcutEnabled(int id, bool enable);
extern void C_ZN7QWidget18setShortcutEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  QSize QWidget::sizeIncrement();
extern void* C_ZNK7QWidget13sizeIncrementEv(void* qthis); // 4
  // proto:  void QWidget::setMouseTracking(bool enable);
extern void C_ZN7QWidget16setMouseTrackingEb(void* qthis, bool arg0); // 2
  // proto:  QString QWidget::whatsThis();
extern void* C_ZNK7QWidget9whatsThisEv(void* qthis); // 4
  // proto:  int QWidget::width();
extern int32_t C_ZNK7QWidget5widthEv(void* qthis); // 2
  // proto:  QRect QWidget::childrenRect();
extern void* C_ZNK7QWidget12childrenRectEv(void* qthis); // 4
  // proto:  QString QWidget::windowIconText();
extern void* C_ZNK7QWidget14windowIconTextEv(void* qthis); // 4
  // proto:  int QWidget::toolTipDuration();
extern int32_t C_ZNK7QWidget15toolTipDurationEv(void* qthis); // 4
  // proto:  void QWidget::setEnabled(bool );
extern void C_ZN7QWidget10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::showMaximized();
extern void C_ZN7QWidget13showMaximizedEv(void* qthis); // 4
  // proto:  void QWidget::ensurePolished();
extern void C_ZNK7QWidget14ensurePolishedEv(void* qthis); // 4
  // proto:  QString QWidget::statusTip();
extern void* C_ZNK7QWidget9statusTipEv(void* qthis); // 4
  // proto:  bool QWidget::acceptDrops();
extern bool C_ZNK7QWidget11acceptDropsEv(void* qthis); // 4
  // proto:  bool QWidget::isFullScreen();
extern bool C_ZNK7QWidget12isFullScreenEv(void* qthis); // 4
  // proto:  QWidget * QWidget::nextInFocusChain();
extern void* C_ZNK7QWidget16nextInFocusChainEv(void* qthis); // 4
  // proto:  QString QWidget::styleSheet();
extern void* C_ZNK7QWidget10styleSheetEv(void* qthis); // 4
  // proto:  void QWidget::setSizeIncrement(const QSize & );
extern void C_ZN7QWidget16setSizeIncrementERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QWidget::setSizeIncrement(int w, int h);
extern void C_ZN7QWidget16setSizeIncrementEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::grabKeyboard();
extern void C_ZN7QWidget12grabKeyboardEv(void* qthis); // 4
  // proto:  QPoint QWidget::mapTo(const QWidget * , const QPoint & );
extern void* C_ZNK7QWidget5mapToEPKS_RK6QPoint(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QWidget::resize(int w, int h);
extern void C_ZN7QWidget6resizeEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QWidget::resize(const QSize & );
extern void C_ZN7QWidget6resizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setFocus();
extern void C_ZN7QWidget8setFocusEv(void* qthis); // 2
  // proto:  QPoint QWidget::mapFromGlobal(const QPoint & );
extern void* C_ZNK7QWidget13mapFromGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QSize QWidget::frameSize();
extern void* C_ZNK7QWidget9frameSizeEv(void* qthis); // 4
  // proto:  void QWidget::releaseShortcut(int id);
extern void C_ZN7QWidget15releaseShortcutEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QWidget::minimumSizeHint();
extern void* C_ZNK7QWidget15minimumSizeHintEv(void* qthis); // 4
  // proto:  bool QWidget::isHidden();
extern bool C_ZNK7QWidget8isHiddenEv(void* qthis); // 2
  // proto:  bool QWidget::hasHeightForWidth();
extern bool C_ZNK7QWidget17hasHeightForWidthEv(void* qthis); // 4
  // proto:  void QWidget::setFixedSize(const QSize & );
extern void C_ZN7QWidget12setFixedSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setFixedSize(int w, int h);
extern void C_ZN7QWidget12setFixedSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto: static void QWidget::setTabOrder(QWidget * , QWidget * );
extern void C_ZN7QWidget11setTabOrderEPS_S0_(void* arg0, void* arg1); // 4
  // proto:  void QWidget::setWindowIcon(const QIcon & icon);
extern void C_ZN7QWidget13setWindowIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isLeftToRight();
extern bool C_ZNK7QWidget13isLeftToRightEv(void* qthis); // 2
  // proto:  bool QWidget::isRightToLeft();
extern bool C_ZNK7QWidget13isRightToLeftEv(void* qthis); // 2
  // proto:  void QWidget::setVisible(bool visible);
extern void C_ZN7QWidget10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::setCursor(const QCursor & );
extern void C_ZN7QWidget9setCursorERK7QCursor(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setFixedWidth(int w);
extern void C_ZN7QWidget13setFixedWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QPoint QWidget::mapFrom(const QWidget * , const QPoint & );
extern void* C_ZNK7QWidget7mapFromEPKS_RK6QPoint(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QWidget::isWindow();
extern bool C_ZNK7QWidget8isWindowEv(void* qthis); // 2
  // proto:  bool QWidget::close();
extern bool C_ZN7QWidget5closeEv(void* qthis); // 4
  // proto:  QFontMetrics QWidget::fontMetrics();
extern void* C_ZNK7QWidget11fontMetricsEv(void* qthis); // 2
  // proto:  QMargins QWidget::contentsMargins();
extern void* C_ZNK7QWidget15contentsMarginsEv(void* qthis); // 4
  // proto:  void QWidget::~QWidget();
extern void C_ZN7QWidgetD2Ev(void* qthis); // 4
  // proto:  void QWidget::setUpdatesEnabled(bool enable);
extern void C_ZN7QWidget17setUpdatesEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::showMinimized();
extern void C_ZN7QWidget13showMinimizedEv(void* qthis); // 4
  // proto:  QRegion QWidget::childrenRegion();
extern void* C_ZNK7QWidget14childrenRegionEv(void* qthis); // 4
  // proto:  Qt::WindowModality QWidget::windowModality();
extern void C_ZNK7QWidget14windowModalityEv(void* qthis); // 4
  // proto:  QFontInfo QWidget::fontInfo();
extern void* C_ZNK7QWidget8fontInfoEv(void* qthis); // 2
  // proto:  void QWidget::releaseKeyboard();
extern void C_ZN7QWidget15releaseKeyboardEv(void* qthis); // 4
  // proto:  QWidget * QWidget::previousInFocusChain();
extern void* C_ZNK7QWidget20previousInFocusChainEv(void* qthis); // 4
  // proto:  void QWidget::showFullScreen();
extern void C_ZN7QWidget14showFullScreenEv(void* qthis); // 4
  // proto:  QSize QWidget::baseSize();
extern void* C_ZNK7QWidget8baseSizeEv(void* qthis); // 4
  // proto:  bool QWidget::hasMouseTracking();
extern bool C_ZNK7QWidget16hasMouseTrackingEv(void* qthis); // 2
  // proto:  bool QWidget::restoreGeometry(const QByteArray & geometry);
extern bool C_ZN7QWidget15restoreGeometryERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setDisabled(bool );
extern void C_ZN7QWidget11setDisabledEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::showNormal();
extern void C_ZN7QWidget10showNormalEv(void* qthis); // 4
  // proto:  void QWidget::setFixedHeight(int h);
extern void C_ZN7QWidget14setFixedHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::setBaseSize(int basew, int baseh);
extern void C_ZN7QWidget11setBaseSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::setBaseSize(const QSize & );
extern void C_ZN7QWidget11setBaseSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QWidget::setAutoFillBackground(bool enabled);
extern void C_ZN7QWidget21setAutoFillBackgroundEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::updateGeometry();
extern void C_ZN7QWidget14updateGeometryEv(void* qthis); // 4
  // proto:  int QWidget::minimumWidth();
extern int32_t C_ZNK7QWidget12minimumWidthEv(void* qthis); // 2
  // proto:  int QWidget::maximumHeight();
extern int32_t C_ZNK7QWidget13maximumHeightEv(void* qthis); // 2
  // proto:  void QWidget::update();
extern void C_ZN7QWidget6updateEv(void* qthis); // 4
  // proto:  void QWidget::update(int x, int y, int w, int h);
extern void C_ZN7QWidget6updateEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QWidget::update(const QRect & );
extern void C_ZN7QWidget6updateERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QWidget::update(const QRegion & );
extern void C_ZN7QWidget6updateERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QWidget::adjustSize();
extern void C_ZN7QWidget10adjustSizeEv(void* qthis); // 4
  // proto:  QGraphicsProxyWidget * QWidget::graphicsProxyWidget();
extern void C_ZNK7QWidget19graphicsProxyWidgetEv(void* qthis); // 4
  // proto:  bool QWidget::isTopLevel();
extern bool C_ZNK7QWidget10isTopLevelEv(void* qthis); // 2
  // proto:  void QWidget::setMask(const QRegion & );
extern void C_ZN7QWidget7setMaskERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMask(const QBitmap & );
extern void C_ZN7QWidget7setMaskERK7QBitmap(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isMaximized();
extern bool C_ZNK7QWidget11isMaximizedEv(void* qthis); // 4
  // proto:  void QWidget::lower();
extern void C_ZN7QWidget5lowerEv(void* qthis); // 4
  // proto:  QPoint QWidget::mapToGlobal(const QPoint & );
extern void* C_ZNK7QWidget11mapToGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMinimumWidth(int minw);
extern void C_ZN7QWidget15setMinimumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::getContentsMargins(int * left, int * top, int * right, int * bottom);
extern void C_ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  bool QWidget::underMouse();
extern bool C_ZNK7QWidget10underMouseEv(void* qthis); // 2
  // proto:  QRegion QWidget::mask();
extern void* C_ZNK7QWidget4maskEv(void* qthis); // 4
  // proto:  qreal QWidget::windowOpacity();
extern double C_ZNK7QWidget13windowOpacityEv(void* qthis); // 4
  // proto:  int QWidget::height();
extern int32_t C_ZNK7QWidget6heightEv(void* qthis); // 2
  // proto:  QWidget * QWidget::nativeParentWidget();
extern void* C_ZNK7QWidget18nativeParentWidgetEv(void* qthis); // 4
  // proto:  QString QWidget::accessibleDescription();
extern void* C_ZNK7QWidget21accessibleDescriptionEv(void* qthis); // 4
  // proto:  int QWidget::y();
extern void C_ZNK7QWidget1yEv(void* qthis); // 4
  // proto:  QPalette::ColorRole QWidget::backgroundRole();
extern void C_ZNK7QWidget14backgroundRoleEv(void* qthis); // 4
  // proto:  void QWidget::scroll(int dx, int dy);
extern void C_ZN7QWidget6scrollEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::scroll(int dx, int dy, const QRect & );
extern void C_ZN7QWidget6scrollEiiRK5QRect(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QWidget::setWhatsThis(const QString & );
extern void C_ZN7QWidget12setWhatsThisERK7QString(void* qthis, void* arg0); // 4
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

// class sizeof(QWidget)=1
type QWidget struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _windowIconChanged QWidget_windowIconChanged_signal;
//  _windowTitleChanged QWidget_windowTitleChanged_signal;
//  _customContextMenuRequested QWidget_customContextMenuRequested_signal;
//  _windowIconTextChanged QWidget_windowIconTextChanged_signal;
}

// class sizeof(QWidgetData)=1
type QWidgetData struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setWindowIconText(const class QString &)
func (this *QWidget) Setwindowicontext(args ...interface{}) () {
  // setWindowIconText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setWindowIconTextERK7QString
    // invoke: void setWindowIconText(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget17setWindowIconTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowIconText", args)
  }

  return
}

// show()
func (this *QWidget) Show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget4showEv
    // invoke: void show()
    C.C_ZN7QWidget4showEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "show", args)
  }

  return
}

// hasFocus()
func (this *QWidget) Hasfocus(args ...interface{}) (ret interface{}) {
  // hasFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8hasFocusEv
    // invoke: bool hasFocus()
    var ret0 = C.C_ZNK7QWidget8hasFocusEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "hasFocus", args)
  }

  return
}

// insertAction(class QAction *, class QAction *)
func (this *QWidget) Insertaction(args ...interface{}) () {
  // insertAction(class QAction *, class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12insertActionEP7QActionS1_
    // invoke: void insertAction(class QAction *, class QAction *)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QAction).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget12insertActionEP7QActionS1_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "insertAction", args)
  }

  return
}

// move(int, int)
func (this *QWidget) Move_(args ...interface{}) () {
  // move(int, int)
  // move(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget4moveEii
    // invoke: void move(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget4moveEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget4moveERK6QPoint
    // invoke: void move(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget4moveERK6QPoint(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "move", args)
  }

  return
}

// actions()
func (this *QWidget) Actions(args ...interface{}) () {
  // actions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7actionsEv
    // invoke: QList<QAction *> actions()
    C.C_ZNK7QWidget7actionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "actions", args)
  }

  return
}

// setStyle(class QStyle *)
func (this *QWidget) Setstyle(args ...interface{}) () {
  // setStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget8setStyleEP6QStyle
    // invoke: void setStyle(class QStyle *)
    var arg0 = args[0].(*QStyle).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget8setStyleEP6QStyle(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setStyle", args)
  }

  return
}

// keyboardGrabber()
func (this *QWidget) Keyboardgrabber_S(args ...interface{}) (ret interface{}) {
  // keyboardGrabber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15keyboardGrabberEv
    // invoke: QWidget * keyboardGrabber()
    var ret0 = C.C_ZN7QWidget15keyboardGrabberEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "keyboardGrabber", args)
  }

  return
}

// setHidden(_Bool)
func (this *QWidget) Sethidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setHiddenEb
    // invoke: void setHidden(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget9setHiddenEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setHidden", args)
  }

  return
}

// childAt(int, int)
func (this *QWidget) Childat(args ...interface{}) (ret interface{}) {
  // childAt(int, int)
  // childAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7childAtEii
    // invoke: QWidget * childAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QWidget7childAtEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QWidget7childAtERK6QPoint
    // invoke: QWidget * childAt(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget7childAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "childAt", args)
  }

  return
}

// focusWidget()
func (this *QWidget) Focuswidget(args ...interface{}) (ret interface{}) {
  // focusWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11focusWidgetEv
    // invoke: QWidget * focusWidget()
    var ret0 = C.C_ZNK7QWidget11focusWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "focusWidget", args)
  }

  return
}

// minimumHeight()
func (this *QWidget) Minimumheight(args ...interface{}) (ret interface{}) {
  // minimumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13minimumHeightEv
    // invoke: int minimumHeight()
    var ret0 = C.C_ZNK7QWidget13minimumHeightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "minimumHeight", args)
  }

  return
}

// font()
func (this *QWidget) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4fontEv
    // invoke: const QFont & font()
    var ret0 = C.C_ZNK7QWidget4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "const QFont &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "font", args)
  }

  return
}

// setLayout(class QLayout *)
func (this *QWidget) Setlayout(args ...interface{}) () {
  // setLayout(class QLayout *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayout{}) // "QLayout *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setLayoutEP7QLayout
    // invoke: void setLayout(class QLayout *)
    var arg0 = args[0].(*QLayout).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget9setLayoutEP7QLayout(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setLayout", args)
  }

  return
}

// layout()
func (this *QWidget) Layout(args ...interface{}) (ret interface{}) {
  // layout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6layoutEv
    // invoke: QLayout * layout()
    var ret0 = C.C_ZNK7QWidget6layoutEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLayout{}) // "QLayout *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "layout", args)
  }

  return
}

// activateWindow()
func (this *QWidget) Activatewindow(args ...interface{}) () {
  // activateWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14activateWindowEv
    // invoke: void activateWindow()
    C.C_ZN7QWidget14activateWindowEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "activateWindow", args)
  }

  return
}

// unsetLayoutDirection()
func (this *QWidget) Unsetlayoutdirection(args ...interface{}) () {
  // unsetLayoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget20unsetLayoutDirectionEv
    // invoke: void unsetLayoutDirection()
    C.C_ZN7QWidget20unsetLayoutDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "unsetLayoutDirection", args)
  }

  return
}

// isEnabled()
func (this *QWidget) Isenabled(args ...interface{}) (ret interface{}) {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9isEnabledEv
    // invoke: bool isEnabled()
    var ret0 = C.C_ZNK7QWidget9isEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isEnabled", args)
  }

  return
}

// contextMenuPolicy()
func (this *QWidget) Contextmenupolicy(args ...interface{}) () {
  // contextMenuPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget17contextMenuPolicyEv
    // invoke: Qt::ContextMenuPolicy contextMenuPolicy()
    C.C_ZNK7QWidget17contextMenuPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "contextMenuPolicy", args)
  }

  return
}

// setWindowFilePath(const class QString &)
func (this *QWidget) Setwindowfilepath(args ...interface{}) () {
  // setWindowFilePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setWindowFilePathERK7QString
    // invoke: void setWindowFilePath(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget17setWindowFilePathERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowFilePath", args)
  }

  return
}

// accessibleName()
func (this *QWidget) Accessiblename(args ...interface{}) (ret interface{}) {
  // accessibleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14accessibleNameEv
    // invoke: QString accessibleName()
    var ret0 = C.C_ZNK7QWidget14accessibleNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "accessibleName", args)
  }

  return
}

// window()
func (this *QWidget) Window(args ...interface{}) (ret interface{}) {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6windowEv
    // invoke: QWidget * window()
    var ret0 = C.C_ZNK7QWidget6windowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "window", args)
  }

  return
}

// internalWinId()
func (this *QWidget) Internalwinid(args ...interface{}) () {
  // internalWinId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13internalWinIdEv
    // invoke: WId internalWinId()
    C.C_ZNK7QWidget13internalWinIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "internalWinId", args)
  }

  return
}

// setWindowTitle(const class QString &)
func (this *QWidget) Setwindowtitle(args ...interface{}) () {
  // setWindowTitle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setWindowTitleERK7QString
    // invoke: void setWindowTitle(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget14setWindowTitleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowTitle", args)
  }

  return
}

// mouseGrabber()
func (this *QWidget) Mousegrabber_S(args ...interface{}) (ret interface{}) {
  // mouseGrabber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12mouseGrabberEv
    // invoke: QWidget * mouseGrabber()
    var ret0 = C.C_ZN7QWidget12mouseGrabberEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mouseGrabber", args)
  }

  return
}

// isModal()
func (this *QWidget) Ismodal(args ...interface{}) (ret interface{}) {
  // isModal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7isModalEv
    // invoke: bool isModal()
    var ret0 = C.C_ZNK7QWidget7isModalEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isModal", args)
  }

  return
}

// foregroundRole()
func (this *QWidget) Foregroundrole(args ...interface{}) () {
  // foregroundRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14foregroundRoleEv
    // invoke: QPalette::ColorRole foregroundRole()
    C.C_ZNK7QWidget14foregroundRoleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "foregroundRole", args)
  }

  return
}

// setShortcutAutoRepeat(int, _Bool)
func (this *QWidget) Setshortcutautorepeat(args ...interface{}) () {
  // setShortcutAutoRepeat(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget21setShortcutAutoRepeatEib
    // invoke: void setShortcutAutoRepeat(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget21setShortcutAutoRepeatEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setShortcutAutoRepeat", args)
  }

  return
}

// setGraphicsEffect(class QGraphicsEffect *)
func (this *QWidget) Setgraphicseffect(args ...interface{}) () {
  // setGraphicsEffect(class QGraphicsEffect *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QGraphicsEffect{}) // "QGraphicsEffect *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect
    // invoke: void setGraphicsEffect(class QGraphicsEffect *)
    var arg0 = args[0].(*QGraphicsEffect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setGraphicsEffect", args)
  }

  return
}

// setAccessibleName(const class QString &)
func (this *QWidget) Setaccessiblename(args ...interface{}) () {
  // setAccessibleName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setAccessibleNameERK7QString
    // invoke: void setAccessibleName(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget17setAccessibleNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAccessibleName", args)
  }

  return
}

// setMinimumHeight(int)
func (this *QWidget) Setminimumheight(args ...interface{}) () {
  // setMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setMinimumHeightEi
    // invoke: void setMinimumHeight(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget16setMinimumHeightEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumHeight", args)
  }

  return
}

// mapToParent(const class QPoint &)
func (this *QWidget) Maptoparent(args ...interface{}) (ret interface{}) {
  // mapToParent(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11mapToParentERK6QPoint
    // invoke: QPoint mapToParent(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget11mapToParentERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mapToParent", args)
  }

  return
}

// setGeometry(int, int, int, int)
func (this *QWidget) Setgeometry(args ...interface{}) () {
  // setGeometry(int, int, int, int)
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11setGeometryEiiii
    // invoke: void setGeometry(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN7QWidget11setGeometryEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QWidget11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget11setGeometryERK5QRect(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setGeometry", args)
  }

  return
}

// pos()
func (this *QWidget) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget3posEv
    // invoke: QPoint pos()
    var ret0 = C.C_ZNK7QWidget3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "pos", args)
  }

  return
}

// sizeHint()
func (this *QWidget) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK7QWidget8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "sizeHint", args)
  }

  return
}

// windowType()
func (this *QWidget) Windowtype(args ...interface{}) () {
  // windowType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10windowTypeEv
    // invoke: Qt::WindowType windowType()
    C.C_ZNK7QWidget10windowTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowType", args)
  }

  return
}

// removeAction(class QAction *)
func (this *QWidget) Removeaction(args ...interface{}) () {
  // removeAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12removeActionEP7QAction
    // invoke: void removeAction(class QAction *)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget12removeActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "removeAction", args)
  }

  return
}

// grabMouse()
func (this *QWidget) Grabmouse(args ...interface{}) () {
  // grabMouse()
  // grabMouse(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9grabMouseEv
    // invoke: void grabMouse()
    C.C_ZN7QWidget9grabMouseEv(this.Qclsinst)
  case 1:
    // invoke: _ZN7QWidget9grabMouseERK7QCursor
    // invoke: void grabMouse(const class QCursor &)
    var arg0 = args[0].(*QCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget9grabMouseERK7QCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "grabMouse", args)
  }

  return
}

// frameGeometry()
func (this *QWidget) Framegeometry(args ...interface{}) (ret interface{}) {
  // frameGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13frameGeometryEv
    // invoke: QRect frameGeometry()
    var ret0 = C.C_ZNK7QWidget13frameGeometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "frameGeometry", args)
  }

  return
}

// setWindowRole(const class QString &)
func (this *QWidget) Setwindowrole(args ...interface{}) () {
  // setWindowRole(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setWindowRoleERK7QString
    // invoke: void setWindowRole(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget13setWindowRoleERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowRole", args)
  }

  return
}

// repaint(const class QRect &)
func (this *QWidget) Repaint(args ...interface{}) () {
  // repaint(const class QRect &)
  // repaint()
  // repaint(int, int, int, int)
  // repaint(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget7repaintERK5QRect
    // invoke: void repaint(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget7repaintERK5QRect(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget7repaintEv
    // invoke: void repaint()
    C.C_ZN7QWidget7repaintEv(this.Qclsinst)
  case 2:
    // invoke: _ZN7QWidget7repaintEiiii
    // invoke: void repaint(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN7QWidget7repaintEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZN7QWidget7repaintERK7QRegion
    // invoke: void repaint(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget7repaintERK7QRegion(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "repaint", args)
  }

  return
}

// inputMethodHints()
func (this *QWidget) Inputmethodhints(args ...interface{}) () {
  // inputMethodHints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget16inputMethodHintsEv
    // invoke: Qt::InputMethodHints inputMethodHints()
    C.C_ZNK7QWidget16inputMethodHintsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "inputMethodHints", args)
  }

  return
}

// metaObject()
func (this *QWidget) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "metaObject", args)
  }

  return
}

// topLevelWidget()
func (this *QWidget) Toplevelwidget(args ...interface{}) (ret interface{}) {
  // topLevelWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14topLevelWidgetEv
    // invoke: QWidget * topLevelWidget()
    var ret0 = C.C_ZNK7QWidget14topLevelWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "topLevelWidget", args)
  }

  return
}

// releaseMouse()
func (this *QWidget) Releasemouse(args ...interface{}) () {
  // releaseMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12releaseMouseEv
    // invoke: void releaseMouse()
    C.C_ZN7QWidget12releaseMouseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "releaseMouse", args)
  }

  return
}

// setLocale(const class QLocale &)
func (this *QWidget) Setlocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setLocaleERK7QLocale
    // invoke: void setLocale(const class QLocale &)
    var arg0 = args[0].(*QLocale).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget9setLocaleERK7QLocale(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setLocale", args)
  }

  return
}

// contentsRect()
func (this *QWidget) Contentsrect(args ...interface{}) (ret interface{}) {
  // contentsRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12contentsRectEv
    // invoke: QRect contentsRect()
    var ret0 = C.C_ZNK7QWidget12contentsRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "contentsRect", args)
  }

  return
}

// focusPolicy()
func (this *QWidget) Focuspolicy(args ...interface{}) () {
  // focusPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11focusPolicyEv
    // invoke: Qt::FocusPolicy focusPolicy()
    C.C_ZNK7QWidget11focusPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "focusPolicy", args)
  }

  return
}

// isVisibleTo(const class QWidget *)
func (this *QWidget) Isvisibleto(args ...interface{}) (ret interface{}) {
  // isVisibleTo(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isVisibleToEPKS_
    // invoke: bool isVisibleTo(const class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget11isVisibleToEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isVisibleTo", args)
  }

  return
}

// updatesEnabled()
func (this *QWidget) Updatesenabled(args ...interface{}) (ret interface{}) {
  // updatesEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14updatesEnabledEv
    // invoke: bool updatesEnabled()
    var ret0 = C.C_ZNK7QWidget14updatesEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "updatesEnabled", args)
  }

  return
}

// windowState()
func (this *QWidget) Windowstate(args ...interface{}) () {
  // windowState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11windowStateEv
    // invoke: Qt::WindowStates windowState()
    C.C_ZNK7QWidget11windowStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowState", args)
  }

  return
}

// isWindowModified()
func (this *QWidget) Iswindowmodified(args ...interface{}) (ret interface{}) {
  // isWindowModified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget16isWindowModifiedEv
    // invoke: bool isWindowModified()
    var ret0 = C.C_ZNK7QWidget16isWindowModifiedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isWindowModified", args)
  }

  return
}

// devType()
func (this *QWidget) Devtype(args ...interface{}) (ret interface{}) {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7devTypeEv
    // invoke: int devType()
    var ret0 = C.C_ZNK7QWidget7devTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "devType", args)
  }

  return
}

// clearFocus()
func (this *QWidget) Clearfocus(args ...interface{}) () {
  // clearFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10clearFocusEv
    // invoke: void clearFocus()
    C.C_ZN7QWidget10clearFocusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "clearFocus", args)
  }

  return
}

// setStyleSheet(const class QString &)
func (this *QWidget) Setstylesheet(args ...interface{}) () {
  // setStyleSheet(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setStyleSheetERK7QString
    // invoke: void setStyleSheet(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget13setStyleSheetERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setStyleSheet", args)
  }

  return
}

// setMaximumHeight(int)
func (this *QWidget) Setmaximumheight(args ...interface{}) () {
  // setMaximumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setMaximumHeightEi
    // invoke: void setMaximumHeight(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget16setMaximumHeightEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumHeight", args)
  }

  return
}

// x()
func (this *QWidget) X(args ...interface{}) () {
  // x()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget1xEv
    // invoke: int x()
    C.C_ZNK7QWidget1xEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "x", args)
  }

  return
}

// grab(const class QRect &)
func (this *QWidget) Grab(args ...interface{}) (ret interface{}) {
  // grab(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget4grabERK5QRect
    // invoke: QPixmap grab(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QWidget4grabERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "grab", args)
  }

  return
}

// maximumSize()
func (this *QWidget) Maximumsize(args ...interface{}) (ret interface{}) {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11maximumSizeEv
    // invoke: QSize maximumSize()
    var ret0 = C.C_ZNK7QWidget11maximumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "maximumSize", args)
  }

  return
}

// locale()
func (this *QWidget) Locale(args ...interface{}) (ret interface{}) {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6localeEv
    // invoke: QLocale locale()
    var ret0 = C.C_ZNK7QWidget6localeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLocale{}) // "QLocale"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "locale", args)
  }

  return
}

// minimumSize()
func (this *QWidget) Minimumsize(args ...interface{}) (ret interface{}) {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11minimumSizeEv
    // invoke: QSize minimumSize()
    var ret0 = C.C_ZNK7QWidget11minimumSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "minimumSize", args)
  }

  return
}

// windowFilePath()
func (this *QWidget) Windowfilepath(args ...interface{}) (ret interface{}) {
  // windowFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14windowFilePathEv
    // invoke: QString windowFilePath()
    var ret0 = C.C_ZNK7QWidget14windowFilePathEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "windowFilePath", args)
  }

  return
}

// setStatusTip(const class QString &)
func (this *QWidget) Setstatustip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget12setStatusTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setStatusTip", args)
  }

  return
}

// focusProxy()
func (this *QWidget) Focusproxy(args ...interface{}) (ret interface{}) {
  // focusProxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10focusProxyEv
    // invoke: QWidget * focusProxy()
    var ret0 = C.C_ZNK7QWidget10focusProxyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "focusProxy", args)
  }

  return
}

// createWinId()
func (this *QWidget) Createwinid(args ...interface{}) () {
  // createWinId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11createWinIdEv
    // invoke: void createWinId()
    C.C_ZN7QWidget11createWinIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "createWinId", args)
  }

  return
}

// paintEngine()
func (this *QWidget) Paintengine(args ...interface{}) (ret interface{}) {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    var ret0 = C.C_ZNK7QWidget11paintEngineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintEngine{}) // "QPaintEngine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "paintEngine", args)
  }

  return
}

// saveGeometry()
func (this *QWidget) Savegeometry(args ...interface{}) (ret interface{}) {
  // saveGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12saveGeometryEv
    // invoke: QByteArray saveGeometry()
    var ret0 = C.C_ZNK7QWidget12saveGeometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "saveGeometry", args)
  }

  return
}

// backingStore()
func (this *QWidget) Backingstore(args ...interface{}) (ret interface{}) {
  // backingStore()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12backingStoreEv
    // invoke: QBackingStore * backingStore()
    var ret0 = C.C_ZNK7QWidget12backingStoreEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBackingStore{}) // "QBackingStore *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "backingStore", args)
  }

  return
}

// setToolTipDuration(int)
func (this *QWidget) Settooltipduration(args ...interface{}) () {
  // setToolTipDuration(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget18setToolTipDurationEi
    // invoke: void setToolTipDuration(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget18setToolTipDurationEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setToolTipDuration", args)
  }

  return
}

// cursor()
func (this *QWidget) Cursor(args ...interface{}) (ret interface{}) {
  // cursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6cursorEv
    // invoke: QCursor cursor()
    var ret0 = C.C_ZNK7QWidget6cursorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCursor{}) // "QCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "cursor", args)
  }

  return
}

// palette()
func (this *QWidget) Palette(args ...interface{}) (ret interface{}) {
  // palette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7paletteEv
    // invoke: const QPalette & palette()
    var ret0 = C.C_ZNK7QWidget7paletteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPalette{}) // "const QPalette &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "palette", args)
  }

  return
}

// hide()
func (this *QWidget) Hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget4hideEv
    // invoke: void hide()
    C.C_ZN7QWidget4hideEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "hide", args)
  }

  return
}

// windowTitle()
func (this *QWidget) Windowtitle(args ...interface{}) (ret interface{}) {
  // windowTitle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11windowTitleEv
    // invoke: QString windowTitle()
    var ret0 = C.C_ZNK7QWidget11windowTitleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "windowTitle", args)
  }

  return
}

// setFocusProxy(class QWidget *)
func (this *QWidget) Setfocusproxy(args ...interface{}) () {
  // setFocusProxy(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setFocusProxyEPS_
    // invoke: void setFocusProxy(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget13setFocusProxyEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFocusProxy", args)
  }

  return
}

// setParent(class QWidget *)
func (this *QWidget) Setparent(args ...interface{}) () {
  // setParent(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setParentEPS_
    // invoke: void setParent(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget9setParentEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setParent", args)
  }

  return
}

// sizePolicy()
func (this *QWidget) Sizepolicy(args ...interface{}) (ret interface{}) {
  // sizePolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10sizePolicyEv
    // invoke: QSizePolicy sizePolicy()
    var ret0 = C.C_ZNK7QWidget10sizePolicyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizePolicy{}) // "QSizePolicy"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "sizePolicy", args)
  }

  return
}

// visibleRegion()
func (this *QWidget) Visibleregion(args ...interface{}) (ret interface{}) {
  // visibleRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13visibleRegionEv
    // invoke: QRegion visibleRegion()
    var ret0 = C.C_ZNK7QWidget13visibleRegionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "visibleRegion", args)
  }

  return
}

// windowFlags()
func (this *QWidget) Windowflags(args ...interface{}) () {
  // windowFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11windowFlagsEv
    // invoke: Qt::WindowFlags windowFlags()
    C.C_ZNK7QWidget11windowFlagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowFlags", args)
  }

  return
}

// setContentsMargins(int, int, int, int)
func (this *QWidget) Setcontentsmargins(args ...interface{}) () {
  // setContentsMargins(int, int, int, int)
  // setContentsMargins(const class QMargins &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QMargins{}) // "const QMargins &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget18setContentsMarginsEiiii
    // invoke: void setContentsMargins(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN7QWidget18setContentsMarginsEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QWidget18setContentsMarginsERK8QMargins
    // invoke: void setContentsMargins(const class QMargins &)
    var arg0 = args[0].(*QMargins).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget18setContentsMarginsERK8QMargins(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setContentsMargins", args)
  }

  return
}

// graphicsEffect()
func (this *QWidget) Graphicseffect(args ...interface{}) () {
  // graphicsEffect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14graphicsEffectEv
    // invoke: QGraphicsEffect * graphicsEffect()
    C.C_ZNK7QWidget14graphicsEffectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "graphicsEffect", args)
  }

  return
}

// mapFromParent(const class QPoint &)
func (this *QWidget) Mapfromparent(args ...interface{}) (ret interface{}) {
  // mapFromParent(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13mapFromParentERK6QPoint
    // invoke: QPoint mapFromParent(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget13mapFromParentERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mapFromParent", args)
  }

  return
}

// windowRole()
func (this *QWidget) Windowrole(args ...interface{}) (ret interface{}) {
  // windowRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10windowRoleEv
    // invoke: QString windowRole()
    var ret0 = C.C_ZNK7QWidget10windowRoleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "windowRole", args)
  }

  return
}

// style()
func (this *QWidget) Style(args ...interface{}) (ret interface{}) {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget5styleEv
    // invoke: QStyle * style()
    var ret0 = C.C_ZNK7QWidget5styleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "style", args)
  }

  return
}

// toolTip()
func (this *QWidget) Tooltip(args ...interface{}) (ret interface{}) {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7toolTipEv
    // invoke: QString toolTip()
    var ret0 = C.C_ZNK7QWidget7toolTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "toolTip", args)
  }

  return
}

// maximumWidth()
func (this *QWidget) Maximumwidth(args ...interface{}) (ret interface{}) {
  // maximumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12maximumWidthEv
    // invoke: int maximumWidth()
    var ret0 = C.C_ZNK7QWidget12maximumWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "maximumWidth", args)
  }

  return
}

// isEnabledToTLW()
func (this *QWidget) Isenabledtotlw(args ...interface{}) (ret interface{}) {
  // isEnabledToTLW()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14isEnabledToTLWEv
    // invoke: bool isEnabledToTLW()
    var ret0 = C.C_ZNK7QWidget14isEnabledToTLWEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isEnabledToTLW", args)
  }

  return
}

// windowIcon()
func (this *QWidget) Windowicon(args ...interface{}) (ret interface{}) {
  // windowIcon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10windowIconEv
    // invoke: QIcon windowIcon()
    var ret0 = C.C_ZNK7QWidget10windowIconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "windowIcon", args)
  }

  return
}

// isMinimized()
func (this *QWidget) Isminimized(args ...interface{}) (ret interface{}) {
  // isMinimized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isMinimizedEv
    // invoke: bool isMinimized()
    var ret0 = C.C_ZNK7QWidget11isMinimizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isMinimized", args)
  }

  return
}

// rect()
func (this *QWidget) Rect(args ...interface{}) (ret interface{}) {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4rectEv
    // invoke: QRect rect()
    var ret0 = C.C_ZNK7QWidget4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "rect", args)
  }

  return
}

// raise()
func (this *QWidget) Raise(args ...interface{}) () {
  // raise()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget5raiseEv
    // invoke: void raise()
    C.C_ZN7QWidget5raiseEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "raise", args)
  }

  return
}

// stackUnder(class QWidget *)
func (this *QWidget) Stackunder(args ...interface{}) () {
  // stackUnder(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10stackUnderEPS_
    // invoke: void stackUnder(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget10stackUnderEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "stackUnder", args)
  }

  return
}

// parentWidget()
func (this *QWidget) Parentwidget(args ...interface{}) (ret interface{}) {
  // parentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12parentWidgetEv
    // invoke: QWidget * parentWidget()
    var ret0 = C.C_ZNK7QWidget12parentWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "parentWidget", args)
  }

  return
}

// effectiveWinId()
func (this *QWidget) Effectivewinid(args ...interface{}) () {
  // effectiveWinId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14effectiveWinIdEv
    // invoke: WId effectiveWinId()
    C.C_ZNK7QWidget14effectiveWinIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "effectiveWinId", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QWidget) Settooltip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget10setToolTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setToolTip", args)
  }

  return
}

// setSizePolicy(class QSizePolicy)
func (this *QWidget) Setsizepolicy(args ...interface{}) () {
  // setSizePolicy(class QSizePolicy)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSizePolicy{}) // "QSizePolicy"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setSizePolicyE11QSizePolicy
    // invoke: void setSizePolicy(class QSizePolicy)
    var arg0 = args[0].(*QSizePolicy).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget13setSizePolicyE11QSizePolicy(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setSizePolicy", args)
  }

  return
}

// geometry()
func (this *QWidget) Geometry(args ...interface{}) (ret interface{}) {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8geometryEv
    // invoke: const QRect & geometry()
    var ret0 = C.C_ZNK7QWidget8geometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "const QRect &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "geometry", args)
  }

  return
}

// windowHandle()
func (this *QWidget) Windowhandle(args ...interface{}) (ret interface{}) {
  // windowHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12windowHandleEv
    // invoke: QWindow * windowHandle()
    var ret0 = C.C_ZNK7QWidget12windowHandleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "windowHandle", args)
  }

  return
}

// setAcceptDrops(_Bool)
func (this *QWidget) Setacceptdrops(args ...interface{}) () {
  // setAcceptDrops(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setAcceptDropsEb
    // invoke: void setAcceptDrops(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget14setAcceptDropsEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAcceptDrops", args)
  }

  return
}

// isEnabledTo(const class QWidget *)
func (this *QWidget) Isenabledto(args ...interface{}) (ret interface{}) {
  // isEnabledTo(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isEnabledToEPKS_
    // invoke: bool isEnabledTo(const class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget11isEnabledToEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isEnabledTo", args)
  }

  return
}

// isVisible()
func (this *QWidget) Isvisible(args ...interface{}) (ret interface{}) {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9isVisibleEv
    // invoke: bool isVisible()
    var ret0 = C.C_ZNK7QWidget9isVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isVisible", args)
  }

  return
}

// setWindowModified(_Bool)
func (this *QWidget) Setwindowmodified(args ...interface{}) () {
  // setWindowModified(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setWindowModifiedEb
    // invoke: void setWindowModified(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget17setWindowModifiedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowModified", args)
  }

  return
}

// size()
func (this *QWidget) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4sizeEv
    // invoke: QSize size()
    var ret0 = C.C_ZNK7QWidget4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "size", args)
  }

  return
}

// setMaximumWidth(int)
func (this *QWidget) Setmaximumwidth(args ...interface{}) () {
  // setMaximumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15setMaximumWidthEi
    // invoke: void setMaximumWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget15setMaximumWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumWidth", args)
  }

  return
}

// addAction(class QAction *)
func (this *QWidget) Addaction(args ...interface{}) () {
  // addAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9addActionEP7QAction
    // invoke: void addAction(class QAction *)
    var arg0 = args[0].(*QAction).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget9addActionEP7QAction(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "addAction", args)
  }

  return
}

// normalGeometry()
func (this *QWidget) Normalgeometry(args ...interface{}) (ret interface{}) {
  // normalGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14normalGeometryEv
    // invoke: QRect normalGeometry()
    var ret0 = C.C_ZNK7QWidget14normalGeometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "normalGeometry", args)
  }

  return
}

// setMinimumSize(int, int)
func (this *QWidget) Setminimumsize(args ...interface{}) () {
  // setMinimumSize(int, int)
  // setMinimumSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setMinimumSizeEii
    // invoke: void setMinimumSize(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget14setMinimumSizeEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget14setMinimumSizeERK5QSize
    // invoke: void setMinimumSize(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget14setMinimumSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumSize", args)
  }

  return
}

// heightForWidth(int)
func (this *QWidget) Heightforwidth(args ...interface{}) (ret interface{}) {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget14heightForWidthEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "heightForWidth", args)
  }

  return
}

// setPalette(const class QPalette &)
func (this *QWidget) Setpalette(args ...interface{}) () {
  // setPalette(const class QPalette &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "const QPalette &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setPaletteERK8QPalette
    // invoke: void setPalette(const class QPalette &)
    var arg0 = args[0].(*QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget10setPaletteERK8QPalette(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setPalette", args)
  }

  return
}

// isAncestorOf(const class QWidget *)
func (this *QWidget) Isancestorof(args ...interface{}) (ret interface{}) {
  // isAncestorOf(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12isAncestorOfEPKS_
    // invoke: bool isAncestorOf(const class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget12isAncestorOfEPKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isAncestorOf", args)
  }

  return
}

// clearMask()
func (this *QWidget) Clearmask(args ...interface{}) () {
  // clearMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9clearMaskEv
    // invoke: void clearMask()
    C.C_ZN7QWidget9clearMaskEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "clearMask", args)
  }

  return
}

// setWindowOpacity(qreal)
func (this *QWidget) Setwindowopacity(args ...interface{}) () {
  // setWindowOpacity(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setWindowOpacityEd
    // invoke: void setWindowOpacity(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget16setWindowOpacityEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowOpacity", args)
  }

  return
}

// isActiveWindow()
func (this *QWidget) Isactivewindow(args ...interface{}) (ret interface{}) {
  // isActiveWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14isActiveWindowEv
    // invoke: bool isActiveWindow()
    var ret0 = C.C_ZNK7QWidget14isActiveWindowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isActiveWindow", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QWidget) Setfont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFont", args)
  }

  return
}

// setAccessibleDescription(const class QString &)
func (this *QWidget) Setaccessibledescription(args ...interface{}) () {
  // setAccessibleDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget24setAccessibleDescriptionERK7QString
    // invoke: void setAccessibleDescription(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget24setAccessibleDescriptionERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAccessibleDescription", args)
  }

  return
}

// setMaximumSize(const class QSize &)
func (this *QWidget) Setmaximumsize(args ...interface{}) () {
  // setMaximumSize(const class QSize &)
  // setMaximumSize(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setMaximumSizeERK5QSize
    // invoke: void setMaximumSize(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget14setMaximumSizeERK5QSize(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget14setMaximumSizeEii
    // invoke: void setMaximumSize(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget14setMaximumSizeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumSize", args)
  }

  return
}

// unsetLocale()
func (this *QWidget) Unsetlocale(args ...interface{}) () {
  // unsetLocale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11unsetLocaleEv
    // invoke: void unsetLocale()
    C.C_ZN7QWidget11unsetLocaleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "unsetLocale", args)
  }

  return
}

// autoFillBackground()
func (this *QWidget) Autofillbackground(args ...interface{}) (ret interface{}) {
  // autoFillBackground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget18autoFillBackgroundEv
    // invoke: bool autoFillBackground()
    var ret0 = C.C_ZNK7QWidget18autoFillBackgroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "autoFillBackground", args)
  }

  return
}

// unsetCursor()
func (this *QWidget) Unsetcursor(args ...interface{}) () {
  // unsetCursor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11unsetCursorEv
    // invoke: void unsetCursor()
    C.C_ZN7QWidget11unsetCursorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "unsetCursor", args)
  }

  return
}

// winId()
func (this *QWidget) Winid(args ...interface{}) () {
  // winId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget5winIdEv
    // invoke: WId winId()
    C.C_ZNK7QWidget5winIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "winId", args)
  }

  return
}

// layoutDirection()
func (this *QWidget) Layoutdirection(args ...interface{}) () {
  // layoutDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget15layoutDirectionEv
    // invoke: Qt::LayoutDirection layoutDirection()
    C.C_ZNK7QWidget15layoutDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "layoutDirection", args)
  }

  return
}

// setShortcutEnabled(int, _Bool)
func (this *QWidget) Setshortcutenabled(args ...interface{}) () {
  // setShortcutEnabled(int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget18setShortcutEnabledEib
    // invoke: void setShortcutEnabled(int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget18setShortcutEnabledEib(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setShortcutEnabled", args)
  }

  return
}

// sizeIncrement()
func (this *QWidget) Sizeincrement(args ...interface{}) (ret interface{}) {
  // sizeIncrement()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13sizeIncrementEv
    // invoke: QSize sizeIncrement()
    var ret0 = C.C_ZNK7QWidget13sizeIncrementEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "sizeIncrement", args)
  }

  return
}

// setMouseTracking(_Bool)
func (this *QWidget) Setmousetracking(args ...interface{}) () {
  // setMouseTracking(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setMouseTrackingEb
    // invoke: void setMouseTracking(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget16setMouseTrackingEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMouseTracking", args)
  }

  return
}

// whatsThis()
func (this *QWidget) Whatsthis(args ...interface{}) (ret interface{}) {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9whatsThisEv
    // invoke: QString whatsThis()
    var ret0 = C.C_ZNK7QWidget9whatsThisEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "whatsThis", args)
  }

  return
}

// width()
func (this *QWidget) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget5widthEv
    // invoke: int width()
    var ret0 = C.C_ZNK7QWidget5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "width", args)
  }

  return
}

// childrenRect()
func (this *QWidget) Childrenrect(args ...interface{}) (ret interface{}) {
  // childrenRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12childrenRectEv
    // invoke: QRect childrenRect()
    var ret0 = C.C_ZNK7QWidget12childrenRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "childrenRect", args)
  }

  return
}

// windowIconText()
func (this *QWidget) Windowicontext(args ...interface{}) (ret interface{}) {
  // windowIconText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14windowIconTextEv
    // invoke: QString windowIconText()
    var ret0 = C.C_ZNK7QWidget14windowIconTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "windowIconText", args)
  }

  return
}

// toolTipDuration()
func (this *QWidget) Tooltipduration(args ...interface{}) (ret interface{}) {
  // toolTipDuration()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget15toolTipDurationEv
    // invoke: int toolTipDuration()
    var ret0 = C.C_ZNK7QWidget15toolTipDurationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "toolTipDuration", args)
  }

  return
}

// setEnabled(_Bool)
func (this *QWidget) Setenabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget10setEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setEnabled", args)
  }

  return
}

// showMaximized()
func (this *QWidget) Showmaximized(args ...interface{}) () {
  // showMaximized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13showMaximizedEv
    // invoke: void showMaximized()
    C.C_ZN7QWidget13showMaximizedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showMaximized", args)
  }

  return
}

// ensurePolished()
func (this *QWidget) Ensurepolished(args ...interface{}) () {
  // ensurePolished()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14ensurePolishedEv
    // invoke: void ensurePolished()
    C.C_ZNK7QWidget14ensurePolishedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "ensurePolished", args)
  }

  return
}

// statusTip()
func (this *QWidget) Statustip(args ...interface{}) (ret interface{}) {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9statusTipEv
    // invoke: QString statusTip()
    var ret0 = C.C_ZNK7QWidget9statusTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "statusTip", args)
  }

  return
}

// acceptDrops()
func (this *QWidget) Acceptdrops(args ...interface{}) (ret interface{}) {
  // acceptDrops()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11acceptDropsEv
    // invoke: bool acceptDrops()
    var ret0 = C.C_ZNK7QWidget11acceptDropsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "acceptDrops", args)
  }

  return
}

// isFullScreen()
func (this *QWidget) Isfullscreen(args ...interface{}) (ret interface{}) {
  // isFullScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12isFullScreenEv
    // invoke: bool isFullScreen()
    var ret0 = C.C_ZNK7QWidget12isFullScreenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isFullScreen", args)
  }

  return
}

// nextInFocusChain()
func (this *QWidget) Nextinfocuschain(args ...interface{}) (ret interface{}) {
  // nextInFocusChain()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget16nextInFocusChainEv
    // invoke: QWidget * nextInFocusChain()
    var ret0 = C.C_ZNK7QWidget16nextInFocusChainEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "nextInFocusChain", args)
  }

  return
}

// styleSheet()
func (this *QWidget) Stylesheet(args ...interface{}) (ret interface{}) {
  // styleSheet()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10styleSheetEv
    // invoke: QString styleSheet()
    var ret0 = C.C_ZNK7QWidget10styleSheetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "styleSheet", args)
  }

  return
}

// setSizeIncrement(const class QSize &)
func (this *QWidget) Setsizeincrement(args ...interface{}) () {
  // setSizeIncrement(const class QSize &)
  // setSizeIncrement(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget16setSizeIncrementERK5QSize
    // invoke: void setSizeIncrement(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget16setSizeIncrementERK5QSize(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget16setSizeIncrementEii
    // invoke: void setSizeIncrement(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget16setSizeIncrementEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setSizeIncrement", args)
  }

  return
}

// grabKeyboard()
func (this *QWidget) Grabkeyboard(args ...interface{}) () {
  // grabKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12grabKeyboardEv
    // invoke: void grabKeyboard()
    C.C_ZN7QWidget12grabKeyboardEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "grabKeyboard", args)
  }

  return
}

// mapTo(const class QWidget *, const class QPoint &)
func (this *QWidget) Mapto(args ...interface{}) (ret interface{}) {
  // mapTo(const class QWidget *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget5mapToEPKS_RK6QPoint
    // invoke: QPoint mapTo(const class QWidget *, const class QPoint &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QWidget5mapToEPKS_RK6QPoint(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mapTo", args)
  }

  return
}

// resize(int, int)
func (this *QWidget) Resize(args ...interface{}) () {
  // resize(int, int)
  // resize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget6resizeEii
    // invoke: void resize(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget6resizeEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget6resizeERK5QSize
    // invoke: void resize(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget6resizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "resize", args)
  }

  return
}

// setFocus()
func (this *QWidget) Setfocus(args ...interface{}) () {
  // setFocus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget8setFocusEv
    // invoke: void setFocus()
    C.C_ZN7QWidget8setFocusEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "setFocus", args)
  }

  return
}

// mapFromGlobal(const class QPoint &)
func (this *QWidget) Mapfromglobal(args ...interface{}) (ret interface{}) {
  // mapFromGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13mapFromGlobalERK6QPoint
    // invoke: QPoint mapFromGlobal(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget13mapFromGlobalERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mapFromGlobal", args)
  }

  return
}

// frameSize()
func (this *QWidget) Framesize(args ...interface{}) (ret interface{}) {
  // frameSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget9frameSizeEv
    // invoke: QSize frameSize()
    var ret0 = C.C_ZNK7QWidget9frameSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "frameSize", args)
  }

  return
}

// releaseShortcut(int)
func (this *QWidget) Releaseshortcut(args ...interface{}) () {
  // releaseShortcut(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15releaseShortcutEi
    // invoke: void releaseShortcut(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget15releaseShortcutEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "releaseShortcut", args)
  }

  return
}

// minimumSizeHint()
func (this *QWidget) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK7QWidget15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "minimumSizeHint", args)
  }

  return
}

// isHidden()
func (this *QWidget) Ishidden(args ...interface{}) (ret interface{}) {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8isHiddenEv
    // invoke: bool isHidden()
    var ret0 = C.C_ZNK7QWidget8isHiddenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isHidden", args)
  }

  return
}

// hasHeightForWidth()
func (this *QWidget) Hasheightforwidth(args ...interface{}) (ret interface{}) {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    var ret0 = C.C_ZNK7QWidget17hasHeightForWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "hasHeightForWidth", args)
  }

  return
}

// setFixedSize(const class QSize &)
func (this *QWidget) Setfixedsize(args ...interface{}) () {
  // setFixedSize(const class QSize &)
  // setFixedSize(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12setFixedSizeERK5QSize
    // invoke: void setFixedSize(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget12setFixedSizeERK5QSize(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget12setFixedSizeEii
    // invoke: void setFixedSize(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget12setFixedSizeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setFixedSize", args)
  }

  return
}

// setTabOrder(class QWidget *, class QWidget *)
func (this *QWidget) Settaborder_S(args ...interface{}) () {
  // setTabOrder(class QWidget *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11setTabOrderEPS_S0_
    // invoke: void setTabOrder(class QWidget *, class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget11setTabOrderEPS_S0_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setTabOrder", args)
  }

  return
}

// setWindowIcon(const class QIcon &)
func (this *QWidget) Setwindowicon(args ...interface{}) () {
  // setWindowIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setWindowIconERK5QIcon
    // invoke: void setWindowIcon(const class QIcon &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget13setWindowIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowIcon", args)
  }

  return
}

// isLeftToRight()
func (this *QWidget) Islefttoright(args ...interface{}) (ret interface{}) {
  // isLeftToRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13isLeftToRightEv
    // invoke: bool isLeftToRight()
    var ret0 = C.C_ZNK7QWidget13isLeftToRightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isLeftToRight", args)
  }

  return
}

// isRightToLeft()
func (this *QWidget) Isrighttoleft(args ...interface{}) (ret interface{}) {
  // isRightToLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13isRightToLeftEv
    // invoke: bool isRightToLeft()
    var ret0 = C.C_ZNK7QWidget13isRightToLeftEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isRightToLeft", args)
  }

  return
}

// setVisible(_Bool)
func (this *QWidget) Setvisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setVisible", args)
  }

  return
}

// setCursor(const class QCursor &)
func (this *QWidget) Setcursor(args ...interface{}) () {
  // setCursor(const class QCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCursor{}) // "const QCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget9setCursorERK7QCursor
    // invoke: void setCursor(const class QCursor &)
    var arg0 = args[0].(*QCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget9setCursorERK7QCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setCursor", args)
  }

  return
}

// setFixedWidth(int)
func (this *QWidget) Setfixedwidth(args ...interface{}) () {
  // setFixedWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13setFixedWidthEi
    // invoke: void setFixedWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget13setFixedWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFixedWidth", args)
  }

  return
}

// mapFrom(const class QWidget *, const class QPoint &)
func (this *QWidget) Mapfrom(args ...interface{}) (ret interface{}) {
  // mapFrom(const class QWidget *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[0][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget7mapFromEPKS_RK6QPoint
    // invoke: QPoint mapFrom(const class QWidget *, const class QPoint &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QWidget7mapFromEPKS_RK6QPoint(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mapFrom", args)
  }

  return
}

// isWindow()
func (this *QWidget) Iswindow(args ...interface{}) (ret interface{}) {
  // isWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8isWindowEv
    // invoke: bool isWindow()
    var ret0 = C.C_ZNK7QWidget8isWindowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isWindow", args)
  }

  return
}

// close()
func (this *QWidget) Close(args ...interface{}) (ret interface{}) {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget5closeEv
    // invoke: bool close()
    var ret0 = C.C_ZN7QWidget5closeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "close", args)
  }

  return
}

// fontMetrics()
func (this *QWidget) Fontmetrics(args ...interface{}) (ret interface{}) {
  // fontMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11fontMetricsEv
    // invoke: QFontMetrics fontMetrics()
    var ret0 = C.C_ZNK7QWidget11fontMetricsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFontMetrics{}) // "QFontMetrics"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "fontMetrics", args)
  }

  return
}

// contentsMargins()
func (this *QWidget) Contentsmargins(args ...interface{}) (ret interface{}) {
  // contentsMargins()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget15contentsMarginsEv
    // invoke: QMargins contentsMargins()
    var ret0 = C.C_ZNK7QWidget15contentsMarginsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMargins{}) // "QMargins"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "contentsMargins", args)
  }

  return
}

// ~QWidget()
func (this *QWidget) Freeqwidget(args ...interface{}) () {
  // ~QWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidgetD0Ev
    // invoke: void ~QWidget()
    C.C_ZN7QWidgetD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "~QWidget", args)
  }

  return
}

// setUpdatesEnabled(_Bool)
func (this *QWidget) Setupdatesenabled(args ...interface{}) () {
  // setUpdatesEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget17setUpdatesEnabledEb
    // invoke: void setUpdatesEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget17setUpdatesEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setUpdatesEnabled", args)
  }

  return
}

// showMinimized()
func (this *QWidget) Showminimized(args ...interface{}) () {
  // showMinimized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget13showMinimizedEv
    // invoke: void showMinimized()
    C.C_ZN7QWidget13showMinimizedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showMinimized", args)
  }

  return
}

// childrenRegion()
func (this *QWidget) Childrenregion(args ...interface{}) (ret interface{}) {
  // childrenRegion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14childrenRegionEv
    // invoke: QRegion childrenRegion()
    var ret0 = C.C_ZNK7QWidget14childrenRegionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "childrenRegion", args)
  }

  return
}

// windowModality()
func (this *QWidget) Windowmodality(args ...interface{}) () {
  // windowModality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14windowModalityEv
    // invoke: Qt::WindowModality windowModality()
    C.C_ZNK7QWidget14windowModalityEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowModality", args)
  }

  return
}

// fontInfo()
func (this *QWidget) Fontinfo(args ...interface{}) (ret interface{}) {
  // fontInfo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8fontInfoEv
    // invoke: QFontInfo fontInfo()
    var ret0 = C.C_ZNK7QWidget8fontInfoEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFontInfo{}) // "QFontInfo"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "fontInfo", args)
  }

  return
}

// releaseKeyboard()
func (this *QWidget) Releasekeyboard(args ...interface{}) () {
  // releaseKeyboard()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15releaseKeyboardEv
    // invoke: void releaseKeyboard()
    C.C_ZN7QWidget15releaseKeyboardEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "releaseKeyboard", args)
  }

  return
}

// previousInFocusChain()
func (this *QWidget) Previousinfocuschain(args ...interface{}) (ret interface{}) {
  // previousInFocusChain()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget20previousInFocusChainEv
    // invoke: QWidget * previousInFocusChain()
    var ret0 = C.C_ZNK7QWidget20previousInFocusChainEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "previousInFocusChain", args)
  }

  return
}

// showFullScreen()
func (this *QWidget) Showfullscreen(args ...interface{}) () {
  // showFullScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14showFullScreenEv
    // invoke: void showFullScreen()
    C.C_ZN7QWidget14showFullScreenEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showFullScreen", args)
  }

  return
}

// baseSize()
func (this *QWidget) Basesize(args ...interface{}) (ret interface{}) {
  // baseSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget8baseSizeEv
    // invoke: QSize baseSize()
    var ret0 = C.C_ZNK7QWidget8baseSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "baseSize", args)
  }

  return
}

// hasMouseTracking()
func (this *QWidget) Hasmousetracking(args ...interface{}) (ret interface{}) {
  // hasMouseTracking()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget16hasMouseTrackingEv
    // invoke: bool hasMouseTracking()
    var ret0 = C.C_ZNK7QWidget16hasMouseTrackingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "hasMouseTracking", args)
  }

  return
}

// restoreGeometry(const class QByteArray &)
func (this *QWidget) Restoregeometry(args ...interface{}) (ret interface{}) {
  // restoreGeometry(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15restoreGeometryERK10QByteArray
    // invoke: bool restoreGeometry(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QWidget15restoreGeometryERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "restoreGeometry", args)
  }

  return
}

// setDisabled(_Bool)
func (this *QWidget) Setdisabled(args ...interface{}) () {
  // setDisabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11setDisabledEb
    // invoke: void setDisabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget11setDisabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setDisabled", args)
  }

  return
}

// showNormal()
func (this *QWidget) Shownormal(args ...interface{}) () {
  // showNormal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10showNormalEv
    // invoke: void showNormal()
    C.C_ZN7QWidget10showNormalEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showNormal", args)
  }

  return
}

// setFixedHeight(int)
func (this *QWidget) Setfixedheight(args ...interface{}) () {
  // setFixedHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14setFixedHeightEi
    // invoke: void setFixedHeight(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget14setFixedHeightEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFixedHeight", args)
  }

  return
}

// setBaseSize(int, int)
func (this *QWidget) Setbasesize(args ...interface{}) () {
  // setBaseSize(int, int)
  // setBaseSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget11setBaseSizeEii
    // invoke: void setBaseSize(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget11setBaseSizeEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget11setBaseSizeERK5QSize
    // invoke: void setBaseSize(const class QSize &)
    var arg0 = args[0].(*QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget11setBaseSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setBaseSize", args)
  }

  return
}

// setAutoFillBackground(_Bool)
func (this *QWidget) Setautofillbackground(args ...interface{}) () {
  // setAutoFillBackground(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget21setAutoFillBackgroundEb
    // invoke: void setAutoFillBackground(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget21setAutoFillBackgroundEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAutoFillBackground", args)
  }

  return
}

// updateGeometry()
func (this *QWidget) Updategeometry(args ...interface{}) () {
  // updateGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget14updateGeometryEv
    // invoke: void updateGeometry()
    C.C_ZN7QWidget14updateGeometryEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "updateGeometry", args)
  }

  return
}

// minimumWidth()
func (this *QWidget) Minimumwidth(args ...interface{}) (ret interface{}) {
  // minimumWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget12minimumWidthEv
    // invoke: int minimumWidth()
    var ret0 = C.C_ZNK7QWidget12minimumWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "minimumWidth", args)
  }

  return
}

// maximumHeight()
func (this *QWidget) Maximumheight(args ...interface{}) (ret interface{}) {
  // maximumHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13maximumHeightEv
    // invoke: int maximumHeight()
    var ret0 = C.C_ZNK7QWidget13maximumHeightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "maximumHeight", args)
  }

  return
}

// update()
func (this *QWidget) Update(args ...interface{}) () {
  // update()
  // update(int, int, int, int)
  // update(const class QRect &)
  // update(const class QRegion &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget6updateEv
    // invoke: void update()
    C.C_ZN7QWidget6updateEv(this.Qclsinst)
  case 1:
    // invoke: _ZN7QWidget6updateEiiii
    // invoke: void update(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN7QWidget6updateEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN7QWidget6updateERK5QRect
    // invoke: void update(const class QRect &)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget6updateERK5QRect(this.Qclsinst, arg0)
  case 3:
    // invoke: _ZN7QWidget6updateERK7QRegion
    // invoke: void update(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget6updateERK7QRegion(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "update", args)
  }

  return
}

// adjustSize()
func (this *QWidget) Adjustsize(args ...interface{}) () {
  // adjustSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget10adjustSizeEv
    // invoke: void adjustSize()
    C.C_ZN7QWidget10adjustSizeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "adjustSize", args)
  }

  return
}

// graphicsProxyWidget()
func (this *QWidget) Graphicsproxywidget(args ...interface{}) () {
  // graphicsProxyWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget19graphicsProxyWidgetEv
    // invoke: QGraphicsProxyWidget * graphicsProxyWidget()
    C.C_ZNK7QWidget19graphicsProxyWidgetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "graphicsProxyWidget", args)
  }

  return
}

// isTopLevel()
func (this *QWidget) Istoplevel(args ...interface{}) (ret interface{}) {
  // isTopLevel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10isTopLevelEv
    // invoke: bool isTopLevel()
    var ret0 = C.C_ZNK7QWidget10isTopLevelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isTopLevel", args)
  }

  return
}

// setMask(const class QRegion &)
func (this *QWidget) Setmask(args ...interface{}) () {
  // setMask(const class QRegion &)
  // setMask(const class QBitmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegion{}) // "const QRegion &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget7setMaskERK7QRegion
    // invoke: void setMask(const class QRegion &)
    var arg0 = args[0].(*QRegion).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget7setMaskERK7QRegion(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget7setMaskERK7QBitmap
    // invoke: void setMask(const class QBitmap &)
    var arg0 = args[0].(*QBitmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget7setMaskERK7QBitmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMask", args)
  }

  return
}

// isMaximized()
func (this *QWidget) Ismaximized(args ...interface{}) (ret interface{}) {
  // isMaximized()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11isMaximizedEv
    // invoke: bool isMaximized()
    var ret0 = C.C_ZNK7QWidget11isMaximizedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "isMaximized", args)
  }

  return
}

// lower()
func (this *QWidget) Lower(args ...interface{}) () {
  // lower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget5lowerEv
    // invoke: void lower()
    C.C_ZN7QWidget5lowerEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "lower", args)
  }

  return
}

// mapToGlobal(const class QPoint &)
func (this *QWidget) Maptoglobal(args ...interface{}) (ret interface{}) {
  // mapToGlobal(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget11mapToGlobalERK6QPoint
    // invoke: QPoint mapToGlobal(const class QPoint &)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QWidget11mapToGlobalERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mapToGlobal", args)
  }

  return
}

// setMinimumWidth(int)
func (this *QWidget) Setminimumwidth(args ...interface{}) () {
  // setMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget15setMinimumWidthEi
    // invoke: void setMinimumWidth(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget15setMinimumWidthEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumWidth", args)
  }

  return
}

// getContentsMargins(int *, int *, int *, int *)
func (this *QWidget) Getcontentsmargins(args ...interface{}) () {
  // getContentsMargins(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_
    // invoke: void getContentsMargins(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QWidget", "getContentsMargins", args)
  }

  return
}

// underMouse()
func (this *QWidget) Undermouse(args ...interface{}) (ret interface{}) {
  // underMouse()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget10underMouseEv
    // invoke: bool underMouse()
    var ret0 = C.C_ZNK7QWidget10underMouseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "underMouse", args)
  }

  return
}

// mask()
func (this *QWidget) Mask(args ...interface{}) (ret interface{}) {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget4maskEv
    // invoke: QRegion mask()
    var ret0 = C.C_ZNK7QWidget4maskEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRegion{}) // "QRegion"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "mask", args)
  }

  return
}

// windowOpacity()
func (this *QWidget) Windowopacity(args ...interface{}) (ret interface{}) {
  // windowOpacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget13windowOpacityEv
    // invoke: qreal windowOpacity()
    var ret0 = C.C_ZNK7QWidget13windowOpacityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "windowOpacity", args)
  }

  return
}

// height()
func (this *QWidget) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget6heightEv
    // invoke: int height()
    var ret0 = C.C_ZNK7QWidget6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "height", args)
  }

  return
}

// nativeParentWidget()
func (this *QWidget) Nativeparentwidget(args ...interface{}) (ret interface{}) {
  // nativeParentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget18nativeParentWidgetEv
    // invoke: QWidget * nativeParentWidget()
    var ret0 = C.C_ZNK7QWidget18nativeParentWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "nativeParentWidget", args)
  }

  return
}

// accessibleDescription()
func (this *QWidget) Accessibledescription(args ...interface{}) (ret interface{}) {
  // accessibleDescription()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget21accessibleDescriptionEv
    // invoke: QString accessibleDescription()
    var ret0 = C.C_ZNK7QWidget21accessibleDescriptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidget", "accessibleDescription", args)
  }

  return
}

// y()
func (this *QWidget) Y(args ...interface{}) () {
  // y()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget1yEv
    // invoke: int y()
    C.C_ZNK7QWidget1yEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "y", args)
  }

  return
}

// backgroundRole()
func (this *QWidget) Backgroundrole(args ...interface{}) () {
  // backgroundRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QWidget14backgroundRoleEv
    // invoke: QPalette::ColorRole backgroundRole()
    C.C_ZNK7QWidget14backgroundRoleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "backgroundRole", args)
  }

  return
}

// scroll(int, int)
func (this *QWidget) Scroll(args ...interface{}) () {
  // scroll(int, int)
  // scroll(int, int, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget6scrollEii
    // invoke: void scroll(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN7QWidget6scrollEii(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget6scrollEiiRK5QRect
    // invoke: void scroll(int, int, const class QRect &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QRect).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN7QWidget6scrollEiiRK5QRect(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QWidget", "scroll", args)
  }

  return
}

// setWhatsThis(const class QString &)
func (this *QWidget) Setwhatsthis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QWidget12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QWidget12setWhatsThisERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWhatsThis", args)
  }

  return
}

// <= body block end

