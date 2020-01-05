//
// Created by zhandong on 2020-01-03.
//

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


class MergeSort {
public:
    int *mergeSort(int *A, int n) {
        // write code here
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
