#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void revarray(void *base, unsigned long nel, unsigned long width)
{
    char t;
    for (unsigned long i = 0; i < (nel / 2); i++) {
        for (unsigned long j = 0; j < width; j++) {
            t = *((char*)base);
            *((char*)base) = *((char*)((char*)base + (nel*width - width - i * 2 * width)));
            *((char*)((char*)base + (nel*width - width - i * 2 * width))) = t;
            base = (char*)base + 1;
        }
    }
}

int main(int argc, char ** argv) {
    int *arr = (int*)malloc(sizeof(int) * 10);
    for(int i = 0; i < 10; ++i) {
        scanf("%d", &arr[i]);
    }
    revarray(arr, 10, sizeof(int));
    for(int i = 0; i < 10; ++i) {
        printf("%d ", arr[i]);
    }
    free(arr);
    return 0;
}