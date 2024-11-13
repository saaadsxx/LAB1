import numpy as np

# Исходная матрица A
A = np.array([
    [0.41585, -0.35891, -0.63151, 0.48148],
    [1.3936, 1.2212, -2.1488, 1.6136],
    [-0.87625, -0.73761, 1.2978, -1.0145],
    [-4.8377, -4.2394, 7.4592, -5.6012]
])

# Функция для нахождения собственных значений методом Фадеева
def fadeev_method(A):
    n = A.shape[0]
    eigenvalues = []
    B = np.copy(A)
    for i in range(n):
        # Находим характеристический многочлен для матрицы B
        p = np.poly(B)
        eigenvalue = p[-2] / p[-1]  # Вычисляем собственное значение
        eigenvalues.append(eigenvalue)
        # Обновляем матрицу B для следующего шага
        B = B - eigenvalue * np.eye(n)
    return eigenvalues

# Функция для нахождения собственных векторов
def eigenvectors(A, eigenvalues):
    eigenvectors = []
    for eigenvalue in eigenvalues:
        # Решаем (A - λI)v = 0 для каждого собственного значения
        eigvec = np.linalg.solve(A - eigenvalue * np.eye(A.shape[0]), np.zeros(A.shape[0]))
        eigenvectors.append(eigvec)
    return eigenvectors

# Нахождение собственных значений методом Фадеева
eigvals = fadeev_method(A)
print("Собственные значения:", eigvals)

# Нахождение собственных векторов
eigvecs = eigenvectors(A, eigvals)
print("Собственные векторы:")
for vec in eigvecs:
    print(vec)
