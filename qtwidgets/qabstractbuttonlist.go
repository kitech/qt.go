// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qabstractbutton.h
// #include <qabstractbutton.h>
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

func init_unused_10127() {
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
type QAbstractButtonList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QAbstractButtonList) Operator_equal0() *QAbstractButtonList {
	// QAbstractButtonList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QAbstractButtonList) Operator_equal1() *QAbstractButtonList {
	// QAbstractButtonList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QAbstractButtonList) Swap0() {
	// QAbstractButtonList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QAbstractButtonList) Operator_equal_equal0() bool {
	// QAbstractButtonList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QAbstractButtonList) Operator_not_equal0() bool {
	// QAbstractButtonList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QAbstractButtonList) Size0() int {
	// QAbstractButtonList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QAbstractButtonList) Detach0() {
	// QAbstractButtonList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QAbstractButtonList) DetachShared0() {
	// QAbstractButtonList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QAbstractButtonList) IsDetached0() bool {
	// QAbstractButtonList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QAbstractButtonList) SetSharable0() {
	// QAbstractButtonList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QAbstractButtonList) IsSharedWith0() bool {
	// QAbstractButtonList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QAbstractButtonList) IsEmpty0() bool {
	// QAbstractButtonList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QAbstractButtonList) Clear0() {
	// QAbstractButtonList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QAbstractButtonList) At0() *QAbstractButton {
	// QAbstractButtonList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & operator[](int)
func (this *QAbstractButtonList) Operator_get_index0() *QAbstractButton {
	// QAbstractButtonList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T & operator[](int)
func (this *QAbstractButtonList) Operator_get_index1() *QAbstractButton {
	// QAbstractButtonList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void reserve(int)
func (this *QAbstractButtonList) Reserve0() {
	// QAbstractButtonList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QAbstractButtonList) Append0() {
	// QAbstractButtonList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QAbstractButtonList) Append1() {
	// QAbstractButtonList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QAbstractButtonList) Prepend0() {
	// QAbstractButtonList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QAbstractButtonList) Insert0() {
	// QAbstractButtonList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QAbstractButtonList) Replace0() {
	// QAbstractButtonList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QAbstractButtonList) RemoveAt0() {
	// QAbstractButtonList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QAbstractButtonList) RemoveAll0() int {
	// QAbstractButtonList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QAbstractButtonList) RemoveOne0() bool {
	// QAbstractButtonList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QAbstractButtonList) TakeAt0() *QAbstractButton {
	// QAbstractButtonList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T takeFirst()
func (this *QAbstractButtonList) TakeFirst0() *QAbstractButton {
	// QAbstractButtonList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T takeLast()
func (this *QAbstractButtonList) TakeLast0() *QAbstractButton {
	// QAbstractButtonList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void move(int, int)
func (this *QAbstractButtonList) Move0() {
	// QAbstractButtonList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QAbstractButtonList) Swap1() {
	// QAbstractButtonList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QAbstractButtonList) IndexOf0() int {
	// QAbstractButtonList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QAbstractButtonList) LastIndexOf0() int {
	// QAbstractButtonList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QAbstractButtonList) Contains0() bool {
	// QAbstractButtonList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QAbstractButtonList) Count0() int {
	// QAbstractButtonList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QAbstractButtonList) Begin0() {
	// QAbstractButtonList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QAbstractButtonList) Begin1() {
	// QAbstractButtonList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QAbstractButtonList) Cbegin0() {
	// QAbstractButtonList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QAbstractButtonList) ConstBegin0() {
	// QAbstractButtonList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QAbstractButtonList) End0() {
	// QAbstractButtonList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QAbstractButtonList) End1() {
	// QAbstractButtonList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QAbstractButtonList) Cend0() {
	// QAbstractButtonList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QAbstractButtonList) ConstEnd0() {
	// QAbstractButtonList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QAbstractButtonList) Rbegin0() {
	// QAbstractButtonList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QAbstractButtonList) Rend0() {
	// QAbstractButtonList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QAbstractButtonList) Rbegin1() {
	// QAbstractButtonList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QAbstractButtonList) Rend1() {
	// QAbstractButtonList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QAbstractButtonList) Crbegin0() {
	// QAbstractButtonList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QAbstractButtonList) Crend0() {
	// QAbstractButtonList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QAbstractButtonList) Insert1() {
	// QAbstractButtonList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QAbstractButtonList) Erase0() {
	// QAbstractButtonList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QAbstractButtonList) Erase1() {
	// QAbstractButtonList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QAbstractButtonList) Count1() int {
	// QAbstractButtonList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QAbstractButtonList) Length0() int {
	// QAbstractButtonList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QAbstractButtonList) First0() *QAbstractButton {
	// QAbstractButtonList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & constFirst()
func (this *QAbstractButtonList) ConstFirst0() *QAbstractButton {
	// QAbstractButtonList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & first()
func (this *QAbstractButtonList) First1() *QAbstractButton {
	// QAbstractButtonList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T & last()
func (this *QAbstractButtonList) Last0() *QAbstractButton {
	// QAbstractButtonList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & last()
func (this *QAbstractButtonList) Last1() *QAbstractButton {
	// QAbstractButtonList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & constLast()
func (this *QAbstractButtonList) ConstLast0() *QAbstractButton {
	// QAbstractButtonList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void removeFirst()
func (this *QAbstractButtonList) RemoveFirst0() {
	// QAbstractButtonList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QAbstractButtonList) RemoveLast0() {
	// QAbstractButtonList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QAbstractButtonList) StartsWith0() bool {
	// QAbstractButtonList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QAbstractButtonList) EndsWith0() bool {
	// QAbstractButtonList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QAbstractButtonList) Mid0() *QAbstractButtonList {
	// QAbstractButtonList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QAbstractButtonList) Value0() *QAbstractButton {
	// QAbstractButtonList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T value(int, const T &)
func (this *QAbstractButtonList) Value1() *QAbstractButton {
	// QAbstractButtonList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void push_back(const T &)
func (this *QAbstractButtonList) Push_back0() {
	// QAbstractButtonList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QAbstractButtonList) Push_front0() {
	// QAbstractButtonList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QAbstractButtonList) Front0() *QAbstractButton {
	// QAbstractButtonList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & front()
func (this *QAbstractButtonList) Front1() *QAbstractButton {
	// QAbstractButtonList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// T & back()
func (this *QAbstractButtonList) Back0() *QAbstractButton {
	// QAbstractButtonList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// const T & back()
func (this *QAbstractButtonList) Back1() *QAbstractButton {
	// QAbstractButtonList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractButton{}
}

// void pop_front()
func (this *QAbstractButtonList) Pop_front0() {
	// QAbstractButtonList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QAbstractButtonList) Pop_back0() {
	// QAbstractButtonList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QAbstractButtonList) Empty0() bool {
	// QAbstractButtonList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QAbstractButtonList) Operator_add_equal0() *QAbstractButtonList {
	// QAbstractButtonList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QAbstractButtonList) Operator_add0() *QAbstractButtonList {
	// QAbstractButtonList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QAbstractButtonList) Operator_add_equal1() *QAbstractButtonList {
	// QAbstractButtonList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QAbstractButtonList) Operator_left_shift0() *QAbstractButtonList {
	// QAbstractButtonList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QAbstractButtonList) Operator_left_shift1() *QAbstractButtonList {
	// QAbstractButtonList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QAbstractButtonList) ToVector0() {
	// QAbstractButtonList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QAbstractButtonList) ToSet0() {
	// QAbstractButtonList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QAbstractButtonList) FromVector0() *QAbstractButtonList {
	// QAbstractButtonList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QAbstractButtonList) FromSet0() *QAbstractButtonList {
	// QAbstractButtonList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QAbstractButtonList) FromStdList0() *QAbstractButtonList {
	// QAbstractButtonList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QAbstractButtonList) ToStdList0() {
	// QAbstractButtonList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QAbstractButtonList) Detach_helper_grow0() {
	// QAbstractButtonList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QAbstractButtonList) Detach_helper0() {
	// QAbstractButtonList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QAbstractButtonList) Detach_helper1() {
	// QAbstractButtonList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QAbstractButtonList) Dealloc0() {
	// QAbstractButtonList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QAbstractButtonList) Node_construct0() {
	// QAbstractButtonList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QAbstractButtonList) Node_destruct0() {
	// QAbstractButtonList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QAbstractButtonList) Node_copy0() {
	// QAbstractButtonList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QAbstractButtonList) Node_destruct1() {
	// QAbstractButtonList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QAbstractButtonList) IsValidIterator0() bool {
	// QAbstractButtonList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractButtonList) Op_eq_impl0() bool {
	// QAbstractButtonList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QAbstractButtonList) Op_eq_impl1() bool {
	// QAbstractButtonList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractButtonList) Contains_impl0() bool {
	// QAbstractButtonList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAbstractButtonList) Contains_impl1() bool {
	// QAbstractButtonList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractButtonList) Count_impl0() int {
	// QAbstractButtonList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAbstractButtonList) Count_impl1() int {
	// QAbstractButtonList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractButtonList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
