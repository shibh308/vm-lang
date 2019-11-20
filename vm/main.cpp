#include <iostream>
#include "vm.h"

int main(){
    Vm vm;
    vm.run("../gen/out.scbc");
    return 0;
}