(load "trace.scm")
(load "unit-test.scm")

(define (find-ind word program i)
  (if (equal? (vector-ref program i) word)
      i
      (find-ind word program (+ i 1))))

(define (interpret program stack)
  (let loop((res-stack stack)
            (i 0)
            (ret-stack '())
            (definitions '()))
    (if (>= i (vector-length program))
        res-stack
        (let ((symbol (vector-ref program i)))
          (cond ((equal? symbol '+) (loop (cons (+ (car res-stack) (cadr res-stack)) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol '-) (loop (cons (- (cadr res-stack) (car res-stack)) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol '*) (loop (cons (* (car res-stack) (cadr res-stack)) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol '/) (loop (cons (quotient (cadr res-stack) (car res-stack)) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol 'mod) (loop (cons (remainder (cadr res-stack) (car res-stack)) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol 'neg) (loop (cons (- (car res-stack)) (cdr res-stack)) (+ i 1) ret-stack definitions))
                ((number? symbol) (loop (cons symbol res-stack) (+ i 1) ret-stack definitions))
                ((equal? symbol '=) (loop (cons (if (= (car res-stack) (cadr res-stack))
                                                    -1
                                                    0) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol '>) (loop (cons (if (< (car res-stack) (cadr res-stack))
                                                    -1
                                                    0) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol '<) (loop (cons (if (> (car res-stack) (cadr res-stack))
                                                    -1
                                                    0) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol 'not) (loop (cons (if (= (car res-stack) 0)
                                                      -1
                                                      0) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol 'and) (loop (cons (if (or (= (car res-stack) 0) (= (cadr res-stack) 0))
                                                      0
                                                      -1) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol 'or) (loop (cons (if (and (= (car res-stack) 0) (= (cadr res-stack) 0))
                                                     0
                                                     -1) (cddr res-stack)) (+ i 1) ret-stack definitions))
                ((equal? symbol 'drop) (loop (cdr res-stack) (+ i 1) ret-stack definitions))
                ((equal? symbol 'swap) (loop (cons (cadr res-stack) (cons (car res-stack) (cddr res-stack))) (+ i 1) ret-stack definitions))
                ((equal? symbol 'dup) (loop (cons (car res-stack) res-stack) (+ i 1) ret-stack definitions))
                ((equal? symbol 'over) (loop (cons (cadr res-stack) res-stack) (+ i 1) ret-stack definitions))
                ((equal? symbol 'rot) (loop (cons (caddr res-stack) (cons (cadr res-stack) (cons (car res-stack) (cdddr res-stack))))
                                            (+ i 1) ret-stack definitions))
                ((equal? symbol 'depth) (loop (cons (vector-length (list->vector res-stack)) res-stack) (+ i 1) ret-stack definitions))
                ((equal? symbol 'if) (if (= (car res-stack) 0)
                                         (loop (cdr res-stack) (+ (find-ind 'endif program i) 1) ret-stack definitions)
                                         (loop (cdr res-stack) (+ i 1) ret-stack definitions)))
                ((equal? symbol 'endif) (loop res-stack (+ i 1) ret-stack definitions))
                ((equal? symbol 'define) (loop res-stack (+ (find-ind 'end program i) 1) ret-stack
                                               (cons (list (vector-ref program (+ i 1)) (+ i 2)) definitions)))
                ((or (equal? symbol 'end) (equal? symbol 'exit)) (loop res-stack (car ret-stack) (cdr ret-stack) definitions))
                (else (loop res-stack (cadr (assoc symbol definitions)) (cons (+ i 1) ret-stack) definitions)))))))
                
                                         


(define the-tests
  (list (test (interpret #(   define abs
                         dup 0 <
                         if neg endif
                         end
                         9 abs
                         -9 abs      ) (quote ())) '(9 9))
        (test (interpret #(   define =0? dup 0 = end
                define <0? dup 0 < end
                define signum
                    =0? if exit endif
                    <0? if drop -1 exit endif
                    drop
                    1
                end
                 0 signum
                -5 signum
                10 signum       ) (quote ())) '(1 -1 0))
        (test (interpret #(   define -- 1 - end
                         define =0? dup 0 = end
                         define =1? dup 1 = end
                         define factorial
                         =0? if drop 1 exit endif
                         =1? if drop 1 exit endif
                         dup --
                         factorial
                         *
                         end
                         0 factorial
                         1 factorial
                         2 factorial
                         3 factorial
                         4 factorial     ) (quote ())) '(24 6 2 1 1))
        (test (interpret #(   define =0? dup 0 = end
                define =1? dup 1 = end
                define -- 1 - end
                define fib
                    =0? if drop 0 exit endif
                    =1? if drop 1 exit endif
                    -- dup
                    -- fib
                    swap fib
                    +
                end
                define make-fib
                    dup 0 < if drop exit endif
                    dup fib
                    swap --
                    make-fib
                end
                10 make-fib     ) (quote ())) '(0 1 1 2 3 5 8 13 21 34 55))
        (test (interpret #(   define =0? dup 0 = end
                define gcd
                    =0? if drop exit endif
                    swap over mod
                    gcd
                end
                90 99 gcd
                234 8100 gcd    ) '()) '(18 9))))

(run-tests the-tests)
