(define (assert cond?)
  (if (not cond?)
      (begin
        (display "FAILED: ")
        (write ))))

(define (1/x x)
  (assert (not (zero? x)))
  (/ 1 x))

(define (save-data data path)
  (with-output-to-file path
    (lambda ()
      (write data))))

(define (load-data path)
  (with-input-from-file path
    (lambda ()
      (read))))

(define (count x xs)
  (let loop((res 0) (xs xs))
    (if (null? xs)
        res
        (if (equal? x (car xs))
            (loop (+ res 1) (cdr xs))
            (loop res (cdr xs))))))

(define (count-strings path)
  (let ((data (load-data path)))
    (+ 1 (count #\newline (string->list data)) (count #\return (string->list data)))))

(define (trib n)
  (if (= n 0)
      0
      (if (= n 1)
          0
          (if (= n 2)
              1
              (+ (trib (- n 1)) (trib (- n 2)) (trib (- n 3)))))))

(define trib-memo
  (let ((known-results '((0 0) (1 0) (2 1))))
    (lambda (x)
      (let* ((arg x)
             (res (assoc arg known-results)))
        (if res
            (cadr res)
            (let ((res (+ (trib-memo (- x 1)) (trib-memo (- x 2)) (trib-memo (- x 3)))))
              (set! known-results (cons (list arg res) known-results))
              res))))))

(define-syntax my-if
  (syntax-rules ()
    ((my-if cond expr1 expr2)
     (force (or (and cond (delay expr1))
                (delay expr2))))))

(define-syntax my-let
  (syntax-rules ()
    ((my-let ((var expr)) action)
     ((lambda (var) action) expr))
    ((my-let ((var expr) . other-vars) action)
     (my-let other-vars ((lambda (var) action) expr)))))

(define-syntax my-let*
  (syntax-rules ()
    ((my-let* ((var expr)) action)
     ((lambda (var) action) expr))
    ((my-let* ((var expr) . other-vars) action)
     ((lambda (var) (my-let* other-vars action)) expr))))

(define-syntax when
  (syntax-rules ()
    ((when cond? expr)
     (and cond? expr))
    ((when cond? expr . other-exprs)
     (and cond? expr (begin . other-exprs)))))

(define-syntax unless
  (syntax-rules ()
    ((unless cond? expr)
     (and (not cond?) expr))
    ((unless cond? expr . other-exprs)
     (and (not cond?) expr (begin . other-exprs)))))

(define-syntax for
  (syntax-rules (in as)
    ((for x in xs . actions)
     (let loop((ys xs))
       (if (not (null? ys))
           (let ((x (car ys)))
             (begin
               (begin . actions)
               (loop (cdr ys)))))))
    ((for xs as x . actions)
     (for x in xs . actions))))

(define-syntax while
  (syntax-rules ()
    ((while cond? . actions)
     (letrec ((loop (lambda ()
                      (if cond?
                          (begin
                            (begin . actions)
                            (loop)))))) (loop)))))

(define-syntax repeat
  (syntax-rules (until)
    ((repeat (action . actions) until cond?)
     (begin
       (begin
         action
         (begin . actions))
       (letrec ((loop (lambda ()
                        (if (not cond?)
                            (begin
                              action
                              (begin . actions)
                              (loop)))))) (loop))))))

(define-syntax cout
  (syntax-rules (<<)
    ((cout << action)
     (if (equal? 'action 'endl)
         (newline)
         (display action)))
    ((cout << action . other-actions)
     (begin
       (if (equal? 'action 'endl)
           (newline)
           (display action))
       (cout . other-actions)))))
