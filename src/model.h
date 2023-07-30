#ifndef MODEL_H
#define MODEL_H

#include <string>
#include <stack>

class Calculator {
public:
    Calculator();

    void pushNumber(double num);
    void pushOperator(char op);
    double calculate();

private:
    std::stack<double> numberStack;
    std::stack<char> operatorStack;
};

#endif // MODEL_H
