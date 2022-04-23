# Langkah-langkah

* Buat akun okteto
* Download kube-config
* Update `~/.kube/config` atau `C:\\Users\<user-name>\.kube\config`
    - Download kubectl
    - Bikin akun okteto
    - Hubungkan kubectl ke cluster okteto dengan cara:
        - Menjalankan ini di terminal:
            - Windows: `$Env:KUBECONFIG=("c:\\User\<user-name>\Downloads\okteto-kube.config;$Env:KUBECONFIG;$HOME\.kube\config")`
            - Linux: `export KUBECONFIG=/mnt/c/Users/<user-name>/Downloads/okteto-kube.config:$KUBECONFIG:$HOME/.kube/config`
        - Supaya tidak perlu menjalankan perintah di atas setiap kali hendak manage okteto:
        Tambahkan script di atas ke `~/.bashrc` atau `~/.zshrc` (tergantung terminal yg dipakai)
        - Jika yang di atas gagal, salin isi okteto-kube.config ke `~/.kube/config`
* Pastikan `kubectl config get-contexts` mengarah ke okteto, jika tidak lakukan `kubectl config use-context <context-name>`
* Deploy mysql (`./mysql/deploy.sh`)
* Deploy go app (`./go-app/deploy.sh`)


# Catatan

* Cloud provider lain, seperti AWS dan GCP memiliki mekanisme sendiri untuk update kubeconfig. Biasanya harus menginstall CLI tool yang mereka sediakan
* Untuk MySQL sebaiknya menggunakan `statefulSet`, bukan `deployment`, namun dalam contoh ini kita menggunakan `deployment` saja, karena okteto free plan tidak mengijinkan akses `volume` API
* Supaya service bisa diakses dari luar cluster, kita perlu menggunakan `ingress`

