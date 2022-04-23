#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

int max(int a, int b) {
    return a > b ? a : b;
}

int **Delta(char *S, int size) {
    int length = strlen(S);
    int **delta1 = (int**)malloc(length * sizeof(int*));
    for(int i = 0; i < length; ++i) {
        delta1[i] = (int*)malloc(sizeof(int) * size);
        for(int j = 0; j < size; ++j) delta1[i][j] = length - i;
    }
    for(int i = 0; i < length; ++i) {
        for(int j = 0; j < (length - i - 1); ++j) delta1[i][(int)S[j] - 33] = length - j - i - 1;
    }
    return delta1;
}

int SimpleBMSubst(char *S, char *T, int size) {
    int **delta = Delta(S, size);
    int length_S = strlen(S), length_T = strlen(T);
    int k = length_S - 1;
    while(k < length_T) {
        int i = length_S - 1;
        int k1 = k;
        while(T[k1] == S[i]) {
            if(i == 0) {
                for(int j = 0; j < length_S; ++j) {
                    free(delta[j]);
                }
                free(delta);
                return k1;
            }
            i--;
            k1--;
        }
        k = k + delta[length_S - i - 1][(int)T[k1] - 33];
    }
    for(int i = 0; i < length_S; ++i) {
        free(delta[i]);
    }
    free(delta);
    k = length_T;
    return k;
}

int main(int argc, char ** argv) {
    char *S, *T;
    /*S = (char*)malloc(200);
    T = (char*)malloc(200);
    gets(S);
    gets(T);
    printf("%d", SimpleBMSubst(S, T, 94));
    free(S);
    free(T);*/
    S = argv[1];
    T = argv[2];
    printf("%d", SimpleBMSubst(S, T, 94));
    return 0;
}