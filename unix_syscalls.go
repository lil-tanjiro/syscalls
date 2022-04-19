package syscalls

import "C"
import (
	"fmt"
	"unsafe"
)

//#include <semaphore.h>
//#include <errno.h>
//#include <sys/ipc.h>
//#include <sys/shm.h>
//
//char* get_error_msg()
//{
//	return strerror(errno);
//}
import "C"
import (
	"errors"
)

// Semaphore

func UnixInitUnnamedSem(addr unsafe.Pointer, pshared int, value int) error {
	var (
		res C.int
	)

	res = C.sem_init((*C.sem_t)(addr), C.int(pshared), C.uint(value))

	if res == -1 {
		return errors.New(C.GoString(C.get_error_msg()))
	}

	return nil
}

func UnixSemWait(addr unsafe.Pointer) error {
	var (
		res C.int
	)

	res = C.sem_wait((*C.sem_t)(addr))

	if res == -1 {
		return errors.New(C.GoString(C.get_error_msg()))
	}

	return nil
}

func UnixSemPost(addr unsafe.Pointer) error {
	var (
		res C.int
	)

	res = C.sem_post((*C.sem_t)(addr))

	if res == -1 {
		return errors.New(C.GoString(C.get_error_msg()))
	}

	return nil
}

func UnixSemDestory(addr unsafe.Pointer) {
	C.sem_destroy((*C.sem_t)(addr))
}

// SharedMemory

func UnixShmGet(file string, id int, size uint64) (int, error) {
	var (
		key   C.key_t
		shmId C.int
	)

	key = C.ftok(C.CString(file), C.int(id))

	if key == -1 {
		return 0, errors.New(C.GoString(C.get_error_msg()))
	}

	shmId = C.shmget(key, C.size_t(size), 0666|C.IPC_CREAT)

	if shmId == -1 {
		fmt.Println()

		return 0, errors.New(C.GoString(C.get_error_msg()))
	}

	return int(shmId), nil
}

func UnixShmAttach(id int) (unsafe.Pointer, error) {
	var (
		r unsafe.Pointer
	)

	r = C.shmat(C.int(id), unsafe.Pointer(nil), 0)

	if uintptr(r) <= 0 {
		return unsafe.Pointer(nil), errors.New(C.GoString(C.get_error_msg()))
	}

	return r, nil
}
