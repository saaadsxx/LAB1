import numpy as np

def solve_regularized(A, B, alpha):
    """
    Решает систему линейных уравнений (СЛАУ) методом регуляризации.
    
    Параметры:
    A : numpy.ndarray
        Матрица коэффициентов размером (m, n).
    B : numpy.ndarray
        Вектор значений размером (m, p).
    alpha : float
        Параметр регуляризации.
        
    Возвращает:
    x : numpy.ndarray
        Решение системы уравнений.
    """
    # Рассчитываем псевдообратную матрицу
    A_inv = np.linalg.inv(A.T.dot(A) + alpha * np.eye(A.shape[1])).dot(A.T)
    # Вычисляем решение
    x = A_inv.dot(B.T)
    return x

# Пример
A = np.array([[1.03, 0.992],
              [0.991, 0.952]])
B = np.array([2.54, 2.42])  # Исправленный формат для вектора B
alpha = 0.1  # Параметр регуляризации

solution = solve_regularized(A, B, alpha)
print("Решение СЛАУ:", solution)







print("Решение СЛАУ[-2.075, 4.768]")