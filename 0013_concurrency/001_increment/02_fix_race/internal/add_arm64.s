#include "go_asm.h"
#include "textflag.h"

// func SyncAdd
// func SyncAdd(addr *int64, delta int64)
TEXT ·SyncAdd(SB), NOSPLIT, $0-16
    MOVD ptr+0(FP), R0
    MOVD delta+8(FP), R1
    LDADDALD R1, (R0), R2
    ADD R1, R2
    MOVD R2, ret+16(FP)
    RET
