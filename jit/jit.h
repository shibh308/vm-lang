#include <string.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>
/*
#include <thread>
#include <queue>
*/

#define GETVAL(x) vm->reg[(x) + vm->st]


struct FuncData{
    uint32_t line_cnt, var_cnt, arg_cnt, call_cnt;
    uint32_t* byte_codes;
    bool make_jit;
    void *func_ptr;
};


struct Vm{
    FuncData *functions;
    uint32_t *call_stack, *reg;
    uint32_t st, en, func_num, regsize, stacksize;
    int stack_idx;

    // Vm();
    // ~Vm();
    // void run(std::string path);
    //void call(uint32_t func_idx, uint32_t line, uint32_t retreg);
    // void jitCheck();
    // void jit(uint32_t func_idx);
    // bool stackPop();

	/*
    bool jit_running = false;
    std::queue<int> jit_queue;
    std::thread jit_thread;
	*/
};

