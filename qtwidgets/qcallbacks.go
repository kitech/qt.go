//  package block begin
package qtwidgets

/*
#include <stdint.h>
#include <stdbool.h>
//  package block end

//  extern block begin
extern void callback_ZN7QWidget5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget14mouseMoveEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget10wheelEventEP11QWheelEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget15keyReleaseEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget10enterEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget10leaveEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget9moveEventEP10QMoveEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget10closeEventEP11QCloseEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget11tabletEventEP12QTabletEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget11actionEventEP12QActionEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget14dragEnterEventEP15QDragEnterEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget9dropEventEP10QDropEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget9showEventEP10QShowEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget9hideEventEP10QHideEvent(void* fnptr , void* event);
extern void callback_ZN7QWidget11nativeEventERK10QByteArrayPvPl(void* fnptr , void* eventType, void* message, void* result);
extern void callback_ZN7QWidget11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE(void* fnptr , int arg0);
extern void callback_ZNK7QWidget11initPainterEP8QPainter(void* fnptr , void* painter);
extern void callback_ZNK7QWidget10redirectedEP6QPoint(void* fnptr , void* offset);
extern void callback_ZNK7QWidget13sharedPainterEv(void* fnptr );
extern void callback_ZN7QWidget16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* arg0);
extern void callback_ZN7QWidget16updateMicroFocusEv(void* fnptr );
extern void callback_ZN7QWidget6createEybb(void* fnptr , uint64_t arg0, bool initializeWindow, bool destroyOldWindow);
extern void callback_ZN7QWidget7destroyEbb(void* fnptr , bool destroyWindow, bool destroySubWindows);
extern void callback_ZN7QWidget18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN7QWidget14focusNextChildEv(void* fnptr );
extern void callback_ZN7QWidget18focusPreviousChildEv(void* fnptr );
extern void callback_ZN15QAbstractButton10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZNK15QAbstractButton9hitButtonERK6QPoint(void* fnptr , void* pos);
extern void callback_ZN15QAbstractButton13checkStateSetEv(void* fnptr );
extern void callback_ZN15QAbstractButton14nextCheckStateEv(void* fnptr );
extern void callback_ZN15QAbstractButton5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton13keyPressEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton15mousePressEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton12focusInEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton13focusOutEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton11changeEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractButton10timerEventEP11QTimerEvent(void* fnptr , void* e);
extern void callback_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox9hideEventEP10QHideEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZN16QAbstractSpinBox9showEventEP10QShowEvent(void* fnptr , void* event);
extern void callback_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox(void* fnptr , void* option);
extern void callback_ZNK16QAbstractSpinBox8lineEditEv(void* fnptr );
extern void callback_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit(void* fnptr , void* edit);
extern void callback_ZNK16QAbstractSpinBox11stepEnabledEv(void* fnptr );
extern void callback_ZN15QAbstractSlider5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii(void* fnptr , int action, int thresholdTime, int repeatTime);
extern void callback_ZNK15QAbstractSlider12repeatActionEv(void* fnptr );
extern void callback_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE(void* fnptr , int change);
extern void callback_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent(void* fnptr , void* ev);
extern void callback_ZN15QAbstractSlider10timerEventEP11QTimerEvent(void* fnptr , void* arg0);
extern void callback_ZN15QAbstractSlider10wheelEventEP11QWheelEvent(void* fnptr , void* e);
extern void callback_ZN15QAbstractSlider11changeEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN7QSlider10paintEventEP11QPaintEvent(void* fnptr , void* ev);
extern void callback_ZN7QSlider15mousePressEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN7QSlider17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN7QSlider14mouseMoveEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider(void* fnptr , void* option);
extern void callback_ZNK7QTabBar11tabSizeHintEi(void* fnptr , int index);
extern void callback_ZNK7QTabBar18minimumTabSizeHintEi(void* fnptr , int index);
extern void callback_ZN7QTabBar11tabInsertedEi(void* fnptr , int index);
extern void callback_ZN7QTabBar10tabRemovedEi(void* fnptr , int index);
extern void callback_ZN7QTabBar15tabLayoutChangeEv(void* fnptr );
extern void callback_ZN7QTabBar5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar9hideEventEP10QHideEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar10wheelEventEP11QWheelEvent(void* fnptr , void* event);
extern void callback_ZN7QTabBar13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN7QTabBar10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZNK7QTabBar15initStyleOptionEP15QStyleOptionTabi(void* fnptr , void* option, int tabIndex);
extern void callback_ZN10QTabWidget11tabInsertedEi(void* fnptr , int index);
extern void callback_ZN10QTabWidget10tabRemovedEi(void* fnptr , int index);
extern void callback_ZN10QTabWidget9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN10QTabWidget11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN10QTabWidget13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN10QTabWidget10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN10QTabWidget9setTabBarEP7QTabBar(void* fnptr , void* arg0);
extern void callback_ZN10QTabWidget11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN10QTabWidget5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame(void* fnptr , void* option);
extern void callback_ZN11QRubberBand5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QRubberBand10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN11QRubberBand11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN11QRubberBand9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN11QRubberBand11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN11QRubberBand9moveEventEP10QMoveEvent(void* fnptr , void* arg0);
extern void callback_ZNK11QRubberBand15initStyleOptionEP22QStyleOptionRubberBand(void* fnptr , void* option);
extern void callback_ZN6QFrame5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN6QFrame10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN6QFrame11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN6QFrame9drawFrameEP8QPainter(void* fnptr , void* arg0);
extern void callback_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame(void* fnptr , void* option);
extern void callback_ZN21QStyleOptionFocusRectC1Ei(void* fnptr , int version);
extern void callback_ZN17QStyleOptionFrameC1Ei(void* fnptr , int version);
extern void callback_ZN26QStyleOptionTabWidgetFrameC1Ei(void* fnptr , int version);
extern void callback_ZN22QStyleOptionTabBarBaseC1Ei(void* fnptr , int version);
extern void callback_ZN18QStyleOptionHeaderC1Ei(void* fnptr , int version);
extern void callback_ZN18QStyleOptionButtonC1Ei(void* fnptr , int version);
extern void callback_ZN15QStyleOptionTabC1Ei(void* fnptr , int version);
extern void callback_ZN19QStyleOptionToolBarC1Ei(void* fnptr , int version);
extern void callback_ZN23QStyleOptionProgressBarC1Ei(void* fnptr , int version);
extern void callback_ZN20QStyleOptionMenuItemC1Ei(void* fnptr , int version);
extern void callback_ZN22QStyleOptionDockWidgetC1Ei(void* fnptr , int version);
extern void callback_ZN20QStyleOptionViewItemC1Ei(void* fnptr , int version);
extern void callback_ZN19QStyleOptionToolBoxC1Ei(void* fnptr , int version);
extern void callback_ZN22QStyleOptionRubberBandC1Ei(void* fnptr , int version);
extern void callback_ZN18QStyleOptionSliderC1Ei(void* fnptr , int version);
extern void callback_ZN19QStyleOptionSpinBoxC1Ei(void* fnptr , int version);
extern void callback_ZN22QStyleOptionToolButtonC1Ei(void* fnptr , int version);
extern void callback_ZN20QStyleOptionComboBoxC1Ei(void* fnptr , int version);
extern void callback_ZN20QStyleOptionTitleBarC1Ei(void* fnptr , int version);
extern void callback_ZN20QStyleOptionGroupBoxC1Ei(void* fnptr , int version);
extern void callback_ZN20QStyleOptionSizeGripC1Ei(void* fnptr , int version);
extern void callback_ZN24QStyleOptionGraphicsItemC1Ei(void* fnptr , int version);
extern void callback_ZN19QAbstractScrollArea18setViewportMarginsEiiii(void* fnptr , int left, int top, int right, int bottom);
extern void callback_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins(void* fnptr , void* margins);
extern void callback_ZNK19QAbstractScrollArea15viewportMarginsEv(void* fnptr );
extern void callback_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent(void* fnptr , void* arg0, void* arg1);
extern void callback_ZN19QAbstractScrollArea5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea13viewportEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea9dropEventEP10QDropEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN19QAbstractScrollArea16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZNK19QAbstractScrollArea16viewportSizeHintEv(void* fnptr );
extern void callback_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_(void* fnptr , void* selected, void* deselected);
extern void callback_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_(void* fnptr , void* current, void* previous);
extern void callback_ZN17QAbstractItemView16updateEditorDataEv(void* fnptr );
extern void callback_ZN17QAbstractItemView22updateEditorGeometriesEv(void* fnptr );
extern void callback_ZN17QAbstractItemView16updateGeometriesEv(void* fnptr );
extern void callback_ZN17QAbstractItemView23verticalScrollbarActionEi(void* fnptr , int action);
extern void callback_ZN17QAbstractItemView25horizontalScrollbarActionEi(void* fnptr , int action);
extern void callback_ZN17QAbstractItemView29verticalScrollbarValueChangedEi(void* fnptr , int value);
extern void callback_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi(void* fnptr , int value);
extern void callback_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE(void* fnptr , void* editor, int hint);
extern void callback_ZN17QAbstractItemView10commitDataEP7QWidget(void* fnptr , void* editor);
extern void callback_ZN17QAbstractItemView15editorDestroyedEP7QObject(void* fnptr , void* editor);
extern void callback_ZN17QAbstractItemView25setHorizontalStepsPerItemEi(void* fnptr , int steps);
extern void callback_ZNK17QAbstractItemView22horizontalStepsPerItemEv(void* fnptr );
extern void callback_ZN17QAbstractItemView23setVerticalStepsPerItemEi(void* fnptr , int steps);
extern void callback_ZNK17QAbstractItemView20verticalStepsPerItemEv(void* fnptr );
extern void callback_ZNK17QAbstractItemView16horizontalOffsetEv(void* fnptr );
extern void callback_ZNK17QAbstractItemView14verticalOffsetEv(void* fnptr );
extern void callback_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(void* fnptr , void* rect, int command);
extern void callback_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection(void* fnptr , void* selection);
extern void callback_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent(void* fnptr , void* index, int trigger, void* event);
extern void callback_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent(void* fnptr , void* index, void* event);
extern void callback_ZNK17QAbstractItemView11viewOptionsEv(void* fnptr );
extern void callback_ZNK17QAbstractItemView5stateEv(void* fnptr );
extern void callback_ZN17QAbstractItemView8setStateENS_5StateE(void* fnptr , int state);
extern void callback_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv(void* fnptr );
extern void callback_ZN17QAbstractItemView25executeDelayedItemsLayoutEv(void* fnptr );
extern void callback_ZN17QAbstractItemView14setDirtyRegionERK7QRegion(void* fnptr , void* region);
extern void callback_ZN17QAbstractItemView17scrollDirtyRegionEii(void* fnptr , int dx, int dy);
extern void callback_ZNK17QAbstractItemView17dirtyRegionOffsetEv(void* fnptr );
extern void callback_ZN17QAbstractItemView15startAutoScrollEv(void* fnptr );
extern void callback_ZN17QAbstractItemView14stopAutoScrollEv(void* fnptr );
extern void callback_ZN17QAbstractItemView12doAutoScrollEv(void* fnptr );
extern void callback_ZN17QAbstractItemView18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN17QAbstractItemView5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView13viewportEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView9dropEventEP10QDropEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* event);
extern void callback_ZNK17QAbstractItemView21dropIndicatorPositionEv(void* fnptr );
extern void callback_ZNK17QAbstractItemView16viewportSizeHintEv(void* fnptr );
extern void callback_ZN17QAccessibleWidgetD1Ev(void* fnptr );
extern void callback_ZNK17QAccessibleWidget6widgetEv(void* fnptr );
extern void callback_ZNK17QAccessibleWidget12parentObjectEv(void* fnptr );
extern void callback_ZN17QAccessibleWidget20addControllingSignalERK7QString(void* fnptr , void* signal);
extern void callback_ZN7QAction5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN12QApplication5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN7QLayout11widgetEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN7QLayout10childEventEP11QChildEvent(void* fnptr , void* e);
extern void callback_ZN7QLayout14addChildLayoutEPS_(void* fnptr , void* l);
extern void callback_ZN7QLayout14addChildWidgetEP7QWidget(void* fnptr , void* w);
extern void callback_ZN7QLayout11adoptLayoutEPS_(void* fnptr , void* layout);
extern void callback_ZNK7QLayout13alignmentRectERK5QRect(void* fnptr , void* arg0);
extern void callback_ZN11QGridLayout7addItemEP11QLayoutItem(void* fnptr , void* arg0);
extern void callback_ZN15QCalendarWidget5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent(void* fnptr , void* watched, void* event);
extern void callback_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN15QCalendarWidget11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate(void* fnptr , void* painter, void* rect, void* date);
extern void callback_ZN15QCalendarWidget10updateCellERK5QDate(void* fnptr , void* date);
extern void callback_ZN15QCalendarWidget11updateCellsEv(void* fnptr );
extern void callback_ZN9QCheckBox5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZNK9QCheckBox9hitButtonERK6QPoint(void* fnptr , void* pos);
extern void callback_ZN9QCheckBox13checkStateSetEv(void* fnptr );
extern void callback_ZN9QCheckBox14nextCheckStateEv(void* fnptr );
extern void callback_ZN9QCheckBox10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton(void* fnptr , void* option);
extern void callback_ZN7QDialog13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN7QDialog10closeEventEP11QCloseEvent(void* fnptr , void* arg0);
extern void callback_ZN7QDialog9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN7QDialog11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN7QDialog16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* arg0);
extern void callback_ZN7QDialog11eventFilterEP7QObjectP6QEvent(void* fnptr , void* arg0, void* arg1);
extern void callback_ZN7QDialog14adjustPositionEP7QWidget(void* fnptr , void* arg0);
extern void callback_ZN12QColorDialog11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN12QColorDialog4doneEi(void* fnptr , int result);
extern void callback_ZNK11QColumnView13isIndexHiddenERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN11QColumnView11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN11QColumnView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(void* fnptr , void* rect, int command);
extern void callback_ZNK11QColumnView24visualRegionForSelectionERK14QItemSelection(void* fnptr , void* selection);
extern void callback_ZNK11QColumnView16horizontalOffsetEv(void* fnptr );
extern void callback_ZNK11QColumnView14verticalOffsetEv(void* fnptr );
extern void callback_ZN11QColumnView12rowsInsertedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZN11QColumnView14currentChangedERK11QModelIndexS2_(void* fnptr , void* current, void* previous);
extern void callback_ZN11QColumnView16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN11QColumnView12createColumnERK11QModelIndex(void* fnptr , void* rootIndex);
extern void callback_ZNK11QColumnView16initializeColumnEP17QAbstractItemView(void* fnptr , void* column);
extern void callback_ZN9QComboBox12focusInEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox13focusOutEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox11changeEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox11resizeEventEP12QResizeEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox9showEventEP10QShowEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox9hideEventEP10QHideEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox15mousePressEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox13keyPressEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox15keyReleaseEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox10wheelEventEP11QWheelEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* e);
extern void callback_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* arg0);
extern void callback_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox(void* fnptr , void* option);
extern void callback_ZN11QPushButton5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QPushButton10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN11QPushButton13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN11QPushButton12focusInEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZN11QPushButton13focusOutEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton(void* fnptr , void* option);
extern void callback_ZNK18QCommandLinkButton8sizeHintEv(void* fnptr );
extern void callback_ZNK18QCommandLinkButton14heightForWidthEi(void* fnptr , int arg0);
extern void callback_ZNK18QCommandLinkButton15minimumSizeHintEv(void* fnptr );
extern void callback_ZN18QCommandLinkButton5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN18QCommandLinkButton10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN10QCompleter11eventFilterEP7QObjectP6QEvent(void* fnptr , void* o, void* e);
extern void callback_ZN10QCompleter5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent(void* fnptr , void* event);
extern void callback_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN13QDateTimeEdit18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZNK13QDateTimeEdit8validateER7QStringRi(void* fnptr , void* input, void* pos);
extern void callback_ZNK13QDateTimeEdit5fixupER7QString(void* fnptr , void* input);
extern void callback_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString(void* fnptr , void* text);
extern void callback_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime(void* fnptr , void* dt);
extern void callback_ZNK13QDateTimeEdit11stepEnabledEv(void* fnptr );
extern void callback_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QDateTimeEdit10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox(void* fnptr , void* option);
extern void callback_ZN13QDateTimeEditC1ERK8QVariantNS0_4TypeEP7QWidget(void* fnptr , void* val, int parserType, void* parent);
extern void callback_ZN14QDesktopWidget11resizeEventEP12QResizeEvent(void* fnptr , void* e);
extern void callback_ZN5QDial5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN5QDial11resizeEventEP12QResizeEvent(void* fnptr , void* re);
extern void callback_ZN5QDial10paintEventEP11QPaintEvent(void* fnptr , void* pe);
extern void callback_ZN5QDial15mousePressEventEP11QMouseEvent(void* fnptr , void* me);
extern void callback_ZN5QDial17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* me);
extern void callback_ZN5QDial14mouseMoveEventEP11QMouseEvent(void* fnptr , void* me);
extern void callback_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE(void* fnptr , int change);
extern void callback_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider(void* fnptr , void* option);
extern void callback_ZN16QDialogButtonBox11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN16QDialogButtonBox5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN11QDockWidget11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN11QDockWidget10closeEventEP11QCloseEvent(void* fnptr , void* event);
extern void callback_ZN11QDockWidget10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZN11QDockWidget5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget(void* fnptr , void* option);
extern void callback_ZN13QErrorMessage4doneEi(void* fnptr , int arg0);
extern void callback_ZN13QErrorMessage11changeEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QFileDialog4doneEi(void* fnptr , int result);
extern void callback_ZN11QFileDialog6acceptEv(void* fnptr );
extern void callback_ZN11QFileDialog11changeEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN16QFileSystemModel10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZN16QFileSystemModel5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN11QFocusFrame5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent(void* fnptr , void* arg0, void* arg1);
extern void callback_ZN11QFocusFrame10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption(void* fnptr , void* option);
extern void callback_ZN13QFontComboBox5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QFontDialog11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN11QFontDialog4doneEi(void* fnptr , int result);
extern void callback_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent(void* fnptr , void* object, void* event);
extern void callback_ZN13QGraphicsItem16updateMicroFocusEv(void* fnptr );
extern void callback_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent(void* fnptr , void* watched, void* event);
extern void callback_ZN13QGraphicsItem10sceneEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* event);
extern void callback_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE(void* fnptr , int query);
extern void callback_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant(void* fnptr , int change, void* value);
extern void callback_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE(void* fnptr , int extension);
extern void callback_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK13QGraphicsItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZN13QGraphicsItem10addToIndexEv(void* fnptr );
extern void callback_ZN13QGraphicsItem15removeFromIndexEv(void* fnptr );
extern void callback_ZN13QGraphicsItem21prepareGeometryChangeEv(void* fnptr );
extern void callback_ZNK17QGraphicsPathItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN17QGraphicsPathItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK17QGraphicsPathItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK17QGraphicsRectItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZNK20QGraphicsEllipseItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN20QGraphicsEllipseItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK20QGraphicsEllipseItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZNK20QGraphicsPolygonItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN20QGraphicsPolygonItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK20QGraphicsPolygonItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZNK17QGraphicsLineItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN17QGraphicsLineItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK17QGraphicsLineItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZNK19QGraphicsPixmapItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN19QGraphicsPixmapItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK19QGraphicsPixmapItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZN17QGraphicsTextItem10sceneEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE(void* fnptr , int query);
extern void callback_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK17QGraphicsTextItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE(void* fnptr , int extension);
extern void callback_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(void* fnptr , int extension, void* variant);
extern void callback_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant(void* fnptr , void* variant);
extern void callback_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem(void* fnptr , void* item);
extern void callback_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb(void* fnptr , bool ownedByLayout);
extern void callback_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF(void* fnptr , int which, void* constraint);
extern void callback_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem(void* fnptr , void* layoutItem);
extern void callback_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF(void* fnptr , int which, void* constraint);
extern void callback_ZN15QGraphicsEffect4drawEP8QPainter(void* fnptr , void* painter);
extern void callback_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE(void* fnptr , int flags);
extern void callback_ZN15QGraphicsEffect18updateBoundingRectEv(void* fnptr );
extern void callback_ZNK15QGraphicsEffect14sourceIsPixmapEv(void* fnptr );
extern void callback_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE(void* fnptr , int system);
extern void callback_ZN15QGraphicsEffect10drawSourceEP8QPainter(void* fnptr , void* painter);
extern void callback_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE(void* fnptr , int system, void* offset, int mode);
extern void callback_ZN23QGraphicsColorizeEffect4drawEP8QPainter(void* fnptr , void* painter);
extern void callback_ZN19QGraphicsBlurEffect4drawEP8QPainter(void* fnptr , void* painter);
extern void callback_ZN25QGraphicsDropShadowEffect4drawEP8QPainter(void* fnptr , void* painter);
extern void callback_ZN22QGraphicsOpacityEffect4drawEP8QPainter(void* fnptr , void* painter);
extern void callback_ZN22QGraphicsItemAnimation19beforeAnimationStepEd(void* fnptr , double step);
extern void callback_ZN22QGraphicsItemAnimation18afterAnimationStepEd(void* fnptr , double step);
extern void callback_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption(void* fnptr , void* option);
extern void callback_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF(void* fnptr , int which, void* constraint);
extern void callback_ZN15QGraphicsWidget14updateGeometryEv(void* fnptr );
extern void callback_ZN15QGraphicsWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant(void* fnptr , int change, void* value);
extern void callback_ZN15QGraphicsWidget14propertyChangeERK7QStringRK8QVariant(void* fnptr , void* propertyName, void* value);
extern void callback_ZN15QGraphicsWidget10sceneEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget16windowFrameEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF(void* fnptr , void* pos);
extern void callback_ZN15QGraphicsWidget5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget10closeEventEP11QCloseEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget9hideEventEP10QHideEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget11polishEventEv(void* fnptr );
extern void callback_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget9showEventEP10QShowEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget14grabMouseEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant(void* fnptr , int change, void* value);
extern void callback_ZN20QGraphicsProxyWidget5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget11eventFilterEP7QObjectP6QEvent(void* fnptr , void* object, void* event);
extern void callback_ZN20QGraphicsProxyWidget9showEventEP10QShowEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget9hideEventEP10QHideEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget16contextMenuEventEP30QGraphicsSceneContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget15hoverEnterEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget14grabMouseEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget16ungrabMouseEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget14mouseMoveEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget15mousePressEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget17mouseReleaseEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget15keyReleaseEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZNK20QGraphicsProxyWidget16inputMethodQueryEN2Qt16InputMethodQueryE(void* fnptr , int query);
extern void callback_ZN20QGraphicsProxyWidget16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* event);
extern void callback_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF(void* fnptr , int which, void* constraint);
extern void callback_ZN20QGraphicsProxyWidget11resizeEventEP25QGraphicsSceneResizeEvent(void* fnptr , void* event);
extern void callback_ZN20QGraphicsProxyWidget14newProxyWidgetEPK7QWidget(void* fnptr , void* arg0);
extern void callback_ZN14QGraphicsScene5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent(void* fnptr , void* watched, void* event);
extern void callback_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* event);
extern void callback_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF(void* fnptr , void* painter, void* rect);
extern void callback_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF(void* fnptr , void* painter, void* rect);
extern void callback_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget(void* fnptr , void* painter, int numItems, void* items, void* options, void* widget);
extern void callback_ZN14QGraphicsScene18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN18QGraphicsTransform6updateEv(void* fnptr );
extern void callback_ZN11QScrollArea5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent(void* fnptr , void* arg0, void* arg1);
extern void callback_ZN11QScrollArea11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN11QScrollArea16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZNK11QScrollArea16viewportSizeHintEv(void* fnptr );
extern void callback_ZN13QGraphicsView13setupViewportEP7QWidget(void* fnptr , void* widget);
extern void callback_ZN13QGraphicsView5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView13viewportEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView9dropEventEP10QDropEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN13QGraphicsView13focusOutEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView10wheelEventEP11QWheelEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN13QGraphicsView9showEventEP10QShowEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* event);
extern void callback_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF(void* fnptr , void* painter, void* rect);
extern void callback_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF(void* fnptr , void* painter, void* rect);
extern void callback_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem(void* fnptr , void* painter, int numItems, void* items, void* options);
extern void callback_ZN9QGroupBox5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox10childEventEP11QChildEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox12focusInEventEP11QFocusEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox(void* fnptr , void* option);
extern void callback_ZN11QHeaderView13updateSectionEi(void* fnptr , int logicalIndex);
extern void callback_ZN11QHeaderView14resizeSectionsEv(void* fnptr );
extern void callback_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii(void* fnptr , void* parent, int logicalFirst, int logicalLast);
extern void callback_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii(void* fnptr , void* parent, int logicalFirst, int logicalLast);
extern void callback_ZN11QHeaderView10initializeEv(void* fnptr );
extern void callback_ZN11QHeaderView18initializeSectionsEv(void* fnptr );
extern void callback_ZN11QHeaderView18initializeSectionsEii(void* fnptr , int start, int end);
extern void callback_ZN11QHeaderView14currentChangedERK11QModelIndexS2_(void* fnptr , void* current, void* old);
extern void callback_ZN11QHeaderView5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QHeaderView10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZN11QHeaderView15mousePressEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN11QHeaderView13viewportEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti(void* fnptr , void* painter, void* rect, int logicalIndex);
extern void callback_ZNK11QHeaderView23sectionSizeFromContentsEi(void* fnptr , int logicalIndex);
extern void callback_ZNK11QHeaderView16horizontalOffsetEv(void* fnptr );
extern void callback_ZNK11QHeaderView14verticalOffsetEv(void* fnptr );
extern void callback_ZN11QHeaderView16updateGeometriesEv(void* fnptr );
extern void callback_ZN11QHeaderView16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN11QHeaderView12rowsInsertedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZNK11QHeaderView10visualRectERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE(void* fnptr , void* index, int hint);
extern void callback_ZNK11QHeaderView7indexAtERK6QPoint(void* fnptr , void* p);
extern void callback_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(void* fnptr , void* rect, int flags);
extern void callback_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection(void* fnptr , void* selection);
extern void callback_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader(void* fnptr , void* option);
extern void callback_ZN9QLineEdit15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit12focusInEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit13focusOutEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* e);
extern void callback_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* e);
extern void callback_ZN9QLineEdit9dropEventEP10QDropEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* arg0);
extern void callback_ZN9QLineEdit16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* arg0);
extern void callback_ZNK9QLineEdit15initStyleOptionEP17QStyleOptionFrame(void* fnptr , void* option);
extern void callback_ZNK9QLineEdit10cursorRectEv(void* fnptr );
extern void callback_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString(void* fnptr , void* painter, void* option, void* rect, void* text);
extern void callback_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap(void* fnptr , void* painter, void* option, void* rect, void* pixmap);
extern void callback_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect(void* fnptr , void* painter, void* option, void* rect);
extern void callback_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE(void* fnptr , void* painter, void* option, void* rect, int state);
extern void callback_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(void* fnptr , void* painter, void* option, void* index);
extern void callback_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b(void* fnptr , void* option, void* checkRect, void* iconRect, void* textRect, bool hint);
extern void callback_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi(void* fnptr , void* option, void* index, int role);
extern void callback_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent(void* fnptr , void* object, void* event);
extern void callback_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex(void* fnptr , void* event, void* model, void* option, void* index);
extern void callback_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem(void* fnptr , void* index, void* option);
extern void callback_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant(void* fnptr , void* option, void* variant);
extern void callback_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb(void* fnptr , void* pixmap, void* palette, bool enabled);
extern void callback_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant(void* fnptr , void* option, void* bounding, void* variant);
extern void callback_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString(void* fnptr , void* painter, void* rect, void* font, void* text);
extern void callback_ZN19QKeyEventTransition12onTransitionEP6QEvent(void* fnptr , void* event);
extern void callback_ZN19QKeyEventTransition9eventTestEP6QEvent(void* fnptr , void* event);
extern void callback_ZN16QKeySequenceEdit5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent(void* fnptr , void* arg0);
extern void callback_ZN6QLabel5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN6QLabel13keyPressEventEP9QKeyEvent(void* fnptr , void* ev);
extern void callback_ZN6QLabel10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN6QLabel11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN6QLabel15mousePressEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN6QLabel14mouseMoveEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN6QLabel17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN6QLabel16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* ev);
extern void callback_ZN6QLabel12focusInEventEP11QFocusEvent(void* fnptr , void* ev);
extern void callback_ZN6QLabel13focusOutEventEP11QFocusEvent(void* fnptr , void* ev);
extern void callback_ZN6QLabel18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN10QLCDNumber5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN10QLCDNumber10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN9QListView5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN9QListView16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN9QListView14resizeContentsEii(void* fnptr , int width, int height);
extern void callback_ZNK9QListView12contentsSizeEv(void* fnptr );
extern void callback_ZN9QListView12rowsInsertedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZN9QListView14mouseMoveEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QListView17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QListView10wheelEventEP11QWheelEvent(void* fnptr , void* e);
extern void callback_ZN9QListView10timerEventEP11QTimerEvent(void* fnptr , void* e);
extern void callback_ZN9QListView11resizeEventEP12QResizeEvent(void* fnptr , void* e);
extern void callback_ZN9QListView13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* e);
extern void callback_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* e);
extern void callback_ZN9QListView9dropEventEP10QDropEvent(void* fnptr , void* e);
extern void callback_ZNK9QListView11viewOptionsEv(void* fnptr );
extern void callback_ZN9QListView10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZNK9QListView16horizontalOffsetEv(void* fnptr );
extern void callback_ZNK9QListView14verticalOffsetEv(void* fnptr );
extern void callback_ZNK9QListView12rectForIndexERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex(void* fnptr , void* position, void* index);
extern void callback_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(void* fnptr , void* rect, int command);
extern void callback_ZNK9QListView24visualRegionForSelectionERK14QItemSelection(void* fnptr , void* selection);
extern void callback_ZN9QListView16updateGeometriesEv(void* fnptr );
extern void callback_ZNK9QListView13isIndexHiddenERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN9QListView16selectionChangedERK14QItemSelectionS2_(void* fnptr , void* selected, void* deselected);
extern void callback_ZN9QListView14currentChangedERK11QModelIndexS2_(void* fnptr , void* current, void* previous);
extern void callback_ZNK9QListView16viewportSizeHintEv(void* fnptr );
extern void callback_ZN11QListWidget5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE(void* fnptr , int index, void* data, int action);
extern void callback_ZNK11QListWidget20supportedDropActionsEv(void* fnptr );
extern void callback_ZNK11QListWidget13indexFromItemEP15QListWidgetItem(void* fnptr , void* item);
extern void callback_ZNK11QListWidget13itemFromIndexERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* event);
extern void callback_ZN11QMainWindow5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN8QMdiArea13setupViewportEP7QWidget(void* fnptr , void* viewport);
extern void callback_ZN8QMdiArea5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent(void* fnptr , void* object, void* event);
extern void callback_ZN8QMdiArea10paintEventEP11QPaintEvent(void* fnptr , void* paintEvent);
extern void callback_ZN8QMdiArea10childEventEP11QChildEvent(void* fnptr , void* childEvent);
extern void callback_ZN8QMdiArea11resizeEventEP12QResizeEvent(void* fnptr , void* resizeEvent);
extern void callback_ZN8QMdiArea10timerEventEP11QTimerEvent(void* fnptr , void* timerEvent);
extern void callback_ZN8QMdiArea9showEventEP10QShowEvent(void* fnptr , void* showEvent);
extern void callback_ZN8QMdiArea13viewportEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN8QMdiArea16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent(void* fnptr , void* object, void* event);
extern void callback_ZN13QMdiSubWindow5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QMdiSubWindow9showEventEP10QShowEvent(void* fnptr , void* showEvent);
extern void callback_ZN13QMdiSubWindow9hideEventEP10QHideEvent(void* fnptr , void* hideEvent);
extern void callback_ZN13QMdiSubWindow11changeEventEP6QEvent(void* fnptr , void* changeEvent);
extern void callback_ZN13QMdiSubWindow10closeEventEP11QCloseEvent(void* fnptr , void* closeEvent);
extern void callback_ZN13QMdiSubWindow10leaveEventEP6QEvent(void* fnptr , void* leaveEvent);
extern void callback_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent(void* fnptr , void* resizeEvent);
extern void callback_ZN13QMdiSubWindow10timerEventEP11QTimerEvent(void* fnptr , void* timerEvent);
extern void callback_ZN13QMdiSubWindow9moveEventEP10QMoveEvent(void* fnptr , void* moveEvent);
extern void callback_ZN13QMdiSubWindow10paintEventEP11QPaintEvent(void* fnptr , void* paintEvent);
extern void callback_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent(void* fnptr , void* mouseEvent);
extern void callback_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* mouseEvent);
extern void callback_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* mouseEvent);
extern void callback_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent(void* fnptr , void* mouseEvent);
extern void callback_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent(void* fnptr , void* keyEvent);
extern void callback_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* contextMenuEvent);
extern void callback_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent(void* fnptr , void* focusInEvent);
extern void callback_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent(void* fnptr , void* focusOutEvent);
extern void callback_ZN13QMdiSubWindow10childEventEP11QChildEvent(void* fnptr , void* childEvent);
extern void callback_ZNK5QMenu11columnCountEv(void* fnptr );
extern void callback_ZN5QMenu11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu10wheelEventEP11QWheelEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu10enterEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu10leaveEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu9hideEventEP10QHideEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu11actionEventEP12QActionEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu10timerEventEP11QTimerEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN5QMenu18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction(void* fnptr , void* option, void* action);
extern void callback_ZN8QMenuBar11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar13keyPressEventEP9QKeyEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar10leaveEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar11actionEventEP12QActionEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar13focusOutEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar12focusInEventEP11QFocusEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar10timerEventEP11QTimerEvent(void* fnptr , void* arg0);
extern void callback_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent(void* fnptr , void* arg0, void* arg1);
extern void callback_ZN8QMenuBar5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction(void* fnptr , void* option, void* action);
extern void callback_ZN11QMessageBox5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QMessageBox11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN11QMessageBox9showEventEP10QShowEvent(void* fnptr , void* event);
extern void callback_ZN11QMessageBox10closeEventEP11QCloseEvent(void* fnptr , void* event);
extern void callback_ZN11QMessageBox13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN11QMessageBox11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN21QMouseEventTransition12onTransitionEP6QEvent(void* fnptr , void* event);
extern void callback_ZN21QMouseEventTransition9eventTestEP6QEvent(void* fnptr , void* event);
extern void callback_ZN9QTextEdit5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit10timerEventEP11QTimerEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit13keyPressEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit11resizeEventEP12QResizeEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit15mousePressEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit9dropEventEP10QDropEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit12focusInEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit13focusOutEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN9QTextEdit11changeEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN9QTextEdit10wheelEventEP11QWheelEvent(void* fnptr , void* e);
extern void callback_ZNK9QTextEdit27createMimeDataFromSelectionEv(void* fnptr );
extern void callback_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData(void* fnptr , void* source);
extern void callback_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData(void* fnptr , void* source);
extern void callback_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* arg0);
extern void callback_ZN9QTextEdit16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN9QTextEdit15doSetTextCursorERK11QTextCursor(void* fnptr , void* cursor);
extern void callback_ZN9QTextEdit7zoomInFEf(void* fnptr , float range_);
extern void callback_ZN14QPlainTextEdit5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit10timerEventEP11QTimerEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit9dropEventEP10QDropEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN14QPlainTextEdit11changeEventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent(void* fnptr , void* e);
extern void callback_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv(void* fnptr );
extern void callback_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData(void* fnptr , void* source);
extern void callback_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData(void* fnptr , void* source);
extern void callback_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent(void* fnptr , void* arg0);
extern void callback_ZN14QPlainTextEdit16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor(void* fnptr , void* cursor);
extern void callback_ZNK14QPlainTextEdit17firstVisibleBlockEv(void* fnptr );
extern void callback_ZNK14QPlainTextEdit13contentOffsetEv(void* fnptr );
extern void callback_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock(void* fnptr , void* block);
extern void callback_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock(void* fnptr , void* block);
extern void callback_ZNK14QPlainTextEdit15getPaintContextEv(void* fnptr );
extern void callback_ZN14QPlainTextEdit7zoomInFEf(void* fnptr , float range_);
extern void callback_ZN24QPlainTextDocumentLayout15documentChangedEiii(void* fnptr , int from, int arg1, int charsAdded);
extern void callback_ZN12QProgressBar5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN12QProgressBar10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZNK12QProgressBar15initStyleOptionEP23QStyleOptionProgressBar(void* fnptr , void* option);
extern void callback_ZN15QProgressDialog11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN15QProgressDialog10closeEventEP11QCloseEvent(void* fnptr , void* event);
extern void callback_ZN15QProgressDialog11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN15QProgressDialog9showEventEP10QShowEvent(void* fnptr , void* event);
extern void callback_ZN15QProgressDialog9forceShowEv(void* fnptr );
extern void callback_ZN11QProxyStyle5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN12QRadioButton5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZNK12QRadioButton9hitButtonERK6QPoint(void* fnptr , void* arg0);
extern void callback_ZN12QRadioButton10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton(void* fnptr , void* button);
extern void callback_ZN10QScrollBar10wheelEventEP11QWheelEvent(void* fnptr , void* arg0);
extern void callback_ZN10QScrollBar10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN10QScrollBar15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN10QScrollBar9hideEventEP10QHideEvent(void* fnptr , void* arg0);
extern void callback_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE(void* fnptr , int change);
extern void callback_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent(void* fnptr , void* arg0);
extern void callback_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider(void* fnptr , void* option);
extern void callback_ZN9QShortcut5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN9QSizeGrip10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSizeGrip15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSizeGrip14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSizeGrip17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* mouseEvent);
extern void callback_ZN9QSizeGrip9moveEventEP10QMoveEvent(void* fnptr , void* moveEvent);
extern void callback_ZN9QSizeGrip9showEventEP10QShowEvent(void* fnptr , void* showEvent);
extern void callback_ZN9QSizeGrip9hideEventEP10QHideEvent(void* fnptr , void* hideEvent);
extern void callback_ZN9QSizeGrip11eventFilterEP7QObjectP6QEvent(void* fnptr , void* arg0, void* arg1);
extern void callback_ZN9QSizeGrip5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN8QSpinBox5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZNK8QSpinBox8validateER7QStringRi(void* fnptr , void* input, void* pos);
extern void callback_ZNK8QSpinBox13valueFromTextERK7QString(void* fnptr , void* text);
extern void callback_ZNK8QSpinBox13textFromValueEi(void* fnptr , int val);
extern void callback_ZNK8QSpinBox5fixupER7QString(void* fnptr , void* str);
extern void callback_ZN13QSplashScreen5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN13QSplashScreen12drawContentsEP8QPainter(void* fnptr , void* painter);
extern void callback_ZN13QSplashScreen15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSplitter12createHandleEv(void* fnptr );
extern void callback_ZN9QSplitter10childEventEP11QChildEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSplitter5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSplitter11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSplitter11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN9QSplitter12moveSplitterEii(void* fnptr , int pos, int index);
extern void callback_ZN9QSplitter13setRubberBandEi(void* fnptr , int position);
extern void callback_ZN9QSplitter20closestLegalPositionEii(void* fnptr , int arg0, int arg1);
extern void callback_ZN15QSplitterHandle10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN15QSplitterHandle11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN15QSplitterHandle5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN15QSplitterHandle12moveSplitterEi(void* fnptr , int p);
extern void callback_ZN15QSplitterHandle20closestLegalPositionEi(void* fnptr , int p);
extern void callback_ZN14QStackedWidget5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN10QStatusBar9showEventEP10QShowEvent(void* fnptr , void* arg0);
extern void callback_ZN10QStatusBar10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN10QStatusBar11resizeEventEP12QResizeEvent(void* fnptr , void* arg0);
extern void callback_ZN10QStatusBar8reformatEv(void* fnptr );
extern void callback_ZN10QStatusBar10hideOrShowEv(void* fnptr );
extern void callback_ZN10QStatusBar5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex(void* fnptr , void* option, void* index);
extern void callback_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent(void* fnptr , void* object, void* event);
extern void callback_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex(void* fnptr , void* event, void* model, void* option, void* index);
extern void callback_ZN15QSystemTrayIcon5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN10QTableView8rowMovedEiii(void* fnptr , int row, int oldIndex, int newIndex);
extern void callback_ZN10QTableView11columnMovedEiii(void* fnptr , int column, int oldIndex, int newIndex);
extern void callback_ZN10QTableView10rowResizedEiii(void* fnptr , int row, int oldHeight, int newHeight);
extern void callback_ZN10QTableView13columnResizedEiii(void* fnptr , int column, int oldWidth, int newWidth);
extern void callback_ZN10QTableView15rowCountChangedEii(void* fnptr , int oldCount, int newCount);
extern void callback_ZN10QTableView18columnCountChangedEii(void* fnptr , int oldCount, int newCount);
extern void callback_ZN10QTableView16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZNK10QTableView11viewOptionsEv(void* fnptr );
extern void callback_ZN10QTableView10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZN10QTableView10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZNK10QTableView16horizontalOffsetEv(void* fnptr );
extern void callback_ZNK10QTableView14verticalOffsetEv(void* fnptr );
extern void callback_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(void* fnptr , void* rect, int command);
extern void callback_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection(void* fnptr , void* selection);
extern void callback_ZN10QTableView16updateGeometriesEv(void* fnptr );
extern void callback_ZNK10QTableView16viewportSizeHintEv(void* fnptr );
extern void callback_ZNK10QTableView14sizeHintForRowEi(void* fnptr , int row);
extern void callback_ZNK10QTableView17sizeHintForColumnEi(void* fnptr , int column);
extern void callback_ZN10QTableView23verticalScrollbarActionEi(void* fnptr , int action);
extern void callback_ZN10QTableView25horizontalScrollbarActionEi(void* fnptr , int action);
extern void callback_ZNK10QTableView13isIndexHiddenERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN10QTableView16selectionChangedERK14QItemSelectionS2_(void* fnptr , void* selected, void* deselected);
extern void callback_ZN10QTableView14currentChangedERK11QModelIndexS2_(void* fnptr , void* current, void* previous);
extern void callback_ZN12QTableWidget5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE(void* fnptr , int row, int column, void* data, int action);
extern void callback_ZNK12QTableWidget20supportedDropActionsEv(void* fnptr );
extern void callback_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem(void* fnptr , void* item);
extern void callback_ZNK12QTableWidget13itemFromIndexERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN12QTableWidget9dropEventEP10QDropEvent(void* fnptr , void* event);
extern void callback_ZN12QTextBrowser5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN12QTextBrowser13keyPressEventEP9QKeyEvent(void* fnptr , void* ev);
extern void callback_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN12QTextBrowser15mousePressEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* ev);
extern void callback_ZN12QTextBrowser13focusOutEventEP11QFocusEvent(void* fnptr , void* ev);
extern void callback_ZN12QTextBrowser18focusNextPrevChildEb(void* fnptr , bool next);
extern void callback_ZN12QTextBrowser10paintEventEP11QPaintEvent(void* fnptr , void* e);
extern void callback_ZN8QToolBar11actionEventEP12QActionEvent(void* fnptr , void* event);
extern void callback_ZN8QToolBar11changeEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN8QToolBar10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZN8QToolBar5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar(void* fnptr , void* option);
extern void callback_ZN8QToolBox5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN8QToolBox12itemInsertedEi(void* fnptr , int index);
extern void callback_ZN8QToolBox11itemRemovedEi(void* fnptr , int index);
extern void callback_ZN8QToolBox9showEventEP10QShowEvent(void* fnptr , void* e);
extern void callback_ZN8QToolBox11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QToolButton15mousePressEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton10paintEventEP11QPaintEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton11actionEventEP12QActionEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton10enterEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton10leaveEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton10timerEventEP11QTimerEvent(void* fnptr , void* arg0);
extern void callback_ZN11QToolButton11changeEventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZNK11QToolButton9hitButtonERK6QPoint(void* fnptr , void* pos);
extern void callback_ZN11QToolButton14nextCheckStateEv(void* fnptr );
extern void callback_ZNK11QToolButton15initStyleOptionEP22QStyleOptionToolButton(void* fnptr , void* option);
extern void callback_ZN9QTreeView13columnResizedEiii(void* fnptr , int column, int oldSize, int newSize);
extern void callback_ZN9QTreeView18columnCountChangedEii(void* fnptr , int oldCount, int newCount);
extern void callback_ZN9QTreeView11columnMovedEv(void* fnptr );
extern void callback_ZN9QTreeView8reexpandEv(void* fnptr );
extern void callback_ZN9QTreeView11rowsRemovedERK11QModelIndexii(void* fnptr , void* parent, int first, int last);
extern void callback_ZN9QTreeView16scrollContentsByEii(void* fnptr , int dx, int dy);
extern void callback_ZN9QTreeView12rowsInsertedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii(void* fnptr , void* parent, int start, int end);
extern void callback_ZNK9QTreeView16horizontalOffsetEv(void* fnptr );
extern void callback_ZNK9QTreeView14verticalOffsetEv(void* fnptr );
extern void callback_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(void* fnptr , void* rect, int command);
extern void callback_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection(void* fnptr , void* selection);
extern void callback_ZN9QTreeView10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion(void* fnptr , void* painter, void* region);
extern void callback_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(void* fnptr , void* painter, void* options, void* index);
extern void callback_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex(void* fnptr , void* painter, void* rect, void* index);
extern void callback_ZN9QTreeView15mousePressEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView14mouseMoveEventEP11QMouseEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView13keyPressEventEP9QKeyEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView13viewportEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN9QTreeView16updateGeometriesEv(void* fnptr );
extern void callback_ZNK9QTreeView16viewportSizeHintEv(void* fnptr );
extern void callback_ZNK9QTreeView17sizeHintForColumnEi(void* fnptr , int column);
extern void callback_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZNK9QTreeView9rowHeightERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN9QTreeView25horizontalScrollbarActionEi(void* fnptr , int action);
extern void callback_ZNK9QTreeView13isIndexHiddenERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_(void* fnptr , void* selected, void* deselected);
extern void callback_ZN9QTreeView14currentChangedERK11QModelIndexS2_(void* fnptr , void* current, void* previous);
extern void callback_ZN15QTreeWidgetItem15emitDataChangedEv(void* fnptr );
extern void callback_ZN11QTreeWidget5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE(void* fnptr , void* parent, int index, void* data, int action);
extern void callback_ZNK11QTreeWidget20supportedDropActionsEv(void* fnptr );
extern void callback_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi(void* fnptr , void* item, int column);
extern void callback_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi(void* fnptr , void* item, int column);
extern void callback_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex(void* fnptr , void* index);
extern void callback_ZN11QTreeWidget9dropEventEP10QDropEvent(void* fnptr , void* event);
extern void callback_ZN13QWidgetAction5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent(void* fnptr , void* arg0, void* arg1);
extern void callback_ZN13QWidgetAction12createWidgetEP7QWidget(void* fnptr , void* parent);
extern void callback_ZN13QWidgetAction12deleteWidgetEP7QWidget(void* fnptr , void* widget);
extern void callback_ZN7QWizard5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN7QWizard11resizeEventEP12QResizeEvent(void* fnptr , void* event);
extern void callback_ZN7QWizard10paintEventEP11QPaintEvent(void* fnptr , void* event);
extern void callback_ZN7QWizard4doneEi(void* fnptr , int result);
extern void callback_ZN7QWizard14initializePageEi(void* fnptr , int id);
extern void callback_ZN7QWizard11cleanupPageEi(void* fnptr , int id);
extern void callback_ZN11QWizardPage8setFieldERK7QStringRK8QVariant(void* fnptr , void* name, void* value);
extern void callback_ZNK11QWizardPage5fieldERK7QString(void* fnptr , void* name);
extern void callback_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_(void* fnptr , void* name, void* widget, void* property, void* changedSignal);
extern void callback_ZNK11QWizardPage6wizardEv(void* fnptr );
//  extern block end

//  header block begin
*/
import "C"
import "unsafe"
import "qt.go/cffiqt"
import "gopp"

// import "log"
//  header block end

//  body block begin
// bool event(class QEvent *)
//export callback_ZN7QWidget5eventEP6QEvent
func callback_ZN7QWidget5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget5eventEP6QEvent", C.callback_ZN7QWidget5eventEP6QEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN7QWidget15mousePressEventEP11QMouseEvent
func callback_ZN7QWidget15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget15mousePressEventEP11QMouseEvent", C.callback_ZN7QWidget15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN7QWidget17mouseReleaseEventEP11QMouseEvent
func callback_ZN7QWidget17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget17mouseReleaseEventEP11QMouseEvent", C.callback_ZN7QWidget17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN7QWidget14mouseMoveEventEP11QMouseEvent
func callback_ZN7QWidget14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget14mouseMoveEventEP11QMouseEvent", C.callback_ZN7QWidget14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN7QWidget10wheelEventEP11QWheelEvent
func callback_ZN7QWidget10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget10wheelEventEP11QWheelEvent", C.callback_ZN7QWidget10wheelEventEP11QWheelEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN7QWidget13keyPressEventEP9QKeyEvent
func callback_ZN7QWidget13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget13keyPressEventEP9QKeyEvent", C.callback_ZN7QWidget13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN7QWidget15keyReleaseEventEP9QKeyEvent
func callback_ZN7QWidget15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget15keyReleaseEventEP9QKeyEvent", C.callback_ZN7QWidget15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN7QWidget12focusInEventEP11QFocusEvent
func callback_ZN7QWidget12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget12focusInEventEP11QFocusEvent", C.callback_ZN7QWidget12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN7QWidget13focusOutEventEP11QFocusEvent
func callback_ZN7QWidget13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget13focusOutEventEP11QFocusEvent", C.callback_ZN7QWidget13focusOutEventEP11QFocusEvent /*nil*/)
}

// void enterEvent(class QEvent *)
//export callback_ZN7QWidget10enterEventEP6QEvent
func callback_ZN7QWidget10enterEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.enterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "enterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget10enterEventEP6QEvent", C.callback_ZN7QWidget10enterEventEP6QEvent /*nil*/)
}

// void leaveEvent(class QEvent *)
//export callback_ZN7QWidget10leaveEventEP6QEvent
func callback_ZN7QWidget10leaveEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.leaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "leaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget10leaveEventEP6QEvent", C.callback_ZN7QWidget10leaveEventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN7QWidget10paintEventEP11QPaintEvent
func callback_ZN7QWidget10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget10paintEventEP11QPaintEvent", C.callback_ZN7QWidget10paintEventEP11QPaintEvent /*nil*/)
}

// void moveEvent(class QMoveEvent *)
//export callback_ZN7QWidget9moveEventEP10QMoveEvent
func callback_ZN7QWidget9moveEventEP10QMoveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.moveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget9moveEventEP10QMoveEvent", C.callback_ZN7QWidget9moveEventEP10QMoveEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN7QWidget11resizeEventEP12QResizeEvent
func callback_ZN7QWidget11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget11resizeEventEP12QResizeEvent", C.callback_ZN7QWidget11resizeEventEP12QResizeEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN7QWidget10closeEventEP11QCloseEvent
func callback_ZN7QWidget10closeEventEP11QCloseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget10closeEventEP11QCloseEvent", C.callback_ZN7QWidget10closeEventEP11QCloseEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN7QWidget16contextMenuEventEP17QContextMenuEvent
func callback_ZN7QWidget16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget16contextMenuEventEP17QContextMenuEvent", C.callback_ZN7QWidget16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void tabletEvent(class QTabletEvent *)
//export callback_ZN7QWidget11tabletEventEP12QTabletEvent
func callback_ZN7QWidget11tabletEventEP12QTabletEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.tabletEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabletEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget11tabletEventEP12QTabletEvent", C.callback_ZN7QWidget11tabletEventEP12QTabletEvent /*nil*/)
}

// void actionEvent(class QActionEvent *)
//export callback_ZN7QWidget11actionEventEP12QActionEvent
func callback_ZN7QWidget11actionEventEP12QActionEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.actionEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "actionEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget11actionEventEP12QActionEvent", C.callback_ZN7QWidget11actionEventEP12QActionEvent /*nil*/)
}

// void dragEnterEvent(class QDragEnterEvent *)
//export callback_ZN7QWidget14dragEnterEventEP15QDragEnterEvent
func callback_ZN7QWidget14dragEnterEventEP15QDragEnterEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget14dragEnterEventEP15QDragEnterEvent", C.callback_ZN7QWidget14dragEnterEventEP15QDragEnterEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN7QWidget13dragMoveEventEP14QDragMoveEvent
func callback_ZN7QWidget13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget13dragMoveEventEP14QDragMoveEvent", C.callback_ZN7QWidget13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN7QWidget9dropEventEP10QDropEvent
func callback_ZN7QWidget9dropEventEP10QDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget9dropEventEP10QDropEvent", C.callback_ZN7QWidget9dropEventEP10QDropEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN7QWidget9showEventEP10QShowEvent
func callback_ZN7QWidget9showEventEP10QShowEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget9showEventEP10QShowEvent", C.callback_ZN7QWidget9showEventEP10QShowEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN7QWidget9hideEventEP10QHideEvent
func callback_ZN7QWidget9hideEventEP10QHideEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget9hideEventEP10QHideEvent", C.callback_ZN7QWidget9hideEventEP10QHideEvent /*nil*/)
}

// bool nativeEvent(const class QByteArray &, void *, long *)
//export callback_ZN7QWidget11nativeEventERK10QByteArrayPvPl
func callback_ZN7QWidget11nativeEventERK10QByteArrayPvPl(cthis unsafe.Pointer, eventType unsafe.Pointer /*555*/, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.nativeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "nativeEvent", eventType, message, result)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget11nativeEventERK10QByteArrayPvPl", C.callback_ZN7QWidget11nativeEventERK10QByteArrayPvPl /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN7QWidget11changeEventEP6QEvent
func callback_ZN7QWidget11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget11changeEventEP6QEvent", C.callback_ZN7QWidget11changeEventEP6QEvent /*nil*/)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
//export callback_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE
func callback_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QWidget.metric")
	rvx := ffiqt.CallbackAllInherits(cthis, "metric", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE", C.callback_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE /*nil*/)
}

// void initPainter(class QPainter *)
//export callback_ZNK7QWidget11initPainterEP8QPainter
func callback_ZNK7QWidget11initPainterEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.initPainter")
	rvx := ffiqt.CallbackAllInherits(cthis, "initPainter", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QWidget11initPainterEP8QPainter", C.callback_ZNK7QWidget11initPainterEP8QPainter /*nil*/)
}

// QPaintDevice * redirected(class QPoint *)
//export callback_ZNK7QWidget10redirectedEP6QPoint
func callback_ZNK7QWidget10redirectedEP6QPoint(cthis unsafe.Pointer, offset unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.redirected")
	rvx := ffiqt.CallbackAllInherits(cthis, "redirected", offset)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QWidget10redirectedEP6QPoint", C.callback_ZNK7QWidget10redirectedEP6QPoint /*nil*/)
}

// QPainter * sharedPainter()
//export callback_ZNK7QWidget13sharedPainterEv
func callback_ZNK7QWidget13sharedPainterEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QWidget.sharedPainter")
	rvx := ffiqt.CallbackAllInherits(cthis, "sharedPainter")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QWidget13sharedPainterEv", C.callback_ZNK7QWidget13sharedPainterEv /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN7QWidget16inputMethodEventEP17QInputMethodEvent
func callback_ZN7QWidget16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidget.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget16inputMethodEventEP17QInputMethodEvent", C.callback_ZN7QWidget16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void updateMicroFocus()
//export callback_ZN7QWidget16updateMicroFocusEv
func callback_ZN7QWidget16updateMicroFocusEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QWidget.updateMicroFocus")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateMicroFocus")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget16updateMicroFocusEv", C.callback_ZN7QWidget16updateMicroFocusEv /*nil*/)
}

// void create(WId, _Bool, _Bool)
//export callback_ZN7QWidget6createEybb
func callback_ZN7QWidget6createEybb(cthis unsafe.Pointer, arg0 C.uint64_t, initializeWindow C.bool, destroyOldWindow C.bool) {
	// log.Println(cthis, "QWidget.create")
	rvx := ffiqt.CallbackAllInherits(cthis, "create", arg0, initializeWindow, destroyOldWindow)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget6createEybb", C.callback_ZN7QWidget6createEybb /*nil*/)
}

// void destroy(_Bool, _Bool)
//export callback_ZN7QWidget7destroyEbb
func callback_ZN7QWidget7destroyEbb(cthis unsafe.Pointer, destroyWindow C.bool, destroySubWindows C.bool) {
	// log.Println(cthis, "QWidget.destroy")
	rvx := ffiqt.CallbackAllInherits(cthis, "destroy", destroyWindow, destroySubWindows)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget7destroyEbb", C.callback_ZN7QWidget7destroyEbb /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN7QWidget18focusNextPrevChildEb
func callback_ZN7QWidget18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QWidget.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget18focusNextPrevChildEb", C.callback_ZN7QWidget18focusNextPrevChildEb /*nil*/)
}

// bool focusNextChild()
//export callback_ZN7QWidget14focusNextChildEv
func callback_ZN7QWidget14focusNextChildEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QWidget.focusNextChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextChild")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget14focusNextChildEv", C.callback_ZN7QWidget14focusNextChildEv /*nil*/)
}

// bool focusPreviousChild()
//export callback_ZN7QWidget18focusPreviousChildEv
func callback_ZN7QWidget18focusPreviousChildEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QWidget.focusPreviousChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusPreviousChild")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWidget18focusPreviousChildEv", C.callback_ZN7QWidget18focusPreviousChildEv /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN15QAbstractButton10paintEventEP11QPaintEvent
func callback_ZN15QAbstractButton10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton10paintEventEP11QPaintEvent", C.callback_ZN15QAbstractButton10paintEventEP11QPaintEvent /*nil*/)
}

// bool hitButton(const class QPoint &)
//export callback_ZNK15QAbstractButton9hitButtonERK6QPoint
func callback_ZNK15QAbstractButton9hitButtonERK6QPoint(cthis unsafe.Pointer, pos unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractButton.hitButton")
	rvx := ffiqt.CallbackAllInherits(cthis, "hitButton", pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QAbstractButton9hitButtonERK6QPoint", C.callback_ZNK15QAbstractButton9hitButtonERK6QPoint /*nil*/)
}

// void checkStateSet()
//export callback_ZN15QAbstractButton13checkStateSetEv
func callback_ZN15QAbstractButton13checkStateSetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractButton.checkStateSet")
	rvx := ffiqt.CallbackAllInherits(cthis, "checkStateSet")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton13checkStateSetEv", C.callback_ZN15QAbstractButton13checkStateSetEv /*nil*/)
}

// void nextCheckState()
//export callback_ZN15QAbstractButton14nextCheckStateEv
func callback_ZN15QAbstractButton14nextCheckStateEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractButton.nextCheckState")
	rvx := ffiqt.CallbackAllInherits(cthis, "nextCheckState")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton14nextCheckStateEv", C.callback_ZN15QAbstractButton14nextCheckStateEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QAbstractButton5eventEP6QEvent
func callback_ZN15QAbstractButton5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton5eventEP6QEvent", C.callback_ZN15QAbstractButton5eventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN15QAbstractButton13keyPressEventEP9QKeyEvent
func callback_ZN15QAbstractButton13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton13keyPressEventEP9QKeyEvent", C.callback_ZN15QAbstractButton13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent
func callback_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent", C.callback_ZN15QAbstractButton15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN15QAbstractButton15mousePressEventEP11QMouseEvent
func callback_ZN15QAbstractButton15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton15mousePressEventEP11QMouseEvent", C.callback_ZN15QAbstractButton15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent
func callback_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent", C.callback_ZN15QAbstractButton17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent
func callback_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent", C.callback_ZN15QAbstractButton14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN15QAbstractButton12focusInEventEP11QFocusEvent
func callback_ZN15QAbstractButton12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton12focusInEventEP11QFocusEvent", C.callback_ZN15QAbstractButton12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN15QAbstractButton13focusOutEventEP11QFocusEvent
func callback_ZN15QAbstractButton13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton13focusOutEventEP11QFocusEvent", C.callback_ZN15QAbstractButton13focusOutEventEP11QFocusEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN15QAbstractButton11changeEventEP6QEvent
func callback_ZN15QAbstractButton11changeEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton11changeEventEP6QEvent", C.callback_ZN15QAbstractButton11changeEventEP6QEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN15QAbstractButton10timerEventEP11QTimerEvent
func callback_ZN15QAbstractButton10timerEventEP11QTimerEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractButton.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractButton10timerEventEP11QTimerEvent", C.callback_ZN15QAbstractButton10timerEventEP11QTimerEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent
func callback_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent", C.callback_ZN16QAbstractSpinBox11resizeEventEP12QResizeEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent
func callback_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent", C.callback_ZN16QAbstractSpinBox13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent
func callback_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent", C.callback_ZN16QAbstractSpinBox15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent
func callback_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent", C.callback_ZN16QAbstractSpinBox10wheelEventEP11QWheelEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent
func callback_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent", C.callback_ZN16QAbstractSpinBox12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent
func callback_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent", C.callback_ZN16QAbstractSpinBox13focusOutEventEP11QFocusEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent
func callback_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent", C.callback_ZN16QAbstractSpinBox16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN16QAbstractSpinBox11changeEventEP6QEvent
func callback_ZN16QAbstractSpinBox11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox11changeEventEP6QEvent", C.callback_ZN16QAbstractSpinBox11changeEventEP6QEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent
func callback_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent", C.callback_ZN16QAbstractSpinBox10closeEventEP11QCloseEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN16QAbstractSpinBox9hideEventEP10QHideEvent
func callback_ZN16QAbstractSpinBox9hideEventEP10QHideEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox9hideEventEP10QHideEvent", C.callback_ZN16QAbstractSpinBox9hideEventEP10QHideEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent
func callback_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent", C.callback_ZN16QAbstractSpinBox15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent
func callback_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent", C.callback_ZN16QAbstractSpinBox17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent
func callback_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent", C.callback_ZN16QAbstractSpinBox14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent
func callback_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent", C.callback_ZN16QAbstractSpinBox10timerEventEP11QTimerEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent
func callback_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent", C.callback_ZN16QAbstractSpinBox10paintEventEP11QPaintEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN16QAbstractSpinBox9showEventEP10QShowEvent
func callback_ZN16QAbstractSpinBox9showEventEP10QShowEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox9showEventEP10QShowEvent", C.callback_ZN16QAbstractSpinBox9showEventEP10QShowEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionSpinBox *)
//export callback_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox
func callback_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox", C.callback_ZNK16QAbstractSpinBox15initStyleOptionEP19QStyleOptionSpinBox /*nil*/)
}

// QLineEdit * lineEdit()
//export callback_ZNK16QAbstractSpinBox8lineEditEv
func callback_ZNK16QAbstractSpinBox8lineEditEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractSpinBox.lineEdit")
	rvx := ffiqt.CallbackAllInherits(cthis, "lineEdit")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK16QAbstractSpinBox8lineEditEv", C.callback_ZNK16QAbstractSpinBox8lineEditEv /*nil*/)
}

// void setLineEdit(class QLineEdit *)
//export callback_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit
func callback_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit(cthis unsafe.Pointer, edit unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSpinBox.setLineEdit")
	rvx := ffiqt.CallbackAllInherits(cthis, "setLineEdit", edit)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit", C.callback_ZN16QAbstractSpinBox11setLineEditEP9QLineEdit /*nil*/)
}

// QAbstractSpinBox::StepEnabled stepEnabled()
//export callback_ZNK16QAbstractSpinBox11stepEnabledEv
func callback_ZNK16QAbstractSpinBox11stepEnabledEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractSpinBox.stepEnabled")
	rvx := ffiqt.CallbackAllInherits(cthis, "stepEnabled")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK16QAbstractSpinBox11stepEnabledEv", C.callback_ZNK16QAbstractSpinBox11stepEnabledEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QAbstractSlider5eventEP6QEvent
func callback_ZN15QAbstractSlider5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSlider.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractSlider5eventEP6QEvent", C.callback_ZN15QAbstractSlider5eventEP6QEvent /*nil*/)
}

// void setRepeatAction(enum QAbstractSlider::SliderAction, int, int)
//export callback_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii
func callback_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii(cthis unsafe.Pointer, action C.int, thresholdTime C.int, repeatTime C.int) {
	// log.Println(cthis, "QAbstractSlider.setRepeatAction")
	rvx := ffiqt.CallbackAllInherits(cthis, "setRepeatAction", action, thresholdTime, repeatTime)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii", C.callback_ZN15QAbstractSlider15setRepeatActionENS_12SliderActionEii /*nil*/)
}

// QAbstractSlider::SliderAction repeatAction()
//export callback_ZNK15QAbstractSlider12repeatActionEv
func callback_ZNK15QAbstractSlider12repeatActionEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractSlider.repeatAction")
	rvx := ffiqt.CallbackAllInherits(cthis, "repeatAction")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QAbstractSlider12repeatActionEv", C.callback_ZNK15QAbstractSlider12repeatActionEv /*nil*/)
}

// void sliderChange(enum QAbstractSlider::SliderChange)
//export callback_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE
func callback_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE(cthis unsafe.Pointer, change C.int) {
	// log.Println(cthis, "QAbstractSlider.sliderChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "sliderChange", change)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE", C.callback_ZN15QAbstractSlider12sliderChangeENS_12SliderChangeE /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent
func callback_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSlider.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent", C.callback_ZN15QAbstractSlider13keyPressEventEP9QKeyEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN15QAbstractSlider10timerEventEP11QTimerEvent
func callback_ZN15QAbstractSlider10timerEventEP11QTimerEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSlider.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractSlider10timerEventEP11QTimerEvent", C.callback_ZN15QAbstractSlider10timerEventEP11QTimerEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN15QAbstractSlider10wheelEventEP11QWheelEvent
func callback_ZN15QAbstractSlider10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSlider.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractSlider10wheelEventEP11QWheelEvent", C.callback_ZN15QAbstractSlider10wheelEventEP11QWheelEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN15QAbstractSlider11changeEventEP6QEvent
func callback_ZN15QAbstractSlider11changeEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractSlider.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAbstractSlider11changeEventEP6QEvent", C.callback_ZN15QAbstractSlider11changeEventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN7QSlider10paintEventEP11QPaintEvent
func callback_ZN7QSlider10paintEventEP11QPaintEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSlider.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QSlider10paintEventEP11QPaintEvent", C.callback_ZN7QSlider10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN7QSlider15mousePressEventEP11QMouseEvent
func callback_ZN7QSlider15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSlider.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QSlider15mousePressEventEP11QMouseEvent", C.callback_ZN7QSlider15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN7QSlider17mouseReleaseEventEP11QMouseEvent
func callback_ZN7QSlider17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSlider.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QSlider17mouseReleaseEventEP11QMouseEvent", C.callback_ZN7QSlider17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN7QSlider14mouseMoveEventEP11QMouseEvent
func callback_ZN7QSlider14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSlider.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QSlider14mouseMoveEventEP11QMouseEvent", C.callback_ZN7QSlider14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionSlider *)
//export callback_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider
func callback_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSlider.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider", C.callback_ZNK7QSlider15initStyleOptionEP18QStyleOptionSlider /*nil*/)
}

// QSize tabSizeHint(int)
//export callback_ZNK7QTabBar11tabSizeHintEi
func callback_ZNK7QTabBar11tabSizeHintEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QTabBar.tabSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabSizeHint", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QTabBar11tabSizeHintEi", C.callback_ZNK7QTabBar11tabSizeHintEi /*nil*/)
}

// QSize minimumTabSizeHint(int)
//export callback_ZNK7QTabBar18minimumTabSizeHintEi
func callback_ZNK7QTabBar18minimumTabSizeHintEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QTabBar.minimumTabSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "minimumTabSizeHint", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QTabBar18minimumTabSizeHintEi", C.callback_ZNK7QTabBar18minimumTabSizeHintEi /*nil*/)
}

// void tabInserted(int)
//export callback_ZN7QTabBar11tabInsertedEi
func callback_ZN7QTabBar11tabInsertedEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QTabBar.tabInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabInserted", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar11tabInsertedEi", C.callback_ZN7QTabBar11tabInsertedEi /*nil*/)
}

// void tabRemoved(int)
//export callback_ZN7QTabBar10tabRemovedEi
func callback_ZN7QTabBar10tabRemovedEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QTabBar.tabRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabRemoved", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar10tabRemovedEi", C.callback_ZN7QTabBar10tabRemovedEi /*nil*/)
}

// void tabLayoutChange()
//export callback_ZN7QTabBar15tabLayoutChangeEv
func callback_ZN7QTabBar15tabLayoutChangeEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTabBar.tabLayoutChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabLayoutChange")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar15tabLayoutChangeEv", C.callback_ZN7QTabBar15tabLayoutChangeEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN7QTabBar5eventEP6QEvent
func callback_ZN7QTabBar5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar5eventEP6QEvent", C.callback_ZN7QTabBar5eventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN7QTabBar11resizeEventEP12QResizeEvent
func callback_ZN7QTabBar11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar11resizeEventEP12QResizeEvent", C.callback_ZN7QTabBar11resizeEventEP12QResizeEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN7QTabBar9showEventEP10QShowEvent
func callback_ZN7QTabBar9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar9showEventEP10QShowEvent", C.callback_ZN7QTabBar9showEventEP10QShowEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN7QTabBar9hideEventEP10QHideEvent
func callback_ZN7QTabBar9hideEventEP10QHideEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar9hideEventEP10QHideEvent", C.callback_ZN7QTabBar9hideEventEP10QHideEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN7QTabBar10paintEventEP11QPaintEvent
func callback_ZN7QTabBar10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar10paintEventEP11QPaintEvent", C.callback_ZN7QTabBar10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN7QTabBar15mousePressEventEP11QMouseEvent
func callback_ZN7QTabBar15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar15mousePressEventEP11QMouseEvent", C.callback_ZN7QTabBar15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN7QTabBar14mouseMoveEventEP11QMouseEvent
func callback_ZN7QTabBar14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar14mouseMoveEventEP11QMouseEvent", C.callback_ZN7QTabBar14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN7QTabBar17mouseReleaseEventEP11QMouseEvent
func callback_ZN7QTabBar17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar17mouseReleaseEventEP11QMouseEvent", C.callback_ZN7QTabBar17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN7QTabBar10wheelEventEP11QWheelEvent
func callback_ZN7QTabBar10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar10wheelEventEP11QWheelEvent", C.callback_ZN7QTabBar10wheelEventEP11QWheelEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN7QTabBar13keyPressEventEP9QKeyEvent
func callback_ZN7QTabBar13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar13keyPressEventEP9QKeyEvent", C.callback_ZN7QTabBar13keyPressEventEP9QKeyEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN7QTabBar11changeEventEP6QEvent
func callback_ZN7QTabBar11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar11changeEventEP6QEvent", C.callback_ZN7QTabBar11changeEventEP6QEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN7QTabBar10timerEventEP11QTimerEvent
func callback_ZN7QTabBar10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabBar.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QTabBar10timerEventEP11QTimerEvent", C.callback_ZN7QTabBar10timerEventEP11QTimerEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionTab *, int)
//export callback_ZNK7QTabBar15initStyleOptionEP15QStyleOptionTabi
func callback_ZNK7QTabBar15initStyleOptionEP15QStyleOptionTabi(cthis unsafe.Pointer, option unsafe.Pointer /*666*/, tabIndex C.int) {
	// log.Println(cthis, "QTabBar.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option, tabIndex)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QTabBar15initStyleOptionEP15QStyleOptionTabi", C.callback_ZNK7QTabBar15initStyleOptionEP15QStyleOptionTabi /*nil*/)
}

// void tabInserted(int)
//export callback_ZN10QTabWidget11tabInsertedEi
func callback_ZN10QTabWidget11tabInsertedEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QTabWidget.tabInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabInserted", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget11tabInsertedEi", C.callback_ZN10QTabWidget11tabInsertedEi /*nil*/)
}

// void tabRemoved(int)
//export callback_ZN10QTabWidget10tabRemovedEi
func callback_ZN10QTabWidget10tabRemovedEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QTabWidget.tabRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "tabRemoved", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget10tabRemovedEi", C.callback_ZN10QTabWidget10tabRemovedEi /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN10QTabWidget9showEventEP10QShowEvent
func callback_ZN10QTabWidget9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget9showEventEP10QShowEvent", C.callback_ZN10QTabWidget9showEventEP10QShowEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN10QTabWidget11resizeEventEP12QResizeEvent
func callback_ZN10QTabWidget11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget11resizeEventEP12QResizeEvent", C.callback_ZN10QTabWidget11resizeEventEP12QResizeEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN10QTabWidget13keyPressEventEP9QKeyEvent
func callback_ZN10QTabWidget13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget13keyPressEventEP9QKeyEvent", C.callback_ZN10QTabWidget13keyPressEventEP9QKeyEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN10QTabWidget10paintEventEP11QPaintEvent
func callback_ZN10QTabWidget10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget10paintEventEP11QPaintEvent", C.callback_ZN10QTabWidget10paintEventEP11QPaintEvent /*nil*/)
}

// void setTabBar(class QTabBar *)
//export callback_ZN10QTabWidget9setTabBarEP7QTabBar
func callback_ZN10QTabWidget9setTabBarEP7QTabBar(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.setTabBar")
	rvx := ffiqt.CallbackAllInherits(cthis, "setTabBar", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget9setTabBarEP7QTabBar", C.callback_ZN10QTabWidget9setTabBarEP7QTabBar /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN10QTabWidget11changeEventEP6QEvent
func callback_ZN10QTabWidget11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget11changeEventEP6QEvent", C.callback_ZN10QTabWidget11changeEventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN10QTabWidget5eventEP6QEvent
func callback_ZN10QTabWidget5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTabWidget5eventEP6QEvent", C.callback_ZN10QTabWidget5eventEP6QEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionTabWidgetFrame *)
//export callback_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame
func callback_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTabWidget.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame", C.callback_ZNK10QTabWidget15initStyleOptionEP26QStyleOptionTabWidgetFrame /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QRubberBand5eventEP6QEvent
func callback_ZN11QRubberBand5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRubberBand.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QRubberBand5eventEP6QEvent", C.callback_ZN11QRubberBand5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN11QRubberBand10paintEventEP11QPaintEvent
func callback_ZN11QRubberBand10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRubberBand.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QRubberBand10paintEventEP11QPaintEvent", C.callback_ZN11QRubberBand10paintEventEP11QPaintEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN11QRubberBand11changeEventEP6QEvent
func callback_ZN11QRubberBand11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRubberBand.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QRubberBand11changeEventEP6QEvent", C.callback_ZN11QRubberBand11changeEventEP6QEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN11QRubberBand9showEventEP10QShowEvent
func callback_ZN11QRubberBand9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRubberBand.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QRubberBand9showEventEP10QShowEvent", C.callback_ZN11QRubberBand9showEventEP10QShowEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN11QRubberBand11resizeEventEP12QResizeEvent
func callback_ZN11QRubberBand11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRubberBand.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QRubberBand11resizeEventEP12QResizeEvent", C.callback_ZN11QRubberBand11resizeEventEP12QResizeEvent /*nil*/)
}

// void moveEvent(class QMoveEvent *)
//export callback_ZN11QRubberBand9moveEventEP10QMoveEvent
func callback_ZN11QRubberBand9moveEventEP10QMoveEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRubberBand.moveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QRubberBand9moveEventEP10QMoveEvent", C.callback_ZN11QRubberBand9moveEventEP10QMoveEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionRubberBand *)
//export callback_ZNK11QRubberBand15initStyleOptionEP22QStyleOptionRubberBand
func callback_ZNK11QRubberBand15initStyleOptionEP22QStyleOptionRubberBand(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRubberBand.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QRubberBand15initStyleOptionEP22QStyleOptionRubberBand", C.callback_ZNK11QRubberBand15initStyleOptionEP22QStyleOptionRubberBand /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN6QFrame5eventEP6QEvent
func callback_ZN6QFrame5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFrame.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QFrame5eventEP6QEvent", C.callback_ZN6QFrame5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN6QFrame10paintEventEP11QPaintEvent
func callback_ZN6QFrame10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFrame.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QFrame10paintEventEP11QPaintEvent", C.callback_ZN6QFrame10paintEventEP11QPaintEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN6QFrame11changeEventEP6QEvent
func callback_ZN6QFrame11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFrame.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QFrame11changeEventEP6QEvent", C.callback_ZN6QFrame11changeEventEP6QEvent /*nil*/)
}

// void drawFrame(class QPainter *)
//export callback_ZN6QFrame9drawFrameEP8QPainter
func callback_ZN6QFrame9drawFrameEP8QPainter(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFrame.drawFrame")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawFrame", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QFrame9drawFrameEP8QPainter", C.callback_ZN6QFrame9drawFrameEP8QPainter /*nil*/)
}

// void initStyleOption(class QStyleOptionFrame *)
//export callback_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame
func callback_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFrame.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame", C.callback_ZNK6QFrame15initStyleOptionEP17QStyleOptionFrame /*nil*/)
}

// void QStyleOptionFocusRect(int)
//export callback_ZN21QStyleOptionFocusRectC1Ei
func callback_ZN21QStyleOptionFocusRectC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionFocusRect.QStyleOptionFocusRect")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionFocusRect", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QStyleOptionFocusRectC1Ei", C.callback_ZN21QStyleOptionFocusRectC1Ei /*nil*/)
}

// void QStyleOptionFrame(int)
//export callback_ZN17QStyleOptionFrameC1Ei
func callback_ZN17QStyleOptionFrameC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionFrame.QStyleOptionFrame")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionFrame", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QStyleOptionFrameC1Ei", C.callback_ZN17QStyleOptionFrameC1Ei /*nil*/)
}

// void QStyleOptionTabWidgetFrame(int)
//export callback_ZN26QStyleOptionTabWidgetFrameC1Ei
func callback_ZN26QStyleOptionTabWidgetFrameC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionTabWidgetFrame.QStyleOptionTabWidgetFrame")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionTabWidgetFrame", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN26QStyleOptionTabWidgetFrameC1Ei", C.callback_ZN26QStyleOptionTabWidgetFrameC1Ei /*nil*/)
}

// void QStyleOptionTabBarBase(int)
//export callback_ZN22QStyleOptionTabBarBaseC1Ei
func callback_ZN22QStyleOptionTabBarBaseC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionTabBarBase.QStyleOptionTabBarBase")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionTabBarBase", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN22QStyleOptionTabBarBaseC1Ei", C.callback_ZN22QStyleOptionTabBarBaseC1Ei /*nil*/)
}

// void QStyleOptionHeader(int)
//export callback_ZN18QStyleOptionHeaderC1Ei
func callback_ZN18QStyleOptionHeaderC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionHeader.QStyleOptionHeader")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionHeader", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QStyleOptionHeaderC1Ei", C.callback_ZN18QStyleOptionHeaderC1Ei /*nil*/)
}

// void QStyleOptionButton(int)
//export callback_ZN18QStyleOptionButtonC1Ei
func callback_ZN18QStyleOptionButtonC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionButton.QStyleOptionButton")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionButton", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QStyleOptionButtonC1Ei", C.callback_ZN18QStyleOptionButtonC1Ei /*nil*/)
}

// void QStyleOptionTab(int)
//export callback_ZN15QStyleOptionTabC1Ei
func callback_ZN15QStyleOptionTabC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionTab.QStyleOptionTab")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionTab", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QStyleOptionTabC1Ei", C.callback_ZN15QStyleOptionTabC1Ei /*nil*/)
}

// void QStyleOptionToolBar(int)
//export callback_ZN19QStyleOptionToolBarC1Ei
func callback_ZN19QStyleOptionToolBarC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionToolBar.QStyleOptionToolBar")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionToolBar", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QStyleOptionToolBarC1Ei", C.callback_ZN19QStyleOptionToolBarC1Ei /*nil*/)
}

// void QStyleOptionProgressBar(int)
//export callback_ZN23QStyleOptionProgressBarC1Ei
func callback_ZN23QStyleOptionProgressBarC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionProgressBar.QStyleOptionProgressBar")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionProgressBar", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN23QStyleOptionProgressBarC1Ei", C.callback_ZN23QStyleOptionProgressBarC1Ei /*nil*/)
}

// void QStyleOptionMenuItem(int)
//export callback_ZN20QStyleOptionMenuItemC1Ei
func callback_ZN20QStyleOptionMenuItemC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionMenuItem.QStyleOptionMenuItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionMenuItem", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QStyleOptionMenuItemC1Ei", C.callback_ZN20QStyleOptionMenuItemC1Ei /*nil*/)
}

// void QStyleOptionDockWidget(int)
//export callback_ZN22QStyleOptionDockWidgetC1Ei
func callback_ZN22QStyleOptionDockWidgetC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionDockWidget.QStyleOptionDockWidget")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionDockWidget", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN22QStyleOptionDockWidgetC1Ei", C.callback_ZN22QStyleOptionDockWidgetC1Ei /*nil*/)
}

// void QStyleOptionViewItem(int)
//export callback_ZN20QStyleOptionViewItemC1Ei
func callback_ZN20QStyleOptionViewItemC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionViewItem.QStyleOptionViewItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionViewItem", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QStyleOptionViewItemC1Ei", C.callback_ZN20QStyleOptionViewItemC1Ei /*nil*/)
}

// void QStyleOptionToolBox(int)
//export callback_ZN19QStyleOptionToolBoxC1Ei
func callback_ZN19QStyleOptionToolBoxC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionToolBox.QStyleOptionToolBox")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionToolBox", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QStyleOptionToolBoxC1Ei", C.callback_ZN19QStyleOptionToolBoxC1Ei /*nil*/)
}

// void QStyleOptionRubberBand(int)
//export callback_ZN22QStyleOptionRubberBandC1Ei
func callback_ZN22QStyleOptionRubberBandC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionRubberBand.QStyleOptionRubberBand")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionRubberBand", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN22QStyleOptionRubberBandC1Ei", C.callback_ZN22QStyleOptionRubberBandC1Ei /*nil*/)
}

// void QStyleOptionSlider(int)
//export callback_ZN18QStyleOptionSliderC1Ei
func callback_ZN18QStyleOptionSliderC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionSlider.QStyleOptionSlider")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionSlider", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QStyleOptionSliderC1Ei", C.callback_ZN18QStyleOptionSliderC1Ei /*nil*/)
}

// void QStyleOptionSpinBox(int)
//export callback_ZN19QStyleOptionSpinBoxC1Ei
func callback_ZN19QStyleOptionSpinBoxC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionSpinBox.QStyleOptionSpinBox")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionSpinBox", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QStyleOptionSpinBoxC1Ei", C.callback_ZN19QStyleOptionSpinBoxC1Ei /*nil*/)
}

// void QStyleOptionToolButton(int)
//export callback_ZN22QStyleOptionToolButtonC1Ei
func callback_ZN22QStyleOptionToolButtonC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionToolButton.QStyleOptionToolButton")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionToolButton", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN22QStyleOptionToolButtonC1Ei", C.callback_ZN22QStyleOptionToolButtonC1Ei /*nil*/)
}

// void QStyleOptionComboBox(int)
//export callback_ZN20QStyleOptionComboBoxC1Ei
func callback_ZN20QStyleOptionComboBoxC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionComboBox.QStyleOptionComboBox")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionComboBox", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QStyleOptionComboBoxC1Ei", C.callback_ZN20QStyleOptionComboBoxC1Ei /*nil*/)
}

// void QStyleOptionTitleBar(int)
//export callback_ZN20QStyleOptionTitleBarC1Ei
func callback_ZN20QStyleOptionTitleBarC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionTitleBar.QStyleOptionTitleBar")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionTitleBar", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QStyleOptionTitleBarC1Ei", C.callback_ZN20QStyleOptionTitleBarC1Ei /*nil*/)
}

// void QStyleOptionGroupBox(int)
//export callback_ZN20QStyleOptionGroupBoxC1Ei
func callback_ZN20QStyleOptionGroupBoxC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionGroupBox.QStyleOptionGroupBox")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionGroupBox", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QStyleOptionGroupBoxC1Ei", C.callback_ZN20QStyleOptionGroupBoxC1Ei /*nil*/)
}

// void QStyleOptionSizeGrip(int)
//export callback_ZN20QStyleOptionSizeGripC1Ei
func callback_ZN20QStyleOptionSizeGripC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionSizeGrip.QStyleOptionSizeGrip")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionSizeGrip", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QStyleOptionSizeGripC1Ei", C.callback_ZN20QStyleOptionSizeGripC1Ei /*nil*/)
}

// void QStyleOptionGraphicsItem(int)
//export callback_ZN24QStyleOptionGraphicsItemC1Ei
func callback_ZN24QStyleOptionGraphicsItemC1Ei(cthis unsafe.Pointer, version C.int) {
	// log.Println(cthis, "QStyleOptionGraphicsItem.QStyleOptionGraphicsItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStyleOptionGraphicsItem", version)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN24QStyleOptionGraphicsItemC1Ei", C.callback_ZN24QStyleOptionGraphicsItemC1Ei /*nil*/)
}

// void setViewportMargins(int, int, int, int)
//export callback_ZN19QAbstractScrollArea18setViewportMarginsEiiii
func callback_ZN19QAbstractScrollArea18setViewportMarginsEiiii(cthis unsafe.Pointer, left C.int, top C.int, right C.int, bottom C.int) {
	// log.Println(cthis, "QAbstractScrollArea.setViewportMargins")
	rvx := ffiqt.CallbackAllInherits(cthis, "setViewportMargins", left, top, right, bottom)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea18setViewportMarginsEiiii", C.callback_ZN19QAbstractScrollArea18setViewportMarginsEiiii /*nil*/)
}

// void setViewportMargins(const class QMargins &)
//export callback_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins
func callback_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins(cthis unsafe.Pointer, margins unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractScrollArea.setViewportMargins")
	rvx := ffiqt.CallbackAllInherits(cthis, "setViewportMargins", margins)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins", C.callback_ZN19QAbstractScrollArea18setViewportMarginsERK8QMargins /*nil*/)
}

// QMargins viewportMargins()
//export callback_ZNK19QAbstractScrollArea15viewportMarginsEv
func callback_ZNK19QAbstractScrollArea15viewportMarginsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractScrollArea.viewportMargins")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportMargins")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK19QAbstractScrollArea15viewportMarginsEv", C.callback_ZNK19QAbstractScrollArea15viewportMarginsEv /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent
func callback_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent", C.callback_ZN19QAbstractScrollArea11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN19QAbstractScrollArea5eventEP6QEvent
func callback_ZN19QAbstractScrollArea5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea5eventEP6QEvent", C.callback_ZN19QAbstractScrollArea5eventEP6QEvent /*nil*/)
}

// bool viewportEvent(class QEvent *)
//export callback_ZN19QAbstractScrollArea13viewportEventEP6QEvent
func callback_ZN19QAbstractScrollArea13viewportEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.viewportEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea13viewportEventEP6QEvent", C.callback_ZN19QAbstractScrollArea13viewportEventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent
func callback_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent", C.callback_ZN19QAbstractScrollArea11resizeEventEP12QResizeEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent
func callback_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent", C.callback_ZN19QAbstractScrollArea10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent
func callback_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent", C.callback_ZN19QAbstractScrollArea15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent
func callback_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent", C.callback_ZN19QAbstractScrollArea17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN19QAbstractScrollArea21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent
func callback_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent", C.callback_ZN19QAbstractScrollArea14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent
func callback_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent", C.callback_ZN19QAbstractScrollArea10wheelEventEP11QWheelEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent
func callback_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent", C.callback_ZN19QAbstractScrollArea16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void dragEnterEvent(class QDragEnterEvent *)
//export callback_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent
func callback_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent", C.callback_ZN19QAbstractScrollArea14dragEnterEventEP15QDragEnterEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent
func callback_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent", C.callback_ZN19QAbstractScrollArea13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN19QAbstractScrollArea14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN19QAbstractScrollArea9dropEventEP10QDropEvent
func callback_ZN19QAbstractScrollArea9dropEventEP10QDropEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea9dropEventEP10QDropEvent", C.callback_ZN19QAbstractScrollArea9dropEventEP10QDropEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent
func callback_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractScrollArea.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent", C.callback_ZN19QAbstractScrollArea13keyPressEventEP9QKeyEvent /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN19QAbstractScrollArea16scrollContentsByEii
func callback_ZN19QAbstractScrollArea16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QAbstractScrollArea.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractScrollArea16scrollContentsByEii", C.callback_ZN19QAbstractScrollArea16scrollContentsByEii /*nil*/)
}

// QSize viewportSizeHint()
//export callback_ZNK19QAbstractScrollArea16viewportSizeHintEv
func callback_ZNK19QAbstractScrollArea16viewportSizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractScrollArea.viewportSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportSizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK19QAbstractScrollArea16viewportSizeHintEv", C.callback_ZNK19QAbstractScrollArea16viewportSizeHintEv /*nil*/)
}

// void rowsInserted(const class QModelIndex &, int, int)
//export callback_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii
func callback_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QAbstractItemView.rowsInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsInserted", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii", C.callback_ZN17QAbstractItemView12rowsInsertedERK11QModelIndexii /*nil*/)
}

// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
//export callback_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii
func callback_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QAbstractItemView.rowsAboutToBeRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsAboutToBeRemoved", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii", C.callback_ZN17QAbstractItemView20rowsAboutToBeRemovedERK11QModelIndexii /*nil*/)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
//export callback_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_
func callback_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_(cthis unsafe.Pointer, selected unsafe.Pointer /*555*/, deselected unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractItemView.selectionChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "selectionChanged", selected, deselected)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_", C.callback_ZN17QAbstractItemView16selectionChangedERK14QItemSelectionS2_ /*nil*/)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
//export callback_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_
func callback_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_(cthis unsafe.Pointer, current unsafe.Pointer /*555*/, previous unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractItemView.currentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentChanged", current, previous)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_", C.callback_ZN17QAbstractItemView14currentChangedERK11QModelIndexS2_ /*nil*/)
}

// void updateEditorData()
//export callback_ZN17QAbstractItemView16updateEditorDataEv
func callback_ZN17QAbstractItemView16updateEditorDataEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.updateEditorData")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateEditorData")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView16updateEditorDataEv", C.callback_ZN17QAbstractItemView16updateEditorDataEv /*nil*/)
}

// void updateEditorGeometries()
//export callback_ZN17QAbstractItemView22updateEditorGeometriesEv
func callback_ZN17QAbstractItemView22updateEditorGeometriesEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.updateEditorGeometries")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateEditorGeometries")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView22updateEditorGeometriesEv", C.callback_ZN17QAbstractItemView22updateEditorGeometriesEv /*nil*/)
}

// void updateGeometries()
//export callback_ZN17QAbstractItemView16updateGeometriesEv
func callback_ZN17QAbstractItemView16updateGeometriesEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.updateGeometries")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateGeometries")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView16updateGeometriesEv", C.callback_ZN17QAbstractItemView16updateGeometriesEv /*nil*/)
}

// void verticalScrollbarAction(int)
//export callback_ZN17QAbstractItemView23verticalScrollbarActionEi
func callback_ZN17QAbstractItemView23verticalScrollbarActionEi(cthis unsafe.Pointer, action C.int) {
	// log.Println(cthis, "QAbstractItemView.verticalScrollbarAction")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalScrollbarAction", action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView23verticalScrollbarActionEi", C.callback_ZN17QAbstractItemView23verticalScrollbarActionEi /*nil*/)
}

// void horizontalScrollbarAction(int)
//export callback_ZN17QAbstractItemView25horizontalScrollbarActionEi
func callback_ZN17QAbstractItemView25horizontalScrollbarActionEi(cthis unsafe.Pointer, action C.int) {
	// log.Println(cthis, "QAbstractItemView.horizontalScrollbarAction")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalScrollbarAction", action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView25horizontalScrollbarActionEi", C.callback_ZN17QAbstractItemView25horizontalScrollbarActionEi /*nil*/)
}

// void verticalScrollbarValueChanged(int)
//export callback_ZN17QAbstractItemView29verticalScrollbarValueChangedEi
func callback_ZN17QAbstractItemView29verticalScrollbarValueChangedEi(cthis unsafe.Pointer, value C.int) {
	// log.Println(cthis, "QAbstractItemView.verticalScrollbarValueChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalScrollbarValueChanged", value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView29verticalScrollbarValueChangedEi", C.callback_ZN17QAbstractItemView29verticalScrollbarValueChangedEi /*nil*/)
}

// void horizontalScrollbarValueChanged(int)
//export callback_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi
func callback_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi(cthis unsafe.Pointer, value C.int) {
	// log.Println(cthis, "QAbstractItemView.horizontalScrollbarValueChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalScrollbarValueChanged", value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi", C.callback_ZN17QAbstractItemView31horizontalScrollbarValueChangedEi /*nil*/)
}

// void closeEditor(class QWidget *, class QAbstractItemDelegate::EndEditHint)
//export callback_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE
func callback_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE(cthis unsafe.Pointer, editor unsafe.Pointer /*666*/, hint C.int) {
	// log.Println(cthis, "QAbstractItemView.closeEditor")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEditor", editor, hint)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE", C.callback_ZN17QAbstractItemView11closeEditorEP7QWidgetN21QAbstractItemDelegate11EndEditHintE /*nil*/)
}

// void commitData(class QWidget *)
//export callback_ZN17QAbstractItemView10commitDataEP7QWidget
func callback_ZN17QAbstractItemView10commitDataEP7QWidget(cthis unsafe.Pointer, editor unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.commitData")
	rvx := ffiqt.CallbackAllInherits(cthis, "commitData", editor)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView10commitDataEP7QWidget", C.callback_ZN17QAbstractItemView10commitDataEP7QWidget /*nil*/)
}

// void editorDestroyed(class QObject *)
//export callback_ZN17QAbstractItemView15editorDestroyedEP7QObject
func callback_ZN17QAbstractItemView15editorDestroyedEP7QObject(cthis unsafe.Pointer, editor unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.editorDestroyed")
	rvx := ffiqt.CallbackAllInherits(cthis, "editorDestroyed", editor)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView15editorDestroyedEP7QObject", C.callback_ZN17QAbstractItemView15editorDestroyedEP7QObject /*nil*/)
}

// void setHorizontalStepsPerItem(int)
//export callback_ZN17QAbstractItemView25setHorizontalStepsPerItemEi
func callback_ZN17QAbstractItemView25setHorizontalStepsPerItemEi(cthis unsafe.Pointer, steps C.int) {
	// log.Println(cthis, "QAbstractItemView.setHorizontalStepsPerItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "setHorizontalStepsPerItem", steps)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView25setHorizontalStepsPerItemEi", C.callback_ZN17QAbstractItemView25setHorizontalStepsPerItemEi /*nil*/)
}

// int horizontalStepsPerItem()
//export callback_ZNK17QAbstractItemView22horizontalStepsPerItemEv
func callback_ZNK17QAbstractItemView22horizontalStepsPerItemEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.horizontalStepsPerItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalStepsPerItem")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView22horizontalStepsPerItemEv", C.callback_ZNK17QAbstractItemView22horizontalStepsPerItemEv /*nil*/)
}

// void setVerticalStepsPerItem(int)
//export callback_ZN17QAbstractItemView23setVerticalStepsPerItemEi
func callback_ZN17QAbstractItemView23setVerticalStepsPerItemEi(cthis unsafe.Pointer, steps C.int) {
	// log.Println(cthis, "QAbstractItemView.setVerticalStepsPerItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "setVerticalStepsPerItem", steps)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView23setVerticalStepsPerItemEi", C.callback_ZN17QAbstractItemView23setVerticalStepsPerItemEi /*nil*/)
}

// int verticalStepsPerItem()
//export callback_ZNK17QAbstractItemView20verticalStepsPerItemEv
func callback_ZNK17QAbstractItemView20verticalStepsPerItemEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.verticalStepsPerItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalStepsPerItem")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView20verticalStepsPerItemEv", C.callback_ZNK17QAbstractItemView20verticalStepsPerItemEv /*nil*/)
}

// int horizontalOffset()
//export callback_ZNK17QAbstractItemView16horizontalOffsetEv
func callback_ZNK17QAbstractItemView16horizontalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.horizontalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView16horizontalOffsetEv", C.callback_ZNK17QAbstractItemView16horizontalOffsetEv /*nil*/)
}

// int verticalOffset()
//export callback_ZNK17QAbstractItemView14verticalOffsetEv
func callback_ZNK17QAbstractItemView14verticalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.verticalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView14verticalOffsetEv", C.callback_ZNK17QAbstractItemView14verticalOffsetEv /*nil*/)
}

// bool isIndexHidden(const class QModelIndex &)
//export callback_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex
func callback_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractItemView.isIndexHidden")
	rvx := ffiqt.CallbackAllInherits(cthis, "isIndexHidden", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex", C.callback_ZNK17QAbstractItemView13isIndexHiddenERK11QModelIndex /*nil*/)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
//export callback_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE
func callback_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(cthis unsafe.Pointer, rect unsafe.Pointer /*555*/, command C.int) {
	// log.Println(cthis, "QAbstractItemView.setSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "setSelection", rect, command)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", C.callback_ZN17QAbstractItemView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE /*nil*/)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
//export callback_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection
func callback_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection(cthis unsafe.Pointer, selection unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractItemView.visualRegionForSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "visualRegionForSelection", selection)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection", C.callback_ZNK17QAbstractItemView24visualRegionForSelectionERK14QItemSelection /*nil*/)
}

// bool edit(const class QModelIndex &, enum QAbstractItemView::EditTrigger, class QEvent *)
//export callback_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent
func callback_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent(cthis unsafe.Pointer, index unsafe.Pointer /*555*/, trigger C.int, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.edit")
	rvx := ffiqt.CallbackAllInherits(cthis, "edit", index, trigger, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent", C.callback_ZN17QAbstractItemView4editERK11QModelIndexNS_11EditTriggerEP6QEvent /*nil*/)
}

// QItemSelectionModel::SelectionFlags selectionCommand(const class QModelIndex &, const class QEvent *)
//export callback_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent
func callback_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent(cthis unsafe.Pointer, index unsafe.Pointer /*555*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.selectionCommand")
	rvx := ffiqt.CallbackAllInherits(cthis, "selectionCommand", index, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent", C.callback_ZNK17QAbstractItemView16selectionCommandERK11QModelIndexPK6QEvent /*nil*/)
}

// QStyleOptionViewItem viewOptions()
//export callback_ZNK17QAbstractItemView11viewOptionsEv
func callback_ZNK17QAbstractItemView11viewOptionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.viewOptions")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewOptions")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView11viewOptionsEv", C.callback_ZNK17QAbstractItemView11viewOptionsEv /*nil*/)
}

// QAbstractItemView::State state()
//export callback_ZNK17QAbstractItemView5stateEv
func callback_ZNK17QAbstractItemView5stateEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.state")
	rvx := ffiqt.CallbackAllInherits(cthis, "state")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView5stateEv", C.callback_ZNK17QAbstractItemView5stateEv /*nil*/)
}

// void setState(enum QAbstractItemView::State)
//export callback_ZN17QAbstractItemView8setStateENS_5StateE
func callback_ZN17QAbstractItemView8setStateENS_5StateE(cthis unsafe.Pointer, state C.int) {
	// log.Println(cthis, "QAbstractItemView.setState")
	rvx := ffiqt.CallbackAllInherits(cthis, "setState", state)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView8setStateENS_5StateE", C.callback_ZN17QAbstractItemView8setStateENS_5StateE /*nil*/)
}

// void scheduleDelayedItemsLayout()
//export callback_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv
func callback_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.scheduleDelayedItemsLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "scheduleDelayedItemsLayout")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv", C.callback_ZN17QAbstractItemView26scheduleDelayedItemsLayoutEv /*nil*/)
}

// void executeDelayedItemsLayout()
//export callback_ZN17QAbstractItemView25executeDelayedItemsLayoutEv
func callback_ZN17QAbstractItemView25executeDelayedItemsLayoutEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.executeDelayedItemsLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "executeDelayedItemsLayout")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView25executeDelayedItemsLayoutEv", C.callback_ZN17QAbstractItemView25executeDelayedItemsLayoutEv /*nil*/)
}

// void setDirtyRegion(const class QRegion &)
//export callback_ZN17QAbstractItemView14setDirtyRegionERK7QRegion
func callback_ZN17QAbstractItemView14setDirtyRegionERK7QRegion(cthis unsafe.Pointer, region unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractItemView.setDirtyRegion")
	rvx := ffiqt.CallbackAllInherits(cthis, "setDirtyRegion", region)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView14setDirtyRegionERK7QRegion", C.callback_ZN17QAbstractItemView14setDirtyRegionERK7QRegion /*nil*/)
}

// void scrollDirtyRegion(int, int)
//export callback_ZN17QAbstractItemView17scrollDirtyRegionEii
func callback_ZN17QAbstractItemView17scrollDirtyRegionEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QAbstractItemView.scrollDirtyRegion")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollDirtyRegion", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView17scrollDirtyRegionEii", C.callback_ZN17QAbstractItemView17scrollDirtyRegionEii /*nil*/)
}

// QPoint dirtyRegionOffset()
//export callback_ZNK17QAbstractItemView17dirtyRegionOffsetEv
func callback_ZNK17QAbstractItemView17dirtyRegionOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.dirtyRegionOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "dirtyRegionOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView17dirtyRegionOffsetEv", C.callback_ZNK17QAbstractItemView17dirtyRegionOffsetEv /*nil*/)
}

// void startAutoScroll()
//export callback_ZN17QAbstractItemView15startAutoScrollEv
func callback_ZN17QAbstractItemView15startAutoScrollEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.startAutoScroll")
	rvx := ffiqt.CallbackAllInherits(cthis, "startAutoScroll")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView15startAutoScrollEv", C.callback_ZN17QAbstractItemView15startAutoScrollEv /*nil*/)
}

// void stopAutoScroll()
//export callback_ZN17QAbstractItemView14stopAutoScrollEv
func callback_ZN17QAbstractItemView14stopAutoScrollEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.stopAutoScroll")
	rvx := ffiqt.CallbackAllInherits(cthis, "stopAutoScroll")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView14stopAutoScrollEv", C.callback_ZN17QAbstractItemView14stopAutoScrollEv /*nil*/)
}

// void doAutoScroll()
//export callback_ZN17QAbstractItemView12doAutoScrollEv
func callback_ZN17QAbstractItemView12doAutoScrollEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.doAutoScroll")
	rvx := ffiqt.CallbackAllInherits(cthis, "doAutoScroll")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView12doAutoScrollEv", C.callback_ZN17QAbstractItemView12doAutoScrollEv /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN17QAbstractItemView18focusNextPrevChildEb
func callback_ZN17QAbstractItemView18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QAbstractItemView.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView18focusNextPrevChildEb", C.callback_ZN17QAbstractItemView18focusNextPrevChildEb /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN17QAbstractItemView5eventEP6QEvent
func callback_ZN17QAbstractItemView5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView5eventEP6QEvent", C.callback_ZN17QAbstractItemView5eventEP6QEvent /*nil*/)
}

// bool viewportEvent(class QEvent *)
//export callback_ZN17QAbstractItemView13viewportEventEP6QEvent
func callback_ZN17QAbstractItemView13viewportEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.viewportEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView13viewportEventEP6QEvent", C.callback_ZN17QAbstractItemView13viewportEventEP6QEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent
func callback_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent", C.callback_ZN17QAbstractItemView15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent
func callback_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent", C.callback_ZN17QAbstractItemView14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent
func callback_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent", C.callback_ZN17QAbstractItemView17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN17QAbstractItemView21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void dragEnterEvent(class QDragEnterEvent *)
//export callback_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent
func callback_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent", C.callback_ZN17QAbstractItemView14dragEnterEventEP15QDragEnterEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent
func callback_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent", C.callback_ZN17QAbstractItemView13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN17QAbstractItemView14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN17QAbstractItemView9dropEventEP10QDropEvent
func callback_ZN17QAbstractItemView9dropEventEP10QDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView9dropEventEP10QDropEvent", C.callback_ZN17QAbstractItemView9dropEventEP10QDropEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN17QAbstractItemView12focusInEventEP11QFocusEvent
func callback_ZN17QAbstractItemView12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView12focusInEventEP11QFocusEvent", C.callback_ZN17QAbstractItemView12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent
func callback_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent", C.callback_ZN17QAbstractItemView13focusOutEventEP11QFocusEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent
func callback_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent", C.callback_ZN17QAbstractItemView13keyPressEventEP9QKeyEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN17QAbstractItemView11resizeEventEP12QResizeEvent
func callback_ZN17QAbstractItemView11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView11resizeEventEP12QResizeEvent", C.callback_ZN17QAbstractItemView11resizeEventEP12QResizeEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN17QAbstractItemView10timerEventEP11QTimerEvent
func callback_ZN17QAbstractItemView10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView10timerEventEP11QTimerEvent", C.callback_ZN17QAbstractItemView10timerEventEP11QTimerEvent /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent
func callback_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemView.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent", C.callback_ZN17QAbstractItemView16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// QAbstractItemView::DropIndicatorPosition dropIndicatorPosition()
//export callback_ZNK17QAbstractItemView21dropIndicatorPositionEv
func callback_ZNK17QAbstractItemView21dropIndicatorPositionEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.dropIndicatorPosition")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropIndicatorPosition")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView21dropIndicatorPositionEv", C.callback_ZNK17QAbstractItemView21dropIndicatorPositionEv /*nil*/)
}

// QSize viewportSizeHint()
//export callback_ZNK17QAbstractItemView16viewportSizeHintEv
func callback_ZNK17QAbstractItemView16viewportSizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemView.viewportSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportSizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAbstractItemView16viewportSizeHintEv", C.callback_ZNK17QAbstractItemView16viewportSizeHintEv /*nil*/)
}

// void ~QAccessibleWidget()
//export callback_ZN17QAccessibleWidgetD1Ev
func callback_ZN17QAccessibleWidgetD1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAccessibleWidget.~QAccessibleWidget")
	rvx := ffiqt.CallbackAllInherits(cthis, "~QAccessibleWidget")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAccessibleWidgetD1Ev", C.callback_ZN17QAccessibleWidgetD1Ev /*nil*/)
}

// QWidget * widget()
//export callback_ZNK17QAccessibleWidget6widgetEv
func callback_ZNK17QAccessibleWidget6widgetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAccessibleWidget.widget")
	rvx := ffiqt.CallbackAllInherits(cthis, "widget")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAccessibleWidget6widgetEv", C.callback_ZNK17QAccessibleWidget6widgetEv /*nil*/)
}

// QObject * parentObject()
//export callback_ZNK17QAccessibleWidget12parentObjectEv
func callback_ZNK17QAccessibleWidget12parentObjectEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAccessibleWidget.parentObject")
	rvx := ffiqt.CallbackAllInherits(cthis, "parentObject")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QAccessibleWidget12parentObjectEv", C.callback_ZNK17QAccessibleWidget12parentObjectEv /*nil*/)
}

// void addControllingSignal(const class QString &)
//export callback_ZN17QAccessibleWidget20addControllingSignalERK7QString
func callback_ZN17QAccessibleWidget20addControllingSignalERK7QString(cthis unsafe.Pointer, signal unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAccessibleWidget.addControllingSignal")
	rvx := ffiqt.CallbackAllInherits(cthis, "addControllingSignal", signal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QAccessibleWidget20addControllingSignalERK7QString", C.callback_ZN17QAccessibleWidget20addControllingSignalERK7QString /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN7QAction5eventEP6QEvent
func callback_ZN7QAction5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAction.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QAction5eventEP6QEvent", C.callback_ZN7QAction5eventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN12QApplication5eventEP6QEvent
func callback_ZN12QApplication5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QApplication.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QApplication5eventEP6QEvent", C.callback_ZN12QApplication5eventEP6QEvent /*nil*/)
}

// void widgetEvent(class QEvent *)
//export callback_ZN7QLayout11widgetEventEP6QEvent
func callback_ZN7QLayout11widgetEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLayout.widgetEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "widgetEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QLayout11widgetEventEP6QEvent", C.callback_ZN7QLayout11widgetEventEP6QEvent /*nil*/)
}

// void childEvent(class QChildEvent *)
//export callback_ZN7QLayout10childEventEP11QChildEvent
func callback_ZN7QLayout10childEventEP11QChildEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLayout.childEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "childEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QLayout10childEventEP11QChildEvent", C.callback_ZN7QLayout10childEventEP11QChildEvent /*nil*/)
}

// void addChildLayout(class QLayout *)
//export callback_ZN7QLayout14addChildLayoutEPS_
func callback_ZN7QLayout14addChildLayoutEPS_(cthis unsafe.Pointer, l unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLayout.addChildLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "addChildLayout", l)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QLayout14addChildLayoutEPS_", C.callback_ZN7QLayout14addChildLayoutEPS_ /*nil*/)
}

// void addChildWidget(class QWidget *)
//export callback_ZN7QLayout14addChildWidgetEP7QWidget
func callback_ZN7QLayout14addChildWidgetEP7QWidget(cthis unsafe.Pointer, w unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLayout.addChildWidget")
	rvx := ffiqt.CallbackAllInherits(cthis, "addChildWidget", w)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QLayout14addChildWidgetEP7QWidget", C.callback_ZN7QLayout14addChildWidgetEP7QWidget /*nil*/)
}

// bool adoptLayout(class QLayout *)
//export callback_ZN7QLayout11adoptLayoutEPS_
func callback_ZN7QLayout11adoptLayoutEPS_(cthis unsafe.Pointer, layout unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLayout.adoptLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "adoptLayout", layout)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QLayout11adoptLayoutEPS_", C.callback_ZN7QLayout11adoptLayoutEPS_ /*nil*/)
}

// QRect alignmentRect(const class QRect &)
//export callback_ZNK7QLayout13alignmentRectERK5QRect
func callback_ZNK7QLayout13alignmentRectERK5QRect(cthis unsafe.Pointer, arg0 unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QLayout.alignmentRect")
	rvx := ffiqt.CallbackAllInherits(cthis, "alignmentRect", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QLayout13alignmentRectERK5QRect", C.callback_ZNK7QLayout13alignmentRectERK5QRect /*nil*/)
}

// void addItem(class QLayoutItem *)
//export callback_ZN11QGridLayout7addItemEP11QLayoutItem
func callback_ZN11QGridLayout7addItemEP11QLayoutItem(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGridLayout.addItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "addItem", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QGridLayout7addItemEP11QLayoutItem", C.callback_ZN11QGridLayout7addItemEP11QLayoutItem /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QCalendarWidget5eventEP6QEvent
func callback_ZN15QCalendarWidget5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCalendarWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QCalendarWidget5eventEP6QEvent", C.callback_ZN15QCalendarWidget5eventEP6QEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent
func callback_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, watched unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCalendarWidget.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", watched, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent", C.callback_ZN15QCalendarWidget11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent
func callback_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCalendarWidget.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent", C.callback_ZN15QCalendarWidget15mousePressEventEP11QMouseEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN15QCalendarWidget11resizeEventEP12QResizeEvent
func callback_ZN15QCalendarWidget11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCalendarWidget.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QCalendarWidget11resizeEventEP12QResizeEvent", C.callback_ZN15QCalendarWidget11resizeEventEP12QResizeEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent
func callback_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCalendarWidget.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent", C.callback_ZN15QCalendarWidget13keyPressEventEP9QKeyEvent /*nil*/)
}

// void paintCell(class QPainter *, const class QRect &, const class QDate &)
//export callback_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate
func callback_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/, date unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QCalendarWidget.paintCell")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintCell", painter, rect, date)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate", C.callback_ZNK15QCalendarWidget9paintCellEP8QPainterRK5QRectRK5QDate /*nil*/)
}

// void updateCell(const class QDate &)
//export callback_ZN15QCalendarWidget10updateCellERK5QDate
func callback_ZN15QCalendarWidget10updateCellERK5QDate(cthis unsafe.Pointer, date unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QCalendarWidget.updateCell")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCell", date)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QCalendarWidget10updateCellERK5QDate", C.callback_ZN15QCalendarWidget10updateCellERK5QDate /*nil*/)
}

// void updateCells()
//export callback_ZN15QCalendarWidget11updateCellsEv
func callback_ZN15QCalendarWidget11updateCellsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QCalendarWidget.updateCells")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCells")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QCalendarWidget11updateCellsEv", C.callback_ZN15QCalendarWidget11updateCellsEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QCheckBox5eventEP6QEvent
func callback_ZN9QCheckBox5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCheckBox.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QCheckBox5eventEP6QEvent", C.callback_ZN9QCheckBox5eventEP6QEvent /*nil*/)
}

// bool hitButton(const class QPoint &)
//export callback_ZNK9QCheckBox9hitButtonERK6QPoint
func callback_ZNK9QCheckBox9hitButtonERK6QPoint(cthis unsafe.Pointer, pos unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QCheckBox.hitButton")
	rvx := ffiqt.CallbackAllInherits(cthis, "hitButton", pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QCheckBox9hitButtonERK6QPoint", C.callback_ZNK9QCheckBox9hitButtonERK6QPoint /*nil*/)
}

// void checkStateSet()
//export callback_ZN9QCheckBox13checkStateSetEv
func callback_ZN9QCheckBox13checkStateSetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QCheckBox.checkStateSet")
	rvx := ffiqt.CallbackAllInherits(cthis, "checkStateSet")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QCheckBox13checkStateSetEv", C.callback_ZN9QCheckBox13checkStateSetEv /*nil*/)
}

// void nextCheckState()
//export callback_ZN9QCheckBox14nextCheckStateEv
func callback_ZN9QCheckBox14nextCheckStateEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QCheckBox.nextCheckState")
	rvx := ffiqt.CallbackAllInherits(cthis, "nextCheckState")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QCheckBox14nextCheckStateEv", C.callback_ZN9QCheckBox14nextCheckStateEv /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QCheckBox10paintEventEP11QPaintEvent
func callback_ZN9QCheckBox10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCheckBox.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QCheckBox10paintEventEP11QPaintEvent", C.callback_ZN9QCheckBox10paintEventEP11QPaintEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent
func callback_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCheckBox.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent", C.callback_ZN9QCheckBox14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionButton *)
//export callback_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton
func callback_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCheckBox.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton", C.callback_ZNK9QCheckBox15initStyleOptionEP18QStyleOptionButton /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN7QDialog13keyPressEventEP9QKeyEvent
func callback_ZN7QDialog13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialog.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QDialog13keyPressEventEP9QKeyEvent", C.callback_ZN7QDialog13keyPressEventEP9QKeyEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN7QDialog10closeEventEP11QCloseEvent
func callback_ZN7QDialog10closeEventEP11QCloseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialog.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QDialog10closeEventEP11QCloseEvent", C.callback_ZN7QDialog10closeEventEP11QCloseEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN7QDialog9showEventEP10QShowEvent
func callback_ZN7QDialog9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialog.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QDialog9showEventEP10QShowEvent", C.callback_ZN7QDialog9showEventEP10QShowEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN7QDialog11resizeEventEP12QResizeEvent
func callback_ZN7QDialog11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialog.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QDialog11resizeEventEP12QResizeEvent", C.callback_ZN7QDialog11resizeEventEP12QResizeEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN7QDialog16contextMenuEventEP17QContextMenuEvent
func callback_ZN7QDialog16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialog.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QDialog16contextMenuEventEP17QContextMenuEvent", C.callback_ZN7QDialog16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN7QDialog11eventFilterEP7QObjectP6QEvent
func callback_ZN7QDialog11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialog.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QDialog11eventFilterEP7QObjectP6QEvent", C.callback_ZN7QDialog11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void adjustPosition(class QWidget *)
//export callback_ZN7QDialog14adjustPositionEP7QWidget
func callback_ZN7QDialog14adjustPositionEP7QWidget(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialog.adjustPosition")
	rvx := ffiqt.CallbackAllInherits(cthis, "adjustPosition", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QDialog14adjustPositionEP7QWidget", C.callback_ZN7QDialog14adjustPositionEP7QWidget /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN12QColorDialog11changeEventEP6QEvent
func callback_ZN12QColorDialog11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QColorDialog.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QColorDialog11changeEventEP6QEvent", C.callback_ZN12QColorDialog11changeEventEP6QEvent /*nil*/)
}

// void done(int)
//export callback_ZN12QColorDialog4doneEi
func callback_ZN12QColorDialog4doneEi(cthis unsafe.Pointer, result C.int) {
	// log.Println(cthis, "QColorDialog.done")
	rvx := ffiqt.CallbackAllInherits(cthis, "done", result)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QColorDialog4doneEi", C.callback_ZN12QColorDialog4doneEi /*nil*/)
}

// bool isIndexHidden(const class QModelIndex &)
//export callback_ZNK11QColumnView13isIndexHiddenERK11QModelIndex
func callback_ZNK11QColumnView13isIndexHiddenERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QColumnView.isIndexHidden")
	rvx := ffiqt.CallbackAllInherits(cthis, "isIndexHidden", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QColumnView13isIndexHiddenERK11QModelIndex", C.callback_ZNK11QColumnView13isIndexHiddenERK11QModelIndex /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN11QColumnView11resizeEventEP12QResizeEvent
func callback_ZN11QColumnView11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QColumnView.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QColumnView11resizeEventEP12QResizeEvent", C.callback_ZN11QColumnView11resizeEventEP12QResizeEvent /*nil*/)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
//export callback_ZN11QColumnView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE
func callback_ZN11QColumnView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(cthis unsafe.Pointer, rect unsafe.Pointer /*555*/, command C.int) {
	// log.Println(cthis, "QColumnView.setSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "setSelection", rect, command)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QColumnView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", C.callback_ZN11QColumnView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE /*nil*/)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
//export callback_ZNK11QColumnView24visualRegionForSelectionERK14QItemSelection
func callback_ZNK11QColumnView24visualRegionForSelectionERK14QItemSelection(cthis unsafe.Pointer, selection unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QColumnView.visualRegionForSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "visualRegionForSelection", selection)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QColumnView24visualRegionForSelectionERK14QItemSelection", C.callback_ZNK11QColumnView24visualRegionForSelectionERK14QItemSelection /*nil*/)
}

// int horizontalOffset()
//export callback_ZNK11QColumnView16horizontalOffsetEv
func callback_ZNK11QColumnView16horizontalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QColumnView.horizontalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QColumnView16horizontalOffsetEv", C.callback_ZNK11QColumnView16horizontalOffsetEv /*nil*/)
}

// int verticalOffset()
//export callback_ZNK11QColumnView14verticalOffsetEv
func callback_ZNK11QColumnView14verticalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QColumnView.verticalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QColumnView14verticalOffsetEv", C.callback_ZNK11QColumnView14verticalOffsetEv /*nil*/)
}

// void rowsInserted(const class QModelIndex &, int, int)
//export callback_ZN11QColumnView12rowsInsertedERK11QModelIndexii
func callback_ZN11QColumnView12rowsInsertedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QColumnView.rowsInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsInserted", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QColumnView12rowsInsertedERK11QModelIndexii", C.callback_ZN11QColumnView12rowsInsertedERK11QModelIndexii /*nil*/)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
//export callback_ZN11QColumnView14currentChangedERK11QModelIndexS2_
func callback_ZN11QColumnView14currentChangedERK11QModelIndexS2_(cthis unsafe.Pointer, current unsafe.Pointer /*555*/, previous unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QColumnView.currentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentChanged", current, previous)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QColumnView14currentChangedERK11QModelIndexS2_", C.callback_ZN11QColumnView14currentChangedERK11QModelIndexS2_ /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN11QColumnView16scrollContentsByEii
func callback_ZN11QColumnView16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QColumnView.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QColumnView16scrollContentsByEii", C.callback_ZN11QColumnView16scrollContentsByEii /*nil*/)
}

// QAbstractItemView * createColumn(const class QModelIndex &)
//export callback_ZN11QColumnView12createColumnERK11QModelIndex
func callback_ZN11QColumnView12createColumnERK11QModelIndex(cthis unsafe.Pointer, rootIndex unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QColumnView.createColumn")
	rvx := ffiqt.CallbackAllInherits(cthis, "createColumn", rootIndex)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QColumnView12createColumnERK11QModelIndex", C.callback_ZN11QColumnView12createColumnERK11QModelIndex /*nil*/)
}

// void initializeColumn(class QAbstractItemView *)
//export callback_ZNK11QColumnView16initializeColumnEP17QAbstractItemView
func callback_ZNK11QColumnView16initializeColumnEP17QAbstractItemView(cthis unsafe.Pointer, column unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QColumnView.initializeColumn")
	rvx := ffiqt.CallbackAllInherits(cthis, "initializeColumn", column)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QColumnView16initializeColumnEP17QAbstractItemView", C.callback_ZNK11QColumnView16initializeColumnEP17QAbstractItemView /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN9QComboBox12focusInEventEP11QFocusEvent
func callback_ZN9QComboBox12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox12focusInEventEP11QFocusEvent", C.callback_ZN9QComboBox12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN9QComboBox13focusOutEventEP11QFocusEvent
func callback_ZN9QComboBox13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox13focusOutEventEP11QFocusEvent", C.callback_ZN9QComboBox13focusOutEventEP11QFocusEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN9QComboBox11changeEventEP6QEvent
func callback_ZN9QComboBox11changeEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox11changeEventEP6QEvent", C.callback_ZN9QComboBox11changeEventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN9QComboBox11resizeEventEP12QResizeEvent
func callback_ZN9QComboBox11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox11resizeEventEP12QResizeEvent", C.callback_ZN9QComboBox11resizeEventEP12QResizeEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QComboBox10paintEventEP11QPaintEvent
func callback_ZN9QComboBox10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox10paintEventEP11QPaintEvent", C.callback_ZN9QComboBox10paintEventEP11QPaintEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN9QComboBox9showEventEP10QShowEvent
func callback_ZN9QComboBox9showEventEP10QShowEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox9showEventEP10QShowEvent", C.callback_ZN9QComboBox9showEventEP10QShowEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN9QComboBox9hideEventEP10QHideEvent
func callback_ZN9QComboBox9hideEventEP10QHideEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox9hideEventEP10QHideEvent", C.callback_ZN9QComboBox9hideEventEP10QHideEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN9QComboBox15mousePressEventEP11QMouseEvent
func callback_ZN9QComboBox15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox15mousePressEventEP11QMouseEvent", C.callback_ZN9QComboBox15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent
func callback_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent", C.callback_ZN9QComboBox17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN9QComboBox13keyPressEventEP9QKeyEvent
func callback_ZN9QComboBox13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox13keyPressEventEP9QKeyEvent", C.callback_ZN9QComboBox13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN9QComboBox15keyReleaseEventEP9QKeyEvent
func callback_ZN9QComboBox15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox15keyReleaseEventEP9QKeyEvent", C.callback_ZN9QComboBox15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN9QComboBox10wheelEventEP11QWheelEvent
func callback_ZN9QComboBox10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox10wheelEventEP11QWheelEvent", C.callback_ZN9QComboBox10wheelEventEP11QWheelEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent
func callback_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent", C.callback_ZN9QComboBox16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent
func callback_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent", C.callback_ZN9QComboBox16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionComboBox *)
//export callback_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox
func callback_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QComboBox.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox", C.callback_ZNK9QComboBox15initStyleOptionEP20QStyleOptionComboBox /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QPushButton5eventEP6QEvent
func callback_ZN11QPushButton5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPushButton.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QPushButton5eventEP6QEvent", C.callback_ZN11QPushButton5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN11QPushButton10paintEventEP11QPaintEvent
func callback_ZN11QPushButton10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPushButton.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QPushButton10paintEventEP11QPaintEvent", C.callback_ZN11QPushButton10paintEventEP11QPaintEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN11QPushButton13keyPressEventEP9QKeyEvent
func callback_ZN11QPushButton13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPushButton.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QPushButton13keyPressEventEP9QKeyEvent", C.callback_ZN11QPushButton13keyPressEventEP9QKeyEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN11QPushButton12focusInEventEP11QFocusEvent
func callback_ZN11QPushButton12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPushButton.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QPushButton12focusInEventEP11QFocusEvent", C.callback_ZN11QPushButton12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN11QPushButton13focusOutEventEP11QFocusEvent
func callback_ZN11QPushButton13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPushButton.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QPushButton13focusOutEventEP11QFocusEvent", C.callback_ZN11QPushButton13focusOutEventEP11QFocusEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionButton *)
//export callback_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton
func callback_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPushButton.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton", C.callback_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton /*nil*/)
}

// QSize sizeHint()
//export callback_ZNK18QCommandLinkButton8sizeHintEv
func callback_ZNK18QCommandLinkButton8sizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QCommandLinkButton.sizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QCommandLinkButton8sizeHintEv", C.callback_ZNK18QCommandLinkButton8sizeHintEv /*nil*/)
}

// int heightForWidth(int)
//export callback_ZNK18QCommandLinkButton14heightForWidthEi
func callback_ZNK18QCommandLinkButton14heightForWidthEi(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QCommandLinkButton.heightForWidth")
	rvx := ffiqt.CallbackAllInherits(cthis, "heightForWidth", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QCommandLinkButton14heightForWidthEi", C.callback_ZNK18QCommandLinkButton14heightForWidthEi /*nil*/)
}

// QSize minimumSizeHint()
//export callback_ZNK18QCommandLinkButton15minimumSizeHintEv
func callback_ZNK18QCommandLinkButton15minimumSizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QCommandLinkButton.minimumSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "minimumSizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QCommandLinkButton15minimumSizeHintEv", C.callback_ZNK18QCommandLinkButton15minimumSizeHintEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN18QCommandLinkButton5eventEP6QEvent
func callback_ZN18QCommandLinkButton5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCommandLinkButton.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QCommandLinkButton5eventEP6QEvent", C.callback_ZN18QCommandLinkButton5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN18QCommandLinkButton10paintEventEP11QPaintEvent
func callback_ZN18QCommandLinkButton10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCommandLinkButton.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QCommandLinkButton10paintEventEP11QPaintEvent", C.callback_ZN18QCommandLinkButton10paintEventEP11QPaintEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN10QCompleter11eventFilterEP7QObjectP6QEvent
func callback_ZN10QCompleter11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, o unsafe.Pointer /*666*/, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCompleter.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", o, e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QCompleter11eventFilterEP7QObjectP6QEvent", C.callback_ZN10QCompleter11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN10QCompleter5eventEP6QEvent
func callback_ZN10QCompleter5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCompleter.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QCompleter5eventEP6QEvent", C.callback_ZN10QCompleter5eventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent
func callback_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDateTimeEdit.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent", C.callback_ZN13QDateTimeEdit13keyPressEventEP9QKeyEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent
func callback_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDateTimeEdit.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent", C.callback_ZN13QDateTimeEdit10wheelEventEP11QWheelEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent
func callback_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDateTimeEdit.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent", C.callback_ZN13QDateTimeEdit12focusInEventEP11QFocusEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN13QDateTimeEdit18focusNextPrevChildEb
func callback_ZN13QDateTimeEdit18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QDateTimeEdit.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QDateTimeEdit18focusNextPrevChildEb", C.callback_ZN13QDateTimeEdit18focusNextPrevChildEb /*nil*/)
}

// QValidator::State validate(class QString &, int &)
//export callback_ZNK13QDateTimeEdit8validateER7QStringRi
func callback_ZNK13QDateTimeEdit8validateER7QStringRi(cthis unsafe.Pointer, input unsafe.Pointer /*555*/, pos unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QDateTimeEdit.validate")
	rvx := ffiqt.CallbackAllInherits(cthis, "validate", input, pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QDateTimeEdit8validateER7QStringRi", C.callback_ZNK13QDateTimeEdit8validateER7QStringRi /*nil*/)
}

// void fixup(class QString &)
//export callback_ZNK13QDateTimeEdit5fixupER7QString
func callback_ZNK13QDateTimeEdit5fixupER7QString(cthis unsafe.Pointer, input unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QDateTimeEdit.fixup")
	rvx := ffiqt.CallbackAllInherits(cthis, "fixup", input)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QDateTimeEdit5fixupER7QString", C.callback_ZNK13QDateTimeEdit5fixupER7QString /*nil*/)
}

// QDateTime dateTimeFromText(const class QString &)
//export callback_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString
func callback_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString(cthis unsafe.Pointer, text unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QDateTimeEdit.dateTimeFromText")
	rvx := ffiqt.CallbackAllInherits(cthis, "dateTimeFromText", text)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString", C.callback_ZNK13QDateTimeEdit16dateTimeFromTextERK7QString /*nil*/)
}

// QString textFromDateTime(const class QDateTime &)
//export callback_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime
func callback_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime(cthis unsafe.Pointer, dt unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QDateTimeEdit.textFromDateTime")
	rvx := ffiqt.CallbackAllInherits(cthis, "textFromDateTime", dt)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime", C.callback_ZNK13QDateTimeEdit16textFromDateTimeERK9QDateTime /*nil*/)
}

// QAbstractSpinBox::StepEnabled stepEnabled()
//export callback_ZNK13QDateTimeEdit11stepEnabledEv
func callback_ZNK13QDateTimeEdit11stepEnabledEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QDateTimeEdit.stepEnabled")
	rvx := ffiqt.CallbackAllInherits(cthis, "stepEnabled")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QDateTimeEdit11stepEnabledEv", C.callback_ZNK13QDateTimeEdit11stepEnabledEv /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent
func callback_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDateTimeEdit.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent", C.callback_ZN13QDateTimeEdit15mousePressEventEP11QMouseEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN13QDateTimeEdit10paintEventEP11QPaintEvent
func callback_ZN13QDateTimeEdit10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDateTimeEdit.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QDateTimeEdit10paintEventEP11QPaintEvent", C.callback_ZN13QDateTimeEdit10paintEventEP11QPaintEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionSpinBox *)
//export callback_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox
func callback_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDateTimeEdit.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox", C.callback_ZNK13QDateTimeEdit15initStyleOptionEP19QStyleOptionSpinBox /*nil*/)
}

// void QDateTimeEdit(const class QVariant &, class QVariant::Type, class QWidget *)
//export callback_ZN13QDateTimeEditC1ERK8QVariantNS0_4TypeEP7QWidget
func callback_ZN13QDateTimeEditC1ERK8QVariantNS0_4TypeEP7QWidget(cthis unsafe.Pointer, val unsafe.Pointer /*555*/, parserType C.int, parent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDateTimeEdit.QDateTimeEdit")
	rvx := ffiqt.CallbackAllInherits(cthis, "QDateTimeEdit", val, parserType, parent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QDateTimeEditC1ERK8QVariantNS0_4TypeEP7QWidget", C.callback_ZN13QDateTimeEditC1ERK8QVariantNS0_4TypeEP7QWidget /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN14QDesktopWidget11resizeEventEP12QResizeEvent
func callback_ZN14QDesktopWidget11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDesktopWidget.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QDesktopWidget11resizeEventEP12QResizeEvent", C.callback_ZN14QDesktopWidget11resizeEventEP12QResizeEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN5QDial5eventEP6QEvent
func callback_ZN5QDial5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDial.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QDial5eventEP6QEvent", C.callback_ZN5QDial5eventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN5QDial11resizeEventEP12QResizeEvent
func callback_ZN5QDial11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, re unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDial.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", re)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QDial11resizeEventEP12QResizeEvent", C.callback_ZN5QDial11resizeEventEP12QResizeEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN5QDial10paintEventEP11QPaintEvent
func callback_ZN5QDial10paintEventEP11QPaintEvent(cthis unsafe.Pointer, pe unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDial.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", pe)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QDial10paintEventEP11QPaintEvent", C.callback_ZN5QDial10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN5QDial15mousePressEventEP11QMouseEvent
func callback_ZN5QDial15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, me unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDial.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", me)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QDial15mousePressEventEP11QMouseEvent", C.callback_ZN5QDial15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN5QDial17mouseReleaseEventEP11QMouseEvent
func callback_ZN5QDial17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, me unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDial.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", me)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QDial17mouseReleaseEventEP11QMouseEvent", C.callback_ZN5QDial17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN5QDial14mouseMoveEventEP11QMouseEvent
func callback_ZN5QDial14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, me unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDial.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", me)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QDial14mouseMoveEventEP11QMouseEvent", C.callback_ZN5QDial14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void sliderChange(enum QAbstractSlider::SliderChange)
//export callback_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE
func callback_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE(cthis unsafe.Pointer, change C.int) {
	// log.Println(cthis, "QDial.sliderChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "sliderChange", change)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE", C.callback_ZN5QDial12sliderChangeEN15QAbstractSlider12SliderChangeE /*nil*/)
}

// void initStyleOption(class QStyleOptionSlider *)
//export callback_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider
func callback_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDial.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider", C.callback_ZNK5QDial15initStyleOptionEP18QStyleOptionSlider /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN16QDialogButtonBox11changeEventEP6QEvent
func callback_ZN16QDialogButtonBox11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialogButtonBox.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QDialogButtonBox11changeEventEP6QEvent", C.callback_ZN16QDialogButtonBox11changeEventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN16QDialogButtonBox5eventEP6QEvent
func callback_ZN16QDialogButtonBox5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDialogButtonBox.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QDialogButtonBox5eventEP6QEvent", C.callback_ZN16QDialogButtonBox5eventEP6QEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN11QDockWidget11changeEventEP6QEvent
func callback_ZN11QDockWidget11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDockWidget.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QDockWidget11changeEventEP6QEvent", C.callback_ZN11QDockWidget11changeEventEP6QEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN11QDockWidget10closeEventEP11QCloseEvent
func callback_ZN11QDockWidget10closeEventEP11QCloseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDockWidget.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QDockWidget10closeEventEP11QCloseEvent", C.callback_ZN11QDockWidget10closeEventEP11QCloseEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN11QDockWidget10paintEventEP11QPaintEvent
func callback_ZN11QDockWidget10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDockWidget.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QDockWidget10paintEventEP11QPaintEvent", C.callback_ZN11QDockWidget10paintEventEP11QPaintEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QDockWidget5eventEP6QEvent
func callback_ZN11QDockWidget5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDockWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QDockWidget5eventEP6QEvent", C.callback_ZN11QDockWidget5eventEP6QEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionDockWidget *)
//export callback_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget
func callback_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QDockWidget.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget", C.callback_ZNK11QDockWidget15initStyleOptionEP22QStyleOptionDockWidget /*nil*/)
}

// void done(int)
//export callback_ZN13QErrorMessage4doneEi
func callback_ZN13QErrorMessage4doneEi(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QErrorMessage.done")
	rvx := ffiqt.CallbackAllInherits(cthis, "done", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QErrorMessage4doneEi", C.callback_ZN13QErrorMessage4doneEi /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN13QErrorMessage11changeEventEP6QEvent
func callback_ZN13QErrorMessage11changeEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QErrorMessage.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QErrorMessage11changeEventEP6QEvent", C.callback_ZN13QErrorMessage11changeEventEP6QEvent /*nil*/)
}

// void done(int)
//export callback_ZN11QFileDialog4doneEi
func callback_ZN11QFileDialog4doneEi(cthis unsafe.Pointer, result C.int) {
	// log.Println(cthis, "QFileDialog.done")
	rvx := ffiqt.CallbackAllInherits(cthis, "done", result)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFileDialog4doneEi", C.callback_ZN11QFileDialog4doneEi /*nil*/)
}

// void accept()
//export callback_ZN11QFileDialog6acceptEv
func callback_ZN11QFileDialog6acceptEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QFileDialog.accept")
	rvx := ffiqt.CallbackAllInherits(cthis, "accept")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFileDialog6acceptEv", C.callback_ZN11QFileDialog6acceptEv /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN11QFileDialog11changeEventEP6QEvent
func callback_ZN11QFileDialog11changeEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFileDialog.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFileDialog11changeEventEP6QEvent", C.callback_ZN11QFileDialog11changeEventEP6QEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN16QFileSystemModel10timerEventEP11QTimerEvent
func callback_ZN16QFileSystemModel10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFileSystemModel.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QFileSystemModel10timerEventEP11QTimerEvent", C.callback_ZN16QFileSystemModel10timerEventEP11QTimerEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN16QFileSystemModel5eventEP6QEvent
func callback_ZN16QFileSystemModel5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFileSystemModel.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QFileSystemModel5eventEP6QEvent", C.callback_ZN16QFileSystemModel5eventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QFocusFrame5eventEP6QEvent
func callback_ZN11QFocusFrame5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFocusFrame.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFocusFrame5eventEP6QEvent", C.callback_ZN11QFocusFrame5eventEP6QEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent
func callback_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFocusFrame.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent", C.callback_ZN11QFocusFrame11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN11QFocusFrame10paintEventEP11QPaintEvent
func callback_ZN11QFocusFrame10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFocusFrame.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFocusFrame10paintEventEP11QPaintEvent", C.callback_ZN11QFocusFrame10paintEventEP11QPaintEvent /*nil*/)
}

// void initStyleOption(class QStyleOption *)
//export callback_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption
func callback_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFocusFrame.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption", C.callback_ZNK11QFocusFrame15initStyleOptionEP12QStyleOption /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN13QFontComboBox5eventEP6QEvent
func callback_ZN13QFontComboBox5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFontComboBox.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QFontComboBox5eventEP6QEvent", C.callback_ZN13QFontComboBox5eventEP6QEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN11QFontDialog11changeEventEP6QEvent
func callback_ZN11QFontDialog11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFontDialog.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFontDialog11changeEventEP6QEvent", C.callback_ZN11QFontDialog11changeEventEP6QEvent /*nil*/)
}

// void done(int)
//export callback_ZN11QFontDialog4doneEi
func callback_ZN11QFontDialog4doneEi(cthis unsafe.Pointer, result C.int) {
	// log.Println(cthis, "QFontDialog.done")
	rvx := ffiqt.CallbackAllInherits(cthis, "done", result)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFontDialog4doneEi", C.callback_ZN11QFontDialog4doneEi /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent
func callback_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, object unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFontDialog.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", object, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent", C.callback_ZN11QFontDialog11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void updateMicroFocus()
//export callback_ZN13QGraphicsItem16updateMicroFocusEv
func callback_ZN13QGraphicsItem16updateMicroFocusEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsItem.updateMicroFocus")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateMicroFocus")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem16updateMicroFocusEv", C.callback_ZN13QGraphicsItem16updateMicroFocusEv /*nil*/)
}

// bool sceneEventFilter(class QGraphicsItem *, class QEvent *)
//export callback_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent
func callback_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent(cthis unsafe.Pointer, watched unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.sceneEventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "sceneEventFilter", watched, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent", C.callback_ZN13QGraphicsItem16sceneEventFilterEPS_P6QEvent /*nil*/)
}

// bool sceneEvent(class QEvent *)
//export callback_ZN13QGraphicsItem10sceneEventEP6QEvent
func callback_ZN13QGraphicsItem10sceneEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.sceneEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "sceneEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem10sceneEventEP6QEvent", C.callback_ZN13QGraphicsItem10sceneEventEP6QEvent /*nil*/)
}

// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
//export callback_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent
func callback_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", C.callback_ZN13QGraphicsItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent /*nil*/)
}

// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent
func callback_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN13QGraphicsItem14dragEnterEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN13QGraphicsItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN13QGraphicsItem13dragMoveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dropEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent
func callback_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN13QGraphicsItem9dropEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN13QGraphicsItem12focusInEventEP11QFocusEvent
func callback_ZN13QGraphicsItem12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem12focusInEventEP11QFocusEvent", C.callback_ZN13QGraphicsItem12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent
func callback_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent", C.callback_ZN13QGraphicsItem13focusOutEventEP11QFocusEvent /*nil*/)
}

// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent
func callback_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.hoverEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", C.callback_ZN13QGraphicsItem15hoverEnterEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent
func callback_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.hoverMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN13QGraphicsItem14hoverMoveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent
func callback_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.hoverLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN13QGraphicsItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent
func callback_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent", C.callback_ZN13QGraphicsItem13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent
func callback_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent", C.callback_ZN13QGraphicsItem15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void mousePressEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent
func callback_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent", C.callback_ZN13QGraphicsItem15mousePressEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent
func callback_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", C.callback_ZN13QGraphicsItem14mouseMoveEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent
func callback_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", C.callback_ZN13QGraphicsItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent
func callback_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", C.callback_ZN13QGraphicsItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void wheelEvent(class QGraphicsSceneWheelEvent *)
//export callback_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent
func callback_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent", C.callback_ZN13QGraphicsItem10wheelEventEP24QGraphicsSceneWheelEvent /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent
func callback_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsItem.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent", C.callback_ZN13QGraphicsItem16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
//export callback_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE
func callback_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE(cthis unsafe.Pointer, query C.int) {
	// log.Println(cthis, "QGraphicsItem.inputMethodQuery")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodQuery", query)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE", C.callback_ZNK13QGraphicsItem16inputMethodQueryEN2Qt16InputMethodQueryE /*nil*/)
}

// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
//export callback_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant
func callback_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant(cthis unsafe.Pointer, change C.int, value unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsItem.itemChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemChange", change, value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant", C.callback_ZN13QGraphicsItem10itemChangeENS_18GraphicsItemChangeERK8QVariant /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE
func callback_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE", C.callback_ZNK13QGraphicsItem17supportsExtensionENS_9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant
func callback_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant", C.callback_ZN13QGraphicsItem12setExtensionENS_9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK13QGraphicsItem9extensionERK8QVariant
func callback_ZNK13QGraphicsItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QGraphicsItem9extensionERK8QVariant", C.callback_ZNK13QGraphicsItem9extensionERK8QVariant /*nil*/)
}

// void addToIndex()
//export callback_ZN13QGraphicsItem10addToIndexEv
func callback_ZN13QGraphicsItem10addToIndexEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsItem.addToIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "addToIndex")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem10addToIndexEv", C.callback_ZN13QGraphicsItem10addToIndexEv /*nil*/)
}

// void removeFromIndex()
//export callback_ZN13QGraphicsItem15removeFromIndexEv
func callback_ZN13QGraphicsItem15removeFromIndexEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsItem.removeFromIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "removeFromIndex")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem15removeFromIndexEv", C.callback_ZN13QGraphicsItem15removeFromIndexEv /*nil*/)
}

// void prepareGeometryChange()
//export callback_ZN13QGraphicsItem21prepareGeometryChangeEv
func callback_ZN13QGraphicsItem21prepareGeometryChangeEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsItem.prepareGeometryChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "prepareGeometryChange")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsItem21prepareGeometryChangeEv", C.callback_ZN13QGraphicsItem21prepareGeometryChangeEv /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK17QGraphicsPathItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK17QGraphicsPathItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsPathItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsPathItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK17QGraphicsPathItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN17QGraphicsPathItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN17QGraphicsPathItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsPathItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsPathItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN17QGraphicsPathItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK17QGraphicsPathItem9extensionERK8QVariant
func callback_ZNK17QGraphicsPathItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsPathItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsPathItem9extensionERK8QVariant", C.callback_ZNK17QGraphicsPathItem9extensionERK8QVariant /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsRectItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsRectItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK17QGraphicsRectItem9extensionERK8QVariant
func callback_ZNK17QGraphicsRectItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsRectItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsRectItem9extensionERK8QVariant", C.callback_ZNK17QGraphicsRectItem9extensionERK8QVariant /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK20QGraphicsEllipseItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK20QGraphicsEllipseItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsEllipseItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QGraphicsEllipseItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK20QGraphicsEllipseItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN20QGraphicsEllipseItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN20QGraphicsEllipseItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsEllipseItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsEllipseItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN20QGraphicsEllipseItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK20QGraphicsEllipseItem9extensionERK8QVariant
func callback_ZNK20QGraphicsEllipseItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsEllipseItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QGraphicsEllipseItem9extensionERK8QVariant", C.callback_ZNK20QGraphicsEllipseItem9extensionERK8QVariant /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK20QGraphicsPolygonItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK20QGraphicsPolygonItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsPolygonItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QGraphicsPolygonItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK20QGraphicsPolygonItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN20QGraphicsPolygonItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN20QGraphicsPolygonItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsPolygonItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsPolygonItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN20QGraphicsPolygonItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK20QGraphicsPolygonItem9extensionERK8QVariant
func callback_ZNK20QGraphicsPolygonItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsPolygonItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QGraphicsPolygonItem9extensionERK8QVariant", C.callback_ZNK20QGraphicsPolygonItem9extensionERK8QVariant /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK17QGraphicsLineItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK17QGraphicsLineItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsLineItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsLineItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK17QGraphicsLineItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN17QGraphicsLineItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN17QGraphicsLineItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsLineItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsLineItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN17QGraphicsLineItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK17QGraphicsLineItem9extensionERK8QVariant
func callback_ZNK17QGraphicsLineItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsLineItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsLineItem9extensionERK8QVariant", C.callback_ZNK17QGraphicsLineItem9extensionERK8QVariant /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK19QGraphicsPixmapItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK19QGraphicsPixmapItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsPixmapItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK19QGraphicsPixmapItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK19QGraphicsPixmapItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN19QGraphicsPixmapItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN19QGraphicsPixmapItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsPixmapItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QGraphicsPixmapItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN19QGraphicsPixmapItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK19QGraphicsPixmapItem9extensionERK8QVariant
func callback_ZNK19QGraphicsPixmapItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsPixmapItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK19QGraphicsPixmapItem9extensionERK8QVariant", C.callback_ZNK19QGraphicsPixmapItem9extensionERK8QVariant /*nil*/)
}

// bool sceneEvent(class QEvent *)
//export callback_ZN17QGraphicsTextItem10sceneEventEP6QEvent
func callback_ZN17QGraphicsTextItem10sceneEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.sceneEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "sceneEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem10sceneEventEP6QEvent", C.callback_ZN17QGraphicsTextItem10sceneEventEP6QEvent /*nil*/)
}

// void mousePressEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent
func callback_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent", C.callback_ZN17QGraphicsTextItem15mousePressEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent
func callback_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent", C.callback_ZN17QGraphicsTextItem14mouseMoveEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent
func callback_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent", C.callback_ZN17QGraphicsTextItem17mouseReleaseEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent
func callback_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", C.callback_ZN17QGraphicsTextItem21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
//export callback_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent
func callback_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent", C.callback_ZN17QGraphicsTextItem16contextMenuEventEP30QGraphicsSceneContextMenuEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent
func callback_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent", C.callback_ZN17QGraphicsTextItem13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent
func callback_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent", C.callback_ZN17QGraphicsTextItem15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent
func callback_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent", C.callback_ZN17QGraphicsTextItem12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent
func callback_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent", C.callback_ZN17QGraphicsTextItem13focusOutEventEP11QFocusEvent /*nil*/)
}

// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent
func callback_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN17QGraphicsTextItem14dragEnterEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN17QGraphicsTextItem14dragLeaveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN17QGraphicsTextItem13dragMoveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dropEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent
func callback_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN17QGraphicsTextItem9dropEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent
func callback_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent", C.callback_ZN17QGraphicsTextItem16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent
func callback_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.hoverEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent", C.callback_ZN17QGraphicsTextItem15hoverEnterEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent
func callback_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.hoverMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN17QGraphicsTextItem14hoverMoveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent
func callback_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsTextItem.hoverLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN17QGraphicsTextItem15hoverLeaveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
//export callback_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE
func callback_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE(cthis unsafe.Pointer, query C.int) {
	// log.Println(cthis, "QGraphicsTextItem.inputMethodQuery")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodQuery", query)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE", C.callback_ZNK17QGraphicsTextItem16inputMethodQueryEN2Qt16InputMethodQueryE /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsTextItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK17QGraphicsTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsTextItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN17QGraphicsTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK17QGraphicsTextItem9extensionERK8QVariant
func callback_ZNK17QGraphicsTextItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsTextItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QGraphicsTextItem9extensionERK8QVariant", C.callback_ZNK17QGraphicsTextItem9extensionERK8QVariant /*nil*/)
}

// bool supportsExtension(enum QGraphicsItem::Extension)
//export callback_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE
func callback_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE(cthis unsafe.Pointer, extension C.int) {
	// log.Println(cthis, "QGraphicsSimpleTextItem.supportsExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportsExtension", extension)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", C.callback_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE /*nil*/)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
//export callback_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant
func callback_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant(cthis unsafe.Pointer, extension C.int, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsSimpleTextItem.setExtension")
	rvx := ffiqt.CallbackAllInherits(cthis, "setExtension", extension, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", C.callback_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant /*nil*/)
}

// QVariant extension(const class QVariant &)
//export callback_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant
func callback_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant(cthis unsafe.Pointer, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsSimpleTextItem.extension")
	rvx := ffiqt.CallbackAllInherits(cthis, "extension", variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant", C.callback_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant /*nil*/)
}

// void setGraphicsItem(class QGraphicsItem *)
//export callback_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem
func callback_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem(cthis unsafe.Pointer, item unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsLayoutItem.setGraphicsItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "setGraphicsItem", item)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem", C.callback_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem /*nil*/)
}

// void setOwnedByLayout(_Bool)
//export callback_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb
func callback_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb(cthis unsafe.Pointer, ownedByLayout C.bool) {
	// log.Println(cthis, "QGraphicsLayoutItem.setOwnedByLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "setOwnedByLayout", ownedByLayout)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb", C.callback_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb /*nil*/)
}

// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
//export callback_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF
func callback_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF(cthis unsafe.Pointer, which C.int, constraint unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsLayoutItem.sizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHint", which, constraint)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF", C.callback_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF /*nil*/)
}

// void addChildLayoutItem(class QGraphicsLayoutItem *)
//export callback_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem
func callback_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem(cthis unsafe.Pointer, layoutItem unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsLayout.addChildLayoutItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "addChildLayoutItem", layoutItem)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem", C.callback_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem /*nil*/)
}

// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
//export callback_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF
func callback_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF(cthis unsafe.Pointer, which C.int, constraint unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsAnchorLayout.sizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHint", which, constraint)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", C.callback_ZNK21QGraphicsAnchorLayout8sizeHintEN2Qt8SizeHintERK6QSizeF /*nil*/)
}

// void draw(class QPainter *)
//export callback_ZN15QGraphicsEffect4drawEP8QPainter
func callback_ZN15QGraphicsEffect4drawEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsEffect.draw")
	rvx := ffiqt.CallbackAllInherits(cthis, "draw", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsEffect4drawEP8QPainter", C.callback_ZN15QGraphicsEffect4drawEP8QPainter /*nil*/)
}

// void sourceChanged(QGraphicsEffect::ChangeFlags)
//export callback_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE
func callback_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE(cthis unsafe.Pointer, flags C.int) {
	// log.Println(cthis, "QGraphicsEffect.sourceChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "sourceChanged", flags)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE", C.callback_ZN15QGraphicsEffect13sourceChangedE6QFlagsINS_10ChangeFlagEE /*nil*/)
}

// void updateBoundingRect()
//export callback_ZN15QGraphicsEffect18updateBoundingRectEv
func callback_ZN15QGraphicsEffect18updateBoundingRectEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsEffect.updateBoundingRect")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateBoundingRect")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsEffect18updateBoundingRectEv", C.callback_ZN15QGraphicsEffect18updateBoundingRectEv /*nil*/)
}

// bool sourceIsPixmap()
//export callback_ZNK15QGraphicsEffect14sourceIsPixmapEv
func callback_ZNK15QGraphicsEffect14sourceIsPixmapEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsEffect.sourceIsPixmap")
	rvx := ffiqt.CallbackAllInherits(cthis, "sourceIsPixmap")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QGraphicsEffect14sourceIsPixmapEv", C.callback_ZNK15QGraphicsEffect14sourceIsPixmapEv /*nil*/)
}

// QRectF sourceBoundingRect(Qt::CoordinateSystem)
//export callback_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE
func callback_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE(cthis unsafe.Pointer, system C.int) {
	// log.Println(cthis, "QGraphicsEffect.sourceBoundingRect")
	rvx := ffiqt.CallbackAllInherits(cthis, "sourceBoundingRect", system)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE", C.callback_ZNK15QGraphicsEffect18sourceBoundingRectEN2Qt16CoordinateSystemE /*nil*/)
}

// void drawSource(class QPainter *)
//export callback_ZN15QGraphicsEffect10drawSourceEP8QPainter
func callback_ZN15QGraphicsEffect10drawSourceEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsEffect.drawSource")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawSource", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsEffect10drawSourceEP8QPainter", C.callback_ZN15QGraphicsEffect10drawSourceEP8QPainter /*nil*/)
}

// QPixmap sourcePixmap(Qt::CoordinateSystem, class QPoint *, enum QGraphicsEffect::PixmapPadMode)
//export callback_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE
func callback_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE(cthis unsafe.Pointer, system C.int, offset unsafe.Pointer /*666*/, mode C.int) {
	// log.Println(cthis, "QGraphicsEffect.sourcePixmap")
	rvx := ffiqt.CallbackAllInherits(cthis, "sourcePixmap", system, offset, mode)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE", C.callback_ZNK15QGraphicsEffect12sourcePixmapEN2Qt16CoordinateSystemEP6QPointNS_13PixmapPadModeE /*nil*/)
}

// void draw(class QPainter *)
//export callback_ZN23QGraphicsColorizeEffect4drawEP8QPainter
func callback_ZN23QGraphicsColorizeEffect4drawEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsColorizeEffect.draw")
	rvx := ffiqt.CallbackAllInherits(cthis, "draw", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN23QGraphicsColorizeEffect4drawEP8QPainter", C.callback_ZN23QGraphicsColorizeEffect4drawEP8QPainter /*nil*/)
}

// void draw(class QPainter *)
//export callback_ZN19QGraphicsBlurEffect4drawEP8QPainter
func callback_ZN19QGraphicsBlurEffect4drawEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsBlurEffect.draw")
	rvx := ffiqt.CallbackAllInherits(cthis, "draw", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QGraphicsBlurEffect4drawEP8QPainter", C.callback_ZN19QGraphicsBlurEffect4drawEP8QPainter /*nil*/)
}

// void draw(class QPainter *)
//export callback_ZN25QGraphicsDropShadowEffect4drawEP8QPainter
func callback_ZN25QGraphicsDropShadowEffect4drawEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsDropShadowEffect.draw")
	rvx := ffiqt.CallbackAllInherits(cthis, "draw", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN25QGraphicsDropShadowEffect4drawEP8QPainter", C.callback_ZN25QGraphicsDropShadowEffect4drawEP8QPainter /*nil*/)
}

// void draw(class QPainter *)
//export callback_ZN22QGraphicsOpacityEffect4drawEP8QPainter
func callback_ZN22QGraphicsOpacityEffect4drawEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsOpacityEffect.draw")
	rvx := ffiqt.CallbackAllInherits(cthis, "draw", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN22QGraphicsOpacityEffect4drawEP8QPainter", C.callback_ZN22QGraphicsOpacityEffect4drawEP8QPainter /*nil*/)
}

// void beforeAnimationStep(qreal)
//export callback_ZN22QGraphicsItemAnimation19beforeAnimationStepEd
func callback_ZN22QGraphicsItemAnimation19beforeAnimationStepEd(cthis unsafe.Pointer, step C.double) {
	// log.Println(cthis, "QGraphicsItemAnimation.beforeAnimationStep")
	rvx := ffiqt.CallbackAllInherits(cthis, "beforeAnimationStep", step)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN22QGraphicsItemAnimation19beforeAnimationStepEd", C.callback_ZN22QGraphicsItemAnimation19beforeAnimationStepEd /*nil*/)
}

// void afterAnimationStep(qreal)
//export callback_ZN22QGraphicsItemAnimation18afterAnimationStepEd
func callback_ZN22QGraphicsItemAnimation18afterAnimationStepEd(cthis unsafe.Pointer, step C.double) {
	// log.Println(cthis, "QGraphicsItemAnimation.afterAnimationStep")
	rvx := ffiqt.CallbackAllInherits(cthis, "afterAnimationStep", step)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN22QGraphicsItemAnimation18afterAnimationStepEd", C.callback_ZN22QGraphicsItemAnimation18afterAnimationStepEd /*nil*/)
}

// void initStyleOption(class QStyleOption *)
//export callback_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption
func callback_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption", C.callback_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption /*nil*/)
}

// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
//export callback_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF
func callback_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF(cthis unsafe.Pointer, which C.int, constraint unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsWidget.sizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHint", which, constraint)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", C.callback_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF /*nil*/)
}

// void updateGeometry()
//export callback_ZN15QGraphicsWidget14updateGeometryEv
func callback_ZN15QGraphicsWidget14updateGeometryEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsWidget.updateGeometry")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateGeometry")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget14updateGeometryEv", C.callback_ZN15QGraphicsWidget14updateGeometryEv /*nil*/)
}

// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
//export callback_ZN15QGraphicsWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant
func callback_ZN15QGraphicsWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant(cthis unsafe.Pointer, change C.int, value unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsWidget.itemChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemChange", change, value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant", C.callback_ZN15QGraphicsWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant /*nil*/)
}

// QVariant propertyChange(const class QString &, const class QVariant &)
//export callback_ZN15QGraphicsWidget14propertyChangeERK7QStringRK8QVariant
func callback_ZN15QGraphicsWidget14propertyChangeERK7QStringRK8QVariant(cthis unsafe.Pointer, propertyName unsafe.Pointer /*555*/, value unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsWidget.propertyChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "propertyChange", propertyName, value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget14propertyChangeERK7QStringRK8QVariant", C.callback_ZN15QGraphicsWidget14propertyChangeERK7QStringRK8QVariant /*nil*/)
}

// bool sceneEvent(class QEvent *)
//export callback_ZN15QGraphicsWidget10sceneEventEP6QEvent
func callback_ZN15QGraphicsWidget10sceneEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.sceneEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "sceneEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget10sceneEventEP6QEvent", C.callback_ZN15QGraphicsWidget10sceneEventEP6QEvent /*nil*/)
}

// bool windowFrameEvent(class QEvent *)
//export callback_ZN15QGraphicsWidget16windowFrameEventEP6QEvent
func callback_ZN15QGraphicsWidget16windowFrameEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.windowFrameEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "windowFrameEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget16windowFrameEventEP6QEvent", C.callback_ZN15QGraphicsWidget16windowFrameEventEP6QEvent /*nil*/)
}

// Qt::WindowFrameSection windowFrameSectionAt(const class QPointF &)
//export callback_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF
func callback_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF(cthis unsafe.Pointer, pos unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsWidget.windowFrameSectionAt")
	rvx := ffiqt.CallbackAllInherits(cthis, "windowFrameSectionAt", pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF", C.callback_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QGraphicsWidget5eventEP6QEvent
func callback_ZN15QGraphicsWidget5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget5eventEP6QEvent", C.callback_ZN15QGraphicsWidget5eventEP6QEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN15QGraphicsWidget11changeEventEP6QEvent
func callback_ZN15QGraphicsWidget11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget11changeEventEP6QEvent", C.callback_ZN15QGraphicsWidget11changeEventEP6QEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN15QGraphicsWidget10closeEventEP11QCloseEvent
func callback_ZN15QGraphicsWidget10closeEventEP11QCloseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget10closeEventEP11QCloseEvent", C.callback_ZN15QGraphicsWidget10closeEventEP11QCloseEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent
func callback_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent", C.callback_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN15QGraphicsWidget18focusNextPrevChildEb
func callback_ZN15QGraphicsWidget18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QGraphicsWidget.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget18focusNextPrevChildEb", C.callback_ZN15QGraphicsWidget18focusNextPrevChildEb /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent
func callback_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent", C.callback_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN15QGraphicsWidget9hideEventEP10QHideEvent
func callback_ZN15QGraphicsWidget9hideEventEP10QHideEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget9hideEventEP10QHideEvent", C.callback_ZN15QGraphicsWidget9hideEventEP10QHideEvent /*nil*/)
}

// void moveEvent(class QGraphicsSceneMoveEvent *)
//export callback_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent
func callback_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.moveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent", C.callback_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent /*nil*/)
}

// void polishEvent()
//export callback_ZN15QGraphicsWidget11polishEventEv
func callback_ZN15QGraphicsWidget11polishEventEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsWidget.polishEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "polishEvent")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget11polishEventEv", C.callback_ZN15QGraphicsWidget11polishEventEv /*nil*/)
}

// void resizeEvent(class QGraphicsSceneResizeEvent *)
//export callback_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent
func callback_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent", C.callback_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN15QGraphicsWidget9showEventEP10QShowEvent
func callback_ZN15QGraphicsWidget9showEventEP10QShowEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget9showEventEP10QShowEvent", C.callback_ZN15QGraphicsWidget9showEventEP10QShowEvent /*nil*/)
}

// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent
func callback_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.hoverMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent
func callback_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.hoverLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void grabMouseEvent(class QEvent *)
//export callback_ZN15QGraphicsWidget14grabMouseEventEP6QEvent
func callback_ZN15QGraphicsWidget14grabMouseEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.grabMouseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "grabMouseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget14grabMouseEventEP6QEvent", C.callback_ZN15QGraphicsWidget14grabMouseEventEP6QEvent /*nil*/)
}

// void ungrabMouseEvent(class QEvent *)
//export callback_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent
func callback_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.ungrabMouseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "ungrabMouseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent", C.callback_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent /*nil*/)
}

// void grabKeyboardEvent(class QEvent *)
//export callback_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent
func callback_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.grabKeyboardEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "grabKeyboardEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent", C.callback_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent /*nil*/)
}

// void ungrabKeyboardEvent(class QEvent *)
//export callback_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent
func callback_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsWidget.ungrabKeyboardEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "ungrabKeyboardEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent", C.callback_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent /*nil*/)
}

// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
//export callback_ZN20QGraphicsProxyWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant
func callback_ZN20QGraphicsProxyWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant(cthis unsafe.Pointer, change C.int, value unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.itemChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemChange", change, value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant", C.callback_ZN20QGraphicsProxyWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN20QGraphicsProxyWidget5eventEP6QEvent
func callback_ZN20QGraphicsProxyWidget5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget5eventEP6QEvent", C.callback_ZN20QGraphicsProxyWidget5eventEP6QEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN20QGraphicsProxyWidget11eventFilterEP7QObjectP6QEvent
func callback_ZN20QGraphicsProxyWidget11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, object unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", object, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget11eventFilterEP7QObjectP6QEvent", C.callback_ZN20QGraphicsProxyWidget11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN20QGraphicsProxyWidget9showEventEP10QShowEvent
func callback_ZN20QGraphicsProxyWidget9showEventEP10QShowEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget9showEventEP10QShowEvent", C.callback_ZN20QGraphicsProxyWidget9showEventEP10QShowEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN20QGraphicsProxyWidget9hideEventEP10QHideEvent
func callback_ZN20QGraphicsProxyWidget9hideEventEP10QHideEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget9hideEventEP10QHideEvent", C.callback_ZN20QGraphicsProxyWidget9hideEventEP10QHideEvent /*nil*/)
}

// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
//export callback_ZN20QGraphicsProxyWidget16contextMenuEventEP30QGraphicsSceneContextMenuEvent
func callback_ZN20QGraphicsProxyWidget16contextMenuEventEP30QGraphicsSceneContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget16contextMenuEventEP30QGraphicsSceneContextMenuEvent", C.callback_ZN20QGraphicsProxyWidget16contextMenuEventEP30QGraphicsSceneContextMenuEvent /*nil*/)
}

// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent
func callback_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dropEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent
func callback_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN20QGraphicsProxyWidget15hoverEnterEventEP24QGraphicsSceneHoverEvent
func callback_ZN20QGraphicsProxyWidget15hoverEnterEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.hoverEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget15hoverEnterEventEP24QGraphicsSceneHoverEvent", C.callback_ZN20QGraphicsProxyWidget15hoverEnterEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN20QGraphicsProxyWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent
func callback_ZN20QGraphicsProxyWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.hoverLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN20QGraphicsProxyWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
//export callback_ZN20QGraphicsProxyWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent
func callback_ZN20QGraphicsProxyWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.hoverMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hoverMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent", C.callback_ZN20QGraphicsProxyWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent /*nil*/)
}

// void grabMouseEvent(class QEvent *)
//export callback_ZN20QGraphicsProxyWidget14grabMouseEventEP6QEvent
func callback_ZN20QGraphicsProxyWidget14grabMouseEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.grabMouseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "grabMouseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget14grabMouseEventEP6QEvent", C.callback_ZN20QGraphicsProxyWidget14grabMouseEventEP6QEvent /*nil*/)
}

// void ungrabMouseEvent(class QEvent *)
//export callback_ZN20QGraphicsProxyWidget16ungrabMouseEventEP6QEvent
func callback_ZN20QGraphicsProxyWidget16ungrabMouseEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.ungrabMouseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "ungrabMouseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget16ungrabMouseEventEP6QEvent", C.callback_ZN20QGraphicsProxyWidget16ungrabMouseEventEP6QEvent /*nil*/)
}

// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN20QGraphicsProxyWidget14mouseMoveEventEP24QGraphicsSceneMouseEvent
func callback_ZN20QGraphicsProxyWidget14mouseMoveEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget14mouseMoveEventEP24QGraphicsSceneMouseEvent", C.callback_ZN20QGraphicsProxyWidget14mouseMoveEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mousePressEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN20QGraphicsProxyWidget15mousePressEventEP24QGraphicsSceneMouseEvent
func callback_ZN20QGraphicsProxyWidget15mousePressEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget15mousePressEventEP24QGraphicsSceneMouseEvent", C.callback_ZN20QGraphicsProxyWidget15mousePressEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN20QGraphicsProxyWidget17mouseReleaseEventEP24QGraphicsSceneMouseEvent
func callback_ZN20QGraphicsProxyWidget17mouseReleaseEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget17mouseReleaseEventEP24QGraphicsSceneMouseEvent", C.callback_ZN20QGraphicsProxyWidget17mouseReleaseEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN20QGraphicsProxyWidget21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent
func callback_ZN20QGraphicsProxyWidget21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", C.callback_ZN20QGraphicsProxyWidget21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void wheelEvent(class QGraphicsSceneWheelEvent *)
//export callback_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent
func callback_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent", C.callback_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN20QGraphicsProxyWidget13keyPressEventEP9QKeyEvent
func callback_ZN20QGraphicsProxyWidget13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget13keyPressEventEP9QKeyEvent", C.callback_ZN20QGraphicsProxyWidget13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN20QGraphicsProxyWidget15keyReleaseEventEP9QKeyEvent
func callback_ZN20QGraphicsProxyWidget15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget15keyReleaseEventEP9QKeyEvent", C.callback_ZN20QGraphicsProxyWidget15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN20QGraphicsProxyWidget12focusInEventEP11QFocusEvent
func callback_ZN20QGraphicsProxyWidget12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget12focusInEventEP11QFocusEvent", C.callback_ZN20QGraphicsProxyWidget12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN20QGraphicsProxyWidget13focusOutEventEP11QFocusEvent
func callback_ZN20QGraphicsProxyWidget13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget13focusOutEventEP11QFocusEvent", C.callback_ZN20QGraphicsProxyWidget13focusOutEventEP11QFocusEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN20QGraphicsProxyWidget18focusNextPrevChildEb
func callback_ZN20QGraphicsProxyWidget18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QGraphicsProxyWidget.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget18focusNextPrevChildEb", C.callback_ZN20QGraphicsProxyWidget18focusNextPrevChildEb /*nil*/)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
//export callback_ZNK20QGraphicsProxyWidget16inputMethodQueryEN2Qt16InputMethodQueryE
func callback_ZNK20QGraphicsProxyWidget16inputMethodQueryEN2Qt16InputMethodQueryE(cthis unsafe.Pointer, query C.int) {
	// log.Println(cthis, "QGraphicsProxyWidget.inputMethodQuery")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodQuery", query)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QGraphicsProxyWidget16inputMethodQueryEN2Qt16InputMethodQueryE", C.callback_ZNK20QGraphicsProxyWidget16inputMethodQueryEN2Qt16InputMethodQueryE /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN20QGraphicsProxyWidget16inputMethodEventEP17QInputMethodEvent
func callback_ZN20QGraphicsProxyWidget16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget16inputMethodEventEP17QInputMethodEvent", C.callback_ZN20QGraphicsProxyWidget16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
//export callback_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF
func callback_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF(cthis unsafe.Pointer, which C.int, constraint unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.sizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHint", which, constraint)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", C.callback_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF /*nil*/)
}

// void resizeEvent(class QGraphicsSceneResizeEvent *)
//export callback_ZN20QGraphicsProxyWidget11resizeEventEP25QGraphicsSceneResizeEvent
func callback_ZN20QGraphicsProxyWidget11resizeEventEP25QGraphicsSceneResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget11resizeEventEP25QGraphicsSceneResizeEvent", C.callback_ZN20QGraphicsProxyWidget11resizeEventEP25QGraphicsSceneResizeEvent /*nil*/)
}

// QGraphicsProxyWidget * newProxyWidget(const class QWidget *)
//export callback_ZN20QGraphicsProxyWidget14newProxyWidgetEPK7QWidget
func callback_ZN20QGraphicsProxyWidget14newProxyWidgetEPK7QWidget(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsProxyWidget.newProxyWidget")
	rvx := ffiqt.CallbackAllInherits(cthis, "newProxyWidget", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN20QGraphicsProxyWidget14newProxyWidgetEPK7QWidget", C.callback_ZN20QGraphicsProxyWidget14newProxyWidgetEPK7QWidget /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN14QGraphicsScene5eventEP6QEvent
func callback_ZN14QGraphicsScene5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene5eventEP6QEvent", C.callback_ZN14QGraphicsScene5eventEP6QEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent
func callback_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, watched unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", watched, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent", C.callback_ZN14QGraphicsScene11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
//export callback_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent
func callback_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent", C.callback_ZN14QGraphicsScene16contextMenuEventEP30QGraphicsSceneContextMenuEvent /*nil*/)
}

// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent
func callback_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN14QGraphicsScene14dragEnterEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN14QGraphicsScene13dragMoveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent
func callback_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN14QGraphicsScene14dragLeaveEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void dropEvent(class QGraphicsSceneDragDropEvent *)
//export callback_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent
func callback_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent", C.callback_ZN14QGraphicsScene9dropEventEP27QGraphicsSceneDragDropEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN14QGraphicsScene12focusInEventEP11QFocusEvent
func callback_ZN14QGraphicsScene12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene12focusInEventEP11QFocusEvent", C.callback_ZN14QGraphicsScene12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent
func callback_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent", C.callback_ZN14QGraphicsScene13focusOutEventEP11QFocusEvent /*nil*/)
}

// void helpEvent(class QGraphicsSceneHelpEvent *)
//export callback_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent
func callback_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.helpEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "helpEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent", C.callback_ZN14QGraphicsScene9helpEventEP23QGraphicsSceneHelpEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent
func callback_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent", C.callback_ZN14QGraphicsScene13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent
func callback_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent", C.callback_ZN14QGraphicsScene15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void mousePressEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent
func callback_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent", C.callback_ZN14QGraphicsScene15mousePressEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent
func callback_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent", C.callback_ZN14QGraphicsScene14mouseMoveEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent
func callback_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent", C.callback_ZN14QGraphicsScene17mouseReleaseEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
//export callback_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent
func callback_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", C.callback_ZN14QGraphicsScene21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent /*nil*/)
}

// void wheelEvent(class QGraphicsSceneWheelEvent *)
//export callback_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent
func callback_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent", C.callback_ZN14QGraphicsScene10wheelEventEP24QGraphicsSceneWheelEvent /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent
func callback_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent", C.callback_ZN14QGraphicsScene16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void drawBackground(class QPainter *, const class QRectF &)
//export callback_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF
func callback_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsScene.drawBackground")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawBackground", painter, rect)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF", C.callback_ZN14QGraphicsScene14drawBackgroundEP8QPainterRK6QRectF /*nil*/)
}

// void drawForeground(class QPainter *, const class QRectF &)
//export callback_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF
func callback_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsScene.drawForeground")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawForeground", painter, rect)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF", C.callback_ZN14QGraphicsScene14drawForegroundEP8QPainterRK6QRectF /*nil*/)
}

// void drawItems(class QPainter *, int, class QGraphicsItem **, const class QStyleOptionGraphicsItem *, class QWidget *)
//export callback_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget
func callback_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, numItems C.int, items unsafe.Pointer /*222*/, options unsafe.Pointer /*222*/, widget unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsScene.drawItems")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawItems", painter, numItems, items, options, widget)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget", C.callback_ZN14QGraphicsScene9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItemP7QWidget /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN14QGraphicsScene18focusNextPrevChildEb
func callback_ZN14QGraphicsScene18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QGraphicsScene.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QGraphicsScene18focusNextPrevChildEb", C.callback_ZN14QGraphicsScene18focusNextPrevChildEb /*nil*/)
}

// void update()
//export callback_ZN18QGraphicsTransform6updateEv
func callback_ZN18QGraphicsTransform6updateEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QGraphicsTransform.update")
	rvx := ffiqt.CallbackAllInherits(cthis, "update")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QGraphicsTransform6updateEv", C.callback_ZN18QGraphicsTransform6updateEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QScrollArea5eventEP6QEvent
func callback_ZN11QScrollArea5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollArea.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QScrollArea5eventEP6QEvent", C.callback_ZN11QScrollArea5eventEP6QEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent
func callback_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollArea.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent", C.callback_ZN11QScrollArea11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN11QScrollArea11resizeEventEP12QResizeEvent
func callback_ZN11QScrollArea11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollArea.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QScrollArea11resizeEventEP12QResizeEvent", C.callback_ZN11QScrollArea11resizeEventEP12QResizeEvent /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN11QScrollArea16scrollContentsByEii
func callback_ZN11QScrollArea16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QScrollArea.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QScrollArea16scrollContentsByEii", C.callback_ZN11QScrollArea16scrollContentsByEii /*nil*/)
}

// QSize viewportSizeHint()
//export callback_ZNK11QScrollArea16viewportSizeHintEv
func callback_ZNK11QScrollArea16viewportSizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QScrollArea.viewportSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportSizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QScrollArea16viewportSizeHintEv", C.callback_ZNK11QScrollArea16viewportSizeHintEv /*nil*/)
}

// void setupViewport(class QWidget *)
//export callback_ZN13QGraphicsView13setupViewportEP7QWidget
func callback_ZN13QGraphicsView13setupViewportEP7QWidget(cthis unsafe.Pointer, widget unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.setupViewport")
	rvx := ffiqt.CallbackAllInherits(cthis, "setupViewport", widget)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView13setupViewportEP7QWidget", C.callback_ZN13QGraphicsView13setupViewportEP7QWidget /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN13QGraphicsView5eventEP6QEvent
func callback_ZN13QGraphicsView5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView5eventEP6QEvent", C.callback_ZN13QGraphicsView5eventEP6QEvent /*nil*/)
}

// bool viewportEvent(class QEvent *)
//export callback_ZN13QGraphicsView13viewportEventEP6QEvent
func callback_ZN13QGraphicsView13viewportEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.viewportEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView13viewportEventEP6QEvent", C.callback_ZN13QGraphicsView13viewportEventEP6QEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent
func callback_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent", C.callback_ZN13QGraphicsView16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void dragEnterEvent(class QDragEnterEvent *)
//export callback_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent
func callback_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent", C.callback_ZN13QGraphicsView14dragEnterEventEP15QDragEnterEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN13QGraphicsView14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent
func callback_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent", C.callback_ZN13QGraphicsView13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN13QGraphicsView9dropEventEP10QDropEvent
func callback_ZN13QGraphicsView9dropEventEP10QDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView9dropEventEP10QDropEvent", C.callback_ZN13QGraphicsView9dropEventEP10QDropEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN13QGraphicsView12focusInEventEP11QFocusEvent
func callback_ZN13QGraphicsView12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView12focusInEventEP11QFocusEvent", C.callback_ZN13QGraphicsView12focusInEventEP11QFocusEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN13QGraphicsView18focusNextPrevChildEb
func callback_ZN13QGraphicsView18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QGraphicsView.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView18focusNextPrevChildEb", C.callback_ZN13QGraphicsView18focusNextPrevChildEb /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN13QGraphicsView13focusOutEventEP11QFocusEvent
func callback_ZN13QGraphicsView13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView13focusOutEventEP11QFocusEvent", C.callback_ZN13QGraphicsView13focusOutEventEP11QFocusEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN13QGraphicsView13keyPressEventEP9QKeyEvent
func callback_ZN13QGraphicsView13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView13keyPressEventEP9QKeyEvent", C.callback_ZN13QGraphicsView13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent
func callback_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent", C.callback_ZN13QGraphicsView15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN13QGraphicsView21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN13QGraphicsView15mousePressEventEP11QMouseEvent
func callback_ZN13QGraphicsView15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView15mousePressEventEP11QMouseEvent", C.callback_ZN13QGraphicsView15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent
func callback_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent", C.callback_ZN13QGraphicsView14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent
func callback_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent", C.callback_ZN13QGraphicsView17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN13QGraphicsView10wheelEventEP11QWheelEvent
func callback_ZN13QGraphicsView10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView10wheelEventEP11QWheelEvent", C.callback_ZN13QGraphicsView10wheelEventEP11QWheelEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN13QGraphicsView10paintEventEP11QPaintEvent
func callback_ZN13QGraphicsView10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView10paintEventEP11QPaintEvent", C.callback_ZN13QGraphicsView10paintEventEP11QPaintEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN13QGraphicsView11resizeEventEP12QResizeEvent
func callback_ZN13QGraphicsView11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView11resizeEventEP12QResizeEvent", C.callback_ZN13QGraphicsView11resizeEventEP12QResizeEvent /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN13QGraphicsView16scrollContentsByEii
func callback_ZN13QGraphicsView16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QGraphicsView.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView16scrollContentsByEii", C.callback_ZN13QGraphicsView16scrollContentsByEii /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN13QGraphicsView9showEventEP10QShowEvent
func callback_ZN13QGraphicsView9showEventEP10QShowEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView9showEventEP10QShowEvent", C.callback_ZN13QGraphicsView9showEventEP10QShowEvent /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent
func callback_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGraphicsView.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent", C.callback_ZN13QGraphicsView16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void drawBackground(class QPainter *, const class QRectF &)
//export callback_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF
func callback_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsView.drawBackground")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawBackground", painter, rect)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF", C.callback_ZN13QGraphicsView14drawBackgroundEP8QPainterRK6QRectF /*nil*/)
}

// void drawForeground(class QPainter *, const class QRectF &)
//export callback_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF
func callback_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QGraphicsView.drawForeground")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawForeground", painter, rect)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF", C.callback_ZN13QGraphicsView14drawForegroundEP8QPainterRK6QRectF /*nil*/)
}

// void drawItems(class QPainter *, int, class QGraphicsItem **, const class QStyleOptionGraphicsItem *)
//export callback_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem
func callback_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, numItems C.int, items unsafe.Pointer /*222*/, options unsafe.Pointer /*222*/) {
	// log.Println(cthis, "QGraphicsView.drawItems")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawItems", painter, numItems, items, options)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem", C.callback_ZN13QGraphicsView9drawItemsEP8QPainteriPP13QGraphicsItemPK24QStyleOptionGraphicsItem /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QGroupBox5eventEP6QEvent
func callback_ZN9QGroupBox5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox5eventEP6QEvent", C.callback_ZN9QGroupBox5eventEP6QEvent /*nil*/)
}

// void childEvent(class QChildEvent *)
//export callback_ZN9QGroupBox10childEventEP11QChildEvent
func callback_ZN9QGroupBox10childEventEP11QChildEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.childEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "childEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox10childEventEP11QChildEvent", C.callback_ZN9QGroupBox10childEventEP11QChildEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN9QGroupBox11resizeEventEP12QResizeEvent
func callback_ZN9QGroupBox11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox11resizeEventEP12QResizeEvent", C.callback_ZN9QGroupBox11resizeEventEP12QResizeEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QGroupBox10paintEventEP11QPaintEvent
func callback_ZN9QGroupBox10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox10paintEventEP11QPaintEvent", C.callback_ZN9QGroupBox10paintEventEP11QPaintEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN9QGroupBox12focusInEventEP11QFocusEvent
func callback_ZN9QGroupBox12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox12focusInEventEP11QFocusEvent", C.callback_ZN9QGroupBox12focusInEventEP11QFocusEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN9QGroupBox11changeEventEP6QEvent
func callback_ZN9QGroupBox11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox11changeEventEP6QEvent", C.callback_ZN9QGroupBox11changeEventEP6QEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN9QGroupBox15mousePressEventEP11QMouseEvent
func callback_ZN9QGroupBox15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox15mousePressEventEP11QMouseEvent", C.callback_ZN9QGroupBox15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent
func callback_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent", C.callback_ZN9QGroupBox14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent
func callback_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent", C.callback_ZN9QGroupBox17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionGroupBox *)
//export callback_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox
func callback_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QGroupBox.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox", C.callback_ZNK9QGroupBox15initStyleOptionEP20QStyleOptionGroupBox /*nil*/)
}

// void updateSection(int)
//export callback_ZN11QHeaderView13updateSectionEi
func callback_ZN11QHeaderView13updateSectionEi(cthis unsafe.Pointer, logicalIndex C.int) {
	// log.Println(cthis, "QHeaderView.updateSection")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateSection", logicalIndex)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView13updateSectionEi", C.callback_ZN11QHeaderView13updateSectionEi /*nil*/)
}

// void resizeSections()
//export callback_ZN11QHeaderView14resizeSectionsEv
func callback_ZN11QHeaderView14resizeSectionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QHeaderView.resizeSections")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeSections")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView14resizeSectionsEv", C.callback_ZN11QHeaderView14resizeSectionsEv /*nil*/)
}

// void sectionsInserted(const class QModelIndex &, int, int)
//export callback_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii
func callback_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, logicalFirst C.int, logicalLast C.int) {
	// log.Println(cthis, "QHeaderView.sectionsInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "sectionsInserted", parent, logicalFirst, logicalLast)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii", C.callback_ZN11QHeaderView16sectionsInsertedERK11QModelIndexii /*nil*/)
}

// void sectionsAboutToBeRemoved(const class QModelIndex &, int, int)
//export callback_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii
func callback_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, logicalFirst C.int, logicalLast C.int) {
	// log.Println(cthis, "QHeaderView.sectionsAboutToBeRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "sectionsAboutToBeRemoved", parent, logicalFirst, logicalLast)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii", C.callback_ZN11QHeaderView24sectionsAboutToBeRemovedERK11QModelIndexii /*nil*/)
}

// void initialize()
//export callback_ZN11QHeaderView10initializeEv
func callback_ZN11QHeaderView10initializeEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QHeaderView.initialize")
	rvx := ffiqt.CallbackAllInherits(cthis, "initialize")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView10initializeEv", C.callback_ZN11QHeaderView10initializeEv /*nil*/)
}

// void initializeSections()
//export callback_ZN11QHeaderView18initializeSectionsEv
func callback_ZN11QHeaderView18initializeSectionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QHeaderView.initializeSections")
	rvx := ffiqt.CallbackAllInherits(cthis, "initializeSections")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView18initializeSectionsEv", C.callback_ZN11QHeaderView18initializeSectionsEv /*nil*/)
}

// void initializeSections(int, int)
//export callback_ZN11QHeaderView18initializeSectionsEii
func callback_ZN11QHeaderView18initializeSectionsEii(cthis unsafe.Pointer, start C.int, end C.int) {
	// log.Println(cthis, "QHeaderView.initializeSections")
	rvx := ffiqt.CallbackAllInherits(cthis, "initializeSections", start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView18initializeSectionsEii", C.callback_ZN11QHeaderView18initializeSectionsEii /*nil*/)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
//export callback_ZN11QHeaderView14currentChangedERK11QModelIndexS2_
func callback_ZN11QHeaderView14currentChangedERK11QModelIndexS2_(cthis unsafe.Pointer, current unsafe.Pointer /*555*/, old unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QHeaderView.currentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentChanged", current, old)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView14currentChangedERK11QModelIndexS2_", C.callback_ZN11QHeaderView14currentChangedERK11QModelIndexS2_ /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QHeaderView5eventEP6QEvent
func callback_ZN11QHeaderView5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView5eventEP6QEvent", C.callback_ZN11QHeaderView5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN11QHeaderView10paintEventEP11QPaintEvent
func callback_ZN11QHeaderView10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView10paintEventEP11QPaintEvent", C.callback_ZN11QHeaderView10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN11QHeaderView15mousePressEventEP11QMouseEvent
func callback_ZN11QHeaderView15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView15mousePressEventEP11QMouseEvent", C.callback_ZN11QHeaderView15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent
func callback_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent", C.callback_ZN11QHeaderView14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent
func callback_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent", C.callback_ZN11QHeaderView17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN11QHeaderView21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// bool viewportEvent(class QEvent *)
//export callback_ZN11QHeaderView13viewportEventEP6QEvent
func callback_ZN11QHeaderView13viewportEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.viewportEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView13viewportEventEP6QEvent", C.callback_ZN11QHeaderView13viewportEventEP6QEvent /*nil*/)
}

// void paintSection(class QPainter *, const class QRect &, int)
//export callback_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti
func callback_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/, logicalIndex C.int) {
	// log.Println(cthis, "QHeaderView.paintSection")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintSection", painter, rect, logicalIndex)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti", C.callback_ZNK11QHeaderView12paintSectionEP8QPainterRK5QRecti /*nil*/)
}

// QSize sectionSizeFromContents(int)
//export callback_ZNK11QHeaderView23sectionSizeFromContentsEi
func callback_ZNK11QHeaderView23sectionSizeFromContentsEi(cthis unsafe.Pointer, logicalIndex C.int) {
	// log.Println(cthis, "QHeaderView.sectionSizeFromContents")
	rvx := ffiqt.CallbackAllInherits(cthis, "sectionSizeFromContents", logicalIndex)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView23sectionSizeFromContentsEi", C.callback_ZNK11QHeaderView23sectionSizeFromContentsEi /*nil*/)
}

// int horizontalOffset()
//export callback_ZNK11QHeaderView16horizontalOffsetEv
func callback_ZNK11QHeaderView16horizontalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QHeaderView.horizontalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView16horizontalOffsetEv", C.callback_ZNK11QHeaderView16horizontalOffsetEv /*nil*/)
}

// int verticalOffset()
//export callback_ZNK11QHeaderView14verticalOffsetEv
func callback_ZNK11QHeaderView14verticalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QHeaderView.verticalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView14verticalOffsetEv", C.callback_ZNK11QHeaderView14verticalOffsetEv /*nil*/)
}

// void updateGeometries()
//export callback_ZN11QHeaderView16updateGeometriesEv
func callback_ZN11QHeaderView16updateGeometriesEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QHeaderView.updateGeometries")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateGeometries")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView16updateGeometriesEv", C.callback_ZN11QHeaderView16updateGeometriesEv /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN11QHeaderView16scrollContentsByEii
func callback_ZN11QHeaderView16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QHeaderView.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView16scrollContentsByEii", C.callback_ZN11QHeaderView16scrollContentsByEii /*nil*/)
}

// void rowsInserted(const class QModelIndex &, int, int)
//export callback_ZN11QHeaderView12rowsInsertedERK11QModelIndexii
func callback_ZN11QHeaderView12rowsInsertedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QHeaderView.rowsInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsInserted", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView12rowsInsertedERK11QModelIndexii", C.callback_ZN11QHeaderView12rowsInsertedERK11QModelIndexii /*nil*/)
}

// QRect visualRect(const class QModelIndex &)
//export callback_ZNK11QHeaderView10visualRectERK11QModelIndex
func callback_ZNK11QHeaderView10visualRectERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QHeaderView.visualRect")
	rvx := ffiqt.CallbackAllInherits(cthis, "visualRect", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView10visualRectERK11QModelIndex", C.callback_ZNK11QHeaderView10visualRectERK11QModelIndex /*nil*/)
}

// void scrollTo(const class QModelIndex &, enum QAbstractItemView::ScrollHint)
//export callback_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE
func callback_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE(cthis unsafe.Pointer, index unsafe.Pointer /*555*/, hint C.int) {
	// log.Println(cthis, "QHeaderView.scrollTo")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollTo", index, hint)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE", C.callback_ZN11QHeaderView8scrollToERK11QModelIndexN17QAbstractItemView10ScrollHintE /*nil*/)
}

// QModelIndex indexAt(const class QPoint &)
//export callback_ZNK11QHeaderView7indexAtERK6QPoint
func callback_ZNK11QHeaderView7indexAtERK6QPoint(cthis unsafe.Pointer, p unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QHeaderView.indexAt")
	rvx := ffiqt.CallbackAllInherits(cthis, "indexAt", p)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView7indexAtERK6QPoint", C.callback_ZNK11QHeaderView7indexAtERK6QPoint /*nil*/)
}

// bool isIndexHidden(const class QModelIndex &)
//export callback_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex
func callback_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QHeaderView.isIndexHidden")
	rvx := ffiqt.CallbackAllInherits(cthis, "isIndexHidden", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex", C.callback_ZNK11QHeaderView13isIndexHiddenERK11QModelIndex /*nil*/)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
//export callback_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE
func callback_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(cthis unsafe.Pointer, rect unsafe.Pointer /*555*/, flags C.int) {
	// log.Println(cthis, "QHeaderView.setSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "setSelection", rect, flags)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", C.callback_ZN11QHeaderView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE /*nil*/)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
//export callback_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection
func callback_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection(cthis unsafe.Pointer, selection unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QHeaderView.visualRegionForSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "visualRegionForSelection", selection)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection", C.callback_ZNK11QHeaderView24visualRegionForSelectionERK14QItemSelection /*nil*/)
}

// void initStyleOption(class QStyleOptionHeader *)
//export callback_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader
func callback_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHeaderView.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader", C.callback_ZNK11QHeaderView15initStyleOptionEP18QStyleOptionHeader /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN9QLineEdit15mousePressEventEP11QMouseEvent
func callback_ZN9QLineEdit15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit15mousePressEventEP11QMouseEvent", C.callback_ZN9QLineEdit15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN9QLineEdit14mouseMoveEventEP11QMouseEvent
func callback_ZN9QLineEdit14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit14mouseMoveEventEP11QMouseEvent", C.callback_ZN9QLineEdit14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN9QLineEdit17mouseReleaseEventEP11QMouseEvent
func callback_ZN9QLineEdit17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit17mouseReleaseEventEP11QMouseEvent", C.callback_ZN9QLineEdit17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN9QLineEdit21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN9QLineEdit21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN9QLineEdit21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN9QLineEdit13keyPressEventEP9QKeyEvent
func callback_ZN9QLineEdit13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit13keyPressEventEP9QKeyEvent", C.callback_ZN9QLineEdit13keyPressEventEP9QKeyEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN9QLineEdit12focusInEventEP11QFocusEvent
func callback_ZN9QLineEdit12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit12focusInEventEP11QFocusEvent", C.callback_ZN9QLineEdit12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN9QLineEdit13focusOutEventEP11QFocusEvent
func callback_ZN9QLineEdit13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit13focusOutEventEP11QFocusEvent", C.callback_ZN9QLineEdit13focusOutEventEP11QFocusEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QLineEdit10paintEventEP11QPaintEvent
func callback_ZN9QLineEdit10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit10paintEventEP11QPaintEvent", C.callback_ZN9QLineEdit10paintEventEP11QPaintEvent /*nil*/)
}

// void dragEnterEvent(class QDragEnterEvent *)
//export callback_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent
func callback_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent", C.callback_ZN9QLineEdit14dragEnterEventEP15QDragEnterEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent
func callback_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent", C.callback_ZN9QLineEdit13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN9QLineEdit14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN9QLineEdit9dropEventEP10QDropEvent
func callback_ZN9QLineEdit9dropEventEP10QDropEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit9dropEventEP10QDropEvent", C.callback_ZN9QLineEdit9dropEventEP10QDropEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN9QLineEdit11changeEventEP6QEvent
func callback_ZN9QLineEdit11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit11changeEventEP6QEvent", C.callback_ZN9QLineEdit11changeEventEP6QEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN9QLineEdit16contextMenuEventEP17QContextMenuEvent
func callback_ZN9QLineEdit16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit16contextMenuEventEP17QContextMenuEvent", C.callback_ZN9QLineEdit16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN9QLineEdit16inputMethodEventEP17QInputMethodEvent
func callback_ZN9QLineEdit16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QLineEdit16inputMethodEventEP17QInputMethodEvent", C.callback_ZN9QLineEdit16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionFrame *)
//export callback_ZNK9QLineEdit15initStyleOptionEP17QStyleOptionFrame
func callback_ZNK9QLineEdit15initStyleOptionEP17QStyleOptionFrame(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLineEdit.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QLineEdit15initStyleOptionEP17QStyleOptionFrame", C.callback_ZNK9QLineEdit15initStyleOptionEP17QStyleOptionFrame /*nil*/)
}

// QRect cursorRect()
//export callback_ZNK9QLineEdit10cursorRectEv
func callback_ZNK9QLineEdit10cursorRectEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QLineEdit.cursorRect")
	rvx := ffiqt.CallbackAllInherits(cthis, "cursorRect")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QLineEdit10cursorRectEv", C.callback_ZNK9QLineEdit10cursorRectEv /*nil*/)
}

// void drawDisplay(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QString &)
//export callback_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString
func callback_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, option unsafe.Pointer /*555*/, rect unsafe.Pointer /*555*/, text unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.drawDisplay")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawDisplay", painter, option, rect, text)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString", C.callback_ZNK13QItemDelegate11drawDisplayEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QString /*nil*/)
}

// void drawDecoration(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, const class QPixmap &)
//export callback_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap
func callback_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, option unsafe.Pointer /*555*/, rect unsafe.Pointer /*555*/, pixmap unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.drawDecoration")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawDecoration", painter, option, rect, pixmap)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap", C.callback_ZNK13QItemDelegate14drawDecorationEP8QPainterRK20QStyleOptionViewItemRK5QRectRK7QPixmap /*nil*/)
}

// void drawFocus(class QPainter *, const class QStyleOptionViewItem &, const class QRect &)
//export callback_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect
func callback_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, option unsafe.Pointer /*555*/, rect unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.drawFocus")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawFocus", painter, option, rect)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect", C.callback_ZNK13QItemDelegate9drawFocusEP8QPainterRK20QStyleOptionViewItemRK5QRect /*nil*/)
}

// void drawCheck(class QPainter *, const class QStyleOptionViewItem &, const class QRect &, Qt::CheckState)
//export callback_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE
func callback_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, option unsafe.Pointer /*555*/, rect unsafe.Pointer /*555*/, state C.int) {
	// log.Println(cthis, "QItemDelegate.drawCheck")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawCheck", painter, option, rect, state)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE", C.callback_ZNK13QItemDelegate9drawCheckEP8QPainterRK20QStyleOptionViewItemRK5QRectN2Qt10CheckStateE /*nil*/)
}

// void drawBackground(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
//export callback_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex
func callback_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, option unsafe.Pointer /*555*/, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.drawBackground")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawBackground", painter, option, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", C.callback_ZNK13QItemDelegate14drawBackgroundEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex /*nil*/)
}

// void doLayout(const class QStyleOptionViewItem &, class QRect *, class QRect *, class QRect *, _Bool)
//export callback_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b
func callback_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b(cthis unsafe.Pointer, option unsafe.Pointer /*555*/, checkRect unsafe.Pointer /*666*/, iconRect unsafe.Pointer /*666*/, textRect unsafe.Pointer /*666*/, hint C.bool) {
	// log.Println(cthis, "QItemDelegate.doLayout")
	rvx := ffiqt.CallbackAllInherits(cthis, "doLayout", option, checkRect, iconRect, textRect, hint)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b", C.callback_ZNK13QItemDelegate8doLayoutERK20QStyleOptionViewItemP5QRectS4_S4_b /*nil*/)
}

// QRect rect(const class QStyleOptionViewItem &, const class QModelIndex &, int)
//export callback_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi
func callback_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi(cthis unsafe.Pointer, option unsafe.Pointer /*555*/, index unsafe.Pointer /*555*/, role C.int) {
	// log.Println(cthis, "QItemDelegate.rect")
	rvx := ffiqt.CallbackAllInherits(cthis, "rect", option, index, role)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi", C.callback_ZNK13QItemDelegate4rectERK20QStyleOptionViewItemRK11QModelIndexi /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent
func callback_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, object unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QItemDelegate.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", object, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent", C.callback_ZN13QItemDelegate11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
//export callback_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex
func callback_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex(cthis unsafe.Pointer, event unsafe.Pointer /*666*/, model unsafe.Pointer /*666*/, option unsafe.Pointer /*555*/, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.editorEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "editorEvent", event, model, option, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", C.callback_ZN13QItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex /*nil*/)
}

// QStyleOptionViewItem setOptions(const class QModelIndex &, const class QStyleOptionViewItem &)
//export callback_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem
func callback_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem(cthis unsafe.Pointer, index unsafe.Pointer /*555*/, option unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.setOptions")
	rvx := ffiqt.CallbackAllInherits(cthis, "setOptions", index, option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem", C.callback_ZNK13QItemDelegate10setOptionsERK11QModelIndexRK20QStyleOptionViewItem /*nil*/)
}

// QPixmap decoration(const class QStyleOptionViewItem &, const class QVariant &)
//export callback_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant
func callback_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant(cthis unsafe.Pointer, option unsafe.Pointer /*555*/, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.decoration")
	rvx := ffiqt.CallbackAllInherits(cthis, "decoration", option, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant", C.callback_ZNK13QItemDelegate10decorationERK20QStyleOptionViewItemRK8QVariant /*nil*/)
}

// QPixmap * selected(const class QPixmap &, const class QPalette &, _Bool)
//export callback_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb
func callback_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb(cthis unsafe.Pointer, pixmap unsafe.Pointer /*555*/, palette unsafe.Pointer /*555*/, enabled C.bool) {
	// log.Println(cthis, "QItemDelegate.selected")
	rvx := ffiqt.CallbackAllInherits(cthis, "selected", pixmap, palette, enabled)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb", C.callback_ZNK13QItemDelegate8selectedERK7QPixmapRK8QPaletteb /*nil*/)
}

// QRect doCheck(const class QStyleOptionViewItem &, const class QRect &, const class QVariant &)
//export callback_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant
func callback_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant(cthis unsafe.Pointer, option unsafe.Pointer /*555*/, bounding unsafe.Pointer /*555*/, variant unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.doCheck")
	rvx := ffiqt.CallbackAllInherits(cthis, "doCheck", option, bounding, variant)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant", C.callback_ZNK13QItemDelegate7doCheckERK20QStyleOptionViewItemRK5QRectRK8QVariant /*nil*/)
}

// QRect textRectangle(class QPainter *, const class QRect &, const class QFont &, const class QString &)
//export callback_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString
func callback_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/, font unsafe.Pointer /*555*/, text unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemDelegate.textRectangle")
	rvx := ffiqt.CallbackAllInherits(cthis, "textRectangle", painter, rect, font, text)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString", C.callback_ZNK13QItemDelegate13textRectangleEP8QPainterRK5QRectRK5QFontRK7QString /*nil*/)
}

// void onTransition(class QEvent *)
//export callback_ZN19QKeyEventTransition12onTransitionEP6QEvent
func callback_ZN19QKeyEventTransition12onTransitionEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QKeyEventTransition.onTransition")
	rvx := ffiqt.CallbackAllInherits(cthis, "onTransition", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QKeyEventTransition12onTransitionEP6QEvent", C.callback_ZN19QKeyEventTransition12onTransitionEP6QEvent /*nil*/)
}

// bool eventTest(class QEvent *)
//export callback_ZN19QKeyEventTransition9eventTestEP6QEvent
func callback_ZN19QKeyEventTransition9eventTestEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QKeyEventTransition.eventTest")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventTest", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QKeyEventTransition9eventTestEP6QEvent", C.callback_ZN19QKeyEventTransition9eventTestEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN16QKeySequenceEdit5eventEP6QEvent
func callback_ZN16QKeySequenceEdit5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QKeySequenceEdit.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QKeySequenceEdit5eventEP6QEvent", C.callback_ZN16QKeySequenceEdit5eventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent
func callback_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QKeySequenceEdit.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent", C.callback_ZN16QKeySequenceEdit13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent
func callback_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QKeySequenceEdit.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent", C.callback_ZN16QKeySequenceEdit15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent
func callback_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QKeySequenceEdit.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent", C.callback_ZN16QKeySequenceEdit10timerEventEP11QTimerEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN6QLabel5eventEP6QEvent
func callback_ZN6QLabel5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel5eventEP6QEvent", C.callback_ZN6QLabel5eventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN6QLabel13keyPressEventEP9QKeyEvent
func callback_ZN6QLabel13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel13keyPressEventEP9QKeyEvent", C.callback_ZN6QLabel13keyPressEventEP9QKeyEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN6QLabel10paintEventEP11QPaintEvent
func callback_ZN6QLabel10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel10paintEventEP11QPaintEvent", C.callback_ZN6QLabel10paintEventEP11QPaintEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN6QLabel11changeEventEP6QEvent
func callback_ZN6QLabel11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel11changeEventEP6QEvent", C.callback_ZN6QLabel11changeEventEP6QEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN6QLabel15mousePressEventEP11QMouseEvent
func callback_ZN6QLabel15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel15mousePressEventEP11QMouseEvent", C.callback_ZN6QLabel15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN6QLabel14mouseMoveEventEP11QMouseEvent
func callback_ZN6QLabel14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel14mouseMoveEventEP11QMouseEvent", C.callback_ZN6QLabel14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN6QLabel17mouseReleaseEventEP11QMouseEvent
func callback_ZN6QLabel17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel17mouseReleaseEventEP11QMouseEvent", C.callback_ZN6QLabel17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN6QLabel16contextMenuEventEP17QContextMenuEvent
func callback_ZN6QLabel16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel16contextMenuEventEP17QContextMenuEvent", C.callback_ZN6QLabel16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN6QLabel12focusInEventEP11QFocusEvent
func callback_ZN6QLabel12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel12focusInEventEP11QFocusEvent", C.callback_ZN6QLabel12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN6QLabel13focusOutEventEP11QFocusEvent
func callback_ZN6QLabel13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLabel.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel13focusOutEventEP11QFocusEvent", C.callback_ZN6QLabel13focusOutEventEP11QFocusEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN6QLabel18focusNextPrevChildEb
func callback_ZN6QLabel18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QLabel.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QLabel18focusNextPrevChildEb", C.callback_ZN6QLabel18focusNextPrevChildEb /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN10QLCDNumber5eventEP6QEvent
func callback_ZN10QLCDNumber5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLCDNumber.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QLCDNumber5eventEP6QEvent", C.callback_ZN10QLCDNumber5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN10QLCDNumber10paintEventEP11QPaintEvent
func callback_ZN10QLCDNumber10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QLCDNumber.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QLCDNumber10paintEventEP11QPaintEvent", C.callback_ZN10QLCDNumber10paintEventEP11QPaintEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QListView5eventEP6QEvent
func callback_ZN9QListView5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView5eventEP6QEvent", C.callback_ZN9QListView5eventEP6QEvent /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN9QListView16scrollContentsByEii
func callback_ZN9QListView16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QListView.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView16scrollContentsByEii", C.callback_ZN9QListView16scrollContentsByEii /*nil*/)
}

// void resizeContents(int, int)
//export callback_ZN9QListView14resizeContentsEii
func callback_ZN9QListView14resizeContentsEii(cthis unsafe.Pointer, width C.int, height C.int) {
	// log.Println(cthis, "QListView.resizeContents")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeContents", width, height)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView14resizeContentsEii", C.callback_ZN9QListView14resizeContentsEii /*nil*/)
}

// QSize contentsSize()
//export callback_ZNK9QListView12contentsSizeEv
func callback_ZNK9QListView12contentsSizeEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QListView.contentsSize")
	rvx := ffiqt.CallbackAllInherits(cthis, "contentsSize")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView12contentsSizeEv", C.callback_ZNK9QListView12contentsSizeEv /*nil*/)
}

// void rowsInserted(const class QModelIndex &, int, int)
//export callback_ZN9QListView12rowsInsertedERK11QModelIndexii
func callback_ZN9QListView12rowsInsertedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QListView.rowsInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsInserted", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView12rowsInsertedERK11QModelIndexii", C.callback_ZN9QListView12rowsInsertedERK11QModelIndexii /*nil*/)
}

// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
//export callback_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii
func callback_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QListView.rowsAboutToBeRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsAboutToBeRemoved", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii", C.callback_ZN9QListView20rowsAboutToBeRemovedERK11QModelIndexii /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN9QListView14mouseMoveEventEP11QMouseEvent
func callback_ZN9QListView14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView14mouseMoveEventEP11QMouseEvent", C.callback_ZN9QListView14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN9QListView17mouseReleaseEventEP11QMouseEvent
func callback_ZN9QListView17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView17mouseReleaseEventEP11QMouseEvent", C.callback_ZN9QListView17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN9QListView10wheelEventEP11QWheelEvent
func callback_ZN9QListView10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView10wheelEventEP11QWheelEvent", C.callback_ZN9QListView10wheelEventEP11QWheelEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN9QListView10timerEventEP11QTimerEvent
func callback_ZN9QListView10timerEventEP11QTimerEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView10timerEventEP11QTimerEvent", C.callback_ZN9QListView10timerEventEP11QTimerEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN9QListView11resizeEventEP12QResizeEvent
func callback_ZN9QListView11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView11resizeEventEP12QResizeEvent", C.callback_ZN9QListView11resizeEventEP12QResizeEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN9QListView13dragMoveEventEP14QDragMoveEvent
func callback_ZN9QListView13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView13dragMoveEventEP14QDragMoveEvent", C.callback_ZN9QListView13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN9QListView14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN9QListView9dropEventEP10QDropEvent
func callback_ZN9QListView9dropEventEP10QDropEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView9dropEventEP10QDropEvent", C.callback_ZN9QListView9dropEventEP10QDropEvent /*nil*/)
}

// QStyleOptionViewItem viewOptions()
//export callback_ZNK9QListView11viewOptionsEv
func callback_ZNK9QListView11viewOptionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QListView.viewOptions")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewOptions")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView11viewOptionsEv", C.callback_ZNK9QListView11viewOptionsEv /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QListView10paintEventEP11QPaintEvent
func callback_ZN9QListView10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListView.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView10paintEventEP11QPaintEvent", C.callback_ZN9QListView10paintEventEP11QPaintEvent /*nil*/)
}

// int horizontalOffset()
//export callback_ZNK9QListView16horizontalOffsetEv
func callback_ZNK9QListView16horizontalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QListView.horizontalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView16horizontalOffsetEv", C.callback_ZNK9QListView16horizontalOffsetEv /*nil*/)
}

// int verticalOffset()
//export callback_ZNK9QListView14verticalOffsetEv
func callback_ZNK9QListView14verticalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QListView.verticalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView14verticalOffsetEv", C.callback_ZNK9QListView14verticalOffsetEv /*nil*/)
}

// QRect rectForIndex(const class QModelIndex &)
//export callback_ZNK9QListView12rectForIndexERK11QModelIndex
func callback_ZNK9QListView12rectForIndexERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QListView.rectForIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "rectForIndex", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView12rectForIndexERK11QModelIndex", C.callback_ZNK9QListView12rectForIndexERK11QModelIndex /*nil*/)
}

// void setPositionForIndex(const class QPoint &, const class QModelIndex &)
//export callback_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex
func callback_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex(cthis unsafe.Pointer, position unsafe.Pointer /*555*/, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QListView.setPositionForIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "setPositionForIndex", position, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex", C.callback_ZN9QListView19setPositionForIndexERK6QPointRK11QModelIndex /*nil*/)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
//export callback_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE
func callback_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(cthis unsafe.Pointer, rect unsafe.Pointer /*555*/, command C.int) {
	// log.Println(cthis, "QListView.setSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "setSelection", rect, command)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", C.callback_ZN9QListView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE /*nil*/)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
//export callback_ZNK9QListView24visualRegionForSelectionERK14QItemSelection
func callback_ZNK9QListView24visualRegionForSelectionERK14QItemSelection(cthis unsafe.Pointer, selection unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QListView.visualRegionForSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "visualRegionForSelection", selection)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView24visualRegionForSelectionERK14QItemSelection", C.callback_ZNK9QListView24visualRegionForSelectionERK14QItemSelection /*nil*/)
}

// void updateGeometries()
//export callback_ZN9QListView16updateGeometriesEv
func callback_ZN9QListView16updateGeometriesEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QListView.updateGeometries")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateGeometries")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView16updateGeometriesEv", C.callback_ZN9QListView16updateGeometriesEv /*nil*/)
}

// bool isIndexHidden(const class QModelIndex &)
//export callback_ZNK9QListView13isIndexHiddenERK11QModelIndex
func callback_ZNK9QListView13isIndexHiddenERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QListView.isIndexHidden")
	rvx := ffiqt.CallbackAllInherits(cthis, "isIndexHidden", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView13isIndexHiddenERK11QModelIndex", C.callback_ZNK9QListView13isIndexHiddenERK11QModelIndex /*nil*/)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
//export callback_ZN9QListView16selectionChangedERK14QItemSelectionS2_
func callback_ZN9QListView16selectionChangedERK14QItemSelectionS2_(cthis unsafe.Pointer, selected unsafe.Pointer /*555*/, deselected unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QListView.selectionChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "selectionChanged", selected, deselected)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView16selectionChangedERK14QItemSelectionS2_", C.callback_ZN9QListView16selectionChangedERK14QItemSelectionS2_ /*nil*/)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
//export callback_ZN9QListView14currentChangedERK11QModelIndexS2_
func callback_ZN9QListView14currentChangedERK11QModelIndexS2_(cthis unsafe.Pointer, current unsafe.Pointer /*555*/, previous unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QListView.currentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentChanged", current, previous)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QListView14currentChangedERK11QModelIndexS2_", C.callback_ZN9QListView14currentChangedERK11QModelIndexS2_ /*nil*/)
}

// QSize viewportSizeHint()
//export callback_ZNK9QListView16viewportSizeHintEv
func callback_ZNK9QListView16viewportSizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QListView.viewportSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportSizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QListView16viewportSizeHintEv", C.callback_ZNK9QListView16viewportSizeHintEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QListWidget5eventEP6QEvent
func callback_ZN11QListWidget5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QListWidget5eventEP6QEvent", C.callback_ZN11QListWidget5eventEP6QEvent /*nil*/)
}

// bool dropMimeData(int, const class QMimeData *, Qt::DropAction)
//export callback_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE
func callback_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE(cthis unsafe.Pointer, index C.int, data unsafe.Pointer /*666*/, action C.int) {
	// log.Println(cthis, "QListWidget.dropMimeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropMimeData", index, data, action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE", C.callback_ZN11QListWidget12dropMimeDataEiPK9QMimeDataN2Qt10DropActionE /*nil*/)
}

// Qt::DropActions supportedDropActions()
//export callback_ZNK11QListWidget20supportedDropActionsEv
func callback_ZNK11QListWidget20supportedDropActionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QListWidget.supportedDropActions")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportedDropActions")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QListWidget20supportedDropActionsEv", C.callback_ZNK11QListWidget20supportedDropActionsEv /*nil*/)
}

// QModelIndex indexFromItem(class QListWidgetItem *)
//export callback_ZNK11QListWidget13indexFromItemEP15QListWidgetItem
func callback_ZNK11QListWidget13indexFromItemEP15QListWidgetItem(cthis unsafe.Pointer, item unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QListWidget.indexFromItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "indexFromItem", item)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QListWidget13indexFromItemEP15QListWidgetItem", C.callback_ZNK11QListWidget13indexFromItemEP15QListWidgetItem /*nil*/)
}

// QListWidgetItem * itemFromIndex(const class QModelIndex &)
//export callback_ZNK11QListWidget13itemFromIndexERK11QModelIndex
func callback_ZNK11QListWidget13itemFromIndexERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QListWidget.itemFromIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemFromIndex", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QListWidget13itemFromIndexERK11QModelIndex", C.callback_ZNK11QListWidget13itemFromIndexERK11QModelIndex /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent
func callback_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMainWindow.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent", C.callback_ZN11QMainWindow16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QMainWindow5eventEP6QEvent
func callback_ZN11QMainWindow5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMainWindow.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMainWindow5eventEP6QEvent", C.callback_ZN11QMainWindow5eventEP6QEvent /*nil*/)
}

// void setupViewport(class QWidget *)
//export callback_ZN8QMdiArea13setupViewportEP7QWidget
func callback_ZN8QMdiArea13setupViewportEP7QWidget(cthis unsafe.Pointer, viewport unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.setupViewport")
	rvx := ffiqt.CallbackAllInherits(cthis, "setupViewport", viewport)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea13setupViewportEP7QWidget", C.callback_ZN8QMdiArea13setupViewportEP7QWidget /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN8QMdiArea5eventEP6QEvent
func callback_ZN8QMdiArea5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea5eventEP6QEvent", C.callback_ZN8QMdiArea5eventEP6QEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent
func callback_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, object unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", object, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent", C.callback_ZN8QMdiArea11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN8QMdiArea10paintEventEP11QPaintEvent
func callback_ZN8QMdiArea10paintEventEP11QPaintEvent(cthis unsafe.Pointer, paintEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", paintEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea10paintEventEP11QPaintEvent", C.callback_ZN8QMdiArea10paintEventEP11QPaintEvent /*nil*/)
}

// void childEvent(class QChildEvent *)
//export callback_ZN8QMdiArea10childEventEP11QChildEvent
func callback_ZN8QMdiArea10childEventEP11QChildEvent(cthis unsafe.Pointer, childEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.childEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "childEvent", childEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea10childEventEP11QChildEvent", C.callback_ZN8QMdiArea10childEventEP11QChildEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN8QMdiArea11resizeEventEP12QResizeEvent
func callback_ZN8QMdiArea11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, resizeEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", resizeEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea11resizeEventEP12QResizeEvent", C.callback_ZN8QMdiArea11resizeEventEP12QResizeEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN8QMdiArea10timerEventEP11QTimerEvent
func callback_ZN8QMdiArea10timerEventEP11QTimerEvent(cthis unsafe.Pointer, timerEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", timerEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea10timerEventEP11QTimerEvent", C.callback_ZN8QMdiArea10timerEventEP11QTimerEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN8QMdiArea9showEventEP10QShowEvent
func callback_ZN8QMdiArea9showEventEP10QShowEvent(cthis unsafe.Pointer, showEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", showEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea9showEventEP10QShowEvent", C.callback_ZN8QMdiArea9showEventEP10QShowEvent /*nil*/)
}

// bool viewportEvent(class QEvent *)
//export callback_ZN8QMdiArea13viewportEventEP6QEvent
func callback_ZN8QMdiArea13viewportEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiArea.viewportEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea13viewportEventEP6QEvent", C.callback_ZN8QMdiArea13viewportEventEP6QEvent /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN8QMdiArea16scrollContentsByEii
func callback_ZN8QMdiArea16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QMdiArea.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMdiArea16scrollContentsByEii", C.callback_ZN8QMdiArea16scrollContentsByEii /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent
func callback_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, object unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", object, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent", C.callback_ZN13QMdiSubWindow11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN13QMdiSubWindow5eventEP6QEvent
func callback_ZN13QMdiSubWindow5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow5eventEP6QEvent", C.callback_ZN13QMdiSubWindow5eventEP6QEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN13QMdiSubWindow9showEventEP10QShowEvent
func callback_ZN13QMdiSubWindow9showEventEP10QShowEvent(cthis unsafe.Pointer, showEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", showEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow9showEventEP10QShowEvent", C.callback_ZN13QMdiSubWindow9showEventEP10QShowEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN13QMdiSubWindow9hideEventEP10QHideEvent
func callback_ZN13QMdiSubWindow9hideEventEP10QHideEvent(cthis unsafe.Pointer, hideEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", hideEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow9hideEventEP10QHideEvent", C.callback_ZN13QMdiSubWindow9hideEventEP10QHideEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN13QMdiSubWindow11changeEventEP6QEvent
func callback_ZN13QMdiSubWindow11changeEventEP6QEvent(cthis unsafe.Pointer, changeEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", changeEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow11changeEventEP6QEvent", C.callback_ZN13QMdiSubWindow11changeEventEP6QEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN13QMdiSubWindow10closeEventEP11QCloseEvent
func callback_ZN13QMdiSubWindow10closeEventEP11QCloseEvent(cthis unsafe.Pointer, closeEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", closeEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow10closeEventEP11QCloseEvent", C.callback_ZN13QMdiSubWindow10closeEventEP11QCloseEvent /*nil*/)
}

// void leaveEvent(class QEvent *)
//export callback_ZN13QMdiSubWindow10leaveEventEP6QEvent
func callback_ZN13QMdiSubWindow10leaveEventEP6QEvent(cthis unsafe.Pointer, leaveEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.leaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "leaveEvent", leaveEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow10leaveEventEP6QEvent", C.callback_ZN13QMdiSubWindow10leaveEventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent
func callback_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, resizeEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", resizeEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent", C.callback_ZN13QMdiSubWindow11resizeEventEP12QResizeEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN13QMdiSubWindow10timerEventEP11QTimerEvent
func callback_ZN13QMdiSubWindow10timerEventEP11QTimerEvent(cthis unsafe.Pointer, timerEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", timerEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow10timerEventEP11QTimerEvent", C.callback_ZN13QMdiSubWindow10timerEventEP11QTimerEvent /*nil*/)
}

// void moveEvent(class QMoveEvent *)
//export callback_ZN13QMdiSubWindow9moveEventEP10QMoveEvent
func callback_ZN13QMdiSubWindow9moveEventEP10QMoveEvent(cthis unsafe.Pointer, moveEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.moveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveEvent", moveEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow9moveEventEP10QMoveEvent", C.callback_ZN13QMdiSubWindow9moveEventEP10QMoveEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN13QMdiSubWindow10paintEventEP11QPaintEvent
func callback_ZN13QMdiSubWindow10paintEventEP11QPaintEvent(cthis unsafe.Pointer, paintEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", paintEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow10paintEventEP11QPaintEvent", C.callback_ZN13QMdiSubWindow10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent
func callback_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, mouseEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", mouseEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent", C.callback_ZN13QMdiSubWindow15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, mouseEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", mouseEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN13QMdiSubWindow21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent
func callback_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, mouseEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", mouseEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent", C.callback_ZN13QMdiSubWindow17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent
func callback_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, mouseEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", mouseEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent", C.callback_ZN13QMdiSubWindow14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent
func callback_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, keyEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", keyEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent", C.callback_ZN13QMdiSubWindow13keyPressEventEP9QKeyEvent /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent
func callback_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, contextMenuEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", contextMenuEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent", C.callback_ZN13QMdiSubWindow16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent
func callback_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, focusInEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", focusInEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent", C.callback_ZN13QMdiSubWindow12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent
func callback_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, focusOutEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", focusOutEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent", C.callback_ZN13QMdiSubWindow13focusOutEventEP11QFocusEvent /*nil*/)
}

// void childEvent(class QChildEvent *)
//export callback_ZN13QMdiSubWindow10childEventEP11QChildEvent
func callback_ZN13QMdiSubWindow10childEventEP11QChildEvent(cthis unsafe.Pointer, childEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMdiSubWindow.childEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "childEvent", childEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QMdiSubWindow10childEventEP11QChildEvent", C.callback_ZN13QMdiSubWindow10childEventEP11QChildEvent /*nil*/)
}

// int columnCount()
//export callback_ZNK5QMenu11columnCountEv
func callback_ZNK5QMenu11columnCountEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QMenu.columnCount")
	rvx := ffiqt.CallbackAllInherits(cthis, "columnCount")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK5QMenu11columnCountEv", C.callback_ZNK5QMenu11columnCountEv /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN5QMenu11changeEventEP6QEvent
func callback_ZN5QMenu11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu11changeEventEP6QEvent", C.callback_ZN5QMenu11changeEventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN5QMenu13keyPressEventEP9QKeyEvent
func callback_ZN5QMenu13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu13keyPressEventEP9QKeyEvent", C.callback_ZN5QMenu13keyPressEventEP9QKeyEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN5QMenu17mouseReleaseEventEP11QMouseEvent
func callback_ZN5QMenu17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu17mouseReleaseEventEP11QMouseEvent", C.callback_ZN5QMenu17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN5QMenu15mousePressEventEP11QMouseEvent
func callback_ZN5QMenu15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu15mousePressEventEP11QMouseEvent", C.callback_ZN5QMenu15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN5QMenu14mouseMoveEventEP11QMouseEvent
func callback_ZN5QMenu14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu14mouseMoveEventEP11QMouseEvent", C.callback_ZN5QMenu14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN5QMenu10wheelEventEP11QWheelEvent
func callback_ZN5QMenu10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu10wheelEventEP11QWheelEvent", C.callback_ZN5QMenu10wheelEventEP11QWheelEvent /*nil*/)
}

// void enterEvent(class QEvent *)
//export callback_ZN5QMenu10enterEventEP6QEvent
func callback_ZN5QMenu10enterEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.enterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "enterEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu10enterEventEP6QEvent", C.callback_ZN5QMenu10enterEventEP6QEvent /*nil*/)
}

// void leaveEvent(class QEvent *)
//export callback_ZN5QMenu10leaveEventEP6QEvent
func callback_ZN5QMenu10leaveEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.leaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "leaveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu10leaveEventEP6QEvent", C.callback_ZN5QMenu10leaveEventEP6QEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN5QMenu9hideEventEP10QHideEvent
func callback_ZN5QMenu9hideEventEP10QHideEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu9hideEventEP10QHideEvent", C.callback_ZN5QMenu9hideEventEP10QHideEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN5QMenu10paintEventEP11QPaintEvent
func callback_ZN5QMenu10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu10paintEventEP11QPaintEvent", C.callback_ZN5QMenu10paintEventEP11QPaintEvent /*nil*/)
}

// void actionEvent(class QActionEvent *)
//export callback_ZN5QMenu11actionEventEP12QActionEvent
func callback_ZN5QMenu11actionEventEP12QActionEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.actionEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "actionEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu11actionEventEP12QActionEvent", C.callback_ZN5QMenu11actionEventEP12QActionEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN5QMenu10timerEventEP11QTimerEvent
func callback_ZN5QMenu10timerEventEP11QTimerEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu10timerEventEP11QTimerEvent", C.callback_ZN5QMenu10timerEventEP11QTimerEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN5QMenu5eventEP6QEvent
func callback_ZN5QMenu5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu5eventEP6QEvent", C.callback_ZN5QMenu5eventEP6QEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN5QMenu18focusNextPrevChildEb
func callback_ZN5QMenu18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QMenu.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN5QMenu18focusNextPrevChildEb", C.callback_ZN5QMenu18focusNextPrevChildEb /*nil*/)
}

// void initStyleOption(class QStyleOptionMenuItem *, const class QAction *)
//export callback_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction
func callback_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction(cthis unsafe.Pointer, option unsafe.Pointer /*666*/, action unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenu.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option, action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", C.callback_ZNK5QMenu15initStyleOptionEP20QStyleOptionMenuItemPK7QAction /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN8QMenuBar11changeEventEP6QEvent
func callback_ZN8QMenuBar11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar11changeEventEP6QEvent", C.callback_ZN8QMenuBar11changeEventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN8QMenuBar13keyPressEventEP9QKeyEvent
func callback_ZN8QMenuBar13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar13keyPressEventEP9QKeyEvent", C.callback_ZN8QMenuBar13keyPressEventEP9QKeyEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent
func callback_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent", C.callback_ZN8QMenuBar17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN8QMenuBar15mousePressEventEP11QMouseEvent
func callback_ZN8QMenuBar15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar15mousePressEventEP11QMouseEvent", C.callback_ZN8QMenuBar15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent
func callback_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent", C.callback_ZN8QMenuBar14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void leaveEvent(class QEvent *)
//export callback_ZN8QMenuBar10leaveEventEP6QEvent
func callback_ZN8QMenuBar10leaveEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.leaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "leaveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar10leaveEventEP6QEvent", C.callback_ZN8QMenuBar10leaveEventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN8QMenuBar10paintEventEP11QPaintEvent
func callback_ZN8QMenuBar10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar10paintEventEP11QPaintEvent", C.callback_ZN8QMenuBar10paintEventEP11QPaintEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN8QMenuBar11resizeEventEP12QResizeEvent
func callback_ZN8QMenuBar11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar11resizeEventEP12QResizeEvent", C.callback_ZN8QMenuBar11resizeEventEP12QResizeEvent /*nil*/)
}

// void actionEvent(class QActionEvent *)
//export callback_ZN8QMenuBar11actionEventEP12QActionEvent
func callback_ZN8QMenuBar11actionEventEP12QActionEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.actionEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "actionEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar11actionEventEP12QActionEvent", C.callback_ZN8QMenuBar11actionEventEP12QActionEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN8QMenuBar13focusOutEventEP11QFocusEvent
func callback_ZN8QMenuBar13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar13focusOutEventEP11QFocusEvent", C.callback_ZN8QMenuBar13focusOutEventEP11QFocusEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN8QMenuBar12focusInEventEP11QFocusEvent
func callback_ZN8QMenuBar12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar12focusInEventEP11QFocusEvent", C.callback_ZN8QMenuBar12focusInEventEP11QFocusEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN8QMenuBar10timerEventEP11QTimerEvent
func callback_ZN8QMenuBar10timerEventEP11QTimerEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar10timerEventEP11QTimerEvent", C.callback_ZN8QMenuBar10timerEventEP11QTimerEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent
func callback_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent", C.callback_ZN8QMenuBar11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN8QMenuBar5eventEP6QEvent
func callback_ZN8QMenuBar5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QMenuBar5eventEP6QEvent", C.callback_ZN8QMenuBar5eventEP6QEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionMenuItem *, const class QAction *)
//export callback_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction
func callback_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction(cthis unsafe.Pointer, option unsafe.Pointer /*666*/, action unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMenuBar.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option, action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction", C.callback_ZNK8QMenuBar15initStyleOptionEP20QStyleOptionMenuItemPK7QAction /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QMessageBox5eventEP6QEvent
func callback_ZN11QMessageBox5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMessageBox.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMessageBox5eventEP6QEvent", C.callback_ZN11QMessageBox5eventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN11QMessageBox11resizeEventEP12QResizeEvent
func callback_ZN11QMessageBox11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMessageBox.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMessageBox11resizeEventEP12QResizeEvent", C.callback_ZN11QMessageBox11resizeEventEP12QResizeEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN11QMessageBox9showEventEP10QShowEvent
func callback_ZN11QMessageBox9showEventEP10QShowEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMessageBox.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMessageBox9showEventEP10QShowEvent", C.callback_ZN11QMessageBox9showEventEP10QShowEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN11QMessageBox10closeEventEP11QCloseEvent
func callback_ZN11QMessageBox10closeEventEP11QCloseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMessageBox.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMessageBox10closeEventEP11QCloseEvent", C.callback_ZN11QMessageBox10closeEventEP11QCloseEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN11QMessageBox13keyPressEventEP9QKeyEvent
func callback_ZN11QMessageBox13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMessageBox.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMessageBox13keyPressEventEP9QKeyEvent", C.callback_ZN11QMessageBox13keyPressEventEP9QKeyEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN11QMessageBox11changeEventEP6QEvent
func callback_ZN11QMessageBox11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMessageBox.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QMessageBox11changeEventEP6QEvent", C.callback_ZN11QMessageBox11changeEventEP6QEvent /*nil*/)
}

// void onTransition(class QEvent *)
//export callback_ZN21QMouseEventTransition12onTransitionEP6QEvent
func callback_ZN21QMouseEventTransition12onTransitionEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMouseEventTransition.onTransition")
	rvx := ffiqt.CallbackAllInherits(cthis, "onTransition", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QMouseEventTransition12onTransitionEP6QEvent", C.callback_ZN21QMouseEventTransition12onTransitionEP6QEvent /*nil*/)
}

// bool eventTest(class QEvent *)
//export callback_ZN21QMouseEventTransition9eventTestEP6QEvent
func callback_ZN21QMouseEventTransition9eventTestEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QMouseEventTransition.eventTest")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventTest", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QMouseEventTransition9eventTestEP6QEvent", C.callback_ZN21QMouseEventTransition9eventTestEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QTextEdit5eventEP6QEvent
func callback_ZN9QTextEdit5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit5eventEP6QEvent", C.callback_ZN9QTextEdit5eventEP6QEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN9QTextEdit10timerEventEP11QTimerEvent
func callback_ZN9QTextEdit10timerEventEP11QTimerEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit10timerEventEP11QTimerEvent", C.callback_ZN9QTextEdit10timerEventEP11QTimerEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN9QTextEdit13keyPressEventEP9QKeyEvent
func callback_ZN9QTextEdit13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit13keyPressEventEP9QKeyEvent", C.callback_ZN9QTextEdit13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent
func callback_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent", C.callback_ZN9QTextEdit15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN9QTextEdit11resizeEventEP12QResizeEvent
func callback_ZN9QTextEdit11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit11resizeEventEP12QResizeEvent", C.callback_ZN9QTextEdit11resizeEventEP12QResizeEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QTextEdit10paintEventEP11QPaintEvent
func callback_ZN9QTextEdit10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit10paintEventEP11QPaintEvent", C.callback_ZN9QTextEdit10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN9QTextEdit15mousePressEventEP11QMouseEvent
func callback_ZN9QTextEdit15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit15mousePressEventEP11QMouseEvent", C.callback_ZN9QTextEdit15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent
func callback_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent", C.callback_ZN9QTextEdit14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent
func callback_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent", C.callback_ZN9QTextEdit17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN9QTextEdit21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN9QTextEdit18focusNextPrevChildEb
func callback_ZN9QTextEdit18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QTextEdit.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit18focusNextPrevChildEb", C.callback_ZN9QTextEdit18focusNextPrevChildEb /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent
func callback_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent", C.callback_ZN9QTextEdit16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void dragEnterEvent(class QDragEnterEvent *)
//export callback_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent
func callback_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent", C.callback_ZN9QTextEdit14dragEnterEventEP15QDragEnterEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN9QTextEdit14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent
func callback_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent", C.callback_ZN9QTextEdit13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN9QTextEdit9dropEventEP10QDropEvent
func callback_ZN9QTextEdit9dropEventEP10QDropEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit9dropEventEP10QDropEvent", C.callback_ZN9QTextEdit9dropEventEP10QDropEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN9QTextEdit12focusInEventEP11QFocusEvent
func callback_ZN9QTextEdit12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit12focusInEventEP11QFocusEvent", C.callback_ZN9QTextEdit12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN9QTextEdit13focusOutEventEP11QFocusEvent
func callback_ZN9QTextEdit13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit13focusOutEventEP11QFocusEvent", C.callback_ZN9QTextEdit13focusOutEventEP11QFocusEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN9QTextEdit9showEventEP10QShowEvent
func callback_ZN9QTextEdit9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit9showEventEP10QShowEvent", C.callback_ZN9QTextEdit9showEventEP10QShowEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN9QTextEdit11changeEventEP6QEvent
func callback_ZN9QTextEdit11changeEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit11changeEventEP6QEvent", C.callback_ZN9QTextEdit11changeEventEP6QEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN9QTextEdit10wheelEventEP11QWheelEvent
func callback_ZN9QTextEdit10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit10wheelEventEP11QWheelEvent", C.callback_ZN9QTextEdit10wheelEventEP11QWheelEvent /*nil*/)
}

// QMimeData * createMimeDataFromSelection()
//export callback_ZNK9QTextEdit27createMimeDataFromSelectionEv
func callback_ZNK9QTextEdit27createMimeDataFromSelectionEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTextEdit.createMimeDataFromSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "createMimeDataFromSelection")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTextEdit27createMimeDataFromSelectionEv", C.callback_ZNK9QTextEdit27createMimeDataFromSelectionEv /*nil*/)
}

// bool canInsertFromMimeData(const class QMimeData *)
//export callback_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData
func callback_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData(cthis unsafe.Pointer, source unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.canInsertFromMimeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "canInsertFromMimeData", source)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData", C.callback_ZNK9QTextEdit21canInsertFromMimeDataEPK9QMimeData /*nil*/)
}

// void insertFromMimeData(const class QMimeData *)
//export callback_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData
func callback_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData(cthis unsafe.Pointer, source unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.insertFromMimeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "insertFromMimeData", source)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData", C.callback_ZN9QTextEdit18insertFromMimeDataEPK9QMimeData /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent
func callback_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextEdit.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent", C.callback_ZN9QTextEdit16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN9QTextEdit16scrollContentsByEii
func callback_ZN9QTextEdit16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QTextEdit.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit16scrollContentsByEii", C.callback_ZN9QTextEdit16scrollContentsByEii /*nil*/)
}

// void doSetTextCursor(const class QTextCursor &)
//export callback_ZN9QTextEdit15doSetTextCursorERK11QTextCursor
func callback_ZN9QTextEdit15doSetTextCursorERK11QTextCursor(cthis unsafe.Pointer, cursor unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTextEdit.doSetTextCursor")
	rvx := ffiqt.CallbackAllInherits(cthis, "doSetTextCursor", cursor)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit15doSetTextCursorERK11QTextCursor", C.callback_ZN9QTextEdit15doSetTextCursorERK11QTextCursor /*nil*/)
}

// void zoomInF(float)
//export callback_ZN9QTextEdit7zoomInFEf
func callback_ZN9QTextEdit7zoomInFEf(cthis unsafe.Pointer, range_ C.float) {
	// log.Println(cthis, "QTextEdit.zoomInF")
	rvx := ffiqt.CallbackAllInherits(cthis, "zoomInF", range_)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTextEdit7zoomInFEf", C.callback_ZN9QTextEdit7zoomInFEf /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN14QPlainTextEdit5eventEP6QEvent
func callback_ZN14QPlainTextEdit5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit5eventEP6QEvent", C.callback_ZN14QPlainTextEdit5eventEP6QEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN14QPlainTextEdit10timerEventEP11QTimerEvent
func callback_ZN14QPlainTextEdit10timerEventEP11QTimerEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit10timerEventEP11QTimerEvent", C.callback_ZN14QPlainTextEdit10timerEventEP11QTimerEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent
func callback_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent", C.callback_ZN14QPlainTextEdit13keyPressEventEP9QKeyEvent /*nil*/)
}

// void keyReleaseEvent(class QKeyEvent *)
//export callback_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent
func callback_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.keyReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent", C.callback_ZN14QPlainTextEdit15keyReleaseEventEP9QKeyEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent
func callback_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent", C.callback_ZN14QPlainTextEdit11resizeEventEP12QResizeEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN14QPlainTextEdit10paintEventEP11QPaintEvent
func callback_ZN14QPlainTextEdit10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit10paintEventEP11QPaintEvent", C.callback_ZN14QPlainTextEdit10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent
func callback_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent", C.callback_ZN14QPlainTextEdit15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent
func callback_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent", C.callback_ZN14QPlainTextEdit14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent
func callback_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent", C.callback_ZN14QPlainTextEdit17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN14QPlainTextEdit21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN14QPlainTextEdit18focusNextPrevChildEb
func callback_ZN14QPlainTextEdit18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QPlainTextEdit.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit18focusNextPrevChildEb", C.callback_ZN14QPlainTextEdit18focusNextPrevChildEb /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent
func callback_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent", C.callback_ZN14QPlainTextEdit16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void dragEnterEvent(class QDragEnterEvent *)
//export callback_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent
func callback_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.dragEnterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragEnterEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent", C.callback_ZN14QPlainTextEdit14dragEnterEventEP15QDragEnterEvent /*nil*/)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
//export callback_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent
func callback_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.dragLeaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragLeaveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent", C.callback_ZN14QPlainTextEdit14dragLeaveEventEP15QDragLeaveEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent
func callback_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent", C.callback_ZN14QPlainTextEdit13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN14QPlainTextEdit9dropEventEP10QDropEvent
func callback_ZN14QPlainTextEdit9dropEventEP10QDropEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit9dropEventEP10QDropEvent", C.callback_ZN14QPlainTextEdit9dropEventEP10QDropEvent /*nil*/)
}

// void focusInEvent(class QFocusEvent *)
//export callback_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent
func callback_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.focusInEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusInEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent", C.callback_ZN14QPlainTextEdit12focusInEventEP11QFocusEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent
func callback_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent", C.callback_ZN14QPlainTextEdit13focusOutEventEP11QFocusEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN14QPlainTextEdit9showEventEP10QShowEvent
func callback_ZN14QPlainTextEdit9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit9showEventEP10QShowEvent", C.callback_ZN14QPlainTextEdit9showEventEP10QShowEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN14QPlainTextEdit11changeEventEP6QEvent
func callback_ZN14QPlainTextEdit11changeEventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit11changeEventEP6QEvent", C.callback_ZN14QPlainTextEdit11changeEventEP6QEvent /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent
func callback_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent", C.callback_ZN14QPlainTextEdit10wheelEventEP11QWheelEvent /*nil*/)
}

// QMimeData * createMimeDataFromSelection()
//export callback_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv
func callback_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPlainTextEdit.createMimeDataFromSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "createMimeDataFromSelection")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv", C.callback_ZNK14QPlainTextEdit27createMimeDataFromSelectionEv /*nil*/)
}

// bool canInsertFromMimeData(const class QMimeData *)
//export callback_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData
func callback_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData(cthis unsafe.Pointer, source unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.canInsertFromMimeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "canInsertFromMimeData", source)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData", C.callback_ZNK14QPlainTextEdit21canInsertFromMimeDataEPK9QMimeData /*nil*/)
}

// void insertFromMimeData(const class QMimeData *)
//export callback_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData
func callback_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData(cthis unsafe.Pointer, source unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.insertFromMimeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "insertFromMimeData", source)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData", C.callback_ZN14QPlainTextEdit18insertFromMimeDataEPK9QMimeData /*nil*/)
}

// void inputMethodEvent(class QInputMethodEvent *)
//export callback_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent
func callback_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPlainTextEdit.inputMethodEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "inputMethodEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent", C.callback_ZN14QPlainTextEdit16inputMethodEventEP17QInputMethodEvent /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN14QPlainTextEdit16scrollContentsByEii
func callback_ZN14QPlainTextEdit16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QPlainTextEdit.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit16scrollContentsByEii", C.callback_ZN14QPlainTextEdit16scrollContentsByEii /*nil*/)
}

// void doSetTextCursor(const class QTextCursor &)
//export callback_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor
func callback_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor(cthis unsafe.Pointer, cursor unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QPlainTextEdit.doSetTextCursor")
	rvx := ffiqt.CallbackAllInherits(cthis, "doSetTextCursor", cursor)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor", C.callback_ZN14QPlainTextEdit15doSetTextCursorERK11QTextCursor /*nil*/)
}

// QTextBlock firstVisibleBlock()
//export callback_ZNK14QPlainTextEdit17firstVisibleBlockEv
func callback_ZNK14QPlainTextEdit17firstVisibleBlockEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPlainTextEdit.firstVisibleBlock")
	rvx := ffiqt.CallbackAllInherits(cthis, "firstVisibleBlock")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK14QPlainTextEdit17firstVisibleBlockEv", C.callback_ZNK14QPlainTextEdit17firstVisibleBlockEv /*nil*/)
}

// QPointF contentOffset()
//export callback_ZNK14QPlainTextEdit13contentOffsetEv
func callback_ZNK14QPlainTextEdit13contentOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPlainTextEdit.contentOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "contentOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK14QPlainTextEdit13contentOffsetEv", C.callback_ZNK14QPlainTextEdit13contentOffsetEv /*nil*/)
}

// QRectF blockBoundingRect(const class QTextBlock &)
//export callback_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock
func callback_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock(cthis unsafe.Pointer, block unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QPlainTextEdit.blockBoundingRect")
	rvx := ffiqt.CallbackAllInherits(cthis, "blockBoundingRect", block)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock", C.callback_ZNK14QPlainTextEdit17blockBoundingRectERK10QTextBlock /*nil*/)
}

// QRectF blockBoundingGeometry(const class QTextBlock &)
//export callback_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock
func callback_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock(cthis unsafe.Pointer, block unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QPlainTextEdit.blockBoundingGeometry")
	rvx := ffiqt.CallbackAllInherits(cthis, "blockBoundingGeometry", block)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock", C.callback_ZNK14QPlainTextEdit21blockBoundingGeometryERK10QTextBlock /*nil*/)
}

// QAbstractTextDocumentLayout::PaintContext getPaintContext()
//export callback_ZNK14QPlainTextEdit15getPaintContextEv
func callback_ZNK14QPlainTextEdit15getPaintContextEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QPlainTextEdit.getPaintContext")
	rvx := ffiqt.CallbackAllInherits(cthis, "getPaintContext")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK14QPlainTextEdit15getPaintContextEv", C.callback_ZNK14QPlainTextEdit15getPaintContextEv /*nil*/)
}

// void zoomInF(float)
//export callback_ZN14QPlainTextEdit7zoomInFEf
func callback_ZN14QPlainTextEdit7zoomInFEf(cthis unsafe.Pointer, range_ C.float) {
	// log.Println(cthis, "QPlainTextEdit.zoomInF")
	rvx := ffiqt.CallbackAllInherits(cthis, "zoomInF", range_)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QPlainTextEdit7zoomInFEf", C.callback_ZN14QPlainTextEdit7zoomInFEf /*nil*/)
}

// void documentChanged(int, int, int)
//export callback_ZN24QPlainTextDocumentLayout15documentChangedEiii
func callback_ZN24QPlainTextDocumentLayout15documentChangedEiii(cthis unsafe.Pointer, from C.int, arg1 C.int, charsAdded C.int) {
	// log.Println(cthis, "QPlainTextDocumentLayout.documentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "documentChanged", from, arg1, charsAdded)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN24QPlainTextDocumentLayout15documentChangedEiii", C.callback_ZN24QPlainTextDocumentLayout15documentChangedEiii /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN12QProgressBar5eventEP6QEvent
func callback_ZN12QProgressBar5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProgressBar.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QProgressBar5eventEP6QEvent", C.callback_ZN12QProgressBar5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN12QProgressBar10paintEventEP11QPaintEvent
func callback_ZN12QProgressBar10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProgressBar.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QProgressBar10paintEventEP11QPaintEvent", C.callback_ZN12QProgressBar10paintEventEP11QPaintEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionProgressBar *)
//export callback_ZNK12QProgressBar15initStyleOptionEP23QStyleOptionProgressBar
func callback_ZNK12QProgressBar15initStyleOptionEP23QStyleOptionProgressBar(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProgressBar.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QProgressBar15initStyleOptionEP23QStyleOptionProgressBar", C.callback_ZNK12QProgressBar15initStyleOptionEP23QStyleOptionProgressBar /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN15QProgressDialog11resizeEventEP12QResizeEvent
func callback_ZN15QProgressDialog11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProgressDialog.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QProgressDialog11resizeEventEP12QResizeEvent", C.callback_ZN15QProgressDialog11resizeEventEP12QResizeEvent /*nil*/)
}

// void closeEvent(class QCloseEvent *)
//export callback_ZN15QProgressDialog10closeEventEP11QCloseEvent
func callback_ZN15QProgressDialog10closeEventEP11QCloseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProgressDialog.closeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "closeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QProgressDialog10closeEventEP11QCloseEvent", C.callback_ZN15QProgressDialog10closeEventEP11QCloseEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN15QProgressDialog11changeEventEP6QEvent
func callback_ZN15QProgressDialog11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProgressDialog.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QProgressDialog11changeEventEP6QEvent", C.callback_ZN15QProgressDialog11changeEventEP6QEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN15QProgressDialog9showEventEP10QShowEvent
func callback_ZN15QProgressDialog9showEventEP10QShowEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProgressDialog.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QProgressDialog9showEventEP10QShowEvent", C.callback_ZN15QProgressDialog9showEventEP10QShowEvent /*nil*/)
}

// void forceShow()
//export callback_ZN15QProgressDialog9forceShowEv
func callback_ZN15QProgressDialog9forceShowEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QProgressDialog.forceShow")
	rvx := ffiqt.CallbackAllInherits(cthis, "forceShow")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QProgressDialog9forceShowEv", C.callback_ZN15QProgressDialog9forceShowEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QProxyStyle5eventEP6QEvent
func callback_ZN11QProxyStyle5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QProxyStyle.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QProxyStyle5eventEP6QEvent", C.callback_ZN11QProxyStyle5eventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN12QRadioButton5eventEP6QEvent
func callback_ZN12QRadioButton5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRadioButton.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QRadioButton5eventEP6QEvent", C.callback_ZN12QRadioButton5eventEP6QEvent /*nil*/)
}

// bool hitButton(const class QPoint &)
//export callback_ZNK12QRadioButton9hitButtonERK6QPoint
func callback_ZNK12QRadioButton9hitButtonERK6QPoint(cthis unsafe.Pointer, arg0 unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QRadioButton.hitButton")
	rvx := ffiqt.CallbackAllInherits(cthis, "hitButton", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QRadioButton9hitButtonERK6QPoint", C.callback_ZNK12QRadioButton9hitButtonERK6QPoint /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN12QRadioButton10paintEventEP11QPaintEvent
func callback_ZN12QRadioButton10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRadioButton.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QRadioButton10paintEventEP11QPaintEvent", C.callback_ZN12QRadioButton10paintEventEP11QPaintEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent
func callback_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRadioButton.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent", C.callback_ZN12QRadioButton14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionButton *)
//export callback_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton
func callback_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton(cthis unsafe.Pointer, button unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QRadioButton.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", button)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton", C.callback_ZNK12QRadioButton15initStyleOptionEP18QStyleOptionButton /*nil*/)
}

// void wheelEvent(class QWheelEvent *)
//export callback_ZN10QScrollBar10wheelEventEP11QWheelEvent
func callback_ZN10QScrollBar10wheelEventEP11QWheelEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.wheelEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "wheelEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar10wheelEventEP11QWheelEvent", C.callback_ZN10QScrollBar10wheelEventEP11QWheelEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN10QScrollBar10paintEventEP11QPaintEvent
func callback_ZN10QScrollBar10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar10paintEventEP11QPaintEvent", C.callback_ZN10QScrollBar10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN10QScrollBar15mousePressEventEP11QMouseEvent
func callback_ZN10QScrollBar15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar15mousePressEventEP11QMouseEvent", C.callback_ZN10QScrollBar15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent
func callback_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent", C.callback_ZN10QScrollBar17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent
func callback_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent", C.callback_ZN10QScrollBar14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN10QScrollBar9hideEventEP10QHideEvent
func callback_ZN10QScrollBar9hideEventEP10QHideEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar9hideEventEP10QHideEvent", C.callback_ZN10QScrollBar9hideEventEP10QHideEvent /*nil*/)
}

// void sliderChange(enum QAbstractSlider::SliderChange)
//export callback_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE
func callback_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE(cthis unsafe.Pointer, change C.int) {
	// log.Println(cthis, "QScrollBar.sliderChange")
	rvx := ffiqt.CallbackAllInherits(cthis, "sliderChange", change)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE", C.callback_ZN10QScrollBar12sliderChangeEN15QAbstractSlider12SliderChangeE /*nil*/)
}

// void contextMenuEvent(class QContextMenuEvent *)
//export callback_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent
func callback_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.contextMenuEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "contextMenuEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent", C.callback_ZN10QScrollBar16contextMenuEventEP17QContextMenuEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionSlider *)
//export callback_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider
func callback_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QScrollBar.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider", C.callback_ZNK10QScrollBar15initStyleOptionEP18QStyleOptionSlider /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QShortcut5eventEP6QEvent
func callback_ZN9QShortcut5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QShortcut.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QShortcut5eventEP6QEvent", C.callback_ZN9QShortcut5eventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QSizeGrip10paintEventEP11QPaintEvent
func callback_ZN9QSizeGrip10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip10paintEventEP11QPaintEvent", C.callback_ZN9QSizeGrip10paintEventEP11QPaintEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN9QSizeGrip15mousePressEventEP11QMouseEvent
func callback_ZN9QSizeGrip15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip15mousePressEventEP11QMouseEvent", C.callback_ZN9QSizeGrip15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN9QSizeGrip14mouseMoveEventEP11QMouseEvent
func callback_ZN9QSizeGrip14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip14mouseMoveEventEP11QMouseEvent", C.callback_ZN9QSizeGrip14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN9QSizeGrip17mouseReleaseEventEP11QMouseEvent
func callback_ZN9QSizeGrip17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, mouseEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", mouseEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip17mouseReleaseEventEP11QMouseEvent", C.callback_ZN9QSizeGrip17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void moveEvent(class QMoveEvent *)
//export callback_ZN9QSizeGrip9moveEventEP10QMoveEvent
func callback_ZN9QSizeGrip9moveEventEP10QMoveEvent(cthis unsafe.Pointer, moveEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.moveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveEvent", moveEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip9moveEventEP10QMoveEvent", C.callback_ZN9QSizeGrip9moveEventEP10QMoveEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN9QSizeGrip9showEventEP10QShowEvent
func callback_ZN9QSizeGrip9showEventEP10QShowEvent(cthis unsafe.Pointer, showEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", showEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip9showEventEP10QShowEvent", C.callback_ZN9QSizeGrip9showEventEP10QShowEvent /*nil*/)
}

// void hideEvent(class QHideEvent *)
//export callback_ZN9QSizeGrip9hideEventEP10QHideEvent
func callback_ZN9QSizeGrip9hideEventEP10QHideEvent(cthis unsafe.Pointer, hideEvent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.hideEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideEvent", hideEvent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip9hideEventEP10QHideEvent", C.callback_ZN9QSizeGrip9hideEventEP10QHideEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN9QSizeGrip11eventFilterEP7QObjectP6QEvent
func callback_ZN9QSizeGrip11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip11eventFilterEP7QObjectP6QEvent", C.callback_ZN9QSizeGrip11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QSizeGrip5eventEP6QEvent
func callback_ZN9QSizeGrip5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSizeGrip.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSizeGrip5eventEP6QEvent", C.callback_ZN9QSizeGrip5eventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN8QSpinBox5eventEP6QEvent
func callback_ZN8QSpinBox5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSpinBox.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QSpinBox5eventEP6QEvent", C.callback_ZN8QSpinBox5eventEP6QEvent /*nil*/)
}

// QValidator::State validate(class QString &, int &)
//export callback_ZNK8QSpinBox8validateER7QStringRi
func callback_ZNK8QSpinBox8validateER7QStringRi(cthis unsafe.Pointer, input unsafe.Pointer /*555*/, pos unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSpinBox.validate")
	rvx := ffiqt.CallbackAllInherits(cthis, "validate", input, pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QSpinBox8validateER7QStringRi", C.callback_ZNK8QSpinBox8validateER7QStringRi /*nil*/)
}

// int valueFromText(const class QString &)
//export callback_ZNK8QSpinBox13valueFromTextERK7QString
func callback_ZNK8QSpinBox13valueFromTextERK7QString(cthis unsafe.Pointer, text unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSpinBox.valueFromText")
	rvx := ffiqt.CallbackAllInherits(cthis, "valueFromText", text)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QSpinBox13valueFromTextERK7QString", C.callback_ZNK8QSpinBox13valueFromTextERK7QString /*nil*/)
}

// QString textFromValue(int)
//export callback_ZNK8QSpinBox13textFromValueEi
func callback_ZNK8QSpinBox13textFromValueEi(cthis unsafe.Pointer, val C.int) {
	// log.Println(cthis, "QSpinBox.textFromValue")
	rvx := ffiqt.CallbackAllInherits(cthis, "textFromValue", val)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QSpinBox13textFromValueEi", C.callback_ZNK8QSpinBox13textFromValueEi /*nil*/)
}

// void fixup(class QString &)
//export callback_ZNK8QSpinBox5fixupER7QString
func callback_ZNK8QSpinBox5fixupER7QString(cthis unsafe.Pointer, str unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSpinBox.fixup")
	rvx := ffiqt.CallbackAllInherits(cthis, "fixup", str)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QSpinBox5fixupER7QString", C.callback_ZNK8QSpinBox5fixupER7QString /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN13QSplashScreen5eventEP6QEvent
func callback_ZN13QSplashScreen5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplashScreen.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QSplashScreen5eventEP6QEvent", C.callback_ZN13QSplashScreen5eventEP6QEvent /*nil*/)
}

// void drawContents(class QPainter *)
//export callback_ZN13QSplashScreen12drawContentsEP8QPainter
func callback_ZN13QSplashScreen12drawContentsEP8QPainter(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplashScreen.drawContents")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawContents", painter)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QSplashScreen12drawContentsEP8QPainter", C.callback_ZN13QSplashScreen12drawContentsEP8QPainter /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN13QSplashScreen15mousePressEventEP11QMouseEvent
func callback_ZN13QSplashScreen15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplashScreen.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QSplashScreen15mousePressEventEP11QMouseEvent", C.callback_ZN13QSplashScreen15mousePressEventEP11QMouseEvent /*nil*/)
}

// QSplitterHandle * createHandle()
//export callback_ZN9QSplitter12createHandleEv
func callback_ZN9QSplitter12createHandleEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QSplitter.createHandle")
	rvx := ffiqt.CallbackAllInherits(cthis, "createHandle")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter12createHandleEv", C.callback_ZN9QSplitter12createHandleEv /*nil*/)
}

// void childEvent(class QChildEvent *)
//export callback_ZN9QSplitter10childEventEP11QChildEvent
func callback_ZN9QSplitter10childEventEP11QChildEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitter.childEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "childEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter10childEventEP11QChildEvent", C.callback_ZN9QSplitter10childEventEP11QChildEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QSplitter5eventEP6QEvent
func callback_ZN9QSplitter5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitter.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter5eventEP6QEvent", C.callback_ZN9QSplitter5eventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN9QSplitter11resizeEventEP12QResizeEvent
func callback_ZN9QSplitter11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitter.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter11resizeEventEP12QResizeEvent", C.callback_ZN9QSplitter11resizeEventEP12QResizeEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN9QSplitter11changeEventEP6QEvent
func callback_ZN9QSplitter11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitter.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter11changeEventEP6QEvent", C.callback_ZN9QSplitter11changeEventEP6QEvent /*nil*/)
}

// void moveSplitter(int, int)
//export callback_ZN9QSplitter12moveSplitterEii
func callback_ZN9QSplitter12moveSplitterEii(cthis unsafe.Pointer, pos C.int, index C.int) {
	// log.Println(cthis, "QSplitter.moveSplitter")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveSplitter", pos, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter12moveSplitterEii", C.callback_ZN9QSplitter12moveSplitterEii /*nil*/)
}

// void setRubberBand(int)
//export callback_ZN9QSplitter13setRubberBandEi
func callback_ZN9QSplitter13setRubberBandEi(cthis unsafe.Pointer, position C.int) {
	// log.Println(cthis, "QSplitter.setRubberBand")
	rvx := ffiqt.CallbackAllInherits(cthis, "setRubberBand", position)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter13setRubberBandEi", C.callback_ZN9QSplitter13setRubberBandEi /*nil*/)
}

// int closestLegalPosition(int, int)
//export callback_ZN9QSplitter20closestLegalPositionEii
func callback_ZN9QSplitter20closestLegalPositionEii(cthis unsafe.Pointer, arg0 C.int, arg1 C.int) {
	// log.Println(cthis, "QSplitter.closestLegalPosition")
	rvx := ffiqt.CallbackAllInherits(cthis, "closestLegalPosition", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSplitter20closestLegalPositionEii", C.callback_ZN9QSplitter20closestLegalPositionEii /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN15QSplitterHandle10paintEventEP11QPaintEvent
func callback_ZN15QSplitterHandle10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitterHandle.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle10paintEventEP11QPaintEvent", C.callback_ZN15QSplitterHandle10paintEventEP11QPaintEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent
func callback_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitterHandle.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent", C.callback_ZN15QSplitterHandle14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent
func callback_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitterHandle.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent", C.callback_ZN15QSplitterHandle15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent
func callback_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitterHandle.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent", C.callback_ZN15QSplitterHandle17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN15QSplitterHandle11resizeEventEP12QResizeEvent
func callback_ZN15QSplitterHandle11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitterHandle.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle11resizeEventEP12QResizeEvent", C.callback_ZN15QSplitterHandle11resizeEventEP12QResizeEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QSplitterHandle5eventEP6QEvent
func callback_ZN15QSplitterHandle5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSplitterHandle.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle5eventEP6QEvent", C.callback_ZN15QSplitterHandle5eventEP6QEvent /*nil*/)
}

// void moveSplitter(int)
//export callback_ZN15QSplitterHandle12moveSplitterEi
func callback_ZN15QSplitterHandle12moveSplitterEi(cthis unsafe.Pointer, p C.int) {
	// log.Println(cthis, "QSplitterHandle.moveSplitter")
	rvx := ffiqt.CallbackAllInherits(cthis, "moveSplitter", p)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle12moveSplitterEi", C.callback_ZN15QSplitterHandle12moveSplitterEi /*nil*/)
}

// int closestLegalPosition(int)
//export callback_ZN15QSplitterHandle20closestLegalPositionEi
func callback_ZN15QSplitterHandle20closestLegalPositionEi(cthis unsafe.Pointer, p C.int) {
	// log.Println(cthis, "QSplitterHandle.closestLegalPosition")
	rvx := ffiqt.CallbackAllInherits(cthis, "closestLegalPosition", p)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSplitterHandle20closestLegalPositionEi", C.callback_ZN15QSplitterHandle20closestLegalPositionEi /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN14QStackedWidget5eventEP6QEvent
func callback_ZN14QStackedWidget5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStackedWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QStackedWidget5eventEP6QEvent", C.callback_ZN14QStackedWidget5eventEP6QEvent /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN10QStatusBar9showEventEP10QShowEvent
func callback_ZN10QStatusBar9showEventEP10QShowEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStatusBar.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QStatusBar9showEventEP10QShowEvent", C.callback_ZN10QStatusBar9showEventEP10QShowEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN10QStatusBar10paintEventEP11QPaintEvent
func callback_ZN10QStatusBar10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStatusBar.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QStatusBar10paintEventEP11QPaintEvent", C.callback_ZN10QStatusBar10paintEventEP11QPaintEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN10QStatusBar11resizeEventEP12QResizeEvent
func callback_ZN10QStatusBar11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStatusBar.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QStatusBar11resizeEventEP12QResizeEvent", C.callback_ZN10QStatusBar11resizeEventEP12QResizeEvent /*nil*/)
}

// void reformat()
//export callback_ZN10QStatusBar8reformatEv
func callback_ZN10QStatusBar8reformatEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QStatusBar.reformat")
	rvx := ffiqt.CallbackAllInherits(cthis, "reformat")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QStatusBar8reformatEv", C.callback_ZN10QStatusBar8reformatEv /*nil*/)
}

// void hideOrShow()
//export callback_ZN10QStatusBar10hideOrShowEv
func callback_ZN10QStatusBar10hideOrShowEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QStatusBar.hideOrShow")
	rvx := ffiqt.CallbackAllInherits(cthis, "hideOrShow")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QStatusBar10hideOrShowEv", C.callback_ZN10QStatusBar10hideOrShowEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN10QStatusBar5eventEP6QEvent
func callback_ZN10QStatusBar5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStatusBar.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QStatusBar5eventEP6QEvent", C.callback_ZN10QStatusBar5eventEP6QEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionViewItem *, const class QModelIndex &)
//export callback_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex
func callback_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex(cthis unsafe.Pointer, option unsafe.Pointer /*666*/, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QStyledItemDelegate.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex", C.callback_ZNK19QStyledItemDelegate15initStyleOptionEP20QStyleOptionViewItemRK11QModelIndex /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent
func callback_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, object unsafe.Pointer /*666*/, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStyledItemDelegate.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", object, event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent", C.callback_ZN19QStyledItemDelegate11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// bool editorEvent(class QEvent *, class QAbstractItemModel *, const class QStyleOptionViewItem &, const class QModelIndex &)
//export callback_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex
func callback_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex(cthis unsafe.Pointer, event unsafe.Pointer /*666*/, model unsafe.Pointer /*666*/, option unsafe.Pointer /*555*/, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QStyledItemDelegate.editorEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "editorEvent", event, model, option, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex", C.callback_ZN19QStyledItemDelegate11editorEventEP6QEventP18QAbstractItemModelRK20QStyleOptionViewItemRK11QModelIndex /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QSystemTrayIcon5eventEP6QEvent
func callback_ZN15QSystemTrayIcon5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSystemTrayIcon.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSystemTrayIcon5eventEP6QEvent", C.callback_ZN15QSystemTrayIcon5eventEP6QEvent /*nil*/)
}

// void rowMoved(int, int, int)
//export callback_ZN10QTableView8rowMovedEiii
func callback_ZN10QTableView8rowMovedEiii(cthis unsafe.Pointer, row C.int, oldIndex C.int, newIndex C.int) {
	// log.Println(cthis, "QTableView.rowMoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowMoved", row, oldIndex, newIndex)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView8rowMovedEiii", C.callback_ZN10QTableView8rowMovedEiii /*nil*/)
}

// void columnMoved(int, int, int)
//export callback_ZN10QTableView11columnMovedEiii
func callback_ZN10QTableView11columnMovedEiii(cthis unsafe.Pointer, column C.int, oldIndex C.int, newIndex C.int) {
	// log.Println(cthis, "QTableView.columnMoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "columnMoved", column, oldIndex, newIndex)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView11columnMovedEiii", C.callback_ZN10QTableView11columnMovedEiii /*nil*/)
}

// void rowResized(int, int, int)
//export callback_ZN10QTableView10rowResizedEiii
func callback_ZN10QTableView10rowResizedEiii(cthis unsafe.Pointer, row C.int, oldHeight C.int, newHeight C.int) {
	// log.Println(cthis, "QTableView.rowResized")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowResized", row, oldHeight, newHeight)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView10rowResizedEiii", C.callback_ZN10QTableView10rowResizedEiii /*nil*/)
}

// void columnResized(int, int, int)
//export callback_ZN10QTableView13columnResizedEiii
func callback_ZN10QTableView13columnResizedEiii(cthis unsafe.Pointer, column C.int, oldWidth C.int, newWidth C.int) {
	// log.Println(cthis, "QTableView.columnResized")
	rvx := ffiqt.CallbackAllInherits(cthis, "columnResized", column, oldWidth, newWidth)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView13columnResizedEiii", C.callback_ZN10QTableView13columnResizedEiii /*nil*/)
}

// void rowCountChanged(int, int)
//export callback_ZN10QTableView15rowCountChangedEii
func callback_ZN10QTableView15rowCountChangedEii(cthis unsafe.Pointer, oldCount C.int, newCount C.int) {
	// log.Println(cthis, "QTableView.rowCountChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowCountChanged", oldCount, newCount)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView15rowCountChangedEii", C.callback_ZN10QTableView15rowCountChangedEii /*nil*/)
}

// void columnCountChanged(int, int)
//export callback_ZN10QTableView18columnCountChangedEii
func callback_ZN10QTableView18columnCountChangedEii(cthis unsafe.Pointer, oldCount C.int, newCount C.int) {
	// log.Println(cthis, "QTableView.columnCountChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "columnCountChanged", oldCount, newCount)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView18columnCountChangedEii", C.callback_ZN10QTableView18columnCountChangedEii /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN10QTableView16scrollContentsByEii
func callback_ZN10QTableView16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QTableView.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView16scrollContentsByEii", C.callback_ZN10QTableView16scrollContentsByEii /*nil*/)
}

// QStyleOptionViewItem viewOptions()
//export callback_ZNK10QTableView11viewOptionsEv
func callback_ZNK10QTableView11viewOptionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTableView.viewOptions")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewOptions")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView11viewOptionsEv", C.callback_ZNK10QTableView11viewOptionsEv /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN10QTableView10paintEventEP11QPaintEvent
func callback_ZN10QTableView10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTableView.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView10paintEventEP11QPaintEvent", C.callback_ZN10QTableView10paintEventEP11QPaintEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN10QTableView10timerEventEP11QTimerEvent
func callback_ZN10QTableView10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTableView.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView10timerEventEP11QTimerEvent", C.callback_ZN10QTableView10timerEventEP11QTimerEvent /*nil*/)
}

// int horizontalOffset()
//export callback_ZNK10QTableView16horizontalOffsetEv
func callback_ZNK10QTableView16horizontalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTableView.horizontalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView16horizontalOffsetEv", C.callback_ZNK10QTableView16horizontalOffsetEv /*nil*/)
}

// int verticalOffset()
//export callback_ZNK10QTableView14verticalOffsetEv
func callback_ZNK10QTableView14verticalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTableView.verticalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView14verticalOffsetEv", C.callback_ZNK10QTableView14verticalOffsetEv /*nil*/)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
//export callback_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE
func callback_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(cthis unsafe.Pointer, rect unsafe.Pointer /*555*/, command C.int) {
	// log.Println(cthis, "QTableView.setSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "setSelection", rect, command)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", C.callback_ZN10QTableView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE /*nil*/)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
//export callback_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection
func callback_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection(cthis unsafe.Pointer, selection unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTableView.visualRegionForSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "visualRegionForSelection", selection)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection", C.callback_ZNK10QTableView24visualRegionForSelectionERK14QItemSelection /*nil*/)
}

// void updateGeometries()
//export callback_ZN10QTableView16updateGeometriesEv
func callback_ZN10QTableView16updateGeometriesEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTableView.updateGeometries")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateGeometries")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView16updateGeometriesEv", C.callback_ZN10QTableView16updateGeometriesEv /*nil*/)
}

// QSize viewportSizeHint()
//export callback_ZNK10QTableView16viewportSizeHintEv
func callback_ZNK10QTableView16viewportSizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTableView.viewportSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportSizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView16viewportSizeHintEv", C.callback_ZNK10QTableView16viewportSizeHintEv /*nil*/)
}

// int sizeHintForRow(int)
//export callback_ZNK10QTableView14sizeHintForRowEi
func callback_ZNK10QTableView14sizeHintForRowEi(cthis unsafe.Pointer, row C.int) {
	// log.Println(cthis, "QTableView.sizeHintForRow")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHintForRow", row)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView14sizeHintForRowEi", C.callback_ZNK10QTableView14sizeHintForRowEi /*nil*/)
}

// int sizeHintForColumn(int)
//export callback_ZNK10QTableView17sizeHintForColumnEi
func callback_ZNK10QTableView17sizeHintForColumnEi(cthis unsafe.Pointer, column C.int) {
	// log.Println(cthis, "QTableView.sizeHintForColumn")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHintForColumn", column)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView17sizeHintForColumnEi", C.callback_ZNK10QTableView17sizeHintForColumnEi /*nil*/)
}

// void verticalScrollbarAction(int)
//export callback_ZN10QTableView23verticalScrollbarActionEi
func callback_ZN10QTableView23verticalScrollbarActionEi(cthis unsafe.Pointer, action C.int) {
	// log.Println(cthis, "QTableView.verticalScrollbarAction")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalScrollbarAction", action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView23verticalScrollbarActionEi", C.callback_ZN10QTableView23verticalScrollbarActionEi /*nil*/)
}

// void horizontalScrollbarAction(int)
//export callback_ZN10QTableView25horizontalScrollbarActionEi
func callback_ZN10QTableView25horizontalScrollbarActionEi(cthis unsafe.Pointer, action C.int) {
	// log.Println(cthis, "QTableView.horizontalScrollbarAction")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalScrollbarAction", action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView25horizontalScrollbarActionEi", C.callback_ZN10QTableView25horizontalScrollbarActionEi /*nil*/)
}

// bool isIndexHidden(const class QModelIndex &)
//export callback_ZNK10QTableView13isIndexHiddenERK11QModelIndex
func callback_ZNK10QTableView13isIndexHiddenERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTableView.isIndexHidden")
	rvx := ffiqt.CallbackAllInherits(cthis, "isIndexHidden", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK10QTableView13isIndexHiddenERK11QModelIndex", C.callback_ZNK10QTableView13isIndexHiddenERK11QModelIndex /*nil*/)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
//export callback_ZN10QTableView16selectionChangedERK14QItemSelectionS2_
func callback_ZN10QTableView16selectionChangedERK14QItemSelectionS2_(cthis unsafe.Pointer, selected unsafe.Pointer /*555*/, deselected unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTableView.selectionChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "selectionChanged", selected, deselected)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView16selectionChangedERK14QItemSelectionS2_", C.callback_ZN10QTableView16selectionChangedERK14QItemSelectionS2_ /*nil*/)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
//export callback_ZN10QTableView14currentChangedERK11QModelIndexS2_
func callback_ZN10QTableView14currentChangedERK11QModelIndexS2_(cthis unsafe.Pointer, current unsafe.Pointer /*555*/, previous unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTableView.currentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentChanged", current, previous)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN10QTableView14currentChangedERK11QModelIndexS2_", C.callback_ZN10QTableView14currentChangedERK11QModelIndexS2_ /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN12QTableWidget5eventEP6QEvent
func callback_ZN12QTableWidget5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTableWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTableWidget5eventEP6QEvent", C.callback_ZN12QTableWidget5eventEP6QEvent /*nil*/)
}

// bool dropMimeData(int, int, const class QMimeData *, Qt::DropAction)
//export callback_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE
func callback_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE(cthis unsafe.Pointer, row C.int, column C.int, data unsafe.Pointer /*666*/, action C.int) {
	// log.Println(cthis, "QTableWidget.dropMimeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropMimeData", row, column, data, action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE", C.callback_ZN12QTableWidget12dropMimeDataEiiPK9QMimeDataN2Qt10DropActionE /*nil*/)
}

// Qt::DropActions supportedDropActions()
//export callback_ZNK12QTableWidget20supportedDropActionsEv
func callback_ZNK12QTableWidget20supportedDropActionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTableWidget.supportedDropActions")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportedDropActions")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QTableWidget20supportedDropActionsEv", C.callback_ZNK12QTableWidget20supportedDropActionsEv /*nil*/)
}

// QModelIndex indexFromItem(class QTableWidgetItem *)
//export callback_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem
func callback_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem(cthis unsafe.Pointer, item unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTableWidget.indexFromItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "indexFromItem", item)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem", C.callback_ZNK12QTableWidget13indexFromItemEP16QTableWidgetItem /*nil*/)
}

// QTableWidgetItem * itemFromIndex(const class QModelIndex &)
//export callback_ZNK12QTableWidget13itemFromIndexERK11QModelIndex
func callback_ZNK12QTableWidget13itemFromIndexERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTableWidget.itemFromIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemFromIndex", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK12QTableWidget13itemFromIndexERK11QModelIndex", C.callback_ZNK12QTableWidget13itemFromIndexERK11QModelIndex /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN12QTableWidget9dropEventEP10QDropEvent
func callback_ZN12QTableWidget9dropEventEP10QDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTableWidget.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTableWidget9dropEventEP10QDropEvent", C.callback_ZN12QTableWidget9dropEventEP10QDropEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN12QTextBrowser5eventEP6QEvent
func callback_ZN12QTextBrowser5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBrowser.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser5eventEP6QEvent", C.callback_ZN12QTextBrowser5eventEP6QEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN12QTextBrowser13keyPressEventEP9QKeyEvent
func callback_ZN12QTextBrowser13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBrowser.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser13keyPressEventEP9QKeyEvent", C.callback_ZN12QTextBrowser13keyPressEventEP9QKeyEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent
func callback_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBrowser.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent", C.callback_ZN12QTextBrowser14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN12QTextBrowser15mousePressEventEP11QMouseEvent
func callback_ZN12QTextBrowser15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBrowser.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser15mousePressEventEP11QMouseEvent", C.callback_ZN12QTextBrowser15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent
func callback_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBrowser.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent", C.callback_ZN12QTextBrowser17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void focusOutEvent(class QFocusEvent *)
//export callback_ZN12QTextBrowser13focusOutEventEP11QFocusEvent
func callback_ZN12QTextBrowser13focusOutEventEP11QFocusEvent(cthis unsafe.Pointer, ev unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBrowser.focusOutEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusOutEvent", ev)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser13focusOutEventEP11QFocusEvent", C.callback_ZN12QTextBrowser13focusOutEventEP11QFocusEvent /*nil*/)
}

// bool focusNextPrevChild(_Bool)
//export callback_ZN12QTextBrowser18focusNextPrevChildEb
func callback_ZN12QTextBrowser18focusNextPrevChildEb(cthis unsafe.Pointer, next C.bool) {
	// log.Println(cthis, "QTextBrowser.focusNextPrevChild")
	rvx := ffiqt.CallbackAllInherits(cthis, "focusNextPrevChild", next)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser18focusNextPrevChildEb", C.callback_ZN12QTextBrowser18focusNextPrevChildEb /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN12QTextBrowser10paintEventEP11QPaintEvent
func callback_ZN12QTextBrowser10paintEventEP11QPaintEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTextBrowser.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN12QTextBrowser10paintEventEP11QPaintEvent", C.callback_ZN12QTextBrowser10paintEventEP11QPaintEvent /*nil*/)
}

// void actionEvent(class QActionEvent *)
//export callback_ZN8QToolBar11actionEventEP12QActionEvent
func callback_ZN8QToolBar11actionEventEP12QActionEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBar.actionEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "actionEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBar11actionEventEP12QActionEvent", C.callback_ZN8QToolBar11actionEventEP12QActionEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN8QToolBar11changeEventEP6QEvent
func callback_ZN8QToolBar11changeEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBar.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBar11changeEventEP6QEvent", C.callback_ZN8QToolBar11changeEventEP6QEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN8QToolBar10paintEventEP11QPaintEvent
func callback_ZN8QToolBar10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBar.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBar10paintEventEP11QPaintEvent", C.callback_ZN8QToolBar10paintEventEP11QPaintEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN8QToolBar5eventEP6QEvent
func callback_ZN8QToolBar5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBar.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBar5eventEP6QEvent", C.callback_ZN8QToolBar5eventEP6QEvent /*nil*/)
}

// void initStyleOption(class QStyleOptionToolBar *)
//export callback_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar
func callback_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBar.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar", C.callback_ZNK8QToolBar15initStyleOptionEP19QStyleOptionToolBar /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN8QToolBox5eventEP6QEvent
func callback_ZN8QToolBox5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBox.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBox5eventEP6QEvent", C.callback_ZN8QToolBox5eventEP6QEvent /*nil*/)
}

// void itemInserted(int)
//export callback_ZN8QToolBox12itemInsertedEi
func callback_ZN8QToolBox12itemInsertedEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QToolBox.itemInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemInserted", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBox12itemInsertedEi", C.callback_ZN8QToolBox12itemInsertedEi /*nil*/)
}

// void itemRemoved(int)
//export callback_ZN8QToolBox11itemRemovedEi
func callback_ZN8QToolBox11itemRemovedEi(cthis unsafe.Pointer, index C.int) {
	// log.Println(cthis, "QToolBox.itemRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemRemoved", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBox11itemRemovedEi", C.callback_ZN8QToolBox11itemRemovedEi /*nil*/)
}

// void showEvent(class QShowEvent *)
//export callback_ZN8QToolBox9showEventEP10QShowEvent
func callback_ZN8QToolBox9showEventEP10QShowEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBox.showEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "showEvent", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBox9showEventEP10QShowEvent", C.callback_ZN8QToolBox9showEventEP10QShowEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN8QToolBox11changeEventEP6QEvent
func callback_ZN8QToolBox11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolBox.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QToolBox11changeEventEP6QEvent", C.callback_ZN8QToolBox11changeEventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QToolButton5eventEP6QEvent
func callback_ZN11QToolButton5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton5eventEP6QEvent", C.callback_ZN11QToolButton5eventEP6QEvent /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN11QToolButton15mousePressEventEP11QMouseEvent
func callback_ZN11QToolButton15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton15mousePressEventEP11QMouseEvent", C.callback_ZN11QToolButton15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN11QToolButton17mouseReleaseEventEP11QMouseEvent
func callback_ZN11QToolButton17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton17mouseReleaseEventEP11QMouseEvent", C.callback_ZN11QToolButton17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN11QToolButton10paintEventEP11QPaintEvent
func callback_ZN11QToolButton10paintEventEP11QPaintEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton10paintEventEP11QPaintEvent", C.callback_ZN11QToolButton10paintEventEP11QPaintEvent /*nil*/)
}

// void actionEvent(class QActionEvent *)
//export callback_ZN11QToolButton11actionEventEP12QActionEvent
func callback_ZN11QToolButton11actionEventEP12QActionEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.actionEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "actionEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton11actionEventEP12QActionEvent", C.callback_ZN11QToolButton11actionEventEP12QActionEvent /*nil*/)
}

// void enterEvent(class QEvent *)
//export callback_ZN11QToolButton10enterEventEP6QEvent
func callback_ZN11QToolButton10enterEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.enterEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "enterEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton10enterEventEP6QEvent", C.callback_ZN11QToolButton10enterEventEP6QEvent /*nil*/)
}

// void leaveEvent(class QEvent *)
//export callback_ZN11QToolButton10leaveEventEP6QEvent
func callback_ZN11QToolButton10leaveEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.leaveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "leaveEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton10leaveEventEP6QEvent", C.callback_ZN11QToolButton10leaveEventEP6QEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN11QToolButton10timerEventEP11QTimerEvent
func callback_ZN11QToolButton10timerEventEP11QTimerEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton10timerEventEP11QTimerEvent", C.callback_ZN11QToolButton10timerEventEP11QTimerEvent /*nil*/)
}

// void changeEvent(class QEvent *)
//export callback_ZN11QToolButton11changeEventEP6QEvent
func callback_ZN11QToolButton11changeEventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.changeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "changeEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton11changeEventEP6QEvent", C.callback_ZN11QToolButton11changeEventEP6QEvent /*nil*/)
}

// bool hitButton(const class QPoint &)
//export callback_ZNK11QToolButton9hitButtonERK6QPoint
func callback_ZNK11QToolButton9hitButtonERK6QPoint(cthis unsafe.Pointer, pos unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QToolButton.hitButton")
	rvx := ffiqt.CallbackAllInherits(cthis, "hitButton", pos)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QToolButton9hitButtonERK6QPoint", C.callback_ZNK11QToolButton9hitButtonERK6QPoint /*nil*/)
}

// void nextCheckState()
//export callback_ZN11QToolButton14nextCheckStateEv
func callback_ZN11QToolButton14nextCheckStateEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QToolButton.nextCheckState")
	rvx := ffiqt.CallbackAllInherits(cthis, "nextCheckState")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QToolButton14nextCheckStateEv", C.callback_ZN11QToolButton14nextCheckStateEv /*nil*/)
}

// void initStyleOption(class QStyleOptionToolButton *)
//export callback_ZNK11QToolButton15initStyleOptionEP22QStyleOptionToolButton
func callback_ZNK11QToolButton15initStyleOptionEP22QStyleOptionToolButton(cthis unsafe.Pointer, option unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QToolButton.initStyleOption")
	rvx := ffiqt.CallbackAllInherits(cthis, "initStyleOption", option)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QToolButton15initStyleOptionEP22QStyleOptionToolButton", C.callback_ZNK11QToolButton15initStyleOptionEP22QStyleOptionToolButton /*nil*/)
}

// void columnResized(int, int, int)
//export callback_ZN9QTreeView13columnResizedEiii
func callback_ZN9QTreeView13columnResizedEiii(cthis unsafe.Pointer, column C.int, oldSize C.int, newSize C.int) {
	// log.Println(cthis, "QTreeView.columnResized")
	rvx := ffiqt.CallbackAllInherits(cthis, "columnResized", column, oldSize, newSize)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView13columnResizedEiii", C.callback_ZN9QTreeView13columnResizedEiii /*nil*/)
}

// void columnCountChanged(int, int)
//export callback_ZN9QTreeView18columnCountChangedEii
func callback_ZN9QTreeView18columnCountChangedEii(cthis unsafe.Pointer, oldCount C.int, newCount C.int) {
	// log.Println(cthis, "QTreeView.columnCountChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "columnCountChanged", oldCount, newCount)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView18columnCountChangedEii", C.callback_ZN9QTreeView18columnCountChangedEii /*nil*/)
}

// void columnMoved()
//export callback_ZN9QTreeView11columnMovedEv
func callback_ZN9QTreeView11columnMovedEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeView.columnMoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "columnMoved")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView11columnMovedEv", C.callback_ZN9QTreeView11columnMovedEv /*nil*/)
}

// void reexpand()
//export callback_ZN9QTreeView8reexpandEv
func callback_ZN9QTreeView8reexpandEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeView.reexpand")
	rvx := ffiqt.CallbackAllInherits(cthis, "reexpand")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView8reexpandEv", C.callback_ZN9QTreeView8reexpandEv /*nil*/)
}

// void rowsRemoved(const class QModelIndex &, int, int)
//export callback_ZN9QTreeView11rowsRemovedERK11QModelIndexii
func callback_ZN9QTreeView11rowsRemovedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, first C.int, last C.int) {
	// log.Println(cthis, "QTreeView.rowsRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsRemoved", parent, first, last)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView11rowsRemovedERK11QModelIndexii", C.callback_ZN9QTreeView11rowsRemovedERK11QModelIndexii /*nil*/)
}

// void scrollContentsBy(int, int)
//export callback_ZN9QTreeView16scrollContentsByEii
func callback_ZN9QTreeView16scrollContentsByEii(cthis unsafe.Pointer, dx C.int, dy C.int) {
	// log.Println(cthis, "QTreeView.scrollContentsBy")
	rvx := ffiqt.CallbackAllInherits(cthis, "scrollContentsBy", dx, dy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView16scrollContentsByEii", C.callback_ZN9QTreeView16scrollContentsByEii /*nil*/)
}

// void rowsInserted(const class QModelIndex &, int, int)
//export callback_ZN9QTreeView12rowsInsertedERK11QModelIndexii
func callback_ZN9QTreeView12rowsInsertedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QTreeView.rowsInserted")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsInserted", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView12rowsInsertedERK11QModelIndexii", C.callback_ZN9QTreeView12rowsInsertedERK11QModelIndexii /*nil*/)
}

// void rowsAboutToBeRemoved(const class QModelIndex &, int, int)
//export callback_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii
func callback_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, start C.int, end C.int) {
	// log.Println(cthis, "QTreeView.rowsAboutToBeRemoved")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowsAboutToBeRemoved", parent, start, end)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii", C.callback_ZN9QTreeView20rowsAboutToBeRemovedERK11QModelIndexii /*nil*/)
}

// int horizontalOffset()
//export callback_ZNK9QTreeView16horizontalOffsetEv
func callback_ZNK9QTreeView16horizontalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeView.horizontalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView16horizontalOffsetEv", C.callback_ZNK9QTreeView16horizontalOffsetEv /*nil*/)
}

// int verticalOffset()
//export callback_ZNK9QTreeView14verticalOffsetEv
func callback_ZNK9QTreeView14verticalOffsetEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeView.verticalOffset")
	rvx := ffiqt.CallbackAllInherits(cthis, "verticalOffset")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView14verticalOffsetEv", C.callback_ZNK9QTreeView14verticalOffsetEv /*nil*/)
}

// void setSelection(const class QRect &, class QItemSelectionModel::SelectionFlags)
//export callback_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE
func callback_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE(cthis unsafe.Pointer, rect unsafe.Pointer /*555*/, command C.int) {
	// log.Println(cthis, "QTreeView.setSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "setSelection", rect, command)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE", C.callback_ZN9QTreeView12setSelectionERK5QRect6QFlagsIN19QItemSelectionModel13SelectionFlagEE /*nil*/)
}

// QRegion visualRegionForSelection(const class QItemSelection &)
//export callback_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection
func callback_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection(cthis unsafe.Pointer, selection unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.visualRegionForSelection")
	rvx := ffiqt.CallbackAllInherits(cthis, "visualRegionForSelection", selection)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection", C.callback_ZNK9QTreeView24visualRegionForSelectionERK14QItemSelection /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN9QTreeView10timerEventEP11QTimerEvent
func callback_ZN9QTreeView10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView10timerEventEP11QTimerEvent", C.callback_ZN9QTreeView10timerEventEP11QTimerEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN9QTreeView10paintEventEP11QPaintEvent
func callback_ZN9QTreeView10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView10paintEventEP11QPaintEvent", C.callback_ZN9QTreeView10paintEventEP11QPaintEvent /*nil*/)
}

// void drawTree(class QPainter *, const class QRegion &)
//export callback_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion
func callback_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, region unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.drawTree")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawTree", painter, region)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion", C.callback_ZNK9QTreeView8drawTreeEP8QPainterRK7QRegion /*nil*/)
}

// void drawRow(class QPainter *, const class QStyleOptionViewItem &, const class QModelIndex &)
//export callback_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex
func callback_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, options unsafe.Pointer /*555*/, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.drawRow")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawRow", painter, options, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex", C.callback_ZNK9QTreeView7drawRowEP8QPainterRK20QStyleOptionViewItemRK11QModelIndex /*nil*/)
}

// void drawBranches(class QPainter *, const class QRect &, const class QModelIndex &)
//export callback_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex
func callback_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex(cthis unsafe.Pointer, painter unsafe.Pointer /*666*/, rect unsafe.Pointer /*555*/, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.drawBranches")
	rvx := ffiqt.CallbackAllInherits(cthis, "drawBranches", painter, rect, index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex", C.callback_ZNK9QTreeView12drawBranchesEP8QPainterRK5QRectRK11QModelIndex /*nil*/)
}

// void mousePressEvent(class QMouseEvent *)
//export callback_ZN9QTreeView15mousePressEventEP11QMouseEvent
func callback_ZN9QTreeView15mousePressEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.mousePressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mousePressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView15mousePressEventEP11QMouseEvent", C.callback_ZN9QTreeView15mousePressEventEP11QMouseEvent /*nil*/)
}

// void mouseReleaseEvent(class QMouseEvent *)
//export callback_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent
func callback_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.mouseReleaseEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseReleaseEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent", C.callback_ZN9QTreeView17mouseReleaseEventEP11QMouseEvent /*nil*/)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
//export callback_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent
func callback_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.mouseDoubleClickEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseDoubleClickEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent", C.callback_ZN9QTreeView21mouseDoubleClickEventEP11QMouseEvent /*nil*/)
}

// void mouseMoveEvent(class QMouseEvent *)
//export callback_ZN9QTreeView14mouseMoveEventEP11QMouseEvent
func callback_ZN9QTreeView14mouseMoveEventEP11QMouseEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.mouseMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "mouseMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView14mouseMoveEventEP11QMouseEvent", C.callback_ZN9QTreeView14mouseMoveEventEP11QMouseEvent /*nil*/)
}

// void keyPressEvent(class QKeyEvent *)
//export callback_ZN9QTreeView13keyPressEventEP9QKeyEvent
func callback_ZN9QTreeView13keyPressEventEP9QKeyEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.keyPressEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "keyPressEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView13keyPressEventEP9QKeyEvent", C.callback_ZN9QTreeView13keyPressEventEP9QKeyEvent /*nil*/)
}

// void dragMoveEvent(class QDragMoveEvent *)
//export callback_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent
func callback_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.dragMoveEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dragMoveEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent", C.callback_ZN9QTreeView13dragMoveEventEP14QDragMoveEvent /*nil*/)
}

// bool viewportEvent(class QEvent *)
//export callback_ZN9QTreeView13viewportEventEP6QEvent
func callback_ZN9QTreeView13viewportEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeView.viewportEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView13viewportEventEP6QEvent", C.callback_ZN9QTreeView13viewportEventEP6QEvent /*nil*/)
}

// void updateGeometries()
//export callback_ZN9QTreeView16updateGeometriesEv
func callback_ZN9QTreeView16updateGeometriesEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeView.updateGeometries")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateGeometries")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView16updateGeometriesEv", C.callback_ZN9QTreeView16updateGeometriesEv /*nil*/)
}

// QSize viewportSizeHint()
//export callback_ZNK9QTreeView16viewportSizeHintEv
func callback_ZNK9QTreeView16viewportSizeHintEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeView.viewportSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "viewportSizeHint")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView16viewportSizeHintEv", C.callback_ZNK9QTreeView16viewportSizeHintEv /*nil*/)
}

// int sizeHintForColumn(int)
//export callback_ZNK9QTreeView17sizeHintForColumnEi
func callback_ZNK9QTreeView17sizeHintForColumnEi(cthis unsafe.Pointer, column C.int) {
	// log.Println(cthis, "QTreeView.sizeHintForColumn")
	rvx := ffiqt.CallbackAllInherits(cthis, "sizeHintForColumn", column)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView17sizeHintForColumnEi", C.callback_ZNK9QTreeView17sizeHintForColumnEi /*nil*/)
}

// int indexRowSizeHint(const class QModelIndex &)
//export callback_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex
func callback_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.indexRowSizeHint")
	rvx := ffiqt.CallbackAllInherits(cthis, "indexRowSizeHint", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex", C.callback_ZNK9QTreeView16indexRowSizeHintERK11QModelIndex /*nil*/)
}

// int rowHeight(const class QModelIndex &)
//export callback_ZNK9QTreeView9rowHeightERK11QModelIndex
func callback_ZNK9QTreeView9rowHeightERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.rowHeight")
	rvx := ffiqt.CallbackAllInherits(cthis, "rowHeight", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView9rowHeightERK11QModelIndex", C.callback_ZNK9QTreeView9rowHeightERK11QModelIndex /*nil*/)
}

// void horizontalScrollbarAction(int)
//export callback_ZN9QTreeView25horizontalScrollbarActionEi
func callback_ZN9QTreeView25horizontalScrollbarActionEi(cthis unsafe.Pointer, action C.int) {
	// log.Println(cthis, "QTreeView.horizontalScrollbarAction")
	rvx := ffiqt.CallbackAllInherits(cthis, "horizontalScrollbarAction", action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView25horizontalScrollbarActionEi", C.callback_ZN9QTreeView25horizontalScrollbarActionEi /*nil*/)
}

// bool isIndexHidden(const class QModelIndex &)
//export callback_ZNK9QTreeView13isIndexHiddenERK11QModelIndex
func callback_ZNK9QTreeView13isIndexHiddenERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.isIndexHidden")
	rvx := ffiqt.CallbackAllInherits(cthis, "isIndexHidden", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QTreeView13isIndexHiddenERK11QModelIndex", C.callback_ZNK9QTreeView13isIndexHiddenERK11QModelIndex /*nil*/)
}

// void selectionChanged(const class QItemSelection &, const class QItemSelection &)
//export callback_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_
func callback_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_(cthis unsafe.Pointer, selected unsafe.Pointer /*555*/, deselected unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.selectionChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "selectionChanged", selected, deselected)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_", C.callback_ZN9QTreeView16selectionChangedERK14QItemSelectionS2_ /*nil*/)
}

// void currentChanged(const class QModelIndex &, const class QModelIndex &)
//export callback_ZN9QTreeView14currentChangedERK11QModelIndexS2_
func callback_ZN9QTreeView14currentChangedERK11QModelIndexS2_(cthis unsafe.Pointer, current unsafe.Pointer /*555*/, previous unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeView.currentChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "currentChanged", current, previous)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTreeView14currentChangedERK11QModelIndexS2_", C.callback_ZN9QTreeView14currentChangedERK11QModelIndexS2_ /*nil*/)
}

// void emitDataChanged()
//export callback_ZN15QTreeWidgetItem15emitDataChangedEv
func callback_ZN15QTreeWidgetItem15emitDataChangedEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeWidgetItem.emitDataChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "emitDataChanged")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QTreeWidgetItem15emitDataChangedEv", C.callback_ZN15QTreeWidgetItem15emitDataChangedEv /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QTreeWidget5eventEP6QEvent
func callback_ZN11QTreeWidget5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeWidget.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QTreeWidget5eventEP6QEvent", C.callback_ZN11QTreeWidget5eventEP6QEvent /*nil*/)
}

// bool dropMimeData(class QTreeWidgetItem *, int, const class QMimeData *, Qt::DropAction)
//export callback_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE
func callback_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE(cthis unsafe.Pointer, parent unsafe.Pointer /*666*/, index C.int, data unsafe.Pointer /*666*/, action C.int) {
	// log.Println(cthis, "QTreeWidget.dropMimeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropMimeData", parent, index, data, action)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE", C.callback_ZN11QTreeWidget12dropMimeDataEP15QTreeWidgetItemiPK9QMimeDataN2Qt10DropActionE /*nil*/)
}

// Qt::DropActions supportedDropActions()
//export callback_ZNK11QTreeWidget20supportedDropActionsEv
func callback_ZNK11QTreeWidget20supportedDropActionsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTreeWidget.supportedDropActions")
	rvx := ffiqt.CallbackAllInherits(cthis, "supportedDropActions")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QTreeWidget20supportedDropActionsEv", C.callback_ZNK11QTreeWidget20supportedDropActionsEv /*nil*/)
}

// QModelIndex indexFromItem(const class QTreeWidgetItem *, int)
//export callback_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi
func callback_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi(cthis unsafe.Pointer, item unsafe.Pointer /*666*/, column C.int) {
	// log.Println(cthis, "QTreeWidget.indexFromItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "indexFromItem", item, column)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi", C.callback_ZNK11QTreeWidget13indexFromItemEPK15QTreeWidgetItemi /*nil*/)
}

// QModelIndex indexFromItem(class QTreeWidgetItem *, int)
//export callback_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi
func callback_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi(cthis unsafe.Pointer, item unsafe.Pointer /*666*/, column C.int) {
	// log.Println(cthis, "QTreeWidget.indexFromItem")
	rvx := ffiqt.CallbackAllInherits(cthis, "indexFromItem", item, column)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi", C.callback_ZNK11QTreeWidget13indexFromItemEP15QTreeWidgetItemi /*nil*/)
}

// QTreeWidgetItem * itemFromIndex(const class QModelIndex &)
//export callback_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex
func callback_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex(cthis unsafe.Pointer, index unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QTreeWidget.itemFromIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "itemFromIndex", index)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex", C.callback_ZNK11QTreeWidget13itemFromIndexERK11QModelIndex /*nil*/)
}

// void dropEvent(class QDropEvent *)
//export callback_ZN11QTreeWidget9dropEventEP10QDropEvent
func callback_ZN11QTreeWidget9dropEventEP10QDropEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTreeWidget.dropEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "dropEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QTreeWidget9dropEventEP10QDropEvent", C.callback_ZN11QTreeWidget9dropEventEP10QDropEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN13QWidgetAction5eventEP6QEvent
func callback_ZN13QWidgetAction5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidgetAction.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QWidgetAction5eventEP6QEvent", C.callback_ZN13QWidgetAction5eventEP6QEvent /*nil*/)
}

// bool eventFilter(class QObject *, class QEvent *)
//export callback_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent
func callback_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/, arg1 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidgetAction.eventFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventFilter", arg0, arg1)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent", C.callback_ZN13QWidgetAction11eventFilterEP7QObjectP6QEvent /*nil*/)
}

// QWidget * createWidget(class QWidget *)
//export callback_ZN13QWidgetAction12createWidgetEP7QWidget
func callback_ZN13QWidgetAction12createWidgetEP7QWidget(cthis unsafe.Pointer, parent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidgetAction.createWidget")
	rvx := ffiqt.CallbackAllInherits(cthis, "createWidget", parent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QWidgetAction12createWidgetEP7QWidget", C.callback_ZN13QWidgetAction12createWidgetEP7QWidget /*nil*/)
}

// void deleteWidget(class QWidget *)
//export callback_ZN13QWidgetAction12deleteWidgetEP7QWidget
func callback_ZN13QWidgetAction12deleteWidgetEP7QWidget(cthis unsafe.Pointer, widget unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWidgetAction.deleteWidget")
	rvx := ffiqt.CallbackAllInherits(cthis, "deleteWidget", widget)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QWidgetAction12deleteWidgetEP7QWidget", C.callback_ZN13QWidgetAction12deleteWidgetEP7QWidget /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN7QWizard5eventEP6QEvent
func callback_ZN7QWizard5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWizard.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWizard5eventEP6QEvent", C.callback_ZN7QWizard5eventEP6QEvent /*nil*/)
}

// void resizeEvent(class QResizeEvent *)
//export callback_ZN7QWizard11resizeEventEP12QResizeEvent
func callback_ZN7QWizard11resizeEventEP12QResizeEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWizard.resizeEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "resizeEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWizard11resizeEventEP12QResizeEvent", C.callback_ZN7QWizard11resizeEventEP12QResizeEvent /*nil*/)
}

// void paintEvent(class QPaintEvent *)
//export callback_ZN7QWizard10paintEventEP11QPaintEvent
func callback_ZN7QWizard10paintEventEP11QPaintEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWizard.paintEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "paintEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWizard10paintEventEP11QPaintEvent", C.callback_ZN7QWizard10paintEventEP11QPaintEvent /*nil*/)
}

// void done(int)
//export callback_ZN7QWizard4doneEi
func callback_ZN7QWizard4doneEi(cthis unsafe.Pointer, result C.int) {
	// log.Println(cthis, "QWizard.done")
	rvx := ffiqt.CallbackAllInherits(cthis, "done", result)
	gopp.ErrPrint(nil, rvx)
}
func init() { ffiqt.SetInheritCallback2c("_ZN7QWizard4doneEi", C.callback_ZN7QWizard4doneEi /*nil*/) }

// void initializePage(int)
//export callback_ZN7QWizard14initializePageEi
func callback_ZN7QWizard14initializePageEi(cthis unsafe.Pointer, id C.int) {
	// log.Println(cthis, "QWizard.initializePage")
	rvx := ffiqt.CallbackAllInherits(cthis, "initializePage", id)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWizard14initializePageEi", C.callback_ZN7QWizard14initializePageEi /*nil*/)
}

// void cleanupPage(int)
//export callback_ZN7QWizard11cleanupPageEi
func callback_ZN7QWizard11cleanupPageEi(cthis unsafe.Pointer, id C.int) {
	// log.Println(cthis, "QWizard.cleanupPage")
	rvx := ffiqt.CallbackAllInherits(cthis, "cleanupPage", id)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QWizard11cleanupPageEi", C.callback_ZN7QWizard11cleanupPageEi /*nil*/)
}

// void setField(const class QString &, const class QVariant &)
//export callback_ZN11QWizardPage8setFieldERK7QStringRK8QVariant
func callback_ZN11QWizardPage8setFieldERK7QStringRK8QVariant(cthis unsafe.Pointer, name unsafe.Pointer /*555*/, value unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QWizardPage.setField")
	rvx := ffiqt.CallbackAllInherits(cthis, "setField", name, value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QWizardPage8setFieldERK7QStringRK8QVariant", C.callback_ZN11QWizardPage8setFieldERK7QStringRK8QVariant /*nil*/)
}

// QVariant field(const class QString &)
//export callback_ZNK11QWizardPage5fieldERK7QString
func callback_ZNK11QWizardPage5fieldERK7QString(cthis unsafe.Pointer, name unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QWizardPage.field")
	rvx := ffiqt.CallbackAllInherits(cthis, "field", name)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QWizardPage5fieldERK7QString", C.callback_ZNK11QWizardPage5fieldERK7QString /*nil*/)
}

// void registerField(const class QString &, class QWidget *, const char *, const char *)
//export callback_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_
func callback_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_(cthis unsafe.Pointer, name unsafe.Pointer /*555*/, widget unsafe.Pointer /*666*/, property unsafe.Pointer /*666*/, changedSignal unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QWizardPage.registerField")
	rvx := ffiqt.CallbackAllInherits(cthis, "registerField", name, widget, property, changedSignal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_", C.callback_ZN11QWizardPage13registerFieldERK7QStringP7QWidgetPKcS6_ /*nil*/)
}

// QWizard * wizard()
//export callback_ZNK11QWizardPage6wizardEv
func callback_ZNK11QWizardPage6wizardEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QWizardPage.wizard")
	rvx := ffiqt.CallbackAllInherits(cthis, "wizard")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK11QWizardPage6wizardEv", C.callback_ZNK11QWizardPage6wizardEv /*nil*/)
}

//  body block end
