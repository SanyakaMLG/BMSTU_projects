(define (count x xs)
  (let loop((res 0) (xs xs))
    (if (null? xs)
        res
        (if (equal? x (car xs))
            (loop (+ res 1) (cdr xs))
            (loop res (cdr xs))))))

(define (delete pred? xs)
  (let loop((res (list)) (xs xs))
    (if (null? xs)
        res
        (if (pred? (car xs))
            (loop res (cdr xs))
            (loop (append res (list (car xs))) (cdr xs))))))

(define (iterate f x n)
  (let loop((res (list)) (x x) (n n))
    (if (= n 0)
        res
        (loop (append res (list x)) (f x) (- n 1)))))

(define (intersperse e xs)
  (let loop((res (list)) (xs xs))
    (if (null? xs)
        res
        (if (null? (cdr xs))
            (append res (list (car xs)))
            (loop (append res (list (car xs)) (list e)) (cdr xs))))))

(define (any? pred? xs)
  (let loop((res #f) (xs xs))
    (if (null? xs)
        res
        (loop (or res (pred? (car xs))) (cdr xs)))))

(define (all? pred? xs)
  (let loop((res #t) (xs xs))
    (if (null? xs)
        res
        (loop (and res (pred? (car xs))) (cdr xs)))))