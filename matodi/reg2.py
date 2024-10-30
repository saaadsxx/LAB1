
A = [[1.03, 0.991],
     [0.991, 0.943]]

b = [2.57,2.49]

lambda_reg = 0.01


x = [0.0, 0.0]


def compute_gradient(A, b, x, lambda_reg):
    gradient = [0.0, 0.0]
    for i in range(len(A)):

        prediction = sum(A[i][j] * x[j] for j in range(len(x)))
        error = prediction - b[i]
        for j in range(len(x)):
            gradient[j] += error * A[i][j]
    

    for j in range(len(x)):
        gradient[j] += lambda_reg * x[j]
    
    return gradient

# Градиентный спуск
learning_rate = 0.01
num_iterations = 1000

for iteration in range(num_iterations):
    gradient = compute_gradient(A, b, x, lambda_reg)
    for j in range(len(x)):
        x[j] -= learning_rate * gradient[j]

print("Решение системы:", x)
