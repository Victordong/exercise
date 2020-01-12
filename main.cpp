#include <iostream>
#include <vector>
#include <stack>

using namespace std;

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

int main() {
    vector<int> a;
    int b[4] = {2,4,1,3};
    for(int i=0;i<4;i++) {
        a.push_back(b[i]);
    }
    MaxTree l;
    l.buildMaxTree(a, 4);
    return 0;
}