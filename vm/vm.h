#ifndef UNTITLED_VM_H
#define UNTITLED_VM_H

#include <string.h>
#include <vector>
#include <string>
#include <iterator>
#include <fstream>
#include <iostream>
#include <cassert>

#define opExtra 0
#define opRead 1
#define opPrint 2
#define opCopy 3
#define opAdd 4
#define opSub 5
#define opMul 6
#define opDiv 7
#define opMod 8
#define opEq 9
#define opNeq 10
#define opGr 11
#define opLe 12
#define opGreq 13
#define opLeeq 14
#define opJump 15
#define opIf 16
#define opCall 17
#define opReturn 18
#define opAssign 19
#define opGet 20
#define opSet 21

#define getIdx(x) ((x) + st)
#define getOpCode(x) ((x) & 63u)
#define getReg1(x) (((x) >> 6u) & 511u)
#define getReg2(x) (((x) >> 15u) & 511u)
#define getReg3(x) (((x) >> 24u) & 511u)
#define getOption1(x) ((x) >> 6u)
#define getOption2(x) ((x) >> 15u)
#define getOption3(x) ((x) >> 24u)

class Vm{
public:
    Vm();
    void run(std::string path);
    uint32_t func_num, line_num, regsize;
    uint32_t *byte_codes, *arg_nums, *var_nums, *def_lines, *call_counts, *reg;
    
};


#endif //UNTITLED_VM_H
