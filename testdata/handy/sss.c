/*****
 boxes.c
 長方形とひし形を描く
 Tonan Shogo
*****/
 
#include <stdio.h>
#include <handy.h>
 
int main() {
    int inputx0;
    int inputy0;
    int inputx1;
    int inputy1;
 
    printf("inputx0:\n");
    scanf("%d",&inputx0);
    printf("inputy0:\n");
    scanf("%d",&inputy0);
    printf("inputx1:\n");
    scanf("%d",&inputx1);
    printf("inputy1:\n");
    scanf("%d",&inputy1);
    
    HgOpen(600,600);
    
    
    HgBox(inputx0,inputy0,inputx1-inputx0,inputy1-inputy0);
    
    HgBox(inputx0/2,inputy0/4*3,inputx1,inputy1/4+inputy0/4*3);
    
    HgLine(inputx0/2,inputy1/2+inputy0/2,inputx1/2+inputx0/2,inputy0/4*3+inputy1/4+inputy0/4*3);
    
    HgLine(inputx0/2,inputy1/2+inputy0/2,inputx1/2+inputx0/2,inputy0/4*3);
    
    HgLine(inputx1/2+inputx0/2,inputy0/2*3+inputy1/4,inputx1+inputx0/2,inputy1/2+inputy0/2);
    
    HgLine(inputx1/2+inputx0/2,inputy0/4*3,inputx1+inputx0/2,inputy1/2+inputy0/2);
    
    HgGetChar();
    HgClose();
    return 0;
}