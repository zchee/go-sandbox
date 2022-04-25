#include <stddef.h>
#include <stdint.h>
#include <stdbool.h>

#include <x86intrin.h>

/***************
* Taken more or less as-is from the chromium project
****************/

/**
 * \file
 * <PRE>
 * High performance base64 encoder / decoder
 * Version 1.3 -- 17-Mar-2006
 *
 * Copyright &copy; 2005, 2006, Nick Galbreath -- nickg [at] modp [dot] com
 * All rights reserved.
 *
 * http://modp.com/release/base64
 *
 * Released under bsd license.  See modp_b64.c for details.
 * </pre>
 *
 * The default implementation is the standard b64 encoding with padding.
 * It's easy to change this to use "URL safe" characters and to remove
 * padding.  See the modp_b64.c source code for details.
 *
 */

#ifndef MODP_B64
#define MODP_B64

#ifdef __cplusplus
extern "C" {
#endif

#define MODP_B64_ERROR ((size_t)-1)
/**
 * Encode a raw binary string into base 64.
 * src contains the bytes
 * len contains the number of bytes in the src
 * dest should be allocated by the caller to contain
 *   at least chromium_base64_encode_len(len) bytes (see below)
 *   This will contain the null-terminated b64 encoded result
 * returns length of the destination string plus the ending null byte
 *    i.e.  the result will be equal to strlen(dest) + 1
 *
 * Example
 *
 * \code
 * char* src = ...;
 * int srclen = ...; //the length of number of bytes in src
 * char* dest = (char*) malloc(chromium_base64_encode_len(srclen));
 * int len = chromium_base64_encode(dest, src, sourcelen);
 * if (len == MODP_B64_ERROR) {
 *   printf("Error\n");
 * } else {
 *   printf("b64 = %s\n", dest);
 * }
 * \endcode
 *
 */
extern size_t chromium_base64_encode(char* dest, const char* str, size_t len);

/**
 * Decode a base64 encoded string
 *
 *
 * src should contain exactly len bytes of b64 characters.
 *     if src contains -any- non-base characters (such as white
 *     space, MODP_B64_ERROR is returned.
 *
 * dest should be allocated by the caller to contain at least
 *    len * 3 / 4 bytes.
 *
 * Returns the length (strlen) of the output, or MODP_B64_ERROR if unable to
 * decode
 *
 * \code
 * char* src = ...;
 * int srclen = ...; // or if you don't know use strlen(src)
 * char* dest = (char*) malloc(chromium_base64_decode_len(srclen));
 * int len = chromium_base64_decode(dest, src, sourcelen);
 * if (len == MODP_B64_ERROR) { error }
 * \endcode
 */
size_t chromium_base64_decode(char* dest, const char* src, size_t len);

/**
 * Given a source string of length len, this returns the amount of
 * memory the destination string should have.
 *
 * remember, this is integer math
 * 3 bytes turn into 4 chars
 * ceiling[len / 3] * 4 + 1
 *
 * +1 is for any extra null.
 */
#define chromium_base64_encode_len(A) ((A+2)/3 * 4 + 1)

/**
 * Given a base64 string of length len,
 *   this returns the amount of memory required for output string
 *  It maybe be more than the actual number of bytes written.
 * NOTE: remember this is integer math
 * this allocates a bit more memory than traditional versions of b64
 * decode  4 chars turn into 3 bytes
 * floor[len * 3/4] + 2
 */
#define chromium_base64_decode_len(A) (A / 4 * 3 + 2)

/**
 * Will return the strlen of the output from encoding.
 * This may be less than the required number of bytes allocated.
 *
 * This allows you to 'deserialized' a struct
 * \code
 * char* b64encoded = "...";
 * int len = strlen(b64encoded);
 *
 * struct datastuff foo;
 * if (chromium_base64_encode_strlen(sizeof(struct datastuff)) != len) {
 *    // wrong size
 *    return false;
 * } else {
 *    // safe to do;
 *    if (chromium_base64_encode((char*) &foo, b64encoded, len) == MODP_B64_ERROR) {
 *      // bad characters
 *      return false;
 *    }
 * }
 * // foo is filled out now
 * \endcode
 */
#define chromium_base64_encode_strlen(A) ((A + 2)/ 3 * 4)



#ifdef __cplusplus
}

#include <string>


/**
 * base 64 decode a string (self-modifing)
 * On failure, the string is empty.
 *
 * This function is for C++ only (duh)
 *
 * \param[in,out] s the string to be decoded
 * \return a reference to the input string
 */
extern std::string& chromium_base64_encode(std::string& s)
{
    std::string x(chromium_base64_encode_len(s.size()), '\0');
    size_t d = chromium_base64_encode(const_cast<char*>(x.data()), s.data(), (int)s.size());
    if (d == MODP_B64_ERROR) {
        x.clear();
    } else {
        x.erase(d, std::string::npos);
    }
    s.swap(x);
    return s;
}

#endif /* __cplusplus */
#endif

#ifdef __cplusplus
extern "C" {
#endif /* __cplusplus */

size_t fast_avx512bw_base64_decode(char *out, const char *src, size_t srclen);
size_t fast_avx512bw_base64_encode(char* dest, const char* str, size_t len);

#ifdef __cplusplus
}
#endif /* __cplusplus */

#define CHAR62 '+'
#define CHAR63 '/'
#define CHARPAD '='
static const char e0[256] = {
 'A',  'A',  'A',  'A',  'B',  'B',  'B',  'B',  'C',  'C',
 'C',  'C',  'D',  'D',  'D',  'D',  'E',  'E',  'E',  'E',
 'F',  'F',  'F',  'F',  'G',  'G',  'G',  'G',  'H',  'H',
 'H',  'H',  'I',  'I',  'I',  'I',  'J',  'J',  'J',  'J',
 'K',  'K',  'K',  'K',  'L',  'L',  'L',  'L',  'M',  'M',
 'M',  'M',  'N',  'N',  'N',  'N',  'O',  'O',  'O',  'O',
 'P',  'P',  'P',  'P',  'Q',  'Q',  'Q',  'Q',  'R',  'R',
 'R',  'R',  'S',  'S',  'S',  'S',  'T',  'T',  'T',  'T',
 'U',  'U',  'U',  'U',  'V',  'V',  'V',  'V',  'W',  'W',
 'W',  'W',  'X',  'X',  'X',  'X',  'Y',  'Y',  'Y',  'Y',
 'Z',  'Z',  'Z',  'Z',  'a',  'a',  'a',  'a',  'b',  'b',
 'b',  'b',  'c',  'c',  'c',  'c',  'd',  'd',  'd',  'd',
 'e',  'e',  'e',  'e',  'f',  'f',  'f',  'f',  'g',  'g',
 'g',  'g',  'h',  'h',  'h',  'h',  'i',  'i',  'i',  'i',
 'j',  'j',  'j',  'j',  'k',  'k',  'k',  'k',  'l',  'l',
 'l',  'l',  'm',  'm',  'm',  'm',  'n',  'n',  'n',  'n',
 'o',  'o',  'o',  'o',  'p',  'p',  'p',  'p',  'q',  'q',
 'q',  'q',  'r',  'r',  'r',  'r',  's',  's',  's',  's',
 't',  't',  't',  't',  'u',  'u',  'u',  'u',  'v',  'v',
 'v',  'v',  'w',  'w',  'w',  'w',  'x',  'x',  'x',  'x',
 'y',  'y',  'y',  'y',  'z',  'z',  'z',  'z',  '0',  '0',
 '0',  '0',  '1',  '1',  '1',  '1',  '2',  '2',  '2',  '2',
 '3',  '3',  '3',  '3',  '4',  '4',  '4',  '4',  '5',  '5',
 '5',  '5',  '6',  '6',  '6',  '6',  '7',  '7',  '7',  '7',
 '8',  '8',  '8',  '8',  '9',  '9',  '9',  '9',  '+',  '+',
 '+',  '+',  '/',  '/',  '/',  '/'
};

static const char e1[256] = {
 'A',  'B',  'C',  'D',  'E',  'F',  'G',  'H',  'I',  'J',
 'K',  'L',  'M',  'N',  'O',  'P',  'Q',  'R',  'S',  'T',
 'U',  'V',  'W',  'X',  'Y',  'Z',  'a',  'b',  'c',  'd',
 'e',  'f',  'g',  'h',  'i',  'j',  'k',  'l',  'm',  'n',
 'o',  'p',  'q',  'r',  's',  't',  'u',  'v',  'w',  'x',
 'y',  'z',  '0',  '1',  '2',  '3',  '4',  '5',  '6',  '7',
 '8',  '9',  '+',  '/',  'A',  'B',  'C',  'D',  'E',  'F',
 'G',  'H',  'I',  'J',  'K',  'L',  'M',  'N',  'O',  'P',
 'Q',  'R',  'S',  'T',  'U',  'V',  'W',  'X',  'Y',  'Z',
 'a',  'b',  'c',  'd',  'e',  'f',  'g',  'h',  'i',  'j',
 'k',  'l',  'm',  'n',  'o',  'p',  'q',  'r',  's',  't',
 'u',  'v',  'w',  'x',  'y',  'z',  '0',  '1',  '2',  '3',
 '4',  '5',  '6',  '7',  '8',  '9',  '+',  '/',  'A',  'B',
 'C',  'D',  'E',  'F',  'G',  'H',  'I',  'J',  'K',  'L',
 'M',  'N',  'O',  'P',  'Q',  'R',  'S',  'T',  'U',  'V',
 'W',  'X',  'Y',  'Z',  'a',  'b',  'c',  'd',  'e',  'f',
 'g',  'h',  'i',  'j',  'k',  'l',  'm',  'n',  'o',  'p',
 'q',  'r',  's',  't',  'u',  'v',  'w',  'x',  'y',  'z',
 '0',  '1',  '2',  '3',  '4',  '5',  '6',  '7',  '8',  '9',
 '+',  '/',  'A',  'B',  'C',  'D',  'E',  'F',  'G',  'H',
 'I',  'J',  'K',  'L',  'M',  'N',  'O',  'P',  'Q',  'R',
 'S',  'T',  'U',  'V',  'W',  'X',  'Y',  'Z',  'a',  'b',
 'c',  'd',  'e',  'f',  'g',  'h',  'i',  'j',  'k',  'l',
 'm',  'n',  'o',  'p',  'q',  'r',  's',  't',  'u',  'v',
 'w',  'x',  'y',  'z',  '0',  '1',  '2',  '3',  '4',  '5',
 '6',  '7',  '8',  '9',  '+',  '/'
};

static const char e2[256] = {
 'A',  'B',  'C',  'D',  'E',  'F',  'G',  'H',  'I',  'J',
 'K',  'L',  'M',  'N',  'O',  'P',  'Q',  'R',  'S',  'T',
 'U',  'V',  'W',  'X',  'Y',  'Z',  'a',  'b',  'c',  'd',
 'e',  'f',  'g',  'h',  'i',  'j',  'k',  'l',  'm',  'n',
 'o',  'p',  'q',  'r',  's',  't',  'u',  'v',  'w',  'x',
 'y',  'z',  '0',  '1',  '2',  '3',  '4',  '5',  '6',  '7',
 '8',  '9',  '+',  '/',  'A',  'B',  'C',  'D',  'E',  'F',
 'G',  'H',  'I',  'J',  'K',  'L',  'M',  'N',  'O',  'P',
 'Q',  'R',  'S',  'T',  'U',  'V',  'W',  'X',  'Y',  'Z',
 'a',  'b',  'c',  'd',  'e',  'f',  'g',  'h',  'i',  'j',
 'k',  'l',  'm',  'n',  'o',  'p',  'q',  'r',  's',  't',
 'u',  'v',  'w',  'x',  'y',  'z',  '0',  '1',  '2',  '3',
 '4',  '5',  '6',  '7',  '8',  '9',  '+',  '/',  'A',  'B',
 'C',  'D',  'E',  'F',  'G',  'H',  'I',  'J',  'K',  'L',
 'M',  'N',  'O',  'P',  'Q',  'R',  'S',  'T',  'U',  'V',
 'W',  'X',  'Y',  'Z',  'a',  'b',  'c',  'd',  'e',  'f',
 'g',  'h',  'i',  'j',  'k',  'l',  'm',  'n',  'o',  'p',
 'q',  'r',  's',  't',  'u',  'v',  'w',  'x',  'y',  'z',
 '0',  '1',  '2',  '3',  '4',  '5',  '6',  '7',  '8',  '9',
 '+',  '/',  'A',  'B',  'C',  'D',  'E',  'F',  'G',  'H',
 'I',  'J',  'K',  'L',  'M',  'N',  'O',  'P',  'Q',  'R',
 'S',  'T',  'U',  'V',  'W',  'X',  'Y',  'Z',  'a',  'b',
 'c',  'd',  'e',  'f',  'g',  'h',  'i',  'j',  'k',  'l',
 'm',  'n',  'o',  'p',  'q',  'r',  's',  't',  'u',  'v',
 'w',  'x',  'y',  'z',  '0',  '1',  '2',  '3',  '4',  '5',
 '6',  '7',  '8',  '9',  '+',  '/'
};



/* SPECIAL DECODE TABLES FOR LITTLE ENDIAN (INTEL) CPUS */

static const uint32_t d0[256] = {
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x000000f8, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x000000fc,
0x000000d0, 0x000000d4, 0x000000d8, 0x000000dc, 0x000000e0, 0x000000e4,
0x000000e8, 0x000000ec, 0x000000f0, 0x000000f4, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x00000000,
0x00000004, 0x00000008, 0x0000000c, 0x00000010, 0x00000014, 0x00000018,
0x0000001c, 0x00000020, 0x00000024, 0x00000028, 0x0000002c, 0x00000030,
0x00000034, 0x00000038, 0x0000003c, 0x00000040, 0x00000044, 0x00000048,
0x0000004c, 0x00000050, 0x00000054, 0x00000058, 0x0000005c, 0x00000060,
0x00000064, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x00000068, 0x0000006c, 0x00000070, 0x00000074, 0x00000078,
0x0000007c, 0x00000080, 0x00000084, 0x00000088, 0x0000008c, 0x00000090,
0x00000094, 0x00000098, 0x0000009c, 0x000000a0, 0x000000a4, 0x000000a8,
0x000000ac, 0x000000b0, 0x000000b4, 0x000000b8, 0x000000bc, 0x000000c0,
0x000000c4, 0x000000c8, 0x000000cc, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff
};


static const uint32_t d1[256] = {
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x0000e003, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x0000f003,
0x00004003, 0x00005003, 0x00006003, 0x00007003, 0x00008003, 0x00009003,
0x0000a003, 0x0000b003, 0x0000c003, 0x0000d003, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x00000000,
0x00001000, 0x00002000, 0x00003000, 0x00004000, 0x00005000, 0x00006000,
0x00007000, 0x00008000, 0x00009000, 0x0000a000, 0x0000b000, 0x0000c000,
0x0000d000, 0x0000e000, 0x0000f000, 0x00000001, 0x00001001, 0x00002001,
0x00003001, 0x00004001, 0x00005001, 0x00006001, 0x00007001, 0x00008001,
0x00009001, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x0000a001, 0x0000b001, 0x0000c001, 0x0000d001, 0x0000e001,
0x0000f001, 0x00000002, 0x00001002, 0x00002002, 0x00003002, 0x00004002,
0x00005002, 0x00006002, 0x00007002, 0x00008002, 0x00009002, 0x0000a002,
0x0000b002, 0x0000c002, 0x0000d002, 0x0000e002, 0x0000f002, 0x00000003,
0x00001003, 0x00002003, 0x00003003, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff
};


static const uint32_t d2[256] = {
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x00800f00, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x00c00f00,
0x00000d00, 0x00400d00, 0x00800d00, 0x00c00d00, 0x00000e00, 0x00400e00,
0x00800e00, 0x00c00e00, 0x00000f00, 0x00400f00, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x00000000,
0x00400000, 0x00800000, 0x00c00000, 0x00000100, 0x00400100, 0x00800100,
0x00c00100, 0x00000200, 0x00400200, 0x00800200, 0x00c00200, 0x00000300,
0x00400300, 0x00800300, 0x00c00300, 0x00000400, 0x00400400, 0x00800400,
0x00c00400, 0x00000500, 0x00400500, 0x00800500, 0x00c00500, 0x00000600,
0x00400600, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x00800600, 0x00c00600, 0x00000700, 0x00400700, 0x00800700,
0x00c00700, 0x00000800, 0x00400800, 0x00800800, 0x00c00800, 0x00000900,
0x00400900, 0x00800900, 0x00c00900, 0x00000a00, 0x00400a00, 0x00800a00,
0x00c00a00, 0x00000b00, 0x00400b00, 0x00800b00, 0x00c00b00, 0x00000c00,
0x00400c00, 0x00800c00, 0x00c00c00, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff
};


static const uint32_t d3[256] = {
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x003e0000, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x003f0000,
0x00340000, 0x00350000, 0x00360000, 0x00370000, 0x00380000, 0x00390000,
0x003a0000, 0x003b0000, 0x003c0000, 0x003d0000, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x00000000,
0x00010000, 0x00020000, 0x00030000, 0x00040000, 0x00050000, 0x00060000,
0x00070000, 0x00080000, 0x00090000, 0x000a0000, 0x000b0000, 0x000c0000,
0x000d0000, 0x000e0000, 0x000f0000, 0x00100000, 0x00110000, 0x00120000,
0x00130000, 0x00140000, 0x00150000, 0x00160000, 0x00170000, 0x00180000,
0x00190000, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x001a0000, 0x001b0000, 0x001c0000, 0x001d0000, 0x001e0000,
0x001f0000, 0x00200000, 0x00210000, 0x00220000, 0x00230000, 0x00240000,
0x00250000, 0x00260000, 0x00270000, 0x00280000, 0x00290000, 0x002a0000,
0x002b0000, 0x002c0000, 0x002d0000, 0x002e0000, 0x002f0000, 0x00300000,
0x00310000, 0x00320000, 0x00330000, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff,
0x01ffffff, 0x01ffffff, 0x01ffffff, 0x01ffffff
};



#define BADCHAR 0x01FFFFFF

/**
 * you can control if we use padding by commenting out this
 * next line.  However, I highly recommend you use padding and not
 * using it should only be for compatability with a 3rd party.
 * Also, 'no padding' is not tested!
 */
#define DOPAD 1

/*
 * if we aren't doing padding
 * set the pad character to NULL
 */
#ifndef DOPAD
#undef CHARPAD
#define CHARPAD '\0'
#endif

extern size_t chromium_base64_encode(char* dest, const char* str, size_t len)
{
    size_t i = 0;
    uint8_t* p = (uint8_t*) dest;

    /* unsigned here is important! */
    uint8_t t1, t2, t3;

    if (len > 2) {
        for (; i < len - 2; i += 3) {
            t1 = str[i]; t2 = str[i+1]; t3 = str[i+2];
            *p++ = e0[t1];
            *p++ = e1[((t1 & 0x03) << 4) | ((t2 >> 4) & 0x0F)];
            *p++ = e1[((t2 & 0x0F) << 2) | ((t3 >> 6) & 0x03)];
            *p++ = e2[t3];
        }
    }

    switch (len - i) {
    case 0:
        break;
    case 1:
        t1 = str[i];
        *p++ = e0[t1];
        *p++ = e1[(t1 & 0x03) << 4];
        *p++ = CHARPAD;
        *p++ = CHARPAD;
        break;
    default: /* case 2 */
        t1 = str[i]; t2 = str[i+1];
        *p++ = e0[t1];
        *p++ = e1[((t1 & 0x03) << 4) | ((t2 >> 4) & 0x0F)];
        *p++ = e2[(t2 & 0x0F) << 2];
        *p++ = CHARPAD;
    }

    *p = '\0';
    return p - (uint8_t*)dest;
}


size_t chromium_base64_decode(char* dest, const char* src, size_t len)
{
    if (len == 0) return 0;

#ifdef DOPAD
    /*
     * if padding is used, then the message must be at least
     * 4 chars and be a multiple of 4
     */
    if (len < 4 || (len % 4 != 0)) {
      return MODP_B64_ERROR; /* error */
    }
    /* there can be at most 2 pad chars at the end */
    if (src[len-1] == CHARPAD) {
        len--;
        if (src[len -1] == CHARPAD) {
            len--;
        }
    }
#endif

    size_t i;
    int leftover = len % 4;
    size_t chunks = (leftover == 0) ? len / 4 - 1 : len /4;

    uint8_t* p = (uint8_t*)dest;
    uint32_t x = 0;
    const uint8_t* y = (uint8_t*)src;
    for (i = 0; i < chunks; ++i, y += 4) {
        x = d0[y[0]] | d1[y[1]] | d2[y[2]] | d3[y[3]];
        if (x >= BADCHAR) return MODP_B64_ERROR;
        *p++ =  ((uint8_t*)(&x))[0];
        *p++ =  ((uint8_t*)(&x))[1];
        *p++ =  ((uint8_t*)(&x))[2];
    }

    switch (leftover) {
    case 0:
        x = d0[y[0]] | d1[y[1]] | d2[y[2]] | d3[y[3]];

        if (x >= BADCHAR) return MODP_B64_ERROR;
        *p++ =  ((uint8_t*)(&x))[0];
        *p++ =  ((uint8_t*)(&x))[1];
        *p =    ((uint8_t*)(&x))[2];
        return (chunks+1)*3;
        break;
    case 1:  /* with padding this is an impossible case */
        x = d0[y[0]];
        *p = *((uint8_t*)(&x)); // i.e. first char/byte in int
        break;
    case 2: // * case 2, 1  output byte */
        x = d0[y[0]] | d1[y[1]];
        *p = *((uint8_t*)(&x)); // i.e. first char
        break;
    default: /* case 3, 2 output bytes */
        x = d0[y[0]] | d1[y[1]] | d2[y[2]];  /* 0x3c */
        *p++ =  ((uint8_t*)(&x))[0];
        *p =  ((uint8_t*)(&x))[1];
        break;
    }

    if (x >= BADCHAR) return MODP_B64_ERROR;

    return 3*chunks + (6*leftover)/8;
}

static inline __m512i enc_reshuffle(const __m512i input) {

    // from https://github.com/WojciechMula/base64simd/blob/master/encode/encode.avx512bw.cpp

    // place each 12-byte subarray in seprate 128-bit lane
    // tmp1 = [?? ?? ?? ??|D2 D1 D0 C2|C1 C0 B2 B1|B0 A2 A1 A0] x 4
    //           ignored
    const __m512i tmp1 = _mm512_permutexvar_epi32(
        _mm512_set_epi32(-1, 11, 10, 9, -1, 8, 7, 6, -1, 5, 4, 3, -1, 2, 1, 0),
        input
    );

    // reshuffle bytes within 128-bit lanes to format required by
    // AVX512BW unpack procedure
    // tmp2 = [D1 D2 D0 D1|C1 C2 C0 C1|B1 B2 B0 B1|A1 A2 A0 A1] x 4
    //         10 11 9  10 7  8  6  7  4  5  3  4  1  2  0  1
    const __m512i tmp2 = _mm512_shuffle_epi8(
        tmp1,
        _mm512_set4_epi32(0x0a0b090a, 0x07080607, 0x04050304, 0x01020001)
    );

    return tmp2;
}

static inline __m512i enc_translate(const __m512i input) {

    // from https://github.com/WojciechMula/base64simd/blob/master/encode/lookup.avx512bw.cpp

    // reduce  0..51 -> 0
    //        52..61 -> 1 .. 10
    //            62 -> 11
    //            63 -> 12
    __m512i result = _mm512_subs_epu8(input, _mm512_set1_epi8(51));

    // distinguish between ranges 0..25 and 26..51:
    //         0 .. 25 -> remains 0
    //        26 .. 51 -> becomes 13
    const __mmask64 less = _mm512_cmpgt_epi8_mask(_mm512_set1_epi8(26), input);
    result = _mm512_mask_mov_epi8(result, less, _mm512_set1_epi8(13));

    /* the SSE lookup
        const __m128i shift_LUT = _mm_setr_epi8(
            'a' - 26, '0' - 52, '0' - 52, '0' - 52, '0' - 52, '0' - 52,
            '0' - 52, '0' - 52, '0' - 52, '0' - 52, '0' - 52, '+' - 62,
            '/' - 63, 'A', 0, 0
        );
        which is:
            0x47, 0xfc, 0xfc, 0xfc,
            0xfc, 0xfc, 0xfc, 0xfc,
            0xfc, 0xfc, 0xfc, 0xed,
            0xf0, 0x41, 0x00, 0x00

        Note that the order of above list is reserved (due to _mm_setr_epi8),
        so the invocation _mm512_set4_epi32 looks... odd.
    */
    const __m512i shift_LUT = _mm512_set4_epi32(
        0x000041f0,
        0xedfcfcfc,
        0xfcfcfcfc,
        0xfcfcfc47
    );

    // read shift
    result = _mm512_shuffle_epi8(shift_LUT, result);

    return _mm512_add_epi8(result, input);
}

static inline __m512i dec_reshuffle(__m512i input) {

    // from https://github.com/WojciechMula/base64simd/blob/master/decode/pack.avx512bw.cpp
    const __m512i merge_ab_and_bc = _mm512_maddubs_epi16(input, _mm512_set1_epi32(0x01400140));

    return _mm512_madd_epi16(merge_ab_and_bc, _mm512_set1_epi32(0x00011000));
}


extern size_t fast_avx512bw_base64_encode(char* dest, const char* str, size_t len) {
    const char* const dest_orig = dest;
    __m512i inputvector;
    while (len >= 64) {
        inputvector = _mm512_loadu_si512((__m512i *)(str));
        inputvector = enc_reshuffle(inputvector);
        inputvector = enc_translate(inputvector);
        _mm512_storeu_si512((__m512i *)dest, inputvector);
        str  += 48;
        dest += 64;
        len -= 48;
    }
    size_t scalarret = chromium_base64_encode(dest, str, len);
    if(scalarret == MODP_B64_ERROR) return MODP_B64_ERROR;
    return (dest - dest_orig) + scalarret;
}

#define build_dword(b0, b1, b2, b3)     \
     (((uint32_t)((uint8_t)(b0)) << 0*8)    \
    | ((uint32_t)((uint8_t)(b1)) << 1*8)    \
    | ((uint32_t)((uint8_t)(b2)) << 2*8)    \
    | ((uint32_t)((uint8_t)(b3)) << 3*8))
    
#define _mm512_set4lanes_epi8(b0, b1, b2, b3, b4, b5, b6, b7, b8, b9, b10, b11, b12, b13, b14, b15) \
    _mm512_setr4_epi32(                  \
        build_dword( b0,  b1,  b2,  b3), \
        build_dword( b4,  b5,  b6,  b7), \
        build_dword( b8,  b9, b10, b11), \
        build_dword(b12, b13, b14, b15))

extern size_t fast_avx512bw_base64_decode(char *out, const char *src, size_t srclen) {
  char* out_orig = out;
  while (srclen >= 88) {

    // load
    const __m512i input = _mm512_loadu_si512((const __m512i*)(src));
        
    // translate from ASCII
    //  and https://github.com/WojciechMula/base64simd/blob/master/decode/lookup.avx512bw.cpp
    const __m512i higher_nibble = _mm512_and_si512(_mm512_srli_epi32(input, 4), _mm512_set1_epi8(0x0f));
    const __m512i lower_nibble  = _mm512_and_si512(input, _mm512_set1_epi8(0x0f));

    const __m512i shiftLUT = _mm512_set4lanes_epi8(
        0,   0,  19,   4, -65, -65, -71, -71,
        0,   0,   0,   0,   0,   0,   0,   0);

    const __m512i maskLUT  = _mm512_set4lanes_epi8(
        /* 0        : 0b1010_1000*/ 0xa8,
        /* 1 .. 9   : 0b1111_1000*/ 0xf8, 0xf8, 0xf8, 0xf8,
                                    0xf8, 0xf8, 0xf8, 0xf8,
                                    0xf8,
        /* 10       : 0b1111_0000*/ 0xf0,
        /* 11       : 0b0101_0100*/ 0x54,
        /* 12 .. 14 : 0b0101_0000*/ 0x50, 0x50, 0x50,
        /* 15       : 0b0101_0100*/ 0x54
    );

    const __m512i bitposLUT = _mm512_set4lanes_epi8(
        0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80,
        0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00
    );

    const __m512i   sh      = _mm512_shuffle_epi8(shiftLUT,  higher_nibble);
    const __mmask64 eq_2f   = _mm512_cmpeq_epi8_mask(input, _mm512_set1_epi8(0x2f));
    const __m512i   shift   = _mm512_mask_mov_epi8(sh, eq_2f, _mm512_set1_epi8(16));

    const __m512i M         = _mm512_shuffle_epi8(maskLUT,   lower_nibble);
    const __m512i bit       = _mm512_shuffle_epi8(bitposLUT, higher_nibble);

    const uint64_t match    = _mm512_test_epi8_mask(M, bit);

    if (match != (uint64_t)(-1)) {
        // some characters do not match the valid range
        return MODP_B64_ERROR;
    }

    const __m512i translated = _mm512_add_epi8(input, shift);

    const __m512i packed = dec_reshuffle(translated);

    //  and https://github.com/WojciechMula/base64simd/blob/master/decode/decode.avx512bw.cpp
    // 
    const __m512i t1 = _mm512_shuffle_epi8(
        packed,
        _mm512_set4lanes_epi8(
             2,  1,  0,
             6,  5,  4,
            10,  9,  8,
            14, 13, 12,
            -1, -1, -1, -1)
    );

    // shuffle bytes
    const __m512i s6 = _mm512_setr_epi32(
         0,  1,  2,
         4,  5,  6,
         8,  9, 10,
        12, 13, 14,
        // unused
         0,  0,  0, 0);

    const __m512i t2 = _mm512_permutexvar_epi32(s6, t1);

    _mm512_storeu_si512((__m512i*)(out), t2);

    srclen -= 64;
    src += 64;
    out += 48;
  }
  size_t scalarret = chromium_base64_decode(out, src, srclen);
  if(scalarret == MODP_B64_ERROR) return MODP_B64_ERROR;
  return (out - out_orig) + scalarret;
}
