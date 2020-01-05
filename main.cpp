#include <iostream>

int main() {
    int A[13] = {54,35,48,36,27,12,44,44,8,14,26,17,28};
    QuickSort q;
    int* B = q.quickSort(A, 13);
    for(int i=0;i<13;i++) {
        std::cout<<B[i]<<",";
    }
    std::cout<<std::endl;
    return 0;
}