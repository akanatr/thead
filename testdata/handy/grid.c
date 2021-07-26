#include <handy.h>
#include <stdio.h>

int main() {
    int i;        //カウンタ変数
    int num = 0;  //入力値
    int dist;     //線と線の距離

    printf("Input number:");
    scanf("%d", &num);

    dist = 600 / num;  //線と線の距離

    HgOpen(600, 600);

    for (i = 0; i < num; i++) {
        HgLine(0, i * dist, 600, i * dist);  //横線
        HgLine(i * dist, 0, i * dist, 600);  //横線
    }

    HgGetChar();
    HgClose();
    return 0;
}