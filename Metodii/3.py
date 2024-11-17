import numpy as np
import matplotlib.pyplot as plt


x = np.array([1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0])
y = np.array([2.23, 2.29, 2.27, 2.62, 2.72, 2.82, 3.13, 3.49, 3.82, 3.95, 4.22, 4.48, 5.06, 5.50, 5.68, 6.19, 6.42, 7.04, 7.57, 8.10])


def divided_differences(x, y):
    n = len(y)
    coef = np.zeros((n, n))
    coef[:, 0] = y  
    
    
    for j in range(1, n):
        for i in range(n - j):
            coef[i, j] = (coef[i + 1, j - 1] - coef[i, j - 1]) / (x[i + j] - x[i])
    
    return coef[0]  

def newton_polynomial(x_val, x, coeffs):
    n = len(coeffs)
    result = coeffs[0]
    term = 1
    for i in range(1, n):
        term *= (x_val - x[i - 1])
        result += coeffs[i] * term
    return result


coeffs = divided_differences(x, y)


x_new = np.linspace(1, 10, 100)
y_new = [newton_polynomial(xi, x, coeffs) for xi in x_new]


plt.scatter(x, y, color="red", label="Nodes")
plt.plot(x_new, y_new, label="Newton Polynomial", color="blue")
plt.title("Interpolation Using Newton Polynomial")
plt.xlabel('x')
plt.ylabel('y')
plt.legend()
plt.grid()
plt.show()


x_values = np.arange(1, 10.5, 0.5)
for x_val in x_values:
    y_val = newton_polynomial(x_val, x, coeffs)
    print(f"x: {x_val:.1f}, y: {y_val:.2f}")
