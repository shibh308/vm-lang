#ifndef UNTITLED_VM_H
#define UNTITLED_VM_H

#include <cstring>
#include <dlfcn.h>
#include <vector>
#include <string>
#include <iterator>
#include <fstream>
#include <iostream>
#include <cassert>
#include <thread>
#include <queue>

#define TMPDIR std::string("./jit/")

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

#define OUT(x) file << x << std::endl

#define JIT

const std::string func_prefix = "func_";
const std::string label_prefix = "line_";


class Vm;

struct FuncData{
    uint32_t line_cnt, var_cnt, arg_cnt, call_cnt;
    uint32_t* byte_codes;
    bool make_jit;
    void (*func)(Vm*, uint32_t);
};


class Vm{
public:
    FuncData *functions;
    uint32_t *call_stack, *reg;
    uint32_t st, en, func_num, regsize, stacksize;
    int stack_idx;
    Vm();
    ~Vm();
    void run(std::string path);
    void call(uint32_t func_idx, uint32_t line);
    void jitCheck();
    void jit(uint32_t func_idx);
    bool stackPop();

    bool jit_running = false;
    std::queue<int> jit_queue;
    std::thread jit_thread;
};


#endif //UNTITLED_VM_H
