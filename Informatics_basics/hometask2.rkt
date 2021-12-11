(define (my-range a b d)
  (if (>= a b)
      '()
      (cons a (my-range (+ a d) b d))))

(define (my-flatten xs)
  (if (null? xs)
      '()
      (if (list? (car xs))
          (append (my-flatten (car xs)) (my-flatten (cdr xs)))
          (cons (car xs) (my-flatten (cdr xs))))))

(define (my-element? x xs)
  (let loop((res #f) (xs xs))
    (if (null? xs)
        res
        (loop (or res (equal? (car xs) x)) (cdr xs)))))

(define (my-filter pred? xs)
  (if (null? xs)
      '()
      (if (pred? (car xs))
          (cons (car xs) (my-filter pred? (cdr xs)))
          (my-filter pred? (cdr xs)))))

(define (my-fold-left op xs)
  (if (null? (cdr xs))
      (car xs)
      (my-fold-left op (cons (op (car xs) (cadr xs)) (cddr xs)))))

(define (my-fold-right op xs)
  (if (null? (cdr xs))
      (car xs)
      (op (car xs) (my-fold-right op (cdr xs)))))

(define (list->set xs)
  (if (null? xs)
      '()
      (if (my-element? (car xs) (cdr xs))
          (list->set (cdr xs))
          (cons (car xs) (list->set (cdr xs))))))

(define (set? xs)
  (or (null? xs)
      (and (not (my-element? (car xs) (cdr xs)))
           (set? (cdr xs)))))

(define (union xs ys)
  (if (null? xs)
      ys
      (if (my-element? (car xs) ys)
          (union (cdr xs) ys)
          (union (cdr xs) (cons (car xs) ys)))))

(define (intersection xs ys)
  (if (null? xs)
      '()
      (if (my-element? (car xs) ys)
          (cons (car xs) (intersection (cdr xs) ys))
          (intersection (cdr xs) ys))))

(define (difference xs ys)
  (if (null? xs)
      '()
      (if (my-element? (car xs) ys)
          (difference (cdr xs) ys)
          (cons (car xs) (difference (cdr xs) ys)))))

(define (symmetric-difference xs ys)
  (union (difference xs ys) (difference ys xs)))

(define (set-eq? xs ys)
  (null? (symmetric-difference xs ys)))

(define (f x) (+ x 2))
(define (g x) (* x 3))
(define (h x) (- x))

(define (o . xs)
  (let loop((xs xs))
    (if (null? xs)
        (lambda (x) x)
        (lambda (x) ((car xs) ((loop (cdr xs)) x))))))