package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZN7QWidget17setWindowIconTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::show();
extern void _ZN7QWidget4showEv(void* qthis); // 4
  // proto:  bool QWidget::hasFocus();
extern void _ZNK7QWidget8hasFocusEv(void* qthis); // 4
  // proto:  void QWidget::insertAction(QAction * before, QAction * action);
extern void _ZN7QWidget12insertActionEP7QActionS1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QWidget::move(int x, int y);
extern void _ZN7QWidget4moveEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QWidget::move(const QPoint & );
extern void _ZN7QWidget4moveERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QList<QAction *> QWidget::actions();
extern void _ZNK7QWidget7actionsEv(void* qthis); // 4
  // proto:  void QWidget::setStyle(QStyle * );
extern void _ZN7QWidget8setStyleEP6QStyle(void* qthis, void* arg0); // 4
  // proto: static QWidget * QWidget::keyboardGrabber();
extern void _ZN7QWidget15keyboardGrabberEv(); // 4
  // proto:  void QWidget::setHidden(bool hidden);
extern void _ZN7QWidget9setHiddenEb(void* qthis, bool arg0); // 4
  // proto:  QWidget * QWidget::childAt(int x, int y);
extern void _ZNK7QWidget7childAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QWidget * QWidget::childAt(const QPoint & p);
extern void _ZNK7QWidget7childAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QWidget * QWidget::focusWidget();
extern void _ZNK7QWidget11focusWidgetEv(void* qthis); // 4
  // proto:  int QWidget::minimumHeight();
extern void _ZNK7QWidget13minimumHeightEv(void* qthis); // 2
  // proto:  const QFont & QWidget::font();
extern void _ZNK7QWidget4fontEv(void* qthis); // 2
  // proto:  void QWidget::setLayout(QLayout * );
extern void _ZN7QWidget9setLayoutEP7QLayout(void* qthis, void* arg0); // 4
  // proto:  QLayout * QWidget::layout();
extern void _ZNK7QWidget6layoutEv(void* qthis); // 4
  // proto:  void QWidget::activateWindow();
extern void _ZN7QWidget14activateWindowEv(void* qthis); // 4
  // proto:  void QWidget::unsetLayoutDirection();
extern void _ZN7QWidget20unsetLayoutDirectionEv(void* qthis); // 4
  // proto:  bool QWidget::isEnabled();
extern void _ZNK7QWidget9isEnabledEv(void* qthis); // 2
  // proto:  Qt::ContextMenuPolicy QWidget::contextMenuPolicy();
extern void _ZNK7QWidget17contextMenuPolicyEv(void* qthis); // 4
  // proto:  void QWidget::setWindowFilePath(const QString & filePath);
extern void _ZN7QWidget17setWindowFilePathERK7QString(void* qthis, void* arg0); // 4
  // proto:  QString QWidget::accessibleName();
extern void _ZNK7QWidget14accessibleNameEv(void* qthis); // 4
  // proto:  QWidget * QWidget::window();
extern void _ZNK7QWidget6windowEv(void* qthis); // 4
  // proto:  WId QWidget::internalWinId();
extern void _ZNK7QWidget13internalWinIdEv(void* qthis); // 2
  // proto:  void QWidget::setWindowTitle(const QString & );
extern void _ZN7QWidget14setWindowTitleERK7QString(void* qthis, void* arg0); // 4
  // proto: static QWidget * QWidget::mouseGrabber();
extern void _ZN7QWidget12mouseGrabberEv(); // 4
  // proto:  bool QWidget::isModal();
extern void _ZNK7QWidget7isModalEv(void* qthis); // 2
  // proto:  QPalette::ColorRole QWidget::foregroundRole();
extern void _ZNK7QWidget14foregroundRoleEv(void* qthis); // 4
  // proto:  void QWidget::setShortcutAutoRepeat(int id, bool enable);
extern void _ZN7QWidget21setShortcutAutoRepeatEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  void QWidget::setGraphicsEffect(QGraphicsEffect * effect);
extern void _ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setAccessibleName(const QString & name);
extern void _ZN7QWidget17setAccessibleNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMinimumHeight(int minh);
extern void _ZN7QWidget16setMinimumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  QPoint QWidget::mapToParent(const QPoint & );
extern void _ZNK7QWidget11mapToParentERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setGeometry(int x, int y, int w, int h);
extern void _ZN7QWidget11setGeometryEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QWidget::setGeometry(const QRect & );
extern void _ZN7QWidget11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QPoint QWidget::pos();
extern void _ZNK7QWidget3posEv(void* qthis); // 4
  // proto:  QSize QWidget::sizeHint();
extern void _ZNK7QWidget8sizeHintEv(void* qthis); // 4
  // proto:  Qt::WindowType QWidget::windowType();
extern void _ZNK7QWidget10windowTypeEv(void* qthis); // 2
  // proto:  void QWidget::removeAction(QAction * action);
extern void _ZN7QWidget12removeActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  void QWidget::grabMouse();
extern void _ZN7QWidget9grabMouseEv(void* qthis); // 4
  // proto:  void QWidget::grabMouse(const QCursor & );
extern void _ZN7QWidget9grabMouseERK7QCursor(void* qthis, void* arg0); // 4
  // proto:  QRect QWidget::frameGeometry();
extern void _ZNK7QWidget13frameGeometryEv(void* qthis); // 4
  // proto:  void QWidget::setWindowRole(const QString & );
extern void _ZN7QWidget13setWindowRoleERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::repaint(const QRect & );
extern void _ZN7QWidget7repaintERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QWidget::repaint();
extern void _ZN7QWidget7repaintEv(void* qthis); // 4
  // proto:  void QWidget::repaint(int x, int y, int w, int h);
extern void _ZN7QWidget7repaintEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QWidget::repaint(const QRegion & );
extern void _ZN7QWidget7repaintERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  Qt::InputMethodHints QWidget::inputMethodHints();
extern void _ZNK7QWidget16inputMethodHintsEv(void* qthis); // 4
  // proto:  const QMetaObject * QWidget::metaObject();
extern void _ZNK7QWidget10metaObjectEv(void* qthis); // 4
  // proto:  QWidget * QWidget::topLevelWidget();
extern void _ZNK7QWidget14topLevelWidgetEv(void* qthis); // 2
  // proto:  void QWidget::releaseMouse();
extern void _ZN7QWidget12releaseMouseEv(void* qthis); // 4
  // proto:  void QWidget::setLocale(const QLocale & locale);
extern void _ZN7QWidget9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  QRect QWidget::contentsRect();
extern void _ZNK7QWidget12contentsRectEv(void* qthis); // 4
  // proto:  Qt::FocusPolicy QWidget::focusPolicy();
extern void _ZNK7QWidget11focusPolicyEv(void* qthis); // 4
  // proto:  bool QWidget::isVisibleTo(const QWidget * );
extern void _ZNK7QWidget11isVisibleToEPKS_(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::updatesEnabled();
extern void _ZNK7QWidget14updatesEnabledEv(void* qthis); // 2
  // proto:  Qt::WindowStates QWidget::windowState();
extern void _ZNK7QWidget11windowStateEv(void* qthis); // 4
  // proto:  bool QWidget::isWindowModified();
extern void _ZNK7QWidget16isWindowModifiedEv(void* qthis); // 4
  // proto:  int QWidget::devType();
extern void _ZNK7QWidget7devTypeEv(void* qthis); // 4
  // proto:  void QWidget::clearFocus();
extern void _ZN7QWidget10clearFocusEv(void* qthis); // 4
  // proto:  void QWidget::setStyleSheet(const QString & styleSheet);
extern void _ZN7QWidget13setStyleSheetERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMaximumHeight(int maxh);
extern void _ZN7QWidget16setMaximumHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  int QWidget::x();
extern void _ZNK7QWidget1xEv(void* qthis); // 4
  // proto:  QPixmap QWidget::grab(const QRect & rectangle);
extern void _ZN7QWidget4grabERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QSize QWidget::maximumSize();
extern void _ZNK7QWidget11maximumSizeEv(void* qthis); // 4
  // proto:  QLocale QWidget::locale();
extern void _ZNK7QWidget6localeEv(void* qthis); // 4
  // proto:  QSize QWidget::minimumSize();
extern void _ZNK7QWidget11minimumSizeEv(void* qthis); // 4
  // proto:  QString QWidget::windowFilePath();
extern void _ZNK7QWidget14windowFilePathEv(void* qthis); // 4
  // proto:  void QWidget::setStatusTip(const QString & );
extern void _ZN7QWidget12setStatusTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  QWidget * QWidget::focusProxy();
extern void _ZNK7QWidget10focusProxyEv(void* qthis); // 4
  // proto:  void QWidget::createWinId();
extern void _ZN7QWidget11createWinIdEv(void* qthis); // 4
  // proto:  QPaintEngine * QWidget::paintEngine();
extern void _ZNK7QWidget11paintEngineEv(void* qthis); // 4
  // proto:  QByteArray QWidget::saveGeometry();
extern void _ZNK7QWidget12saveGeometryEv(void* qthis); // 4
  // proto:  QBackingStore * QWidget::backingStore();
extern void _ZNK7QWidget12backingStoreEv(void* qthis); // 4
  // proto:  void QWidget::setToolTipDuration(int msec);
extern void _ZN7QWidget18setToolTipDurationEi(void* qthis, int32_t arg0); // 4
  // proto:  QCursor QWidget::cursor();
extern void _ZNK7QWidget6cursorEv(void* qthis); // 4
  // proto:  const QPalette & QWidget::palette();
extern void _ZNK7QWidget7paletteEv(void* qthis); // 4
  // proto:  void QWidget::hide();
extern void _ZN7QWidget4hideEv(void* qthis); // 4
  // proto:  QString QWidget::windowTitle();
extern void _ZNK7QWidget11windowTitleEv(void* qthis); // 4
  // proto:  void QWidget::setFocusProxy(QWidget * );
extern void _ZN7QWidget13setFocusProxyEPS_(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setParent(QWidget * parent);
extern void _ZN7QWidget9setParentEPS_(void* qthis, void* arg0); // 4
  // proto:  QSizePolicy QWidget::sizePolicy();
extern void _ZNK7QWidget10sizePolicyEv(void* qthis); // 4
  // proto:  QRegion QWidget::visibleRegion();
extern void _ZNK7QWidget13visibleRegionEv(void* qthis); // 4
  // proto:  Qt::WindowFlags QWidget::windowFlags();
extern void _ZNK7QWidget11windowFlagsEv(void* qthis); // 2
  // proto:  void QWidget::setContentsMargins(int left, int top, int right, int bottom);
extern void _ZN7QWidget18setContentsMarginsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QWidget::setContentsMargins(const QMargins & margins);
extern void _ZN7QWidget18setContentsMarginsERK8QMargins(void* qthis, void* arg0); // 4
  // proto:  QGraphicsEffect * QWidget::graphicsEffect();
extern void _ZNK7QWidget14graphicsEffectEv(void* qthis); // 4
  // proto:  QPoint QWidget::mapFromParent(const QPoint & );
extern void _ZNK7QWidget13mapFromParentERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QString QWidget::windowRole();
extern void _ZNK7QWidget10windowRoleEv(void* qthis); // 4
  // proto:  QStyle * QWidget::style();
extern void _ZNK7QWidget5styleEv(void* qthis); // 4
  // proto:  QString QWidget::toolTip();
extern void _ZNK7QWidget7toolTipEv(void* qthis); // 4
  // proto:  int QWidget::maximumWidth();
extern void _ZNK7QWidget12maximumWidthEv(void* qthis); // 2
  // proto:  bool QWidget::isEnabledToTLW();
extern void _ZNK7QWidget14isEnabledToTLWEv(void* qthis); // 2
  // proto:  QIcon QWidget::windowIcon();
extern void _ZNK7QWidget10windowIconEv(void* qthis); // 4
  // proto:  bool QWidget::isMinimized();
extern void _ZNK7QWidget11isMinimizedEv(void* qthis); // 4
  // proto:  QRect QWidget::rect();
extern void _ZNK7QWidget4rectEv(void* qthis); // 2
  // proto:  void QWidget::raise();
extern void _ZN7QWidget5raiseEv(void* qthis); // 4
  // proto:  void QWidget::stackUnder(QWidget * );
extern void _ZN7QWidget10stackUnderEPS_(void* qthis, void* arg0); // 4
  // proto:  QWidget * QWidget::parentWidget();
extern void _ZNK7QWidget12parentWidgetEv(void* qthis); // 2
  // proto:  WId QWidget::effectiveWinId();
extern void _ZNK7QWidget14effectiveWinIdEv(void* qthis); // 4
  // proto:  void QWidget::setToolTip(const QString & );
extern void _ZN7QWidget10setToolTipERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setSizePolicy(QSizePolicy );
extern void _ZN7QWidget13setSizePolicyE11QSizePolicy(void* qthis, void* arg0); // 4
  // proto:  const QRect & QWidget::geometry();
extern void _ZNK7QWidget8geometryEv(void* qthis); // 2
  // proto:  QWindow * QWidget::windowHandle();
extern void _ZNK7QWidget12windowHandleEv(void* qthis); // 4
  // proto:  void QWidget::setAcceptDrops(bool on);
extern void _ZN7QWidget14setAcceptDropsEb(void* qthis, bool arg0); // 4
  // proto:  bool QWidget::isEnabledTo(const QWidget * );
extern void _ZNK7QWidget11isEnabledToEPKS_(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isVisible();
extern void _ZNK7QWidget9isVisibleEv(void* qthis); // 2
  // proto:  void QWidget::setWindowModified(bool );
extern void _ZN7QWidget17setWindowModifiedEb(void* qthis, bool arg0); // 4
  // proto:  QSize QWidget::size();
extern void _ZNK7QWidget4sizeEv(void* qthis); // 2
  // proto:  void QWidget::setMaximumWidth(int maxw);
extern void _ZN7QWidget15setMaximumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::addAction(QAction * action);
extern void _ZN7QWidget9addActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QRect QWidget::normalGeometry();
extern void _ZNK7QWidget14normalGeometryEv(void* qthis); // 4
  // proto:  void QWidget::setMinimumSize(int minw, int minh);
extern void _ZN7QWidget14setMinimumSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::setMinimumSize(const QSize & );
extern void _ZN7QWidget14setMinimumSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  int QWidget::heightForWidth(int );
extern void _ZNK7QWidget14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::setPalette(const QPalette & );
extern void _ZN7QWidget10setPaletteERK8QPalette(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isAncestorOf(const QWidget * child);
extern void _ZNK7QWidget12isAncestorOfEPKS_(void* qthis, void* arg0); // 4
  // proto:  void QWidget::clearMask();
extern void _ZN7QWidget9clearMaskEv(void* qthis); // 4
  // proto:  void QWidget::setWindowOpacity(qreal level);
extern void _ZN7QWidget16setWindowOpacityEd(void* qthis, double arg0); // 4
  // proto:  bool QWidget::isActiveWindow();
extern void _ZNK7QWidget14isActiveWindowEv(void* qthis); // 4
  // proto:  void QWidget::setFont(const QFont & );
extern void _ZN7QWidget7setFontERK5QFont(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setAccessibleDescription(const QString & description);
extern void _ZN7QWidget24setAccessibleDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMaximumSize(const QSize & );
extern void _ZN7QWidget14setMaximumSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QWidget::setMaximumSize(int maxw, int maxh);
extern void _ZN7QWidget14setMaximumSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::unsetLocale();
extern void _ZN7QWidget11unsetLocaleEv(void* qthis); // 4
  // proto:  bool QWidget::autoFillBackground();
extern void _ZNK7QWidget18autoFillBackgroundEv(void* qthis); // 4
  // proto:  void QWidget::unsetCursor();
extern void _ZN7QWidget11unsetCursorEv(void* qthis); // 4
  // proto:  WId QWidget::winId();
extern void _ZNK7QWidget5winIdEv(void* qthis); // 4
  // proto:  Qt::LayoutDirection QWidget::layoutDirection();
extern void _ZNK7QWidget15layoutDirectionEv(void* qthis); // 4
  // proto:  void QWidget::setShortcutEnabled(int id, bool enable);
extern void _ZN7QWidget18setShortcutEnabledEib(void* qthis, int32_t arg0, bool arg1); // 4
  // proto:  QSize QWidget::sizeIncrement();
extern void _ZNK7QWidget13sizeIncrementEv(void* qthis); // 4
  // proto:  void QWidget::setMouseTracking(bool enable);
extern void _ZN7QWidget16setMouseTrackingEb(void* qthis, bool arg0); // 2
  // proto:  QString QWidget::whatsThis();
extern void _ZNK7QWidget9whatsThisEv(void* qthis); // 4
  // proto:  int QWidget::width();
extern void _ZNK7QWidget5widthEv(void* qthis); // 2
  // proto:  QRect QWidget::childrenRect();
extern void _ZNK7QWidget12childrenRectEv(void* qthis); // 4
  // proto:  QString QWidget::windowIconText();
extern void _ZNK7QWidget14windowIconTextEv(void* qthis); // 4
  // proto:  int QWidget::toolTipDuration();
extern void _ZNK7QWidget15toolTipDurationEv(void* qthis); // 4
  // proto:  void QWidget::setEnabled(bool );
extern void _ZN7QWidget10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::showMaximized();
extern void _ZN7QWidget13showMaximizedEv(void* qthis); // 4
  // proto:  void QWidget::ensurePolished();
extern void _ZNK7QWidget14ensurePolishedEv(void* qthis); // 4
  // proto:  QString QWidget::statusTip();
extern void _ZNK7QWidget9statusTipEv(void* qthis); // 4
  // proto:  bool QWidget::acceptDrops();
extern void _ZNK7QWidget11acceptDropsEv(void* qthis); // 4
  // proto:  bool QWidget::isFullScreen();
extern void _ZNK7QWidget12isFullScreenEv(void* qthis); // 4
  // proto:  QWidget * QWidget::nextInFocusChain();
extern void _ZNK7QWidget16nextInFocusChainEv(void* qthis); // 4
  // proto:  QString QWidget::styleSheet();
extern void _ZNK7QWidget10styleSheetEv(void* qthis); // 4
  // proto:  void QWidget::setSizeIncrement(const QSize & );
extern void _ZN7QWidget16setSizeIncrementERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QWidget::setSizeIncrement(int w, int h);
extern void _ZN7QWidget16setSizeIncrementEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::grabKeyboard();
extern void _ZN7QWidget12grabKeyboardEv(void* qthis); // 4
  // proto:  QPoint QWidget::mapTo(const QWidget * , const QPoint & );
extern void _ZNK7QWidget5mapToEPKS_RK6QPoint(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QWidget::resize(int w, int h);
extern void _ZN7QWidget6resizeEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QWidget::resize(const QSize & );
extern void _ZN7QWidget6resizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setFocus();
extern void _ZN7QWidget8setFocusEv(void* qthis); // 2
  // proto:  QPoint QWidget::mapFromGlobal(const QPoint & );
extern void _ZNK7QWidget13mapFromGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QSize QWidget::frameSize();
extern void _ZNK7QWidget9frameSizeEv(void* qthis); // 4
  // proto:  void QWidget::releaseShortcut(int id);
extern void _ZN7QWidget15releaseShortcutEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QWidget::minimumSizeHint();
extern void _ZNK7QWidget15minimumSizeHintEv(void* qthis); // 4
  // proto:  bool QWidget::isHidden();
extern void _ZNK7QWidget8isHiddenEv(void* qthis); // 2
  // proto:  bool QWidget::hasHeightForWidth();
extern void _ZNK7QWidget17hasHeightForWidthEv(void* qthis); // 4
  // proto:  void QWidget::setFixedSize(const QSize & );
extern void _ZN7QWidget12setFixedSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setFixedSize(int w, int h);
extern void _ZN7QWidget12setFixedSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto: static void QWidget::setTabOrder(QWidget * , QWidget * );
extern void _ZN7QWidget11setTabOrderEPS_S0_(void* arg0, void* arg1); // 4
  // proto:  void QWidget::setWindowIcon(const QIcon & icon);
extern void _ZN7QWidget13setWindowIconERK5QIcon(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isLeftToRight();
extern void _ZNK7QWidget13isLeftToRightEv(void* qthis); // 2
  // proto:  bool QWidget::isRightToLeft();
extern void _ZNK7QWidget13isRightToLeftEv(void* qthis); // 2
  // proto:  void QWidget::setVisible(bool visible);
extern void _ZN7QWidget10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::setCursor(const QCursor & );
extern void _ZN7QWidget9setCursorERK7QCursor(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setFixedWidth(int w);
extern void _ZN7QWidget13setFixedWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QPoint QWidget::mapFrom(const QWidget * , const QPoint & );
extern void _ZNK7QWidget7mapFromEPKS_RK6QPoint(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QWidget::isWindow();
extern void _ZNK7QWidget8isWindowEv(void* qthis); // 2
  // proto:  bool QWidget::close();
extern void _ZN7QWidget5closeEv(void* qthis); // 4
  // proto:  QFontMetrics QWidget::fontMetrics();
extern void _ZNK7QWidget11fontMetricsEv(void* qthis); // 2
  // proto:  QMargins QWidget::contentsMargins();
extern void _ZNK7QWidget15contentsMarginsEv(void* qthis); // 4
  // proto:  void QWidget::~QWidget();
extern void _ZN7QWidgetD2Ev(void* qthis); // 4
  // proto:  void QWidget::setUpdatesEnabled(bool enable);
extern void _ZN7QWidget17setUpdatesEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::showMinimized();
extern void _ZN7QWidget13showMinimizedEv(void* qthis); // 4
  // proto:  QRegion QWidget::childrenRegion();
extern void _ZNK7QWidget14childrenRegionEv(void* qthis); // 4
  // proto:  Qt::WindowModality QWidget::windowModality();
extern void _ZNK7QWidget14windowModalityEv(void* qthis); // 4
  // proto:  QFontInfo QWidget::fontInfo();
extern void _ZNK7QWidget8fontInfoEv(void* qthis); // 2
  // proto:  void QWidget::releaseKeyboard();
extern void _ZN7QWidget15releaseKeyboardEv(void* qthis); // 4
  // proto:  QWidget * QWidget::previousInFocusChain();
extern void _ZNK7QWidget20previousInFocusChainEv(void* qthis); // 4
  // proto:  void QWidget::showFullScreen();
extern void _ZN7QWidget14showFullScreenEv(void* qthis); // 4
  // proto:  QSize QWidget::baseSize();
extern void _ZNK7QWidget8baseSizeEv(void* qthis); // 4
  // proto:  bool QWidget::hasMouseTracking();
extern void _ZNK7QWidget16hasMouseTrackingEv(void* qthis); // 2
  // proto:  bool QWidget::restoreGeometry(const QByteArray & geometry);
extern void _ZN7QWidget15restoreGeometryERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setDisabled(bool );
extern void _ZN7QWidget11setDisabledEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::showNormal();
extern void _ZN7QWidget10showNormalEv(void* qthis); // 4
  // proto:  void QWidget::setFixedHeight(int h);
extern void _ZN7QWidget14setFixedHeightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::setBaseSize(int basew, int baseh);
extern void _ZN7QWidget11setBaseSizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::setBaseSize(const QSize & );
extern void _ZN7QWidget11setBaseSizeERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QWidget::setAutoFillBackground(bool enabled);
extern void _ZN7QWidget21setAutoFillBackgroundEb(void* qthis, bool arg0); // 4
  // proto:  void QWidget::updateGeometry();
extern void _ZN7QWidget14updateGeometryEv(void* qthis); // 4
  // proto:  int QWidget::minimumWidth();
extern void _ZNK7QWidget12minimumWidthEv(void* qthis); // 2
  // proto:  int QWidget::maximumHeight();
extern void _ZNK7QWidget13maximumHeightEv(void* qthis); // 2
  // proto:  void QWidget::update();
extern void _ZN7QWidget6updateEv(void* qthis); // 4
  // proto:  void QWidget::update(int x, int y, int w, int h);
extern void _ZN7QWidget6updateEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  void QWidget::update(const QRect & );
extern void _ZN7QWidget6updateERK5QRect(void* qthis, void* arg0); // 4
  // proto:  void QWidget::update(const QRegion & );
extern void _ZN7QWidget6updateERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QWidget::adjustSize();
extern void _ZN7QWidget10adjustSizeEv(void* qthis); // 4
  // proto:  QGraphicsProxyWidget * QWidget::graphicsProxyWidget();
extern void _ZNK7QWidget19graphicsProxyWidgetEv(void* qthis); // 4
  // proto:  bool QWidget::isTopLevel();
extern void _ZNK7QWidget10isTopLevelEv(void* qthis); // 2
  // proto:  void QWidget::setMask(const QRegion & );
extern void _ZN7QWidget7setMaskERK7QRegion(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMask(const QBitmap & );
extern void _ZN7QWidget7setMaskERK7QBitmap(void* qthis, void* arg0); // 4
  // proto:  bool QWidget::isMaximized();
extern void _ZNK7QWidget11isMaximizedEv(void* qthis); // 4
  // proto:  void QWidget::lower();
extern void _ZN7QWidget5lowerEv(void* qthis); // 4
  // proto:  QPoint QWidget::mapToGlobal(const QPoint & );
extern void _ZNK7QWidget11mapToGlobalERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QWidget::setMinimumWidth(int minw);
extern void _ZN7QWidget15setMinimumWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QWidget::getContentsMargins(int * left, int * top, int * right, int * bottom);
extern void _ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  bool QWidget::underMouse();
extern void _ZNK7QWidget10underMouseEv(void* qthis); // 2
  // proto:  QRegion QWidget::mask();
extern void _ZNK7QWidget4maskEv(void* qthis); // 4
  // proto:  qreal QWidget::windowOpacity();
extern void _ZNK7QWidget13windowOpacityEv(void* qthis); // 4
  // proto:  int QWidget::height();
extern void _ZNK7QWidget6heightEv(void* qthis); // 2
  // proto:  QWidget * QWidget::nativeParentWidget();
extern void _ZNK7QWidget18nativeParentWidgetEv(void* qthis); // 4
  // proto:  QString QWidget::accessibleDescription();
extern void _ZNK7QWidget21accessibleDescriptionEv(void* qthis); // 4
  // proto:  int QWidget::y();
extern void _ZNK7QWidget1yEv(void* qthis); // 4
  // proto:  QPalette::ColorRole QWidget::backgroundRole();
extern void _ZNK7QWidget14backgroundRoleEv(void* qthis); // 4
  // proto:  void QWidget::scroll(int dx, int dy);
extern void _ZN7QWidget6scrollEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QWidget::scroll(int dx, int dy, const QRect & );
extern void _ZN7QWidget6scrollEiiRK5QRect(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QWidget::setWhatsThis(const QString & );
extern void _ZN7QWidget12setWhatsThisERK7QString(void* qthis, void* arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _windowIconChanged QWidget_windowIconChanged_signal;
//  _windowTitleChanged QWidget_windowTitleChanged_signal;
//  _customContextMenuRequested QWidget_customContextMenuRequested_signal;
//  _windowIconTextChanged QWidget_windowIconTextChanged_signal;
}

// class sizeof(QWidgetData)=1
type QWidgetData struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setWindowIconText(const class QString &)
func (this *QWidget) setWindowIconText(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget17setWindowIconTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowIconText", args)
  }

}

// show()
func (this *QWidget) show(args ...interface{}) () {
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
    C._ZN7QWidget4showEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "show", args)
  }

}

// hasFocus()
func (this *QWidget) hasFocus(args ...interface{}) () {
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
    C._ZNK7QWidget8hasFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "hasFocus", args)
  }

}

// insertAction(class QAction *, class QAction *)
func (this *QWidget) insertAction(args ...interface{}) () {
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
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QAction).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QWidget12insertActionEP7QActionS1_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "insertAction", args)
  }

}

// move(int, int)
func (this *QWidget) move_(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget4moveEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget4moveERK6QPoint
    // invoke: void move(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget4moveERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "move", args)
  }

}

// actions()
func (this *QWidget) actions(args ...interface{}) () {
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
    C._ZNK7QWidget7actionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "actions", args)
  }

}

// setStyle(class QStyle *)
func (this *QWidget) setStyle(args ...interface{}) () {
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
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget8setStyleEP6QStyle(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setStyle", args)
  }

}

// keyboardGrabber()
func (this *QWidget) keyboardGrabber_s(args ...interface{}) () {
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
    C._ZN7QWidget15keyboardGrabberEv()
  default:
    qtrt.ErrorResolve("QWidget", "keyboardGrabber", args)
  }

}

// setHidden(_Bool)
func (this *QWidget) setHidden(args ...interface{}) () {
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
    C._ZN7QWidget9setHiddenEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setHidden", args)
  }

}

// childAt(int, int)
func (this *QWidget) childAt(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QWidget7childAtEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK7QWidget7childAtERK6QPoint
    // invoke: QWidget * childAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget7childAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "childAt", args)
  }

}

// focusWidget()
func (this *QWidget) focusWidget(args ...interface{}) () {
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
    C._ZNK7QWidget11focusWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "focusWidget", args)
  }

}

// minimumHeight()
func (this *QWidget) minimumHeight(args ...interface{}) () {
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
    C._ZNK7QWidget13minimumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "minimumHeight", args)
  }

}

// font()
func (this *QWidget) font(args ...interface{}) () {
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
    C._ZNK7QWidget4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "font", args)
  }

}

// setLayout(class QLayout *)
func (this *QWidget) setLayout(args ...interface{}) () {
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
    var arg0 = args[0].(QLayout).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget9setLayoutEP7QLayout(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setLayout", args)
  }

}

// layout()
func (this *QWidget) layout(args ...interface{}) () {
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
    C._ZNK7QWidget6layoutEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "layout", args)
  }

}

// activateWindow()
func (this *QWidget) activateWindow(args ...interface{}) () {
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
    C._ZN7QWidget14activateWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "activateWindow", args)
  }

}

// unsetLayoutDirection()
func (this *QWidget) unsetLayoutDirection(args ...interface{}) () {
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
    C._ZN7QWidget20unsetLayoutDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "unsetLayoutDirection", args)
  }

}

// isEnabled()
func (this *QWidget) isEnabled(args ...interface{}) () {
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
    C._ZNK7QWidget9isEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isEnabled", args)
  }

}

// contextMenuPolicy()
func (this *QWidget) contextMenuPolicy(args ...interface{}) () {
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
    C._ZNK7QWidget17contextMenuPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "contextMenuPolicy", args)
  }

}

// setWindowFilePath(const class QString &)
func (this *QWidget) setWindowFilePath(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget17setWindowFilePathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowFilePath", args)
  }

}

// accessibleName()
func (this *QWidget) accessibleName(args ...interface{}) () {
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
    C._ZNK7QWidget14accessibleNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "accessibleName", args)
  }

}

// window()
func (this *QWidget) window(args ...interface{}) () {
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
    C._ZNK7QWidget6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "window", args)
  }

}

// internalWinId()
func (this *QWidget) internalWinId(args ...interface{}) () {
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
    C._ZNK7QWidget13internalWinIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "internalWinId", args)
  }

}

// setWindowTitle(const class QString &)
func (this *QWidget) setWindowTitle(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget14setWindowTitleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowTitle", args)
  }

}

// mouseGrabber()
func (this *QWidget) mouseGrabber_s(args ...interface{}) () {
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
    C._ZN7QWidget12mouseGrabberEv()
  default:
    qtrt.ErrorResolve("QWidget", "mouseGrabber", args)
  }

}

// isModal()
func (this *QWidget) isModal(args ...interface{}) () {
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
    C._ZNK7QWidget7isModalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isModal", args)
  }

}

// foregroundRole()
func (this *QWidget) foregroundRole(args ...interface{}) () {
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
    C._ZNK7QWidget14foregroundRoleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "foregroundRole", args)
  }

}

// setShortcutAutoRepeat(int, _Bool)
func (this *QWidget) setShortcutAutoRepeat(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget21setShortcutAutoRepeatEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setShortcutAutoRepeat", args)
  }

}

// setGraphicsEffect(class QGraphicsEffect *)
func (this *QWidget) setGraphicsEffect(args ...interface{}) () {
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
    var arg0 = args[0].(QGraphicsEffect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setGraphicsEffect", args)
  }

}

// setAccessibleName(const class QString &)
func (this *QWidget) setAccessibleName(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget17setAccessibleNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAccessibleName", args)
  }

}

// setMinimumHeight(int)
func (this *QWidget) setMinimumHeight(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget16setMinimumHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumHeight", args)
  }

}

// mapToParent(const class QPoint &)
func (this *QWidget) mapToParent(args ...interface{}) () {
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget11mapToParentERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "mapToParent", args)
  }

}

// setGeometry(int, int, int, int)
func (this *QWidget) setGeometry(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN7QWidget11setGeometryEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QWidget11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setGeometry", args)
  }

}

// pos()
func (this *QWidget) pos(args ...interface{}) () {
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
    C._ZNK7QWidget3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "pos", args)
  }

}

// sizeHint()
func (this *QWidget) sizeHint(args ...interface{}) () {
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
    C._ZNK7QWidget8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "sizeHint", args)
  }

}

// windowType()
func (this *QWidget) windowType(args ...interface{}) () {
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
    C._ZNK7QWidget10windowTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowType", args)
  }

}

// removeAction(class QAction *)
func (this *QWidget) removeAction(args ...interface{}) () {
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
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget12removeActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "removeAction", args)
  }

}

// grabMouse()
func (this *QWidget) grabMouse(args ...interface{}) () {
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
    C._ZN7QWidget9grabMouseEv(this.qclsinst)
  case 1:
    // invoke: _ZN7QWidget9grabMouseERK7QCursor
    // invoke: void grabMouse(const class QCursor &)
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget9grabMouseERK7QCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "grabMouse", args)
  }

}

// frameGeometry()
func (this *QWidget) frameGeometry(args ...interface{}) () {
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
    C._ZNK7QWidget13frameGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "frameGeometry", args)
  }

}

// setWindowRole(const class QString &)
func (this *QWidget) setWindowRole(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget13setWindowRoleERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowRole", args)
  }

}

// repaint(const class QRect &)
func (this *QWidget) repaint(args ...interface{}) () {
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
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget7repaintERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget7repaintEv
    // invoke: void repaint()
    C._ZN7QWidget7repaintEv(this.qclsinst)
  case 2:
    // invoke: _ZN7QWidget7repaintEiiii
    // invoke: void repaint(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN7QWidget7repaintEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 3:
    // invoke: _ZN7QWidget7repaintERK7QRegion
    // invoke: void repaint(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget7repaintERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "repaint", args)
  }

}

// inputMethodHints()
func (this *QWidget) inputMethodHints(args ...interface{}) () {
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
    C._ZNK7QWidget16inputMethodHintsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "inputMethodHints", args)
  }

}

// metaObject()
func (this *QWidget) metaObject(args ...interface{}) () {
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
    C._ZNK7QWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "metaObject", args)
  }

}

// topLevelWidget()
func (this *QWidget) topLevelWidget(args ...interface{}) () {
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
    C._ZNK7QWidget14topLevelWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "topLevelWidget", args)
  }

}

// releaseMouse()
func (this *QWidget) releaseMouse(args ...interface{}) () {
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
    C._ZN7QWidget12releaseMouseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "releaseMouse", args)
  }

}

// setLocale(const class QLocale &)
func (this *QWidget) setLocale(args ...interface{}) () {
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
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget9setLocaleERK7QLocale(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setLocale", args)
  }

}

// contentsRect()
func (this *QWidget) contentsRect(args ...interface{}) () {
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
    C._ZNK7QWidget12contentsRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "contentsRect", args)
  }

}

// focusPolicy()
func (this *QWidget) focusPolicy(args ...interface{}) () {
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
    C._ZNK7QWidget11focusPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "focusPolicy", args)
  }

}

// isVisibleTo(const class QWidget *)
func (this *QWidget) isVisibleTo(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget11isVisibleToEPKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "isVisibleTo", args)
  }

}

// updatesEnabled()
func (this *QWidget) updatesEnabled(args ...interface{}) () {
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
    C._ZNK7QWidget14updatesEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "updatesEnabled", args)
  }

}

// windowState()
func (this *QWidget) windowState(args ...interface{}) () {
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
    C._ZNK7QWidget11windowStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowState", args)
  }

}

// isWindowModified()
func (this *QWidget) isWindowModified(args ...interface{}) () {
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
    C._ZNK7QWidget16isWindowModifiedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isWindowModified", args)
  }

}

// devType()
func (this *QWidget) devType(args ...interface{}) () {
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
    C._ZNK7QWidget7devTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "devType", args)
  }

}

// clearFocus()
func (this *QWidget) clearFocus(args ...interface{}) () {
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
    C._ZN7QWidget10clearFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "clearFocus", args)
  }

}

// setStyleSheet(const class QString &)
func (this *QWidget) setStyleSheet(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget13setStyleSheetERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setStyleSheet", args)
  }

}

// setMaximumHeight(int)
func (this *QWidget) setMaximumHeight(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget16setMaximumHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumHeight", args)
  }

}

// x()
func (this *QWidget) x(args ...interface{}) () {
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
    C._ZNK7QWidget1xEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "x", args)
  }

}

// grab(const class QRect &)
func (this *QWidget) grab(args ...interface{}) () {
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
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget4grabERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "grab", args)
  }

}

// maximumSize()
func (this *QWidget) maximumSize(args ...interface{}) () {
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
    C._ZNK7QWidget11maximumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "maximumSize", args)
  }

}

// locale()
func (this *QWidget) locale(args ...interface{}) () {
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
    C._ZNK7QWidget6localeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "locale", args)
  }

}

// minimumSize()
func (this *QWidget) minimumSize(args ...interface{}) () {
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
    C._ZNK7QWidget11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "minimumSize", args)
  }

}

// windowFilePath()
func (this *QWidget) windowFilePath(args ...interface{}) () {
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
    C._ZNK7QWidget14windowFilePathEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowFilePath", args)
  }

}

// setStatusTip(const class QString &)
func (this *QWidget) setStatusTip(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget12setStatusTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setStatusTip", args)
  }

}

// focusProxy()
func (this *QWidget) focusProxy(args ...interface{}) () {
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
    C._ZNK7QWidget10focusProxyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "focusProxy", args)
  }

}

// createWinId()
func (this *QWidget) createWinId(args ...interface{}) () {
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
    C._ZN7QWidget11createWinIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "createWinId", args)
  }

}

// paintEngine()
func (this *QWidget) paintEngine(args ...interface{}) () {
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
    C._ZNK7QWidget11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "paintEngine", args)
  }

}

// saveGeometry()
func (this *QWidget) saveGeometry(args ...interface{}) () {
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
    C._ZNK7QWidget12saveGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "saveGeometry", args)
  }

}

// backingStore()
func (this *QWidget) backingStore(args ...interface{}) () {
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
    C._ZNK7QWidget12backingStoreEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "backingStore", args)
  }

}

// setToolTipDuration(int)
func (this *QWidget) setToolTipDuration(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget18setToolTipDurationEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setToolTipDuration", args)
  }

}

// cursor()
func (this *QWidget) cursor(args ...interface{}) () {
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
    C._ZNK7QWidget6cursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "cursor", args)
  }

}

// palette()
func (this *QWidget) palette(args ...interface{}) () {
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
    C._ZNK7QWidget7paletteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "palette", args)
  }

}

// hide()
func (this *QWidget) hide(args ...interface{}) () {
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
    C._ZN7QWidget4hideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "hide", args)
  }

}

// windowTitle()
func (this *QWidget) windowTitle(args ...interface{}) () {
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
    C._ZNK7QWidget11windowTitleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowTitle", args)
  }

}

// setFocusProxy(class QWidget *)
func (this *QWidget) setFocusProxy(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget13setFocusProxyEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFocusProxy", args)
  }

}

// setParent(class QWidget *)
func (this *QWidget) setParent(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget9setParentEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setParent", args)
  }

}

// sizePolicy()
func (this *QWidget) sizePolicy(args ...interface{}) () {
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
    C._ZNK7QWidget10sizePolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "sizePolicy", args)
  }

}

// visibleRegion()
func (this *QWidget) visibleRegion(args ...interface{}) () {
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
    C._ZNK7QWidget13visibleRegionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "visibleRegion", args)
  }

}

// windowFlags()
func (this *QWidget) windowFlags(args ...interface{}) () {
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
    C._ZNK7QWidget11windowFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowFlags", args)
  }

}

// setContentsMargins(int, int, int, int)
func (this *QWidget) setContentsMargins(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN7QWidget18setContentsMarginsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN7QWidget18setContentsMarginsERK8QMargins
    // invoke: void setContentsMargins(const class QMargins &)
    var arg0 = args[0].(QMargins).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget18setContentsMarginsERK8QMargins(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setContentsMargins", args)
  }

}

// graphicsEffect()
func (this *QWidget) graphicsEffect(args ...interface{}) () {
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
    C._ZNK7QWidget14graphicsEffectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "graphicsEffect", args)
  }

}

// mapFromParent(const class QPoint &)
func (this *QWidget) mapFromParent(args ...interface{}) () {
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget13mapFromParentERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "mapFromParent", args)
  }

}

// windowRole()
func (this *QWidget) windowRole(args ...interface{}) () {
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
    C._ZNK7QWidget10windowRoleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowRole", args)
  }

}

// style()
func (this *QWidget) style(args ...interface{}) () {
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
    C._ZNK7QWidget5styleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "style", args)
  }

}

// toolTip()
func (this *QWidget) toolTip(args ...interface{}) () {
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
    C._ZNK7QWidget7toolTipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "toolTip", args)
  }

}

// maximumWidth()
func (this *QWidget) maximumWidth(args ...interface{}) () {
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
    C._ZNK7QWidget12maximumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "maximumWidth", args)
  }

}

// isEnabledToTLW()
func (this *QWidget) isEnabledToTLW(args ...interface{}) () {
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
    C._ZNK7QWidget14isEnabledToTLWEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isEnabledToTLW", args)
  }

}

// windowIcon()
func (this *QWidget) windowIcon(args ...interface{}) () {
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
    C._ZNK7QWidget10windowIconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowIcon", args)
  }

}

// isMinimized()
func (this *QWidget) isMinimized(args ...interface{}) () {
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
    C._ZNK7QWidget11isMinimizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isMinimized", args)
  }

}

// rect()
func (this *QWidget) rect(args ...interface{}) () {
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
    C._ZNK7QWidget4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "rect", args)
  }

}

// raise()
func (this *QWidget) raise(args ...interface{}) () {
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
    C._ZN7QWidget5raiseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "raise", args)
  }

}

// stackUnder(class QWidget *)
func (this *QWidget) stackUnder(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget10stackUnderEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "stackUnder", args)
  }

}

// parentWidget()
func (this *QWidget) parentWidget(args ...interface{}) () {
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
    C._ZNK7QWidget12parentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "parentWidget", args)
  }

}

// effectiveWinId()
func (this *QWidget) effectiveWinId(args ...interface{}) () {
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
    C._ZNK7QWidget14effectiveWinIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "effectiveWinId", args)
  }

}

// setToolTip(const class QString &)
func (this *QWidget) setToolTip(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setToolTip", args)
  }

}

// setSizePolicy(class QSizePolicy)
func (this *QWidget) setSizePolicy(args ...interface{}) () {
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
    var arg0 = args[0].(QSizePolicy).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget13setSizePolicyE11QSizePolicy(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setSizePolicy", args)
  }

}

// geometry()
func (this *QWidget) geometry(args ...interface{}) () {
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
    C._ZNK7QWidget8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "geometry", args)
  }

}

// windowHandle()
func (this *QWidget) windowHandle(args ...interface{}) () {
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
    C._ZNK7QWidget12windowHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowHandle", args)
  }

}

// setAcceptDrops(_Bool)
func (this *QWidget) setAcceptDrops(args ...interface{}) () {
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
    C._ZN7QWidget14setAcceptDropsEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAcceptDrops", args)
  }

}

// isEnabledTo(const class QWidget *)
func (this *QWidget) isEnabledTo(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget11isEnabledToEPKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "isEnabledTo", args)
  }

}

// isVisible()
func (this *QWidget) isVisible(args ...interface{}) () {
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
    C._ZNK7QWidget9isVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isVisible", args)
  }

}

// setWindowModified(_Bool)
func (this *QWidget) setWindowModified(args ...interface{}) () {
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
    C._ZN7QWidget17setWindowModifiedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowModified", args)
  }

}

// size()
func (this *QWidget) size(args ...interface{}) () {
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
    C._ZNK7QWidget4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "size", args)
  }

}

// setMaximumWidth(int)
func (this *QWidget) setMaximumWidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget15setMaximumWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumWidth", args)
  }

}

// addAction(class QAction *)
func (this *QWidget) addAction(args ...interface{}) () {
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
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget9addActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "addAction", args)
  }

}

// normalGeometry()
func (this *QWidget) normalGeometry(args ...interface{}) () {
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
    C._ZNK7QWidget14normalGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "normalGeometry", args)
  }

}

// setMinimumSize(int, int)
func (this *QWidget) setMinimumSize(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget14setMinimumSizeEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget14setMinimumSizeERK5QSize
    // invoke: void setMinimumSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget14setMinimumSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumSize", args)
  }

}

// heightForWidth(int)
func (this *QWidget) heightForWidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "heightForWidth", args)
  }

}

// setPalette(const class QPalette &)
func (this *QWidget) setPalette(args ...interface{}) () {
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
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget10setPaletteERK8QPalette(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setPalette", args)
  }

}

// isAncestorOf(const class QWidget *)
func (this *QWidget) isAncestorOf(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget12isAncestorOfEPKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "isAncestorOf", args)
  }

}

// clearMask()
func (this *QWidget) clearMask(args ...interface{}) () {
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
    C._ZN7QWidget9clearMaskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "clearMask", args)
  }

}

// setWindowOpacity(qreal)
func (this *QWidget) setWindowOpacity(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget16setWindowOpacityEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowOpacity", args)
  }

}

// isActiveWindow()
func (this *QWidget) isActiveWindow(args ...interface{}) () {
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
    C._ZNK7QWidget14isActiveWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isActiveWindow", args)
  }

}

// setFont(const class QFont &)
func (this *QWidget) setFont(args ...interface{}) () {
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
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFont", args)
  }

}

// setAccessibleDescription(const class QString &)
func (this *QWidget) setAccessibleDescription(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget24setAccessibleDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAccessibleDescription", args)
  }

}

// setMaximumSize(const class QSize &)
func (this *QWidget) setMaximumSize(args ...interface{}) () {
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
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget14setMaximumSizeERK5QSize(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget14setMaximumSizeEii
    // invoke: void setMaximumSize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget14setMaximumSizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setMaximumSize", args)
  }

}

// unsetLocale()
func (this *QWidget) unsetLocale(args ...interface{}) () {
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
    C._ZN7QWidget11unsetLocaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "unsetLocale", args)
  }

}

// autoFillBackground()
func (this *QWidget) autoFillBackground(args ...interface{}) () {
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
    C._ZNK7QWidget18autoFillBackgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "autoFillBackground", args)
  }

}

// unsetCursor()
func (this *QWidget) unsetCursor(args ...interface{}) () {
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
    C._ZN7QWidget11unsetCursorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "unsetCursor", args)
  }

}

// winId()
func (this *QWidget) winId(args ...interface{}) () {
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
    C._ZNK7QWidget5winIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "winId", args)
  }

}

// layoutDirection()
func (this *QWidget) layoutDirection(args ...interface{}) () {
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
    C._ZNK7QWidget15layoutDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "layoutDirection", args)
  }

}

// setShortcutEnabled(int, _Bool)
func (this *QWidget) setShortcutEnabled(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget18setShortcutEnabledEib(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setShortcutEnabled", args)
  }

}

// sizeIncrement()
func (this *QWidget) sizeIncrement(args ...interface{}) () {
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
    C._ZNK7QWidget13sizeIncrementEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "sizeIncrement", args)
  }

}

// setMouseTracking(_Bool)
func (this *QWidget) setMouseTracking(args ...interface{}) () {
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
    C._ZN7QWidget16setMouseTrackingEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMouseTracking", args)
  }

}

// whatsThis()
func (this *QWidget) whatsThis(args ...interface{}) () {
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
    C._ZNK7QWidget9whatsThisEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "whatsThis", args)
  }

}

// width()
func (this *QWidget) width(args ...interface{}) () {
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
    C._ZNK7QWidget5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "width", args)
  }

}

// childrenRect()
func (this *QWidget) childrenRect(args ...interface{}) () {
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
    C._ZNK7QWidget12childrenRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "childrenRect", args)
  }

}

// windowIconText()
func (this *QWidget) windowIconText(args ...interface{}) () {
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
    C._ZNK7QWidget14windowIconTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowIconText", args)
  }

}

// toolTipDuration()
func (this *QWidget) toolTipDuration(args ...interface{}) () {
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
    C._ZNK7QWidget15toolTipDurationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "toolTipDuration", args)
  }

}

// setEnabled(_Bool)
func (this *QWidget) setEnabled(args ...interface{}) () {
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
    C._ZN7QWidget10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setEnabled", args)
  }

}

// showMaximized()
func (this *QWidget) showMaximized(args ...interface{}) () {
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
    C._ZN7QWidget13showMaximizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showMaximized", args)
  }

}

// ensurePolished()
func (this *QWidget) ensurePolished(args ...interface{}) () {
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
    C._ZNK7QWidget14ensurePolishedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "ensurePolished", args)
  }

}

// statusTip()
func (this *QWidget) statusTip(args ...interface{}) () {
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
    C._ZNK7QWidget9statusTipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "statusTip", args)
  }

}

// acceptDrops()
func (this *QWidget) acceptDrops(args ...interface{}) () {
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
    C._ZNK7QWidget11acceptDropsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "acceptDrops", args)
  }

}

// isFullScreen()
func (this *QWidget) isFullScreen(args ...interface{}) () {
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
    C._ZNK7QWidget12isFullScreenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isFullScreen", args)
  }

}

// nextInFocusChain()
func (this *QWidget) nextInFocusChain(args ...interface{}) () {
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
    C._ZNK7QWidget16nextInFocusChainEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "nextInFocusChain", args)
  }

}

// styleSheet()
func (this *QWidget) styleSheet(args ...interface{}) () {
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
    C._ZNK7QWidget10styleSheetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "styleSheet", args)
  }

}

// setSizeIncrement(const class QSize &)
func (this *QWidget) setSizeIncrement(args ...interface{}) () {
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
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget16setSizeIncrementERK5QSize(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget16setSizeIncrementEii
    // invoke: void setSizeIncrement(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget16setSizeIncrementEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setSizeIncrement", args)
  }

}

// grabKeyboard()
func (this *QWidget) grabKeyboard(args ...interface{}) () {
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
    C._ZN7QWidget12grabKeyboardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "grabKeyboard", args)
  }

}

// mapTo(const class QWidget *, const class QPoint &)
func (this *QWidget) mapTo(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QWidget5mapToEPKS_RK6QPoint(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "mapTo", args)
  }

}

// resize(int, int)
func (this *QWidget) resize(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget6resizeEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget6resizeERK5QSize
    // invoke: void resize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget6resizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "resize", args)
  }

}

// setFocus()
func (this *QWidget) setFocus(args ...interface{}) () {
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
    C._ZN7QWidget8setFocusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "setFocus", args)
  }

}

// mapFromGlobal(const class QPoint &)
func (this *QWidget) mapFromGlobal(args ...interface{}) () {
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget13mapFromGlobalERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "mapFromGlobal", args)
  }

}

// frameSize()
func (this *QWidget) frameSize(args ...interface{}) () {
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
    C._ZNK7QWidget9frameSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "frameSize", args)
  }

}

// releaseShortcut(int)
func (this *QWidget) releaseShortcut(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget15releaseShortcutEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "releaseShortcut", args)
  }

}

// minimumSizeHint()
func (this *QWidget) minimumSizeHint(args ...interface{}) () {
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
    C._ZNK7QWidget15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "minimumSizeHint", args)
  }

}

// isHidden()
func (this *QWidget) isHidden(args ...interface{}) () {
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
    C._ZNK7QWidget8isHiddenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isHidden", args)
  }

}

// hasHeightForWidth()
func (this *QWidget) hasHeightForWidth(args ...interface{}) () {
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
    C._ZNK7QWidget17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "hasHeightForWidth", args)
  }

}

// setFixedSize(const class QSize &)
func (this *QWidget) setFixedSize(args ...interface{}) () {
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
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget12setFixedSizeERK5QSize(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget12setFixedSizeEii
    // invoke: void setFixedSize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget12setFixedSizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setFixedSize", args)
  }

}

// setTabOrder(class QWidget *, class QWidget *)
func (this *QWidget) setTabOrder_s(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QWidget11setTabOrderEPS_S0_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "setTabOrder", args)
  }

}

// setWindowIcon(const class QIcon &)
func (this *QWidget) setWindowIcon(args ...interface{}) () {
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
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget13setWindowIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWindowIcon", args)
  }

}

// isLeftToRight()
func (this *QWidget) isLeftToRight(args ...interface{}) () {
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
    C._ZNK7QWidget13isLeftToRightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isLeftToRight", args)
  }

}

// isRightToLeft()
func (this *QWidget) isRightToLeft(args ...interface{}) () {
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
    C._ZNK7QWidget13isRightToLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isRightToLeft", args)
  }

}

// setVisible(_Bool)
func (this *QWidget) setVisible(args ...interface{}) () {
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
    C._ZN7QWidget10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setVisible", args)
  }

}

// setCursor(const class QCursor &)
func (this *QWidget) setCursor(args ...interface{}) () {
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
    var arg0 = args[0].(QCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget9setCursorERK7QCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setCursor", args)
  }

}

// setFixedWidth(int)
func (this *QWidget) setFixedWidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget13setFixedWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFixedWidth", args)
  }

}

// mapFrom(const class QWidget *, const class QPoint &)
func (this *QWidget) mapFrom(args ...interface{}) () {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QWidget7mapFromEPKS_RK6QPoint(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QWidget", "mapFrom", args)
  }

}

// isWindow()
func (this *QWidget) isWindow(args ...interface{}) () {
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
    C._ZNK7QWidget8isWindowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isWindow", args)
  }

}

// close()
func (this *QWidget) close(args ...interface{}) () {
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
    C._ZN7QWidget5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "close", args)
  }

}

// fontMetrics()
func (this *QWidget) fontMetrics(args ...interface{}) () {
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
    C._ZNK7QWidget11fontMetricsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "fontMetrics", args)
  }

}

// contentsMargins()
func (this *QWidget) contentsMargins(args ...interface{}) () {
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
    C._ZNK7QWidget15contentsMarginsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "contentsMargins", args)
  }

}

// ~QWidget()
func (this *QWidget) FreeQWidget(args ...interface{}) () {
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
    C._ZN7QWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "~QWidget", args)
  }

}

// setUpdatesEnabled(_Bool)
func (this *QWidget) setUpdatesEnabled(args ...interface{}) () {
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
    C._ZN7QWidget17setUpdatesEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setUpdatesEnabled", args)
  }

}

// showMinimized()
func (this *QWidget) showMinimized(args ...interface{}) () {
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
    C._ZN7QWidget13showMinimizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showMinimized", args)
  }

}

// childrenRegion()
func (this *QWidget) childrenRegion(args ...interface{}) () {
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
    C._ZNK7QWidget14childrenRegionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "childrenRegion", args)
  }

}

// windowModality()
func (this *QWidget) windowModality(args ...interface{}) () {
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
    C._ZNK7QWidget14windowModalityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowModality", args)
  }

}

// fontInfo()
func (this *QWidget) fontInfo(args ...interface{}) () {
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
    C._ZNK7QWidget8fontInfoEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "fontInfo", args)
  }

}

// releaseKeyboard()
func (this *QWidget) releaseKeyboard(args ...interface{}) () {
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
    C._ZN7QWidget15releaseKeyboardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "releaseKeyboard", args)
  }

}

// previousInFocusChain()
func (this *QWidget) previousInFocusChain(args ...interface{}) () {
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
    C._ZNK7QWidget20previousInFocusChainEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "previousInFocusChain", args)
  }

}

// showFullScreen()
func (this *QWidget) showFullScreen(args ...interface{}) () {
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
    C._ZN7QWidget14showFullScreenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showFullScreen", args)
  }

}

// baseSize()
func (this *QWidget) baseSize(args ...interface{}) () {
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
    C._ZNK7QWidget8baseSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "baseSize", args)
  }

}

// hasMouseTracking()
func (this *QWidget) hasMouseTracking(args ...interface{}) () {
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
    C._ZNK7QWidget16hasMouseTrackingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "hasMouseTracking", args)
  }

}

// restoreGeometry(const class QByteArray &)
func (this *QWidget) restoreGeometry(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget15restoreGeometryERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "restoreGeometry", args)
  }

}

// setDisabled(_Bool)
func (this *QWidget) setDisabled(args ...interface{}) () {
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
    C._ZN7QWidget11setDisabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setDisabled", args)
  }

}

// showNormal()
func (this *QWidget) showNormal(args ...interface{}) () {
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
    C._ZN7QWidget10showNormalEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "showNormal", args)
  }

}

// setFixedHeight(int)
func (this *QWidget) setFixedHeight(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget14setFixedHeightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setFixedHeight", args)
  }

}

// setBaseSize(int, int)
func (this *QWidget) setBaseSize(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget11setBaseSizeEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget11setBaseSizeERK5QSize
    // invoke: void setBaseSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget11setBaseSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setBaseSize", args)
  }

}

// setAutoFillBackground(_Bool)
func (this *QWidget) setAutoFillBackground(args ...interface{}) () {
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
    C._ZN7QWidget21setAutoFillBackgroundEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setAutoFillBackground", args)
  }

}

// updateGeometry()
func (this *QWidget) updateGeometry(args ...interface{}) () {
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
    C._ZN7QWidget14updateGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "updateGeometry", args)
  }

}

// minimumWidth()
func (this *QWidget) minimumWidth(args ...interface{}) () {
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
    C._ZNK7QWidget12minimumWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "minimumWidth", args)
  }

}

// maximumHeight()
func (this *QWidget) maximumHeight(args ...interface{}) () {
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
    C._ZNK7QWidget13maximumHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "maximumHeight", args)
  }

}

// update()
func (this *QWidget) update(args ...interface{}) () {
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
    C._ZN7QWidget6updateEv(this.qclsinst)
  case 1:
    // invoke: _ZN7QWidget6updateEiiii
    // invoke: void update(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN7QWidget6updateEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZN7QWidget6updateERK5QRect
    // invoke: void update(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget6updateERK5QRect(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN7QWidget6updateERK7QRegion
    // invoke: void update(const class QRegion &)
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget6updateERK7QRegion(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "update", args)
  }

}

// adjustSize()
func (this *QWidget) adjustSize(args ...interface{}) () {
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
    C._ZN7QWidget10adjustSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "adjustSize", args)
  }

}

// graphicsProxyWidget()
func (this *QWidget) graphicsProxyWidget(args ...interface{}) () {
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
    C._ZNK7QWidget19graphicsProxyWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "graphicsProxyWidget", args)
  }

}

// isTopLevel()
func (this *QWidget) isTopLevel(args ...interface{}) () {
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
    C._ZNK7QWidget10isTopLevelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isTopLevel", args)
  }

}

// setMask(const class QRegion &)
func (this *QWidget) setMask(args ...interface{}) () {
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
    var arg0 = args[0].(QRegion).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget7setMaskERK7QRegion(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QWidget7setMaskERK7QBitmap
    // invoke: void setMask(const class QBitmap &)
    var arg0 = args[0].(QBitmap).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget7setMaskERK7QBitmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMask", args)
  }

}

// isMaximized()
func (this *QWidget) isMaximized(args ...interface{}) () {
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
    C._ZNK7QWidget11isMaximizedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "isMaximized", args)
  }

}

// lower()
func (this *QWidget) lower(args ...interface{}) () {
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
    C._ZN7QWidget5lowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "lower", args)
  }

}

// mapToGlobal(const class QPoint &)
func (this *QWidget) mapToGlobal(args ...interface{}) () {
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
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QWidget11mapToGlobalERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "mapToGlobal", args)
  }

}

// setMinimumWidth(int)
func (this *QWidget) setMinimumWidth(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QWidget15setMinimumWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setMinimumWidth", args)
  }

}

// getContentsMargins(int *, int *, int *, int *)
func (this *QWidget) getContentsMargins(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QWidget", "getContentsMargins", args)
  }

}

// underMouse()
func (this *QWidget) underMouse(args ...interface{}) () {
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
    C._ZNK7QWidget10underMouseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "underMouse", args)
  }

}

// mask()
func (this *QWidget) mask(args ...interface{}) () {
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
    C._ZNK7QWidget4maskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "mask", args)
  }

}

// windowOpacity()
func (this *QWidget) windowOpacity(args ...interface{}) () {
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
    C._ZNK7QWidget13windowOpacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "windowOpacity", args)
  }

}

// height()
func (this *QWidget) height(args ...interface{}) () {
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
    C._ZNK7QWidget6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "height", args)
  }

}

// nativeParentWidget()
func (this *QWidget) nativeParentWidget(args ...interface{}) () {
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
    C._ZNK7QWidget18nativeParentWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "nativeParentWidget", args)
  }

}

// accessibleDescription()
func (this *QWidget) accessibleDescription(args ...interface{}) () {
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
    C._ZNK7QWidget21accessibleDescriptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "accessibleDescription", args)
  }

}

// y()
func (this *QWidget) y(args ...interface{}) () {
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
    C._ZNK7QWidget1yEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "y", args)
  }

}

// backgroundRole()
func (this *QWidget) backgroundRole(args ...interface{}) () {
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
    C._ZNK7QWidget14backgroundRoleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidget", "backgroundRole", args)
  }

}

// scroll(int, int)
func (this *QWidget) scroll(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QWidget6scrollEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QWidget6scrollEiiRK5QRect
    // invoke: void scroll(int, int, const class QRect &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRect).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN7QWidget6scrollEiiRK5QRect(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QWidget", "scroll", args)
  }

}

// setWhatsThis(const class QString &)
func (this *QWidget) setWhatsThis(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QWidget12setWhatsThisERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidget", "setWhatsThis", args)
  }

}

// <= body block end

