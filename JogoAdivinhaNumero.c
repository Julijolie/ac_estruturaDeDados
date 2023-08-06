#include <stdio.h>
#include <stdlib.h>

int main() {

    const int numeroSecreto = 45;
    int tentativa = 1;
    int chute;
    int ganhou = 0; //0 representa false
    double pontos = 1000;

    printf("***********************************\n");
    printf("* Bem-vindo ao jogo de adivinhacao *\n");
    printf("***********************************\n");

    while (1) { //1 representa o loop infinito
        printf("\n***********************************");

        printf("\nTentativa %d", tentativa);

        printf("\nSeu chute eh: ", &chute);
        scanf("%d", &chute);

        printf("\nSeu chute foi %d ", chute);

        if (chute < 0) {
            printf("Voce nao pode chutar numeros negativos!\n");
            continue; //para nao considerar se é maior ou menor que numero reservado
        }
        else {
            int acertou = (chute == numeroSecreto);
            int maior = (chute > numeroSecreto);

            if (acertou) {
                printf("\nParabens, voce acertou o numero secreto!\n");
                break;  //quebra do loop
            }
            else if (maior) {
                printf("\nO numero que voce escolheu eh MAIOR que o numero secreto.\n");
            }
            else {
                printf("\nO numero que voce escolheu eh MENOR que o numero secreto.\n");
            }
        }
        tentativa++; // Incrementa a tentativa a cada chute

        double pontosPerdidos = abs (chute - numeroSecreto) / 2.0;
        pontos = pontos - pontosPerdidos;
    }
    printf("\nFim de jogo!");
    printf("Voce acertou em %d tentativas\n", tentativa);
    printf("Total de pontos perdidos: %.1f" , pontos);
    return 0;
}
