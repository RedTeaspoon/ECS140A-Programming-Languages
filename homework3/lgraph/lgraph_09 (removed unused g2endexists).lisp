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
      ((null current-edge) nil)
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
      ((and (null (car current-path)) (equal target current-node))
         nil
      )
      ;; at end of current path & not reached target
      ((and (null (car current-path)) (not (equal target current-node)))
         t
      )
      ((null (apply g2 current-node)) t)
      (t
         (check-edges g2 (apply g2 current-node) current-path target)
      ) 
   )
)
;; return t if no edges have the current path 
(defun check-edges (g2 edges current-path target)
  (cond 
      ((null edges) t)
      (t
         (cond 
            ((equal (car (car edges)) (car current-path)) 
               (check-path g2 (cdr (car edges)) (cdr current-path) target))
            (t (check-edges g2 (cdr edges) current-path target))
         )
      )
   )
)

(defun find-sequence (g1 g2 start target k)
   (let ((g1StartEdges (apply g1 (list start))) (g1EndExists (apply g1 (list target))) (g2StartEdges (apply g2 (list start))))
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
                     (t nil)
                  )
               )
            )
         )
         (t nil)
      )
   )
)