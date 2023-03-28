#include <stdio.h>
#include <stdlib.h>

/*
int a[3] = {3, 1, 2};
int compare(unsigned long i, unsigned long j) {
    if(a[i] < a[j]) return -1;
    if(a[i] == a[j]) return 0;
    return 1;
}
void swap(unsigned long i, unsigned long j) {
    int tmp = a[i];
    a[i] = a[j];
    a[j] = tmp;
    return;
}
*/

void shellsort(unsigned long nel,
               int (*compare)(unsigned long i, unsigned long j),
               void (*swap)(unsigned long i, unsigned long j))
{
    if(nel == 1) return;
    unsigned long fib[200];
    fib[0] = 1;
    fib[1] = 2;
    int counter = 1;
    while(fib[counter] < nel) {
        ++counter;
        fib[counter] = fib[counter - 1] + fib[counter - 2];
    }
    --counter;
    for(int i = counter; i >= 0; --i) {
        unsigned long j = fib[i];
        while(j < nel) {
            unsigned long loc = j;
            while(loc >= fib[i] && compare(loc, (loc - fib[i])) == -1) {
                swap(loc, (loc - fib[i]));
                loc -= fib[i];
            }
            ++j;
        }
    }
}

int main(int argc, char ** argv) {
    /* shellsort(3, compare, swap);
       for(int i = 0; i < 3; ++i) {
           printf("%d ", a[i]);
       } */
    return 0;
}