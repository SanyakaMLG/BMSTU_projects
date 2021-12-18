(load "unit-test.scm")

;<Program>  ::= <Articles> <Body> .
;<Articles> ::= <Article> <Articles> | .
;<Article>  ::= define word <Body> end .
;<Body>     ::= if <Body> endif <Body> | integer <Body> | word <Body> | .

(define force-return 0)

(define (exit)
  (force-return #f))

(define (find-tail xs target)
  (if (equal? (car xs) target)
      (cdr xs)
      (find-tail (cdr xs) target)))

(define (find-head xs target)
  (let loop ((xs xs)
             (res '()))
    (and (not (null? xs))
         (if (equal? (car xs) target)
             (reverse res)
             (loop (cdr xs) (cons (car xs) res))))))

(define (last xs)
  (list-ref xs (- (length xs) 1)))

(define (tail-endif program)
  (let loop ((program program)
             (balance -1))
    
    (and (not (null? program))
         (let ((word (car program)))
           (cond
             ((and (equal? word 'endif) (= balance 0)) (cdr program))
             ((equal? word 'endif) (loop (cdr program) (- balance 1)))
             ((equal? word 'if) (loop (cdr program) (+ balance 1)))
             (else (loop (cdr program) balance)))))))

(define (head xs n)
  (if (or (= n -1) (null? xs))
      '()
      (cons (car xs) (head (cdr xs) (- n 1)))))

(define (parse program)
  (call-with-current-continuation
   (lambda (stack)
     (set! force-return stack)
     (let ((program (vector->list program)))
       (if (equal? (car program) 'define)
           (let ((articles (parse-articles program)))
             (cons (head articles (- (length articles) 2))
                   (list (parse-body (last articles)))))
           (cons '() (list (parse-body program))))))))


(define (parse-articles program)
  (let loop ((program program))
    (if (not (null? program))
        (let ((word (car program))
              (other (cdr program)))
          (if (equal? word 'define)
              (if (null? other) (exit)
                  (if (or (equal? (car other) 'if)
                          (equal? (car other) 'endif))
                      (exit)
                      (let ((head (find-head (cdr other) 'end)))
                        (if (not head) (exit)
                            (cons (cons (car other) (list (parse-body head)))
                                  (loop (find-tail (cdr other) 'end)))))))
              (list program)))
        (list program))))
            
(define (parse-body program)
  (let loop ((program program) (parsed '()) (stack '()))
    (if (not (null? program))
        (let ((word (car program)))
          (cond
            ((equal? word 'if)
             (let ((tail (tail-endif program)))
               (if (not tail)
                   (exit)
                   (loop tail (cons (list 'if (loop (cdr program) '() (cons 'if stack))) parsed) stack))))
            ((equal? word 'endif)
             (if (and (not (null? stack)) (equal? (car stack) 'if))
                 (reverse parsed)
                 (exit)))
            ((or (equal? word 'define) (equal? word 'end)) (exit))
            (else (loop (cdr program) (cons word parsed) stack))))
        (reverse parsed))))

(define the-tests
  (list (test (parse #(1 2 +)) (() (1 2 +)))
        (test (parse #(x dup 0 swap if drop -1 endif)) (() (x dup 0 swap (if (drop -1)))))
        (test (parse #( define -- 1 - end
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
                         4 factorial ))

              (((-- (1 -))
                (=0? (dup 0 =))
                (=1? (dup 1 =))
                (factorial
                 (=0? (if (drop 1 exit)) =1? (if (drop 1 exit)) dup -- factorial *)))
               (0 factorial 1 factorial 2 factorial 3 factorial 4 factorial)))

        (test (parse #(define word w1 w2 w3)) #f)))

(run-tests the-tests)