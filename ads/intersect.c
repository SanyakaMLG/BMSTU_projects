#include <stdio.h>

int main(int argc, char ** argv) {
    int n, m, a = 0, b = 0, a32 = 0, b32 = 0;
    scanf("%d", &n);
    for (int i = 0; i < n; ++i) {
        scanf("%d", &m);
        if(m != 31) {
            a += (1 << m);
        } else {
            a32 = 1;
        }
    }
    scanf("%d", &n);
    for(int i = 0; i < n; ++i) {
        scanf("%d", &m);
        if(m != 31) {
            b += (1 << m);
        } else {
            b32 = 1;
        }
    }
    for(int i = 0; i < 31; ++i) {
        if((a >> i) % 2 && (b >> i) % 2) {
            printf("%d ", i);
        }
    }
    if(a32 && b32) printf("%d", 31);
    return 0;
}
