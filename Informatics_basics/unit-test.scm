(define-syntax test
  (syntax-rules ()
    ((test func res)
     `(test func res))))

(define-syntax run-tests
  (syntax-rules ()
    ((run-tests xs)
     (let loop((ys xs)
               (res #t))
       (if (null? ys)
           res
           (begin
             (write (cadar ys))
             (let ((exp (caddar ys)) (ret (eval (cadar ys) (interaction-environment))))
               (if (equal? exp ret)
                   (begin
                     (display " ok")
                     (newline)
                     (loop (cdr ys) (and res #t)))
                   (begin
                     (display " FAIL")
                     (newline)
                     (display "  Expected: ")
                     (write exp)
                     (newline)
                     (display "  Returned: ")
                     (write ret)
                     (newline)
                     (loop (cdr ys) (and #f res)))))))))))
  
