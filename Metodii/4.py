import numpy as np
import matplotlib.pyplot as plt


x = np.array([1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0])
y = np.array([2.23, 2.29, 2.27, 2.62, 2.72, 2.82, 3.13, 3.49, 3.82, 3.95, 4.22, 4.48, 5.06, 5.50, 5.68, 6.19, 6.42, 7.04, 7.57, 8.10])


def polynomial_fit(x, y):
    n = len(x)
    
    A = np.vstack([np.ones(n), x, x ** 2]).T  
    b = y  
  
    coeffs = np.linalg.lstsq(A, b, rcond=None)[0]
    return coeffs


coeffs = polynomial_fit(x, y)
print("Coefficients of the second-degree approximating polynomial:")
print(f"a0 = {coeffs[0]:.4f}, a1 = {coeffs[1]:.4f}, a2 = {coeffs[2]:.4f}")


def polynomial(x_val, coeffs):
    return coeffs[0] + coeffs[1] * x_val + coeffs[2] * x_val ** 2


x_values = np.arange(1, 10.5, 0.5)
print("\nPolynomial values for x in the range from 1 to 10 with step 0.5:")
for x_val in x_values:
    y_val = polynomial(x_val, coeffs)
    print(f"x: {x_val:.1f}, y: {y_val:.2f}")

x_values_plot = np.linspace(1, 20, 100)  
y_values_plot = polynomial(x_values_plot, coeffs)

plt.figure(figsize=(8, 6))
plt.scatter(x, y, color="red", label="Original Data")
plt.plot(x_values_plot, y_values_plot, color="blue", label="Approximated Curve (2nd-degree polynomial)")

# Customize the plot
plt.title("Data Approximation Using Second-Degree Polynomial")
plt.xlabel("x")
plt.ylabel("y")
plt.legend()
plt.grid(True)


plt.show()
