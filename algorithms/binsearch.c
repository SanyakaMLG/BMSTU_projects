#include <stdio.h>

unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i)) {
    int L = -1, R = nel - 1;
    while(R != L + 1) {
        if(compare(R - (R   - L)/2) == 0) {
            return R - (R - L)/2;
        } else {
            if(compare(R - (R - L)/2) == -1) {
                L = R - (R - L)/2;
            } else {
                R = R - (R - L)/2;
            }
        }
    }
    if(compare(R - (R   - L)/2) == 0) {
        return R - (R - L)/2;
    }
    return nel;
}

int main(int argc, char ** argv) {
    return 0;
}