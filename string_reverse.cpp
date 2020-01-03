//
// Created by zhandong on 2020-01-03.
//

#include <string>
using namespace std;

class Rotation {
public:
    bool chkRotation(string A, int lena, string B, int lenb) {
        // write code here
        if(lena!= lenb) {
            return false;
        }
        string new_string = B+B;
        return new_string.find(A)!= string::npos;
    }
};