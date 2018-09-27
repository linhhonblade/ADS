#include <stdio.h>
#include <stdlib.h>

struct nat {
  int val;
};

struct nat* newN(int i){
  struct nat *n = (struct nat *)malloc(sizeof(struct nat));
  n->val = i;
  return n;
}

struct nat* addN(struct nat *n1, struct nat *n2){
  struct nat *n = (struct nat*)malloc(sizeof(struct nat));
  n->val = n1->val + n2->val;
  return n;
}
void display(struct nat* n){
  printf("%d\n",n->val );
}
int main(){
  struct nat *n3;
  struct nat *n1 = newN(13);
  struct nat *n2 = newN(14);
  n3 = addN(n1, n2);
  display(n3);
}
