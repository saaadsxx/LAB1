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
            print(f"iteration: {i + 1} ")
            return new_guess
        current_guess = new_guess

    raise ValueError("Jacobi method did not converge within the specified tolerance.")

# Создание матрицы A и вектора B
A = np.array([[27.7000, 0.1171, 0.1403, 0.1636],
              [0.0888, 26.8000, 0.1353, 0.1586],
              [0.0838, 0.1070, 25.9000, 0.1536],
              [0.0788, 0.1020, 0.1253, 25.0000]])
B = np.array([67.6682, 70.7478, 73.4601, 75.8051])

# Решение СЛАУ методом Якоби
solution = jacobi(A, B)

print("Jacoby method:")
print(solution)
