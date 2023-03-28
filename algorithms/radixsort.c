#include <stdio.h>
#include <stdlib.h>
#include <string.h>

union Int32 {
    int x;
    unsigned char bytes[4];
};

void radixsort(union Int32 *nums, int n) {
    for(int ind = 0; ind < 4; ++ind) {
        int count[256];
        for (int i = 0; i < 256; ++i) {
            count[i] = 0;
        }
        for (int i = 0; i < n; ++i) {
            if (ind == 3) {
                count[(nums[i].bytes[ind]) ^ 128]++;
            } else {
                count[(nums[i].bytes[ind])]++;
            }
        }
        for (int i = 1; i < 256; i++) {
            count[i] += count[i - 1];
        }
        union Int32 sorted[n];
        for (int i = n - 1; i >= 0; --i) {
            if (ind == 3) {
                int j = count[(nums[i].bytes[ind]) ^ 128] - 1;
                count[(nums[i].bytes[ind]) ^ 128] = j;
                sorted[j] = nums[i];
            } else {
                int j = count[(nums[i].bytes[ind])] - 1;
                count[(nums[i].bytes[ind])] = j;
                sorted[j] = nums[i];
            }
        }
        for (int i = 0; i < n; ++i) {
            nums[i] = sorted[i];
        }
    }
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    union Int32 nums[n];
    for(int i = 0; i < n; i++) {
        scanf("%d", &nums[i].x);
    }
    radixsort(nums, n);
    for(int i = 0; i < n; ++i) {
        printf("%d ", nums[i].x);
    }
    return 0;
}