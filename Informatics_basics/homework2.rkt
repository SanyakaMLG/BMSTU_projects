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

(define (string-trim-left a)
  (let loop ((a (string->list a)))
    (if (or (null? a)
            (not (char-whitespace? (car a))))
        (list->string a)
        (loop (cdr a)))))

(define (string-trim-right a)
  (list->string (reverse (string->list (string-trim-left (list->string (reverse (string->list a))))))))

(define (string-trim a)
  (string-trim-left (string-trim-right a)))

(define (string-prefix? a b)
  (let loop((a (reverse (string->list a)))
            (b (reverse (string->list b))))
    (and (not (null? b))
         (or (equal? a b)
             (loop a (cdr b))))))

(define (string-suffix? a b)
  (let loop((a (string->list a))
            (b (string->list b)))
    (and (not (null? b))
         (or (equal? a b)
             (loop a (cdr b))))))

(define (string-infix? a b)
  (and (not (null? (string->list b)))
       (or (string-suffix? a b)
           (or (string-prefix? a b)
               (string-infix? a (list->string (cdr (string->list b))))))))

(define (string-split str sep)
  (let loop((str str)
            (res '())
            (ress '()))
    (if (null? (string->list str))
        (reverse (cons (list->string (reverse ress)) res))
        (if (string-prefix? sep str)
            (loop (substring str (string-length sep)) (cons (list->string (reverse ress)) res) '())
            (loop (substring str 1) res (cons (car (string->list str)) ress))))))

(string-trim-left  "\t\tabc def")
(string-trim-right "abc def\t")
(string-trim       "\t abc def \n")

(string-prefix? "abc" "abcdef")
(string-prefix? "bcd" "abcdef")
(string-prefix? "abcdef" "abc")

(string-suffix? "def" "abcdef")
(string-suffix? "bcd" "abcdef")

(string-infix? "def" "abcdefgh")
(string-infix? "abc" "abcdefgh")
(string-infix? "fgh" "abcdefgh")
(string-infix? "ijk" "abcdefgh")
(string-infix? "bcd" "abc")

(string-split "x;y;z" ";")
(string-split "x-->y-->z" "-->")

(define (size sizes)
  (let loop((res 1)
            (sizes sizes))
    (if (null? sizes)
        res
        (loop (* res (car sizes)) (cdr sizes)))))

(define (number indices multi-vector)
  (let loop((res 1)
            (ress 0)
            (indices indices)
            (multi-vector (vector-ref multi-vector 0)))
    (if (null? indices)
        ress
        (loop (* res (car multi-vector)) (+ ress (* res (car indices))) (cdr indices) (cdr multi-vector)))))

(define (make-multi-vector . xs)
  (if (null? (cdr xs))
      (make-multi-vector1 (car xs))
      (make-multi-vector2 (car xs) (cadr xs))))

(define (make-multi-vector1 sizes)
  (define tmp (make-vector 2 sizes))
  (vector-set! tmp 1 (make-vector (size sizes)))
  tmp)

(define (make-multi-vector2 sizes fill)
  (define tmp (make-vector 2 sizes))
  (vector-set! tmp 1 (make-vector (size sizes) fill))
  tmp)

(define (multi-vector? multi-vec)
  (and (vector? multi-vec)
       (= (vector-length multi-vec) 2)
       (list? (vector-ref multi-vec 0))
       (vector? (vector-ref multi-vec 1))))

(define (multi-vector-set! multi-vec indices el)
  (vector-set! (vector-ref multi-vec 1) (number indices multi-vec) el))

(define (multi-vector-ref multi-vec indices)
  (vector-ref (vector-ref multi-vec 1) (number indices multi-vec)))

(define (o . xs)
  (let loop((xs xs))
    (if (null? xs)
        (lambda (x) x)
        (lambda (x) ((car xs) ((loop (cdr xs)) x))))))
