#include <iostream>
#include <vector>
#include <stack>
#include <queue>
using namespace std;

class ScaleSort {
public:
    void heap(vector<int> &A, int root, int k)
    {
        int smallest = root, left = root*2 + 1, right = left + 1;
        if(left < k && A[smallest] > A[left])
            smallest = left;
        if(right < k && A[smallest] > A[right])
            smallest = right;
        if(smallest != root)
        {
            swap(A[root], A[smallest]);
            heap(A, smallest, k);
        }
    }

    vector<int> sortElement(vector<int> A, int n, int k) {
        // write code here

        int i, j;
        vector<int> help(k, 0);
        for(i = 0; i < k; ++i)
            help[i] = A[i];
        for(i = k/2-1; i >= 0; --i)
            heap(help, i, k);
        for(i = 0; i < n-k; ++i)
        {
            A[i] = help[0];
            help[0] = A[i+k];
            heap(help, 0, k);
        }
        for(i = k; i > 0; --i)
        {
            A[n-i] = help[0];
            swap(help[0], help[i-1]);
            heap(help, 0, i-1);
        }
        return A;
    }
};
int main() {
    int a[] = {2,1,4,3,6,5,8,7,10,9};
    vector<int>A;
    for(int i=0;i<15;i++) {
        A.push_back(a[i]);
    }
    ScaleSort l;
    l.sortElement(A, 10, 2);

    return 0;
}