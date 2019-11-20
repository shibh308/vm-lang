#include "vm.h"

Vm::Vm(){
}
void Vm::run(std::string path){
    std::ifstream file(path, std::ios::in | std::ifstream::binary);
    
    if(file.fail()){
        std::cerr << "failed to open binary file" << std::endl;
        exit(1);
    }
    while(true){
        uint32_t inp;
        file.read(reinterpret_cast<char*>(&inp), sizeof(inp));
        if(file.eof())
            break;
        byte_codes.emplace_back(inp);
    }
    
    func_num = byte_codes[0];
    arg_nums.resize(func_num);
    var_nums.resize(func_num);
    def_lines.resize(func_num);
    call_counts.resize(func_num, 0);
    call_counts[0] = 1;
    uint32_t line = 1;
    for(int func_idx = 0; func_idx < func_num; ++func_idx){
        var_nums[func_idx] = byte_codes[line] & ((1 << 16) - 1);
        arg_nums[func_idx] = (byte_codes[line++] >> 16) & ((1 << 16) - 1);
        def_lines[func_idx] = byte_codes[line++] + func_num * 2 + 1;
    }
    
    uint32_t st = 0;
    uint32_t en = var_nums[0] + 4;
    std::vector<int> reg(var_nums[0] + 4);
    reg[0] = byte_codes.size();
    reg[2] = 1;
    reg[3] = 0;
    /* TODO: argument */
    
    while(line < byte_codes.size()){
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
            /* TODO*/
        }
        else if(op_code == opPrint){
            /* TODO*/
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
            if(comp)
                line = label;
            continue;
        }
        else if(op_code == opCall){
            uint32_t copy_st = getReg1(bc);
            uint32_t dst = getReg2(bc);
            uint32_t def = getOption3(bc);
            ++call_counts[def];
            if(reg.size() < en + var_nums[def] + 4)
                reg.resize(en + var_nums[def] + 4);
            reg[en] = line;
            reg[en + 2] = dst;
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
            reg[st + ret_reg] = ret;
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
    std::cout << reg[1] << std::endl;
}
