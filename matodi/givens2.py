import math

def givens_rotation(a, b):

    r = math.sqrt(pow(a,2) + pow(b,2))  # Исправлено: добавлено возведение в квадрат
    if r == 0:
        return 0, 0
    c = a / r
    s = -b / r
    return c, s

def apply_givens_rotation(A, b, i, j):
    c, s = givens_rotation(A[i][j], A[j][j])
    # Применяем вращение к строкам матрицы A
    for k in range(len(A)):
        temp = c * A[i][k] + s * A[j][k]
        A[j][k] = -s * A[i][k] + c * A[j][k]
        A[i][k] = temp

    # Применяем вращение к вектору b
    temp = c * b[i] + s * b[j]
    b[j] = -s * b[i] + c * b[j]
    b[i] = temp

def solve_using_givens(A, b):
    n = len(A)
    for j in range(n - 1):
        for i in range(j + 1, n):
            if A[i][j] != 0:
                apply_givens_rotation(A, b, i, j)

    # Обратная подстановка
    x = [0] * n
    for i in range(n - 1, -1, -1):
        x[i] = b[i]
        for j in range(i + 1, n):
            x[i] -= A[i][j] * x[j]
        x[i] /= A[i][i]

    return x

A = [[1.03, 0.991],
     [0.991, 0.943]]

b = [2.57,2.49]


solution = solve_using_givens(A, b)

print("Решение системы уравнений:", solution)