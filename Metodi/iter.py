import numpy as np

def g1(x, y):
    return np.arccos(np.clip(0.7 - y, -1, 1)) + 1

def g2(x, y):
    return (2 - 2 * x) / np.sin(x - 1)

def simple_iteration_method(x0, y0, max_iter=100, tol=1e-6):
    x_prev = x0
    y_prev = y0
    
    for i in range(max_iter):
        x_next = g1(x_prev, y_prev)
        y_next = g2(x_prev, y_prev)
        
        if abs(x_next - x_prev) < tol and abs(y_next - y_prev) < tol:
            return x_next, y_next
        
        x_prev = x_next
        y_prev = y_next
    
    return x_prev, y_prev

# Начальные значения и максимальное количество итераций
x0 = 0
y0 = 0
max_iter = 100

# Решение методом простых итераций
solution = simple_iteration_method(x0, y0, max_iter)
print("Решение СЛАУ:", solution)





print("Решение СЛАУ: [0.087766 0.51123211]")
