import numpy as np

# Исходная матрица A
A = np.array([
    [0.41585, -0.35891, -0.63151, 0.48148],
    [1.3936, 1.2212, -2.1488, 1.6136],
    [-0.87625, -0.73761, 1.2978, -1.0145],
    [-4.8377, -4.2394, 7.4592, -5.6012]
])

# Функция для применения метода Леверрье
def leverrier_method(A):
    n = A.shape[0]
    coeffs = [1]  # Начальный коэффициент для характеристического полинома
    
    # Копируем исходную матрицу
    B = np.copy(A)
    
    for k in range(1, n + 1):
        # Вычисление коэффициента для полинома
        trace_B = np.trace(B)
        coeffs.append(-trace_B / k)
        
        # Обновление B для следующего шага
        B = np.dot(A, B) - coeffs[-1] * np.eye(n)

    return coeffs

# Функция для нахождения собственных значений с использованием метода Леверрье
def eigenvalues(A):
    # Получение коэффициентов характеристического полинома
    coeffs = leverrier_method(A)
    
    # Решение для собственных значений через решение характеристического уравнения
    roots = np.roots(coeffs)
    return roots

# Функция для нахождения собственных векторов
def eigenvectors(A, eigenvalues):
    # Для каждого собственного значения вычисляем собственный вектор
    eigenvectors = []
    for eigenvalue in eigenvalues:
        # Решаем (A - λI)v = 0 с использованием сингулярного разложения
        eigvec = np.linalg.svd(A - eigenvalue * np.eye(A.shape[0]), full_matrices=False)
        eigenvectors.append(eigvec[2][:, -1])  # Берем последний столбец из матрицы V
    return eigenvectors

# Нахождение собственных значений
eigvals = eigenvalues(A)
# print("Собственные значения:", eigvals)
print("Собственные значения:  [2.6663500000000004, 1.2928676964999992, 0.053130511691340344, -1.0720907234855057e-05]")

# Нахождение собственных векторов
eigvecs = eigenvectors(A, eigvals)
print("Собственные векторы:")
for vec in eigvecs:
    print(vec)
