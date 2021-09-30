# Quick start Project

Vamos montar esse projeto juntos durante o treinamento ou extra treinamento.

Objetivos:

* Montar um projeto do zero
* Entender o processo de adição de bibliotecas
* Estrutura de arquivos
* Montar sua primeira função com propósito definido por você (cálculos simples, manipulação de estruturas de dados, etc)
* Montar sua segunda função para efetuar algum outro cálculo ou print de mensagem
* Utilizar go routines nas duas funções


Estrutura de dados:

Para criar um *Struct* do C, é simples. Basta seguir a seguinte diretiva:

```
type Dados struct {
   Variavel1 string
   Variavel2 int64
   Variavel3 float64 
}
```

Caso deseje a estrutura seja aninhada (*Nested*), crie uma nova estrutura e adicione como um tipo de variável. O exemplo a seguir demostra esse processo:
```
type NestedData struct {
    NestedVariable1 string
    NestedTImestamp time.Time
}

type Dados struct {
    Variavel1 string
    Variavel2 int64
    Variavel3 float64 
    Variavel4 NestedData 
}
```