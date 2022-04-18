#include <iostream>
#include "declaration.h"

using namespace std;

int main() {
    int size;
    cin >> size;
    Queue x(size);
    for(int i = 0; i < 5; i++) {
        x.Enqueue(i);
    }
    /*cout << x.getSize() << endl;
    cout << x.Dequeue() << endl;
    cout << x.Dequeue() << endl;
    cout << x.getSize() << endl;
    cout << x.Dequeue() << endl;
    x.Enqueue(10);
    cout << x.Dequeue() << endl;*/
    x.Enqueue(10);
    x.Dequeue();
    cout << x.getSize() << endl;
    for(int i = 0; i < 5; i++) {
        cout << x[i] << " ";
    }
    cout << endl;
    cout << x.Dequeue() << endl;
    return 0;
}
