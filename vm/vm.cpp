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
        uint32_t src, dst, src1, src2, label, comp, copy_st, def, before_st, ret, ret_reg, val;
        switch(op_code){
            case opExtra:
                /* TODO*/
                break;
            case opRead:
                dst = getReg1(bc);
                scanf("%d", &reg[getIdx(dst)]);
                break;
            case opPrint:
                src = getReg1(bc);
                printf("%d\n", reg[getIdx(src)]);
                break;
            case opCopy:
                src = getReg1(bc);
                dst = getReg2(bc);
                reg[getIdx(dst)] = reg[getIdx(src)];
                break;
            case opAdd:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] + reg[getIdx(src2)];
                break;
            case opSub:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] - reg[getIdx(src2)];
                break;
            case opMul:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] * reg[getIdx(src2)];
                break;
            case opDiv:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] / reg[getIdx(src2)];
            case opMod:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] % reg[getIdx(src2)];
                break;
            case opEq:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] == reg[getIdx(src2)];
                break;
            case opNeq:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[st + dst] = reg[st + src1] != reg[st + src2];
                reg[getIdx(dst)] = reg[getIdx(src1)] != reg[getIdx(src2)];
                break;
            case opGr:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] > reg[getIdx(src2)];
                break;
            case opLe:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] < reg[getIdx(src2)];
                break;
            case opGreq:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] >= reg[getIdx(src2)];
                break;
            case opLeeq:
                src1 = getReg1(bc);
                src2 = getReg2(bc);
                dst = getReg3(bc);
                reg[getIdx(dst)] = reg[getIdx(src1)] <= reg[getIdx(src2)];
                break;
            case opJump:
                label = getOption1(bc);
                line = label - 1;
                break;
            case opIf:
                comp = getReg1(bc);
                label = getOption2(bc);
                if(reg[getIdx(comp)] == 0)
                    line = label;
                break;
            case opCall:
                copy_st = getReg1(bc);
                dst = getReg2(bc);
                def = getOption3(bc);
                ++call_counts[def];
                if(regsize < en + var_nums[def] + 4){
                    uint32_t new_size = (en + var_nums[def] + 4) * 2;
                    auto new_reg = (uint32_t*)malloc(new_size * sizeof(uint32_t));
                    memcpy(new_reg, reg, regsize * sizeof(uint32_t));
                    regsize = new_size;
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
                line = def_lines[def] - 1;
                break;
            case opReturn:
                before_st = reg[st + 3];
                ret = reg[st + 1];
                ret_reg = reg[st + 2];
                line = reg[st];
                en = st;
                st = before_st;
                reg[ret_reg] = ret;
                break;
            case opAssign:
                dst = getReg1(bc);
                val = getOption2(bc);
                reg[getIdx(dst)] = val;
                break;
            case opGet:
                /* TODO*/
                break;
            case opSet:
                /* TODO*/
                break;
                
        }
        ++line;
    }
}
