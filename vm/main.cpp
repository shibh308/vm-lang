#include <iostream>
#include "vm.h"

#define TEST

int main(int argc, char** argv){
#ifdef TEST
    std::vector<std::string> files{
        /*
        "../gen/out_s1.scbc",
        "../gen/out_s2.scbc",
        "../gen/out_s3.scbc",
        "../gen/out_s4.scbc",
         */
        "../out.scbc",
    };
    for(auto& s : files){
        Vm vm;
        vm.run(s);
    }
#else
    Vm vm;
    vm.run(argv[1]);
#endif
    return 0;
}