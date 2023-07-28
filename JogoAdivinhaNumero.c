#include <stdio.h>

int main() {
    const int NUMERO_DE_TENTATIVAS = 5;
    const int numeroSecreto = 45;
    int i;
    int chute;

    printf("***********************************\n");
    printf("* Bem-vindo ao jogo de adivinhacao *\n");
    printf("***********************************\n");

    for (i = 1; i <= NUMERO_DE_TENTATIVAS; i++) {
        printf("\n***********************************");
        printf("\nTentativa %d/%d", i, NUMERO_DE_TENTATIVAS);
        printf("\nSeu chute eh: ");
        scanf("%d", &chute);

        if (chute < 0) {
            printf("Voce nao pode chutar numeros negativos!\n");
            i--; // Reduz a tentativa atual para permitir uma nova tentativa válida
            continue; //para nao considerar se é maior ou menor que numero reservado
        } else {
            int acertou = (chute == numeroSecreto);
            int maior = (chute > numeroSecreto);
            int menor = (chute < numeroSecreto);

            if (acertou) {
                printf("Parabens, voce acertou o numero secreto!\n");
                break;
            } else if (maior) {
                printf("O numero que voce escolheu eh MAIOR que o numero secreto.\n");
            } else {
                printf("O numero que voce escolheu eh MENOR que o numero secreto.\n");
            }
        }
    }

    if (i > NUMERO_DE_TENTATIVAS) {
        printf("\nSuas tentativas acabaram. O numero secreto era %d.\n", numeroSecreto);
    }

    return 0;
}
