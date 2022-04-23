#include <stdio.h>

int main(int argc, char ** argv) {
    long long x, f1 = 1, f2 = 2, f = 2;
    int k = 0;
    scanf("%lld", &x);
    if(x == 0) {
        printf("0");
        return 0;
    }
    while(f <= x) {
        if(f2 == 7540113804746346429) {
            f1 = f2;
            ++k;
            break;
        }
        f = f1 + f2;
        f1 = f2;
        f2 = f;
        ++k;
    }
    int arr[k + 1];
    for(int i = 0; i < k + 1; ++i) {
        arr[i] = 0;
    }
    int n = k + 1;
    arr[n - k - 1] = 1;
    x -= f1;
    k =  0;
    f1 = 1;
    f2 = f = 2;
    while(x > 0) {
        while(f <= x) {
            f = f1 + f2;
            f1 = f2;
            f2 = f;
            ++k;
        }
        arr[n - k - 1] = 1;
        x -= f1;
        k =  0;
        f1 = 1;
        f2 = f = 2;
    }
    for(int i = 0; i < n; ++i) {
        printf("%d", arr[i]);
    }
    return 0;
}
