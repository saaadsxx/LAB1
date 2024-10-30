import numpy as np

def tikhonov_regularization(A, b, alpha=0.01):
    """
    Решение плохо обусловленной системы Ax = b с помощью метода регуляризации Тихонова.
    
    Параметры:
    A : ndarray, 2D массив, матрица системы
    b : ndarray, 1D массив, вектор правой части
    alpha : float, параметр регуляризации (чем больше, тем сильнее регуляризация)
    
    Возвращает:
    x : ndarray, решение системы
    """
    A = np.array(A, dtype=float)
    b = np.array(b, dtype=float)
    
    # Создаем регуляризованную матрицу (A^T * A + alpha * I)
    I = np.eye(A.shape[1])
    regularized_matrix = A.T @ A + alpha * I
    regularized_vector = A.T @ b
    
    # Решаем систему для регуляризованной матрицы
    x = np.linalg.solve(regularized_matrix, regularized_vector)
    return x

# Входные данные
A = [[1.03, 0.991], [0.991, 0.943]]
b = [2.57, 2.49]

# Решение системы с регуляризацией
alpha = 0.01  # параметр регуляризации, можно настроить
x = tikhonov_regularization(A, b, alpha)
print("regularization:", x)
