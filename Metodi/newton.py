import numpy as np
from scipy.optimize import fsolve

def equations(vars):
    x, y = vars
    f1 = np.sin(x - 1) + y - 1.5
    f2 = x - np.sin(y + 1) - 1
    return [f1, f2]

# Начальное приближение
initial_guess = [0, 0]

# Решение системы уравнений методом Ньютона
solution = fsolve(equations, initial_guess)

print("Решение СЛАУ:", solution)
