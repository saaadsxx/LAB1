import numpy as np

def givens_rotation(A, b):
    """
    Решение СЛАУ методом вращения Гивенса
    :param A: Матрица коэффициентов
    :param b: Вектор правой части
    :return: Вектор решений x
    """
    # Размер матрицы A
    m, n = A.shape
    A = A.astype(float)
    b = b.astype(float)
    
    # Применение вращений Гивенса для приведения A к верхнетреугольной форме
    for j in range(n):
        for i in range(m-1, j, -1):
            # Вычисляем значения cos и sin для матрицы вращения
            a, b_ = A[i-1, j], A[i, j]
            r = np.hypot(a, b_)
            c, s = a / r, -b_ / r
            
            # Применение матрицы вращения Гивенса к матрице A
            G = np.eye(m)
            G[i-1, i-1] = c
            G[i, i] = c
            G[i-1, i] = s
            G[i, i-1] = -s
            
            A = G @ A
            b = G @ b
    
    # Решение верхнетреугольной системы
    x = np.zeros(n)
    for i in range(n-1, -1, -1):
        x[i] = (b[i] - np.dot(A[i, i+1:], x[i+1:])) / A[i, i]
    
    return x

# Входные данные
A = np.array([[1.03, 0.991], 
              [0.991, 0.943]])
b = np.array([2.57, 2.49])

# Решение системы методом Гивенса
x_solution = givens_rotation(A, b)
print("Решение x:", x_solution)
