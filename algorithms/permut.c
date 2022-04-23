#include <stdio.h>

int main(int argc, char ** argv) {
    long long a[8], b[8];
    for(int i = 0; i < 8; ++i) {
        scanf("%lld", &a[i]);
    }
    for(int i = 0; i < 8; ++i) {
        scanf("%lld", &b[i]);
    }
    int k1 = 0, k2 = 0, check;
    for(int i = 0; i < 8; ++i) {
        for(int j = 0; j < 8; ++j) {
            if(a[i] == b[j]) ++k2;
        }
        check = a[i];
        for(int q = 0; q < 8; ++q) {
            if(a[q] == check) ++k1;
        }
        if(k1 != k2 || k2 == 0) {
            printf("no");
            return 0;
        }
        k1 = k2 = 0;
    }
    printf("yes");
    return 0;
}
