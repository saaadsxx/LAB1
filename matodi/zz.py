import numpy as np

def relaxation(A, B, omega=1.0, initial_guess=None, tolerance=1e-10, max_iterations=1000):
    n = len(B)
    if initial_guess is None:
        initial_guess = np.zeros(n)

    D = np.diag(np.diag(A))
    L = -np.tril(A, -1)
    U = -np.triu(A, 1)

    inv_DL = np.linalg.inv(D + omega * L)
    T = np.dot(inv_DL, ((1 - omega) * D - omega * U))
    C = omega * np.dot(inv_DL, B)

    current_guess = initial_guess
    for i in range(max_iterations):
        new_guess = np.dot(T, current_guess) + C
        if np.linalg.norm(new_guess - current_guess) < tolerance:
            print(f"Метод релаксации сошелся за {i + 1} итераций.")
            return new_guess
        current_guess = new_guess

    raise ValueError("Relaxation method did not converge within the specified tolerance.")

# Создание матрицы A и вектора B
A = np.array([[26.4000, 0.1117, 0.1339, 0.1562],
              [0.0844, 25.5000, 0.1289, 0.1512],
              [0.0794, 0.1017, 24.6000, 0.1461],
              [0.0744, 0.0966, 0.1189, 23.7000]])
B = np.array([61.8139, 64.7307, 67.2806, 69.4636])

# Решение СЛАУ методом релаксации
solution = relaxation(A, B)

print("Решение СЛАУ методом релаксации:")
print(solution)
