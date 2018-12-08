package qtgui

// /usr/include/qt/QtGui/qkeysequence.h
// #include <qkeysequence.h>
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
type QKeySequenceList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QKeySequenceList) Operator_equal0() *QKeySequenceList {
	// QKeySequenceList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QKeySequenceList) Operator_equal1() *QKeySequenceList {
	// QKeySequenceList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QKeySequenceList) Swap0() {
	// QKeySequenceList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QKeySequenceList) Operator_equal_equal0() bool {
	// QKeySequenceList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QKeySequenceList) Operator_not_equal0() bool {
	// QKeySequenceList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QKeySequenceList) Size0() int {
	// QKeySequenceList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QKeySequenceList) Detach0() {
	// QKeySequenceList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QKeySequenceList) DetachShared0() {
	// QKeySequenceList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QKeySequenceList) IsDetached0() bool {
	// QKeySequenceList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QKeySequenceList) SetSharable0() {
	// QKeySequenceList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QKeySequenceList) IsSharedWith0() bool {
	// QKeySequenceList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QKeySequenceList) IsEmpty0() bool {
	// QKeySequenceList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QKeySequenceList) Clear0() {
	// QKeySequenceList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QKeySequenceList) At0() *QKeySequence {
	// QKeySequenceList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// const T & operator[](int)
func (this *QKeySequenceList) Operator_get_index0() *QKeySequence {
	// QKeySequenceList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// T & operator[](int)
func (this *QKeySequenceList) Operator_get_index1() *QKeySequence {
	// QKeySequenceList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// void reserve(int)
func (this *QKeySequenceList) Reserve0() {
	// QKeySequenceList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QKeySequenceList) Append0() {
	// QKeySequenceList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QKeySequenceList) Append1() {
	// QKeySequenceList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QKeySequenceList) Prepend0() {
	// QKeySequenceList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QKeySequenceList) Insert0() {
	// QKeySequenceList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QKeySequenceList) Replace0() {
	// QKeySequenceList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QKeySequenceList) RemoveAt0() {
	// QKeySequenceList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QKeySequenceList) RemoveAll0() int {
	// QKeySequenceList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QKeySequenceList) RemoveOne0() bool {
	// QKeySequenceList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QKeySequenceList) TakeAt0() *QKeySequence {
	// QKeySequenceList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// T takeFirst()
func (this *QKeySequenceList) TakeFirst0() *QKeySequence {
	// QKeySequenceList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// T takeLast()
func (this *QKeySequenceList) TakeLast0() *QKeySequence {
	// QKeySequenceList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// void move(int, int)
func (this *QKeySequenceList) Move0() {
	// QKeySequenceList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QKeySequenceList) Swap1() {
	// QKeySequenceList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QKeySequenceList) IndexOf0() int {
	// QKeySequenceList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QKeySequenceList) LastIndexOf0() int {
	// QKeySequenceList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QKeySequenceList) Contains0() bool {
	// QKeySequenceList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QKeySequenceList) Count0() int {
	// QKeySequenceList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QKeySequenceList) Begin0() {
	// QKeySequenceList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QKeySequenceList) Begin1() {
	// QKeySequenceList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QKeySequenceList) Cbegin0() {
	// QKeySequenceList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QKeySequenceList) ConstBegin0() {
	// QKeySequenceList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QKeySequenceList) End0() {
	// QKeySequenceList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QKeySequenceList) End1() {
	// QKeySequenceList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QKeySequenceList) Cend0() {
	// QKeySequenceList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QKeySequenceList) ConstEnd0() {
	// QKeySequenceList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QKeySequenceList) Rbegin0() {
	// QKeySequenceList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QKeySequenceList) Rend0() {
	// QKeySequenceList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QKeySequenceList) Rbegin1() {
	// QKeySequenceList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QKeySequenceList) Rend1() {
	// QKeySequenceList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QKeySequenceList) Crbegin0() {
	// QKeySequenceList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QKeySequenceList) Crend0() {
	// QKeySequenceList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QKeySequenceList) Insert1() {
	// QKeySequenceList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QKeySequenceList) Erase0() {
	// QKeySequenceList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QKeySequenceList) Erase1() {
	// QKeySequenceList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QKeySequenceList) Count1() int {
	// QKeySequenceList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QKeySequenceList) Length0() int {
	// QKeySequenceList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QKeySequenceList) First0() *QKeySequence {
	// QKeySequenceList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// const T & constFirst()
func (this *QKeySequenceList) ConstFirst0() *QKeySequence {
	// QKeySequenceList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// const T & first()
func (this *QKeySequenceList) First1() *QKeySequence {
	// QKeySequenceList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// T & last()
func (this *QKeySequenceList) Last0() *QKeySequence {
	// QKeySequenceList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// const T & last()
func (this *QKeySequenceList) Last1() *QKeySequence {
	// QKeySequenceList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// const T & constLast()
func (this *QKeySequenceList) ConstLast0() *QKeySequence {
	// QKeySequenceList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// void removeFirst()
func (this *QKeySequenceList) RemoveFirst0() {
	// QKeySequenceList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QKeySequenceList) RemoveLast0() {
	// QKeySequenceList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QKeySequenceList) StartsWith0() bool {
	// QKeySequenceList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QKeySequenceList) EndsWith0() bool {
	// QKeySequenceList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QKeySequenceList) Mid0() *QKeySequenceList {
	// QKeySequenceList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QKeySequenceList) Value0() *QKeySequence {
	// QKeySequenceList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// T value(int, const T &)
func (this *QKeySequenceList) Value1() *QKeySequence {
	// QKeySequenceList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// void push_back(const T &)
func (this *QKeySequenceList) Push_back0() {
	// QKeySequenceList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QKeySequenceList) Push_front0() {
	// QKeySequenceList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QKeySequenceList) Front0() *QKeySequence {
	// QKeySequenceList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// const T & front()
func (this *QKeySequenceList) Front1() *QKeySequence {
	// QKeySequenceList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// T & back()
func (this *QKeySequenceList) Back0() *QKeySequence {
	// QKeySequenceList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// const T & back()
func (this *QKeySequenceList) Back1() *QKeySequence {
	// QKeySequenceList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QKeySequence{}
}

// void pop_front()
func (this *QKeySequenceList) Pop_front0() {
	// QKeySequenceList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QKeySequenceList) Pop_back0() {
	// QKeySequenceList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QKeySequenceList) Empty0() bool {
	// QKeySequenceList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QKeySequenceList) Operator_add_equal0() *QKeySequenceList {
	// QKeySequenceList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QKeySequenceList) Operator_add0() *QKeySequenceList {
	// QKeySequenceList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QKeySequenceList) Operator_add_equal1() *QKeySequenceList {
	// QKeySequenceList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QKeySequenceList) Operator_left_shift0() *QKeySequenceList {
	// QKeySequenceList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QKeySequenceList) Operator_left_shift1() *QKeySequenceList {
	// QKeySequenceList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QKeySequenceList) ToVector0() {
	// QKeySequenceList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QKeySequenceList) ToSet0() {
	// QKeySequenceList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QKeySequenceList) FromVector0() *QKeySequenceList {
	// QKeySequenceList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QKeySequenceList) FromSet0() *QKeySequenceList {
	// QKeySequenceList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QKeySequenceList) FromStdList0() *QKeySequenceList {
	// QKeySequenceList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QKeySequenceList) ToStdList0() {
	// QKeySequenceList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QKeySequenceList) Detach_helper_grow0() {
	// QKeySequenceList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QKeySequenceList) Detach_helper0() {
	// QKeySequenceList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QKeySequenceList) Detach_helper1() {
	// QKeySequenceList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QKeySequenceList) Dealloc0() {
	// QKeySequenceList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QKeySequenceList) Node_construct0() {
	// QKeySequenceList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QKeySequenceList) Node_destruct0() {
	// QKeySequenceList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QKeySequenceList) Node_copy0() {
	// QKeySequenceList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QKeySequenceList) Node_destruct1() {
	// QKeySequenceList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QKeySequenceList) IsValidIterator0() bool {
	// QKeySequenceList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QKeySequenceList) Op_eq_impl0() bool {
	// QKeySequenceList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QKeySequenceList) Op_eq_impl1() bool {
	// QKeySequenceList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QKeySequenceList) Contains_impl0() bool {
	// QKeySequenceList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QKeySequenceList) Contains_impl1() bool {
	// QKeySequenceList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QKeySequenceList) Count_impl0() int {
	// QKeySequenceList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QKeySequenceList) Count_impl1() int {
	// QKeySequenceList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QKeySequenceList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
