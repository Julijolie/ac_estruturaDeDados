#include <stdio.h>

int main() {

    const int numeroSecreto = 45;
    int tentativa = 1;
    int chute;
    int ganhou = 0; //0 representa false

    printf("***********************************\n");
    printf("* Bem-vindo ao jogo de adivinhacao *\n");
    printf("***********************************\n");

    while (ganhou == 0) {
        printf("\n***********************************");
        printf("\nTentativa %d", tentativa);
        printf("\nSeu chute eh: ", &chute);
        scanf("%d", &chute);
        printf("\nSeu chute foi %d ", &chute);

        if (chute < 0) {
            printf("Voce nao pode chutar numeros negativos!\n");
            continue; //para nao considerar se é maior ou menor que numero reservado
        }
        else {
            int acertou = (chute == numeroSecreto);
            int maior = (chute > numeroSecreto);

            if (acertou) {
                printf("Parabens, voce acertou o numero secreto!\n");
                ganhou = 1;
            }
            else if (maior) {
                printf("O numero que voce escolheu eh MAIOR que o numero secreto.\n");
            }
            else {
                printf("O numero que voce escolheu eh MENOR que o numero secreto.\n");
            }
        }
        tentativa++; // Incrementa a tentativa a cada chute
    }
    printf("Fim de jogo!");
    printf("Voce acertou em %d tentativas", tentativa);
    return 0;
}
