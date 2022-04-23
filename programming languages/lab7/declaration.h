//
// Created by Alexander on 28.03.2022.
//

#ifndef LAB7_DECLARATION_H
#define LAB7_DECLARATION_H

#endif //LAB7_DECLARATION_H

class Queue {
private:
    class DoubleStack {
    private:
        int *arr;
        int top1, top2, cap;
    public:
        explicit DoubleStack(int);
        void Push1(int);
        void Push2(int);
        int Pop1();
        int Pop2();
        bool StackEmpty1();
        bool StackEmpty2();
        int& operator [] (int idx);
    };
    DoubleStack queue = DoubleStack(0);
    int numbers = 0;
public:
    explicit Queue(int);
    void Enqueue(int);
    int Dequeue();
    int getSize();
    int& operator [] (int idx);
};
