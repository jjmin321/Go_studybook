/*#include <stdio.h>

int main() {
    node* tail = NULL;
    int input;

    while(1) {
        scanf_s("%d", &input);
        if (input==0)
        break;
        insert(&tail, input);
    }
    return 0;
}

void insert(node ** tail, int input){
    node* newNode;
    newNode = (node)*malloc(sizeof(newnode));
    newnode->data = input;
    newNode -> next = NULL;

    if(*tail == NULL){
        *tail =newNode;
        newNode -> next = newNode;
    }
    else{
        newNode -> next = (*tail)->next;
        (*tail)->next=newNode;
    }
}*/

#include <stdio.h>

int main() {
    node *tail = NULL;
    node2* head = (node2*)malloc(sizeof(node2));
    
}