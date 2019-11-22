#include "vm.h"

Vm::Vm(){
}

void Vm::run(std::string path){
    std::ifstream file(path, std::ios::in | std::ifstream::binary);
    
    if(file.fail()){
        std::cerr << "failed to open binary file" << std::endl;
        exit(1);
    }
    file.read(reinterpret_cast<char*>(&func_num), sizeof(func_num));
    
    assert(!file.eof());
    
    functions = (FuncData*) malloc(func_num * sizeof(FuncData));
    
    uint32_t inp;
    
    for(int i = 0; i < func_num; ++i){
        FuncData *f = &functions[i];
        file.read(reinterpret_cast<char*>(&inp), sizeof(inp));
        f->line_cnt = inp;
        file.read(reinterpret_cast<char*>(&inp), sizeof(inp));
        f->var_cnt = inp & ((1u << 16u) - 1);
        f->arg_cnt = (inp >> 16u) & ((1u << 16u) - 1);
        f->byte_codes = (uint32_t*) malloc(f->line_cnt * sizeof(uint32_t));
        f->call_cnt = 0;
        for(int j = 0; j < f->line_cnt; ++j)
            file.read(reinterpret_cast<char*>(&f->byte_codes[j]), sizeof(uint32_t));
    }
    
    file.read(reinterpret_cast<char*>(&inp), sizeof(uint32_t));
    assert(file.eof());
    
    st = 0;
    en = functions[0].var_cnt + 1;
    regsize = 1024;
    reg = (uint32_t*) malloc(1024 * sizeof(uint32_t));
    stacksize = 1024;
    call_stack = (uint32_t*) malloc(1024 * sizeof(uint32_t));
    
    stack_idx = 0;
    call_stack[0] = 0;
    call_stack[1] = 0;
    call_stack[2] = 0;
    while(stackPop());
}

bool Vm::stackPop(){
    if(stack_idx < 0)
        return false;
    uint32_t idx, line, retreg;
    idx = call_stack[3 * stack_idx];
    line = call_stack[3 * stack_idx + 1];
    retreg = call_stack[3 * stack_idx + 2];
    call(idx, line, retreg);
    return true;
}

void Vm::call(uint32_t func_idx, uint32_t line, uint32_t retreg){
    auto f = &functions[func_idx];
    uint32_t* byte_codes = f->byte_codes;
    ++f->call_cnt;
    
    while(line < f->line_cnt){
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
        uint32_t src, dst, src1, src2, label, comp, copy_st, def, ret_reg, val;
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
                
                if(regsize < en + functions[def].var_cnt + 1){
                    uint32_t new_size = (en + functions[def].var_cnt + 1) * 2;
                    auto new_reg = (uint32_t*)malloc(new_size * sizeof(uint32_t));
                    memcpy(new_reg, reg, regsize * sizeof(uint32_t));
                    regsize = new_size;
                    free(reg);
                    reg = new_reg;
                }
                
                for(int i = 0; i < functions[def].arg_cnt; ++i)
                    reg[en + i + 1] = reg[st + copy_st + i];
        
                st = en;
                en += functions[def].var_cnt + 1;
        
        
                if(stacksize < stack_idx * 3 + 6){
                    uint32_t new_size = (stack_idx * 3 + 6) * 2;
                    auto new_stack = (uint32_t*)malloc(new_size * sizeof(uint32_t));
                    memcpy(new_stack, call_stack, stacksize * sizeof(uint32_t));
                    stacksize = new_size;
                    free(call_stack);
                    call_stack = new_stack;
                }
                call_stack[stack_idx * 3] = func_idx;
                call_stack[stack_idx * 3 + 1] = line + 1;
                call_stack[stack_idx * 3 + 2] = retreg;
                call_stack[stack_idx * 3 + 3] = def;
                call_stack[stack_idx * 3 + 4] = 0;
                call_stack[stack_idx * 3 + 5] = dst;
                ++stack_idx;
                return ;
            case opReturn:
                ret_reg = call_stack[stack_idx * 3 + 2];
                --stack_idx;
                def = call_stack[stack_idx * 3];
                if(st == 0)
                    return ;
                en = st;
                st -= (functions[def].var_cnt + 1);
                reg[getIdx(ret_reg)] = reg[en];
                return ;
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

void Vm::jit(uint32_t func_idx){
}
