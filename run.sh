rm jit/*.cpp jit/*.so 2> /dev/null
./gen/gen $1 ./out.scbc && time ./vm/vm ./out.scbc
