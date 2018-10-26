#include <stdio.h>
#include <stdlib.h>

//each node is a ring
//data is the ring size
struct Node {
  int data;
  struct Node* next;
};

//each stack is a rod
//size is number of rings in that rod
//rod_num is the name of the rod
struct Stack {
  int size;
  struct Node* top;
  char rod_num;
};

//initialize a hand(temp0) and three rods with no rings
//initially I did not intend to set rod1, rod2 and rod3 as global variables
//However when I add a new funtion showRod(), I see it is easier to implement
struct Node*temp0 = NULL;
struct Stack rod1 = {0, NULL, 1};
struct Stack rod2 = {0, NULL, 2};
struct Stack rod3 = {0, NULL, 3};

//use hand to pick a ring
void pop(struct Stack* stack){
  temp0 = stack->top;
  stack->top = stack->top->next;
  stack->size--;
}

//use hand to put the ring to other rod
void push(struct Stack* stack){
  temp0->next = stack->top;
  stack->top = temp0;
  temp0 = NULL;
  stack->size++;
}

//add rings to a rod before start playing
void initStack(struct Stack* stack,int size){
  int i;
  if(size ==0)
    return;
  for(i = size; i>0; i--){
    struct Node*temp = (struct Node*)malloc(sizeof(struct Node));
    temp->data = i;
    temp->next = stack->top;
    stack->top = temp;
    stack->size++;
  }
  return;
}

//This function is only to give insight how everything is going
void showRod(struct Stack stack){
  struct Node* temp = stack.top;
  printf("Rod %d: ",stack.rod_num );
  if(temp == NULL)
    printf("nothing in this rod");
  while(temp!=NULL){
    printf("%d ",temp->data );
    temp = temp->next;
  }
  printf("\n");
}

//function to solve the tower of hanoi
void solve(int n, struct Stack* from_rod, struct Stack* to_rod, struct Stack* temp_rod){
	if(n==1){
    printf("move ring #%d from rod %d to rod %d\n", from_rod->top->data, from_rod->rod_num, to_rod->rod_num);
		pop(from_rod);
		push(to_rod);
    showRod(rod1);
    showRod(rod2);
    showRod(rod3);
		return;
	}
	solve(n-1, from_rod, temp_rod, to_rod);
	solve(1,from_rod, to_rod, temp_rod);
	solve(n-1, temp_rod, to_rod, from_rod);
}



int main(){
  int n;
  printf("Enter the number of rings: ");
  scanf("%d",&n);
  initStack(&rod1,n);
  showRod(rod1);
  showRod(rod2);
  showRod(rod3);
  solve(n, &rod1, &rod3, &rod2);
  printf("DONE!!!\n");
  return 0;


}
