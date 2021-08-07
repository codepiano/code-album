; 王垠根据 《The Little Schemer》用 scheme 实现的推导版本
; reinventing-the-ycombinator: https://www.slideshare.net/yinwang0/reinventing-the-ycombinator

; 原始代码
(define length
  (lambda (ls)
    (cond
      ((null? ls) 0)
      (else (add1 (length (cdr ls)))))))

; 不能用递归，所以需要把 define 换成 lambda 表达式
(lambda (length)
  (lambda (ls)
    (cond
      ((null? ls) 0)
      (else (add1 (length (cdr ls)))))))

; 复制现有表达式，将之作为参数传递给自己，以绑定 length 为自身
; 注意，此处改变了函数的性质，把自己当做参数传递给自己，把 length 提升成了一个高阶函数，函数体里面的 length 需要修改，见下一步
((lambda (length) ; 函数调用
   (lambda (ls)
     (cond
       ((null? ls) 0)
       (else (add1 (length (cdr ls)))))))
 (lambda (length) ; 将自己作为参数传递给自己
   (lambda (ls)
     (cond
       ((null? ls) 0)
       (else (add1 (length (cdr ls))))))))

; length 被提升为高阶函数，修复
; 修复后的版本已经可以正常递归调用，递归 pattern 已经显露，现在需要提取 pattern，让其适合所有满足 "shape" 的函数
; poor man's Y
((lambda (length) ; 函数调用
   (lambda (ls)
     (cond
       ((null? ls) 0)
       (else (add1 ((length length) (cdr ls)))))))
 (lambda (length) ; 将自己作为参数传递给自己
   (lambda (ls)
     (cond
       ((null? ls) 0)
       (else (add1 ((length length) (cdr ls))))))))

; 需要消除 (length length) 结构
; 上面的代码中有三次自己作为参数调用自身，两次在内部，一次在外部(传递参数），要把这三次 pattern 都提取出来
; pattern: (lambda (u) (u u))
; 用编译原理的术语叫： common subexpression elimination
((lambda (u) (u u)) ;提取后的 pattern
 (lambda (length) ; 将自己作为参数传递给自己
   (lambda (ls)
     (cond
       ((null? ls) 0)
       (else (add1 ((length length) (cdr ls))))))))

; 现在在代码中只剩下一个 (length length) 结构了
; 通过再次的提升抽象等级（高阶函数），把这个结构抽到外层
((lambda (u) (u u)) ;提取后的 pattern
 (lambda (length) ; 将自己作为参数传递给自己
   ((lambda (g) ; 这个函数已经和上面第二步的 length lambda 表达式的定义一致了
      (lambda (ls)
        (cond
          ((null? ls) 0)
          (else (add1 (g (cdr ls)))))))
    (length length)))) ; (length length) 结构在 call-by-value 的语言中会造成无限循环，需要改进

; 改进 call-by-value 的问题
((lambda (u) (u u)) ;提取后的 pattern
 (lambda (length) ; 将自己作为参数传递给自己
   ((lambda (g) ; 这个函数已经和上面第二步的 length lambda 表达式的定义一致了
      (lambda (ls)
        (cond
          ((null? ls) 0)
          (else (add1 (g (cdr ls)))))))
    (lambda (v) (length length) v)))) ; (length length) 结构在 call-by-value 的语言中会造成无限循环，需要改进

; 再次提升，把 length lambda 结构提取到外层，通过参数传递 f 进去
((lambda (f)
   ((lambda (u) (u u)) ;提取后的 pattern
    (lambda (length) ; 将自己作为参数传递给自己
      (f
        (lambda (v) (length length) v)))))
 (lambda (g) ; 这个函数已经和上面第二步的 length lambda 表达式的定义一致了
   (lambda (ls)
     (cond
       ((null? ls) 0)
       (else (add1 (g (cdr ls)))))))) ; (length length) 结构在 call-by-value 的语言中会造成无限循环，需要改进

; 我们把 length 计算的结构提取出来，通过参数 f 传递之后，剩余的结构可以适用于任何函数了，只要当做 f 传递进去即可
(lambda (f)
  ((lambda (u) (u u)) ;提取后的 pattern
   (lambda (x) ; 将自己作为参数传递给自己
     (f (lambda (v) (x x) v)))))

(lambda (g) ; 这个函数已经和上面第二步的 length lambda 表达式的定义一致了
  (lambda (ls)
    (cond
      ((null? ls) 0)
      (else (add1 (g (cdr ls)))))))

