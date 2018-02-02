//  package block begin
package qtcore

/*
#include <stdint.h>
#include <stdbool.h>
//  package block end

//  extern block begin
extern void callback_ZNK7QObject6senderEv(void* fnptr );
extern void callback_ZNK7QObject17senderSignalIndexEv(void* fnptr );
extern void callback_ZNK7QObject9receiversEPKc(void* fnptr , void* signal);
extern void callback_ZNK7QObject17isSignalConnectedERK11QMetaMethod(void* fnptr , void* signal);
extern void callback_ZN7QObject10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZN7QObject10childEventEP11QChildEvent(void* fnptr , void* event);
extern void callback_ZN7QObject11customEventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN7QObject13connectNotifyERK11QMetaMethod(void* fnptr , void* signal);
extern void callback_ZN7QObject16disconnectNotifyERK11QMetaMethod(void* fnptr , void* signal);
extern void callback_ZN18QAbstractAnimation5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN18QAbstractAnimation17updateCurrentTimeEi(void* fnptr , int currentTime);
extern void callback_ZN18QAbstractAnimation11updateStateENS_5StateES0_(void* fnptr , int newState, int oldState);
extern void callback_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE(void* fnptr , int direction);
extern void callback_ZN16QAnimationDriver16advanceAnimationEx(void* fnptr , int64_t timeStep);
extern void callback_ZN16QAnimationDriver5startEv(void* fnptr );
extern void callback_ZN16QAnimationDriver4stopEv(void* fnptr );
extern void callback_ZN8QVariant6createEiPKv(void* fnptr , int type_, void* copy);
extern void callback_ZNK8QVariant3cmpERKS_(void* fnptr , void* other);
extern void callback_ZNK8QVariant7compareERKS_(void* fnptr , void* other);
extern void callback_ZNK8QVariant7convertEiPv(void* fnptr , int t, void* ptr);
extern void callback_ZN18QAbstractItemModel17resetInternalDataEv(void* fnptr );
extern void callback_ZNK18QAbstractItemModel11createIndexEiiPv(void* fnptr , int row, int column, void* data);
extern void callback_ZNK18QAbstractItemModel11createIndexEiiy(void* fnptr , int row, int column, uint64_t id);
extern void callback_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream(void* fnptr , int row, int column, void* parent, void* stream);
extern void callback_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii(void* fnptr , void* parent, int first, int last);
extern void callback_ZN18QAbstractItemModel13endInsertRowsEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii(void* fnptr , void* parent, int first, int last);
extern void callback_ZN18QAbstractItemModel13endRemoveRowsEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i(void* fnptr , void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow);
extern void callback_ZN18QAbstractItemModel11endMoveRowsEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii(void* fnptr , void* parent, int first, int last);
extern void callback_ZN18QAbstractItemModel16endInsertColumnsEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii(void* fnptr , void* parent, int first, int last);
extern void callback_ZN18QAbstractItemModel16endRemoveColumnsEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i(void* fnptr , void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn);
extern void callback_ZN18QAbstractItemModel14endMoveColumnsEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel15beginResetModelEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel13endResetModelEv(void* fnptr );
extern void callback_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_(void* fnptr , void* from, void* to);
extern void callback_ZN19QAbstractProxyModel17resetInternalDataEv(void* fnptr );
extern void callback_ZN14QAbstractStateC1EP6QState(void* fnptr , void* parent);
extern void callback_ZN14QAbstractState7onEntryEP6QEvent(void* fnptr , void* event);
extern void callback_ZN14QAbstractState6onExitEP6QEvent(void* fnptr , void* event);
extern void callback_ZN14QAbstractState5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN19QAbstractTransition9eventTestEP6QEvent(void* fnptr , void* event);
extern void callback_ZN19QAbstractTransition12onTransitionEP6QEvent(void* fnptr , void* event);
extern void callback_ZN19QAbstractTransition5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN15QAnimationGroup5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN9QIODevice8readDataEPcx(void* fnptr , void* data, int64_t maxlen);
extern void callback_ZN9QIODevice12readLineDataEPcx(void* fnptr , void* data, int64_t maxlen);
extern void callback_ZN9QIODevice9writeDataEPKcx(void* fnptr , void* data, int64_t len);
extern void callback_ZN9QIODevice14setErrorStringERK7QString(void* fnptr , void* errorString);
extern void callback_ZN7QBuffer13connectNotifyERK11QMetaMethod(void* fnptr , void* arg0);
extern void callback_ZN7QBuffer16disconnectNotifyERK11QMetaMethod(void* fnptr , void* arg0);
extern void callback_ZN7QBuffer8readDataEPcx(void* fnptr , void* data, int64_t maxlen);
extern void callback_ZN7QBuffer9writeDataEPKcx(void* fnptr , void* data, int64_t len);
extern void callback_ZN27QStaticByteArrayMatcherBaseC1EPKcj(void* fnptr , void* pattern, unsigned int n);
extern void callback_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii(void* fnptr , void* needle, unsigned int nlen, void* haystack, int hlen, int from);
extern void callback_ZN16QCoreApplication5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZN16QEventTransition9eventTestEP6QEvent(void* fnptr , void* event);
extern void callback_ZN16QEventTransition12onTransitionEP6QEvent(void* fnptr , void* event);
extern void callback_ZN16QEventTransition5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN11QFinalState7onEntryEP6QEvent(void* fnptr , void* event);
extern void callback_ZN11QFinalState6onExitEP6QEvent(void* fnptr , void* event);
extern void callback_ZN11QFinalState5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZNK20QFutureInterfaceBase4refTEv(void* fnptr );
extern void callback_ZNK20QFutureInterfaceBase6derefTEv(void* fnptr );
extern void callback_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod(void* fnptr , void* signal);
extern void callback_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod(void* fnptr , void* signal);
extern void callback_ZN18QFutureWatcherBase22connectOutputInterfaceEv(void* fnptr );
extern void callback_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb(void* fnptr , bool pendingAssignment);
extern void callback_ZN13QHistoryState7onEntryEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QHistoryState6onExitEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QHistoryState5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN19QItemSelectionModel20emitSelectionChangedERK14QItemSelectionS2_(void* fnptr , void* newSelection, void* oldSelection);
extern void callback_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE(void* fnptr , void* mimetype, int preferredType);
extern void callback_ZN23QParallelAnimationGroup5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN23QParallelAnimationGroup17updateCurrentTimeEi(void* fnptr , int currentTime);
extern void callback_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_(void* fnptr , int newState, int oldState);
extern void callback_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE(void* fnptr , int direction);
extern void callback_ZN15QPauseAnimation5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN15QPauseAnimation17updateCurrentTimeEi(void* fnptr , int arg0);
extern void callback_ZN8QProcess15setProcessStateENS_12ProcessStateE(void* fnptr , int state);
extern void callback_ZN8QProcess17setupChildProcessEv(void* fnptr );
extern void callback_ZN8QProcess8readDataEPcx(void* fnptr , void* data, int64_t maxlen);
extern void callback_ZN8QProcess9writeDataEPKcx(void* fnptr , void* data, int64_t len);
extern void callback_ZN17QVariantAnimation5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN17QVariantAnimation17updateCurrentTimeEi(void* fnptr , int arg0);
extern void callback_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_(void* fnptr , int newState, int oldState);
extern void callback_ZN17QVariantAnimation18updateCurrentValueERK8QVariant(void* fnptr , void* value);
extern void callback_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d(void* fnptr , void* from, void* to, double progress);
extern void callback_ZN18QPropertyAnimation5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN18QPropertyAnimation18updateCurrentValueERK8QVariant(void* fnptr , void* value);
extern void callback_ZN18QPropertyAnimation11updateStateEN18QAbstractAnimation5StateES1_(void* fnptr , int newState, int oldState);
extern void callback_ZN16QRandomGeneratorC1ENS_6SystemE(void* fnptr , int arg0);
extern void callback_ZNK9QResource5isDirEv(void* fnptr );
extern void callback_ZNK9QResource6isFileEv(void* fnptr );
extern void callback_ZN9QSaveFile9writeDataEPKcx(void* fnptr , void* data, int64_t len);
extern void callback_ZN25QSequentialAnimationGroup5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN25QSequentialAnimationGroup17updateCurrentTimeEi(void* fnptr , int arg0);
extern void callback_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_(void* fnptr , int newState, int oldState);
extern void callback_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE(void* fnptr , int direction);
extern void callback_ZN9QSettings5eventEP6QEvent(void* fnptr , void* event);
extern void callback_ZN17QSignalTransition9eventTestEP6QEvent(void* fnptr , void* event);
extern void callback_ZN17QSignalTransition12onTransitionEP6QEvent(void* fnptr , void* event);
extern void callback_ZN17QSignalTransition5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN15QSocketNotifier5eventEP6QEvent(void* fnptr , void* arg0);
extern void callback_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex(void* fnptr , int source_row, void* source_parent);
extern void callback_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex(void* fnptr , int source_column, void* source_parent);
extern void callback_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_(void* fnptr , void* source_left, void* source_right);
extern void callback_ZN21QSortFilterProxyModel13filterChangedEv(void* fnptr );
extern void callback_ZN21QSortFilterProxyModel16invalidateFilterEv(void* fnptr );
extern void callback_ZN6QState7onEntryEP6QEvent(void* fnptr , void* event);
extern void callback_ZN6QState6onExitEP6QEvent(void* fnptr , void* event);
extern void callback_ZN6QState5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN13QStateMachine7onEntryEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QStateMachine6onExitEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QStateMachine22beginSelectTransitionsEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QStateMachine20endSelectTransitionsEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QStateMachine14beginMicrostepEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QStateMachine12endMicrostepEP6QEvent(void* fnptr , void* event);
extern void callback_ZN13QStateMachine5eventEP6QEvent(void* fnptr , void* e);
extern void callback_ZN21QAbstractConcatenable16convertFromAsciiEPKciRP5QChar(void* fnptr , void* a, int len, void* out);
extern void callback_ZN21QAbstractConcatenable16convertFromAsciiEcRP5QChar(void* fnptr , int8_t a, void* out);
extern void callback_ZN21QAbstractConcatenable14appendLatin1ToEPKciP5QChar(void* fnptr , void* a, int len, void* out);
extern void callback_ZN10QTextCodecC1Ev(void* fnptr );
extern void callback_ZN10QTextCodecD1Ev(void* fnptr );
extern void callback_ZN7QThread3runEv(void* fnptr );
extern void callback_ZN7QThread4execEv(void* fnptr );
extern void callback_ZN7QThread21setTerminationEnabledEb(void* fnptr , bool enabled);
extern void callback_ZN9QTimeLine10timerEventEP11QTimerEvent(void* fnptr , void* event);
extern void callback_ZN6QTimer10timerEventEP11QTimerEvent(void* fnptr , void* arg0);
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
// QObject * sender()
//export callback_ZNK7QObject6senderEv
func callback_ZNK7QObject6senderEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QObject.sender")
	rvx := ffiqt.CallbackAllInherits(cthis, "sender")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QObject6senderEv", C.callback_ZNK7QObject6senderEv /*nil*/)
}

// int senderSignalIndex()
//export callback_ZNK7QObject17senderSignalIndexEv
func callback_ZNK7QObject17senderSignalIndexEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QObject.senderSignalIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "senderSignalIndex")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QObject17senderSignalIndexEv", C.callback_ZNK7QObject17senderSignalIndexEv /*nil*/)
}

// int receivers(const char *)
//export callback_ZNK7QObject9receiversEPKc
func callback_ZNK7QObject9receiversEPKc(cthis unsafe.Pointer, signal unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QObject.receivers")
	rvx := ffiqt.CallbackAllInherits(cthis, "receivers", signal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QObject9receiversEPKc", C.callback_ZNK7QObject9receiversEPKc /*nil*/)
}

// bool isSignalConnected(const class QMetaMethod &)
//export callback_ZNK7QObject17isSignalConnectedERK11QMetaMethod
func callback_ZNK7QObject17isSignalConnectedERK11QMetaMethod(cthis unsafe.Pointer, signal unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QObject.isSignalConnected")
	rvx := ffiqt.CallbackAllInherits(cthis, "isSignalConnected", signal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK7QObject17isSignalConnectedERK11QMetaMethod", C.callback_ZNK7QObject17isSignalConnectedERK11QMetaMethod /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN7QObject10timerEventEP11QTimerEvent
func callback_ZN7QObject10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QObject.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QObject10timerEventEP11QTimerEvent", C.callback_ZN7QObject10timerEventEP11QTimerEvent /*nil*/)
}

// void childEvent(class QChildEvent *)
//export callback_ZN7QObject10childEventEP11QChildEvent
func callback_ZN7QObject10childEventEP11QChildEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QObject.childEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "childEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QObject10childEventEP11QChildEvent", C.callback_ZN7QObject10childEventEP11QChildEvent /*nil*/)
}

// void customEvent(class QEvent *)
//export callback_ZN7QObject11customEventEP6QEvent
func callback_ZN7QObject11customEventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QObject.customEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "customEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QObject11customEventEP6QEvent", C.callback_ZN7QObject11customEventEP6QEvent /*nil*/)
}

// void connectNotify(const class QMetaMethod &)
//export callback_ZN7QObject13connectNotifyERK11QMetaMethod
func callback_ZN7QObject13connectNotifyERK11QMetaMethod(cthis unsafe.Pointer, signal unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QObject.connectNotify")
	rvx := ffiqt.CallbackAllInherits(cthis, "connectNotify", signal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QObject13connectNotifyERK11QMetaMethod", C.callback_ZN7QObject13connectNotifyERK11QMetaMethod /*nil*/)
}

// void disconnectNotify(const class QMetaMethod &)
//export callback_ZN7QObject16disconnectNotifyERK11QMetaMethod
func callback_ZN7QObject16disconnectNotifyERK11QMetaMethod(cthis unsafe.Pointer, signal unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QObject.disconnectNotify")
	rvx := ffiqt.CallbackAllInherits(cthis, "disconnectNotify", signal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QObject16disconnectNotifyERK11QMetaMethod", C.callback_ZN7QObject16disconnectNotifyERK11QMetaMethod /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN18QAbstractAnimation5eventEP6QEvent
func callback_ZN18QAbstractAnimation5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractAnimation.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractAnimation5eventEP6QEvent", C.callback_ZN18QAbstractAnimation5eventEP6QEvent /*nil*/)
}

// void updateCurrentTime(int)
//export callback_ZN18QAbstractAnimation17updateCurrentTimeEi
func callback_ZN18QAbstractAnimation17updateCurrentTimeEi(cthis unsafe.Pointer, currentTime C.int) {
	// log.Println(cthis, "QAbstractAnimation.updateCurrentTime")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCurrentTime", currentTime)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractAnimation17updateCurrentTimeEi", C.callback_ZN18QAbstractAnimation17updateCurrentTimeEi /*nil*/)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
//export callback_ZN18QAbstractAnimation11updateStateENS_5StateES0_
func callback_ZN18QAbstractAnimation11updateStateENS_5StateES0_(cthis unsafe.Pointer, newState C.int, oldState C.int) {
	// log.Println(cthis, "QAbstractAnimation.updateState")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateState", newState, oldState)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractAnimation11updateStateENS_5StateES0_", C.callback_ZN18QAbstractAnimation11updateStateENS_5StateES0_ /*nil*/)
}

// void updateDirection(class QAbstractAnimation::Direction)
//export callback_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE
func callback_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE(cthis unsafe.Pointer, direction C.int) {
	// log.Println(cthis, "QAbstractAnimation.updateDirection")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateDirection", direction)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE", C.callback_ZN18QAbstractAnimation15updateDirectionENS_9DirectionE /*nil*/)
}

// void advanceAnimation(qint64)
//export callback_ZN16QAnimationDriver16advanceAnimationEx
func callback_ZN16QAnimationDriver16advanceAnimationEx(cthis unsafe.Pointer, timeStep C.int64_t) {
	// log.Println(cthis, "QAnimationDriver.advanceAnimation")
	rvx := ffiqt.CallbackAllInherits(cthis, "advanceAnimation", timeStep)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAnimationDriver16advanceAnimationEx", C.callback_ZN16QAnimationDriver16advanceAnimationEx /*nil*/)
}

// void start()
//export callback_ZN16QAnimationDriver5startEv
func callback_ZN16QAnimationDriver5startEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAnimationDriver.start")
	rvx := ffiqt.CallbackAllInherits(cthis, "start")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAnimationDriver5startEv", C.callback_ZN16QAnimationDriver5startEv /*nil*/)
}

// void stop()
//export callback_ZN16QAnimationDriver4stopEv
func callback_ZN16QAnimationDriver4stopEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAnimationDriver.stop")
	rvx := ffiqt.CallbackAllInherits(cthis, "stop")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QAnimationDriver4stopEv", C.callback_ZN16QAnimationDriver4stopEv /*nil*/)
}

// void create(int, const void *)
//export callback_ZN8QVariant6createEiPKv
func callback_ZN8QVariant6createEiPKv(cthis unsafe.Pointer, type_ C.int, copy unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QVariant.create")
	rvx := ffiqt.CallbackAllInherits(cthis, "create", type_, copy)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QVariant6createEiPKv", C.callback_ZN8QVariant6createEiPKv /*nil*/)
}

// bool cmp(const class QVariant &)
//export callback_ZNK8QVariant3cmpERKS_
func callback_ZNK8QVariant3cmpERKS_(cthis unsafe.Pointer, other unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QVariant.cmp")
	rvx := ffiqt.CallbackAllInherits(cthis, "cmp", other)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QVariant3cmpERKS_", C.callback_ZNK8QVariant3cmpERKS_ /*nil*/)
}

// int compare(const class QVariant &)
//export callback_ZNK8QVariant7compareERKS_
func callback_ZNK8QVariant7compareERKS_(cthis unsafe.Pointer, other unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QVariant.compare")
	rvx := ffiqt.CallbackAllInherits(cthis, "compare", other)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QVariant7compareERKS_", C.callback_ZNK8QVariant7compareERKS_ /*nil*/)
}

// bool convert(const int, void *)
//export callback_ZNK8QVariant7convertEiPv
func callback_ZNK8QVariant7convertEiPv(cthis unsafe.Pointer, t C.int, ptr unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QVariant.convert")
	rvx := ffiqt.CallbackAllInherits(cthis, "convert", t, ptr)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK8QVariant7convertEiPv", C.callback_ZNK8QVariant7convertEiPv /*nil*/)
}

// void resetInternalData()
//export callback_ZN18QAbstractItemModel17resetInternalDataEv
func callback_ZN18QAbstractItemModel17resetInternalDataEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.resetInternalData")
	rvx := ffiqt.CallbackAllInherits(cthis, "resetInternalData")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel17resetInternalDataEv", C.callback_ZN18QAbstractItemModel17resetInternalDataEv /*nil*/)
}

// QModelIndex createIndex(int, int, void *)
//export callback_ZNK18QAbstractItemModel11createIndexEiiPv
func callback_ZNK18QAbstractItemModel11createIndexEiiPv(cthis unsafe.Pointer, row C.int, column C.int, data unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractItemModel.createIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "createIndex", row, column, data)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QAbstractItemModel11createIndexEiiPv", C.callback_ZNK18QAbstractItemModel11createIndexEiiPv /*nil*/)
}

// QModelIndex createIndex(int, int, quintptr)
//export callback_ZNK18QAbstractItemModel11createIndexEiiy
func callback_ZNK18QAbstractItemModel11createIndexEiiy(cthis unsafe.Pointer, row C.int, column C.int, id C.uint64_t) {
	// log.Println(cthis, "QAbstractItemModel.createIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "createIndex", row, column, id)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK18QAbstractItemModel11createIndexEiiy", C.callback_ZNK18QAbstractItemModel11createIndexEiiy /*nil*/)
}

// bool decodeData(int, int, const class QModelIndex &, class QDataStream &)
//export callback_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream
func callback_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream(cthis unsafe.Pointer, row C.int, column C.int, parent unsafe.Pointer /*555*/, stream unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractItemModel.decodeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "decodeData", row, column, parent, stream)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream", C.callback_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream /*nil*/)
}

// void beginInsertRows(const class QModelIndex &, int, int)
//export callback_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii
func callback_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, first C.int, last C.int) {
	// log.Println(cthis, "QAbstractItemModel.beginInsertRows")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginInsertRows", parent, first, last)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii", C.callback_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii /*nil*/)
}

// void endInsertRows()
//export callback_ZN18QAbstractItemModel13endInsertRowsEv
func callback_ZN18QAbstractItemModel13endInsertRowsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.endInsertRows")
	rvx := ffiqt.CallbackAllInherits(cthis, "endInsertRows")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel13endInsertRowsEv", C.callback_ZN18QAbstractItemModel13endInsertRowsEv /*nil*/)
}

// void beginRemoveRows(const class QModelIndex &, int, int)
//export callback_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii
func callback_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, first C.int, last C.int) {
	// log.Println(cthis, "QAbstractItemModel.beginRemoveRows")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginRemoveRows", parent, first, last)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii", C.callback_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii /*nil*/)
}

// void endRemoveRows()
//export callback_ZN18QAbstractItemModel13endRemoveRowsEv
func callback_ZN18QAbstractItemModel13endRemoveRowsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.endRemoveRows")
	rvx := ffiqt.CallbackAllInherits(cthis, "endRemoveRows")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel13endRemoveRowsEv", C.callback_ZN18QAbstractItemModel13endRemoveRowsEv /*nil*/)
}

// bool beginMoveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
//export callback_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i
func callback_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i(cthis unsafe.Pointer, sourceParent unsafe.Pointer /*555*/, sourceFirst C.int, sourceLast C.int, destinationParent unsafe.Pointer /*555*/, destinationRow C.int) {
	// log.Println(cthis, "QAbstractItemModel.beginMoveRows")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginMoveRows", sourceParent, sourceFirst, sourceLast, destinationParent, destinationRow)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i", C.callback_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i /*nil*/)
}

// void endMoveRows()
//export callback_ZN18QAbstractItemModel11endMoveRowsEv
func callback_ZN18QAbstractItemModel11endMoveRowsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.endMoveRows")
	rvx := ffiqt.CallbackAllInherits(cthis, "endMoveRows")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel11endMoveRowsEv", C.callback_ZN18QAbstractItemModel11endMoveRowsEv /*nil*/)
}

// void beginInsertColumns(const class QModelIndex &, int, int)
//export callback_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii
func callback_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, first C.int, last C.int) {
	// log.Println(cthis, "QAbstractItemModel.beginInsertColumns")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginInsertColumns", parent, first, last)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii", C.callback_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii /*nil*/)
}

// void endInsertColumns()
//export callback_ZN18QAbstractItemModel16endInsertColumnsEv
func callback_ZN18QAbstractItemModel16endInsertColumnsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.endInsertColumns")
	rvx := ffiqt.CallbackAllInherits(cthis, "endInsertColumns")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel16endInsertColumnsEv", C.callback_ZN18QAbstractItemModel16endInsertColumnsEv /*nil*/)
}

// void beginRemoveColumns(const class QModelIndex &, int, int)
//export callback_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii
func callback_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii(cthis unsafe.Pointer, parent unsafe.Pointer /*555*/, first C.int, last C.int) {
	// log.Println(cthis, "QAbstractItemModel.beginRemoveColumns")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginRemoveColumns", parent, first, last)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii", C.callback_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii /*nil*/)
}

// void endRemoveColumns()
//export callback_ZN18QAbstractItemModel16endRemoveColumnsEv
func callback_ZN18QAbstractItemModel16endRemoveColumnsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.endRemoveColumns")
	rvx := ffiqt.CallbackAllInherits(cthis, "endRemoveColumns")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel16endRemoveColumnsEv", C.callback_ZN18QAbstractItemModel16endRemoveColumnsEv /*nil*/)
}

// bool beginMoveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
//export callback_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i
func callback_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i(cthis unsafe.Pointer, sourceParent unsafe.Pointer /*555*/, sourceFirst C.int, sourceLast C.int, destinationParent unsafe.Pointer /*555*/, destinationColumn C.int) {
	// log.Println(cthis, "QAbstractItemModel.beginMoveColumns")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginMoveColumns", sourceParent, sourceFirst, sourceLast, destinationParent, destinationColumn)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i", C.callback_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i /*nil*/)
}

// void endMoveColumns()
//export callback_ZN18QAbstractItemModel14endMoveColumnsEv
func callback_ZN18QAbstractItemModel14endMoveColumnsEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.endMoveColumns")
	rvx := ffiqt.CallbackAllInherits(cthis, "endMoveColumns")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel14endMoveColumnsEv", C.callback_ZN18QAbstractItemModel14endMoveColumnsEv /*nil*/)
}

// void beginResetModel()
//export callback_ZN18QAbstractItemModel15beginResetModelEv
func callback_ZN18QAbstractItemModel15beginResetModelEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.beginResetModel")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginResetModel")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel15beginResetModelEv", C.callback_ZN18QAbstractItemModel15beginResetModelEv /*nil*/)
}

// void endResetModel()
//export callback_ZN18QAbstractItemModel13endResetModelEv
func callback_ZN18QAbstractItemModel13endResetModelEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractItemModel.endResetModel")
	rvx := ffiqt.CallbackAllInherits(cthis, "endResetModel")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel13endResetModelEv", C.callback_ZN18QAbstractItemModel13endResetModelEv /*nil*/)
}

// void changePersistentIndex(const class QModelIndex &, const class QModelIndex &)
//export callback_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_
func callback_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_(cthis unsafe.Pointer, from unsafe.Pointer /*555*/, to unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractItemModel.changePersistentIndex")
	rvx := ffiqt.CallbackAllInherits(cthis, "changePersistentIndex", from, to)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_", C.callback_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_ /*nil*/)
}

// void resetInternalData()
//export callback_ZN19QAbstractProxyModel17resetInternalDataEv
func callback_ZN19QAbstractProxyModel17resetInternalDataEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QAbstractProxyModel.resetInternalData")
	rvx := ffiqt.CallbackAllInherits(cthis, "resetInternalData")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractProxyModel17resetInternalDataEv", C.callback_ZN19QAbstractProxyModel17resetInternalDataEv /*nil*/)
}

// void QAbstractState(class QState *)
//export callback_ZN14QAbstractStateC1EP6QState
func callback_ZN14QAbstractStateC1EP6QState(cthis unsafe.Pointer, parent unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractState.QAbstractState")
	rvx := ffiqt.CallbackAllInherits(cthis, "QAbstractState", parent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QAbstractStateC1EP6QState", C.callback_ZN14QAbstractStateC1EP6QState /*nil*/)
}

// void onEntry(class QEvent *)
//export callback_ZN14QAbstractState7onEntryEP6QEvent
func callback_ZN14QAbstractState7onEntryEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractState.onEntry")
	rvx := ffiqt.CallbackAllInherits(cthis, "onEntry", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QAbstractState7onEntryEP6QEvent", C.callback_ZN14QAbstractState7onEntryEP6QEvent /*nil*/)
}

// void onExit(class QEvent *)
//export callback_ZN14QAbstractState6onExitEP6QEvent
func callback_ZN14QAbstractState6onExitEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractState.onExit")
	rvx := ffiqt.CallbackAllInherits(cthis, "onExit", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QAbstractState6onExitEP6QEvent", C.callback_ZN14QAbstractState6onExitEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN14QAbstractState5eventEP6QEvent
func callback_ZN14QAbstractState5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractState.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN14QAbstractState5eventEP6QEvent", C.callback_ZN14QAbstractState5eventEP6QEvent /*nil*/)
}

// bool eventTest(class QEvent *)
//export callback_ZN19QAbstractTransition9eventTestEP6QEvent
func callback_ZN19QAbstractTransition9eventTestEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractTransition.eventTest")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventTest", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractTransition9eventTestEP6QEvent", C.callback_ZN19QAbstractTransition9eventTestEP6QEvent /*nil*/)
}

// void onTransition(class QEvent *)
//export callback_ZN19QAbstractTransition12onTransitionEP6QEvent
func callback_ZN19QAbstractTransition12onTransitionEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractTransition.onTransition")
	rvx := ffiqt.CallbackAllInherits(cthis, "onTransition", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractTransition12onTransitionEP6QEvent", C.callback_ZN19QAbstractTransition12onTransitionEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN19QAbstractTransition5eventEP6QEvent
func callback_ZN19QAbstractTransition5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractTransition.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QAbstractTransition5eventEP6QEvent", C.callback_ZN19QAbstractTransition5eventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QAnimationGroup5eventEP6QEvent
func callback_ZN15QAnimationGroup5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAnimationGroup.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QAnimationGroup5eventEP6QEvent", C.callback_ZN15QAnimationGroup5eventEP6QEvent /*nil*/)
}

// qint64 readData(char *, qint64)
//export callback_ZN9QIODevice8readDataEPcx
func callback_ZN9QIODevice8readDataEPcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, maxlen C.int64_t) {
	// log.Println(cthis, "QIODevice.readData")
	rvx := ffiqt.CallbackAllInherits(cthis, "readData", data, maxlen)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QIODevice8readDataEPcx", C.callback_ZN9QIODevice8readDataEPcx /*nil*/)
}

// qint64 readLineData(char *, qint64)
//export callback_ZN9QIODevice12readLineDataEPcx
func callback_ZN9QIODevice12readLineDataEPcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, maxlen C.int64_t) {
	// log.Println(cthis, "QIODevice.readLineData")
	rvx := ffiqt.CallbackAllInherits(cthis, "readLineData", data, maxlen)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QIODevice12readLineDataEPcx", C.callback_ZN9QIODevice12readLineDataEPcx /*nil*/)
}

// qint64 writeData(const char *, qint64)
//export callback_ZN9QIODevice9writeDataEPKcx
func callback_ZN9QIODevice9writeDataEPKcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, len C.int64_t) {
	// log.Println(cthis, "QIODevice.writeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "writeData", data, len)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QIODevice9writeDataEPKcx", C.callback_ZN9QIODevice9writeDataEPKcx /*nil*/)
}

// void setErrorString(const class QString &)
//export callback_ZN9QIODevice14setErrorStringERK7QString
func callback_ZN9QIODevice14setErrorStringERK7QString(cthis unsafe.Pointer, errorString unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QIODevice.setErrorString")
	rvx := ffiqt.CallbackAllInherits(cthis, "setErrorString", errorString)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QIODevice14setErrorStringERK7QString", C.callback_ZN9QIODevice14setErrorStringERK7QString /*nil*/)
}

// void connectNotify(const class QMetaMethod &)
//export callback_ZN7QBuffer13connectNotifyERK11QMetaMethod
func callback_ZN7QBuffer13connectNotifyERK11QMetaMethod(cthis unsafe.Pointer, arg0 unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QBuffer.connectNotify")
	rvx := ffiqt.CallbackAllInherits(cthis, "connectNotify", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QBuffer13connectNotifyERK11QMetaMethod", C.callback_ZN7QBuffer13connectNotifyERK11QMetaMethod /*nil*/)
}

// void disconnectNotify(const class QMetaMethod &)
//export callback_ZN7QBuffer16disconnectNotifyERK11QMetaMethod
func callback_ZN7QBuffer16disconnectNotifyERK11QMetaMethod(cthis unsafe.Pointer, arg0 unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QBuffer.disconnectNotify")
	rvx := ffiqt.CallbackAllInherits(cthis, "disconnectNotify", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QBuffer16disconnectNotifyERK11QMetaMethod", C.callback_ZN7QBuffer16disconnectNotifyERK11QMetaMethod /*nil*/)
}

// qint64 readData(char *, qint64)
//export callback_ZN7QBuffer8readDataEPcx
func callback_ZN7QBuffer8readDataEPcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, maxlen C.int64_t) {
	// log.Println(cthis, "QBuffer.readData")
	rvx := ffiqt.CallbackAllInherits(cthis, "readData", data, maxlen)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QBuffer8readDataEPcx", C.callback_ZN7QBuffer8readDataEPcx /*nil*/)
}

// qint64 writeData(const char *, qint64)
//export callback_ZN7QBuffer9writeDataEPKcx
func callback_ZN7QBuffer9writeDataEPKcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, len C.int64_t) {
	// log.Println(cthis, "QBuffer.writeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "writeData", data, len)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QBuffer9writeDataEPKcx", C.callback_ZN7QBuffer9writeDataEPKcx /*nil*/)
}

// void QStaticByteArrayMatcherBase(const char *, uint)
//export callback_ZN27QStaticByteArrayMatcherBaseC1EPKcj
func callback_ZN27QStaticByteArrayMatcherBaseC1EPKcj(cthis unsafe.Pointer, pattern unsafe.Pointer /*666*/, n C.uint) {
	// log.Println(cthis, "QStaticByteArrayMatcherBase.QStaticByteArrayMatcherBase")
	rvx := ffiqt.CallbackAllInherits(cthis, "QStaticByteArrayMatcherBase", pattern, n)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN27QStaticByteArrayMatcherBaseC1EPKcj", C.callback_ZN27QStaticByteArrayMatcherBaseC1EPKcj /*nil*/)
}

// int indexOfIn(const char *, uint, const char *, int, int)
//export callback_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii
func callback_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii(cthis unsafe.Pointer, needle unsafe.Pointer /*666*/, nlen C.uint, haystack unsafe.Pointer /*666*/, hlen C.int, from C.int) {
	// log.Println(cthis, "QStaticByteArrayMatcherBase.indexOfIn")
	rvx := ffiqt.CallbackAllInherits(cthis, "indexOfIn", needle, nlen, haystack, hlen, from)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii", C.callback_ZNK27QStaticByteArrayMatcherBase9indexOfInEPKcjS1_ii /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN16QCoreApplication5eventEP6QEvent
func callback_ZN16QCoreApplication5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QCoreApplication.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QCoreApplication5eventEP6QEvent", C.callback_ZN16QCoreApplication5eventEP6QEvent /*nil*/)
}

// bool eventTest(class QEvent *)
//export callback_ZN16QEventTransition9eventTestEP6QEvent
func callback_ZN16QEventTransition9eventTestEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QEventTransition.eventTest")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventTest", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QEventTransition9eventTestEP6QEvent", C.callback_ZN16QEventTransition9eventTestEP6QEvent /*nil*/)
}

// void onTransition(class QEvent *)
//export callback_ZN16QEventTransition12onTransitionEP6QEvent
func callback_ZN16QEventTransition12onTransitionEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QEventTransition.onTransition")
	rvx := ffiqt.CallbackAllInherits(cthis, "onTransition", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QEventTransition12onTransitionEP6QEvent", C.callback_ZN16QEventTransition12onTransitionEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN16QEventTransition5eventEP6QEvent
func callback_ZN16QEventTransition5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QEventTransition.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QEventTransition5eventEP6QEvent", C.callback_ZN16QEventTransition5eventEP6QEvent /*nil*/)
}

// void onEntry(class QEvent *)
//export callback_ZN11QFinalState7onEntryEP6QEvent
func callback_ZN11QFinalState7onEntryEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFinalState.onEntry")
	rvx := ffiqt.CallbackAllInherits(cthis, "onEntry", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFinalState7onEntryEP6QEvent", C.callback_ZN11QFinalState7onEntryEP6QEvent /*nil*/)
}

// void onExit(class QEvent *)
//export callback_ZN11QFinalState6onExitEP6QEvent
func callback_ZN11QFinalState6onExitEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFinalState.onExit")
	rvx := ffiqt.CallbackAllInherits(cthis, "onExit", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFinalState6onExitEP6QEvent", C.callback_ZN11QFinalState6onExitEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN11QFinalState5eventEP6QEvent
func callback_ZN11QFinalState5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QFinalState.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN11QFinalState5eventEP6QEvent", C.callback_ZN11QFinalState5eventEP6QEvent /*nil*/)
}

// bool refT()
//export callback_ZNK20QFutureInterfaceBase4refTEv
func callback_ZNK20QFutureInterfaceBase4refTEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QFutureInterfaceBase.refT")
	rvx := ffiqt.CallbackAllInherits(cthis, "refT")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QFutureInterfaceBase4refTEv", C.callback_ZNK20QFutureInterfaceBase4refTEv /*nil*/)
}

// bool derefT()
//export callback_ZNK20QFutureInterfaceBase6derefTEv
func callback_ZNK20QFutureInterfaceBase6derefTEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QFutureInterfaceBase.derefT")
	rvx := ffiqt.CallbackAllInherits(cthis, "derefT")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK20QFutureInterfaceBase6derefTEv", C.callback_ZNK20QFutureInterfaceBase6derefTEv /*nil*/)
}

// void connectNotify(const class QMetaMethod &)
//export callback_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod
func callback_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod(cthis unsafe.Pointer, signal unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QFutureWatcherBase.connectNotify")
	rvx := ffiqt.CallbackAllInherits(cthis, "connectNotify", signal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod", C.callback_ZN18QFutureWatcherBase13connectNotifyERK11QMetaMethod /*nil*/)
}

// void disconnectNotify(const class QMetaMethod &)
//export callback_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod
func callback_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod(cthis unsafe.Pointer, signal unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QFutureWatcherBase.disconnectNotify")
	rvx := ffiqt.CallbackAllInherits(cthis, "disconnectNotify", signal)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod", C.callback_ZN18QFutureWatcherBase16disconnectNotifyERK11QMetaMethod /*nil*/)
}

// void connectOutputInterface()
//export callback_ZN18QFutureWatcherBase22connectOutputInterfaceEv
func callback_ZN18QFutureWatcherBase22connectOutputInterfaceEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QFutureWatcherBase.connectOutputInterface")
	rvx := ffiqt.CallbackAllInherits(cthis, "connectOutputInterface")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QFutureWatcherBase22connectOutputInterfaceEv", C.callback_ZN18QFutureWatcherBase22connectOutputInterfaceEv /*nil*/)
}

// void disconnectOutputInterface(_Bool)
//export callback_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb
func callback_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb(cthis unsafe.Pointer, pendingAssignment C.bool) {
	// log.Println(cthis, "QFutureWatcherBase.disconnectOutputInterface")
	rvx := ffiqt.CallbackAllInherits(cthis, "disconnectOutputInterface", pendingAssignment)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb", C.callback_ZN18QFutureWatcherBase25disconnectOutputInterfaceEb /*nil*/)
}

// void onEntry(class QEvent *)
//export callback_ZN13QHistoryState7onEntryEP6QEvent
func callback_ZN13QHistoryState7onEntryEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHistoryState.onEntry")
	rvx := ffiqt.CallbackAllInherits(cthis, "onEntry", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QHistoryState7onEntryEP6QEvent", C.callback_ZN13QHistoryState7onEntryEP6QEvent /*nil*/)
}

// void onExit(class QEvent *)
//export callback_ZN13QHistoryState6onExitEP6QEvent
func callback_ZN13QHistoryState6onExitEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHistoryState.onExit")
	rvx := ffiqt.CallbackAllInherits(cthis, "onExit", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QHistoryState6onExitEP6QEvent", C.callback_ZN13QHistoryState6onExitEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN13QHistoryState5eventEP6QEvent
func callback_ZN13QHistoryState5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QHistoryState.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QHistoryState5eventEP6QEvent", C.callback_ZN13QHistoryState5eventEP6QEvent /*nil*/)
}

// void emitSelectionChanged(const class QItemSelection &, const class QItemSelection &)
//export callback_ZN19QItemSelectionModel20emitSelectionChangedERK14QItemSelectionS2_
func callback_ZN19QItemSelectionModel20emitSelectionChangedERK14QItemSelectionS2_(cthis unsafe.Pointer, newSelection unsafe.Pointer /*555*/, oldSelection unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QItemSelectionModel.emitSelectionChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "emitSelectionChanged", newSelection, oldSelection)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN19QItemSelectionModel20emitSelectionChangedERK14QItemSelectionS2_", C.callback_ZN19QItemSelectionModel20emitSelectionChangedERK14QItemSelectionS2_ /*nil*/)
}

// QVariant retrieveData(const class QString &, class QVariant::Type)
//export callback_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE
func callback_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE(cthis unsafe.Pointer, mimetype unsafe.Pointer /*555*/, preferredType C.int) {
	// log.Println(cthis, "QMimeData.retrieveData")
	rvx := ffiqt.CallbackAllInherits(cthis, "retrieveData", mimetype, preferredType)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE", C.callback_ZNK9QMimeData12retrieveDataERK7QStringN8QVariant4TypeE /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN23QParallelAnimationGroup5eventEP6QEvent
func callback_ZN23QParallelAnimationGroup5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QParallelAnimationGroup.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN23QParallelAnimationGroup5eventEP6QEvent", C.callback_ZN23QParallelAnimationGroup5eventEP6QEvent /*nil*/)
}

// void updateCurrentTime(int)
//export callback_ZN23QParallelAnimationGroup17updateCurrentTimeEi
func callback_ZN23QParallelAnimationGroup17updateCurrentTimeEi(cthis unsafe.Pointer, currentTime C.int) {
	// log.Println(cthis, "QParallelAnimationGroup.updateCurrentTime")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCurrentTime", currentTime)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN23QParallelAnimationGroup17updateCurrentTimeEi", C.callback_ZN23QParallelAnimationGroup17updateCurrentTimeEi /*nil*/)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
//export callback_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_
func callback_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_(cthis unsafe.Pointer, newState C.int, oldState C.int) {
	// log.Println(cthis, "QParallelAnimationGroup.updateState")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateState", newState, oldState)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", C.callback_ZN23QParallelAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_ /*nil*/)
}

// void updateDirection(class QAbstractAnimation::Direction)
//export callback_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE
func callback_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE(cthis unsafe.Pointer, direction C.int) {
	// log.Println(cthis, "QParallelAnimationGroup.updateDirection")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateDirection", direction)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", C.callback_ZN23QParallelAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QPauseAnimation5eventEP6QEvent
func callback_ZN15QPauseAnimation5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPauseAnimation.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QPauseAnimation5eventEP6QEvent", C.callback_ZN15QPauseAnimation5eventEP6QEvent /*nil*/)
}

// void updateCurrentTime(int)
//export callback_ZN15QPauseAnimation17updateCurrentTimeEi
func callback_ZN15QPauseAnimation17updateCurrentTimeEi(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QPauseAnimation.updateCurrentTime")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCurrentTime", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QPauseAnimation17updateCurrentTimeEi", C.callback_ZN15QPauseAnimation17updateCurrentTimeEi /*nil*/)
}

// void setProcessState(enum QProcess::ProcessState)
//export callback_ZN8QProcess15setProcessStateENS_12ProcessStateE
func callback_ZN8QProcess15setProcessStateENS_12ProcessStateE(cthis unsafe.Pointer, state C.int) {
	// log.Println(cthis, "QProcess.setProcessState")
	rvx := ffiqt.CallbackAllInherits(cthis, "setProcessState", state)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QProcess15setProcessStateENS_12ProcessStateE", C.callback_ZN8QProcess15setProcessStateENS_12ProcessStateE /*nil*/)
}

// void setupChildProcess()
//export callback_ZN8QProcess17setupChildProcessEv
func callback_ZN8QProcess17setupChildProcessEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QProcess.setupChildProcess")
	rvx := ffiqt.CallbackAllInherits(cthis, "setupChildProcess")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QProcess17setupChildProcessEv", C.callback_ZN8QProcess17setupChildProcessEv /*nil*/)
}

// qint64 readData(char *, qint64)
//export callback_ZN8QProcess8readDataEPcx
func callback_ZN8QProcess8readDataEPcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, maxlen C.int64_t) {
	// log.Println(cthis, "QProcess.readData")
	rvx := ffiqt.CallbackAllInherits(cthis, "readData", data, maxlen)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QProcess8readDataEPcx", C.callback_ZN8QProcess8readDataEPcx /*nil*/)
}

// qint64 writeData(const char *, qint64)
//export callback_ZN8QProcess9writeDataEPKcx
func callback_ZN8QProcess9writeDataEPKcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, len C.int64_t) {
	// log.Println(cthis, "QProcess.writeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "writeData", data, len)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN8QProcess9writeDataEPKcx", C.callback_ZN8QProcess9writeDataEPKcx /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN17QVariantAnimation5eventEP6QEvent
func callback_ZN17QVariantAnimation5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QVariantAnimation.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QVariantAnimation5eventEP6QEvent", C.callback_ZN17QVariantAnimation5eventEP6QEvent /*nil*/)
}

// void updateCurrentTime(int)
//export callback_ZN17QVariantAnimation17updateCurrentTimeEi
func callback_ZN17QVariantAnimation17updateCurrentTimeEi(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QVariantAnimation.updateCurrentTime")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCurrentTime", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QVariantAnimation17updateCurrentTimeEi", C.callback_ZN17QVariantAnimation17updateCurrentTimeEi /*nil*/)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
//export callback_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_
func callback_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_(cthis unsafe.Pointer, newState C.int, oldState C.int) {
	// log.Println(cthis, "QVariantAnimation.updateState")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateState", newState, oldState)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_", C.callback_ZN17QVariantAnimation11updateStateEN18QAbstractAnimation5StateES1_ /*nil*/)
}

// void updateCurrentValue(const class QVariant &)
//export callback_ZN17QVariantAnimation18updateCurrentValueERK8QVariant
func callback_ZN17QVariantAnimation18updateCurrentValueERK8QVariant(cthis unsafe.Pointer, value unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QVariantAnimation.updateCurrentValue")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCurrentValue", value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QVariantAnimation18updateCurrentValueERK8QVariant", C.callback_ZN17QVariantAnimation18updateCurrentValueERK8QVariant /*nil*/)
}

// QVariant interpolated(const class QVariant &, const class QVariant &, qreal)
//export callback_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d
func callback_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d(cthis unsafe.Pointer, from unsafe.Pointer /*555*/, to unsafe.Pointer /*555*/, progress C.double) {
	// log.Println(cthis, "QVariantAnimation.interpolated")
	rvx := ffiqt.CallbackAllInherits(cthis, "interpolated", from, to, progress)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d", C.callback_ZNK17QVariantAnimation12interpolatedERK8QVariantS2_d /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN18QPropertyAnimation5eventEP6QEvent
func callback_ZN18QPropertyAnimation5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QPropertyAnimation.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QPropertyAnimation5eventEP6QEvent", C.callback_ZN18QPropertyAnimation5eventEP6QEvent /*nil*/)
}

// void updateCurrentValue(const class QVariant &)
//export callback_ZN18QPropertyAnimation18updateCurrentValueERK8QVariant
func callback_ZN18QPropertyAnimation18updateCurrentValueERK8QVariant(cthis unsafe.Pointer, value unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QPropertyAnimation.updateCurrentValue")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCurrentValue", value)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QPropertyAnimation18updateCurrentValueERK8QVariant", C.callback_ZN18QPropertyAnimation18updateCurrentValueERK8QVariant /*nil*/)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
//export callback_ZN18QPropertyAnimation11updateStateEN18QAbstractAnimation5StateES1_
func callback_ZN18QPropertyAnimation11updateStateEN18QAbstractAnimation5StateES1_(cthis unsafe.Pointer, newState C.int, oldState C.int) {
	// log.Println(cthis, "QPropertyAnimation.updateState")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateState", newState, oldState)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN18QPropertyAnimation11updateStateEN18QAbstractAnimation5StateES1_", C.callback_ZN18QPropertyAnimation11updateStateEN18QAbstractAnimation5StateES1_ /*nil*/)
}

// void QRandomGenerator(enum QRandomGenerator::System)
//export callback_ZN16QRandomGeneratorC1ENS_6SystemE
func callback_ZN16QRandomGeneratorC1ENS_6SystemE(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QRandomGenerator.QRandomGenerator")
	rvx := ffiqt.CallbackAllInherits(cthis, "QRandomGenerator", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN16QRandomGeneratorC1ENS_6SystemE", C.callback_ZN16QRandomGeneratorC1ENS_6SystemE /*nil*/)
}

// bool isDir()
//export callback_ZNK9QResource5isDirEv
func callback_ZNK9QResource5isDirEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QResource.isDir")
	rvx := ffiqt.CallbackAllInherits(cthis, "isDir")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QResource5isDirEv", C.callback_ZNK9QResource5isDirEv /*nil*/)
}

// bool isFile()
//export callback_ZNK9QResource6isFileEv
func callback_ZNK9QResource6isFileEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QResource.isFile")
	rvx := ffiqt.CallbackAllInherits(cthis, "isFile")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK9QResource6isFileEv", C.callback_ZNK9QResource6isFileEv /*nil*/)
}

// qint64 writeData(const char *, qint64)
//export callback_ZN9QSaveFile9writeDataEPKcx
func callback_ZN9QSaveFile9writeDataEPKcx(cthis unsafe.Pointer, data unsafe.Pointer /*666*/, len C.int64_t) {
	// log.Println(cthis, "QSaveFile.writeData")
	rvx := ffiqt.CallbackAllInherits(cthis, "writeData", data, len)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSaveFile9writeDataEPKcx", C.callback_ZN9QSaveFile9writeDataEPKcx /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN25QSequentialAnimationGroup5eventEP6QEvent
func callback_ZN25QSequentialAnimationGroup5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSequentialAnimationGroup.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN25QSequentialAnimationGroup5eventEP6QEvent", C.callback_ZN25QSequentialAnimationGroup5eventEP6QEvent /*nil*/)
}

// void updateCurrentTime(int)
//export callback_ZN25QSequentialAnimationGroup17updateCurrentTimeEi
func callback_ZN25QSequentialAnimationGroup17updateCurrentTimeEi(cthis unsafe.Pointer, arg0 C.int) {
	// log.Println(cthis, "QSequentialAnimationGroup.updateCurrentTime")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateCurrentTime", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN25QSequentialAnimationGroup17updateCurrentTimeEi", C.callback_ZN25QSequentialAnimationGroup17updateCurrentTimeEi /*nil*/)
}

// void updateState(class QAbstractAnimation::State, class QAbstractAnimation::State)
//export callback_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_
func callback_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_(cthis unsafe.Pointer, newState C.int, oldState C.int) {
	// log.Println(cthis, "QSequentialAnimationGroup.updateState")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateState", newState, oldState)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_", C.callback_ZN25QSequentialAnimationGroup11updateStateEN18QAbstractAnimation5StateES1_ /*nil*/)
}

// void updateDirection(class QAbstractAnimation::Direction)
//export callback_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE
func callback_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE(cthis unsafe.Pointer, direction C.int) {
	// log.Println(cthis, "QSequentialAnimationGroup.updateDirection")
	rvx := ffiqt.CallbackAllInherits(cthis, "updateDirection", direction)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE", C.callback_ZN25QSequentialAnimationGroup15updateDirectionEN18QAbstractAnimation9DirectionE /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN9QSettings5eventEP6QEvent
func callback_ZN9QSettings5eventEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSettings.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QSettings5eventEP6QEvent", C.callback_ZN9QSettings5eventEP6QEvent /*nil*/)
}

// bool eventTest(class QEvent *)
//export callback_ZN17QSignalTransition9eventTestEP6QEvent
func callback_ZN17QSignalTransition9eventTestEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSignalTransition.eventTest")
	rvx := ffiqt.CallbackAllInherits(cthis, "eventTest", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QSignalTransition9eventTestEP6QEvent", C.callback_ZN17QSignalTransition9eventTestEP6QEvent /*nil*/)
}

// void onTransition(class QEvent *)
//export callback_ZN17QSignalTransition12onTransitionEP6QEvent
func callback_ZN17QSignalTransition12onTransitionEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSignalTransition.onTransition")
	rvx := ffiqt.CallbackAllInherits(cthis, "onTransition", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QSignalTransition12onTransitionEP6QEvent", C.callback_ZN17QSignalTransition12onTransitionEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN17QSignalTransition5eventEP6QEvent
func callback_ZN17QSignalTransition5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSignalTransition.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN17QSignalTransition5eventEP6QEvent", C.callback_ZN17QSignalTransition5eventEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN15QSocketNotifier5eventEP6QEvent
func callback_ZN15QSocketNotifier5eventEP6QEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QSocketNotifier.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN15QSocketNotifier5eventEP6QEvent", C.callback_ZN15QSocketNotifier5eventEP6QEvent /*nil*/)
}

// bool filterAcceptsRow(int, const class QModelIndex &)
//export callback_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex
func callback_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex(cthis unsafe.Pointer, source_row C.int, source_parent unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSortFilterProxyModel.filterAcceptsRow")
	rvx := ffiqt.CallbackAllInherits(cthis, "filterAcceptsRow", source_row, source_parent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex", C.callback_ZNK21QSortFilterProxyModel16filterAcceptsRowEiRK11QModelIndex /*nil*/)
}

// bool filterAcceptsColumn(int, const class QModelIndex &)
//export callback_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex
func callback_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex(cthis unsafe.Pointer, source_column C.int, source_parent unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSortFilterProxyModel.filterAcceptsColumn")
	rvx := ffiqt.CallbackAllInherits(cthis, "filterAcceptsColumn", source_column, source_parent)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex", C.callback_ZNK21QSortFilterProxyModel19filterAcceptsColumnEiRK11QModelIndex /*nil*/)
}

// bool lessThan(const class QModelIndex &, const class QModelIndex &)
//export callback_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_
func callback_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_(cthis unsafe.Pointer, source_left unsafe.Pointer /*555*/, source_right unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QSortFilterProxyModel.lessThan")
	rvx := ffiqt.CallbackAllInherits(cthis, "lessThan", source_left, source_right)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_", C.callback_ZNK21QSortFilterProxyModel8lessThanERK11QModelIndexS2_ /*nil*/)
}

// void filterChanged()
//export callback_ZN21QSortFilterProxyModel13filterChangedEv
func callback_ZN21QSortFilterProxyModel13filterChangedEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QSortFilterProxyModel.filterChanged")
	rvx := ffiqt.CallbackAllInherits(cthis, "filterChanged")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QSortFilterProxyModel13filterChangedEv", C.callback_ZN21QSortFilterProxyModel13filterChangedEv /*nil*/)
}

// void invalidateFilter()
//export callback_ZN21QSortFilterProxyModel16invalidateFilterEv
func callback_ZN21QSortFilterProxyModel16invalidateFilterEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QSortFilterProxyModel.invalidateFilter")
	rvx := ffiqt.CallbackAllInherits(cthis, "invalidateFilter")
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QSortFilterProxyModel16invalidateFilterEv", C.callback_ZN21QSortFilterProxyModel16invalidateFilterEv /*nil*/)
}

// void onEntry(class QEvent *)
//export callback_ZN6QState7onEntryEP6QEvent
func callback_ZN6QState7onEntryEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QState.onEntry")
	rvx := ffiqt.CallbackAllInherits(cthis, "onEntry", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QState7onEntryEP6QEvent", C.callback_ZN6QState7onEntryEP6QEvent /*nil*/)
}

// void onExit(class QEvent *)
//export callback_ZN6QState6onExitEP6QEvent
func callback_ZN6QState6onExitEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QState.onExit")
	rvx := ffiqt.CallbackAllInherits(cthis, "onExit", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QState6onExitEP6QEvent", C.callback_ZN6QState6onExitEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN6QState5eventEP6QEvent
func callback_ZN6QState5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QState.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QState5eventEP6QEvent", C.callback_ZN6QState5eventEP6QEvent /*nil*/)
}

// void onEntry(class QEvent *)
//export callback_ZN13QStateMachine7onEntryEP6QEvent
func callback_ZN13QStateMachine7onEntryEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStateMachine.onEntry")
	rvx := ffiqt.CallbackAllInherits(cthis, "onEntry", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStateMachine7onEntryEP6QEvent", C.callback_ZN13QStateMachine7onEntryEP6QEvent /*nil*/)
}

// void onExit(class QEvent *)
//export callback_ZN13QStateMachine6onExitEP6QEvent
func callback_ZN13QStateMachine6onExitEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStateMachine.onExit")
	rvx := ffiqt.CallbackAllInherits(cthis, "onExit", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStateMachine6onExitEP6QEvent", C.callback_ZN13QStateMachine6onExitEP6QEvent /*nil*/)
}

// void beginSelectTransitions(class QEvent *)
//export callback_ZN13QStateMachine22beginSelectTransitionsEP6QEvent
func callback_ZN13QStateMachine22beginSelectTransitionsEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStateMachine.beginSelectTransitions")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginSelectTransitions", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStateMachine22beginSelectTransitionsEP6QEvent", C.callback_ZN13QStateMachine22beginSelectTransitionsEP6QEvent /*nil*/)
}

// void endSelectTransitions(class QEvent *)
//export callback_ZN13QStateMachine20endSelectTransitionsEP6QEvent
func callback_ZN13QStateMachine20endSelectTransitionsEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStateMachine.endSelectTransitions")
	rvx := ffiqt.CallbackAllInherits(cthis, "endSelectTransitions", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStateMachine20endSelectTransitionsEP6QEvent", C.callback_ZN13QStateMachine20endSelectTransitionsEP6QEvent /*nil*/)
}

// void beginMicrostep(class QEvent *)
//export callback_ZN13QStateMachine14beginMicrostepEP6QEvent
func callback_ZN13QStateMachine14beginMicrostepEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStateMachine.beginMicrostep")
	rvx := ffiqt.CallbackAllInherits(cthis, "beginMicrostep", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStateMachine14beginMicrostepEP6QEvent", C.callback_ZN13QStateMachine14beginMicrostepEP6QEvent /*nil*/)
}

// void endMicrostep(class QEvent *)
//export callback_ZN13QStateMachine12endMicrostepEP6QEvent
func callback_ZN13QStateMachine12endMicrostepEP6QEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStateMachine.endMicrostep")
	rvx := ffiqt.CallbackAllInherits(cthis, "endMicrostep", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStateMachine12endMicrostepEP6QEvent", C.callback_ZN13QStateMachine12endMicrostepEP6QEvent /*nil*/)
}

// bool event(class QEvent *)
//export callback_ZN13QStateMachine5eventEP6QEvent
func callback_ZN13QStateMachine5eventEP6QEvent(cthis unsafe.Pointer, e unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QStateMachine.event")
	rvx := ffiqt.CallbackAllInherits(cthis, "event", e)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN13QStateMachine5eventEP6QEvent", C.callback_ZN13QStateMachine5eventEP6QEvent /*nil*/)
}

// void convertFromAscii(const char *, int, class QChar *&)
//export callback_ZN21QAbstractConcatenable16convertFromAsciiEPKciRP5QChar
func callback_ZN21QAbstractConcatenable16convertFromAsciiEPKciRP5QChar(cthis unsafe.Pointer, a unsafe.Pointer /*666*/, len C.int, out unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractConcatenable.convertFromAscii")
	rvx := ffiqt.CallbackAllInherits(cthis, "convertFromAscii", a, len, out)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QAbstractConcatenable16convertFromAsciiEPKciRP5QChar", C.callback_ZN21QAbstractConcatenable16convertFromAsciiEPKciRP5QChar /*nil*/)
}

// void convertFromAscii(char, class QChar *&)
//export callback_ZN21QAbstractConcatenable16convertFromAsciiEcRP5QChar
func callback_ZN21QAbstractConcatenable16convertFromAsciiEcRP5QChar(cthis unsafe.Pointer, a C.int8_t, out unsafe.Pointer /*555*/) {
	// log.Println(cthis, "QAbstractConcatenable.convertFromAscii")
	rvx := ffiqt.CallbackAllInherits(cthis, "convertFromAscii", a, out)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QAbstractConcatenable16convertFromAsciiEcRP5QChar", C.callback_ZN21QAbstractConcatenable16convertFromAsciiEcRP5QChar /*nil*/)
}

// void appendLatin1To(const char *, int, class QChar *)
//export callback_ZN21QAbstractConcatenable14appendLatin1ToEPKciP5QChar
func callback_ZN21QAbstractConcatenable14appendLatin1ToEPKciP5QChar(cthis unsafe.Pointer, a unsafe.Pointer /*666*/, len C.int, out unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QAbstractConcatenable.appendLatin1To")
	rvx := ffiqt.CallbackAllInherits(cthis, "appendLatin1To", a, len, out)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN21QAbstractConcatenable14appendLatin1ToEPKciP5QChar", C.callback_ZN21QAbstractConcatenable14appendLatin1ToEPKciP5QChar /*nil*/)
}

// void QTextCodec()
//export callback_ZN10QTextCodecC1Ev
func callback_ZN10QTextCodecC1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTextCodec.QTextCodec")
	rvx := ffiqt.CallbackAllInherits(cthis, "QTextCodec")
	gopp.ErrPrint(nil, rvx)
}
func init() { ffiqt.SetInheritCallback2c("_ZN10QTextCodecC1Ev", C.callback_ZN10QTextCodecC1Ev /*nil*/) }

// void ~QTextCodec()
//export callback_ZN10QTextCodecD1Ev
func callback_ZN10QTextCodecD1Ev(cthis unsafe.Pointer) {
	// log.Println(cthis, "QTextCodec.~QTextCodec")
	rvx := ffiqt.CallbackAllInherits(cthis, "~QTextCodec")
	gopp.ErrPrint(nil, rvx)
}
func init() { ffiqt.SetInheritCallback2c("_ZN10QTextCodecD1Ev", C.callback_ZN10QTextCodecD1Ev /*nil*/) }

// void run()
//export callback_ZN7QThread3runEv
func callback_ZN7QThread3runEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QThread.run")
	rvx := ffiqt.CallbackAllInherits(cthis, "run")
	gopp.ErrPrint(nil, rvx)
}
func init() { ffiqt.SetInheritCallback2c("_ZN7QThread3runEv", C.callback_ZN7QThread3runEv /*nil*/) }

// int exec()
//export callback_ZN7QThread4execEv
func callback_ZN7QThread4execEv(cthis unsafe.Pointer) {
	// log.Println(cthis, "QThread.exec")
	rvx := ffiqt.CallbackAllInherits(cthis, "exec")
	gopp.ErrPrint(nil, rvx)
}
func init() { ffiqt.SetInheritCallback2c("_ZN7QThread4execEv", C.callback_ZN7QThread4execEv /*nil*/) }

// void setTerminationEnabled(_Bool)
//export callback_ZN7QThread21setTerminationEnabledEb
func callback_ZN7QThread21setTerminationEnabledEb(cthis unsafe.Pointer, enabled C.bool) {
	// log.Println(cthis, "QThread.setTerminationEnabled")
	rvx := ffiqt.CallbackAllInherits(cthis, "setTerminationEnabled", enabled)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN7QThread21setTerminationEnabledEb", C.callback_ZN7QThread21setTerminationEnabledEb /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN9QTimeLine10timerEventEP11QTimerEvent
func callback_ZN9QTimeLine10timerEventEP11QTimerEvent(cthis unsafe.Pointer, event unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTimeLine.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", event)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN9QTimeLine10timerEventEP11QTimerEvent", C.callback_ZN9QTimeLine10timerEventEP11QTimerEvent /*nil*/)
}

// void timerEvent(class QTimerEvent *)
//export callback_ZN6QTimer10timerEventEP11QTimerEvent
func callback_ZN6QTimer10timerEventEP11QTimerEvent(cthis unsafe.Pointer, arg0 unsafe.Pointer /*666*/) {
	// log.Println(cthis, "QTimer.timerEvent")
	rvx := ffiqt.CallbackAllInherits(cthis, "timerEvent", arg0)
	gopp.ErrPrint(nil, rvx)
}
func init() {
	ffiqt.SetInheritCallback2c("_ZN6QTimer10timerEventEP11QTimerEvent", C.callback_ZN6QTimer10timerEventEP11QTimerEvent /*nil*/)
}

//  body block end
