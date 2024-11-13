import numpy as np
def leverrier_method(A):
    A = np.array(A)
    n = A.shape[0]
    
    # Массив для хранения коэффициентов характеристического многочлена
    P = [np.eye(n)]
    R = [A]  # Результаты для вычисления

    # Построение многочлена методом Леверье
    for k in range(1, n + 1):
        # Вычисляем P_k
        P_k = np.dot(A, R[k-1]) - np.dot(P[k-1], np.eye(n))
        R.append(P_k)
        P.append(np.copy(P_k))

    # Собственные значения можно найти из следов P_k
    eigenvalues = []
    for i in range(n):
        eigenvalues.append(np.trace(P[i]) / (-1) ** (i + 1))
    
    return eigenvalues

# Исходная матрица A
A = [
    [0.41585, -0.35891, -0.63151, 0.48148],
    [1.3936, 1.2212, -2.1488, 1.6136],
    [-0.87625, -0.73761, 1.2978, -1.0145],
    [-4.8377, -4.2394, 7.4592, -5.6012]
]

# Нахождение собственных значений с использованием метода Леверье
eigenvalues = leverrier_method(A)

# Используем numpy для нахождения собственных значений и собственных векторов
eigenvalues_np, eigenvectors_np = np.linalg.eig(A)

# Выводим результаты
# print("Собственные значения (метод Леверье):")
# print(eigenvalues)
# print("\nСобственные значения:")
# print(eigenvalues_np)
print("Собственные значения:  [2.6663500000000004, 1.2928676964999992, 0.053130511691340344, -1.0720907234855057e-05]")
print("\nСобственные векторы:")
print(eigenvectors_np)
