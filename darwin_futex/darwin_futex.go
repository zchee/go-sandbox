package main

/*
#include "ulock.h"
#include "darwin_futex.h"

#include <unistd.h>    // for close(), etc.
#include <time.h>

// The private syscalls. Normally declared in <sys/ulock.h>.
// Some toolchains require you to explicitly declare them.
extern int __ulock_wait(uint32_t operation, void *addr, uint64_t value, uint32_t timeout);
extern int __ulock_wake(uint32_t operation, void *addr, uint64_t wake_value);

// Helper: Convert struct timespec (relative) to microseconds.
// If 'timeout' is NULL, we return 0 (meaning "no timeout" to __ulock_wait).
static uint32_t timespec_to_microseconds(const struct timespec *timeout)
{
    if (!timeout) {
        return 0;  // no timeout
    }
    // Convert seconds + nanoseconds => microseconds
    uint64_t us = (uint64_t)timeout->tv_sec * 1000000ULL +
                  (uint64_t)timeout->tv_nsec / 1000ULL;

    // If large, clamp to max 32-bit
    if (us > 0xFFFFFFFFULL) {
        us = 0xFFFFFFFFULL;
    }
    return (uint32_t)us;
}

// darwin_futex_wait32:
//   Emulate "futex(uaddr, FUTEX_WAIT, val, timeout)" on Linux.
//   Only sleep if *uaddr == val. If mismatch, returns -EAGAIN immediately.
//   If timed out, returns -ETIMEDOUT. If woken, returns 0.
int darwin_futex_wait32(volatile _Atomic uint32_t *uaddr,
                        uint32_t val,
                        const struct timespec *timeout)
{
    // 1) Check if memory still matches 'val'. If not, return EAGAIN.
    //    (This is the typical futex "val mismatch" check.)
    if (__atomic_load_n(&uaddr, __ATOMIC_RELAXED) != val) {
        return -EAGAIN;
    }

    // 2) Prepare arguments for __ulock_wait
    uint32_t microseconds = timespec_to_microseconds(timeout);
    // The 'value' parameter for UL_COMPARE_AND_WAIT is the 32-bit expected value.
    uint64_t compare_value = (uint64_t)val;
    uint32_t op = UL_COMPARE_AND_WAIT | ULF_NO_ERRNO;  // no errno; returns errors directly

    // 3) Call the private syscall
    int rc = __ulock_wait(op, (void *)uaddr, compare_value, microseconds);
    if (rc == 0) {
        // Woken up normally
        return 0;
    }
    // __ulock_wait returns negative errno or a positive value?
    // On newer SDKs, it returns -errno directly if ULF_NO_ERRNO is used.
    // But be mindful of different behaviors in older OS versions.
    return rc;  // e.g. -EINTR, -ETIMEDOUT, etc.
}

// darwin_futex_wake32:
//   Emulate "futex(uaddr, FUTEX_WAKE, wake_count)" on Linux.
//   Wakes up to 'wake_count' waiters that did UL_COMPARE_AND_WAIT on *uaddr.
//
//   Returns number of threads actually woken, or negative errno on error.
int darwin_futex_wake32(volatile _Atomic(uint32_t) uaddr,
                        int wake_count)
{
    // Apple’s doc for __ulock_wake says the "wake_value" parameter is:
    //   - 0 = wake one
    //   - >0 = wake up to that many waiters
    //   - ~0ULL = wake all
    uint64_t wake_value;
    if (wake_count <= 0) {
        // Typically means "wake one" if wake_count == 1 on Linux,
        // but if user passes 0 or negative, we do nothing.
        return 0;
    } else if (wake_count >= INT32_MAX) {
        // "wake all"
        wake_value = (uint64_t)(-1); // ~0ULL
    } else {
        wake_value = (uint64_t)wake_count;
    }

    uint32_t op = UL_COMPARE_AND_WAIT | ULF_NO_ERRNO;
    int rc = __ulock_wake(op, (void *)uaddr, wake_value);
    if (rc < 0) {
        // If using ULF_NO_ERRNO, returns negative errno, e.g. -EINVAL.
        return rc;
    }
    // rc should be the number of threads woken
    return rc;
}

#include <stdio.h>
#include <pthread.h>
#include <stdatomic.h>
#include <unistd.h>

static volatile _Atomic(uint32_t) g_shared = 0;

void *worker_thread(void *arg)
{
    (void)arg;

    // Wait until g_shared == 1, or time out after ~2 seconds
    struct timespec ts;
    ts.tv_sec = 2;
    ts.tv_nsec = 0;

    int ret = darwin_futex_wait32(&g_shared, 0, &ts);
    if (ret == 0) {
        printf("Worker: Woken up!\n");
    } else {
        printf("Worker: darwin_futex_wait32 returned %d\n", ret);
    }
    return NULL;
}

int Main(void)
{
    pthread_t thread;
    pthread_create(&thread, NULL, worker_thread, NULL);

    // Sleep a bit, then set g_shared = 1 and wake
    sleep(1);
    atomic_store(&g_shared, 1);  // new value
    printf("Main: Setting g_shared=1 and waking futex\n");

    int num_woken = darwin_futex_wake32(g_shared, 1);
    if (num_woken < 0) {
        printf("Wake returned error: %d\n", num_woken);
    } else {
        printf("Woke %d thread(s)\n", num_woken);
    }

    pthread_join(thread, NULL);
    return 0;
}
*/
import "C"

func main() {
	C.Main()
}
