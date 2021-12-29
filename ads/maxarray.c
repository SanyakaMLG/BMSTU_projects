#include <stdio.h>

int comp(void * elem1, void * elem2)
{
    int f = *((int*)elem1);
    int s = *((int*)elem2);
    if (f > s) return  1;
    if (f < s) return -1;
    return 0;
}

int maxarray(void *base, size_t nel, size_t width,
             int(*compare)(void *a, void *b))
{
    void *p = base;
    int *max = (int *)p;
    for(int i = 0; i < nel; ++i, p += width) {
        if(compare(p, max) == 1) max = p;
    }
    return *max;
}

int main(int argc, char ** argv) {
    double arr[10];
    for(int i = 0; i < 10; ++i) {
        scanf("%d", &arr[i]);
    }
    printf("%d", maxarray(arr, 10, sizeof(int), comp));
    return 0;
}
