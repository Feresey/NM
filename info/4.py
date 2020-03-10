import numpy as np

def getIdxMaxEl(M):
    i_max, j_max, m = 0, 0, 0
    for i in range(len(M)):
        for j in range(i+1, len(M)):
            if np.abs(M[i, j]) > m:
                m = np.abs(M[i, j])
                i_max, j_max = i, j 
    return i_max, j_max

def criterion(A):
    return np.sqrt(sum([A[i, j]**2 for i in range(len(A)) for j in range(len(A)) if i != j]))

def getOrt(A, i , j):
    
    if A[i, i] == A[j, j]:
        phi = np.pi / 4
    else:
        phi = .5 * np.arctan(2*A[i, j] / (A[i, i] - A[j, j]))
        
    U = np.eye(A.shape[0])
    U[i, j] = -np.sin(phi)
    U[j, i] = np.sin(phi)
    U[i, i] = np.cos(phi)
    U[j, j] = np.cos(phi)
    return U

def jacobiProcess(A, eps=0.01):
    num_it = 1
    U_final = np.eye(A.shape[0])
    print("Default criterion: ", criterion(A))
    while criterion(A) > eps:
        i, j = getIdxMaxEl(A)
        U = getOrt(A, i, j)
        U_final = np.dot(U_final, U)
        A = np.linalg.multi_dot([U.T, A, U])
        print(f"It: {num_it}, norm: {criterion(A)}")
        num_it += 1
 
    return A, U_final

if __name__ == "__main__":
    
    A = np.array([
        [8, -3, 9],
        [-3, 8, -2],
        [9, -2, -8],
    ])
    
    J, U = jacobiProcess(A)
    eigenvalues = [J[i, i] for i in range(len(J))]
    eigenvectores = [U[:, i] for i in range(len(U))]
    
    print("After iterations:\n", J)
    print("\nEigenvalues: ", eigenvalues)
    print("Values:\n", eigenvectores)
    
    print("\nOrotogonality:")
    for i in range(len(eigenvectores)):
        for j in range(i+1, len(eigenvectores)):
            print(f"(x{i}, x{j}) = ", np.dot(eigenvectores[i], eigenvectores[j]))
            
    print("\nAx = lambda*x")
    for e, v in zip(eigenvalues, eigenvectores):
        print(f"{np.dot(A, v)} = {(e * v)}")
