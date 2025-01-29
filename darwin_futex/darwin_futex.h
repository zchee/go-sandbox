#ifndef DARWIN_FUTEX_H
#define DARWIN_FUTEX_H

#include <stdint.h>
#include <sys/time.h>     // for struct timespec
#include <errno.h>
#include "ulock.h"    // for __ulock_wait, __ulock_wake, UL_COMPARE_AND_WAIT, etc.

// Apple’s ulock API “operation” constants (32-bit compare-and-wait)
#ifndef UL_COMPARE_AND_WAIT
#define UL_COMPARE_AND_WAIT             0x00000002
#endif

// Typically added to the operation to say "don’t set errno, return errors directly"
#ifndef ULF_NO_ERRNO
#define ULF_NO_ERRNO                    0x01000000
#endif

#ifdef __cplusplus
extern "C" {
#endif

// darwin_futex_wait32:
//   Wait if *uaddr still equals 'val'.
//   If 'timeout' is non-null, treat it as a relative timeout (in ns).
//
// Returns 0 on success, or a negative errno code on failure
// (e.g. -EAGAIN if the value does not match, -ETIMEDOUT, etc.).
int darwin_futex_wait32(volatile _Atomic uint32_t *uaddr,
                        uint32_t val,
                        const struct timespec *timeout);

// darwin_futex_wake32:
//   Wake up to 'wake_count' waiters that called darwin_futex_wait32 on *uaddr.
//
// Returns the number of threads woken on success, or a negative errno code.
int darwin_futex_wake32(volatile _Atomic(uint32_t) uaddr,
                        int wake_count);

#ifdef __cplusplus
}
#endif
#endif
