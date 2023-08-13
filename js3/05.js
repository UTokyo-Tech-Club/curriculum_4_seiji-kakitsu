const fibo = (N) => {
    let f = [0, 1]
    for (let i = 1; i < N; i++) {
        // f の最後（n番目）に f[n-1] + f[n-2] を追加する
        const f_i = f[i-1] + f[i];
        f.push(f_i);
    }
    return f  // 第0項から第N項までが順に格納された配列
}