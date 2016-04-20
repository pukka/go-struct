Go言語の型や構造体について
=====

このコードではgoでサポートされている独自の型定義と、  
構造体について書いてあります。
  
goは次のように独自に型を定義することができます。  
>type ID int  
>type TaskID int  

どちらも本質はint型であるものの、  
IDとTaskIDという独自の型を定義しています。  
これにより、例えば同じ型の引数を必要とあるメソッドが存在した際にも、  
型をそれぞれ指定しておくことでコンパイル時にエラーとして検知できます。  

goでは、javaのクラスのような構造体と呼ばれる機能を提供しています。  
構造体は次のように定義し、  
また次のようにアクセスします。  
`
    var task Task = Task{  
        Code:   1,  
        Detail: "enjoy go",  
        done:   true,  
    }  
  
    fmt.Println(task.Code)  
    fmt.Println(task.Detail)  
    fmt.Println(task.done)  
`
