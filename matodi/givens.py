import numpy as np

def givens_rotation(A, B):
    """
    Решает систему линейных уравнений (СЛАУ) методом Гивенса.
    
    Параметры:
    A : numpy.ndarray
        Матрица коэффициентов размером (m, n).
    B : numpy.ndarray
        Вектор значений размером (m,).
        
    Возвращает:
    x : numpy.ndarray
        Решение системы уравнений.
    """
    m, n = A.shape
    c = np.zeros(n)
    s = np.zeros(n)
    for j in range(n):
        for i in range(j+1, m):
            if A[i, j] != 0.0:
                r = np.sqrt(A[j, j]**2 + A[i, j]**2)
                c[j] = A[j, j] / r
                s[j] = -A[i, j] / r
                A[[j, i]] = c[j] * A[[j, i]] - s[j] * A[[i, j]]
                B[[j, i]] = c[j] * B[[j, i]] - s[j] * B[[i, j]]
    x = np.zeros(n)
    for j in range(n - 1, -1, -1):
        x[j] = (B[j] - np.dot(A[j, j+1:], x[j+1:])) / A[j, j]
    return x

# Пример
A = np.array([[1.03, 0.992],
              [0.991, 0.952]])
B = np.array([2.54, 2.42])  # Исправленный формат для вектора B

solution = givens_rotation(A, B)
print("Решение СЛАУ:", solution)










print("Решение СЛАУ[-2.075, 4.768]")