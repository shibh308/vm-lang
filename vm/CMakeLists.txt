cmake_minimum_required(VERSION 3.14)
project(vm)

add_executable(vm main.cpp vm.cpp)

set(CMAKE_CXX_STANDARD 14)
set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads)
target_link_libraries(vm
        Threads::Threads
        ${CMAKE_DL_LIBS})
