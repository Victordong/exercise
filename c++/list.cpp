//
// Created by zhandong on 2020-01-12.
//

#include<vector>

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
        ListNode *new_node = new ListNode(val);
        if (A.empty()) {
            new_node->next = new_node;
            return new_node;
        }
        ListNode *head = new ListNode(A[0]);
        ListNode *cur = head;
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
        return head;
    }
};
class Remove {
public:
    ListNode* removeNode(ListNode* pHead, int delVal) {
        // write code here
        if(pHead->val == delVal) {
            ListNode *node = pHead->next;
            pHead->next = NULL;
            return node;
        }
        ListNode *pre = pHead;
        ListNode *cur = pHead->next;
        while(cur != NULL)
        {
            if(cur->val == delVal)
            {
                pre->next = cur->next;
                cur->next = NULL;
                break;
            }
            pre = pre->next;
            cur = cur->next;
        }
        return pHead;
    }
};