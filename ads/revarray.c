#include <stdio.h>

void revarray(void *base, size_t nel, size_t width) {
    void *p, *q;
    int i, tmp;
    for(i = 0, p = base, q = base + (nel - 1) * width; i < nel/2; p += width, q -= width, ++i) {
        tmp = *(int *)p;
        *(int *)p = *(int *)q;
        *(int *)q = tmp;
    }
}

int main(int argc, char ** argv) {
    return 0;
}
