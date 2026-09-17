; You may define helper functions here

(defun match (pattern assertion)  
  (cond 
    ;; end of both lists -> T
    ((and (null (car pattern)) (null (car assertion))) t)
    ;; single match
    ((or (equal (car pattern) (car assertion)) (and (equal (car pattern) '?) (not (null (car assertion)))))
      (match (cdr pattern) (cdr assertion))
    )
    ;; pattern '! match only 1
    ((and (equal (car pattern) '!)
        (or (equal (car (cdr pattern)) (car (cdr assertion))) (and (equal (car (cdr pattern)) '?) (not (null (car (cdr assertion))))) 
        (and (null (car (cdr pattern))) (null (car (cdr assertion))))))
      (match (cdr pattern) (cdr assertion))
    )
    ;; pattern '! match possibly multiple
    ((and (equal (car pattern) '!) (not (null (car assertion))))
      (match pattern (cdr assertion))
    )
    (t nil)
  )
)

