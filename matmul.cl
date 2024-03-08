__kernel void matmul(const int M, const int N, const int K, 
                     const __global float* A, const __global float* B, __global float* C) {
    int row = get_global_id(0);
    int col = get_global_id(1);
    float sum = 0.0;
    for (int k = 0; k < K; ++k) {
        sum += A[row * K + k] * B[k * N + col];
    }
    C[row * N + col] = sum;
}