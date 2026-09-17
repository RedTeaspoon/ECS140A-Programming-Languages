;; You may define helper functions here

;; graph g1, target, allowed depth ;; current-edge, nodes visited so far, depth so far
;; make sure to pass in current-edge & target as a list
(defun find-target (g1 current-edge target k visited depth)
  (cond
      ;; at k & reached target
      ((and (equal k depth) (equal target current-edge))
         (list visited)
      )
      ;; at k & not reached target
      ((and (equal k depth) (not (equal target (cdr current-edge))))
         nil
      )
      ;; otherwise for each edge in edges adjacent to current-edge's node apply find-target -> sum up results in list
      (t
         (apply 'append (mapcar 
               (lambda (edge)
                  (find-target g1 (cdr edge) target k (append visited (list (car edge))) (+ depth 1)))
               (apply g1 current-edge) 
            )
         )
      ) 

   )
)
;; for each path in paths check if it's a valid path in g2, returns first invalid path
(defun call-check-path (g2 start target paths)
   (cond 
      ((null paths) nil)
      (t
         (cond 
            ((check-path g2 start (car paths) target) (car paths))
            (t (call-check-path g2 start target (cdr paths)))
         )
      )
   )
)

(defun check-path (g2 current-node current-path target)
   (cond
      ;; at end of current path & reached target
      ((and (null (car current-path)) (equal target (apply g2 current-node)))
         nil
      )
      ;; at end of current path & not reached target
      ((and (null (car current-path)) (not (equal target (apply g2 current-node))))
         t
      )
      ((null (apply g2 current-node)) t)
      (t
         ;; should be only 1 edge with the same label as start of current-path -> should only end up with 1 possible T
         (or (mapcar
               (lambda (labeled-edge)
                  (cond 
                     ((equal (car labeled-edge) (car current-path))
                        (check-path g2 (cdr labeled-edge) (list (cdr current-path)) target))
                     (t (list t))                  
                  )
               )
               (apply g2 current-node)
            )
         )
      ) 
   )
)

(defun find-sequence (g1 g2 start target k)
   (let ((g1StartEdges (apply g1 (list start))) (g1EndExists (apply g1 (list target))) (g2StartEdges (apply g2 (list start))) (g2EndExists (apply g2 (list target))))
      ;; start == target, g1Start exists, but g2Start doesn't exist
      (cond 
         ((and (equal start target) (not (null g1StartEdges)) (null g2StartEdges)) 
            (cons nil t))
         ;; k >= 1 && g1StartExists && g1EndExists --> return (cons S nil)
         ((and (not (zerop k)) (not (null g1StartEdges)) (not (null g1EndExists)))
            ;; pass start & target as lists for comparison purposes
            (let ((paths (find-target g1 (list start) (list target) k nil 0))) 
               (let ((valid (call-check-path g2 (list start) (list target) paths)))
                  (cond 
                     ((null paths) nil)
                     ((not (null valid)) (cons valid t))
                     (t '(error))
                  )
               )
            )
         )
         (t nil)
      )
   )
)