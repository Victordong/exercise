#include <iostream>
#include <vector>
#include <stack>
#include <queue>
using namespace std;

class HeapSort {
public:
    void swap(int* A, int first, int second) {
        int swap = A[first];
        A[first] = A[second];
        A[second] = swap;
    }
    void mergePart(int* A, int start, int end) {
        int father = start;
        int son = (father * 2) +1;
        while(son<=end) {
            if(son+1<=end && A[son]<A[son+1]) {
                son++;
            }
            if(A[father]>A[son]) {
                return;
            } else {
                swap(A, father, son);
                father = son;
                son = father*2+1;
            }
        }
    }

    int* heapSort(int* A, int n) {
        // write code here
        for(int i=(n/2)-1;i>=0;i--) {
            mergePart(A, i, n-1);
        }
        for(int i=n-1;i>=0;i--) {
            swap(A, 0, i);
            mergePart(A, 0, i-1);
        }
        for(int i=0;i<n;i++) {
            cout<<A[i]<<" ";
        }
        return A;
    }
};
int main() {
    int a[] = {54,35,48,36,27,12,44,44,8,14,26,17,28};
    HeapSort l;
    l.heapSort(a, 13);

    return 0;
}