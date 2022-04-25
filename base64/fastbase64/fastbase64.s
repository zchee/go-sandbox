	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 12, 0	sdk_version 12, 3
	.globl	_chromium_base64_encode         ## -- Begin function chromium_base64_encode
	.p2align	4, 0x90
_chromium_base64_encode:                ## @chromium_base64_encode
## %bb.0:
	pushq	%rbp
	movq	%rsp, %rbp
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	andq	$-8, %rsp
	xorl	%r11d, %r11d
	cmpq	$3, %rdx
	jb	LBB0_1
## %bb.2:
	movq	%rdx, %r8
	movq	%rdi, %rax
	addq	$-2, %r8
	je	LBB0_5
## %bb.3:
	xorl	%r11d, %r11d
	leaq	_e0(%rip), %r9
	leaq	_e2(%rip), %r10
	movq	%rdi, %rax
	.p2align	4, 0x90
LBB0_4:                                 ## =>This Inner Loop Header: Depth=1
	movzbl	(%rsi,%r11), %ecx
	movzbl	1(%rsi,%r11), %r14d
	movzbl	2(%rsi,%r11), %r15d
	movzbl	(%rcx,%r9), %ebx
	movb	%bl, (%rax)
	shll	$4, %ecx
	andl	$48, %ecx
	movq	%r14, %rbx
	shrq	$4, %rbx
	orq	%rcx, %rbx
	movzbl	(%r10,%rbx), %ecx
	movb	%cl, 1(%rax)
	andl	$15, %r14d
	movzbl	(%r15,%r10), %ecx
	shrq	$6, %r15
	leaq	(%r15,%r14,4), %rbx
	movzbl	(%r10,%rbx), %ebx
	movb	%bl, 2(%rax)
	movb	%cl, 3(%rax)
	addq	$4, %rax
	addq	$3, %r11
	cmpq	%r8, %r11
	jb	LBB0_4
LBB0_5:
	subq	%r11, %rdx
	je	LBB0_10
LBB0_6:
	cmpq	$1, %rdx
	jne	LBB0_8
## %bb.7:
	movzbl	(%rsi,%r11), %ecx
	leaq	_e0(%rip), %rdx
	movb	(%rcx,%rdx), %dl
	movb	%dl, (%rax)
	shll	$4, %ecx
	andl	$48, %ecx
	leaq	_e2(%rip), %rdx
	movb	(%rcx,%rdx), %cl
	movb	%cl, 1(%rax)
	movb	$61, 2(%rax)
	jmp	LBB0_9
LBB0_8:
	movzbl	(%rsi,%r11), %ecx
	movzbl	1(%rsi,%r11), %edx
	leaq	_e0(%rip), %rsi
	movb	(%rcx,%rsi), %bl
	movb	%bl, (%rax)
	shll	$4, %ecx
	andl	$48, %ecx
	movq	%rdx, %rsi
	shrq	$4, %rsi
	orq	%rcx, %rsi
	leaq	_e2(%rip), %rcx
	movb	(%rcx,%rsi), %bl
	movb	%bl, 1(%rax)
	andl	$15, %edx
	movb	(%rcx,%rdx,4), %cl
	movb	%cl, 2(%rax)
LBB0_9:
	movb	$61, 3(%rax)
	addq	$4, %rax
LBB0_10:
	movb	$0, (%rax)
	subq	%rdi, %rax
	leaq	-24(%rbp), %rsp
	popq	%rbx
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
LBB0_1:
	movq	%rdi, %rax
	subq	%r11, %rdx
	jne	LBB0_6
	jmp	LBB0_10
                                        ## -- End function
	.globl	_chromium_base64_decode         ## -- Begin function chromium_base64_decode
	.p2align	4, 0x90
_chromium_base64_decode:                ## @chromium_base64_decode
## %bb.0:
	pushq	%rbp
	movq	%rsp, %rbp
	pushq	%r15
	pushq	%r14
	pushq	%rbx
	andq	$-8, %rsp
	testq	%rdx, %rdx
	je	LBB1_1
## %bb.2:
	movq	$-1, %rax
	cmpq	$4, %rdx
	jb	LBB1_21
## %bb.3:
	movl	%edx, %ecx
	andl	$3, %ecx
	testq	%rcx, %rcx
	jne	LBB1_21
## %bb.4:
	cmpb	$61, -1(%rsi,%rdx)
	jne	LBB1_6
## %bb.5:
	leaq	-2(%rdx), %rcx
	cmpb	$61, -2(%rsi,%rdx)
	leaq	-1(%rdx), %rdx
	cmoveq	%rcx, %rdx
LBB1_6:
	xorl	%ecx, %ecx
	movl	%edx, %r9d
	andl	$3, %r9d
	sete	%cl
	shrq	$2, %rdx
	movq	%rdx, %r8
	subq	%rcx, %r8
	je	LBB1_10
## %bb.7:
	negq	%rcx
	addq	%rcx, %rdx
	leaq	_d0(%rip), %r10
	leaq	_d1(%rip), %r11
	leaq	_d2(%rip), %r14
	leaq	_d3(%rip), %r15
	.p2align	4, 0x90
LBB1_8:                                 ## =>This Inner Loop Header: Depth=1
	movzbl	(%rsi), %ebx
	movzbl	1(%rsi), %ecx
	movl	(%r11,%rcx,4), %ecx
	orl	(%r10,%rbx,4), %ecx
	movzbl	2(%rsi), %ebx
	orl	(%r14,%rbx,4), %ecx
	movzbl	3(%rsi), %ebx
	orl	(%r15,%rbx,4), %ecx
	cmpl	$33554430, %ecx                 ## imm = 0x1FFFFFE
	ja	LBB1_21
## %bb.9:                               ##   in Loop: Header=BB1_8 Depth=1
	movl	%ecx, %ebx
	shrl	$16, %ebx
	movb	%cl, (%rdi)
	movb	%ch, 1(%rdi)
	movb	%bl, 2(%rdi)
	addq	$3, %rdi
	addq	$4, %rsi
	decq	%rdx
	jne	LBB1_8
LBB1_10:
	cmpl	$2, %r9d
	je	LBB1_17
## %bb.11:
	cmpl	$1, %r9d
	je	LBB1_15
## %bb.12:
	testl	%r9d, %r9d
	jne	LBB1_18
## %bb.13:
	movzbl	(%rsi), %edx
	leaq	_d0(%rip), %r9
	movzbl	1(%rsi), %ecx
	leaq	_d1(%rip), %rbx
	movl	(%rbx,%rcx,4), %ecx
	orl	(%r9,%rdx,4), %ecx
	movzbl	2(%rsi), %edx
	leaq	_d2(%rip), %rbx
	orl	(%rbx,%rdx,4), %ecx
	movzbl	3(%rsi), %edx
	leaq	_d3(%rip), %rsi
	orl	(%rsi,%rdx,4), %ecx
	cmpl	$33554430, %ecx                 ## imm = 0x1FFFFFE
	ja	LBB1_21
## %bb.14:
	movl	%ecx, %eax
	shrl	$16, %eax
	movb	%cl, (%rdi)
	movb	%ch, 1(%rdi)
	movb	%al, 2(%rdi)
	leaq	(%r8,%r8,2), %rax
	addq	$3, %rax
	jmp	LBB1_21
LBB1_1:
	xorl	%eax, %eax
LBB1_21:
	leaq	-24(%rbp), %rsp
	popq	%rbx
	popq	%r14
	popq	%r15
	popq	%rbp
	retq
LBB1_17:
	movzbl	(%rsi), %edx
	leaq	_d0(%rip), %rbx
	movzbl	1(%rsi), %ecx
	leaq	_d1(%rip), %rsi
	movl	(%rsi,%rcx,4), %ecx
	orl	(%rbx,%rdx,4), %ecx
	jmp	LBB1_16
LBB1_15:
	movzbl	(%rsi), %ecx
	leaq	_d0(%rip), %rdx
	movl	(%rdx,%rcx,4), %ecx
LBB1_16:
	movl	%ecx, %edx
	shrl	$8, %edx
	movb	%cl, (%rdi)
	jmp	LBB1_19
LBB1_18:
	movzbl	(%rsi), %edx
	leaq	_d0(%rip), %r10
	movzbl	1(%rsi), %ecx
	leaq	_d1(%rip), %rbx
	movl	(%rbx,%rcx,4), %ecx
	orl	(%r10,%rdx,4), %ecx
	movzbl	2(%rsi), %edx
	leaq	_d2(%rip), %rsi
	orl	(%rsi,%rdx,4), %ecx
	movl	%ecx, %edx
	shrl	$8, %edx
	movw	%cx, (%rdi)
LBB1_19:
	shll	$8, %edx
	movzwl	%dx, %edx
	andl	$-65281, %ecx                   ## imm = 0xFFFF00FF
	orl	%edx, %ecx
	cmpl	$33554430, %ecx                 ## imm = 0x1FFFFFE
	ja	LBB1_21
## %bb.20:
	leaq	(%r8,%r8,2), %rcx
	addl	%r9d, %r9d
	leal	(%r9,%r9,2), %eax
	shrl	$3, %eax
	addq	%rcx, %rax
	jmp	LBB1_21
                                        ## -- End function
	.section	__TEXT,__const
	.p2align	6                               ## -- Begin function fast_avx512bw_base64_encode
LCPI2_0:
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	4                               ## 0x4
	.byte	3                               ## 0x3
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	7                               ## 0x7
	.byte	6                               ## 0x6
	.byte	8                               ## 0x8
	.byte	7                               ## 0x7
	.byte	10                              ## 0xa
	.byte	9                               ## 0x9
	.byte	11                              ## 0xb
	.byte	10                              ## 0xa
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	6                               ## 0x6
	.byte	5                               ## 0x5
	.byte	8                               ## 0x8
	.byte	7                               ## 0x7
	.byte	9                               ## 0x9
	.byte	8                               ## 0x8
	.byte	11                              ## 0xb
	.byte	10                              ## 0xa
	.byte	12                              ## 0xc
	.byte	11                              ## 0xb
	.byte	14                              ## 0xe
	.byte	13                              ## 0xd
	.byte	15                              ## 0xf
	.byte	14                              ## 0xe
	.space	32
LCPI2_3:
	.space	64,51
LCPI2_4:
	.space	64,26
LCPI2_5:
	.space	64,13
LCPI2_6:
	.byte	71                              ## 0x47
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	237                             ## 0xed
	.byte	240                             ## 0xf0
	.byte	65                              ## 0x41
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	71                              ## 0x47
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	237                             ## 0xed
	.byte	240                             ## 0xf0
	.byte	65                              ## 0x41
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	71                              ## 0x47
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	237                             ## 0xed
	.byte	240                             ## 0xf0
	.byte	65                              ## 0x41
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	71                              ## 0x47
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	252                             ## 0xfc
	.byte	237                             ## 0xed
	.byte	240                             ## 0xf0
	.byte	65                              ## 0x41
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.section	__TEXT,__literal16,16byte_literals
	.p2align	4
LCPI2_1:
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	4                               ## 0x4
	.byte	3                               ## 0x3
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	7                               ## 0x7
	.byte	6                               ## 0x6
	.byte	8                               ## 0x8
	.byte	7                               ## 0x7
	.byte	10                              ## 0xa
	.byte	9                               ## 0x9
	.byte	11                              ## 0xb
	.byte	10                              ## 0xa
LCPI2_2:
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	6                               ## 0x6
	.byte	5                               ## 0x5
	.byte	8                               ## 0x8
	.byte	7                               ## 0x7
	.byte	9                               ## 0x9
	.byte	8                               ## 0x8
	.byte	11                              ## 0xb
	.byte	10                              ## 0xa
	.byte	12                              ## 0xc
	.byte	11                              ## 0xb
	.byte	14                              ## 0xe
	.byte	13                              ## 0xd
	.byte	15                              ## 0xf
	.byte	14                              ## 0xe
	.section	__TEXT,__text,regular,pure_instructions
	.globl	_fast_avx512bw_base64_encode
	.p2align	4, 0x90
_fast_avx512bw_base64_encode:           ## @fast_avx512bw_base64_encode
## %bb.0:
	pushq	%rbp
	movq	%rsp, %rbp
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	andq	$-8, %rsp
	movq	%rdx, %r11
	cmpq	$64, %rdx
	jb	LBB2_1
## %bb.2:
	leaq	-64(%r11), %rdx
	movabsq	$-6148914691236517205, %rax     ## imm = 0xAAAAAAAAAAAAAAAB
	mulxq	%rax, %rax, %rax
	btl	$5, %eax
	jb	LBB2_3
## %bb.4:
	vpermq	$148, (%rsi), %ymm0             ## ymm0 = mem[0,1,1,2]
	vpshufb	LCPI2_0(%rip), %ymm0, %ymm0     ## ymm0 = ymm0[1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10,21,20,22,21,24,23,25,24,27,26,28,27,30,29,31,30]
	vmovdqu	32(%rsi), %xmm1
	vpalignr	$8, 16(%rsi), %xmm1, %xmm2      ## xmm2 = mem[8,9,10,11,12,13,14,15],xmm1[0,1,2,3,4,5,6,7]
	vpshufb	LCPI2_1(%rip), %xmm2, %xmm2     ## xmm2 = xmm2[1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10]
	vpshufb	LCPI2_2(%rip), %xmm1, %xmm1     ## xmm1 = xmm1[5,4,6,5,8,7,9,8,11,10,12,11,14,13,15,14]
	vinserti128	$1, %xmm1, %ymm2, %ymm1
	vinserti64x4	$1, %ymm1, %zmm0, %zmm0
	vpsubusb	LCPI2_3(%rip), %zmm0, %zmm1
	vpcmpltb	LCPI2_4(%rip), %zmm0, %k1
	vmovdqu8	LCPI2_5(%rip), %zmm1 {%k1}
	vmovdqa64	LCPI2_6(%rip), %zmm2    ## zmm2 = [71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0]
	vpshufb	%zmm1, %zmm2, %zmm1
	vpaddb	%zmm1, %zmm0, %zmm0
	vmovdqu64	%zmm0, (%rdi)
	addq	$48, %rsi
	leaq	64(%rdi), %r8
	addq	$-48, %r11
	cmpq	$48, %rdx
	jae	LBB2_6
	jmp	LBB2_8
LBB2_1:
	movq	%rdi, %r8
	jmp	LBB2_8
LBB2_3:
	movq	%rdi, %r8
	cmpq	$48, %rdx
	jb	LBB2_8
LBB2_6:
	vmovdqa	LCPI2_2(%rip), %xmm8            ## xmm8 = [5,4,6,5,8,7,9,8,11,10,12,11,14,13,15,14]
	vmovdqa	LCPI2_1(%rip), %xmm1            ## xmm1 = [1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10]
	vmovdqa	LCPI2_0(%rip), %ymm2            ## ymm2 = [1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10,5,4,6,5,8,7,9,8,11,10,12,11,14,13,15,14]
	vmovdqa64	LCPI2_3(%rip), %zmm3    ## zmm3 = [51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51]
	vmovdqa64	LCPI2_4(%rip), %zmm4    ## zmm4 = [26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26]
	vmovdqa64	LCPI2_5(%rip), %zmm5    ## zmm5 = [13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13]
	vmovdqa64	LCPI2_6(%rip), %zmm6    ## zmm6 = [71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0]
	.p2align	4, 0x90
LBB2_7:                                 ## =>This Inner Loop Header: Depth=1
	vmovdqu	32(%rsi), %xmm7
	vpalignr	$8, 16(%rsi), %xmm7, %xmm0      ## xmm0 = mem[8,9,10,11,12,13,14,15],xmm7[0,1,2,3,4,5,6,7]
	vpshufb	%xmm8, %xmm7, %xmm7
	vpshufb	%xmm1, %xmm0, %xmm0
	vinserti128	$1, %xmm7, %ymm0, %ymm0
	vpermq	$148, (%rsi), %ymm7             ## ymm7 = mem[0,1,1,2]
	vpshufb	%ymm2, %ymm7, %ymm7
	vinserti64x4	$1, %ymm0, %zmm7, %zmm0
	vpsubusb	%zmm3, %zmm0, %zmm7
	vpcmpgtb	%zmm0, %zmm4, %k1
	vmovdqu8	%zmm5, %zmm7 {%k1}
	vpshufb	%zmm7, %zmm6, %zmm7
	vpaddb	%zmm7, %zmm0, %zmm0
	vmovdqu64	%zmm0, (%r8)
	vmovdqu	80(%rsi), %xmm0
	vpalignr	$8, 64(%rsi), %xmm0, %xmm7      ## xmm7 = mem[8,9,10,11,12,13,14,15],xmm0[0,1,2,3,4,5,6,7]
	vpshufb	%xmm8, %xmm0, %xmm0
	vpshufb	%xmm1, %xmm7, %xmm7
	vinserti128	$1, %xmm0, %ymm7, %ymm0
	vpermq	$148, 48(%rsi), %ymm7           ## ymm7 = mem[0,1,1,2]
	vpshufb	%ymm2, %ymm7, %ymm7
	vinserti64x4	$1, %ymm0, %zmm7, %zmm0
	vpsubusb	%zmm3, %zmm0, %zmm7
	vpcmpgtb	%zmm0, %zmm4, %k1
	vmovdqu8	%zmm5, %zmm7 {%k1}
	vpshufb	%zmm7, %zmm6, %zmm7
	vpaddb	%zmm7, %zmm0, %zmm0
	vmovdqu64	%zmm0, 64(%r8)
	addq	$96, %rsi
	subq	$-128, %r8
	addq	$-96, %r11
	cmpq	$63, %r11
	ja	LBB2_7
LBB2_8:
	xorl	%r12d, %r12d
	cmpq	$3, %r11
	jb	LBB2_9
## %bb.10:
	movq	%r11, %r9
	movq	%r8, %rax
	addq	$-2, %r9
	je	LBB2_13
## %bb.11:
	xorl	%r12d, %r12d
	leaq	_e0(%rip), %r10
	leaq	_e2(%rip), %r14
	movq	%r8, %rax
	.p2align	4, 0x90
LBB2_12:                                ## =>This Inner Loop Header: Depth=1
	movzbl	(%rsi,%r12), %ebx
	movzbl	1(%rsi,%r12), %r15d
	movzbl	2(%rsi,%r12), %ecx
	movzbl	(%rbx,%r10), %edx
	movb	%dl, (%rax)
	shll	$4, %ebx
	andl	$48, %ebx
	movq	%r15, %rdx
	shrq	$4, %rdx
	orq	%rbx, %rdx
	movzbl	(%r14,%rdx), %edx
	movb	%dl, 1(%rax)
	andl	$15, %r15d
	movzbl	(%rcx,%r14), %edx
	shrq	$6, %rcx
	leaq	(%rcx,%r15,4), %rcx
	movzbl	(%r14,%rcx), %ecx
	movb	%cl, 2(%rax)
	movb	%dl, 3(%rax)
	addq	$4, %rax
	addq	$3, %r12
	cmpq	%r9, %r12
	jb	LBB2_12
LBB2_13:
	subq	%r12, %r11
	je	LBB2_18
LBB2_14:
	cmpq	$1, %r11
	jne	LBB2_16
## %bb.15:
	movzbl	(%rsi,%r12), %ecx
	leaq	_e0(%rip), %rdx
	movb	(%rcx,%rdx), %dl
	movb	%dl, (%rax)
	shll	$4, %ecx
	andl	$48, %ecx
	leaq	_e2(%rip), %rdx
	movb	(%rcx,%rdx), %cl
	movb	%cl, 1(%rax)
	movb	$61, 2(%rax)
	jmp	LBB2_17
LBB2_16:
	movzbl	(%rsi,%r12), %ecx
	movzbl	1(%rsi,%r12), %edx
	leaq	_e0(%rip), %rsi
	movb	(%rcx,%rsi), %bl
	movb	%bl, (%rax)
	shll	$4, %ecx
	andl	$48, %ecx
	movq	%rdx, %rsi
	shrq	$4, %rsi
	orq	%rcx, %rsi
	leaq	_e2(%rip), %rcx
	movb	(%rcx,%rsi), %bl
	movb	%bl, 1(%rax)
	andl	$15, %edx
	movb	(%rcx,%rdx,4), %cl
	movb	%cl, 2(%rax)
LBB2_17:
	movb	$61, 3(%rax)
	addq	$4, %rax
LBB2_18:
	movb	$0, (%rax)
	movq	%rax, %rcx
	subq	%r8, %rcx
	subq	%rdi, %rax
	cmpq	$-1, %rcx
	cmoveq	%rcx, %rax
	leaq	-32(%rbp), %rsp
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	popq	%rbp
	vzeroupper
	retq
LBB2_9:
	movq	%r8, %rax
	subq	%r12, %r11
	jne	LBB2_14
	jmp	LBB2_18
                                        ## -- End function
	.section	__TEXT,__const
	.p2align	6                               ## -- Begin function fast_avx512bw_base64_decode
LCPI3_0:
	.space	64,15
LCPI3_1:
	.byte	168                             ## 0xa8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	240                             ## 0xf0
	.byte	84                              ## 0x54
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	84                              ## 0x54
	.byte	168                             ## 0xa8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	240                             ## 0xf0
	.byte	84                              ## 0x54
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	84                              ## 0x54
	.byte	168                             ## 0xa8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	240                             ## 0xf0
	.byte	84                              ## 0x54
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	84                              ## 0x54
	.byte	168                             ## 0xa8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	248                             ## 0xf8
	.byte	240                             ## 0xf0
	.byte	84                              ## 0x54
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	80                              ## 0x50
	.byte	84                              ## 0x54
LCPI3_2:
	.byte	1                               ## 0x1
	.byte	2                               ## 0x2
	.byte	4                               ## 0x4
	.byte	8                               ## 0x8
	.byte	16                              ## 0x10
	.byte	32                              ## 0x20
	.byte	64                              ## 0x40
	.byte	128                             ## 0x80
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	1                               ## 0x1
	.byte	2                               ## 0x2
	.byte	4                               ## 0x4
	.byte	8                               ## 0x8
	.byte	16                              ## 0x10
	.byte	32                              ## 0x20
	.byte	64                              ## 0x40
	.byte	128                             ## 0x80
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	1                               ## 0x1
	.byte	2                               ## 0x2
	.byte	4                               ## 0x4
	.byte	8                               ## 0x8
	.byte	16                              ## 0x10
	.byte	32                              ## 0x20
	.byte	64                              ## 0x40
	.byte	128                             ## 0x80
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	1                               ## 0x1
	.byte	2                               ## 0x2
	.byte	4                               ## 0x4
	.byte	8                               ## 0x8
	.byte	16                              ## 0x10
	.byte	32                              ## 0x20
	.byte	64                              ## 0x40
	.byte	128                             ## 0x80
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
LCPI3_3:
	.space	64,47
LCPI3_4:
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	19                              ## 0x13
	.byte	4                               ## 0x4
	.byte	191                             ## 0xbf
	.byte	191                             ## 0xbf
	.byte	185                             ## 0xb9
	.byte	185                             ## 0xb9
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	19                              ## 0x13
	.byte	4                               ## 0x4
	.byte	191                             ## 0xbf
	.byte	191                             ## 0xbf
	.byte	185                             ## 0xb9
	.byte	185                             ## 0xb9
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	19                              ## 0x13
	.byte	4                               ## 0x4
	.byte	191                             ## 0xbf
	.byte	191                             ## 0xbf
	.byte	185                             ## 0xb9
	.byte	185                             ## 0xb9
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	19                              ## 0x13
	.byte	4                               ## 0x4
	.byte	191                             ## 0xbf
	.byte	191                             ## 0xbf
	.byte	185                             ## 0xb9
	.byte	185                             ## 0xb9
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
	.byte	0                               ## 0x0
LCPI3_5:
	.space	64,16
LCPI3_6:
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
	.byte	64                              ## 0x40
	.byte	1                               ## 0x1
LCPI3_7:
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
	.short	4096                            ## 0x1000
	.short	1                               ## 0x1
LCPI3_8:
	.space	1
	.space	1
	.space	1
	.space	1
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	10                              ## 0xa
	.byte	9                               ## 0x9
	.byte	8                               ## 0x8
	.byte	14                              ## 0xe
	.byte	13                              ## 0xd
	.byte	12                              ## 0xc
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
LCPI3_11:
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	10                              ## 0xa
	.byte	9                               ## 0x9
	.byte	8                               ## 0x8
	.byte	14                              ## 0xe
	.byte	13                              ## 0xd
	.byte	12                              ## 0xc
	.space	1
	.space	1
	.space	1
	.space	1
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	10                              ## 0xa
	.byte	9                               ## 0x9
	.byte	8                               ## 0x8
	.byte	14                              ## 0xe
	.byte	13                              ## 0xd
	.byte	12                              ## 0xc
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.section	__TEXT,__literal16,16byte_literals
	.p2align	4
LCPI3_9:
	.byte	8                               ## 0x8
	.byte	14                              ## 0xe
	.byte	13                              ## 0xd
	.byte	12                              ## 0xc
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
LCPI3_10:
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
LCPI3_12:
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.space	1
	.byte	2                               ## 0x2
	.byte	1                               ## 0x1
	.byte	0                               ## 0x0
	.byte	6                               ## 0x6
	.byte	5                               ## 0x5
	.byte	4                               ## 0x4
	.byte	10                              ## 0xa
	.byte	9                               ## 0x9
	.section	__TEXT,__text,regular,pure_instructions
	.globl	_fast_avx512bw_base64_decode
	.p2align	4, 0x90
_fast_avx512bw_base64_decode:           ## @fast_avx512bw_base64_decode
## %bb.0:
	pushq	%rbp
	movq	%rsp, %rbp
	pushq	%r15
	pushq	%r14
	pushq	%r12
	pushq	%rbx
	andq	$-8, %rsp
	movq	%rdi, %r12
	cmpq	$88, %rdx
	jb	LBB3_1
## %bb.2:
	vmovdqa64	LCPI3_0(%rip), %zmm0    ## zmm0 = [15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15]
	vmovdqa64	LCPI3_1(%rip), %zmm9    ## zmm9 = [168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84,168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84,168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84,168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84]
	vmovdqa64	LCPI3_2(%rip), %zmm16   ## zmm16 = [1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0,1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0,1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0,1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0]
	vmovdqa64	LCPI3_3(%rip), %zmm8    ## zmm8 = [47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47]
	vmovdqa64	LCPI3_4(%rip), %zmm11   ## zmm11 = [0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0,0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0,0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0,0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0]
	vmovdqa64	LCPI3_6(%rip), %zmm17   ## zmm17 = [64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1]
	vmovdqa64	LCPI3_7(%rip), %zmm6    ## zmm6 = [4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1]
	vmovdqa	LCPI3_8(%rip), %ymm7            ## ymm7 = <u,u,u,u,2,1,0,6,5,4,10,9,8,14,13,12,2,1,0,6,2,1,0,6,2,1,0,6,2,1,0,6>
	vmovdqa	LCPI3_9(%rip), %xmm3            ## xmm3 = <8,14,13,12,u,u,u,u,u,u,u,u,u,u,u,u>
	vmovdqa	LCPI3_10(%rip), %xmm1           ## xmm1 = <u,u,u,u,u,u,u,u,u,u,u,u,2,1,0,6>
	vmovdqa	LCPI3_11(%rip), %ymm10          ## ymm10 = <2,1,0,6,5,4,10,9,8,14,13,12,u,u,u,u,5,4,10,9,8,14,13,12,u,u,u,u,u,u,u,u>
	vmovdqa	LCPI3_12(%rip), %xmm4           ## xmm4 = <u,u,u,u,u,u,u,u,2,1,0,6,5,4,10,9>
	movq	%r12, %rax
	.p2align	4, 0x90
LBB3_3:                                 ## =>This Inner Loop Header: Depth=1
	vmovdqu64	(%rsi), %zmm12
	vpsrld	$4, %zmm12, %zmm13
	vpandq	%zmm0, %zmm13, %zmm13
	vpandq	%zmm0, %zmm12, %zmm14
	vpshufb	%zmm14, %zmm9, %zmm14
	vpshufb	%zmm13, %zmm16, %zmm15
	vptestmb	%zmm14, %zmm15, %k0
	kortestq	%k0, %k0
	jae	LBB3_29
## %bb.4:                               ##   in Loop: Header=BB3_3 Depth=1
	vpcmpeqb	%zmm8, %zmm12, %k1
	vpshufb	%zmm13, %zmm11, %zmm13
	vmovdqu8	LCPI3_5(%rip), %zmm13 {%k1}
	vpaddb	%zmm12, %zmm13, %zmm12
	vpmaddubsw	%zmm17, %zmm12, %zmm12
	vpmaddwd	%zmm6, %zmm12, %zmm12
	vextracti64x4	$1, %zmm12, %ymm13
	vperm2i128	$33, %ymm12, %ymm13, %ymm14 ## ymm14 = ymm13[2,3],ymm12[0,1]
	vpshufb	%ymm7, %ymm14, %ymm14
	vpshufb	%xmm3, %xmm13, %xmm15
	vpblendd	$1, %ymm15, %ymm14, %ymm14      ## ymm14 = ymm15[0],ymm14[1,2,3,4,5,6,7]
	vextracti128	$1, %ymm12, %xmm2
	vpshufb	%xmm1, %xmm2, %xmm2
	vpshufb	%ymm10, %ymm12, %ymm12
	vpblendd	$8, %ymm2, %ymm12, %ymm2        ## ymm2 = ymm12[0,1,2],ymm2[3],ymm12[4,5,6,7]
	vpshufb	%xmm4, %xmm13, %xmm5
	vinserti128	$1, %xmm5, %ymm0, %ymm5
	vpblendd	$192, %ymm5, %ymm2, %ymm2       ## ymm2 = ymm2[0,1,2,3,4,5],ymm5[6,7]
	vinserti64x4	$1, %ymm14, %zmm2, %zmm2
	vmovdqu64	%zmm2, (%rax)
	addq	$-64, %rdx
	addq	$64, %rsi
	addq	$48, %rax
	cmpq	$87, %rdx
	ja	LBB3_3
## %bb.5:
	testq	%rdx, %rdx
	je	LBB3_6
LBB3_7:
	cmpq	$4, %rdx
	jb	LBB3_29
## %bb.8:
	movl	%edx, %ecx
	andl	$3, %ecx
	testq	%rcx, %rcx
	jne	LBB3_29
## %bb.9:
	cmpb	$61, -1(%rsi,%rdx)
	jne	LBB3_11
## %bb.10:
	leaq	-2(%rdx), %rcx
	cmpb	$61, -2(%rsi,%rdx)
	leaq	-1(%rdx), %rdx
	cmoveq	%rcx, %rdx
LBB3_11:
	xorl	%ecx, %ecx
	movl	%edx, %r9d
	andl	$3, %r9d
	sete	%cl
	shrq	$2, %rdx
	movq	%rdx, %r8
	subq	%rcx, %r8
	je	LBB3_12
## %bb.13:
	negq	%rcx
	addq	%rcx, %rdx
	leaq	_d0(%rip), %r10
	leaq	_d1(%rip), %r11
	leaq	_d2(%rip), %r14
	leaq	_d3(%rip), %r15
	movq	%rax, %rdi
	.p2align	4, 0x90
LBB3_14:                                ## =>This Inner Loop Header: Depth=1
	movzbl	(%rsi), %ebx
	movzbl	1(%rsi), %ecx
	movl	(%r11,%rcx,4), %ecx
	orl	(%r10,%rbx,4), %ecx
	movzbl	2(%rsi), %ebx
	orl	(%r14,%rbx,4), %ecx
	movzbl	3(%rsi), %ebx
	orl	(%r15,%rbx,4), %ecx
	cmpl	$33554430, %ecx                 ## imm = 0x1FFFFFE
	ja	LBB3_29
## %bb.15:                              ##   in Loop: Header=BB3_14 Depth=1
	movl	%ecx, %ebx
	shrl	$16, %ebx
	movb	%cl, (%rdi)
	movb	%ch, 1(%rdi)
	movb	%bl, 2(%rdi)
	addq	$3, %rdi
	addq	$4, %rsi
	decq	%rdx
	jne	LBB3_14
## %bb.16:
	cmpl	$2, %r9d
	je	LBB3_23
LBB3_17:
	cmpl	$1, %r9d
	je	LBB3_21
## %bb.18:
	testl	%r9d, %r9d
	jne	LBB3_24
## %bb.19:
	movzbl	(%rsi), %edx
	leaq	_d0(%rip), %r9
	movzbl	1(%rsi), %ecx
	leaq	_d1(%rip), %rbx
	movl	(%rbx,%rcx,4), %ecx
	orl	(%r9,%rdx,4), %ecx
	movzbl	2(%rsi), %edx
	leaq	_d2(%rip), %rbx
	orl	(%rbx,%rdx,4), %ecx
	movzbl	3(%rsi), %edx
	leaq	_d3(%rip), %rsi
	orl	(%rsi,%rdx,4), %ecx
	cmpl	$33554430, %ecx                 ## imm = 0x1FFFFFE
	ja	LBB3_29
## %bb.20:
	movl	%ecx, %edx
	shrl	$16, %edx
	movb	%cl, (%rdi)
	movb	%ch, 1(%rdi)
	movb	%dl, 2(%rdi)
	leaq	(%r8,%r8,2), %rcx
	addq	$3, %rcx
	jmp	LBB3_27
LBB3_1:
	movq	%r12, %rax
	testq	%rdx, %rdx
	jne	LBB3_7
LBB3_6:
	subq	%r12, %rax
	jmp	LBB3_30
LBB3_12:
	movq	%rax, %rdi
	cmpl	$2, %r9d
	jne	LBB3_17
LBB3_23:
	movzbl	(%rsi), %edx
	leaq	_d0(%rip), %rbx
	movzbl	1(%rsi), %ecx
	leaq	_d1(%rip), %rsi
	movl	(%rsi,%rcx,4), %ecx
	orl	(%rbx,%rdx,4), %ecx
	jmp	LBB3_22
LBB3_21:
	movzbl	(%rsi), %ecx
	leaq	_d0(%rip), %rdx
	movl	(%rdx,%rcx,4), %ecx
LBB3_22:
	movl	%ecx, %edx
	shrl	$8, %edx
	movb	%cl, (%rdi)
	jmp	LBB3_25
LBB3_24:
	movzbl	(%rsi), %edx
	leaq	_d0(%rip), %r10
	movzbl	1(%rsi), %ecx
	leaq	_d1(%rip), %rbx
	movl	(%rbx,%rcx,4), %ecx
	orl	(%r10,%rdx,4), %ecx
	movzbl	2(%rsi), %edx
	leaq	_d2(%rip), %rsi
	orl	(%rsi,%rdx,4), %ecx
	movl	%ecx, %edx
	shrl	$8, %edx
	movw	%cx, (%rdi)
LBB3_25:
	shll	$8, %edx
	movzwl	%dx, %edx
	andl	$-65281, %ecx                   ## imm = 0xFFFF00FF
	orl	%edx, %ecx
	cmpl	$33554430, %ecx                 ## imm = 0x1FFFFFE
	ja	LBB3_29
## %bb.26:
	leaq	(%r8,%r8,2), %rdx
	addl	%r9d, %r9d
	leal	(%r9,%r9,2), %ecx
	shrl	$3, %ecx
	addq	%rdx, %rcx
LBB3_27:
	cmpq	$-1, %rcx
	je	LBB3_29
## %bb.28:
	subq	%r12, %rax
	addq	%rcx, %rax
	jmp	LBB3_30
LBB3_29:
	movq	$-1, %rax
LBB3_30:
	leaq	-32(%rbp), %rsp
	popq	%rbx
	popq	%r12
	popq	%r14
	popq	%r15
	popq	%rbp
	vzeroupper
	retq
                                        ## -- End function
	.section	__TEXT,__const
	.p2align	4                               ## @e0
_e0:
	.ascii	"AAAABBBBCCCCDDDDEEEEFFFFGGGGHHHHIIIIJJJJKKKKLLLLMMMMNNNNOOOOPPPPQQQQRRRRSSSSTTTTUUUUVVVVWWWWXXXXYYYYZZZZaaaabbbbccccddddeeeeffffgggghhhhiiiijjjjkkkkllllmmmmnnnnooooppppqqqqrrrrssssttttuuuuvvvvwwwwxxxxyyyyzzzz0000111122223333444455556666777788889999++++////"

	.p2align	4                               ## @e2
_e2:
	.ascii	"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	.p2align	4                               ## @d0
_d0:
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	248                             ## 0xf8
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	252                             ## 0xfc
	.long	208                             ## 0xd0
	.long	212                             ## 0xd4
	.long	216                             ## 0xd8
	.long	220                             ## 0xdc
	.long	224                             ## 0xe0
	.long	228                             ## 0xe4
	.long	232                             ## 0xe8
	.long	236                             ## 0xec
	.long	240                             ## 0xf0
	.long	244                             ## 0xf4
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	0                               ## 0x0
	.long	4                               ## 0x4
	.long	8                               ## 0x8
	.long	12                              ## 0xc
	.long	16                              ## 0x10
	.long	20                              ## 0x14
	.long	24                              ## 0x18
	.long	28                              ## 0x1c
	.long	32                              ## 0x20
	.long	36                              ## 0x24
	.long	40                              ## 0x28
	.long	44                              ## 0x2c
	.long	48                              ## 0x30
	.long	52                              ## 0x34
	.long	56                              ## 0x38
	.long	60                              ## 0x3c
	.long	64                              ## 0x40
	.long	68                              ## 0x44
	.long	72                              ## 0x48
	.long	76                              ## 0x4c
	.long	80                              ## 0x50
	.long	84                              ## 0x54
	.long	88                              ## 0x58
	.long	92                              ## 0x5c
	.long	96                              ## 0x60
	.long	100                             ## 0x64
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	104                             ## 0x68
	.long	108                             ## 0x6c
	.long	112                             ## 0x70
	.long	116                             ## 0x74
	.long	120                             ## 0x78
	.long	124                             ## 0x7c
	.long	128                             ## 0x80
	.long	132                             ## 0x84
	.long	136                             ## 0x88
	.long	140                             ## 0x8c
	.long	144                             ## 0x90
	.long	148                             ## 0x94
	.long	152                             ## 0x98
	.long	156                             ## 0x9c
	.long	160                             ## 0xa0
	.long	164                             ## 0xa4
	.long	168                             ## 0xa8
	.long	172                             ## 0xac
	.long	176                             ## 0xb0
	.long	180                             ## 0xb4
	.long	184                             ## 0xb8
	.long	188                             ## 0xbc
	.long	192                             ## 0xc0
	.long	196                             ## 0xc4
	.long	200                             ## 0xc8
	.long	204                             ## 0xcc
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff

	.p2align	4                               ## @d1
_d1:
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	57347                           ## 0xe003
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	61443                           ## 0xf003
	.long	16387                           ## 0x4003
	.long	20483                           ## 0x5003
	.long	24579                           ## 0x6003
	.long	28675                           ## 0x7003
	.long	32771                           ## 0x8003
	.long	36867                           ## 0x9003
	.long	40963                           ## 0xa003
	.long	45059                           ## 0xb003
	.long	49155                           ## 0xc003
	.long	53251                           ## 0xd003
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	0                               ## 0x0
	.long	4096                            ## 0x1000
	.long	8192                            ## 0x2000
	.long	12288                           ## 0x3000
	.long	16384                           ## 0x4000
	.long	20480                           ## 0x5000
	.long	24576                           ## 0x6000
	.long	28672                           ## 0x7000
	.long	32768                           ## 0x8000
	.long	36864                           ## 0x9000
	.long	40960                           ## 0xa000
	.long	45056                           ## 0xb000
	.long	49152                           ## 0xc000
	.long	53248                           ## 0xd000
	.long	57344                           ## 0xe000
	.long	61440                           ## 0xf000
	.long	1                               ## 0x1
	.long	4097                            ## 0x1001
	.long	8193                            ## 0x2001
	.long	12289                           ## 0x3001
	.long	16385                           ## 0x4001
	.long	20481                           ## 0x5001
	.long	24577                           ## 0x6001
	.long	28673                           ## 0x7001
	.long	32769                           ## 0x8001
	.long	36865                           ## 0x9001
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	40961                           ## 0xa001
	.long	45057                           ## 0xb001
	.long	49153                           ## 0xc001
	.long	53249                           ## 0xd001
	.long	57345                           ## 0xe001
	.long	61441                           ## 0xf001
	.long	2                               ## 0x2
	.long	4098                            ## 0x1002
	.long	8194                            ## 0x2002
	.long	12290                           ## 0x3002
	.long	16386                           ## 0x4002
	.long	20482                           ## 0x5002
	.long	24578                           ## 0x6002
	.long	28674                           ## 0x7002
	.long	32770                           ## 0x8002
	.long	36866                           ## 0x9002
	.long	40962                           ## 0xa002
	.long	45058                           ## 0xb002
	.long	49154                           ## 0xc002
	.long	53250                           ## 0xd002
	.long	57346                           ## 0xe002
	.long	61442                           ## 0xf002
	.long	3                               ## 0x3
	.long	4099                            ## 0x1003
	.long	8195                            ## 0x2003
	.long	12291                           ## 0x3003
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff

	.p2align	4                               ## @d2
_d2:
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	8392448                         ## 0x800f00
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	12586752                        ## 0xc00f00
	.long	3328                            ## 0xd00
	.long	4197632                         ## 0x400d00
	.long	8391936                         ## 0x800d00
	.long	12586240                        ## 0xc00d00
	.long	3584                            ## 0xe00
	.long	4197888                         ## 0x400e00
	.long	8392192                         ## 0x800e00
	.long	12586496                        ## 0xc00e00
	.long	3840                            ## 0xf00
	.long	4198144                         ## 0x400f00
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	0                               ## 0x0
	.long	4194304                         ## 0x400000
	.long	8388608                         ## 0x800000
	.long	12582912                        ## 0xc00000
	.long	256                             ## 0x100
	.long	4194560                         ## 0x400100
	.long	8388864                         ## 0x800100
	.long	12583168                        ## 0xc00100
	.long	512                             ## 0x200
	.long	4194816                         ## 0x400200
	.long	8389120                         ## 0x800200
	.long	12583424                        ## 0xc00200
	.long	768                             ## 0x300
	.long	4195072                         ## 0x400300
	.long	8389376                         ## 0x800300
	.long	12583680                        ## 0xc00300
	.long	1024                            ## 0x400
	.long	4195328                         ## 0x400400
	.long	8389632                         ## 0x800400
	.long	12583936                        ## 0xc00400
	.long	1280                            ## 0x500
	.long	4195584                         ## 0x400500
	.long	8389888                         ## 0x800500
	.long	12584192                        ## 0xc00500
	.long	1536                            ## 0x600
	.long	4195840                         ## 0x400600
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	8390144                         ## 0x800600
	.long	12584448                        ## 0xc00600
	.long	1792                            ## 0x700
	.long	4196096                         ## 0x400700
	.long	8390400                         ## 0x800700
	.long	12584704                        ## 0xc00700
	.long	2048                            ## 0x800
	.long	4196352                         ## 0x400800
	.long	8390656                         ## 0x800800
	.long	12584960                        ## 0xc00800
	.long	2304                            ## 0x900
	.long	4196608                         ## 0x400900
	.long	8390912                         ## 0x800900
	.long	12585216                        ## 0xc00900
	.long	2560                            ## 0xa00
	.long	4196864                         ## 0x400a00
	.long	8391168                         ## 0x800a00
	.long	12585472                        ## 0xc00a00
	.long	2816                            ## 0xb00
	.long	4197120                         ## 0x400b00
	.long	8391424                         ## 0x800b00
	.long	12585728                        ## 0xc00b00
	.long	3072                            ## 0xc00
	.long	4197376                         ## 0x400c00
	.long	8391680                         ## 0x800c00
	.long	12585984                        ## 0xc00c00
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff

	.p2align	4                               ## @d3
_d3:
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	4063232                         ## 0x3e0000
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	4128768                         ## 0x3f0000
	.long	3407872                         ## 0x340000
	.long	3473408                         ## 0x350000
	.long	3538944                         ## 0x360000
	.long	3604480                         ## 0x370000
	.long	3670016                         ## 0x380000
	.long	3735552                         ## 0x390000
	.long	3801088                         ## 0x3a0000
	.long	3866624                         ## 0x3b0000
	.long	3932160                         ## 0x3c0000
	.long	3997696                         ## 0x3d0000
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	0                               ## 0x0
	.long	65536                           ## 0x10000
	.long	131072                          ## 0x20000
	.long	196608                          ## 0x30000
	.long	262144                          ## 0x40000
	.long	327680                          ## 0x50000
	.long	393216                          ## 0x60000
	.long	458752                          ## 0x70000
	.long	524288                          ## 0x80000
	.long	589824                          ## 0x90000
	.long	655360                          ## 0xa0000
	.long	720896                          ## 0xb0000
	.long	786432                          ## 0xc0000
	.long	851968                          ## 0xd0000
	.long	917504                          ## 0xe0000
	.long	983040                          ## 0xf0000
	.long	1048576                         ## 0x100000
	.long	1114112                         ## 0x110000
	.long	1179648                         ## 0x120000
	.long	1245184                         ## 0x130000
	.long	1310720                         ## 0x140000
	.long	1376256                         ## 0x150000
	.long	1441792                         ## 0x160000
	.long	1507328                         ## 0x170000
	.long	1572864                         ## 0x180000
	.long	1638400                         ## 0x190000
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	1703936                         ## 0x1a0000
	.long	1769472                         ## 0x1b0000
	.long	1835008                         ## 0x1c0000
	.long	1900544                         ## 0x1d0000
	.long	1966080                         ## 0x1e0000
	.long	2031616                         ## 0x1f0000
	.long	2097152                         ## 0x200000
	.long	2162688                         ## 0x210000
	.long	2228224                         ## 0x220000
	.long	2293760                         ## 0x230000
	.long	2359296                         ## 0x240000
	.long	2424832                         ## 0x250000
	.long	2490368                         ## 0x260000
	.long	2555904                         ## 0x270000
	.long	2621440                         ## 0x280000
	.long	2686976                         ## 0x290000
	.long	2752512                         ## 0x2a0000
	.long	2818048                         ## 0x2b0000
	.long	2883584                         ## 0x2c0000
	.long	2949120                         ## 0x2d0000
	.long	3014656                         ## 0x2e0000
	.long	3080192                         ## 0x2f0000
	.long	3145728                         ## 0x300000
	.long	3211264                         ## 0x310000
	.long	3276800                         ## 0x320000
	.long	3342336                         ## 0x330000
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff
	.long	33554431                        ## 0x1ffffff

.subsections_via_symbols
