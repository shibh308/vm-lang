#include "vm.h"

Vm::Vm(){
}

void Vm::run(std::string path){
    std::ifstream file(path, std::ios::in | std::ifstream::binary);
    
    if(file.fail()){
        std::cerr << "failed to open binary file" << std::endl;
        exit(1);
    }
    file.read(reinterpret_cast<char*>(&line_num), sizeof(line_num));
    assert(!file.eof());
    file.read(reinterpret_cast<char*>(&func_num), sizeof(func_num));
    assert(!file.eof());
    var_nums = (uint32_t*)malloc(func_num * sizeof(uint32_t));
    arg_nums = (uint32_t*)malloc(func_num * sizeof(uint32_t));
    def_lines = (uint32_t*)malloc(func_num * sizeof(uint32_t));
    call_counts = (uint32_t*)calloc(func_num, sizeof(uint32_t));
    call_counts[0] = 1;
    uint32_t inp;
    for(int i = 0; i < func_num; ++i){
        file.read(reinterpret_cast<char*>(&inp), sizeof(inp));
        def_lines[i] = inp;
        file.read(reinterpret_cast<char*>(&inp), sizeof(inp));
        var_nums[i] = inp & ((1u << 16u) - 1);
        arg_nums[i] = (inp >> 16u) & ((1u << 16u) - 1);
    }
    
    byte_codes = (uint32_t*)malloc(line_num * sizeof(uint32_t));
    for(int i = 0; i < line_num; ++i)
        file.read(reinterpret_cast<char*>(&byte_codes[i]), sizeof(uint32_t));
    file.read(reinterpret_cast<char*>(&inp), sizeof(uint32_t));
    assert(file.eof());
    
    uint32_t line = 0;
    uint32_t st = 0;
    uint32_t en = var_nums[0] + 4;
    regsize = 1024;
    reg = (uint32_t*)malloc(1024 * sizeof(uint32_t));
    reg[0] = line_num;
    reg[2] = 1;
    reg[3] = 0;
    /* TODO: argument */
    
    while(line < line_num){
        uint32_t bc = byte_codes[line];
    
        /*
        std::cout << "reg: ";
        for(auto& x : reg)
            std::cout << x << " ";
        std::cout << std::endl;
        std::cout << std::endl;
        std::cout << line << ": ";
        std::cout << getOpCode(bc) << " " << getReg1(bc) << " " << getReg2(bc) << " " << getReg3(bc) << std::endl;
         */
        
        uint32_t op_code = getOpCode(bc);
        if(op_code == opExtra){
            /* TODO*/
        }
        else if(op_code == opRead){
            uint32_t dst = getReg1(bc);
            scanf("%d", &reg[getIdx(dst)]);
        }
        else if(op_code == opPrint){
            uint32_t src = getReg1(bc);
            printf("%d\n", reg[getIdx(src)]);
        }
        else if(op_code == opCopy){
            uint32_t src = getReg1(bc);
            uint32_t dst = getReg2(bc);
            reg[getIdx(dst)] = reg[getIdx(src)];
        }
        else if(op_code == opAdd){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] + reg[getIdx(src2)];
        }
        else if(op_code == opSub){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] - reg[getIdx(src2)];
        }
        else if(op_code == opMul){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] * reg[getIdx(src2)];
        }
        else if(op_code == opDiv){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] / reg[getIdx(src2)];
        }
        else if(op_code == opMod){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[dst] = reg[src1] % reg[src2];
            reg[getIdx(dst)] = reg[getIdx(src1)] % reg[getIdx(src2)];
        }
        else if(op_code == opEq){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] == reg[getIdx(src2)];
        }
        else if(op_code == opNeq){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[st + dst] = reg[st + src1] != reg[st + src2];
            reg[getIdx(dst)] = reg[getIdx(src1)] != reg[getIdx(src2)];
        }
        else if(op_code == opGr){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] > reg[getIdx(src2)];
        }
        else if(op_code == opLe){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] < reg[getIdx(src2)];
        }
        else if(op_code == opGreq){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] >= reg[getIdx(src2)];
        }
        else if(op_code == opLeeq){
            uint32_t src1 = getReg1(bc);
            uint32_t src2 = getReg2(bc);
            uint32_t dst = getReg3(bc);
            reg[getIdx(dst)] = reg[getIdx(src1)] <= reg[getIdx(src2)];
        }
        else if(op_code == opJump){
            uint32_t label = getOption1(bc);
            line = label;
            continue;
        }
        else if(op_code == opIf){
            uint32_t comp = getReg1(bc);
            uint32_t label = getOption2(bc);
            if(reg[getIdx(comp)] == 0)
                line = label;
        }
        else if(op_code == opCall){
            uint32_t copy_st = getReg1(bc);
            uint32_t dst = getReg2(bc);
            uint32_t def = getOption3(bc);
            ++call_counts[def];
            if(regsize < en + var_nums[def] + 4){
                auto new_reg = (uint32_t*)malloc((en + var_nums[def] + 4) * 2);
                memcpy(new_reg, reg, sizeof(uint32_t));
                free(reg);
                reg = new_reg;
            }
            reg[en] = line;
            reg[en + 2] = getIdx(dst);
            reg[en + 3] = st;
            for(int i = 0; i < arg_nums[def]; ++i)
                reg[en + i + 4] = reg[st + copy_st + i];
            st = en;
            en += var_nums[def] + 4;
            line = def_lines[def];
            continue;
        }
        else if(op_code == opReturn){
            uint32_t before_st = reg[st + 3];
            uint32_t ret = reg[st + 1];
            uint32_t ret_reg = reg[st + 2];
            line = reg[st];
            en = st;
            st = before_st;
            reg[ret_reg] = ret;
        }
        else if(op_code == opAssign){
            uint32_t dst = getReg1(bc);
            uint32_t val = getOption2(bc);
            reg[getIdx(dst)] = val;
        }
        else if(op_code == opGet){
            /* TODO*/
        }
        else if(op_code == opSet){
            /* TODO*/
        }
        ++line;
    }
}
