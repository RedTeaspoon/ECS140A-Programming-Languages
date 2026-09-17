;; You may define helper functions here

;; check if sequence contains rune label, returns index of edge containing label
defun edges-contains-label (edges, label) int{
	for index, edge := range edges{
		if edge.label == label{
			return index
		}
	}
	return -1
}

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
               (g1 (car current-edge)) 
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
            ((check-path g2 (car paths) target) (car paths))
            (t (call-check-path g2 start target (cdr paths)))
         )
      )
   )
)

(defun check-path (g2 current-path target)

)

(defun find-sequence (g1 g2 start target k)
   (let ((g1StartEdges (g1 start)) (g1EndExists (g1 target)) (g2StartEdges (g2 start)) g2EndExists (g2 target))
      ;; start == target, g1Start exists, but g2Start doesn't exist
      (cond 
         ((and (equal start target) (not (null car g1StartEdges)) (null g2StartExists)) 
            (cons nil t))
         ;; k >= 1 && g1StartExists && g1EndExists --> return (cons S nil)
         ((and (zerop k) (not (null g1StartEdges)) (not (null g1EndExists)))
            (let* ((paths (find-target g1 (list start) (list target) k nil 0)))
               (cond ((null paths) nil)
                  ((check-path) nil)
               )
            )
         )
         (t nil)
      )
   )
)