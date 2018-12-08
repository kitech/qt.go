package qtgui

// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
// #include <QtGui>

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

//  ext block end

//  keep block begin

func init() {
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
}

//  keep block end

//  body block begin
type QWindowList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QWindowList) Operator_equal0() *QWindowList {
	// QWindowList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QWindowList) Operator_equal1() *QWindowList {
	// QWindowList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QWindowList) Swap0() {
	// QWindowList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QWindowList) Operator_equal_equal0() bool {
	// QWindowList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QWindowList) Operator_not_equal0() bool {
	// QWindowList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QWindowList) Size0() int {
	// QWindowList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QWindowList) Detach0() {
	// QWindowList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QWindowList) DetachShared0() {
	// QWindowList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QWindowList) IsDetached0() bool {
	// QWindowList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QWindowList) SetSharable0() {
	// QWindowList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QWindowList) IsSharedWith0() bool {
	// QWindowList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QWindowList) IsEmpty0() bool {
	// QWindowList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QWindowList) Clear0() {
	// QWindowList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QWindowList) At0() *QWindow {
	// QWindowList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & operator[](int)
func (this *QWindowList) Operator_get_index0() *QWindow {
	// QWindowList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T & operator[](int)
func (this *QWindowList) Operator_get_index1() *QWindow {
	// QWindowList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void reserve(int)
func (this *QWindowList) Reserve0() {
	// QWindowList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QWindowList) Append0() {
	// QWindowList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QWindowList) Append1() {
	// QWindowList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QWindowList) Prepend0() {
	// QWindowList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QWindowList) Insert0() {
	// QWindowList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QWindowList) Replace0() {
	// QWindowList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QWindowList) RemoveAt0() {
	// QWindowList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QWindowList) RemoveAll0() int {
	// QWindowList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QWindowList) RemoveOne0() bool {
	// QWindowList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QWindowList) TakeAt0() *QWindow {
	// QWindowList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T takeFirst()
func (this *QWindowList) TakeFirst0() *QWindow {
	// QWindowList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T takeLast()
func (this *QWindowList) TakeLast0() *QWindow {
	// QWindowList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void move(int, int)
func (this *QWindowList) Move0() {
	// QWindowList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QWindowList) Swap1() {
	// QWindowList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QWindowList) IndexOf0() int {
	// QWindowList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QWindowList) LastIndexOf0() int {
	// QWindowList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QWindowList) Contains0() bool {
	// QWindowList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QWindowList) Count0() int {
	// QWindowList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QWindowList) Begin0() {
	// QWindowList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QWindowList) Begin1() {
	// QWindowList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QWindowList) Cbegin0() {
	// QWindowList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QWindowList) ConstBegin0() {
	// QWindowList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QWindowList) End0() {
	// QWindowList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QWindowList) End1() {
	// QWindowList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QWindowList) Cend0() {
	// QWindowList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QWindowList) ConstEnd0() {
	// QWindowList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QWindowList) Rbegin0() {
	// QWindowList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QWindowList) Rend0() {
	// QWindowList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QWindowList) Rbegin1() {
	// QWindowList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QWindowList) Rend1() {
	// QWindowList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QWindowList) Crbegin0() {
	// QWindowList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QWindowList) Crend0() {
	// QWindowList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QWindowList) Insert1() {
	// QWindowList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QWindowList) Erase0() {
	// QWindowList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QWindowList) Erase1() {
	// QWindowList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QWindowList) Count1() int {
	// QWindowList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QWindowList) Length0() int {
	// QWindowList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QWindowList) First0() *QWindow {
	// QWindowList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & constFirst()
func (this *QWindowList) ConstFirst0() *QWindow {
	// QWindowList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & first()
func (this *QWindowList) First1() *QWindow {
	// QWindowList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T & last()
func (this *QWindowList) Last0() *QWindow {
	// QWindowList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & last()
func (this *QWindowList) Last1() *QWindow {
	// QWindowList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & constLast()
func (this *QWindowList) ConstLast0() *QWindow {
	// QWindowList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void removeFirst()
func (this *QWindowList) RemoveFirst0() {
	// QWindowList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QWindowList) RemoveLast0() {
	// QWindowList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QWindowList) StartsWith0() bool {
	// QWindowList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QWindowList) EndsWith0() bool {
	// QWindowList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QWindowList) Mid0() *QWindowList {
	// QWindowList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QWindowList) Value0() *QWindow {
	// QWindowList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T value(int, const T &)
func (this *QWindowList) Value1() *QWindow {
	// QWindowList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void push_back(const T &)
func (this *QWindowList) Push_back0() {
	// QWindowList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QWindowList) Push_front0() {
	// QWindowList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QWindowList) Front0() *QWindow {
	// QWindowList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & front()
func (this *QWindowList) Front1() *QWindow {
	// QWindowList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// T & back()
func (this *QWindowList) Back0() *QWindow {
	// QWindowList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// const T & back()
func (this *QWindowList) Back1() *QWindow {
	// QWindowList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QWindow{}
}

// void pop_front()
func (this *QWindowList) Pop_front0() {
	// QWindowList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QWindowList) Pop_back0() {
	// QWindowList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QWindowList) Empty0() bool {
	// QWindowList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QWindowList) Operator_add_equal0() *QWindowList {
	// QWindowList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QWindowList) Operator_add0() *QWindowList {
	// QWindowList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QWindowList) Operator_add_equal1() *QWindowList {
	// QWindowList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QWindowList) Operator_left_shift0() *QWindowList {
	// QWindowList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QWindowList) Operator_left_shift1() *QWindowList {
	// QWindowList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QWindowList) ToVector0() {
	// QWindowList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QWindowList) ToSet0() {
	// QWindowList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QWindowList) FromVector0() *QWindowList {
	// QWindowList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QWindowList) FromSet0() *QWindowList {
	// QWindowList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QWindowList) FromStdList0() *QWindowList {
	// QWindowList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QWindowList) ToStdList0() {
	// QWindowList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QWindowList) Detach_helper_grow0() {
	// QWindowList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QWindowList) Detach_helper0() {
	// QWindowList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QWindowList) Detach_helper1() {
	// QWindowList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QWindowList) Dealloc0() {
	// QWindowList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QWindowList) Node_construct0() {
	// QWindowList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QWindowList) Node_destruct0() {
	// QWindowList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QWindowList) Node_copy0() {
	// QWindowList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QWindowList) Node_destruct1() {
	// QWindowList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QWindowList) IsValidIterator0() bool {
	// QWindowList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QWindowList) Op_eq_impl0() bool {
	// QWindowList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QWindowList) Op_eq_impl1() bool {
	// QWindowList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QWindowList) Contains_impl0() bool {
	// QWindowList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QWindowList) Contains_impl1() bool {
	// QWindowList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QWindowList) Count_impl0() int {
	// QWindowList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QWindowList) Count_impl1() int {
	// QWindowList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
