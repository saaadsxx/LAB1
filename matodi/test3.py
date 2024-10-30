import numpy as np

def givens_rotation(A, b):
    # Преобразуем A и b к типу float для точных вычислений
    A = np.array(A, dtype=float)
    b = np.array(b, dtype=float)
    n = len(b)

    # Применение вращений Гивенса для приведения A к верхнетреугольной форме
    for j in range(n - 1):
        for i in range(j + 1, n):
            if A[i, j] != 0:
                r = np.hypot(A[j, j], A[i, j])
                c = A[j, j] / r
                s = -A[i, j] / r

                # Применение вращения Гивенса к строкам j и i
                G = np.eye(n)
                G[j, j], G[i, i] = c, c
                G[j, i], G[i, j] = s, -s

                # Модификация матрицы A и вектора b
                A = G @ A
                b = G @ b

    # Решение треугольной системы
    x = np.zeros(n)
    for i in range(n - 1, -1, -1):
        x[i] = (b[i] - np.dot(A[i, i + 1:], x[i + 1:])) / A[i, i]

    return x

# Входные данные
A = [[1.03, 0.991], [0.991, 0.943]]
b = [2.57, 2.49]

# Решение системы
x = givens_rotation(A, b)
print("givens:", x)
