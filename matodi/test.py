import numpy as np
from numpy.linalg import inv, norm

def tikhonov_regularization(A, b, alpha):
    """
    Решение СЛАУ методом регуляризации Тихонова
    :param A: Матрица коэффициентов
    :param b: Вектор правой части
    :param alpha: Параметр регуляризации
    :return: Вектор решений x
    """
    # Определим регуляризованную матрицу
    I = np.eye(A.shape[1])
    A_tikhonov = A.T @ A + alpha * I
    b_tikhonov = A.T @ b

    # Решаем регуляризованную систему
    x = inv(A_tikhonov) @ b_tikhonov
    return x

# Входные данные
A = np.array([[1.03, 0.991], 
              [0.991, 0.943]])
b = np.array([2.57, 2.49])
alpha = 0.01  # Регуляризационный параметр 

x_solution = tikhonov_regularization(A, b, alpha)
print("Решение x:", x_solution)
print("Норма невязки:", norm(A @ x_solution - b))
