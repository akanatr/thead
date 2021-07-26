#include<stdio.h>
#include<handy.h>

int main(){
    int x, y, r;
    int distance, radius;

    printf("input x:");
    scanf("%d", &x);
    printf("input y:");
    scanf("%d", &y);
    printf("input r:");
    scanf("%d", &r);

    distance = (200-x)*(200-x)+(200-y)*(200-y);
    radius = (150+r)*(150+r);

    HgOpen(400, 400);
    
    HgCircle(200, 200, 150);//円1

    if(distance < radius){
        HgSetFillColor(HG_RED);
    }else if(distance == radius || distance > radius){
        HgSetFillColor(HG_BLUE);
    }
    HgCircleFill(x, y, r, 0);//円2

    HgGetChar();
    HgClose();
    return 0;
}

