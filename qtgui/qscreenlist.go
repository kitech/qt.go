package qtgui

// /usr/include/qt/QtGui/qscreen.h
// #include <qscreen.h>
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

func init_unused_10155() {
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
type QScreenList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QScreenList) Operator_equal0() *QScreenList {
	// QScreenList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QScreenList) Operator_equal1() *QScreenList {
	// QScreenList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QScreenList) Swap0() {
	// QScreenList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QScreenList) Operator_equal_equal0() bool {
	// QScreenList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QScreenList) Operator_not_equal0() bool {
	// QScreenList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QScreenList) Size0() int {
	// QScreenList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QScreenList) Detach0() {
	// QScreenList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QScreenList) DetachShared0() {
	// QScreenList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QScreenList) IsDetached0() bool {
	// QScreenList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QScreenList) SetSharable0() {
	// QScreenList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QScreenList) IsSharedWith0() bool {
	// QScreenList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QScreenList) IsEmpty0() bool {
	// QScreenList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QScreenList) Clear0() {
	// QScreenList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QScreenList) At0() *QScreen {
	// QScreenList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & operator[](int)
func (this *QScreenList) Operator_get_index0() *QScreen {
	// QScreenList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T & operator[](int)
func (this *QScreenList) Operator_get_index1() *QScreen {
	// QScreenList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void reserve(int)
func (this *QScreenList) Reserve0() {
	// QScreenList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QScreenList) Append0() {
	// QScreenList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QScreenList) Append1() {
	// QScreenList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QScreenList) Prepend0() {
	// QScreenList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QScreenList) Insert0() {
	// QScreenList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QScreenList) Replace0() {
	// QScreenList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QScreenList) RemoveAt0() {
	// QScreenList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QScreenList) RemoveAll0() int {
	// QScreenList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QScreenList) RemoveOne0() bool {
	// QScreenList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QScreenList) TakeAt0() *QScreen {
	// QScreenList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T takeFirst()
func (this *QScreenList) TakeFirst0() *QScreen {
	// QScreenList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T takeLast()
func (this *QScreenList) TakeLast0() *QScreen {
	// QScreenList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void move(int, int)
func (this *QScreenList) Move0() {
	// QScreenList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QScreenList) Swap1() {
	// QScreenList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QScreenList) IndexOf0() int {
	// QScreenList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QScreenList) LastIndexOf0() int {
	// QScreenList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QScreenList) Contains0() bool {
	// QScreenList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QScreenList) Count0() int {
	// QScreenList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QScreenList) Begin0() {
	// QScreenList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QScreenList) Begin1() {
	// QScreenList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QScreenList) Cbegin0() {
	// QScreenList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QScreenList) ConstBegin0() {
	// QScreenList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QScreenList) End0() {
	// QScreenList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QScreenList) End1() {
	// QScreenList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QScreenList) Cend0() {
	// QScreenList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QScreenList) ConstEnd0() {
	// QScreenList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QScreenList) Rbegin0() {
	// QScreenList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QScreenList) Rend0() {
	// QScreenList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QScreenList) Rbegin1() {
	// QScreenList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QScreenList) Rend1() {
	// QScreenList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QScreenList) Crbegin0() {
	// QScreenList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QScreenList) Crend0() {
	// QScreenList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QScreenList) Insert1() {
	// QScreenList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QScreenList) Erase0() {
	// QScreenList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QScreenList) Erase1() {
	// QScreenList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QScreenList) Count1() int {
	// QScreenList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QScreenList) Length0() int {
	// QScreenList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QScreenList) First0() *QScreen {
	// QScreenList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & constFirst()
func (this *QScreenList) ConstFirst0() *QScreen {
	// QScreenList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & first()
func (this *QScreenList) First1() *QScreen {
	// QScreenList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T & last()
func (this *QScreenList) Last0() *QScreen {
	// QScreenList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & last()
func (this *QScreenList) Last1() *QScreen {
	// QScreenList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & constLast()
func (this *QScreenList) ConstLast0() *QScreen {
	// QScreenList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void removeFirst()
func (this *QScreenList) RemoveFirst0() {
	// QScreenList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QScreenList) RemoveLast0() {
	// QScreenList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QScreenList) StartsWith0() bool {
	// QScreenList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QScreenList) EndsWith0() bool {
	// QScreenList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QScreenList) Mid0() *QScreenList {
	// QScreenList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QScreenList) Value0() *QScreen {
	// QScreenList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T value(int, const T &)
func (this *QScreenList) Value1() *QScreen {
	// QScreenList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void push_back(const T &)
func (this *QScreenList) Push_back0() {
	// QScreenList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QScreenList) Push_front0() {
	// QScreenList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QScreenList) Front0() *QScreen {
	// QScreenList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & front()
func (this *QScreenList) Front1() *QScreen {
	// QScreenList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T & back()
func (this *QScreenList) Back0() *QScreen {
	// QScreenList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & back()
func (this *QScreenList) Back1() *QScreen {
	// QScreenList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void pop_front()
func (this *QScreenList) Pop_front0() {
	// QScreenList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QScreenList) Pop_back0() {
	// QScreenList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QScreenList) Empty0() bool {
	// QScreenList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QScreenList) Operator_add_equal0() *QScreenList {
	// QScreenList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QScreenList) Operator_add0() *QScreenList {
	// QScreenList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QScreenList) Operator_add_equal1() *QScreenList {
	// QScreenList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QScreenList) Operator_left_shift0() *QScreenList {
	// QScreenList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QScreenList) Operator_left_shift1() *QScreenList {
	// QScreenList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QScreenList) ToVector0() {
	// QScreenList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QScreenList) ToSet0() {
	// QScreenList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QScreenList) FromVector0() *QScreenList {
	// QScreenList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QScreenList) FromSet0() *QScreenList {
	// QScreenList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QScreenList) FromStdList0() *QScreenList {
	// QScreenList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QScreenList) ToStdList0() {
	// QScreenList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QScreenList) Detach_helper_grow0() {
	// QScreenList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QScreenList) Detach_helper0() {
	// QScreenList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QScreenList) Detach_helper1() {
	// QScreenList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QScreenList) Dealloc0() {
	// QScreenList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QScreenList) Node_construct0() {
	// QScreenList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QScreenList) Node_destruct0() {
	// QScreenList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QScreenList) Node_copy0() {
	// QScreenList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QScreenList) Node_destruct1() {
	// QScreenList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QScreenList) IsValidIterator0() bool {
	// QScreenList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QScreenList) Op_eq_impl0() bool {
	// QScreenList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QScreenList) Op_eq_impl1() bool {
	// QScreenList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QScreenList) Contains_impl0() bool {
	// QScreenList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QScreenList) Contains_impl1() bool {
	// QScreenList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QScreenList) Count_impl0() int {
	// QScreenList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QScreenList) Count_impl1() int {
	// QScreenList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
