import numpy as np

def qr_method(A):
    Q, R = np.linalg.qr(A)
    return np.linalg.solve(R, np.dot(Q.T, np.eye(A.shape[0])))

# Матрица из вашего примера
A = np.array( [
    [18.736, -0.080167, -0.25998, 0.16584],
    [-29.757, 0.057914, 0.18781, -0.26340],
    [31.857, -0.12683, -0.41130, 0.28199],
    [-45.869, 0.089269, 0.28950, -0.40601]
])

solution = qr_method(A)
print("Solution:")
print(solution)