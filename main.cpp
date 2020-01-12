#include <iostream>
#include <vector>
#include <stack>
#include <queue>
using namespace std;

class RadixSort {
public:
    int* radixSort(int* A, int n) {
        // write code here
        queue<int> a[10];
        int i;
        int k=1;
        for(int m=0;m<3;m++) {
            for(i=0;i<n;i++) {
                int cur = (A[i]/k)%10;
                a[cur].push(A[i]);
            }
            i=0;

            while(i<n) {
                for(int j=0;j<10;j++) {
                    while(!a[j].empty()) {
                        A[i] = a[j].front();
                        a[j].pop();
                        i++;
                    }
                }
            }

            k = k*10;
        }
        return A;
    }
};
int main() {
    int a[] = {54,35,48,36,27,12,44,44,8,14,26,17,28};
    RadixSort l;
    l.radixSort(a, 13);
    return 0;
}