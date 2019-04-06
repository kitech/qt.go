package qtwidgets

// /usr/include/qt/QtWidgets/qgesture.h
// #include <qgesture.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  keep block begin

func init_unused_10131() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end

//  body block begin
type QGestureList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QGestureList) Operator_equal0() *QGestureList {
	// QGestureList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QGestureList) Operator_equal1() *QGestureList {
	// QGestureList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QGestureList) Swap0() {
	// QGestureList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QGestureList) Operator_equal_equal0() bool {
	// QGestureList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QGestureList) Operator_not_equal0() bool {
	// QGestureList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QGestureList) Size0() int {
	// QGestureList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QGestureList) Detach0() {
	// QGestureList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QGestureList) DetachShared0() {
	// QGestureList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QGestureList) IsDetached0() bool {
	// QGestureList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QGestureList) SetSharable0() {
	// QGestureList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QGestureList) IsSharedWith0() bool {
	// QGestureList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QGestureList) IsEmpty0() bool {
	// QGestureList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QGestureList) Clear0() {
	// QGestureList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QGestureList) At0() *QGesture {
	// QGestureList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// const T & operator[](int)
func (this *QGestureList) Operator_get_index0() *QGesture {
	// QGestureList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// T & operator[](int)
func (this *QGestureList) Operator_get_index1() *QGesture {
	// QGestureList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// void reserve(int)
func (this *QGestureList) Reserve0() {
	// QGestureList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QGestureList) Append0() {
	// QGestureList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QGestureList) Append1() {
	// QGestureList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QGestureList) Prepend0() {
	// QGestureList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QGestureList) Insert0() {
	// QGestureList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QGestureList) Replace0() {
	// QGestureList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QGestureList) RemoveAt0() {
	// QGestureList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QGestureList) RemoveAll0() int {
	// QGestureList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QGestureList) RemoveOne0() bool {
	// QGestureList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QGestureList) TakeAt0() *QGesture {
	// QGestureList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// T takeFirst()
func (this *QGestureList) TakeFirst0() *QGesture {
	// QGestureList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// T takeLast()
func (this *QGestureList) TakeLast0() *QGesture {
	// QGestureList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// void move(int, int)
func (this *QGestureList) Move0() {
	// QGestureList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QGestureList) Swap1() {
	// QGestureList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QGestureList) IndexOf0() int {
	// QGestureList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QGestureList) LastIndexOf0() int {
	// QGestureList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QGestureList) Contains0() bool {
	// QGestureList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QGestureList) Count0() int {
	// QGestureList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QGestureList) Begin0() {
	// QGestureList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QGestureList) Begin1() {
	// QGestureList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QGestureList) Cbegin0() {
	// QGestureList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QGestureList) ConstBegin0() {
	// QGestureList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QGestureList) End0() {
	// QGestureList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QGestureList) End1() {
	// QGestureList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QGestureList) Cend0() {
	// QGestureList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QGestureList) ConstEnd0() {
	// QGestureList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QGestureList) Rbegin0() {
	// QGestureList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QGestureList) Rend0() {
	// QGestureList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QGestureList) Rbegin1() {
	// QGestureList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QGestureList) Rend1() {
	// QGestureList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QGestureList) Crbegin0() {
	// QGestureList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QGestureList) Crend0() {
	// QGestureList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QGestureList) Insert1() {
	// QGestureList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QGestureList) Erase0() {
	// QGestureList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QGestureList) Erase1() {
	// QGestureList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QGestureList) Count1() int {
	// QGestureList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QGestureList) Length0() int {
	// QGestureList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QGestureList) First0() *QGesture {
	// QGestureList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// const T & constFirst()
func (this *QGestureList) ConstFirst0() *QGesture {
	// QGestureList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// const T & first()
func (this *QGestureList) First1() *QGesture {
	// QGestureList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// T & last()
func (this *QGestureList) Last0() *QGesture {
	// QGestureList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// const T & last()
func (this *QGestureList) Last1() *QGesture {
	// QGestureList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// const T & constLast()
func (this *QGestureList) ConstLast0() *QGesture {
	// QGestureList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// void removeFirst()
func (this *QGestureList) RemoveFirst0() {
	// QGestureList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QGestureList) RemoveLast0() {
	// QGestureList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QGestureList) StartsWith0() bool {
	// QGestureList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QGestureList) EndsWith0() bool {
	// QGestureList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QGestureList) Mid0() *QGestureList {
	// QGestureList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QGestureList) Value0() *QGesture {
	// QGestureList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// T value(int, const T &)
func (this *QGestureList) Value1() *QGesture {
	// QGestureList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// void push_back(const T &)
func (this *QGestureList) Push_back0() {
	// QGestureList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QGestureList) Push_front0() {
	// QGestureList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QGestureList) Front0() *QGesture {
	// QGestureList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// const T & front()
func (this *QGestureList) Front1() *QGesture {
	// QGestureList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// T & back()
func (this *QGestureList) Back0() *QGesture {
	// QGestureList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// const T & back()
func (this *QGestureList) Back1() *QGesture {
	// QGestureList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QGesture{}
}

// void pop_front()
func (this *QGestureList) Pop_front0() {
	// QGestureList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QGestureList) Pop_back0() {
	// QGestureList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QGestureList) Empty0() bool {
	// QGestureList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QGestureList) Operator_add_equal0() *QGestureList {
	// QGestureList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QGestureList) Operator_add0() *QGestureList {
	// QGestureList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QGestureList) Operator_add_equal1() *QGestureList {
	// QGestureList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QGestureList) Operator_left_shift0() *QGestureList {
	// QGestureList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QGestureList) Operator_left_shift1() *QGestureList {
	// QGestureList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QGestureList) ToVector0() {
	// QGestureList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QGestureList) ToSet0() {
	// QGestureList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QGestureList) FromVector0() *QGestureList {
	// QGestureList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QGestureList) FromSet0() *QGestureList {
	// QGestureList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QGestureList) FromStdList0() *QGestureList {
	// QGestureList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QGestureList) ToStdList0() {
	// QGestureList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QGestureList) Detach_helper_grow0() {
	// QGestureList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QGestureList) Detach_helper0() {
	// QGestureList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QGestureList) Detach_helper1() {
	// QGestureList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QGestureList) Dealloc0() {
	// QGestureList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QGestureList) Node_construct0() {
	// QGestureList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QGestureList) Node_destruct0() {
	// QGestureList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QGestureList) Node_copy0() {
	// QGestureList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QGestureList) Node_destruct1() {
	// QGestureList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QGestureList) IsValidIterator0() bool {
	// QGestureList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QGestureList) Op_eq_impl0() bool {
	// QGestureList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QGestureList) Op_eq_impl1() bool {
	// QGestureList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGestureList) Contains_impl0() bool {
	// QGestureList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGestureList) Contains_impl1() bool {
	// QGestureList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QGestureList) Count_impl0() int {
	// QGestureList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QGestureList) Count_impl1() int {
	// QGestureList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QGestureList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
