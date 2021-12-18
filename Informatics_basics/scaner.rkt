(load "stream.scm")
(load "unit-test.scm")

; <fraction> ::= <nomenator><denominator>
; <nomenator> ::= <number with sign>
; <denominator> ::= /<number>
; <number with sign> ::= +<number> | -<number> | <number>
; <number> ::= DIGIT <number> | DIGIT

(define (scan str)
  (let* ((EOF (integer->char 0))
         (stream (make-stream (string->list str) EOF)))
    
    (call-with-current-continuation
     (lambda (error)
       (define result (tokens stream error))
       (and (equal? (peek stream) EOF)
            result)))))

(define (tokens stream error)
  (define (start-token? char)
    (or (char-numeric? char)
        (equal? char #\/)
        (equal? char #\+)
        (equal? char #\-)))
  (cond ((char-whitespace? (peek stream))
         (spaces stream error)
         (cons #\space (tokens stream error)))
        ((start-token? (peek stream))
         (cons (token stream error)
               (tokens stream error)))
        (else '())))

(define (spaces stream error)
  (cond ((char-whitespace? (peek stream))
         ;(if (char-whitespace? (peek stream))
         ;    (next stream)
         ;    (error #f))
         (next stream))
        (else #t)))

(define (token stream error)
  (cond ((equal? (peek stream) #\/) (next stream))
        ((equal? (peek stream) #\+) (next stream))
        ((equal? (peek stream) #\-) (next stream))
        ((char-numeric? (peek stream)) (next stream))
        (else (error #f))))

(define (check-frac str)
  (and (scan str)
       (let* ((EOF (integer->char 0))
              (stream (make-stream (scan str))))
  
         (call-with-current-continuation
          (lambda (error)
            (sequence stream error)
            (equal? (peek stream) #f))))))

(define (find-ind symbol xs)
  (let loop((ys xs)
            (i 0))
    (if (null? ys)
        #f
        (if (equal? symbol (car ys))
            i
            (loop (cdr ys) (+ i 1))))))

(define (count symbol xs)
  (let loop((ys xs)
            (res 0))
    (if (null? ys)
        res
        (if (equal? symbol (car ys))
            (loop (cdr ys) (+ res 1))
            (loop (cdr ys) res)))))

(define (check-frac str)
  (and (scan str)
       (and (not (null? (scan str)))
            (or (equal? (vector-ref (list->vector (scan str)) 0) #\+)
                (equal? (vector-ref (list->vector (scan str)) 0) #\-)
                (char-numeric? (vector-ref (list->vector (scan str)) 0)))
            (or (and (equal? (count #\+ (scan str)) 1)
                     (equal? (count #\- (scan str)) 0))
                (and (equal? (count #\+ (scan str)) 0)
                     (equal? (count #\- (scan str)) 1))
                (and (equal? (count #\+ (scan str)) 0)
                     (equal? (count #\- (scan str)) 0)))
            (equal? (count #\/ (scan str)) 1)
            (> (vector-length (list->vector (scan str))) (+ (find-ind #\/ (scan str)) 1)))))

(define (scan-frac str)
  (and (check-frac str)
       (string->number str)))

(define (scan-many-fracs str)
  (let loop((xs (scan str))
            (res '())
            (ress '())
            (k 0))
    (and xs
         (if (null? xs)
             (if (null? res)
                 (and (not (equal? ress '())) (reverse ress))
                 (if (scan-frac (list->string (reverse res)))
                     (reverse (cons (scan-frac (list->string (reverse res))) ress))
                     #f))
             (if (char-whitespace? (car xs))
                 (if (= k 0)
                     (loop (cdr xs) res ress 0)
                     (if (scan-frac (list->string (reverse res)))
                         (loop (cdr xs) '() (cons (scan-frac (list->string (reverse res))) ress) 0)
                         #f))
                 (loop (cdr xs) (cons (car xs) res) ress (+ k 1)))))))
                 
                 

(define the-tests
  (list (test (check-frac "110/111") #t)
        (test (check-frac "-4/3") #t)
        (test (check-frac "+5/10") #t)
        (test (check-frac "5.0/10") #f)
        (test (check-frac "FF/10") #f)
        (test (check-frac "-") #f)
        (test (check-frac "111/112/113") #f)
        (test (check-frac "+10/4-5") #f)
        (test (check-frac "1/") #f)
        (test (check-frac "/3") #f)
        (test (check-frac "") #f)
        (test (scan-frac "110/111") 110/111)
        (test (scan-frac "-4/3") -4/3)
        (test (scan-frac "+5/10") 1/2)
        (test (scan-frac "5.0/10") #f)
        (test (scan-frac "FF/10") #f)
        (test (scan-frac "-") #f)
        (test (scan-frac "111/112/113") #f)
        (test (scan-frac "+10/4-5") #f)
        (test (scan-frac "1/") #f)
        (test (scan-frac "/3") #f)
        (test (scan-frac "") #f)
        (test (scan-many-fracs "\t1/2 1/3\n\n10/8") (1/2 1/3 5/4))
        (test (scan-many-fracs "\t1/2 1/3\n\n2/-5") #f)))

(run-tests the-tests)
