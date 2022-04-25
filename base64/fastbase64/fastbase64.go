.LCPI0_0:
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   4                               # 0x4
        .byte   3                               # 0x3
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   7                               # 0x7
        .byte   6                               # 0x6
        .byte   8                               # 0x8
        .byte   7                               # 0x7
        .byte   10                              # 0xa
        .byte   9                               # 0x9
        .byte   11                              # 0xb
        .byte   10                              # 0xa
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   6                               # 0x6
        .byte   5                               # 0x5
        .byte   8                               # 0x8
        .byte   7                               # 0x7
        .byte   9                               # 0x9
        .byte   8                               # 0x8
        .byte   11                              # 0xb
        .byte   10                              # 0xa
        .byte   12                              # 0xc
        .byte   11                              # 0xb
        .byte   14                              # 0xe
        .byte   13                              # 0xd
        .byte   15                              # 0xf
        .byte   14                              # 0xe
.LCPI0_1:
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   4                               # 0x4
        .byte   3                               # 0x3
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   7                               # 0x7
        .byte   6                               # 0x6
        .byte   8                               # 0x8
        .byte   7                               # 0x7
        .byte   10                              # 0xa
        .byte   9                               # 0x9
        .byte   11                              # 0xb
        .byte   10                              # 0xa
.LCPI0_2:
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   6                               # 0x6
        .byte   5                               # 0x5
        .byte   8                               # 0x8
        .byte   7                               # 0x7
        .byte   9                               # 0x9
        .byte   8                               # 0x8
        .byte   11                              # 0xb
        .byte   10                              # 0xa
        .byte   12                              # 0xc
        .byte   11                              # 0xb
        .byte   14                              # 0xe
        .byte   13                              # 0xd
        .byte   15                              # 0xf
        .byte   14                              # 0xe
.LCPI0_3:
        .zero   64,51
.LCPI0_4:
        .zero   64,26
.LCPI0_5:
        .zero   64,13
.LCPI0_6:
        .byte   71                              # 0x47
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   237                             # 0xed
        .byte   240                             # 0xf0
        .byte   65                              # 0x41
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   71                              # 0x47
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   237                             # 0xed
        .byte   240                             # 0xf0
        .byte   65                              # 0x41
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   71                              # 0x47
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   237                             # 0xed
        .byte   240                             # 0xf0
        .byte   65                              # 0x41
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   71                              # 0x47
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   252                             # 0xfc
        .byte   237                             # 0xed
        .byte   240                             # 0xf0
        .byte   65                              # 0x41
        .byte   0                               # 0x0
        .byte   0                               # 0x0
fast_avx512bw_base64_encode:            # @fast_avx512bw_base64_encode
        pushq   %r14
        pushq   %rbx
        pushq   %rax
        movq    %rdx, %rax
        movq    %rdi, %r14
        cmpq    $64, %rdx
        jb      .LBB0_1
        leaq    -64(%rax), %rdx
        movabsq $-6148914691236517205, %rcx     # imm = 0xAAAAAAAAAAAAAAAB
        mulxq   %rcx, %rcx, %rcx
        btl     $5, %ecx
        jb      .LBB0_3
        vpermq  $148, (%rsi), %ymm0             # ymm0 = mem[0,1,1,2]
        vpshufb .LCPI0_0(%rip), %ymm0, %ymm0    # ymm0 = ymm0[1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10,21,20,22,21,24,23,25,24,27,26,28,27,30,29,31,30]
        vmovdqu 32(%rsi), %xmm1
        vpalignr        $8, 16(%rsi), %xmm1, %xmm2      # xmm2 = mem[8,9,10,11,12,13,14,15],xmm1[0,1,2,3,4,5,6,7]
        vpshufb .LCPI0_1(%rip), %xmm2, %xmm2    # xmm2 = xmm2[1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10]
        vpshufb .LCPI0_2(%rip), %xmm1, %xmm1    # xmm1 = xmm1[5,4,6,5,8,7,9,8,11,10,12,11,14,13,15,14]
        vinserti128     $1, %xmm1, %ymm2, %ymm1
        vinserti64x4    $1, %ymm1, %zmm0, %zmm0
        vpsubusb        .LCPI0_3(%rip), %zmm0, %zmm1
        vpcmpltb        .LCPI0_4(%rip), %zmm0, %k1
        vmovdqu8        .LCPI0_5(%rip), %zmm1 {%k1}
        vmovdqa64       .LCPI0_6(%rip), %zmm2   # zmm2 = [71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0]
        vpshufb %zmm1, %zmm2, %zmm1
        vpaddb  %zmm1, %zmm0, %zmm0
        vmovdqu64       %zmm0, (%r14)
        addq    $48, %rsi
        leaq    64(%r14), %rbx
        addq    $-48, %rax
        cmpq    $48, %rdx
        jae     .LBB0_6
        jmp     .LBB0_8
.LBB0_1:
        movq    %r14, %rbx
        jmp     .LBB0_8
.LBB0_3:
        movq    %r14, %rbx
        cmpq    $48, %rdx
        jb      .LBB0_8
.LBB0_6:
        vmovdqa .LCPI0_2(%rip), %xmm8           # xmm8 = [5,4,6,5,8,7,9,8,11,10,12,11,14,13,15,14]
        vmovdqa .LCPI0_1(%rip), %xmm1           # xmm1 = [1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10]
        vmovdqa .LCPI0_0(%rip), %ymm2           # ymm2 = [1,0,2,1,4,3,5,4,7,6,8,7,10,9,11,10,5,4,6,5,8,7,9,8,11,10,12,11,14,13,15,14]
        vmovdqa64       .LCPI0_3(%rip), %zmm3   # zmm3 = [51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51,51]
        vmovdqa64       .LCPI0_4(%rip), %zmm4   # zmm4 = [26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26,26]
        vmovdqa64       .LCPI0_5(%rip), %zmm5   # zmm5 = [13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13,13]
        vmovdqa64       .LCPI0_6(%rip), %zmm6   # zmm6 = [71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0,71,252,252,252,252,252,252,252,252,252,252,237,240,65,0,0]
.LBB0_7:                                # =>This Inner Loop Header: Depth=1
        vmovdqu 32(%rsi), %xmm7
        vpalignr        $8, 16(%rsi), %xmm7, %xmm0      # xmm0 = mem[8,9,10,11,12,13,14,15],xmm7[0,1,2,3,4,5,6,7]
        vpshufb %xmm8, %xmm7, %xmm7
        vpshufb %xmm1, %xmm0, %xmm0
        vinserti128     $1, %xmm7, %ymm0, %ymm0
        vpermq  $148, (%rsi), %ymm7             # ymm7 = mem[0,1,1,2]
        vpshufb %ymm2, %ymm7, %ymm7
        vinserti64x4    $1, %ymm0, %zmm7, %zmm0
        vpsubusb        %zmm3, %zmm0, %zmm7
        vpcmpgtb        %zmm0, %zmm4, %k1
        vmovdqu8        %zmm5, %zmm7 {%k1}
        vpshufb %zmm7, %zmm6, %zmm7
        vpaddb  %zmm7, %zmm0, %zmm0
        vmovdqu64       %zmm0, (%rbx)
        vmovdqu 80(%rsi), %xmm0
        vpalignr        $8, 64(%rsi), %xmm0, %xmm7      # xmm7 = mem[8,9,10,11,12,13,14,15],xmm0[0,1,2,3,4,5,6,7]
        vpshufb %xmm8, %xmm0, %xmm0
        vpshufb %xmm1, %xmm7, %xmm7
        vinserti128     $1, %xmm0, %ymm7, %ymm0
        vpermq  $148, 48(%rsi), %ymm7           # ymm7 = mem[0,1,1,2]
        vpshufb %ymm2, %ymm7, %ymm7
        vinserti64x4    $1, %ymm0, %zmm7, %zmm0
        vpsubusb        %zmm3, %zmm0, %zmm7
        vpcmpgtb        %zmm0, %zmm4, %k1
        vmovdqu8        %zmm5, %zmm7 {%k1}
        vpshufb %zmm7, %zmm6, %zmm7
        vpaddb  %zmm7, %zmm0, %zmm0
        vmovdqu64       %zmm0, 64(%rbx)
        addq    $96, %rsi
        subq    $-128, %rbx
        addq    $-96, %rax
        cmpq    $63, %rax
        ja      .LBB0_7
.LBB0_8:
        movq    %rbx, %rdi
        movq    %rax, %rdx
        vzeroupper
        callq   chromium_base64_encode@PLT
        subq    %r14, %rbx
        addq    %rax, %rbx
        cmpq    $-1, %rax
        cmovneq %rbx, %rax
        addq    $8, %rsp
        popq    %rbx
        popq    %r14
        retq
.LCPI1_0:
        .zero   64,15
.LCPI1_1:
        .byte   168                             # 0xa8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   240                             # 0xf0
        .byte   84                              # 0x54
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   84                              # 0x54
        .byte   168                             # 0xa8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   240                             # 0xf0
        .byte   84                              # 0x54
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   84                              # 0x54
        .byte   168                             # 0xa8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   240                             # 0xf0
        .byte   84                              # 0x54
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   84                              # 0x54
        .byte   168                             # 0xa8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   248                             # 0xf8
        .byte   240                             # 0xf0
        .byte   84                              # 0x54
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   80                              # 0x50
        .byte   84                              # 0x54
.LCPI1_2:
        .byte   1                               # 0x1
        .byte   2                               # 0x2
        .byte   4                               # 0x4
        .byte   8                               # 0x8
        .byte   16                              # 0x10
        .byte   32                              # 0x20
        .byte   64                              # 0x40
        .byte   128                             # 0x80
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   1                               # 0x1
        .byte   2                               # 0x2
        .byte   4                               # 0x4
        .byte   8                               # 0x8
        .byte   16                              # 0x10
        .byte   32                              # 0x20
        .byte   64                              # 0x40
        .byte   128                             # 0x80
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   1                               # 0x1
        .byte   2                               # 0x2
        .byte   4                               # 0x4
        .byte   8                               # 0x8
        .byte   16                              # 0x10
        .byte   32                              # 0x20
        .byte   64                              # 0x40
        .byte   128                             # 0x80
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   1                               # 0x1
        .byte   2                               # 0x2
        .byte   4                               # 0x4
        .byte   8                               # 0x8
        .byte   16                              # 0x10
        .byte   32                              # 0x20
        .byte   64                              # 0x40
        .byte   128                             # 0x80
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
.LCPI1_3:
        .zero   64,47
.LCPI1_4:
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   19                              # 0x13
        .byte   4                               # 0x4
        .byte   191                             # 0xbf
        .byte   191                             # 0xbf
        .byte   185                             # 0xb9
        .byte   185                             # 0xb9
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   19                              # 0x13
        .byte   4                               # 0x4
        .byte   191                             # 0xbf
        .byte   191                             # 0xbf
        .byte   185                             # 0xb9
        .byte   185                             # 0xb9
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   19                              # 0x13
        .byte   4                               # 0x4
        .byte   191                             # 0xbf
        .byte   191                             # 0xbf
        .byte   185                             # 0xb9
        .byte   185                             # 0xb9
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   19                              # 0x13
        .byte   4                               # 0x4
        .byte   191                             # 0xbf
        .byte   191                             # 0xbf
        .byte   185                             # 0xb9
        .byte   185                             # 0xb9
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
        .byte   0                               # 0x0
.LCPI1_5:
        .zero   64,16
.LCPI1_6:
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
        .byte   64                              # 0x40
        .byte   1                               # 0x1
.LCPI1_7:
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
        .short  4096                            # 0x1000
        .short  1                               # 0x1
.LCPI1_8:
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   10                              # 0xa
        .byte   9                               # 0x9
        .byte   8                               # 0x8
        .byte   14                              # 0xe
        .byte   13                              # 0xd
        .byte   12                              # 0xc
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
.LCPI1_11:
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   10                              # 0xa
        .byte   9                               # 0x9
        .byte   8                               # 0x8
        .byte   14                              # 0xe
        .byte   13                              # 0xd
        .byte   12                              # 0xc
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   10                              # 0xa
        .byte   9                               # 0x9
        .byte   8                               # 0x8
        .byte   14                              # 0xe
        .byte   13                              # 0xd
        .byte   12                              # 0xc
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
.LCPI1_9:
        .byte   8                               # 0x8
        .byte   14                              # 0xe
        .byte   13                              # 0xd
        .byte   12                              # 0xc
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
.LCPI1_10:
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
.LCPI1_12:
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .zero   1
        .byte   2                               # 0x2
        .byte   1                               # 0x1
        .byte   0                               # 0x0
        .byte   6                               # 0x6
        .byte   5                               # 0x5
        .byte   4                               # 0x4
        .byte   10                              # 0xa
        .byte   9                               # 0x9
fast_avx512bw_base64_decode:            # @fast_avx512bw_base64_decode
        pushq   %r14
        pushq   %rbx
        pushq   %rax
        movq    %rdi, %r14
        cmpq    $88, %rdx
        jb      .LBB1_1
        vmovdqa64       .LCPI1_0(%rip), %zmm0   # zmm0 = [15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15,15]
        vmovdqa64       .LCPI1_1(%rip), %zmm9   # zmm9 = [168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84,168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84,168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84,168,248,248,248,248,248,248,248,248,248,240,84,80,80,80,84]
        vmovdqa64       .LCPI1_2(%rip), %zmm16  # zmm16 = [1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0,1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0,1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0,1,2,4,8,16,32,64,128,0,0,0,0,0,0,0,0]
        vmovdqa64       .LCPI1_3(%rip), %zmm8   # zmm8 = [47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47,47]
        vmovdqa64       .LCPI1_4(%rip), %zmm11  # zmm11 = [0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0,0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0,0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0,0,0,19,4,191,191,185,185,0,0,0,0,0,0,0,0]
        vmovdqa64       .LCPI1_6(%rip), %zmm17  # zmm17 = [64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1,64,1]
        vmovdqa64       .LCPI1_7(%rip), %zmm6   # zmm6 = [4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1,4096,1]
        vmovdqa .LCPI1_8(%rip), %ymm7           # ymm7 = <u,u,u,u,2,1,0,6,5,4,10,9,8,14,13,12,2,1,0,6,2,1,0,6,2,1,0,6,2,1,0,6>
        vmovdqa .LCPI1_9(%rip), %xmm3           # xmm3 = <8,14,13,12,u,u,u,u,u,u,u,u,u,u,u,u>
        vmovdqa .LCPI1_10(%rip), %xmm1          # xmm1 = <u,u,u,u,u,u,u,u,u,u,u,u,2,1,0,6>
        vmovdqa .LCPI1_11(%rip), %ymm10         # ymm10 = <2,1,0,6,5,4,10,9,8,14,13,12,u,u,u,u,5,4,10,9,8,14,13,12,u,u,u,u,u,u,u,u>
        vmovdqa .LCPI1_12(%rip), %xmm4          # xmm4 = <u,u,u,u,u,u,u,u,2,1,0,6,5,4,10,9>
        movq    %r14, %rbx
.LBB1_3:                                # =>This Inner Loop Header: Depth=1
        vmovdqu64       (%rsi), %zmm12
        vpsrld  $4, %zmm12, %zmm13
        vpandq  %zmm0, %zmm13, %zmm13
        vpandq  %zmm0, %zmm12, %zmm14
        vpshufb %zmm14, %zmm9, %zmm14
        vpshufb %zmm13, %zmm16, %zmm15
        vptestnmb       %zmm14, %zmm15, %k0
        kortestq        %k0, %k0
        jne     .LBB1_4
        vpcmpeqb        %zmm8, %zmm12, %k1
        vpshufb %zmm13, %zmm11, %zmm13
        vmovdqu8        .LCPI1_5(%rip), %zmm13 {%k1}
        vpaddb  %zmm12, %zmm13, %zmm12
        vpmaddubsw      %zmm17, %zmm12, %zmm12
        vpmaddwd        %zmm6, %zmm12, %zmm12
        vextracti64x4   $1, %zmm12, %ymm13
        vperm2i128      $33, %ymm12, %ymm13, %ymm14 # ymm14 = ymm13[2,3],ymm12[0,1]
        vpshufb %ymm7, %ymm14, %ymm14
        vpshufb %xmm3, %xmm13, %xmm15
        vpblendd        $1, %ymm15, %ymm14, %ymm14      # ymm14 = ymm15[0],ymm14[1,2,3,4,5,6,7]
        vextracti128    $1, %ymm12, %xmm2
        vpshufb %xmm1, %xmm2, %xmm2
        vpshufb %ymm10, %ymm12, %ymm12
        vpblendd        $8, %ymm2, %ymm12, %ymm2        # ymm2 = ymm12[0,1,2],ymm2[3],ymm12[4,5,6,7]
        vpshufb %xmm4, %xmm13, %xmm5
        vinserti128     $1, %xmm5, %ymm0, %ymm5
        vpblendd        $192, %ymm5, %ymm2, %ymm2       # ymm2 = ymm2[0,1,2,3,4,5],ymm5[6,7]
        vinserti64x4    $1, %ymm14, %zmm2, %zmm2
        vmovdqu64       %zmm2, (%rbx)
        addq    $-64, %rdx
        addq    $64, %rsi
        addq    $48, %rbx
        cmpq    $87, %rdx
        ja      .LBB1_3
        jmp     .LBB1_6
.LBB1_1:
        movq    %r14, %rbx
.LBB1_6:
        movq    %rbx, %rdi
        vzeroupper
        callq   chromium_base64_decode@PLT
        subq    %r14, %rbx
        addq    %rax, %rbx
        cmpq    $-1, %rax
        cmoveq  %rax, %rbx
        jmp     .LBB1_7
.LBB1_4:
        movq    $-1, %rbx
.LBB1_7:
        movq    %rbx, %rax
        addq    $8, %rsp
        popq    %rbx
        popq    %r14
        vzeroupper
        retq
