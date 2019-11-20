#ifndef UNTITLED_VM_H
#define UNTITLED_VM_H

#include <vector>
#include <string>
#include <iterator>
#include <fstream>
#include <iostream>

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
#define getOpCode(x) ((x) & 63)
#define getReg1(x) (((x) >> 6) & 511)
#define getReg2(x) (((x) >> 15) & 511)
#define getReg3(x) (((x) >> 24) & 511)
#define getOption1(x) ((x) >> 6)
#define getOption2(x) ((x) >> 15)
#define getOption3(x) ((x) >> 24)

class Vm{
public:
    Vm();
    void run(std::string path);
    uint32_t func_num;
    std::vector<uint32_t> byte_codes;
    std::vector<uint32_t> arg_nums;
    std::vector<uint32_t> var_nums;
    std::vector<uint32_t> def_lines;
    std::vector<uint32_t> call_counts;
    
};


#endif //UNTITLED_VM_H
