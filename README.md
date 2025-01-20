# Введение
Я прочитал статью и понял, что она недостаточно полна в части, касающейся изменения емкости срезов Golang в зависимости от типов элементов среза.
Ссылка на статью: [habr.com](https://habr.com/ru/articles/874084/) 

# Замысел
Создаем массивы разных типов и фиксируем их емкость.
В итоге получаем такую картину:
- размерность 512
  
![размерность 512](https://github.com/dreddsa5dies/go_capasity_test/blob/main/capacity_512.png)

- размерность 1024
  
![размерность 1024](https://github.com/dreddsa5dies/go_capasity_test/blob/main/capacity_1024.png)

- и далее

![и далее](https://github.com/dreddsa5dies/go_capasity_test/blob/main/capacity_1536.png)

# Исходный код slice.go
Код обработчика находится тут: [slice.go](https://github.com/golang/go/blob/master/src/runtime/slice.go)

# Introduction
I read the article and realized that it is not complete enough in terms of changing the capacity of Golang slices depending on the types of slice elements.
Link to the article: [habr.com ](https://habr.com/ru/articles/874084/) 

# The idea
We create arrays of different types and fix their capacity.
As a result, we get this picture: see before.

# slice.go source code
The handler code is located here: [slice.go](https://github.com/golang/go/blob/master/src/runtime/slice.go)
