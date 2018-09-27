#include <stdio.h>
#include <stdlib.h>
#include <math.h>

struct Node {
  int coeff;
  int degr;
  struct Node* next;
};

int size = 0;   // size of linked list
struct Node* head;

void addTerm(int coeff, int degr){
  struct Node* temp1 = head;
  struct Node* temp2 = NULL;
  struct Node* temp = (struct Node*)malloc(sizeof(struct Node));
  temp->coeff = coeff;
  temp->degr = degr;

///////////// add first Node or Nodes with smallest degree to the beginning of the function //////////////////
  if((temp1 == NULL) || (temp1->degr > degr) && temp2 == NULL){
    temp->next = temp1;
    head = temp;
    size++;
    return;
  }

///////////////////////////////////
  do{
    if(temp1->degr == degr){
      temp1->coeff += coeff; // add to existing Node
      return;
    }
    else if((temp1->degr > degr)&&(temp2->degr < degr)){
      temp->next = temp1;
      temp2->next = temp;   // add new Node between existing Nodes
      size++;
      return;
    }
    else{
      temp2 = temp1;
      temp1 = temp1->next;    // continue to find the right position to add new Node
    }
  } while (temp1 != NULL);

////////////// add Node with highest degree to the end of the function /////////////
  temp->next = temp1;
  temp2->next = temp;
  size++;
  return;
}

//////////////// Display the function ////////////////////////
void Display(){
  struct Node* temp = head;
  printf("%dx^(%d) ",temp->coeff, temp->degr);
  temp = temp->next;
  while(temp!=NULL){
    printf("+ %dx^(%d) ",temp->coeff, temp->degr);
    temp = temp->next;
  }
  printf("\n");
}

/////////////// function to remove a term /////////////////////
void rmTerm(int degr){
  struct Node* temp1 = head;
  struct Node* temp2 = NULL;
  if (temp1 == NULL) {
    printf("There is nothing to remove\n");
    return;
  }
  while(temp1!=NULL){
    if(temp1->degr == degr){
      if(temp2 == NULL){    // term that needs to be removed is the first term
        head = temp1->next;
        free(temp1);
        size--;
        return;
      }
      else{
        temp2->next = temp1->next;
        free(temp1);
        size--;
        return;
      }
    }
    temp2 = temp1;
    temp1 = temp1->next;
  }
  printf("There is no term with degree %d. Fuction remains the same.\n", degr );
}


int calFunc(int x){
  struct Node* temp = head;
  int result = 0;
  if (temp == NULL){
    return 0;
  }
  while(temp != NULL){
    result += temp->coeff*pow(x, temp->degr);
    temp = temp->next;
  }
  return result;
}

int main(){
  head = NULL;
  addTerm(3,4);
  addTerm(2,0);
  addTerm(3,5);
  printf("result = %d\n",calFunc(3));
  Display();
  rmTerm(0);
  printf("result = %d\n",calFunc(3));
  printf("size = %d\n",size );
}
