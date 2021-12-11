;(use-syntax (ice-9 syncase))

(define ie
  (interaction-environment))

(define memoized-factorial
  (let ((known-results '((0 1) (1 1))))
    (lambda (x)
      (let* ((arg x)
             (res (assoc arg known-results)))
        (if res
            (cadr res)
            (let ((res (* x (memoized-factorial (- x 1)))))
              (set! known-results (cons (list arg res) known-results))
              res))))))

(define-syntax lazy-cons
  (syntax-rules ()
    ((lazy-cons a b)
     (cons a (delay b)))))

(define (lazy-car p)
  (car p))

(define (lazy-cdr p)
  (force (cdr p)))

(define (lazy-head xs k)
  (let loop((ys xs)
            (i k)
            (res '()))
    (if (> i 0)
        (loop (lazy-cdr ys) (- i 1) (cons (lazy-car ys) res))
        (reverse res))))

(define (lazy-ref xs k)
  (let loop((ys xs)
            (i k))
    (if (= i 0)
        (lazy-car ys)
        (loop (lazy-cdr ys) (- i 1)))))

(define (naturals start)
  (lazy-cons start (naturals (+ start 1))))

(define (make-lazy-factorial start)
  (lazy-cons (memoized-factorial start) (make-lazy-factorial (+ start 1))))

(define (lazy-factorial n)
  (lazy-ref (make-lazy-factorial 0) n))

(define-syntax define-struct
  (syntax-rules ()
    ((define-struct name (obj ...))
     (begin
       (eval (list 'define (string->symbol (string-append (symbol->string 'make-)
                                                          (symbol->string 'name)))
                   '(lambda (obj ...) (vector 'name (list (list 'obj obj) ...)))) ie)
       (eval (list 'define (string->symbol (string-append (symbol->string 'name)
                                                          (symbol->string '?)))
                   '(lambda (xs) (equal? (vector-ref xs 0) 'name))) ie)
       (eval (list 'define (string->symbol (string-append (symbol->string 'name)
                                                          (symbol->string '-)
                                                          (symbol->string 'obj)))
                   '(lambda (xs) (cadr (assoc 'obj (vector-ref xs 1))))) ie) ...
       (eval (list 'define (string->symbol (string-append (symbol->string 'set-)
                                                          (symbol->string 'name)
                                                          (symbol->string '-)
                                                          (symbol->string 'obj)
                                                          (symbol->string '!)))
                   '(lambda (xs x) (set-car! (cdr (assoc 'obj (vector-ref xs 1))) x))) ie) ...))))

(define-syntax define-data
  (syntax-rules ()
    ((define-data type ((name val ...)
                        ...))
     (begin (eval (list 'define 'name
                 '(lambda (val ...) (vector 'type 'name (list (list 'val val) ...)))) ie) ...
            (eval (list 'define (string->symbol (string-append (symbol->string 'type)
                                                               (symbol->string '?)))
                        '(lambda (xs) (equal? (vector-ref xs 0) 'type))) ie)
            (eval (list 'define (string->symbol (string-append (symbol->string 'name)
                                                               (symbol->string '?)))
                        '(lambda (xs) (equal? (vector-ref xs 1) 'name))) ie) ...))))

