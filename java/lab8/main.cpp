#include <iostream>
#include "Equality.h"

int main() {
    Equality<bool> x(3, true);
    vector<bool> tmp = {true, false, false};
    x.setVector(tmp);
    Equality<int> y(3, 10);
    y.setVector(3);
    vector<int> tmp1 = {1, 2, 3, 4};
    vector<int> tmp2 = {3, 1, 0};
    vector<bool> tmp3 = {true, false, true};


    if(x.isEqual(tmp3)) {
        cout << "Equal" << endl;
    } else {
        cout << "Not equal" << endl;
    }

    if(y.isEqual(tmp1)) {
        cout << "Equal" << endl;
    } else {
        cout << "Not equal" << endl;
    }

    if(y.isEqual(tmp2)) {
        cout << "Equal" << endl;
    } else {
        cout << "Not equal" << endl;
    }
    return 0;
}
