//
// Created by zhandong on 2020-01-03.
//
#include <iostream>
#include <vector>
#include <queue>
using namespace std;

class BubbleSort {
public:
    int *bubbleSort(int *A, int n) {
        // write code here
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n - i - 1; j++) {
                if (A[j] > A[j + 1]) {
                    int swap = A[j];
                    A[j] = A[j + 1];
                    A[j + 1] = swap;
                }
            }
        }
        return A;
    }
};

class SelectionSort {
public:
    int *selectionSort(int *A, int n) {
        // write code here
        for (int i = 0; i < n; i++) {
            int max = 0;
            for (int j = 0; j < n - i; j++) {
                if (A[max] < A[j]) {
                    max = j;
                }
            }
            int swap = A[n - i - 1];
            A[n - i - 1] = A[max];
            A[max] = swap;
        }
        return A;
    }
};


class InsertionSort {
public:
    int *insertionSort(int *A, int n) {
        // write code here
        for (int i = 0; i < n; i++) {
            int j = i;
            int temp = A[j];
            while (j > 0) {
                if (temp <= A[j - 1]) {
                    A[j] = A[j - 1];
                    j--;
                } else {
                    break;
                }

            }
            A[j] = temp;
        }
        return A;
    }
};


class QuickSort {
public:
    int partSort(int *A, int left, int right) {
        int swap = A[left];
        while (left < right) {
            while (A[right]>=swap && right> left) {
                right--;
            }
            A[left] = A[right];
            while (A[left] <=swap && left< right) {
                left++;
            }
            A[right] = A[left];
        }
        A[left] = swap;
        return left;
    }

    void quickSortMain(int *A, int left, int right) {
        if(left<right) {
            int index = partSort(A, left, right);
            quickSortMain(A, left, index-1);
            quickSortMain(A, index+1, right);
        }
    }

    int *quickSort(int *A, int n) {
        quickSortMain(A, 0, n-1);
        return A;
    }
};


class MergeSort {
public:
    void merge_part(int* A, int start, int mid, int end) {
        int temp[1000];
        int i=start;
        int j=mid+1;
        int k=0;
        while(i<=mid&&j<end) {
            if(A[i]<A[j]) {
                temp[k] = A[i];
                i+=1;
            } else {
                temp[k] = A[j];
                j+=1;
            }
            k+=1;
        }
        if(i==mid+1) {
            while(j<end) {
                temp[k] = A[j];
                k+=1;
                j+=1;
            }
        }else if(j==end) {
            while(i<=mid) {
                temp[k] = A[i];
                k+=1;
                i+=1;
            }
        }
        for(i=start,k=0;i<end;i++,k++) {
            A[i] = temp[k];
        }
    }
    void merge(int* A, int start, int end,int n) {
        if(start>=end) {
            return;
        }
        int mid = (start+end)/2;
        merge(A, start, mid, n);
        merge(A, mid+1, end, n);
        printList(A, n);
        merge_part(A, start, mid, end);
    }
    void printList(int *A, int n) {
        for(int i=0;i<n;i++) {
            cout<<A[i]<<" ";
        }
        cout<<endl;
    }
    int* mergeSort(int* A, int n) {
        // write code here
        merge(A,0,n-1, n);
        return A;
    }
};


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

class CountingSort {
public:
    int* countingSort(int* A, int n) {
        // write code here
        vector<int> help;
        int min = A[0];
        int max = A[0];
        for(int i=0;i<n;i++) {
            if(max<A[i])
                max = A[i];
            if(min>A[i])
                min = A[i];
        }
        for(int i=min;i<=max;i++) {
            help.push_back(0);
        }
        for(int i=0;i<n;i++) {
            help[A[i]-min] +=1;
        }
        int m = 0;
        for(int i=0;i<=max-min;i++) {
            int count = help[i];
            while(count>0) {
                A[m] = i+min;
                count-=1;
                m+=1;
            }
        }
        return A;
    }
};


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
        return A;
    }
};