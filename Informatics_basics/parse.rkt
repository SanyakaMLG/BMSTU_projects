(load "stream.scm")
(load "unit-test.scm")

(define finish-symbol (integer->char 0))
(define forbidden-words (map list (list finish-symbol 'endif 'end)))

;<Program>  ::= <Articles> <Body> .
;<Articles> ::= <Article> <Articles> | .
;<Article>  ::= define word <Body> end .
;<Body>     ::= if <Body> endif <Body> | integer <Body> | word <Body> | .

(define (parse program)
  (define stream (make-stream (vector->list program) finish-symbol))
  (parse-program stream))

(define parse-finish-symbol
  (lambda (stream)
    (and (equal? (peek stream) finish-symbol)
         (next stream))))

(define parse-program
  (lambda (stream)
    (let ((articles (parse-articles stream))
          (body (parse-body stream))
          (finish (parse-finish-symbol stream)))
      (and (and articles body finish)
           (list articles body)))))

(define parse-articles
  (lambda (stream)
    (let* ((article (parse-article stream))
           (articles (and (not (equal? article '()))
                          article
                          (parse-articles stream))))
      (if (and article articles (not (equal? article '())))
          (cons article articles)
          article))))

(define parse-define
  (lambda (stream)
    (and (equal? (peek stream) 'define)
         (next stream))))

(define parse-end
  (lambda (stream)
    (and (equal? (peek stream) 'end)
         (next stream))))

(define parse-word
  (lambda (stream)
    (next stream)))

(define parse-article
  (lambda (stream)
    (let* ((def (parse-define stream))
           (word (and def (parse-word stream)))
           (body (and word (parse-body stream)))
           (end (and body (parse-end stream))))
      (if def
          (and def word body end
               (list word body))
          '()))))

(define parse-endif
  (lambda (stream)
    (and (equal? (peek stream) 'endif)
         (next stream))))

(define parse-body
  (lambda (stream)
    (let ((symbol (peek stream)))
      (cond ((equal? symbol 'if)
             (next stream)
             (let* ((body (parse-body stream))
                    (endif (and body (parse-endif stream)))
                    (body-next (and endif (parse-body stream))))
               (and symbol body endif body-next
                    (append (list (append (list symbol body))) body-next))))
            ((number? symbol)
             (next stream)
             (let ((body (parse-body stream)))
               (and symbol body
                    (cons symbol body))))
            ((not (assoc (peek stream) forbidden-words))
             (next stream)
             (let ((body (parse-body stream)))
               (and symbol body
                    (cons symbol body))))
            (else '())))))

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
                         4 factorial )) (((-- (1 -))
                                           (=0? (dup 0 =))
                                           (=1? (dup 1 =))
                                           (factorial
                                            (=0? (if (drop 1 exit)) =1? (if (drop 1 exit)) dup -- factorial *)))
                                          (0 factorial 1 factorial 2 factorial 3 factorial 4 factorial)))
        (test (parse #(define word w1 w2 w3)) #f)))

(run-tests the-tests)        
