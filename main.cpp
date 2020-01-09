#include <iostream>
#include <vector>

using namespace std;

class Backpack {
public:
    int maxValue(vector<int> w, vector<int> v, int n, int cap) {
        int a[1000][1000];
        for(int i=0;i<n;i++) {
            a[i][0] = 0;
        }
        for(int i=1;i<cap;i++) {
            if(i>=w[0]) {
                a[0][i] = v[0];
            } else {
                a[0][i] = 0;
            }
        }
        for(int i=0;i<n;i++) {
            for(int j=0;j<=cap;j++) {
                cout<<a[i][j]<<" ";
            }
            cout<<"fin"<<endl;
        }
        for(int i=1;i<n;i++) {
            for(int j=1;j<=cap;j++) {
                if(j>=w[i]) {
                    a[i][j] = max(a[i-1][j], a[i-1][j-w[i]]+v[i]);
                } else {
                    a[i][j] =a[i-1][j];
                }
            }
        }
//        for(int i=0;i<n;i++) {
//            for(int j=0;j<=cap;j++) {
//                cout<<a[i][j]<<" ";
//            }
//            cout<<"fin"<<endl;
//        }
        return a[n-1][cap];
    }
};

int main() {
    int a[8] = {27,26,41,29,26,25,38};
    int b[8] = {274,153,595,431,534,586,364};
    int n=7;
    vector<int> A;
    vector<int> B;
    for(int i=0;i<n;i++) {
        A.push_back(a[i]);
        B.push_back(b[i]);
    }
    Backpack l;
    cout<<l.maxValue(A, B, n, 470);
    return 0;
}