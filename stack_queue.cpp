//
// Created by zhandong on 2020-01-11.
//
#include <stack>
#include <vector>
using namespace std;

class Solution {
public:
    stack<int> stack_value;
    stack<int> stack_min;
    void push(int value) {
        stack_value.push(value);
        if(stack_min.empty()||value<=stack_min.top()) {
            stack_min.push(value);
        } else {
            stack_min.push(stack_min.top());
        }
    }
    void pop() {
        if(!stack_value.empty()) {
            stack_min.pop();
            stack_value.pop();
        }
    }
    int top() {
        return stack_value.top();
    }
    int min() {
        return stack_min.top();
    }
};


class TwoStacks {
public:
    vector<int> twoStacksSort(vector<int> numbers) {
        // write code here
        stack<int> help;
        while(!numbers.empty()) {
            int cur = numbers.front();
            numbers.erase(numbers.begin());
            if(help.empty()) {
                help.push(cur);
            }else if (cur<=help.top()) {
                help.push(cur);
            } else {
                while(!help.empty()&&help.top()<cur) {
                    int now = help.top();
                    help.pop();
                    numbers.insert(numbers.begin(), now);
                }
                help.push(cur);
            }
        }
        while(!help.empty()) {
            int cur = help.top();
            help.pop();
            numbers.insert(numbers.begin(), cur);
        }
        return numbers;
    }
};

class TwoStack {
public:
    vector<int> twoStack(vector<int> ope, int n) {
        // write code here
        stack<int> stack1, stack2;
        vector<int> result;
        for (int i = 0; i < n; i++){
            if (ope[i] != 0){
                stack1.push(ope[i]);
            }
            else{
                if (!stack2.empty()){
                    result.push_back(stack2.top());
                    stack2.pop();
                }
                else{
                    while(!stack1.empty()){
                        stack2.push(stack1.top());
                        stack1.pop();
                    }
                    result.push_back(stack2.top());
                    stack2.pop();
                }
            }
        }
        return result;
    }
};


class StackReverse {
public:
    int get(vector<int>&A){
        int cur = A.front();
        A.erase(A.begin());
        if(A.empty()) {
            return cur;
        } else {
            int val = get(A);
            A.insert(A.begin(), cur);
            return val;
        }
    }

    void reverse(vector<int>&A) {
        if(A.empty()) {
            return;
        } else {
            int val = get(A);
            reverse(A);
            A.insert(A.begin(), val);
            return;
        }
    }

    vector<int> reverseStack(vector<int> A, int n) {
        reverse(A);
        return A;
    }
};


class SlideWindow {
public:
    vector<int> slide(vector<int> arr, int n, int w) {
        vector<int> stack_value;
        vector<int> result;
        for(int i=0;i<n;i++) {
            if(stack_value.empty()) {
                stack_value.push_back(i);
            } else {
                if(arr[i]<arr[stack_value.back()]) {
                    stack_value.push_back(i);
                } else {
                    while(arr[stack_value.back()]<=arr[i]&&!stack_value.empty()) {
                        stack_value.erase(stack_value.end()-1);
                    }
                    stack_value.push_back(i);
                }
            }
            if(stack_value.front()<=i-w) {
                stack_value.erase(stack_value.begin());
            }
            if(i>=w-1) {
                int val = arr[stack_value.front()];
                result.push_back(val);
            }
        }
        return result;
    }
};

class MaxTree {
public:
    vector<int> buildMaxTree(vector<int> A, int n) {
        // write code here
        stack<int> main;
        stack<int> help;
        vector<int> left;
        vector<int> right;
        vector<int> result;
        for(int i=0;i<n;i++) {
            if(main.empty()) {
                main.push(i);
                left.push_back(-1);
            } else {
                while(!main.empty()&&A[main.top()]<A[i]) {
                    int cur = main.top();
                    main.pop();
                    help.push(cur);
                }
                if(main.empty()) {
                    left.push_back(-1);
                } else {
                    left.push_back(main.top());
                }
                main.push(i);
            }
        }
        while(!main.empty()) {
            main.pop();
        }
        for(int i=n-1;i>=0;i--) {
            if(main.empty()) {
                main.push(i);
                right.push_back(-1);
            } else {
                while(!main.empty()&&A[main.top()]<A[i]) {
                    int cur = main.top();
                    main.pop();
                    help.push(cur);
                }
                if(main.empty()) {
                    right.insert(right.begin(),-1);
                } else {
                    right.insert(right.begin(),main.top());
                }
                main.push(i);
            }
        }
        for(int i=0;i<n;i++) {
            if(left[i]<0 && right[i]<0) {
                result.push_back(-1);
            } else if(right[i]<0&&left[i]>=0) {
                result.push_back(left[i]);
            } else if(right[i]>=0&&left[i]<0) {
                result.push_back(right[i]);
            } else {
                if(A[left[i]]<A[right[i]])
                    result.push_back(left[i]);
                else
                    result.push_back(right[i]);
            }
        }
        return result;
    }
};
