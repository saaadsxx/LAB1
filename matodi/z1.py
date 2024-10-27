import numpy as np

def jacobi(A, B, initial_guess=None, tolerance=1e-10, max_iterations=1000):
    n = len(B)
    if initial_guess is None:
        initial_guess = np.zeros(n)

    D = np.diag(np.diag(A))
    LU = A - D

    inv_D = np.linalg.inv(D)
    T = -np.dot(inv_D, LU)
    C = np.dot(inv_D, B)

    current_guess = initial_guess
    for i in range(max_iterations):
        new_guess = np.dot(T, current_guess) + C
        if np.linalg.norm(new_guess - current_guess) < tolerance:
            print(f"Метод Якоби сошелся за {i + 1} итераций.")
            return new_guess
        current_guess = new_guess

    raise ValueError("Jacobi method did not converge within the specified tolerance.")

# Создание матрицы A и вектора B
A = np.array([[26.4000, 0.1117, 0.1339, 0.1562],
              [0.0844, 25.5000, 0.1289, 0.1512],
              [0.0794, 0.1017, 24.6000, 0.1461],
              [0.0744, 0.0966, 0.1189, 23.7000]])
B = np.array([61.8139, 64.7307, 67.2806, 69.4636])

# Решение СЛАУ методом Якоби
solution = jacobi(A, B)

print("Решение СЛАУ методом Якоби:")
print(solution)
