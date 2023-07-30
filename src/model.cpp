#include "model.h"

Calculator::Calculator() {}

void Calculator::pushNumber(double num) {
    numberStack.push(num);
}

void Calculator::pushOperator(char op) {
    while (!operatorStack.empty() && operatorStack.top() != '(') {
        double num2 = numberStack.top(); numberStack.pop();
        double num1 = numberStack.top(); numberStack.pop();

        if (operatorStack.top() == '+')
            numberStack.push(num1 + num2);
        else if (operatorStack.top() == '-')
            numberStack.push(num1 - num2);
        else if (operatorStack.top() == '*')
            numberStack.push(num1 * num2);
        else if (operatorStack.top() == '/')
            numberStack.push(num1 / num2);

        operatorStack.pop();
    }

    if (op != ')')
        operatorStack.push(op);
    else
        operatorStack.pop();
}

double Calculator::calculate() {
    while (!operatorStack.empty()) {
        double num2 = numberStack.top(); numberStack.pop();
        double num1 = numberStack.top(); numberStack.pop();

        if (operatorStack.top() == '+')
            numberStack.push(num1 + num2);
        else if (operatorStack.top() == '-')
            numberStack.push(num1 - num2);
        else if (operatorStack.top() == '*')
            numberStack.push(num1 * num2);
        else if (operatorStack.top() == '/')
            numberStack.push(num1 / num2);

        operatorStack.pop();
    }

    return numberStack.top();
}
