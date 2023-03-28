#include <iostream>
#include "SparseMatrix.h"

using namespace std;

int main() {
    SparseMatrix<int, 5, 5> m1, m2;
    cout << "M1:" << endl;
    m1.setElem(1, 3, 3);
    m1.setElem(2, 1, 4);
    m1.setElem(2, 3, 5);
    m1.setElem(3, 1, 6);
    m1.setElem(3, 2, 7);
    m1.setElem(3, 3, 8);
    m1.setElem(4, 0, 9);
    m1.show();
    cout << endl << "M2:" << endl;
    m2.setElem(3, 1, 5);
    m2.setElem(0, 0, 2);
    m2.setElem(3, 3, 3);
    m2.setElem(2, 4, 7);
    m2.setElem(0, 4, 1);
    m2.show();
    cout << endl << "M1 + M2:" << endl;
    (m1 + m2).show();
    cout << endl << "M1 * M2" << endl;
    (m1 * m2).show();
    cout << endl << "M1 * 3" << endl;
    (m1 * 3).show();
    cout << endl << "M1 == M2:" << endl;
    if(m1 == m2) {
        cout << "true";
    } else {
        cout << "false";
    }
    return 0;
}