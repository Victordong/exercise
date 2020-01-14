#include <iostream>
#include <vector>
#include <stack>
#include <queue>
using namespace std;

struct ListNode {
    int val;
    struct ListNode *next;

    ListNode(int x) : val(x), next(NULL) {}
};

class InsertValue {
public:
    ListNode* insert(vector<int> A, vector<int> nxt, int val) {
        // write code here
        auto new_node = new ListNode(val);
        if (A.empty()) {
            new_node->next = new_node;
            return new_node;
        }
        auto head = new ListNode(A[0]);
        ListNode* cur = head;
        for (int i = 0; i < A.size(); i++) {
            ListNode *l;
            if (nxt[i] != 0)
                l = new ListNode(A[nxt[i]]);
            else
                l = head;
            cur->next = l;
            cur = l;
        }
        cur = head;
        ListNode* next = head->next;
        while(next!=head) {
            if (val <= next->val && val >= cur->val) {
                break;
            }
            cur = next;
            next = cur->next;
        }
        if(next==head&&val<head->val) {
            head = new_node;
        }
        new_node->next = next;
        cur->next = new_node;
        cur = head;
        while(cur->next!=head) {
            cout<<cur->val<<" ";
            cur = cur->next;
        }
        cout<<cur->val<<" "<<cur->next->val<<endl;
        return head;
    }
};

int main() {
    int a[] = {1,3,4,5,7};
    int b[] = {1,2,3,4,0};
    vector<int> A;
    vector<int> B;
    for(int i=0;i<5;i++) {
        A.push_back(a[i]);
        B.push_back(b[i]);
    }
    InsertValue i;
    i.insert(A,B,2);
//    r.reverse(A, 4,9);
    return 0;
}