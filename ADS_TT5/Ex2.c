#include <stdio.h>
#include <stdlib.h>
#include <time.h>

struct Node {
  int num;
  struct Node* next;
};

//function to print a list
void Print(struct Node* head){
  while(head!=NULL){
    printf("%d ",head->num );
    head = head->next;
  }
  printf("\n");
  return;
}

//function to generate a list with random value
struct Node* makeList(){
  int n, i, high, low;
  struct Node* head = NULL;
  printf("How many nodes do you want? ");
  scanf("%d",&n );
  printf("Enter the smallest value you want: ");
  scanf("%d",&low );
  printf("Enter the largest value you want: ");
  scanf("%d",&high );
  for(i = 0; i < n; i++){
    struct Node* temp = (struct Node*)malloc(sizeof(struct Node));
    temp->num = low + rand() % (high-low+1);
    temp->next = head;
    head = temp;
  }
  return head;
}


void Divide(struct Node* head, struct Node** mid){
//find the mid point of the list
  if(head->next == NULL) return;
  struct Node* findTail = head;
  *mid = head;
  while(findTail->next != NULL && findTail->next->next != NULL){//=_=//
    (*mid) = (*mid)->next;
    findTail = findTail->next->next;
  }
//mid now is before midpoint
//devide the list
  findTail = *mid;
  *mid = (*mid)->next;
  findTail->next = NULL;
  return;
}

//function to add node to a list
//node is the node we want to add
//to node is the head node of the list that we add node to
void addN(struct Node** node, struct Node** tonode){
  if(*tonode == NULL) *tonode = *node;
  else {(*tonode)->next = *node; *tonode = *node;}
}

/*function to merge two sorted lists
  head1 and head1 is the head pointers of two sorted lists
  this function returns the pointer to the head of the new list
  which is merged from two initial lists*/
struct Node* merge(struct Node* head1, struct Node* head2){
  struct Node* result = NULL;
  if(head1->num < head2->num) result = head1;//if heah1 || head2 ==NULL
  else result = head2;
  struct Node* temp = NULL;
  while(1){
    if(head1 == NULL){
      while(head2 != NULL){
        addN(&head2, &temp);
        head2 = head2->next;
      }
      return result;
    }
    if(head2 == NULL){
      while(head1 != NULL){
        addN(&head1, &temp);
        head1 = head1->next;
      }
      return result;
    }
    if(head1->num < head2->num){
      addN(&head1, &temp);
      head1 = head1->next;
    }
    else{addN(&head2, &temp); head2 = head2->next;}
  }
}

void mergeSort(struct Node** head){
  struct Node* mid = NULL;
  Divide(*head, &mid); //find mid
  if(mid == NULL) return; //base case
  mergeSort(head);
  mergeSort(&mid);
  *head = merge(*head,mid); //merge two sorted lists
  return;
}


int main(){
  struct Node* head = NULL;
  srand(time(NULL));
  head = makeList();
  printf("Initial list:\n");
  Print(head);
  clock_t begin = clock();
  mergeSort(&head);
  clock_t end = clock();
  printf("Sorted list:\n");
  Print(head);
  printf("Time taken: %lfs \n",(double)(end-begin)/CLOCKS_PER_SEC );
  return 0;
}


/*
############# COMPLEXITY ###################
- The linked list implementation for merge sort
does not affect the time complexity of the
merge sort with array implenentation but there
is difference in space complexity.
- Because we don't need to copy data into any new
array, we just work on the intitial list so the
space complexity of this implementation is O(1)
instead of O(n)
- And the time complexity remains O(nlogn) for
both best and average cases
*/
