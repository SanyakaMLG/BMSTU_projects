#include <iostream>
#include "Matrix.h"

using namespace std;

int main() {
    cout << "Source matrix:\n";
    Matrix<3, 3> m;
    m.fillMatrix();
    auto it = m.begin();
    cout << "GCDS:\n";
    while(it != m.end()) {
        cout << *it << endl;
        it++;
    }
    it = m.begin();
    while(it != m.end()) {
        it = 3;
        it++;
    }
    cout << "Matrix after transformation gcds:\n";
    m.printMatrix();
    return 0;
}
