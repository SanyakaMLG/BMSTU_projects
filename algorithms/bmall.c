#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

int max(int a, int b) {
    return a > b ? a : b;
}

int *Delta(char *S, int size) {
    int *delta1 = (int*)malloc(size * sizeof(int));
    int length = strlen(S);
    for(int i = 0; i < size; ++i) {
        delta1[i] = length;
    }
    for(int i = 0; i < length; ++i) {
        delta1[(int)S[i] - 33] = length - i - 1;
    }
    return delta1;
}

void SimpleBMSubst(char *S, char *T, int size) {
    int *delta = Delta(S, size);
    int length_S = strlen(S), length_T = strlen(T);
    int k = length_S - 1;
    while(k < length_T) {
        int i = length_S - 1;
        while(T[k] == S[i]) {
            if(i == 0) {
                printf("%d ", k);
                break;
            }
            i--;
            k--;
        }
        k = k + max(delta[(int)T[k] - 33], length_S - i);
    }
    free(delta);
}

int main(int argc, char ** argv) {
    char *S, *T;
    S = argv[1];
    T = argv[2];
    SimpleBMSubst(S, T, 94);
    return 0;
}