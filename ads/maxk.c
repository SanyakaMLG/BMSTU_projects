#include <stdio.h>

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    int arr[n];
    for(int i = 0; i < n; ++i) {
        scanf("%d", &arr[i]);
    }
    int k, sum = 0, tmpsum = 0;
    scanf("%d", &k);
    for(int i = 0; i < k; ++i) {
        sum += arr[i];
    }
    tmpsum = sum;
    for(int i = k; i < n; ++i) {
        tmpsum = tmpsum - arr[(i/k - 1)*k + i%k] + arr[i/k * k + i%k];
        if(tmpsum > sum) sum = tmpsum;
    }
    printf("%d", sum);
    return 0;
}
