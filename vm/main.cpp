#include <iostream>
#include "vm.h"

int main(){
    std::vector<std::string> files{
        "../gen/out_s1.scbc",
        "../gen/out_s2.scbc",
        "../gen/out_s3.scbc",
        "../gen/out_s4.scbc",
    };
    for(auto& s : files){
        Vm vm;
        vm.run(s);
    }
    return 0;
}