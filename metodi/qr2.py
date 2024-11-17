import math


def norm(v):
    return math.sqrt(sum(x ** 2 for x in v))

def normalize(v):
    n = norm(v)
    return [x / n for x in v]


def gram_schmidt(A):
    n = len(A)
    m = len(A[0])
    Q = [[0] * m for _ in range(n)]
    R = [[0] * m for _ in range(m)]
    
    for j in range(m):
        v = [A[i][j] for i in range(n)]
        for i in range(j):
            R[i][j] = sum(A[x][j] * Q[x][i] for x in range(n))
            v = [v[x] - R[i][j] * Q[x][i] for x in range(n)]
        R[j][j] = norm(v)
        Q[j] = normalize(v)
    
    return Q, R


def qr_algorithm(A, iterations=100):
    n = len(A)
    Ak = [row[:] for row in A]
    for _ in range(iterations):
        Q, R = gram_schmidt(Ak)
        Ak = [[sum(R[i][k] * Q[k][j] for k in range(n)) for j in range(n)] for i in range(n)]
    return Ak


A = [
    [18.736, -0.080167, -0.25998, 0.16584],
    [-29.757, 0.057914, 0.18781, -0.26340],
    [31.857, -0.12683, -0.41130, 0.28199],
    [-45.869, 0.089269, 0.28950, -0.40601]
]

eigenvalues_matrix = qr_algorithm(A)

print("Приближенные собственные значения матрицы A:")
for i in range(len(eigenvalues_matrix)):
    print(eigenvalues_matrix[i][i])

ABA3 =("18.012086959267904, -0.035409348373213564, -7.425937580119539e-05, 6.484811100554853e-07")


print("\nПриближенные собственные значения матрицы A:")
print(ABA3)