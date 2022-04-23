#include <stdio.h>
#include <stdlib.h>

/*int a[0];

int compare(unsigned long i, unsigned long j) {
    if(a[i] < a[j]) return -1;
    if(a[i] == a[j]) return 0;
    return 1;
}

void swap(unsigned long i, unsigned long j) {
    int tmp;
    tmp = a[i];
    a[i] = a[j];
    a[j] = tmp;
}
*/

void bubblesort(unsigned long nel,
                int  (*compare)(unsigned long i, unsigned long j),
                void (*swap)(unsigned long i, unsigned long j))
{
    if(nel == 0) return;
    unsigned long tLEFT = 0;
    unsigned long tRIGHT = nel - 1;
    unsigned long bound;
    unsigned long count = -1;
    while(count != 0) {
        count = 0;
        bound = tRIGHT;
        tRIGHT = 0;
        unsigned long i = tLEFT;
        while(i < bound) {
            if(compare(i, i + 1) == 1) {
                swap(i, i + 1);
                tRIGHT = i;
                count++;
            }
            ++i;
        }
        i = tRIGHT;
        bound = tLEFT;
        tLEFT = nel - 1;
        while(i > bound) {
            if(compare(i, i - 1) == -1) {
                swap(i, i - 1);
                tLEFT = i;
                count++;
            }
            --i;
        }
    }
}

int main()
{
/*(    bubblesort(0, compare, swap);
    for(int i = 0; i < 0; ++i) {
        printf("%d ", a[i]);
    }*/
    return 0;
}