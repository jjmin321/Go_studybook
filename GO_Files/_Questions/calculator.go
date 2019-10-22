int main() {
	char infix[100] = {0, };
	char postfix[100] = {0, };
	printf("계산식 입력(중위표기법) : ");
	scanf_s("%s", infix, sizeof(infix));

	printf("\n후위표기법으로 변환 : ");
	infixToPostfix(infix, postfix);
	printf("%s\n", postfix);

	printf("\n계산결과 : %.3f", eval(postfix));
}