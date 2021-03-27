package main

import (
  "fmt"
	"io"
)

type Vaga struct {
  tamanho int
  inicio int
  fim int
}

func criarVaga() Vaga {
  vaga := Vaga{}
  vaga.inicio =  -1
  vaga.fim = -1
  vaga.tamanho = 0
  return vaga
}

func maiorVaga(estacionamento []int, tamanho int, comparador int) Vaga{	
    vaga := criarVaga();
    for i:= 0; i < tamanho; i++ {
       if estacionamento[i] == 0 {
         if vaga.inicio == -1  {
            vaga.inicio = i;
         }
         vaga.fim = i;
         vaga.tamanho++;
         if vaga.tamanho == comparador {
            return vaga
         }
       }
    }
    return vaga
}

func colocar(estacionamento []int, tamanho int, placa int, vaga Vaga) []int {
   for i:= 0; i <= vaga.fim; i++ {
     estacionamento[i] = placa
   }
   return estacionamento
}

func retirar(estacionamento []int, C int, placa int) []int {
    for i:= 0; i < C; i++ {
      if estacionamento[i] == placa {
         estacionamento[i] = 0;
      }
    }
   return estacionamento
}

func main() {
	var N,C,i,comparador,faturamento,placa int
	var vaga Vaga 
	var e string 
	var estacionamento []int
	for {
		fmt.Printf("for")
		_,err := fmt.Scanf("%d %d", &C, &N)
		if err == io.EOF {
			break
		}
		for i = 0; i < C; i++ {
			estacionamento = append(estacionamento, 0)
		}
	     	for i = 0; i < N; i++ { 
			fmt.Scanf("%s %d", &e, &placa)
			if e == "C" {
				fmt.Scanf("%d", &comparador)
				vaga = maiorVaga(estacionamento, C, comparador)
				if vaga.tamanho >= comparador {
					faturamento += 10
					estacionamento = colocar(estacionamento,C,placa, vaga)
				}
			} else {
				estacionamento = retirar(estacionamento,C,placa)
			}
		}
	}
	fmt.Printf("%d\n",faturamento)
}


