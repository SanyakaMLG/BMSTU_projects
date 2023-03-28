#include <stdio.h>

unsigned long peak(unsigned long nel,
                   int (*less)(unsigned long i, unsigned long j)) {
    if(nel == 1) return 0;
    if(less(1, 0)) return 0;
    for(int i = 1; i < nel - 1; ++i) {
        if(less(i - 1, i) && less(i + 1, i)) return i;
    }
    if(less(nel - 2, nel - 1)) return nel - 1;
}

int main(int argc, char ** argv) {
    return 0;
}
