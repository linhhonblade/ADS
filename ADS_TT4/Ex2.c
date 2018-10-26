#include <stdio.h>
#include <stdlib.h>

struct Node {
  double data;
  struct Node* next;
};

struct Node* head;

double sumLink(struct Node* head){
    if(head->next == NULL)
      return head->data;
    return head->data+sumLiny(head->next);
}
void addN(){
  double n;
  printf("enter a number: ");
  scanf("%lf",&n );
  //create a node to store data
  struct Node* temp = (struct Node*)malloc(sizeof(struct Node));
  temp->data = n;
  temp->next = head;
  head = temp;
  //ask if you want to add more node
  char c;
  printf("Do you want to add more? y/n?");
  scanf("\n%c",&c );
  if(c=='y'){
    addN();
  }
  else
    return;
}

int main(){
  head = NULL;
  addN();
  printf("sum of all elements = %.2lf\n",sumLink(head));
  return 0;
}
