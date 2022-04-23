#include <stdio.h>

int main(int argc, char ** argv) {
    int a, b, x;
    scanf("%d%d", &a, &b);
    int arr1[b], arr2[a], arr3[b], arr4[a];
    for(int i = 0; i < b; ++i) {
        arr1[i] = (1 << 31) - 1;
    }
    for(int i = 0; i < a; ++i) {
        arr2[i] = (1 << 31);
    }
    for(int i = 0; i < a; ++i) {
        for(int j = 0; j < b; ++j) {
            scanf("%d", &x);
            if(x < arr1[j]) {
                arr1[j] = x;
                arr3[j] = j;
            }
            if(x > arr2[i]) {
                arr2[i] = x;
                arr4[i] = i;
            }
        }
    }
    char flag = 1;
    for(int i = 0; i < a; ++i) {
        for(int j = 0; j < b; ++j) {
            if(arr1[j] == arr2[i]) {
                printf("%d %d", arr4[i], arr3[j]);
                flag = 0;
                break;
            }
        }
    }
    if(flag) printf("none");
    return 0;
}
